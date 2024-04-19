package serviceprovider

import (
	notifierservice "route256.ozon.ru/project/notifier/internal/service/notifier_service"
)

type services struct {
	notifierService notifierservice.Service
}

func (sp *ServiceProvider) GetNotifierService() notifierservice.Service {
	if sp.services.notifierService == nil {
		sp.services.notifierService = notifierservice.NewService()
	}
	return sp.services.notifierService
}
