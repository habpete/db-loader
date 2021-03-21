package externalServices

type extServices struct{}

func InitExtServices() *extServices {}

func (this *extServices) Serf() {
	actions := make(IExternalService, 3)

	actions[0] = profile
	actions[1] = review
	actions[2] = score
}
