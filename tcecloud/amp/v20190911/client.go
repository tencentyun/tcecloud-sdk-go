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

package v20190911

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-09-11"

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


func NewGetAlertHistoryRequest() (request *GetAlertHistoryRequest) {
    request = &GetAlertHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetAlertHistory")
    return
}

func NewGetAlertHistoryResponse() (response *GetAlertHistoryResponse) {
    response = &GetAlertHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警历史
func (c *Client) GetAlertHistory(request *GetAlertHistoryRequest) (response *GetAlertHistoryResponse, err error) {
    if request == nil {
        request = NewGetAlertHistoryRequest()
    }
    response = NewGetAlertHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewSendAlarmRequest() (request *SendAlarmRequest) {
    request = &SendAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "SendAlarm")
    return
}

func NewSendAlarmResponse() (response *SendAlarmResponse) {
    response = &SendAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发送告警接口
func (c *Client) SendAlarm(request *SendAlarmRequest) (response *SendAlarmResponse, err error) {
    if request == nil {
        request = NewSendAlarmRequest()
    }
    response = NewSendAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateHandleRequest() (request *UpdateHandleRequest) {
    request = &UpdateHandleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "UpdateHandle")
    return
}

func NewUpdateHandleResponse() (response *UpdateHandleResponse) {
    response = &UpdateHandleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改处理规则
func (c *Client) UpdateHandle(request *UpdateHandleRequest) (response *UpdateHandleResponse, err error) {
    if request == nil {
        request = NewUpdateHandleRequest()
    }
    response = NewUpdateHandleResponse()
    err = c.Send(request, response)
    return
}
