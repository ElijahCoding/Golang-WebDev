package main

import (
	"encoding/json"
	"fmt"
)

var authors []Author = []Author{
	Author{
		Id:        "author-1",
		Firstname: "Nic",
		Lastname:  "Raboy",
		Username:  "nraboy",
		Password:  "pass",
	},
	Author{
		Id:        "author-2",
		Firstname: "Maria",
		Lastname:  "Raboy",
		Username:  "mraboy",
		Password:  "abc123",
	},
}

var articles []Article = []Article{
	Article{
		Id:      "article-1",
		Author:  "author-1",
		Title:   "Blog Post 1",
		Content: "This is an example blog article",
	},
}

func main() {
	fmt.Println("starting")
	fmt.Println(authors)
	//fmt.Println(articles)

	data, _ := json.Marshal(authors)
	print(string(data))
}
