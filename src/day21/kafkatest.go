package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func writeByConn() {
	topic := "my-topic"
	partition := 0

	// 连接至kafka集群的leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader", err)
	}

	// 设置发送超时
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	// 发送消息
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one")},
		kafka.Message{Value: []byte("two")},
		kafka.Message{Value: []byte("three")},
	)
	if err != nil {
		log.Fatal("fail to write message: ", err)
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer: ", err)
	}
}

// 连接到kafka后接收消息
func readByConn() {
	// 指定topic，partition
	topic := "my-topic"
	partition := 0
	// 连接至kafka的leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 设置读取超时
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	// 批量读取,最小10kB,最大1MB
	batch := conn.ReadBatch(10e3, 1e6)
	// 遍历读取消息
	b := make([]byte, 10e3) // buffer大小为10kB,这个buffer大小需要合理设置
	for {
		n, err := batch.Read(b) // 把batch中的数据读取到b中
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))

		// msg, err := batch.ReadMessage()	// 读取效率低一些
		// if err != nil {
		// 	break
		// }
		// fmt.Println(string(msg.Value))
	}

	// 关闭batch
	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}
	// 关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close conntetion:", err)
	}
}

func topicList() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}
	// 遍历所有分区取topic
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}

}
