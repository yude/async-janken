package main

import (
	"log"
)

func main() {
	InitDB()
	log.Fatal(Handler().Listen(":3000"))
}
