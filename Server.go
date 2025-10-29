package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// Message struct represents a chat message
type Message struct {
	Text string
}

// ChatRoom struct holds all messages
type ChatRoom struct {
	messages []Message
	mu       sync.Mutex
}

// SendMessage RPC method stores a message
func (c *ChatRoom) SendMessage(msg *Message, reply *string) error {
	if msg.Text == "" {
		return errors.New("empty message")
	}
	c.mu.Lock()
	c.messages = append(c.messages, *msg)
	c.mu.Unlock()
	*reply = "Message sent successfully"
	return nil
}

// GetHistory RPC method returns all chat history
func (c *ChatRoom) GetHistory(dummy *int, reply *[]Message) error {
	c.mu.Lock()
	*reply = make([]Message, len(c.messages))
	copy(*reply, c.messages)
	c.mu.Unlock()
	return nil
}

func main() {
	chat := new(ChatRoom)
	rpc.Register(chat)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Chat server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
