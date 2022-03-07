package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")

	// ioutil

	// WriteFileで書き込み
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	// ReadFileで読み込み
	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))

	// os.File

	// Createで作成してWriteで書き込み
	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	// Openで開いてReadで読み込み
	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes fromm file\n", bytes)
	fmt.Println(string(read2))
}
