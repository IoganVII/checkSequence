package main

import "fmt"

func main() {

	a := []int{1, 3, 6, 4, 2, 8, 5, 7}
	result := Solution(a)
	fmt.Println(result)

}

// Вот это уже сосвем флагами обмазываюсь, как так-то(
func Solution(inputArray []int) int {
	minElement := inputArray[0]
	maxElement := inputArray[0]
	isSequence := true

	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] < minElement {
			minElement = inputArray[i]
		}
		if inputArray[i] > maxElement {
			maxElement = inputArray[i]
		}
	}
	var countInput int
	for elemtnValue := minElement; elemtnValue <= maxElement; elemtnValue++ {
		countInput = 0
		for j := 0; j < len(inputArray); j++ {
			if inputArray[j] == elemtnValue {
				countInput++
				if countInput > 1 {
					break
				}
			}
		}
		if countInput == 0 || countInput > 1 {
			isSequence = false
			break
		}
	}

	if isSequence {
		return 1
	} else {
		return 0
	}

}
