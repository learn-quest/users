package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// injecting database session into middlewares so that it will accessible from everywhere
func DbSession(session *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// setting session
		c.Set("dbSession", session)
		c.Next() // Proceed to the next handler
	}
}
