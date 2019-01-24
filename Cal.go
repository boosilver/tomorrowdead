package main

import (
	// "encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type request struct {
	Operator string `json:"operator"`
	Operate1 int    `json:"operate1"`
	Operate2 int    `json:"operate2"`
}

func timestamp() {
	fmt.Println(time.Now().Format("03:04:0555"))
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

// func main() {
	
// }
