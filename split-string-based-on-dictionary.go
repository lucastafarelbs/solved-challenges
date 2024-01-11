package main

import (
	"fmt"
	"strings"
)

func checkWordExistence(word string, dictionary []string) bool {
	if word == "" || len(dictionary) <= 0 {
		return false
	}

	for _, v := range dictionary {
		if word == v {
			return true
		}
	}
	return false
}

func validateArray(strArr []string) string {
	dictionary := strings.Split(strArr[1], ",")

	var foundWords []string
	word := ""
	for _, v := range strArr[0] {
		word = fmt.Sprintf("%s%c", word, v)

		exists := checkWordExistence(word, dictionary)
		if exists {
			foundWords = append(foundWords, word)
		}
	}

	for i := range strArr[0] {
		word = strArr[0][i:]
		exists := checkWordExistence(word, dictionary)
		if exists {
			foundWords = append(foundWords, word)
		}
	}

	var possibleCombinations []string
	for i, v := range foundWords {
		firstWord := v
		for j := i + 1; j < len(foundWords); j++ {
			lastWord := foundWords[j]
			combination := firstWord + lastWord
			if len(combination) == len(strArr[0]) {
				commaSeparatedCombination := fmt.Sprintf("%s,%s", v, foundWords[j])
				possibleCombinations = append(possibleCombinations, commaSeparatedCombination)
			}
		}
	}

	if len(possibleCombinations) > 0 {
		possibleSplit := possibleCombinations[0]
		fmt.Println("word", word)
		fmt.Println("dictionary", dictionary)
		fmt.Println("possible split: ", possibleSplit)
		return possibleSplit
	}

	return "not possible"

}

func main() {
	// the challenge is:
	// using an array of two indexes containg a word and dictionary
	// split the word in two other strings that matches with words in dictionary
	// using this example, the possible split for "baseball" will be "base,ball"
	// if a possible split was found, return it, if not, return "not possible"

	strArr := []string{"baseball", "b,base,ball,code,x,z,y,lucas"}
	fmt.Println(validateArray(strArr))

}
