package service

import "context"

// Implement yor service methods methods.
// e.x: Foo(ctx context.Context,s string)(rs string, err error)
type PersisterService interface {
	Foo(ctx context.Context, s string) (rs string, err error)
}

type stubPersisterService struct{}

// Get a new instance of the service.
// If you want to add service middleware this is the place to put them.
func New() (s *stubPersisterService) {
	s = &stubPersisterService{}
	return s
}

// Implement the business logic of Foo
func (pe *stubPersisterService) Foo(ctx context.Context, s string) (rs string, err error) {
	return rs, err
}
