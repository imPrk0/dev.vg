package util

import (
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomWords(count int) []string {
	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = randomdata.FirstName(randomdata.RandomGender)
	}
	return words
}

func GenerateRandomSentence(wordCount int) []string {
	var sentence []string
	for i := 0; i < wordCount; i++ {
		sentences := strings.Split(
			strings.Replace(
				randomdata.Paragraph(),
				". ",
				".",
				-1,
			),
			".",
		)
		sentence = append(
			sentence,
			sentences[0]+".",
		)
	}
	return sentence
}

func GenerateRandomParagraph(sentenceCount int) []string {
	var paragraph []string
	for i := 0; i < sentenceCount; i++ {
		paragraph = append(
			paragraph,
			randomdata.Paragraph(),
		)
	}
	return paragraph
}

func GenerateRandomText(count int) []string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]string, 1)
	randomString := make([]byte, count)
	for i := 0; i < count; i++ {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	result[0] = string(randomString)

	return result
}
