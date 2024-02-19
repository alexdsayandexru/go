package main

import "fmt"

func myprint(ii ...interface{}) {
	for _, i := range ii {
		fmt.Println(i)
	}

}

func main() {
	myprint(1, 2, 3, "4")
}

//go run main.go
//go build main.go
//.\main.exe
