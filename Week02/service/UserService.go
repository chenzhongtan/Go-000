package service

import "work/dao"

func GetUserInfoService() (interface{} , error) {
	return dao.GetUserInfoModel()
}
