package lasagna

const OvenTime = 40

func RemainingOvenTime(alreadyIn int) int {
	return OvenTime - alreadyIn
}

func PreparationTime(layers int) int {
	return layers * 2
}

func ElapsedTime(layers, alreadyIn int) int {
	return alreadyIn + PreparationTime(layers)
}
