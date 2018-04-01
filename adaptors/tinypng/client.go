package tinypng

import (
	"os"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	"github.com/matt-major/imgit/utils"
)

const (
	tinyPNGURL = "https://api.tinify.com/shrink"
	tinyPngMethod = "POST"
)

// Create Request.Create sends new image to TinyPNG, returns parsed response body
func (r Request) Create() *Response {
	var tinyPngKey = os.Getenv("IMGIT_TINY_KEY")

	serializedBodyReader := bytes.NewReader(r.body)

	request, err := http.NewRequest(tinyPngMethod, tinyPNGURL, serializedBodyReader)
	
	if err != nil {
		log.Printf("Failed to create TinyPNG request. %s", err)
		return nil
	}

	encodedCredentials := utils.ConvertToBase64(tinyPngKey)
	authHeader := fmt.Sprintf("Basic %s", encodedCredentials)
	request.Header.Add("Authorization", authHeader)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Printf("Failed to upload image to TinyPNG. %s", err)
		return nil
	}

	defer response.Body.Close()

	parsed := &Response{}
	err = json.NewDecoder(response.Body).Decode(parsed)
	

	return parsed
}
