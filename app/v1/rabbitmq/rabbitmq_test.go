package rabbitmq

// import (
// 	"testing"

// 	"github.com/rabbitmq/amqp091-go/mocks"
// 	"gotest.tools/assert"
// )

// func TestGetRabbitMQQueues(t *testing.T) {
// 	// Create a mock connection
// 	mockConn := mocks.NewConnection()
// 	defer mockConn.Close()

// 	// Create a mock channel
// 	mockChannel := mockConn.NewChannel()
// 	defer mockChannel.Close()

// 	// Define the list of queues you want to mock
// 	mockQueue1 := mocks.Queue{
// 		Name:      "queue1",
// 		Messages:  10,
// 		Consumers: 5,
// 	}
// 	mockQueue2 := mocks.Queue{
// 		Name:      "queue2",
// 		Messages:  20,
// 		Consumers: 3,
// 	}

// 	// Set up the mock channel's behavior
// 	mockChannel.On("ListQueues").Return([]mocks.Queue{mockQueue1, mockQueue2}, nil)

// 	// Call the function that retrieves queues
// 	queues, err := getRabbitMQQueues(mockChannel)
// 	if err != nil {
// 		t.Fatalf("Expected no error, but got: %v", err)
// 	}

// 	// Assert the expected queue information
// 	assert.Equal(t, 2, len(queues))
// 	assert.Equal(t, "queue1", queues[0].Name)
// 	assert.Equal(t, 10, queues[0].Messages)
// 	assert.Equal(t, 5, queues[0].Consumers)
// 	assert.Equal(t, "queue2", queues[1].Name)
// 	assert.Equal(t, 20, queues[1].Messages)
// 	assert.Equal(t, 3, queues[1].Consumers)

// 	// Assert that the expected method was called on the mock channel
// 	mockChannel.AssertExpectations(t)
// }
