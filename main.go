package main

import (
	"flag"
	"fmt"
	"kafka-go-benchmark/segmentio"
	"runtime"
	"time"
)

var (
	strategy string
	brokers  string
	topic    string
	driver   string
	records  int
)

func main() {
	flag.StringVar(&brokers, "brokers", "localhost:9092", "0")
	flag.StringVar(&strategy, "strategy", "consume", "produce, consume")
	flag.IntVar(&records, "records", 250000, "number of records to read from kafka")
	flag.StringVar(&driver, "driver", "segmentio", "segmentio,sarama")
	flag.Parse()

	fmt.Println("Strategy: ", strategy, "Records:", records, "Driver:", driver)
	if strategy == "produce" {
		testProduce()
	} else if strategy == "consume" {
		testConsume()
	}
}

func testProduce() {
	topic = fmt.Sprintf("%d", time.Now().Unix())
	segmentio.CreateTopic(brokers, topic)

	produce()
	fmt.Println("topic", topic)
}

func testConsume() {
	topic = fmt.Sprintf("%d", time.Now().Unix())
	consume()
}

func PrintMemUsage(title string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("[%s] TotalAlloc = %v MiB\n", title, m.TotalAlloc/1024/1024)
}
