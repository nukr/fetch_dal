package main

import "github.com/nukr/grafetch"

func createService(serviceName, accessToken string) string {
	query := `
  mutation {
    createService(name: "{{.ServiceName}}") {
      id
      serviceName
    }
  }
  `
	s := struct {
		ServiceName string
	}{
		ServiceName: serviceName,
	}
	queryString := templateParse(query, s)

	graphql := grafetch.New("http://localhost:12345/graphql")
	graphql.SetHeader("x-meepcloud-access-token", accessToken)
	graphql.SetQuery(grafetch.GraphQLQuery{
		Query: queryString,
	})
	var resp struct {
		Data struct {
			CreateService struct {
				ID          string `json:"id"`
				ServiceName string `json:"serviceName"`
			} `json:"createService"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
	}
	graphql.Fetch(&resp)
	return resp.Data.CreateService.ServiceName
}
