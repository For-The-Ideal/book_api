package model

type Login struct {
	accout string `gorm:type varchar(20);not null`
}
