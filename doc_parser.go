package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//Doc is the representation of a real document. Using JSON by now.
type Doc struct {
	ID    int `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func LoadDocs(path string) ([]Doc, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	d := struct {
		Docs []Doc `json:"docs"`
	}{}
	dec := json.NewDecoder(f)

	err = dec.Decode(&d)

	if err != nil {
		return nil, err
	}

	return d.Docs, nil
}

func genDocID() string {
	src := rand.NewSource(time.Now().UnixNano())
	rand.Seed(src.Int63())
	i1 := rand.Int63()
	rand.Seed(i1 % 31)
	x1 := rand.Int63()
	rand.Seed(x1 % 13)
	i2 := rand.Int63()
	rand.Seed(i2 % 59)
	x2 := rand.Int63()
	s1 := fmt.Sprintf("%d", i1)
	s2 := fmt.Sprintf("%x", x1)
	s3 := fmt.Sprintf("%d", i2)
	s4 := fmt.Sprintf("%x", x2)
	return fmt.Sprintf("%.6s-%.3s-%.6s-%.6s", s1, s2, s3, s4)
}
