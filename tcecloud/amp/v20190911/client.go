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


func NewAddConvergenceRequest() (request *AddConvergenceRequest) {
    request = &AddConvergenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "AddConvergence")
    return
}

func NewAddConvergenceResponse() (response *AddConvergenceResponse) {
    response = &AddConvergenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加收敛规则
func (c *Client) AddConvergence(request *AddConvergenceRequest) (response *AddConvergenceResponse, err error) {
    if request == nil {
        request = NewAddConvergenceRequest()
    }
    response = NewAddConvergenceResponse()
    err = c.Send(request, response)
    return
}

func NewAddHandleRequest() (request *AddHandleRequest) {
    request = &AddHandleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "AddHandle")
    return
}

func NewAddHandleResponse() (response *AddHandleResponse) {
    response = &AddHandleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加处理规则
func (c *Client) AddHandle(request *AddHandleRequest) (response *AddHandleResponse, err error) {
    if request == nil {
        request = NewAddHandleRequest()
    }
    response = NewAddHandleResponse()
    err = c.Send(request, response)
    return
}

func NewAddShieldRequest() (request *AddShieldRequest) {
    request = &AddShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "AddShield")
    return
}

func NewAddShieldResponse() (response *AddShieldResponse) {
    response = &AddShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加屏蔽规则
func (c *Client) AddShield(request *AddShieldRequest) (response *AddShieldResponse, err error) {
    if request == nil {
        request = NewAddShieldRequest()
    }
    response = NewAddShieldResponse()
    err = c.Send(request, response)
    return
}

func NewAddSubRequest() (request *AddSubRequest) {
    request = &AddSubRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "AddSub")
    return
}

func NewAddSubResponse() (response *AddSubResponse) {
    response = &AddSubResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加订阅
func (c *Client) AddSub(request *AddSubRequest) (response *AddSubResponse, err error) {
    if request == nil {
        request = NewAddSubRequest()
    }
    response = NewAddSubResponse()
    err = c.Send(request, response)
    return
}

func NewAddTopicRequest() (request *AddTopicRequest) {
    request = &AddTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "AddTopic")
    return
}

