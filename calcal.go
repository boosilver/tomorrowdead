package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type getData struct {
	Resultfloat      float64 `json:"resultfloat"`
	Resultseientific string  `json:"resultseientific"`
	Tx_id            int     `json:"tx_id"`
}

type data struct {
	Operator string `json:"operator"`
	Operand1 int    `json:"operand1"`
	Operand2 int    `json:"operand2"`
	TX_ID    int    `json:"tx_id"`
}

var mydata = [][]string{
	{"ID", "TIME", " PROBLEM", "RESULT", "RESULTSCI"},
}
var putdata = [][]string{
	{"Resultseientific", "Resultseientific", "TX_ID"},
}

func timestamp() string {
	t := time.Now().Format("03:04:05.000")
	return t
}

func calculator(Operator, operand1, operand2 string) string {
	var r1, s1 = 0.0, " "
	f, _ := strconv.ParseFloat(operand1, 32)

	f2, _ := strconv.ParseFloat(operand2, 32)

	if Operator == "+" {
		r1 = f + f2
		s1 = fmt.Sprintf("%f", r1)
		return s1

	}
	if Operator == "-" {
		r1 = f - f2
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}
	if Operator == "*" {
		r1 = f * f2
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}
	if Operator == "/" {
		r1 = f / f2
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}

	if Operator == "DIV" {
		r1 = math.Floor(f / f2)
		s1 = fmt.Sprintf("%f", r1)
		return s1
	}
	if Operator == "MOD" {
		r1 := math.Mod(f, f2)
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
	request, _ := http.NewRequest("POST", "http://a9183ce3d200511e9a6250a2c719c0b1-1242495179.us-east-1.elb.amazonaws.com:3000/api/cal", bytes.NewBuffer(post_data))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	//fmt.Println(err)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		//
		fmt.Println(string(data)) //ข้อมูลที่รับกลับมาทั้งหมด
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

	file, err := os.Create("result.csv")
	checkError("connot creat file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush() //writer1

	for _, value := range mydata { //head of table
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

	file2, err := os.Create("Get result.csv")
	checkError("connot creat file", err)
	defer file2.Close()
	writer2 := csv.NewWriter(file2)
	defer writer2.Flush() //writer2

	for _, value := range putdata { //head of table 2
		err := writer2.Write(value)
		checkError("Cannot write to file", err)
	}

	now := time.Now()
	after := now.Add(5 * time.Second) // fmt.Println("\nAdd 1 Minute:", after)
	for {
		time.Sleep(0 * time.Second)

		for i := 1; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(0 * time.Millisecond)
			a := randomInt(-100, 100) //get an int in the 1...n range
			o := randomoperator()
			b := randomInt(-100, 100) //get an int in the 1...n range
			e := strconv.Itoa(i)
			t := timestamp()
			astring := strconv.Itoa(a)
			bstring := strconv.Itoa(b)
			text := astring + o + bstring
			ans := calculator(o, astring, bstring)
			ans2, _ := strconv.ParseFloat(ans, 64)
			anssci := fmt.Sprintf("%.4e\n", ans2)
			fmt.Println("ans value ", ans)
			fmt.Println("v value", anssci)
			row := []string{e, t, text, ans, anssci}
			err := writer.Write(row)
			checkError("Cannot write to file", err)

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

		} //for i
		now = time.Now()
		if now.After(after) {
			break
		}
	}
	go func() {
		wg.Wait() //wait wg.=0
		close(ch)
	}()
	for res := range ch {
		body_string := getData{}
		json.Unmarshal([]byte(res), &body_string)
		//fmt.Println(res)
		aaa := body_string.Resultfloat
		bbb := body_string.Resultseientific
		ccc := body_string.Tx_id
		a2 := fmt.Sprintf("%f", aaa)
		c2 := strconv.Itoa(ccc)
		row2 := []string{a2, bbb, c2}
		err := writer2.Write(row2)
		checkError("Cannot write to file", err)

	}

	//fmt.Print(response)
} //main

//----------------------------------------------------------------------------------------------------------------------------------
// operator string
// operand1 int
// operand2 int
// tx_id  int
