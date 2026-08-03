package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	r := gin.Default()

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "reminder_user"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "reminder_password"
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "medicine-reminder-db"
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		panic(err)
	}

	createTables(db)
	service := NewService(db)

	r.GET("check", func(c *gin.Context) {
		c.JSON(200, "server is running")
	})

	r.GET("addCourseTable", func(context *gin.Context) {
		err := createPillCourseTable(db)
		if err != nil {
			context.JSON(500, err)
		}
		context.JSON(200, nil)
	})

	reminderGroup := r.Group("api/reminder")
	reminderGroup.POST("", service.CreateReminder)
	reminderGroup.GET("{:id}", service.FindReminderById)
	reminderGroup.GET("findByUserId", service.FindAllRemindersByUserId)
	reminderGroup.DELETE("{:id}", service.DeleteById)
	reminderGroup.PUT("{:reminderDateId}", service.UpdateStatus)

	userGroup := r.Group("api/user")
	userGroup.POST("register", service.RegisterUser)

	pillCourseGroup := r.Group("api/course")
	pillCourseGroup.GET("", service.GetCoursesByUserId)
	pillCourseGroup.POST("", service.CreatePillCourse)
	pillCourseGroup.PUT("", service.AddReminderDateToCourse)

	// Bind to all interfaces (0.0.0.0) for Docker compatibility
	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

func createPillCourseTable(db *sqlx.DB) error {
	query := `create table if not exists pillcourse(
		id uuid primary key,
		name varchar(20) not null,
		user_id string not null
		);
		
		alter table reminderDate add column pill_course_id uuid;
		`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createTables(db *sqlx.DB) {

	var exists bool
	checkTablesQuery := `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users')`

	err := db.Get(&exists, checkTablesQuery)
	if err != nil {
		panic(err)
	}
	if exists {
		return
	}

	createTablesQuery := `
		create table if not exists _User(
		    id varchar primary key
		);
		
		create table if not exists Reminder(
		    id uuid primary key,
		    medicine_name varchar(20) not null,
		    dosage varchar(20) not null,
		    description varchar(20) not null,
		    user_id string not null
		);
		
		create table if not exists ReminderDate(
		    id uuid primary key,
		    remind_date Date not null,
		    status BOOLEAN not null,
		    reminder_id uuid not null 
		);
		`
	_, err = db.Exec(createTablesQuery)
	if err != nil {
		panic(err)
	}

}
