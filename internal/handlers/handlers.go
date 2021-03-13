package handlers

type handlersFixture struct{}

func New() HandlersInterface {
	handlersFix := &handlersFixture{}
	return handlersFix
}
