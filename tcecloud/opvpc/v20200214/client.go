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

package v20200214

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-02-14"

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


func NewDescribeAclListRequest() (request *DescribeAclListRequest) {
    request = &DescribeAclListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeAclList")
    return
}

func NewDescribeAclListResponse() (response *DescribeAclListResponse) {
    response = &DescribeAclListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取acl列表
func (c *Client) DescribeAclList(request *DescribeAclListRequest) (response *DescribeAclListResponse, err error) {
    if request == nil {
        request = NewDescribeAclListRequest()
    }
    response = NewDescribeAclListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCvmForSecurityGroupRequest() (request *DescribeCvmForSecurityGroupRequest) {
    request = &DescribeCvmForSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeCvmForSecurityGroup")
    return
}

func NewDescribeCvmForSecurityGroupResponse() (response *DescribeCvmForSecurityGroupResponse) {
    response = &DescribeCvmForSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取安全组关联实例列表
func (c *Client) DescribeCvmForSecurityGroup(request *DescribeCvmForSecurityGroupRequest) (response *DescribeCvmForSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeCvmForSecurityGroupRequest()
    }
    response = NewDescribeCvmForSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4RuleExRequest() (request *DescribeL4RuleExRequest) {
    request = &DescribeL4RuleExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeL4RuleEx")
    return
}

func NewDescribeL4RuleExResponse() (response *DescribeL4RuleExResponse) {
    response = &DescribeL4RuleExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取四层规则列表
func (c *Client) DescribeL4RuleEx(request *DescribeL4RuleExRequest) (response *DescribeL4RuleExResponse, err error) {
    if request == nil {
        request = NewDescribeL4RuleExRequest()
    }
    response = NewDescribeL4RuleExResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7RuleExRequest() (request *DescribeL7RuleExRequest) {
    request = &DescribeL7RuleExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeL7RuleEx")
    return
}

func NewDescribeL7RuleExResponse() (response *DescribeL7RuleExResponse) {
    response = &DescribeL7RuleExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取七层规则列表
func (c *Client) DescribeL7RuleEx(request *DescribeL7RuleExRequest) (response *DescribeL7RuleExResponse, err error) {
    if request == nil {
        request = NewDescribeL7RuleExRequest()
    }
    response = NewDescribeL7RuleExResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLbLimitRequest() (request *DescribeLbLimitRequest) {
    request = &DescribeLbLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeLbLimit")
    return
}

func NewDescribeLbLimitResponse() (response *DescribeLbLimitResponse) {
    response = &DescribeLbLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取lb的配额
func (c *Client) DescribeLbLimit(request *DescribeLbLimitRequest) (response *DescribeLbLimitResponse, err error) {
    if request == nil {
        request = NewDescribeLbLimitRequest()
    }
    response = NewDescribeLbLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
    request = &DescribeNatGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeNatGateways")
    return
}

func NewDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
    response = &DescribeNatGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取nat 网关列表
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaysRequest()
    }
    response = NewDescribeNatGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRtbSubnetsRequest() (request *DescribeRtbSubnetsRequest) {
    request = &DescribeRtbSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeRtbSubnets")
    return
}

func NewDescribeRtbSubnetsResponse() (response *DescribeRtbSubnetsResponse) {
    response = &DescribeRtbSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查路由表对应的全量子网
func (c *Client) DescribeRtbSubnets(request *DescribeRtbSubnetsRequest) (response *DescribeRtbSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeRtbSubnetsRequest()
    }
    response = NewDescribeRtbSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRtbsExRequest() (request *DescribeRtbsExRequest) {
    request = &DescribeRtbsExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeRtbsEx")
    return
}

func NewDescribeRtbsExResponse() (response *DescribeRtbsExResponse) {
    response = &DescribeRtbsExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询路由列表
func (c *Client) DescribeRtbsEx(request *DescribeRtbsExRequest) (response *DescribeRtbsExResponse, err error) {
    if request == nil {
        request = NewDescribeRtbsExRequest()
    }
    response = NewDescribeRtbsExResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
    request = &DescribeSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeSecurityGroup")
    return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
    response = &DescribeSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取安全组列表
func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) (response *DescribeSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupRequest()
    }
    response = NewDescribeSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupForCvmRequest() (request *DescribeSecurityGroupForCvmRequest) {
    request = &DescribeSecurityGroupForCvmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeSecurityGroupForCvm")
    return
}

func NewDescribeSecurityGroupForCvmResponse() (response *DescribeSecurityGroupForCvmResponse) {
    response = &DescribeSecurityGroupForCvmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取vm关联的安全组
func (c *Client) DescribeSecurityGroupForCvm(request *DescribeSecurityGroupForCvmRequest) (response *DescribeSecurityGroupForCvmResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupForCvmRequest()
    }
    response = NewDescribeSecurityGroupForCvmResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupPolicyRequest() (request *DescribeSecurityGroupPolicyRequest) {
    request = &DescribeSecurityGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeSecurityGroupPolicy")
    return
}

func NewDescribeSecurityGroupPolicyResponse() (response *DescribeSecurityGroupPolicyResponse) {
    response = &DescribeSecurityGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取安全组策略
func (c *Client) DescribeSecurityGroupPolicy(request *DescribeSecurityGroupPolicyRequest) (response *DescribeSecurityGroupPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupPolicyRequest()
    }
    response = NewDescribeSecurityGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetsExRequest() (request *DescribeSubnetsExRequest) {
    request = &DescribeSubnetsExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeSubnetsEx")
    return
}

func NewDescribeSubnetsExResponse() (response *DescribeSubnetsExResponse) {
    response = &DescribeSubnetsExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取子网列表
func (c *Client) DescribeSubnetsEx(request *DescribeSubnetsExRequest) (response *DescribeSubnetsExResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetsExRequest()
    }
    response = NewDescribeSubnetsExResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcsExRequest() (request *DescribeVpcsExRequest) {
    request = &DescribeVpcsExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcsEx")
    return
}

func NewDescribeVpcsExResponse() (response *DescribeVpcsExResponse) {
    response = &DescribeVpcsExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取vpc列表
func (c *Client) DescribeVpcsEx(request *DescribeVpcsExRequest) (response *DescribeVpcsExResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsExRequest()
    }
    response = NewDescribeVpcsExResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLbLimitRequest() (request *ModifyLbLimitRequest) {
    request = &ModifyLbLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opvpc", APIVersion, "ModifyLbLimit")
    return
}

func NewModifyLbLimitResponse() (response *ModifyLbLimitResponse) {
    response = &ModifyLbLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改LB配额限制
func (c *Client) ModifyLbLimit(request *ModifyLbLimitRequest) (response *ModifyLbLimitResponse, err error) {
    if request == nil {
        request = NewModifyLbLimitRequest()
    }
    response = NewModifyLbLimitResponse()
    err = c.Send(request, response)
    return
}
