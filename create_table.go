package main

import "github.com/nukr/grafetch"
import "github.com/nukr/template_parse"

func createTable(tableName, token, serviceName string) string {
	query := `
  mutation {
    createTable(name: "{{.TableName}}") {
      db
    }
  }
  `
	s := struct {
		TableName string
	}{
		TableName: tableName,
	}
	queryString := tp.TemplateParse(query, s)

	graphql := grafetch.New("http://localhost:12345/graphql")
	graphql.SetHeader("x-meepcloud-access-token", token)
	graphql.SetHeader("x-meepcloud-service-name", serviceName)
	graphql.SetQuery(grafetch.GraphQLQuery{
		Query: queryString,
	})
	var resp struct {
		Data struct {
			CreateTable struct {
				DB string `json:"db"`
			} `json:"createTable"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
	}
	graphql.Fetch(&resp)
	return resp.Data.CreateTable.DB
}
