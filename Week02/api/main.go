package main

import (
	"fmt"
	"work/service"
)

func main() {
	user_info , err := service.GetUserInfoService()

	if err != nil {
		fmt.Printf("query error %+v\n",err)
	}

	fmt.Println(user_info)
}
