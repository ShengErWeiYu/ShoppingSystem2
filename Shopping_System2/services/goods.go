package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
	"Shopping_System/untils"
)

func GetGoodDetail(GoodDetail *model.GetGoodDetail) (interface{}, error) {
	if err := mysql.CheckGid(GoodDetail.Gid); err != nil {
		return nil, err
	}
	return mysql.GetGoodDetail(GoodDetail.Gid)
}
func ChangeTheSizeOfGood(ChangeSizeGood *model.ChangeSizeGood, myclaims untils.MyClaims) error {
	if err := mysql.CheckGid(ChangeSizeGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(myclaims.Uid, ChangeSizeGood.Gid); err != nil {
		return err
	}
	return mysql.ChangeTheSizeOfGood(ChangeSizeGood.Gid, ChangeSizeGood.NewSize)
}
func ChangeTheIntroductionOfGood(ChangeIntroductionGood *model.ChangeIntroductionGood, myclaims untils.MyClaims) error {
	if err := mysql.CheckGid(ChangeIntroductionGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(myclaims.Uid, ChangeIntroductionGood.Gid); err != nil {
		return err
	}
	return mysql.ChangeTheIntroductionOfGood(ChangeIntroductionGood.Gid, ChangeIntroductionGood.NewIntroduction)
}
func ChangeThePriceOfGood(ChangePriceGood *model.ChangePriceGood, myclaims untils.MyClaims) error {
	if err := mysql.CheckGid(ChangePriceGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(myclaims.Uid, ChangePriceGood.Gid); err != nil {
		return err
	}
	return mysql.ChangeThePriceOfGood(ChangePriceGood.Gid, ChangePriceGood.NewPrice)
}
func ChangeTheNameOfGood(ChangeNameGood *model.ChangeNameGood, myclaims untils.MyClaims) error {
	if err := mysql.CheckGid(ChangeNameGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(myclaims.Uid, ChangeNameGood.Gid); err != nil {
		return err
	}
	return mysql.ChangeTheNameOfGood(ChangeNameGood.Gid, ChangeNameGood.NewGoodName)
}
func DeleteAGood(DeleteGood *model.GoodsDelete, myclaims untils.MyClaims) error {
	if err := mysql.CheckGid(DeleteGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(myclaims.Uid, DeleteGood.Gid); err != nil {
		return err
	}
	return mysql.DeleteAGood(DeleteGood.Gid)
}
func SearchMyGoods(myclaims untils.MyClaims) (interface{}, error) {
	return mysql.SearchMyGoods(myclaims.Uid)
}
func PublishAGood(Good *model.GoodsPublish, myclaims untils.MyClaims) error {
	return mysql.PublishAGood(Good.GoodName, Good.Price, Good.Introduction, Good.Size, myclaims.Uid)
}
