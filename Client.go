package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Text string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message (or type 'exit'): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		msg := Message{Text: input}
		var reply string
		err = client.Call("ChatRoom.SendMessage", &msg, &reply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}
		fmt.Println(reply)

		// Fetch chat history
		var history []Message
		err = client.Call("ChatRoom.GetHistory", new(int), &history)
		if err != nil {
			fmt.Println("Error fetching history:", err)
			continue
		}

		fmt.Println("Chat history:")
		for i, m := range history {
			fmt.Printf("%d: %s\n", i+1, m.Text)
		}
	}
}
