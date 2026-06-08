package lasagnamaster

func PreparationTime(sliceLayers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		return len(sliceLayers) * 2
	}
	return len(sliceLayers) * avgPrepTime
}

func Quantities(sliceLayers []string) (int, float64) {
	souceQty, noodlesQty := 0.0, 0
	for i := 0; i < len(sliceLayers); i++ {
		if sliceLayers[i] == "noodles" {
			noodlesQty++
		} else if sliceLayers[i] == "sauce" {
			souceQty++
		}
	}
	return noodlesQty * 50, souceQty * 0.2
}

func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(qtyNeeded []float64, portionsNum int) []float64 {
	scale := (float64(portionsNum) / 2.0)
	var newQtyNeeded []float64
	for i := 0; i < len(qtyNeeded); i++ {
		newQtyNeeded = append(newQtyNeeded, qtyNeeded[i]*(scale))
	}
	return newQtyNeeded
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
