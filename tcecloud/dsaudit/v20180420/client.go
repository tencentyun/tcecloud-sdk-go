// All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180420

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-20"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dsaudit", APIVersion, "CreateInstance")
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建数据安全审计实例
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dsaudit", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取实例列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductTypesRequest() (request *DescribeProductTypesRequest) {
    request = &DescribeProductTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dsaudit", APIVersion, "DescribeProductTypes")
    return
}

func NewDescribeProductTypesResponse() (response *DescribeProductTypesResponse) {
    response = &DescribeProductTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取数据安全审计产品规格列表
func (c *Client) DescribeProductTypes(request *DescribeProductTypesRequest) (response *DescribeProductTypesResponse, err error) {
    if request == nil {
        request = NewDescribeProductTypesRequest()
    }
    response = NewDescribeProductTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSaleInfoRequest() (request *DescribeSaleInfoRequest) {
    request = &DescribeSaleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dsaudit", APIVersion, "DescribeSaleInfo")
    return
}

func NewDescribeSaleInfoResponse() (response *DescribeSaleInfoResponse) {
    response = &DescribeSaleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取产品的销售信息，规格，时长
func (c *Client) DescribeSaleInfo(request *DescribeSaleInfoRequest) (response *DescribeSaleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSaleInfoRequest()
    }
    response = NewDescribeSaleInfoResponse()
    err = c.Send(request, response)
    return
}
