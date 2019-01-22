package main

import (
	"encoding/csv"
	"log"
	"os"
)

var a, b, c, d = "2", "100", "7+10", "17"
var data = [][]string{
	{"1", "50", "1+1", "2"},
	{a, b, c, d},
}
var testdata = [][]string{
	{"ID", "TIME", "EX", "RESULT"},
}

func main() {
	file, err := os.Create("result.csv")
	checkError("connot creat file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range testdata { //head of table
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

	for _, value := range data { // data table
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(testdata)
	w.WriteAll(data) // calls Flush internally

}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
