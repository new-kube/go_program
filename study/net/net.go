package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	ifaceEth0 = "eth0"
	ifaceEn0  = "en0"
)

// 0-hostname, 1-host_no, 2-ipv4, 3-env
type MachineID [4]string

// 不同环境的机器数量不同，这里简单的使用hostname来区分不同的机器。
// 开发、测试、灰度环境通常是一台机器
// 生产环境通常是多台机器
// 不同环境可以使用相同的id。

func GetMachineID(env string) (MachineID, error) {
	mid := MachineID{}

	hostname, err := os.Hostname()
	if err != nil {
		return mid, fmt.Errorf("get hostname failed, err=%v", err)
	}

	mno := 0
	switch env {
	case "dev":
	case "test":
	case "gray":
	case "prod":
		mno, err = getMachineNo(hostname)
	}
	if err != nil {
		return mid, fmt.Errorf("parse hostname=%s no failed, err=%v", hostname, err)
	}

	ifaces, err := getIFaces()
	if err != nil {
		return mid, fmt.Errorf("get ifaces failed, err=%v", err)
	}

	ip := ""
	if x, ok := ifaces[ifaceEth0]; ok {
		ip = x
	}
	if ip == "" {
		if x, ok := ifaces[ifaceEn0]; ok {
			ip = x
		}
	}
	if ip == "" {
		return mid, fmt.Errorf("not found iface[%s, %s] exist", ifaceEth0, ifaceEn0)
	}

	// 机器号
	mid[0] = hostname
	mid[1] = strconv.Itoa(mno)
	mid[2] = ip
	mid[3] = env

	return mid, nil
}

func getIFaces() (map[string]string, error) {
	ret := make(map[string]string)
	ifaces, err := net.Interfaces()
	if err != nil {
		return ret, fmt.Errorf("get interfaces failed! err: %v", err)
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return ret, fmt.Errorf("get addresses failed! err: %v", err)
		}

		ips := make([]string, 0)
		for _, addr := range addrs {
			if x, ok := addr.(*net.IPNet); ok && !x.IP.IsLoopback() {
				if x.IP.To4() != nil {
					ips = append(ips, x.IP.String())
				}
			}
		}
		if len(ips) == 0 {
			continue
		}

		// 一般不会出现一个网卡绑定多个ip地址，除非特殊设置
		ret[i.Name] = ips[0]
	}
	return ret, nil
}

func getMachineNo(hostname string) (int, error) {
	x := strings.Split(hostname, "-")
	if len(x) == 1 {
		return 0, fmt.Errorf("hostname=%s not like xxx-xxx-xxx format", hostname)
	}

	no, err := strconv.Atoi(x[len(x)-1])
	if err != nil {
		return 0, fmt.Errorf("hostname=%s not end with -xxx parse failed, err=%v", hostname, err)
	}

	return no, nil
}

func main() {
	env := "test"
	mid, err := GetMachineID(env)
	if err != nil {
		fmt.Printf("get machine id failed, err=%v\n", err)
		return
	}

	fmt.Printf("machine id=%+v\n", mid)
}
