package main

import (
	"fmt"

	"github.com/josephphyo/words/quote"
	"github.com/josephphyo/words/word"
)

func main() {
	fmt.Println("Total Word Count:", word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
