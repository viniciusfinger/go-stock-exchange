package main

import (
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/viniciusfinger/bolsa-valores/internal/infra/kafka"
	"github.com/viniciusfinger/bolsa-valores/internal/market/dto"
	"github.com/viniciusfinger/bolsa-valores/internal/market/entity"
	"github.com/viniciusfinger/bolsa-valores/internal/market/transformer"
)

func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)

	kafkaMsgChannel := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}

	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewConsumer(configMap, []string{"input"})

	//cria nova thread para rodar o consumer sem travar a execucao do restante do programa
	go kafka.Consume(kafkaMsgChannel)

	//recebe do kafka, joga no input channel, processa, joga pro output channel e publica no kafka
	book := entity.NewBook(ordersIn, ordersOut, wg)

	//cria nova thread para processar as ordens
	go book.Trade()

	//funcao anonima pra receber as mensagens do kafka
	go func() {
		for msg := range kafkaMsgChannel {
			wg.Add(1)
			fmt.Println(string(msg.Value))
			//dessereializar a mensagem de json -> dto
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)

			if err != nil {
				panic(err)
			}

			order := transformer.TransformInput(tradeInput)

			//joga as orders para o canal de entrada do book
			ordersIn <- order
		}
	}()

	for res := range ordersOut {
		orderOutput := transformer.TransformOutput(res)
		outputJson, err := json.Marshal(orderOutput)

		if err != nil {
			fmt.Println(err)
		}

		//public as orders processadas de volta no kafka
		producer.Publish(outputJson, []byte("orders"), "output")

	}
}
