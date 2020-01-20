package main

import (
	"context"
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common/dto"
	lemon_cloud_user_sdk "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-sdk"
	client "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-sdk/client"
)

func main() {
	lemon_cloud_user_sdk.InitSdk("localhost:33385")
	rsp, err := client.GetUserLoginServiceClient().LoginByNumber(context.Background(), &dto.LoginByNumberReq{
		ClientInfo: nil,
		Number:     "123",
		Password:   "234",
	})
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(rsp.Token)
}
