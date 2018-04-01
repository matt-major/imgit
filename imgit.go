package main

import (
	"bufio"
	"fmt"
	"os"
	"log"

	"github.com/matt-major/imgit/adaptors/tinypng"
)

func main() {
	var tinyPngKey = os.Getenv("IMGIT_TINY_KEY")
	
	if tinyPngKey == "" {
		log.Fatal("No TinyPNG API Key. Please export IMGIT_TINY_KEY=api:<api_key> and try again.")
	}

	imageFile, err := os.Open("test.png")
	
	if err != nil {
		log.Printf("Failed to load image file. %s", err)
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