func NewAddTopicResponse() (response *AddTopicResponse) {
    response = &AddTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加告警源
func (c *Client) AddTopic(request *AddTopicRequest) (response *AddTopicResponse, err error) {
    if request == nil {
        request = NewAddTopicRequest()
    }
    response = NewAddTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConvergencesRequest() (request *DeleteConvergencesRequest) {
    request = &DeleteConvergencesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "DeleteConvergences")
    return
}

func NewDeleteConvergencesResponse() (response *DeleteConvergencesResponse) {
    response = &DeleteConvergencesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除多个收敛规则
func (c *Client) DeleteConvergences(request *DeleteConvergencesRequest) (response *DeleteConvergencesResponse, err error) {
    if request == nil {
        request = NewDeleteConvergencesRequest()
    }
    response = NewDeleteConvergencesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHandlesRequest() (request *DeleteHandlesRequest) {
    request = &DeleteHandlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "DeleteHandles")
    return
}

func NewDeleteHandlesResponse() (response *DeleteHandlesResponse) {
    response = &DeleteHandlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除多个处理规则
func (c *Client) DeleteHandles(request *DeleteHandlesRequest) (response *DeleteHandlesResponse, err error) {
    if request == nil {
        request = NewDeleteHandlesRequest()
    }
    response = NewDeleteHandlesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShieldsRequest() (request *DeleteShieldsRequest) {
    request = &DeleteShieldsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "DeleteShields")
    return
}

func NewDeleteShieldsResponse() (response *DeleteShieldsResponse) {
    response = &DeleteShieldsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除多个屏蔽规则
func (c *Client) DeleteShields(request *DeleteShieldsRequest) (response *DeleteShieldsResponse, err error) {
    if request == nil {
        request = NewDeleteShieldsRequest()
    }
    response = NewDeleteShieldsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubsRequest() (request *DeleteSubsRequest) {
    request = &DeleteSubsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "DeleteSubs")
    return
}

func NewDeleteSubsResponse() (response *DeleteSubsResponse) {
    response = &DeleteSubsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除多个订阅
func (c *Client) DeleteSubs(request *DeleteSubsRequest) (response *DeleteSubsResponse, err error) {
    if request == nil {
        request = NewDeleteSubsRequest()
    }
    response = NewDeleteSubsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicsRequest() (request *DeleteTopicsRequest) {
    request = &DeleteTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "DeleteTopics")
    return
}

func NewDeleteTopicsResponse() (response *DeleteTopicsResponse) {
    response = &DeleteTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除多个告警主题
func (c *Client) DeleteTopics(request *DeleteTopicsRequest) (response *DeleteTopicsResponse, err error) {
    if request == nil {
        request = NewDeleteTopicsRequest()
    }
    response = NewDeleteTopicsResponse()
    err = c.Send(request, response)
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

func NewGetAlertHistoryByIdRequest() (request *GetAlertHistoryByIdRequest) {
    request = &GetAlertHistoryByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetAlertHistoryById")
    return
}

func NewGetAlertHistoryByIdResponse() (response *GetAlertHistoryByIdResponse) {
    response = &GetAlertHistoryByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警历史详情
func (c *Client) GetAlertHistoryById(request *GetAlertHistoryByIdRequest) (response *GetAlertHistoryByIdResponse, err error) {
    if request == nil {
        request = NewGetAlertHistoryByIdRequest()
    }
    response = NewGetAlertHistoryByIdResponse()
    err = c.Send(request, response)
    return
}

func NewGetConfigHistoryRequest() (request *GetConfigHistoryRequest) {
    request = &GetConfigHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetConfigHistory")
    return
}

func NewGetConfigHistoryResponse() (response *GetConfigHistoryResponse) {
    response = &GetConfigHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取修改历史
func (c *Client) GetConfigHistory(request *GetConfigHistoryRequest) (response *GetConfigHistoryResponse, err error) {
    if request == nil {
        request = NewGetConfigHistoryRequest()
    }
    response = NewGetConfigHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewGetConvergenceRequest() (request *GetConvergenceRequest) {
    request = &GetConvergenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetConvergence")
    return
}

func NewGetConvergenceResponse() (response *GetConvergenceResponse) {
    response = &GetConvergenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个收敛规则
func (c *Client) GetConvergence(request *GetConvergenceRequest) (response *GetConvergenceResponse, err error) {
    if request == nil {
        request = NewGetConvergenceRequest()
    }
    response = NewGetConvergenceResponse()
    err = c.Send(request, response)
    return
}

func NewGetConvergencesRequest() (request *GetConvergencesRequest) {
    request = &GetConvergencesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetConvergences")
    return
}

func NewGetConvergencesResponse() (response *GetConvergencesResponse) {
    response = &GetConvergencesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取收敛规则
func (c *Client) GetConvergences(request *GetConvergencesRequest) (response *GetConvergencesResponse, err error) {
    if request == nil {
        request = NewGetConvergencesRequest()
    }
    response = NewGetConvergencesResponse()
    err = c.Send(request, response)
    return
}

func NewGetHandleRequest() (request *GetHandleRequest) {
    request = &GetHandleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetHandle")
    return
}

func NewGetHandleResponse() (response *GetHandleResponse) {
    response = &GetHandleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个告警处理
func (c *Client) GetHandle(request *GetHandleRequest) (response *GetHandleResponse, err error) {
    if request == nil {
        request = NewGetHandleRequest()
    }
    response = NewGetHandleResponse()
    err = c.Send(request, response)
    return
}

func NewGetHandlesRequest() (request *GetHandlesRequest) {
    request = &GetHandlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetHandles")
    return
}

func NewGetHandlesResponse() (response *GetHandlesResponse) {
    response = &GetHandlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取处理动作
func (c *Client) GetHandles(request *GetHandlesRequest) (response *GetHandlesResponse, err error) {
    if request == nil {
        request = NewGetHandlesRequest()
    }
    response = NewGetHandlesResponse()
    err = c.Send(request, response)
    return
}

func NewGetShieldRequest() (request *GetShieldRequest) {
    request = &GetShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetShield")
    return
}

func NewGetShieldResponse() (response *GetShieldResponse) {
    response = &GetShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个屏蔽
func (c *Client) GetShield(request *GetShieldRequest) (response *GetShieldResponse, err error) {
    if request == nil {
        request = NewGetShieldRequest()
    }
    response = NewGetShieldResponse()
    err = c.Send(request, response)
    return
}

func NewGetShieldsRequest() (request *GetShieldsRequest) {
    request = &GetShieldsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetShields")
    return
}

func NewGetShieldsResponse() (response *GetShieldsResponse) {
    response = &GetShieldsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取屏蔽规则
func (c *Client) GetShields(request *GetShieldsRequest) (response *GetShieldsResponse, err error) {
    if request == nil {
        request = NewGetShieldsRequest()
    }
    response = NewGetShieldsResponse()
    err = c.Send(request, response)
    return
}

func NewGetSubRequest() (request *GetSubRequest) {
    request = &GetSubRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetSub")
    return
}

func NewGetSubResponse() (response *GetSubResponse) {
    response = &GetSubResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个订阅规则
func (c *Client) GetSub(request *GetSubRequest) (response *GetSubResponse, err error) {
    if request == nil {
        request = NewGetSubRequest()
    }
    response = NewGetSubResponse()
    err = c.Send(request, response)
    return
}

func NewGetSubsRequest() (request *GetSubsRequest) {
    request = &GetSubsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetSubs")
    return
}

func NewGetSubsResponse() (response *GetSubsResponse) {
    response = &GetSubsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取订阅
func (c *Client) GetSubs(request *GetSubsRequest) (response *GetSubsResponse, err error) {
    if request == nil {
        request = NewGetSubsRequest()
    }
    response = NewGetSubsResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicRequest() (request *GetTopicRequest) {
    request = &GetTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetTopic")
    return
}

func NewGetTopicResponse() (response *GetTopicResponse) {
    response = &GetTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个告警主题
func (c *Client) GetTopic(request *GetTopicRequest) (response *GetTopicResponse, err error) {
    if request == nil {
        request = NewGetTopicRequest()
    }
    response = NewGetTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicFieldTemplateRequest() (request *GetTopicFieldTemplateRequest) {
    request = &GetTopicFieldTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetTopicFieldTemplate")
    return
}

func NewGetTopicFieldTemplateResponse() (response *GetTopicFieldTemplateResponse) {
    response = &GetTopicFieldTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警参数信息
func (c *Client) GetTopicFieldTemplate(request *GetTopicFieldTemplateRequest) (response *GetTopicFieldTemplateResponse, err error) {
    if request == nil {
        request = NewGetTopicFieldTemplateRequest()
    }
    response = NewGetTopicFieldTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicsRequest() (request *GetTopicsRequest) {
    request = &GetTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "GetTopics")
    return
}

func NewGetTopicsResponse() (response *GetTopicsResponse) {
    response = &GetTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取所有告警主题
func (c *Client) GetTopics(request *GetTopicsRequest) (response *GetTopicsResponse, err error) {
    if request == nil {
        request = NewGetTopicsRequest()
    }
    response = NewGetTopicsResponse()
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

func NewUpdateConvergenceRequest() (request *UpdateConvergenceRequest) {
    request = &UpdateConvergenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "UpdateConvergence")
    return
}

func NewUpdateConvergenceResponse() (response *UpdateConvergenceResponse) {
    response = &UpdateConvergenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改收敛规则
func (c *Client) UpdateConvergence(request *UpdateConvergenceRequest) (response *UpdateConvergenceResponse, err error) {
    if request == nil {
        request = NewUpdateConvergenceRequest()
    }
    response = NewUpdateConvergenceResponse()
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

func NewUpdateSubRequest() (request *UpdateSubRequest) {
    request = &UpdateSubRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "UpdateSub")
    return
}

func NewUpdateSubResponse() (response *UpdateSubResponse) {
    response = &UpdateSubResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改订阅
func (c *Client) UpdateSub(request *UpdateSubRequest) (response *UpdateSubResponse, err error) {
    if request == nil {
        request = NewUpdateSubRequest()
    }
    response = NewUpdateSubResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTopicRequest() (request *UpdateTopicRequest) {
    request = &UpdateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("amp", APIVersion, "UpdateTopic")
    return
}

func NewUpdateTopicResponse() (response *UpdateTopicResponse) {
    response = &UpdateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改告警主题
func (c *Client) UpdateTopic(request *UpdateTopicRequest) (response *UpdateTopicResponse, err error) {
    if request == nil {
        request = NewUpdateTopicRequest()
    }
    response = NewUpdateTopicResponse()
    err = c.Send(request, response)
    return
}
