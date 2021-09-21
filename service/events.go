package service

import (
	"regexp"

	"github.com/d-kuznetsov/calendar-backend/dto"
)

func (service *Service) CreateEvent(eventData dto.Event) (string, error) {
	if !checkDate(eventData.Date) ||
		!checkTime(eventData.StartTime) ||
		!checkTime(eventData.EndTime) ||
		eventData.StartTime > eventData.EndTime ||
		eventData.Content == "" {
		return "", ErrIncorrectData
	}
	return service.repository.CreateEvent(eventData)
}

func (service *Service) GetEvents(params dto.PeriodParams) ([]dto.Event, error) {
	if !checkDate(params.PeriodStart) ||
		!checkDate(params.PeriodEnd) ||
		params.PeriodStart > params.PeriodEnd ||
		params.UserId == "" {
		return make([]dto.Event, 0), ErrIncorrectData
	}
	return service.repository.GetEvents(params)
}

func (service *Service) UpdateEvent(eventData dto.Event) error {
	if eventData.Id == "" ||
		!checkDate(eventData.Date) ||
		!checkTime(eventData.StartTime) ||
		!checkTime(eventData.EndTime) ||
		eventData.StartTime > eventData.EndTime ||
		eventData.Content == "" {
		return ErrIncorrectData
	}
	return service.repository.UpdateEvent(eventData)
}

func (service *Service) DeleteEvent(id string) error {
	if id == "" {
		return ErrIncorrectData
	}
	return service.repository.DeleteEvent(id)
}

var dateRegexp = regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`)

func checkDate(date string) bool {
	res := dateRegexp.MatchString(date)
	return res
}

var timeRegexp = regexp.MustCompile(`^([01]\d|2[0-3]):([0-5]\d)$`)

func checkTime(time string) bool {
	res := timeRegexp.MatchString(time)
	return res
}
