package lasagna


// TODO: define the 'OvenTime' constant
const OvenTime = 40 // lasanha deve ficar no forno

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven
	panic("RemainingOvenTime not implemented") // retorna o tempo que a lasanha ficou no forno - ovenTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    return numberOfLayers * 2
	panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    return PreparationTime(numberOfLayers) + actualMinutesInOven
	panic("ElapsedTime not implemented")
}
