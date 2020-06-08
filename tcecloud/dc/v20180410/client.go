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

package v20180410

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-04-10"

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


func NewApproveDirectConnectTunnelRequest() (request *ApproveDirectConnectTunnelRequest) {
    request = &ApproveDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ApproveDirectConnectTunnel")
    return
}

func NewApproveDirectConnectTunnelResponse() (response *ApproveDirectConnectTunnelResponse) {
    response = &ApproveDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 专线所有者审批共享通道申请
func (c *Client) ApproveDirectConnectTunnel(request *ApproveDirectConnectTunnelRequest) (response *ApproveDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewApproveDirectConnectTunnelRequest()
    }
    response = NewApproveDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDirectConnectRequest() (request *CreateDirectConnectRequest) {
    request = &CreateDirectConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "CreateDirectConnect")
    return
}

func NewCreateDirectConnectResponse() (response *CreateDirectConnectResponse) {
    response = &CreateDirectConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 申请物理专线接入。
// 调用该接口时，请注意：
// 账号要进行实名认证，否则不允许申请物理专线；
// 若账户下存在欠费状态的物理专线，则不能申请更多的物理专线。
func (c *Client) CreateDirectConnect(request *CreateDirectConnectRequest) (response *CreateDirectConnectResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectRequest()
    }
    response = NewCreateDirectConnectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDirectConnectTunnelRequest() (request *CreateDirectConnectTunnelRequest) {
    request = &CreateDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "CreateDirectConnectTunnel")
    return
}

func NewCreateDirectConnectTunnelResponse() (response *CreateDirectConnectTunnelResponse) {
    response = &CreateDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于创建专用通道的接口
func (c *Client) CreateDirectConnectTunnel(request *CreateDirectConnectTunnelRequest) (response *CreateDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectTunnelRequest()
    }
    response = NewCreateDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDirectConnectRequest() (request *DeleteDirectConnectRequest) {
    request = &DeleteDirectConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DeleteDirectConnect")
    return
}

func NewDeleteDirectConnectResponse() (response *DeleteDirectConnectResponse) {
    response = &DeleteDirectConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除物理专线。
// 只能删除处于状态的物理专线。
func (c *Client) DeleteDirectConnect(request *DeleteDirectConnectRequest) (response *DeleteDirectConnectResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectRequest()
    }
    response = NewDeleteDirectConnectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDirectConnectTunnelRequest() (request *DeleteDirectConnectTunnelRequest) {
    request = &DeleteDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DeleteDirectConnectTunnel")
    return
}

func NewDeleteDirectConnectTunnelResponse() (response *DeleteDirectConnectTunnelResponse) {
    response = &DeleteDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除专用通道
func (c *Client) DeleteDirectConnectTunnel(request *DeleteDirectConnectTunnelRequest) (response *DeleteDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectTunnelRequest()
    }
    response = NewDeleteDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessPointsRequest() (request *DescribeAccessPointsRequest) {
    request = &DescribeAccessPointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeAccessPoints")
    return
}

func NewDescribeAccessPointsResponse() (response *DescribeAccessPointsResponse) {
    response = &DescribeAccessPointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询物理专线接入点
func (c *Client) DescribeAccessPoints(request *DescribeAccessPointsRequest) (response *DescribeAccessPointsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessPointsRequest()
    }
    response = NewDescribeAccessPointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDirectConnectTunnelsRequest() (request *DescribeDirectConnectTunnelsRequest) {
    request = &DescribeDirectConnectTunnelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeDirectConnectTunnels")
    return
}

func NewDescribeDirectConnectTunnelsResponse() (response *DescribeDirectConnectTunnelsResponse) {
    response = &DescribeDirectConnectTunnelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询专用通道列表。
func (c *Client) DescribeDirectConnectTunnels(request *DescribeDirectConnectTunnelsRequest) (response *DescribeDirectConnectTunnelsResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectTunnelsRequest()
    }
    response = NewDescribeDirectConnectTunnelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDirectConnectTunnelAttributeRequest() (request *ModifyDirectConnectTunnelAttributeRequest) {
    request = &ModifyDirectConnectTunnelAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ModifyDirectConnectTunnelAttribute")
    return
}

func NewModifyDirectConnectTunnelAttributeResponse() (response *ModifyDirectConnectTunnelAttributeResponse) {
    response = &ModifyDirectConnectTunnelAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改专用通道属性
func (c *Client) ModifyDirectConnectTunnelAttribute(request *ModifyDirectConnectTunnelAttributeRequest) (response *ModifyDirectConnectTunnelAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDirectConnectTunnelAttributeRequest()
    }
    response = NewModifyDirectConnectTunnelAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateVifAssociatedRequest() (request *UpdateVifAssociatedRequest) {
    request = &UpdateVifAssociatedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "UpdateVifAssociated")
    return
}

func NewUpdateVifAssociatedResponse() (response *UpdateVifAssociatedResponse) {
    response = &UpdateVifAssociatedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新通道冗余模式
func (c *Client) UpdateVifAssociated(request *UpdateVifAssociatedRequest) (response *UpdateVifAssociatedResponse, err error) {
    if request == nil {
        request = NewUpdateVifAssociatedRequest()
    }
    response = NewUpdateVifAssociatedResponse()
    err = c.Send(request, response)
    return
}
