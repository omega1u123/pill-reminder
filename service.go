package main

import (
	"fmt"

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

	err := c.Bind(&requestBody)
	if err != nil {
		c.JSON(500, err)
	}

	reminder := ReminderEntity{
		Id:           uuid.New(),
		MedicineName: requestBody.MedicineName,
		Dosage:       requestBody.Dosage,
		Description:  requestBody.Description,
	}

	var remindDates []ReminderDate
	for _, v := range requestBody.RemindDate {
		remindDate := ReminderDate{
			Id:          uuid.New(),
			Date:        v.UTC(),
			IsCompleted: false,
			ReminderId:  reminder.Id,
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
		Id:           reminder.Id,
		MedicineName: reminder.MedicineName,
		Dosage:       reminder.Dosage,
		RemindDate:   remindDates,
		Description:  reminder.Description,
	}

	c.JSON(200, &reminderResponse)
}

func (s *PillService) FindReminderById(c *gin.Context) {
	var id uuid.UUID

	err := c.Bind(&id)
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
		Id:           reminder.Id,
		MedicineName: reminder.MedicineName,
		Dosage:       reminder.Dosage,
		RemindDate:   reminderDates,
		Description:  reminder.Description,
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
		remindersId = append(remindersId, v.Id)
	}

	var reminderDates []ReminderDate
	err = s.db.Select(&reminderDates, "select * from ReminderDates where reminder_id = any($1)", remindersId)
	if err != nil {
		c.JSON(500, err)
	}

	var reminderResponse []ReminderResponse
	for _, v := range reminders {
		reminder := ReminderResponse{
			Id:           v.Id,
			MedicineName: v.MedicineName,
			Dosage:       v.Dosage,
			Description:  v.Description,
		}

		for _, k := range reminderDates {
			if k.ReminderId == v.Id {
				reminder.RemindDate = append(reminder.RemindDate, k)
			}
		}

		reminderResponse = append(reminderResponse, reminder)
	}

	c.JSON(200, &reminderResponse)
}

func (s *PillService) DeleteById(c *gin.Context) {
	var reminderId uuid.UUID

	err := c.Bind(&reminderId)
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

	err := c.Bind(&req)
	if err != nil {
		c.JSON(500, err)
	}

	_, err = s.db.Exec("update reminderdate set isCompleted = $1 where id = $2", req.IsCompleted, req.Id)
	if err != nil {
		c.JSON(500, err)
	}
}

func (s *PillService) CreatePillCourse(c *gin.Context) {
	var req CreatePillCourseReq
	err := c.Bind(&req)
	if err != nil {
		c.JSON(500, err)
	}

	course := PillCourse{
		Id:   uuid.New(),
		Name: req.Name,
	}

	query := `insert into pillcourse (id, name, user_id) values ($1, $2, $3)`
	_, err = s.db.Exec(query, course.Id, course.Name, req.UserId)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, course)
}

func (s *PillService) AddReminderDateToCourse(c *gin.Context) {
	var req AddReminderDateToCourse
	err := c.Bind(&req)
	if err != nil {
		c.JSON(500, err)
	}

	var isExists bool
	query := `select exists (select 1 from pillcourse where id = $1`
	err = s.db.Get(&isExists, query, req.CourseId)
	if !isExists {
		c.JSON(400, fmt.Sprintf("course with id = %d not found", req.CourseId))
	}
	if err != nil {
		c.JSON(500, err)
	}

	query = `select exists (select 1 from reminderdate where id = $1`
	err = s.db.Get(&isExists, query, req.ReminderDateId)
	if !isExists {
		c.JSON(400, fmt.Sprintf("reminderDate with id = %d not found", req.ReminderDateId))
	}
	if err != nil {
		c.JSON(500, err)
	}

	query = `insert into pillcourse(reminder_date_id) values ($1)`
	_, err = s.db.Exec(query, req.ReminderDateId)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, nil)
}

func (s *PillService) GetCoursesByUserId(c *gin.Context) {
	var userId string
	err := c.Bind(&userId)
	if err != nil {
		c.JSON(500, err)
	}

	var courseList []PillCourse
	query := `select * from pillcourse where user_id = $1`
	err = s.db.Select(&courseList, query, userId)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, courseList)
}
