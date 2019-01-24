package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
    "time"
    "strconv"
)

type data struct {
	Operator string `json:"operator"`
	Operand1 int    `json:"operand1"`
	Operand2 int    `json:"operand2"`
	TX_ID    int    `json:"tx_id"`

	//Time     string `json:"time"`
}

// func timestamp() string {
// 	t := time.Now().Format("03:04:05.000")
// 	return t
//}
// func calculator(Operator, operand1,operand2 string int int )(string ) {
// 	var r1,s1 = 0.0," "

// 	x := float32(operand1) //int to float
// 	y := float32(operand2)
// 	if Operator == "+" {
// 		r1 = x+y
// 		s1 = fmt.Sprintf("%f", r1)
// 		return s1

// 	}
// 	else if Operator == "-"{
// 		r1 = x-y
// 		s1 = fmt.Sprintf("%f", r1)
// 		return s1
// 	}
// 	else if Operator == "*"{
// 		r1 = x*y
// 		s1 = fmt.Sprintf("%f", r1)
// 		return s1
// 	}
// 	else if Operator == "/"{
// 		r1 = x/y
// 		s1 = fmt.Sprintf("%f", r1)
// 		return s1
// 	}

// 	else if Operator == "DIV"{
// 		r1 = math.Floor(x/y)
// 		s1 = fmt.Sprintf("%f", r1)
// 	}
// 	else if Operator == "MOD"{
// 		r1 := math.Mod(x,y)
// 		s1 = fmt.Sprintf("%f", r1)

// 	}

// 	}

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
	for i := 1; i < 5; i++ {

		rand.Seed(time.Now().UnixNano())
		time.Sleep(0 * time.Millisecond)
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


        aa := strconv.Itoa(a)
        oo := string(o)
        bb := strconv.Itoa(b)
        fmt.Println(aa)
        fmt.Println(bb)
        fmt.Println(oo)
        ii := string(i)
        abc2:=aa+oo+bb
        abc:=[]string{abc2,aa,oo,bb}
        iii:=[]string{ii}
        // ooo:=[]string{oo}
        // bbb:=[]string{bb}
        datas := [][]string{iii,abc,abc,abc}
        save(datas)
        // fmt.Println(datas)
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
// package main

// import (
//     "bytes"
//     // "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "net/http"
//     "time"
//     // "encoding/json"
// 	// "math/rand"
	
// )

// var a =5
// func main() {
//     t := time.Now()
    
//     fmt.Println("start  ",t.Format("03.04.05.0000000000"))

//     response, err := http.Get("http://localhost:3754/date")
//     if err != nil {
//         fmt.Printf(" HTTP request failed %s\n", err)
//     } else {
//         data, _ := ioutil.ReadAll(response.Body,)
//         fmt.Println(string(data))
//     }

// var b =1


// // for i := 0; i < 10; i++ {
// //     rand.Seed(time.Now().UnixNano())
// //     time.Sleep(500 * time.Millisecond)
// //     a := randomInt(1, 100) //get an int in the 1...n range
// //     o := randomoperator()
// //     b := randomInt(1, 100) //get an int in the 1...n range
// //     var request = request{
// //         Operator: o,
// //         Operate1: a,
// //         Operate2: b,
// //     }
// //     body, _ := json.Marshal(request)
// //     fmt.Println(string(body))
// // }


//     for b=a;b>3;b--{
        
//         post_data := []byte(`{"text": a}`)
//         request, _ := http.NewRequest("POST", "http://localhost:3754/postdate", bytes.NewBuffer(post_data))
//         request.Header.Set("Content-Type", "application/json")
//         client := &http.Client{}
//         response, err = client.Do(request)
//         if err != nil {
//             fmt.Printf("HTTP request failed r %s\n", err)
//         } else {
//             data, _ := ioutil.ReadAll(response.Body)
//             t := time.Now()
//             fmt.Println("Time =   ", t.Format("03.04.05.0000000000"))
//             fmt.Println(string(data))
//             // if string(data) == "Bad Requestpeet"{a=-1}
//             // fmt.Println(a)
//         }
// 	}
    
    

    
    
//     fmt.Println( response)
//     fmt.Println("start    ", t.Format("03.04.05.0000000000"))
// }
// // func sumab()int{
//     //    sum := `{"text": 5}`
//     //    return sum
//     // }

//     // func randomoperator() string {
//     //     rand.Seed(time.Now().Unix())
//     //     operator := []string{
//     //         "+", "-", "*", "/", "DIV", "MOD",
//     //     }
//     //     n := rand.Int() % len(operator)
//     //     // fmt.Println("Operate is ", operator[n])
//     //     return operator[n]
//     // }
    