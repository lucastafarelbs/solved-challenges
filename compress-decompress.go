package main

import (
	"fmt"
	"strconv"
)

// receive: luuccaass
// give: l1u2c2a2s2

func compress(s string) string {
	current := ""
	currentCount := 0
	newS := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == current || current == "" {
			current = string(s[i])
			currentCount++
			continue
		}

		newS = fmt.Sprintf("%s%s%d", newS, current, currentCount)
		current = string(s[i])
		currentCount = 1
	}
	newS = fmt.Sprintf("%s%s%d", newS, current, currentCount)

	return newS
}

func decompress(s string) string {
	if s == "" {
		return s
	}

	count := 0
	newS := ""
	for i := 0; i < len(s)-1; i++ {
		count, _ = strconv.Atoi(string(s[i+1]))
		for j := 0; j < count; j++ {
			newS = fmt.Sprintf("%s%s", newS, string(s[i]))
		}
		i++
	}
	return newS
}

func main() {
	fmt.Println("--- ---")
	s := "luuccaass"
	s2 := "aannooottheeeeeerrrrrrrr aaastringgg222"
	sComp := compress(s)
	sComp2 := compress(s2)
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(sComp)
	fmt.Println(sComp2)
	fmt.Println(decompress(sComp))
	fmt.Println(decompress(sComp2))
}
