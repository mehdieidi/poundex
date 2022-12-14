package main

import (
	"fmt"
	"os"

	"github.com/mehdieidi/poundex/soundex"
	"github.com/mehdieidi/poundex/tokenizer"
)

func main() {
	tweetsFile, err := os.Open("tweets/Tweets.5000.json")
	if err != nil {
		panic(err)
	}
	defer tweetsFile.Close()

	lines, err := tokenizer.ReadLines(tweetsFile)
	if err != nil {
		panic(err)
	}

	var tweets []tokenizer.Tweet
	for _, l := range lines {
		t, err := tokenizer.ReadJSON(l)
		if err != nil {
			panic(err)
		}
		tweets = append(tweets, t)
	}

	tokensFile, err := os.OpenFile("tweets/tokens.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer tokensFile.Close()

	tokens := []string{}

	for _, t := range tweets {
		fileTokens, err := tokenizer.Tokenize(tokensFile, t.Text)
		if err != nil {
			fmt.Printf("err: %+v\n", err)
			continue
		}
		tokens = append(tokens, fileTokens...)
	}

	soundexFile, err := os.OpenFile("output/soundex.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer soundexFile.Close()

	soundexToPersian := map[string]map[string]struct{}{}
	soundexes := map[string]struct{}{}

	for _, t := range tokens {
		s, err := soundex.Get(t)
		if err != nil {
			continue
		}

		soundexes[s] = struct{}{}

		if soundexToPersian[s] == nil {
			soundexToPersian[s] = map[string]struct{}{}
		}

		soundexToPersian[s][t] = struct{}{}
	}

	for s := range soundexes {
		soundexFile.WriteString(fmt.Sprintf("%s\n", s))

		for persian := range soundexToPersian[s] {
			soundexFile.WriteString(fmt.Sprintf("[%s]", persian))
		}

		soundexFile.WriteString("\n----------------------------------------------------\n")
	}
}
