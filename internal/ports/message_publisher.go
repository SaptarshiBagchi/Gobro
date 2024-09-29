package ports

import "sync"

type PublisherConfig interface {
	Publish(content any)
	Close()
}

type MessagePublisher struct {
	adapter PublisherConfig
}

var messagePublisherInstance *MessagePublisher
var messagePublisherMutex = &sync.Mutex{}

func GetMessagingInstance(infra PublisherConfig) *MessagePublisher {
	messagePublisherMutex.Lock()
	defer messagePublisherMutex.Unlock()
	if messagePublisherInstance == nil {
		messagePublisherInstance = &MessagePublisher{adapter: infra}
	}

	return messagePublisherInstance

}

func (p *MessagePublisher) Publish(content any) {
	//publish the content
	p.adapter.Publish(content)
}
