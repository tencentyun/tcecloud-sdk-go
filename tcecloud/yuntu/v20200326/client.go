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

package v20200326

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-26"

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


func NewGetModuleNamesRequest() (request *GetModuleNamesRequest) {
    request = &GetModuleNamesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yuntu", APIVersion, "GetModuleNames")
    return
}

func NewGetModuleNamesResponse() (response *GetModuleNamesResponse) {
    response = &GetModuleNamesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回所有的模块名
func (c *Client) GetModuleNames(request *GetModuleNamesRequest) (response *GetModuleNamesResponse, err error) {
    if request == nil {
        request = NewGetModuleNamesRequest()
    }
    response = NewGetModuleNamesResponse()
    err = c.Send(request, response)
    return
}

func NewQueryApiInfoFilterRequest() (request *QueryApiInfoFilterRequest) {
    request = &QueryApiInfoFilterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yuntu", APIVersion, "QueryApiInfoFilter")
    return
}

func NewQueryApiInfoFilterResponse() (response *QueryApiInfoFilterResponse) {
    response = &QueryApiInfoFilterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分页、搜索列表API接口
func (c *Client) QueryApiInfoFilter(request *QueryApiInfoFilterRequest) (response *QueryApiInfoFilterResponse, err error) {
    if request == nil {
        request = NewQueryApiInfoFilterRequest()
    }
    response = NewQueryApiInfoFilterResponse()
    err = c.Send(request, response)
    return
}
