package slinky_test

import (
	"fmt"

	"github.com/FurqanSoftware/slinky"
)

func ExampleParse() {
	u, err := slinky.Parse("https://github.com/hjr265")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Service)
	fmt.Println(u.Type)
	fmt.Println(u.ID)
	fmt.Println(u.Data["username"])
	// Output:
	// GitHub
	// User
	// hjr265
	// hjr265
}
