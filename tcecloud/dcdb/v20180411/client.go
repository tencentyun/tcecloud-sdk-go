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

package v20180411

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-11"

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


func NewActiveDedicatedDBInstanceRequest() (request *ActiveDedicatedDBInstanceRequest) {
    request = &ActiveDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ActiveDedicatedDBInstance")
    return
}

func NewActiveDedicatedDBInstanceResponse() (response *ActiveDedicatedDBInstanceResponse) {
    response = &ActiveDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ActiveDedicatedDBInstance）用于恢复已隔离的独享分布式数据库实例。
func (c *Client) ActiveDedicatedDBInstance(request *ActiveDedicatedDBInstanceRequest) (response *ActiveDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewActiveDedicatedDBInstanceRequest()
    }
    response = NewActiveDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewActiveHourDCDBInstanceRequest() (request *ActiveHourDCDBInstanceRequest) {
    request = &ActiveHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ActiveHourDCDBInstance")
    return
}

func NewActiveHourDCDBInstanceResponse() (response *ActiveHourDCDBInstanceResponse) {
    response = &ActiveHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ActiveHourDCDBInstance）用于恢复按量计费DCDB实例。
func (c *Client) ActiveHourDCDBInstance(request *ActiveHourDCDBInstanceRequest) (response *ActiveHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewActiveHourDCDBInstanceRequest()
    }
    response = NewActiveHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewAuthenticateCAMRequest() (request *AuthenticateCAMRequest) {
    request = &AuthenticateCAMRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "AuthenticateCAM")
    return
}

func NewAuthenticateCAMResponse() (response *AuthenticateCAMResponse) {
    response = &AuthenticateCAMResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AuthenticateCAM）为控制台提供预鉴权
func (c *Client) AuthenticateCAM(request *AuthenticateCAMRequest) (response *AuthenticateCAMResponse, err error) {
    if request == nil {
        request = NewAuthenticateCAMRequest()
    }
    response = NewAuthenticateCAMResponse()
    err = c.Send(request, response)
    return
}

func NewCheckIpStatusRequest() (request *CheckIpStatusRequest) {
    request = &CheckIpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CheckIpStatus")
    return
}

func NewCheckIpStatusResponse() (response *CheckIpStatusResponse) {
    response = &CheckIpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CheckIpStatus)用于查询指定的私有网络中的虚拟IP是否可用。
func (c *Client) CheckIpStatus(request *CheckIpStatusRequest) (response *CheckIpStatusResponse, err error) {
    if request == nil {
        request = NewCheckIpStatusRequest()
    }
    response = NewCheckIpStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCloneAccountRequest() (request *CloneAccountRequest) {
    request = &CloneAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CloneAccount")
    return
}

func NewCloneAccountResponse() (response *CloneAccountResponse) {
    response = &CloneAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CloneAccount）用于克隆实例账户。
func (c *Client) CloneAccount(request *CloneAccountRequest) (response *CloneAccountResponse, err error) {
    if request == nil {
        request = NewCloneAccountRequest()
    }
    response = NewCloneAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCloseDBExtranetAccessRequest() (request *CloseDBExtranetAccessRequest) {
    request = &CloseDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CloseDBExtranetAccess")
    return
}

func NewCloseDBExtranetAccessResponse() (response *CloseDBExtranetAccessResponse) {
    response = &CloseDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CloseDBExtranetAccess)用于关闭云数据库实例的外网访问。关闭外网访问后，外网地址将不可访问，查询实例列表接口将不返回对应实例的外网域名和端口信息。
func (c *Client) CloseDBExtranetAccess(request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseDBExtranetAccessRequest()
    }
    response = NewCloseDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCopyAccountPrivilegesRequest() (request *CopyAccountPrivilegesRequest) {
    request = &CopyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CopyAccountPrivileges")
    return
}

func NewCopyAccountPrivilegesResponse() (response *CopyAccountPrivilegesResponse) {
    response = &CopyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CopyAccountPrivileges）用于复制云数据库账号的权限。
// 注意：相同用户名，不同Host是不同的账号，Readonly属性相同的账号之间才能复制权限。
func (c *Client) CopyAccountPrivileges(request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewCopyAccountPrivilegesRequest()
    }
    response = NewCopyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateAccount")
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateAccount）用于创建云数据库账号。一个实例可以创建多个不同的账号，相同的用户名+不同的host是不同的账号。
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDCDBInstanceRequest() (request *CreateDCDBInstanceRequest) {
    request = &CreateDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateDCDBInstance")
    return
}

func NewCreateDCDBInstanceResponse() (response *CreateDCDBInstanceResponse) {
    response = &CreateDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateDCDBInstance）用于创建包年包月的云数据库实例，可通过传入实例规格、数据库版本号、购买时长等信息创建云数据库实例。
func (c *Client) CreateDCDBInstance(request *CreateDCDBInstanceRequest) (response *CreateDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDCDBInstanceRequest()
    }
    response = NewCreateDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHourDCDBInstanceRequest() (request *CreateHourDCDBInstanceRequest) {
    request = &CreateHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateHourDCDBInstance")
    return
}

func NewCreateHourDCDBInstanceResponse() (response *CreateHourDCDBInstanceResponse) {
    response = &CreateHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateHourDCDBInstance）用于创建按量计费的DCDB实例。
func (c *Client) CreateHourDCDBInstance(request *CreateHourDCDBInstanceRequest) (response *CreateHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateHourDCDBInstanceRequest()
    }
    response = NewCreateHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTmpDCDBInstanceRequest() (request *CreateTmpDCDBInstanceRequest) {
    request = &CreateTmpDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateTmpDCDBInstance")
    return
}

func NewCreateTmpDCDBInstanceResponse() (response *CreateTmpDCDBInstanceResponse) {
    response = &CreateTmpDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateTmpDCDBInstance）用于创建临时实例。
func (c *Client) CreateTmpDCDBInstance(request *CreateTmpDCDBInstanceRequest) (response *CreateTmpDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateTmpDCDBInstanceRequest()
    }
    response = NewCreateTmpDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DeleteAccount")
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteAccount）用于删除云数据库账号。用户名+host唯一确定一个账号。
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTmpInstanceRequest() (request *DeleteTmpInstanceRequest) {
    request = &DeleteTmpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DeleteTmpInstance")
    return
}

func NewDeleteTmpInstanceResponse() (response *DeleteTmpInstanceResponse) {
    response = &DeleteTmpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteTmpInstance）用于删除临时实例。
func (c *Client) DeleteTmpInstance(request *DeleteTmpInstanceRequest) (response *DeleteTmpInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteTmpInstanceRequest()
    }
    response = NewDeleteTmpInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAccountPrivileges")
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAccountPrivileges）用于查询云数据库账号权限。
// 注意：注意：相同用户名，不同Host是不同的账号。
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    response = NewDescribeAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAccounts）用于查询指定云数据库实例的账号列表。
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogsRequest() (request *DescribeAuditLogsRequest) {
    request = &DescribeAuditLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAuditLogs")
    return
}

func NewDescribeAuditLogsResponse() (response *DescribeAuditLogsResponse) {
    response = &DescribeAuditLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAuditLogs）用于查询审计日志列表。
func (c *Client) DescribeAuditLogs(request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogsRequest()
    }
    response = NewDescribeAuditLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRuleDetailRequest() (request *DescribeAuditRuleDetailRequest) {
    request = &DescribeAuditRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAuditRuleDetail")
    return
}

func NewDescribeAuditRuleDetailResponse() (response *DescribeAuditRuleDetailResponse) {
    response = &DescribeAuditRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAuditRuleDetail）用于查询审计规则详情。
func (c *Client) DescribeAuditRuleDetail(request *DescribeAuditRuleDetailRequest) (response *DescribeAuditRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleDetailRequest()
    }
    response = NewDescribeAuditRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRulesRequest() (request *DescribeAuditRulesRequest) {
    request = &DescribeAuditRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAuditRules")
    return
}

func NewDescribeAuditRulesResponse() (response *DescribeAuditRulesResponse) {
    response = &DescribeAuditRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAuditRules）用于查询审计规则列表。
func (c *Client) DescribeAuditRules(request *DescribeAuditRulesRequest) (response *DescribeAuditRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRulesRequest()
    }
    response = NewDescribeAuditRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditStrategiesRequest() (request *DescribeAuditStrategiesRequest) {
    request = &DescribeAuditStrategiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAuditStrategies")
    return
}

func NewDescribeAuditStrategiesResponse() (response *DescribeAuditStrategiesResponse) {
    response = &DescribeAuditStrategiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAuditStrategies）用于查询某一个实例的审计策略列表。
func (c *Client) DescribeAuditStrategies(request *DescribeAuditStrategiesRequest) (response *DescribeAuditStrategiesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditStrategiesRequest()
    }
    response = NewDescribeAuditStrategiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableExclusiveGroupsRequest() (request *DescribeAvailableExclusiveGroupsRequest) {
    request = &DescribeAvailableExclusiveGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAvailableExclusiveGroups")
    return
}

func NewDescribeAvailableExclusiveGroupsResponse() (response *DescribeAvailableExclusiveGroupsResponse) {
    response = &DescribeAvailableExclusiveGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeAvailableExclusiveGroups)用于拉取独享资源池信息
func (c *Client) DescribeAvailableExclusiveGroups(request *DescribeAvailableExclusiveGroupsRequest) (response *DescribeAvailableExclusiveGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableExclusiveGroupsRequest()
    }
    response = NewDescribeAvailableExclusiveGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchDCDBRenewalPriceRequest() (request *DescribeBatchDCDBRenewalPriceRequest) {
    request = &DescribeBatchDCDBRenewalPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeBatchDCDBRenewalPrice")
    return
}

func NewDescribeBatchDCDBRenewalPriceResponse() (response *DescribeBatchDCDBRenewalPriceResponse) {
    response = &DescribeBatchDCDBRenewalPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeBatchDCDBRenewalPrice）用于在续费分布式数据库实例时，批量查询续费的价格。
func (c *Client) DescribeBatchDCDBRenewalPrice(request *DescribeBatchDCDBRenewalPriceRequest) (response *DescribeBatchDCDBRenewalPriceResponse, err error) {
    if request == nil {
        request = NewDescribeBatchDCDBRenewalPriceRequest()
    }
    response = NewDescribeBatchDCDBRenewalPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDetailMetricsRequest() (request *DescribeDBDetailMetricsRequest) {
    request = &DescribeDBDetailMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBDetailMetrics")
    return
}

func NewDescribeDBDetailMetricsResponse() (response *DescribeDBDetailMetricsResponse) {
    response = &DescribeDBDetailMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBDetailMetrics)用于查询云数据库主备详细监控指标。
func (c *Client) DescribeDBDetailMetrics(request *DescribeDBDetailMetricsRequest) (response *DescribeDBDetailMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeDBDetailMetricsRequest()
    }
    response = NewDescribeDBDetailMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBEnginesRequest() (request *DescribeDBEnginesRequest) {
    request = &DescribeDBEnginesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBEngines")
    return
}

func NewDescribeDBEnginesResponse() (response *DescribeDBEnginesResponse) {
    response = &DescribeDBEnginesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于在调用创建接口前查询当前支持的数据库引擎列表
func (c *Client) DescribeDBEngines(request *DescribeDBEnginesRequest) (response *DescribeDBEnginesResponse, err error) {
    if request == nil {
        request = NewDescribeDBEnginesRequest()
    }
    response = NewDescribeDBEnginesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceRsipRequest() (request *DescribeDBInstanceRsipRequest) {
    request = &DescribeDBInstanceRsipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBInstanceRsip")
    return
}

func NewDescribeDBInstanceRsipResponse() (response *DescribeDBInstanceRsipResponse) {
    response = &DescribeDBInstanceRsipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBInstanceRsip）用于获取实例Rsip
func (c *Client) DescribeDBInstanceRsip(request *DescribeDBInstanceRsipRequest) (response *DescribeDBInstanceRsipResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceRsipRequest()
    }
    response = NewDescribeDBInstanceRsipResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
    request = &DescribeDBLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBLogFiles")
    return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
    response = &DescribeDBLogFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBLogFiles)用于获取数据库的各种日志列表，包括冷备、binlog、errlog和slowlog。
func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDBLogFilesRequest()
    }
    response = NewDescribeDBLogFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBMetricsRequest() (request *DescribeDBMetricsRequest) {
    request = &DescribeDBMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBMetrics")
    return
}

func NewDescribeDBMetricsResponse() (response *DescribeDBMetricsResponse) {
    response = &DescribeDBMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBMetrics)用于查询云数据库监控指标。
func (c *Client) DescribeDBMetrics(request *DescribeDBMetricsRequest) (response *DescribeDBMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeDBMetricsRequest()
    }
    response = NewDescribeDBMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBParametersRequest() (request *DescribeDBParametersRequest) {
    request = &DescribeDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBParameters")
    return
}

func NewDescribeDBParametersResponse() (response *DescribeDBParametersResponse) {
    response = &DescribeDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBParameters)用于获取数据库的当前参数设置。
func (c *Client) DescribeDBParameters(request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDBParametersRequest()
    }
    response = NewDescribeDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSecurityGroups")
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBSecurityGroups)用于查询实例的安全组详情。
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSlowLogAnalysisRequest() (request *DescribeDBSlowLogAnalysisRequest) {
    request = &DescribeDBSlowLogAnalysisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSlowLogAnalysis")
    return
}

func NewDescribeDBSlowLogAnalysisResponse() (response *DescribeDBSlowLogAnalysisResponse) {
    response = &DescribeDBSlowLogAnalysisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBSlowLogAnalysis)用于获取慢查询记录详情。
func (c *Client) DescribeDBSlowLogAnalysis(request *DescribeDBSlowLogAnalysisRequest) (response *DescribeDBSlowLogAnalysisResponse, err error) {
    if request == nil {
        request = NewDescribeDBSlowLogAnalysisRequest()
    }
    response = NewDescribeDBSlowLogAnalysisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSlowLogsRequest() (request *DescribeDBSlowLogsRequest) {
    request = &DescribeDBSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSlowLogs")
    return
}

func NewDescribeDBSlowLogsResponse() (response *DescribeDBSlowLogsResponse) {
    response = &DescribeDBSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBSlowLogs)用于查询慢查询日志列表。
func (c *Client) DescribeDBSlowLogs(request *DescribeDBSlowLogsRequest) (response *DescribeDBSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSlowLogsRequest()
    }
    response = NewDescribeDBSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSyncModeRequest() (request *DescribeDBSyncModeRequest) {
    request = &DescribeDBSyncModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSyncMode")
    return
}

func NewDescribeDBSyncModeResponse() (response *DescribeDBSyncModeResponse) {
    response = &DescribeDBSyncModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBSyncMode）用于查询云数据库实例的同步模式。
func (c *Client) DescribeDBSyncMode(request *DescribeDBSyncModeRequest) (response *DescribeDBSyncModeResponse, err error) {
    if request == nil {
        request = NewDescribeDBSyncModeRequest()
    }
    response = NewDescribeDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBTmpInstancesRequest() (request *DescribeDBTmpInstancesRequest) {
    request = &DescribeDBTmpInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBTmpInstances")
    return
}

func NewDescribeDBTmpInstancesResponse() (response *DescribeDBTmpInstancesResponse) {
    response = &DescribeDBTmpInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBTmpInstances）用于获取实例回档生成的临时实例
func (c *Client) DescribeDBTmpInstances(request *DescribeDBTmpInstancesRequest) (response *DescribeDBTmpInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBTmpInstancesRequest()
    }
    response = NewDescribeDBTmpInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBBinlogTimeRequest() (request *DescribeDCDBBinlogTimeRequest) {
    request = &DescribeDCDBBinlogTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBBinlogTime")
    return
}

func NewDescribeDCDBBinlogTimeResponse() (response *DescribeDCDBBinlogTimeResponse) {
    response = &DescribeDCDBBinlogTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBBinlogTime）用于查询可回档时间范围。
func (c *Client) DescribeDCDBBinlogTime(request *DescribeDCDBBinlogTimeRequest) (response *DescribeDCDBBinlogTimeResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBBinlogTimeRequest()
    }
    response = NewDescribeDCDBBinlogTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBInstanceDetailRequest() (request *DescribeDCDBInstanceDetailRequest) {
    request = &DescribeDCDBInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstanceDetail")
    return
}

func NewDescribeDCDBInstanceDetailResponse() (response *DescribeDCDBInstanceDetailResponse) {
    response = &DescribeDCDBInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBInstanceDetail）用于获取DCDB实例详情
func (c *Client) DescribeDCDBInstanceDetail(request *DescribeDCDBInstanceDetailRequest) (response *DescribeDCDBInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstanceDetailRequest()
    }
    response = NewDescribeDCDBInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBInstancesRequest() (request *DescribeDCDBInstancesRequest) {
    request = &DescribeDCDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstances")
    return
}

func NewDescribeDCDBInstancesResponse() (response *DescribeDCDBInstancesResponse) {
    response = &DescribeDCDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询云数据库实例列表，支持通过项目ID、实例ID、内网地址、实例名称等来筛选实例。
// 如果不指定任何筛选条件，则默认返回10条实例记录，单次请求最多支持返回100条实例记录。
func (c *Client) DescribeDCDBInstances(request *DescribeDCDBInstancesRequest) (response *DescribeDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstancesRequest()
    }
    response = NewDescribeDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBPriceRequest() (request *DescribeDCDBPriceRequest) {
    request = &DescribeDCDBPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBPrice")
    return
}

func NewDescribeDCDBPriceResponse() (response *DescribeDCDBPriceResponse) {
    response = &DescribeDCDBPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBPrice）用于在购买实例前，查询实例的价格。
func (c *Client) DescribeDCDBPrice(request *DescribeDCDBPriceRequest) (response *DescribeDCDBPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBPriceRequest()
    }
    response = NewDescribeDCDBPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBRenewalPriceRequest() (request *DescribeDCDBRenewalPriceRequest) {
    request = &DescribeDCDBRenewalPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBRenewalPrice")
    return
}

func NewDescribeDCDBRenewalPriceResponse() (response *DescribeDCDBRenewalPriceResponse) {
    response = &DescribeDCDBRenewalPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBRenewalPrice）用于在续费分布式数据库实例时，查询续费的价格。
func (c *Client) DescribeDCDBRenewalPrice(request *DescribeDCDBRenewalPriceRequest) (response *DescribeDCDBRenewalPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBRenewalPriceRequest()
    }
    response = NewDescribeDCDBRenewalPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBSaleInfoRequest() (request *DescribeDCDBSaleInfoRequest) {
    request = &DescribeDCDBSaleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBSaleInfo")
    return
}

func NewDescribeDCDBSaleInfoResponse() (response *DescribeDCDBSaleInfoResponse) {
    response = &DescribeDCDBSaleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDCDBSaleInfo)用于查询分布式数据库可售卖的地域和可用区信息。
func (c *Client) DescribeDCDBSaleInfo(request *DescribeDCDBSaleInfoRequest) (response *DescribeDCDBSaleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBSaleInfoRequest()
    }
    response = NewDescribeDCDBSaleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBShardsRequest() (request *DescribeDCDBShardsRequest) {
    request = &DescribeDCDBShardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBShards")
    return
}

func NewDescribeDCDBShardsResponse() (response *DescribeDCDBShardsResponse) {
    response = &DescribeDCDBShardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBShards）用于查询云数据库实例的分片信息。
func (c *Client) DescribeDCDBShards(request *DescribeDCDBShardsRequest) (response *DescribeDCDBShardsResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBShardsRequest()
    }
    response = NewDescribeDCDBShardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBUpgradePriceRequest() (request *DescribeDCDBUpgradePriceRequest) {
    request = &DescribeDCDBUpgradePriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBUpgradePrice")
    return
}

func NewDescribeDCDBUpgradePriceResponse() (response *DescribeDCDBUpgradePriceResponse) {
    response = &DescribeDCDBUpgradePriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDCDBUpgradePrice）用于查询升级分布式数据库实例价格。
func (c *Client) DescribeDCDBUpgradePrice(request *DescribeDCDBUpgradePriceRequest) (response *DescribeDCDBUpgradePriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBUpgradePriceRequest()
    }
    response = NewDescribeDCDBUpgradePriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseObjectsRequest() (request *DescribeDatabaseObjectsRequest) {
    request = &DescribeDatabaseObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabaseObjects")
    return
}

func NewDescribeDatabaseObjectsResponse() (response *DescribeDatabaseObjectsResponse) {
    response = &DescribeDatabaseObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDatabaseObjects）用于查询云数据库实例的数据库中的对象列表，包含表、存储过程、视图和函数。
func (c *Client) DescribeDatabaseObjects(request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseObjectsRequest()
    }
    response = NewDescribeDatabaseObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseTableRequest() (request *DescribeDatabaseTableRequest) {
    request = &DescribeDatabaseTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabaseTable")
    return
}

func NewDescribeDatabaseTableResponse() (response *DescribeDatabaseTableResponse) {
    response = &DescribeDatabaseTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDatabaseTable）用于查询云数据库实例的表信息。
func (c *Client) DescribeDatabaseTable(request *DescribeDatabaseTableRequest) (response *DescribeDatabaseTableResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseTableRequest()
    }
    response = NewDescribeDatabaseTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabases")
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDatabases）用于查询云数据库实例的数据库列表。
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFenceShardSpecRequest() (request *DescribeFenceShardSpecRequest) {
    request = &DescribeFenceShardSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeFenceShardSpec")
    return
}

func NewDescribeFenceShardSpecResponse() (response *DescribeFenceShardSpecResponse) {
    response = &DescribeFenceShardSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeFenceShardSpec)用于查询可创建的独享分布式数据库实例的规格配置。
func (c *Client) DescribeFenceShardSpec(request *DescribeFenceShardSpecRequest) (response *DescribeFenceShardSpecResponse, err error) {
    if request == nil {
        request = NewDescribeFenceShardSpecRequest()
    }
    response = NewDescribeFenceShardSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeFlow")
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeFlow）用于查询流程状态
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstances）用于云监控拉取实例列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLatestCloudDBAReportRequest() (request *DescribeLatestCloudDBAReportRequest) {
    request = &DescribeLatestCloudDBAReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeLatestCloudDBAReport")
    return
}

func NewDescribeLatestCloudDBAReportResponse() (response *DescribeLatestCloudDBAReportResponse) {
    response = &DescribeLatestCloudDBAReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeLatestCloudDBAReport）用于获取最新的性能检测报告
func (c *Client) DescribeLatestCloudDBAReport(request *DescribeLatestCloudDBAReportRequest) (response *DescribeLatestCloudDBAReportResponse, err error) {
    if request == nil {
        request = NewDescribeLatestCloudDBAReportRequest()
    }
    response = NewDescribeLatestCloudDBAReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogFileRetentionPeriodRequest() (request *DescribeLogFileRetentionPeriodRequest) {
    request = &DescribeLogFileRetentionPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeLogFileRetentionPeriod")
    return
}

func NewDescribeLogFileRetentionPeriodResponse() (response *DescribeLogFileRetentionPeriodResponse) {
    response = &DescribeLogFileRetentionPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeLogFileRetentionPeriod)用于查看数据库备份日志的备份天数的设置情况。
func (c *Client) DescribeLogFileRetentionPeriod(request *DescribeLogFileRetentionPeriodRequest) (response *DescribeLogFileRetentionPeriodResponse, err error) {
    if request == nil {
        request = NewDescribeLogFileRetentionPeriodRequest()
    }
    response = NewDescribeLogFileRetentionPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeOrders")
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeOrders）用于查询分布式数据库订单信息。传入订单Id来查询订单关联的分布式数据库实例，和对应的任务流程ID。
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeProjectSecurityGroups")
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeProjectSecurityGroups)用于查询项目的安全组详情。
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeProjects")
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeProjects）用于查询项目列表
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShardSpecRequest() (request *DescribeShardSpecRequest) {
    request = &DescribeShardSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeShardSpec")
    return
}

func NewDescribeShardSpecResponse() (response *DescribeShardSpecResponse) {
    response = &DescribeShardSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询可创建的分布式数据库可售卖的分片规格配置。
func (c *Client) DescribeShardSpec(request *DescribeShardSpecRequest) (response *DescribeShardSpecResponse, err error) {
    if request == nil {
        request = NewDescribeShardSpecRequest()
    }
    response = NewDescribeShardSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSqlLogsRequest() (request *DescribeSqlLogsRequest) {
    request = &DescribeSqlLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeSqlLogs")
    return
}

func NewDescribeSqlLogsResponse() (response *DescribeSqlLogsResponse) {
    response = &DescribeSqlLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSqlLogs）用于获取实例SQL日志。
func (c *Client) DescribeSqlLogs(request *DescribeSqlLogsRequest) (response *DescribeSqlLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSqlLogsRequest()
    }
    response = NewDescribeSqlLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyHourDCDBInstanceRequest() (request *DestroyHourDCDBInstanceRequest) {
    request = &DestroyHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DestroyHourDCDBInstance")
    return
}

func NewDestroyHourDCDBInstanceResponse() (response *DestroyHourDCDBInstanceResponse) {
    response = &DestroyHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DestroyHourDCDBInstance）用于销毁按量计费实例。
func (c *Client) DestroyHourDCDBInstance(request *DestroyHourDCDBInstanceRequest) (response *DestroyHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyHourDCDBInstanceRequest()
    }
    response = NewDestroyHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewGrantAccountPrivilegesRequest() (request *GrantAccountPrivilegesRequest) {
    request = &GrantAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "GrantAccountPrivileges")
    return
}

func NewGrantAccountPrivilegesResponse() (response *GrantAccountPrivilegesResponse) {
    response = &GrantAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GrantAccountPrivileges）用于给云数据库账号赋权。
// 注意：相同用户名，不同Host是不同的账号。
func (c *Client) GrantAccountPrivileges(request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewGrantAccountPrivilegesRequest()
    }
    response = NewGrantAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewInitDCDBInstancesRequest() (request *InitDCDBInstancesRequest) {
    request = &InitDCDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "InitDCDBInstances")
    return
}

func NewInitDCDBInstancesResponse() (response *InitDCDBInstancesResponse) {
    response = &InitDCDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(InitDCDBInstances)用于初始化云数据库实例，包括设置默认字符集、表名大小写敏感等。
func (c *Client) InitDCDBInstances(request *InitDCDBInstancesRequest) (response *InitDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDCDBInstancesRequest()
    }
    response = NewInitDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDedicatedDBInstanceRequest() (request *IsolateDedicatedDBInstanceRequest) {
    request = &IsolateDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "IsolateDedicatedDBInstance")
    return
}

func NewIsolateDedicatedDBInstanceResponse() (response *IsolateDedicatedDBInstanceResponse) {
    response = &IsolateDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（IsolateDedicatedDBInstance）用于隔离独享云数据库实例。
func (c *Client) IsolateDedicatedDBInstance(request *IsolateDedicatedDBInstanceRequest) (response *IsolateDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDedicatedDBInstanceRequest()
    }
    response = NewIsolateDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateHourDCDBInstanceRequest() (request *IsolateHourDCDBInstanceRequest) {
    request = &IsolateHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "IsolateHourDCDBInstance")
    return
}

func NewIsolateHourDCDBInstanceResponse() (response *IsolateHourDCDBInstanceResponse) {
    response = &IsolateHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（IsolateHourDCDBInstance）用于隔离按量计费DCDB实例。
func (c *Client) IsolateHourDCDBInstance(request *IsolateHourDCDBInstanceRequest) (response *IsolateHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateHourDCDBInstanceRequest()
    }
    response = NewIsolateHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyAccountDescription")
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyAccountDescription）用于修改云数据库账号备注。
// 注意：相同用户名，不同Host是不同的账号。
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoRenewFlagRequest() (request *ModifyAutoRenewFlagRequest) {
    request = &ModifyAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyAutoRenewFlag")
    return
}

func NewModifyAutoRenewFlagResponse() (response *ModifyAutoRenewFlagResponse) {
    response = &ModifyAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyAutoRenewFlag）用于修改自动续费标记。
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    response = NewModifyAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstanceName")
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstanceName）用于修改实例名字
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyDBInstanceSecurityGroups)用于修改实例绑定的安全组
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstancesProjectRequest() (request *ModifyDBInstancesProjectRequest) {
    request = &ModifyDBInstancesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstancesProject")
    return
}

func NewModifyDBInstancesProjectResponse() (response *ModifyDBInstancesProjectResponse) {
    response = &ModifyDBInstancesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstancesProject）用于修改云数据库实例所属项目。
func (c *Client) ModifyDBInstancesProject(request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstancesProjectRequest()
    }
    response = NewModifyDBInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBParametersRequest() (request *ModifyDBParametersRequest) {
    request = &ModifyDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBParameters")
    return
}

func NewModifyDBParametersResponse() (response *ModifyDBParametersResponse) {
    response = &ModifyDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyDBParameters)用于修改数据库参数。
func (c *Client) ModifyDBParameters(request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBParametersRequest()
    }
    response = NewModifyDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBSyncModeRequest() (request *ModifyDBSyncModeRequest) {
    request = &ModifyDBSyncModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBSyncMode")
    return
}

func NewModifyDBSyncModeResponse() (response *ModifyDBSyncModeResponse) {
    response = &ModifyDBSyncModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBSyncMode）用于修改云数据库实例的同步模式。
func (c *Client) ModifyDBSyncMode(request *ModifyDBSyncModeRequest) (response *ModifyDBSyncModeResponse, err error) {
    if request == nil {
        request = NewModifyDBSyncModeRequest()
    }
    response = NewModifyDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNetworkRequest() (request *ModifyInstanceNetworkRequest) {
    request = &ModifyInstanceNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceNetwork")
    return
}

func NewModifyInstanceNetworkResponse() (response *ModifyInstanceNetworkResponse) {
    response = &ModifyInstanceNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyInstanceNetwork）用于修改实例所属网络。
func (c *Client) ModifyInstanceNetwork(request *ModifyInstanceNetworkRequest) (response *ModifyInstanceNetworkResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNetworkRequest()
    }
    response = NewModifyInstanceNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRemarkRequest() (request *ModifyInstanceRemarkRequest) {
    request = &ModifyInstanceRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceRemark")
    return
}

func NewModifyInstanceRemarkResponse() (response *ModifyInstanceRemarkResponse) {
    response = &ModifyInstanceRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyInstanceRemark）用于修改实例备注。
func (c *Client) ModifyInstanceRemark(request *ModifyInstanceRemarkRequest) (response *ModifyInstanceRemarkResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRemarkRequest()
    }
    response = NewModifyInstanceRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceVipRequest() (request *ModifyInstanceVipRequest) {
    request = &ModifyInstanceVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceVip")
    return
}

func NewModifyInstanceVipResponse() (response *ModifyInstanceVipResponse) {
    response = &ModifyInstanceVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyInstanceVip）用于修改实例Vip
func (c *Client) ModifyInstanceVip(request *ModifyInstanceVipRequest) (response *ModifyInstanceVipResponse, err error) {
    if request == nil {
        request = NewModifyInstanceVipRequest()
    }
    response = NewModifyInstanceVipResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogFileRetentionPeriodRequest() (request *ModifyLogFileRetentionPeriodRequest) {
    request = &ModifyLogFileRetentionPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyLogFileRetentionPeriod")
    return
}

func NewModifyLogFileRetentionPeriodResponse() (response *ModifyLogFileRetentionPeriodResponse) {
    response = &ModifyLogFileRetentionPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyLogFileRetentionPeriod)用于修改数据库备份日志保存天数。
func (c *Client) ModifyLogFileRetentionPeriod(request *ModifyLogFileRetentionPeriodRequest) (response *ModifyLogFileRetentionPeriodResponse, err error) {
    if request == nil {
        request = NewModifyLogFileRetentionPeriodRequest()
    }
    response = NewModifyLogFileRetentionPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewOpenDBExtranetAccessRequest() (request *OpenDBExtranetAccessRequest) {
    request = &OpenDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "OpenDBExtranetAccess")
    return
}

func NewOpenDBExtranetAccessResponse() (response *OpenDBExtranetAccessResponse) {
    response = &OpenDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（OpenDBExtranetAccess）用于开通云数据库实例的外网访问。开通外网访问后，您可通过外网域名和端口访问实例，可使用查询实例列表接口获取外网域名和端口信息。
func (c *Client) OpenDBExtranetAccess(request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenDBExtranetAccessRequest()
    }
    response = NewOpenDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDCDBInstanceRequest() (request *RenewDCDBInstanceRequest) {
    request = &RenewDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "RenewDCDBInstance")
    return
}

func NewRenewDCDBInstanceResponse() (response *RenewDCDBInstanceResponse) {
    response = &RenewDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RenewDCDBInstance）用于续费分布式数据库实例。
func (c *Client) RenewDCDBInstance(request *RenewDCDBInstanceRequest) (response *RenewDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewDCDBInstanceRequest()
    }
    response = NewRenewDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ResetAccountPassword")
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ResetAccountPassword）用于重置云数据库账号的密码。
// 注意：相同用户名，不同Host是不同的账号。
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewStartSmartDBARequest() (request *StartSmartDBARequest) {
    request = &StartSmartDBARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "StartSmartDBA")
    return
}

func NewStartSmartDBAResponse() (response *StartSmartDBAResponse) {
    response = &StartSmartDBAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartSmartDBA）用于启动性能检测任务。
func (c *Client) StartSmartDBA(request *StartSmartDBARequest) (response *StartSmartDBAResponse, err error) {
    if request == nil {
        request = NewStartSmartDBARequest()
    }
    response = NewStartSmartDBAResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchRollbackInstanceRequest() (request *SwitchRollbackInstanceRequest) {
    request = &SwitchRollbackInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "SwitchRollbackInstance")
    return
}

func NewSwitchRollbackInstanceResponse() (response *SwitchRollbackInstanceResponse) {
    response = &SwitchRollbackInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SwitchRollbackInstance）用于切换回档实例。
func (c *Client) SwitchRollbackInstance(request *SwitchRollbackInstanceRequest) (response *SwitchRollbackInstanceResponse, err error) {
    if request == nil {
        request = NewSwitchRollbackInstanceRequest()
    }
    response = NewSwitchRollbackInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDedicatedDBInstanceRequest() (request *TerminateDedicatedDBInstanceRequest) {
    request = &TerminateDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "TerminateDedicatedDBInstance")
    return
}

func NewTerminateDedicatedDBInstanceResponse() (response *TerminateDedicatedDBInstanceResponse) {
    response = &TerminateDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（IsolateDedicatedDBInstance）用于销毁已隔离的独享分布式数据库实例。
func (c *Client) TerminateDedicatedDBInstance(request *TerminateDedicatedDBInstanceRequest) (response *TerminateDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateDedicatedDBInstanceRequest()
    }
    response = NewTerminateDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDCDBInstanceRequest() (request *UpgradeDCDBInstanceRequest) {
    request = &UpgradeDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "UpgradeDCDBInstance")
    return
}

func NewUpgradeDCDBInstanceResponse() (response *UpgradeDCDBInstanceResponse) {
    response = &UpgradeDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeDCDBInstance）用于升级分布式数据库实例。本接口完成下单和支付两个动作，如果发生支付失败的错误，调用用户账户相关接口中的支付订单接口（PayDeals）重新支付即可。
func (c *Client) UpgradeDCDBInstance(request *UpgradeDCDBInstanceRequest) (response *UpgradeDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDCDBInstanceRequest()
    }
    response = NewUpgradeDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDedicatedDCDBInstanceRequest() (request *UpgradeDedicatedDCDBInstanceRequest) {
    request = &UpgradeDedicatedDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "UpgradeDedicatedDCDBInstance")
    return
}

func NewUpgradeDedicatedDCDBInstanceResponse() (response *UpgradeDedicatedDCDBInstanceResponse) {
    response = &UpgradeDedicatedDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeDedicatedDCDBInstance）用于升级独享DCDB实例
func (c *Client) UpgradeDedicatedDCDBInstance(request *UpgradeDedicatedDCDBInstanceRequest) (response *UpgradeDedicatedDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDedicatedDCDBInstanceRequest()
    }
    response = NewUpgradeDedicatedDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeHourDCDBInstanceRequest() (request *UpgradeHourDCDBInstanceRequest) {
    request = &UpgradeHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "UpgradeHourDCDBInstance")
    return
}

func NewUpgradeHourDCDBInstanceResponse() (response *UpgradeHourDCDBInstanceResponse) {
    response = &UpgradeHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeHourDCDBInstance）用于升级升级按量计费DCDB实例
func (c *Client) UpgradeHourDCDBInstance(request *UpgradeHourDCDBInstanceRequest) (response *UpgradeHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeHourDCDBInstanceRequest()
    }
    response = NewUpgradeHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}
