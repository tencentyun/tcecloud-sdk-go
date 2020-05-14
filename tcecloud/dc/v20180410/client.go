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
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
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


func NewAcceptDirectConnectRequest() (request *AcceptDirectConnectRequest) {
    request = &AcceptDirectConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "AcceptDirectConnect")
    return
}

func NewAcceptDirectConnectResponse() (response *AcceptDirectConnectResponse) {
    response = &AcceptDirectConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 审批通过专线接入申请，调用时传入PortId，则使用指定的端口创建专线，否则自动分配交换机端口
func (c *Client) AcceptDirectConnect(request *AcceptDirectConnectRequest) (response *AcceptDirectConnectResponse, err error) {
    if request == nil {
        request = NewAcceptDirectConnectRequest()
    }
    response = NewAcceptDirectConnectResponse()
    err = c.Send(request, response)
    return
}

func NewAcceptDirectConnectTunnelRequest() (request *AcceptDirectConnectTunnelRequest) {
    request = &AcceptDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "AcceptDirectConnectTunnel")
    return
}

func NewAcceptDirectConnectTunnelResponse() (response *AcceptDirectConnectTunnelResponse) {
    response = &AcceptDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 接受专用通道申请
func (c *Client) AcceptDirectConnectTunnel(request *AcceptDirectConnectTunnelRequest) (response *AcceptDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewAcceptDirectConnectTunnelRequest()
    }
    response = NewAcceptDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessPointRequest() (request *CreateAccessPointRequest) {
    request = &CreateAccessPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "CreateAccessPoint")
    return
}

