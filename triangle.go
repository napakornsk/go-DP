package main

import (
	"encoding/json"
	"log"
	"os"
)

func getTriangle() [][]int {

	data, err := os.ReadFile("./json/hard.json")
	if err != nil {
		log.Fatal(err)
	}
	var triangle [][]int

	err = json.Unmarshal(data, &triangle)
	if err != nil {
		log.Fatal(err)
	}
	return triangle
}
