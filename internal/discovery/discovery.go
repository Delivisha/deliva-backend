package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrNotFound = errors.New("no service addresses found")
)

type Registry interface {
	Register(ctx context.Context, instanceID, serviceName, hostPort string) error
	Deregister(ctx context.Context, instanceID, serviceName string) error
	ServiceAddresses(ctx context.Context, serviceID string) ([]string, error)
	ReportHealthState(instanceID, serviceName string) error
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
