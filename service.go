package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PillService struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *PillService {
	return &PillService{
		db: db,
	}
}

func (s *PillService) Check(c *gin.Context) {
	c.JSON(200, "running")
}

func (s *PillService) RegisterUser(c *gin.Context) {
	var requestBody string
	err := c.Bind(&requestBody)
	if err != nil {
		c.JSON(500, err)
	}

	user := User{
		id: requestBody,
	}

	_, err = s.db.Exec("insert into User (id, name) values ($1, $2)", user)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, nil)
}

func (s *PillService) CreateReminder(c *gin.Context) {
	var requestBody CreateReminderReq

	err := c.Bind(requestBody)
	if err != nil {
		c.JSON(500, err)
	}

	reminder := ReminderEntity{
		id:           uuid.New(),
		medicineName: requestBody.MedicineName,
		dosage:       requestBody.Dosage,
		description:  requestBody.Description,
	}

	var remindDates []ReminderDate
	for _, v := range requestBody.RemindDate {
		remindDate := ReminderDate{
			id:          uuid.New(),
			date:        v.UTC(),
			isCompleted: false,
			reminderId:  reminder.id,
		}
		remindDates = append(remindDates, remindDate)
	}

	_, err = s.db.Exec("insert into ReminderDates (id, date, status, reminder_id) values ($1, $2, $3, $4)", remindDates)
	if err != nil {
		c.JSON(500, err)
	}

	_, err = s.db.Exec("insert into Reminder (id, medicine_name, dosage, description) values ($1, $2, $3, $4)", reminder)

	if err != nil {
		c.JSON(500, err)
	}

	reminderResponse := ReminderResponse{
		Id:           reminder.id,
		MedicineName: reminder.medicineName,
		Dosage:       reminder.dosage,
		RemindDate:   remindDates,
		Description:  reminder.description,
	}

	c.JSON(200, &reminderResponse)
}

func (s *PillService) FindReminderById(c *gin.Context) {
	var id uuid.UUID

	err := c.Bind(id)
	if err != nil {
		c.JSON(500, err)
	}

	var reminder ReminderEntity
	err = s.db.Select(&reminder, "select * from Reminder where id = $1 ", id)
	if err != nil {
		c.JSON(500, err)
	}

	var reminderDates []ReminderDate
	err = s.db.Select(&reminderDates, "select * from ReminderDates where reminder_id = $1", id)
	if err != nil {
		c.JSON(500, err)
	}

	reminderResponse := ReminderResponse{
		Id:           reminder.id,
		MedicineName: reminder.medicineName,
		Dosage:       reminder.dosage,
		RemindDate:   reminderDates,
		Description:  reminder.description,
	}

	c.JSON(200, &reminderResponse)
}

func (s *PillService) FindAllRemindersByUserId(c *gin.Context) {
	userId, err := uuid.Parse(c.Query("userId"))
	if err != nil {
		c.JSON(500, err)
	}

	var reminders []ReminderEntity

	err = s.db.Select(&reminders, "select * from Reminder where user_id = $1", userId)
	if err != nil {
		c.JSON(500, err)
	}

	var remindersId []uuid.UUID
	for _, v := range reminders {
		remindersId = append(remindersId, v.id)
	}

	var reminderDates []ReminderDate
	err = s.db.Select(&reminderDates, "select * from ReminderDates where reminder_id = any($1)", remindersId)
	if err != nil {
		c.JSON(500, err)
	}

	var reminderResponse []ReminderResponse
	for _, v := range reminders {
		reminder := ReminderResponse{
			Id:           v.id,
			MedicineName: v.medicineName,
			Dosage:       v.dosage,
			Description:  v.description,
		}

		for _, k := range reminderDates {
			if k.reminderId == v.id {
				reminder.RemindDate = append(reminder.RemindDate, k)
			}
		}

		reminderResponse = append(reminderResponse, reminder)
	}

	c.JSON(200, &reminderResponse)
}

func (s *PillService) DeleteById(c *gin.Context) {
	var reminderId uuid.UUID

	err := c.Bind(reminderId)
	if err != nil {
		c.JSON(500, err)
	}

	_, err = s.db.Exec("delete from Reminder where id = $1", reminderId)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, "deleted")
}

func (s *PillService) UpdateStatus(c *gin.Context) {
	var req UpdateReminderDateStatus

	err := c.Bind(req)
	if err != nil {
		c.JSON(500, err)
	}

	_, err = s.db.Exec("update reminderdate set isCompleted = $1 where id = $2", req.IsCompleted, req.Id)
	if err != nil {
		c.JSON(500, err)
	}
}
