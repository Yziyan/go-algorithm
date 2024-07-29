// @Author: Ciusyan 2024/7/28

package main

import (
	"bufio"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
)

func main() {
	// Kafka 配置
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	// 创建生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, cfg)
	if err != nil {
		log.Fatalf("Failed to start producer: %s", err)
	}
	defer producer.Close()

	// 从控制台读取输入并发送消息
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter message to send to Kafka (type 'exit' to quit):")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1] // 移除换行符
		if text == "exit" {
			break
		}

		msg := &sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.StringEncoder(text),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to send message: %s", err)
		} else {
			fmt.Printf("Message is stored in partition %d, offset %d\n", partition, offset)
		}
	}
}
