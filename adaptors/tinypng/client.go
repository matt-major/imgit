package tinypng

import (
	"os"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matt-major/imgit/utils"
)

const (
	tinyPNGURL = "https://api.tinify.com/shrink"
	tinyPngMethod = "POST"
)

func (r Request) Create() *Response {
	var tinyPngKey = os.Getenv("IMGIT_TINY_KEY")
	if tinyPngKey == "" {
		panic("NO TINYPNG KEY!")
	}

	serializedBodyReader := bytes.NewReader(r.body)

	request, err := http.NewRequest(tinyPngMethod, tinyPNGURL, serializedBodyReader)
	if err != nil {
		return nil
	}

	encodedCredentials := utils.ConvertToBase64(tinyPngKey)
	authHeader := fmt.Sprintf("Basic %s", encodedCredentials)
	request.Header.Add("Authorization", authHeader)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil
	}

	parsed := &Response{}
	err = json.NewDecoder(response.Body).Decode(parsed)

	defer func() {
		response.Body.Close()
	}()

	return parsed
}
