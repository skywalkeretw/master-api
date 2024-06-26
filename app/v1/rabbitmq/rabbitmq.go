package rabbitmq

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/skywalkeretw/master-api/app/utils"
	"github.com/skywalkeretw/master-api/app/v1/kubernetes"
)

// import (
// 	"fmt"
// 	"log"

// 	amqp "github.com/rabbitmq/amqp091-go"
// )

// func init() {

// }

// func getRabbitMQQueues() {
// 	// RabbitMQ connection parameters
// 	rabbitMQURL := "amqp://guest:guest@localhost:5672/"

// 	// Establish a connection to RabbitMQ
// 	conn, err := amqp.Dial(rabbitMQURL)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
// 	}
// 	defer conn.Close()

// 	// Create a channel
// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatalf("Failed to open a channel: %v", err)
// 	}
// 	defer ch.Close()

// 	// List all queues
// 	queues, err := ch.ListQueues()
// 	if err != nil {
// 		log.Fatalf("Failed to list queues: %v", err)
// 	}

//		fmt.Println("List of RabbitMQ queues:")
//		for _, queue := range queues {
//			fmt.Printf("Queue Name: %s\n", queue.Name)
//			fmt.Printf("Queue Messages: %d\n", queue.Messages)
//			fmt.Printf("Queue Consumers: %d\n", queue.Consumers)
//			fmt.Println("---------------------------")
//		}
//	}
var conn *amqp.Connection

func init() {
	username := "default_user_8wLijMkNdrNBoRADgib"
	password := "GfsPBslwHprNY7jx_jJFJRZBDZBf2UCn"
	host := "hello-world"
	port := 5672
	dialUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port)
	var err error
	conn, err = amqp.Dial(dialUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type AutoGenerated []struct {
	Arguments struct {
	} `json:"arguments"`
	AutoDelete         bool `json:"auto_delete"`
	BackingQueueStatus struct {
		AvgAckEgressRate  float64 `json:"avg_ack_egress_rate"`
		AvgAckIngressRate float64 `json:"avg_ack_ingress_rate"`
		AvgEgressRate     float64 `json:"avg_egress_rate"`
		AvgIngressRate    float64 `json:"avg_ingress_rate"`
		Delta             []any   `json:"delta"`
		Len               int     `json:"len"`
		Mode              string  `json:"mode"`
		NextDeliverSeqID  int     `json:"next_deliver_seq_id"`
		NextSeqID         int     `json:"next_seq_id"`
		NumPendingAcks    int     `json:"num_pending_acks"`
		NumUnconfirmed    int     `json:"num_unconfirmed"`
		Q1                int     `json:"q1"`
		Q2                int     `json:"q2"`
		Q3                int     `json:"q3"`
		Q4                int     `json:"q4"`
		QsBufferSize      int     `json:"qs_buffer_size"`
		TargetRAMCount    string  `json:"target_ram_count"`
		Version           int     `json:"version"`
	} `json:"backing_queue_status"`
	ConsumerCapacity          int  `json:"consumer_capacity"`
	ConsumerUtilisation       int  `json:"consumer_utilisation"`
	Consumers                 int  `json:"consumers"`
	Durable                   bool `json:"durable"`
	EffectivePolicyDefinition struct {
	} `json:"effective_policy_definition"`
	Exclusive            bool `json:"exclusive"`
	ExclusiveConsumerTag any  `json:"exclusive_consumer_tag"`
	GarbageCollection    struct {
		FullsweepAfter  int `json:"fullsweep_after"`
		MaxHeapSize     int `json:"max_heap_size"`
		MinBinVheapSize int `json:"min_bin_vheap_size"`
		MinHeapSize     int `json:"min_heap_size"`
		MinorGcs        int `json:"minor_gcs"`
	} `json:"garbage_collection"`
	HeadMessageTimestamp       any       `json:"head_message_timestamp"`
	IdleSince                  time.Time `json:"idle_since"`
	Memory                     int       `json:"memory"`
	MessageBytes               int       `json:"message_bytes"`
	MessageBytesPagedOut       int       `json:"message_bytes_paged_out"`
	MessageBytesPersistent     int       `json:"message_bytes_persistent"`
	MessageBytesRAM            int       `json:"message_bytes_ram"`
	MessageBytesReady          int       `json:"message_bytes_ready"`
	MessageBytesUnacknowledged int       `json:"message_bytes_unacknowledged"`
	Messages                   int       `json:"messages"`
	MessagesDetails            struct {
		Rate float64 `json:"rate"`
	} `json:"messages_details"`
	MessagesPagedOut     int `json:"messages_paged_out"`
	MessagesPersistent   int `json:"messages_persistent"`
	MessagesRAM          int `json:"messages_ram"`
	MessagesReady        int `json:"messages_ready"`
	MessagesReadyDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_ready_details"`
	MessagesReadyRAM              int `json:"messages_ready_ram"`
	MessagesUnacknowledged        int `json:"messages_unacknowledged"`
	MessagesUnacknowledgedDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_unacknowledged_details"`
	MessagesUnacknowledgedRAM int    `json:"messages_unacknowledged_ram"`
	Name                      string `json:"name"`
	Node                      string `json:"node"`
	OperatorPolicy            any    `json:"operator_policy"`
	Policy                    any    `json:"policy"`
	RecoverableSlaves         any    `json:"recoverable_slaves"`
	Reductions                int    `json:"reductions"`
	ReductionsDetails         struct {
		Rate float64 `json:"rate"`
	} `json:"reductions_details"`
	SingleActiveConsumerTag any    `json:"single_active_consumer_tag"`
	State                   string `json:"state"`
	Type                    string `json:"type"`
	Vhost                   string `json:"vhost"`
}

func ListQueues() {

	// Replace these variables with your username, password, and URL
	username := "default_user_8wLijMkNdrNBoRADgib"
	password := "GfsPBslwHprNY7jx_jJFJRZBDZBf2UCn"
	url := "http://hello-world:15672/api/queues"

	// Create the HTTP client with basic authentication
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Failed to create HTTP request: %v\n", err)
		return
	}

	// Set basic authentication header
	auth := username + ":" + password
	b64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+b64Auth)

	// Make the HTTP GET request
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("HTTP GET request failed: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)
		return
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return
	}

	// Print the response body as a string
	fmt.Println("Response:")
	fmt.Println(string(body))
}

