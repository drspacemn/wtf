package main

import (
	"os"
	"fmt"
	"flag"
	"net/url"
	"net/http"
	"encoding/json"
)

var KEY, HOST string
	
type Word struct {
	Word    string `json:"word"`
	Results []struct {
		Definition   string   `json:"definition"`
		PartOfSpeech string   `json:"partOfSpeech"`
		Synonyms     []string `json:"synonyms"`
		TypeOf       []string `json:"typeOf"`
		InCategory   []string `json:"inCategory"`
		HasTypes     []string `json:"hasTypes,omitempty"`
		Derivation   []string `json:"derivation,omitempty"`
		Examples     []string `json:"examples,omitempty"`
	} `json:"results"`
	Syllables struct {
		Count int      `json:"count"`
		List  []string `json:"list"`
	} `json:"syllables"`
	Pronunciation struct {
		All string `json:"all"`
	} `json:"pronunciation"`
	Frequency float64 `json:"frequency"`
}

func init() {
	KEY = os.Getenv("WORDS_KEY")
	HOST = os.Getenv("WORDS_HOST")
}

func main() {
	topicCmd := flag.NewFlagSet("topic", flag.ExitOnError)
	topicWord := topicCmd.String("is", "", "topic you want to explore(snake_case)")
	
	subtopicCmd := flag.NewFlagSet("subtopic", flag.ExitOnError)
	underTopic := subtopicCmd.String("under", "", "top level topic(snake_case)")
	subtopicWord := subtopicCmd.String("is", "", "subtopic to create examples(snake_case)")

	// graphCmd := flag.NewFlagSet("graph", flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s topic|subtopic|graph\n", os.Args[0])
		flag.PrintDefaults()
	}


	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "topic":
		topicCmd.Parse(os.Args[2:])
		fmt.Println("GOT TOPIC: ", *topicWord)
		w := &Word{Word: *topicWord}
		w.Full()
	
		for _, val := range w.Results {
			fmt.Println("ALL: ", val.InCategory)
			fmt.Println("ALL: ", val.HasTypes)
			fmt.Println("ALL: ", val.TypeOf)
		}

	case "subtopic":
		subtopicCmd.Parse(os.Args[2:])
		fmt.Println("GOT TOPLEVEL: ", *underTopic)
		fmt.Println("GOT SUBTOPIC: ", *subtopicWord)
		w := &Word{Word: *subtopicWord}
		w.Full()
	
		for _, val := range w.Results {
			fmt.Println("ALL: ", val.InCategory)
			fmt.Println("ALL: ", val.HasTypes)
			fmt.Println("ALL: ", val.TypeOf)
		}

	case "graph":
		fmt.Println("GOT GRAPH")
	
	default:
		flag.Usage()
		os.Exit(1)
	}


}

func (w *Word) Full() {
	url := "https://wordsapiv1.p.rapidapi.com/words/" + url.PathEscape(w.Word)
	fmt.Println("URL: ", url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "wordsapiv1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "bb6c04d1b6msh944a28816ddce50p10b2a1jsnfa8d4c4e82ca")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(w)
}