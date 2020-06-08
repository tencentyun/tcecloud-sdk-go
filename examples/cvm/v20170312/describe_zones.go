package main

import (
	"fmt"
	"os"

	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/errors"
	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
	cvm "github.com/tencentyun/tcecloud-sdk-go/tcecloud/cvm/v20170312"
)

func main() {
	// 实例化一个认证对象，入参需要传入TCE账户secretId，secretKey
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值
	credential := common.NewCredential(
		os.Getenv("TCECLOUD_SECRET_ID"),
		os.Getenv("TCECLOUD_SECRET_KEY"),
	)

	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	cpf.HttpProfile.ReqTimeout = 5

	// 实例化要请求产品(以cvm为例)的client对象
	client, _ := cvm.NewClient(credential, "ap-beijing", cpf)
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	request := cvm.NewDescribeZonesRequest()
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.DescribeZones(request)
	// 处理异常
	if _, ok := err.(*errors.TceCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// unexpected errors
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
}
