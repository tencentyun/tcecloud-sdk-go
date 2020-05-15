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

package v20190411

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-04-11"

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


func NewAddTopicRequest() (request *AddTopicRequest) {
    request = &AddTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcemsg", APIVersion, "AddTopic")
    return
}

func NewAddTopicResponse() (response *AddTopicResponse) {
    response = &AddTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加订阅类型
func (c *Client) AddTopic(request *AddTopicRequest) (response *AddTopicResponse, err error) {
    if request == nil {
        request = NewAddTopicRequest()
    }
    response = NewAddTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDelTopicRequest() (request *DelTopicRequest) {
    request = &DelTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcemsg", APIVersion, "DelTopic")
    return
}

func NewDelTopicResponse() (response *DelTopicResponse) {
    response = &DelTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除订阅类型
func (c *Client) DelTopic(request *DelTopicRequest) (response *DelTopicResponse, err error) {
    if request == nil {
        request = NewDelTopicRequest()
    }
    response = NewDelTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDelUserCmgtSiteMsgRequest() (request *DelUserCmgtSiteMsgRequest) {
    request = &DelUserCmgtSiteMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcemsg", APIVersion, "DelUserCmgtSiteMsg")
    return
}

func NewDelUserCmgtSiteMsgResponse() (response *DelUserCmgtSiteMsgResponse) {
    response = &DelUserCmgtSiteMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除运营端用户站内信
func (c *Client) DelUserCmgtSiteMsg(request *DelUserCmgtSiteMsgRequest) (response *DelUserCmgtSiteMsgResponse, err error) {
    if request == nil {
        request = NewDelUserCmgtSiteMsgRequest()
    }
    response = NewDelUserCmgtSiteMsgResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicRequest() (request *GetTopicRequest) {
    request = &GetTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcemsg", APIVersion, "GetTopic")
    return
}

func NewGetTopicResponse() (response *GetTopicResponse) {
    response = &GetTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询订阅类型
func (c *Client) GetTopic(request *GetTopicRequest) (response *GetTopicResponse, err error) {
    if request == nil {
        request = NewGetTopicRequest()
    }
    response = NewGetTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserCmgtSiteMsgRequest() (request *GetUserCmgtSiteMsgRequest) {
    request = &GetUserCmgtSiteMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcemsg", APIVersion, "GetUserCmgtSiteMsg")
    return
}

func NewGetUserCmgtSiteMsgResponse() (response *GetUserCmgtSiteMsgResponse) {
    response = &GetUserCmgtSiteMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询运营端用户站内信
func (c *Client) GetUserCmgtSiteMsg(request *GetUserCmgtSiteMsgRequest) (response *GetUserCmgtSiteMsgResponse, err error) {
    if request == nil {
        request = NewGetUserCmgtSiteMsgRequest()
    }
    response = NewGetUserCmgtSiteMsgResponse()
    err = c.Send(request, response)
    return
}

func NewReadUserCmgtSiteMsgRequest() (request *ReadUserCmgtSiteMsgRequest) {
    request = &ReadUserCmgtSiteMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcemsg", APIVersion, "ReadUserCmgtSiteMsg")
    return
}

func NewReadUserCmgtSiteMsgResponse() (response *ReadUserCmgtSiteMsgResponse) {
    response = &ReadUserCmgtSiteMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新运营端用户站内信状态
func (c *Client) ReadUserCmgtSiteMsg(request *ReadUserCmgtSiteMsgRequest) (response *ReadUserCmgtSiteMsgResponse, err error) {
    if request == nil {
        request = NewReadUserCmgtSiteMsgRequest()
    }
    response = NewReadUserCmgtSiteMsgResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTopicRequest() (request *UpdateTopicRequest) {
    request = &UpdateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcemsg", APIVersion, "UpdateTopic")
    return
}

func NewUpdateTopicResponse() (response *UpdateTopicResponse) {
    response = &UpdateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改订阅类型
func (c *Client) UpdateTopic(request *UpdateTopicRequest) (response *UpdateTopicResponse, err error) {
    if request == nil {
        request = NewUpdateTopicRequest()
    }
    response = NewUpdateTopicResponse()
    err = c.Send(request, response)
    return
}
