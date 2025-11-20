package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	RABBITMQ_URL = "amqp://guest:guest@rabbitmq:5672/"
	QUEUE_NAME   = "weather_queue"
	API_URL      = "http://backend:3000/api/weather/logs"
)

func main() {
	var conn *amqp.Connection
	var err error
	
	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial(RABBITMQ_URL)
		if err == nil {
			break
		}
		log.Printf("RabbitMQ ainda não está pronto... esperando 5s. Erro: %v", err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("Falha ao conectar no RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Falha ao abrir canal: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		QUEUE_NAME, 
		true,       
		false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Falha ao declarar fila: %v", err)
	}


	msgs, err := ch.Consume(
		q.Name, 
		"",     
		false,  
		false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Falha ao registrar consumidor: %v", err)
	}

	log.Println(" [*] Worker Go rodando. Esperando mensagens...")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [>] Recebido: %s", d.Body)

			err := sendToAPI(d.Body)
			if err != nil {
				log.Printf(" [!] Erro ao enviar para API (API pode estar offline): %v", err)
				d.Ack(false) 
			} else {
				log.Println(" [V] Sucesso! Enviado para API.")
				d.Ack(false) 
			}
		}
	}()

	<-forever
}

func sendToAPI(jsonData []byte) error {
	resp, err := http.Post(API_URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}