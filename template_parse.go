package main

import (
	"bytes"
	"html/template"
	"log"
)

func templateParse(s string, variable interface{}) string {
	buf := make([]byte, 1024)
	writer := bytes.NewBuffer(buf)
	tmpl, err := template.New("graphql").Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(writer, variable)
	queryString := string(bytes.TrimLeft(writer.Bytes(), "\x00"))
	return queryString
}
