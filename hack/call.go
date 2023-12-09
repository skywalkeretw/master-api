package main

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func CallFunction(service, function, mode string, data interface{}) interface{} {
// 	infoUrl := fmt.Sprintf("http://%s:%d", service, 8080)
// 	http.Get(infoUrl)

// 	switch mode {
// 	case "http":
// 		var bodyReader *bytes.Reader

// 		// Check if data is provided and serialize it to JSON
// 		if data != nil {
// 			jsonData, err := json.Marshal(data)
// 			if err != nil {
// 				fmt.Println("Error encoding JSON data:", err)
// 				return nil
// 			}
// 			bodyReader = bytes.NewReader(jsonData)
// 		}

// 		// Make the HTTP request
// 		request, err := http.NewRequest(strings.ToUpper(http.MethodPost), "url", bodyReader)
// 		if err != nil {
// 			fmt.Println("Error creating HTTP request:", err)
// 			return nil
// 		}

// 		// Set the Content-Type header for JSON data
// 		if bodyReader != nil {
// 			request.Header.Set("Content-Type", "application/json")
// 		}

// 		// Perform the HTTP request
// 		response, err := http.DefaultClient.Do(request)
// 		if err != nil {
// 			fmt.Println("Error making HTTP request:", err)
// 			return nil
// 		}
// 		defer response.Body.Close()

// 		// Read the response body
// 		responseBody, err := ioutil.ReadAll(response.Body)
// 		if err != nil {
// 			fmt.Println("Error reading response body:", err)
// 			return nil
// 		}

// 		// Print the response body
// 		fmt.Println("Response Body:", string(responseBody))
// 	case "rabbitmq":
// 		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 		if err != nil {
// 			return nil, err
// 		}

// 		channel, err := conn.Channel()
// 		if err != nil {
// 			return nil, err
// 		}

// 		queue, err := channel.QueueDeclare(
// 			"",    // name
// 			false, // durable
// 			false, // delete when unused
// 			true,  // exclusive
// 			false, // noWait
// 			nil,   // arguments
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		conn.Close()
// 		channel.Close()
// 		// return &RPCClient{
// 		// 	conn:    conn,
// 		// 	channel: channel,
// 		// 	queue:   queue,
// 		// }, nil

// 		corrID := randomString(32)

// 		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 		defer cancel()

// 		err := c.channel.PublishWithContext(ctx,
// 			"",          // exchange
// 			"rpc_queue", // routing key
// 			false,       // mandatory
// 			false,       // immediate
// 			amqp.Publishing{
// 				ContentType:   "text/plain",
// 				CorrelationId: corrID,
// 				ReplyTo:       c.queue.Name,
// 				Body:          []byte(strconv.Itoa(n)),
// 			})
// 		if err != nil {
// 			return 0, err
// 		}

// 		msgs, err := c.channel.Consume(
// 			c.queue.Name, // queue
// 			"",           // consumer
// 			true,         // auto-ack
// 			false,        // exclusive
// 			false,        // no-local
// 			false,        // no-wait
// 			nil,          // args
// 		)
// 		if err != nil {
// 			return 0, err
// 		}

// 		for d := range msgs {
// 			if corrID == d.CorrelationId {
// 				res, err := strconv.Atoi(string(d.Body))
// 				if err != nil {
// 					return 0, err
// 				}
// 				return res, nil
// 			}
// 		}
// 	default:
// 		fmt.Println("Not a valid day.")
// 	}
// }

// // RPCClient represents the RabbitMQ RPC client
// type RPCClient struct {
// 	conn    *amqp.Connection
// 	channel *amqp.Channel
// 	queue   amqp.Queue
// }

// // NewRPCClient initializes a new RPCClient
// func NewRPCClient() (*RPCClient, error) {

// }

// // Close closes the RPCClient connections
// func (c *RPCClient) Close() {

// }

// // Call performs the RPC call
// func (c *RPCClient) Call(n int) (int, error) {

// 	return 0, nil
// }
