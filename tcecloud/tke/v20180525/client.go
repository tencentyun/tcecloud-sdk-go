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

package v20180525

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-05-25"

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


func NewAddAlarmPolicyRequest() (request *AddAlarmPolicyRequest) {
    request = &AddAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "AddAlarmPolicy")
    return
}

func NewAddAlarmPolicyResponse() (response *AddAlarmPolicyResponse) {
    response = &AddAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加告警策略
func (c *Client) AddAlarmPolicy(request *AddAlarmPolicyRequest) (response *AddAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewAddAlarmPolicyRequest()
    }
    response = NewAddAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewAddClusterCIDRToCcnRequest() (request *AddClusterCIDRToCcnRequest) {
    request = &AddClusterCIDRToCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "AddClusterCIDRToCcn")
    return
}

func NewAddClusterCIDRToCcnResponse() (response *AddClusterCIDRToCcnResponse) {
    response = &AddClusterCIDRToCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加TKE集群CIDR到云联网
func (c *Client) AddClusterCIDRToCcn(request *AddClusterCIDRToCcnRequest) (response *AddClusterCIDRToCcnResponse, err error) {
    if request == nil {
        request = NewAddClusterCIDRToCcnRequest()
    }
    response = NewAddClusterCIDRToCcnResponse()
    err = c.Send(request, response)
    return
}

func NewAddClusterCIDRToVbcRequest() (request *AddClusterCIDRToVbcRequest) {
    request = &AddClusterCIDRToVbcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "AddClusterCIDRToVbc")
    return
}

func NewAddClusterCIDRToVbcResponse() (response *AddClusterCIDRToVbcResponse) {
    response = &AddClusterCIDRToVbcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布tke集群cidr到云联网
func (c *Client) AddClusterCIDRToVbc(request *AddClusterCIDRToVbcRequest) (response *AddClusterCIDRToVbcResponse, err error) {
    if request == nil {
        request = NewAddClusterCIDRToVbcRequest()
    }
    response = NewAddClusterCIDRToVbcResponse()
    err = c.Send(request, response)
    return
}

func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
    request = &AddClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "AddClusterInstances")
    return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
    response = &AddClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 扩展集群节点，API 3.0
func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
    if request == nil {
        request = NewAddClusterInstancesRequest()
    }
    response = NewAddClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAddExistedInstancesRequest() (request *AddExistedInstancesRequest) {
    request = &AddExistedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "AddExistedInstances")
    return
}

