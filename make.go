package main

import (
	"encoding/csv"
	// "log"
	"os"
)

// var a, b, c, d = "2", "100", "7+10", "1777777"
// var data = [][]string{
// 	{a, b, c, d},
// }
var testdata = [][]string{
	{"ID", "TIME", "EX", "RESULT"},
}

func save(data [][]string)string {
	var s string = "succes"
	file, err := os.Create("result.csv")
	checkError("connot creat file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// data2 := string(data)

	for _, value := range testdata { //head of table
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

	for _, value := range data { // data table
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

	// w := csv.NewWriter(os.Stdout)
	// w.WriteAll(testdata)
	// w.WriteAll(data) // calls Flush internally //แค่ให้มันอ่านไฟล์ที่กำหนด ให้มันมาโชว์บน consolelog
	return s
}
// func checkError(message string, err error) {
// 	if err != nil {
// 		log.Fatal(message, err)
// 	}
// }
