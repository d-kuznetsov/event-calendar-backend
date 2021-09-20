package service

import (
	"github.com/d-kuznetsov/calendar-backend/dto"
)

func (service *Service) CreateEvent(eventData dto.Event) (string, error) {
	return service.repository.CreateEvent(eventData)
}

func (service *Service) GetEvents(params dto.PeriodParams) ([]dto.Event, error) {
	return service.repository.GetEvents(params)
}

func (service *Service) UpdateEvent(eventData dto.Event) error {
	return service.repository.UpdateEvent(eventData)
}

func (service *Service) DeleteEvent(id string) error {
	return service.repository.DeleteEvent(id)
}
