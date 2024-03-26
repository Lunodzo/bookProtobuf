package main

import (
	"encoding/json"
	"fmt"
	"go/protobuf/model/model"
	"os"

	"google.golang.org/protobuf/proto"
)

func main(){
	book := &model.Book{
		Id: 1,
		Title: "Go Programming",
		Author: []*model.Author{
			{Id: 1, Name: "Lunodzo"},
		},
		Category: model.Category_FICTION,
	}
	data, err := proto.Marshal(book)
	if err != nil {
		panic(err)
	}
	
	// Write to file
	os.WriteFile("book.proto", data, 0644)

	// Create json
	data, err = json.Marshal(book)
	if err != nil {
		panic(err)
	}
	os.WriteFile("book.json", data, 0644)


	// Decode the data from protobuf bytes
	data, err = os.ReadFile("book.proto")
	if err != nil {
		panic(err)
	}

	readFile := model.Book{}
	if err := proto.Unmarshal(data, &readFile); err != nil {
		panic(err)
	}
	println("All Content: ", readFile.String())
	fmt.Printf("Book Title: %s\n", readFile.GetTitle())
}