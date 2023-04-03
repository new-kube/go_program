package main

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
)

type msgConsumerGroup struct {
	name string
}

func (msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		// todo::

		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V2_2_0_0 // specify appropriate version
	//consumerConfig.Version = sarama.V0_11_0_2 // specify appropriate version
	consumerConfig.Consumer.Return.Errors = true
	//consumerConfig.Consumer.Offsets.AutoCommit.Enable = true      // 禁用自动提交，改为手动
	//consumerConfig.Consumer.Offsets.AutoCommit.Interval = time.Second * 1 // 测试3秒自动提交
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	addrs := []string{
		//"10.90.73.100:9092",
		"alikafka-pre-cn-7mz2si1lk005-1-vpc.alikafka.aliyuncs.com:9092",
		"alikafka-pre-cn-7mz2si1lk005-2-vpc.alikafka.aliyuncs.com:9092",
		"alikafka-pre-cn-7mz2si1lk005-3-vpc.alikafka.aliyuncs.com:9092",
	}
	group := "dm_supplier_group1"
	topics := []string{
		"xes_dm_supplier_adapter_test",
	}

	cGroup, err := sarama.NewConsumerGroup(addrs, group, consumerConfig)
	if err != nil {
		panic(err)
	}

	consumerGroup := &msgConsumerGroup{}

	go func() {
		for err := range cGroup.Errors() {
			fmt.Printf("kafka running error: %s\n", err.Error())
		}
	}()

	for {
		err := cGroup.Consume(context.Background(), topics, consumerGroup)
		if err != nil {
			fmt.Printf("error: %v\n", err.Error())
			break
		}
		fmt.Printf("xxxxx\n")
	}

	_ = cGroup.Close()
}
