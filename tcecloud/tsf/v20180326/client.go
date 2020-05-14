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

package v20180326

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-26"

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


func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
    request = &AddClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AddClusterInstances")
    return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
    response = &AddClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加云主机节点至TSF集群
func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
    if request == nil {
        request = NewAddClusterInstancesRequest()
    }
    response = NewAddClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAddInstanceRequest() (request *AddInstanceRequest) {
    request = &AddInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AddInstance")
    return
}

func NewAddInstanceResponse() (response *AddInstanceResponse) {
    response = &AddInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加机器节点至TSF集群
func (c *Client) AddInstance(request *AddInstanceRequest) (response *AddInstanceResponse, err error) {
    if request == nil {
        request = NewAddInstanceRequest()
    }
    response = NewAddInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewAddInstancesRequest() (request *AddInstancesRequest) {
    request = &AddInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AddInstances")
    return
}

func NewAddInstancesResponse() (response *AddInstancesResponse) {
    response = &AddInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加云主机节点至TSF集群
func (c *Client) AddInstances(request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    if request == nil {
        request = NewAddInstancesRequest()
    }
    response = NewAddInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAnalyzeLogSchemaRequest() (request *AnalyzeLogSchemaRequest) {
    request = &AnalyzeLogSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AnalyzeLogSchema")
    return
}

func NewAnalyzeLogSchemaResponse() (response *AnalyzeLogSchemaResponse) {
    response = &AnalyzeLogSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 模拟解析日志 
func (c *Client) AnalyzeLogSchema(request *AnalyzeLogSchemaRequest) (response *AnalyzeLogSchemaResponse, err error) {
    if request == nil {
        request = NewAnalyzeLogSchemaRequest()
    }
    response = NewAnalyzeLogSchemaResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateBusinessLogConfigRequest() (request *AssociateBusinessLogConfigRequest) {
    request = &AssociateBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AssociateBusinessLogConfig")
    return
}

func NewAssociateBusinessLogConfigResponse() (response *AssociateBusinessLogConfigResponse) {
    response = &AssociateBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 关联日志配置项到应用
func (c *Client) AssociateBusinessLogConfig(request *AssociateBusinessLogConfigRequest) (response *AssociateBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewAssociateBusinessLogConfigRequest()
    }
    response = NewAssociateBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewAutoRetryTransactionRequest() (request *AutoRetryTransactionRequest) {
    request = &AutoRetryTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AutoRetryTransaction")
    return
}

func NewAutoRetryTransactionResponse() (response *AutoRetryTransactionResponse) {
    response = &AutoRetryTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开启事务自动重试
func (c *Client) AutoRetryTransaction(request *AutoRetryTransactionRequest) (response *AutoRetryTransactionResponse, err error) {
    if request == nil {
        request = NewAutoRetryTransactionRequest()
    }
    response = NewAutoRetryTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewBindApiGroupRequest() (request *BindApiGroupRequest) {
    request = &BindApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "BindApiGroup")
    return
}

func NewBindApiGroupResponse() (response *BindApiGroupResponse) {
    response = &BindApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网关与API分组批量绑定
func (c *Client) BindApiGroup(request *BindApiGroupRequest) (response *BindApiGroupResponse, err error) {
    if request == nil {
        request = NewBindApiGroupRequest()
    }
    response = NewBindApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindPluginRequest() (request *BindPluginRequest) {
    request = &BindPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "BindPlugin")
    return
}

func NewBindPluginResponse() (response *BindPluginResponse) {
    response = &BindPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 插件与网关分组/API批量绑定
func (c *Client) BindPlugin(request *BindPluginRequest) (response *BindPluginResponse, err error) {
    if request == nil {
        request = NewBindPluginRequest()
    }
    response = NewBindPluginResponse()
    err = c.Send(request, response)
    return
}

func NewChangeApiUsableStatusRequest() (request *ChangeApiUsableStatusRequest) {
    request = &ChangeApiUsableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ChangeApiUsableStatus")
    return
}

func NewChangeApiUsableStatusResponse() (response *ChangeApiUsableStatusResponse) {
    response = &ChangeApiUsableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用或禁用API
func (c *Client) ChangeApiUsableStatus(request *ChangeApiUsableStatusRequest) (response *ChangeApiUsableStatusResponse, err error) {
    if request == nil {
        request = NewChangeApiUsableStatusRequest()
    }
    response = NewChangeApiUsableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewChangeContainerReplicasRequest() (request *ChangeContainerReplicasRequest) {
    request = &ChangeContainerReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ChangeContainerReplicas")
    return
}

func NewChangeContainerReplicasResponse() (response *ChangeContainerReplicasResponse) {
    response = &ChangeContainerReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 容器扩缩容操作
func (c *Client) ChangeContainerReplicas(request *ChangeContainerReplicasRequest) (response *ChangeContainerReplicasResponse, err error) {
    if request == nil {
        request = NewChangeContainerReplicasRequest()
    }
    response = NewChangeContainerReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewChangeSecretStatusRequest() (request *ChangeSecretStatusRequest) {
    request = &ChangeSecretStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ChangeSecretStatus")
    return
}

func NewChangeSecretStatusResponse() (response *ChangeSecretStatusResponse) {
    response = &ChangeSecretStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用（禁用）秘钥
func (c *Client) ChangeSecretStatus(request *ChangeSecretStatusRequest) (response *ChangeSecretStatusResponse, err error) {
    if request == nil {
        request = NewChangeSecretStatusRequest()
    }
    response = NewChangeSecretStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCheckClusterCIDRRequest() (request *CheckClusterCIDRRequest) {
    request = &CheckClusterCIDRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CheckClusterCIDR")
    return
}

func NewCheckClusterCIDRResponse() (response *CheckClusterCIDRResponse) {
    response = &CheckClusterCIDRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查集群CIDR是否可用
func (c *Client) CheckClusterCIDR(request *CheckClusterCIDRRequest) (response *CheckClusterCIDRResponse, err error) {
    if request == nil {
        request = NewCheckClusterCIDRRequest()
    }
    response = NewCheckClusterCIDRResponse()
    err = c.Send(request, response)
    return
}

func NewClearResourceDataRequest() (request *ClearResourceDataRequest) {
    request = &ClearResourceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ClearResourceData")
    return
}

func NewClearResourceDataResponse() (response *ClearResourceDataResponse) {
    response = &ClearResourceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 清理资源
func (c *Client) ClearResourceData(request *ClearResourceDataRequest) (response *ClearResourceDataResponse, err error) {
    if request == nil {
        request = NewClearResourceDataRequest()
    }
    response = NewClearResourceDataResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmPolicyRequest() (request *CreateAlarmPolicyRequest) {
    request = &CreateAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateAlarmPolicy")
    return
}

func NewCreateAlarmPolicyResponse() (response *CreateAlarmPolicyResponse) {
    response = &CreateAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新建告警策略
func (c *Client) CreateAlarmPolicy(request *CreateAlarmPolicyRequest) (response *CreateAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAlarmPolicyRequest()
    }
    response = NewCreateAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmReceiverRequest() (request *CreateAlarmReceiverRequest) {
    request = &CreateAlarmReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateAlarmReceiver")
    return
}

func NewCreateAlarmReceiverResponse() (response *CreateAlarmReceiverResponse) {
    response = &CreateAlarmReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增绑定人员
func (c *Client) CreateAlarmReceiver(request *CreateAlarmReceiverRequest) (response *CreateAlarmReceiverResponse, err error) {
    if request == nil {
        request = NewCreateAlarmReceiverRequest()
    }
    response = NewCreateAlarmReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndDownloadTemplateRequest() (request *CreateAndDownloadTemplateRequest) {
    request = &CreateAndDownloadTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateAndDownloadTemplate")
    return
}

func NewCreateAndDownloadTemplateResponse() (response *CreateAndDownloadTemplateResponse) {
    response = &CreateAndDownloadTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 保存并生成模板
func (c *Client) CreateAndDownloadTemplate(request *CreateAndDownloadTemplateRequest) (response *CreateAndDownloadTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAndDownloadTemplateRequest()
    }
    response = NewCreateAndDownloadTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiAccessRequest() (request *CreateApiAccessRequest) {
    request = &CreateApiAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiAccess")
    return
}

func NewCreateApiAccessResponse() (response *CreateApiAccessResponse) {
    response = &CreateApiAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开启Serverless应用的外网访问
func (c *Client) CreateApiAccess(request *CreateApiAccessRequest) (response *CreateApiAccessResponse, err error) {
    if request == nil {
        request = NewCreateApiAccessRequest()
    }
    response = NewCreateApiAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiGroupRequest() (request *CreateApiGroupRequest) {
    request = &CreateApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiGroup")
    return
}

func NewCreateApiGroupResponse() (response *CreateApiGroupResponse) {
    response = &CreateApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建API分组
func (c *Client) CreateApiGroup(request *CreateApiGroupRequest) (response *CreateApiGroupResponse, err error) {
    if request == nil {
        request = NewCreateApiGroupRequest()
    }
    response = NewCreateApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRateLimitRuleRequest() (request *CreateApiRateLimitRuleRequest) {
    request = &CreateApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiRateLimitRule")
    return
}

func NewCreateApiRateLimitRuleResponse() (response *CreateApiRateLimitRuleResponse) {
    response = &CreateApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建API限流规则
func (c *Client) CreateApiRateLimitRule(request *CreateApiRateLimitRuleRequest) (response *CreateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewCreateApiRateLimitRuleRequest()
    }
    response = NewCreateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApplication")
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建应用
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuthorizationRequest() (request *CreateAuthorizationRequest) {
    request = &CreateAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateAuthorization")
    return
}

func NewCreateAuthorizationResponse() (response *CreateAuthorizationResponse) {
    response = &CreateAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增服务权限配置
func (c *Client) CreateAuthorization(request *CreateAuthorizationRequest) (response *CreateAuthorizationResponse, err error) {
    if request == nil {
        request = NewCreateAuthorizationRequest()
    }
    response = NewCreateAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBusinessLogConfigRequest() (request *CreateBusinessLogConfigRequest) {
    request = &CreateBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateBusinessLogConfig")
    return
}

func NewCreateBusinessLogConfigResponse() (response *CreateBusinessLogConfigResponse) {
    response = &CreateBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建业务日志配置项
func (c *Client) CreateBusinessLogConfig(request *CreateBusinessLogConfigRequest) (response *CreateBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateBusinessLogConfigRequest()
    }
    response = NewCreateBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCircuitBreakerRuleRequest() (request *CreateCircuitBreakerRuleRequest) {
    request = &CreateCircuitBreakerRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateCircuitBreakerRule")
    return
}

func NewCreateCircuitBreakerRuleResponse() (response *CreateCircuitBreakerRuleResponse) {
    response = &CreateCircuitBreakerRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建熔断规则
func (c *Client) CreateCircuitBreakerRule(request *CreateCircuitBreakerRuleRequest) (response *CreateCircuitBreakerRuleResponse, err error) {
    if request == nil {
        request = NewCreateCircuitBreakerRuleRequest()
    }
    response = NewCreateCircuitBreakerRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCkafkaRequest() (request *CreateCkafkaRequest) {
    request = &CreateCkafkaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateCkafka")
    return
}

func NewCreateCkafkaResponse() (response *CreateCkafkaResponse) {
    response = &CreateCkafkaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建ckafka
func (c *Client) CreateCkafka(request *CreateCkafkaRequest) (response *CreateCkafkaResponse, err error) {
    if request == nil {
        request = NewCreateCkafkaRequest()
    }
    response = NewCreateCkafkaResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateCluster")
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

func NewCreateConfigRequest() (request *CreateConfigRequest) {
    request = &CreateConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateConfig")
    return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
    response = &CreateConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建配置项
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    if request == nil {
        request = NewCreateConfigRequest()
    }
    response = NewCreateConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigTemplateRequest() (request *CreateConfigTemplateRequest) {
    request = &CreateConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateConfigTemplate")
    return
}

func NewCreateConfigTemplateResponse() (response *CreateConfigTemplateResponse) {
    response = &CreateConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建参数模板
func (c *Client) CreateConfigTemplate(request *CreateConfigTemplateRequest) (response *CreateConfigTemplateResponse, err error) {
    if request == nil {
        request = NewCreateConfigTemplateRequest()
    }
    response = NewCreateConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContainGroupRequest() (request *CreateContainGroupRequest) {
    request = &CreateContainGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateContainGroup")
    return
}

func NewCreateContainGroupResponse() (response *CreateContainGroupResponse) {
    response = &CreateContainGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建容器部署组
func (c *Client) CreateContainGroup(request *CreateContainGroupRequest) (response *CreateContainGroupResponse, err error) {
    if request == nil {
        request = NewCreateContainGroupRequest()
    }
    response = NewCreateContainGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileConfigRequest() (request *CreateFileConfigRequest) {
    request = &CreateFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateFileConfig")
    return
}

func NewCreateFileConfigResponse() (response *CreateFileConfigResponse) {
    response = &CreateFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建文件配置项
func (c *Client) CreateFileConfig(request *CreateFileConfigRequest) (response *CreateFileConfigResponse, err error) {
    if request == nil {
        request = NewCreateFileConfigRequest()
    }
    response = NewCreateFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatewayApiRequest() (request *CreateGatewayApiRequest) {
    request = &CreateGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayApi")
    return
}

func NewCreateGatewayApiResponse() (response *CreateGatewayApiResponse) {
    response = &CreateGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量导入API至api分组(也支持新建API到分组)
func (c *Client) CreateGatewayApi(request *CreateGatewayApiRequest) (response *CreateGatewayApiResponse, err error) {
    if request == nil {
        request = NewCreateGatewayApiRequest()
    }
    response = NewCreateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatewayJwtPluginRequest() (request *CreateGatewayJwtPluginRequest) {
    request = &CreateGatewayJwtPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayJwtPlugin")
    return
}

func NewCreateGatewayJwtPluginResponse() (response *CreateGatewayJwtPluginResponse) {
    response = &CreateGatewayJwtPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增网关JWT认证插件
func (c *Client) CreateGatewayJwtPlugin(request *CreateGatewayJwtPluginRequest) (response *CreateGatewayJwtPluginResponse, err error) {
    if request == nil {
        request = NewCreateGatewayJwtPluginRequest()
    }
    response = NewCreateGatewayJwtPluginResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatewayOAuthPluginRequest() (request *CreateGatewayOAuthPluginRequest) {
    request = &CreateGatewayOAuthPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayOAuthPlugin")
    return
}

func NewCreateGatewayOAuthPluginResponse() (response *CreateGatewayOAuthPluginResponse) {
    response = &CreateGatewayOAuthPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增网关OAuth认证插件
func (c *Client) CreateGatewayOAuthPlugin(request *CreateGatewayOAuthPluginRequest) (response *CreateGatewayOAuthPluginResponse, err error) {
    if request == nil {
        request = NewCreateGatewayOAuthPluginRequest()
    }
    response = NewCreateGatewayOAuthPluginResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatewayTagPluginRequest() (request *CreateGatewayTagPluginRequest) {
    request = &CreateGatewayTagPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayTagPlugin")
    return
}

func NewCreateGatewayTagPluginResponse() (response *CreateGatewayTagPluginResponse) {
    response = &CreateGatewayTagPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增网关Tag转化插件
func (c *Client) CreateGatewayTagPlugin(request *CreateGatewayTagPluginRequest) (response *CreateGatewayTagPluginResponse, err error) {
    if request == nil {
        request = NewCreateGatewayTagPluginRequest()
    }
    response = NewCreateGatewayTagPluginResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGroup")
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建容器部署组
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupSecretRequest() (request *CreateGroupSecretRequest) {
    request = &CreateGroupSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGroupSecret")
    return
}

func NewCreateGroupSecretResponse() (response *CreateGroupSecretResponse) {
    response = &CreateGroupSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建API分组的秘钥
func (c *Client) CreateGroupSecret(request *CreateGroupSecretRequest) (response *CreateGroupSecretResponse, err error) {
    if request == nil {
        request = NewCreateGroupSecretRequest()
    }
    response = NewCreateGroupSecretResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMachinesRequest() (request *CreateMachinesRequest) {
    request = &CreateMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateMachines")
    return
}

func NewCreateMachinesResponse() (response *CreateMachinesResponse) {
    response = &CreateMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加机器列表
func (c *Client) CreateMachines(request *CreateMachinesRequest) (response *CreateMachinesResponse, err error) {
    if request == nil {
        request = NewCreateMachinesRequest()
    }
    response = NewCreateMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMicroserviceRequest() (request *CreateMicroserviceRequest) {
    request = &CreateMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateMicroservice")
    return
}

func NewCreateMicroserviceResponse() (response *CreateMicroserviceResponse) {
    response = &CreateMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增微服务
func (c *Client) CreateMicroservice(request *CreateMicroserviceRequest) (response *CreateMicroserviceResponse, err error) {
    if request == nil {
        request = NewCreateMicroserviceRequest()
    }
    response = NewCreateMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMonitorStatisticsPolicyRequest() (request *CreateMonitorStatisticsPolicyRequest) {
    request = &CreateMonitorStatisticsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateMonitorStatisticsPolicy")
    return
}

func NewCreateMonitorStatisticsPolicyResponse() (response *CreateMonitorStatisticsPolicyResponse) {
    response = &CreateMonitorStatisticsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建监控统计策略
func (c *Client) CreateMonitorStatisticsPolicy(request *CreateMonitorStatisticsPolicyRequest) (response *CreateMonitorStatisticsPolicyResponse, err error) {
    if request == nil {
        request = NewCreateMonitorStatisticsPolicyRequest()
    }
    response = NewCreateMonitorStatisticsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateNamespace")
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建命名空间
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyDocumentRequest() (request *CreatePolicyDocumentRequest) {
    request = &CreatePolicyDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePolicyDocument")
    return
}

func NewCreatePolicyDocumentResponse() (response *CreatePolicyDocumentResponse) {
    response = &CreatePolicyDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建CAM策略文档
func (c *Client) CreatePolicyDocument(request *CreatePolicyDocumentRequest) (response *CreatePolicyDocumentResponse, err error) {
    if request == nil {
        request = NewCreatePolicyDocumentRequest()
    }
    response = NewCreatePolicyDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProgramRequest() (request *CreateProgramRequest) {
    request = &CreateProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateProgram")
    return
}

func NewCreateProgramResponse() (response *CreateProgramResponse) {
    response = &CreateProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建数据集
func (c *Client) CreateProgram(request *CreateProgramRequest) (response *CreateProgramResponse, err error) {
    if request == nil {
        request = NewCreateProgramRequest()
    }
    response = NewCreateProgramResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublicConfigRequest() (request *CreatePublicConfigRequest) {
    request = &CreatePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePublicConfig")
    return
}

func NewCreatePublicConfigResponse() (response *CreatePublicConfigResponse) {
    response = &CreatePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建公共配置项
func (c *Client) CreatePublicConfig(request *CreatePublicConfigRequest) (response *CreatePublicConfigResponse, err error) {
    if request == nil {
        request = NewCreatePublicConfigRequest()
    }
    response = NewCreatePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRatelimitRequest() (request *CreateRatelimitRequest) {
    request = &CreateRatelimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateRatelimit")
    return
}

func NewCreateRatelimitResponse() (response *CreateRatelimitResponse) {
    response = &CreateRatelimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加限流规则
func (c *Client) CreateRatelimit(request *CreateRatelimitRequest) (response *CreateRatelimitResponse, err error) {
    if request == nil {
        request = NewCreateRatelimitRequest()
    }
    response = NewCreateRatelimitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRegionRequest() (request *CreateRegionRequest) {
    request = &CreateRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateRegion")
    return
}

func NewCreateRegionResponse() (response *CreateRegionResponse) {
    response = &CreateRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加业务地域
func (c *Client) CreateRegion(request *CreateRegionRequest) (response *CreateRegionResponse, err error) {
    if request == nil {
        request = NewCreateRegionRequest()
    }
    response = NewCreateRegionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateRole")
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建角色
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    response = NewCreateRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
    request = &CreateRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateRoute")
    return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
    response = &CreateRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建路由
func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
    if request == nil {
        request = NewCreateRouteRequest()
    }
    response = NewCreateRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRouteRuleRequest() (request *CreateRouteRuleRequest) {
    request = &CreateRouteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateRouteRule")
    return
}

func NewCreateRouteRuleResponse() (response *CreateRouteRuleResponse) {
    response = &CreateRouteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建路由规则
func (c *Client) CreateRouteRule(request *CreateRouteRuleRequest) (response *CreateRouteRuleResponse, err error) {
    if request == nil {
        request = NewCreateRouteRuleRequest()
    }
    response = NewCreateRouteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScalableRuleRequest() (request *CreateScalableRuleRequest) {
    request = &CreateScalableRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateScalableRule")
    return
}

func NewCreateScalableRuleResponse() (response *CreateScalableRuleResponse) {
    response = &CreateScalableRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建弹性扩缩容规则 
func (c *Client) CreateScalableRule(request *CreateScalableRuleRequest) (response *CreateScalableRuleResponse, err error) {
    if request == nil {
        request = NewCreateScalableRuleRequest()
    }
    response = NewCreateScalableRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerlessGroupRequest() (request *CreateServerlessGroupRequest) {
    request = &CreateServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateServerlessGroup")
    return
}

func NewCreateServerlessGroupResponse() (response *CreateServerlessGroupResponse) {
    response = &CreateServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建Serverless部署组
func (c *Client) CreateServerlessGroup(request *CreateServerlessGroupRequest) (response *CreateServerlessGroupResponse, err error) {
    if request == nil {
        request = NewCreateServerlessGroupRequest()
    }
    response = NewCreateServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceInstanceRequest() (request *CreateServiceInstanceRequest) {
    request = &CreateServiceInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateServiceInstance")
    return
}

func NewCreateServiceInstanceResponse() (response *CreateServiceInstanceResponse) {
    response = &CreateServiceInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加模块实例节点
func (c *Client) CreateServiceInstance(request *CreateServiceInstanceRequest) (response *CreateServiceInstanceResponse, err error) {
    if request == nil {
        request = NewCreateServiceInstanceRequest()
    }
    response = NewCreateServiceInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTemplateRequest() (request *CreateTemplateRequest) {
    request = &CreateTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateTemplate")
    return
}

func NewCreateTemplateResponse() (response *CreateTemplateResponse) {
    response = &CreateTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 保存模板草稿
func (c *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTemplateRequest()
    }
    response = NewCreateTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTsfZoneRequest() (request *CreateTsfZoneRequest) {
    request = &CreateTsfZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateTsfZone")
    return
}

func NewCreateTsfZoneResponse() (response *CreateTsfZoneResponse) {
    response = &CreateTsfZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建基础可用区
func (c *Client) CreateTsfZone(request *CreateTsfZoneRequest) (response *CreateTsfZoneResponse, err error) {
    if request == nil {
        request = NewCreateTsfZoneRequest()
    }
    response = NewCreateTsfZoneResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
    request = &CreateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateZone")
    return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
    response = &CreateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建业务可用区
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    if request == nil {
        request = NewCreateZoneRequest()
    }
    response = NewCreateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmPolicyRequest() (request *DeleteAlarmPolicyRequest) {
    request = &DeleteAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteAlarmPolicy")
    return
}

func NewDeleteAlarmPolicyResponse() (response *DeleteAlarmPolicyResponse) {
    response = &DeleteAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除告警策略
func (c *Client) DeleteAlarmPolicy(request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmPolicyRequest()
    }
    response = NewDeleteAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmReceiversRequest() (request *DeleteAlarmReceiversRequest) {
    request = &DeleteAlarmReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteAlarmReceivers")
    return
}

func NewDeleteAlarmReceiversResponse() (response *DeleteAlarmReceiversResponse) {
    response = &DeleteAlarmReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除人员
func (c *Client) DeleteAlarmReceivers(request *DeleteAlarmReceiversRequest) (response *DeleteAlarmReceiversResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmReceiversRequest()
    }
    response = NewDeleteAlarmReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiAccessRequest() (request *DeleteApiAccessRequest) {
    request = &DeleteApiAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApiAccess")
    return
}

func NewDeleteApiAccessResponse() (response *DeleteApiAccessResponse) {
    response = &DeleteApiAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 关闭Serverless应用外网访问
func (c *Client) DeleteApiAccess(request *DeleteApiAccessRequest) (response *DeleteApiAccessResponse, err error) {
    if request == nil {
        request = NewDeleteApiAccessRequest()
    }
    response = NewDeleteApiAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiGroupRequest() (request *DeleteApiGroupRequest) {
    request = &DeleteApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApiGroup")
    return
}

func NewDeleteApiGroupResponse() (response *DeleteApiGroupResponse) {
    response = &DeleteApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除Api分组
func (c *Client) DeleteApiGroup(request *DeleteApiGroupRequest) (response *DeleteApiGroupResponse, err error) {
    if request == nil {
        request = NewDeleteApiGroupRequest()
    }
    response = NewDeleteApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApplication")
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除应用
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuthorizationRequest() (request *DeleteAuthorizationRequest) {
    request = &DeleteAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteAuthorization")
    return
}

func NewDeleteAuthorizationResponse() (response *DeleteAuthorizationResponse) {
    response = &DeleteAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除服务权限配置
func (c *Client) DeleteAuthorization(request *DeleteAuthorizationRequest) (response *DeleteAuthorizationResponse, err error) {
    if request == nil {
        request = NewDeleteAuthorizationRequest()
    }
    response = NewDeleteAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBusinessLogConfigRequest() (request *DeleteBusinessLogConfigRequest) {
    request = &DeleteBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteBusinessLogConfig")
    return
}

func NewDeleteBusinessLogConfigResponse() (response *DeleteBusinessLogConfigResponse) {
    response = &DeleteBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除业务日志配置项
func (c *Client) DeleteBusinessLogConfig(request *DeleteBusinessLogConfigRequest) (response *DeleteBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewDeleteBusinessLogConfigRequest()
    }
    response = NewDeleteBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCircuitBreakerRuleRequest() (request *DeleteCircuitBreakerRuleRequest) {
    request = &DeleteCircuitBreakerRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteCircuitBreakerRule")
    return
}

func NewDeleteCircuitBreakerRuleResponse() (response *DeleteCircuitBreakerRuleResponse) {
    response = &DeleteCircuitBreakerRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除熔断规则
func (c *Client) DeleteCircuitBreakerRule(request *DeleteCircuitBreakerRuleRequest) (response *DeleteCircuitBreakerRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCircuitBreakerRuleRequest()
    }
    response = NewDeleteCircuitBreakerRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteCluster")
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除集群
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
    request = &DeleteConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteConfig")
    return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
    response = &DeleteConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除配置项
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    if request == nil {
        request = NewDeleteConfigRequest()
    }
    response = NewDeleteConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigTemplateRequest() (request *DeleteConfigTemplateRequest) {
    request = &DeleteConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteConfigTemplate")
    return
}

func NewDeleteConfigTemplateResponse() (response *DeleteConfigTemplateResponse) {
    response = &DeleteConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除模板
func (c *Client) DeleteConfigTemplate(request *DeleteConfigTemplateRequest) (response *DeleteConfigTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteConfigTemplateRequest()
    }
    response = NewDeleteConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContainerGroupRequest() (request *DeleteContainerGroupRequest) {
    request = &DeleteContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteContainerGroup")
    return
}

func NewDeleteContainerGroupResponse() (response *DeleteContainerGroupResponse) {
    response = &DeleteContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除容器部署组
func (c *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) (response *DeleteContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteContainerGroupRequest()
    }
    response = NewDeleteContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileConfigRequest() (request *DeleteFileConfigRequest) {
    request = &DeleteFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteFileConfig")
    return
}

func NewDeleteFileConfigResponse() (response *DeleteFileConfigResponse) {
    response = &DeleteFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除文件配置项
func (c *Client) DeleteFileConfig(request *DeleteFileConfigRequest) (response *DeleteFileConfigResponse, err error) {
    if request == nil {
        request = NewDeleteFileConfigRequest()
    }
    response = NewDeleteFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGatewayApiRequest() (request *DeleteGatewayApiRequest) {
    request = &DeleteGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteGatewayApi")
    return
}

func NewDeleteGatewayApiResponse() (response *DeleteGatewayApiResponse) {
    response = &DeleteGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除API
func (c *Client) DeleteGatewayApi(request *DeleteGatewayApiRequest) (response *DeleteGatewayApiResponse, err error) {
    if request == nil {
        request = NewDeleteGatewayApiRequest()
    }
    response = NewDeleteGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGatewayPluginRequest() (request *DeleteGatewayPluginRequest) {
    request = &DeleteGatewayPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteGatewayPlugin")
    return
}

func NewDeleteGatewayPluginResponse() (response *DeleteGatewayPluginResponse) {
    response = &DeleteGatewayPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除网关插件
func (c *Client) DeleteGatewayPlugin(request *DeleteGatewayPluginRequest) (response *DeleteGatewayPluginResponse, err error) {
    if request == nil {
        request = NewDeleteGatewayPluginRequest()
    }
    response = NewDeleteGatewayPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteGroup")
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除容器部署组
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupSecretRequest() (request *DeleteGroupSecretRequest) {
    request = &DeleteGroupSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteGroupSecret")
    return
}

func NewDeleteGroupSecretResponse() (response *DeleteGroupSecretResponse) {
    response = &DeleteGroupSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除API分组的秘钥
func (c *Client) DeleteGroupSecret(request *DeleteGroupSecretRequest) (response *DeleteGroupSecretResponse, err error) {
    if request == nil {
        request = NewDeleteGroupSecretRequest()
    }
    response = NewDeleteGroupSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageTagRequest() (request *DeleteImageTagRequest) {
    request = &DeleteImageTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteImageTag")
    return
}

func NewDeleteImageTagResponse() (response *DeleteImageTagResponse) {
    response = &DeleteImageTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除镜像版本 
func (c *Client) DeleteImageTag(request *DeleteImageTagRequest) (response *DeleteImageTagResponse, err error) {
    if request == nil {
        request = NewDeleteImageTagRequest()
    }
    response = NewDeleteImageTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageTagsRequest() (request *DeleteImageTagsRequest) {
    request = &DeleteImageTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteImageTags")
    return
}

func NewDeleteImageTagsResponse() (response *DeleteImageTagsResponse) {
    response = &DeleteImageTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除镜像版本
func (c *Client) DeleteImageTags(request *DeleteImageTagsRequest) (response *DeleteImageTagsResponse, err error) {
    if request == nil {
        request = NewDeleteImageTagsRequest()
    }
    response = NewDeleteImageTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachinesRequest() (request *DeleteMachinesRequest) {
    request = &DeleteMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteMachines")
    return
}

func NewDeleteMachinesResponse() (response *DeleteMachinesResponse) {
    response = &DeleteMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除机器
func (c *Client) DeleteMachines(request *DeleteMachinesRequest) (response *DeleteMachinesResponse, err error) {
    if request == nil {
        request = NewDeleteMachinesRequest()
    }
    response = NewDeleteMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMicroserviceRequest() (request *DeleteMicroserviceRequest) {
    request = &DeleteMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteMicroservice")
    return
}

func NewDeleteMicroserviceResponse() (response *DeleteMicroserviceResponse) {
    response = &DeleteMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除微服务
func (c *Client) DeleteMicroservice(request *DeleteMicroserviceRequest) (response *DeleteMicroserviceResponse, err error) {
    if request == nil {
        request = NewDeleteMicroserviceRequest()
    }
    response = NewDeleteMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMonitorStatisticsPolicyRequest() (request *DeleteMonitorStatisticsPolicyRequest) {
    request = &DeleteMonitorStatisticsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteMonitorStatisticsPolicy")
    return
}

func NewDeleteMonitorStatisticsPolicyResponse() (response *DeleteMonitorStatisticsPolicyResponse) {
    response = &DeleteMonitorStatisticsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除监控统计策略
func (c *Client) DeleteMonitorStatisticsPolicy(request *DeleteMonitorStatisticsPolicyRequest) (response *DeleteMonitorStatisticsPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteMonitorStatisticsPolicyRequest()
    }
    response = NewDeleteMonitorStatisticsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteNamespace")
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除命名空间
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePkgRequest() (request *DeletePkgRequest) {
    request = &DeletePkgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePkg")
    return
}

func NewDeletePkgResponse() (response *DeletePkgResponse) {
    response = &DeletePkgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除包
func (c *Client) DeletePkg(request *DeletePkgRequest) (response *DeletePkgResponse, err error) {
    if request == nil {
        request = NewDeletePkgRequest()
    }
    response = NewDeletePkgResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePkgsRequest() (request *DeletePkgsRequest) {
    request = &DeletePkgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePkgs")
    return
}

func NewDeletePkgsResponse() (response *DeletePkgsResponse) {
    response = &DeletePkgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从软件仓库批量删除程序包。
// 一次最多支持删除1000个包，数量超过1000，返回UpperDeleteLimit错误。
func (c *Client) DeletePkgs(request *DeletePkgsRequest) (response *DeletePkgsResponse, err error) {
    if request == nil {
        request = NewDeletePkgsRequest()
    }
    response = NewDeletePkgsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePodRequest() (request *DeletePodRequest) {
    request = &DeletePodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePod")
    return
}

func NewDeletePodResponse() (response *DeletePodResponse) {
    response = &DeletePodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除容器服务实例
func (c *Client) DeletePod(request *DeletePodRequest) (response *DeletePodResponse, err error) {
    if request == nil {
        request = NewDeletePodRequest()
    }
    response = NewDeletePodResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProgramRequest() (request *DeleteProgramRequest) {
    request = &DeleteProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteProgram")
    return
}

func NewDeleteProgramResponse() (response *DeleteProgramResponse) {
    response = &DeleteProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除数据集
func (c *Client) DeleteProgram(request *DeleteProgramRequest) (response *DeleteProgramResponse, err error) {
    if request == nil {
        request = NewDeleteProgramRequest()
    }
    response = NewDeleteProgramResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublicConfigRequest() (request *DeletePublicConfigRequest) {
    request = &DeletePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePublicConfig")
    return
}

func NewDeletePublicConfigResponse() (response *DeletePublicConfigResponse) {
    response = &DeletePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除公共配置项
func (c *Client) DeletePublicConfig(request *DeletePublicConfigRequest) (response *DeletePublicConfigResponse, err error) {
    if request == nil {
        request = NewDeletePublicConfigRequest()
    }
    response = NewDeletePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRatelimitRequest() (request *DeleteRatelimitRequest) {
    request = &DeleteRatelimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRatelimit")
    return
}

func NewDeleteRatelimitResponse() (response *DeleteRatelimitResponse) {
    response = &DeleteRatelimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除限流规则
func (c *Client) DeleteRatelimit(request *DeleteRatelimitRequest) (response *DeleteRatelimitResponse, err error) {
    if request == nil {
        request = NewDeleteRatelimitRequest()
    }
    response = NewDeleteRatelimitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRegionRequest() (request *DeleteRegionRequest) {
    request = &DeleteRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRegion")
    return
}

func NewDeleteRegionResponse() (response *DeleteRegionResponse) {
    response = &DeleteRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除业务区域
func (c *Client) DeleteRegion(request *DeleteRegionRequest) (response *DeleteRegionResponse, err error) {
    if request == nil {
        request = NewDeleteRegionRequest()
    }
    response = NewDeleteRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRelatedGroupRequest() (request *DeleteRelatedGroupRequest) {
    request = &DeleteRelatedGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRelatedGroup")
    return
}

func NewDeleteRelatedGroupResponse() (response *DeleteRelatedGroupResponse) {
    response = &DeleteRelatedGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除关联部署组 
func (c *Client) DeleteRelatedGroup(request *DeleteRelatedGroupRequest) (response *DeleteRelatedGroupResponse, err error) {
    if request == nil {
        request = NewDeleteRelatedGroupRequest()
    }
    response = NewDeleteRelatedGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
    request = &DeleteRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRole")
    return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
    response = &DeleteRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除角色
func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    if request == nil {
        request = NewDeleteRoleRequest()
    }
    response = NewDeleteRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
    request = &DeleteRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRoute")
    return
}

func NewDeleteRouteResponse() (response *DeleteRouteResponse) {
    response = &DeleteRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除路由
func (c *Client) DeleteRoute(request *DeleteRouteRequest) (response *DeleteRouteResponse, err error) {
    if request == nil {
        request = NewDeleteRouteRequest()
    }
    response = NewDeleteRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteRuleRequest() (request *DeleteRouteRuleRequest) {
    request = &DeleteRouteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRouteRule")
    return
}

func NewDeleteRouteRuleResponse() (response *DeleteRouteRuleResponse) {
    response = &DeleteRouteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除路由规则
func (c *Client) DeleteRouteRule(request *DeleteRouteRuleRequest) (response *DeleteRouteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRouteRuleRequest()
    }
    response = NewDeleteRouteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScalableRuleRequest() (request *DeleteScalableRuleRequest) {
    request = &DeleteScalableRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteScalableRule")
    return
}

func NewDeleteScalableRuleResponse() (response *DeleteScalableRuleResponse) {
    response = &DeleteScalableRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除弹性伸缩规则 
func (c *Client) DeleteScalableRule(request *DeleteScalableRuleRequest) (response *DeleteScalableRuleResponse, err error) {
    if request == nil {
        request = NewDeleteScalableRuleRequest()
    }
    response = NewDeleteScalableRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServerlessGroupRequest() (request *DeleteServerlessGroupRequest) {
    request = &DeleteServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteServerlessGroup")
    return
}

func NewDeleteServerlessGroupResponse() (response *DeleteServerlessGroupResponse) {
    response = &DeleteServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除Serverless部署组
func (c *Client) DeleteServerlessGroup(request *DeleteServerlessGroupRequest) (response *DeleteServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessGroupRequest()
    }
    response = NewDeleteServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceInstancesRequest() (request *DeleteServiceInstancesRequest) {
    request = &DeleteServiceInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteServiceInstances")
    return
}

func NewDeleteServiceInstancesResponse() (response *DeleteServiceInstancesResponse) {
    response = &DeleteServiceInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除模块实例节点
func (c *Client) DeleteServiceInstances(request *DeleteServiceInstancesRequest) (response *DeleteServiceInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteServiceInstancesRequest()
    }
    response = NewDeleteServiceInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
    request = &DeleteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteTemplate")
    return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
    response = &DeleteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除模板
func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTemplateRequest()
    }
    response = NewDeleteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTsfZoneRequest() (request *DeleteTsfZoneRequest) {
    request = &DeleteTsfZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteTsfZone")
    return
}

func NewDeleteTsfZoneResponse() (response *DeleteTsfZoneResponse) {
    response = &DeleteTsfZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除基础可用区
func (c *Client) DeleteTsfZone(request *DeleteTsfZoneRequest) (response *DeleteTsfZoneResponse, err error) {
    if request == nil {
        request = NewDeleteTsfZoneRequest()
    }
    response = NewDeleteTsfZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
    request = &DeleteZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteZone")
    return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
    response = &DeleteZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除业务可用区
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    if request == nil {
        request = NewDeleteZoneRequest()
    }
    response = NewDeleteZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeployContainGroupRequest() (request *DeployContainGroupRequest) {
    request = &DeployContainGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployContainGroup")
    return
}

func NewDeployContainGroupResponse() (response *DeployContainGroupResponse) {
    response = &DeployContainGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署容器应用
func (c *Client) DeployContainGroup(request *DeployContainGroupRequest) (response *DeployContainGroupResponse, err error) {
    if request == nil {
        request = NewDeployContainGroupRequest()
    }
    response = NewDeployContainGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployContainerGroupRequest() (request *DeployContainerGroupRequest) {
    request = &DeployContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployContainerGroup")
    return
}

func NewDeployContainerGroupResponse() (response *DeployContainerGroupResponse) {
    response = &DeployContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署容器应用
func (c *Client) DeployContainerGroup(request *DeployContainerGroupRequest) (response *DeployContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeployContainerGroupRequest()
    }
    response = NewDeployContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployGroupRequest() (request *DeployGroupRequest) {
    request = &DeployGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployGroup")
    return
}

func NewDeployGroupResponse() (response *DeployGroupResponse) {
    response = &DeployGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署虚拟机部署组应用
func (c *Client) DeployGroup(request *DeployGroupRequest) (response *DeployGroupResponse, err error) {
    if request == nil {
        request = NewDeployGroupRequest()
    }
    response = NewDeployGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployInstanceRequest() (request *DeployInstanceRequest) {
    request = &DeployInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployInstance")
    return
}

func NewDeployInstanceResponse() (response *DeployInstanceResponse) {
    response = &DeployInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署机器
func (c *Client) DeployInstance(request *DeployInstanceRequest) (response *DeployInstanceResponse, err error) {
    if request == nil {
        request = NewDeployInstanceRequest()
    }
    response = NewDeployInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeployServerlessGroupRequest() (request *DeployServerlessGroupRequest) {
    request = &DeployServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployServerlessGroup")
    return
}

func NewDeployServerlessGroupResponse() (response *DeployServerlessGroupResponse) {
    response = &DeployServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署Serverless应用
func (c *Client) DeployServerlessGroup(request *DeployServerlessGroupRequest) (response *DeployServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeployServerlessGroupRequest()
    }
    response = NewDeployServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployTsfModulesRequest() (request *DeployTsfModulesRequest) {
    request = &DeployTsfModulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployTsfModules")
    return
}

func NewDeployTsfModulesResponse() (response *DeployTsfModulesResponse) {
    response = &DeployTsfModulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 模块部署
func (c *Client) DeployTsfModules(request *DeployTsfModulesRequest) (response *DeployTsfModulesResponse, err error) {
    if request == nil {
        request = NewDeployTsfModulesRequest()
    }
    response = NewDeployTsfModulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddibleInstancesRequest() (request *DescribeAddibleInstancesRequest) {
    request = &DescribeAddibleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAddibleInstances")
    return
}

func NewDescribeAddibleInstancesResponse() (response *DescribeAddibleInstancesResponse) {
    response = &DescribeAddibleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群可添加的机器列表
func (c *Client) DescribeAddibleInstances(request *DescribeAddibleInstancesRequest) (response *DescribeAddibleInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeAddibleInstancesRequest()
    }
    response = NewDescribeAddibleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmOverviewListRequest() (request *DescribeAlarmOverviewListRequest) {
    request = &DescribeAlarmOverviewListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAlarmOverviewList")
    return
}

func NewDescribeAlarmOverviewListResponse() (response *DescribeAlarmOverviewListResponse) {
    response = &DescribeAlarmOverviewListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询概览页告警信息列表
func (c *Client) DescribeAlarmOverviewList(request *DescribeAlarmOverviewListRequest) (response *DescribeAlarmOverviewListResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmOverviewListRequest()
    }
    response = NewDescribeAlarmOverviewListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmPolicyRequest() (request *DescribeAlarmPolicyRequest) {
    request = &DescribeAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAlarmPolicy")
    return
}

func NewDescribeAlarmPolicyResponse() (response *DescribeAlarmPolicyResponse) {
    response = &DescribeAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询告警策略
func (c *Client) DescribeAlarmPolicy(request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPolicyRequest()
    }
    response = NewDescribeAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiAccessRequest() (request *DescribeApiAccessRequest) {
    request = &DescribeApiAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiAccess")
    return
}

func NewDescribeApiAccessResponse() (response *DescribeApiAccessResponse) {
    response = &DescribeApiAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Serverless应用的外网访问权限状态
func (c *Client) DescribeApiAccess(request *DescribeApiAccessRequest) (response *DescribeApiAccessResponse, err error) {
    if request == nil {
        request = NewDescribeApiAccessRequest()
    }
    response = NewDescribeApiAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiDetailRequest() (request *DescribeApiDetailRequest) {
    request = &DescribeApiDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiDetail")
    return
}

func NewDescribeApiDetailResponse() (response *DescribeApiDetailResponse) {
    response = &DescribeApiDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API详情
func (c *Client) DescribeApiDetail(request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiDetailRequest()
    }
    response = NewDescribeApiDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiGroupRequest() (request *DescribeApiGroupRequest) {
    request = &DescribeApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroup")
    return
}

func NewDescribeApiGroupResponse() (response *DescribeApiGroupResponse) {
    response = &DescribeApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API分组
func (c *Client) DescribeApiGroup(request *DescribeApiGroupRequest) (response *DescribeApiGroupResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupRequest()
    }
    response = NewDescribeApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiGroupsRequest() (request *DescribeApiGroupsRequest) {
    request = &DescribeApiGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroups")
    return
}

func NewDescribeApiGroupsResponse() (response *DescribeApiGroupsResponse) {
    response = &DescribeApiGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API 分组信息列表
func (c *Client) DescribeApiGroups(request *DescribeApiGroupsRequest) (response *DescribeApiGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupsRequest()
    }
    response = NewDescribeApiGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiRateLimitRulesRequest() (request *DescribeApiRateLimitRulesRequest) {
    request = &DescribeApiRateLimitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiRateLimitRules")
    return
}

func NewDescribeApiRateLimitRulesResponse() (response *DescribeApiRateLimitRulesResponse) {
    response = &DescribeApiRateLimitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API限流规则
func (c *Client) DescribeApiRateLimitRules(request *DescribeApiRateLimitRulesRequest) (response *DescribeApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeApiRateLimitRulesRequest()
    }
    response = NewDescribeApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiUseDetailRequest() (request *DescribeApiUseDetailRequest) {
    request = &DescribeApiUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiUseDetail")
    return
}

func NewDescribeApiUseDetailResponse() (response *DescribeApiUseDetailResponse) {
    response = &DescribeApiUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关API监控明细数据
func (c *Client) DescribeApiUseDetail(request *DescribeApiUseDetailRequest) (response *DescribeApiUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiUseDetailRequest()
    }
    response = NewDescribeApiUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiVersionsRequest() (request *DescribeApiVersionsRequest) {
    request = &DescribeApiVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiVersions")
    return
}

func NewDescribeApiVersionsResponse() (response *DescribeApiVersionsResponse) {
    response = &DescribeApiVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API 版本
func (c *Client) DescribeApiVersions(request *DescribeApiVersionsRequest) (response *DescribeApiVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeApiVersionsRequest()
    }
    response = NewDescribeApiVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApisWithPluginRequest() (request *DescribeApisWithPluginRequest) {
    request = &DescribeApisWithPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApisWithPlugin")
    return
}

func NewDescribeApisWithPluginResponse() (response *DescribeApisWithPluginResponse) {
    response = &DescribeApisWithPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某个插件下可用于绑定的API
func (c *Client) DescribeApisWithPlugin(request *DescribeApisWithPluginRequest) (response *DescribeApisWithPluginResponse, err error) {
    if request == nil {
        request = NewDescribeApisWithPluginRequest()
    }
    response = NewDescribeApisWithPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationRequest() (request *DescribeApplicationRequest) {
    request = &DescribeApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplication")
    return
}

func NewDescribeApplicationResponse() (response *DescribeApplicationResponse) {
    response = &DescribeApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取应用详情
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    response = NewDescribeApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationAttributeRequest() (request *DescribeApplicationAttributeRequest) {
    request = &DescribeApplicationAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationAttribute")
    return
}

func NewDescribeApplicationAttributeResponse() (response *DescribeApplicationAttributeResponse) {
    response = &DescribeApplicationAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取应用列表其它字段，如实例数量信息等
func (c *Client) DescribeApplicationAttribute(request *DescribeApplicationAttributeRequest) (response *DescribeApplicationAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationAttributeRequest()
    }
    response = NewDescribeApplicationAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationBusinessLogConfigRequest() (request *DescribeApplicationBusinessLogConfigRequest) {
    request = &DescribeApplicationBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationBusinessLogConfig")
    return
}

func NewDescribeApplicationBusinessLogConfigResponse() (response *DescribeApplicationBusinessLogConfigResponse) {
    response = &DescribeApplicationBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询应用关联日志配置项信息
func (c *Client) DescribeApplicationBusinessLogConfig(request *DescribeApplicationBusinessLogConfigRequest) (response *DescribeApplicationBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationBusinessLogConfigRequest()
    }
    response = NewDescribeApplicationBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationInstancesRequest() (request *DescribeApplicationInstancesRequest) {
    request = &DescribeApplicationInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationInstances")
    return
}

func NewDescribeApplicationInstancesResponse() (response *DescribeApplicationInstancesResponse) {
    response = &DescribeApplicationInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询应用机器列表
func (c *Client) DescribeApplicationInstances(request *DescribeApplicationInstancesRequest) (response *DescribeApplicationInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationInstancesRequest()
    }
    response = NewDescribeApplicationInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
    request = &DescribeApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplications")
    return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
    response = &DescribeApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取应用列表
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsAttributeRequest() (request *DescribeApplicationsAttributeRequest) {
    request = &DescribeApplicationsAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationsAttribute")
    return
}

func NewDescribeApplicationsAttributeResponse() (response *DescribeApplicationsAttributeResponse) {
    response = &DescribeApplicationsAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) DescribeApplicationsAttribute(request *DescribeApplicationsAttributeRequest) (response *DescribeApplicationsAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsAttributeRequest()
    }
    response = NewDescribeApplicationsAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsOtherRequest() (request *DescribeApplicationsOtherRequest) {
    request = &DescribeApplicationsOtherRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationsOther")
    return
}

func NewDescribeApplicationsOtherResponse() (response *DescribeApplicationsOtherResponse) {
    response = &DescribeApplicationsOtherResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) DescribeApplicationsOther(request *DescribeApplicationsOtherRequest) (response *DescribeApplicationsOtherResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsOtherRequest()
    }
    response = NewDescribeApplicationsOtherResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicatoinServerLogRequest() (request *DescribeApplicatoinServerLogRequest) {
    request = &DescribeApplicatoinServerLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicatoinServerLog")
    return
}

func NewDescribeApplicatoinServerLogResponse() (response *DescribeApplicatoinServerLogResponse) {
    response = &DescribeApplicatoinServerLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取应用服务器的日志
func (c *Client) DescribeApplicatoinServerLog(request *DescribeApplicatoinServerLogRequest) (response *DescribeApplicatoinServerLogResponse, err error) {
    if request == nil {
        request = NewDescribeApplicatoinServerLogRequest()
    }
    response = NewDescribeApplicatoinServerLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthNamespacesRequest() (request *DescribeAuthNamespacesRequest) {
    request = &DescribeAuthNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthNamespaces")
    return
}

func NewDescribeAuthNamespacesResponse() (response *DescribeAuthNamespacesResponse) {
    response = &DescribeAuthNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询租户下可用于授权的命名空间
func (c *Client) DescribeAuthNamespaces(request *DescribeAuthNamespacesRequest) (response *DescribeAuthNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeAuthNamespacesRequest()
    }
    response = NewDescribeAuthNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthorizationRequest() (request *DescribeAuthorizationRequest) {
    request = &DescribeAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorization")
    return
}

func NewDescribeAuthorizationResponse() (response *DescribeAuthorizationResponse) {
    response = &DescribeAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务权限配置（废弃）
func (c *Client) DescribeAuthorization(request *DescribeAuthorizationRequest) (response *DescribeAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeAuthorizationRequest()
    }
    response = NewDescribeAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthorizationInfoRequest() (request *DescribeAuthorizationInfoRequest) {
    request = &DescribeAuthorizationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorizationInfo")
    return
}

func NewDescribeAuthorizationInfoResponse() (response *DescribeAuthorizationInfoResponse) {
    response = &DescribeAuthorizationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务权限配置
func (c *Client) DescribeAuthorizationInfo(request *DescribeAuthorizationInfoRequest) (response *DescribeAuthorizationInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAuthorizationInfoRequest()
    }
    response = NewDescribeAuthorizationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthorizationMicroservicesRequest() (request *DescribeAuthorizationMicroservicesRequest) {
    request = &DescribeAuthorizationMicroservicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorizationMicroservices")
    return
}

func NewDescribeAuthorizationMicroservicesResponse() (response *DescribeAuthorizationMicroservicesResponse) {
    response = &DescribeAuthorizationMicroservicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务权限的微服务列表
func (c *Client) DescribeAuthorizationMicroservices(request *DescribeAuthorizationMicroservicesRequest) (response *DescribeAuthorizationMicroservicesResponse, err error) {
    if request == nil {
        request = NewDescribeAuthorizationMicroservicesRequest()
    }
    response = NewDescribeAuthorizationMicroservicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthorizationTypeRequest() (request *DescribeAuthorizationTypeRequest) {
    request = &DescribeAuthorizationTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorizationType")
    return
}

func NewDescribeAuthorizationTypeResponse() (response *DescribeAuthorizationTypeResponse) {
    response = &DescribeAuthorizationTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询鉴权类型，返回：D：未启用；B：黑名单模式；W：白名单模式
func (c *Client) DescribeAuthorizationType(request *DescribeAuthorizationTypeRequest) (response *DescribeAuthorizationTypeResponse, err error) {
    if request == nil {
        request = NewDescribeAuthorizationTypeRequest()
    }
    response = NewDescribeAuthorizationTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthorizationsRequest() (request *DescribeAuthorizationsRequest) {
    request = &DescribeAuthorizationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorizations")
    return
}

func NewDescribeAuthorizationsResponse() (response *DescribeAuthorizationsResponse) {
    response = &DescribeAuthorizationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务权限配置列表
func (c *Client) DescribeAuthorizations(request *DescribeAuthorizationsRequest) (response *DescribeAuthorizationsResponse, err error) {
    if request == nil {
        request = NewDescribeAuthorizationsRequest()
    }
    response = NewDescribeAuthorizationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBusinessLogConfigRequest() (request *DescribeBusinessLogConfigRequest) {
    request = &DescribeBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeBusinessLogConfig")
    return
}

func NewDescribeBusinessLogConfigResponse() (response *DescribeBusinessLogConfigResponse) {
    response = &DescribeBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询业务日志配置项信息
func (c *Client) DescribeBusinessLogConfig(request *DescribeBusinessLogConfigRequest) (response *DescribeBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBusinessLogConfigRequest()
    }
    response = NewDescribeBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBusinessLogConfigsRequest() (request *DescribeBusinessLogConfigsRequest) {
    request = &DescribeBusinessLogConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeBusinessLogConfigs")
    return
}

func NewDescribeBusinessLogConfigsResponse() (response *DescribeBusinessLogConfigsResponse) {
    response = &DescribeBusinessLogConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询日志配置项列表
func (c *Client) DescribeBusinessLogConfigs(request *DescribeBusinessLogConfigsRequest) (response *DescribeBusinessLogConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeBusinessLogConfigsRequest()
    }
    response = NewDescribeBusinessLogConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCircuitBreakerRuleRequest() (request *DescribeCircuitBreakerRuleRequest) {
    request = &DescribeCircuitBreakerRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeCircuitBreakerRule")
    return
}

func NewDescribeCircuitBreakerRuleResponse() (response *DescribeCircuitBreakerRuleResponse) {
    response = &DescribeCircuitBreakerRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询单条熔断规则
func (c *Client) DescribeCircuitBreakerRule(request *DescribeCircuitBreakerRuleRequest) (response *DescribeCircuitBreakerRuleResponse, err error) {
    if request == nil {
        request = NewDescribeCircuitBreakerRuleRequest()
    }
    response = NewDescribeCircuitBreakerRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCircuitBreakerRulesRequest() (request *DescribeCircuitBreakerRulesRequest) {
    request = &DescribeCircuitBreakerRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeCircuitBreakerRules")
    return
}

func NewDescribeCircuitBreakerRulesResponse() (response *DescribeCircuitBreakerRulesResponse) {
    response = &DescribeCircuitBreakerRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询熔断规则列表
func (c *Client) DescribeCircuitBreakerRules(request *DescribeCircuitBreakerRulesRequest) (response *DescribeCircuitBreakerRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCircuitBreakerRulesRequest()
    }
    response = NewDescribeCircuitBreakerRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudMonitorGroupsRequest() (request *DescribeCloudMonitorGroupsRequest) {
    request = &DescribeCloudMonitorGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeCloudMonitorGroups")
    return
}

func NewDescribeCloudMonitorGroupsResponse() (response *DescribeCloudMonitorGroupsResponse) {
    response = &DescribeCloudMonitorGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取部署组列表，提供给云监控
func (c *Client) DescribeCloudMonitorGroups(request *DescribeCloudMonitorGroupsRequest) (response *DescribeCloudMonitorGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudMonitorGroupsRequest()
    }
    response = NewDescribeCloudMonitorGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRequest() (request *DescribeClusterRequest) {
    request = &DescribeClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeCluster")
    return
}

func NewDescribeClusterResponse() (response *DescribeClusterResponse) {
    response = &DescribeClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群信息
func (c *Client) DescribeCluster(request *DescribeClusterRequest) (response *DescribeClusterResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRequest()
    }
    response = NewDescribeClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstanceCountRequest() (request *DescribeClusterInstanceCountRequest) {
    request = &DescribeClusterInstanceCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterInstanceCount")
    return
}

func NewDescribeClusterInstanceCountResponse() (response *DescribeClusterInstanceCountResponse) {
    response = &DescribeClusterInstanceCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群机器统计信息
func (c *Client) DescribeClusterInstanceCount(request *DescribeClusterInstanceCountRequest) (response *DescribeClusterInstanceCountResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstanceCountRequest()
    }
    response = NewDescribeClusterInstanceCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterInstances")
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群实例
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterLimitResourceRequest() (request *DescribeClusterLimitResourceRequest) {
    request = &DescribeClusterLimitResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterLimitResource")
    return
}

func NewDescribeClusterLimitResourceResponse() (response *DescribeClusterLimitResourceResponse) {
    response = &DescribeClusterLimitResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群剩余资源
func (c *Client) DescribeClusterLimitResource(request *DescribeClusterLimitResourceRequest) (response *DescribeClusterLimitResourceResponse, err error) {
    if request == nil {
        request = NewDescribeClusterLimitResourceRequest()
    }
    response = NewDescribeClusterLimitResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterSchedulabilityRequest() (request *DescribeClusterSchedulabilityRequest) {
    request = &DescribeClusterSchedulabilityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterSchedulability")
    return
}

func NewDescribeClusterSchedulabilityResponse() (response *DescribeClusterSchedulabilityResponse) {
    response = &DescribeClusterSchedulabilityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建容器部署组时，根据用户所需部署组资源确定集群剩余资源是否满足调度需求
func (c *Client) DescribeClusterSchedulability(request *DescribeClusterSchedulabilityRequest) (response *DescribeClusterSchedulabilityResponse, err error) {
    if request == nil {
        request = NewDescribeClusterSchedulabilityRequest()
    }
    response = NewDescribeClusterSchedulabilityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusters")
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群列表
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCommonPkgRequest() (request *DescribeCommonPkgRequest) {
    request = &DescribeCommonPkgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeCommonPkg")
    return
}

func NewDescribeCommonPkgResponse() (response *DescribeCommonPkgResponse) {
    response = &DescribeCommonPkgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取公共包信息
func (c *Client) DescribeCommonPkg(request *DescribeCommonPkgRequest) (response *DescribeCommonPkgResponse, err error) {
    if request == nil {
        request = NewDescribeCommonPkgRequest()
    }
    response = NewDescribeCommonPkgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
    request = &DescribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfig")
    return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
    response = &DescribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigReleaseLogsRequest() (request *DescribeConfigReleaseLogsRequest) {
    request = &DescribeConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleaseLogs")
    return
}

func NewDescribeConfigReleaseLogsResponse() (response *DescribeConfigReleaseLogsResponse) {
    response = &DescribeConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置发布历史
func (c *Client) DescribeConfigReleaseLogs(request *DescribeConfigReleaseLogsRequest) (response *DescribeConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleaseLogsRequest()
    }
    response = NewDescribeConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigReleasesRequest() (request *DescribeConfigReleasesRequest) {
    request = &DescribeConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleases")
    return
}

func NewDescribeConfigReleasesResponse() (response *DescribeConfigReleasesResponse) {
    response = &DescribeConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置发布信息
func (c *Client) DescribeConfigReleases(request *DescribeConfigReleasesRequest) (response *DescribeConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleasesRequest()
    }
    response = NewDescribeConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigSummaryRequest() (request *DescribeConfigSummaryRequest) {
    request = &DescribeConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigSummary")
    return
}

func NewDescribeConfigSummaryResponse() (response *DescribeConfigSummaryResponse) {
    response = &DescribeConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置汇总列表
func (c *Client) DescribeConfigSummary(request *DescribeConfigSummaryRequest) (response *DescribeConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeConfigSummaryRequest()
    }
    response = NewDescribeConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigTemplateRequest() (request *DescribeConfigTemplateRequest) {
    request = &DescribeConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigTemplate")
    return
}

func NewDescribeConfigTemplateResponse() (response *DescribeConfigTemplateResponse) {
    response = &DescribeConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导入配置
func (c *Client) DescribeConfigTemplate(request *DescribeConfigTemplateRequest) (response *DescribeConfigTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeConfigTemplateRequest()
    }
    response = NewDescribeConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigTemplatesRequest() (request *DescribeConfigTemplatesRequest) {
    request = &DescribeConfigTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigTemplates")
    return
}

func NewDescribeConfigTemplatesResponse() (response *DescribeConfigTemplatesResponse) {
    response = &DescribeConfigTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置模板
func (c *Client) DescribeConfigTemplates(request *DescribeConfigTemplatesRequest) (response *DescribeConfigTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigTemplatesRequest()
    }
    response = NewDescribeConfigTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
    request = &DescribeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigs")
    return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
    response = &DescribeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置项列表
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    response = NewDescribeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupAttributeRequest() (request *DescribeContainerGroupAttributeRequest) {
    request = &DescribeContainerGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupAttribute")
    return
}

func NewDescribeContainerGroupAttributeResponse() (response *DescribeContainerGroupAttributeResponse) {
    response = &DescribeContainerGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取部署组其他字段-用于前端并发调用
func (c *Client) DescribeContainerGroupAttribute(request *DescribeContainerGroupAttributeRequest) (response *DescribeContainerGroupAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupAttributeRequest()
    }
    response = NewDescribeContainerGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupDeployInfoRequest() (request *DescribeContainerGroupDeployInfoRequest) {
    request = &DescribeContainerGroupDeployInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDeployInfo")
    return
}

func NewDescribeContainerGroupDeployInfoResponse() (response *DescribeContainerGroupDeployInfoResponse) {
    response = &DescribeContainerGroupDeployInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  获取部署组详情
func (c *Client) DescribeContainerGroupDeployInfo(request *DescribeContainerGroupDeployInfoRequest) (response *DescribeContainerGroupDeployInfoResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDeployInfoRequest()
    }
    response = NewDescribeContainerGroupDeployInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupDetailRequest() (request *DescribeContainerGroupDetailRequest) {
    request = &DescribeContainerGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDetail")
    return
}

func NewDescribeContainerGroupDetailResponse() (response *DescribeContainerGroupDetailResponse) {
    response = &DescribeContainerGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  容器部署组详情
func (c *Client) DescribeContainerGroupDetail(request *DescribeContainerGroupDetailRequest) (response *DescribeContainerGroupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDetailRequest()
    }
    response = NewDescribeContainerGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupsRequest() (request *DescribeContainerGroupsRequest) {
    request = &DescribeContainerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroups")
    return
}

func NewDescribeContainerGroupsResponse() (response *DescribeContainerGroupsResponse) {
    response = &DescribeContainerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 容器部署组列表
func (c *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupsRequest()
    }
    response = NewDescribeContainerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerMicroServiceListRequest() (request *DescribeContainerMicroServiceListRequest) {
    request = &DescribeContainerMicroServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerMicroServiceList")
    return
}

func NewDescribeContainerMicroServiceListResponse() (response *DescribeContainerMicroServiceListResponse) {
    response = &DescribeContainerMicroServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询容器微服务详情
func (c *Client) DescribeContainerMicroServiceList(request *DescribeContainerMicroServiceListRequest) (response *DescribeContainerMicroServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeContainerMicroServiceListRequest()
    }
    response = NewDescribeContainerMicroServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerTasksRequest() (request *DescribeContainerTasksRequest) {
    request = &DescribeContainerTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerTasks")
    return
}

func NewDescribeContainerTasksResponse() (response *DescribeContainerTasksResponse) {
    response = &DescribeContainerTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 变更记录任务列表
func (c *Client) DescribeContainerTasks(request *DescribeContainerTasksRequest) (response *DescribeContainerTasksResponse, err error) {
    if request == nil {
        request = NewDescribeContainerTasksRequest()
    }
    response = NewDescribeContainerTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployGroupRequest() (request *DescribeDeployGroupRequest) {
    request = &DescribeDeployGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeployGroup")
    return
}

func NewDescribeDeployGroupResponse() (response *DescribeDeployGroupResponse) {
    response = &DescribeDeployGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询部署组详情
func (c *Client) DescribeDeployGroup(request *DescribeDeployGroupRequest) (response *DescribeDeployGroupResponse, err error) {
    if request == nil {
        request = NewDescribeDeployGroupRequest()
    }
    response = NewDescribeDeployGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployGroupsRequest() (request *DescribeDeployGroupsRequest) {
    request = &DescribeDeployGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeployGroups")
    return
}

func NewDescribeDeployGroupsResponse() (response *DescribeDeployGroupsResponse) {
    response = &DescribeDeployGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询部署组列表
func (c *Client) DescribeDeployGroups(request *DescribeDeployGroupsRequest) (response *DescribeDeployGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDeployGroupsRequest()
    }
    response = NewDescribeDeployGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDockerForUseRequest() (request *DescribeDockerForUseRequest) {
    request = &DescribeDockerForUseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDockerForUse")
    return
}

func NewDescribeDockerForUseResponse() (response *DescribeDockerForUseResponse) {
    response = &DescribeDockerForUseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取docker使用指引
func (c *Client) DescribeDockerForUse(request *DescribeDockerForUseRequest) (response *DescribeDockerForUseResponse, err error) {
    if request == nil {
        request = NewDescribeDockerForUseRequest()
    }
    response = NewDescribeDockerForUseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDownloadInfoRequest() (request *DescribeDownloadInfoRequest) {
    request = &DescribeDownloadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDownloadInfo")
    return
}

func NewDescribeDownloadInfoResponse() (response *DescribeDownloadInfoResponse) {
    response = &DescribeDownloadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TSF上传的程序包存放在Tce对象存储（COS）中，通过该API可以获取从COS下载程序包需要的信息，包括包所在的桶、存储路径、鉴权信息等，之后使用COS API（或SDK）进行下载。
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
func (c *Client) DescribeDownloadInfo(request *DescribeDownloadInfoRequest) (response *DescribeDownloadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadInfoRequest()
    }
    response = NewDescribeDownloadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileConfigRequest() (request *DescribeFileConfigRequest) {
    request = &DescribeFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfig")
    return
}

func NewDescribeFileConfigResponse() (response *DescribeFileConfigResponse) {
    response = &DescribeFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询文件配置项
func (c *Client) DescribeFileConfig(request *DescribeFileConfigRequest) (response *DescribeFileConfigResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigRequest()
    }
    response = NewDescribeFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileConfigReleaseLogsRequest() (request *DescribeFileConfigReleaseLogsRequest) {
    request = &DescribeFileConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigReleaseLogs")
    return
}

func NewDescribeFileConfigReleaseLogsResponse() (response *DescribeFileConfigReleaseLogsResponse) {
    response = &DescribeFileConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询文件配置发布历史
func (c *Client) DescribeFileConfigReleaseLogs(request *DescribeFileConfigReleaseLogsRequest) (response *DescribeFileConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigReleaseLogsRequest()
    }
    response = NewDescribeFileConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileConfigReleasesRequest() (request *DescribeFileConfigReleasesRequest) {
    request = &DescribeFileConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigReleases")
    return
}

func NewDescribeFileConfigReleasesResponse() (response *DescribeFileConfigReleasesResponse) {
    response = &DescribeFileConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询文件配置项发布信息
func (c *Client) DescribeFileConfigReleases(request *DescribeFileConfigReleasesRequest) (response *DescribeFileConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigReleasesRequest()
    }
    response = NewDescribeFileConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileConfigSummaryRequest() (request *DescribeFileConfigSummaryRequest) {
    request = &DescribeFileConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigSummary")
    return
}

func NewDescribeFileConfigSummaryResponse() (response *DescribeFileConfigSummaryResponse) {
    response = &DescribeFileConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询文件配置汇总列表
func (c *Client) DescribeFileConfigSummary(request *DescribeFileConfigSummaryRequest) (response *DescribeFileConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigSummaryRequest()
    }
    response = NewDescribeFileConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileConfigsRequest() (request *DescribeFileConfigsRequest) {
    request = &DescribeFileConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigs")
    return
}

func NewDescribeFileConfigsResponse() (response *DescribeFileConfigsResponse) {
    response = &DescribeFileConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询文件配置项列表
func (c *Client) DescribeFileConfigs(request *DescribeFileConfigsRequest) (response *DescribeFileConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigsRequest()
    }
    response = NewDescribeFileConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGWOverviewInvocationRequest() (request *DescribeGWOverviewInvocationRequest) {
    request = &DescribeGWOverviewInvocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGWOverviewInvocation")
    return
}

func NewDescribeGWOverviewInvocationResponse() (response *DescribeGWOverviewInvocationResponse) {
    response = &DescribeGWOverviewInvocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关分组或者api监控数据
func (c *Client) DescribeGWOverviewInvocation(request *DescribeGWOverviewInvocationRequest) (response *DescribeGWOverviewInvocationResponse, err error) {
    if request == nil {
        request = NewDescribeGWOverviewInvocationRequest()
    }
    response = NewDescribeGWOverviewInvocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayAllGroupApisRequest() (request *DescribeGatewayAllGroupApisRequest) {
    request = &DescribeGatewayAllGroupApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayAllGroupApis")
    return
}

func NewDescribeGatewayAllGroupApisResponse() (response *DescribeGatewayAllGroupApisResponse) {
    response = &DescribeGatewayAllGroupApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关所有分组下Api列表
func (c *Client) DescribeGatewayAllGroupApis(request *DescribeGatewayAllGroupApisRequest) (response *DescribeGatewayAllGroupApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayAllGroupApisRequest()
    }
    response = NewDescribeGatewayAllGroupApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayApisRequest() (request *DescribeGatewayApisRequest) {
    request = &DescribeGatewayApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayApis")
    return
}

func NewDescribeGatewayApisResponse() (response *DescribeGatewayApisResponse) {
    response = &DescribeGatewayApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API分组下的Api列表信息
func (c *Client) DescribeGatewayApis(request *DescribeGatewayApisRequest) (response *DescribeGatewayApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayApisRequest()
    }
    response = NewDescribeGatewayApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayDailyUseStatisticsRequest() (request *DescribeGatewayDailyUseStatisticsRequest) {
    request = &DescribeGatewayDailyUseStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayDailyUseStatistics")
    return
}

func NewDescribeGatewayDailyUseStatisticsResponse() (response *DescribeGatewayDailyUseStatisticsResponse) {
    response = &DescribeGatewayDailyUseStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关分组或者api日使用统计数据
func (c *Client) DescribeGatewayDailyUseStatistics(request *DescribeGatewayDailyUseStatisticsRequest) (response *DescribeGatewayDailyUseStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayDailyUseStatisticsRequest()
    }
    response = NewDescribeGatewayDailyUseStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayJwtPluginRequest() (request *DescribeGatewayJwtPluginRequest) {
    request = &DescribeGatewayJwtPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayJwtPlugin")
    return
}

func NewDescribeGatewayJwtPluginResponse() (response *DescribeGatewayJwtPluginResponse) {
    response = &DescribeGatewayJwtPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询jwt插件详情信息
func (c *Client) DescribeGatewayJwtPlugin(request *DescribeGatewayJwtPluginRequest) (response *DescribeGatewayJwtPluginResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayJwtPluginRequest()
    }
    response = NewDescribeGatewayJwtPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayMonitorOverviewRequest() (request *DescribeGatewayMonitorOverviewRequest) {
    request = &DescribeGatewayMonitorOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayMonitorOverview")
    return
}

func NewDescribeGatewayMonitorOverviewResponse() (response *DescribeGatewayMonitorOverviewResponse) {
    response = &DescribeGatewayMonitorOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关监控概览
func (c *Client) DescribeGatewayMonitorOverview(request *DescribeGatewayMonitorOverviewRequest) (response *DescribeGatewayMonitorOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayMonitorOverviewRequest()
    }
    response = NewDescribeGatewayMonitorOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayOAuthPluginRequest() (request *DescribeGatewayOAuthPluginRequest) {
    request = &DescribeGatewayOAuthPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayOAuthPlugin")
    return
}

func NewDescribeGatewayOAuthPluginResponse() (response *DescribeGatewayOAuthPluginResponse) {
    response = &DescribeGatewayOAuthPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询OAuth插件详情信息
func (c *Client) DescribeGatewayOAuthPlugin(request *DescribeGatewayOAuthPluginRequest) (response *DescribeGatewayOAuthPluginResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayOAuthPluginRequest()
    }
    response = NewDescribeGatewayOAuthPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayPluginTypesRequest() (request *DescribeGatewayPluginTypesRequest) {
    request = &DescribeGatewayPluginTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayPluginTypes")
    return
}

func NewDescribeGatewayPluginTypesResponse() (response *DescribeGatewayPluginTypesResponse) {
    response = &DescribeGatewayPluginTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关插件支持的类型列表
func (c *Client) DescribeGatewayPluginTypes(request *DescribeGatewayPluginTypesRequest) (response *DescribeGatewayPluginTypesResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayPluginTypesRequest()
    }
    response = NewDescribeGatewayPluginTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayPluginsRequest() (request *DescribeGatewayPluginsRequest) {
    request = &DescribeGatewayPluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayPlugins")
    return
}

func NewDescribeGatewayPluginsResponse() (response *DescribeGatewayPluginsResponse) {
    response = &DescribeGatewayPluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分页查询网关插件列表
func (c *Client) DescribeGatewayPlugins(request *DescribeGatewayPluginsRequest) (response *DescribeGatewayPluginsResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayPluginsRequest()
    }
    response = NewDescribeGatewayPluginsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayTagPluginRequest() (request *DescribeGatewayTagPluginRequest) {
    request = &DescribeGatewayTagPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayTagPlugin")
    return
}

func NewDescribeGatewayTagPluginResponse() (response *DescribeGatewayTagPluginResponse) {
    response = &DescribeGatewayTagPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Tag插件详情信息
func (c *Client) DescribeGatewayTagPlugin(request *DescribeGatewayTagPluginRequest) (response *DescribeGatewayTagPluginResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayTagPluginRequest()
    }
    response = NewDescribeGatewayTagPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroup")
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询虚拟机部署组详情
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupAddibleInstancesRequest() (request *DescribeGroupAddibleInstancesRequest) {
    request = &DescribeGroupAddibleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupAddibleInstances")
    return
}

func NewDescribeGroupAddibleInstancesResponse() (response *DescribeGroupAddibleInstancesResponse) {
    response = &DescribeGroupAddibleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  查询分组可添加的机器列表
func (c *Client) DescribeGroupAddibleInstances(request *DescribeGroupAddibleInstancesRequest) (response *DescribeGroupAddibleInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupAddibleInstancesRequest()
    }
    response = NewDescribeGroupAddibleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupAttributeRequest() (request *DescribeGroupAttributeRequest) {
    request = &DescribeGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupAttribute")
    return
}

func NewDescribeGroupAttributeResponse() (response *DescribeGroupAttributeResponse) {
    response = &DescribeGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取部署组其他属性
func (c *Client) DescribeGroupAttribute(request *DescribeGroupAttributeRequest) (response *DescribeGroupAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeGroupAttributeRequest()
    }
    response = NewDescribeGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupBindedGatewaysRequest() (request *DescribeGroupBindedGatewaysRequest) {
    request = &DescribeGroupBindedGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupBindedGateways")
    return
}

func NewDescribeGroupBindedGatewaysResponse() (response *DescribeGroupBindedGatewaysResponse) {
    response = &DescribeGroupBindedGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某个API分组已绑定的网关部署组信息列表
func (c *Client) DescribeGroupBindedGateways(request *DescribeGroupBindedGatewaysRequest) (response *DescribeGroupBindedGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupBindedGatewaysRequest()
    }
    response = NewDescribeGroupBindedGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupBusinessLogConfigsRequest() (request *DescribeGroupBusinessLogConfigsRequest) {
    request = &DescribeGroupBusinessLogConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupBusinessLogConfigs")
    return
}

func NewDescribeGroupBusinessLogConfigsResponse() (response *DescribeGroupBusinessLogConfigsResponse) {
    response = &DescribeGroupBusinessLogConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询分组管理日志配置列表
func (c *Client) DescribeGroupBusinessLogConfigs(request *DescribeGroupBusinessLogConfigsRequest) (response *DescribeGroupBusinessLogConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupBusinessLogConfigsRequest()
    }
    response = NewDescribeGroupBusinessLogConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupGatewaysRequest() (request *DescribeGroupGatewaysRequest) {
    request = &DescribeGroupGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupGateways")
    return
}

func NewDescribeGroupGatewaysResponse() (response *DescribeGroupGatewaysResponse) {
    response = &DescribeGroupGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某个网关绑定的API 分组信息列表
func (c *Client) DescribeGroupGateways(request *DescribeGroupGatewaysRequest) (response *DescribeGroupGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupGatewaysRequest()
    }
    response = NewDescribeGroupGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupInstancesRequest() (request *DescribeGroupInstancesRequest) {
    request = &DescribeGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupInstances")
    return
}

func NewDescribeGroupInstancesResponse() (response *DescribeGroupInstancesResponse) {
    response = &DescribeGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询虚拟机部署组云主机列表
func (c *Client) DescribeGroupInstances(request *DescribeGroupInstancesRequest) (response *DescribeGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupInstancesRequest()
    }
    response = NewDescribeGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupOtherRequest() (request *DescribeGroupOtherRequest) {
    request = &DescribeGroupOtherRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupOther")
    return
}

func NewDescribeGroupOtherResponse() (response *DescribeGroupOtherResponse) {
    response = &DescribeGroupOtherResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取部署组其他字段
func (c *Client) DescribeGroupOther(request *DescribeGroupOtherRequest) (response *DescribeGroupOtherResponse, err error) {
    if request == nil {
        request = NewDescribeGroupOtherRequest()
    }
    response = NewDescribeGroupOtherResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupSecretsRequest() (request *DescribeGroupSecretsRequest) {
    request = &DescribeGroupSecretsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupSecrets")
    return
}

func NewDescribeGroupSecretsResponse() (response *DescribeGroupSecretsResponse) {
    response = &DescribeGroupSecretsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询秘钥列表
func (c *Client) DescribeGroupSecrets(request *DescribeGroupSecretsRequest) (response *DescribeGroupSecretsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupSecretsRequest()
    }
    response = NewDescribeGroupSecretsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupUseDetailRequest() (request *DescribeGroupUseDetailRequest) {
    request = &DescribeGroupUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupUseDetail")
    return
}

func NewDescribeGroupUseDetailResponse() (response *DescribeGroupUseDetailResponse) {
    response = &DescribeGroupUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询网关分组监控明细数据
func (c *Client) DescribeGroupUseDetail(request *DescribeGroupUseDetailRequest) (response *DescribeGroupUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGroupUseDetailRequest()
    }
    response = NewDescribeGroupUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
    request = &DescribeGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroups")
    return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
    response = &DescribeGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取虚拟机部署组列表
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsWithPluginRequest() (request *DescribeGroupsWithPluginRequest) {
    request = &DescribeGroupsWithPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupsWithPlugin")
    return
}

func NewDescribeGroupsWithPluginResponse() (response *DescribeGroupsWithPluginResponse) {
    response = &DescribeGroupsWithPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某个插件下绑定或未绑定的API分组
func (c *Client) DescribeGroupsWithPlugin(request *DescribeGroupsWithPluginRequest) (response *DescribeGroupsWithPluginResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsWithPluginRequest()
    }
    response = NewDescribeGroupsWithPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRepositoryRequest() (request *DescribeImageRepositoryRequest) {
    request = &DescribeImageRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageRepository")
    return
}

func NewDescribeImageRepositoryResponse() (response *DescribeImageRepositoryResponse) {
    response = &DescribeImageRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 镜像仓库列表 
func (c *Client) DescribeImageRepository(request *DescribeImageRepositoryRequest) (response *DescribeImageRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeImageRepositoryRequest()
    }
    response = NewDescribeImageRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageTagsRequest() (request *DescribeImageTagsRequest) {
    request = &DescribeImageTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageTags")
    return
}

func NewDescribeImageTagsResponse() (response *DescribeImageTagsResponse) {
    response = &DescribeImageTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 镜像版本列表
func (c *Client) DescribeImageTags(request *DescribeImageTagsRequest) (response *DescribeImageTagsResponse, err error) {
    if request == nil {
        request = NewDescribeImageTagsRequest()
    }
    response = NewDescribeImageTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageUserIsExistsRequest() (request *DescribeImageUserIsExistsRequest) {
    request = &DescribeImageUserIsExistsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageUserIsExists")
    return
}

func NewDescribeImageUserIsExistsResponse() (response *DescribeImageUserIsExistsResponse) {
    response = &DescribeImageUserIsExistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 镜像仓库用户是否已经开通
func (c *Client) DescribeImageUserIsExists(request *DescribeImageUserIsExistsRequest) (response *DescribeImageUserIsExistsResponse, err error) {
    if request == nil {
        request = NewDescribeImageUserIsExistsRequest()
    }
    response = NewDescribeImageUserIsExistsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImages")
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询镜像列表
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInovcationIndicatorsRequest() (request *DescribeInovcationIndicatorsRequest) {
    request = &DescribeInovcationIndicatorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInovcationIndicators")
    return
}

func NewDescribeInovcationIndicatorsResponse() (response *DescribeInovcationIndicatorsResponse) {
    response = &DescribeInovcationIndicatorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用监控指标
func (c *Client) DescribeInovcationIndicators(request *DescribeInovcationIndicatorsRequest) (response *DescribeInovcationIndicatorsResponse, err error) {
    if request == nil {
        request = NewDescribeInovcationIndicatorsRequest()
    }
    response = NewDescribeInovcationIndicatorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationStatisticsRequest() (request *DescribeInvocationStatisticsRequest) {
    request = &DescribeInvocationStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationStatistics")
    return
}

func NewDescribeInvocationStatisticsResponse() (response *DescribeInvocationStatisticsResponse) {
    response = &DescribeInvocationStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用监控统计
func (c *Client) DescribeInvocationStatistics(request *DescribeInvocationStatisticsRequest) (response *DescribeInvocationStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationStatisticsRequest()
    }
    response = NewDescribeInvocationStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationStatisticsRatioRequest() (request *DescribeInvocationStatisticsRatioRequest) {
    request = &DescribeInvocationStatisticsRatioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationStatisticsRatio")
    return
}

func NewDescribeInvocationStatisticsRatioResponse() (response *DescribeInvocationStatisticsRatioResponse) {
    response = &DescribeInvocationStatisticsRatioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用统计数据对比
func (c *Client) DescribeInvocationStatisticsRatio(request *DescribeInvocationStatisticsRatioRequest) (response *DescribeInvocationStatisticsRatioResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationStatisticsRatioRequest()
    }
    response = NewDescribeInvocationStatisticsRatioResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseApplicationsRequest() (request *DescribeLicenseApplicationsRequest) {
    request = &DescribeLicenseApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeLicenseApplications")
    return
}

func NewDescribeLicenseApplicationsResponse() (response *DescribeLicenseApplicationsResponse) {
    response = &DescribeLicenseApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取许可申请列表
func (c *Client) DescribeLicenseApplications(request *DescribeLicenseApplicationsRequest) (response *DescribeLicenseApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseApplicationsRequest()
    }
    response = NewDescribeLicenseApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroserviceRequest() (request *DescribeMicroserviceRequest) {
    request = &DescribeMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservice")
    return
}

func NewDescribeMicroserviceResponse() (response *DescribeMicroserviceResponse) {
    response = &DescribeMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询微服务详情
func (c *Client) DescribeMicroservice(request *DescribeMicroserviceRequest) (response *DescribeMicroserviceResponse, err error) {
    if request == nil {
        request = NewDescribeMicroserviceRequest()
    }
    response = NewDescribeMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroservicesRequest() (request *DescribeMicroservicesRequest) {
    request = &DescribeMicroservicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservices")
    return
}

func NewDescribeMicroservicesResponse() (response *DescribeMicroservicesResponse) {
    response = &DescribeMicroservicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取微服务列表
func (c *Client) DescribeMicroservices(request *DescribeMicroservicesRequest) (response *DescribeMicroservicesResponse, err error) {
    if request == nil {
        request = NewDescribeMicroservicesRequest()
    }
    response = NewDescribeMicroservicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroservicesByAssociateApplicationRequest() (request *DescribeMicroservicesByAssociateApplicationRequest) {
    request = &DescribeMicroservicesByAssociateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservicesByAssociateApplication")
    return
}

func NewDescribeMicroservicesByAssociateApplicationResponse() (response *DescribeMicroservicesByAssociateApplicationResponse) {
    response = &DescribeMicroservicesByAssociateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按关联的应用 ID 获取服务列表
func (c *Client) DescribeMicroservicesByAssociateApplication(request *DescribeMicroservicesByAssociateApplicationRequest) (response *DescribeMicroservicesByAssociateApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeMicroservicesByAssociateApplicationRequest()
    }
    response = NewDescribeMicroservicesByAssociateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorStatisticsPolicyRequest() (request *DescribeMonitorStatisticsPolicyRequest) {
    request = &DescribeMonitorStatisticsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMonitorStatisticsPolicy")
    return
}

func NewDescribeMonitorStatisticsPolicyResponse() (response *DescribeMonitorStatisticsPolicyResponse) {
    response = &DescribeMonitorStatisticsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定id的监控统计策略
func (c *Client) DescribeMonitorStatisticsPolicy(request *DescribeMonitorStatisticsPolicyRequest) (response *DescribeMonitorStatisticsPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorStatisticsPolicyRequest()
    }
    response = NewDescribeMonitorStatisticsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsApiListRequest() (request *DescribeMsApiListRequest) {
    request = &DescribeMsApiListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsApiList")
    return
}

func NewDescribeMsApiListResponse() (response *DescribeMsApiListResponse) {
    response = &DescribeMsApiListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务API列表
func (c *Client) DescribeMsApiList(request *DescribeMsApiListRequest) (response *DescribeMsApiListResponse, err error) {
    if request == nil {
        request = NewDescribeMsApiListRequest()
    }
    response = NewDescribeMsApiListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsRouteFallbackRequest() (request *DescribeMsRouteFallbackRequest) {
    request = &DescribeMsRouteFallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsRouteFallback")
    return
}

func NewDescribeMsRouteFallbackResponse() (response *DescribeMsRouteFallbackResponse) {
    response = &DescribeMsRouteFallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询微服务路由保护策略启停状态
func (c *Client) DescribeMsRouteFallback(request *DescribeMsRouteFallbackRequest) (response *DescribeMsRouteFallbackResponse, err error) {
    if request == nil {
        request = NewDescribeMsRouteFallbackRequest()
    }
    response = NewDescribeMsRouteFallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsRunningApplicationsRequest() (request *DescribeMsRunningApplicationsRequest) {
    request = &DescribeMsRunningApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsRunningApplications")
    return
}

func NewDescribeMsRunningApplicationsResponse() (response *DescribeMsRunningApplicationsResponse) {
    response = &DescribeMsRunningApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询微服务运行态应用列表
func (c *Client) DescribeMsRunningApplications(request *DescribeMsRunningApplicationsRequest) (response *DescribeMsRunningApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeMsRunningApplicationsRequest()
    }
    response = NewDescribeMsRunningApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespaceAffinitiesRequest() (request *DescribeNamespaceAffinitiesRequest) {
    request = &DescribeNamespaceAffinitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaceAffinities")
    return
}

func NewDescribeNamespaceAffinitiesResponse() (response *DescribeNamespaceAffinitiesResponse) {
    response = &DescribeNamespaceAffinitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询命名空间就近访问策略列表
func (c *Client) DescribeNamespaceAffinities(request *DescribeNamespaceAffinitiesRequest) (response *DescribeNamespaceAffinitiesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespaceAffinitiesRequest()
    }
    response = NewDescribeNamespaceAffinitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespaceAffinityRequest() (request *DescribeNamespaceAffinityRequest) {
    request = &DescribeNamespaceAffinityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaceAffinity")
    return
}

func NewDescribeNamespaceAffinityResponse() (response *DescribeNamespaceAffinityResponse) {
    response = &DescribeNamespaceAffinityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询命名空间就近访问策略
func (c *Client) DescribeNamespaceAffinity(request *DescribeNamespaceAffinityRequest) (response *DescribeNamespaceAffinityResponse, err error) {
    if request == nil {
        request = NewDescribeNamespaceAffinityRequest()
    }
    response = NewDescribeNamespaceAffinityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespacesRequest() (request *DescribeNamespacesRequest) {
    request = &DescribeNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaces")
    return
}

func NewDescribeNamespacesResponse() (response *DescribeNamespacesResponse) {
    response = &DescribeNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  获取命名空间列表 
func (c *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacesRequest()
    }
    response = NewDescribeNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewInvocationRequest() (request *DescribeOverviewInvocationRequest) {
    request = &DescribeOverviewInvocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeOverviewInvocation")
    return
}

func NewDescribeOverviewInvocationResponse() (response *DescribeOverviewInvocationResponse) {
    response = &DescribeOverviewInvocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务调用监控统计概览
func (c *Client) DescribeOverviewInvocation(request *DescribeOverviewInvocationRequest) (response *DescribeOverviewInvocationResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewInvocationRequest()
    }
    response = NewDescribeOverviewInvocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewResourceRequest() (request *DescribeOverviewResourceRequest) {
    request = &DescribeOverviewResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeOverviewResource")
    return
}

func NewDescribeOverviewResourceResponse() (response *DescribeOverviewResourceResponse) {
    response = &DescribeOverviewResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 概览页资源信息
func (c *Client) DescribeOverviewResource(request *DescribeOverviewResourceRequest) (response *DescribeOverviewResourceResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewResourceRequest()
    }
    response = NewDescribeOverviewResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviweServiceRequest() (request *DescribeOverviweServiceRequest) {
    request = &DescribeOverviweServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeOverviweService")
    return
}

func NewDescribeOverviweServiceResponse() (response *DescribeOverviweServiceResponse) {
    response = &DescribeOverviweServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 概览页微服务信息接口
func (c *Client) DescribeOverviweService(request *DescribeOverviweServiceRequest) (response *DescribeOverviweServiceResponse, err error) {
    if request == nil {
        request = NewDescribeOverviweServiceRequest()
    }
    response = NewDescribeOverviweServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePermissionCategoriesRequest() (request *DescribePermissionCategoriesRequest) {
    request = &DescribePermissionCategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePermissionCategories")
    return
}

func NewDescribePermissionCategoriesResponse() (response *DescribePermissionCategoriesResponse) {
    response = &DescribePermissionCategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询权限组列表
func (c *Client) DescribePermissionCategories(request *DescribePermissionCategoriesRequest) (response *DescribePermissionCategoriesResponse, err error) {
    if request == nil {
        request = NewDescribePermissionCategoriesRequest()
    }
    response = NewDescribePermissionCategoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePkgsRequest() (request *DescribePkgsRequest) {
    request = &DescribePkgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePkgs")
    return
}

func NewDescribePkgsResponse() (response *DescribePkgsResponse) {
    response = &DescribePkgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) DescribePkgs(request *DescribePkgsRequest) (response *DescribePkgsResponse, err error) {
    if request == nil {
        request = NewDescribePkgsRequest()
    }
    response = NewDescribePkgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginInstancesRequest() (request *DescribePluginInstancesRequest) {
    request = &DescribePluginInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePluginInstances")
    return
}

func NewDescribePluginInstancesResponse() (response *DescribePluginInstancesResponse) {
    response = &DescribePluginInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分页查询网关分组/API绑定（或未绑定）的插件列表
func (c *Client) DescribePluginInstances(request *DescribePluginInstancesRequest) (response *DescribePluginInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePluginInstancesRequest()
    }
    response = NewDescribePluginInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePodInstancesRequest() (request *DescribePodInstancesRequest) {
    request = &DescribePodInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePodInstances")
    return
}

func NewDescribePodInstancesResponse() (response *DescribePodInstancesResponse) {
    response = &DescribePodInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取部署组实例列表
func (c *Client) DescribePodInstances(request *DescribePodInstancesRequest) (response *DescribePodInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePodInstancesRequest()
    }
    response = NewDescribePodInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProgramRequest() (request *DescribeProgramRequest) {
    request = &DescribeProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeProgram")
    return
}

func NewDescribeProgramResponse() (response *DescribeProgramResponse) {
    response = &DescribeProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询数据集
func (c *Client) DescribeProgram(request *DescribeProgramRequest) (response *DescribeProgramResponse, err error) {
    if request == nil {
        request = NewDescribeProgramRequest()
    }
    response = NewDescribeProgramResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProgramsRequest() (request *DescribeProgramsRequest) {
    request = &DescribeProgramsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePrograms")
    return
}

func NewDescribeProgramsResponse() (response *DescribeProgramsResponse) {
    response = &DescribeProgramsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询数据集列表
func (c *Client) DescribePrograms(request *DescribeProgramsRequest) (response *DescribeProgramsResponse, err error) {
    if request == nil {
        request = NewDescribeProgramsRequest()
    }
    response = NewDescribeProgramsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigRequest() (request *DescribePublicConfigRequest) {
    request = &DescribePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfig")
    return
}

func NewDescribePublicConfigResponse() (response *DescribePublicConfigResponse) {
    response = &DescribePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置（单条）
func (c *Client) DescribePublicConfig(request *DescribePublicConfigRequest) (response *DescribePublicConfigResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigRequest()
    }
    response = NewDescribePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigReleaseLogsRequest() (request *DescribePublicConfigReleaseLogsRequest) {
    request = &DescribePublicConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleaseLogs")
    return
}

func NewDescribePublicConfigReleaseLogsResponse() (response *DescribePublicConfigReleaseLogsResponse) {
    response = &DescribePublicConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置发布历史
func (c *Client) DescribePublicConfigReleaseLogs(request *DescribePublicConfigReleaseLogsRequest) (response *DescribePublicConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleaseLogsRequest()
    }
    response = NewDescribePublicConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigReleasesRequest() (request *DescribePublicConfigReleasesRequest) {
    request = &DescribePublicConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleases")
    return
}

func NewDescribePublicConfigReleasesResponse() (response *DescribePublicConfigReleasesResponse) {
    response = &DescribePublicConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置发布信息
func (c *Client) DescribePublicConfigReleases(request *DescribePublicConfigReleasesRequest) (response *DescribePublicConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleasesRequest()
    }
    response = NewDescribePublicConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigSummaryRequest() (request *DescribePublicConfigSummaryRequest) {
    request = &DescribePublicConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigSummary")
    return
}

func NewDescribePublicConfigSummaryResponse() (response *DescribePublicConfigSummaryResponse) {
    response = &DescribePublicConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置汇总列表
func (c *Client) DescribePublicConfigSummary(request *DescribePublicConfigSummaryRequest) (response *DescribePublicConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigSummaryRequest()
    }
    response = NewDescribePublicConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigsRequest() (request *DescribePublicConfigsRequest) {
    request = &DescribePublicConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigs")
    return
}

func NewDescribePublicConfigsResponse() (response *DescribePublicConfigsResponse) {
    response = &DescribePublicConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共配置项列表
func (c *Client) DescribePublicConfigs(request *DescribePublicConfigsRequest) (response *DescribePublicConfigsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigsRequest()
    }
    response = NewDescribePublicConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRatelimitRequest() (request *DescribeRatelimitRequest) {
    request = &DescribeRatelimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRatelimit")
    return
}

func NewDescribeRatelimitResponse() (response *DescribeRatelimitResponse) {
    response = &DescribeRatelimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 读取限流规则列表
func (c *Client) DescribeRatelimit(request *DescribeRatelimitRequest) (response *DescribeRatelimitResponse, err error) {
    if request == nil {
        request = NewDescribeRatelimitRequest()
    }
    response = NewDescribeRatelimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRatelimitCommitConfigRequest() (request *DescribeRatelimitCommitConfigRequest) {
    request = &DescribeRatelimitCommitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRatelimitCommitConfig")
    return
}

func NewDescribeRatelimitCommitConfigResponse() (response *DescribeRatelimitCommitConfigResponse) {
    response = &DescribeRatelimitCommitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRatelimitCommitConfig
func (c *Client) DescribeRatelimitCommitConfig(request *DescribeRatelimitCommitConfigRequest) (response *DescribeRatelimitCommitConfigResponse, err error) {
    if request == nil {
        request = NewDescribeRatelimitCommitConfigRequest()
    }
    response = NewDescribeRatelimitCommitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordCodesRequest() (request *DescribeRecordCodesRequest) {
    request = &DescribeRecordCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRecordCodes")
    return
}

func NewDescribeRecordCodesResponse() (response *DescribeRecordCodesResponse) {
    response = &DescribeRecordCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询审计记录术语描述
func (c *Client) DescribeRecordCodes(request *DescribeRecordCodesRequest) (response *DescribeRecordCodesResponse, err error) {
    if request == nil {
        request = NewDescribeRecordCodesRequest()
    }
    response = NewDescribeRecordCodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordsRequest() (request *DescribeRecordsRequest) {
    request = &DescribeRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRecords")
    return
}

func NewDescribeRecordsResponse() (response *DescribeRecordsResponse) {
    response = &DescribeRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询操作记录列表
func (c *Client) DescribeRecords(request *DescribeRecordsRequest) (response *DescribeRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeRecordsRequest()
    }
    response = NewDescribeRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询业务地域
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleasedConfigRequest() (request *DescribeReleasedConfigRequest) {
    request = &DescribeReleasedConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasedConfig")
    return
}

func NewDescribeReleasedConfigResponse() (response *DescribeReleasedConfigResponse) {
    response = &DescribeReleasedConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询group发布的配置
func (c *Client) DescribeReleasedConfig(request *DescribeReleasedConfigRequest) (response *DescribeReleasedConfigResponse, err error) {
    if request == nil {
        request = NewDescribeReleasedConfigRequest()
    }
    response = NewDescribeReleasedConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceConfigRequest() (request *DescribeResourceConfigRequest) {
    request = &DescribeResourceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceConfig")
    return
}

func NewDescribeResourceConfigResponse() (response *DescribeResourceConfigResponse) {
    response = &DescribeResourceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) DescribeResourceConfig(request *DescribeResourceConfigRequest) (response *DescribeResourceConfigResponse, err error) {
    if request == nil {
        request = NewDescribeResourceConfigRequest()
    }
    response = NewDescribeResourceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceUsageRateRequest() (request *DescribeResourceUsageRateRequest) {
    request = &DescribeResourceUsageRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceUsageRate")
    return
}

func NewDescribeResourceUsageRateResponse() (response *DescribeResourceUsageRateResponse) {
    response = &DescribeResourceUsageRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看资源使用情况
func (c *Client) DescribeResourceUsageRate(request *DescribeResourceUsageRateRequest) (response *DescribeResourceUsageRateResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUsageRateRequest()
    }
    response = NewDescribeResourceUsageRateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
    request = &DescribeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeResources")
    return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
    response = &DescribeResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询资源列表
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesRequest()
    }
    response = NewDescribeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoleRequest() (request *DescribeRoleRequest) {
    request = &DescribeRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRole")
    return
}

func NewDescribeRoleResponse() (response *DescribeRoleResponse) {
    response = &DescribeRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询角色
func (c *Client) DescribeRole(request *DescribeRoleRequest) (response *DescribeRoleResponse, err error) {
    if request == nil {
        request = NewDescribeRoleRequest()
    }
    response = NewDescribeRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRolesRequest() (request *DescribeRolesRequest) {
    request = &DescribeRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRoles")
    return
}

func NewDescribeRolesResponse() (response *DescribeRolesResponse) {
    response = &DescribeRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询角色列表
func (c *Client) DescribeRoles(request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
    if request == nil {
        request = NewDescribeRolesRequest()
    }
    response = NewDescribeRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteRequest() (request *DescribeRouteRequest) {
    request = &DescribeRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRoute")
    return
}

func NewDescribeRouteResponse() (response *DescribeRouteResponse) {
    response = &DescribeRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询路由详情
func (c *Client) DescribeRoute(request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
    if request == nil {
        request = NewDescribeRouteRequest()
    }
    response = NewDescribeRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteReleaseHistoryRequest() (request *DescribeRouteReleaseHistoryRequest) {
    request = &DescribeRouteReleaseHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRouteReleaseHistory")
    return
}

func NewDescribeRouteReleaseHistoryResponse() (response *DescribeRouteReleaseHistoryResponse) {
    response = &DescribeRouteReleaseHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询路由规则启用 停用记录
func (c *Client) DescribeRouteReleaseHistory(request *DescribeRouteReleaseHistoryRequest) (response *DescribeRouteReleaseHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeRouteReleaseHistoryRequest()
    }
    response = NewDescribeRouteReleaseHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteRuleRequest() (request *DescribeRouteRuleRequest) {
    request = &DescribeRouteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRouteRule")
    return
}

func NewDescribeRouteRuleResponse() (response *DescribeRouteRuleResponse) {
    response = &DescribeRouteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询路由规则详情
func (c *Client) DescribeRouteRule(request *DescribeRouteRuleRequest) (response *DescribeRouteRuleResponse, err error) {
    if request == nil {
        request = NewDescribeRouteRuleRequest()
    }
    response = NewDescribeRouteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteRulesRequest() (request *DescribeRouteRulesRequest) {
    request = &DescribeRouteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRouteRules")
    return
}

func NewDescribeRouteRulesResponse() (response *DescribeRouteRulesResponse) {
    response = &DescribeRouteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取路由规则列表
func (c *Client) DescribeRouteRules(request *DescribeRouteRulesRequest) (response *DescribeRouteRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteRulesRequest()
    }
    response = NewDescribeRouteRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoutesRequest() (request *DescribeRoutesRequest) {
    request = &DescribeRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRoutes")
    return
}

func NewDescribeRoutesResponse() (response *DescribeRoutesResponse) {
    response = &DescribeRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询路由列表
func (c *Client) DescribeRoutes(request *DescribeRoutesRequest) (response *DescribeRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeRoutesRequest()
    }
    response = NewDescribeRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScalableRuleAttributeRequest() (request *DescribeScalableRuleAttributeRequest) {
    request = &DescribeScalableRuleAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeScalableRuleAttribute")
    return
}

func NewDescribeScalableRuleAttributeResponse() (response *DescribeScalableRuleAttributeResponse) {
    response = &DescribeScalableRuleAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取弹性扩缩容规则属性 
func (c *Client) DescribeScalableRuleAttribute(request *DescribeScalableRuleAttributeRequest) (response *DescribeScalableRuleAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeScalableRuleAttributeRequest()
    }
    response = NewDescribeScalableRuleAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScalableRuleAttributeByGroupRequest() (request *DescribeScalableRuleAttributeByGroupRequest) {
    request = &DescribeScalableRuleAttributeByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeScalableRuleAttributeByGroup")
    return
}

func NewDescribeScalableRuleAttributeByGroupResponse() (response *DescribeScalableRuleAttributeByGroupResponse) {
    response = &DescribeScalableRuleAttributeByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 显示关联的弹性伸缩规则
func (c *Client) DescribeScalableRuleAttributeByGroup(request *DescribeScalableRuleAttributeByGroupRequest) (response *DescribeScalableRuleAttributeByGroupResponse, err error) {
    if request == nil {
        request = NewDescribeScalableRuleAttributeByGroupRequest()
    }
    response = NewDescribeScalableRuleAttributeByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessGroupRequest() (request *DescribeServerlessGroupRequest) {
    request = &DescribeServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeServerlessGroup")
    return
}

func NewDescribeServerlessGroupResponse() (response *DescribeServerlessGroupResponse) {
    response = &DescribeServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Serverless部署组明细
func (c *Client) DescribeServerlessGroup(request *DescribeServerlessGroupRequest) (response *DescribeServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessGroupRequest()
    }
    response = NewDescribeServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessGroupsRequest() (request *DescribeServerlessGroupsRequest) {
    request = &DescribeServerlessGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeServerlessGroups")
    return
}

func NewDescribeServerlessGroupsResponse() (response *DescribeServerlessGroupsResponse) {
    response = &DescribeServerlessGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Serverless部署组列表
func (c *Client) DescribeServerlessGroups(request *DescribeServerlessGroupsRequest) (response *DescribeServerlessGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessGroupsRequest()
    }
    response = NewDescribeServerlessGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceStatisticsRequest() (request *DescribeServiceStatisticsRequest) {
    request = &DescribeServiceStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeServiceStatistics")
    return
}

func NewDescribeServiceStatisticsResponse() (response *DescribeServiceStatisticsResponse) {
    response = &DescribeServiceStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询微服务监控统计
func (c *Client) DescribeServiceStatistics(request *DescribeServiceStatisticsRequest) (response *DescribeServiceStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceStatisticsRequest()
    }
    response = NewDescribeServiceStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServicesRequest() (request *DescribeServicesRequest) {
    request = &DescribeServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeServices")
    return
}

func NewDescribeServicesResponse() (response *DescribeServicesResponse) {
    response = &DescribeServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询产品列表
func (c *Client) DescribeServices(request *DescribeServicesRequest) (response *DescribeServicesResponse, err error) {
    if request == nil {
        request = NewDescribeServicesRequest()
    }
    response = NewDescribeServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleApplicationsRequest() (request *DescribeSimpleApplicationsRequest) {
    request = &DescribeSimpleApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleApplications")
    return
}

func NewDescribeSimpleApplicationsResponse() (response *DescribeSimpleApplicationsResponse) {
    response = &DescribeSimpleApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单应用列表
func (c *Client) DescribeSimpleApplications(request *DescribeSimpleApplicationsRequest) (response *DescribeSimpleApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleApplicationsRequest()
    }
    response = NewDescribeSimpleApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleClustersRequest() (request *DescribeSimpleClustersRequest) {
    request = &DescribeSimpleClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleClusters")
    return
}

func NewDescribeSimpleClustersResponse() (response *DescribeSimpleClustersResponse) {
    response = &DescribeSimpleClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单集群列表
func (c *Client) DescribeSimpleClusters(request *DescribeSimpleClustersRequest) (response *DescribeSimpleClustersResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleClustersRequest()
    }
    response = NewDescribeSimpleClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleGroupsRequest() (request *DescribeSimpleGroupsRequest) {
    request = &DescribeSimpleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleGroups")
    return
}

func NewDescribeSimpleGroupsResponse() (response *DescribeSimpleGroupsResponse) {
    response = &DescribeSimpleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单部署组列表
func (c *Client) DescribeSimpleGroups(request *DescribeSimpleGroupsRequest) (response *DescribeSimpleGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleGroupsRequest()
    }
    response = NewDescribeSimpleGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleInstancesRequest() (request *DescribeSimpleInstancesRequest) {
    request = &DescribeSimpleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleInstances")
    return
}

func NewDescribeSimpleInstancesResponse() (response *DescribeSimpleInstancesResponse) {
    response = &DescribeSimpleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单机器实例列表
func (c *Client) DescribeSimpleInstances(request *DescribeSimpleInstancesRequest) (response *DescribeSimpleInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleInstancesRequest()
    }
    response = NewDescribeSimpleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleNamespacesRequest() (request *DescribeSimpleNamespacesRequest) {
    request = &DescribeSimpleNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleNamespaces")
    return
}

func NewDescribeSimpleNamespacesResponse() (response *DescribeSimpleNamespacesResponse) {
    response = &DescribeSimpleNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询简单命名空间列表 
func (c *Client) DescribeSimpleNamespaces(request *DescribeSimpleNamespacesRequest) (response *DescribeSimpleNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleNamespacesRequest()
    }
    response = NewDescribeSimpleNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSingleContainerGroupsRequest() (request *DescribeSingleContainerGroupsRequest) {
    request = &DescribeSingleContainerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSingleContainerGroups")
    return
}

func NewDescribeSingleContainerGroupsResponse() (response *DescribeSingleContainerGroupsResponse) {
    response = &DescribeSingleContainerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取分组列表-独立菜单
func (c *Client) DescribeSingleContainerGroups(request *DescribeSingleContainerGroupsRequest) (response *DescribeSingleContainerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSingleContainerGroupsRequest()
    }
    response = NewDescribeSingleContainerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubTasksRequest() (request *DescribeSubTasksRequest) {
    request = &DescribeSubTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSubTasks")
    return
}

func NewDescribeSubTasksResponse() (response *DescribeSubTasksResponse) {
    response = &DescribeSubTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取子任务详情
func (c *Client) DescribeSubTasks(request *DescribeSubTasksRequest) (response *DescribeSubTasksResponse, err error) {
    if request == nil {
        request = NewDescribeSubTasksRequest()
    }
    response = NewDescribeSubTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubTransactionsRequest() (request *DescribeSubTransactionsRequest) {
    request = &DescribeSubTransactionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSubTransactions")
    return
}

func NewDescribeSubTransactionsResponse() (response *DescribeSubTransactionsResponse) {
    response = &DescribeSubTransactionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询子事务列表
func (c *Client) DescribeSubTransactions(request *DescribeSubTransactionsRequest) (response *DescribeSubTransactionsResponse, err error) {
    if request == nil {
        request = NewDescribeSubTransactionsRequest()
    }
    response = NewDescribeSubTransactionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateRequest() (request *DescribeTemplateRequest) {
    request = &DescribeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTemplate")
    return
}

func NewDescribeTemplateResponse() (response *DescribeTemplateResponse) {
    response = &DescribeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取模板
func (c *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateRequest()
    }
    response = NewDescribeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
    request = &DescribeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTemplates")
    return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
    response = &DescribeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取模板列表
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTemplatesRequest()
    }
    response = NewDescribeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTransactionRequest() (request *DescribeTransactionRequest) {
    request = &DescribeTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTransaction")
    return
}

func NewDescribeTransactionResponse() (response *DescribeTransactionResponse) {
    response = &DescribeTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询事务详情
func (c *Client) DescribeTransaction(request *DescribeTransactionRequest) (response *DescribeTransactionResponse, err error) {
    if request == nil {
        request = NewDescribeTransactionRequest()
    }
    response = NewDescribeTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTransactionsRequest() (request *DescribeTransactionsRequest) {
    request = &DescribeTransactionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTransactions")
    return
}

func NewDescribeTransactionsResponse() (response *DescribeTransactionsResponse) {
    response = &DescribeTransactionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事务列表
func (c *Client) DescribeTransactions(request *DescribeTransactionsRequest) (response *DescribeTransactionsResponse, err error) {
    if request == nil {
        request = NewDescribeTransactionsRequest()
    }
    response = NewDescribeTransactionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTsfRegionsRequest() (request *DescribeTsfRegionsRequest) {
    request = &DescribeTsfRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTsfRegions")
    return
}

func NewDescribeTsfRegionsResponse() (response *DescribeTsfRegionsResponse) {
    response = &DescribeTsfRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TSF地域列表
func (c *Client) DescribeTsfRegions(request *DescribeTsfRegionsRequest) (response *DescribeTsfRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeTsfRegionsRequest()
    }
    response = NewDescribeTsfRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTsfZonesRequest() (request *DescribeTsfZonesRequest) {
    request = &DescribeTsfZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTsfZones")
    return
}

func NewDescribeTsfZonesResponse() (response *DescribeTsfZonesResponse) {
    response = &DescribeTsfZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询基础可用区
func (c *Client) DescribeTsfZones(request *DescribeTsfZonesRequest) (response *DescribeTsfZonesResponse, err error) {
    if request == nil {
        request = NewDescribeTsfZonesRequest()
    }
    response = NewDescribeTsfZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTsfmanagerZonesRequest() (request *DescribeTsfmanagerZonesRequest) {
    request = &DescribeTsfmanagerZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTsfmanagerZones")
    return
}

func NewDescribeTsfmanagerZonesResponse() (response *DescribeTsfmanagerZonesResponse) {
    response = &DescribeTsfmanagerZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TSF可用区列表
func (c *Client) DescribeTsfmanagerZones(request *DescribeTsfmanagerZonesRequest) (response *DescribeTsfmanagerZonesResponse, err error) {
    if request == nil {
        request = NewDescribeTsfmanagerZonesRequest()
    }
    response = NewDescribeTsfmanagerZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadInfoRequest() (request *DescribeUploadInfoRequest) {
    request = &DescribeUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUploadInfo")
    return
}

func NewDescribeUploadInfoResponse() (response *DescribeUploadInfoResponse) {
    response = &DescribeUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TSF会将软件包上传到Tce对象存储（COS）。调用此接口获取上传信息，如目标地域，桶，包Id，存储路径，鉴权信息等，之后请使用COS API（或SDK）进行上传。
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
func (c *Client) DescribeUploadInfo(request *DescribeUploadInfoRequest) (response *DescribeUploadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadInfoRequest()
    }
    response = NewDescribeUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsableApisRequest() (request *DescribeUsableApisRequest) {
    request = &DescribeUsableApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUsableApis")
    return
}

func NewDescribeUsableApisResponse() (response *DescribeUsableApisResponse) {
    response = &DescribeUsableApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询可用于被导入的API
func (c *Client) DescribeUsableApis(request *DescribeUsableApisRequest) (response *DescribeUsableApisResponse, err error) {
    if request == nil {
        request = NewDescribeUsableApisRequest()
    }
    response = NewDescribeUsableApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsableGatewayGroupsRequest() (request *DescribeUsableGatewayGroupsRequest) {
    request = &DescribeUsableGatewayGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUsableGatewayGroups")
    return
}

func NewDescribeUsableGatewayGroupsResponse() (response *DescribeUsableGatewayGroupsResponse) {
    response = &DescribeUsableGatewayGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某个API分组下可用于被绑定的网关部署组
func (c *Client) DescribeUsableGatewayGroups(request *DescribeUsableGatewayGroupsRequest) (response *DescribeUsableGatewayGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeUsableGatewayGroupsRequest()
    }
    response = NewDescribeUsableGatewayGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeZones")
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询业务可用区
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableCircuitBreakerRuleRequest() (request *DisableCircuitBreakerRuleRequest) {
    request = &DisableCircuitBreakerRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableCircuitBreakerRule")
    return
}

func NewDisableCircuitBreakerRuleResponse() (response *DisableCircuitBreakerRuleResponse) {
    response = &DisableCircuitBreakerRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 禁用熔断规则
func (c *Client) DisableCircuitBreakerRule(request *DisableCircuitBreakerRuleRequest) (response *DisableCircuitBreakerRuleResponse, err error) {
    if request == nil {
        request = NewDisableCircuitBreakerRuleRequest()
    }
    response = NewDisableCircuitBreakerRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDisableFallbackRouteRequest() (request *DisableFallbackRouteRequest) {
    request = &DisableFallbackRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableFallbackRoute")
    return
}

func NewDisableFallbackRouteResponse() (response *DisableFallbackRouteResponse) {
    response = &DisableFallbackRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停用微服务路由保护策略
func (c *Client) DisableFallbackRoute(request *DisableFallbackRouteRequest) (response *DisableFallbackRouteResponse, err error) {
    if request == nil {
        request = NewDisableFallbackRouteRequest()
    }
    response = NewDisableFallbackRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDisableNamespaceAffinityRequest() (request *DisableNamespaceAffinityRequest) {
    request = &DisableNamespaceAffinityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableNamespaceAffinity")
    return
}

func NewDisableNamespaceAffinityResponse() (response *DisableNamespaceAffinityResponse) {
    response = &DisableNamespaceAffinityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停用命名空间就近访问策略
func (c *Client) DisableNamespaceAffinity(request *DisableNamespaceAffinityRequest) (response *DisableNamespaceAffinityResponse, err error) {
    if request == nil {
        request = NewDisableNamespaceAffinityRequest()
    }
    response = NewDisableNamespaceAffinityResponse()
    err = c.Send(request, response)
    return
}

func NewDisableRouteRequest() (request *DisableRouteRequest) {
    request = &DisableRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableRoute")
    return
}

func NewDisableRouteResponse() (response *DisableRouteResponse) {
    response = &DisableRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停用服务路由
func (c *Client) DisableRoute(request *DisableRouteRequest) (response *DisableRouteResponse, err error) {
    if request == nil {
        request = NewDisableRouteRequest()
    }
    response = NewDisableRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDisableRouteRuleRequest() (request *DisableRouteRuleRequest) {
    request = &DisableRouteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableRouteRule")
    return
}

func NewDisableRouteRuleResponse() (response *DisableRouteRuleResponse) {
    response = &DisableRouteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停用路由规则
func (c *Client) DisableRouteRule(request *DisableRouteRuleRequest) (response *DisableRouteRuleResponse, err error) {
    if request == nil {
        request = NewDisableRouteRuleRequest()
    }
    response = NewDisableRouteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateBusinessLogConfigRequest() (request *DisassociateBusinessLogConfigRequest) {
    request = &DisassociateBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisassociateBusinessLogConfig")
    return
}

func NewDisassociateBusinessLogConfigResponse() (response *DisassociateBusinessLogConfigResponse) {
    response = &DisassociateBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 取消关联业务日志配置项和应用
func (c *Client) DisassociateBusinessLogConfig(request *DisassociateBusinessLogConfigRequest) (response *DisassociateBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewDisassociateBusinessLogConfigRequest()
    }
    response = NewDisassociateBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadMsApiRequest() (request *DownloadMsApiRequest) {
    request = &DownloadMsApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DownloadMsApi")
    return
}

func NewDownloadMsApiResponse() (response *DownloadMsApiResponse) {
    response = &DownloadMsApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导出API列表
func (c *Client) DownloadMsApi(request *DownloadMsApiRequest) (response *DownloadMsApiResponse, err error) {
    if request == nil {
        request = NewDownloadMsApiRequest()
    }
    response = NewDownloadMsApiResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadTemplateRequest() (request *DownloadTemplateRequest) {
    request = &DownloadTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DownloadTemplate")
    return
}

func NewDownloadTemplateResponse() (response *DownloadTemplateResponse) {
    response = &DownloadTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 生成模板
func (c *Client) DownloadTemplate(request *DownloadTemplateRequest) (response *DownloadTemplateResponse, err error) {
    if request == nil {
        request = NewDownloadTemplateRequest()
    }
    response = NewDownloadTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDraftApiGroupRequest() (request *DraftApiGroupRequest) {
    request = &DraftApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DraftApiGroup")
    return
}

func NewDraftApiGroupResponse() (response *DraftApiGroupResponse) {
    response = &DraftApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线Api分组
func (c *Client) DraftApiGroup(request *DraftApiGroupRequest) (response *DraftApiGroupResponse, err error) {
    if request == nil {
        request = NewDraftApiGroupRequest()
    }
    response = NewDraftApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDraftPluginRequest() (request *DraftPluginRequest) {
    request = &DraftPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DraftPlugin")
    return
}

func NewDraftPluginResponse() (response *DraftPluginResponse) {
    response = &DraftPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线插件
func (c *Client) DraftPlugin(request *DraftPluginRequest) (response *DraftPluginResponse, err error) {
    if request == nil {
        request = NewDraftPluginRequest()
    }
    response = NewDraftPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDscribeTasksRequest() (request *DscribeTasksRequest) {
    request = &DscribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DscribeTasks")
    return
}

func NewDscribeTasksResponse() (response *DscribeTasksResponse) {
    response = &DscribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取操作日志列表
func (c *Client) DscribeTasks(request *DscribeTasksRequest) (response *DscribeTasksResponse, err error) {
    if request == nil {
        request = NewDscribeTasksRequest()
    }
    response = NewDscribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewEnableCircuitBreakerRuleRequest() (request *EnableCircuitBreakerRuleRequest) {
    request = &EnableCircuitBreakerRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableCircuitBreakerRule")
    return
}

func NewEnableCircuitBreakerRuleResponse() (response *EnableCircuitBreakerRuleResponse) {
    response = &EnableCircuitBreakerRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用熔断规则
func (c *Client) EnableCircuitBreakerRule(request *EnableCircuitBreakerRuleRequest) (response *EnableCircuitBreakerRuleResponse, err error) {
    if request == nil {
        request = NewEnableCircuitBreakerRuleRequest()
    }
    response = NewEnableCircuitBreakerRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEnableFallbackRouteRequest() (request *EnableFallbackRouteRequest) {
    request = &EnableFallbackRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableFallbackRoute")
    return
}

func NewEnableFallbackRouteResponse() (response *EnableFallbackRouteResponse) {
    response = &EnableFallbackRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用微服务路由保护策略
func (c *Client) EnableFallbackRoute(request *EnableFallbackRouteRequest) (response *EnableFallbackRouteResponse, err error) {
    if request == nil {
        request = NewEnableFallbackRouteRequest()
    }
    response = NewEnableFallbackRouteResponse()
    err = c.Send(request, response)
    return
}

func NewEnableNamespaceAffinityRequest() (request *EnableNamespaceAffinityRequest) {
    request = &EnableNamespaceAffinityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableNamespaceAffinity")
    return
}

func NewEnableNamespaceAffinityResponse() (response *EnableNamespaceAffinityResponse) {
    response = &EnableNamespaceAffinityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用命名空间就近访问策略
func (c *Client) EnableNamespaceAffinity(request *EnableNamespaceAffinityRequest) (response *EnableNamespaceAffinityResponse, err error) {
    if request == nil {
        request = NewEnableNamespaceAffinityRequest()
    }
    response = NewEnableNamespaceAffinityResponse()
    err = c.Send(request, response)
    return
}

func NewEnableRouteRequest() (request *EnableRouteRequest) {
    request = &EnableRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableRoute")
    return
}

func NewEnableRouteResponse() (response *EnableRouteResponse) {
    response = &EnableRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用路由
func (c *Client) EnableRoute(request *EnableRouteRequest) (response *EnableRouteResponse, err error) {
    if request == nil {
        request = NewEnableRouteRequest()
    }
    response = NewEnableRouteResponse()
    err = c.Send(request, response)
    return
}

func NewEnableRouteRuleRequest() (request *EnableRouteRuleRequest) {
    request = &EnableRouteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableRouteRule")
    return
}

func NewEnableRouteRuleResponse() (response *EnableRouteRuleResponse) {
    response = &EnableRouteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用路由规则
func (c *Client) EnableRouteRule(request *EnableRouteRuleRequest) (response *EnableRouteRuleResponse, err error) {
    if request == nil {
        request = NewEnableRouteRuleRequest()
    }
    response = NewEnableRouteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewExpandGroupRequest() (request *ExpandGroupRequest) {
    request = &ExpandGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ExpandGroup")
    return
}

func NewExpandGroupResponse() (response *ExpandGroupResponse) {
    response = &ExpandGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 虚拟机部署组添加实例
func (c *Client) ExpandGroup(request *ExpandGroupRequest) (response *ExpandGroupResponse, err error) {
    if request == nil {
        request = NewExpandGroupRequest()
    }
    response = NewExpandGroupResponse()
    err = c.Send(request, response)
    return
}

func NewExpandNamespaceRequest() (request *ExpandNamespaceRequest) {
    request = &ExpandNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ExpandNamespace")
    return
}

func NewExpandNamespaceResponse() (response *ExpandNamespaceResponse) {
    response = &ExpandNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加机器至命名空间中
func (c *Client) ExpandNamespace(request *ExpandNamespaceRequest) (response *ExpandNamespaceResponse, err error) {
    if request == nil {
        request = NewExpandNamespaceRequest()
    }
    response = NewExpandNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewFindContainerGroupRequest() (request *FindContainerGroupRequest) {
    request = &FindContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "FindContainerGroup")
    return
}

func NewFindContainerGroupResponse() (response *FindContainerGroupResponse) {
    response = &FindContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FindContainerGroup
func (c *Client) FindContainerGroup(request *FindContainerGroupRequest) (response *FindContainerGroupResponse, err error) {
    if request == nil {
        request = NewFindContainerGroupRequest()
    }
    response = NewFindContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewFindContainerGroupsRequest() (request *FindContainerGroupsRequest) {
    request = &FindContainerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "FindContainerGroups")
    return
}

func NewFindContainerGroupsResponse() (response *FindContainerGroupsResponse) {
    response = &FindContainerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 容器部署组列表-独立菜单
func (c *Client) FindContainerGroups(request *FindContainerGroupsRequest) (response *FindContainerGroupsResponse, err error) {
    if request == nil {
        request = NewFindContainerGroupsRequest()
    }
    response = NewFindContainerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewFindDeployModuleParamsRequest() (request *FindDeployModuleParamsRequest) {
    request = &FindDeployModuleParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "FindDeployModuleParams")
    return
}

func NewFindDeployModuleParamsResponse() (response *FindDeployModuleParamsResponse) {
    response = &FindDeployModuleParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询参数配置
func (c *Client) FindDeployModuleParams(request *FindDeployModuleParamsRequest) (response *FindDeployModuleParamsResponse, err error) {
    if request == nil {
        request = NewFindDeployModuleParamsRequest()
    }
    response = NewFindDeployModuleParamsResponse()
    err = c.Send(request, response)
    return
}

func NewFindMonitorObjectRequest() (request *FindMonitorObjectRequest) {
    request = &FindMonitorObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "FindMonitorObject")
    return
}

func NewFindMonitorObjectResponse() (response *FindMonitorObjectResponse) {
    response = &FindMonitorObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FindMonitorObject
func (c *Client) FindMonitorObject(request *FindMonitorObjectRequest) (response *FindMonitorObjectResponse, err error) {
    if request == nil {
        request = NewFindMonitorObjectRequest()
    }
    response = NewFindMonitorObjectResponse()
    err = c.Send(request, response)
    return
}

func NewFindServiceMonitorObjectRequest() (request *FindServiceMonitorObjectRequest) {
    request = &FindServiceMonitorObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "FindServiceMonitorObject")
    return
}

func NewFindServiceMonitorObjectResponse() (response *FindServiceMonitorObjectResponse) {
    response = &FindServiceMonitorObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FindServiceMonitorObject
func (c *Client) FindServiceMonitorObject(request *FindServiceMonitorObjectRequest) (response *FindServiceMonitorObjectResponse, err error) {
    if request == nil {
        request = NewFindServiceMonitorObjectRequest()
    }
    response = NewFindServiceMonitorObjectResponse()
    err = c.Send(request, response)
    return
}

func NewGetApmEsAuthInfoRequest() (request *GetApmEsAuthInfoRequest) {
    request = &GetApmEsAuthInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetApmEsAuthInfo")
    return
}

func NewGetApmEsAuthInfoResponse() (response *GetApmEsAuthInfoResponse) {
    response = &GetApmEsAuthInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取应用性能管理ES权限信息
func (c *Client) GetApmEsAuthInfo(request *GetApmEsAuthInfoRequest) (response *GetApmEsAuthInfoResponse, err error) {
    if request == nil {
        request = NewGetApmEsAuthInfoRequest()
    }
    response = NewGetApmEsAuthInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterLimitResourceRequest() (request *GetClusterLimitResourceRequest) {
    request = &GetClusterLimitResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetClusterLimitResource")
    return
}

func NewGetClusterLimitResourceResponse() (response *GetClusterLimitResourceResponse) {
    response = &GetClusterLimitResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群资源
func (c *Client) GetClusterLimitResource(request *GetClusterLimitResourceRequest) (response *GetClusterLimitResourceResponse, err error) {
    if request == nil {
        request = NewGetClusterLimitResourceRequest()
    }
    response = NewGetClusterLimitResourceResponse()
    err = c.Send(request, response)
    return
}

func NewGetConfigpropsRequest() (request *GetConfigpropsRequest) {
    request = &GetConfigpropsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetConfigprops")
    return
}

func NewGetConfigpropsResponse() (response *GetConfigpropsResponse) {
    response = &GetConfigpropsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetConfigprops
func (c *Client) GetConfigprops(request *GetConfigpropsRequest) (response *GetConfigpropsResponse, err error) {
    if request == nil {
        request = NewGetConfigpropsRequest()
    }
    response = NewGetConfigpropsResponse()
    err = c.Send(request, response)
    return
}

func NewGetContainGroupDeployInfoRequest() (request *GetContainGroupDeployInfoRequest) {
    request = &GetContainGroupDeployInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetContainGroupDeployInfo")
    return
}

func NewGetContainGroupDeployInfoResponse() (response *GetContainGroupDeployInfoResponse) {
    response = &GetContainGroupDeployInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取部署组详情
func (c *Client) GetContainGroupDeployInfo(request *GetContainGroupDeployInfoRequest) (response *GetContainGroupDeployInfoResponse, err error) {
    if request == nil {
        request = NewGetContainGroupDeployInfoRequest()
    }
    response = NewGetContainGroupDeployInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetContainGroupOtherRequest() (request *GetContainGroupOtherRequest) {
    request = &GetContainGroupOtherRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetContainGroupOther")
    return
}

func NewGetContainGroupOtherResponse() (response *GetContainGroupOtherResponse) {
    response = &GetContainGroupOtherResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取部署组其他字段-用于前端并发调用
func (c *Client) GetContainGroupOther(request *GetContainGroupOtherRequest) (response *GetContainGroupOtherResponse, err error) {
    if request == nil {
        request = NewGetContainGroupOtherRequest()
    }
    response = NewGetContainGroupOtherResponse()
    err = c.Send(request, response)
    return
}

func NewGetDefaultValueRequest() (request *GetDefaultValueRequest) {
    request = &GetDefaultValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetDefaultValue")
    return
}

func NewGetDefaultValueResponse() (response *GetDefaultValueResponse) {
    response = &GetDefaultValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据参数获取默认值 
func (c *Client) GetDefaultValue(request *GetDefaultValueRequest) (response *GetDefaultValueResponse, err error) {
    if request == nil {
        request = NewGetDefaultValueRequest()
    }
    response = NewGetDefaultValueResponse()
    err = c.Send(request, response)
    return
}

func NewGetDockerForUseRequest() (request *GetDockerForUseRequest) {
    request = &GetDockerForUseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetDockerForUse")
    return
}

func NewGetDockerForUseResponse() (response *GetDockerForUseResponse) {
    response = &GetDockerForUseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDockerForUse
func (c *Client) GetDockerForUse(request *GetDockerForUseRequest) (response *GetDockerForUseResponse, err error) {
    if request == nil {
        request = NewGetDockerForUseRequest()
    }
    response = NewGetDockerForUseResponse()
    err = c.Send(request, response)
    return
}

func NewGetDownloadInfoRequest() (request *GetDownloadInfoRequest) {
    request = &GetDownloadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetDownloadInfo")
    return
}

func NewGetDownloadInfoResponse() (response *GetDownloadInfoResponse) {
    response = &GetDownloadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取下载信息
func (c *Client) GetDownloadInfo(request *GetDownloadInfoRequest) (response *GetDownloadInfoResponse, err error) {
    if request == nil {
        request = NewGetDownloadInfoRequest()
    }
    response = NewGetDownloadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetDumpRequest() (request *GetDumpRequest) {
    request = &GetDumpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetDump")
    return
}

func NewGetDumpResponse() (response *GetDumpResponse) {
    response = &GetDumpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDump
func (c *Client) GetDump(request *GetDumpRequest) (response *GetDumpResponse, err error) {
    if request == nil {
        request = NewGetDumpRequest()
    }
    response = NewGetDumpResponse()
    err = c.Send(request, response)
    return
}

func NewGetEnvRequest() (request *GetEnvRequest) {
    request = &GetEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetEnv")
    return
}

func NewGetEnvResponse() (response *GetEnvResponse) {
    response = &GetEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetEnv
func (c *Client) GetEnv(request *GetEnvRequest) (response *GetEnvResponse, err error) {
    if request == nil {
        request = NewGetEnvRequest()
    }
    response = NewGetEnvResponse()
    err = c.Send(request, response)
    return
}

func NewGetMetricsRequest() (request *GetMetricsRequest) {
    request = &GetMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetMetrics")
    return
}

func NewGetMetricsResponse() (response *GetMetricsResponse) {
    response = &GetMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取监控端点数据
func (c *Client) GetMetrics(request *GetMetricsRequest) (response *GetMetricsResponse, err error) {
    if request == nil {
        request = NewGetMetricsRequest()
    }
    response = NewGetMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewGetOssTopologyGraphRequest() (request *GetOssTopologyGraphRequest) {
    request = &GetOssTopologyGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetOssTopologyGraph")
    return
}

func NewGetOssTopologyGraphResponse() (response *GetOssTopologyGraphResponse) {
    response = &GetOssTopologyGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询依赖拓扑图
func (c *Client) GetOssTopologyGraph(request *GetOssTopologyGraphRequest) (response *GetOssTopologyGraphResponse, err error) {
    if request == nil {
        request = NewGetOssTopologyGraphRequest()
    }
    response = NewGetOssTopologyGraphResponse()
    err = c.Send(request, response)
    return
}

func NewGetOssTraceInterfacesRequest() (request *GetOssTraceInterfacesRequest) {
    request = &GetOssTraceInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetOssTraceInterfaces")
    return
}

func NewGetOssTraceInterfacesResponse() (response *GetOssTraceInterfacesResponse) {
    response = &GetOssTraceInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用链接口列表
func (c *Client) GetOssTraceInterfaces(request *GetOssTraceInterfacesRequest) (response *GetOssTraceInterfacesResponse, err error) {
    if request == nil {
        request = NewGetOssTraceInterfacesRequest()
    }
    response = NewGetOssTraceInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewGetOssTraceServicesRequest() (request *GetOssTraceServicesRequest) {
    request = &GetOssTraceServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetOssTraceServices")
    return
}

func NewGetOssTraceServicesResponse() (response *GetOssTraceServicesResponse) {
    response = &GetOssTraceServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用链服务列表
func (c *Client) GetOssTraceServices(request *GetOssTraceServicesRequest) (response *GetOssTraceServicesResponse, err error) {
    if request == nil {
        request = NewGetOssTraceServicesRequest()
    }
    response = NewGetOssTraceServicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetOssTraceSpansRequest() (request *GetOssTraceSpansRequest) {
    request = &GetOssTraceSpansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetOssTraceSpans")
    return
}

func NewGetOssTraceSpansResponse() (response *GetOssTraceSpansResponse) {
    response = &GetOssTraceSpansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用TraceId查询调用链Span
func (c *Client) GetOssTraceSpans(request *GetOssTraceSpansRequest) (response *GetOssTraceSpansResponse, err error) {
    if request == nil {
        request = NewGetOssTraceSpansRequest()
    }
    response = NewGetOssTraceSpansResponse()
    err = c.Send(request, response)
    return
}

func NewGetPkgInfoRequest() (request *GetPkgInfoRequest) {
    request = &GetPkgInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetPkgInfo")
    return
}

func NewGetPkgInfoResponse() (response *GetPkgInfoResponse) {
    response = &GetPkgInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询包详情
func (c *Client) GetPkgInfo(request *GetPkgInfoRequest) (response *GetPkgInfoResponse, err error) {
    if request == nil {
        request = NewGetPkgInfoRequest()
    }
    response = NewGetPkgInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceStatisticsRequest() (request *GetServiceStatisticsRequest) {
    request = &GetServiceStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetServiceStatistics")
    return
}

func NewGetServiceStatisticsResponse() (response *GetServiceStatisticsResponse) {
    response = &GetServiceStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务运行状态
func (c *Client) GetServiceStatistics(request *GetServiceStatisticsRequest) (response *GetServiceStatisticsResponse, err error) {
    if request == nil {
        request = NewGetServiceStatisticsRequest()
    }
    response = NewGetServiceStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopAvgTimeCostInterfacesRequest() (request *GetTopAvgTimeCostInterfacesRequest) {
    request = &GetTopAvgTimeCostInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTopAvgTimeCostInterfaces")
    return
}

func NewGetTopAvgTimeCostInterfacesResponse() (response *GetTopAvgTimeCostInterfacesResponse) {
    response = &GetTopAvgTimeCostInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TopN请求平均耗时接口列表
func (c *Client) GetTopAvgTimeCostInterfaces(request *GetTopAvgTimeCostInterfacesRequest) (response *GetTopAvgTimeCostInterfacesResponse, err error) {
    if request == nil {
        request = NewGetTopAvgTimeCostInterfacesRequest()
    }
    response = NewGetTopAvgTimeCostInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopAvgTimeCostServicesRequest() (request *GetTopAvgTimeCostServicesRequest) {
    request = &GetTopAvgTimeCostServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTopAvgTimeCostServices")
    return
}

func NewGetTopAvgTimeCostServicesResponse() (response *GetTopAvgTimeCostServicesResponse) {
    response = &GetTopAvgTimeCostServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TopN请求平均耗时服务列表
func (c *Client) GetTopAvgTimeCostServices(request *GetTopAvgTimeCostServicesRequest) (response *GetTopAvgTimeCostServicesResponse, err error) {
    if request == nil {
        request = NewGetTopAvgTimeCostServicesRequest()
    }
    response = NewGetTopAvgTimeCostServicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopFailureRateInterfacesRequest() (request *GetTopFailureRateInterfacesRequest) {
    request = &GetTopFailureRateInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTopFailureRateInterfaces")
    return
}

func NewGetTopFailureRateInterfacesResponse() (response *GetTopFailureRateInterfacesResponse) {
    response = &GetTopFailureRateInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TopN请求失败率接口列表
func (c *Client) GetTopFailureRateInterfaces(request *GetTopFailureRateInterfacesRequest) (response *GetTopFailureRateInterfacesResponse, err error) {
    if request == nil {
        request = NewGetTopFailureRateInterfacesRequest()
    }
    response = NewGetTopFailureRateInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopFailureRateServicesRequest() (request *GetTopFailureRateServicesRequest) {
    request = &GetTopFailureRateServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTopFailureRateServices")
    return
}

func NewGetTopFailureRateServicesResponse() (response *GetTopFailureRateServicesResponse) {
    response = &GetTopFailureRateServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TopN请求失败率服务列表
func (c *Client) GetTopFailureRateServices(request *GetTopFailureRateServicesRequest) (response *GetTopFailureRateServicesResponse, err error) {
    if request == nil {
        request = NewGetTopFailureRateServicesRequest()
    }
    response = NewGetTopFailureRateServicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopReqAmountInterfacesRequest() (request *GetTopReqAmountInterfacesRequest) {
    request = &GetTopReqAmountInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTopReqAmountInterfaces")
    return
}

func NewGetTopReqAmountInterfacesResponse() (response *GetTopReqAmountInterfacesResponse) {
    response = &GetTopReqAmountInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TopN请求量接口列表
func (c *Client) GetTopReqAmountInterfaces(request *GetTopReqAmountInterfacesRequest) (response *GetTopReqAmountInterfacesResponse, err error) {
    if request == nil {
        request = NewGetTopReqAmountInterfacesRequest()
    }
    response = NewGetTopReqAmountInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopReqAmountServicesRequest() (request *GetTopReqAmountServicesRequest) {
    request = &GetTopReqAmountServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTopReqAmountServices")
    return
}

func NewGetTopReqAmountServicesResponse() (response *GetTopReqAmountServicesResponse) {
    response = &GetTopReqAmountServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TopN请求量服务列表
func (c *Client) GetTopReqAmountServices(request *GetTopReqAmountServicesRequest) (response *GetTopReqAmountServicesResponse, err error) {
    if request == nil {
        request = NewGetTopReqAmountServicesRequest()
    }
    response = NewGetTopReqAmountServicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopologyGraphRequest() (request *GetTopologyGraphRequest) {
    request = &GetTopologyGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTopologyGraph")
    return
}

func NewGetTopologyGraphResponse() (response *GetTopologyGraphResponse) {
    response = &GetTopologyGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询依赖拓扑图
func (c *Client) GetTopologyGraph(request *GetTopologyGraphRequest) (response *GetTopologyGraphResponse, err error) {
    if request == nil {
        request = NewGetTopologyGraphRequest()
    }
    response = NewGetTopologyGraphResponse()
    err = c.Send(request, response)
    return
}

func NewGetTraceRequest() (request *GetTraceRequest) {
    request = &GetTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTrace")
    return
}

func NewGetTraceResponse() (response *GetTraceResponse) {
    response = &GetTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTrace
func (c *Client) GetTrace(request *GetTraceRequest) (response *GetTraceResponse, err error) {
    if request == nil {
        request = NewGetTraceRequest()
    }
    response = NewGetTraceResponse()
    err = c.Send(request, response)
    return
}

func NewGetTraceInterfacesRequest() (request *GetTraceInterfacesRequest) {
    request = &GetTraceInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTraceInterfaces")
    return
}

func NewGetTraceInterfacesResponse() (response *GetTraceInterfacesResponse) {
    response = &GetTraceInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用链接口列表
func (c *Client) GetTraceInterfaces(request *GetTraceInterfacesRequest) (response *GetTraceInterfacesResponse, err error) {
    if request == nil {
        request = NewGetTraceInterfacesRequest()
    }
    response = NewGetTraceInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTraceServicesRequest() (request *GetTraceServicesRequest) {
    request = &GetTraceServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTraceServices")
    return
}

func NewGetTraceServicesResponse() (response *GetTraceServicesResponse) {
    response = &GetTraceServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用链服务列表
func (c *Client) GetTraceServices(request *GetTraceServicesRequest) (response *GetTraceServicesResponse, err error) {
    if request == nil {
        request = NewGetTraceServicesRequest()
    }
    response = NewGetTraceServicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTraceSpansRequest() (request *GetTraceSpansRequest) {
    request = &GetTraceSpansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetTraceSpans")
    return
}

func NewGetTraceSpansResponse() (response *GetTraceSpansResponse) {
    response = &GetTraceSpansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用链Span
func (c *Client) GetTraceSpans(request *GetTraceSpansRequest) (response *GetTraceSpansResponse, err error) {
    if request == nil {
        request = NewGetTraceSpansRequest()
    }
    response = NewGetTraceSpansResponse()
    err = c.Send(request, response)
    return
}

func NewGetUploadInfoRequest() (request *GetUploadInfoRequest) {
    request = &GetUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "GetUploadInfo")
    return
}

func NewGetUploadInfoResponse() (response *GetUploadInfoResponse) {
    response = &GetUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取上传信息
func (c *Client) GetUploadInfo(request *GetUploadInfoRequest) (response *GetUploadInfoResponse, err error) {
    if request == nil {
        request = NewGetUploadInfoRequest()
    }
    response = NewGetUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewImageDeleteTagRequest() (request *ImageDeleteTagRequest) {
    request = &ImageDeleteTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ImageDeleteTag")
    return
}

func NewImageDeleteTagResponse() (response *ImageDeleteTagResponse) {
    response = &ImageDeleteTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除镜像版本
func (c *Client) ImageDeleteTag(request *ImageDeleteTagRequest) (response *ImageDeleteTagResponse, err error) {
    if request == nil {
        request = NewImageDeleteTagRequest()
    }
    response = NewImageDeleteTagResponse()
    err = c.Send(request, response)
    return
}

func NewImageGetRepositoryListRequest() (request *ImageGetRepositoryListRequest) {
    request = &ImageGetRepositoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ImageGetRepositoryList")
    return
}

func NewImageGetRepositoryListResponse() (response *ImageGetRepositoryListResponse) {
    response = &ImageGetRepositoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 镜像仓库列表
func (c *Client) ImageGetRepositoryList(request *ImageGetRepositoryListRequest) (response *ImageGetRepositoryListResponse, err error) {
    if request == nil {
        request = NewImageGetRepositoryListRequest()
    }
    response = NewImageGetRepositoryListResponse()
    err = c.Send(request, response)
    return
}

func NewImageGetTagListRequest() (request *ImageGetTagListRequest) {
    request = &ImageGetTagListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ImageGetTagList")
    return
}

func NewImageGetTagListResponse() (response *ImageGetTagListResponse) {
    response = &ImageGetTagListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 标签页镜像详细信息+部署应用选择镜像 时使用该接口
func (c *Client) ImageGetTagList(request *ImageGetTagListRequest) (response *ImageGetTagListResponse, err error) {
    if request == nil {
        request = NewImageGetTagListRequest()
    }
    response = NewImageGetTagListResponse()
    err = c.Send(request, response)
    return
}

func NewImageRegisterRequest() (request *ImageRegisterRequest) {
    request = &ImageRegisterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ImageRegister")
    return
}

func NewImageRegisterResponse() (response *ImageRegisterResponse) {
    response = &ImageRegisterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 镜像用户注册
func (c *Client) ImageRegister(request *ImageRegisterRequest) (response *ImageRegisterResponse, err error) {
    if request == nil {
        request = NewImageRegisterRequest()
    }
    response = NewImageRegisterResponse()
    err = c.Send(request, response)
    return
}

func NewImageUserIsExistsRequest() (request *ImageUserIsExistsRequest) {
    request = &ImageUserIsExistsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ImageUserIsExists")
    return
}

func NewImageUserIsExistsResponse() (response *ImageUserIsExistsResponse) {
    response = &ImageUserIsExistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 首次点击”镜像标签页”时触发
func (c *Client) ImageUserIsExists(request *ImageUserIsExistsRequest) (response *ImageUserIsExistsResponse, err error) {
    if request == nil {
        request = NewImageUserIsExistsRequest()
    }
    response = NewImageUserIsExistsResponse()
    err = c.Send(request, response)
    return
}

func NewInitializeApmRequest() (request *InitializeApmRequest) {
    request = &InitializeApmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "InitializeApm")
    return
}

func NewInitializeApmResponse() (response *InitializeApmResponse) {
    response = &InitializeApmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 初始化应用性能管理功能，包括默认调用链索引创建，ES权限用户创建和授权等
func (c *Client) InitializeApm(request *InitializeApmRequest) (response *InitializeApmResponse, err error) {
    if request == nil {
        request = NewInitializeApmRequest()
    }
    response = NewInitializeApmResponse()
    err = c.Send(request, response)
    return
}

func NewInstallAgentScriptsRequest() (request *InstallAgentScriptsRequest) {
    request = &InstallAgentScriptsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "InstallAgentScripts")
    return
}

func NewInstallAgentScriptsResponse() (response *InstallAgentScriptsResponse) {
    response = &InstallAgentScriptsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 手动安装agent脚本
func (c *Client) InstallAgentScripts(request *InstallAgentScriptsRequest) (response *InstallAgentScriptsResponse, err error) {
    if request == nil {
        request = NewInstallAgentScriptsRequest()
    }
    response = NewInstallAgentScriptsResponse()
    err = c.Send(request, response)
    return
}

func NewIssueLicenseRequest() (request *IssueLicenseRequest) {
    request = &IssueLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "IssueLicense")
    return
}

func NewIssueLicenseResponse() (response *IssueLicenseResponse) {
    response = &IssueLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 授予许可证
func (c *Client) IssueLicense(request *IssueLicenseRequest) (response *IssueLicenseResponse, err error) {
    if request == nil {
        request = NewIssueLicenseRequest()
    }
    response = NewIssueLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewListAlarmPoliciesRequest() (request *ListAlarmPoliciesRequest) {
    request = &ListAlarmPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListAlarmPolicies")
    return
}

func NewListAlarmPoliciesResponse() (response *ListAlarmPoliciesResponse) {
    response = &ListAlarmPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警策略列表
func (c *Client) ListAlarmPolicies(request *ListAlarmPoliciesRequest) (response *ListAlarmPoliciesResponse, err error) {
    if request == nil {
        request = NewListAlarmPoliciesRequest()
    }
    response = NewListAlarmPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListAlarmReceiversRequest() (request *ListAlarmReceiversRequest) {
    request = &ListAlarmReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListAlarmReceivers")
    return
}

func NewListAlarmReceiversResponse() (response *ListAlarmReceiversResponse) {
    response = &ListAlarmReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警接收者列表
func (c *Client) ListAlarmReceivers(request *ListAlarmReceiversRequest) (response *ListAlarmReceiversResponse, err error) {
    if request == nil {
        request = NewListAlarmReceiversRequest()
    }
    response = NewListAlarmReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewListAppRequest() (request *ListAppRequest) {
    request = &ListAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListApp")
    return
}

func NewListAppResponse() (response *ListAppResponse) {
    response = &ListAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户的应用信息
func (c *Client) ListApp(request *ListAppRequest) (response *ListAppResponse, err error) {
    if request == nil {
        request = NewListAppRequest()
    }
    response = NewListAppResponse()
    err = c.Send(request, response)
    return
}

func NewListAppPkgRequest() (request *ListAppPkgRequest) {
    request = &ListAppPkgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListAppPkg")
    return
}

func NewListAppPkgResponse() (response *ListAppPkgResponse) {
    response = &ListAppPkgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户的应用信息
func (c *Client) ListAppPkg(request *ListAppPkgRequest) (response *ListAppPkgResponse, err error) {
    if request == nil {
        request = NewListAppPkgRequest()
    }
    response = NewListAppPkgResponse()
    err = c.Send(request, response)
    return
}

func NewListApplicationServersRequest() (request *ListApplicationServersRequest) {
    request = &ListApplicationServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListApplicationServers")
    return
}

func NewListApplicationServersResponse() (response *ListApplicationServersResponse) {
    response = &ListApplicationServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 应用服务器列表
func (c *Client) ListApplicationServers(request *ListApplicationServersRequest) (response *ListApplicationServersResponse, err error) {
    if request == nil {
        request = NewListApplicationServersRequest()
    }
    response = NewListApplicationServersResponse()
    err = c.Send(request, response)
    return
}

func NewListCloudMicroServiceFindPagedListRequest() (request *ListCloudMicroServiceFindPagedListRequest) {
    request = &ListCloudMicroServiceFindPagedListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListCloudMicroServiceFindPagedList")
    return
}

func NewListCloudMicroServiceFindPagedListResponse() (response *ListCloudMicroServiceFindPagedListResponse) {
    response = &ListCloudMicroServiceFindPagedListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloudMonitorMicroserviceResult
func (c *Client) ListCloudMicroServiceFindPagedList(request *ListCloudMicroServiceFindPagedListRequest) (response *ListCloudMicroServiceFindPagedListResponse, err error) {
    if request == nil {
        request = NewListCloudMicroServiceFindPagedListRequest()
    }
    response = NewListCloudMicroServiceFindPagedListResponse()
    err = c.Send(request, response)
    return
}

func NewListColudMonitorStatisticsPolicyRequest() (request *ListColudMonitorStatisticsPolicyRequest) {
    request = &ListColudMonitorStatisticsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListColudMonitorStatisticsPolicy")
    return
}

func NewListColudMonitorStatisticsPolicyResponse() (response *ListColudMonitorStatisticsPolicyResponse) {
    response = &ListColudMonitorStatisticsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListColudMonitorStatisticsPolicy
func (c *Client) ListColudMonitorStatisticsPolicy(request *ListColudMonitorStatisticsPolicyRequest) (response *ListColudMonitorStatisticsPolicyResponse, err error) {
    if request == nil {
        request = NewListColudMonitorStatisticsPolicyRequest()
    }
    response = NewListColudMonitorStatisticsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewListContainGroupRequest() (request *ListContainGroupRequest) {
    request = &ListContainGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListContainGroup")
    return
}

func NewListContainGroupResponse() (response *ListContainGroupResponse) {
    response = &ListContainGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 容器部署组列表
func (c *Client) ListContainGroup(request *ListContainGroupRequest) (response *ListContainGroupResponse, err error) {
    if request == nil {
        request = NewListContainGroupRequest()
    }
    response = NewListContainGroupResponse()
    err = c.Send(request, response)
    return
}

func NewListContainerTaskRequest() (request *ListContainerTaskRequest) {
    request = &ListContainerTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListContainerTask")
    return
}

func NewListContainerTaskResponse() (response *ListContainerTaskResponse) {
    response = &ListContainerTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 变更记录任务列表
func (c *Client) ListContainerTask(request *ListContainerTaskRequest) (response *ListContainerTaskResponse, err error) {
    if request == nil {
        request = NewListContainerTaskRequest()
    }
    response = NewListContainerTaskResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupPodRequest() (request *ListGroupPodRequest) {
    request = &ListGroupPodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListGroupPod")
    return
}

func NewListGroupPodResponse() (response *ListGroupPodResponse) {
    response = &ListGroupPodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取容器部署组实例列表
func (c *Client) ListGroupPod(request *ListGroupPodRequest) (response *ListGroupPodResponse, err error) {
    if request == nil {
        request = NewListGroupPodRequest()
    }
    response = NewListGroupPodResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupsByScalableRuleIdRequest() (request *ListGroupsByScalableRuleIdRequest) {
    request = &ListGroupsByScalableRuleIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListGroupsByScalableRuleId")
    return
}

func NewListGroupsByScalableRuleIdResponse() (response *ListGroupsByScalableRuleIdResponse) {
    response = &ListGroupsByScalableRuleIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取弹性扩缩容应用和部署组 
func (c *Client) ListGroupsByScalableRuleId(request *ListGroupsByScalableRuleIdRequest) (response *ListGroupsByScalableRuleIdResponse, err error) {
    if request == nil {
        request = NewListGroupsByScalableRuleIdRequest()
    }
    response = NewListGroupsByScalableRuleIdResponse()
    err = c.Send(request, response)
    return
}

func NewListMachinesRequest() (request *ListMachinesRequest) {
    request = &ListMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListMachines")
    return
}

func NewListMachinesResponse() (response *ListMachinesResponse) {
    response = &ListMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询机器列表
func (c *Client) ListMachines(request *ListMachinesRequest) (response *ListMachinesResponse, err error) {
    if request == nil {
        request = NewListMachinesRequest()
    }
    response = NewListMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewListMonitorStatisticsPolicyRequest() (request *ListMonitorStatisticsPolicyRequest) {
    request = &ListMonitorStatisticsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListMonitorStatisticsPolicy")
    return
}

func NewListMonitorStatisticsPolicyResponse() (response *ListMonitorStatisticsPolicyResponse) {
    response = &ListMonitorStatisticsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询监控统计策略列表
func (c *Client) ListMonitorStatisticsPolicy(request *ListMonitorStatisticsPolicyRequest) (response *ListMonitorStatisticsPolicyResponse, err error) {
    if request == nil {
        request = NewListMonitorStatisticsPolicyRequest()
    }
    response = NewListMonitorStatisticsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewListPkgRequest() (request *ListPkgRequest) {
    request = &ListPkgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListPkg")
    return
}

func NewListPkgResponse() (response *ListPkgResponse) {
    response = &ListPkgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 包列表
func (c *Client) ListPkg(request *ListPkgRequest) (response *ListPkgResponse, err error) {
    if request == nil {
        request = NewListPkgRequest()
    }
    response = NewListPkgResponse()
    err = c.Send(request, response)
    return
}

func NewListScalableRuleRequest() (request *ListScalableRuleRequest) {
    request = &ListScalableRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListScalableRule")
    return
}

func NewListScalableRuleResponse() (response *ListScalableRuleResponse) {
    response = &ListScalableRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// list弹性扩缩容规则 
func (c *Client) ListScalableRule(request *ListScalableRuleRequest) (response *ListScalableRuleResponse, err error) {
    if request == nil {
        request = NewListScalableRuleRequest()
    }
    response = NewListScalableRuleResponse()
    err = c.Send(request, response)
    return
}

func NewListScalableTasksRequest() (request *ListScalableTasksRequest) {
    request = &ListScalableTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListScalableTasks")
    return
}

func NewListScalableTasksResponse() (response *ListScalableTasksResponse) {
    response = &ListScalableTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取弹性扩缩容任务
func (c *Client) ListScalableTasks(request *ListScalableTasksRequest) (response *ListScalableTasksResponse, err error) {
    if request == nil {
        request = NewListScalableTasksRequest()
    }
    response = NewListScalableTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListTsfModulesDetailRequest() (request *ListTsfModulesDetailRequest) {
    request = &ListTsfModulesDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ListTsfModulesDetail")
    return
}

func NewListTsfModulesDetailResponse() (response *ListTsfModulesDetailResponse) {
    response = &ListTsfModulesDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 模块详情列表
func (c *Client) ListTsfModulesDetail(request *ListTsfModulesDetailRequest) (response *ListTsfModulesDetailResponse, err error) {
    if request == nil {
        request = NewListTsfModulesDetailRequest()
    }
    response = NewListTsfModulesDetailResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyRequest() (request *ModifyAlarmPolicyRequest) {
    request = &ModifyAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyAlarmPolicy")
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

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyApplication")
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改应用
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuthorizationRequest() (request *ModifyAuthorizationRequest) {
    request = &ModifyAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyAuthorization")
    return
}

func NewModifyAuthorizationResponse() (response *ModifyAuthorizationResponse) {
    response = &ModifyAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新微服务权限规则
func (c *Client) ModifyAuthorization(request *ModifyAuthorizationRequest) (response *ModifyAuthorizationResponse, err error) {
    if request == nil {
        request = NewModifyAuthorizationRequest()
    }
    response = NewModifyAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuthorizationInfoRequest() (request *ModifyAuthorizationInfoRequest) {
    request = &ModifyAuthorizationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyAuthorizationInfo")
    return
}

func NewModifyAuthorizationInfoResponse() (response *ModifyAuthorizationInfoResponse) {
    response = &ModifyAuthorizationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置服务权限配置
func (c *Client) ModifyAuthorizationInfo(request *ModifyAuthorizationInfoRequest) (response *ModifyAuthorizationInfoResponse, err error) {
    if request == nil {
        request = NewModifyAuthorizationInfoRequest()
    }
    response = NewModifyAuthorizationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuthorizationTypeRequest() (request *ModifyAuthorizationTypeRequest) {
    request = &ModifyAuthorizationTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyAuthorizationType")
    return
}

func NewModifyAuthorizationTypeResponse() (response *ModifyAuthorizationTypeResponse) {
    response = &ModifyAuthorizationTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改鉴权类型
func (c *Client) ModifyAuthorizationType(request *ModifyAuthorizationTypeRequest) (response *ModifyAuthorizationTypeResponse, err error) {
    if request == nil {
        request = NewModifyAuthorizationTypeRequest()
    }
    response = NewModifyAuthorizationTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterRequest() (request *ModifyClusterRequest) {
    request = &ModifyClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyCluster")
    return
}

func NewModifyClusterResponse() (response *ModifyClusterResponse) {
    response = &ModifyClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改集群信息
func (c *Client) ModifyCluster(request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
    if request == nil {
        request = NewModifyClusterRequest()
    }
    response = NewModifyClusterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerGroupRequest() (request *ModifyContainerGroupRequest) {
    request = &ModifyContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerGroup")
    return
}

func NewModifyContainerGroupResponse() (response *ModifyContainerGroupResponse) {
    response = &ModifyContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改容器部署组
func (c *Client) ModifyContainerGroup(request *ModifyContainerGroupRequest) (response *ModifyContainerGroupResponse, err error) {
    if request == nil {
        request = NewModifyContainerGroupRequest()
    }
    response = NewModifyContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerReplicasRequest() (request *ModifyContainerReplicasRequest) {
    request = &ModifyContainerReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerReplicas")
    return
}

func NewModifyContainerReplicasResponse() (response *ModifyContainerReplicasResponse) {
    response = &ModifyContainerReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改容器部署组实例数
func (c *Client) ModifyContainerReplicas(request *ModifyContainerReplicasRequest) (response *ModifyContainerReplicasResponse, err error) {
    if request == nil {
        request = NewModifyContainerReplicasRequest()
    }
    response = NewModifyContainerReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
    request = &ModifyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyGroup")
    return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
    response = &ModifyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新部署组信息
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    response = NewModifyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNamespaceRequest() (request *ModifyInstanceNamespaceRequest) {
    request = &ModifyInstanceNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyInstanceNamespace")
    return
}

func NewModifyInstanceNamespaceResponse() (response *ModifyInstanceNamespaceResponse) {
    response = &ModifyInstanceNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改机器命名空间
func (c *Client) ModifyInstanceNamespace(request *ModifyInstanceNamespaceRequest) (response *ModifyInstanceNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNamespaceRequest()
    }
    response = NewModifyInstanceNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMachinesRequest() (request *ModifyMachinesRequest) {
    request = &ModifyMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyMachines")
    return
}

func NewModifyMachinesResponse() (response *ModifyMachinesResponse) {
    response = &ModifyMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改机器列表
func (c *Client) ModifyMachines(request *ModifyMachinesRequest) (response *ModifyMachinesResponse, err error) {
    if request == nil {
        request = NewModifyMachinesRequest()
    }
    response = NewModifyMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMicroserviceRequest() (request *ModifyMicroserviceRequest) {
    request = &ModifyMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyMicroservice")
    return
}

func NewModifyMicroserviceResponse() (response *ModifyMicroserviceResponse) {
    response = &ModifyMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改微服务详情
func (c *Client) ModifyMicroservice(request *ModifyMicroserviceRequest) (response *ModifyMicroserviceResponse, err error) {
    if request == nil {
        request = NewModifyMicroserviceRequest()
    }
    response = NewModifyMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleConfRequest() (request *ModifyModuleConfRequest) {
    request = &ModifyModuleConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyModuleConf")
    return
}

func NewModifyModuleConfResponse() (response *ModifyModuleConfResponse) {
    response = &ModifyModuleConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改模块配置
func (c *Client) ModifyModuleConf(request *ModifyModuleConfRequest) (response *ModifyModuleConfResponse, err error) {
    if request == nil {
        request = NewModifyModuleConfRequest()
    }
    response = NewModifyModuleConfResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMonitorStatisticsPolicyRequest() (request *ModifyMonitorStatisticsPolicyRequest) {
    request = &ModifyMonitorStatisticsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyMonitorStatisticsPolicy")
    return
}

func NewModifyMonitorStatisticsPolicyResponse() (response *ModifyMonitorStatisticsPolicyResponse) {
    response = &ModifyMonitorStatisticsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改监控统计策略
func (c *Client) ModifyMonitorStatisticsPolicy(request *ModifyMonitorStatisticsPolicyRequest) (response *ModifyMonitorStatisticsPolicyResponse, err error) {
    if request == nil {
        request = NewModifyMonitorStatisticsPolicyRequest()
    }
    response = NewModifyMonitorStatisticsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNamespaceRequest() (request *ModifyNamespaceRequest) {
    request = &ModifyNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyNamespace")
    return
}

func NewModifyNamespaceResponse() (response *ModifyNamespaceResponse) {
    response = &ModifyNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改命名空间
func (c *Client) ModifyNamespace(request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceRequest()
    }
    response = NewModifyNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNamespaceCodeRequest() (request *ModifyNamespaceCodeRequest) {
    request = &ModifyNamespaceCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyNamespaceCode")
    return
}

func NewModifyNamespaceCodeResponse() (response *ModifyNamespaceCodeResponse) {
    response = &ModifyNamespaceCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新命名空间编码
func (c *Client) ModifyNamespaceCode(request *ModifyNamespaceCodeRequest) (response *ModifyNamespaceCodeResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceCodeRequest()
    }
    response = NewModifyNamespaceCodeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProgramRequest() (request *ModifyProgramRequest) {
    request = &ModifyProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyProgram")
    return
}

func NewModifyProgramResponse() (response *ModifyProgramResponse) {
    response = &ModifyProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新数据集
func (c *Client) ModifyProgram(request *ModifyProgramRequest) (response *ModifyProgramResponse, err error) {
    if request == nil {
        request = NewModifyProgramRequest()
    }
    response = NewModifyProgramResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRegionRequest() (request *ModifyRegionRequest) {
    request = &ModifyRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyRegion")
    return
}

func NewModifyRegionResponse() (response *ModifyRegionResponse) {
    response = &ModifyRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改业务地域
func (c *Client) ModifyRegion(request *ModifyRegionRequest) (response *ModifyRegionResponse, err error) {
    if request == nil {
        request = NewModifyRegionRequest()
    }
    response = NewModifyRegionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoleRequest() (request *ModifyRoleRequest) {
    request = &ModifyRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyRole")
    return
}

func NewModifyRoleResponse() (response *ModifyRoleResponse) {
    response = &ModifyRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新角色
func (c *Client) ModifyRole(request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
    if request == nil {
        request = NewModifyRoleRequest()
    }
    response = NewModifyRoleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRouteRequest() (request *ModifyRouteRequest) {
    request = &ModifyRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyRoute")
    return
}

func NewModifyRouteResponse() (response *ModifyRouteResponse) {
    response = &ModifyRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新路由详情
func (c *Client) ModifyRoute(request *ModifyRouteRequest) (response *ModifyRouteResponse, err error) {
    if request == nil {
        request = NewModifyRouteRequest()
    }
    response = NewModifyRouteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRouteRuleRequest() (request *ModifyRouteRuleRequest) {
    request = &ModifyRouteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyRouteRule")
    return
}

func NewModifyRouteRuleResponse() (response *ModifyRouteRuleResponse) {
    response = &ModifyRouteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新服务路由规则
func (c *Client) ModifyRouteRule(request *ModifyRouteRuleRequest) (response *ModifyRouteRuleResponse, err error) {
    if request == nil {
        request = NewModifyRouteRuleRequest()
    }
    response = NewModifyRouteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScalableRuleRequest() (request *ModifyScalableRuleRequest) {
    request = &ModifyScalableRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyScalableRule")
    return
}

func NewModifyScalableRuleResponse() (response *ModifyScalableRuleResponse) {
    response = &ModifyScalableRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改弹性扩缩容规则
func (c *Client) ModifyScalableRule(request *ModifyScalableRuleRequest) (response *ModifyScalableRuleResponse, err error) {
    if request == nil {
        request = NewModifyScalableRuleRequest()
    }
    response = NewModifyScalableRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTsfZoneRequest() (request *ModifyTsfZoneRequest) {
    request = &ModifyTsfZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyTsfZone")
    return
}

func NewModifyTsfZoneResponse() (response *ModifyTsfZoneResponse) {
    response = &ModifyTsfZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改基础区域
func (c *Client) ModifyTsfZone(request *ModifyTsfZoneRequest) (response *ModifyTsfZoneResponse, err error) {
    if request == nil {
        request = NewModifyTsfZoneRequest()
    }
    response = NewModifyTsfZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUploadInfoRequest() (request *ModifyUploadInfoRequest) {
    request = &ModifyUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyUploadInfo")
    return
}

func NewModifyUploadInfoResponse() (response *ModifyUploadInfoResponse) {
    response = &ModifyUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 调用该接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
// 调用此接口完成后，才标志上传包流程结束。
func (c *Client) ModifyUploadInfo(request *ModifyUploadInfoRequest) (response *ModifyUploadInfoResponse, err error) {
    if request == nil {
        request = NewModifyUploadInfoRequest()
    }
    response = NewModifyUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneRequest() (request *ModifyZoneRequest) {
    request = &ModifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyZone")
    return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
    response = &ModifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改业务可用区
func (c *Client) ModifyZone(request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    if request == nil {
        request = NewModifyZoneRequest()
    }
    response = NewModifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewReRelateGroupToScalableRuleRequest() (request *ReRelateGroupToScalableRuleRequest) {
    request = &ReRelateGroupToScalableRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReRelateGroupToScalableRule")
    return
}

func NewReRelateGroupToScalableRuleResponse() (response *ReRelateGroupToScalableRuleResponse) {
    response = &ReRelateGroupToScalableRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重新关联弹性伸缩规则到应用
func (c *Client) ReRelateGroupToScalableRule(request *ReRelateGroupToScalableRuleRequest) (response *ReRelateGroupToScalableRuleResponse, err error) {
    if request == nil {
        request = NewReRelateGroupToScalableRuleRequest()
    }
    response = NewReRelateGroupToScalableRuleResponse()
    err = c.Send(request, response)
    return
}

func NewReassociateBusinessLogConfigRequest() (request *ReassociateBusinessLogConfigRequest) {
    request = &ReassociateBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReassociateBusinessLogConfig")
    return
}

func NewReassociateBusinessLogConfigResponse() (response *ReassociateBusinessLogConfigResponse) {
    response = &ReassociateBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重关联业务日志配置
func (c *Client) ReassociateBusinessLogConfig(request *ReassociateBusinessLogConfigRequest) (response *ReassociateBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewReassociateBusinessLogConfigRequest()
    }
    response = NewReassociateBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterImageUserRequest() (request *RegisterImageUserRequest) {
    request = &RegisterImageUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RegisterImageUser")
    return
}

func NewRegisterImageUserResponse() (response *RegisterImageUserResponse) {
    response = &RegisterImageUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 镜像用户注册 
func (c *Client) RegisterImageUser(request *RegisterImageUserRequest) (response *RegisterImageUserResponse, err error) {
    if request == nil {
        request = NewRegisterImageUserRequest()
    }
    response = NewRegisterImageUserResponse()
    err = c.Send(request, response)
    return
}

func NewRelateGroupToScalableRuleRequest() (request *RelateGroupToScalableRuleRequest) {
    request = &RelateGroupToScalableRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RelateGroupToScalableRule")
    return
}

func NewRelateGroupToScalableRuleResponse() (response *RelateGroupToScalableRuleResponse) {
    response = &RelateGroupToScalableRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 关联弹性伸缩规则到应用 
func (c *Client) RelateGroupToScalableRule(request *RelateGroupToScalableRuleRequest) (response *RelateGroupToScalableRuleResponse, err error) {
    if request == nil {
        request = NewRelateGroupToScalableRuleRequest()
    }
    response = NewRelateGroupToScalableRuleResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseApiGroupRequest() (request *ReleaseApiGroupRequest) {
    request = &ReleaseApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseApiGroup")
    return
}

func NewReleaseApiGroupResponse() (response *ReleaseApiGroupResponse) {
    response = &ReleaseApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布Api分组
func (c *Client) ReleaseApiGroup(request *ReleaseApiGroupRequest) (response *ReleaseApiGroupResponse, err error) {
    if request == nil {
        request = NewReleaseApiGroupRequest()
    }
    response = NewReleaseApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseBusinessLogConfigRequest() (request *ReleaseBusinessLogConfigRequest) {
    request = &ReleaseBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseBusinessLogConfig")
    return
}

func NewReleaseBusinessLogConfigResponse() (response *ReleaseBusinessLogConfigResponse) {
    response = &ReleaseBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布业务日志配置
func (c *Client) ReleaseBusinessLogConfig(request *ReleaseBusinessLogConfigRequest) (response *ReleaseBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewReleaseBusinessLogConfigRequest()
    }
    response = NewReleaseBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseConfigRequest() (request *ReleaseConfigRequest) {
    request = &ReleaseConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseConfig")
    return
}

func NewReleaseConfigResponse() (response *ReleaseConfigResponse) {
    response = &ReleaseConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布配置
func (c *Client) ReleaseConfig(request *ReleaseConfigRequest) (response *ReleaseConfigResponse, err error) {
    if request == nil {
        request = NewReleaseConfigRequest()
    }
    response = NewReleaseConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseFileConfigRequest() (request *ReleaseFileConfigRequest) {
    request = &ReleaseFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseFileConfig")
    return
}

func NewReleaseFileConfigResponse() (response *ReleaseFileConfigResponse) {
    response = &ReleaseFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布文件配置
func (c *Client) ReleaseFileConfig(request *ReleaseFileConfigRequest) (response *ReleaseFileConfigResponse, err error) {
    if request == nil {
        request = NewReleaseFileConfigRequest()
    }
    response = NewReleaseFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReleasePluginRequest() (request *ReleasePluginRequest) {
    request = &ReleasePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleasePlugin")
    return
}

func NewReleasePluginResponse() (response *ReleasePluginResponse) {
    response = &ReleasePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布插件
func (c *Client) ReleasePlugin(request *ReleasePluginRequest) (response *ReleasePluginResponse, err error) {
    if request == nil {
        request = NewReleasePluginRequest()
    }
    response = NewReleasePluginResponse()
    err = c.Send(request, response)
    return
}

func NewReleasePublicConfigRequest() (request *ReleasePublicConfigRequest) {
    request = &ReleasePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleasePublicConfig")
    return
}

func NewReleasePublicConfigResponse() (response *ReleasePublicConfigResponse) {
    response = &ReleasePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布公共配置
func (c *Client) ReleasePublicConfig(request *ReleasePublicConfigRequest) (response *ReleasePublicConfigResponse, err error) {
    if request == nil {
        request = NewReleasePublicConfigRequest()
    }
    response = NewReleasePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveInstanceRequest() (request *RemoveInstanceRequest) {
    request = &RemoveInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RemoveInstance")
    return
}

func NewRemoveInstanceResponse() (response *RemoveInstanceResponse) {
    response = &RemoveInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从tsf集群中移除机器节点
func (c *Client) RemoveInstance(request *RemoveInstanceRequest) (response *RemoveInstanceResponse, err error) {
    if request == nil {
        request = NewRemoveInstanceRequest()
    }
    response = NewRemoveInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveInstancesRequest() (request *RemoveInstancesRequest) {
    request = &RemoveInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RemoveInstances")
    return
}

func NewRemoveInstancesResponse() (response *RemoveInstancesResponse) {
    response = &RemoveInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从 TSF 集群中批量移除云主机节点
func (c *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    if request == nil {
        request = NewRemoveInstancesRequest()
    }
    response = NewRemoveInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRetryTransactionRequest() (request *RetryTransactionRequest) {
    request = &RetryTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RetryTransaction")
    return
}

func NewRetryTransactionResponse() (response *RetryTransactionResponse) {
    response = &RetryTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 手动触发事务超时重试
func (c *Client) RetryTransaction(request *RetryTransactionRequest) (response *RetryTransactionResponse, err error) {
    if request == nil {
        request = NewRetryTransactionRequest()
    }
    response = NewRetryTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewRevocationConfigRequest() (request *RevocationConfigRequest) {
    request = &RevocationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RevocationConfig")
    return
}

func NewRevocationConfigResponse() (response *RevocationConfigResponse) {
    response = &RevocationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 撤回已发布的配置
func (c *Client) RevocationConfig(request *RevocationConfigRequest) (response *RevocationConfigResponse, err error) {
    if request == nil {
        request = NewRevocationConfigRequest()
    }
    response = NewRevocationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRevocationPublicConfigRequest() (request *RevocationPublicConfigRequest) {
    request = &RevocationPublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RevocationPublicConfig")
    return
}

func NewRevocationPublicConfigResponse() (response *RevocationPublicConfigResponse) {
    response = &RevocationPublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 撤回已发布的公共配置
func (c *Client) RevocationPublicConfig(request *RevocationPublicConfigRequest) (response *RevocationPublicConfigResponse, err error) {
    if request == nil {
        request = NewRevocationPublicConfigRequest()
    }
    response = NewRevocationPublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeFileConfigRequest() (request *RevokeFileConfigRequest) {
    request = &RevokeFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RevokeFileConfig")
    return
}

func NewRevokeFileConfigResponse() (response *RevokeFileConfigResponse) {
    response = &RevokeFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 撤回已发布的文件配置
func (c *Client) RevokeFileConfig(request *RevokeFileConfigRequest) (response *RevokeFileConfigResponse, err error) {
    if request == nil {
        request = NewRevokeFileConfigRequest()
    }
    response = NewRevokeFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackConfigRequest() (request *RollbackConfigRequest) {
    request = &RollbackConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RollbackConfig")
    return
}

func NewRollbackConfigResponse() (response *RollbackConfigResponse) {
    response = &RollbackConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回滚配置
func (c *Client) RollbackConfig(request *RollbackConfigRequest) (response *RollbackConfigResponse, err error) {
    if request == nil {
        request = NewRollbackConfigRequest()
    }
    response = NewRollbackConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackFileConfigRequest() (request *RollbackFileConfigRequest) {
    request = &RollbackFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RollbackFileConfig")
    return
}

func NewRollbackFileConfigResponse() (response *RollbackFileConfigResponse) {
    response = &RollbackFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回滚文件配置
func (c *Client) RollbackFileConfig(request *RollbackFileConfigRequest) (response *RollbackFileConfigResponse, err error) {
    if request == nil {
        request = NewRollbackFileConfigRequest()
    }
    response = NewRollbackFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackPublicConfigRequest() (request *RollbackPublicConfigRequest) {
    request = &RollbackPublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RollbackPublicConfig")
    return
}

func NewRollbackPublicConfigResponse() (response *RollbackPublicConfigResponse) {
    response = &RollbackPublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回滚公共配置
func (c *Client) RollbackPublicConfig(request *RollbackPublicConfigRequest) (response *RollbackPublicConfigResponse, err error) {
    if request == nil {
        request = NewRollbackPublicConfigRequest()
    }
    response = NewRollbackPublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRunMsApiRequest() (request *RunMsApiRequest) {
    request = &RunMsApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RunMsApi")
    return
}

func NewRunMsApiResponse() (response *RunMsApiResponse) {
    response = &RunMsApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 调试API
func (c *Client) RunMsApi(request *RunMsApiRequest) (response *RunMsApiResponse, err error) {
    if request == nil {
        request = NewRunMsApiRequest()
    }
    response = NewRunMsApiResponse()
    err = c.Send(request, response)
    return
}

func NewSaveDeployModuleParamsRequest() (request *SaveDeployModuleParamsRequest) {
    request = &SaveDeployModuleParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SaveDeployModuleParams")
    return
}

func NewSaveDeployModuleParamsResponse() (response *SaveDeployModuleParamsResponse) {
    response = &SaveDeployModuleParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 保存参数配置
func (c *Client) SaveDeployModuleParams(request *SaveDeployModuleParamsRequest) (response *SaveDeployModuleParamsResponse, err error) {
    if request == nil {
        request = NewSaveDeployModuleParamsRequest()
    }
    response = NewSaveDeployModuleParamsResponse()
    err = c.Send(request, response)
    return
}

func NewSearchBusinessLogRequest() (request *SearchBusinessLogRequest) {
    request = &SearchBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchBusinessLog")
    return
}

func NewSearchBusinessLogResponse() (response *SearchBusinessLogResponse) {
    response = &SearchBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务日志搜索
func (c *Client) SearchBusinessLog(request *SearchBusinessLogRequest) (response *SearchBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchBusinessLogRequest()
    }
    response = NewSearchBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchOssBusinessLogRequest() (request *SearchOssBusinessLogRequest) {
    request = &SearchOssBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchOssBusinessLog")
    return
}

func NewSearchOssBusinessLogResponse() (response *SearchOssBusinessLogResponse) {
    response = &SearchOssBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务日志搜索
func (c *Client) SearchOssBusinessLog(request *SearchOssBusinessLogRequest) (response *SearchOssBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchOssBusinessLogRequest()
    }
    response = NewSearchOssBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchOssRealtimeBusinessLogRequest() (request *SearchOssRealtimeBusinessLogRequest) {
    request = &SearchOssRealtimeBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchOssRealtimeBusinessLog")
    return
}

func NewSearchOssRealtimeBusinessLogResponse() (response *SearchOssRealtimeBusinessLogResponse) {
    response = &SearchOssRealtimeBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OSS模块群业务日志实时检索
func (c *Client) SearchOssRealtimeBusinessLog(request *SearchOssRealtimeBusinessLogRequest) (response *SearchOssRealtimeBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchOssRealtimeBusinessLogRequest()
    }
    response = NewSearchOssRealtimeBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchOssSpanBusinessLogRequest() (request *SearchOssSpanBusinessLogRequest) {
    request = &SearchOssSpanBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchOssSpanBusinessLog")
    return
}

func NewSearchOssSpanBusinessLogResponse() (response *SearchOssSpanBusinessLogResponse) {
    response = &SearchOssSpanBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务日志搜索（关联调用链Span）
func (c *Client) SearchOssSpanBusinessLog(request *SearchOssSpanBusinessLogRequest) (response *SearchOssSpanBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchOssSpanBusinessLogRequest()
    }
    response = NewSearchOssSpanBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchOssStaticBusinessLogRequest() (request *SearchOssStaticBusinessLogRequest) {
    request = &SearchOssStaticBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchOssStaticBusinessLog")
    return
}

func NewSearchOssStaticBusinessLogResponse() (response *SearchOssStaticBusinessLogResponse) {
    response = &SearchOssStaticBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 搜索日志
func (c *Client) SearchOssStaticBusinessLog(request *SearchOssStaticBusinessLogRequest) (response *SearchOssStaticBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchOssStaticBusinessLogRequest()
    }
    response = NewSearchOssStaticBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchOssSurroundBusinessLogRequest() (request *SearchOssSurroundBusinessLogRequest) {
    request = &SearchOssSurroundBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchOssSurroundBusinessLog")
    return
}

func NewSearchOssSurroundBusinessLogResponse() (response *SearchOssSurroundBusinessLogResponse) {
    response = &SearchOssSurroundBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务日志上下文检索
func (c *Client) SearchOssSurroundBusinessLog(request *SearchOssSurroundBusinessLogRequest) (response *SearchOssSurroundBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchOssSurroundBusinessLogRequest()
    }
    response = NewSearchOssSurroundBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchOssTraceRequest() (request *SearchOssTraceRequest) {
    request = &SearchOssTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchOssTrace")
    return
}

func NewSearchOssTraceResponse() (response *SearchOssTraceResponse) {
    response = &SearchOssTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用链
func (c *Client) SearchOssTrace(request *SearchOssTraceRequest) (response *SearchOssTraceResponse, err error) {
    if request == nil {
        request = NewSearchOssTraceRequest()
    }
    response = NewSearchOssTraceResponse()
    err = c.Send(request, response)
    return
}

func NewSearchRealtimeBusinessLogRequest() (request *SearchRealtimeBusinessLogRequest) {
    request = &SearchRealtimeBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchRealtimeBusinessLog")
    return
}

func NewSearchRealtimeBusinessLogResponse() (response *SearchRealtimeBusinessLogResponse) {
    response = &SearchRealtimeBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务日志实时检索
func (c *Client) SearchRealtimeBusinessLog(request *SearchRealtimeBusinessLogRequest) (response *SearchRealtimeBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchRealtimeBusinessLogRequest()
    }
    response = NewSearchRealtimeBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchRealtimeStdoutLogRequest() (request *SearchRealtimeStdoutLogRequest) {
    request = &SearchRealtimeStdoutLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchRealtimeStdoutLog")
    return
}

func NewSearchRealtimeStdoutLogResponse() (response *SearchRealtimeStdoutLogResponse) {
    response = &SearchRealtimeStdoutLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 标准输出日志实时检索
func (c *Client) SearchRealtimeStdoutLog(request *SearchRealtimeStdoutLogRequest) (response *SearchRealtimeStdoutLogResponse, err error) {
    if request == nil {
        request = NewSearchRealtimeStdoutLogRequest()
    }
    response = NewSearchRealtimeStdoutLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchSpanBusinessLogRequest() (request *SearchSpanBusinessLogRequest) {
    request = &SearchSpanBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchSpanBusinessLog")
    return
}

func NewSearchSpanBusinessLogResponse() (response *SearchSpanBusinessLogResponse) {
    response = &SearchSpanBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务日志搜索（关联调用链Span）
func (c *Client) SearchSpanBusinessLog(request *SearchSpanBusinessLogRequest) (response *SearchSpanBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchSpanBusinessLogRequest()
    }
    response = NewSearchSpanBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchStdoutLogRequest() (request *SearchStdoutLogRequest) {
    request = &SearchStdoutLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchStdoutLog")
    return
}

func NewSearchStdoutLogResponse() (response *SearchStdoutLogResponse) {
    response = &SearchStdoutLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 标准输出日志搜索
func (c *Client) SearchStdoutLog(request *SearchStdoutLogRequest) (response *SearchStdoutLogResponse, err error) {
    if request == nil {
        request = NewSearchStdoutLogRequest()
    }
    response = NewSearchStdoutLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchSurroundBusinessLogRequest() (request *SearchSurroundBusinessLogRequest) {
    request = &SearchSurroundBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchSurroundBusinessLog")
    return
}

func NewSearchSurroundBusinessLogResponse() (response *SearchSurroundBusinessLogResponse) {
    response = &SearchSurroundBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务日志上下文检索
func (c *Client) SearchSurroundBusinessLog(request *SearchSurroundBusinessLogRequest) (response *SearchSurroundBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchSurroundBusinessLogRequest()
    }
    response = NewSearchSurroundBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchTraceRequest() (request *SearchTraceRequest) {
    request = &SearchTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchTrace")
    return
}

func NewSearchTraceResponse() (response *SearchTraceResponse) {
    response = &SearchTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询调用链
func (c *Client) SearchTrace(request *SearchTraceRequest) (response *SearchTraceResponse, err error) {
    if request == nil {
        request = NewSearchTraceRequest()
    }
    response = NewSearchTraceResponse()
    err = c.Send(request, response)
    return
}

func NewSetCkafkaToContainerClusterRequest() (request *SetCkafkaToContainerClusterRequest) {
    request = &SetCkafkaToContainerClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SetCkafkaToContainerCluster")
    return
}

func NewSetCkafkaToContainerClusterResponse() (response *SetCkafkaToContainerClusterResponse) {
    response = &SetCkafkaToContainerClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetCkafkaToContainerCluster
func (c *Client) SetCkafkaToContainerCluster(request *SetCkafkaToContainerClusterRequest) (response *SetCkafkaToContainerClusterResponse, err error) {
    if request == nil {
        request = NewSetCkafkaToContainerClusterRequest()
    }
    response = NewSetCkafkaToContainerClusterResponse()
    err = c.Send(request, response)
    return
}

func NewSetNamespaceCodeRequest() (request *SetNamespaceCodeRequest) {
    request = &SetNamespaceCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SetNamespaceCode")
    return
}

func NewSetNamespaceCodeResponse() (response *SetNamespaceCodeResponse) {
    response = &SetNamespaceCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置命名空间编码
func (c *Client) SetNamespaceCode(request *SetNamespaceCodeRequest) (response *SetNamespaceCodeResponse, err error) {
    if request == nil {
        request = NewSetNamespaceCodeRequest()
    }
    response = NewSetNamespaceCodeResponse()
    err = c.Send(request, response)
    return
}

func NewShirkGroupRequest() (request *ShirkGroupRequest) {
    request = &ShirkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShirkGroup")
    return
}

func NewShirkGroupResponse() (response *ShirkGroupResponse) {
    response = &ShirkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线部署组所有机器实例
func (c *Client) ShirkGroup(request *ShirkGroupRequest) (response *ShirkGroupResponse, err error) {
    if request == nil {
        request = NewShirkGroupRequest()
    }
    response = NewShirkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewShirkInstanceRequest() (request *ShirkInstanceRequest) {
    request = &ShirkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShirkInstance")
    return
}

func NewShirkInstanceResponse() (response *ShirkInstanceResponse) {
    response = &ShirkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线部署组机器
func (c *Client) ShirkInstance(request *ShirkInstanceRequest) (response *ShirkInstanceResponse, err error) {
    if request == nil {
        request = NewShirkInstanceRequest()
    }
    response = NewShirkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewShirkNamespaceRequest() (request *ShirkNamespaceRequest) {
    request = &ShirkNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShirkNamespace")
    return
}

func NewShirkNamespaceResponse() (response *ShirkNamespaceResponse) {
    response = &ShirkNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 移除命名空间机器
func (c *Client) ShirkNamespace(request *ShirkNamespaceRequest) (response *ShirkNamespaceResponse, err error) {
    if request == nil {
        request = NewShirkNamespaceRequest()
    }
    response = NewShirkNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkGroupRequest() (request *ShrinkGroupRequest) {
    request = &ShrinkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkGroup")
    return
}

func NewShrinkGroupResponse() (response *ShrinkGroupResponse) {
    response = &ShrinkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线部署组所有机器实例
func (c *Client) ShrinkGroup(request *ShrinkGroupRequest) (response *ShrinkGroupResponse, err error) {
    if request == nil {
        request = NewShrinkGroupRequest()
    }
    response = NewShrinkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkInstanceRequest() (request *ShrinkInstanceRequest) {
    request = &ShrinkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkInstance")
    return
}

func NewShrinkInstanceResponse() (response *ShrinkInstanceResponse) {
    response = &ShrinkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 虚拟机部署组下线实例
func (c *Client) ShrinkInstance(request *ShrinkInstanceRequest) (response *ShrinkInstanceResponse, err error) {
    if request == nil {
        request = NewShrinkInstanceRequest()
    }
    response = NewShrinkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkInstancesRequest() (request *ShrinkInstancesRequest) {
    request = &ShrinkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkInstances")
    return
}

func NewShrinkInstancesResponse() (response *ShrinkInstancesResponse) {
    response = &ShrinkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 虚拟机部署组下线实例
func (c *Client) ShrinkInstances(request *ShrinkInstancesRequest) (response *ShrinkInstancesResponse, err error) {
    if request == nil {
        request = NewShrinkInstancesRequest()
    }
    response = NewShrinkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartContainerGroupRequest() (request *StartContainerGroupRequest) {
    request = &StartContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StartContainerGroup")
    return
}

func NewStartContainerGroupResponse() (response *StartContainerGroupResponse) {
    response = &StartContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启动容器部署组
func (c *Client) StartContainerGroup(request *StartContainerGroupRequest) (response *StartContainerGroupResponse, err error) {
    if request == nil {
        request = NewStartContainerGroupRequest()
    }
    response = NewStartContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStartGroupRequest() (request *StartGroupRequest) {
    request = &StartGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StartGroup")
    return
}

func NewStartGroupResponse() (response *StartGroupResponse) {
    response = &StartGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启动分组
func (c *Client) StartGroup(request *StartGroupRequest) (response *StartGroupResponse, err error) {
    if request == nil {
        request = NewStartGroupRequest()
    }
    response = NewStartGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstanceRequest() (request *StartInstanceRequest) {
    request = &StartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StartInstance")
    return
}

func NewStartInstanceResponse() (response *StartInstanceResponse) {
    response = &StartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启动机器
func (c *Client) StartInstance(request *StartInstanceRequest) (response *StartInstanceResponse, err error) {
    if request == nil {
        request = NewStartInstanceRequest()
    }
    response = NewStartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStartServerlessGroupRequest() (request *StartServerlessGroupRequest) {
    request = &StartServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StartServerlessGroup")
    return
}

func NewStartServerlessGroupResponse() (response *StartServerlessGroupResponse) {
    response = &StartServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) StartServerlessGroup(request *StartServerlessGroupRequest) (response *StartServerlessGroupResponse, err error) {
    if request == nil {
        request = NewStartServerlessGroupRequest()
    }
    response = NewStartServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopContainerGroupRequest() (request *StopContainerGroupRequest) {
    request = &StopContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopContainerGroup")
    return
}

func NewStopContainerGroupResponse() (response *StopContainerGroupResponse) {
    response = &StopContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止容器部署组
func (c *Client) StopContainerGroup(request *StopContainerGroupRequest) (response *StopContainerGroupResponse, err error) {
    if request == nil {
        request = NewStopContainerGroupRequest()
    }
    response = NewStopContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopGroupRequest() (request *StopGroupRequest) {
    request = &StopGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopGroup")
    return
}

func NewStopGroupResponse() (response *StopGroupResponse) {
    response = &StopGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止虚拟机部署组
func (c *Client) StopGroup(request *StopGroupRequest) (response *StopGroupResponse, err error) {
    if request == nil {
        request = NewStopGroupRequest()
    }
    response = NewStopGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstanceRequest() (request *StopInstanceRequest) {
    request = &StopInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopInstance")
    return
}

func NewStopInstanceResponse() (response *StopInstanceResponse) {
    response = &StopInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止机器
func (c *Client) StopInstance(request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
    if request == nil {
        request = NewStopInstanceRequest()
    }
    response = NewStopInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopServerlessGroupRequest() (request *StopServerlessGroupRequest) {
    request = &StopServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopServerlessGroup")
    return
}

func NewStopServerlessGroupResponse() (response *StopServerlessGroupResponse) {
    response = &StopServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) StopServerlessGroup(request *StopServerlessGroupRequest) (response *StopServerlessGroupResponse, err error) {
    if request == nil {
        request = NewStopServerlessGroupRequest()
    }
    response = NewStopServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSynchronizeContainerClusterRequest() (request *SynchronizeContainerClusterRequest) {
    request = &SynchronizeContainerClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SynchronizeContainerCluster")
    return
}

func NewSynchronizeContainerClusterResponse() (response *SynchronizeContainerClusterResponse) {
    response = &SynchronizeContainerClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 同步容器平台集群到tsf
func (c *Client) SynchronizeContainerCluster(request *SynchronizeContainerClusterRequest) (response *SynchronizeContainerClusterResponse, err error) {
    if request == nil {
        request = NewSynchronizeContainerClusterRequest()
    }
    response = NewSynchronizeContainerClusterResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindApiGroupRequest() (request *UnbindApiGroupRequest) {
    request = &UnbindApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UnbindApiGroup")
    return
}

func NewUnbindApiGroupResponse() (response *UnbindApiGroupResponse) {
    response = &UnbindApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// API分组批量与网关解绑
func (c *Client) UnbindApiGroup(request *UnbindApiGroupRequest) (response *UnbindApiGroupResponse, err error) {
    if request == nil {
        request = NewUnbindApiGroupRequest()
    }
    response = NewUnbindApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindPluginRequest() (request *UnbindPluginRequest) {
    request = &UnbindPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UnbindPlugin")
    return
}

func NewUnbindPluginResponse() (response *UnbindPluginResponse) {
    response = &UnbindPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 插件批量解绑
func (c *Client) UnbindPlugin(request *UnbindPluginRequest) (response *UnbindPluginResponse, err error) {
    if request == nil {
        request = NewUnbindPluginRequest()
    }
    response = NewUnbindPluginResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallAgentScriptsRequest() (request *UninstallAgentScriptsRequest) {
    request = &UninstallAgentScriptsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UninstallAgentScripts")
    return
}

func NewUninstallAgentScriptsResponse() (response *UninstallAgentScriptsResponse) {
    response = &UninstallAgentScriptsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 虚拟机集群导入云主机，移除机器卸载Agent脚本
func (c *Client) UninstallAgentScripts(request *UninstallAgentScriptsRequest) (response *UninstallAgentScriptsResponse, err error) {
    if request == nil {
        request = NewUninstallAgentScriptsRequest()
    }
    response = NewUninstallAgentScriptsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiGroupRequest() (request *UpdateApiGroupRequest) {
    request = &UpdateApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiGroup")
    return
}

func NewUpdateApiGroupResponse() (response *UpdateApiGroupResponse) {
    response = &UpdateApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新Api分组
func (c *Client) UpdateApiGroup(request *UpdateApiGroupRequest) (response *UpdateApiGroupResponse, err error) {
    if request == nil {
        request = NewUpdateApiGroupRequest()
    }
    response = NewUpdateApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiRateLimitRuleRequest() (request *UpdateApiRateLimitRuleRequest) {
    request = &UpdateApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRule")
    return
}

func NewUpdateApiRateLimitRuleResponse() (response *UpdateApiRateLimitRuleResponse) {
    response = &UpdateApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新API限流规则
func (c *Client) UpdateApiRateLimitRule(request *UpdateApiRateLimitRuleRequest) (response *UpdateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRuleRequest()
    }
    response = NewUpdateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApplicationAgentRequest() (request *UpdateApplicationAgentRequest) {
    request = &UpdateApplicationAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApplicationAgent")
    return
}

func NewUpdateApplicationAgentResponse() (response *UpdateApplicationAgentResponse) {
    response = &UpdateApplicationAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新Agent
func (c *Client) UpdateApplicationAgent(request *UpdateApplicationAgentRequest) (response *UpdateApplicationAgentResponse, err error) {
    if request == nil {
        request = NewUpdateApplicationAgentRequest()
    }
    response = NewUpdateApplicationAgentResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateBusinessLogConfigRequest() (request *UpdateBusinessLogConfigRequest) {
    request = &UpdateBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateBusinessLogConfig")
    return
}

func NewUpdateBusinessLogConfigResponse() (response *UpdateBusinessLogConfigResponse) {
    response = &UpdateBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新业务日志配置
func (c *Client) UpdateBusinessLogConfig(request *UpdateBusinessLogConfigRequest) (response *UpdateBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewUpdateBusinessLogConfigRequest()
    }
    response = NewUpdateBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCircuitBreakerRuleRequest() (request *UpdateCircuitBreakerRuleRequest) {
    request = &UpdateCircuitBreakerRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateCircuitBreakerRule")
    return
}

func NewUpdateCircuitBreakerRuleResponse() (response *UpdateCircuitBreakerRuleResponse) {
    response = &UpdateCircuitBreakerRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新熔断规则
func (c *Client) UpdateCircuitBreakerRule(request *UpdateCircuitBreakerRuleRequest) (response *UpdateCircuitBreakerRuleResponse, err error) {
    if request == nil {
        request = NewUpdateCircuitBreakerRuleRequest()
    }
    response = NewUpdateCircuitBreakerRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConfigTemplateRequest() (request *UpdateConfigTemplateRequest) {
    request = &UpdateConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateConfigTemplate")
    return
}

func NewUpdateConfigTemplateResponse() (response *UpdateConfigTemplateResponse) {
    response = &UpdateConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新参数模板
func (c *Client) UpdateConfigTemplate(request *UpdateConfigTemplateRequest) (response *UpdateConfigTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateConfigTemplateRequest()
    }
    response = NewUpdateConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateContainerGroupRequest() (request *UpdateContainerGroupRequest) {
    request = &UpdateContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateContainerGroup")
    return
}

func NewUpdateContainerGroupResponse() (response *UpdateContainerGroupResponse) {
    response = &UpdateContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改容器部署组
func (c *Client) UpdateContainerGroup(request *UpdateContainerGroupRequest) (response *UpdateContainerGroupResponse, err error) {
    if request == nil {
        request = NewUpdateContainerGroupRequest()
    }
    response = NewUpdateContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGatewayJwtPluginRequest() (request *UpdateGatewayJwtPluginRequest) {
    request = &UpdateGatewayJwtPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayJwtPlugin")
    return
}

func NewUpdateGatewayJwtPluginResponse() (response *UpdateGatewayJwtPluginResponse) {
    response = &UpdateGatewayJwtPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改网关Jwt插件
func (c *Client) UpdateGatewayJwtPlugin(request *UpdateGatewayJwtPluginRequest) (response *UpdateGatewayJwtPluginResponse, err error) {
    if request == nil {
        request = NewUpdateGatewayJwtPluginRequest()
    }
    response = NewUpdateGatewayJwtPluginResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGatewayOAuthPluginRequest() (request *UpdateGatewayOAuthPluginRequest) {
    request = &UpdateGatewayOAuthPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayOAuthPlugin")
    return
}

func NewUpdateGatewayOAuthPluginResponse() (response *UpdateGatewayOAuthPluginResponse) {
    response = &UpdateGatewayOAuthPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改网关OAuth插件
func (c *Client) UpdateGatewayOAuthPlugin(request *UpdateGatewayOAuthPluginRequest) (response *UpdateGatewayOAuthPluginResponse, err error) {
    if request == nil {
        request = NewUpdateGatewayOAuthPluginRequest()
    }
    response = NewUpdateGatewayOAuthPluginResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGatewayTagPluginRequest() (request *UpdateGatewayTagPluginRequest) {
    request = &UpdateGatewayTagPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayTagPlugin")
    return
}

func NewUpdateGatewayTagPluginResponse() (response *UpdateGatewayTagPluginResponse) {
    response = &UpdateGatewayTagPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改网关Tag插件
func (c *Client) UpdateGatewayTagPlugin(request *UpdateGatewayTagPluginRequest) (response *UpdateGatewayTagPluginResponse, err error) {
    if request == nil {
        request = NewUpdateGatewayTagPluginRequest()
    }
    response = NewUpdateGatewayTagPluginResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGroupSecretRequest() (request *UpdateGroupSecretRequest) {
    request = &UpdateGroupSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateGroupSecret")
    return
}

func NewUpdateGroupSecretResponse() (response *UpdateGroupSecretResponse) {
    response = &UpdateGroupSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新API分组的秘钥
func (c *Client) UpdateGroupSecret(request *UpdateGroupSecretRequest) (response *UpdateGroupSecretResponse, err error) {
    if request == nil {
        request = NewUpdateGroupSecretRequest()
    }
    response = NewUpdateGroupSecretResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRatelimitRequest() (request *UpdateRatelimitRequest) {
    request = &UpdateRatelimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateRatelimit")
    return
}

func NewUpdateRatelimitResponse() (response *UpdateRatelimitResponse) {
    response = &UpdateRatelimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改限流规则
func (c *Client) UpdateRatelimit(request *UpdateRatelimitRequest) (response *UpdateRatelimitResponse, err error) {
    if request == nil {
        request = NewUpdateRatelimitRequest()
    }
    response = NewUpdateRatelimitResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUploadInfoRequest() (request *UpdateUploadInfoRequest) {
    request = &UpdateUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateUploadInfo")
    return
}

func NewUpdateUploadInfoResponse() (response *UpdateUploadInfoResponse) {
    response = &UpdateUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 调用TSF的GetUploadInfo接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
// 调用此接口完成后，才标志上传包流程结束。
func (c *Client) UpdateUploadInfo(request *UpdateUploadInfoRequest) (response *UpdateUploadInfoResponse, err error) {
    if request == nil {
        request = NewUpdateUploadInfoRequest()
    }
    response = NewUpdateUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewUploadLicenseApplicationRequest() (request *UploadLicenseApplicationRequest) {
    request = &UploadLicenseApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UploadLicenseApplication")
    return
}

func NewUploadLicenseApplicationResponse() (response *UploadLicenseApplicationResponse) {
    response = &UploadLicenseApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 递交许可申请
func (c *Client) UploadLicenseApplication(request *UploadLicenseApplicationRequest) (response *UploadLicenseApplicationResponse, err error) {
    if request == nil {
        request = NewUploadLicenseApplicationRequest()
    }
    response = NewUploadLicenseApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewUploadPkgRequest() (request *UploadPkgRequest) {
    request = &UploadPkgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UploadPkg")
    return
}

func NewUploadPkgResponse() (response *UploadPkgResponse) {
    response = &UploadPkgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 上传包test
func (c *Client) UploadPkg(request *UploadPkgRequest) (response *UploadPkgResponse, err error) {
    if request == nil {
        request = NewUploadPkgRequest()
    }
    response = NewUploadPkgResponse()
    err = c.Send(request, response)
    return
}

func NewValidateDeleteConfigRequest() (request *ValidateDeleteConfigRequest) {
    request = &ValidateDeleteConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ValidateDeleteConfig")
    return
}

func NewValidateDeleteConfigResponse() (response *ValidateDeleteConfigResponse) {
    response = &ValidateDeleteConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 校验配置项可删除
func (c *Client) ValidateDeleteConfig(request *ValidateDeleteConfigRequest) (response *ValidateDeleteConfigResponse, err error) {
    if request == nil {
        request = NewValidateDeleteConfigRequest()
    }
    response = NewValidateDeleteConfigResponse()
    err = c.Send(request, response)
    return
}

func NewValidateDeletePublicConfigRequest() (request *ValidateDeletePublicConfigRequest) {
    request = &ValidateDeletePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ValidateDeletePublicConfig")
    return
}

func NewValidateDeletePublicConfigResponse() (response *ValidateDeletePublicConfigResponse) {
    response = &ValidateDeletePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 校验公共配置项可删除
func (c *Client) ValidateDeletePublicConfig(request *ValidateDeletePublicConfigRequest) (response *ValidateDeletePublicConfigResponse, err error) {
    if request == nil {
        request = NewValidateDeletePublicConfigRequest()
    }
    response = NewValidateDeletePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewValidateLogSchemaRequest() (request *ValidateLogSchemaRequest) {
    request = &ValidateLogSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ValidateLogSchema")
    return
}

func NewValidateLogSchemaResponse() (response *ValidateLogSchemaResponse) {
    response = &ValidateLogSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 校验日志Pattern格式合法性 
func (c *Client) ValidateLogSchema(request *ValidateLogSchemaRequest) (response *ValidateLogSchemaResponse, err error) {
    if request == nil {
        request = NewValidateLogSchemaRequest()
    }
    response = NewValidateLogSchemaResponse()
    err = c.Send(request, response)
    return
}

func NewValidateNamespaceClusterVPCRequest() (request *ValidateNamespaceClusterVPCRequest) {
    request = &ValidateNamespaceClusterVPCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ValidateNamespaceClusterVPC")
    return
}

func NewValidateNamespaceClusterVPCResponse() (response *ValidateNamespaceClusterVPCResponse) {
    response = &ValidateNamespaceClusterVPCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 校验命名空间归属集群VPC一致性
func (c *Client) ValidateNamespaceClusterVPC(request *ValidateNamespaceClusterVPCRequest) (response *ValidateNamespaceClusterVPCResponse, err error) {
    if request == nil {
        request = NewValidateNamespaceClusterVPCRequest()
    }
    response = NewValidateNamespaceClusterVPCResponse()
    err = c.Send(request, response)
    return
}
