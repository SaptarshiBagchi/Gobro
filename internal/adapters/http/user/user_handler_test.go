package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gobro.starter/internal/ports"
	"gobro.starter/internal/ports/services"
)

// MockMessagePublisher is a mock of the MessagePublisher (dependency inside UserService)
type MockMessagePublisher struct {
	mock.Mock
}

func (m *MockMessagePublisher) Publish(content any) {
	m.Called(content) // Record that Publish was called with certain content
}

func (m *MockMessagePublisher) Close() {
	// Implement if necessary
}

func TestUserControllerHandler_Greet_RealMethod(t *testing.T) {
	// Step 1: Create a mock for MessagePublisher (dependency of UserService)
	mockPublisher := new(MockMessagePublisher)

	// Step 2: Set up expectation that Publish will be called with specific content
	mockPublisher.On("Publish", "Publishing Content from UserService").Return()

	messagePublisher := ports.GetMessagingInstance(mockPublisher)
	// Step 3: Create a real instance of UserService, injecting the mock publisher
	userService := services.GetInstance(messagePublisher)

	// Step 4: Inject the real UserService into the UserControllerHandler
	controller := NewUserController(userService)

	// Step 5: Create a new HTTP request to trigger the handler
	req, err := http.NewRequest("GET", "/greet", nil)
	assert.NoError(t, err)

	// Step 6: Create an HTTP response recorder
	rr := httptest.NewRecorder()

	// Step 7: Serve the request using the Greet handler
	handler := http.HandlerFunc(controller.Greet)
	handler.ServeHTTP(rr, req)

	// Step 8: Assert that the mock Publish method was called
	mockPublisher.AssertCalled(t, "Publish", "Publishing Content from UserService")

	// Optionally, assert the status code (this depends on whether your handler sets a status)
	assert.Equal(t, http.StatusOK, rr.Code)
}
