package main

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

type User struct {
	Name  string `color:"blue"`
	Age   int    `color:"red"`
	Email string
}

func PrintWithColor(u interface{}) {
	ru := reflect.ValueOf(u)
	ut := ru.Type()

	fmt.Printf("%s: ", ut.Name())
	for i := 0; i < ut.NumField(); i++ {
		switch clr := ut.Field(i).Tag.Get("color"); clr {
		case "blue":
			color.Set(color.FgBlue)
		case "red":
			color.Set(color.FgRed)
		case "yellow":
			color.Set(color.FgYellow)
		}

		fmt.Printf("%v ", ru.Field(i))
		color.Unset()
	}

	fmt.Println()
}

func main() {
	u1 := User{"Joe", 34, "example@acne.com"}
	u2 := User{"Mary", 46, "mary@acne.com"}
	u3 := User{"Alice", 15, "alice@acne.com"}
	u4 := User{"Bob", 18, "bob@acne.com"}

	PrintWithColor(u1)
	PrintWithColor(u2)
	PrintWithColor(u3)
	PrintWithColor(u4)
}
