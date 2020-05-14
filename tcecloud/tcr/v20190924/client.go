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
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
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


func NewBatchDeleteFavorRepositoryPersonalRequest() (request *BatchDeleteFavorRepositoryPersonalRequest) {
    request = &BatchDeleteFavorRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "BatchDeleteFavorRepositoryPersonal")
    return
}

func NewBatchDeleteFavorRepositoryPersonalResponse() (response *BatchDeleteFavorRepositoryPersonalResponse) {
    response = &BatchDeleteFavorRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除个人收藏仓库
func (c *Client) BatchDeleteFavorRepositoryPersonal(request *BatchDeleteFavorRepositoryPersonalRequest) (response *BatchDeleteFavorRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewBatchDeleteFavorRepositoryPersonalRequest()
    }
    response = NewBatchDeleteFavorRepositoryPersonalResponse()
    err = c.Send(request, response)
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

func NewCreateApplicationTokenPersonalRequest() (request *CreateApplicationTokenPersonalRequest) {
    request = &CreateApplicationTokenPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateApplicationTokenPersonal")
    return
}

func NewCreateApplicationTokenPersonalResponse() (response *CreateApplicationTokenPersonalResponse) {
    response = &CreateApplicationTokenPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于创建第三方应用访问凭证
func (c *Client) CreateApplicationTokenPersonal(request *CreateApplicationTokenPersonalRequest) (response *CreateApplicationTokenPersonalResponse, err error) {
    if request == nil {
        request = NewCreateApplicationTokenPersonalRequest()
    }
    response = NewCreateApplicationTokenPersonalResponse()
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

func NewCreateFavorRepositoryPersonalRequest() (request *CreateFavorRepositoryPersonalRequest) {
    request = &CreateFavorRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateFavorRepositoryPersonal")
    return
}

func NewCreateFavorRepositoryPersonalResponse() (response *CreateFavorRepositoryPersonalResponse) {
    response = &CreateFavorRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建个人收藏仓库
func (c *Client) CreateFavorRepositoryPersonal(request *CreateFavorRepositoryPersonalRequest) (response *CreateFavorRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewCreateFavorRepositoryPersonalRequest()
    }
    response = NewCreateFavorRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageBuildPersonalRequest() (request *CreateImageBuildPersonalRequest) {
    request = &CreateImageBuildPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImageBuildPersonal")
    return
}

func NewCreateImageBuildPersonalResponse() (response *CreateImageBuildPersonalResponse) {
    response = &CreateImageBuildPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建镜像构建规则
func (c *Client) CreateImageBuildPersonal(request *CreateImageBuildPersonalRequest) (response *CreateImageBuildPersonalResponse, err error) {
    if request == nil {
        request = NewCreateImageBuildPersonalRequest()
    }
    response = NewCreateImageBuildPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageBuildTaskDockerPersonalRequest() (request *CreateImageBuildTaskDockerPersonalRequest) {
    request = &CreateImageBuildTaskDockerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImageBuildTaskDockerPersonal")
    return
}

func NewCreateImageBuildTaskDockerPersonalResponse() (response *CreateImageBuildTaskDockerPersonalResponse) {
    response = &CreateImageBuildTaskDockerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建基于Dockerfile的镜像构建任务
func (c *Client) CreateImageBuildTaskDockerPersonal(request *CreateImageBuildTaskDockerPersonalRequest) (response *CreateImageBuildTaskDockerPersonalResponse, err error) {
    if request == nil {
        request = NewCreateImageBuildTaskDockerPersonalRequest()
    }
    response = NewCreateImageBuildTaskDockerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageBuildTaskManuallyPersonalRequest() (request *CreateImageBuildTaskManuallyPersonalRequest) {
    request = &CreateImageBuildTaskManuallyPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImageBuildTaskManuallyPersonal")
    return
}

func NewCreateImageBuildTaskManuallyPersonalResponse() (response *CreateImageBuildTaskManuallyPersonalResponse) {
    response = &CreateImageBuildTaskManuallyPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建手动执行的镜像构建任务
func (c *Client) CreateImageBuildTaskManuallyPersonal(request *CreateImageBuildTaskManuallyPersonalRequest) (response *CreateImageBuildTaskManuallyPersonalResponse, err error) {
    if request == nil {
        request = NewCreateImageBuildTaskManuallyPersonalRequest()
    }
    response = NewCreateImageBuildTaskManuallyPersonalResponse()
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

func NewCreateRepoRequest() (request *CreateRepoRequest) {
    request = &CreateRepoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateRepo")
    return
}

func NewCreateRepoResponse() (response *CreateRepoResponse) {
    response = &CreateRepoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在共享版仓库中创建镜像仓库
func (c *Client) CreateRepo(request *CreateRepoRequest) (response *CreateRepoResponse, err error) {
    if request == nil {
        request = NewCreateRepoRequest()
    }
    response = NewCreateRepoResponse()
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

func NewCreateSecurityPoliciesRequest() (request *CreateSecurityPoliciesRequest) {
    request = &CreateSecurityPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSecurityPolicies")
    return
}

func NewCreateSecurityPoliciesResponse() (response *CreateSecurityPoliciesResponse) {
    response = &CreateSecurityPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建实例公网访问白名单策略
func (c *Client) CreateSecurityPolicies(request *CreateSecurityPoliciesRequest) (response *CreateSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPoliciesRequest()
    }
    response = NewCreateSecurityPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityPolicyRequest() (request *CreateSecurityPolicyRequest) {
    request = &CreateSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSecurityPolicy")
    return
}

func NewCreateSecurityPolicyResponse() (response *CreateSecurityPolicyResponse) {
    response = &CreateSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建实例公网访问白名单策略
func (c *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPolicyRequest()
    }
    response = NewCreateSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSourceCodeAuthPersonalRequest() (request *CreateSourceCodeAuthPersonalRequest) {
    request = &CreateSourceCodeAuthPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSourceCodeAuthPersonal")
    return
}

func NewCreateSourceCodeAuthPersonalResponse() (response *CreateSourceCodeAuthPersonalResponse) {
    response = &CreateSourceCodeAuthPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建源代码授权
func (c *Client) CreateSourceCodeAuthPersonal(request *CreateSourceCodeAuthPersonalRequest) (response *CreateSourceCodeAuthPersonalResponse, err error) {
    if request == nil {
        request = NewCreateSourceCodeAuthPersonalRequest()
    }
    response = NewCreateSourceCodeAuthPersonalResponse()
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

func NewDeleteFavorRepositoryPersonalRequest() (request *DeleteFavorRepositoryPersonalRequest) {
    request = &DeleteFavorRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteFavorRepositoryPersonal")
    return
}

func NewDeleteFavorRepositoryPersonalResponse() (response *DeleteFavorRepositoryPersonalResponse) {
    response = &DeleteFavorRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除个人收藏仓库
func (c *Client) DeleteFavorRepositoryPersonal(request *DeleteFavorRepositoryPersonalRequest) (response *DeleteFavorRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteFavorRepositoryPersonalRequest()
    }
    response = NewDeleteFavorRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFavorRepositoryRegionPersonalRequest() (request *DeleteFavorRepositoryRegionPersonalRequest) {
    request = &DeleteFavorRepositoryRegionPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteFavorRepositoryRegionPersonal")
    return
}

func NewDeleteFavorRepositoryRegionPersonalResponse() (response *DeleteFavorRepositoryRegionPersonalResponse) {
    response = &DeleteFavorRepositoryRegionPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除指定地域所有个人收藏仓库
func (c *Client) DeleteFavorRepositoryRegionPersonal(request *DeleteFavorRepositoryRegionPersonalRequest) (response *DeleteFavorRepositoryRegionPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteFavorRepositoryRegionPersonalRequest()
    }
    response = NewDeleteFavorRepositoryRegionPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageBuildPersonalRequest() (request *DeleteImageBuildPersonalRequest) {
    request = &DeleteImageBuildPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImageBuildPersonal")
    return
}

func NewDeleteImageBuildPersonalResponse() (response *DeleteImageBuildPersonalResponse) {
    response = &DeleteImageBuildPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除镜像构建规则
func (c *Client) DeleteImageBuildPersonal(request *DeleteImageBuildPersonalRequest) (response *DeleteImageBuildPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImageBuildPersonalRequest()
    }
    response = NewDeleteImageBuildPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageBuildTaskPersonalRequest() (request *DeleteImageBuildTaskPersonalRequest) {
    request = &DeleteImageBuildTaskPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImageBuildTaskPersonal")
    return
}

func NewDeleteImageBuildTaskPersonalResponse() (response *DeleteImageBuildTaskPersonalResponse) {
    response = &DeleteImageBuildTaskPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除镜像构建任务
func (c *Client) DeleteImageBuildTaskPersonal(request *DeleteImageBuildTaskPersonalRequest) (response *DeleteImageBuildTaskPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImageBuildTaskPersonalRequest()
    }
    response = NewDeleteImageBuildTaskPersonalResponse()
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

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteInstance")
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除镜像仓库企业版实例
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    response = NewDeleteInstanceResponse()
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

func NewDeleteSecurityPolicyRequest() (request *DeleteSecurityPolicyRequest) {
    request = &DeleteSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteSecurityPolicy")
    return
}

func NewDeleteSecurityPolicyResponse() (response *DeleteSecurityPolicyResponse) {
    response = &DeleteSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除实例公网访问白名单策略
func (c *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityPolicyRequest()
    }
    response = NewDeleteSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSourceCodeAuthPersonalRequest() (request *DeleteSourceCodeAuthPersonalRequest) {
    request = &DeleteSourceCodeAuthPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteSourceCodeAuthPersonal")
    return
}

func NewDeleteSourceCodeAuthPersonalResponse() (response *DeleteSourceCodeAuthPersonalResponse) {
    response = &DeleteSourceCodeAuthPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除源代码授权
func (c *Client) DeleteSourceCodeAuthPersonal(request *DeleteSourceCodeAuthPersonalRequest) (response *DeleteSourceCodeAuthPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteSourceCodeAuthPersonalRequest()
    }
    response = NewDeleteSourceCodeAuthPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationTokenPersonalRequest() (request *DescribeApplicationTokenPersonalRequest) {
    request = &DescribeApplicationTokenPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeApplicationTokenPersonal")
    return
}

func NewDescribeApplicationTokenPersonalResponse() (response *DescribeApplicationTokenPersonalResponse) {
    response = &DescribeApplicationTokenPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于获取第三方应用访问凭证
func (c *Client) DescribeApplicationTokenPersonal(request *DescribeApplicationTokenPersonalRequest) (response *DescribeApplicationTokenPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationTokenPersonalRequest()
    }
    response = NewDescribeApplicationTokenPersonalResponse()
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

func NewDescribeDockerHubImagePersonalRequest() (request *DescribeDockerHubImagePersonalRequest) {
    request = &DescribeDockerHubImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeDockerHubImagePersonal")
    return
}

func NewDescribeDockerHubImagePersonalResponse() (response *DescribeDockerHubImagePersonalResponse) {
    response = &DescribeDockerHubImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询DockerHub镜像列表
func (c *Client) DescribeDockerHubImagePersonal(request *DescribeDockerHubImagePersonalRequest) (response *DescribeDockerHubImagePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeDockerHubImagePersonalRequest()
    }
    response = NewDescribeDockerHubImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDockerHubRepositoryInfoPersonalRequest() (request *DescribeDockerHubRepositoryInfoPersonalRequest) {
    request = &DescribeDockerHubRepositoryInfoPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeDockerHubRepositoryInfoPersonal")
    return
}

func NewDescribeDockerHubRepositoryInfoPersonalResponse() (response *DescribeDockerHubRepositoryInfoPersonalResponse) {
    response = &DescribeDockerHubRepositoryInfoPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询DockerHub仓库信息
func (c *Client) DescribeDockerHubRepositoryInfoPersonal(request *DescribeDockerHubRepositoryInfoPersonalRequest) (response *DescribeDockerHubRepositoryInfoPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeDockerHubRepositoryInfoPersonalRequest()
    }
    response = NewDescribeDockerHubRepositoryInfoPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDockerHubRepositoryPersonalRequest() (request *DescribeDockerHubRepositoryPersonalRequest) {
    request = &DescribeDockerHubRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeDockerHubRepositoryPersonal")
    return
}

func NewDescribeDockerHubRepositoryPersonalResponse() (response *DescribeDockerHubRepositoryPersonalResponse) {
    response = &DescribeDockerHubRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询DockerHub仓库列表
func (c *Client) DescribeDockerHubRepositoryPersonal(request *DescribeDockerHubRepositoryPersonalRequest) (response *DescribeDockerHubRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeDockerHubRepositoryPersonalRequest()
    }
    response = NewDescribeDockerHubRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalEndpointStatusRequest() (request *DescribeExternalEndpointStatusRequest) {
    request = &DescribeExternalEndpointStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeExternalEndpointStatus")
    return
}

func NewDescribeExternalEndpointStatusResponse() (response *DescribeExternalEndpointStatusResponse) {
    response = &DescribeExternalEndpointStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例公网访问入口状态
func (c *Client) DescribeExternalEndpointStatus(request *DescribeExternalEndpointStatusRequest) (response *DescribeExternalEndpointStatusResponse, err error) {
    if request == nil {
        request = NewDescribeExternalEndpointStatusRequest()
    }
    response = NewDescribeExternalEndpointStatusResponse()
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

func NewDescribeImageBuildPersonalRequest() (request *DescribeImageBuildPersonalRequest) {
    request = &DescribeImageBuildPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageBuildPersonal")
    return
}

func NewDescribeImageBuildPersonalResponse() (response *DescribeImageBuildPersonalResponse) {
    response = &DescribeImageBuildPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询镜像构建规则
func (c *Client) DescribeImageBuildPersonal(request *DescribeImageBuildPersonalRequest) (response *DescribeImageBuildPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageBuildPersonalRequest()
    }
    response = NewDescribeImageBuildPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageBuildTaskLogInfoPersonalRequest() (request *DescribeImageBuildTaskLogInfoPersonalRequest) {
    request = &DescribeImageBuildTaskLogInfoPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageBuildTaskLogInfoPersonal")
    return
}

func NewDescribeImageBuildTaskLogInfoPersonalResponse() (response *DescribeImageBuildTaskLogInfoPersonalResponse) {
    response = &DescribeImageBuildTaskLogInfoPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询镜像构建任务日志信息
func (c *Client) DescribeImageBuildTaskLogInfoPersonal(request *DescribeImageBuildTaskLogInfoPersonalRequest) (response *DescribeImageBuildTaskLogInfoPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageBuildTaskLogInfoPersonalRequest()
    }
    response = NewDescribeImageBuildTaskLogInfoPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageBuildTaskLogPersonalRequest() (request *DescribeImageBuildTaskLogPersonalRequest) {
    request = &DescribeImageBuildTaskLogPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageBuildTaskLogPersonal")
    return
}

func NewDescribeImageBuildTaskLogPersonalResponse() (response *DescribeImageBuildTaskLogPersonalResponse) {
    response = &DescribeImageBuildTaskLogPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询镜像构建任务日志
func (c *Client) DescribeImageBuildTaskLogPersonal(request *DescribeImageBuildTaskLogPersonalRequest) (response *DescribeImageBuildTaskLogPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageBuildTaskLogPersonalRequest()
    }
    response = NewDescribeImageBuildTaskLogPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageConfigPersonalRequest() (request *DescribeImageConfigPersonalRequest) {
    request = &DescribeImageConfigPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageConfigPersonal")
    return
}

func NewDescribeImageConfigPersonalResponse() (response *DescribeImageConfigPersonalResponse) {
    response = &DescribeImageConfigPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询镜像版本配置信息
func (c *Client) DescribeImageConfigPersonal(request *DescribeImageConfigPersonalRequest) (response *DescribeImageConfigPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageConfigPersonalRequest()
    }
    response = NewDescribeImageConfigPersonalResponse()
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

func NewDescribeInternalEndpointsRequest() (request *DescribeInternalEndpointsRequest) {
    request = &DescribeInternalEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInternalEndpoints")
    return
}

func NewDescribeInternalEndpointsResponse() (response *DescribeInternalEndpointsResponse) {
    response = &DescribeInternalEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例内网访问VPC链接
func (c *Client) DescribeInternalEndpoints(request *DescribeInternalEndpointsRequest) (response *DescribeInternalEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeInternalEndpointsRequest()
    }
    response = NewDescribeInternalEndpointsResponse()
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

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于在TCR中获取可用区域
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryAllPersonalRequest() (request *DescribeRepositoryAllPersonalRequest) {
    request = &DescribeRepositoryAllPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositoryAllPersonal")
    return
}

func NewDescribeRepositoryAllPersonalResponse() (response *DescribeRepositoryAllPersonalResponse) {
    response = &DescribeRepositoryAllPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询所有可访问的镜像仓库
func (c *Client) DescribeRepositoryAllPersonal(request *DescribeRepositoryAllPersonalRequest) (response *DescribeRepositoryAllPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryAllPersonalRequest()
    }
    response = NewDescribeRepositoryAllPersonalResponse()
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

func NewDescribeSecurityPoliciesRequest() (request *DescribeSecurityPoliciesRequest) {
    request = &DescribeSecurityPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeSecurityPolicies")
    return
}

func NewDescribeSecurityPoliciesResponse() (response *DescribeSecurityPoliciesResponse) {
    response = &DescribeSecurityPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例公网访问白名单策略
func (c *Client) DescribeSecurityPolicies(request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPoliciesRequest()
    }
    response = NewDescribeSecurityPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceCodeAuthPersonalRequest() (request *DescribeSourceCodeAuthPersonalRequest) {
    request = &DescribeSourceCodeAuthPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeSourceCodeAuthPersonal")
    return
}

func NewDescribeSourceCodeAuthPersonalResponse() (response *DescribeSourceCodeAuthPersonalResponse) {
    response = &DescribeSourceCodeAuthPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询源代码授权
func (c *Client) DescribeSourceCodeAuthPersonal(request *DescribeSourceCodeAuthPersonalRequest) (response *DescribeSourceCodeAuthPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeSourceCodeAuthPersonalRequest()
    }
    response = NewDescribeSourceCodeAuthPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceCodeAuthUserInfoPersonalRequest() (request *DescribeSourceCodeAuthUserInfoPersonalRequest) {
    request = &DescribeSourceCodeAuthUserInfoPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeSourceCodeAuthUserInfoPersonal")
    return
}

func NewDescribeSourceCodeAuthUserInfoPersonalResponse() (response *DescribeSourceCodeAuthUserInfoPersonalResponse) {
    response = &DescribeSourceCodeAuthUserInfoPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询源代码授权用户信息
func (c *Client) DescribeSourceCodeAuthUserInfoPersonal(request *DescribeSourceCodeAuthUserInfoPersonalRequest) (response *DescribeSourceCodeAuthUserInfoPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeSourceCodeAuthUserInfoPersonalRequest()
    }
    response = NewDescribeSourceCodeAuthUserInfoPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceCodeRepositoryBranchPersonalRequest() (request *DescribeSourceCodeRepositoryBranchPersonalRequest) {
    request = &DescribeSourceCodeRepositoryBranchPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeSourceCodeRepositoryBranchPersonal")
    return
}

func NewDescribeSourceCodeRepositoryBranchPersonalResponse() (response *DescribeSourceCodeRepositoryBranchPersonalResponse) {
    response = &DescribeSourceCodeRepositoryBranchPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询源代码仓库分支列表
func (c *Client) DescribeSourceCodeRepositoryBranchPersonal(request *DescribeSourceCodeRepositoryBranchPersonalRequest) (response *DescribeSourceCodeRepositoryBranchPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeSourceCodeRepositoryBranchPersonalRequest()
    }
    response = NewDescribeSourceCodeRepositoryBranchPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceCodeRepositoryPersonalRequest() (request *DescribeSourceCodeRepositoryPersonalRequest) {
    request = &DescribeSourceCodeRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeSourceCodeRepositoryPersonal")
    return
}

func NewDescribeSourceCodeRepositoryPersonalResponse() (response *DescribeSourceCodeRepositoryPersonalResponse) {
    response = &DescribeSourceCodeRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询源代码仓库列表
func (c *Client) DescribeSourceCodeRepositoryPersonal(request *DescribeSourceCodeRepositoryPersonalRequest) (response *DescribeSourceCodeRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeSourceCodeRepositoryPersonalRequest()
    }
    response = NewDescribeSourceCodeRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserPersonalRequest() (request *DescribeUserPersonalRequest) {
    request = &DescribeUserPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeUserPersonal")
    return
}

func NewDescribeUserPersonalResponse() (response *DescribeUserPersonalResponse) {
    response = &DescribeUserPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询个人用户信息
func (c *Client) DescribeUserPersonal(request *DescribeUserPersonalRequest) (response *DescribeUserPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeUserPersonalRequest()
    }
    response = NewDescribeUserPersonalResponse()
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

func NewForwardRequestRequest() (request *ForwardRequestRequest) {
    request = &ForwardRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ForwardRequest")
    return
}

func NewForwardRequestResponse() (response *ForwardRequestResponse) {
    response = &ForwardRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TCR 代理转发接口
func (c *Client) ForwardRequest(request *ForwardRequestRequest) (response *ForwardRequestResponse, err error) {
    if request == nil {
        request = NewForwardRequestRequest()
    }
    response = NewForwardRequestResponse()
    err = c.Send(request, response)
    return
}

func NewManageExternalEndpointRequest() (request *ManageExternalEndpointRequest) {
    request = &ManageExternalEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ManageExternalEndpoint")
    return
}

func NewManageExternalEndpointResponse() (response *ManageExternalEndpointResponse) {
    response = &ManageExternalEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理实例公网访问
func (c *Client) ManageExternalEndpoint(request *ManageExternalEndpointRequest) (response *ManageExternalEndpointResponse, err error) {
    if request == nil {
        request = NewManageExternalEndpointRequest()
    }
    response = NewManageExternalEndpointResponse()
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

func NewManageInternalEndpointRequest() (request *ManageInternalEndpointRequest) {
    request = &ManageInternalEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ManageInternalEndpoint")
    return
}

func NewManageInternalEndpointResponse() (response *ManageInternalEndpointResponse) {
    response = &ManageInternalEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理实例内网访问VPC链接
func (c *Client) ManageInternalEndpoint(request *ManageInternalEndpointRequest) (response *ManageInternalEndpointResponse, err error) {
    if request == nil {
        request = NewManageInternalEndpointRequest()
    }
    response = NewManageInternalEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewManageReplicationRequest() (request *ManageReplicationRequest) {
    request = &ManageReplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ManageReplication")
    return
}

func NewManageReplicationResponse() (response *ManageReplicationResponse) {
    response = &ManageReplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理实例同步
func (c *Client) ManageReplication(request *ManageReplicationRequest) (response *ManageReplicationResponse, err error) {
    if request == nil {
        request = NewManageReplicationRequest()
    }
    response = NewManageReplicationResponse()
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

func NewModifyImageBuildPersonalRequest() (request *ModifyImageBuildPersonalRequest) {
    request = &ModifyImageBuildPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyImageBuildPersonal")
    return
}

func NewModifyImageBuildPersonalResponse() (response *ModifyImageBuildPersonalResponse) {
    response = &ModifyImageBuildPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改镜像构建规则
func (c *Client) ModifyImageBuildPersonal(request *ModifyImageBuildPersonalRequest) (response *ModifyImageBuildPersonalResponse, err error) {
    if request == nil {
        request = NewModifyImageBuildPersonalRequest()
    }
    response = NewModifyImageBuildPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyInstance")
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新实例信息
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    response = NewModifyInstanceResponse()
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

func NewModifySecurityPolicyRequest() (request *ModifySecurityPolicyRequest) {
    request = &ModifySecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifySecurityPolicy")
    return
}

func NewModifySecurityPolicyResponse() (response *ModifySecurityPolicyResponse) {
    response = &ModifySecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新实例公网访问白名单
func (c *Client) ModifySecurityPolicy(request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyRequest()
    }
    response = NewModifySecurityPolicyResponse()
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

func NewValidateApplicationTokenPersonalRequest() (request *ValidateApplicationTokenPersonalRequest) {
    request = &ValidateApplicationTokenPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ValidateApplicationTokenPersonal")
    return
}

func NewValidateApplicationTokenPersonalResponse() (response *ValidateApplicationTokenPersonalResponse) {
    response = &ValidateApplicationTokenPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于验证第三方应用访问凭证
func (c *Client) ValidateApplicationTokenPersonal(request *ValidateApplicationTokenPersonalRequest) (response *ValidateApplicationTokenPersonalResponse, err error) {
    if request == nil {
        request = NewValidateApplicationTokenPersonalRequest()
    }
    response = NewValidateApplicationTokenPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewValidateGitHubAuthPersonalRequest() (request *ValidateGitHubAuthPersonalRequest) {
    request = &ValidateGitHubAuthPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ValidateGitHubAuthPersonal")
    return
}

func NewValidateGitHubAuthPersonalResponse() (response *ValidateGitHubAuthPersonalResponse) {
    response = &ValidateGitHubAuthPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 验证GitHub授权
func (c *Client) ValidateGitHubAuthPersonal(request *ValidateGitHubAuthPersonalRequest) (response *ValidateGitHubAuthPersonalResponse, err error) {
    if request == nil {
        request = NewValidateGitHubAuthPersonalRequest()
    }
    response = NewValidateGitHubAuthPersonalResponse()
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

func NewValidateUserPersonalRequest() (request *ValidateUserPersonalRequest) {
    request = &ValidateUserPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ValidateUserPersonal")
    return
}

func NewValidateUserPersonalResponse() (response *ValidateUserPersonalResponse) {
    response = &ValidateUserPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 验证个人用户
func (c *Client) ValidateUserPersonal(request *ValidateUserPersonalRequest) (response *ValidateUserPersonalResponse, err error) {
    if request == nil {
        request = NewValidateUserPersonalRequest()
    }
    response = NewValidateUserPersonalResponse()
    err = c.Send(request, response)
    return
}
