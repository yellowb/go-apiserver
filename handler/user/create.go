package user

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	var req CreateRequest
	var resp CreateResponse

	var err error

	// Check & build err object
	if err = c.Bind(&req); err == nil {
		if req.Username == "" {
			err = errno.New(errno.UserFieldMissed, fmt.Errorf("username filed is missed"))
		}
		if req.Password == "" {
			err = errno.New(errno.PwdFieldMissed, fmt.Errorf("password filed is missed"))
		}
	} else {
		err = errno.New(errno.ErrBind, err)
	}

	// Logging
	if err != nil {
		log.Error("Invalid Request: ", err)
		handler.SendResponse(c, err, nil)
	} else {
		log.Infof("Valid Request: username = %s, password = %s", req.Username, req.Password)
		resp.Username = req.Username
		handler.SendResponse(c, err, resp)
	}



}
