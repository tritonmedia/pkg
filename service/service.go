package service

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Run(ctx context.Context, log logrus.FieldLogger) error
}

type Runner struct {
	services []Service
}

// NewServiceRunner creates a new service runner that launches a service
// in a goroutine and handles termination of other services when one
// fails to launch.
func NewServiceRunner(ctx context.Context, services []Service) *Runner {
	return &Runner{
		services: services,
	}
}

// Run starts the service runner
func (r *Runner) Run(ctx context.Context, log *logrus.Entry) error {
	wg := sync.WaitGroup{}
	errs := make([]error, 0)

	wg.Add(len(r.services))
	for _, s := range r.services {
		go func(s Service) {
			serviceName := reflect.TypeOf(s).String()

			err := s.Run(ctx, logrus.WithFields(logrus.Fields{"service": serviceName}))
			if err != nil {
				errs = append(errs, errors.Wrapf(err, "service '%s' failed", serviceName))
			}

			wg.Done()
		}(s)
	}

	// wait for the workers to finish
	// if our context is cancelled they are notified as well.
	// and _should_ return.
	wg.Wait()

	if len(errs) != 0 {
		return fmt.Errorf("failed to run, got the following errors: %v", errs)
	}

	return nil
}
