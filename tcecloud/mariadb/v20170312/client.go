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

package v20170312

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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
    request.Init().WithApiInfo("mariadb", APIVersion, "ActiveDedicatedDBInstance")
    return
}

func NewActiveDedicatedDBInstanceResponse() (response *ActiveDedicatedDBInstanceResponse) {
    response = &ActiveDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ActiveDedicatedDBInstance）用于恢复已隔离的独享云数据库实例。
func (c *Client) ActiveDedicatedDBInstance(request *ActiveDedicatedDBInstanceRequest) (response *ActiveDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewActiveDedicatedDBInstanceRequest()
    }
    response = NewActiveDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewActiveHourDBInstanceRequest() (request *ActiveHourDBInstanceRequest) {
    request = &ActiveHourDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ActiveHourDBInstance")
    return
}

func NewActiveHourDBInstanceResponse() (response *ActiveHourDBInstanceResponse) {
    response = &ActiveHourDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ActiveHourDBInstance）用于恢复隔离的按量计费实例。
func (c *Client) ActiveHourDBInstance(request *ActiveHourDBInstanceRequest) (response *ActiveHourDBInstanceResponse, err error) {
    if request == nil {
        request = NewActiveHourDBInstanceRequest()
    }
    response = NewActiveHourDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewAuthenticateCAMRequest() (request *AuthenticateCAMRequest) {
    request = &AuthenticateCAMRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "AuthenticateCAM")
    return
}

func NewAuthenticateCAMResponse() (response *AuthenticateCAMResponse) {
    response = &AuthenticateCAMResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口为控制台提供预鉴权
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
    request.Init().WithApiInfo("mariadb", APIVersion, "CheckIpStatus")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "CloneAccount")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "CloseDBExtranetAccess")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "CopyAccountPrivileges")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateAccount")
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

func NewCreateConfigTemplateRequest() (request *CreateConfigTemplateRequest) {
    request = &CreateConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateConfigTemplate")
    return
}

