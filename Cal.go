package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type data struct {
	Operator string `json:"operator"`
	Operand1 int    `json:"operand1"`
	Operand2 int    `json:"operand2"`
	TX_ID    int    `json:"tx_id"`

	//Time     string `json:"time"`
}

func timestamp() string {
	t := time.Now().Format("03:04:05.000")
	return t
}
func calculator(Operator, operand1, operand2 string) string {
	var r1, s1 = 0.0, " "
	f, err := strconv.ParseFloat(operand1, 32)
	if err == nil {
		fmt.Println("operand1 =", f)
	}
	f2, err := strconv.ParseFloat(operand2, 32)
	if err == nil {
		fmt.Println("operand2 =", f2)
	}
	x := float64(f) //int to float
	y := float64(f2)
	if Operator == "+" {
		r1 = x + y
		s1 = fmt.Sprintf("%f", r1)
		return s1

	}
	if Operator == "-" {
		r1 = x - y
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}
	if Operator == "*" {
		r1 = x * y
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}
	if Operator == "/" {
		r1 = x / y
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}

	if Operator == "DIV" {
		r1 = math.Floor(x / y)
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}
	if Operator == "MOD" {
		r1 := math.Mod(x, y)
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}
	return s1
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomoperator() string {
	rand.Seed(time.Now().Unix())
	operator := []string{
		"+", "-", "*", "/", "DIV", "MOD",
	}
	n := rand.Int() % len(operator)
	// fmt.Println("Operate is ", operator[n])
	return operator[n]
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func post(post_data []byte, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done() //clear 1 stak
	request, _ := http.NewRequest("POST", "http://35.202.123.123:3000/api/cal", bytes.NewBuffer(post_data))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	//fmt.Println(err)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		ch <- string(data)
		if string(data) == "Bad Request" {
			fmt.Println("OMG")
		}
		// fmt.Println("")
	} //else

}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	var response []string
	for i := 1; i < 1000; i++ {
		if i > 50 && i < 52 || i > 500 && i < 502 {
			time.Sleep(5000 * time.Millisecond)
		}
		rand.Seed(time.Now().UnixNano())

		a := randomInt(-100, 100) //get an int in the 1...n range
		o := randomoperator()
		b := randomInt(-100, 100) //get an int in the 1...n range
		// t := timestamp()

		var data = data{
			Operator: o,
			Operand1: a,
			Operand2: b,
			TX_ID:    i,
		}
		body, _ := json.Marshal(data)
		fmt.Println(string(body))

		post_data, _ := json.Marshal(data)
		wg.Add(1) //add 1
		go post(post_data, &wg, ch)

	} //for
	go func() {
		wg.Wait() //wait wg.=0
		close(ch)
	}()
	for res := range ch {
		response = append(response, res)
	}
	//fmt.Print(response)
} //main

//----------------------------------------------------------------------------------------------------------------------------------
// operator string
// operand1 int
// operand2 int
// tx_id  int
