package main

import (
    "bytes"
    // "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

var a =5
func main() {
    t := time.Now()
    
    fmt.Println("start  ",t)

    response, err := http.Get("http://localhost:3754/date")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body,)
        fmt.Println(string(data))
    }

var b =1
    for b=a;b>-5000;b--{
        
        post_data := []byte(`{"text": 5}`)
        request, _ := http.NewRequest("POST", "http://localhost:3754/postdate", bytes.NewBuffer(post_data))
        request.Header.Set("Content-Type", "application/json")
        client := &http.Client{}
        response, err = client.Do(request)
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            t := time.Now()
            fmt.Println("Time =   ", t.Format("03.04.05.0000000000"))
            fmt.Println(string(data))
            if string(data) == "Bad Requestpeet"{a=-1}
            // fmt.Println(a)
        }
	}
    
    

    
    
    fmt.Println("end    ", response)
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
    