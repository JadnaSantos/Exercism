package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2 
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
//50 gramas de macarrão noodles
//0,2 litro de molho sauce
func Quantities(layers []string) (noodles int, sauce float64) {
     for _, elem := range layers {
         if elem == "noodles" {
             noodles += 50
         } else if elem == "sauce" {
             sauce += 0.2
         }
     }
    
    return
}
     
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe (qntForTwo []float64, numberPortins int) []float64 {
    output := make([]float64, len(qntForTwo)) 

    for i, elem := range qntForTwo {
        output[i] = elem * float64(numberPortins) / 2
    }

    return output
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
