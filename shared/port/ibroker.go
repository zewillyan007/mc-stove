package port

type IBrokerDeliveryChannel interface {
	IsValidChannel() bool
}

type IBrokerEventHandler func(IBrokerDeliveryChannel)

type IBroker interface {
	Close()
	CreateQueue(name string, options ...map[string]interface{}) error
	CreateExchange(name string, kind string, options ...map[string]interface{}) error
	Consume(queue string, consumer string, options ...map[string]interface{}) (IBrokerDeliveryChannel, error)
	Publish(exchange string, key string, typeMessage string, message []byte, options ...map[string]interface{}) error
}
