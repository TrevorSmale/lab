package main

import "log"

func main() {
	var isTrue bool
	isTrue = true 

	if isTrue {
		log.Println("isTrue", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
}