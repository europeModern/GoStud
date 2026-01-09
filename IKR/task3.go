package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	runes := []rune(word)
	if len(runes) <= 1 {
		return word
	}
	firstChar := string(runes[0])
	restRunes := runes[1:]
	
	for i, j := 0, len(restRunes)-1; i < j; i, j = i+1, j-1 {
		restRunes[i], restRunes[j] = restRunes[j], restRunes[i]
	}

	return firstChar + string(restRunes)
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)
	encryptedWords := make([]string, len(words))

	for i, word := range words {
		encryptedWords[i] = encryptWord(word)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {
	fmt.Println("ЗАДАЧА #3: Шифратор фраз")

	testPhrases := []string{
		"Pepe Schnele is a legend",
		"Hello world",
		"Go programming language",
		"a",
		"test",
		"Привет, Мир!"
	}
	
	for _, phrase := range testPhrases {
		encrypted := encryptPhrase(phrase)
		fmt.Printf("\nИсходная фраза: \"%s\"\n", phrase)
		fmt.Printf("Зашифрованная:  \"%s\"\n", encrypted)
	}
}
