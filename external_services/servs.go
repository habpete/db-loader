package externalServices

import (
	"db-loader/external_services/profile"
	"db-loader/external_services/review"
	"db-loader/external_services/score"
	"sync"
)

const externalServicesCount = 3

type extServices struct{}

func InitExtServices() *extServices {}

func (this *extServices) Serf() {
	actions := make(IExternalService, externalServicesCount)

	actions[0] = profile.ProfileService
	actions[1] = review.ReviewService
	actions[2] = score.ScoreService

	var wg sync.WaitGroup
	wg.Add(externalServicesCount)

	for i := 0; i < externalServicesCount; i++ {
		go func () {
			actions[i].Get()
			wg.Done()
		}
	}

	wg.Wait()
}
