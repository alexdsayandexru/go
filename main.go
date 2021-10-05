package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := "Hello World, Привет Мир, 世界 !!!"
	fmt.Fprintln(w, getHostname()+text)
	fmt.Println(time.Now(), text)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	text := "Hello World, Hello World, Hello World !!!"
	fmt.Fprintln(w, getHostname()+text)
	fmt.Println(time.Now(), text)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	text := "Привет, Привет, Привет !!!"
	fmt.Fprintln(w, getHostname()+text)
	fmt.Println(time.Now(), text)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	text := "世界 世界 世界 ±±±"
	fmt.Fprintln(w, getHostname()+text)
	fmt.Println(time.Now(), text)
}

func handlerAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, getHostname()+"Привет мир, Привет мир, Привет мир!!!")
	fmt.Fprintln(w, getHostname()+"世界 世界 世界 !!!")
	fmt.Fprintln(w, getHostname()+"Hello World, Hello World, Hello World !!!")
	fmt.Fprintln(w, getHostname()+"Hello World, Привет Мир, 世界 !!!")
}

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("hostname:", name)
	return "[" + name + "] "
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/en", handler1)
	http.HandleFunc("/ru", handler2)
	http.HandleFunc("/ch", handler3)
	http.HandleFunc("/all", handlerAll)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
