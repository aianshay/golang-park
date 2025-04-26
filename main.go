package main

import (
	"os"
	"context"
	"log"
	"bufio"
	"fmt"
	"github.com/aprendagolang/agent/brain"
)

func main() {
	agent, err := brain.NewAgent(context.Background())

	if err != nil {
		log.Fatal(err)
	}	 
	agent.Debug = true
	ctx := context.Background()

	for { 
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Eu: ")
		input, _ := reader.ReadString('\n')

		resp, err := agent.SendMessage(ctx, input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Assistente do zap: %s\n", resp.Text())
	}
}