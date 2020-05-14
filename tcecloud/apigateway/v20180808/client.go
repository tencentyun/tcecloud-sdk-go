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

package v20180808

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-08-08"

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


func NewBindEnvironmentRequest() (request *BindEnvironmentRequest) {
    request = &BindEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "BindEnvironment")
    return
}

func NewBindEnvironmentResponse() (response *BindEnvironmentResponse) {
    response = &BindEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用计划绑定环境
func (c *Client) BindEnvironment(request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
    if request == nil {
        request = NewBindEnvironmentRequest()
    }
    response = NewBindEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewBindIPStrategyRequest() (request *BindIPStrategyRequest) {
    request = &BindIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "BindIPStrategy")
    return
}

func NewBindIPStrategyResponse() (response *BindIPStrategyResponse) {
    response = &BindIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 绑定ip策略
func (c *Client) BindIPStrategy(request *BindIPStrategyRequest) (response *BindIPStrategyResponse, err error) {
    if request == nil {
        request = NewBindIPStrategyRequest()
    }
    response = NewBindIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewBindSecretIdsRequest() (request *BindSecretIdsRequest) {
    request = &BindSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "BindSecretIds")
    return
}

func NewBindSecretIdsResponse() (response *BindSecretIdsResponse) {
    response = &BindSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用计划绑定密钥
func (c *Client) BindSecretIds(request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
    if request == nil {
        request = NewBindSecretIdsRequest()
    }
    response = NewBindSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewBindSubDomainRequest() (request *BindSubDomainRequest) {
    request = &BindSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "BindSubDomain")
    return
}

func NewBindSubDomainResponse() (response *BindSubDomainResponse) {
    response = &BindSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 绑定自定义域名
func (c *Client) BindSubDomain(request *BindSubDomainRequest) (response *BindSubDomainResponse, err error) {
    if request == nil {
        request = NewBindSubDomainRequest()
    }
    response = NewBindSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCloneApisRequest() (request *CheckCloneApisRequest) {
    request = &CheckCloneApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CheckCloneApis")
    return
}

func NewCheckCloneApisResponse() (response *CheckCloneApisResponse) {
    response = &CheckCloneApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 校验克隆api 是否满足要求
func (c *Client) CheckCloneApis(request *CheckCloneApisRequest) (response *CheckCloneApisResponse, err error) {
    if request == nil {
        request = NewCheckCloneApisRequest()
    }
    response = NewCheckCloneApisResponse()
    err = c.Send(request, response)
    return
}

func NewCheckServiceConfigRequest() (request *CheckServiceConfigRequest) {
    request = &CheckServiceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CheckServiceConfig")
    return
}

func NewCheckServiceConfigResponse() (response *CheckServiceConfigResponse) {
    response = &CheckServiceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查服务配置合法性
func (c *Client) CheckServiceConfig(request *CheckServiceConfigRequest) (response *CheckServiceConfigResponse, err error) {
    if request == nil {
        request = NewCheckServiceConfigRequest()
    }
    response = NewCheckServiceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCloneApisRequest() (request *CloneApisRequest) {
    request = &CloneApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CloneApis")
    return
}

func NewCloneApisResponse() (response *CloneApisResponse) {
    response = &CloneApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 克隆api
func (c *Client) CloneApis(request *CloneApisRequest) (response *CloneApisResponse, err error) {
    if request == nil {
        request = NewCloneApisRequest()
    }
    response = NewCloneApisResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRequest() (request *CreateApiRequest) {
    request = &CreateApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApi")
    return
}

func NewCreateApiResponse() (response *CreateApiResponse) {
    response = &CreateApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建api
func (c *Client) CreateApi(request *CreateApiRequest) (response *CreateApiResponse, err error) {
    if request == nil {
        request = NewCreateApiRequest()
    }
    response = NewCreateApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
    request = &CreateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApiKey")
    return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
    response = &CreateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建密钥
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
    response = NewCreateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIPStrategyRequest() (request *CreateIPStrategyRequest) {
    request = &CreateIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateIPStrategy")
    return
}

func NewCreateIPStrategyResponse() (response *CreateIPStrategyResponse) {
    response = &CreateIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建ip策略
func (c *Client) CreateIPStrategy(request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
    if request == nil {
        request = NewCreateIPStrategyRequest()
    }
    response = NewCreateIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogRuleRequest() (request *CreateLogRuleRequest) {
    request = &CreateLogRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateLogRule")
    return
}

func NewCreateLogRuleResponse() (response *CreateLogRuleResponse) {
    response = &CreateLogRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用于日志上报的规则。
func (c *Client) CreateLogRule(request *CreateLogRuleRequest) (response *CreateLogRuleResponse, err error) {
    if request == nil {
        request = NewCreateLogRuleRequest()
    }
    response = NewCreateLogRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
    request = &CreateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateService")
    return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
    response = &CreateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建服务 
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    response = NewCreateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTcbScfApiRequest() (request *CreateTcbScfApiRequest) {
    request = &CreateTcbScfApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateTcbScfApi")
    return
}

func NewCreateTcbScfApiResponse() (response *CreateTcbScfApiResponse) {
    response = &CreateTcbScfApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建TCB-SCF-HTTP触发器
func (c *Client) CreateTcbScfApi(request *CreateTcbScfApiRequest) (response *CreateTcbScfApiResponse, err error) {
    if request == nil {
        request = NewCreateTcbScfApiRequest()
    }
    response = NewCreateTcbScfApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsagePlanRequest() (request *CreateUsagePlanRequest) {
    request = &CreateUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateUsagePlan")
    return
}

func NewCreateUsagePlanResponse() (response *CreateUsagePlanResponse) {
    response = &CreateUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建使用计划
func (c *Client) CreateUsagePlan(request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
    if request == nil {
        request = NewCreateUsagePlanRequest()
    }
    response = NewCreateUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiRequest() (request *DeleteApiRequest) {
    request = &DeleteApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApi")
    return
}

func NewDeleteApiResponse() (response *DeleteApiResponse) {
    response = &DeleteApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除api
func (c *Client) DeleteApi(request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
    if request == nil {
        request = NewDeleteApiRequest()
    }
    response = NewDeleteApiResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
    request = &DeleteApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApiKey")
    return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
    response = &DeleteApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除密钥
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
    response = NewDeleteApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApisRequest() (request *DeleteApisRequest) {
    request = &DeleteApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApis")
    return
}

func NewDeleteApisResponse() (response *DeleteApisResponse) {
    response = &DeleteApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除API，支持多个API参数
// 
//         self.serviceId = self.args.get('serviceId')
//         self.apiId = self.args.get('apiId')
func (c *Client) DeleteApis(request *DeleteApisRequest) (response *DeleteApisResponse, err error) {
    if request == nil {
        request = NewDeleteApisRequest()
    }
    response = NewDeleteApisResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIPStrategyRequest() (request *DeleteIPStrategyRequest) {
    request = &DeleteIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteIPStrategy")
    return
}

func NewDeleteIPStrategyResponse() (response *DeleteIPStrategyResponse) {
    response = &DeleteIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除ip策略
func (c *Client) DeleteIPStrategy(request *DeleteIPStrategyRequest) (response *DeleteIPStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteIPStrategyRequest()
    }
    response = NewDeleteIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogRulesRequest() (request *DeleteLogRulesRequest) {
    request = &DeleteLogRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteLogRules")
    return
}

func NewDeleteLogRulesResponse() (response *DeleteLogRulesResponse) {
    response = &DeleteLogRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除日志规则。
func (c *Client) DeleteLogRules(request *DeleteLogRulesRequest) (response *DeleteLogRulesResponse, err error) {
    if request == nil {
        request = NewDeleteLogRulesRequest()
    }
    response = NewDeleteLogRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
    request = &DeleteServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteService")
    return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
    response = &DeleteServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除服务
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    if request == nil {
        request = NewDeleteServiceRequest()
    }
    response = NewDeleteServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceSubDomainMappingRequest() (request *DeleteServiceSubDomainMappingRequest) {
    request = &DeleteServiceSubDomainMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteServiceSubDomainMapping")
    return
}

func NewDeleteServiceSubDomainMappingResponse() (response *DeleteServiceSubDomainMappingResponse) {
    response = &DeleteServiceSubDomainMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除自定义域名的路径映射，需要保证至少有2个路径映射，否则删除会报错。
func (c *Client) DeleteServiceSubDomainMapping(request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
    if request == nil {
        request = NewDeleteServiceSubDomainMappingRequest()
    }
    response = NewDeleteServiceSubDomainMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTcbScfApiRequest() (request *DeleteTcbScfApiRequest) {
    request = &DeleteTcbScfApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteTcbScfApi")
    return
}

func NewDeleteTcbScfApiResponse() (response *DeleteTcbScfApiResponse) {
    response = &DeleteTcbScfApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于删除TCB-SCF-HTTP触发器
func (c *Client) DeleteTcbScfApi(request *DeleteTcbScfApiRequest) (response *DeleteTcbScfApiResponse, err error) {
    if request == nil {
        request = NewDeleteTcbScfApiRequest()
    }
    response = NewDeleteTcbScfApiResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsagePlanRequest() (request *DeleteUsagePlanRequest) {
    request = &DeleteUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteUsagePlan")
    return
}

func NewDeleteUsagePlanResponse() (response *DeleteUsagePlanResponse) {
    response = &DeleteUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除使用计划
func (c *Client) DeleteUsagePlan(request *DeleteUsagePlanRequest) (response *DeleteUsagePlanResponse, err error) {
    if request == nil {
        request = NewDeleteUsagePlanRequest()
    }
    response = NewDeleteUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDemoteServiceUsagePlanRequest() (request *DemoteServiceUsagePlanRequest) {
    request = &DemoteServiceUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DemoteServiceUsagePlan")
    return
}

func NewDemoteServiceUsagePlanResponse() (response *DemoteServiceUsagePlanResponse) {
    response = &DemoteServiceUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用计划降级
func (c *Client) DemoteServiceUsagePlan(request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDemoteServiceUsagePlanRequest()
    }
    response = NewDemoteServiceUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiRequest() (request *DescribeApiRequest) {
    request = &DescribeApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApi")
    return
}

func NewDescribeApiResponse() (response *DescribeApiResponse) {
    response = &DescribeApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API详情
func (c *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
    if request == nil {
        request = NewDescribeApiRequest()
    }
    response = NewDescribeApiResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiEnvironmentApiKeysRequest() (request *DescribeApiEnvironmentApiKeysRequest) {
    request = &DescribeApiEnvironmentApiKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiEnvironmentApiKeys")
    return
}

func NewDescribeApiEnvironmentApiKeysResponse() (response *DescribeApiEnvironmentApiKeysResponse) {
    response = &DescribeApiEnvironmentApiKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询已发布环境api密钥
func (c *Client) DescribeApiEnvironmentApiKeys(request *DescribeApiEnvironmentApiKeysRequest) (response *DescribeApiEnvironmentApiKeysResponse, err error) {
    if request == nil {
        request = NewDescribeApiEnvironmentApiKeysRequest()
    }
    response = NewDescribeApiEnvironmentApiKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiEnvironmentStrategyRequest() (request *DescribeApiEnvironmentStrategyRequest) {
    request = &DescribeApiEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiEnvironmentStrategy")
    return
}

func NewDescribeApiEnvironmentStrategyResponse() (response *DescribeApiEnvironmentStrategyResponse) {
    response = &DescribeApiEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示发布api绑定的ip策略
func (c *Client) DescribeApiEnvironmentStrategy(request *DescribeApiEnvironmentStrategyRequest) (response *DescribeApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeApiEnvironmentStrategyRequest()
    }
    response = NewDescribeApiEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyRequest() (request *DescribeApiKeyRequest) {
    request = &DescribeApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKey")
    return
}

func NewDescribeApiKeyResponse() (response *DescribeApiKeyResponse) {
    response = &DescribeApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示密钥详情
func (c *Client) DescribeApiKey(request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyRequest()
    }
    response = NewDescribeApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeysStatusRequest() (request *DescribeApiKeysStatusRequest) {
    request = &DescribeApiKeysStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKeysStatus")
    return
}

func NewDescribeApiKeysStatusResponse() (response *DescribeApiKeysStatusResponse) {
    response = &DescribeApiKeysStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示密钥列表
func (c *Client) DescribeApiKeysStatus(request *DescribeApiKeysStatusRequest) (response *DescribeApiKeysStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeysStatusRequest()
    }
    response = NewDescribeApiKeysStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiUsagePlanRequest() (request *DescribeApiUsagePlanRequest) {
    request = &DescribeApiUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiUsagePlan")
    return
}

func NewDescribeApiUsagePlanResponse() (response *DescribeApiUsagePlanResponse) {
    response = &DescribeApiUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务下api绑定的使用计划
func (c *Client) DescribeApiUsagePlan(request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeApiUsagePlanRequest()
    }
    response = NewDescribeApiUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApisStatusRequest() (request *DescribeApisStatusRequest) {
    request = &DescribeApisStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApisStatus")
    return
}

func NewDescribeApisStatusResponse() (response *DescribeApisStatusResponse) {
    response = &DescribeApisStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询API状态
func (c *Client) DescribeApisStatus(request *DescribeApisStatusRequest) (response *DescribeApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApisStatusRequest()
    }
    response = NewDescribeApisStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainRequest() (request *DescribeDomainRequest) {
    request = &DescribeDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeDomain")
    return
}

func NewDescribeDomainResponse() (response *DescribeDomainResponse) {
    response = &DescribeDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询域名状态
func (c *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRequest()
    }
    response = NewDescribeDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExclusiveSetRequest() (request *DescribeExclusiveSetRequest) {
    request = &DescribeExclusiveSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeExclusiveSet")
    return
}

func NewDescribeExclusiveSetResponse() (response *DescribeExclusiveSetResponse) {
    response = &DescribeExclusiveSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询独立集群列表
func (c *Client) DescribeExclusiveSet(request *DescribeExclusiveSetRequest) (response *DescribeExclusiveSetResponse, err error) {
    if request == nil {
        request = NewDescribeExclusiveSetRequest()
    }
    response = NewDescribeExclusiveSetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategyRequest() (request *DescribeIPStrategyRequest) {
    request = &DescribeIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategy")
    return
}

func NewDescribeIPStrategyResponse() (response *DescribeIPStrategyResponse) {
    response = &DescribeIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询ip策略详情
func (c *Client) DescribeIPStrategy(request *DescribeIPStrategyRequest) (response *DescribeIPStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyRequest()
    }
    response = NewDescribeIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategyApisStatusRequest() (request *DescribeIPStrategyApisStatusRequest) {
    request = &DescribeIPStrategyApisStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategyApisStatus")
    return
}

func NewDescribeIPStrategyApisStatusResponse() (response *DescribeIPStrategyApisStatusResponse) {
    response = &DescribeIPStrategyApisStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询ip策略绑定api
func (c *Client) DescribeIPStrategyApisStatus(request *DescribeIPStrategyApisStatusRequest) (response *DescribeIPStrategyApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyApisStatusRequest()
    }
    response = NewDescribeIPStrategyApisStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategysStatusRequest() (request *DescribeIPStrategysStatusRequest) {
    request = &DescribeIPStrategysStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategysStatus")
    return
}

func NewDescribeIPStrategysStatusResponse() (response *DescribeIPStrategysStatusResponse) {
    response = &DescribeIPStrategysStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询ip策略列表
func (c *Client) DescribeIPStrategysStatus(request *DescribeIPStrategysStatusRequest) (response *DescribeIPStrategysStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategysStatusRequest()
    }
    response = NewDescribeIPStrategysStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogRulesRequest() (request *DescribeLogRulesRequest) {
    request = &DescribeLogRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeLogRules")
    return
}

func NewDescribeLogRulesResponse() (response *DescribeLogRulesResponse) {
    response = &DescribeLogRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询日志规则列表
func (c *Client) DescribeLogRules(request *DescribeLogRulesRequest) (response *DescribeLogRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLogRulesRequest()
    }
    response = NewDescribeLogRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogSearchRequest() (request *DescribeLogSearchRequest) {
    request = &DescribeLogSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeLogSearch")
    return
}

func NewDescribeLogSearchResponse() (response *DescribeLogSearchResponse) {
    response = &DescribeLogSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于搜索日志
func (c *Client) DescribeLogSearch(request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    if request == nil {
        request = NewDescribeLogSearchRequest()
    }
    response = NewDescribeLogSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
    request = &DescribeServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeService")
    return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
    response = &DescribeServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询服务详情 
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    if request == nil {
        request = NewDescribeServiceRequest()
    }
    response = NewDescribeServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentKeyMonitorUploadRequest() (request *DescribeServiceEnvironmentKeyMonitorUploadRequest) {
    request = &DescribeServiceEnvironmentKeyMonitorUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentKeyMonitorUpload")
    return
}

func NewDescribeServiceEnvironmentKeyMonitorUploadResponse() (response *DescribeServiceEnvironmentKeyMonitorUploadResponse) {
    response = &DescribeServiceEnvironmentKeyMonitorUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务的环境是否进行key维度的监控数据上报
func (c *Client) DescribeServiceEnvironmentKeyMonitorUpload(request *DescribeServiceEnvironmentKeyMonitorUploadRequest) (response *DescribeServiceEnvironmentKeyMonitorUploadResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentKeyMonitorUploadRequest()
    }
    response = NewDescribeServiceEnvironmentKeyMonitorUploadResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentListRequest() (request *DescribeServiceEnvironmentListRequest) {
    request = &DescribeServiceEnvironmentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentList")
    return
}

func NewDescribeServiceEnvironmentListResponse() (response *DescribeServiceEnvironmentListResponse) {
    response = &DescribeServiceEnvironmentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务环境列表
func (c *Client) DescribeServiceEnvironmentList(request *DescribeServiceEnvironmentListRequest) (response *DescribeServiceEnvironmentListResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentListRequest()
    }
    response = NewDescribeServiceEnvironmentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentReleaseHistoryRequest() (request *DescribeServiceEnvironmentReleaseHistoryRequest) {
    request = &DescribeServiceEnvironmentReleaseHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentReleaseHistory")
    return
}

func NewDescribeServiceEnvironmentReleaseHistoryResponse() (response *DescribeServiceEnvironmentReleaseHistoryResponse) {
    response = &DescribeServiceEnvironmentReleaseHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示服务发布历史
func (c *Client) DescribeServiceEnvironmentReleaseHistory(request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentReleaseHistoryRequest()
    }
    response = NewDescribeServiceEnvironmentReleaseHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentStrategyRequest() (request *DescribeServiceEnvironmentStrategyRequest) {
    request = &DescribeServiceEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentStrategy")
    return
}

func NewDescribeServiceEnvironmentStrategyResponse() (response *DescribeServiceEnvironmentStrategyResponse) {
    response = &DescribeServiceEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示环境策略
func (c *Client) DescribeServiceEnvironmentStrategy(request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentStrategyRequest()
    }
    response = NewDescribeServiceEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceReleaseVersionRequest() (request *DescribeServiceReleaseVersionRequest) {
    request = &DescribeServiceReleaseVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceReleaseVersion")
    return
}

func NewDescribeServiceReleaseVersionResponse() (response *DescribeServiceReleaseVersionResponse) {
    response = &DescribeServiceReleaseVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取服务发布版本
func (c *Client) DescribeServiceReleaseVersion(request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
    if request == nil {
        request = NewDescribeServiceReleaseVersionRequest()
    }
    response = NewDescribeServiceReleaseVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceSubDomainMappingsRequest() (request *DescribeServiceSubDomainMappingsRequest) {
    request = &DescribeServiceSubDomainMappingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomainMappings")
    return
}

func NewDescribeServiceSubDomainMappingsResponse() (response *DescribeServiceSubDomainMappingsResponse) {
    response = &DescribeServiceSubDomainMappingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询自定义域名的路径映射
func (c *Client) DescribeServiceSubDomainMappings(request *DescribeServiceSubDomainMappingsRequest) (response *DescribeServiceSubDomainMappingsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainMappingsRequest()
    }
    response = NewDescribeServiceSubDomainMappingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceSubDomainsRequest() (request *DescribeServiceSubDomainsRequest) {
    request = &DescribeServiceSubDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomains")
    return
}

func NewDescribeServiceSubDomainsResponse() (response *DescribeServiceSubDomainsResponse) {
    response = &DescribeServiceSubDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务自定义域名列表
func (c *Client) DescribeServiceSubDomains(request *DescribeServiceSubDomainsRequest) (response *DescribeServiceSubDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainsRequest()
    }
    response = NewDescribeServiceSubDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceUsagePlanRequest() (request *DescribeServiceUsagePlanRequest) {
    request = &DescribeServiceUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceUsagePlan")
    return
}

func NewDescribeServiceUsagePlanResponse() (response *DescribeServiceUsagePlanResponse) {
    response = &DescribeServiceUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务绑定使用计划列表
func (c *Client) DescribeServiceUsagePlan(request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeServiceUsagePlanRequest()
    }
    response = NewDescribeServiceUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServicesStatusRequest() (request *DescribeServicesStatusRequest) {
    request = &DescribeServicesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServicesStatus")
    return
}

func NewDescribeServicesStatusResponse() (response *DescribeServicesStatusResponse) {
    response = &DescribeServicesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务列表
func (c *Client) DescribeServicesStatus(request *DescribeServicesStatusRequest) (response *DescribeServicesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeServicesStatusRequest()
    }
    response = NewDescribeServicesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTcbScfApisRequest() (request *DescribeTcbScfApisRequest) {
    request = &DescribeTcbScfApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeTcbScfApis")
    return
}

func NewDescribeTcbScfApisResponse() (response *DescribeTcbScfApisResponse) {
    response = &DescribeTcbScfApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询TCB-SCF-HTTP触发器列表
func (c *Client) DescribeTcbScfApis(request *DescribeTcbScfApisRequest) (response *DescribeTcbScfApisResponse, err error) {
    if request == nil {
        request = NewDescribeTcbScfApisRequest()
    }
    response = NewDescribeTcbScfApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanRequest() (request *DescribeUsagePlanRequest) {
    request = &DescribeUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlan")
    return
}

func NewDescribeUsagePlanResponse() (response *DescribeUsagePlanResponse) {
    response = &DescribeUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示使用计划详情
func (c *Client) DescribeUsagePlan(request *DescribeUsagePlanRequest) (response *DescribeUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanRequest()
    }
    response = NewDescribeUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanEnvironmentsRequest() (request *DescribeUsagePlanEnvironmentsRequest) {
    request = &DescribeUsagePlanEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanEnvironments")
    return
}

func NewDescribeUsagePlanEnvironmentsResponse() (response *DescribeUsagePlanEnvironmentsResponse) {
    response = &DescribeUsagePlanEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示使用计划绑定详情
func (c *Client) DescribeUsagePlanEnvironments(request *DescribeUsagePlanEnvironmentsRequest) (response *DescribeUsagePlanEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanEnvironmentsRequest()
    }
    response = NewDescribeUsagePlanEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanSecretIdsRequest() (request *DescribeUsagePlanSecretIdsRequest) {
    request = &DescribeUsagePlanSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanSecretIds")
    return
}

func NewDescribeUsagePlanSecretIdsResponse() (response *DescribeUsagePlanSecretIdsResponse) {
    response = &DescribeUsagePlanSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示使用计划绑定的密钥
func (c *Client) DescribeUsagePlanSecretIds(request *DescribeUsagePlanSecretIdsRequest) (response *DescribeUsagePlanSecretIdsResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanSecretIdsRequest()
    }
    response = NewDescribeUsagePlanSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlansStatusRequest() (request *DescribeUsagePlansStatusRequest) {
    request = &DescribeUsagePlansStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlansStatus")
    return
}

func NewDescribeUsagePlansStatusResponse() (response *DescribeUsagePlansStatusResponse) {
    response = &DescribeUsagePlansStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示使用计划列表
func (c *Client) DescribeUsagePlansStatus(request *DescribeUsagePlansStatusRequest) (response *DescribeUsagePlansStatusResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlansStatusRequest()
    }
    response = NewDescribeUsagePlansStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisableApiKeyRequest() (request *DisableApiKeyRequest) {
    request = &DisableApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DisableApiKey")
    return
}

func NewDisableApiKeyResponse() (response *DisableApiKeyResponse) {
    response = &DisableApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 禁用密钥
func (c *Client) DisableApiKey(request *DisableApiKeyRequest) (response *DisableApiKeyResponse, err error) {
    if request == nil {
        request = NewDisableApiKeyRequest()
    }
    response = NewDisableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableApiKeyRequest() (request *EnableApiKeyRequest) {
    request = &EnableApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "EnableApiKey")
    return
}

func NewEnableApiKeyResponse() (response *EnableApiKeyResponse) {
    response = &EnableApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用密钥
func (c *Client) EnableApiKey(request *EnableApiKeyRequest) (response *EnableApiKeyResponse, err error) {
    if request == nil {
        request = NewEnableApiKeyRequest()
    }
    response = NewEnableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateApiDocumentRequest() (request *GenerateApiDocumentRequest) {
    request = &GenerateApiDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "GenerateApiDocument")
    return
}

func NewGenerateApiDocumentResponse() (response *GenerateApiDocumentResponse) {
    response = &GenerateApiDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 生成sdk和api文档
func (c *Client) GenerateApiDocument(request *GenerateApiDocumentRequest) (response *GenerateApiDocumentResponse, err error) {
    if request == nil {
        request = NewGenerateApiDocumentRequest()
    }
    response = NewGenerateApiDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewGetAccountSettingsRequest() (request *GetAccountSettingsRequest) {
    request = &GetAccountSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "GetAccountSettings")
    return
}

func NewGetAccountSettingsResponse() (response *GetAccountSettingsResponse) {
    response = &GetAccountSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户配额
func (c *Client) GetAccountSettings(request *GetAccountSettingsRequest) (response *GetAccountSettingsResponse, err error) {
    if request == nil {
        request = NewGetAccountSettingsRequest()
    }
    response = NewGetAccountSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewImportOpenApiRequest() (request *ImportOpenApiRequest) {
    request = &ImportOpenApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ImportOpenApi")
    return
}

func NewImportOpenApiResponse() (response *ImportOpenApiResponse) {
    response = &ImportOpenApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过OpenAPI规范导入已有的API到API网关。 
func (c *Client) ImportOpenApi(request *ImportOpenApiRequest) (response *ImportOpenApiResponse, err error) {
    if request == nil {
        request = NewImportOpenApiRequest()
    }
    response = NewImportOpenApiResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRequest() (request *InquiryPriceRequest) {
    request = &InquiryPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "InquiryPrice")
    return
}

func NewInquiryPriceResponse() (response *InquiryPriceResponse) {
    response = &InquiryPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建apigateway时，用于询价。
func (c *Client) InquiryPrice(request *InquiryPriceRequest) (response *InquiryPriceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRequest()
    }
    response = NewInquiryPriceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiRequest() (request *ModifyApiRequest) {
    request = &ModifyApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApi")
    return
}

func NewModifyApiResponse() (response *ModifyApiResponse) {
    response = &ModifyApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改api
func (c *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
    if request == nil {
        request = NewModifyApiRequest()
    }
    response = NewModifyApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiEnvironmentStrategyRequest() (request *ModifyApiEnvironmentStrategyRequest) {
    request = &ModifyApiEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiEnvironmentStrategy")
    return
}

func NewModifyApiEnvironmentStrategyResponse() (response *ModifyApiEnvironmentStrategyResponse) {
    response = &ModifyApiEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改api限速
func (c *Client) ModifyApiEnvironmentStrategy(request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyApiEnvironmentStrategyRequest()
    }
    response = NewModifyApiEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiIncrementRequest() (request *ModifyApiIncrementRequest) {
    request = &ModifyApiIncrementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiIncrement")
    return
}

func NewModifyApiIncrementResponse() (response *ModifyApiIncrementResponse) {
    response = &ModifyApiIncrementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供增量更新API能力，主要是给程序调用（区别于ModifyApi，该接口是需要传入API的全量参数，对console使用较友好）
func (c *Client) ModifyApiIncrement(request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    if request == nil {
        request = NewModifyApiIncrementRequest()
    }
    response = NewModifyApiIncrementResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApisAuthRelationApiRequest() (request *ModifyApisAuthRelationApiRequest) {
    request = &ModifyApisAuthRelationApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApisAuthRelationApi")
    return
}

func NewModifyApisAuthRelationApiResponse() (response *ModifyApisAuthRelationApiResponse) {
    response = &ModifyApisAuthRelationApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改关联授权API
func (c *Client) ModifyApisAuthRelationApi(request *ModifyApisAuthRelationApiRequest) (response *ModifyApisAuthRelationApiResponse, err error) {
    if request == nil {
        request = NewModifyApisAuthRelationApiRequest()
    }
    response = NewModifyApisAuthRelationApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIPStrategyRequest() (request *ModifyIPStrategyRequest) {
    request = &ModifyIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyIPStrategy")
    return
}

func NewModifyIPStrategyResponse() (response *ModifyIPStrategyResponse) {
    response = &ModifyIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改ip策略
func (c *Client) ModifyIPStrategy(request *ModifyIPStrategyRequest) (response *ModifyIPStrategyResponse, err error) {
    if request == nil {
        request = NewModifyIPStrategyRequest()
    }
    response = NewModifyIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogRuleRequest() (request *ModifyLogRuleRequest) {
    request = &ModifyLogRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyLogRule")
    return
}

func NewModifyLogRuleResponse() (response *ModifyLogRuleResponse) {
    response = &ModifyLogRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改日志规则
func (c *Client) ModifyLogRule(request *ModifyLogRuleRequest) (response *ModifyLogRuleResponse, err error) {
    if request == nil {
        request = NewModifyLogRuleRequest()
    }
    response = NewModifyLogRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceRequest() (request *ModifyServiceRequest) {
    request = &ModifyServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyService")
    return
}

func NewModifyServiceResponse() (response *ModifyServiceResponse) {
    response = &ModifyServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于修改服务
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
    if request == nil {
        request = NewModifyServiceRequest()
    }
    response = NewModifyServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceEnvironmentKeyMonitorUploadRequest() (request *ModifyServiceEnvironmentKeyMonitorUploadRequest) {
    request = &ModifyServiceEnvironmentKeyMonitorUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyServiceEnvironmentKeyMonitorUpload")
    return
}

func NewModifyServiceEnvironmentKeyMonitorUploadResponse() (response *ModifyServiceEnvironmentKeyMonitorUploadResponse) {
    response = &ModifyServiceEnvironmentKeyMonitorUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置服务的环境是否进行key维度的监控数据上报
func (c *Client) ModifyServiceEnvironmentKeyMonitorUpload(request *ModifyServiceEnvironmentKeyMonitorUploadRequest) (response *ModifyServiceEnvironmentKeyMonitorUploadResponse, err error) {
    if request == nil {
        request = NewModifyServiceEnvironmentKeyMonitorUploadRequest()
    }
    response = NewModifyServiceEnvironmentKeyMonitorUploadResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceEnvironmentStrategyRequest() (request *ModifyServiceEnvironmentStrategyRequest) {
    request = &ModifyServiceEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyServiceEnvironmentStrategy")
    return
}

func NewModifyServiceEnvironmentStrategyResponse() (response *ModifyServiceEnvironmentStrategyResponse) {
    response = &ModifyServiceEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改服务环境限速
func (c *Client) ModifyServiceEnvironmentStrategy(request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyServiceEnvironmentStrategyRequest()
    }
    response = NewModifyServiceEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubDomainRequest() (request *ModifySubDomainRequest) {
    request = &ModifySubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifySubDomain")
    return
}

func NewModifySubDomainResponse() (response *ModifySubDomainResponse) {
    response = &ModifySubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务修改自定义域名
func (c *Client) ModifySubDomain(request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
    if request == nil {
        request = NewModifySubDomainRequest()
    }
    response = NewModifySubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUsagePlanRequest() (request *ModifyUsagePlanRequest) {
    request = &ModifyUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyUsagePlan")
    return
}

func NewModifyUsagePlanResponse() (response *ModifyUsagePlanResponse) {
    response = &ModifyUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改使用计划
func (c *Client) ModifyUsagePlan(request *ModifyUsagePlanRequest) (response *ModifyUsagePlanResponse, err error) {
    if request == nil {
        request = NewModifyUsagePlanRequest()
    }
    response = NewModifyUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewMonitorTopsRequest() (request *MonitorTopsRequest) {
    request = &MonitorTopsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "MonitorTops")
    return
}

func NewMonitorTopsResponse() (response *MonitorTopsResponse) {
    response = &MonitorTopsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询监控top数据
func (c *Client) MonitorTops(request *MonitorTopsRequest) (response *MonitorTopsResponse, err error) {
    if request == nil {
        request = NewMonitorTopsRequest()
    }
    response = NewMonitorTopsResponse()
    err = c.Send(request, response)
    return
}

func NewMonitorTrendRequest() (request *MonitorTrendRequest) {
    request = &MonitorTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "MonitorTrend")
    return
}

func NewMonitorTrendResponse() (response *MonitorTrendResponse) {
    response = &MonitorTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取监控指标趋势，粒度为分钟
func (c *Client) MonitorTrend(request *MonitorTrendRequest) (response *MonitorTrendResponse, err error) {
    if request == nil {
        request = NewMonitorTrendRequest()
    }
    response = NewMonitorTrendResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseServiceRequest() (request *ReleaseServiceRequest) {
    request = &ReleaseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ReleaseService")
    return
}

func NewReleaseServiceResponse() (response *ReleaseServiceResponse) {
    response = &ReleaseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户发布服务 
func (c *Client) ReleaseService(request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
    if request == nil {
        request = NewReleaseServiceRequest()
    }
    response = NewReleaseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewRunApiRequest() (request *RunApiRequest) {
    request = &RunApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "RunApi")
    return
}

func NewRunApiResponse() (response *RunApiResponse) {
    response = &RunApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 调试运行API
func (c *Client) RunApi(request *RunApiRequest) (response *RunApiResponse, err error) {
    if request == nil {
        request = NewRunApiRequest()
    }
    response = NewRunApiResponse()
    err = c.Send(request, response)
    return
}

func NewRunApiForMarketRequest() (request *RunApiForMarketRequest) {
    request = &RunApiForMarketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "RunApiForMarket")
    return
}

func NewRunApiForMarketResponse() (response *RunApiForMarketResponse) {
    response = &RunApiForMarketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运行API（云市场专用）
func (c *Client) RunApiForMarket(request *RunApiForMarketRequest) (response *RunApiForMarketResponse, err error) {
    if request == nil {
        request = NewRunApiForMarketRequest()
    }
    response = NewRunApiForMarketResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindEnvironmentRequest() (request *UnBindEnvironmentRequest) {
    request = &UnBindEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindEnvironment")
    return
}

func NewUnBindEnvironmentResponse() (response *UnBindEnvironmentResponse) {
    response = &UnBindEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑环境
func (c *Client) UnBindEnvironment(request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
    if request == nil {
        request = NewUnBindEnvironmentRequest()
    }
    response = NewUnBindEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindIPStrategyRequest() (request *UnBindIPStrategyRequest) {
    request = &UnBindIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindIPStrategy")
    return
}

func NewUnBindIPStrategyResponse() (response *UnBindIPStrategyResponse) {
    response = &UnBindIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑ip策略
func (c *Client) UnBindIPStrategy(request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
    if request == nil {
        request = NewUnBindIPStrategyRequest()
    }
    response = NewUnBindIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindSecretIdsRequest() (request *UnBindSecretIdsRequest) {
    request = &UnBindSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSecretIds")
    return
}

func NewUnBindSecretIdsResponse() (response *UnBindSecretIdsResponse) {
    response = &UnBindSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用计划解绑密钥
func (c *Client) UnBindSecretIds(request *UnBindSecretIdsRequest) (response *UnBindSecretIdsResponse, err error) {
    if request == nil {
        request = NewUnBindSecretIdsRequest()
    }
    response = NewUnBindSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindSubDomainRequest() (request *UnBindSubDomainRequest) {
    request = &UnBindSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSubDomain")
    return
}

func NewUnBindSubDomainResponse() (response *UnBindSubDomainResponse) {
    response = &UnBindSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑自定义域名，解绑之前需要确认该域名存在且成功绑定。
func (c *Client) UnBindSubDomain(request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
    if request == nil {
        request = NewUnBindSubDomainRequest()
    }
    response = NewUnBindSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUnReleaseServiceRequest() (request *UnReleaseServiceRequest) {
    request = &UnReleaseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnReleaseService")
    return
}

func NewUnReleaseServiceResponse() (response *UnReleaseServiceResponse) {
    response = &UnReleaseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于下线服务
func (c *Client) UnReleaseService(request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
    if request == nil {
        request = NewUnReleaseServiceRequest()
    }
    response = NewUnReleaseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiKeyRequest() (request *UpdateApiKeyRequest) {
    request = &UpdateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateApiKey")
    return
}

func NewUpdateApiKeyResponse() (response *UpdateApiKeyResponse) {
    response = &UpdateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新密钥
func (c *Client) UpdateApiKey(request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApiKeyRequest()
    }
    response = NewUpdateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServiceRequest() (request *UpdateServiceRequest) {
    request = &UpdateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateService")
    return
}

func NewUpdateServiceResponse() (response *UpdateServiceResponse) {
    response = &UpdateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于切换服务上线的版本
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    if request == nil {
        request = NewUpdateServiceRequest()
    }
    response = NewUpdateServiceResponse()
    err = c.Send(request, response)
    return
}
