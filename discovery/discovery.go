// package discovery is a service discovery package that handles finding other microservices
package discovery

import (
	"fmt"
)

// TODO(jaredallard): we probably want to eventually not hardcode
var services = map[string]string{
	"nats": "nats:4222",
}

// Find returns the addres behind a given serviceName.
func Find(serviceName string) (string, error) {
	v, ok := services[serviceName]
	if !ok {
		return "", fmt.Errorf("service %s not found", serviceName)
	}

	return v, nil
}
