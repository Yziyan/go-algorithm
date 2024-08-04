// @Author: Ciusyan 2024/7/28

package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func main() {
	// 创建消费者配置
	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true

	// 创建消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, cfg)
	if err != nil {
		log.Fatalf("Failed to start consumer: %s", err)
	}
	defer consumer.Close()

	// 获取分区主题
	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Failed to start partision consumer: %s", err)
	}
	defer partitionConsumer.Close()

	// 消费消息
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Consumed message: %s product time: %s\n", string(msg.Value), msg.Timestamp)
	}
}
