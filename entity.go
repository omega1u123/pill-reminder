package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id string `db:"id"`
}

type ReminderEntity struct {
	Id           uuid.UUID `db:"id"`
	MedicineName string    `db:"medicine_name"`
	Dosage       string    `db:"dosage"`
	Description  string    `db:"description"`
	UserId       string    `db:"user_id"`
}

type ReminderDate struct {
	Id           uuid.UUID `db:"id"`
	Date         time.Time `db:"date"`
	IsCompleted  bool      `db:"isCompleted"`
	ReminderId   uuid.UUID `db:"reminder_id"`
	PillCourseId uuid.UUID `db:"pill_course_id"`
}

type PillCourse struct {
	Id     uuid.UUID `db:"id"`
	Name   string    `db:"name"`
	UserId string    `db:"user_id"`
}
