package main

import (
	"fmt"
	"strings"
)

func main() {
	test1()
	fmt.Println()
	test2()
	fmt.Println()
	test3()
}

func test1() {
	fmt.Print("Input Number: ")
	var input int
	fmt.Scanf("%d", &input)

	tempArray := []string{}
	for i := 0; i < input; i++ {
		tempArray = append(tempArray, fmt.Sprintf("%v", i+1))
	}

	fmt.Println(strings.Join(tempArray, ","))
}

func test2() {
	fmt.Print("Input Number: ")
	var input int
	fmt.Scanf("%d", &input)
	if input%2 == 0 {
		fmt.Println("Bilangan Genap")
	} else {
		fmt.Println("Bilangan Ganjil")
	}
}

var key = [][]string{
	{"ha", "na", "ca", "ra", "ka"},
	{"da", "ta", "sa", "wa", "la"},
	{"pa", "dha", "ja", "ya", "nya"},
	{"ma", "ga", "ba", "ta", "nga"},
}

func test3() {
	fmt.Print("Input Text: ")
	var inputText string
	fmt.Scanf("%s", &inputText)

	changeVocal, replacedIndex := changeVocal(inputText)
	split := splitString(changeVocal)
	encoding := encodeStringToDagadu(inputText, replacedIndex, split)
	fmt.Println(fmt.Sprintf("Translate: %s", encoding))
}

func changeVocal(word string) (string, []int) {
	runeWord := []rune(word)
	replacedIndex := []int{}
	for i := 0; i < len(word); i++ {
		if string(word[i]) == "i" || string(word[i]) == "u" || string(word[i]) == "e" || string(word[i]) == "o" {
			runeWord[i] = 'a'
			replacedIndex = append(replacedIndex, i)
		}
	}
	if string(word[len(word)-1]) != "a" && string(word[len(word)-1]) != "i" && string(word[len(word)-1]) != "u" && string(word[len(word)-1]) != "e" && string(word[len(word)-1]) != "o" {
		runeWord = append(runeWord, 'a')
	}
	return string(runeWord), replacedIndex
}

func subStringInsideKey(substring string) (string, int, int) {
	for i := 0; i < len(key); i++ {
		for j := 0; j < len(key[i]); j++ {
			if substring == key[i][j] {
				return key[i][j], i, j
			}
		}
	}

	return "", -1, -1
}

func splitString(word string) []string {
	tempArray := []string{}
	index := 0
	for index < len(word) {
		subString := ""
		if index < len(word)-1 {
			subString = string(word[index : index+2])
			result, _, _ := subStringInsideKey(subString)
			if result != "" {
				tempArray = append(tempArray, result)
				index += 2
			} else {
				subString = string(word[index : index+3])
				result, _, _ := subStringInsideKey(subString)
				if result != "" {
					tempArray = append(tempArray, result)
					index += 3
				}
			}
		}
	}
	return tempArray
}

func encodeStringToDagadu(input string, replacedIndex []int, subStrings []string) string {
	result := []string{}
	for i := 0; i < len(subStrings); i++ {
		_, x, y := subStringInsideKey(subStrings[i])

		var newX int

		if x == 0 {
			newX = 2
		} else if x == 1 {
			newX = 3
		} else if x == 2 {
			newX = 0
		} else if x == 3 {
			newX = 1
		}
		result = append(result, key[newX][y])
	}

	resultString := strings.Join(result, "")
	if string(input[len(input)-1]) != "a" && string(input[len(input)-1]) != "i" && string(input[len(input)-1]) != "u" && string(input[len(input)-1]) != "e" && string(input[len(input)-1]) != "o" {
		if len(replacedIndex) > 0 {
			for i := 0; i < len(replacedIndex); i++ {
				replacement := []rune(string(input[replacedIndex[i]]))
				resultString = resultString[:replacedIndex[i]] + string(replacement) + resultString[replacedIndex[i]+1:]
			}
		}
		return resultString[:len(resultString)-1]
	} else {
		if len(replacedIndex) > 0 {
			for i := 0; i < len(replacedIndex); i++ {
				replacement := []rune(string(input[replacedIndex[i]]))
				resultString = resultString[:replacedIndex[i]] + string(replacement) + resultString[replacedIndex[i]+1:]
			}
			return resultString
		} else {
			return resultString[:len(input)]
		}
	}
}
