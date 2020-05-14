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

package v20180701

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-01"

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


func NewCreateOfflineTaskRequest() (request *CreateOfflineTaskRequest) {
    request = &CreateOfflineTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("prm", APIVersion, "CreateOfflineTask")
    return
}

func NewCreateOfflineTaskResponse() (response *CreateOfflineTaskResponse) {
    response = &CreateOfflineTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设备下线
func (c *Client) CreateOfflineTask(request *CreateOfflineTaskRequest) (response *CreateOfflineTaskResponse, err error) {
    if request == nil {
        request = NewCreateOfflineTaskRequest()
    }
    response = NewCreateOfflineTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOnlineTaskRequest() (request *CreateOnlineTaskRequest) {
    request = &CreateOnlineTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("prm", APIVersion, "CreateOnlineTask")
    return
}

func NewCreateOnlineTaskResponse() (response *CreateOnlineTaskResponse) {
    response = &CreateOnlineTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设备上线
func (c *Client) CreateOnlineTask(request *CreateOnlineTaskRequest) (response *CreateOnlineTaskResponse, err error) {
    if request == nil {
        request = NewCreateOnlineTaskRequest()
    }
    response = NewCreateOnlineTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskListRequest() (request *GetTaskListRequest) {
    request = &GetTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("prm", APIVersion, "GetTaskList")
    return
}

func NewGetTaskListResponse() (response *GetTaskListResponse) {
    response = &GetTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询任务列表
func (c *Client) GetTaskList(request *GetTaskListRequest) (response *GetTaskListResponse, err error) {
    if request == nil {
        request = NewGetTaskListRequest()
    }
    response = NewGetTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewGetTmplByIdRequest() (request *GetTmplByIdRequest) {
    request = &GetTmplByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("prm", APIVersion, "GetTmplById")
    return
}

func NewGetTmplByIdResponse() (response *GetTmplByIdResponse) {
    response = &GetTmplByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过ID查询母机上架流程模板
func (c *Client) GetTmplById(request *GetTmplByIdRequest) (response *GetTmplByIdResponse, err error) {
    if request == nil {
        request = NewGetTmplByIdRequest()
    }
    response = NewGetTmplByIdResponse()
    err = c.Send(request, response)
    return
}

func NewGetTmplListRequest() (request *GetTmplListRequest) {
    request = &GetTmplListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("prm", APIVersion, "GetTmplList")
    return
}

func NewGetTmplListResponse() (response *GetTmplListResponse) {
    response = &GetTmplListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询所有母机上架流程模板
func (c *Client) GetTmplList(request *GetTmplListRequest) (response *GetTmplListResponse, err error) {
    if request == nil {
        request = NewGetTmplListRequest()
    }
    response = NewGetTmplListResponse()
    err = c.Send(request, response)
    return
}

func NewMarkFailRequest() (request *MarkFailRequest) {
    request = &MarkFailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("prm", APIVersion, "MarkFail")
    return
}

func NewMarkFailResponse() (response *MarkFailResponse) {
    response = &MarkFailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 置失败
func (c *Client) MarkFail(request *MarkFailRequest) (response *MarkFailResponse, err error) {
    if request == nil {
        request = NewMarkFailRequest()
    }
    response = NewMarkFailResponse()
    err = c.Send(request, response)
    return
}

func NewRetryTaskRequest() (request *RetryTaskRequest) {
    request = &RetryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("prm", APIVersion, "RetryTask")
    return
}

func NewRetryTaskResponse() (response *RetryTaskResponse) {
    response = &RetryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重试
func (c *Client) RetryTask(request *RetryTaskRequest) (response *RetryTaskResponse, err error) {
    if request == nil {
        request = NewRetryTaskRequest()
    }
    response = NewRetryTaskResponse()
    err = c.Send(request, response)
    return
}
