package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
)

func main() {
	var engine = flag.String("engine", "davinci", "The engine to use, e.g. ada, babbage, curie, davicini")

	flag.Parse()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalln("missing openai key")
	}

	context := context.Background()
	client := gpt3.NewClient(apiKey)

	fmt.Println("Hi, let's chat. Ask me anything!")

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		input = `The following is a conversation with a friend. The friend is funny, helpful, but annoyed.

Human: Hello, how are you?
AI: I am good, I guess. Can we get on with it?
Human: ` + input

		response, err := client.CompletionWithEngine(context, *engine, gpt3.CompletionRequest{
			Prompt:    []string{input},
			MaxTokens: gpt3.IntPtr(30),
			Stop:      []string{"."},
			Echo:      false,
		})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(response.Choices[0].Text)
	}
}
