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
var aaa int =1
type data struct {
	Operator string `json:"operator"`
	Operand1 int    `json:"operand1"`
	Operand2 int    `json:"operand2"`
	TX_ID    int    `json:"tx_id"`
}
//
var testdata = [][]string{
	{"ID", "TIME", " PROBLEM", "RESULT"},
}
var testdata2 = [][]string{
	{"ID", "RESULTFLOAT", "RESULTSEIENTIFIC"},
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

// type Result struct {
// 	resultfloat string //`json:"resultfloat"`
// 	resultseientific string //`json:"resultseientific"`
// 	tx_id string //`json:"tx_id"`
// }
type Result struct {
	Resultfloat float64 `json:"resultfloat"`
	Resultseientific string `json:"resultseientific"`
	Tx_id int `json:"tx_id"`
}

func post(e string,post_data []byte, wg *sync.WaitGroup, ch chan<- string)[]byte {
	defer wg.Done() //clear 1 stak
	request, _ := http.NewRequest("POST", "http://a9183ce3d200511e9a6250a2c719c0b1-1242495179.us-east-1.elb.amazonaws.com:3000/api/cal", bytes.NewBuffer(post_data))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	var a []byte 
	// data1, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(err)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		return data
		// fmt.Println(string(data))
		ch <- string(data)
		if string(data) == "Bad Request" {
			fmt.Println("OMG")
		}
		// fmt.Println("")
	} //return data1
	return a
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	var response []string

	file, err := os.Create("result.csv")
	checkError("connot creat file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	file2, err := os.Create("result2.csv")
	checkError("connot creat file", err)
	defer file2.Close()
	writer2 := csv.NewWriter(file2)
	defer writer2.Flush()

	for _, value := range testdata { //head of table
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
	for _, value2 := range testdata2 { //head of table
		err := writer2.Write(value2)
		checkError("Cannot write to file", err)
	}
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

		row := []string{e, t, text, ans}
		err := writer.Write(row)
		checkError("Cannot write to file", err)

		var data = data{
			Operator: o,
			Operand1: a,
			Operand2: b,
			TX_ID:    i,
		}
		// body, _ := json.Marshal(data)
		// fmt.Println(string(body))

		post_data, _ := json.Marshal(data)
		wg.Add(1) //add 1
		// go post(e,post_data, &wg, ch)
		data1 := post(e,post_data, &wg, ch)
		// fmt.Println(string(data1))
		////////////////////////////////////////////////
		res := Result{}
			err = json.Unmarshal(data1, &res)
			if err != nil {
				fmt.Println("There was an error:", err)
				
			}
				fmt.Printf(" = %v\t", res.Resultfloat)
				fmt.Printf(" %v\t", res.Resultseientific)
				fmt.Printf(" %v", res.Tx_id)
				fmt.Println()
		
		r:=res.Resultfloat
			rf := fmt.Sprintf("%.6f", r)
			id := strconv.Itoa(res.Tx_id)
			data2 := []string{id, rf,res.Resultseientific}
			err = writer2.Write(data2)
			checkError("Cannot write to file", err)
			if ans !=rf{
				fmt.Println("error")
			}else {fmt.Println("pass")}
			

	} //for
	go func() {
		wg.Wait() //wait wg.=0
		close(ch)
	}()
	for res := range ch {
		response = append(response, res)
	}
	//fmt.Print(response)
} 
//main

//----------------------------------------------------------------------------------------------------------------------------------
// operator string
// operand1 int
// operand2 int
// tx_id  int
