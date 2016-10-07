package main

import "github.com/nukr/grafetch"

func createAccount(username, password string) (accessToken string) {
	query := `
  mutation {
    createAccount(username: "{{.Username}}", password: "{{.Password}}") {
      accessToken
    }
  }
  `
	s := struct {
		Username string
		Password string
	}{
		Username: username,
		Password: password,
	}
	queryString := templateParse(query, s)

	graphql := grafetch.New("http://localhost:12345/graphql")
	graphql.SetQuery(grafetch.GraphQLQuery{
		Query: queryString,
	})
	var resp struct {
		Data struct {
			CreateAccount struct {
				AccessToken string `json:"accessToken"`
			} `json:"createAccount"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
	}
	graphql.Fetch(&resp)
	return resp.Data.CreateAccount.AccessToken
}
