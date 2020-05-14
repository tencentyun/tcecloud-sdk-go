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

package v20180724

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-24"

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


func NewAmpFrontTunnelRequest() (request *AmpFrontTunnelRequest) {
    request = &AmpFrontTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "AmpFrontTunnel")
    return
}

func NewAmpFrontTunnelResponse() (response *AmpFrontTunnelResponse) {
    response = &AmpFrontTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Amp系统前端统一接入API
func (c *Client) AmpFrontTunnel(request *AmpFrontTunnelRequest) (response *AmpFrontTunnelResponse, err error) {
    if request == nil {
        request = NewAmpFrontTunnelRequest()
    }
    response = NewAmpFrontTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewArgusFrontTunnelRequest() (request *ArgusFrontTunnelRequest) {
    request = &ArgusFrontTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ArgusFrontTunnel")
    return
}

func NewArgusFrontTunnelResponse() (response *ArgusFrontTunnelResponse) {
    response = &ArgusFrontTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus系统前端统一接入API
func (c *Client) ArgusFrontTunnel(request *ArgusFrontTunnelRequest) (response *ArgusFrontTunnelResponse, err error) {
    if request == nil {
        request = NewArgusFrontTunnelRequest()
    }
    response = NewArgusFrontTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewBindingPolicyObjectRequest() (request *BindingPolicyObjectRequest) {
    request = &BindingPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "BindingPolicyObject")
    return
}

func NewBindingPolicyObjectResponse() (response *BindingPolicyObjectResponse) {
    response = &BindingPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将告警策略绑定到特定对象
func (c *Client) BindingPolicyObject(request *BindingPolicyObjectRequest) (response *BindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewBindingPolicyObjectRequest()
    }
    response = NewBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewCLMDescribeMetricSetsRequest() (request *CLMDescribeMetricSetsRequest) {
    request = &CLMDescribeMetricSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CLMDescribeMetricSets")
    return
}

func NewCLMDescribeMetricSetsResponse() (response *CLMDescribeMetricSetsResponse) {
    response = &CLMDescribeMetricSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取指标集列表
func (c *Client) CLMDescribeMetricSets(request *CLMDescribeMetricSetsRequest) (response *CLMDescribeMetricSetsResponse, err error) {
    if request == nil {
        request = NewCLMDescribeMetricSetsRequest()
    }
    response = NewCLMDescribeMetricSetsResponse()
    err = c.Send(request, response)
    return
}

func NewClonePolicyGroupRequest() (request *ClonePolicyGroupRequest) {
    request = &ClonePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ClonePolicyGroup")
    return
}

func NewClonePolicyGroupResponse() (response *ClonePolicyGroupResponse) {
    response = &ClonePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 复制基础告警策略组
func (c *Client) ClonePolicyGroup(request *ClonePolicyGroupRequest) (response *ClonePolicyGroupResponse, err error) {
    if request == nil {
        request = NewClonePolicyGroupRequest()
    }
    response = NewClonePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCopyConditionsTemplateRequest() (request *CopyConditionsTemplateRequest) {
    request = &CopyConditionsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CopyConditionsTemplate")
    return
}

func NewCopyConditionsTemplateResponse() (response *CopyConditionsTemplateResponse) {
    response = &CopyConditionsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 复制触发条件模板
func (c *Client) CopyConditionsTemplate(request *CopyConditionsTemplateRequest) (response *CopyConditionsTemplateResponse, err error) {
    if request == nil {
        request = NewCopyConditionsTemplateRequest()
    }
    response = NewCopyConditionsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCopyInstanceGroupRequest() (request *CopyInstanceGroupRequest) {
    request = &CopyInstanceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CopyInstanceGroup")
    return
}

func NewCopyInstanceGroupResponse() (response *CopyInstanceGroupResponse) {
    response = &CopyInstanceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 复制实例分组
func (c *Client) CopyInstanceGroup(request *CopyInstanceGroupRequest) (response *CopyInstanceGroupResponse, err error) {
    if request == nil {
        request = NewCopyInstanceGroupRequest()
    }
    response = NewCopyInstanceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlertPolicyRequest() (request *CreateAlertPolicyRequest) {
    request = &CreateAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlertPolicy")
    return
}

func NewCreateAlertPolicyResponse() (response *CreateAlertPolicyResponse) {
    response = &CreateAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增告警策略返回对应ID
func (c *Client) CreateAlertPolicy(request *CreateAlertPolicyRequest) (response *CreateAlertPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAlertPolicyRequest()
    }
    response = NewCreateAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAttributesRequest() (request *CreateAttributesRequest) {
    request = &CreateAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAttributes")
    return
}

func NewCreateAttributesResponse() (response *CreateAttributesResponse) {
    response = &CreateAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量创建指标
func (c *Client) CreateAttributes(request *CreateAttributesRequest) (response *CreateAttributesResponse, err error) {
    if request == nil {
        request = NewCreateAttributesRequest()
    }
    response = NewCreateAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCMChartRequest() (request *CreateCCMChartRequest) {
    request = &CreateCCMChartRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateCCMChart")
    return
}

func NewCreateCCMChartResponse() (response *CreateCCMChartResponse) {
    response = &CreateCCMChartResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建分组视图自定义监控图表
func (c *Client) CreateCCMChart(request *CreateCCMChartRequest) (response *CreateCCMChartResponse, err error) {
    if request == nil {
        request = NewCreateCCMChartRequest()
    }
    response = NewCreateCCMChartResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConditionsTemplateRequest() (request *CreateConditionsTemplateRequest) {
    request = &CreateConditionsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateConditionsTemplate")
    return
}

func NewCreateConditionsTemplateResponse() (response *CreateConditionsTemplateResponse) {
    response = &CreateConditionsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建告警条件模板
func (c *Client) CreateConditionsTemplate(request *CreateConditionsTemplateRequest) (response *CreateConditionsTemplateResponse, err error) {
    if request == nil {
        request = NewCreateConditionsTemplateRequest()
    }
    response = NewCreateConditionsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDashboardRequest() (request *CreateDashboardRequest) {
    request = &CreateDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateDashboard")
    return
}

func NewCreateDashboardResponse() (response *CreateDashboardResponse) {
    response = &CreateDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增监控面板
func (c *Client) CreateDashboard(request *CreateDashboardRequest) (response *CreateDashboardResponse, err error) {
    if request == nil {
        request = NewCreateDashboardRequest()
    }
    response = NewCreateDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDashboardViewRequest() (request *CreateDashboardViewRequest) {
    request = &CreateDashboardViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateDashboardView")
    return
}

func NewCreateDashboardViewResponse() (response *CreateDashboardViewResponse) {
    response = &CreateDashboardViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建监控面板视图
func (c *Client) CreateDashboardView(request *CreateDashboardViewRequest) (response *CreateDashboardViewResponse, err error) {
    if request == nil {
        request = NewCreateDashboardViewRequest()
    }
    response = NewCreateDashboardViewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateGroup")
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建子分组
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceGroupRequest() (request *CreateInstanceGroupRequest) {
    request = &CreateInstanceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateInstanceGroup")
    return
}

func NewCreateInstanceGroupResponse() (response *CreateInstanceGroupResponse) {
    response = &CreateInstanceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建实例分组
func (c *Client) CreateInstanceGroup(request *CreateInstanceGroupRequest) (response *CreateInstanceGroupResponse, err error) {
    if request == nil {
        request = NewCreateInstanceGroupRequest()
    }
    response = NewCreateInstanceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateInstances")
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量添加实例列表
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMetricSetRequest() (request *CreateMetricSetRequest) {
    request = &CreateMetricSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateMetricSet")
    return
}

func NewCreateMetricSetResponse() (response *CreateMetricSetResponse) {
    response = &CreateMetricSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建指标集
func (c *Client) CreateMetricSet(request *CreateMetricSetRequest) (response *CreateMetricSetResponse, err error) {
    if request == nil {
        request = NewCreateMetricSetRequest()
    }
    response = NewCreateMetricSetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModuleRequest() (request *CreateModuleRequest) {
    request = &CreateModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateModule")
    return
}

func NewCreateModuleResponse() (response *CreateModuleResponse) {
    response = &CreateModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建叶子分组节点
func (c *Client) CreateModule(request *CreateModuleRequest) (response *CreateModuleResponse, err error) {
    if request == nil {
        request = NewCreateModuleRequest()
    }
    response = NewCreateModuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMsgPolicyRequest() (request *CreateMsgPolicyRequest) {
    request = &CreateMsgPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateMsgPolicy")
    return
}

func NewCreateMsgPolicyResponse() (response *CreateMsgPolicyResponse) {
    response = &CreateMsgPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建自定义消息策略
func (c *Client) CreateMsgPolicy(request *CreateMsgPolicyRequest) (response *CreateMsgPolicyResponse, err error) {
    if request == nil {
        request = NewCreateMsgPolicyRequest()
    }
    response = NewCreateMsgPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyGroupRequest() (request *CreatePolicyGroupRequest) {
    request = &CreatePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePolicyGroup")
    return
}

func NewCreatePolicyGroupResponse() (response *CreatePolicyGroupResponse) {
    response = &CreatePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加策略组
func (c *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (response *CreatePolicyGroupResponse, err error) {
    if request == nil {
        request = NewCreatePolicyGroupRequest()
    }
    response = NewCreatePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
    request = &CreateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateService")
    return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
    response = &CreateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开通服务
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    response = NewCreateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStrategyRequest() (request *CreateStrategyRequest) {
    request = &CreateStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateStrategy")
    return
}

func NewCreateStrategyResponse() (response *CreateStrategyResponse) {
    response = &CreateStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建告警策略
func (c *Client) CreateStrategy(request *CreateStrategyRequest) (response *CreateStrategyResponse, err error) {
    if request == nil {
        request = NewCreateStrategyRequest()
    }
    response = NewCreateStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlertPolicyRequest() (request *DeleteAlertPolicyRequest) {
    request = &DeleteAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlertPolicy")
    return
}

func NewDeleteAlertPolicyResponse() (response *DeleteAlertPolicyResponse) {
    response = &DeleteAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAlertPolicy删除告警策略
func (c *Client) DeleteAlertPolicy(request *DeleteAlertPolicyRequest) (response *DeleteAlertPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAlertPolicyRequest()
    }
    response = NewDeleteAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttributeRequest() (request *DeleteAttributeRequest) {
    request = &DeleteAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAttribute")
    return
}

func NewDeleteAttributeResponse() (response *DeleteAttributeResponse) {
    response = &DeleteAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除指标
func (c *Client) DeleteAttribute(request *DeleteAttributeRequest) (response *DeleteAttributeResponse, err error) {
    if request == nil {
        request = NewDeleteAttributeRequest()
    }
    response = NewDeleteAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttributesRequest() (request *DeleteAttributesRequest) {
    request = &DeleteAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAttributes")
    return
}

func NewDeleteAttributesResponse() (response *DeleteAttributesResponse) {
    response = &DeleteAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除指标
func (c *Client) DeleteAttributes(request *DeleteAttributesRequest) (response *DeleteAttributesResponse, err error) {
    if request == nil {
        request = NewDeleteAttributesRequest()
    }
    response = NewDeleteAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCMChartsRequest() (request *DeleteCCMChartsRequest) {
    request = &DeleteCCMChartsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteCCMCharts")
    return
}

func NewDeleteCCMChartsResponse() (response *DeleteCCMChartsResponse) {
    response = &DeleteCCMChartsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除分组视图自定义监控图表
func (c *Client) DeleteCCMCharts(request *DeleteCCMChartsRequest) (response *DeleteCCMChartsResponse, err error) {
    if request == nil {
        request = NewDeleteCCMChartsRequest()
    }
    response = NewDeleteCCMChartsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConditionsTemplateRequest() (request *DeleteConditionsTemplateRequest) {
    request = &DeleteConditionsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteConditionsTemplate")
    return
}

func NewDeleteConditionsTemplateResponse() (response *DeleteConditionsTemplateResponse) {
    response = &DeleteConditionsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除触发条件模板
func (c *Client) DeleteConditionsTemplate(request *DeleteConditionsTemplateRequest) (response *DeleteConditionsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteConditionsTemplateRequest()
    }
    response = NewDeleteConditionsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDashboardRequest() (request *DeleteDashboardRequest) {
    request = &DeleteDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteDashboard")
    return
}

func NewDeleteDashboardResponse() (response *DeleteDashboardResponse) {
    response = &DeleteDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除监控面板
func (c *Client) DeleteDashboard(request *DeleteDashboardRequest) (response *DeleteDashboardResponse, err error) {
    if request == nil {
        request = NewDeleteDashboardRequest()
    }
    response = NewDeleteDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDashboardViewRequest() (request *DeleteDashboardViewRequest) {
    request = &DeleteDashboardViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteDashboardView")
    return
}

func NewDeleteDashboardViewResponse() (response *DeleteDashboardViewResponse) {
    response = &DeleteDashboardViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除监控面板视图
func (c *Client) DeleteDashboardView(request *DeleteDashboardViewRequest) (response *DeleteDashboardViewResponse, err error) {
    if request == nil {
        request = NewDeleteDashboardViewRequest()
    }
    response = NewDeleteDashboardViewResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteGroup")
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除分组节点
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceGroupRequest() (request *DeleteInstanceGroupRequest) {
    request = &DeleteInstanceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteInstanceGroup")
    return
}

func NewDeleteInstanceGroupResponse() (response *DeleteInstanceGroupResponse) {
    response = &DeleteInstanceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除实例分组
func (c *Client) DeleteInstanceGroup(request *DeleteInstanceGroupRequest) (response *DeleteInstanceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceGroupRequest()
    }
    response = NewDeleteInstanceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstancesRequest() (request *DeleteInstancesRequest) {
    request = &DeleteInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteInstances")
    return
}

func NewDeleteInstancesResponse() (response *DeleteInstancesResponse) {
    response = &DeleteInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除分组列表
func (c *Client) DeleteInstances(request *DeleteInstancesRequest) (response *DeleteInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteInstancesRequest()
    }
    response = NewDeleteInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstancesInInstanceGroupRequest() (request *DeleteInstancesInInstanceGroupRequest) {
    request = &DeleteInstancesInInstanceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteInstancesInInstanceGroup")
    return
}

func NewDeleteInstancesInInstanceGroupResponse() (response *DeleteInstancesInInstanceGroupResponse) {
    response = &DeleteInstancesInInstanceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除实例分组中实例对象
func (c *Client) DeleteInstancesInInstanceGroup(request *DeleteInstancesInInstanceGroupRequest) (response *DeleteInstancesInInstanceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteInstancesInInstanceGroupRequest()
    }
    response = NewDeleteInstancesInInstanceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMetricSetRequest() (request *DeleteMetricSetRequest) {
    request = &DeleteMetricSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteMetricSet")
    return
}

func NewDeleteMetricSetResponse() (response *DeleteMetricSetResponse) {
    response = &DeleteMetricSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMetricSet删除指标集
func (c *Client) DeleteMetricSet(request *DeleteMetricSetRequest) (response *DeleteMetricSetResponse, err error) {
    if request == nil {
        request = NewDeleteMetricSetRequest()
    }
    response = NewDeleteMetricSetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModuleRequest() (request *DeleteModuleRequest) {
    request = &DeleteModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteModule")
    return
}

func NewDeleteModuleResponse() (response *DeleteModuleResponse) {
    response = &DeleteModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除叶子分组节点
func (c *Client) DeleteModule(request *DeleteModuleRequest) (response *DeleteModuleResponse, err error) {
    if request == nil {
        request = NewDeleteModuleRequest()
    }
    response = NewDeleteModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMsgPolicyRequest() (request *DeleteMsgPolicyRequest) {
    request = &DeleteMsgPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteMsgPolicy")
    return
}

func NewDeleteMsgPolicyResponse() (response *DeleteMsgPolicyResponse) {
    response = &DeleteMsgPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除自定义消息策略
func (c *Client) DeleteMsgPolicy(request *DeleteMsgPolicyRequest) (response *DeleteMsgPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteMsgPolicyRequest()
    }
    response = NewDeleteMsgPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyGroupRequest() (request *DeletePolicyGroupRequest) {
    request = &DeletePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePolicyGroup")
    return
}

func NewDeletePolicyGroupResponse() (response *DeletePolicyGroupResponse) {
    response = &DeletePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除告警策略组
func (c *Client) DeletePolicyGroup(request *DeletePolicyGroupRequest) (response *DeletePolicyGroupResponse, err error) {
    if request == nil {
        request = NewDeletePolicyGroupRequest()
    }
    response = NewDeletePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStrategysRequest() (request *DeleteStrategysRequest) {
    request = &DeleteStrategysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteStrategys")
    return
}

func NewDeleteStrategysResponse() (response *DeleteStrategysResponse) {
    response = &DeleteStrategysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除告警策略
func (c *Client) DeleteStrategys(request *DeleteStrategysRequest) (response *DeleteStrategysResponse, err error) {
    if request == nil {
        request = NewDeleteStrategysRequest()
    }
    response = NewDeleteStrategysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalObjectsRequest() (request *DescribeAbnormalObjectsRequest) {
    request = &DescribeAbnormalObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAbnormalObjects")
    return
}

func NewDescribeAbnormalObjectsResponse() (response *DescribeAbnormalObjectsResponse) {
    response = &DescribeAbnormalObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取近二十四小时发生异常的业务实例（告警维度）
func (c *Client) DescribeAbnormalObjects(request *DescribeAbnormalObjectsRequest) (response *DescribeAbnormalObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalObjectsRequest()
    }
    response = NewDescribeAbnormalObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccidentConfigRequest() (request *DescribeAccidentConfigRequest) {
    request = &DescribeAccidentConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAccidentConfig")
    return
}

func NewDescribeAccidentConfigResponse() (response *DescribeAccidentConfigResponse) {
    response = &DescribeAccidentConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取服务健康看板事件配置
func (c *Client) DescribeAccidentConfig(request *DescribeAccidentConfigRequest) (response *DescribeAccidentConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAccidentConfigRequest()
    }
    response = NewDescribeAccidentConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccidentEventListRequest() (request *DescribeAccidentEventListRequest) {
    request = &DescribeAccidentEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAccidentEventList")
    return
}

func NewDescribeAccidentEventListResponse() (response *DescribeAccidentEventListResponse) {
    response = &DescribeAccidentEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取平台事件列表
func (c *Client) DescribeAccidentEventList(request *DescribeAccidentEventListRequest) (response *DescribeAccidentEventListResponse, err error) {
    if request == nil {
        request = NewDescribeAccidentEventListRequest()
    }
    response = NewDescribeAccidentEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentStatusHistoryRequest() (request *DescribeAgentStatusHistoryRequest) {
    request = &DescribeAgentStatusHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAgentStatusHistory")
    return
}

func NewDescribeAgentStatusHistoryResponse() (response *DescribeAgentStatusHistoryResponse) {
    response = &DescribeAgentStatusHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取子机历史状态
func (c *Client) DescribeAgentStatusHistory(request *DescribeAgentStatusHistoryRequest) (response *DescribeAgentStatusHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeAgentStatusHistoryRequest()
    }
    response = NewDescribeAgentStatusHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmBindingInstanceListRequest() (request *DescribeAlarmBindingInstanceListRequest) {
    request = &DescribeAlarmBindingInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmBindingInstanceList")
    return
}

func NewDescribeAlarmBindingInstanceListResponse() (response *DescribeAlarmBindingInstanceListResponse) {
    response = &DescribeAlarmBindingInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警策略绑定实例列表
func (c *Client) DescribeAlarmBindingInstanceList(request *DescribeAlarmBindingInstanceListRequest) (response *DescribeAlarmBindingInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmBindingInstanceListRequest()
    }
    response = NewDescribeAlarmBindingInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmCallbackHistoryRequest() (request *DescribeAlarmCallbackHistoryRequest) {
    request = &DescribeAlarmCallbackHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmCallbackHistory")
    return
}

func NewDescribeAlarmCallbackHistoryResponse() (response *DescribeAlarmCallbackHistoryResponse) {
    response = &DescribeAlarmCallbackHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警回调接口历史记录列表
func (c *Client) DescribeAlarmCallbackHistory(request *DescribeAlarmCallbackHistoryRequest) (response *DescribeAlarmCallbackHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmCallbackHistoryRequest()
    }
    response = NewDescribeAlarmCallbackHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmCallbackVerifyCodeRequest() (request *DescribeAlarmCallbackVerifyCodeRequest) {
    request = &DescribeAlarmCallbackVerifyCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmCallbackVerifyCode")
    return
}

func NewDescribeAlarmCallbackVerifyCodeResponse() (response *DescribeAlarmCallbackVerifyCodeResponse) {
    response = &DescribeAlarmCallbackVerifyCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警回调验证码接口
func (c *Client) DescribeAlarmCallbackVerifyCode(request *DescribeAlarmCallbackVerifyCodeRequest) (response *DescribeAlarmCallbackVerifyCodeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmCallbackVerifyCodeRequest()
    }
    response = NewDescribeAlarmCallbackVerifyCodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmHistoryByAlarmIdRequest() (request *DescribeAlarmHistoryByAlarmIdRequest) {
    request = &DescribeAlarmHistoryByAlarmIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmHistoryByAlarmId")
    return
}

func NewDescribeAlarmHistoryByAlarmIdResponse() (response *DescribeAlarmHistoryByAlarmIdResponse) {
    response = &DescribeAlarmHistoryByAlarmIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据alarmId获取历史告警数据
func (c *Client) DescribeAlarmHistoryByAlarmId(request *DescribeAlarmHistoryByAlarmIdRequest) (response *DescribeAlarmHistoryByAlarmIdResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmHistoryByAlarmIdRequest()
    }
    response = NewDescribeAlarmHistoryByAlarmIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmInfoRequest() (request *DescribeAlarmInfoRequest) {
    request = &DescribeAlarmInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmInfo")
    return
}

func NewDescribeAlarmInfoResponse() (response *DescribeAlarmInfoResponse) {
    response = &DescribeAlarmInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警详情
func (c *Client) DescribeAlarmInfo(request *DescribeAlarmInfoRequest) (response *DescribeAlarmInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmInfoRequest()
    }
    response = NewDescribeAlarmInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmSmsQuotaRequest() (request *DescribeAlarmSmsQuotaRequest) {
    request = &DescribeAlarmSmsQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmSmsQuota")
    return
}

func NewDescribeAlarmSmsQuotaResponse() (response *DescribeAlarmSmsQuotaResponse) {
    response = &DescribeAlarmSmsQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警短信配额
func (c *Client) DescribeAlarmSmsQuota(request *DescribeAlarmSmsQuotaRequest) (response *DescribeAlarmSmsQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmSmsQuotaRequest()
    }
    response = NewDescribeAlarmSmsQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertPoliciesRequest() (request *DescribeAlertPoliciesRequest) {
    request = &DescribeAlertPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlertPolicies")
    return
}

func NewDescribeAlertPoliciesResponse() (response *DescribeAlertPoliciesResponse) {
    response = &DescribeAlertPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询告警策略列表
func (c *Client) DescribeAlertPolicies(request *DescribeAlertPoliciesRequest) (response *DescribeAlertPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAlertPoliciesRequest()
    }
    response = NewDescribeAlertPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertPolicyRequest() (request *DescribeAlertPolicyRequest) {
    request = &DescribeAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlertPolicy")
    return
}

func NewDescribeAlertPolicyResponse() (response *DescribeAlertPolicyResponse) {
    response = &DescribeAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlertPolicyCLM查询告警策略详情
func (c *Client) DescribeAlertPolicy(request *DescribeAlertPolicyRequest) (response *DescribeAlertPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeAlertPolicyRequest()
    }
    response = NewDescribeAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppFlowRequest() (request *DescribeAppFlowRequest) {
    request = &DescribeAppFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAppFlow")
    return
}

func NewDescribeAppFlowResponse() (response *DescribeAppFlowResponse) {
    response = &DescribeAppFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取实例月粒度数据
func (c *Client) DescribeAppFlow(request *DescribeAppFlowRequest) (response *DescribeAppFlowResponse, err error) {
    if request == nil {
        request = NewDescribeAppFlowRequest()
    }
    response = NewDescribeAppFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttributeAggregateDataRequest() (request *DescribeAttributeAggregateDataRequest) {
    request = &DescribeAttributeAggregateDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAttributeAggregateData")
    return
}

func NewDescribeAttributeAggregateDataResponse() (response *DescribeAttributeAggregateDataResponse) {
    response = &DescribeAttributeAggregateDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指标聚合数据
func (c *Client) DescribeAttributeAggregateData(request *DescribeAttributeAggregateDataRequest) (response *DescribeAttributeAggregateDataResponse, err error) {
    if request == nil {
        request = NewDescribeAttributeAggregateDataRequest()
    }
    response = NewDescribeAttributeAggregateDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttributeAllServerRequest() (request *DescribeAttributeAllServerRequest) {
    request = &DescribeAttributeAllServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAttributeAllServer")
    return
}

func NewDescribeAttributeAllServerResponse() (response *DescribeAttributeAllServerResponse) {
    response = &DescribeAttributeAllServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指标关联服务器信息
func (c *Client) DescribeAttributeAllServer(request *DescribeAttributeAllServerRequest) (response *DescribeAttributeAllServerResponse, err error) {
    if request == nil {
        request = NewDescribeAttributeAllServerRequest()
    }
    response = NewDescribeAttributeAllServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttributeUnitsRequest() (request *DescribeAttributeUnitsRequest) {
    request = &DescribeAttributeUnitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAttributeUnits")
    return
}

func NewDescribeAttributeUnitsResponse() (response *DescribeAttributeUnitsResponse) {
    response = &DescribeAttributeUnitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指标单位
func (c *Client) DescribeAttributeUnits(request *DescribeAttributeUnitsRequest) (response *DescribeAttributeUnitsResponse, err error) {
    if request == nil {
        request = NewDescribeAttributeUnitsRequest()
    }
    response = NewDescribeAttributeUnitsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttributesRequest() (request *DescribeAttributesRequest) {
    request = &DescribeAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAttributes")
    return
}

func NewDescribeAttributesResponse() (response *DescribeAttributesResponse) {
    response = &DescribeAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询属性信息
func (c *Client) DescribeAttributes(request *DescribeAttributesRequest) (response *DescribeAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeAttributesRequest()
    }
    response = NewDescribeAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaseMetricsRequest() (request *DescribeBaseMetricsRequest) {
    request = &DescribeBaseMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBaseMetrics")
    return
}

func NewDescribeBaseMetricsResponse() (response *DescribeBaseMetricsResponse) {
    response = &DescribeBaseMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取基础指标详情
func (c *Client) DescribeBaseMetrics(request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeBaseMetricsRequest()
    }
    response = NewDescribeBaseMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaseMetricsForConsoleFontEndRequest() (request *DescribeBaseMetricsForConsoleFontEndRequest) {
    request = &DescribeBaseMetricsForConsoleFontEndRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBaseMetricsForConsoleFontEnd")
    return
}

func NewDescribeBaseMetricsForConsoleFontEndResponse() (response *DescribeBaseMetricsForConsoleFontEndResponse) {
    response = &DescribeBaseMetricsForConsoleFontEndResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 控制台前端调用获取基础指标
func (c *Client) DescribeBaseMetricsForConsoleFontEnd(request *DescribeBaseMetricsForConsoleFontEndRequest) (response *DescribeBaseMetricsForConsoleFontEndResponse, err error) {
    if request == nil {
        request = NewDescribeBaseMetricsForConsoleFontEndRequest()
    }
    response = NewDescribeBaseMetricsForConsoleFontEndResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicAlarmListRequest() (request *DescribeBasicAlarmListRequest) {
    request = &DescribeBasicAlarmListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBasicAlarmList")
    return
}

func NewDescribeBasicAlarmListResponse() (response *DescribeBasicAlarmListResponse) {
    response = &DescribeBasicAlarmListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取基础告警列表
func (c *Client) DescribeBasicAlarmList(request *DescribeBasicAlarmListRequest) (response *DescribeBasicAlarmListResponse, err error) {
    if request == nil {
        request = NewDescribeBasicAlarmListRequest()
    }
    response = NewDescribeBasicAlarmListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindingPolicyObjectListRequest() (request *DescribeBindingPolicyObjectListRequest) {
    request = &DescribeBindingPolicyObjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBindingPolicyObjectList")
    return
}

func NewDescribeBindingPolicyObjectListResponse() (response *DescribeBindingPolicyObjectListResponse) {
    response = &DescribeBindingPolicyObjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取已绑定对象列表
func (c *Client) DescribeBindingPolicyObjectList(request *DescribeBindingPolicyObjectListRequest) (response *DescribeBindingPolicyObjectListResponse, err error) {
    if request == nil {
        request = NewDescribeBindingPolicyObjectListRequest()
    }
    response = NewDescribeBindingPolicyObjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCMChartsRequest() (request *DescribeCCMChartsRequest) {
    request = &DescribeCCMChartsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeCCMCharts")
    return
}

func NewDescribeCCMChartsResponse() (response *DescribeCCMChartsResponse) {
    response = &DescribeCCMChartsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询分组视图自定义监控图表
func (c *Client) DescribeCCMCharts(request *DescribeCCMChartsRequest) (response *DescribeCCMChartsResponse, err error) {
    if request == nil {
        request = NewDescribeCCMChartsRequest()
    }
    response = NewDescribeCCMChartsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCMGroupViewRequest() (request *DescribeCCMGroupViewRequest) {
    request = &DescribeCCMGroupViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeCCMGroupView")
    return
}

func NewDescribeCCMGroupViewResponse() (response *DescribeCCMGroupViewResponse) {
    response = &DescribeCCMGroupViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询分组视图
func (c *Client) DescribeCCMGroupView(request *DescribeCCMGroupViewRequest) (response *DescribeCCMGroupViewResponse, err error) {
    if request == nil {
        request = NewDescribeCCMGroupViewRequest()
    }
    response = NewDescribeCCMGroupViewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCMGroupViewAttributeRequest() (request *DescribeCCMGroupViewAttributeRequest) {
    request = &DescribeCCMGroupViewAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeCCMGroupViewAttribute")
    return
}

func NewDescribeCCMGroupViewAttributeResponse() (response *DescribeCCMGroupViewAttributeResponse) {
    response = &DescribeCCMGroupViewAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询分组视图关联指标
func (c *Client) DescribeCCMGroupViewAttribute(request *DescribeCCMGroupViewAttributeRequest) (response *DescribeCCMGroupViewAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeCCMGroupViewAttributeRequest()
    }
    response = NewDescribeCCMGroupViewAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCMGroupViewStrategyRequest() (request *DescribeCCMGroupViewStrategyRequest) {
    request = &DescribeCCMGroupViewStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeCCMGroupViewStrategy")
    return
}

func NewDescribeCCMGroupViewStrategyResponse() (response *DescribeCCMGroupViewStrategyResponse) {
    response = &DescribeCCMGroupViewStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询分组视图关联告警策略
func (c *Client) DescribeCCMGroupViewStrategy(request *DescribeCCMGroupViewStrategyRequest) (response *DescribeCCMGroupViewStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeCCMGroupViewStrategyRequest()
    }
    response = NewDescribeCCMGroupViewStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCMInstanceDatasRequest() (request *DescribeCCMInstanceDatasRequest) {
    request = &DescribeCCMInstanceDatasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeCCMInstanceDatas")
    return
}

func NewDescribeCCMInstanceDatasResponse() (response *DescribeCCMInstanceDatasResponse) {
    response = &DescribeCCMInstanceDatasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量查询实例指标数据
func (c *Client) DescribeCCMInstanceDatas(request *DescribeCCMInstanceDatasRequest) (response *DescribeCCMInstanceDatasResponse, err error) {
    if request == nil {
        request = NewDescribeCCMInstanceDatasRequest()
    }
    response = NewDescribeCCMInstanceDatasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConditionsTemplateListRequest() (request *DescribeConditionsTemplateListRequest) {
    request = &DescribeConditionsTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeConditionsTemplateList")
    return
}

func NewDescribeConditionsTemplateListResponse() (response *DescribeConditionsTemplateListResponse) {
    response = &DescribeConditionsTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取条件模板列表
func (c *Client) DescribeConditionsTemplateList(request *DescribeConditionsTemplateListRequest) (response *DescribeConditionsTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeConditionsTemplateListRequest()
    }
    response = NewDescribeConditionsTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContactListRequest() (request *DescribeContactListRequest) {
    request = &DescribeContactListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeContactList")
    return
}

func NewDescribeContactListResponse() (response *DescribeContactListResponse) {
    response = &DescribeContactListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取联系人列表
func (c *Client) DescribeContactList(request *DescribeContactListRequest) (response *DescribeContactListResponse, err error) {
    if request == nil {
        request = NewDescribeContactListRequest()
    }
    response = NewDescribeContactListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurrentTimestampRequest() (request *DescribeCurrentTimestampRequest) {
    request = &DescribeCurrentTimestampRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeCurrentTimestamp")
    return
}

func NewDescribeCurrentTimestampResponse() (response *DescribeCurrentTimestampResponse) {
    response = &DescribeCurrentTimestampResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取服务器当前时间戳
func (c *Client) DescribeCurrentTimestamp(request *DescribeCurrentTimestampRequest) (response *DescribeCurrentTimestampResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentTimestampRequest()
    }
    response = NewDescribeCurrentTimestampResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomAlarmListRequest() (request *DescribeCustomAlarmListRequest) {
    request = &DescribeCustomAlarmListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeCustomAlarmList")
    return
}

func NewDescribeCustomAlarmListResponse() (response *DescribeCustomAlarmListResponse) {
    response = &DescribeCustomAlarmListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取自定义消息列表
func (c *Client) DescribeCustomAlarmList(request *DescribeCustomAlarmListRequest) (response *DescribeCustomAlarmListResponse, err error) {
    if request == nil {
        request = NewDescribeCustomAlarmListRequest()
    }
    response = NewDescribeCustomAlarmListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCvmAgentStatusRequest() (request *DescribeCvmAgentStatusRequest) {
    request = &DescribeCvmAgentStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeCvmAgentStatus")
    return
}

func NewDescribeCvmAgentStatusResponse() (response *DescribeCvmAgentStatusResponse) {
    response = &DescribeCvmAgentStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取cvm主机agent的状态（是否安装）
func (c *Client) DescribeCvmAgentStatus(request *DescribeCvmAgentStatusRequest) (response *DescribeCvmAgentStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCvmAgentStatusRequest()
    }
    response = NewDescribeCvmAgentStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardViewsRequest() (request *DescribeDashboardViewsRequest) {
    request = &DescribeDashboardViewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeDashboardViews")
    return
}

func NewDescribeDashboardViewsResponse() (response *DescribeDashboardViewsResponse) {
    response = &DescribeDashboardViewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取监控面板视图
func (c *Client) DescribeDashboardViews(request *DescribeDashboardViewsRequest) (response *DescribeDashboardViewsResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardViewsRequest()
    }
    response = NewDescribeDashboardViewsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardsRequest() (request *DescribeDashboardsRequest) {
    request = &DescribeDashboardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeDashboards")
    return
}

func NewDescribeDashboardsResponse() (response *DescribeDashboardsResponse) {
    response = &DescribeDashboardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Dashboard列表
func (c *Client) DescribeDashboards(request *DescribeDashboardsRequest) (response *DescribeDashboardsResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardsRequest()
    }
    response = NewDescribeDashboardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDimensionAnalysisDataRequest() (request *DescribeDimensionAnalysisDataRequest) {
    request = &DescribeDimensionAnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeDimensionAnalysisData")
    return
}

func NewDescribeDimensionAnalysisDataResponse() (response *DescribeDimensionAnalysisDataResponse) {
    response = &DescribeDimensionAnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取维度分析数据
func (c *Client) DescribeDimensionAnalysisData(request *DescribeDimensionAnalysisDataRequest) (response *DescribeDimensionAnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeDimensionAnalysisDataRequest()
    }
    response = NewDescribeDimensionAnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventPolicyConfigRequest() (request *DescribeEventPolicyConfigRequest) {
    request = &DescribeEventPolicyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeEventPolicyConfig")
    return
}

func NewDescribeEventPolicyConfigResponse() (response *DescribeEventPolicyConfigResponse) {
    response = &DescribeEventPolicyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件策略配置列表
func (c *Client) DescribeEventPolicyConfig(request *DescribeEventPolicyConfigRequest) (response *DescribeEventPolicyConfigResponse, err error) {
    if request == nil {
        request = NewDescribeEventPolicyConfigRequest()
    }
    response = NewDescribeEventPolicyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGraphDataRequest() (request *DescribeGraphDataRequest) {
    request = &DescribeGraphDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGraphData")
    return
}

func NewDescribeGraphDataResponse() (response *DescribeGraphDataResponse) {
    response = &DescribeGraphDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取硬盘分区使用图
func (c *Client) DescribeGraphData(request *DescribeGraphDataRequest) (response *DescribeGraphDataResponse, err error) {
    if request == nil {
        request = NewDescribeGraphDataRequest()
    }
    response = NewDescribeGraphDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGroup")
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取分组树
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdcServerRequest() (request *DescribeIdcServerRequest) {
    request = &DescribeIdcServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeIdcServer")
    return
}

func NewDescribeIdcServerResponse() (response *DescribeIdcServerResponse) {
    response = &DescribeIdcServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询地域下的所有服务器列表
func (c *Client) DescribeIdcServer(request *DescribeIdcServerRequest) (response *DescribeIdcServerResponse, err error) {
    if request == nil {
        request = NewDescribeIdcServerRequest()
    }
    response = NewDescribeIdcServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdcServerCountRequest() (request *DescribeIdcServerCountRequest) {
    request = &DescribeIdcServerCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeIdcServerCount")
    return
}

func NewDescribeIdcServerCountResponse() (response *DescribeIdcServerCountResponse) {
    response = &DescribeIdcServerCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询地域实例统计信息
func (c *Client) DescribeIdcServerCount(request *DescribeIdcServerCountRequest) (response *DescribeIdcServerCountResponse, err error) {
    if request == nil {
        request = NewDescribeIdcServerCountRequest()
    }
    response = NewDescribeIdcServerCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdcTreeRequest() (request *DescribeIdcTreeRequest) {
    request = &DescribeIdcTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeIdcTree")
    return
}

func NewDescribeIdcTreeResponse() (response *DescribeIdcTreeResponse) {
    response = &DescribeIdcTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询IDC树
func (c *Client) DescribeIdcTree(request *DescribeIdcTreeRequest) (response *DescribeIdcTreeResponse, err error) {
    if request == nil {
        request = NewDescribeIdcTreeRequest()
    }
    response = NewDescribeIdcTreeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeInstance")
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例列表
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    response = NewDescribeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceGroupRequest() (request *DescribeInstanceGroupRequest) {
    request = &DescribeInstanceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeInstanceGroup")
    return
}

func NewDescribeInstanceGroupResponse() (response *DescribeInstanceGroupResponse) {
    response = &DescribeInstanceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取实例分组详情
func (c *Client) DescribeInstanceGroup(request *DescribeInstanceGroupRequest) (response *DescribeInstanceGroupResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceGroupRequest()
    }
    response = NewDescribeInstanceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceGroupListRequest() (request *DescribeInstanceGroupListRequest) {
    request = &DescribeInstanceGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeInstanceGroupList")
    return
}

func NewDescribeInstanceGroupListResponse() (response *DescribeInstanceGroupListResponse) {
    response = &DescribeInstanceGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取实例分组列表
func (c *Client) DescribeInstanceGroupList(request *DescribeInstanceGroupListRequest) (response *DescribeInstanceGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceGroupListRequest()
    }
    response = NewDescribeInstanceGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogSetsRequest() (request *DescribeLogSetsRequest) {
    request = &DescribeLogSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeLogSets")
    return
}

func NewDescribeLogSetsResponse() (response *DescribeLogSetsResponse) {
    response = &DescribeLogSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取CLS中的日志集列表
func (c *Client) DescribeLogSets(request *DescribeLogSetsRequest) (response *DescribeLogSetsResponse, err error) {
    if request == nil {
        request = NewDescribeLogSetsRequest()
    }
    response = NewDescribeLogSetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogTopicIndexRequest() (request *DescribeLogTopicIndexRequest) {
    request = &DescribeLogTopicIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeLogTopicIndex")
    return
}

func NewDescribeLogTopicIndexResponse() (response *DescribeLogTopicIndexResponse) {
    response = &DescribeLogTopicIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取CLS日志主题索引信息
func (c *Client) DescribeLogTopicIndex(request *DescribeLogTopicIndexRequest) (response *DescribeLogTopicIndexResponse, err error) {
    if request == nil {
        request = NewDescribeLogTopicIndexRequest()
    }
    response = NewDescribeLogTopicIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogTopicsRequest() (request *DescribeLogTopicsRequest) {
    request = &DescribeLogTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeLogTopics")
    return
}

func NewDescribeLogTopicsResponse() (response *DescribeLogTopicsResponse) {
    response = &DescribeLogTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取CLS日志主题
func (c *Client) DescribeLogTopics(request *DescribeLogTopicsRequest) (response *DescribeLogTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeLogTopicsRequest()
    }
    response = NewDescribeLogTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricAnalysisDataRequest() (request *DescribeMetricAnalysisDataRequest) {
    request = &DescribeMetricAnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMetricAnalysisData")
    return
}

func NewDescribeMetricAnalysisDataResponse() (response *DescribeMetricAnalysisDataResponse) {
    response = &DescribeMetricAnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指标分析数据
func (c *Client) DescribeMetricAnalysisData(request *DescribeMetricAnalysisDataRequest) (response *DescribeMetricAnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeMetricAnalysisDataRequest()
    }
    response = NewDescribeMetricAnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricSetRequest() (request *DescribeMetricSetRequest) {
    request = &DescribeMetricSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMetricSet")
    return
}

func NewDescribeMetricSetResponse() (response *DescribeMetricSetResponse) {
    response = &DescribeMetricSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指标集详情
func (c *Client) DescribeMetricSet(request *DescribeMetricSetRequest) (response *DescribeMetricSetResponse, err error) {
    if request == nil {
        request = NewDescribeMetricSetRequest()
    }
    response = NewDescribeMetricSetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricSetsRequest() (request *DescribeMetricSetsRequest) {
    request = &DescribeMetricSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMetricSets")
    return
}

func NewDescribeMetricSetsResponse() (response *DescribeMetricSetsResponse) {
    response = &DescribeMetricSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取指标集列表DescribeMetricSets
func (c *Client) DescribeMetricSets(request *DescribeMetricSetsRequest) (response *DescribeMetricSetsResponse, err error) {
    if request == nil {
        request = NewDescribeMetricSetsRequest()
    }
    response = NewDescribeMetricSetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModuleRequest() (request *DescribeModuleRequest) {
    request = &DescribeModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeModule")
    return
}

func NewDescribeModuleResponse() (response *DescribeModuleResponse) {
    response = &DescribeModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取叶子分组节点
func (c *Client) DescribeModule(request *DescribeModuleRequest) (response *DescribeModuleResponse, err error) {
    if request == nil {
        request = NewDescribeModuleRequest()
    }
    response = NewDescribeModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorDataByAlarmIDRequest() (request *DescribeMonitorDataByAlarmIDRequest) {
    request = &DescribeMonitorDataByAlarmIDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMonitorDataByAlarmID")
    return
}

func NewDescribeMonitorDataByAlarmIDResponse() (response *DescribeMonitorDataByAlarmIDResponse) {
    response = &DescribeMonitorDataByAlarmIDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据告警ID获取监控数据
func (c *Client) DescribeMonitorDataByAlarmID(request *DescribeMonitorDataByAlarmIDRequest) (response *DescribeMonitorDataByAlarmIDResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorDataByAlarmIDRequest()
    }
    response = NewDescribeMonitorDataByAlarmIDResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorProductByIdsRequest() (request *DescribeMonitorProductByIdsRequest) {
    request = &DescribeMonitorProductByIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMonitorProductByIds")
    return
}

func NewDescribeMonitorProductByIdsResponse() (response *DescribeMonitorProductByIdsResponse) {
    response = &DescribeMonitorProductByIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按照Id查询监控产品列表
func (c *Client) DescribeMonitorProductByIds(request *DescribeMonitorProductByIdsRequest) (response *DescribeMonitorProductByIdsResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorProductByIdsRequest()
    }
    response = NewDescribeMonitorProductByIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorProductsRequest() (request *DescribeMonitorProductsRequest) {
    request = &DescribeMonitorProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMonitorProducts")
    return
}

func NewDescribeMonitorProductsResponse() (response *DescribeMonitorProductsResponse) {
    response = &DescribeMonitorProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询监控产品列表
func (c *Client) DescribeMonitorProducts(request *DescribeMonitorProductsRequest) (response *DescribeMonitorProductsResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorProductsRequest()
    }
    response = NewDescribeMonitorProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsgPolicyInfoRequest() (request *DescribeMsgPolicyInfoRequest) {
    request = &DescribeMsgPolicyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMsgPolicyInfo")
    return
}

func NewDescribeMsgPolicyInfoResponse() (response *DescribeMsgPolicyInfoResponse) {
    response = &DescribeMsgPolicyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取自定义消息策略信息
func (c *Client) DescribeMsgPolicyInfo(request *DescribeMsgPolicyInfoRequest) (response *DescribeMsgPolicyInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMsgPolicyInfoRequest()
    }
    response = NewDescribeMsgPolicyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsgPolicyListRequest() (request *DescribeMsgPolicyListRequest) {
    request = &DescribeMsgPolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMsgPolicyList")
    return
}

func NewDescribeMsgPolicyListResponse() (response *DescribeMsgPolicyListResponse) {
    response = &DescribeMsgPolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取自定义消息列表
func (c *Client) DescribeMsgPolicyList(request *DescribeMsgPolicyListRequest) (response *DescribeMsgPolicyListResponse, err error) {
    if request == nil {
        request = NewDescribeMsgPolicyListRequest()
    }
    response = NewDescribeMsgPolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyConditionListRequest() (request *DescribePolicyConditionListRequest) {
    request = &DescribePolicyConditionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyConditionList")
    return
}

func NewDescribePolicyConditionListResponse() (response *DescribePolicyConditionListResponse) {
    response = &DescribePolicyConditionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取基础告警策略条件
func (c *Client) DescribePolicyConditionList(request *DescribePolicyConditionListRequest) (response *DescribePolicyConditionListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyConditionListRequest()
    }
    response = NewDescribePolicyConditionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupCountRequest() (request *DescribePolicyGroupCountRequest) {
    request = &DescribePolicyGroupCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupCount")
    return
}

func NewDescribePolicyGroupCountResponse() (response *DescribePolicyGroupCountResponse) {
    response = &DescribePolicyGroupCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询对象列表绑定的告警策略组个数
func (c *Client) DescribePolicyGroupCount(request *DescribePolicyGroupCountRequest) (response *DescribePolicyGroupCountResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupCountRequest()
    }
    response = NewDescribePolicyGroupCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupInfoRequest() (request *DescribePolicyGroupInfoRequest) {
    request = &DescribePolicyGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupInfo")
    return
}

func NewDescribePolicyGroupInfoResponse() (response *DescribePolicyGroupInfoResponse) {
    response = &DescribePolicyGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取基础策略组详情
func (c *Client) DescribePolicyGroupInfo(request *DescribePolicyGroupInfoRequest) (response *DescribePolicyGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupInfoRequest()
    }
    response = NewDescribePolicyGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupListRequest() (request *DescribePolicyGroupListRequest) {
    request = &DescribePolicyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupList")
    return
}

func NewDescribePolicyGroupListResponse() (response *DescribePolicyGroupListResponse) {
    response = &DescribePolicyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取基础策略告警组列表
func (c *Client) DescribePolicyGroupList(request *DescribePolicyGroupListRequest) (response *DescribePolicyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupListRequest()
    }
    response = NewDescribePolicyGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyInfoByInstanceRequest() (request *DescribePolicyInfoByInstanceRequest) {
    request = &DescribePolicyInfoByInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyInfoByInstance")
    return
}

func NewDescribePolicyInfoByInstanceResponse() (response *DescribePolicyInfoByInstanceResponse) {
    response = &DescribePolicyInfoByInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过实例获取策略信息
func (c *Client) DescribePolicyInfoByInstance(request *DescribePolicyInfoByInstanceRequest) (response *DescribePolicyInfoByInstanceResponse, err error) {
    if request == nil {
        request = NewDescribePolicyInfoByInstanceRequest()
    }
    response = NewDescribePolicyInfoByInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyObjectCountRequest() (request *DescribePolicyObjectCountRequest) {
    request = &DescribePolicyObjectCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyObjectCount")
    return
}

func NewDescribePolicyObjectCountResponse() (response *DescribePolicyObjectCountResponse) {
    response = &DescribePolicyObjectCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询策略组在每个地域下面绑定的对象数统计
func (c *Client) DescribePolicyObjectCount(request *DescribePolicyObjectCountRequest) (response *DescribePolicyObjectCountResponse, err error) {
    if request == nil {
        request = NewDescribePolicyObjectCountRequest()
    }
    response = NewDescribePolicyObjectCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyQuotaRequest() (request *DescribePolicyQuotaRequest) {
    request = &DescribePolicyQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyQuota")
    return
}

func NewDescribePolicyQuotaResponse() (response *DescribePolicyQuotaResponse) {
    response = &DescribePolicyQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取策略配额
func (c *Client) DescribePolicyQuota(request *DescribePolicyQuotaRequest) (response *DescribePolicyQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePolicyQuotaRequest()
    }
    response = NewDescribePolicyQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicySituationRequest() (request *DescribePolicySituationRequest) {
    request = &DescribePolicySituationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicySituation")
    return
}

func NewDescribePolicySituationResponse() (response *DescribePolicySituationResponse) {
    response = &DescribePolicySituationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取告警策略统计情况
func (c *Client) DescribePolicySituation(request *DescribePolicySituationRequest) (response *DescribePolicySituationResponse, err error) {
    if request == nil {
        request = NewDescribePolicySituationRequest()
    }
    response = NewDescribePolicySituationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyUseListRequest() (request *DescribePolicyUseListRequest) {
    request = &DescribePolicyUseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyUseList")
    return
}

func NewDescribePolicyUseListResponse() (response *DescribePolicyUseListResponse) {
    response = &DescribePolicyUseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取已关联对象列表
func (c *Client) DescribePolicyUseList(request *DescribePolicyUseListRequest) (response *DescribePolicyUseListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyUseListRequest()
    }
    response = NewDescribePolicyUseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductEventListRequest() (request *DescribeProductEventListRequest) {
    request = &DescribeProductEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeProductEventList")
    return
}

func NewDescribeProductEventListResponse() (response *DescribeProductEventListResponse) {
    response = &DescribeProductEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分页获取产品事件的列表
func (c *Client) DescribeProductEventList(request *DescribeProductEventListRequest) (response *DescribeProductEventListResponse, err error) {
    if request == nil {
        request = NewDescribeProductEventListRequest()
    }
    response = NewDescribeProductEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductHealthStatusListRequest() (request *DescribeProductHealthStatusListRequest) {
    request = &DescribeProductHealthStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeProductHealthStatusList")
    return
}

func NewDescribeProductHealthStatusListResponse() (response *DescribeProductHealthStatusListResponse) {
    response = &DescribeProductHealthStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取每个业务近24小时健康情况统计
func (c *Client) DescribeProductHealthStatusList(request *DescribeProductHealthStatusListRequest) (response *DescribeProductHealthStatusListResponse, err error) {
    if request == nil {
        request = NewDescribeProductHealthStatusListRequest()
    }
    response = NewDescribeProductHealthStatusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsListRequest() (request *DescribeProjectsListRequest) {
    request = &DescribeProjectsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeProjectsList")
    return
}

func NewDescribeProjectsListResponse() (response *DescribeProjectsListResponse) {
    response = &DescribeProjectsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获得项目列表
func (c *Client) DescribeProjectsList(request *DescribeProjectsListRequest) (response *DescribeProjectsListResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsListRequest()
    }
    response = NewDescribeProjectsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecommendedTemplateRequest() (request *DescribeRecommendedTemplateRequest) {
    request = &DescribeRecommendedTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeRecommendedTemplate")
    return
}

func NewDescribeRecommendedTemplateResponse() (response *DescribeRecommendedTemplateResponse) {
    response = &DescribeRecommendedTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取默认推荐告警策略模板列表
func (c *Client) DescribeRecommendedTemplate(request *DescribeRecommendedTemplateRequest) (response *DescribeRecommendedTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeRecommendedTemplateRequest()
    }
    response = NewDescribeRecommendedTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerDataRequest() (request *DescribeServerDataRequest) {
    request = &DescribeServerDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeServerData")
    return
}

func NewDescribeServerDataResponse() (response *DescribeServerDataResponse) {
    response = &DescribeServerDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询单机数据
func (c *Client) DescribeServerData(request *DescribeServerDataRequest) (response *DescribeServerDataResponse, err error) {
    if request == nil {
        request = NewDescribeServerDataRequest()
    }
    response = NewDescribeServerDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerDatasRequest() (request *DescribeServerDatasRequest) {
    request = &DescribeServerDatasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeServerDatas")
    return
}

func NewDescribeServerDatasResponse() (response *DescribeServerDatasResponse) {
    response = &DescribeServerDatasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量查询单机数据
func (c *Client) DescribeServerDatas(request *DescribeServerDatasRequest) (response *DescribeServerDatasResponse, err error) {
    if request == nil {
        request = NewDescribeServerDatasRequest()
    }
    response = NewDescribeServerDatasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServersRequest() (request *DescribeServersRequest) {
    request = &DescribeServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeServers")
    return
}

func NewDescribeServersResponse() (response *DescribeServersResponse) {
    response = &DescribeServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务器
func (c *Client) DescribeServers(request *DescribeServersRequest) (response *DescribeServersResponse, err error) {
    if request == nil {
        request = NewDescribeServersRequest()
    }
    response = NewDescribeServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
    request = &DescribeServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeService")
    return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
    response = &DescribeServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务信息
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    if request == nil {
        request = NewDescribeServiceRequest()
    }
    response = NewDescribeServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSortObjectListRequest() (request *DescribeSortObjectListRequest) {
    request = &DescribeSortObjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeSortObjectList")
    return
}

func NewDescribeSortObjectListResponse() (response *DescribeSortObjectListResponse) {
    response = &DescribeSortObjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取按topN排序的某个viewname下面的对象列表
func (c *Client) DescribeSortObjectList(request *DescribeSortObjectListRequest) (response *DescribeSortObjectListResponse, err error) {
    if request == nil {
        request = NewDescribeSortObjectListRequest()
    }
    response = NewDescribeSortObjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageDurationRequest() (request *DescribeStorageDurationRequest) {
    request = &DescribeStorageDurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeStorageDuration")
    return
}

func NewDescribeStorageDurationResponse() (response *DescribeStorageDurationResponse) {
    response = &DescribeStorageDurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取存储周期及持续时间列表
func (c *Client) DescribeStorageDuration(request *DescribeStorageDurationRequest) (response *DescribeStorageDurationResponse, err error) {
    if request == nil {
        request = NewDescribeStorageDurationRequest()
    }
    response = NewDescribeStorageDurationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStrategysRequest() (request *DescribeStrategysRequest) {
    request = &DescribeStrategysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeStrategys")
    return
}

func NewDescribeStrategysResponse() (response *DescribeStrategysResponse) {
    response = &DescribeStrategysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量查询告警策略
func (c *Client) DescribeStrategys(request *DescribeStrategysRequest) (response *DescribeStrategysResponse, err error) {
    if request == nil {
        request = NewDescribeStrategysRequest()
    }
    response = NewDescribeStrategysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeInfoRequest() (request *DescribeSubscribeInfoRequest) {
    request = &DescribeSubscribeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeSubscribeInfo")
    return
}

func NewDescribeSubscribeInfoResponse() (response *DescribeSubscribeInfoResponse) {
    response = &DescribeSubscribeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询订阅告警列表
func (c *Client) DescribeSubscribeInfo(request *DescribeSubscribeInfoRequest) (response *DescribeSubscribeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeInfoRequest()
    }
    response = NewDescribeSubscribeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTransLogRequest() (request *DescribeTransLogRequest) {
    request = &DescribeTransLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeTransLog")
    return
}

func NewDescribeTransLogResponse() (response *DescribeTransLogResponse) {
    response = &DescribeTransLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取实例组变更日志
func (c *Client) DescribeTransLog(request *DescribeTransLogRequest) (response *DescribeTransLogResponse, err error) {
    if request == nil {
        request = NewDescribeTransLogRequest()
    }
    response = NewDescribeTransLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
    request = &DescribeUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeUserInfo")
    return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
    response = &DescribeUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户账号信息
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserInfoRequest()
    }
    response = NewDescribeUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeViewDataRequest() (request *DescribeViewDataRequest) {
    request = &DescribeViewDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeViewData")
    return
}

func NewDescribeViewDataResponse() (response *DescribeViewDataResponse) {
    response = &DescribeViewDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询视图数据
func (c *Client) DescribeViewData(request *DescribeViewDataRequest) (response *DescribeViewDataResponse, err error) {
    if request == nil {
        request = NewDescribeViewDataRequest()
    }
    response = NewDescribeViewDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetArgusMonitorDataRequest() (request *GetArgusMonitorDataRequest) {
    request = &GetArgusMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "GetArgusMonitorData")
    return
}

func NewGetArgusMonitorDataResponse() (response *GetArgusMonitorDataResponse) {
    response = &GetArgusMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Argus监控数据
func (c *Client) GetArgusMonitorData(request *GetArgusMonitorDataRequest) (response *GetArgusMonitorDataResponse, err error) {
    if request == nil {
        request = NewGetArgusMonitorDataRequest()
    }
    response = NewGetArgusMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetMonitorDataRequest() (request *GetMonitorDataRequest) {
    request = &GetMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "GetMonitorData")
    return
}

func NewGetMonitorDataResponse() (response *GetMonitorDataResponse) {
    response = &GetMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取云产品的监控数据。传入产品的命名空间、对象维度描述和监控指标即可获得相应的监控数据。
// 接口调用频率限制为：20次/秒，1200次/分钟。
// 若您需要调用的指标、对象较多，可能存在因限频出现拉取失败的情况，建议尽量将请求按时间维度均摊。
func (c *Client) GetMonitorData(request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
    if request == nil {
        request = NewGetMonitorDataRequest()
    }
    response = NewGetMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetTkeDataRequest() (request *GetTkeDataRequest) {
    request = &GetTkeDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "GetTkeData")
    return
}

func NewGetTkeDataResponse() (response *GetTkeDataResponse) {
    response = &GetTkeDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取TKE监控数据
func (c *Client) GetTkeData(request *GetTkeDataRequest) (response *GetTkeDataResponse, err error) {
    if request == nil {
        request = NewGetTkeDataRequest()
    }
    response = NewGetTkeDataResponse()
    err = c.Send(request, response)
    return
}

func NewInitGroupRequest() (request *InitGroupRequest) {
    request = &InitGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "InitGroup")
    return
}

func NewInitGroupResponse() (response *InitGroupResponse) {
    response = &InitGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 初始化一颗分组树
func (c *Client) InitGroup(request *InitGroupRequest) (response *InitGroupResponse, err error) {
    if request == nil {
        request = NewInitGroupRequest()
    }
    response = NewInitGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmCallbackRequest() (request *ModifyAlarmCallbackRequest) {
    request = &ModifyAlarmCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmCallback")
    return
}

func NewModifyAlarmCallbackResponse() (response *ModifyAlarmCallbackResponse) {
    response = &ModifyAlarmCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新告警回调地址
func (c *Client) ModifyAlarmCallback(request *ModifyAlarmCallbackRequest) (response *ModifyAlarmCallbackResponse, err error) {
    if request == nil {
        request = NewModifyAlarmCallbackRequest()
    }
    response = NewModifyAlarmCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmReceiversRequest() (request *ModifyAlarmReceiversRequest) {
    request = &ModifyAlarmReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmReceivers")
    return
}

func NewModifyAlarmReceiversResponse() (response *ModifyAlarmReceiversResponse) {
    response = &ModifyAlarmReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改告警接收人
func (c *Client) ModifyAlarmReceivers(request *ModifyAlarmReceiversRequest) (response *ModifyAlarmReceiversResponse, err error) {
    if request == nil {
        request = NewModifyAlarmReceiversRequest()
    }
    response = NewModifyAlarmReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlertPolicyRequest() (request *ModifyAlertPolicyRequest) {
    request = &ModifyAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlertPolicy")
    return
}

func NewModifyAlertPolicyResponse() (response *ModifyAlertPolicyResponse) {
    response = &ModifyAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlertPolicy修改告警策略
func (c *Client) ModifyAlertPolicy(request *ModifyAlertPolicyRequest) (response *ModifyAlertPolicyResponse, err error) {
    if request == nil {
        request = NewModifyAlertPolicyRequest()
    }
    response = NewModifyAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlertPolicyStatusRequest() (request *ModifyAlertPolicyStatusRequest) {
    request = &ModifyAlertPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlertPolicyStatus")
    return
}

func NewModifyAlertPolicyStatusResponse() (response *ModifyAlertPolicyStatusResponse) {
    response = &ModifyAlertPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlertPolicyStatus 更新告警策略状态
func (c *Client) ModifyAlertPolicyStatus(request *ModifyAlertPolicyStatusRequest) (response *ModifyAlertPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyAlertPolicyStatusRequest()
    }
    response = NewModifyAlertPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAttributeRequest() (request *ModifyAttributeRequest) {
    request = &ModifyAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAttribute")
    return
}

func NewModifyAttributeResponse() (response *ModifyAttributeResponse) {
    response = &ModifyAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指标
func (c *Client) ModifyAttribute(request *ModifyAttributeRequest) (response *ModifyAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAttributeRequest()
    }
    response = NewModifyAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCMChartRequest() (request *ModifyCCMChartRequest) {
    request = &ModifyCCMChartRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyCCMChart")
    return
}

func NewModifyCCMChartResponse() (response *ModifyCCMChartResponse) {
    response = &ModifyCCMChartResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改分组视图自定义监控图表
func (c *Client) ModifyCCMChart(request *ModifyCCMChartRequest) (response *ModifyCCMChartResponse, err error) {
    if request == nil {
        request = NewModifyCCMChartRequest()
    }
    response = NewModifyCCMChartResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCMGroupStrategyRequest() (request *ModifyCCMGroupStrategyRequest) {
    request = &ModifyCCMGroupStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyCCMGroupStrategy")
    return
}

func NewModifyCCMGroupStrategyResponse() (response *ModifyCCMGroupStrategyResponse) {
    response = &ModifyCCMGroupStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改实例分组部分实例告警历史信息置为失效状态
func (c *Client) ModifyCCMGroupStrategy(request *ModifyCCMGroupStrategyRequest) (response *ModifyCCMGroupStrategyResponse, err error) {
    if request == nil {
        request = NewModifyCCMGroupStrategyRequest()
    }
    response = NewModifyCCMGroupStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCMGroupViewRequest() (request *ModifyCCMGroupViewRequest) {
    request = &ModifyCCMGroupViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyCCMGroupView")
    return
}

func NewModifyCCMGroupViewResponse() (response *ModifyCCMGroupViewResponse) {
    response = &ModifyCCMGroupViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改分组视图
func (c *Client) ModifyCCMGroupView(request *ModifyCCMGroupViewRequest) (response *ModifyCCMGroupViewResponse, err error) {
    if request == nil {
        request = NewModifyCCMGroupViewRequest()
    }
    response = NewModifyCCMGroupViewResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConditionsTemplateRequest() (request *ModifyConditionsTemplateRequest) {
    request = &ModifyConditionsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyConditionsTemplate")
    return
}

func NewModifyConditionsTemplateResponse() (response *ModifyConditionsTemplateResponse) {
    response = &ModifyConditionsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改告警策略模板
func (c *Client) ModifyConditionsTemplate(request *ModifyConditionsTemplateRequest) (response *ModifyConditionsTemplateResponse, err error) {
    if request == nil {
        request = NewModifyConditionsTemplateRequest()
    }
    response = NewModifyConditionsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDashboardRequest() (request *ModifyDashboardRequest) {
    request = &ModifyDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyDashboard")
    return
}

func NewModifyDashboardResponse() (response *ModifyDashboardResponse) {
    response = &ModifyDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改dashboard配置
func (c *Client) ModifyDashboard(request *ModifyDashboardRequest) (response *ModifyDashboardResponse, err error) {
    if request == nil {
        request = NewModifyDashboardRequest()
    }
    response = NewModifyDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDashboardViewRequest() (request *ModifyDashboardViewRequest) {
    request = &ModifyDashboardViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyDashboardView")
    return
}

func NewModifyDashboardViewResponse() (response *ModifyDashboardViewResponse) {
    response = &ModifyDashboardViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改监控面板视图
func (c *Client) ModifyDashboardView(request *ModifyDashboardViewRequest) (response *ModifyDashboardViewResponse, err error) {
    if request == nil {
        request = NewModifyDashboardViewRequest()
    }
    response = NewModifyDashboardViewResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
    request = &ModifyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyGroup")
    return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
    response = &ModifyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新分组节点
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    response = NewModifyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceGroupRequest() (request *ModifyInstanceGroupRequest) {
    request = &ModifyInstanceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyInstanceGroup")
    return
}

func NewModifyInstanceGroupResponse() (response *ModifyInstanceGroupResponse) {
    response = &ModifyInstanceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新实例分组
func (c *Client) ModifyInstanceGroup(request *ModifyInstanceGroupRequest) (response *ModifyInstanceGroupResponse, err error) {
    if request == nil {
        request = NewModifyInstanceGroupRequest()
    }
    response = NewModifyInstanceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMetricSetRequest() (request *ModifyMetricSetRequest) {
    request = &ModifyMetricSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyMetricSet")
    return
}

func NewModifyMetricSetResponse() (response *ModifyMetricSetResponse) {
    response = &ModifyMetricSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMetricSet 修改指标集配置
func (c *Client) ModifyMetricSet(request *ModifyMetricSetRequest) (response *ModifyMetricSetResponse, err error) {
    if request == nil {
        request = NewModifyMetricSetRequest()
    }
    response = NewModifyMetricSetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleRequest() (request *ModifyModuleRequest) {
    request = &ModifyModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyModule")
    return
}

func NewModifyModuleResponse() (response *ModifyModuleResponse) {
    response = &ModifyModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新叶子分组节点
func (c *Client) ModifyModule(request *ModifyModuleRequest) (response *ModifyModuleResponse, err error) {
    if request == nil {
        request = NewModifyModuleRequest()
    }
    response = NewModifyModuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMsgPolicyRequest() (request *ModifyMsgPolicyRequest) {
    request = &ModifyMsgPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyMsgPolicy")
    return
}

func NewModifyMsgPolicyResponse() (response *ModifyMsgPolicyResponse) {
    response = &ModifyMsgPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更改自定义消息策略
func (c *Client) ModifyMsgPolicy(request *ModifyMsgPolicyRequest) (response *ModifyMsgPolicyResponse, err error) {
    if request == nil {
        request = NewModifyMsgPolicyRequest()
    }
    response = NewModifyMsgPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNotifyBatchRequest() (request *ModifyNotifyBatchRequest) {
    request = &ModifyNotifyBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyNotifyBatch")
    return
}

func NewModifyNotifyBatchResponse() (response *ModifyNotifyBatchResponse) {
    response = &ModifyNotifyBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量更新告警渠道
func (c *Client) ModifyNotifyBatch(request *ModifyNotifyBatchRequest) (response *ModifyNotifyBatchResponse, err error) {
    if request == nil {
        request = NewModifyNotifyBatchRequest()
    }
    response = NewModifyNotifyBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPolicyGroupRequest() (request *ModifyPolicyGroupRequest) {
    request = &ModifyPolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPolicyGroup")
    return
}

func NewModifyPolicyGroupResponse() (response *ModifyPolicyGroupResponse) {
    response = &ModifyPolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新策略组
func (c *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (response *ModifyPolicyGroupResponse, err error) {
    if request == nil {
        request = NewModifyPolicyGroupRequest()
    }
    response = NewModifyPolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPolicyGroupInfoRequest() (request *ModifyPolicyGroupInfoRequest) {
    request = &ModifyPolicyGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPolicyGroupInfo")
    return
}

func NewModifyPolicyGroupInfoResponse() (response *ModifyPolicyGroupInfoResponse) {
    response = &ModifyPolicyGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新策略组详情
func (c *Client) ModifyPolicyGroupInfo(request *ModifyPolicyGroupInfoRequest) (response *ModifyPolicyGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyPolicyGroupInfoRequest()
    }
    response = NewModifyPolicyGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecoverNotifyBatchRequest() (request *ModifyRecoverNotifyBatchRequest) {
    request = &ModifyRecoverNotifyBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyRecoverNotifyBatch")
    return
}

func NewModifyRecoverNotifyBatchResponse() (response *ModifyRecoverNotifyBatchResponse) {
    response = &ModifyRecoverNotifyBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量编辑恢复告警渠道
func (c *Client) ModifyRecoverNotifyBatch(request *ModifyRecoverNotifyBatchRequest) (response *ModifyRecoverNotifyBatchResponse, err error) {
    if request == nil {
        request = NewModifyRecoverNotifyBatchRequest()
    }
    response = NewModifyRecoverNotifyBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStrategyRequest() (request *ModifyStrategyRequest) {
    request = &ModifyStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyStrategy")
    return
}

func NewModifyStrategyResponse() (response *ModifyStrategyResponse) {
    response = &ModifyStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改告警策略
func (c *Client) ModifyStrategy(request *ModifyStrategyRequest) (response *ModifyStrategyResponse, err error) {
    if request == nil {
        request = NewModifyStrategyRequest()
    }
    response = NewModifyStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStrategyChannelsRequest() (request *ModifyStrategyChannelsRequest) {
    request = &ModifyStrategyChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyStrategyChannels")
    return
}

func NewModifyStrategyChannelsResponse() (response *ModifyStrategyChannelsResponse) {
    response = &ModifyStrategyChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量修改告警策略告警渠道
func (c *Client) ModifyStrategyChannels(request *ModifyStrategyChannelsRequest) (response *ModifyStrategyChannelsResponse, err error) {
    if request == nil {
        request = NewModifyStrategyChannelsRequest()
    }
    response = NewModifyStrategyChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStrategyStatesRequest() (request *ModifyStrategyStatesRequest) {
    request = &ModifyStrategyStatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyStrategyStates")
    return
}

func NewModifyStrategyStatesResponse() (response *ModifyStrategyStatesResponse) {
    response = &ModifyStrategyStatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量修改告警启停状态
func (c *Client) ModifyStrategyStates(request *ModifyStrategyStatesRequest) (response *ModifyStrategyStatesResponse, err error) {
    if request == nil {
        request = NewModifyStrategyStatesRequest()
    }
    response = NewModifyStrategyStatesResponse()
    err = c.Send(request, response)
    return
}

func NewPutMonitorDataRequest() (request *PutMonitorDataRequest) {
    request = &PutMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "PutMonitorData")
    return
}

func NewPutMonitorDataResponse() (response *PutMonitorDataResponse) {
    response = &PutMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 默认接口请求频率限制：50次/秒。
// 默认单租户指标上限：100个。
// 单次上报最多 30 个指标/值对，请求返回错误时，请求中所有的指标/值均不会被保存。
// 
// 上报的时间戳为期望保存的时间戳，建议构造整数分钟时刻的时间戳。
// 时间戳时间范围必须为当前时间到 300 秒前之间。
// 同一 IP 指标对的数据需按分钟先后顺序上报。
func (c *Client) PutMonitorData(request *PutMonitorDataRequest) (response *PutMonitorDataResponse, err error) {
    if request == nil {
        request = NewPutMonitorDataRequest()
    }
    response = NewPutMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewSearchDimensionValueRequest() (request *SearchDimensionValueRequest) {
    request = &SearchDimensionValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "SearchDimensionValue")
    return
}

func NewSearchDimensionValueResponse() (response *SearchDimensionValueResponse) {
    response = &SearchDimensionValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 搜索维度值
func (c *Client) SearchDimensionValue(request *SearchDimensionValueRequest) (response *SearchDimensionValueResponse, err error) {
    if request == nil {
        request = NewSearchDimensionValueRequest()
    }
    response = NewSearchDimensionValueResponse()
    err = c.Send(request, response)
    return
}

func NewSearchLogRequest() (request *SearchLogRequest) {
    request = &SearchLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "SearchLog")
    return
}

func NewSearchLogResponse() (response *SearchLogResponse) {
    response = &SearchLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 搜索CLS日志主题日志
func (c *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
    if request == nil {
        request = NewSearchLogRequest()
    }
    response = NewSearchLogResponse()
    err = c.Send(request, response)
    return
}

func NewSendCustomAlarmMsgRequest() (request *SendCustomAlarmMsgRequest) {
    request = &SendCustomAlarmMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "SendCustomAlarmMsg")
    return
}

func NewSendCustomAlarmMsgResponse() (response *SendCustomAlarmMsgResponse) {
    response = &SendCustomAlarmMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发送自定义消息告警
func (c *Client) SendCustomAlarmMsg(request *SendCustomAlarmMsgRequest) (response *SendCustomAlarmMsgResponse, err error) {
    if request == nil {
        request = NewSendCustomAlarmMsgRequest()
    }
    response = NewSendCustomAlarmMsgResponse()
    err = c.Send(request, response)
    return
}

func NewSetDefaultPolicyGroupRequest() (request *SetDefaultPolicyGroupRequest) {
    request = &SetDefaultPolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "SetDefaultPolicyGroup")
    return
}

func NewSetDefaultPolicyGroupResponse() (response *SetDefaultPolicyGroupResponse) {
    response = &SetDefaultPolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置默认策略
func (c *Client) SetDefaultPolicyGroup(request *SetDefaultPolicyGroupRequest) (response *SetDefaultPolicyGroupResponse, err error) {
    if request == nil {
        request = NewSetDefaultPolicyGroupRequest()
    }
    response = NewSetDefaultPolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewShieldPolicyAlarmRequest() (request *ShieldPolicyAlarmRequest) {
    request = &ShieldPolicyAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ShieldPolicyAlarm")
    return
}

func NewShieldPolicyAlarmResponse() (response *ShieldPolicyAlarmResponse) {
    response = &ShieldPolicyAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置告警屏蔽状态
func (c *Client) ShieldPolicyAlarm(request *ShieldPolicyAlarmRequest) (response *ShieldPolicyAlarmResponse, err error) {
    if request == nil {
        request = NewShieldPolicyAlarmRequest()
    }
    response = NewShieldPolicyAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewSubscribeEventRequest() (request *SubscribeEventRequest) {
    request = &SubscribeEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "SubscribeEvent")
    return
}

func NewSubscribeEventResponse() (response *SubscribeEventResponse) {
    response = &SubscribeEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 订阅平台事件告警
func (c *Client) SubscribeEvent(request *SubscribeEventRequest) (response *SubscribeEventResponse, err error) {
    if request == nil {
        request = NewSubscribeEventRequest()
    }
    response = NewSubscribeEventResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingAllPolicyObjectRequest() (request *UnBindingAllPolicyObjectRequest) {
    request = &UnBindingAllPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingAllPolicyObject")
    return
}

func NewUnBindingAllPolicyObjectResponse() (response *UnBindingAllPolicyObjectResponse) {
    response = &UnBindingAllPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除全部的关联对象
func (c *Client) UnBindingAllPolicyObject(request *UnBindingAllPolicyObjectRequest) (response *UnBindingAllPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingAllPolicyObjectRequest()
    }
    response = NewUnBindingAllPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingInstanceGroupRequest() (request *UnBindingInstanceGroupRequest) {
    request = &UnBindingInstanceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingInstanceGroup")
    return
}

func NewUnBindingInstanceGroupResponse() (response *UnBindingInstanceGroupResponse) {
    response = &UnBindingInstanceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑实例分组的告警策略
func (c *Client) UnBindingInstanceGroup(request *UnBindingInstanceGroupRequest) (response *UnBindingInstanceGroupResponse, err error) {
    if request == nil {
        request = NewUnBindingInstanceGroupRequest()
    }
    response = NewUnBindingInstanceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingPolicyObjectRequest() (request *UnBindingPolicyObjectRequest) {
    request = &UnBindingPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingPolicyObject")
    return
}

func NewUnBindingPolicyObjectResponse() (response *UnBindingPolicyObjectResponse) {
    response = &UnBindingPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除策略的关联对象
func (c *Client) UnBindingPolicyObject(request *UnBindingPolicyObjectRequest) (response *UnBindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingPolicyObjectRequest()
    }
    response = NewUnBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewUnsubscribeEventRequest() (request *UnsubscribeEventRequest) {
    request = &UnsubscribeEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "UnsubscribeEvent")
    return
}

func NewUnsubscribeEventResponse() (response *UnsubscribeEventResponse) {
    response = &UnsubscribeEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 取消订阅平台事件告警
func (c *Client) UnsubscribeEvent(request *UnsubscribeEventRequest) (response *UnsubscribeEventResponse, err error) {
    if request == nil {
        request = NewUnsubscribeEventRequest()
    }
    response = NewUnsubscribeEventResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyAlarmCallbackRequest() (request *VerifyAlarmCallbackRequest) {
    request = &VerifyAlarmCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "VerifyAlarmCallback")
    return
}

func NewVerifyAlarmCallbackResponse() (response *VerifyAlarmCallbackResponse) {
    response = &VerifyAlarmCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 验证用户的回调Url
func (c *Client) VerifyAlarmCallback(request *VerifyAlarmCallbackRequest) (response *VerifyAlarmCallbackResponse, err error) {
    if request == nil {
        request = NewVerifyAlarmCallbackRequest()
    }
    response = NewVerifyAlarmCallbackResponse()
    err = c.Send(request, response)
    return
}
