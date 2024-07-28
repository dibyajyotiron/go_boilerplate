package discovery

import (
	"github.com/hashicorp/consul/api"
)

type ConsulServiceRegistry struct {
	client *api.Client
}

func NewConsulServiceRegistry(address string) (*ConsulServiceRegistry, error) {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ConsulServiceRegistry{client: client}, nil
}

func (c *ConsulServiceRegistry) Register(serviceID, serviceName, address string, port int) error {
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: address,
		Port:    port,
	}
	return c.client.Agent().ServiceRegister(registration)
}

func (c *ConsulServiceRegistry) Deregister(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}
