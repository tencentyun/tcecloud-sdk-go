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

package v20190314

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-03-14"

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


func NewCancelWeChatNoticeRequest() (request *CancelWeChatNoticeRequest) {
    request = &CancelWeChatNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "CancelWeChatNotice")
    return
}

func NewCancelWeChatNoticeResponse() (response *CancelWeChatNoticeResponse) {
    response = &CancelWeChatNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 取消微信通知，会发验证二维码邮件
func (c *Client) CancelWeChatNotice(request *CancelWeChatNoticeRequest) (response *CancelWeChatNoticeResponse, err error) {
    if request == nil {
        request = NewCancelWeChatNoticeRequest()
    }
    response = NewCancelWeChatNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewGetCustomAccountRequest() (request *GetCustomAccountRequest) {
    request = &GetCustomAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "GetCustomAccount")
    return
}

func NewGetCustomAccountResponse() (response *GetCustomAccountResponse) {
    response = &GetCustomAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取全部租户端主账号
func (c *Client) GetCustomAccount(request *GetCustomAccountRequest) (response *GetCustomAccountResponse, err error) {
    if request == nil {
        request = NewGetCustomAccountRequest()
    }
    response = NewGetCustomAccountResponse()
    err = c.Send(request, response)
    return
}

func NewGetCustomSubAccountRequest() (request *GetCustomSubAccountRequest) {
    request = &GetCustomSubAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "GetCustomSubAccount")
    return
}

func NewGetCustomSubAccountResponse() (response *GetCustomSubAccountResponse) {
    response = &GetCustomSubAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过主账号拉取子账号
func (c *Client) GetCustomSubAccount(request *GetCustomSubAccountRequest) (response *GetCustomSubAccountResponse, err error) {
    if request == nil {
        request = NewGetCustomSubAccountRequest()
    }
    response = NewGetCustomSubAccountResponse()
    err = c.Send(request, response)
    return
}

func NewGetSwitchInfoRequest() (request *GetSwitchInfoRequest) {
    request = &GetSwitchInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "GetSwitchInfo")
    return
}

func NewGetSwitchInfoResponse() (response *GetSwitchInfoResponse) {
    response = &GetSwitchInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询是否开启微信通道
func (c *Client) GetSwitchInfo(request *GetSwitchInfoRequest) (response *GetSwitchInfoResponse, err error) {
    if request == nil {
        request = NewGetSwitchInfoRequest()
    }
    response = NewGetSwitchInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetVerifyQRCodeRequest() (request *GetVerifyQRCodeRequest) {
    request = &GetVerifyQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "GetVerifyQRCode")
    return
}

func NewGetVerifyQRCodeResponse() (response *GetVerifyQRCodeResponse) {
    response = &GetVerifyQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取验证二维码
func (c *Client) GetVerifyQRCode(request *GetVerifyQRCodeRequest) (response *GetVerifyQRCodeResponse, err error) {
    if request == nil {
        request = NewGetVerifyQRCodeRequest()
    }
    response = NewGetVerifyQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWeChatRequest() (request *ModifyWeChatRequest) {
    request = &ModifyWeChatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "ModifyWeChat")
    return
}

func NewModifyWeChatResponse() (response *ModifyWeChatResponse) {
    response = &ModifyWeChatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更换微信，会重发绑定邮件
func (c *Client) ModifyWeChat(request *ModifyWeChatRequest) (response *ModifyWeChatResponse, err error) {
    if request == nil {
        request = NewModifyWeChatRequest()
    }
    response = NewModifyWeChatResponse()
    err = c.Send(request, response)
    return
}

func NewNoAcceptNoticeRequest() (request *NoAcceptNoticeRequest) {
    request = &NoAcceptNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "NoAcceptNotice")
    return
}

func NewNoAcceptNoticeResponse() (response *NoAcceptNoticeResponse) {
    response = &NoAcceptNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 不接收微信通知
func (c *Client) NoAcceptNotice(request *NoAcceptNoticeRequest) (response *NoAcceptNoticeResponse, err error) {
    if request == nil {
        request = NewNoAcceptNoticeRequest()
    }
    response = NewNoAcceptNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustomAccountRequest() (request *QueryCustomAccountRequest) {
    request = &QueryCustomAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "QueryCustomAccount")
    return
}

func NewQueryCustomAccountResponse() (response *QueryCustomAccountResponse) {
    response = &QueryCustomAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运营端客户管理查询租户端账号
func (c *Client) QueryCustomAccount(request *QueryCustomAccountRequest) (response *QueryCustomAccountResponse, err error) {
    if request == nil {
        request = NewQueryCustomAccountRequest()
    }
    response = NewQueryCustomAccountResponse()
    err = c.Send(request, response)
    return
}

func NewSendBindingEmailRequest() (request *SendBindingEmailRequest) {
    request = &SendBindingEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("platform", APIVersion, "SendBindingEmail")
    return
}

func NewSendBindingEmailResponse() (response *SendBindingEmailResponse) {
    response = &SendBindingEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发送二维码绑定邮件
func (c *Client) SendBindingEmail(request *SendBindingEmailRequest) (response *SendBindingEmailResponse, err error) {
    if request == nil {
        request = NewSendBindingEmailRequest()
    }
    response = NewSendBindingEmailResponse()
    err = c.Send(request, response)
    return
}
