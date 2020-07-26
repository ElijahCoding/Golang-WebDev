package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

var authors []Author = []Author{
	Author{
		Id:        "author-1",
		Firstname: "Nic",
		Lastname:  "Raboy",
		Username:  "nraboy",
		Password:  "abc123",
	},
	Author{
		Id:        "author-2",
		Firstname: "Maria",
		Lastname:  "Raboy",
		Username:  "mraboy",
		Password:  "xyzabc",
	},
}

var articles []Article = []Article{
	Article{
		Id:      "article-1",
		Author:  "author-1",
		Title:   "Article #1",
		Content: "This is my first article",
	},
}

var rootQuery *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"authors": &graphql.Field{
			Type: graphql.NewList(authorType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				
			},
		},
	},
})

func main()  {
	fmt.Println("working")
}