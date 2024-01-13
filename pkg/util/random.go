package util

import (
	"github.com/Pallinder/go-randomdata"
	"strings"
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
			randomdata.Paragraph(),
			".",
		)
		sentence = append(
			sentence,
			sentences[randomdata.Number(
				0,
				len(sentences)-1,
			)],
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

func GenerateRandomText(paragraphCount int) []string {
	var text []string
	for i := 0; i < paragraphCount; i++ {
		text = append(
			text,
			randomdata.FirstName(randomdata.RandomGender),
		)
	}
	return text
}
