package main

import "fmt"

func main() {

	m := make(map[string]string)
	fmt.Println(m)

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println(countryCapitalMap)
}
