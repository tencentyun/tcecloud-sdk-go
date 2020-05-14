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

package v20191018

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-18"

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


func NewCreateSCRTaskRequest() (request *CreateSCRTaskRequest) {
    request = &CreateSCRTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "CreateSCRTask")
    return
}

func NewCreateSCRTaskResponse() (response *CreateSCRTaskResponse) {
    response = &CreateSCRTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建敏感内容识别请求的任务
func (c *Client) CreateSCRTask(request *CreateSCRTaskRequest) (response *CreateSCRTaskResponse, err error) {
    if request == nil {
        request = NewCreateSCRTaskRequest()
    }
    response = NewCreateSCRTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDasbCvmConfigRequest() (request *DescribeDasbCvmConfigRequest) {
    request = &DescribeDasbCvmConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbCvmConfig")
    return
}

func NewDescribeDasbCvmConfigResponse() (response *DescribeDasbCvmConfigResponse) {
    response = &DescribeDasbCvmConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询CVM配置信息
func (c *Client) DescribeDasbCvmConfig(request *DescribeDasbCvmConfigRequest) (response *DescribeDasbCvmConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDasbCvmConfigRequest()
    }
    response = NewDescribeDasbCvmConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDasbCvmInstanceRequest() (request *DescribeDasbCvmInstanceRequest) {
    request = &DescribeDasbCvmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbCvmInstance")
    return
}

func NewDescribeDasbCvmInstanceResponse() (response *DescribeDasbCvmInstanceResponse) {
    response = &DescribeDasbCvmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询DASB产品对应的CVM实例信息
func (c *Client) DescribeDasbCvmInstance(request *DescribeDasbCvmInstanceRequest) (response *DescribeDasbCvmInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeDasbCvmInstanceRequest()
    }
    response = NewDescribeDasbCvmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDasbImageIdsRequest() (request *DescribeDasbImageIdsRequest) {
    request = &DescribeDasbImageIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbImageIds")
    return
}

func NewDescribeDasbImageIdsResponse() (response *DescribeDasbImageIdsResponse) {
    response = &DescribeDasbImageIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取镜像列表
func (c *Client) DescribeDasbImageIds(request *DescribeDasbImageIdsRequest) (response *DescribeDasbImageIdsResponse, err error) {
    if request == nil {
        request = NewDescribeDasbImageIdsRequest()
    }
    response = NewDescribeDasbImageIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDasbInstanceAlarmRequest() (request *DescribeDasbInstanceAlarmRequest) {
    request = &DescribeDasbInstanceAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbInstanceAlarm")
    return
}

func NewDescribeDasbInstanceAlarmResponse() (response *DescribeDasbInstanceAlarmResponse) {
    response = &DescribeDasbInstanceAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取DASB实例告警列表
func (c *Client) DescribeDasbInstanceAlarm(request *DescribeDasbInstanceAlarmRequest) (response *DescribeDasbInstanceAlarmResponse, err error) {
    if request == nil {
        request = NewDescribeDasbInstanceAlarmRequest()
    }
    response = NewDescribeDasbInstanceAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDasbServiceRequest() (request *DescribeDasbServiceRequest) {
    request = &DescribeDasbServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbService")
    return
}

func NewDescribeDasbServiceResponse() (response *DescribeDasbServiceResponse) {
    response = &DescribeDasbServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询DASB服务列表
func (c *Client) DescribeDasbService(request *DescribeDasbServiceRequest) (response *DescribeDasbServiceResponse, err error) {
    if request == nil {
        request = NewDescribeDasbServiceRequest()
    }
    response = NewDescribeDasbServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInitialPasswordRequest() (request *DescribeInitialPasswordRequest) {
    request = &DescribeInitialPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeInitialPassword")
    return
}

func NewDescribeInitialPasswordResponse() (response *DescribeInitialPasswordResponse) {
    response = &DescribeInitialPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看初始化密码
func (c *Client) DescribeInitialPassword(request *DescribeInitialPasswordRequest) (response *DescribeInitialPasswordResponse, err error) {
    if request == nil {
        request = NewDescribeInitialPasswordRequest()
    }
    response = NewDescribeInitialPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewGetDasbUserInfosRequest() (request *GetDasbUserInfosRequest) {
    request = &GetDasbUserInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "GetDasbUserInfos")
    return
}

func NewGetDasbUserInfosResponse() (response *GetDasbUserInfosResponse) {
    response = &GetDasbUserInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户信息接口
func (c *Client) GetDasbUserInfos(request *GetDasbUserInfosRequest) (response *GetDasbUserInfosResponse, err error) {
    if request == nil {
        request = NewGetDasbUserInfosRequest()
    }
    response = NewGetDasbUserInfosResponse()
    err = c.Send(request, response)
    return
}

func NewInitDasbInstanceRequest() (request *InitDasbInstanceRequest) {
    request = &InitDasbInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "InitDasbInstance")
    return
}

func NewInitDasbInstanceResponse() (response *InitDasbInstanceResponse) {
    response = &InitDasbInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 初始化用户购买的实例
func (c *Client) InitDasbInstance(request *InitDasbInstanceRequest) (response *InitDasbInstanceResponse, err error) {
    if request == nil {
        request = NewInitDasbInstanceRequest()
    }
    response = NewInitDasbInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRunInstanceRequest() (request *RunInstanceRequest) {
    request = &RunInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "RunInstance")
    return
}

func NewRunInstanceResponse() (response *RunInstanceResponse) {
    response = &RunInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建堡垒机实例
func (c *Client) RunInstance(request *RunInstanceRequest) (response *RunInstanceResponse, err error) {
    if request == nil {
        request = NewRunInstanceRequest()
    }
    response = NewRunInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSetDasbSecretInfosRequest() (request *SetDasbSecretInfosRequest) {
    request = &SetDasbSecretInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "SetDasbSecretInfos")
    return
}

func NewSetDasbSecretInfosResponse() (response *SetDasbSecretInfosResponse) {
    response = &SetDasbSecretInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DASB设置密钥信息接口
func (c *Client) SetDasbSecretInfos(request *SetDasbSecretInfosRequest) (response *SetDasbSecretInfosResponse, err error) {
    if request == nil {
        request = NewSetDasbSecretInfosRequest()
    }
    response = NewSetDasbSecretInfosResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "TerminateInstances")
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 销毁、退还堡垒机实例
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
