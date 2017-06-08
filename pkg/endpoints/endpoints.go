package endpoints

import (
	"context"

	"github.com/notque/persister/pkg/service"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.

type Endpoints struct {
	FooEndpoint endpoint.Endpoint
}
type FooRequest struct {
	S string
}
type FooResponse struct {
	Rs  string
	Err error
}

func New(svc service.PersisterService) (ep Endpoints) {
	ep.FooEndpoint = MakeFooEndpoint(svc)
	return ep
}

// MakeFooEndpoint returns an endpoint that invokes Foo on the service.
// Primarily useful in a server.
func MakeFooEndpoint(svc service.PersisterService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		rs, err := svc.Foo(ctx, req.S)
		return FooResponse{Rs: rs, Err: err}, nil
	}
}
