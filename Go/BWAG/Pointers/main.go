package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changUsingPointer(&myString)
	log.Println("after function call myString is set to", myString)
}

func changUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