func NewAddExistedInstancesResponse() (response *AddExistedInstancesResponse) {
    response = &AddExistedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加已经存在的实例到集群
func (c *Client) AddExistedInstances(request *AddExistedInstancesRequest) (response *AddExistedInstancesResponse, err error) {
    if request == nil {
        request = NewAddExistedInstancesRequest()
    }
    response = NewAddExistedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCheckClusterCIDRRequest() (request *CheckClusterCIDRRequest) {
    request = &CheckClusterCIDRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CheckClusterCIDR")
    return
}

func NewCheckClusterCIDRResponse() (response *CheckClusterCIDRResponse) {
    response = &CheckClusterCIDRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查集群的CIDR是否冲突
func (c *Client) CheckClusterCIDR(request *CheckClusterCIDRRequest) (response *CheckClusterCIDRResponse, err error) {
    if request == nil {
        request = NewCheckClusterCIDRRequest()
    }
    response = NewCheckClusterCIDRResponse()
    err = c.Send(request, response)
    return
}

func NewCheckClusterHostNameRequest() (request *CheckClusterHostNameRequest) {
    request = &CheckClusterHostNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CheckClusterHostName")
    return
}

func NewCheckClusterHostNameResponse() (response *CheckClusterHostNameResponse) {
    response = &CheckClusterHostNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查集群节点主机名称，判断节点主机名称是否符合规则，是否可以加入集群。
func (c *Client) CheckClusterHostName(request *CheckClusterHostNameRequest) (response *CheckClusterHostNameResponse, err error) {
    if request == nil {
        request = NewCheckClusterHostNameRequest()
    }
    response = NewCheckClusterHostNameResponse()
    err = c.Send(request, response)
    return
}

func NewCheckClusterImageRequest() (request *CheckClusterImageRequest) {
    request = &CheckClusterImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CheckClusterImage")
    return
}

func NewCheckClusterImageResponse() (response *CheckClusterImageResponse) {
    response = &CheckClusterImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查镜像是否支持设置为集群镜像
func (c *Client) CheckClusterImage(request *CheckClusterImageRequest) (response *CheckClusterImageResponse, err error) {
    if request == nil {
        request = NewCheckClusterImageRequest()
    }
    response = NewCheckClusterImageResponse()
    err = c.Send(request, response)
    return
}

func NewCheckInstancesUpgradeAbleRequest() (request *CheckInstancesUpgradeAbleRequest) {
    request = &CheckInstancesUpgradeAbleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CheckInstancesUpgradeAble")
    return
}

func NewCheckInstancesUpgradeAbleResponse() (response *CheckInstancesUpgradeAbleResponse) {
    response = &CheckInstancesUpgradeAbleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查给定节点列表中哪些是可升级的 
func (c *Client) CheckInstancesUpgradeAble(request *CheckInstancesUpgradeAbleRequest) (response *CheckInstancesUpgradeAbleResponse, err error) {
    if request == nil {
        request = NewCheckInstancesUpgradeAbleRequest()
    }
    response = NewCheckInstancesUpgradeAbleResponse()
    err = c.Send(request, response)
    return
}

func NewCheckLogCollectorHostPathRequest() (request *CheckLogCollectorHostPathRequest) {
    request = &CheckLogCollectorHostPathRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CheckLogCollectorHostPath")
    return
}

func NewCheckLogCollectorHostPathResponse() (response *CheckLogCollectorHostPathResponse) {
    response = &CheckLogCollectorHostPathResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查主机路径
func (c *Client) CheckLogCollectorHostPath(request *CheckLogCollectorHostPathRequest) (response *CheckLogCollectorHostPathResponse, err error) {
    if request == nil {
        request = NewCheckLogCollectorHostPathRequest()
    }
    response = NewCheckLogCollectorHostPathResponse()
    err = c.Send(request, response)
    return
}

func NewCheckLogCollectorNameRequest() (request *CheckLogCollectorNameRequest) {
    request = &CheckLogCollectorNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CheckLogCollectorName")
    return
}

func NewCheckLogCollectorNameResponse() (response *CheckLogCollectorNameResponse) {
    response = &CheckLogCollectorNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过名称检查日志采集规则是否存在
func (c *Client) CheckLogCollectorName(request *CheckLogCollectorNameRequest) (response *CheckLogCollectorNameResponse, err error) {
    if request == nil {
        request = NewCheckLogCollectorNameRequest()
    }
    response = NewCheckLogCollectorNameResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateCluster")
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterAsGroupRequest() (request *CreateClusterAsGroupRequest) {
    request = &CreateClusterAsGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterAsGroup")
    return
}

func NewCreateClusterAsGroupResponse() (response *CreateClusterAsGroupResponse) {
    response = &CreateClusterAsGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为已经存在的集群创建伸缩组
func (c *Client) CreateClusterAsGroup(request *CreateClusterAsGroupRequest) (response *CreateClusterAsGroupResponse, err error) {
    if request == nil {
        request = NewCreateClusterAsGroupRequest()
    }
    response = NewCreateClusterAsGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterEndpointRequest() (request *CreateClusterEndpointRequest) {
    request = &CreateClusterEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterEndpoint")
    return
}

func NewCreateClusterEndpointResponse() (response *CreateClusterEndpointResponse) {
    response = &CreateClusterEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建集群访问端口(独立集群开启内网/外网访问，托管集群支持开启内网访问)
func (c *Client) CreateClusterEndpoint(request *CreateClusterEndpointRequest) (response *CreateClusterEndpointResponse, err error) {
    if request == nil {
        request = NewCreateClusterEndpointRequest()
    }
    response = NewCreateClusterEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterEndpointVipRequest() (request *CreateClusterEndpointVipRequest) {
    request = &CreateClusterEndpointVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterEndpointVip")
    return
}

func NewCreateClusterEndpointVipResponse() (response *CreateClusterEndpointVipResponse) {
    response = &CreateClusterEndpointVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建托管集群外网访问端口（老的方式，仅支持托管集群外网端口）
func (c *Client) CreateClusterEndpointVip(request *CreateClusterEndpointVipRequest) (response *CreateClusterEndpointVipResponse, err error) {
    if request == nil {
        request = NewCreateClusterEndpointVipRequest()
    }
    response = NewCreateClusterEndpointVipResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterInstancesRequest() (request *CreateClusterInstancesRequest) {
    request = &CreateClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterInstances")
    return
}

func NewCreateClusterInstancesResponse() (response *CreateClusterInstancesResponse) {
    response = &CreateClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 扩展(新建)集群节点
func (c *Client) CreateClusterInstances(request *CreateClusterInstancesRequest) (response *CreateClusterInstancesResponse, err error) {
    if request == nil {
        request = NewCreateClusterInstancesRequest()
    }
    response = NewCreateClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRouteRequest() (request *CreateClusterRouteRequest) {
    request = &CreateClusterRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRoute")
    return
}

func NewCreateClusterRouteResponse() (response *CreateClusterRouteResponse) {
    response = &CreateClusterRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建集群路由
func (c *Client) CreateClusterRoute(request *CreateClusterRouteRequest) (response *CreateClusterRouteResponse, err error) {
    if request == nil {
        request = NewCreateClusterRouteRequest()
    }
    response = NewCreateClusterRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRouteTableRequest() (request *CreateClusterRouteTableRequest) {
    request = &CreateClusterRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRouteTable")
    return
}

func NewCreateClusterRouteTableResponse() (response *CreateClusterRouteTableResponse) {
    response = &CreateClusterRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建集群路由表
func (c *Client) CreateClusterRouteTable(request *CreateClusterRouteTableRequest) (response *CreateClusterRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateClusterRouteTableRequest()
    }
    response = NewCreateClusterRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEKSClusterRequest() (request *CreateEKSClusterRequest) {
    request = &CreateEKSClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateEKSCluster")
    return
}

func NewCreateEKSClusterResponse() (response *CreateEKSClusterResponse) {
    response = &CreateEKSClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建弹性集群
func (c *Client) CreateEKSCluster(request *CreateEKSClusterRequest) (response *CreateEKSClusterResponse, err error) {
    if request == nil {
        request = NewCreateEKSClusterRequest()
    }
    response = NewCreateEKSClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIndependentClusterRequest() (request *CreateIndependentClusterRequest) {
    request = &CreateIndependentClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateIndependentCluster")
    return
}

func NewCreateIndependentClusterResponse() (response *CreateIndependentClusterResponse) {
    response = &CreateIndependentClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建独立集群
func (c *Client) CreateIndependentCluster(request *CreateIndependentClusterRequest) (response *CreateIndependentClusterResponse, err error) {
    if request == nil {
        request = NewCreateIndependentClusterRequest()
    }
    response = NewCreateIndependentClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTKEClusterRequest() (request *CreateTKEClusterRequest) {
    request = &CreateTKEClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "CreateTKECluster")
    return
}

func NewCreateTKEClusterResponse() (response *CreateTKEClusterResponse) {
    response = &CreateTKEClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建TKE集群
func (c *Client) CreateTKECluster(request *CreateTKEClusterRequest) (response *CreateTKEClusterResponse, err error) {
    if request == nil {
        request = NewCreateTKEClusterRequest()
    }
    response = NewCreateTKEClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmPoliciesRequest() (request *DeleteAlarmPoliciesRequest) {
    request = &DeleteAlarmPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteAlarmPolicies")
    return
}

func NewDeleteAlarmPoliciesResponse() (response *DeleteAlarmPoliciesResponse) {
    response = &DeleteAlarmPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除告警策略，支持批量删除
func (c *Client) DeleteAlarmPolicies(request *DeleteAlarmPoliciesRequest) (response *DeleteAlarmPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmPoliciesRequest()
    }
    response = NewDeleteAlarmPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteCluster")
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除集群(YUNAPI V3版本)
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterAsGroupsRequest() (request *DeleteClusterAsGroupsRequest) {
    request = &DeleteClusterAsGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterAsGroups")
    return
}

func NewDeleteClusterAsGroupsResponse() (response *DeleteClusterAsGroupsResponse) {
    response = &DeleteClusterAsGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除集群伸缩组
func (c *Client) DeleteClusterAsGroups(request *DeleteClusterAsGroupsRequest) (response *DeleteClusterAsGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteClusterAsGroupsRequest()
    }
    response = NewDeleteClusterAsGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterCIDRFromCcnRequest() (request *DeleteClusterCIDRFromCcnRequest) {
    request = &DeleteClusterCIDRFromCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterCIDRFromCcn")
    return
}

func NewDeleteClusterCIDRFromCcnResponse() (response *DeleteClusterCIDRFromCcnResponse) {
    response = &DeleteClusterCIDRFromCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从云联网删除集群CIDR的路由
func (c *Client) DeleteClusterCIDRFromCcn(request *DeleteClusterCIDRFromCcnRequest) (response *DeleteClusterCIDRFromCcnResponse, err error) {
    if request == nil {
        request = NewDeleteClusterCIDRFromCcnRequest()
    }
    response = NewDeleteClusterCIDRFromCcnResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterCIDRFromVbcRequest() (request *DeleteClusterCIDRFromVbcRequest) {
    request = &DeleteClusterCIDRFromVbcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterCIDRFromVbc")
    return
}

func NewDeleteClusterCIDRFromVbcResponse() (response *DeleteClusterCIDRFromVbcResponse) {
    response = &DeleteClusterCIDRFromVbcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从云联网删除集群CIDR
func (c *Client) DeleteClusterCIDRFromVbc(request *DeleteClusterCIDRFromVbcRequest) (response *DeleteClusterCIDRFromVbcResponse, err error) {
    if request == nil {
        request = NewDeleteClusterCIDRFromVbcRequest()
    }
    response = NewDeleteClusterCIDRFromVbcResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterEndpointRequest() (request *DeleteClusterEndpointRequest) {
    request = &DeleteClusterEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterEndpoint")
    return
}

func NewDeleteClusterEndpointResponse() (response *DeleteClusterEndpointResponse) {
    response = &DeleteClusterEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除集群访问端口(独立集群开启内网/外网访问，托管集群支持开启内网访问)
func (c *Client) DeleteClusterEndpoint(request *DeleteClusterEndpointRequest) (response *DeleteClusterEndpointResponse, err error) {
    if request == nil {
        request = NewDeleteClusterEndpointRequest()
    }
    response = NewDeleteClusterEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterEndpointVipRequest() (request *DeleteClusterEndpointVipRequest) {
    request = &DeleteClusterEndpointVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterEndpointVip")
    return
}

func NewDeleteClusterEndpointVipResponse() (response *DeleteClusterEndpointVipResponse) {
    response = &DeleteClusterEndpointVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除托管集群外网访问端口（老的方式，仅支持托管集群外网端口）
func (c *Client) DeleteClusterEndpointVip(request *DeleteClusterEndpointVipRequest) (response *DeleteClusterEndpointVipResponse, err error) {
    if request == nil {
        request = NewDeleteClusterEndpointVipRequest()
    }
    response = NewDeleteClusterEndpointVipResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterInstancesRequest() (request *DeleteClusterInstancesRequest) {
    request = &DeleteClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterInstances")
    return
}

func NewDeleteClusterInstancesResponse() (response *DeleteClusterInstancesResponse) {
    response = &DeleteClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除集群中的实例
func (c *Client) DeleteClusterInstances(request *DeleteClusterInstancesRequest) (response *DeleteClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteClusterInstancesRequest()
    }
    response = NewDeleteClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRouteRequest() (request *DeleteClusterRouteRequest) {
    request = &DeleteClusterRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterRoute")
    return
}

func NewDeleteClusterRouteResponse() (response *DeleteClusterRouteResponse) {
    response = &DeleteClusterRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除集群路由
func (c *Client) DeleteClusterRoute(request *DeleteClusterRouteRequest) (response *DeleteClusterRouteResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRouteRequest()
    }
    response = NewDeleteClusterRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRouteTableRequest() (request *DeleteClusterRouteTableRequest) {
    request = &DeleteClusterRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterRouteTable")
    return
}

func NewDeleteClusterRouteTableResponse() (response *DeleteClusterRouteTableResponse) {
    response = &DeleteClusterRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除集群路由表
func (c *Client) DeleteClusterRouteTable(request *DeleteClusterRouteTableRequest) (response *DeleteClusterRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRouteTableRequest()
    }
    response = NewDeleteClusterRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEKSClusterRequest() (request *DeleteEKSClusterRequest) {
    request = &DeleteEKSClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DeleteEKSCluster")
    return
}

func NewDeleteEKSClusterResponse() (response *DeleteEKSClusterResponse) {
    response = &DeleteEKSClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除弹性集群(yunapiv3)
func (c *Client) DeleteEKSCluster(request *DeleteEKSClusterRequest) (response *DeleteEKSClusterResponse, err error) {
    if request == nil {
        request = NewDeleteEKSClusterRequest()
    }
    response = NewDeleteEKSClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmPoliciesRequest() (request *DescribeAlarmPoliciesRequest) {
    request = &DescribeAlarmPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeAlarmPolicies")
    return
}

func NewDescribeAlarmPoliciesResponse() (response *DescribeAlarmPoliciesResponse) {
    response = &DescribeAlarmPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警策略列表
func (c *Client) DescribeAlarmPolicies(request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPoliciesRequest()
    }
    response = NewDescribeAlarmPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableClusterVersionRequest() (request *DescribeAvailableClusterVersionRequest) {
    request = &DescribeAvailableClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableClusterVersion")
    return
}

func NewDescribeAvailableClusterVersionResponse() (response *DescribeAvailableClusterVersionResponse) {
    response = &DescribeAvailableClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群可以升级的所有版本
func (c *Client) DescribeAvailableClusterVersion(request *DescribeAvailableClusterVersionRequest) (response *DescribeAvailableClusterVersionResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableClusterVersionRequest()
    }
    response = NewDescribeAvailableClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnInstancesRequest() (request *DescribeCcnInstancesRequest) {
    request = &DescribeCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeCcnInstances")
    return
}

func NewDescribeCcnInstancesResponse() (response *DescribeCcnInstancesResponse) {
    response = &DescribeCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询vpc是否加入云联网
func (c *Client) DescribeCcnInstances(request *DescribeCcnInstancesRequest) (response *DescribeCcnInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnInstancesRequest()
    }
    response = NewDescribeCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnRoutesRequest() (request *DescribeCcnRoutesRequest) {
    request = &DescribeCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeCcnRoutes")
    return
}

func NewDescribeCcnRoutesResponse() (response *DescribeCcnRoutesResponse) {
    response = &DescribeCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询tke集群CIDR是否加入云联网
func (c *Client) DescribeCcnRoutes(request *DescribeCcnRoutesRequest) (response *DescribeCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnRoutesRequest()
    }
    response = NewDescribeCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClsLogSetsRequest() (request *DescribeClsLogSetsRequest) {
    request = &DescribeClsLogSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClsLogSets")
    return
}

func NewDescribeClsLogSetsResponse() (response *DescribeClsLogSetsResponse) {
    response = &DescribeClsLogSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 列出CLS日志集
func (c *Client) DescribeClsLogSets(request *DescribeClsLogSetsRequest) (response *DescribeClsLogSetsResponse, err error) {
    if request == nil {
        request = NewDescribeClsLogSetsRequest()
    }
    response = NewDescribeClsLogSetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClsLogTopicsRequest() (request *DescribeClsLogTopicsRequest) {
    request = &DescribeClsLogTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClsLogTopics")
    return
}

func NewDescribeClsLogTopicsResponse() (response *DescribeClsLogTopicsResponse) {
    response = &DescribeClsLogTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 列出CLS日志主题
func (c *Client) DescribeClsLogTopics(request *DescribeClsLogTopicsRequest) (response *DescribeClsLogTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeClsLogTopicsRequest()
    }
    response = NewDescribeClsLogTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAsGroupOptionRequest() (request *DescribeClusterAsGroupOptionRequest) {
    request = &DescribeClusterAsGroupOptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAsGroupOption")
    return
}

func NewDescribeClusterAsGroupOptionResponse() (response *DescribeClusterAsGroupOptionResponse) {
    response = &DescribeClusterAsGroupOptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 集群弹性伸缩配置
func (c *Client) DescribeClusterAsGroupOption(request *DescribeClusterAsGroupOptionRequest) (response *DescribeClusterAsGroupOptionResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAsGroupOptionRequest()
    }
    response = NewDescribeClusterAsGroupOptionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAsGroupsRequest() (request *DescribeClusterAsGroupsRequest) {
    request = &DescribeClusterAsGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAsGroups")
    return
}

func NewDescribeClusterAsGroupsResponse() (response *DescribeClusterAsGroupsResponse) {
    response = &DescribeClusterAsGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 集群关联的伸缩组列表
func (c *Client) DescribeClusterAsGroups(request *DescribeClusterAsGroupsRequest) (response *DescribeClusterAsGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAsGroupsRequest()
    }
    response = NewDescribeClusterAsGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAvailableExtraArgsRequest() (request *DescribeClusterAvailableExtraArgsRequest) {
    request = &DescribeClusterAvailableExtraArgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAvailableExtraArgs")
    return
}

func NewDescribeClusterAvailableExtraArgsResponse() (response *DescribeClusterAvailableExtraArgsResponse) {
    response = &DescribeClusterAvailableExtraArgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群可用的自定义参数
func (c *Client) DescribeClusterAvailableExtraArgs(request *DescribeClusterAvailableExtraArgsRequest) (response *DescribeClusterAvailableExtraArgsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAvailableExtraArgsRequest()
    }
    response = NewDescribeClusterAvailableExtraArgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterCreateProgressRequest() (request *DescribeClusterCreateProgressRequest) {
    request = &DescribeClusterCreateProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterCreateProgress")
    return
}

func NewDescribeClusterCreateProgressResponse() (response *DescribeClusterCreateProgressResponse) {
    response = &DescribeClusterCreateProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群创建进度
func (c *Client) DescribeClusterCreateProgress(request *DescribeClusterCreateProgressRequest) (response *DescribeClusterCreateProgressResponse, err error) {
    if request == nil {
        request = NewDescribeClusterCreateProgressRequest()
    }
    response = NewDescribeClusterCreateProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointStatusRequest() (request *DescribeClusterEndpointStatusRequest) {
    request = &DescribeClusterEndpointStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpointStatus")
    return
}

func NewDescribeClusterEndpointStatusResponse() (response *DescribeClusterEndpointStatusResponse) {
    response = &DescribeClusterEndpointStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群访问端口状态(独立集群开启内网/外网访问，托管集群支持开启内网访问)
func (c *Client) DescribeClusterEndpointStatus(request *DescribeClusterEndpointStatusRequest) (response *DescribeClusterEndpointStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointStatusRequest()
    }
    response = NewDescribeClusterEndpointStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointVipStatusRequest() (request *DescribeClusterEndpointVipStatusRequest) {
    request = &DescribeClusterEndpointVipStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpointVipStatus")
    return
}

func NewDescribeClusterEndpointVipStatusResponse() (response *DescribeClusterEndpointVipStatusResponse) {
    response = &DescribeClusterEndpointVipStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群开启端口流程状态(仅支持托管集群外网端口)
func (c *Client) DescribeClusterEndpointVipStatus(request *DescribeClusterEndpointVipStatusRequest) (response *DescribeClusterEndpointVipStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointVipStatusRequest()
    }
    response = NewDescribeClusterEndpointVipStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterExtraArgsRequest() (request *DescribeClusterExtraArgsRequest) {
    request = &DescribeClusterExtraArgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterExtraArgs")
    return
}

func NewDescribeClusterExtraArgsResponse() (response *DescribeClusterExtraArgsResponse) {
    response = &DescribeClusterExtraArgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群自定义参数
func (c *Client) DescribeClusterExtraArgs(request *DescribeClusterExtraArgsRequest) (response *DescribeClusterExtraArgsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterExtraArgsRequest()
    }
    response = NewDescribeClusterExtraArgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterHealthyStatusRequest() (request *DescribeClusterHealthyStatusRequest) {
    request = &DescribeClusterHealthyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterHealthyStatus")
    return
}

func NewDescribeClusterHealthyStatusResponse() (response *DescribeClusterHealthyStatusResponse) {
    response = &DescribeClusterHealthyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 描述集群目前的健康状态 
func (c *Client) DescribeClusterHealthyStatus(request *DescribeClusterHealthyStatusRequest) (response *DescribeClusterHealthyStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterHealthyStatusRequest()
    }
    response = NewDescribeClusterHealthyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInspectionOverviewsRequest() (request *DescribeClusterInspectionOverviewsRequest) {
    request = &DescribeClusterInspectionOverviewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInspectionOverviews")
    return
}

func NewDescribeClusterInspectionOverviewsResponse() (response *DescribeClusterInspectionOverviewsResponse) {
    response = &DescribeClusterInspectionOverviewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获得集群巡检报告列表
func (c *Client) DescribeClusterInspectionOverviews(request *DescribeClusterInspectionOverviewsRequest) (response *DescribeClusterInspectionOverviewsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInspectionOverviewsRequest()
    }
    response = NewDescribeClusterInspectionOverviewsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInspectionReportRequest() (request *DescribeClusterInspectionReportRequest) {
    request = &DescribeClusterInspectionReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInspectionReport")
    return
}

func NewDescribeClusterInspectionReportResponse() (response *DescribeClusterInspectionReportResponse) {
    response = &DescribeClusterInspectionReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获得集群巡检报告页详情
func (c *Client) DescribeClusterInspectionReport(request *DescribeClusterInspectionReportRequest) (response *DescribeClusterInspectionReportResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInspectionReportRequest()
    }
    response = NewDescribeClusterInspectionReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInspectionsRequest() (request *DescribeClusterInspectionsRequest) {
    request = &DescribeClusterInspectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInspections")
    return
}

func NewDescribeClusterInspectionsResponse() (response *DescribeClusterInspectionsResponse) {
    response = &DescribeClusterInspectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 集群巡检概览
func (c *Client) DescribeClusterInspections(request *DescribeClusterInspectionsRequest) (response *DescribeClusterInspectionsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInspectionsRequest()
    }
    response = NewDescribeClusterInspectionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstanceIdsRequest() (request *DescribeClusterInstanceIdsRequest) {
    request = &DescribeClusterInstanceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInstanceIds")
    return
}

func NewDescribeClusterInstanceIdsResponse() (response *DescribeClusterInstanceIdsResponse) {
    response = &DescribeClusterInstanceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群节点ID列表【仅内部使用】
func (c *Client) DescribeClusterInstanceIds(request *DescribeClusterInstanceIdsRequest) (response *DescribeClusterInstanceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstanceIdsRequest()
    }
    response = NewDescribeClusterInstanceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInstances")
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  查询集群下节点实例信息 
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterPodsRequest() (request *DescribeClusterPodsRequest) {
    request = &DescribeClusterPodsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterPods")
    return
}

func NewDescribeClusterPodsResponse() (response *DescribeClusterPodsResponse) {
    response = &DescribeClusterPodsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口获取集群内Pod相关的详细描述信息，参考kubernetes API获取pod，只对内部短期使用
func (c *Client) DescribeClusterPods(request *DescribeClusterPodsRequest) (response *DescribeClusterPodsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPodsRequest()
    }
    response = NewDescribeClusterPodsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRouteTablesRequest() (request *DescribeClusterRouteTablesRequest) {
    request = &DescribeClusterRouteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRouteTables")
    return
}

func NewDescribeClusterRouteTablesResponse() (response *DescribeClusterRouteTablesResponse) {
    response = &DescribeClusterRouteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群路由表
func (c *Client) DescribeClusterRouteTables(request *DescribeClusterRouteTablesRequest) (response *DescribeClusterRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRouteTablesRequest()
    }
    response = NewDescribeClusterRouteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRoutesRequest() (request *DescribeClusterRoutesRequest) {
    request = &DescribeClusterRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRoutes")
    return
}

func NewDescribeClusterRoutesResponse() (response *DescribeClusterRoutesResponse) {
    response = &DescribeClusterRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群路由
func (c *Client) DescribeClusterRoutes(request *DescribeClusterRoutesRequest) (response *DescribeClusterRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRoutesRequest()
    }
    response = NewDescribeClusterRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterSecurityRequest() (request *DescribeClusterSecurityRequest) {
    request = &DescribeClusterSecurityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterSecurity")
    return
}

func NewDescribeClusterSecurityResponse() (response *DescribeClusterSecurityResponse) {
    response = &DescribeClusterSecurityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 集群的密钥信息
func (c *Client) DescribeClusterSecurity(request *DescribeClusterSecurityRequest) (response *DescribeClusterSecurityResponse, err error) {
    if request == nil {
        request = NewDescribeClusterSecurityRequest()
    }
    response = NewDescribeClusterSecurityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterServicesRequest() (request *DescribeClusterServicesRequest) {
    request = &DescribeClusterServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterServices")
    return
}

func NewDescribeClusterServicesResponse() (response *DescribeClusterServicesResponse) {
    response = &DescribeClusterServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口获取集群内Service相关的详细描述信息，参考kubernetes API获取Service，只对内部短期使用
func (c *Client) DescribeClusterServices(request *DescribeClusterServicesRequest) (response *DescribeClusterServicesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterServicesRequest()
    }
    response = NewDescribeClusterServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterStatusRequest() (request *DescribeClusterStatusRequest) {
    request = &DescribeClusterStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterStatus")
    return
}

func NewDescribeClusterStatusResponse() (response *DescribeClusterStatusResponse) {
    response = &DescribeClusterStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看集群状态列表【非对外接口】
func (c *Client) DescribeClusterStatus(request *DescribeClusterStatusRequest) (response *DescribeClusterStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterStatusRequest()
    }
    response = NewDescribeClusterStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusters")
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群列表
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersResourceStatusRequest() (request *DescribeClustersResourceStatusRequest) {
    request = &DescribeClustersResourceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClustersResourceStatus")
    return
}

func NewDescribeClustersResourceStatusResponse() (response *DescribeClustersResourceStatusResponse) {
    response = &DescribeClustersResourceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群资源状态
func (c *Client) DescribeClustersResourceStatus(request *DescribeClustersResourceStatusRequest) (response *DescribeClustersResourceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClustersResourceStatusRequest()
    }
    response = NewDescribeClustersResourceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEKSClusterCredentialRequest() (request *DescribeEKSClusterCredentialRequest) {
    request = &DescribeEKSClusterCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusterCredential")
    return
}

func NewDescribeEKSClusterCredentialResponse() (response *DescribeEKSClusterCredentialResponse) {
    response = &DescribeEKSClusterCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取弹性容器集群的接入认证信息
func (c *Client) DescribeEKSClusterCredential(request *DescribeEKSClusterCredentialRequest) (response *DescribeEKSClusterCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeEKSClusterCredentialRequest()
    }
    response = NewDescribeEKSClusterCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEKSClusterStatusRequest() (request *DescribeEKSClusterStatusRequest) {
    request = &DescribeEKSClusterStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusterStatus")
    return
}

func NewDescribeEKSClusterStatusResponse() (response *DescribeEKSClusterStatusResponse) {
    response = &DescribeEKSClusterStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取弹性容器集群的当前状态以及过程信息
func (c *Client) DescribeEKSClusterStatus(request *DescribeEKSClusterStatusRequest) (response *DescribeEKSClusterStatusResponse, err error) {
    if request == nil {
        request = NewDescribeEKSClusterStatusRequest()
    }
    response = NewDescribeEKSClusterStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEKSClustersRequest() (request *DescribeEKSClustersRequest) {
    request = &DescribeEKSClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusters")
    return
}

func NewDescribeEKSClustersResponse() (response *DescribeEKSClustersResponse) {
    response = &DescribeEKSClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询弹性集群列表
func (c *Client) DescribeEKSClusters(request *DescribeEKSClustersRequest) (response *DescribeEKSClustersResponse, err error) {
    if request == nil {
        request = NewDescribeEKSClustersRequest()
    }
    response = NewDescribeEKSClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExistedInstancesRequest() (request *DescribeExistedInstancesRequest) {
    request = &DescribeExistedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeExistedInstances")
    return
}

func NewDescribeExistedInstancesResponse() (response *DescribeExistedInstancesResponse) {
    response = &DescribeExistedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询已经存在的节点，判断是否可以加入集群
func (c *Client) DescribeExistedInstances(request *DescribeExistedInstancesRequest) (response *DescribeExistedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeExistedInstancesRequest()
    }
    response = NewDescribeExistedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowIdStatusRequest() (request *DescribeFlowIdStatusRequest) {
    request = &DescribeFlowIdStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeFlowIdStatus")
    return
}

func NewDescribeFlowIdStatusResponse() (response *DescribeFlowIdStatusResponse) {
    response = &DescribeFlowIdStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群开启端口流程状态(仅支持托管集群外网端口)
func (c *Client) DescribeFlowIdStatus(request *DescribeFlowIdStatusRequest) (response *DescribeFlowIdStatusResponse, err error) {
    if request == nil {
        request = NewDescribeFlowIdStatusRequest()
    }
    response = NewDescribeFlowIdStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeImages")
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取镜像信息
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceCreateProgressRequest() (request *DescribeInstanceCreateProgressRequest) {
    request = &DescribeInstanceCreateProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeInstanceCreateProgress")
    return
}

func NewDescribeInstanceCreateProgressResponse() (response *DescribeInstanceCreateProgressResponse) {
    response = &DescribeInstanceCreateProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取节点创建进度
func (c *Client) DescribeInstanceCreateProgress(request *DescribeInstanceCreateProgressRequest) (response *DescribeInstanceCreateProgressResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceCreateProgressRequest()
    }
    response = NewDescribeInstanceCreateProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesVersionRequest() (request *DescribeInstancesVersionRequest) {
    request = &DescribeInstancesVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeInstancesVersion")
    return
}

func NewDescribeInstancesVersionResponse() (response *DescribeInstancesVersionResponse) {
    response = &DescribeInstancesVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// worker节点的版本统计
func (c *Client) DescribeInstancesVersion(request *DescribeInstancesVersionRequest) (response *DescribeInstancesVersionResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesVersionRequest()
    }
    response = NewDescribeInstancesVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaRequest() (request *DescribeQuotaRequest) {
    request = &DescribeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeQuota")
    return
}

func NewDescribeQuotaResponse() (response *DescribeQuotaResponse) {
    response = &DescribeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群配额
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaRequest()
    }
    response = NewDescribeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取容器服务支持的所有地域
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteTableConflictsRequest() (request *DescribeRouteTableConflictsRequest) {
    request = &DescribeRouteTableConflictsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeRouteTableConflicts")
    return
}

func NewDescribeRouteTableConflictsResponse() (response *DescribeRouteTableConflictsResponse) {
    response = &DescribeRouteTableConflictsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询路由表冲突列表
func (c *Client) DescribeRouteTableConflicts(request *DescribeRouteTableConflictsRequest) (response *DescribeRouteTableConflictsResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTableConflictsRequest()
    }
    response = NewDescribeRouteTableConflictsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpgradeClusterProgressRequest() (request *DescribeUpgradeClusterProgressRequest) {
    request = &DescribeUpgradeClusterProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeUpgradeClusterProgress")
    return
}

func NewDescribeUpgradeClusterProgressResponse() (response *DescribeUpgradeClusterProgressResponse) {
    response = &DescribeUpgradeClusterProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群升级进度
func (c *Client) DescribeUpgradeClusterProgress(request *DescribeUpgradeClusterProgressRequest) (response *DescribeUpgradeClusterProgressResponse, err error) {
    if request == nil {
        request = NewDescribeUpgradeClusterProgressRequest()
    }
    response = NewDescribeUpgradeClusterProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVersionsRequest() (request *DescribeVersionsRequest) {
    request = &DescribeVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DescribeVersions")
    return
}

func NewDescribeVersionsResponse() (response *DescribeVersionsResponse) {
    response = &DescribeVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群版本信息
func (c *Client) DescribeVersions(request *DescribeVersionsRequest) (response *DescribeVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeVersionsRequest()
    }
    response = NewDescribeVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDrainClusterNodeRequest() (request *DrainClusterNodeRequest) {
    request = &DrainClusterNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "DrainClusterNode")
    return
}

func NewDrainClusterNodeResponse() (response *DrainClusterNodeResponse) {
    response = &DrainClusterNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 驱逐集群中的节点
func (c *Client) DrainClusterNode(request *DrainClusterNodeRequest) (response *DrainClusterNodeResponse, err error) {
    if request == nil {
        request = NewDrainClusterNodeRequest()
    }
    response = NewDrainClusterNodeResponse()
    err = c.Send(request, response)
    return
}

func NewEnableVpcPeerClusterRoutesRequest() (request *EnableVpcPeerClusterRoutesRequest) {
    request = &EnableVpcPeerClusterRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "EnableVpcPeerClusterRoutes")
    return
}

func NewEnableVpcPeerClusterRoutesResponse() (response *EnableVpcPeerClusterRoutesResponse) {
    response = &EnableVpcPeerClusterRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启动对等连接容器路由
func (c *Client) EnableVpcPeerClusterRoutes(request *EnableVpcPeerClusterRoutesRequest) (response *EnableVpcPeerClusterRoutesResponse, err error) {
    if request == nil {
        request = NewEnableVpcPeerClusterRoutesRequest()
    }
    response = NewEnableVpcPeerClusterRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewForwardClusterRequestRequest() (request *ForwardClusterRequestRequest) {
    request = &ForwardClusterRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ForwardClusterRequest")
    return
}

func NewForwardClusterRequestResponse() (response *ForwardClusterRequestResponse) {
    response = &ForwardClusterRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询、新增、删除、编辑TKE集群内资源
func (c *Client) ForwardClusterRequest(request *ForwardClusterRequestRequest) (response *ForwardClusterRequestResponse, err error) {
    if request == nil {
        request = NewForwardClusterRequestRequest()
    }
    response = NewForwardClusterRequestResponse()
    err = c.Send(request, response)
    return
}

func NewForwardRequestRequest() (request *ForwardRequestRequest) {
    request = &ForwardRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ForwardRequest")
    return
}

func NewForwardRequestResponse() (response *ForwardRequestResponse) {
    response = &ForwardRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// YUNAPI 转发请求给TKE APIServer接口
func (c *Client) ForwardRequest(request *ForwardRequestRequest) (response *ForwardRequestResponse, err error) {
    if request == nil {
        request = NewForwardRequestRequest()
    }
    response = NewForwardRequestResponse()
    err = c.Send(request, response)
    return
}

func NewGetUpgradeClusterProgressRequest() (request *GetUpgradeClusterProgressRequest) {
    request = &GetUpgradeClusterProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "GetUpgradeClusterProgress")
    return
}

func NewGetUpgradeClusterProgressResponse() (response *GetUpgradeClusterProgressResponse) {
    response = &GetUpgradeClusterProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回当前集群升级进度
func (c *Client) GetUpgradeClusterProgress(request *GetUpgradeClusterProgressRequest) (response *GetUpgradeClusterProgressResponse, err error) {
    if request == nil {
        request = NewGetUpgradeClusterProgressRequest()
    }
    response = NewGetUpgradeClusterProgressResponse()
    err = c.Send(request, response)
    return
}

func NewGetUpgradeInstanceProgressRequest() (request *GetUpgradeInstanceProgressRequest) {
    request = &GetUpgradeInstanceProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "GetUpgradeInstanceProgress")
    return
}

func NewGetUpgradeInstanceProgressResponse() (response *GetUpgradeInstanceProgressResponse) {
    response = &GetUpgradeInstanceProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获得节点升级当前的进度 
func (c *Client) GetUpgradeInstanceProgress(request *GetUpgradeInstanceProgressRequest) (response *GetUpgradeInstanceProgressResponse, err error) {
    if request == nil {
        request = NewGetUpgradeInstanceProgressRequest()
    }
    response = NewGetUpgradeInstanceProgressResponse()
    err = c.Send(request, response)
    return
}

func NewGetVbcInstanceRequest() (request *GetVbcInstanceRequest) {
    request = &GetVbcInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "GetVbcInstance")
    return
}

func NewGetVbcInstanceResponse() (response *GetVbcInstanceResponse) {
    response = &GetVbcInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询vpc是否加入云联网
func (c *Client) GetVbcInstance(request *GetVbcInstanceRequest) (response *GetVbcInstanceResponse, err error) {
    if request == nil {
        request = NewGetVbcInstanceRequest()
    }
    response = NewGetVbcInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewGetVbcRouteRequest() (request *GetVbcRouteRequest) {
    request = &GetVbcRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "GetVbcRoute")
    return
}

func NewGetVbcRouteResponse() (response *GetVbcRouteResponse) {
    response = &GetVbcRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询tke集群cidr是否加入云联网
func (c *Client) GetVbcRoute(request *GetVbcRouteRequest) (response *GetVbcRouteResponse, err error) {
    if request == nil {
        request = NewGetVbcRouteRequest()
    }
    response = NewGetVbcRouteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyRequest() (request *ModifyAlarmPolicyRequest) {
    request = &ModifyAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ModifyAlarmPolicy")
    return
}

func NewModifyAlarmPolicyResponse() (response *ModifyAlarmPolicyResponse) {
    response = &ModifyAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改告警策略
func (c *Client) ModifyAlarmPolicy(request *ModifyAlarmPolicyRequest) (response *ModifyAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyRequest()
    }
    response = NewModifyAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAsGroupAttributeRequest() (request *ModifyClusterAsGroupAttributeRequest) {
    request = &ModifyClusterAsGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAsGroupAttribute")
    return
}

func NewModifyClusterAsGroupAttributeResponse() (response *ModifyClusterAsGroupAttributeResponse) {
    response = &ModifyClusterAsGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改集群伸缩组属性
func (c *Client) ModifyClusterAsGroupAttribute(request *ModifyClusterAsGroupAttributeRequest) (response *ModifyClusterAsGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAsGroupAttributeRequest()
    }
    response = NewModifyClusterAsGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAsGroupOptionAttributeRequest() (request *ModifyClusterAsGroupOptionAttributeRequest) {
    request = &ModifyClusterAsGroupOptionAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAsGroupOptionAttribute")
    return
}

func NewModifyClusterAsGroupOptionAttributeResponse() (response *ModifyClusterAsGroupOptionAttributeResponse) {
    response = &ModifyClusterAsGroupOptionAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改集群弹性伸缩属性
func (c *Client) ModifyClusterAsGroupOptionAttribute(request *ModifyClusterAsGroupOptionAttributeRequest) (response *ModifyClusterAsGroupOptionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAsGroupOptionAttributeRequest()
    }
    response = NewModifyClusterAsGroupOptionAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAttributeRequest() (request *ModifyClusterAttributeRequest) {
    request = &ModifyClusterAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAttribute")
    return
}

func NewModifyClusterAttributeResponse() (response *ModifyClusterAttributeResponse) {
    response = &ModifyClusterAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改集群属性
func (c *Client) ModifyClusterAttribute(request *ModifyClusterAttributeRequest) (response *ModifyClusterAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAttributeRequest()
    }
    response = NewModifyClusterAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterEndpointSPRequest() (request *ModifyClusterEndpointSPRequest) {
    request = &ModifyClusterEndpointSPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterEndpointSP")
    return
}

func NewModifyClusterEndpointSPResponse() (response *ModifyClusterEndpointSPResponse) {
    response = &ModifyClusterEndpointSPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改托管集群外网端口的安全策略（老的方式，仅支持托管集群外网端口）
func (c *Client) ModifyClusterEndpointSP(request *ModifyClusterEndpointSPRequest) (response *ModifyClusterEndpointSPResponse, err error) {
    if request == nil {
        request = NewModifyClusterEndpointSPRequest()
    }
    response = NewModifyClusterEndpointSPResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterImageRequest() (request *ModifyClusterImageRequest) {
    request = &ModifyClusterImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterImage")
    return
}

func NewModifyClusterImageResponse() (response *ModifyClusterImageResponse) {
    response = &ModifyClusterImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改集群镜像
func (c *Client) ModifyClusterImage(request *ModifyClusterImageRequest) (response *ModifyClusterImageResponse, err error) {
    if request == nil {
        request = NewModifyClusterImageRequest()
    }
    response = NewModifyClusterImageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterInspectionRequest() (request *ModifyClusterInspectionRequest) {
    request = &ModifyClusterInspectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterInspection")
    return
}

func NewModifyClusterInspectionResponse() (response *ModifyClusterInspectionResponse) {
    response = &ModifyClusterInspectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新集群巡检配置
func (c *Client) ModifyClusterInspection(request *ModifyClusterInspectionRequest) (response *ModifyClusterInspectionResponse, err error) {
    if request == nil {
        request = NewModifyClusterInspectionRequest()
    }
    response = NewModifyClusterInspectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterUpgradingStateRequest() (request *ModifyClusterUpgradingStateRequest) {
    request = &ModifyClusterUpgradingStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterUpgradingState")
    return
}

func NewModifyClusterUpgradingStateResponse() (response *ModifyClusterUpgradingStateResponse) {
    response = &ModifyClusterUpgradingStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 暂停或者取消集群升级 
func (c *Client) ModifyClusterUpgradingState(request *ModifyClusterUpgradingStateRequest) (response *ModifyClusterUpgradingStateResponse, err error) {
    if request == nil {
        request = NewModifyClusterUpgradingStateRequest()
    }
    response = NewModifyClusterUpgradingStateResponse()
    err = c.Send(request, response)
    return
}

func NewOpUpgradeClusterInstancesRequest() (request *OpUpgradeClusterInstancesRequest) {
    request = &OpUpgradeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "OpUpgradeClusterInstances")
    return
}

func NewOpUpgradeClusterInstancesResponse() (response *OpUpgradeClusterInstancesResponse) {
    response = &OpUpgradeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于控制节点升级任务 
func (c *Client) OpUpgradeClusterInstances(request *OpUpgradeClusterInstancesRequest) (response *OpUpgradeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewOpUpgradeClusterInstancesRequest()
    }
    response = NewOpUpgradeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewPauseClusterInstancesRequest() (request *PauseClusterInstancesRequest) {
    request = &PauseClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "PauseClusterInstances")
    return
}

func NewPauseClusterInstancesResponse() (response *PauseClusterInstancesResponse) {
    response = &PauseClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 暂停集群节点升级, API 3.0
func (c *Client) PauseClusterInstances(request *PauseClusterInstancesRequest) (response *PauseClusterInstancesResponse, err error) {
    if request == nil {
        request = NewPauseClusterInstancesRequest()
    }
    response = NewPauseClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResumeClusterInstancesRequest() (request *ResumeClusterInstancesRequest) {
    request = &ResumeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ResumeClusterInstances")
    return
}

func NewResumeClusterInstancesResponse() (response *ResumeClusterInstancesResponse) {
    response = &ResumeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 恢复集群节点升级，API 3.0
func (c *Client) ResumeClusterInstances(request *ResumeClusterInstancesRequest) (response *ResumeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewResumeClusterInstancesRequest()
    }
    response = NewResumeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRunClusterInspectionsRequest() (request *RunClusterInspectionsRequest) {
    request = &RunClusterInspectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "RunClusterInspections")
    return
}

func NewRunClusterInspectionsResponse() (response *RunClusterInspectionsResponse) {
    response = &RunClusterInspectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 触发一次集群巡检。
func (c *Client) RunClusterInspections(request *RunClusterInspectionsRequest) (response *RunClusterInspectionsResponse, err error) {
    if request == nil {
        request = NewRunClusterInspectionsRequest()
    }
    response = NewRunClusterInspectionsResponse()
    err = c.Send(request, response)
    return
}

func NewServiceMeshForwardRequestRequest() (request *ServiceMeshForwardRequestRequest) {
    request = &ServiceMeshForwardRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "ServiceMeshForwardRequest")
    return
}

func NewServiceMeshForwardRequestResponse() (response *ServiceMeshForwardRequestResponse) {
    response = &ServiceMeshForwardRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务网格代理转发
func (c *Client) ServiceMeshForwardRequest(request *ServiceMeshForwardRequestRequest) (response *ServiceMeshForwardRequestResponse, err error) {
    if request == nil {
        request = NewServiceMeshForwardRequestRequest()
    }
    response = NewServiceMeshForwardRequestResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateClusterInstancesRequest() (request *UpdateClusterInstancesRequest) {
    request = &UpdateClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterInstances")
    return
}

func NewUpdateClusterInstancesResponse() (response *UpdateClusterInstancesResponse) {
    response = &UpdateClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 集群节点k8s版本升级，API 3.0
func (c *Client) UpdateClusterInstances(request *UpdateClusterInstancesRequest) (response *UpdateClusterInstancesResponse, err error) {
    if request == nil {
        request = NewUpdateClusterInstancesRequest()
    }
    response = NewUpdateClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateClusterVersionRequest() (request *UpdateClusterVersionRequest) {
    request = &UpdateClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterVersion")
    return
}

func NewUpdateClusterVersionResponse() (response *UpdateClusterVersionResponse) {
    response = &UpdateClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 升级集群 Master 组件到指定版本
func (c *Client) UpdateClusterVersion(request *UpdateClusterVersionRequest) (response *UpdateClusterVersionResponse, err error) {
    if request == nil {
        request = NewUpdateClusterVersionRequest()
    }
    response = NewUpdateClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEKSClusterRequest() (request *UpdateEKSClusterRequest) {
    request = &UpdateEKSClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "UpdateEKSCluster")
    return
}

func NewUpdateEKSClusterResponse() (response *UpdateEKSClusterResponse) {
    response = &UpdateEKSClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改弹性集群名称等属性 
func (c *Client) UpdateEKSCluster(request *UpdateEKSClusterRequest) (response *UpdateEKSClusterResponse, err error) {
    if request == nil {
        request = NewUpdateEKSClusterRequest()
    }
    response = NewUpdateEKSClusterResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeClusterInstancesRequest() (request *UpgradeClusterInstancesRequest) {
    request = &UpgradeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tke", APIVersion, "UpgradeClusterInstances")
    return
}

func NewUpgradeClusterInstancesResponse() (response *UpgradeClusterInstancesResponse) {
    response = &UpgradeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 给集群的一批work节点进行升级 
func (c *Client) UpgradeClusterInstances(request *UpgradeClusterInstancesRequest) (response *UpgradeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewUpgradeClusterInstancesRequest()
    }
    response = NewUpgradeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}
