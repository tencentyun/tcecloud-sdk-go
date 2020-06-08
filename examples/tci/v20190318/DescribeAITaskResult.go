package main

import (
	"fmt"

	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/errors"
	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
	tci "github.com/tencentyun/tcecloud-sdk-go/tcecloud/tci/v20190318"
)

func main() {
	credential := common.NewCredential(
		// os.Getenv("TCECLOUD_SECRET_ID"),
		// os.Getenv("TCECLOUD_SECRET_KEY"),
		"", "",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 30
	cpf.HttpProfile.Endpoint = "tci.tcecloudapi.com"
	client, _ := tci.NewClient(credential, "ap-guangzhou", cpf)
	request := tci.NewDescribeAITaskResultRequest()
	request.TaskId = common.Int64Ptr(2716561495)
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.DescribeAITaskResult(request)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TceCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
}
