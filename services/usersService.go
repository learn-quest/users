package services

import (
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/learn-quest/users/models"
)

func InserUser(c *gin.Context, user *models.User) error {
	/* This function will insert the data into the database */

	// getting sesstion from gin context
	dbSession, _ := c.Get("dbSession")
	session := dbSession.(*pgxpool.Pool)

	//  select query to check if any user already exists with given username or email
	selectSqlQuery := "SELECT (_id, name, email, username) FROM users WHERE username = $1 OR email = $2"

	// executing query
	rows, err := session.Query(context.Background(), selectSqlQuery, user.Username, user.Email)
	if err != nil {
		// returning error if any error occurs in execution
		return err
	}
	// closing rows session after execution of function
	defer rows.Close()

	// creating list to store all found users from select query
	var users []models.User

	// looping database records
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user)
		if err != nil {
			// returning error if any error found during scan
			return err
		}
		// appending db user to users list
		users = append(users, user)
		// breaking loop after finding one row since only one user is needed to compare
		break
	}
	// if list len is more than 0 means already an user exists with same username or email
	if len(users) > 0 {
		if user.Username == users[0].Username {
			// returning error if same username found
			return errors.New("an user with this username already exists")
		} else if user.Email == users[0].Email {
			// returning error if same email found
			return errors.New("an user with this email already exists")
		}
	}

	// if no user found inserting user into db
	// generating uuid and storing current time
	uuid, _ := uuid.NewV7()
	currentTime := time.Now().Format(time.RFC3339)

	// assigning default system values
	user.Id = uuid.String()
	user.LastLoggedIn = currentTime
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
	user.IsBanned = false
	user.ProfilePic = ""

	// insert query
	insertSqlQuery := "INSERT INTO users(_id, name, email, username, is_banned, profile_pic, country, last_logged_in, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// executing insert query
	_, err = session.Exec(context.Background(), insertSqlQuery, user.Id, user.Name, user.Email, user.Username, user.IsBanned, user.ProfilePic, user.Country, user.LastLoggedIn, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		// if any error in execution of insert query then returning error
		return err

	}
	// returing nil if no error throughout function execution
	return nil
}
