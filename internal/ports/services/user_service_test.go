package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gobro.starter/internal/ports"
)

// Mocking the PublisherConfig interface
type MockPublisherConfig struct {
	mock.Mock
}

func (m *MockPublisherConfig) Publish(content any) {
	// Record the method call with arguments
	m.Called(content)
}

func (m *MockPublisherConfig) Close() {
	// Optionally mock this if needed
	m.Called()
}

func TestUserService_Greet(t *testing.T) {
	// Step 1: Create a mock PublisherConfig
	mockConfig := new(MockPublisherConfig)

	// Step 2: Set up the expectation (Publish should be called with the correct content)
	expectedContent := "Publishing Content from UserService"
	mockConfig.On("Publish", expectedContent).Return()

	// Step 3: Create a MessagePublisher instance with the mock adapter
	messagePublisher := ports.GetMessagingInstance(mockConfig)

	// Step 4: Inject the MessagePublisher into the UserService singleton
	userService := GetInstance(messagePublisher)

	// Step 5: Call the Greet method
	userService.Greet()

	// Step 6: Verify that the mock's Publish method was called with the expected content
	mockConfig.AssertCalled(t, "Publish", expectedContent)

	// Optionally, use testify's assertion library for better test output
	assert.True(t, mockConfig.AssertCalled(t, "Publish", expectedContent), "Publish was not called with the expected content")
}
