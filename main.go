package main

import (
	"fmt"
	"log"
)

func main() {
	docs, err := LoadDocs("test.json")

	if err != nil {
		log.Fatal(err)
	}

	idx := make(index)
	idx.add(docs)
	ids := idx.search("engine")

	s := fmt.Sprintf("the word is in: %v", ids)
	fmt.Println(s)
}
