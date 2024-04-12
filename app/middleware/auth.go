package middleware

import (
	"context"
	"encoding/json"
	"fuge/app/constant"
	"fuge/app/core"
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		redisKey := constant.UserToken + token
		ctx := context.Background()
		var err error
		userInfoStr, err := core.RedisClient.Get(ctx, redisKey).Result()
		if err != nil {
			c.JSON(403, gin.H{})
			c.Abort()
			return
		}
		userInfo := &models.UserSchema{}
		err = json.Unmarshal([]byte(userInfoStr), userInfo)
		if err != nil {
			c.JSON(403, gin.H{})
			c.Abort()
			return
		}
		user, err := daos.UserDAO.DoGetUserByUserID(userInfo.ID)
		if err != nil {
			c.JSON(403, gin.H{})
			c.Abort()
			return
		}
		if user.Status != models.Complete {
			c.JSON(401, gin.H{})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
