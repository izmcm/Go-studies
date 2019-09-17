package main

import (
	"encoding/csv"
	"encoding/json"
	// "github.com/masnun/gopher-and-rabbit"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var amqpURL string = "amqp://guest:guest@localhost:5672/"
var wg sync.WaitGroup

var beforeTime map[int]time.Time = make(map[int]time.Time)
var afterTime map[int]time.Time = make(map[int]time.Time)

// var rct int = 0

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

func createQueue(name string, channel *amqp.Channel) amqp.Queue {
	queue, err := channel.QueueDeclare(name, true, false, false, false, nil)
	handleError(err, fmt.Sprintf("Could not declare %s queue", name))
	return queue
}

func receiveData(receiveChannel <-chan amqp.Delivery, requests int) {
	stopChan := make(chan bool)
	rct := 0

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range receiveChannel {
			log.Printf("Received a message: %s", d.Body)

			res := &Response{}
			err := json.Unmarshal(d.Body, res)

			afterTime[res.Opid] = time.Now()

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			// return the data
			log.Printf("Result of operation %d -> %d", res.Opid, res.Number)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

			rct += 1
			if rct == requests {
				close(stopChan)
				// wg.Done()
			}
		}
	}()

	// Stop for program termination
	<-stopChan

	wg.Done()
}

func postData(channel *amqp.Channel, queue amqp.Queue) {
	// Random pra gerar os numeros pra colocar na fila
	rand.Seed(time.Now().UnixNano())

	ID := rand.Intn(9999999)
	addTask := AddTask{Number1: rand.Intn(999), Number2: rand.Intn(999), Operation: rand.Intn(4), Opid: ID}
	body, err := json.Marshal(addTask)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	// TODO: Fazer o uso de json no outro middleware pra poder evitar problemas
	// na comparação.

	// Zona de publish
	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	beforeTime[ID] = time.Now()

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	// Agora precisa ver a resposta - AKA mete o code de consumer aqui

	log.Printf("AddTask: %d+%d", addTask.Number1, addTask.Number2)
}

func creatFileAndWrite(times []time.Duration) {
	file, _ := os.Create("result.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var arr []string

	for _, v := range times {
		arr = append(arr, strconv.FormatInt(v.Nanoseconds(), 10))
	}

	// for _, value := range arr {
	writer.Write(arr)
	// }
}

func main() {
	conn, err := amqp.Dial(amqpURL)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	calculatorQueue := createQueue("calculator", amqpChannel)
	ansQueue := createQueue("ans", amqpChannel)
	createQueue("ans", amqpChannel)

	// With a prefetch count greater than zero, the server will deliver that many messages to consumers before acknowledgments are received.
	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	ansChannel, err := amqpChannel.Consume(
		ansQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")
	n := 2000

	for i := 0; i < n; i++ {
		go postData(amqpChannel, calculatorQueue)
	}
	wg.Add(1)
	go receiveData(ansChannel, n)
	wg.Wait()

	times := []time.Duration{}
	for key, value := range afterTime {
		log.Println("\n")
		log.Println(key)
		// log.Println(value)
		// log.Println(beforeTime[key])

		diff, _ := time.ParseDuration(value.Sub(beforeTime[key]).String())
		times = append(times, diff)
	}

	creatFileAndWrite(times)
}
