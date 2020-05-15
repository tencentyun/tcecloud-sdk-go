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

package v20190924

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-09-24"

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


func NewBatchDeleteImagePersonalRequest() (request *BatchDeleteImagePersonalRequest) {
    request = &BatchDeleteImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "BatchDeleteImagePersonal")
    return
}

func NewBatchDeleteImagePersonalResponse() (response *BatchDeleteImagePersonalResponse) {
    response = &BatchDeleteImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版镜像仓库中批量删除Tag
func (c *Client) BatchDeleteImagePersonal(request *BatchDeleteImagePersonalRequest) (response *BatchDeleteImagePersonalResponse, err error) {
    if request == nil {
        request = NewBatchDeleteImagePersonalRequest()
    }
    response = NewBatchDeleteImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteRepositoryPersonalRequest() (request *BatchDeleteRepositoryPersonalRequest) {
    request = &BatchDeleteRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "BatchDeleteRepositoryPersonal")
    return
}

func NewBatchDeleteRepositoryPersonalResponse() (response *BatchDeleteRepositoryPersonalResponse) {
    response = &BatchDeleteRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于个人版镜像仓库中批量删除镜像仓库
func (c *Client) BatchDeleteRepositoryPersonal(request *BatchDeleteRepositoryPersonalRequest) (response *BatchDeleteRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewBatchDeleteRepositoryPersonalRequest()
    }
    response = NewBatchDeleteRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationTriggerPersonalRequest() (request *CreateApplicationTriggerPersonalRequest) {
    request = &CreateApplicationTriggerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateApplicationTriggerPersonal")
    return
}

func NewCreateApplicationTriggerPersonalResponse() (response *CreateApplicationTriggerPersonalResponse) {
    response = &CreateApplicationTriggerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于创建应用更新触发器
func (c *Client) CreateApplicationTriggerPersonal(request *CreateApplicationTriggerPersonalRequest) (response *CreateApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewCreateApplicationTriggerPersonalRequest()
    }
    response = NewCreateApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageLifecyclePersonalRequest() (request *CreateImageLifecyclePersonalRequest) {
    request = &CreateImageLifecyclePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImageLifecyclePersonal")
    return
}

func NewCreateImageLifecyclePersonalResponse() (response *CreateImageLifecyclePersonalResponse) {
    response = &CreateImageLifecyclePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版中创建清理策略
func (c *Client) CreateImageLifecyclePersonal(request *CreateImageLifecyclePersonalRequest) (response *CreateImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewCreateImageLifecyclePersonalRequest()
    }
    response = NewCreateImageLifecyclePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInstance")
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建实例
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceTokenRequest() (request *CreateInstanceTokenRequest) {
    request = &CreateInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInstanceToken")
    return
}

func NewCreateInstanceTokenResponse() (response *CreateInstanceTokenResponse) {
    response = &CreateInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取临时登录密码
func (c *Client) CreateInstanceToken(request *CreateInstanceTokenRequest) (response *CreateInstanceTokenResponse, err error) {
    if request == nil {
        request = NewCreateInstanceTokenRequest()
    }
    response = NewCreateInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespacePersonalRequest() (request *CreateNamespacePersonalRequest) {
    request = &CreateNamespacePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateNamespacePersonal")
    return
}

func NewCreateNamespacePersonalResponse() (response *CreateNamespacePersonalResponse) {
    response = &CreateNamespacePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建个人版镜像仓库命名空间，此命名空间全局唯一
func (c *Client) CreateNamespacePersonal(request *CreateNamespacePersonalRequest) (response *CreateNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewCreateNamespacePersonalRequest()
    }
    response = NewCreateNamespacePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRepositoryPersonalRequest() (request *CreateRepositoryPersonalRequest) {
    request = &CreateRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateRepositoryPersonal")
    return
}

func NewCreateRepositoryPersonalResponse() (response *CreateRepositoryPersonalResponse) {
    response = &CreateRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版仓库中创建镜像仓库
func (c *Client) CreateRepositoryPersonal(request *CreateRepositoryPersonalRequest) (response *CreateRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryPersonalRequest()
    }
    response = NewCreateRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserPersonalRequest() (request *CreateUserPersonalRequest) {
    request = &CreateUserPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateUserPersonal")
    return
}

func NewCreateUserPersonalResponse() (response *CreateUserPersonalResponse) {
    response = &CreateUserPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建个人用户
func (c *Client) CreateUserPersonal(request *CreateUserPersonalRequest) (response *CreateUserPersonalResponse, err error) {
    if request == nil {
        request = NewCreateUserPersonalRequest()
    }
    response = NewCreateUserPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationTriggerPersonalRequest() (request *DeleteApplicationTriggerPersonalRequest) {
    request = &DeleteApplicationTriggerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteApplicationTriggerPersonal")
    return
}

func NewDeleteApplicationTriggerPersonalResponse() (response *DeleteApplicationTriggerPersonalResponse) {
    response = &DeleteApplicationTriggerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于删除应用更新触发器
func (c *Client) DeleteApplicationTriggerPersonal(request *DeleteApplicationTriggerPersonalRequest) (response *DeleteApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationTriggerPersonalRequest()
    }
    response = NewDeleteApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageLifecycleGlobalPersonalRequest() (request *DeleteImageLifecycleGlobalPersonalRequest) {
    request = &DeleteImageLifecycleGlobalPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImageLifecycleGlobalPersonal")
    return
}

func NewDeleteImageLifecycleGlobalPersonalResponse() (response *DeleteImageLifecycleGlobalPersonalResponse) {
    response = &DeleteImageLifecycleGlobalPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于删除个人版全局镜像版本自动清理策略
func (c *Client) DeleteImageLifecycleGlobalPersonal(request *DeleteImageLifecycleGlobalPersonalRequest) (response *DeleteImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImageLifecycleGlobalPersonalRequest()
    }
    response = NewDeleteImageLifecycleGlobalPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageLifecyclePersonalRequest() (request *DeleteImageLifecyclePersonalRequest) {
    request = &DeleteImageLifecyclePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImageLifecyclePersonal")
    return
}

func NewDeleteImageLifecyclePersonalResponse() (response *DeleteImageLifecyclePersonalResponse) {
    response = &DeleteImageLifecyclePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版镜像仓库中删除仓库Tag自动清理策略
func (c *Client) DeleteImageLifecyclePersonal(request *DeleteImageLifecyclePersonalRequest) (response *DeleteImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImageLifecyclePersonalRequest()
    }
    response = NewDeleteImageLifecyclePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImagePersonalRequest() (request *DeleteImagePersonalRequest) {
    request = &DeleteImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImagePersonal")
    return
}

func NewDeleteImagePersonalResponse() (response *DeleteImagePersonalResponse) {
    response = &DeleteImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版中删除tag
func (c *Client) DeleteImagePersonal(request *DeleteImagePersonalRequest) (response *DeleteImagePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImagePersonalRequest()
    }
    response = NewDeleteImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespacePersonalRequest() (request *DeleteNamespacePersonalRequest) {
    request = &DeleteNamespacePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteNamespacePersonal")
    return
}

func NewDeleteNamespacePersonalResponse() (response *DeleteNamespacePersonalResponse) {
    response = &DeleteNamespacePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除共享版命名空间
func (c *Client) DeleteNamespacePersonal(request *DeleteNamespacePersonalRequest) (response *DeleteNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteNamespacePersonalRequest()
    }
    response = NewDeleteNamespacePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryPersonalRequest() (request *DeleteRepositoryPersonalRequest) {
    request = &DeleteRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteRepositoryPersonal")
    return
}

func NewDeleteRepositoryPersonalResponse() (response *DeleteRepositoryPersonalResponse) {
    response = &DeleteRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于个人版镜像仓库中删除
func (c *Client) DeleteRepositoryPersonal(request *DeleteRepositoryPersonalRequest) (response *DeleteRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryPersonalRequest()
    }
    response = NewDeleteRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationTriggerLogPersonalRequest() (request *DescribeApplicationTriggerLogPersonalRequest) {
    request = &DescribeApplicationTriggerLogPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeApplicationTriggerLogPersonal")
    return
}

func NewDescribeApplicationTriggerLogPersonalResponse() (response *DescribeApplicationTriggerLogPersonalResponse) {
    response = &DescribeApplicationTriggerLogPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询应用更新触发器触发日志
func (c *Client) DescribeApplicationTriggerLogPersonal(request *DescribeApplicationTriggerLogPersonalRequest) (response *DescribeApplicationTriggerLogPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationTriggerLogPersonalRequest()
    }
    response = NewDescribeApplicationTriggerLogPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationTriggerPersonalRequest() (request *DescribeApplicationTriggerPersonalRequest) {
    request = &DescribeApplicationTriggerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeApplicationTriggerPersonal")
    return
}

func NewDescribeApplicationTriggerPersonalResponse() (response *DescribeApplicationTriggerPersonalResponse) {
    response = &DescribeApplicationTriggerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询应用更新触发器
func (c *Client) DescribeApplicationTriggerPersonal(request *DescribeApplicationTriggerPersonalRequest) (response *DescribeApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationTriggerPersonalRequest()
    }
    response = NewDescribeApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFavorRepositoryPersonalRequest() (request *DescribeFavorRepositoryPersonalRequest) {
    request = &DescribeFavorRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeFavorRepositoryPersonal")
    return
}

func NewDescribeFavorRepositoryPersonalResponse() (response *DescribeFavorRepositoryPersonalResponse) {
    response = &DescribeFavorRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询个人收藏仓库
func (c *Client) DescribeFavorRepositoryPersonal(request *DescribeFavorRepositoryPersonalRequest) (response *DescribeFavorRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeFavorRepositoryPersonalRequest()
    }
    response = NewDescribeFavorRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageFilterPersonalRequest() (request *DescribeImageFilterPersonalRequest) {
    request = &DescribeImageFilterPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageFilterPersonal")
    return
}

func NewDescribeImageFilterPersonalResponse() (response *DescribeImageFilterPersonalResponse) {
    response = &DescribeImageFilterPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版中查询与指定tag镜像内容相同的tag列表
func (c *Client) DescribeImageFilterPersonal(request *DescribeImageFilterPersonalRequest) (response *DescribeImageFilterPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageFilterPersonalRequest()
    }
    response = NewDescribeImageFilterPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageLifecycleGlobalPersonalRequest() (request *DescribeImageLifecycleGlobalPersonalRequest) {
    request = &DescribeImageLifecycleGlobalPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageLifecycleGlobalPersonal")
    return
}

func NewDescribeImageLifecycleGlobalPersonalResponse() (response *DescribeImageLifecycleGlobalPersonalResponse) {
    response = &DescribeImageLifecycleGlobalPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于获取个人版全局镜像版本自动清理策略
func (c *Client) DescribeImageLifecycleGlobalPersonal(request *DescribeImageLifecycleGlobalPersonalRequest) (response *DescribeImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageLifecycleGlobalPersonalRequest()
    }
    response = NewDescribeImageLifecycleGlobalPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageLifecyclePersonalRequest() (request *DescribeImageLifecyclePersonalRequest) {
    request = &DescribeImageLifecyclePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageLifecyclePersonal")
    return
}

func NewDescribeImageLifecyclePersonalResponse() (response *DescribeImageLifecyclePersonalResponse) {
    response = &DescribeImageLifecyclePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于获取个人版仓库中自动清理策略
func (c *Client) DescribeImageLifecyclePersonal(request *DescribeImageLifecyclePersonalRequest) (response *DescribeImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageLifecyclePersonalRequest()
    }
    response = NewDescribeImageLifecyclePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagePersonalRequest() (request *DescribeImagePersonalRequest) {
    request = &DescribeImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImagePersonal")
    return
}

func NewDescribeImagePersonalResponse() (response *DescribeImagePersonalResponse) {
    response = &DescribeImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于获取个人版镜像仓库tag列表
func (c *Client) DescribeImagePersonal(request *DescribeImagePersonalRequest) (response *DescribeImagePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImagePersonalRequest()
    }
    response = NewDescribeImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceStatusRequest() (request *DescribeInstanceStatusRequest) {
    request = &DescribeInstanceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceStatus")
    return
}

func NewDescribeInstanceStatusResponse() (response *DescribeInstanceStatusResponse) {
    response = &DescribeInstanceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例当前状态以及过程信息
func (c *Client) DescribeInstanceStatus(request *DescribeInstanceStatusRequest) (response *DescribeInstanceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceStatusRequest()
    }
    response = NewDescribeInstanceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例信息
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespacePersonalRequest() (request *DescribeNamespacePersonalRequest) {
    request = &DescribeNamespacePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeNamespacePersonal")
    return
}

func NewDescribeNamespacePersonalResponse() (response *DescribeNamespacePersonalResponse) {
    response = &DescribeNamespacePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询个人版命名空间信息
func (c *Client) DescribeNamespacePersonal(request *DescribeNamespacePersonalRequest) (response *DescribeNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacePersonalRequest()
    }
    response = NewDescribeNamespacePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryFilterPersonalRequest() (request *DescribeRepositoryFilterPersonalRequest) {
    request = &DescribeRepositoryFilterPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositoryFilterPersonal")
    return
}

func NewDescribeRepositoryFilterPersonalResponse() (response *DescribeRepositoryFilterPersonalResponse) {
    response = &DescribeRepositoryFilterPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版镜像仓库中，获取满足输入搜索条件的用户镜像仓库
func (c *Client) DescribeRepositoryFilterPersonal(request *DescribeRepositoryFilterPersonalRequest) (response *DescribeRepositoryFilterPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryFilterPersonalRequest()
    }
    response = NewDescribeRepositoryFilterPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryOwnerPersonalRequest() (request *DescribeRepositoryOwnerPersonalRequest) {
    request = &DescribeRepositoryOwnerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositoryOwnerPersonal")
    return
}

func NewDescribeRepositoryOwnerPersonalResponse() (response *DescribeRepositoryOwnerPersonalResponse) {
    response = &DescribeRepositoryOwnerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版中获取用户全部的镜像仓库列表
func (c *Client) DescribeRepositoryOwnerPersonal(request *DescribeRepositoryOwnerPersonalRequest) (response *DescribeRepositoryOwnerPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryOwnerPersonalRequest()
    }
    response = NewDescribeRepositoryOwnerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryPersonalRequest() (request *DescribeRepositoryPersonalRequest) {
    request = &DescribeRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositoryPersonal")
    return
}

func NewDescribeRepositoryPersonalResponse() (response *DescribeRepositoryPersonalResponse) {
    response = &DescribeRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询个人版仓库信息
func (c *Client) DescribeRepositoryPersonal(request *DescribeRepositoryPersonalRequest) (response *DescribeRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryPersonalRequest()
    }
    response = NewDescribeRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserQuotaPersonalRequest() (request *DescribeUserQuotaPersonalRequest) {
    request = &DescribeUserQuotaPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeUserQuotaPersonal")
    return
}

func NewDescribeUserQuotaPersonalResponse() (response *DescribeUserQuotaPersonalResponse) {
    response = &DescribeUserQuotaPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询个人用户配额
func (c *Client) DescribeUserQuotaPersonal(request *DescribeUserQuotaPersonalRequest) (response *DescribeUserQuotaPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeUserQuotaPersonalRequest()
    }
    response = NewDescribeUserQuotaPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDuplicateImagePersonalRequest() (request *DuplicateImagePersonalRequest) {
    request = &DuplicateImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DuplicateImagePersonal")
    return
}

func NewDuplicateImagePersonalResponse() (response *DuplicateImagePersonalResponse) {
    response = &DuplicateImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版镜像仓库中复制镜像版本
func (c *Client) DuplicateImagePersonal(request *DuplicateImagePersonalRequest) (response *DuplicateImagePersonalResponse, err error) {
    if request == nil {
        request = NewDuplicateImagePersonalRequest()
    }
    response = NewDuplicateImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewManageImageLifecycleGlobalPersonalRequest() (request *ManageImageLifecycleGlobalPersonalRequest) {
    request = &ManageImageLifecycleGlobalPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ManageImageLifecycleGlobalPersonal")
    return
}

func NewManageImageLifecycleGlobalPersonalResponse() (response *ManageImageLifecycleGlobalPersonalResponse) {
    response = &ManageImageLifecycleGlobalPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于设置个人版全局镜像版本自动清理策略
func (c *Client) ManageImageLifecycleGlobalPersonal(request *ManageImageLifecycleGlobalPersonalRequest) (response *ManageImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewManageImageLifecycleGlobalPersonalRequest()
    }
    response = NewManageImageLifecycleGlobalPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationTriggerPersonalRequest() (request *ModifyApplicationTriggerPersonalRequest) {
    request = &ModifyApplicationTriggerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyApplicationTriggerPersonal")
    return
}

func NewModifyApplicationTriggerPersonalResponse() (response *ModifyApplicationTriggerPersonalResponse) {
    response = &ModifyApplicationTriggerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于修改应用更新触发器
func (c *Client) ModifyApplicationTriggerPersonal(request *ModifyApplicationTriggerPersonalRequest) (response *ModifyApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewModifyApplicationTriggerPersonalRequest()
    }
    response = NewModifyApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRepositoryAccessPersonalRequest() (request *ModifyRepositoryAccessPersonalRequest) {
    request = &ModifyRepositoryAccessPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyRepositoryAccessPersonal")
    return
}

func NewModifyRepositoryAccessPersonalResponse() (response *ModifyRepositoryAccessPersonalResponse) {
    response = &ModifyRepositoryAccessPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于更新个人版镜像仓库的访问属性
func (c *Client) ModifyRepositoryAccessPersonal(request *ModifyRepositoryAccessPersonalRequest) (response *ModifyRepositoryAccessPersonalResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryAccessPersonalRequest()
    }
    response = NewModifyRepositoryAccessPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRepositoryInfoPersonalRequest() (request *ModifyRepositoryInfoPersonalRequest) {
    request = &ModifyRepositoryInfoPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyRepositoryInfoPersonal")
    return
}

func NewModifyRepositoryInfoPersonalResponse() (response *ModifyRepositoryInfoPersonalResponse) {
    response = &ModifyRepositoryInfoPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在个人版镜像仓库中更新容器镜像描述
func (c *Client) ModifyRepositoryInfoPersonal(request *ModifyRepositoryInfoPersonalRequest) (response *ModifyRepositoryInfoPersonalResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryInfoPersonalRequest()
    }
    response = NewModifyRepositoryInfoPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserPasswordPersonalRequest() (request *ModifyUserPasswordPersonalRequest) {
    request = &ModifyUserPasswordPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyUserPasswordPersonal")
    return
}

func NewModifyUserPasswordPersonalResponse() (response *ModifyUserPasswordPersonalResponse) {
    response = &ModifyUserPasswordPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改个人用户登录密码
func (c *Client) ModifyUserPasswordPersonal(request *ModifyUserPasswordPersonalRequest) (response *ModifyUserPasswordPersonalResponse, err error) {
    if request == nil {
        request = NewModifyUserPasswordPersonalRequest()
    }
    response = NewModifyUserPasswordPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewValidateNamespaceExistPersonalRequest() (request *ValidateNamespaceExistPersonalRequest) {
    request = &ValidateNamespaceExistPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ValidateNamespaceExistPersonal")
    return
}

func NewValidateNamespaceExistPersonalResponse() (response *ValidateNamespaceExistPersonalResponse) {
    response = &ValidateNamespaceExistPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询个人版用户命名空间是否存在
func (c *Client) ValidateNamespaceExistPersonal(request *ValidateNamespaceExistPersonalRequest) (response *ValidateNamespaceExistPersonalResponse, err error) {
    if request == nil {
        request = NewValidateNamespaceExistPersonalRequest()
    }
    response = NewValidateNamespaceExistPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewValidateRepositoryExistPersonalRequest() (request *ValidateRepositoryExistPersonalRequest) {
    request = &ValidateRepositoryExistPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ValidateRepositoryExistPersonal")
    return
}

func NewValidateRepositoryExistPersonalResponse() (response *ValidateRepositoryExistPersonalResponse) {
    response = &ValidateRepositoryExistPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于判断个人版仓库是否存在
func (c *Client) ValidateRepositoryExistPersonal(request *ValidateRepositoryExistPersonalRequest) (response *ValidateRepositoryExistPersonalResponse, err error) {
    if request == nil {
        request = NewValidateRepositoryExistPersonalRequest()
    }
    response = NewValidateRepositoryExistPersonalResponse()
    err = c.Send(request, response)
    return
}
