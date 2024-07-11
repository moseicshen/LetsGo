package models

import (
	"LetsGo/config"
)

// List TodoList Model
type List struct {
	ID         uint   `gorm:"primary_key;autoIncrement"`
	Note       string `gorm:"not null"`
	IsComplete bool   `gorm:"default:false"`
}

func ShowLists() (lists []*List, err error) {
	err = config.DB.Find(&lists).Error
	if err != nil {
		return nil, err
	}
	return
}

func CreateList(list *List) (err error) {
	err = config.DB.Create(list).Error
	if err != nil {
		return err
	}
	return
}

func DeleteList(id uint) (err error) {
	err = config.DB.Delete(&List{}, id).Error
	if err != nil {
		return err
	}
	return
}

func UpdateList(id uint) {
	list := &List{
		ID: id,
	}
	config.DB.Model(&list).Update("IsComplete", true)
}
