package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id string `db:"id"`
}

type ReminderEntity struct {
	id           uuid.UUID `db:"id"`
	medicineName string    `db:"medicine_name"`
	dosage       string    `db:"dosage"`
	description  string    `db:"description"`
	userId       uuid.UUID `db:"user_id"`
}

type ReminderDate struct {
	id          uuid.UUID `db:"id"`
	date        time.Time `db:"date"`
	isCompleted bool      `db:"isCompleted"`
	reminderId  uuid.UUID `db:"reminder_id"`
}
