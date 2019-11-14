package main

import (
	"fmt"
)

func main() {
	var cityArray = [4]string{"北京", "上海", "深圳", "广州"}
	fmt.Println(cityArray)

	var boolArray = [...]bool{true, false, true, true, false, false}
	fmt.Println(boolArray)

	var languageArray = [...]string{1: "Golang", 3: "Java", 5: "Python", 8: "C"}
	fmt.Println(languageArray)
	fmt.Printf("%T\n", languageArray)

}