func NewCreateAccessPointResponse() (response *CreateAccessPointResponse) {
    response = &CreateAccessPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建专线接入点
func (c *Client) CreateAccessPoint(request *CreateAccessPointRequest) (response *CreateAccessPointResponse, err error) {
    if request == nil {
        request = NewCreateAccessPointRequest()
    }
    response = NewCreateAccessPointResponse()
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

func NewCreatePublicDirectConnectTunnelRequest() (request *CreatePublicDirectConnectTunnelRequest) {
    request = &CreatePublicDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "CreatePublicDirectConnectTunnel")
    return
}

func NewCreatePublicDirectConnectTunnelResponse() (response *CreatePublicDirectConnectTunnelResponse) {
    response = &CreatePublicDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于创建互联网专用通道的接口
func (c *Client) CreatePublicDirectConnectTunnel(request *CreatePublicDirectConnectTunnelRequest) (response *CreatePublicDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewCreatePublicDirectConnectTunnelRequest()
    }
    response = NewCreatePublicDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSwitchRequest() (request *CreateSwitchRequest) {
    request = &CreateSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "CreateSwitch")
    return
}

func NewCreateSwitchResponse() (response *CreateSwitchResponse) {
    response = &CreateSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建专线接入交换机
func (c *Client) CreateSwitch(request *CreateSwitchRequest) (response *CreateSwitchResponse, err error) {
    if request == nil {
        request = NewCreateSwitchRequest()
    }
    response = NewCreateSwitchResponse()
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

func NewDeleteSwitchRequest() (request *DeleteSwitchRequest) {
    request = &DeleteSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DeleteSwitch")
    return
}

func NewDeleteSwitchResponse() (response *DeleteSwitchResponse) {
    response = &DeleteSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除专线接入交换机，仅在交换机上没有正在运行的专线时才能删除该交换机
func (c *Client) DeleteSwitch(request *DeleteSwitchRequest) (response *DeleteSwitchResponse, err error) {
    if request == nil {
        request = NewDeleteSwitchRequest()
    }
    response = NewDeleteSwitchResponse()
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

func NewDescribeAvailableSwitchesRequest() (request *DescribeAvailableSwitchesRequest) {
    request = &DescribeAvailableSwitchesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeAvailableSwitches")
    return
}

func NewDescribeAvailableSwitchesResponse() (response *DescribeAvailableSwitchesResponse) {
    response = &DescribeAvailableSwitchesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从DCOS获取专线接入交换机
func (c *Client) DescribeAvailableSwitches(request *DescribeAvailableSwitchesRequest) (response *DescribeAvailableSwitchesResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableSwitchesRequest()
    }
    response = NewDescribeAvailableSwitchesResponse()
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

func NewDescribeDirectConnectsRequest() (request *DescribeDirectConnectsRequest) {
    request = &DescribeDirectConnectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeDirectConnects")
    return
}

func NewDescribeDirectConnectsResponse() (response *DescribeDirectConnectsResponse) {
    response = &DescribeDirectConnectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询物理专线列表。
func (c *Client) DescribeDirectConnects(request *DescribeDirectConnectsRequest) (response *DescribeDirectConnectsResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectsRequest()
    }
    response = NewDescribeDirectConnectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSwitchPortsRequest() (request *DescribeSwitchPortsRequest) {
    request = &DescribeSwitchPortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeSwitchPorts")
    return
}

func NewDescribeSwitchPortsResponse() (response *DescribeSwitchPortsResponse) {
    response = &DescribeSwitchPortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取交换机可用端口
func (c *Client) DescribeSwitchPorts(request *DescribeSwitchPortsRequest) (response *DescribeSwitchPortsResponse, err error) {
    if request == nil {
        request = NewDescribeSwitchPortsRequest()
    }
    response = NewDescribeSwitchPortsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSwitchesRequest() (request *DescribeSwitchesRequest) {
    request = &DescribeSwitchesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeSwitches")
    return
}

func NewDescribeSwitchesResponse() (response *DescribeSwitchesResponse) {
    response = &DescribeSwitchesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取专线接入交换机
func (c *Client) DescribeSwitches(request *DescribeSwitchesRequest) (response *DescribeSwitchesResponse, err error) {
    if request == nil {
        request = NewDescribeSwitchesRequest()
    }
    response = NewDescribeSwitchesResponse()
    err = c.Send(request, response)
    return
}

func NewImplementDirectConnectRequest() (request *ImplementDirectConnectRequest) {
    request = &ImplementDirectConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ImplementDirectConnect")
    return
}

func NewImplementDirectConnectResponse() (response *ImplementDirectConnectResponse) {
    response = &ImplementDirectConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 专线人工确认实施完成，填写线路编号，保障联系人，保障联系电话
func (c *Client) ImplementDirectConnect(request *ImplementDirectConnectRequest) (response *ImplementDirectConnectResponse, err error) {
    if request == nil {
        request = NewImplementDirectConnectRequest()
    }
    response = NewImplementDirectConnectResponse()
    err = c.Send(request, response)
    return
}

func NewIsSameRegionRequest() (request *IsSameRegionRequest) {
    request = &IsSameRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "IsSameRegion")
    return
}

func NewIsSameRegionResponse() (response *IsSameRegionResponse) {
    response = &IsSameRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 判断两个地域同城
func (c *Client) IsSameRegion(request *IsSameRegionRequest) (response *IsSameRegionResponse, err error) {
    if request == nil {
        request = NewIsSameRegionRequest()
    }
    response = NewIsSameRegionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessPointRequest() (request *ModifyAccessPointRequest) {
    request = &ModifyAccessPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ModifyAccessPoint")
    return
}

func NewModifyAccessPointResponse() (response *ModifyAccessPointResponse) {
    response = &ModifyAccessPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改专线接入点信息
func (c *Client) ModifyAccessPoint(request *ModifyAccessPointRequest) (response *ModifyAccessPointResponse, err error) {
    if request == nil {
        request = NewModifyAccessPointRequest()
    }
    response = NewModifyAccessPointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDirectConnectAttributeRequest() (request *ModifyDirectConnectAttributeRequest) {
    request = &ModifyDirectConnectAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ModifyDirectConnectAttribute")
    return
}

func NewModifyDirectConnectAttributeResponse() (response *ModifyDirectConnectAttributeResponse) {
    response = &ModifyDirectConnectAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改物理专线的属性。
func (c *Client) ModifyDirectConnectAttribute(request *ModifyDirectConnectAttributeRequest) (response *ModifyDirectConnectAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDirectConnectAttributeRequest()
    }
    response = NewModifyDirectConnectAttributeResponse()
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

func NewRejectDirectConnectRequest() (request *RejectDirectConnectRequest) {
    request = &RejectDirectConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "RejectDirectConnect")
    return
}

func NewRejectDirectConnectResponse() (response *RejectDirectConnectResponse) {
    response = &RejectDirectConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运营端拒绝专线申请
func (c *Client) RejectDirectConnect(request *RejectDirectConnectRequest) (response *RejectDirectConnectResponse, err error) {
    if request == nil {
        request = NewRejectDirectConnectRequest()
    }
    response = NewRejectDirectConnectResponse()
    err = c.Send(request, response)
    return
}

func NewRejectDirectConnectTunnelRequest() (request *RejectDirectConnectTunnelRequest) {
    request = &RejectDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "RejectDirectConnectTunnel")
    return
}

func NewRejectDirectConnectTunnelResponse() (response *RejectDirectConnectTunnelResponse) {
    response = &RejectDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拒绝专用通道申请
func (c *Client) RejectDirectConnectTunnel(request *RejectDirectConnectTunnelRequest) (response *RejectDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewRejectDirectConnectTunnelRequest()
    }
    response = NewRejectDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateVifNetDetectRequest() (request *UpdateVifNetDetectRequest) {
    request = &UpdateVifNetDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "UpdateVifNetDetect")
    return
}

func NewUpdateVifNetDetectResponse() (response *UpdateVifNetDetectResponse) {
    response = &UpdateVifNetDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新专用通道关联的网络自定义探测信息
func (c *Client) UpdateVifNetDetect(request *UpdateVifNetDetectRequest) (response *UpdateVifNetDetectResponse, err error) {
    if request == nil {
        request = NewUpdateVifNetDetectRequest()
    }
    response = NewUpdateVifNetDetectResponse()
    err = c.Send(request, response)
    return
}
