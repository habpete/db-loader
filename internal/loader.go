package internal

import "db-loader/internal/handlers"

type internalFixture struct {
	handlers handlers.HandlersInterface
}

func New() *internalFixture {
	intFix := &internalFixture{}
	intFix.handlers = handlers.New()
	return intFix
}
