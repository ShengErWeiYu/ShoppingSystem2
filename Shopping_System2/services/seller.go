package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/untils"
	"log"
)

func CheckOrder(myclaims untils.MyClaims) (interface{}, error) {
	if err := mysql.CheckOrderExist(myclaims.Uid); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return mysql.CheckOrder(myclaims.Uid)
}
