package main

import "log"

func main() {
	myNum := 80
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to Not True")
	} else if myNum < 100 && !isTrue {
		log.Println("myNum is less than 100 and isTrue is set to Not True")
	} else if myNum == 101 || isTrue {
		log.Println("myNum is equal to 101")
	} else if myNum > 1000 && isTrue == false {
		log.Println("myNum is greater than 1000 and isTrue is set to True")
	}
}