package main

import (
	"fmt"
	"time"

	"github.com/nukr/chaos"
	"github.com/nukr/dal-go"
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
	tableName := chaos.Alphabet(15)
	dbName := createTable(tableName, accessToken, serviceName)
	fmt.Printf("tableName %s\n", tableName)
	fmt.Printf("dbName %s\n", dbName)
	fmt.Printf("total cost %s\n", time.Since(start))
	d := dal.DAL{
		URL:         "http://localhost:12345/graphql",
		AccessToken: accessToken,
		ServiceName: serviceName,
	}
	s := struct {
		AA int
	}{
		AA: 1234,
	}

	result := d.CreateObject(tableName, s)
	fmt.Println(result)
}
