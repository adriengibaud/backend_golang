package middleware

import (
	"backend_golang/responses"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
    IDToken string `header:"Authorization"`
}

func BasicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){

		// var user models.User

		// if err := c.BindJSON(&user); err != nil {
    //   return
    // }

    // newUser := models.User{
    //         Id:       primitive.NewObjectID(),
    //         Name:     user.Name,
    //         Location: user.Location,
    //         Title:    user.Title,
    //     }

		// body, _ := ioutil.ReadAll(c.Request.Body)

		// fmt.Println(body)

		ht:= authHeader{}

		c.ShouldBindHeader(&ht)


		token:= strings.Split(ht.IDToken, "Bearer ")

		if len(token) < 2 {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Bad token format"})
			c.Abort()
			return
		}

		
		it := token[1]

		randomToken := "123456789"

	//	header := c.Request.Header.Get("Authorization")

		if it != randomToken {
			fmt.Println(it)
			fmt.Println(randomToken)
			 c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Bad token"})
				c.Abort()
				return
		}

		
		c.Next()
	}
}