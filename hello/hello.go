package hello

import "fmt"

func Hello(name string) {
	result := fmt.Sprintf("hi, %s!!!!!!!!", name)
	fmt.Println(result)
}