package services

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/learn-quest/users/models"
)

func InserUser(c *gin.Context, user *models.User) error {

	dbSession, _ := c.Get("dbSession")
	session := dbSession.(*pgxpool.Pool)

	fmt.Println(session)
	uuid, _ := uuid.NewV7()
	currentTime := time.Now().Format(time.RFC3339)

	user.Id = uuid.String()
	user.LastLoggedIn = currentTime
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
	user.IsBanned = false
	user.ProfilePic = ""

	sqlQuery := "INSERT INTO users(_id, name, email, username, is_banned, profile_pic, country, last_logged_in, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	_, err := session.Exec(context.Background(), sqlQuery, user.Id, user.Name, user.Email, user.Username, user.IsBanned, user.ProfilePic, user.Country, user.LastLoggedIn, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
