package main

import "fmt"

func findLargestRanchStock[K comparable, V int | float64](m map[K]V) K {
	var stock V
	var name K
	for k, v := range m {
		if v > stock {
			stock = v
			name = k
		}
	}
	return name
}

func main() {
	animalStock := map[string]int{
		"Chicken": 5,
		"Cattle":  29,
		"Horses":  4,
	}
	miscStock := map[string]float64{
		"Hay":        5.5,
		"Feed":       1.2,
		"Fertilizer": 3.7,
	}
	largestStockOnRanchInt := findLargestRanchStock(animalStock)
	fmt.Printf("The largest stock item on the ranch is %s\n", largestStockOnRanchInt)

	largestStockOnRanchFloat := findLargestRanchStock(miscStock)
	fmt.Printf("The largest stocked item on the rach is %s\n", largestStockOnRanchFloat)
}
