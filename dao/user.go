package dao

import (
	"github.com/e421083458/gorm"
	"github.com/ztNozdormu/ginweb-framework/public"
	"time"
)

type User struct {
	Id       int       `json:"id" orm:"column(id);auto"`
	Name     string    `json:"name" orm:"column(name);"`
	Addr     string    `json:"addr" orm:"column(addr);"`
	Age      int       `json:"age" orm:"column(age);"`
	Birth    string    `json:"birth" orm:"column(birth);"`
	Sex      int       `json:"sex" orm:"column(sex);"`
	UpdateAt time.Time `json:"update_at" orm:"column(update_at); description:"更新时间"`
	CreateAt time.Time `json:"create_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
}

func (f *User) TableName() string {
	return "user"
}


func (f *User) Del(idSlice []string) error {
	err := public.GormPool.Where("id in (?)", idSlice).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}


func (f *User) Find(id int64) (*User, error) {
	var user User
	err := public.GormPool.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (f *User) PageList(name string, pageNo int, pageSize int) ([]*User, int64, error) {
	var user []*User
	var userCount int64
	//limit offset,pagesize
	offset := (pageNo - 1) * pageSize
	query := public.GormPool
	if name != "" {
		query = query.Where("name = ?", name)
	}

	err := query.Limit(pageSize).Offset(offset).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Table("user").Count(&userCount).Error
	if errCount != nil {
		return nil, 0, err
	}
	return user, userCount, nil
}

func (f *User) Save() error {
	if err:=public.GormPool.Save(f).Error;err!=nil{
		return err
	}
	return nil
}
