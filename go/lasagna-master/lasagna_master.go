package lasagna

// TODO: define the 'PreparationTime()' function

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

func PreparationTime(layers []string, time_per_layer int) int {
	if time_per_layer == 0 {
		time_per_layer = 2
	}
	return len(layers) * time_per_layer
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, x := range layers {
		if x == "noodles" {
			noodles += 50
		}
		if x == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friend, own []string) {
	own[len(own)-1] = friend[len(friend)-1]

}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var ans []float64 = []float64{}
	for _, x := range quantities {
		ans = append(ans, x*(float64(portions)/2.0))
	}
	return ans
}
