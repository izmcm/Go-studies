package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	// gopher_and_rabbit "github.com/masnun/gopher-and-rabbit"
	"github.com/streadway/amqp"
)

var amqpURL string = "amqp://guest:guest@localhost:5672/"
var wg sync.WaitGroup

// 0 - add
// 1 - sub
// 2 - mul
// 3 - div
type AddTask struct {
	Number1   int
	Number2   int
	Operation int
	Opid      int
}

type Response struct {
	Number int
	Opid   int
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func compute(a int, b int, op int) int {
	if op == 1 {
		return a - b
	}
	if op == 2 {
		return a * b
	}
	if op == 3 {
		if b == 0 {
			return 0x3f3f3f3f
		}
		return a / b
	}
	return a + b
}

func createQueue(name string, channel *amqp.Channel) amqp.Queue {
	queue, err := channel.QueueDeclare(name, true, false, false, false, nil)
	handleError(err, fmt.Sprintf("Could not declare %s queue", name))
	return queue
}

func consumeNumbers(receiveChannel <-chan amqp.Delivery, responseChannel *amqp.Channel) {
	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range receiveChannel {
			log.Printf("Received a message: %s", d.Body)

			addTask := &AddTask{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			n1 := addTask.Number1
			n2 := addTask.Number2
			op := addTask.Operation
			id := addTask.Opid
			res := compute(n1, n2, op)
			data := Response{Number: res, Opid: id}

			var s string
			if op == 0 {
				s = "+"
			}
			if op == 1 {
				s = "-"
			}
			if op == 2 {
				s = "*"
			}
			if op == 3 {
				s = "/"
			}
			// return the data
			go outputNumbers(data, responseChannel)
			log.Printf("operation %d: %d %s %d = %d", id, n1, s, n2, res)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()

	// Stop for program termination
	<-stopChan
	wg.Done()
}

func outputNumbers(data Response, responseChannel *amqp.Channel) {
	body, err := json.Marshal(data)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	// TODO: Fazer o uso de json no outro middleware pra poder evitar problemas
	// na comparação.

	// Zona de publish
	err = responseChannel.Publish("", "ans", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}
}

// TODO: Goroutine that publish and dieeees

func main() {
	// connect
	conn, err := amqp.Dial(amqpURL)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	// create rabbitmq channel
	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")
	defer amqpChannel.Close()

	calculatorQueue := createQueue("calculator", amqpChannel)
	// ansQueue := createQueue("ans", amqpChannel)
	createQueue("ans", amqpChannel)

	// With a prefetch count greater than zero, the server will deliver that many messages to consumers before acknowledgments are received.
	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		calculatorQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	wg.Add(1)
	go consumeNumbers(messageChannel, amqpChannel)
	wg.Wait()
}
