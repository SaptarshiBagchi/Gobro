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
var messageSync sync.Once

func GetMessagingInstance(infra PublisherConfig) *MessagePublisher {
	//Using a different way of singleTon to execute this
	messageSync.Do(
		func() {
			messagePublisherInstance = &MessagePublisher{adapter: infra}
		})
	return messagePublisherInstance

}

func (p *MessagePublisher) Publish(content any) {
	//publish the content
	p.adapter.Publish(content)
}
