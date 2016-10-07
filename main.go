package main

import (
	"fmt"
	"time"

	"github.com/nukr/chaos"
)

func main() {
	start := time.Now()
	username := chaos.Alphabet(12)
	password := chaos.Alphabet(12)
	fmt.Printf("username: %s, password %s\n", username, password)
	accessToken := createAccount(username, password)
	fmt.Printf("accessToken %s\n", accessToken)
	serviceName := createService(chaos.Alphabet(15), accessToken)
	fmt.Printf("serviceName %s\n", serviceName)
	tableName := createTable(chaos.Alphabet(15), accessToken, serviceName)
	fmt.Printf("tableName %s\n", tableName)
	fmt.Printf("total cost %s\n", time.Since(start))
}