func CreateQueue(queueName string) error {
	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare a queue
	_, err = ch.QueueDeclare(
		queueName, // Queue name
		false,     // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return err
	}

	fmt.Printf("Queue '%s' created successfully.\n", queueName)
	return nil
}

// Handle Sending of Message to builder and recive status that the container has been built

type RabbitMQDial struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type ImageData struct {
	ImageName string `json:"imagename"`
}

func (r RabbitMQDial) getUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", r.UserName, r.Password, r.Host, r.Port)
}

type FunctionBuildDeployData struct {
	Name        string `json:"name" binding:"required"`
	Language    string `json:"language" binding:"required"`
	Description string `json:"description"`
	SourceCode  string `json:"sourcecode" binding:"required"`
	// InputParameters string                 `json:"inputparameters" binding:"required"`
	// ReturnValue     string                 `json:"returnvalue" binding:"required"`
	// FunctionModes function.FunctionModes `json:"functionmodes" binding:"required"`
	FuncInput     string                   `json:"fucinput" binding:"required"`
	OpenAPIJSON   string                   `json:"openapijson"`
	AsyncAPIJSON  string                   `json:"asyncapijson"`
	FunctionModes kubernetes.FunctionModes `json:"functionmodes" binding:"required"`
}

func RPCclient(data FunctionBuildDeployData) {
	log.Println("Start Creation Process for", data.Name)
	rDial := RabbitMQDial{
		UserName: utils.GetEnvSting("RABBITMQ_USERNAME", "guest"),
		Password: utils.GetEnvSting("RABBITMQ_PASSWORD", "guest"),
		Host:     utils.GetEnvSting("RABBITMQ_HOST", "localhost"),
		Port:     utils.GetEnvInt("RABBITMQ_PORT", 5672),
	}
	conn, err := amqp.Dial(rDial.getUrl())
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	sendQueue := "imagebuilder"
	recieveQueue := "builtimages"

	// Declare or use existing "imageBuilder" queue
	_, err = ch.QueueDeclare(
		sendQueue, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare a queue", err)
	}

	_, err = ch.QueueDeclare(
		recieveQueue, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare a queue", err)
	}

	msgs, err := ch.Consume(
		recieveQueue, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to register a consumer", err)
	}

	corrId := fmt.Sprintf("build-id-%s", uuid.New())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	log.Println("send message")

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	err = ch.PublishWithContext(ctx,
		"",        // exchange
		sendQueue, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       recieveQueue,
			Body:          jsonData,
		})
	if err != nil {
		log.Panicf("%s: %s", "Failed to publish a message", err)
	}

	for d := range msgs {
		fmt.Println("recived: ", d.CorrelationId)
		if corrId == d.CorrelationId {
			var imageData ImageData
			err := json.Unmarshal(d.Body, &imageData)
			if err != nil {
				log.Panicf("%s: %s", "Failed To unmarshal", err)
			}
			fmt.Println("recived data:")
			fmt.Println(imageData)
			tagString := fmt.Sprintf("runtime=%s:version=v1.0.0", data.Language)
			err = kubernetes.CreateKubernetesDeployment(data.Name, "default", imageData.ImageName, data.Description, tagString, data.FunctionModes, 1)

			if err != nil {
				fmt.Println("failed to deploy", err.Error())
			} else {
				fmt.Println("Function has been deployed")
			}
			// kubernetes.CreateKubernetesDeployment("name", "default", 1, v1.PodTemplateSpec{
			// 	ObjectMeta: v1.PodTemplateSpec{},
			// })
			break
		}
	}
}
