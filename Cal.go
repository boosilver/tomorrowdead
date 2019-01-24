package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type request struct {
	Operator string `json:"operator"`
	Operate1 int    `json:"operate1"`
	Operate2 int    `json:"operate2"`
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

func main() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(500 * time.Millisecond)
		a := randomInt(1, 100) //get an int in the 1...n range
		o := randomoperator()
		b := randomInt(1, 100) //get an int in the 1...n range
		var request = request{
			Operator: o,
			Operate1: a,
			Operate2: b,
		}
		body, _ := json.Marshal(request)
		fmt.Println(string(body))
	}
}
