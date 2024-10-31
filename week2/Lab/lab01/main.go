package main

import (
	"fmt"
)

func main() {
	list := []User{
		{"Paul", 19},
		{"John", 21},
		{"Jane", 35},
		{"Abraham", 25},
	}

	sorting(&list)
	for _, user := range list {
		fmt.Println(user.Name)
	}
}
