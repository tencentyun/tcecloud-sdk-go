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

package v20180808

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-08-08"

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


func NewDescribeLogSearchRequest() (request *DescribeLogSearchRequest) {
    request = &DescribeLogSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeLogSearch")
    return
}

func NewDescribeLogSearchResponse() (response *DescribeLogSearchResponse) {
    response = &DescribeLogSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于搜索日志
func (c *Client) DescribeLogSearch(request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    if request == nil {
        request = NewDescribeLogSearchRequest()
    }
    response = NewDescribeLogSearchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiIncrementRequest() (request *ModifyApiIncrementRequest) {
    request = &ModifyApiIncrementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiIncrement")
    return
}

func NewModifyApiIncrementResponse() (response *ModifyApiIncrementResponse) {
    response = &ModifyApiIncrementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供增量更新API能力，主要是给程序调用（区别于ModifyApi，该接口是需要传入API的全量参数，对console使用较友好）
func (c *Client) ModifyApiIncrement(request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    if request == nil {
        request = NewModifyApiIncrementRequest()
    }
    response = NewModifyApiIncrementResponse()
    err = c.Send(request, response)
    return
}
