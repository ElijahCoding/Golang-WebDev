package utils

func BubbleSort(elements []int) []int {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements); i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
	return elements
}