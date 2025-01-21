package main

import "log"

func main()  {
	
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
		
	case "fish":
		log.Println("cat is set to fish")

	case "tree":
		log.Println("cat is set to tree")

	case "car":
		log.Println("cat is set to car")

	default:
		log.Println("cat is set to something else")
	}
}