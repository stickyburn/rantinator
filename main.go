package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Sentence struct {
	Text string `json:"text"`
	Hate float32 `json:"hate"`
}

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/corollari/linusrants/master/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "HTTP GET", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Read", err)
		os.Exit(1)
	}

	var sentences []Sentence
	err = json.Unmarshal(body, &sentences)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Transforming", err)
		os.Exit(1)
	}

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	var rant = sentences[r.Intn(len(sentences))].Text
	fmt.Println(rant, "\n - Linus Torvalds")
}