package main

import (
	"time"

	"github.com/google/uuid"
)

type CreateReminderReq struct {
	MedicineName string      `json:"medicine_name"`
	Dosage       string      `json:"dosage"`
	RemindDate   []time.Time `json:"date"`
	Description  string      `json:"description"`
}

type ReminderResponse struct {
	Id           uuid.UUID      `json:"id"`
	MedicineName string         `json:"medicine_name"`
	Dosage       string         `json:"dosage"`
	RemindDate   []ReminderDate `json:"date"`
	Description  string         `json:"description"`
}

type UpdateReminderDateStatus struct {
	Id          uuid.UUID `json:"reminderDateId"`
	IsCompleted bool      `json:"isCompleted"`
}

type CreatePillCourseReq struct {
	Name   string `json:"name"`
	UserId string `json:"userId"`
}

type AddReminderDateToCourse struct {
	ReminderDateId uuid.UUID `json:"reminderDateId"`
	CourseId       uuid.UUID `json:"courseId"`
}
