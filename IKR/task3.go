package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) <= 1 {
		return word
	}
	
	firstChar := string(word[0])
	restChars := word[1:]
	
	runes := []rune(restChars)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	return firstChar + string(runes)
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
	}
	
	for _, phrase := range testPhrases {
		encrypted := encryptPhrase(phrase)
		fmt.Printf("\nИсходная фраза: \"%s\"\n", phrase)
		fmt.Printf("Зашифрованная:  \"%s\"\n", encrypted)
	}
}
