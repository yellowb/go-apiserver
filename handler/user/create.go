package user

import (
	"apiserver/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error

	if err = c.Bind(&r); err == nil {
		if r.Username == "" {
			err = errno.New(errno.UserFieldMissed, fmt.Errorf("username filed is missed"))
		}
		if r.Password == "" {
			err = errno.New(errno.PwdFieldMissed, fmt.Errorf("password filed is missed"))
		}
	} else {
		err = errno.New(errno.ErrBind, err)
	}

	if err != nil {
		log.Error("Invalid Request: ", err)
	} else {
		log.Infof("Valid Request: username = %s, password = %s", r.Username, r.Password)
	}

	code, message := errno.Decode(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})

}
