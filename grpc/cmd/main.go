package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func main() {
	user := User{1, "Sudakov Aleksandr Aleksandrovich"}

	dataJson, _ := json.Marshal(&user)
	fmt.Println(len(dataJson))

	//dataPb, _ := proto.Marshal(&user)
	//fmt.Println(len(dataPb))
}
