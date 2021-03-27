package externalServices

type IExternalService interface {
	Get()
	SetServiceName(sName string)
}
