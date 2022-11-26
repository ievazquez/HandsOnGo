package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func getRandomProverb() string {
	//fmt.Println(proverbs.Random())
	//fmt.Println(proverbs.Random())
	//fmt.Println(proverbs.Random())
	return proverbs.Random().Saying
}

func main() {
	fmt.Println(getRandomProverb())
}
