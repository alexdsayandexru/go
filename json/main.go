package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type User struct {
	Name       string
	Occupation string
	Born       string
}

func main() {
	filename, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}

	defer filename.Close()

	data, err := io.ReadAll(filename)

	if err != nil {
		log.Fatal(err)
	}

	var result []User

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(result)
}

//go run main.go
//go build main.go
//.\main.exe
