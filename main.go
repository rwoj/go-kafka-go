package main

import (
	"flag"
	"fmt"

	"github.com/rwoj/go-kafka-go/confl/cconsumer"
	"github.com/rwoj/go-kafka-go/confl/cproducer"
	// "github.com/rwoj/go-kafka-go/seg/consumer"
	// "github.com/rwoj/go-kafka-go/seg/producer"
)

func main() {

	libraryFlagValue := flag.String("l", "", "    Use this flag with either \"s  \" = segmentio or \"c \" = confluent")
	optionFlagValue := flag.String("o", "", "    Use this flag with either \"consumer\" or \"producer\"")
	flag.Parse()

	if *libraryFlagValue == "c" {
		if *optionFlagValue == "consumer" {
			cconsumer.StartConsumer()
		} else if *optionFlagValue == "producer" {
			cproducer.StartProducer()
		} else {
			fmt.Print("Usage: \n -l    and either \"s\" or \"producer\"\n -o     and either \"consumer\" or \"producer\"\n")
		}
	} else if *libraryFlagValue == "c" {
		fmt.Print("hello")
	} else {
		fmt.Print("Usage: \n -o     and either \"consumer\" or \"producer\"\n")
	}

}
