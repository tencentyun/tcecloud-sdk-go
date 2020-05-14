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

package v20190325

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-25"

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


func NewAddSubAccountRequest() (request *AddSubAccountRequest) {
    request = &AddSubAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("account", APIVersion, "AddSubAccount")
    return
}

func NewAddSubAccountResponse() (response *AddSubAccountResponse) {
    response = &AddSubAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加子账号
func (c *Client) AddSubAccount(request *AddSubAccountRequest) (response *AddSubAccountResponse, err error) {
    if request == nil {
        request = NewAddSubAccountRequest()
    }
    response = NewAddSubAccountResponse()
    err = c.Send(request, response)
    return
}

func NewChangeMailPasswordRequest() (request *ChangeMailPasswordRequest) {
    request = &ChangeMailPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("account", APIVersion, "ChangeMailPassword")
    return
}

func NewChangeMailPasswordResponse() (response *ChangeMailPasswordResponse) {
    response = &ChangeMailPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改邮箱账号密码
func (c *Client) ChangeMailPassword(request *ChangeMailPasswordRequest) (response *ChangeMailPasswordResponse, err error) {
    if request == nil {
        request = NewChangeMailPasswordRequest()
    }
    response = NewChangeMailPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewChangeSubAccountPasswordRequest() (request *ChangeSubAccountPasswordRequest) {
    request = &ChangeSubAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("account", APIVersion, "ChangeSubAccountPassword")
    return
}

func NewChangeSubAccountPasswordResponse() (response *ChangeSubAccountPasswordResponse) {
    response = &ChangeSubAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改子账号密码
func (c *Client) ChangeSubAccountPassword(request *ChangeSubAccountPasswordRequest) (response *ChangeSubAccountPasswordResponse, err error) {
    if request == nil {
        request = NewChangeSubAccountPasswordRequest()
    }
    response = NewChangeSubAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewGetMultiFactorParasRequest() (request *GetMultiFactorParasRequest) {
    request = &GetMultiFactorParasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("account", APIVersion, "GetMultiFactorParas")
    return
}

func NewGetMultiFactorParasResponse() (response *GetMultiFactorParasResponse) {
    response = &GetMultiFactorParasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取多因子登录参数
func (c *Client) GetMultiFactorParas(request *GetMultiFactorParasRequest) (response *GetMultiFactorParasResponse, err error) {
    if request == nil {
        request = NewGetMultiFactorParasRequest()
    }
    response = NewGetMultiFactorParasResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserInfoRequest() (request *GetUserInfoRequest) {
    request = &GetUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("account", APIVersion, "GetUserInfo")
    return
}

func NewGetUserInfoResponse() (response *GetUserInfoResponse) {
    response = &GetUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户信息
func (c *Client) GetUserInfo(request *GetUserInfoRequest) (response *GetUserInfoResponse, err error) {
    if request == nil {
        request = NewGetUserInfoRequest()
    }
    response = NewGetUserInfoResponse()
    err = c.Send(request, response)
    return
}
