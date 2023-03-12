package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	adjectives := []string{
		"happy", "funny", "sad", "silly", "brave", "crazy", "angry", "shiny",
		"spooky", "tiny", "fierce", "cute", "jolly", "funky", "wild", "yummy",
		"hairy", "zesty", "lively", "smiling",
	}

	colors := []string{
		"red", "blue", "green", "yellow", "purple", "orange", "pink", "black",
		"white", "brown", "gray", "gold", "silver", "beige", "turquoise",
		"lavender", "maroon", "teal", "magenta", "navy",
	}

	animals := []string{
		"cat", "dog", "bird", "fish", "rabbit", "hamster", "lion", "tiger",
		"bear", "elephant", "monkey", "giraffe", "wolf", "snake", "crocodile",
		"shark", "whale", "octopus", "penguin", "koala",
	}

	rand.Seed(time.Now().UnixNano())

	rando := func(words []string) string { return words[rand.Intn(len(words))] }

	w1 := rando(adjectives)
	w2 := rando(colors)
	w3 := rando(animals)

	fmt.Printf("%s-%s-%s\n", w1, w2, w3)
}
