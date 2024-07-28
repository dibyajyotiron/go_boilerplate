package discovery

type ServiceRegistry interface {
	Register(serviceID, serviceName, address string, port int) error
	Deregister(serviceID string) error
}
