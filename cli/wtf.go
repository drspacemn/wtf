package main

import (
	"os"
	"fmt"
	"flag"
	"bytes"
	"net/http"
	"strings"
	"encoding/json"
)

var OPENAI_API_KEY string

type Prompt struct {
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
	Engine string `json:"-"`
}

type Completion struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
}

func init() {
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
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
		c, err := TopicOverview(*topicWord)
		WriteLine(*topicWord, fmt.Sprintf("## WTF is %s?\n", strings.Title(strings.Replace(*topicWord, "_", " ", -1))))
		line := strings.Replace(c.Choices[0].Text, "\n\n", "\n", -1)
		WriteLine(*topicWord, line+"\n\n")

		WriteLine(*topicWord, "#### Sources:\n\n")
		WriteLine(*topicWord, "#### Examples:\n\n")
		
		WriteLine(*topicWord, "<hr>\n")
		c, err = TopicOuterRelated(*topicWord)
		fmt.Println("LINE:", c.Choices[0].Text)
		line = strings.Replace(c.Choices[0].Text, "-", ",", -1)
		line = strings.Replace(line, "\n", "", -1)
		WriteLine(*topicWord, "outerTags:")
		WriteLine(*topicWord, fmt.Sprintf("[%s]\n", strings.Replace(strings.ToLower(line), ",", "", 1)))
		
		c, err = TopicInnerRelated(*topicWord)
		fmt.Println("LINE:", c.Choices[0].Text)
		line = strings.Replace(c.Choices[0].Text, "-", ",", -1)
		line = strings.Replace(line, "\n", "", -1)
		WriteLine(*topicWord, "innerTags:")
		WriteLine(*topicWord, fmt.Sprintf("[%s]\n", strings.Replace(strings.ToLower(line), ",", "", 1)))
		fmt.Println("ERR: ", err)

	case "subtopic":
		subtopicCmd.Parse(os.Args[2:])
		fmt.Println("GOT TOPLEVEL: ", *underTopic)
		fmt.Println("GOT SUBTOPIC: ", *subtopicWord)
		// w := &Word{Word: *subtopicWord}

	case "graph":
		fmt.Println("GOT GRAPH")
	
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func TopicOverview(topic string) (Completion, error) {
	p := Prompt{
		Prompt: fmt.Sprintf("What are computer %ss?", topic),
		Temperature: 0.5,
		TopP: 0.4,
		MaxTokens: 100,
		FrequencyPenalty: 0,
		PresencePenalty: 0,
		Engine: "https://api.openai.com/v1/engines/text-babbage-001/completions",
	}
	return p.Complete()
}

func TopicOuterRelated(topic string) (Completion, error) {
	p := Prompt{
		Prompt: fmt.Sprintf("List the larger concepts that include \"Computer %s\" with dashes.\nList:\n-", topic),
		Temperature: 0.1,
		TopP: 0.9,
		MaxTokens: 100,
		FrequencyPenalty: 0,
		PresencePenalty: 0,
		Engine: "https://api.openai.com/v1/engines/text-davinci-001/completions",
	}
	return p.Complete()
}

func TopicInnerRelated(topic string) (Completion, error) {
	p := Prompt{
		Prompt: fmt.Sprintf("List the concepts that make up \"Computer %s\". \nList:\n-", topic),
		Temperature: 0,
		TopP: 1,
		MaxTokens: 100,
		FrequencyPenalty: 0,
		PresencePenalty: 0,
		Engine: "https://api.openai.com/v1/engines/text-davinci-001/completions",
	}
	return p.Complete()
}

func (p Prompt) Complete() (comp Completion, err error) {
	raw, err := json.Marshal(p)
	if err != nil {
		return comp, err
	}
	fmt.Println("RAW: ", string(raw))
	req, err := http.NewRequest("POST", p.Engine, bytes.NewReader(raw))
	if err != nil {
		return comp, err
	}

	req.Header.Add("Authorization", "Bearer "+OPENAI_API_KEY)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return comp, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&comp)
	return comp, err
}

func WriteLine(dir, line string) {
	filePath := fmt.Sprintf("is_%s/README.md", dir)
	var _, err = os.Stat(filePath)
    // create file if not exists
    if os.IsNotExist(err) {
		err = os.Mkdir(fmt.Sprintf("is_%s", dir), os.ModePerm)
        var file, err = os.Create(filePath)
        if err != nil {
			panic(err.Error())
        }
        defer file.Close()
		if _, err := file.WriteString(line); err != nil {
			panic(err.Error())
		}
    } else {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()
		if _, err := f.WriteString(line); err != nil {
			err.Error()
		}
	}	
}