module github.com/lemon-cloud/service/lemon-cloud-message/lemon-cloud-message-service

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common v0.0.0-20200119084700-2e9816bca4a8
	github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-sdk v0.0.0-20200119084700-2e9816bca4a8
	google.golang.org/grpc v1.26.0
)

replace github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-sdk => ../../lemon-cloud-user/lemon-cloud-user-sdk

replace github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common => ../../lemon-cloud-user/lemon-cloud-user-common
