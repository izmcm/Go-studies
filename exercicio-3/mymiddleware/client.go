package main

import (
	"fmt"
	// "invoker"
	// "log"
	"clientproxy"
	"encoding/csv"
	"log"
	"math/rand"
	"naming"
	"os"
	"proxies"
	"strconv"
	"time"
)

const REPETICOES = 15000

var times [REPETICOES][]string

func main() {
	simpleTest()
}

func simpleTest() {
	namingServer := naming.NamingService{Repository: make(map[string]clientproxy.ClientProxy)}
	namingServer.Register("Calculator", proxies.NewCalculatorProxy())

	calculator := proxies.CalculatorProxy(namingServer.Lookup("Calculator"))
	REPETICOES := 3

	for i := 0; i < REPETICOES; i++ {
		start := time.Now()

		op := rand.Intn(4)
		var value int
		if op == 0 {
			value = calculator.Add(1, 2)
		} else if op == 1 {
			value = calculator.Mul(1, 2)
		} else if op == 2 {
			value = calculator.Sub(4, 2)
		} else if op == 3 {
			value = calculator.Div(4, 2)
		}

		now := time.Now()
		timeAll, _ := time.ParseDuration(now.Sub(start).String())
		times[i] = []string{strconv.FormatInt(timeAll.Nanoseconds(), 10)}
		fmt.Println(value)
	}
}

func makeCommonTest() {
	namingServer := naming.NamingService{Repository: make(map[string]clientproxy.ClientProxy)}
	namingServer.Register("Calculator", proxies.NewCalculatorProxy())

	calculator := proxies.CalculatorProxy(namingServer.Lookup("Calculator"))

	for i := 0; i < REPETICOES; i++ {
		start := time.Now()

		op := rand.Intn(4)
		// var value int
		if op == 0 {
			calculator.Add(1, 2)
		} else if op == 1 {
			calculator.Mul(1, 2)
		} else if op == 2 {
			calculator.Sub(4, 2)
		} else if op == 3 {
			calculator.Div(4, 2)
		}

		now := time.Now()
		timeAll, _ := time.ParseDuration(now.Sub(start).String())
		times[i] = []string{strconv.FormatInt(timeAll.Nanoseconds(), 10)}
		// fmt.Println(value)
	}

	creatFileAndWrite("warmup")
}

func makeBurstTest() {
	namingServer := naming.NamingService{Repository: make(map[string]clientproxy.ClientProxy)}
	namingServer.Register("Calculator", proxies.NewCalculatorProxy())

	calculator := proxies.CalculatorProxy(namingServer.Lookup("Calculator"))

	for i := 0; i < REPETICOES; i++ {
		if i%1000 == 0 {
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}

		start := time.Now()

		op := rand.Intn(4)
		// var value int
		if op == 0 {
			calculator.Add(1, 2)
		} else if op == 1 {
			calculator.Mul(1, 2)
		} else if op == 2 {
			calculator.Sub(4, 2)
		} else if op == 3 {
			calculator.Div(4, 2)
		}

		now := time.Now()
		timeAll, _ := time.ParseDuration(now.Sub(start).String())
		times[i] = []string{strconv.FormatInt(timeAll.Nanoseconds(), 10)}
	}

	creatFileAndWrite("burst")
}

func makeSparseTest() {
	namingServer := naming.NamingService{Repository: make(map[string]clientproxy.ClientProxy)}
	namingServer.Register("Calculator", proxies.NewCalculatorProxy())

	calculator := proxies.CalculatorProxy(namingServer.Lookup("Calculator"))

	for i := 0; i < REPETICOES; i++ {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		start := time.Now()

		op := rand.Intn(4)
		// var value int
		if op == 0 {
			calculator.Add(1, 2)
		} else if op == 1 {
			calculator.Mul(1, 2)
		} else if op == 2 {
			calculator.Sub(4, 2)
		} else if op == 3 {
			calculator.Div(4, 2)
		}

		now := time.Now()
		timeAll, _ := time.ParseDuration(now.Sub(start).String())
		times[i] = []string{strconv.FormatInt(timeAll.Nanoseconds(), 10)}
	}

	creatFileAndWrite("esparco")
}

func creatFileAndWrite(testType string) {
	id := rand.Intn(99999)
	file, err := os.Create(fmt.Sprintf("benchmark/"+testType+"/result%d.csv", id))
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range times {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