func NewCreateConfigTemplateResponse() (response *CreateConfigTemplateResponse) {
    response = &CreateConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateConfigTemplate）用于创建参数模板。
func (c *Client) CreateConfigTemplate(request *CreateConfigTemplateRequest) (response *CreateConfigTemplateResponse, err error) {
    if request == nil {
        request = NewCreateConfigTemplateRequest()
    }
    response = NewCreateConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
    request = &CreateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateDBInstance")
    return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
    response = &CreateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateDBInstance）用于创建包年包月的云数据库实例，可通过传入实例规格、数据库版本号、购买时长和数量等信息创建云数据库实例。
func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceRequest()
    }
    response = NewCreateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHourDBInstanceRequest() (request *CreateHourDBInstanceRequest) {
    request = &CreateHourDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateHourDBInstance")
    return
}

func NewCreateHourDBInstanceResponse() (response *CreateHourDBInstanceResponse) {
    response = &CreateHourDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateHourDBInstance）用于创建按量计费实例。
func (c *Client) CreateHourDBInstance(request *CreateHourDBInstanceRequest) (response *CreateHourDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateHourDBInstanceRequest()
    }
    response = NewCreateHourDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTmpInstancesRequest() (request *CreateTmpInstancesRequest) {
    request = &CreateTmpInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateTmpInstances")
    return
}

func NewCreateTmpInstancesResponse() (response *CreateTmpInstancesResponse) {
    response = &CreateTmpInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateTmpInstances）用于创建临时实例。
func (c *Client) CreateTmpInstances(request *CreateTmpInstancesRequest) (response *CreateTmpInstancesResponse, err error) {
    if request == nil {
        request = NewCreateTmpInstancesRequest()
    }
    response = NewCreateTmpInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DeleteAccount")
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

func NewDeleteConfigTemplateRequest() (request *DeleteConfigTemplateRequest) {
    request = &DeleteConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DeleteConfigTemplate")
    return
}

func NewDeleteConfigTemplateResponse() (response *DeleteConfigTemplateResponse) {
    response = &DeleteConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteConfigTemplate）用于删除参数模板。
func (c *Client) DeleteConfigTemplate(request *DeleteConfigTemplateRequest) (response *DeleteConfigTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteConfigTemplateRequest()
    }
    response = NewDeleteConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTmpInstanceRequest() (request *DeleteTmpInstanceRequest) {
    request = &DeleteTmpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DeleteTmpInstance")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAccountPrivileges")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAccounts")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAuditLogs")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAuditRuleDetail")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAuditRules")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAuditStrategies")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAvailableExclusiveGroups")
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

func NewDescribeBackupTimeRequest() (request *DescribeBackupTimeRequest) {
    request = &DescribeBackupTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeBackupTime")
    return
}

func NewDescribeBackupTimeResponse() (response *DescribeBackupTimeResponse) {
    response = &DescribeBackupTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeBackupTime）用于获取云数据库的备份时间。后台系统将根据此配置定期进行实例备份。
func (c *Client) DescribeBackupTime(request *DescribeBackupTimeRequest) (response *DescribeBackupTimeResponse, err error) {
    if request == nil {
        request = NewDescribeBackupTimeRequest()
    }
    response = NewDescribeBackupTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchRenewalPriceRequest() (request *DescribeBatchRenewalPriceRequest) {
    request = &DescribeBatchRenewalPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeBatchRenewalPrice")
    return
}

func NewDescribeBatchRenewalPriceResponse() (response *DescribeBatchRenewalPriceResponse) {
    response = &DescribeBatchRenewalPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeBatchRenewalPrice）用于批量实例续费询价
func (c *Client) DescribeBatchRenewalPrice(request *DescribeBatchRenewalPriceRequest) (response *DescribeBatchRenewalPriceResponse, err error) {
    if request == nil {
        request = NewDescribeBatchRenewalPriceRequest()
    }
    response = NewDescribeBatchRenewalPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogTimeRequest() (request *DescribeBinlogTimeRequest) {
    request = &DescribeBinlogTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeBinlogTime")
    return
}

func NewDescribeBinlogTimeResponse() (response *DescribeBinlogTimeResponse) {
    response = &DescribeBinlogTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeBinlogTime）用于查询可回档时间范围。
func (c *Client) DescribeBinlogTime(request *DescribeBinlogTimeRequest) (response *DescribeBinlogTimeResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogTimeRequest()
    }
    response = NewDescribeBinlogTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigHistoriesRequest() (request *DescribeConfigHistoriesRequest) {
    request = &DescribeConfigHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeConfigHistories")
    return
}

func NewDescribeConfigHistoriesResponse() (response *DescribeConfigHistoriesResponse) {
    response = &DescribeConfigHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeConfigHistories）用于查询配置历史列表。
func (c *Client) DescribeConfigHistories(request *DescribeConfigHistoriesRequest) (response *DescribeConfigHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigHistoriesRequest()
    }
    response = NewDescribeConfigHistoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigTemplateRequest() (request *DescribeConfigTemplateRequest) {
    request = &DescribeConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeConfigTemplate")
    return
}

func NewDescribeConfigTemplateResponse() (response *DescribeConfigTemplateResponse) {
    response = &DescribeConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeConfigTemplate）用于查询参数模板详情。
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeConfigTemplates")
    return
}

func NewDescribeConfigTemplatesResponse() (response *DescribeConfigTemplatesResponse) {
    response = &DescribeConfigTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeConfigTemplates）用于查询参数模板列表。
func (c *Client) DescribeConfigTemplates(request *DescribeConfigTemplatesRequest) (response *DescribeConfigTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigTemplatesRequest()
    }
    response = NewDescribeConfigTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDetailMetricsRequest() (request *DescribeDBDetailMetricsRequest) {
    request = &DescribeDBDetailMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBDetailMetrics")
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

func NewDescribeDBEncryptAttributesRequest() (request *DescribeDBEncryptAttributesRequest) {
    request = &DescribeDBEncryptAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBEncryptAttributes")
    return
}

func NewDescribeDBEncryptAttributesResponse() (response *DescribeDBEncryptAttributesResponse) {
    response = &DescribeDBEncryptAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBEncryptAttributes)用于查询实例数据加密状态。
func (c *Client) DescribeDBEncryptAttributes(request *DescribeDBEncryptAttributesRequest) (response *DescribeDBEncryptAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeDBEncryptAttributesRequest()
    }
    response = NewDescribeDBEncryptAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBEnginesRequest() (request *DescribeDBEnginesRequest) {
    request = &DescribeDBEnginesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBEngines")
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

func NewDescribeDBInstanceDetailRequest() (request *DescribeDBInstanceDetailRequest) {
    request = &DescribeDBInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBInstanceDetail")
    return
}

func NewDescribeDBInstanceDetailResponse() (response *DescribeDBInstanceDetailResponse) {
    response = &DescribeDBInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBInstanceDetail)用于查询指定实例的详细信息。
func (c *Client) DescribeDBInstanceDetail(request *DescribeDBInstanceDetailRequest) (response *DescribeDBInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceDetailRequest()
    }
    response = NewDescribeDBInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceHAInfoRequest() (request *DescribeDBInstanceHAInfoRequest) {
    request = &DescribeDBInstanceHAInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBInstanceHAInfo")
    return
}

func NewDescribeDBInstanceHAInfoResponse() (response *DescribeDBInstanceHAInfoResponse) {
    response = &DescribeDBInstanceHAInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBInstanceHAInfo）用于查询数据库实例的当前主可用区及主备切换状态。
func (c *Client) DescribeDBInstanceHAInfo(request *DescribeDBInstanceHAInfoRequest) (response *DescribeDBInstanceHAInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceHAInfoRequest()
    }
    response = NewDescribeDBInstanceHAInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceRsipRequest() (request *DescribeDBInstanceRsipRequest) {
    request = &DescribeDBInstanceRsipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBInstanceRsip")
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

func NewDescribeDBInstanceSpecsRequest() (request *DescribeDBInstanceSpecsRequest) {
    request = &DescribeDBInstanceSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBInstanceSpecs")
    return
}

func NewDescribeDBInstanceSpecsResponse() (response *DescribeDBInstanceSpecsResponse) {
    response = &DescribeDBInstanceSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBInstanceSpecs)用于查询可创建的云数据库可售卖的规格配置。
func (c *Client) DescribeDBInstanceSpecs(request *DescribeDBInstanceSpecsRequest) (response *DescribeDBInstanceSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceSpecsRequest()
    }
    response = NewDescribeDBInstanceSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBInstances）用于查询云数据库实例列表，支持通过项目ID、实例ID、内网地址、实例名称等来筛选实例。
// 如果不指定任何筛选条件，则默认返回20条实例记录，单次请求最多支持返回100条实例记录。
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
    request = &DescribeDBLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBLogFiles")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBMetrics")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBParameters")
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

func NewDescribeDBPerformanceRequest() (request *DescribeDBPerformanceRequest) {
    request = &DescribeDBPerformanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBPerformance")
    return
}

func NewDescribeDBPerformanceResponse() (response *DescribeDBPerformanceResponse) {
    response = &DescribeDBPerformanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBPerformance)用于查看数据库实例当前性能数据。
func (c *Client) DescribeDBPerformance(request *DescribeDBPerformanceRequest) (response *DescribeDBPerformanceResponse, err error) {
    if request == nil {
        request = NewDescribeDBPerformanceRequest()
    }
    response = NewDescribeDBPerformanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBPerformanceDetailsRequest() (request *DescribeDBPerformanceDetailsRequest) {
    request = &DescribeDBPerformanceDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBPerformanceDetails")
    return
}

func NewDescribeDBPerformanceDetailsResponse() (response *DescribeDBPerformanceDetailsResponse) {
    response = &DescribeDBPerformanceDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBPerformanceDetails)用于查看实例性能数据详情。
func (c *Client) DescribeDBPerformanceDetails(request *DescribeDBPerformanceDetailsRequest) (response *DescribeDBPerformanceDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeDBPerformanceDetailsRequest()
    }
    response = NewDescribeDBPerformanceDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBResourceUsageRequest() (request *DescribeDBResourceUsageRequest) {
    request = &DescribeDBResourceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBResourceUsage")
    return
}

func NewDescribeDBResourceUsageResponse() (response *DescribeDBResourceUsageResponse) {
    response = &DescribeDBResourceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBResourceUsage)用于查看数据库实例资源的使用情况。
func (c *Client) DescribeDBResourceUsage(request *DescribeDBResourceUsageRequest) (response *DescribeDBResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeDBResourceUsageRequest()
    }
    response = NewDescribeDBResourceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBResourceUsageDetailsRequest() (request *DescribeDBResourceUsageDetailsRequest) {
    request = &DescribeDBResourceUsageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBResourceUsageDetails")
    return
}

func NewDescribeDBResourceUsageDetailsResponse() (response *DescribeDBResourceUsageDetailsResponse) {
    response = &DescribeDBResourceUsageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBResourceUsageDetails)用于查看数据库实例当前性能数据。
func (c *Client) DescribeDBResourceUsageDetails(request *DescribeDBResourceUsageDetailsRequest) (response *DescribeDBResourceUsageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeDBResourceUsageDetailsRequest()
    }
    response = NewDescribeDBResourceUsageDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBSecurityGroups")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBSlowLogAnalysis")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBSlowLogs")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBSyncMode")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBTmpInstances")
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

func NewDescribeDatabaseObjectsRequest() (request *DescribeDatabaseObjectsRequest) {
    request = &DescribeDatabaseObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDatabaseObjects")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDatabaseTable")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDatabases")
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

func NewDescribeDefaultConfigTemplateRequest() (request *DescribeDefaultConfigTemplateRequest) {
    request = &DescribeDefaultConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDefaultConfigTemplate")
    return
}

func NewDescribeDefaultConfigTemplateResponse() (response *DescribeDefaultConfigTemplateResponse) {
    response = &DescribeDefaultConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDefaultConfigTemplate）用于查询默认参数模板信息。
func (c *Client) DescribeDefaultConfigTemplate(request *DescribeDefaultConfigTemplateRequest) (response *DescribeDefaultConfigTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultConfigTemplateRequest()
    }
    response = NewDescribeDefaultConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFenceDBInstanceSpecsRequest() (request *DescribeFenceDBInstanceSpecsRequest) {
    request = &DescribeFenceDBInstanceSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeFenceDBInstanceSpecs")
    return
}

func NewDescribeFenceDBInstanceSpecsResponse() (response *DescribeFenceDBInstanceSpecsResponse) {
    response = &DescribeFenceDBInstanceSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeFenceShardSpec)用于查询可创建的独享云数据库实例的规格配置。
func (c *Client) DescribeFenceDBInstanceSpecs(request *DescribeFenceDBInstanceSpecsRequest) (response *DescribeFenceDBInstanceSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeFenceDBInstanceSpecsRequest()
    }
    response = NewDescribeFenceDBInstanceSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeFlow")
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeFlow）用于查询流程状态。
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceProxyConfigRequest() (request *DescribeInstanceProxyConfigRequest) {
    request = &DescribeInstanceProxyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeInstanceProxyConfig")
    return
}

func NewDescribeInstanceProxyConfigResponse() (response *DescribeInstanceProxyConfigResponse) {
    response = &DescribeInstanceProxyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstanceProxyConfig）用于拉取实例网关配置
func (c *Client) DescribeInstanceProxyConfig(request *DescribeInstanceProxyConfigRequest) (response *DescribeInstanceProxyConfigResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceProxyConfigRequest()
    }
    response = NewDescribeInstanceProxyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSSLAttributesRequest() (request *DescribeInstanceSSLAttributesRequest) {
    request = &DescribeInstanceSSLAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeInstanceSSLAttributes")
    return
}

func NewDescribeInstanceSSLAttributesResponse() (response *DescribeInstanceSSLAttributesResponse) {
    response = &DescribeInstanceSSLAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstanceSSLAttributes）用于拉取实例SSL认证属性
func (c *Client) DescribeInstanceSSLAttributes(request *DescribeInstanceSSLAttributesRequest) (response *DescribeInstanceSSLAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSSLAttributesRequest()
    }
    response = NewDescribeInstanceSSLAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstances）用于拉取实例信息，Barad使用。
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeLatestCloudDBAReport")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeLogFileRetentionPeriod")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeOrders")
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeOrders）用于查询云数据库订单信息。传入订单Id来查询订单关联的云数据库实例，和对应的任务流程ID。
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePriceRequest() (request *DescribePriceRequest) {
    request = &DescribePriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribePrice")
    return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
    response = &DescribePriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribePrice）用于在购买实例前，查询实例的价格。
func (c *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
    if request == nil {
        request = NewDescribePriceRequest()
    }
    response = NewDescribePriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeProjectSecurityGroups")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeProjects")
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeProjects）用于查询项目列表。
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRenewalPriceRequest() (request *DescribeRenewalPriceRequest) {
    request = &DescribeRenewalPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeRenewalPrice")
    return
}

func NewDescribeRenewalPriceResponse() (response *DescribeRenewalPriceResponse) {
    response = &DescribeRenewalPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRenewalPrice）用于在续费云数据库实例时，查询续费的价格。
func (c *Client) DescribeRenewalPrice(request *DescribeRenewalPriceRequest) (response *DescribeRenewalPriceResponse, err error) {
    if request == nil {
        request = NewDescribeRenewalPriceRequest()
    }
    response = NewDescribeRenewalPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSaleInfoRequest() (request *DescribeSaleInfoRequest) {
    request = &DescribeSaleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeSaleInfo")
    return
}

func NewDescribeSaleInfoResponse() (response *DescribeSaleInfoResponse) {
    response = &DescribeSaleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeSaleInfo)用于查询云数据库可售卖的地域和可用区信息。
func (c *Client) DescribeSaleInfo(request *DescribeSaleInfoRequest) (response *DescribeSaleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSaleInfoRequest()
    }
    response = NewDescribeSaleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSqlLogsRequest() (request *DescribeSqlLogsRequest) {
    request = &DescribeSqlLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeSqlLogs")
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

func NewDescribeSyncTasksRequest() (request *DescribeSyncTasksRequest) {
    request = &DescribeSyncTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeSyncTasks")
    return
}

func NewDescribeSyncTasksResponse() (response *DescribeSyncTasksResponse) {
    response = &DescribeSyncTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSyncTasks）用于拉取多源同步任务列表
func (c *Client) DescribeSyncTasks(request *DescribeSyncTasksRequest) (response *DescribeSyncTasksResponse, err error) {
    if request == nil {
        request = NewDescribeSyncTasksRequest()
    }
    response = NewDescribeSyncTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpgradePriceRequest() (request *DescribeUpgradePriceRequest) {
    request = &DescribeUpgradePriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeUpgradePrice")
    return
}

func NewDescribeUpgradePriceResponse() (response *DescribeUpgradePriceResponse) {
    response = &DescribeUpgradePriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeUpgradePrice）用于在扩容云数据库实例时，查询扩容的价格。
func (c *Client) DescribeUpgradePrice(request *DescribeUpgradePriceRequest) (response *DescribeUpgradePriceResponse, err error) {
    if request == nil {
        request = NewDescribeUpgradePriceRequest()
    }
    response = NewDescribeUpgradePriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserTasksRequest() (request *DescribeUserTasksRequest) {
    request = &DescribeUserTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeUserTasks")
    return
}

func NewDescribeUserTasksResponse() (response *DescribeUserTasksResponse) {
    response = &DescribeUserTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeUserTasks）用于查询用户任务列表。
func (c *Client) DescribeUserTasks(request *DescribeUserTasksRequest) (response *DescribeUserTasksResponse, err error) {
    if request == nil {
        request = NewDescribeUserTasksRequest()
    }
    response = NewDescribeUserTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyHourDBInstanceRequest() (request *DestroyHourDBInstanceRequest) {
    request = &DestroyHourDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DestroyHourDBInstance")
    return
}

func NewDestroyHourDBInstanceResponse() (response *DestroyHourDBInstanceResponse) {
    response = &DestroyHourDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DestroyHourDBInstance）用于销毁按量计费实例。
func (c *Client) DestroyHourDBInstance(request *DestroyHourDBInstanceRequest) (response *DestroyHourDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyHourDBInstanceRequest()
    }
    response = NewDestroyHourDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewGrantAccountPrivilegesRequest() (request *GrantAccountPrivilegesRequest) {
    request = &GrantAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "GrantAccountPrivileges")
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

func NewInitDBInstancesRequest() (request *InitDBInstancesRequest) {
    request = &InitDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "InitDBInstances")
    return
}

func NewInitDBInstancesResponse() (response *InitDBInstancesResponse) {
    response = &InitDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(InitDBInstances)用于初始化云数据库实例，包括设置默认字符集、表名大小写敏感等。
func (c *Client) InitDBInstances(request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDBInstancesRequest()
    }
    response = NewInitDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDedicatedDBInstanceRequest() (request *IsolateDedicatedDBInstanceRequest) {
    request = &IsolateDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "IsolateDedicatedDBInstance")
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

func NewIsolateHourDBInstanceRequest() (request *IsolateHourDBInstanceRequest) {
    request = &IsolateHourDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "IsolateHourDBInstance")
    return
}

func NewIsolateHourDBInstanceResponse() (response *IsolateHourDBInstanceResponse) {
    response = &IsolateHourDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（IsolateHourDBInstance）用于隔离按量计费实例。
func (c *Client) IsolateHourDBInstance(request *IsolateHourDBInstanceRequest) (response *IsolateHourDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateHourDBInstanceRequest()
    }
    response = NewIsolateHourDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyAccountDescription")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyAutoRenewFlag")
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

func NewModifyBackupTimeRequest() (request *ModifyBackupTimeRequest) {
    request = &ModifyBackupTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyBackupTime")
    return
}

func NewModifyBackupTimeResponse() (response *ModifyBackupTimeResponse) {
    response = &ModifyBackupTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyBackupTime）用于设置云数据库实例的备份时间。后台系统将根据此配置定期进行实例备份。
func (c *Client) ModifyBackupTime(request *ModifyBackupTimeRequest) (response *ModifyBackupTimeResponse, err error) {
    if request == nil {
        request = NewModifyBackupTimeRequest()
    }
    response = NewModifyBackupTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConfigTemplateRequest() (request *ModifyConfigTemplateRequest) {
    request = &ModifyConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyConfigTemplate")
    return
}

func NewModifyConfigTemplateResponse() (response *ModifyConfigTemplateResponse) {
    response = &ModifyConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyConfigTemplate）用于修改参数模板。
func (c *Client) ModifyConfigTemplate(request *ModifyConfigTemplateRequest) (response *ModifyConfigTemplateResponse, err error) {
    if request == nil {
        request = NewModifyConfigTemplateRequest()
    }
    response = NewModifyConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBEncryptAttributesRequest() (request *ModifyDBEncryptAttributesRequest) {
    request = &ModifyDBEncryptAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBEncryptAttributes")
    return
}

func NewModifyDBEncryptAttributesResponse() (response *ModifyDBEncryptAttributesResponse) {
    response = &ModifyDBEncryptAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyDBEncryptAttributes)用于修改实例数据加密。
func (c *Client) ModifyDBEncryptAttributes(request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDBEncryptAttributesRequest()
    }
    response = NewModifyDBEncryptAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBInstanceName")
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstanceName）用于修改云数据库实例的名称。
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
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBInstanceSecurityGroups")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBInstancesProject")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBParameters")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBSyncMode")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyInstanceNetwork")
    return
}

func NewModifyInstanceNetworkResponse() (response *ModifyInstanceNetworkResponse) {
    response = &ModifyInstanceNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyInstanceNetwork）用于修改实例所属网络
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
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyInstanceRemark")
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

func NewModifyInstanceSSLAttributesRequest() (request *ModifyInstanceSSLAttributesRequest) {
    request = &ModifyInstanceSSLAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyInstanceSSLAttributes")
    return
}

func NewModifyInstanceSSLAttributesResponse() (response *ModifyInstanceSSLAttributesResponse) {
    response = &ModifyInstanceSSLAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口  （ModifyInstanceSSLAttributes）用于修改实例SSL认证功能属性
func (c *Client) ModifyInstanceSSLAttributes(request *ModifyInstanceSSLAttributesRequest) (response *ModifyInstanceSSLAttributesResponse, err error) {
    if request == nil {
        request = NewModifyInstanceSSLAttributesRequest()
    }
    response = NewModifyInstanceSSLAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceVipRequest() (request *ModifyInstanceVipRequest) {
    request = &ModifyInstanceVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyInstanceVip")
    return
}

func NewModifyInstanceVipResponse() (response *ModifyInstanceVipResponse) {
    response = &ModifyInstanceVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyInstanceVip）用于修改实例VIP
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
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyLogFileRetentionPeriod")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "OpenDBExtranetAccess")
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

func NewRenewDBInstanceRequest() (request *RenewDBInstanceRequest) {
    request = &RenewDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "RenewDBInstance")
    return
}

func NewRenewDBInstanceResponse() (response *RenewDBInstanceResponse) {
    response = &RenewDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RenewDBInstance）用于续费云数据库实例。
func (c *Client) RenewDBInstance(request *RenewDBInstanceRequest) (response *RenewDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewDBInstanceRequest()
    }
    response = NewRenewDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ResetAccountPassword")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "StartSmartDBA")
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

func NewSwitchDBInstanceHARequest() (request *SwitchDBInstanceHARequest) {
    request = &SwitchDBInstanceHARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "SwitchDBInstanceHA")
    return
}

func NewSwitchDBInstanceHAResponse() (response *SwitchDBInstanceHAResponse) {
    response = &SwitchDBInstanceHAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SwitchDBInstanceHA）用于发起实例主备切换。
func (c *Client) SwitchDBInstanceHA(request *SwitchDBInstanceHARequest) (response *SwitchDBInstanceHAResponse, err error) {
    if request == nil {
        request = NewSwitchDBInstanceHARequest()
    }
    response = NewSwitchDBInstanceHAResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchRollbackInstanceRequest() (request *SwitchRollbackInstanceRequest) {
    request = &SwitchRollbackInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "SwitchRollbackInstance")
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
    request.Init().WithApiInfo("mariadb", APIVersion, "TerminateDedicatedDBInstance")
    return
}

func NewTerminateDedicatedDBInstanceResponse() (response *TerminateDedicatedDBInstanceResponse) {
    response = &TerminateDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（IsolateDedicatedDBInstance）用于销毁已隔离的独享云数据库实例。
func (c *Client) TerminateDedicatedDBInstance(request *TerminateDedicatedDBInstanceRequest) (response *TerminateDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateDedicatedDBInstanceRequest()
    }
    response = NewTerminateDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "UpgradeDBInstance")
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(UpgradeDBInstance)用于扩容云数据库实例。本接口完成下单和支付两个动作，如果发生支付失败的错误，调用用户账户相关接口中的支付订单接口（PayDeals）重新支付即可。
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDedicatedDBInstanceRequest() (request *UpgradeDedicatedDBInstanceRequest) {
    request = &UpgradeDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "UpgradeDedicatedDBInstance")
    return
}

func NewUpgradeDedicatedDBInstanceResponse() (response *UpgradeDedicatedDBInstanceResponse) {
    response = &UpgradeDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(UpgradeDedicatedDBInstance)用于扩容独享云数据库实例。
func (c *Client) UpgradeDedicatedDBInstance(request *UpgradeDedicatedDBInstanceRequest) (response *UpgradeDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDedicatedDBInstanceRequest()
    }
    response = NewUpgradeDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeHourDBInstanceRequest() (request *UpgradeHourDBInstanceRequest) {
    request = &UpgradeHourDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "UpgradeHourDBInstance")
    return
}

func NewUpgradeHourDBInstanceResponse() (response *UpgradeHourDBInstanceResponse) {
    response = &UpgradeHourDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeHourDBInstance）用于升级按量计费的规格。
func (c *Client) UpgradeHourDBInstance(request *UpgradeHourDBInstanceRequest) (response *UpgradeHourDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeHourDBInstanceRequest()
    }
    response = NewUpgradeHourDBInstanceResponse()
    err = c.Send(request, response)
    return
}
