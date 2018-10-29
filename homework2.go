package main

import (
	"fmt"
)

var fairyAnimals = []string{"fly", "spider", "bird", "cat", "dog", "cow", "horse"}

var swallowAnimal = map[string]string{
	"spider": "That wriggled and wiggled and tickled inside her.",
	"bird":   "How absurd to swallow a bird.",
	"cat":    "Fancy that to swallow a cat!",
	"dog":    "What a hog, to swallow a dog!",
	"cow":    "I don't know how she swallowed a cow!",
	"horse":  "...She's dead, of course!",
}

var animalsMap = map[string]int{
	"fly":    0,
	"spider": 1,
	"bird":   2,
	"cat":    3,
	"dog":    4,
	"cow":    5,
	"horse":  6,
}

var animalsMap2 = map[int]string{
	0: "fly",
	1: "spider",
	2: "bird",
	3: "cat",
	4: "dog",
	5: "cow",
	6: "horse",
}

func main() {
	story := getSongs(fairyAnimals)
	fmt.Println(story)
}

func getSongs(animals []string) []string {
	var songs []string
	for _, animal := range animals {
		song := getSong(animal)
		songs = append(songs, song)
	}
	return songs
}

func catchAnimal(animal string) string {
	var returnValues string
	if animalsMap[animal] != 0 && animalsMap[animal] != 6 {
		for i := animalsMap[animal]; i > 0; i-- {
			returnValue := fmt.Sprintf("She swallowed the %s to catch the %s,\n", animalsMap2[i], animalsMap2[i-1])
			returnValues += returnValue
		}
	}
	return returnValues
}

func getSong(animal string) string {
	var song string
	song1 := "There was an old lady who swallowed a " + animal + ". "
	song2 := swallowAnimal[animal]
	song3 := catchAnimal(animal)
	song4 := "I don't know why she swallowed a fly - perhaps she'll die!"

	if animal == "horse" {
		song = fmt.Sprintf("%s\n %s\n", song1, song2)
	} else if animal == "fly" {
		song = fmt.Sprintf("%s\n %s\n", song1, song4)
	} else {
		song = fmt.Sprintf("\n%s\n%s\n%s\n%s\n", song1, song2, song3, song4)
	}
	return song
}
