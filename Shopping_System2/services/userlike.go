package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
	"Shopping_System/untils"
)

func SearchGoodInLike(myclaims untils.MyClaims) (interface{}, error) {

	return mysql.SearchGoodInLike(myclaims.Uid)
}
func LikeAGood(LikeGood *model.UserLike, myclaims untils.MyClaims) error {
	if err := mysql.CheckGid(LikeGood.Gid); err != nil {
		return err
	}
	return mysql.LikeAGood(LikeGood.Gid, myclaims.Uid)
}
