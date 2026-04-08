package main

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	highestCount := 0
	popularWord := "n/a"

	for key, value := range words {

		if value > highestCount {
			highestCount = value
			popularWord = key
		}

	}

	println("Most popular word: ", popularWord)
	println("With a count of: ", highestCount)

}
