package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
	"Shopping_System/untils"
)

func LookComment(Look *model.Look) (interface{}, error) {
	if err := mysql.CheckGid(Look.Gid); err != nil {
		return nil, err
	}
	return mysql.LookComment(Look.Gid)
}
func PublishComment(PB *model.Comments, myclaims untils.MyClaims) error {
	if err := mysql.CheckGoodOrderExist(PB.Gid); err != nil { //查询在order表里有没有购买这个商品的记录
		return err
	}
	if err := mysql.CheckOrderWithUid(myclaims.Uid, PB.Gid); err != nil {
		return err
	} //查看该用户是否购买该商品
	return mysql.PubulishComment(myclaims.Uid, PB.Gid, PB.Star, PB.Comment)
}
