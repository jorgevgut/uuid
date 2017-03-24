package main

import (
	"fmt"
	"log"
)

func main() {
	uuid, err := NewUUID()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(uuid)
}
