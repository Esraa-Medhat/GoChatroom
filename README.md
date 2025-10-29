# Go Chatroom (Client/Server)

A simple chatroom implemented in Go using RPC.  
The client can send messages to the server and fetch the full chat history.

---

## **Files**

- `server.go` – The chat server that stores all messages.
- `client.go` – The chat client that sends messages and fetches history.

---

## **How to Run**

### 1. Start the Server
Open terminal in the project folder and run:

```bash
go run server.go

You should see:
Chat server running on port 1234...

2. Start the Client
Open another terminal in the project folder and run:
go run client.go
You will see:
Enter message (or type 'exit'):

Type a message and press Enter to send it.

The client will show all chat history after sending each message.

To exit, type exit or press Ctrl+C.

Project Features

Server stores all messages in a long list.

Client can fetch the full chat history using RPC.

Thread-safe server using sync.Mutex.

Handles messages with spaces using bufio.NewReader.

Video Demo
Watch the chatroom in action:[Go Chatroom Demo Video](https://drive.google.com/file/d/1fNd88FCLWDcCvk0VR4auZ4WxKdeH_pPz/view?usp=sharing)



