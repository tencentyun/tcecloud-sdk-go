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

package v20190116

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-01-16"

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


func NewAttachRolePoliciesRequest() (request *AttachRolePoliciesRequest) {
    request = &AttachRolePoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "AttachRolePolicies")
    return
}

func NewAttachRolePoliciesResponse() (response *AttachRolePoliciesResponse) {
    response = &AttachRolePoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 绑定多个策略到角色 
func (c *Client) AttachRolePolicies(request *AttachRolePoliciesRequest) (response *AttachRolePoliciesResponse, err error) {
    if request == nil {
        request = NewAttachRolePoliciesRequest()
    }
    response = NewAttachRolePoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewAttachRolePolicyRequest() (request *AttachRolePolicyRequest) {
    request = &AttachRolePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "AttachRolePolicy")
    return
}

func NewAttachRolePolicyResponse() (response *AttachRolePolicyResponse) {
    response = &AttachRolePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AttachRolePolicy）用于绑定策略到角色。
func (c *Client) AttachRolePolicy(request *AttachRolePolicyRequest) (response *AttachRolePolicyResponse, err error) {
    if request == nil {
        request = NewAttachRolePolicyRequest()
    }
    response = NewAttachRolePolicyResponse()
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

func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreateRole")
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateRole）用于创建角色。
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    response = NewCreateRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
    request = &DeleteRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteRole")
    return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
    response = &DeleteRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteRole）用于删除指定角色。
func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    if request == nil {
        request = NewDeleteRoleRequest()
    }
    response = NewDeleteRoleResponse()
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

func NewDescribeRoleListRequest() (request *DescribeRoleListRequest) {
    request = &DescribeRoleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DescribeRoleList")
    return
}

func NewDescribeRoleListResponse() (response *DescribeRoleListResponse) {
    response = &DescribeRoleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRoleList）用于获取账号下的角色列表。
func (c *Client) DescribeRoleList(request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    if request == nil {
        request = NewDescribeRoleListRequest()
    }
    response = NewDescribeRoleListResponse()
    err = c.Send(request, response)
    return
}

func NewDetachGroupsPolicyRequest() (request *DetachGroupsPolicyRequest) {
    request = &DetachGroupsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DetachGroupsPolicy")
    return
}

func NewDetachGroupsPolicyResponse() (response *DetachGroupsPolicyResponse) {
    response = &DetachGroupsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解除绑定策略到多个用户组
func (c *Client) DetachGroupsPolicy(request *DetachGroupsPolicyRequest) (response *DetachGroupsPolicyResponse, err error) {
    if request == nil {
        request = NewDetachGroupsPolicyRequest()
    }
    response = NewDetachGroupsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDetachUsersPolicyRequest() (request *DetachUsersPolicyRequest) {
    request = &DetachUsersPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DetachUsersPolicy")
    return
}

func NewDetachUsersPolicyResponse() (response *DetachUsersPolicyResponse) {
    response = &DetachUsersPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解除绑定策略到多个用户 
func (c *Client) DetachUsersPolicy(request *DetachUsersPolicyRequest) (response *DetachUsersPolicyResponse, err error) {
    if request == nil {
        request = NewDetachUsersPolicyRequest()
    }
    response = NewDetachUsersPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewGetPolicyRequest() (request *GetPolicyRequest) {
    request = &GetPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetPolicy")
    return
}

func NewGetPolicyResponse() (response *GetPolicyResponse) {
    response = &GetPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetPolicy）可用于查询查看策略详情。
func (c *Client) GetPolicy(request *GetPolicyRequest) (response *GetPolicyResponse, err error) {
    if request == nil {
        request = NewGetPolicyRequest()
    }
    response = NewGetPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoleRequest() (request *GetRoleRequest) {
    request = &GetRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetRole")
    return
}

func NewGetRoleResponse() (response *GetRoleResponse) {
    response = &GetRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetRole）用于获取指定角色的详细信息。
func (c *Client) GetRole(request *GetRoleRequest) (response *GetRoleResponse, err error) {
    if request == nil {
        request = NewGetRoleRequest()
    }
    response = NewGetRoleResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceApiListRequest() (request *GetServiceApiListRequest) {
    request = &GetServiceApiListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetServiceApiList")
    return
}

func NewGetServiceApiListResponse() (response *GetServiceApiListResponse) {
    response = &GetServiceApiListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取服务及其API列表 
func (c *Client) GetServiceApiList(request *GetServiceApiListRequest) (response *GetServiceApiListResponse, err error) {
    if request == nil {
        request = NewGetServiceApiListRequest()
    }
    response = NewGetServiceApiListResponse()
    err = c.Send(request, response)
    return
}

func NewListAttachedRolePoliciesRequest() (request *ListAttachedRolePoliciesRequest) {
    request = &ListAttachedRolePoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListAttachedRolePolicies")
    return
}

func NewListAttachedRolePoliciesResponse() (response *ListAttachedRolePoliciesResponse) {
    response = &ListAttachedRolePoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListAttachedRolePolicies）用于获取角色绑定的策略列表。
func (c *Client) ListAttachedRolePolicies(request *ListAttachedRolePoliciesRequest) (response *ListAttachedRolePoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedRolePoliciesRequest()
    }
    response = NewListAttachedRolePoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListEntitiesForPolicyRequest() (request *ListEntitiesForPolicyRequest) {
    request = &ListEntitiesForPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListEntitiesForPolicy")
    return
}

func NewListEntitiesForPolicyResponse() (response *ListEntitiesForPolicyResponse) {
    response = &ListEntitiesForPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListEntitiesForPolicy）可用于查询策略关联的实体列表。
func (c *Client) ListEntitiesForPolicy(request *ListEntitiesForPolicyRequest) (response *ListEntitiesForPolicyResponse, err error) {
    if request == nil {
        request = NewListEntitiesForPolicyRequest()
    }
    response = NewListEntitiesForPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewListPoliciesRequest() (request *ListPoliciesRequest) {
    request = &ListPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListPolicies")
    return
}

func NewListPoliciesResponse() (response *ListPoliciesResponse) {
    response = &ListPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListPolicies）可用于查询策略列表。
func (c *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
    if request == nil {
        request = NewListPoliciesRequest()
    }
    response = NewListPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAssumeRolePolicyRequest() (request *UpdateAssumeRolePolicyRequest) {
    request = &UpdateAssumeRolePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateAssumeRolePolicy")
    return
}

func NewUpdateAssumeRolePolicyResponse() (response *UpdateAssumeRolePolicyResponse) {
    response = &UpdateAssumeRolePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateAssumeRolePolicy）用于修改角色信任策略的策略文档。
func (c *Client) UpdateAssumeRolePolicy(request *UpdateAssumeRolePolicyRequest) (response *UpdateAssumeRolePolicyResponse, err error) {
    if request == nil {
        request = NewUpdateAssumeRolePolicyRequest()
    }
    response = NewUpdateAssumeRolePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
    request = &UpdatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdatePolicy")
    return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
    response = &UpdatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdatePolicy ）可用于更新策略。
// 如果已存在策略版本，本接口会直接更新策略的默认版本，不会创建新版本，如果不存在任何策略版本，则直接创建一个默认版本。
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRoleDescriptionRequest() (request *UpdateRoleDescriptionRequest) {
    request = &UpdateRoleDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateRoleDescription")
    return
}

func NewUpdateRoleDescriptionResponse() (response *UpdateRoleDescriptionResponse) {
    response = &UpdateRoleDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateRoleDescription）用于修改角色的描述信息。
func (c *Client) UpdateRoleDescription(request *UpdateRoleDescriptionRequest) (response *UpdateRoleDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateRoleDescriptionRequest()
    }
    response = NewUpdateRoleDescriptionResponse()
    err = c.Send(request, response)
    return
}
