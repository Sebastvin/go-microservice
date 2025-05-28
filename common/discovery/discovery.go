package discovery

import "context"

type Registry interface {
	Register(ctx context.Context, instanceID, serverName, hostPort string) error
	ReRegister(ctx context.Context, instanceID, serverName string) error
	Discover(ctx context.Context, serviceName string) ([]string, error)
	HealthCheck(ctx context.Context, instanceID, serviceName string) error
}
