package main

import "fmt"

func mostRepeatedWords(words []string) {

	repeatedWord := ""
	mostRepeat := 0

	// range ile ilk önce değerleri alıyorum.
	for _, currentWord := range words {
		repeat := 0
		// burda aldığım değerlerin sayısını buluyorum.
		for _, sameWord := range words {
			if currentWord == sameWord {
				repeat++
			}
		}
		// burda mostRepeat guncelliyorum. mostRepeat scope dışı olduğu için en son değer yukarda tutuluyor.
		if repeat > mostRepeat {
			mostRepeat = repeat
			repeatedWord = currentWord
		}
	}
	fmt.Println("Most Repeated Word is: ", repeatedWord)
}

func main() {

	words := []string{"apple", "pie", "apple", "red", "red", "red"}

	mostRepeatedWords(words)

}
