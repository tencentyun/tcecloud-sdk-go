package main

import (
	"fmt"
	"os"

	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/errors"
	"github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
	vpc "github.com/tencentyun/tcecloud-sdk-go/tcecloud/vpc/v20170312"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TCECLOUD_SECRET_ID"),
		os.Getenv("TCECLOUD_SECRET_KEY"),
	)
	client, _ := vpc.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())
	sgreq := vpc.NewDescribeSecurityGroupsRequest()
	sgresp, err := client.DescribeSecurityGroups(sgreq)
	if _, ok := err.(*errors.TceCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if *sgresp.Response.TotalCount < 1 {
		fmt.Printf("No security group found.")
		return
	}
	request := vpc.NewDescribeSecurityGroupPoliciesRequest()
	// we don't care which one, hence with no filter and just use the first one
	request.SecurityGroupId = sgresp.Response.SecurityGroupSet[0].SecurityGroupId
	response, err := client.DescribeSecurityGroupPolicies(request)
	if _, ok := err.(*errors.TceCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
