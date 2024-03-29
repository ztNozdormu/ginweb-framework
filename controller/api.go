package controller

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ztNozdormu/ginweb-framework/dao"
	"github.com/ztNozdormu/ginweb-framework/dto"
	"github.com/ztNozdormu/ginweb-framework/public"
	"strconv"
	"strings"
	"time"
)

type Api struct {
}

func (demo *Api) Login(c *gin.Context) {
	api:=&dto.LoginInput{}
	if err:=api.BindingValidParams(c);err!=nil{
		public.ResponseError(c,2001,err)
		return
	}
	//todo 这块应该从db验证
	if api.Username=="admin" && api.Password=="123456"{
		session := sessions.Default(c)
		session.Set("user", api.Username)
		session.Save()
		public.ResponseSuccess(c, api,"")
	}else{
		public.ResponseError(c,2002,errors.New("账号或密码错误"))
	}
	return
}


func (demo *Api) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	return
}


func (demo *Api) ListPage(c *gin.Context) {
	listInput:=&dto.ListPageInput{}
	if err:=listInput.BindingValidParams(c);err!=nil{
		public.ResponseError(c,2003,err)
		return
	}
	user:=&dao.User{}
	pageInt,err:=strconv.ParseInt(listInput.Page,10,64)
	if err!=nil{
		public.ResponseError(c,2004,err)
		return
	}
	if userList,total,err:=user.PageList(listInput.Name,int(pageInt),20);err!=nil{
		public.ResponseError(c,2005,err)
		return
	}else{
		m:=map[string]interface{}{
			"list":userList,
			"total":total,
		}
		public.ResponseSuccess(c, m,"")
	}
	return
}

func (demo *Api) AddUser(c *gin.Context) {
	addInput:=&dto.AddUserInput{}
	if err:=addInput.BindingValidParams(c);err!=nil{
		public.ResponseError(c,2006,err)
		return
	}
	user:=&dao.User{}
	user.Name = addInput.Name
	user.Sex = addInput.Sex
	user.Age = addInput.Age
	user.Birth = addInput.Birth
	user.Addr = addInput.Addr
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()
	if err:=user.Save();err!=nil{
		public.ResponseError(c,2007,err)
		return
	}
	public.ResponseSuccess(c,"","")
	return
}

func (demo *Api) EditUser(c *gin.Context) {
	editInput:=&dto.EditUserInput{}
	if err:=editInput.BindingValidParams(c);err!=nil{
		public.ResponseError(c,2006,err)
		return
	}

	user:=&dao.User{}
	if userDb,err:=user.Find(int64(editInput.Id));err!=nil{
		public.ResponseError(c,2006,err)
		return
	}else{
		user = userDb
	}
	user.Name = editInput.Name
	user.Sex = editInput.Sex
	user.Age = editInput.Age
	user.Birth = editInput.Birth
	user.Addr = editInput.Addr
	user.UpdateAt = time.Now()
	if err:=user.Save();err!=nil{
		public.ResponseError(c,2007,err)
		return
	}
	public.ResponseSuccess(c,"","")
	return
}


func (demo *Api) RemoveUser(c *gin.Context) {
	removeInput:=&dto.RemoveUserInput{}
	if err:=removeInput.BindingValidParams(c);err!=nil{
		public.ResponseError(c,2006,err)
		return
	}

	user:=&dao.User{}
	if err:=user.Del(strings.Split(removeInput.IDS,","));err!=nil{
		public.ResponseError(c,2007,err)
		return
	}
	public.ResponseSuccess(c,"","")
	return
}
