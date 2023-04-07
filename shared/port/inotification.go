package port

type INotificationFactoryStruct interface {
	IsValidParams(params map[string]interface{}) bool
	Manufacture(version string, params map[string]interface{}) (interface{}, error)
}

type INotification interface {
	MakeServices()
	RegisterFactoryStruct(topic, event string, factory INotificationFactoryStruct)
	RegisterFactoryStructEvents(topic string, events []string, factory INotificationFactoryStruct)
	Notify(topic, event, version string, params map[string]interface{}) error
	WithTransaction(transaction ITransaction) INotification
}
