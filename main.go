package main

import (
	"fmt"

	"github.com/josephphyo/LearningGo/NinjaLevelExercises/ninja-level10-test/exe2/quote"
	"github.com/josephphyo/LearningGo/NinjaLevelExercises/ninja-level10-test/exe2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
