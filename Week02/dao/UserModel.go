package dao

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func GetUserInfoModel() (interface{}, error) {
	db,err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test") //连接数据库
	defer db.Close()
	user_info , err := db.Query("select * from user where uid = 10")
	if err != nil {
		return user_info,errors.Wrap(sql.ErrNoRows,"NOT DATA")
	}

	return user_info , nil
}
