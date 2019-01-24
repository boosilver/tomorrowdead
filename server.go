package main

import (
    "bytes"
    // "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    // "encoding/json"
	// "math/rand"
	
)

var a =5
func main() {
    t := time.Now()
    
    fmt.Println("start  ",t.Format("03.04.05.0000000000"))

    response, err := http.Get("http://localhost:3754/date")
    if err != nil {
        fmt.Printf(" HTTP request failed %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body,)
        fmt.Println(string(data))
    }

var b =1


// for i := 0; i < 10; i++ {
//     rand.Seed(time.Now().UnixNano())
//     time.Sleep(500 * time.Millisecond)
//     a := randomInt(1, 100) //get an int in the 1...n range
//     o := randomoperator()
//     b := randomInt(1, 100) //get an int in the 1...n range
//     var request = request{
//         Operator: o,
//         Operate1: a,
//         Operate2: b,
//     }
//     body, _ := json.Marshal(request)
//     fmt.Println(string(body))
// }


    for b=a;b>3;b--{
        
        post_data := []byte(`{"text": a}`)
        request, _ := http.NewRequest("POST", "http://localhost:3754/postdate", bytes.NewBuffer(post_data))
        request.Header.Set("Content-Type", "application/json")
        client := &http.Client{}
        response, err = client.Do(request)
        if err != nil {
            fmt.Printf("HTTP request failed r %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            t := time.Now()
            fmt.Println("Time =   ", t.Format("03.04.05.0000000000"))
            fmt.Println(string(data))
            // if string(data) == "Bad Requestpeet"{a=-1}
            // fmt.Println(a)
        }
	}
    
    

    
    
    fmt.Println( response)
    fmt.Println("start    ", t.Format("03.04.05.0000000000"))
}
// func sumab()int{
    //    sum := `{"text": 5}`
    //    return sum
    // }

    // func randomoperator() string {
    //     rand.Seed(time.Now().Unix())
    //     operator := []string{
    //         "+", "-", "*", "/", "DIV", "MOD",
    //     }
    //     n := rand.Int() % len(operator)
    //     // fmt.Println("Operate is ", operator[n])
    //     return operator[n]
    // }
    