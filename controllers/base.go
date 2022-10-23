package controllers

import (
	"errors"
	"github.com/colincheng/servicedesk/models"
	"github.com/gin-gonic/gin"
)

var (
	// UserControl 所有的controller类声明都在这儿
	UserControl = &UserController{} //先注释，在编写完相应程序之后再打开
	//MachineControl = &MachineController{} //先注释，在编写完相应程序之后再打开
)

type Response struct {
	Message string
}

type ResponseLogin struct {
	Response
	Token string
}

type ResponseListMachines struct {
	Response
	Machines models.Machines
}

type MachinesStatus struct {
	Mem float64
	Cpu float64
}

type ResponseMachineStatus struct {
	Response
	Data []MachinesStatus
}

func Auth(c *gin.Context) (*models.User, error) {
	token := c.GetHeader("token")
	if token == "" {
		return nil, errors.New("token为空")
	}
	var user models.User
	user.Token = token
	if err := user.FindUserByToken(); err != nil {
		c.JSON(400, Response{Message: "没有权限"})
		return nil, errors.New("没有权限")
	}
	return &user, nil
}
