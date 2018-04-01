package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matt-major/imgit/adaptors/tinypng"
)

func main() {
	imageFile, err := os.Open("test.png")
	
	if err != nil {
		fmt.Println("Failed to open image!")
		fmt.Println(err)
	}
	
	defer imageFile.Close()

	fileStats, _ := imageFile.Stat()
	var size = fileStats.Size()
	var bytes = make([]byte, size)

	buffer := bufio.NewReader(imageFile)
	_, err = buffer.Read(bytes)

	request := tinypng.NewRequest()
	request.SetBody(bytes)

	response := request.Create()
	fmt.Println(response)
}
