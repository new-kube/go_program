package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/spf13/cast"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 等待所有follower都回复ack，确保Kafka不会丢消息
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner // 对Key进行Hash，同样的Key每次都落到一个分区，这样消息是有序的

	addrs := []string{
		"alikafka-pre-cn-tl32zhevr007-1-vpc.alikafka.aliyuncs.com:9092",
		"alikafka-pre-cn-tl32zhevr007-2-vpc.alikafka.aliyuncs.com:9092",
		"alikafka-pre-cn-tl32zhevr007-3-vpc.alikafka.aliyuncs.com:9092",
	}
	topic := "xes_dm_supplier_adapter_order_event"
	//topic = "xes_dm_supplier_adapter_after_sale_event"

	// 使用同步producer，异步模式下有更高的性能，但是处理更复杂，这里建议先从简单的入手
	producer, err := sarama.NewSyncProducer(addrs, config)
	defer func() {
		_ = producer.Close()
	}()
	if err != nil {
		panic(err.Error())
	}

	type Message[T any] struct {
		Tag   string `json:"tag"`
		MsgID string `json:"msg_id"`
		Data  T      `json:"data"`
	}

	type Forward struct {
		SupplierID  int      `json:"supplier_id"`
		OrderID     string   `json:"p_id"`         // 父订单id
		SubOrderIDs []string `json:"s_ids"`        // 子订单id列表
		OrderStatus int      `json:"order_status"` // 父订单状态
	}

	type Backward struct {
		SupplierID      int      `json:"supplier_id"`
		OrderID         string   `json:"p_id"`             // 父订单id
		SubOrderID      []string `json:"s_ids"`            // 关联子订单id
		AfterSaleID     string   `json:"aftersale_id"`     // 售后单id
		AfterSaleStatus int      `json:"aftersale_status"` // 售后单状态
		AfterSaleType   int      `json:"aftersale_type"`   // 售后单类型
	}

	// orderID := "12230207113235541643097"
	// key := cast.ToString(cast.ToInt(orderID) % 6)

	// order := Forward{
	// 	SupplierID:  2,
	// 	OrderID:     orderID,
	// 	SubOrderIDs: []string{"22230207113235376373097"},
	// 	OrderStatus: 20,
	// }

	// orderPaied := Forward{
	// 	SupplierID:  1,
	// 	OrderID:     orderID,
	// 	SubOrderIDs: []string{subOrderID},
	// 	OrderStatus: 11, // 已支付
	// }

	// {"tag":"206","msg_id":"dayu_c3b9483d798ef8c5df3ff96ac00d0421","data":
	// {
	// 	   "supplier_id":2,
	//     "p_id":"12230208183342126093097",
	//     "s_ids":["22230208183342456783097"],
	//     "aftersale_id":"32230208183701827213097",
	//     "aftersale_status":71,
	//     "aftersale_type":2
	// }}
	topic = "dm_adapter_boku_sales"

	afterSaleID := "32230208183701827213097"

	key := cast.ToString(cast.ToInt(afterSaleID) % 6)
	sales := Backward{
		SupplierID:      2,
		OrderID:         "12230208183342126093097",
		SubOrderID:      []string{"22230208183342456783097"},
		AfterSaleID:     afterSaleID,
		AfterSaleStatus: 71,
		AfterSaleType:   2,
	}

	message := Message[Backward]{
		Tag:   "206",
		MsgID: "hello-xxxxxx",
		Data:  sales,
	}

	data, _ := json.Marshal(message)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(data),
	}

	t1 := time.Now().Nanosecond()
	partition, offset, err := producer.SendMessage(msg)
	t2 := time.Now().Nanosecond()

	time.Sleep(time.Millisecond * 10)

	if err == nil {
		fmt.Println("produce success, partition:", partition, ",offset:", offset, ",cost:", (t2-t1)/(1000*1000), " ms")
	} else {
		fmt.Println(err.Error())
	}

}
