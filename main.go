package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// https://twitter.com/arufa_faru/status/850689005746143232
	rand.Seed(time.Now().UnixNano())
	homestr := []string{
		"すごい",
		"イカす",
		"かっこいい",
		"COOL",
		"かわいい",
		"バッチグー",
		"ハンサム",
		"リーダーシップがある",
		"意外とモテるでしょ？",
		"しびぃ～",
	}
	n := flag.Int("n", 10, "ホメ回数")
	p := flag.String("p", "すごいひとは", "ほめたいもの")
	flag.Parse()

	for i := 0; i < *n; i++ {
		fmt.Printf("%s%s\n", *p, homestr[rand.Intn(len(homestr))])
	}

}
