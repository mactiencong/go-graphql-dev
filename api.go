package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func apiHandler(resp http.ResponseWriter, req *http.Request) {
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: querySchema,
	})
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: req.URL.Query().Get("query"),
	})
	json.NewEncoder(resp).Encode(result)
}

var querySchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"person": &graphql.Field{
			Type: personType,
			Args: graphql.FieldConfigArgument{
				"ID": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
				return getPerson(rp.Args["ID"].(int)), nil
			},
		},
	},
})

// {
// 	person {
// 		id,
// 		name,
// 		age
// 	}
// }
var personType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Person",
	Description: "Person detail",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.Int,
			Description: "Person ID",
		},
		"Name": &graphql.Field{
			Type:        graphql.String,
			Description: "Person name",
		},
		"Age": &graphql.Field{
			Type:        graphql.Int,
			Description: "Person age",
		},
	},
})

// Person model
type Person struct {
	ID   int
	Name string
	Age  int
}

func getPerson(id int) Person {
	return Person{id, "matico", 29}
}

func main() {
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe("localhost:8080", nil)
}
