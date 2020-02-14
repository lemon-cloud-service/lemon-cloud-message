package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common/dto"
	lemon_cloud_user_sdk "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-sdk"
	client "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-sdk/client"
	"google.golang.org/grpc"
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

	// 从原理测试

	con, err := grpc.Dial("localhost:33385", grpc.WithInsecure())
	if err != nil {
		fmt.Println("初始化失败")
	}
	con.Invoke(context.Background(), "/service.UserLoginService/LoginByNumber", in, out, opts...)
	proto.Marshal()

}
