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

package v20200226

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-26"

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


func NewCheckUserPolicyAttachmentRequest() (request *CheckUserPolicyAttachmentRequest) {
    request = &CheckUserPolicyAttachmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CheckUserPolicyAttachment")
    return
}

func NewCheckUserPolicyAttachmentResponse() (response *CheckUserPolicyAttachmentResponse) {
    response = &CheckUserPolicyAttachmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CheckUserPolicyAttachment）用于查询用户是否关联策略。
func (c *Client) CheckUserPolicyAttachment(request *CheckUserPolicyAttachmentRequest) (response *CheckUserPolicyAttachmentResponse, err error) {
    if request == nil {
        request = NewCheckUserPolicyAttachmentRequest()
    }
    response = NewCheckUserPolicyAttachmentResponse()
    err = c.Send(request, response)
    return
}

func NewConfirmCASProviderRequest() (request *ConfirmCASProviderRequest) {
    request = &ConfirmCASProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ConfirmCASProvider")
    return
}

func NewConfirmCASProviderResponse() (response *ConfirmCASProviderResponse) {
    response = &ConfirmCASProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 确认CAS身份提供商
func (c *Client) ConfirmCASProvider(request *ConfirmCASProviderRequest) (response *ConfirmCASProviderResponse, err error) {
    if request == nil {
        request = NewConfirmCASProviderRequest()
    }
    response = NewConfirmCASProviderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCASProviderRequest() (request *CreateCASProviderRequest) {
    request = &CreateCASProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreateCASProvider")
    return
}

func NewCreateCASProviderResponse() (response *CreateCASProviderResponse) {
    response = &CreateCASProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建CAS身份提供商
func (c *Client) CreateCASProvider(request *CreateCASProviderRequest) (response *CreateCASProviderResponse, err error) {
    if request == nil {
        request = NewCreateCASProviderRequest()
    }
    response = NewCreateCASProviderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCASProviderRequest() (request *DeleteCASProviderRequest) {
    request = &DeleteCASProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteCASProvider")
    return
}

func NewDeleteCASProviderResponse() (response *DeleteCASProviderResponse) {
    response = &DeleteCASProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除cas身份提供商信息
func (c *Client) DeleteCASProvider(request *DeleteCASProviderRequest) (response *DeleteCASProviderResponse, err error) {
    if request == nil {
        request = NewDeleteCASProviderRequest()
    }
    response = NewDeleteCASProviderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubAccountRequest() (request *DeleteSubAccountRequest) {
    request = &DeleteSubAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteSubAccount")
    return
}

func NewDeleteSubAccountResponse() (response *DeleteSubAccountResponse) {
    response = &DeleteSubAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除子帐号
func (c *Client) DeleteSubAccount(request *DeleteSubAccountRequest) (response *DeleteSubAccountResponse, err error) {
    if request == nil {
        request = NewDeleteSubAccountRequest()
    }
    response = NewDeleteSubAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasProviderRequest() (request *DescribeCasProviderRequest) {
    request = &DescribeCasProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DescribeCasProvider")
    return
}

func NewDescribeCasProviderResponse() (response *DescribeCasProviderResponse) {
    response = &DescribeCasProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取cas身份提供商信息
func (c *Client) DescribeCasProvider(request *DescribeCasProviderRequest) (response *DescribeCasProviderResponse, err error) {
    if request == nil {
        request = NewDescribeCasProviderRequest()
    }
    response = NewDescribeCasProviderResponse()
    err = c.Send(request, response)
    return
}

func NewDisableCASProviderRequest() (request *DisableCASProviderRequest) {
    request = &DisableCASProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DisableCASProvider")
    return
}

func NewDisableCASProviderResponse() (response *DisableCASProviderResponse) {
    response = &DisableCASProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 禁用CAS身份提供商
func (c *Client) DisableCASProvider(request *DisableCASProviderRequest) (response *DisableCASProviderResponse, err error) {
    if request == nil {
        request = NewDisableCASProviderRequest()
    }
    response = NewDisableCASProviderResponse()
    err = c.Send(request, response)
    return
}

func NewListSubAccountsRequest() (request *ListSubAccountsRequest) {
    request = &ListSubAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListSubAccounts")
    return
}

func NewListSubAccountsResponse() (response *ListSubAccountsResponse) {
    response = &ListSubAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取子帐号列表
func (c *Client) ListSubAccounts(request *ListSubAccountsRequest) (response *ListSubAccountsResponse, err error) {
    if request == nil {
        request = NewListSubAccountsRequest()
    }
    response = NewListSubAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersForPolicyRequest() (request *ListUsersForPolicyRequest) {
    request = &ListUsersForPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListUsersForPolicy")
    return
}

func NewListUsersForPolicyResponse() (response *ListUsersForPolicyResponse) {
    response = &ListUsersForPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListUsersForPolicy）用于列出策略关联的用户列表（包括随组关联）。
func (c *Client) ListUsersForPolicy(request *ListUsersForPolicyRequest) (response *ListUsersForPolicyResponse, err error) {
    if request == nil {
        request = NewListUsersForPolicyRequest()
    }
    response = NewListUsersForPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCASProviderRequest() (request *UpdateCASProviderRequest) {
    request = &UpdateCASProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateCASProvider")
    return
}

func NewUpdateCASProviderResponse() (response *UpdateCASProviderResponse) {
    response = &UpdateCASProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新CAS身份提供商
func (c *Client) UpdateCASProvider(request *UpdateCASProviderRequest) (response *UpdateCASProviderResponse, err error) {
    if request == nil {
        request = NewUpdateCASProviderRequest()
    }
    response = NewUpdateCASProviderResponse()
    err = c.Send(request, response)
    return
}
