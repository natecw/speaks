package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type strand struct {
	one   string
	two   string
	three string
}

func main() {
	fileArg := flag.String("f", "default.txt", "input text file")
	lenArg := flag.Int("size", 1000, "length of text to generate")
	flag.Parse()

	chain := make_chain(*fileArg)
	fmt.Println(generate(chain, *lenArg))
}

func (s strand) String() string {
	return s.one + " " + s.two + " " + s.three
}

func generate(chain map[strand][]string, length int) string {
	curr_key := rand_key(chain)
	var text strings.Builder
	text.WriteString(curr_key.String())
	fmt.Println("curr_key=", curr_key)
	size := 2
	for size < length {
		words := chain[curr_key]
		next_word := ""
		if len(words) != 0 {
			next_word = words[rand.Intn(len(words))]
		} else {
			next_word = rand_key(chain).three
		}

		curr_key = strand{curr_key.two, curr_key.three, next_word}
		text.WriteString(" " + next_word)
		size++
	}
	return text.String()
}

func rand_key(chain map[strand][]string) strand {
	n := rand.Int31n(int32(len(chain)))
	i := int32(0)
	for k := range chain {
		if i == n {
			return k
		}
		i++
	}
	return strand{}
}

func make_chain(filePath string) map[strand][]string {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	replacer := strings.NewReplacer("\n", "", "\"", "")
	words := strings.Split(string(buf), " ")
	chain := make(map[strand][]string)
	for i := 2; i < len(words)-1; i++ {
		key := strand{
			replacer.Replace(words[i-2]),
			replacer.Replace(words[i-1]),
			replacer.Replace(words[i]),
		}
		if i+1 < len(words)-1 {
			chain[key] = append(chain[key], words[i+1])
		}
	}
	return chain
}
