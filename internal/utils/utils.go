package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(data))
}

func ApiRequest(method, url, body string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	return resp
}

func ExtractField(json, field string) string {
	start := strings.Index(json, `"`+field+`": "`) + len(field) + 4
	end := strings.Index(json[start:], `"`) + start
	return json[start:end]
}

func GetControllerAddress() string {
	address := os.Getenv("ZT_CONTROLLER_ADDRESS")
	if address == "" {
		address = "http://localhost:9993"
	}
	return address
}
