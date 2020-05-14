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

package v20180228

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-02-28"

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


func NewAddLoginWhiteListRequest() (request *AddLoginWhiteListRequest) {
    request = &AddLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "AddLoginWhiteList")
    return
}

func NewAddLoginWhiteListResponse() (response *AddLoginWhiteListResponse) {
    response = &AddLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AddLoginWhiteList）用于添加白名单规则
func (c *Client) AddLoginWhiteList(request *AddLoginWhiteListRequest) (response *AddLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewAddLoginWhiteListRequest()
    }
    response = NewAddLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddMachineTagRequest() (request *AddMachineTagRequest) {
    request = &AddMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "AddMachineTag")
    return
}

func NewAddMachineTagResponse() (response *AddMachineTagResponse) {
    response = &AddMachineTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加机器关联标签
func (c *Client) AddMachineTag(request *AddMachineTagRequest) (response *AddMachineTagResponse, err error) {
    if request == nil {
        request = NewAddMachineTagRequest()
    }
    response = NewAddMachineTagResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProVersionRequest() (request *CloseProVersionRequest) {
    request = &CloseProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CloseProVersion")
    return
}

func NewCloseProVersionResponse() (response *CloseProVersionResponse) {
    response = &CloseProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (CloseProVersion) 用于关闭专业版。
func (c *Client) CloseProVersion(request *CloseProVersionRequest) (response *CloseProVersionResponse, err error) {
    if request == nil {
        request = NewCloseProVersionRequest()
    }
    response = NewCloseProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBanRecordRequest() (request *CreateBanRecordRequest) {
    request = &CreateBanRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateBanRecord")
    return
}

func NewCreateBanRecordResponse() (response *CreateBanRecordResponse) {
    response = &CreateBanRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加阻断白名单
func (c *Client) CreateBanRecord(request *CreateBanRecordRequest) (response *CreateBanRecordResponse, err error) {
    if request == nil {
        request = NewCreateBanRecordRequest()
    }
    response = NewCreateBanRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBanWhiteListRequest() (request *CreateBanWhiteListRequest) {
    request = &CreateBanWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateBanWhiteList")
    return
}

func NewCreateBanWhiteListResponse() (response *CreateBanWhiteListResponse) {
    response = &CreateBanWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加阻断白名单列表
func (c *Client) CreateBanWhiteList(request *CreateBanWhiteListRequest) (response *CreateBanWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateBanWhiteListRequest()
    }
    response = NewCreateBanWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenPortTaskRequest() (request *CreateOpenPortTaskRequest) {
    request = &CreateOpenPortTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateOpenPortTask")
    return
}

func NewCreateOpenPortTaskResponse() (response *CreateOpenPortTaskResponse) {
    response = &CreateOpenPortTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (CreateOpenPortTask) 用于创建实时获取端口任务。
func (c *Client) CreateOpenPortTask(request *CreateOpenPortTaskRequest) (response *CreateOpenPortTaskResponse, err error) {
    if request == nil {
        request = NewCreateOpenPortTaskRequest()
    }
    response = NewCreateOpenPortTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcessTaskRequest() (request *CreateProcessTaskRequest) {
    request = &CreateProcessTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateProcessTask")
    return
}

func NewCreateProcessTaskResponse() (response *CreateProcessTaskResponse) {
    response = &CreateProcessTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (CreateProcessTask) 用于创建实时拉取进程任务。
func (c *Client) CreateProcessTask(request *CreateProcessTaskRequest) (response *CreateProcessTaskResponse, err error) {
    if request == nil {
        request = NewCreateProcessTaskRequest()
    }
    response = NewCreateProcessTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsualLoginPlacesRequest() (request *CreateUsualLoginPlacesRequest) {
    request = &CreateUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateUsualLoginPlaces")
    return
}

func NewCreateUsualLoginPlacesResponse() (response *CreateUsualLoginPlacesResponse) {
    response = &CreateUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口（CreateUsualLoginPlaces）用于添加常用登录地。
func (c *Client) CreateUsualLoginPlaces(request *CreateUsualLoginPlacesRequest) (response *CreateUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewCreateUsualLoginPlacesRequest()
    }
    response = NewCreateUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttackLogsRequest() (request *DeleteAttackLogsRequest) {
    request = &DeleteAttackLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteAttackLogs")
    return
}

func NewDeleteAttackLogsResponse() (response *DeleteAttackLogsResponse) {
    response = &DeleteAttackLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除网络攻击日志
func (c *Client) DeleteAttackLogs(request *DeleteAttackLogsRequest) (response *DeleteAttackLogsResponse, err error) {
    if request == nil {
        request = NewDeleteAttackLogsRequest()
    }
    response = NewDeleteAttackLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBanListRequest() (request *DeleteBanListRequest) {
    request = &DeleteBanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBanList")
    return
}

func NewDeleteBanListResponse() (response *DeleteBanListResponse) {
    response = &DeleteBanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除阻断白名单
func (c *Client) DeleteBanList(request *DeleteBanListRequest) (response *DeleteBanListResponse, err error) {
    if request == nil {
        request = NewDeleteBanListRequest()
    }
    response = NewDeleteBanListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBanWhiteListRequest() (request *DeleteBanWhiteListRequest) {
    request = &DeleteBanWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBanWhiteList")
    return
}

func NewDeleteBanWhiteListResponse() (response *DeleteBanWhiteListResponse) {
    response = &DeleteBanWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBanWhiteList
func (c *Client) DeleteBanWhiteList(request *DeleteBanWhiteListRequest) (response *DeleteBanWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteBanWhiteListRequest()
    }
    response = NewDeleteBanWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBashEventsRequest() (request *DeleteBashEventsRequest) {
    request = &DeleteBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBashEvents")
    return
}

func NewDeleteBashEventsResponse() (response *DeleteBashEventsResponse) {
    response = &DeleteBashEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据Ids删除高危命令事件
func (c *Client) DeleteBashEvents(request *DeleteBashEventsRequest) (response *DeleteBashEventsResponse, err error) {
    if request == nil {
        request = NewDeleteBashEventsRequest()
    }
    response = NewDeleteBashEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBashRulesRequest() (request *DeleteBashRulesRequest) {
    request = &DeleteBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBashRules")
    return
}

func NewDeleteBashRulesResponse() (response *DeleteBashRulesResponse) {
    response = &DeleteBashRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除高危命令规则
func (c *Client) DeleteBashRules(request *DeleteBashRulesRequest) (response *DeleteBashRulesResponse, err error) {
    if request == nil {
        request = NewDeleteBashRulesRequest()
    }
    response = NewDeleteBashRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBruteAttacksRequest() (request *DeleteBruteAttacksRequest) {
    request = &DeleteBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBruteAttacks")
    return
}

func NewDeleteBruteAttacksResponse() (response *DeleteBruteAttacksResponse) {
    response = &DeleteBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DeleteBruteAttacks) 用于删除暴力破解记录。
func (c *Client) DeleteBruteAttacks(request *DeleteBruteAttacksRequest) (response *DeleteBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDeleteBruteAttacksRequest()
    }
    response = NewDeleteBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoginWhiteListRequest() (request *DeleteLoginWhiteListRequest) {
    request = &DeleteLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteLoginWhiteList")
    return
}

func NewDeleteLoginWhiteListResponse() (response *DeleteLoginWhiteListResponse) {
    response = &DeleteLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除白名单规则
func (c *Client) DeleteLoginWhiteList(request *DeleteLoginWhiteListRequest) (response *DeleteLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteLoginWhiteListRequest()
    }
    response = NewDeleteLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineRequest() (request *DeleteMachineRequest) {
    request = &DeleteMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMachine")
    return
}

func NewDeleteMachineResponse() (response *DeleteMachineResponse) {
    response = &DeleteMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteMachine）用于卸载云镜客户端。
func (c *Client) DeleteMachine(request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    if request == nil {
        request = NewDeleteMachineRequest()
    }
    response = NewDeleteMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineTagRequest() (request *DeleteMachineTagRequest) {
    request = &DeleteMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMachineTag")
    return
}

func NewDeleteMachineTagResponse() (response *DeleteMachineTagResponse) {
    response = &DeleteMachineTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除服务器关联的标签
func (c *Client) DeleteMachineTag(request *DeleteMachineTagRequest) (response *DeleteMachineTagResponse, err error) {
    if request == nil {
        request = NewDeleteMachineTagRequest()
    }
    response = NewDeleteMachineTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMaliciousRequestsRequest() (request *DeleteMaliciousRequestsRequest) {
    request = &DeleteMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMaliciousRequests")
    return
}

func NewDeleteMaliciousRequestsResponse() (response *DeleteMaliciousRequestsResponse) {
    response = &DeleteMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DeleteMaliciousRequests) 用于删除恶意请求记录。
func (c *Client) DeleteMaliciousRequests(request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDeleteMaliciousRequestsRequest()
    }
    response = NewDeleteMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMalwaresRequest() (request *DeleteMalwaresRequest) {
    request = &DeleteMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMalwares")
    return
}

func NewDeleteMalwaresResponse() (response *DeleteMalwaresResponse) {
    response = &DeleteMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DeleteMalwares) 用于删除木马记录。
func (c *Client) DeleteMalwares(request *DeleteMalwaresRequest) (response *DeleteMalwaresResponse, err error) {
    if request == nil {
        request = NewDeleteMalwaresRequest()
    }
    response = NewDeleteMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNonlocalLoginPlacesRequest() (request *DeleteNonlocalLoginPlacesRequest) {
    request = &DeleteNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteNonlocalLoginPlaces")
    return
}

func NewDeleteNonlocalLoginPlacesResponse() (response *DeleteNonlocalLoginPlacesResponse) {
    response = &DeleteNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DeleteNonlocalLoginPlaces) 用于删除异地登录记录。
func (c *Client) DeleteNonlocalLoginPlaces(request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteNonlocalLoginPlacesRequest()
    }
    response = NewDeleteNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivilegeEventsRequest() (request *DeletePrivilegeEventsRequest) {
    request = &DeletePrivilegeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeletePrivilegeEvents")
    return
}

func NewDeletePrivilegeEventsResponse() (response *DeletePrivilegeEventsResponse) {
    response = &DeletePrivilegeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据Ids删除本地提权
func (c *Client) DeletePrivilegeEvents(request *DeletePrivilegeEventsRequest) (response *DeletePrivilegeEventsResponse, err error) {
    if request == nil {
        request = NewDeletePrivilegeEventsRequest()
    }
    response = NewDeletePrivilegeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivilegeRulesRequest() (request *DeletePrivilegeRulesRequest) {
    request = &DeletePrivilegeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeletePrivilegeRules")
    return
}

func NewDeletePrivilegeRulesResponse() (response *DeletePrivilegeRulesResponse) {
    response = &DeletePrivilegeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除本地提权规则
func (c *Client) DeletePrivilegeRules(request *DeletePrivilegeRulesRequest) (response *DeletePrivilegeRulesResponse, err error) {
    if request == nil {
        request = NewDeletePrivilegeRulesRequest()
    }
    response = NewDeletePrivilegeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReverseShellEventsRequest() (request *DeleteReverseShellEventsRequest) {
    request = &DeleteReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteReverseShellEvents")
    return
}

func NewDeleteReverseShellEventsResponse() (response *DeleteReverseShellEventsResponse) {
    response = &DeleteReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据Ids删除反弹Shell事件
func (c *Client) DeleteReverseShellEvents(request *DeleteReverseShellEventsRequest) (response *DeleteReverseShellEventsResponse, err error) {
    if request == nil {
        request = NewDeleteReverseShellEventsRequest()
    }
    response = NewDeleteReverseShellEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReverseShellRulesRequest() (request *DeleteReverseShellRulesRequest) {
    request = &DeleteReverseShellRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteReverseShellRules")
    return
}

func NewDeleteReverseShellRulesResponse() (response *DeleteReverseShellRulesResponse) {
    response = &DeleteReverseShellRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除反弹Shell规则
func (c *Client) DeleteReverseShellRules(request *DeleteReverseShellRulesRequest) (response *DeleteReverseShellRulesResponse, err error) {
    if request == nil {
        request = NewDeleteReverseShellRulesRequest()
    }
    response = NewDeleteReverseShellRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTagsRequest() (request *DeleteTagsRequest) {
    request = &DeleteTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteTags")
    return
}

func NewDeleteTagsResponse() (response *DeleteTagsResponse) {
    response = &DeleteTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除标签
func (c *Client) DeleteTags(request *DeleteTagsRequest) (response *DeleteTagsResponse, err error) {
    if request == nil {
        request = NewDeleteTagsRequest()
    }
    response = NewDeleteTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsualLoginPlacesRequest() (request *DeleteUsualLoginPlacesRequest) {
    request = &DeleteUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteUsualLoginPlaces")
    return
}

func NewDeleteUsualLoginPlacesResponse() (response *DeleteUsualLoginPlacesResponse) {
    response = &DeleteUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteUsualLoginPlaces）用于删除常用登录地。
func (c *Client) DeleteUsualLoginPlaces(request *DeleteUsualLoginPlacesRequest) (response *DeleteUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteUsualLoginPlacesRequest()
    }
    response = NewDeleteUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbTestStatusRequest() (request *DescribeAbTestStatusRequest) {
    request = &DescribeAbTestStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAbTestStatus")
    return
}

func NewDescribeAbTestStatusResponse() (response *DescribeAbTestStatusResponse) {
    response = &DescribeAbTestStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAbTestStatus）用户获取用户当前灰度状态
func (c *Client) DescribeAbTestStatus(request *DescribeAbTestStatusRequest) (response *DescribeAbTestStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAbTestStatusRequest()
    }
    response = NewDescribeAbTestStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountStatisticsRequest() (request *DescribeAccountStatisticsRequest) {
    request = &DescribeAccountStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAccountStatistics")
    return
}

func NewDescribeAccountStatisticsResponse() (response *DescribeAccountStatisticsResponse) {
    response = &DescribeAccountStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeAccountStatistics) 用于获取帐号统计列表数据。
func (c *Client) DescribeAccountStatistics(request *DescribeAccountStatisticsRequest) (response *DescribeAccountStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountStatisticsRequest()
    }
    response = NewDescribeAccountStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeAccounts) 用于获取帐号列表数据。
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentVulsRequest() (request *DescribeAgentVulsRequest) {
    request = &DescribeAgentVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAgentVuls")
    return
}

func NewDescribeAgentVulsResponse() (response *DescribeAgentVulsResponse) {
    response = &DescribeAgentVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeAgentVuls) 用于获取单台主机的漏洞列表。
func (c *Client) DescribeAgentVuls(request *DescribeAgentVulsRequest) (response *DescribeAgentVulsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentVulsRequest()
    }
    response = NewDescribeAgentVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmAttributeRequest() (request *DescribeAlarmAttributeRequest) {
    request = &DescribeAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAlarmAttribute")
    return
}

func NewDescribeAlarmAttributeResponse() (response *DescribeAlarmAttributeResponse) {
    response = &DescribeAlarmAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeAlarmAttribute) 用于获取告警设置。
func (c *Client) DescribeAlarmAttribute(request *DescribeAlarmAttributeRequest) (response *DescribeAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmAttributeRequest()
    }
    response = NewDescribeAlarmAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllMachinesRequest() (request *DescribeAllMachinesRequest) {
    request = &DescribeAllMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAllMachines")
    return
}

func NewDescribeAllMachinesResponse() (response *DescribeAllMachinesResponse) {
    response = &DescribeAllMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAllMachines）按所有获取机器列表。
func (c *Client) DescribeAllMachines(request *DescribeAllMachinesRequest) (response *DescribeAllMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeAllMachinesRequest()
    }
    response = NewDescribeAllMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackLogInfoRequest() (request *DescribeAttackLogInfoRequest) {
    request = &DescribeAttackLogInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAttackLogInfo")
    return
}

func NewDescribeAttackLogInfoResponse() (response *DescribeAttackLogInfoResponse) {
    response = &DescribeAttackLogInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络攻击日志详情
func (c *Client) DescribeAttackLogInfo(request *DescribeAttackLogInfoRequest) (response *DescribeAttackLogInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAttackLogInfoRequest()
    }
    response = NewDescribeAttackLogInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackLogsRequest() (request *DescribeAttackLogsRequest) {
    request = &DescribeAttackLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAttackLogs")
    return
}

func NewDescribeAttackLogsResponse() (response *DescribeAttackLogsResponse) {
    response = &DescribeAttackLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按分页形式展示网络攻击日志列表
func (c *Client) DescribeAttackLogs(request *DescribeAttackLogsRequest) (response *DescribeAttackLogsResponse, err error) {
    if request == nil {
        request = NewDescribeAttackLogsRequest()
    }
    response = NewDescribeAttackLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoQuaraInfoRequest() (request *DescribeAutoQuaraInfoRequest) {
    request = &DescribeAutoQuaraInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAutoQuaraInfo")
    return
}

func NewDescribeAutoQuaraInfoResponse() (response *DescribeAutoQuaraInfoResponse) {
    response = &DescribeAutoQuaraInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取自动隔离配置
func (c *Client) DescribeAutoQuaraInfo(request *DescribeAutoQuaraInfoRequest) (response *DescribeAutoQuaraInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAutoQuaraInfoRequest()
    }
    response = NewDescribeAutoQuaraInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBanListRequest() (request *DescribeBanListRequest) {
    request = &DescribeBanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBanList")
    return
}

func NewDescribeBanListResponse() (response *DescribeBanListResponse) {
    response = &DescribeBanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取阻断白名单列表
func (c *Client) DescribeBanList(request *DescribeBanListRequest) (response *DescribeBanListResponse, err error) {
    if request == nil {
        request = NewDescribeBanListRequest()
    }
    response = NewDescribeBanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBanStatusRequest() (request *DescribeBanStatusRequest) {
    request = &DescribeBanStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBanStatus")
    return
}

func NewDescribeBanStatusResponse() (response *DescribeBanStatusResponse) {
    response = &DescribeBanStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取阻断按钮状态
func (c *Client) DescribeBanStatus(request *DescribeBanStatusRequest) (response *DescribeBanStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBanStatusRequest()
    }
    response = NewDescribeBanStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBanWhiteListRequest() (request *DescribeBanWhiteListRequest) {
    request = &DescribeBanWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBanWhiteList")
    return
}

func NewDescribeBanWhiteListResponse() (response *DescribeBanWhiteListResponse) {
    response = &DescribeBanWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取阻断白名单列表
func (c *Client) DescribeBanWhiteList(request *DescribeBanWhiteListRequest) (response *DescribeBanWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeBanWhiteListRequest()
    }
    response = NewDescribeBanWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBashEventsRequest() (request *DescribeBashEventsRequest) {
    request = &DescribeBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBashEvents")
    return
}

func NewDescribeBashEventsResponse() (response *DescribeBashEventsResponse) {
    response = &DescribeBashEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取高危命令列表
func (c *Client) DescribeBashEvents(request *DescribeBashEventsRequest) (response *DescribeBashEventsResponse, err error) {
    if request == nil {
        request = NewDescribeBashEventsRequest()
    }
    response = NewDescribeBashEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBashRulesRequest() (request *DescribeBashRulesRequest) {
    request = &DescribeBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBashRules")
    return
}

func NewDescribeBashRulesResponse() (response *DescribeBashRulesResponse) {
    response = &DescribeBashRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取高危命令规则列表
func (c *Client) DescribeBashRules(request *DescribeBashRulesRequest) (response *DescribeBashRulesResponse, err error) {
    if request == nil {
        request = NewDescribeBashRulesRequest()
    }
    response = NewDescribeBashRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockStatusRequest() (request *DescribeBlockStatusRequest) {
    request = &DescribeBlockStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBlockStatus")
    return
}

func NewDescribeBlockStatusResponse() (response *DescribeBlockStatusResponse) {
    response = &DescribeBlockStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取阻断状态
func (c *Client) DescribeBlockStatus(request *DescribeBlockStatusRequest) (response *DescribeBlockStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBlockStatusRequest()
    }
    response = NewDescribeBlockStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBruteAttacksRequest() (request *DescribeBruteAttacksRequest) {
    request = &DescribeBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBruteAttacks")
    return
}

func NewDescribeBruteAttacksResponse() (response *DescribeBruteAttacksResponse) {
    response = &DescribeBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口{DescribeBruteAttacks}用于获取暴力破解事件列表。
func (c *Client) DescribeBruteAttacks(request *DescribeBruteAttacksRequest) (response *DescribeBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttacksRequest()
    }
    response = NewDescribeBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentInfoRequest() (request *DescribeComponentInfoRequest) {
    request = &DescribeComponentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponentInfo")
    return
}

func NewDescribeComponentInfoResponse() (response *DescribeComponentInfoResponse) {
    response = &DescribeComponentInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeComponentInfo) 用于获取组件信息数据。
func (c *Client) DescribeComponentInfo(request *DescribeComponentInfoRequest) (response *DescribeComponentInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComponentInfoRequest()
    }
    response = NewDescribeComponentInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentStatisticsRequest() (request *DescribeComponentStatisticsRequest) {
    request = &DescribeComponentStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponentStatistics")
    return
}

func NewDescribeComponentStatisticsResponse() (response *DescribeComponentStatisticsResponse) {
    response = &DescribeComponentStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeComponentStatistics) 用于获取组件统计列表数据。
func (c *Client) DescribeComponentStatistics(request *DescribeComponentStatisticsRequest) (response *DescribeComponentStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentStatisticsRequest()
    }
    response = NewDescribeComponentStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentsRequest() (request *DescribeComponentsRequest) {
    request = &DescribeComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponents")
    return
}

func NewDescribeComponentsResponse() (response *DescribeComponentsResponse) {
    response = &DescribeComponentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeComponents) 用于获取组件列表数据。
func (c *Client) DescribeComponents(request *DescribeComponentsRequest) (response *DescribeComponentsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentsRequest()
    }
    response = NewDescribeComponentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventListRequest() (request *DescribeEventListRequest) {
    request = &DescribeEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeEventList")
    return
}

func NewDescribeEventListResponse() (response *DescribeEventListResponse) {
    response = &DescribeEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 漏洞管理模块，获取事件列表
func (c *Client) DescribeEventList(request *DescribeEventListRequest) (response *DescribeEventListResponse, err error) {
    if request == nil {
        request = NewDescribeEventListRequest()
    }
    response = NewDescribeEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventlogRequest() (request *DescribeEventlogRequest) {
    request = &DescribeEventlogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeEventlog")
    return
}

func NewDescribeEventlogResponse() (response *DescribeEventlogResponse) {
    response = &DescribeEventlogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全事件日志列表
func (c *Client) DescribeEventlog(request *DescribeEventlogRequest) (response *DescribeEventlogResponse, err error) {
    if request == nil {
        request = NewDescribeEventlogRequest()
    }
    response = NewDescribeEventlogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryAccountsRequest() (request *DescribeHistoryAccountsRequest) {
    request = &DescribeHistoryAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeHistoryAccounts")
    return
}

func NewDescribeHistoryAccountsResponse() (response *DescribeHistoryAccountsResponse) {
    response = &DescribeHistoryAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeHistoryAccounts) 用于获取帐号变更历史列表数据。
func (c *Client) DescribeHistoryAccounts(request *DescribeHistoryAccountsRequest) (response *DescribeHistoryAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryAccountsRequest()
    }
    response = NewDescribeHistoryAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImpactedHostsRequest() (request *DescribeImpactedHostsRequest) {
    request = &DescribeImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeImpactedHosts")
    return
}

func NewDescribeImpactedHostsResponse() (response *DescribeImpactedHostsResponse) {
    response = &DescribeImpactedHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeImpactedHosts) 用于获取漏洞受影响机器列表。
func (c *Client) DescribeImpactedHosts(request *DescribeImpactedHostsRequest) (response *DescribeImpactedHostsResponse, err error) {
    if request == nil {
        request = NewDescribeImpactedHostsRequest()
    }
    response = NewDescribeImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginWhiteListRequest() (request *DescribeLoginWhiteListRequest) {
    request = &DescribeLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeLoginWhiteList")
    return
}

func NewDescribeLoginWhiteListResponse() (response *DescribeLoginWhiteListResponse) {
    response = &DescribeLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取异地登录白名单列表
func (c *Client) DescribeLoginWhiteList(request *DescribeLoginWhiteListRequest) (response *DescribeLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeLoginWhiteListRequest()
    }
    response = NewDescribeLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineInfoRequest() (request *DescribeMachineInfoRequest) {
    request = &DescribeMachineInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachineInfo")
    return
}

func NewDescribeMachineInfoResponse() (response *DescribeMachineInfoResponse) {
    response = &DescribeMachineInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeMachineInfo）用于获取机器详细信息。
func (c *Client) DescribeMachineInfo(request *DescribeMachineInfoRequest) (response *DescribeMachineInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMachineInfoRequest()
    }
    response = NewDescribeMachineInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineRegionsRequest() (request *DescribeMachineRegionsRequest) {
    request = &DescribeMachineRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachineRegions")
    return
}

func NewDescribeMachineRegionsResponse() (response *DescribeMachineRegionsResponse) {
    response = &DescribeMachineRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取机器地域列表
func (c *Client) DescribeMachineRegions(request *DescribeMachineRegionsRequest) (response *DescribeMachineRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeMachineRegionsRequest()
    }
    response = NewDescribeMachineRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
    request = &DescribeMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachines")
    return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
    response = &DescribeMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeMachines) 用于获取区域主机列表。
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeMachinesRequest()
    }
    response = NewDescribeMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaliciousRequestsRequest() (request *DescribeMaliciousRequestsRequest) {
    request = &DescribeMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMaliciousRequests")
    return
}

func NewDescribeMaliciousRequestsResponse() (response *DescribeMaliciousRequestsResponse) {
    response = &DescribeMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeMaliciousRequests) 用于获取恶意请求数据。
func (c *Client) DescribeMaliciousRequests(request *DescribeMaliciousRequestsRequest) (response *DescribeMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDescribeMaliciousRequestsRequest()
    }
    response = NewDescribeMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwaresRequest() (request *DescribeMalwaresRequest) {
    request = &DescribeMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMalwares")
    return
}

func NewDescribeMalwaresResponse() (response *DescribeMalwaresResponse) {
    response = &DescribeMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeMalwares）用于获取木马事件列表。
func (c *Client) DescribeMalwares(request *DescribeMalwaresRequest) (response *DescribeMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeMalwaresRequest()
    }
    response = NewDescribeMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNonlocalLoginPlacesRequest() (request *DescribeNonlocalLoginPlacesRequest) {
    request = &DescribeNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeNonlocalLoginPlaces")
    return
}

func NewDescribeNonlocalLoginPlacesResponse() (response *DescribeNonlocalLoginPlacesResponse) {
    response = &DescribeNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeNonlocalLoginPlaces)用于获取异地登录事件。
func (c *Client) DescribeNonlocalLoginPlaces(request *DescribeNonlocalLoginPlacesRequest) (response *DescribeNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeNonlocalLoginPlacesRequest()
    }
    response = NewDescribeNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortStatisticsRequest() (request *DescribeOpenPortStatisticsRequest) {
    request = &DescribeOpenPortStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPortStatistics")
    return
}

func NewDescribeOpenPortStatisticsResponse() (response *DescribeOpenPortStatisticsResponse) {
    response = &DescribeOpenPortStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeOpenPortStatistics) 用于获取端口统计列表。
func (c *Client) DescribeOpenPortStatistics(request *DescribeOpenPortStatisticsRequest) (response *DescribeOpenPortStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortStatisticsRequest()
    }
    response = NewDescribeOpenPortStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortTaskStatusRequest() (request *DescribeOpenPortTaskStatusRequest) {
    request = &DescribeOpenPortTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPortTaskStatus")
    return
}

func NewDescribeOpenPortTaskStatusResponse() (response *DescribeOpenPortTaskStatusResponse) {
    response = &DescribeOpenPortTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeOpenPortTaskStatus) 用于获取实时拉取端口任务状态。
func (c *Client) DescribeOpenPortTaskStatus(request *DescribeOpenPortTaskStatusRequest) (response *DescribeOpenPortTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortTaskStatusRequest()
    }
    response = NewDescribeOpenPortTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortsRequest() (request *DescribeOpenPortsRequest) {
    request = &DescribeOpenPortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPorts")
    return
}

func NewDescribeOpenPortsResponse() (response *DescribeOpenPortsResponse) {
    response = &DescribeOpenPortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeOpenPorts) 用于获取端口列表数据。
func (c *Client) DescribeOpenPorts(request *DescribeOpenPortsRequest) (response *DescribeOpenPortsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortsRequest()
    }
    response = NewDescribeOpenPortsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewStatisticsRequest() (request *DescribeOverviewStatisticsRequest) {
    request = &DescribeOverviewStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOverviewStatistics")
    return
}

func NewDescribeOverviewStatisticsResponse() (response *DescribeOverviewStatisticsResponse) {
    response = &DescribeOverviewStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于（DescribeOverviewStatistics）获取概览统计数据。
func (c *Client) DescribeOverviewStatistics(request *DescribeOverviewStatisticsRequest) (response *DescribeOverviewStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewStatisticsRequest()
    }
    response = NewDescribeOverviewStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyRequest() (request *DescribePolicyRequest) {
    request = &DescribePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribePolicy")
    return
}

func NewDescribePolicyResponse() (response *DescribePolicyResponse) {
    response = &DescribePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 自定义策略详情
func (c *Client) DescribePolicy(request *DescribePolicyRequest) (response *DescribePolicyResponse, err error) {
    if request == nil {
        request = NewDescribePolicyRequest()
    }
    response = NewDescribePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivilegeEventsRequest() (request *DescribePrivilegeEventsRequest) {
    request = &DescribePrivilegeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribePrivilegeEvents")
    return
}

func NewDescribePrivilegeEventsResponse() (response *DescribePrivilegeEventsResponse) {
    response = &DescribePrivilegeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取本地提权事件列表
func (c *Client) DescribePrivilegeEvents(request *DescribePrivilegeEventsRequest) (response *DescribePrivilegeEventsResponse, err error) {
    if request == nil {
        request = NewDescribePrivilegeEventsRequest()
    }
    response = NewDescribePrivilegeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivilegeRulesRequest() (request *DescribePrivilegeRulesRequest) {
    request = &DescribePrivilegeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribePrivilegeRules")
    return
}

func NewDescribePrivilegeRulesResponse() (response *DescribePrivilegeRulesResponse) {
    response = &DescribePrivilegeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取本地提权规则列表
func (c *Client) DescribePrivilegeRules(request *DescribePrivilegeRulesRequest) (response *DescribePrivilegeRulesResponse, err error) {
    if request == nil {
        request = NewDescribePrivilegeRulesRequest()
    }
    response = NewDescribePrivilegeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProVersionInfoRequest() (request *DescribeProVersionInfoRequest) {
    request = &DescribeProVersionInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProVersionInfo")
    return
}

func NewDescribeProVersionInfoResponse() (response *DescribeProVersionInfoResponse) {
    response = &DescribeProVersionInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeProVersionInfo) 用于获取专业版信息。
func (c *Client) DescribeProVersionInfo(request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProVersionInfoRequest()
    }
    response = NewDescribeProVersionInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProVersionStatusRequest() (request *DescribeProVersionStatusRequest) {
    request = &DescribeProVersionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProVersionStatus")
    return
}

func NewDescribeProVersionStatusResponse() (response *DescribeProVersionStatusResponse) {
    response = &DescribeProVersionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeProVersionStatus) 用于获取单台主机或所有主机是否开通专业版状态。
func (c *Client) DescribeProVersionStatus(request *DescribeProVersionStatusRequest) (response *DescribeProVersionStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProVersionStatusRequest()
    }
    response = NewDescribeProVersionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessStatisticsRequest() (request *DescribeProcessStatisticsRequest) {
    request = &DescribeProcessStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcessStatistics")
    return
}

func NewDescribeProcessStatisticsResponse() (response *DescribeProcessStatisticsResponse) {
    response = &DescribeProcessStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeProcessStatistics) 用于获取进程统计列表数据。
func (c *Client) DescribeProcessStatistics(request *DescribeProcessStatisticsRequest) (response *DescribeProcessStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProcessStatisticsRequest()
    }
    response = NewDescribeProcessStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessTaskStatusRequest() (request *DescribeProcessTaskStatusRequest) {
    request = &DescribeProcessTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcessTaskStatus")
    return
}

func NewDescribeProcessTaskStatusResponse() (response *DescribeProcessTaskStatusResponse) {
    response = &DescribeProcessTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeProcessTaskStatus) 用于获取实时拉取进程任务状态。
func (c *Client) DescribeProcessTaskStatus(request *DescribeProcessTaskStatusRequest) (response *DescribeProcessTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProcessTaskStatusRequest()
    }
    response = NewDescribeProcessTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessesRequest() (request *DescribeProcessesRequest) {
    request = &DescribeProcessesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcesses")
    return
}

func NewDescribeProcessesResponse() (response *DescribeProcessesResponse) {
    response = &DescribeProcessesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeProcesses) 用于获取进程列表数据。
func (c *Client) DescribeProcesses(request *DescribeProcessesRequest) (response *DescribeProcessesResponse, err error) {
    if request == nil {
        request = NewDescribeProcessesRequest()
    }
    response = NewDescribeProcessesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
    request = &DescribeReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeReverseShellEvents")
    return
}

func NewDescribeReverseShellEventsResponse() (response *DescribeReverseShellEventsResponse) {
    response = &DescribeReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取反弹Shell列表
func (c *Client) DescribeReverseShellEvents(request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellEventsRequest()
    }
    response = NewDescribeReverseShellEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellRulesRequest() (request *DescribeReverseShellRulesRequest) {
    request = &DescribeReverseShellRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeReverseShellRules")
    return
}

func NewDescribeReverseShellRulesResponse() (response *DescribeReverseShellRulesResponse) {
    response = &DescribeReverseShellRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取反弹Shell规则列表
func (c *Client) DescribeReverseShellRules(request *DescribeReverseShellRulesRequest) (response *DescribeReverseShellRulesResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellRulesRequest()
    }
    response = NewDescribeReverseShellRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityDynamicsRequest() (request *DescribeSecurityDynamicsRequest) {
    request = &DescribeSecurityDynamicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeSecurityDynamics")
    return
}

func NewDescribeSecurityDynamicsResponse() (response *DescribeSecurityDynamicsResponse) {
    response = &DescribeSecurityDynamicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeSecurityDynamics) 用于获取安全事件消息数据。
func (c *Client) DescribeSecurityDynamics(request *DescribeSecurityDynamicsRequest) (response *DescribeSecurityDynamicsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityDynamicsRequest()
    }
    response = NewDescribeSecurityDynamicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityTrendsRequest() (request *DescribeSecurityTrendsRequest) {
    request = &DescribeSecurityTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeSecurityTrends")
    return
}

func NewDescribeSecurityTrendsResponse() (response *DescribeSecurityTrendsResponse) {
    response = &DescribeSecurityTrendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeSecurityTrends) 用于获取安全事件统计数据。
func (c *Client) DescribeSecurityTrends(request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityTrendsRequest()
    }
    response = NewDescribeSecurityTrendsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStatRequest() (request *DescribeStatRequest) {
    request = &DescribeStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeStat")
    return
}

func NewDescribeStatResponse() (response *DescribeStatResponse) {
    response = &DescribeStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 包括在线主机数，专业版主机数，所有安全事件，病毒库，POC的更新时间等。
func (c *Client) DescribeStat(request *DescribeStatRequest) (response *DescribeStatResponse, err error) {
    if request == nil {
        request = NewDescribeStatRequest()
    }
    response = NewDescribeStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagMachinesRequest() (request *DescribeTagMachinesRequest) {
    request = &DescribeTagMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeTagMachines")
    return
}

func NewDescribeTagMachinesResponse() (response *DescribeTagMachinesResponse) {
    response = &DescribeTagMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定标签关联的服务器信息
func (c *Client) DescribeTagMachines(request *DescribeTagMachinesRequest) (response *DescribeTagMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeTagMachinesRequest()
    }
    response = NewDescribeTagMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
    request = &DescribeTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeTags")
    return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
    response = &DescribeTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取所有主机标签
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsRequest()
    }
    response = NewDescribeTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUndoVulCountsRequest() (request *DescribeUndoVulCountsRequest) {
    request = &DescribeUndoVulCountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeUndoVulCounts")
    return
}

func NewDescribeUndoVulCountsResponse() (response *DescribeUndoVulCountsResponse) {
    response = &DescribeUndoVulCountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 漏洞管理模块，获取指定类型的待处理漏洞数、主机数和非专业版主机数量
func (c *Client) DescribeUndoVulCounts(request *DescribeUndoVulCountsRequest) (response *DescribeUndoVulCountsResponse, err error) {
    if request == nil {
        request = NewDescribeUndoVulCountsRequest()
    }
    response = NewDescribeUndoVulCountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsualLoginPlacesRequest() (request *DescribeUsualLoginPlacesRequest) {
    request = &DescribeUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeUsualLoginPlaces")
    return
}

func NewDescribeUsualLoginPlacesResponse() (response *DescribeUsualLoginPlacesResponse) {
    response = &DescribeUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口（DescribeUsualLoginPlaces）用于查询常用登录地。
func (c *Client) DescribeUsualLoginPlaces(request *DescribeUsualLoginPlacesRequest) (response *DescribeUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeUsualLoginPlacesRequest()
    }
    response = NewDescribeUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVersionStatisticsRequest() (request *DescribeVersionStatisticsRequest) {
    request = &DescribeVersionStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVersionStatistics")
    return
}

func NewDescribeVersionStatisticsResponse() (response *DescribeVersionStatisticsResponse) {
    response = &DescribeVersionStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeVersionStatistics) 用于统计专业版和基础版机器数。
func (c *Client) DescribeVersionStatistics(request *DescribeVersionStatisticsRequest) (response *DescribeVersionStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeVersionStatisticsRequest()
    }
    response = NewDescribeVersionStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulCountByDatesRequest() (request *DescribeVulCountByDatesRequest) {
    request = &DescribeVulCountByDatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulCountByDates")
    return
}

func NewDescribeVulCountByDatesResponse() (response *DescribeVulCountByDatesResponse) {
    response = &DescribeVulCountByDatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 漏洞管理模块，批量获取近日指定类型的漏洞数量和主机数量
func (c *Client) DescribeVulCountByDates(request *DescribeVulCountByDatesRequest) (response *DescribeVulCountByDatesResponse, err error) {
    if request == nil {
        request = NewDescribeVulCountByDatesRequest()
    }
    response = NewDescribeVulCountByDatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulInfoRequest() (request *DescribeVulInfoRequest) {
    request = &DescribeVulInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulInfo")
    return
}

func NewDescribeVulInfoResponse() (response *DescribeVulInfoResponse) {
    response = &DescribeVulInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeVulInfo) 用于获取漏洞详情。
func (c *Client) DescribeVulInfo(request *DescribeVulInfoRequest) (response *DescribeVulInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVulInfoRequest()
    }
    response = NewDescribeVulInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulListRequest() (request *DescribeVulListRequest) {
    request = &DescribeVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulList")
    return
}

func NewDescribeVulListResponse() (response *DescribeVulListResponse) {
    response = &DescribeVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 漏洞管理模块，获取漏洞列表接口V2
func (c *Client) DescribeVulList(request *DescribeVulListRequest) (response *DescribeVulListResponse, err error) {
    if request == nil {
        request = NewDescribeVulListRequest()
    }
    response = NewDescribeVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulScanResultRequest() (request *DescribeVulScanResultRequest) {
    request = &DescribeVulScanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulScanResult")
    return
}

func NewDescribeVulScanResultResponse() (response *DescribeVulScanResultResponse) {
    response = &DescribeVulScanResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeVulScanResult) 用于获取漏洞检测结果。
func (c *Client) DescribeVulScanResult(request *DescribeVulScanResultRequest) (response *DescribeVulScanResultResponse, err error) {
    if request == nil {
        request = NewDescribeVulScanResultRequest()
    }
    response = NewDescribeVulScanResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulsRequest() (request *DescribeVulsRequest) {
    request = &DescribeVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVuls")
    return
}

func NewDescribeVulsResponse() (response *DescribeVulsResponse) {
    response = &DescribeVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeVuls) 用于获取漏洞列表数据。
func (c *Client) DescribeVuls(request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
    if request == nil {
        request = NewDescribeVulsRequest()
    }
    response = NewDescribeVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportBruteAttacksRequest() (request *DescribeWeeklyReportBruteAttacksRequest) {
    request = &DescribeWeeklyReportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportBruteAttacks")
    return
}

func NewDescribeWeeklyReportBruteAttacksResponse() (response *DescribeWeeklyReportBruteAttacksResponse) {
    response = &DescribeWeeklyReportBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeWeeklyReportBruteAttacks) 用于获取专业周报密码破解数据。
func (c *Client) DescribeWeeklyReportBruteAttacks(request *DescribeWeeklyReportBruteAttacksRequest) (response *DescribeWeeklyReportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportBruteAttacksRequest()
    }
    response = NewDescribeWeeklyReportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportInfoRequest() (request *DescribeWeeklyReportInfoRequest) {
    request = &DescribeWeeklyReportInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportInfo")
    return
}

func NewDescribeWeeklyReportInfoResponse() (response *DescribeWeeklyReportInfoResponse) {
    response = &DescribeWeeklyReportInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeWeeklyReportInfo) 用于获取专业周报详情数据。
func (c *Client) DescribeWeeklyReportInfo(request *DescribeWeeklyReportInfoRequest) (response *DescribeWeeklyReportInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportInfoRequest()
    }
    response = NewDescribeWeeklyReportInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportMalwaresRequest() (request *DescribeWeeklyReportMalwaresRequest) {
    request = &DescribeWeeklyReportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportMalwares")
    return
}

func NewDescribeWeeklyReportMalwaresResponse() (response *DescribeWeeklyReportMalwaresResponse) {
    response = &DescribeWeeklyReportMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeWeeklyReportMalwares) 用于获取专业周报木马数据。
func (c *Client) DescribeWeeklyReportMalwares(request *DescribeWeeklyReportMalwaresRequest) (response *DescribeWeeklyReportMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportMalwaresRequest()
    }
    response = NewDescribeWeeklyReportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesRequest() (request *DescribeWeeklyReportNonlocalLoginPlacesRequest) {
    request = &DescribeWeeklyReportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportNonlocalLoginPlaces")
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesResponse() (response *DescribeWeeklyReportNonlocalLoginPlacesResponse) {
    response = &DescribeWeeklyReportNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeWeeklyReportNonlocalLoginPlaces) 用于获取专业周报异地登录数据。
func (c *Client) DescribeWeeklyReportNonlocalLoginPlaces(request *DescribeWeeklyReportNonlocalLoginPlacesRequest) (response *DescribeWeeklyReportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportNonlocalLoginPlacesRequest()
    }
    response = NewDescribeWeeklyReportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportVulsRequest() (request *DescribeWeeklyReportVulsRequest) {
    request = &DescribeWeeklyReportVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportVuls")
    return
}

func NewDescribeWeeklyReportVulsResponse() (response *DescribeWeeklyReportVulsResponse) {
    response = &DescribeWeeklyReportVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeWeeklyReportVuls) 用于专业版周报漏洞数据。
func (c *Client) DescribeWeeklyReportVuls(request *DescribeWeeklyReportVulsRequest) (response *DescribeWeeklyReportVulsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportVulsRequest()
    }
    response = NewDescribeWeeklyReportVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportsRequest() (request *DescribeWeeklyReportsRequest) {
    request = &DescribeWeeklyReportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReports")
    return
}

func NewDescribeWeeklyReportsResponse() (response *DescribeWeeklyReportsResponse) {
    response = &DescribeWeeklyReportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeWeeklyReports) 用于获取周报列表数据。
func (c *Client) DescribeWeeklyReports(request *DescribeWeeklyReportsRequest) (response *DescribeWeeklyReportsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportsRequest()
    }
    response = NewDescribeWeeklyReportsResponse()
    err = c.Send(request, response)
    return
}

func NewEditBashRuleRequest() (request *EditBashRuleRequest) {
    request = &EditBashRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "EditBashRule")
    return
}

func NewEditBashRuleResponse() (response *EditBashRuleResponse) {
    response = &EditBashRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增或修改高危命令规则
func (c *Client) EditBashRule(request *EditBashRuleRequest) (response *EditBashRuleResponse, err error) {
    if request == nil {
        request = NewEditBashRuleRequest()
    }
    response = NewEditBashRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditPrivilegeRuleRequest() (request *EditPrivilegeRuleRequest) {
    request = &EditPrivilegeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "EditPrivilegeRule")
    return
}

func NewEditPrivilegeRuleResponse() (response *EditPrivilegeRuleResponse) {
    response = &EditPrivilegeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增或修改本地提权规则
func (c *Client) EditPrivilegeRule(request *EditPrivilegeRuleRequest) (response *EditPrivilegeRuleResponse, err error) {
    if request == nil {
        request = NewEditPrivilegeRuleRequest()
    }
    response = NewEditPrivilegeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditReverseShellRuleRequest() (request *EditReverseShellRuleRequest) {
    request = &EditReverseShellRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "EditReverseShellRule")
    return
}

func NewEditReverseShellRuleResponse() (response *EditReverseShellRuleResponse) {
    response = &EditReverseShellRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 编辑反弹Shell规则
func (c *Client) EditReverseShellRule(request *EditReverseShellRuleRequest) (response *EditReverseShellRuleResponse, err error) {
    if request == nil {
        request = NewEditReverseShellRuleRequest()
    }
    response = NewEditReverseShellRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditTagsRequest() (request *EditTagsRequest) {
    request = &EditTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "EditTags")
    return
}

func NewEditTagsResponse() (response *EditTagsResponse) {
    response = &EditTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增或编辑标签
func (c *Client) EditTags(request *EditTagsRequest) (response *EditTagsResponse, err error) {
    if request == nil {
        request = NewEditTagsRequest()
    }
    response = NewEditTagsResponse()
    err = c.Send(request, response)
    return
}

func NewExportAttackLogsRequest() (request *ExportAttackLogsRequest) {
    request = &ExportAttackLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportAttackLogs")
    return
}

func NewExportAttackLogsResponse() (response *ExportAttackLogsResponse) {
    response = &ExportAttackLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导出网络攻击日志
func (c *Client) ExportAttackLogs(request *ExportAttackLogsRequest) (response *ExportAttackLogsResponse, err error) {
    if request == nil {
        request = NewExportAttackLogsRequest()
    }
    response = NewExportAttackLogsResponse()
    err = c.Send(request, response)
    return
}

func NewExportBashEventsRequest() (request *ExportBashEventsRequest) {
    request = &ExportBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportBashEvents")
    return
}

func NewExportBashEventsResponse() (response *ExportBashEventsResponse) {
    response = &ExportBashEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导出高危命令事件
func (c *Client) ExportBashEvents(request *ExportBashEventsRequest) (response *ExportBashEventsResponse, err error) {
    if request == nil {
        request = NewExportBashEventsRequest()
    }
    response = NewExportBashEventsResponse()
    err = c.Send(request, response)
    return
}

func NewExportBruteAttacksRequest() (request *ExportBruteAttacksRequest) {
    request = &ExportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportBruteAttacks")
    return
}

func NewExportBruteAttacksResponse() (response *ExportBruteAttacksResponse) {
    response = &ExportBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ExportBruteAttacks) 用于导出密码破解记录成CSV文件。
func (c *Client) ExportBruteAttacks(request *ExportBruteAttacksRequest) (response *ExportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewExportBruteAttacksRequest()
    }
    response = NewExportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewExportEventlogRequest() (request *ExportEventlogRequest) {
    request = &ExportEventlogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportEventlog")
    return
}

func NewExportEventlogResponse() (response *ExportEventlogResponse) {
    response = &ExportEventlogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导出安全事件列表
func (c *Client) ExportEventlog(request *ExportEventlogRequest) (response *ExportEventlogResponse, err error) {
    if request == nil {
        request = NewExportEventlogRequest()
    }
    response = NewExportEventlogResponse()
    err = c.Send(request, response)
    return
}

func NewExportMaliciousRequestsRequest() (request *ExportMaliciousRequestsRequest) {
    request = &ExportMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportMaliciousRequests")
    return
}

func NewExportMaliciousRequestsResponse() (response *ExportMaliciousRequestsResponse) {
    response = &ExportMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ExportMaliciousRequests) 用于导出下载恶意请求文件。
func (c *Client) ExportMaliciousRequests(request *ExportMaliciousRequestsRequest) (response *ExportMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewExportMaliciousRequestsRequest()
    }
    response = NewExportMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewExportMalwaresRequest() (request *ExportMalwaresRequest) {
    request = &ExportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportMalwares")
    return
}

func NewExportMalwaresResponse() (response *ExportMalwaresResponse) {
    response = &ExportMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ExportMalwares) 用于导出木马记录CSV文件。
func (c *Client) ExportMalwares(request *ExportMalwaresRequest) (response *ExportMalwaresResponse, err error) {
    if request == nil {
        request = NewExportMalwaresRequest()
    }
    response = NewExportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewExportNonlocalLoginPlacesRequest() (request *ExportNonlocalLoginPlacesRequest) {
    request = &ExportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportNonlocalLoginPlaces")
    return
}

func NewExportNonlocalLoginPlacesResponse() (response *ExportNonlocalLoginPlacesResponse) {
    response = &ExportNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ExportNonlocalLoginPlaces) 用于导出异地登录事件记录CSV文件。
func (c *Client) ExportNonlocalLoginPlaces(request *ExportNonlocalLoginPlacesRequest) (response *ExportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewExportNonlocalLoginPlacesRequest()
    }
    response = NewExportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewExportPrivilegeEventsRequest() (request *ExportPrivilegeEventsRequest) {
    request = &ExportPrivilegeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportPrivilegeEvents")
    return
}

func NewExportPrivilegeEventsResponse() (response *ExportPrivilegeEventsResponse) {
    response = &ExportPrivilegeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导出本地提权事件
func (c *Client) ExportPrivilegeEvents(request *ExportPrivilegeEventsRequest) (response *ExportPrivilegeEventsResponse, err error) {
    if request == nil {
        request = NewExportPrivilegeEventsRequest()
    }
    response = NewExportPrivilegeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewExportReverseShellEventsRequest() (request *ExportReverseShellEventsRequest) {
    request = &ExportReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportReverseShellEvents")
    return
}

func NewExportReverseShellEventsResponse() (response *ExportReverseShellEventsResponse) {
    response = &ExportReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导出反弹Shell事件
func (c *Client) ExportReverseShellEvents(request *ExportReverseShellEventsRequest) (response *ExportReverseShellEventsResponse, err error) {
    if request == nil {
        request = NewExportReverseShellEventsRequest()
    }
    response = NewExportReverseShellEventsResponse()
    err = c.Send(request, response)
    return
}

func NewIgnoreImpactedHostsRequest() (request *IgnoreImpactedHostsRequest) {
    request = &IgnoreImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "IgnoreImpactedHosts")
    return
}

func NewIgnoreImpactedHostsResponse() (response *IgnoreImpactedHostsResponse) {
    response = &IgnoreImpactedHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (IgnoreImpactedHosts) 用于忽略漏洞。
func (c *Client) IgnoreImpactedHosts(request *IgnoreImpactedHostsRequest) (response *IgnoreImpactedHostsResponse, err error) {
    if request == nil {
        request = NewIgnoreImpactedHostsRequest()
    }
    response = NewIgnoreImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceOpenProVersionPrepaidRequest() (request *InquiryPriceOpenProVersionPrepaidRequest) {
    request = &InquiryPriceOpenProVersionPrepaidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "InquiryPriceOpenProVersionPrepaid")
    return
}

func NewInquiryPriceOpenProVersionPrepaidResponse() (response *InquiryPriceOpenProVersionPrepaidResponse) {
    response = &InquiryPriceOpenProVersionPrepaidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (InquiryPriceOpenProVersionPrepaid) 用于开通专业版询价(预付费)。
func (c *Client) InquiryPriceOpenProVersionPrepaid(request *InquiryPriceOpenProVersionPrepaidRequest) (response *InquiryPriceOpenProVersionPrepaidResponse, err error) {
    if request == nil {
        request = NewInquiryPriceOpenProVersionPrepaidRequest()
    }
    response = NewInquiryPriceOpenProVersionPrepaidResponse()
    err = c.Send(request, response)
    return
}

func NewMisAlarmNonlocalLoginPlacesRequest() (request *MisAlarmNonlocalLoginPlacesRequest) {
    request = &MisAlarmNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "MisAlarmNonlocalLoginPlaces")
    return
}

func NewMisAlarmNonlocalLoginPlacesResponse() (response *MisAlarmNonlocalLoginPlacesResponse) {
    response = &MisAlarmNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口{MisAlarmNonlocalLoginPlaces}将设置当前地点为常用登录地。
func (c *Client) MisAlarmNonlocalLoginPlaces(request *MisAlarmNonlocalLoginPlacesRequest) (response *MisAlarmNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewMisAlarmNonlocalLoginPlacesRequest()
    }
    response = NewMisAlarmNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmAttributeRequest() (request *ModifyAlarmAttributeRequest) {
    request = &ModifyAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyAlarmAttribute")
    return
}

func NewModifyAlarmAttributeResponse() (response *ModifyAlarmAttributeResponse) {
    response = &ModifyAlarmAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyAlarmAttribute）用于修改告警设置。
func (c *Client) ModifyAlarmAttribute(request *ModifyAlarmAttributeRequest) (response *ModifyAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmAttributeRequest()
    }
    response = NewModifyAlarmAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoOpenProVersionConfigRequest() (request *ModifyAutoOpenProVersionConfigRequest) {
    request = &ModifyAutoOpenProVersionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyAutoOpenProVersionConfig")
    return
}

func NewModifyAutoOpenProVersionConfigResponse() (response *ModifyAutoOpenProVersionConfigResponse) {
    response = &ModifyAutoOpenProVersionConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ModifyAutoOpenProVersionConfig) 用于设置新增主机自动开通专业版配置。
func (c *Client) ModifyAutoOpenProVersionConfig(request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoOpenProVersionConfigRequest()
    }
    response = NewModifyAutoOpenProVersionConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBanRecordRequest() (request *ModifyBanRecordRequest) {
    request = &ModifyBanRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyBanRecord")
    return
}

func NewModifyBanRecordResponse() (response *ModifyBanRecordResponse) {
    response = &ModifyBanRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改阻断白名单
func (c *Client) ModifyBanRecord(request *ModifyBanRecordRequest) (response *ModifyBanRecordResponse, err error) {
    if request == nil {
        request = NewModifyBanRecordRequest()
    }
    response = NewModifyBanRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBanStatusRequest() (request *ModifyBanStatusRequest) {
    request = &ModifyBanStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyBanStatus")
    return
}

func NewModifyBanStatusResponse() (response *ModifyBanStatusResponse) {
    response = &ModifyBanStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置阻断开关状态
func (c *Client) ModifyBanStatus(request *ModifyBanStatusRequest) (response *ModifyBanStatusResponse, err error) {
    if request == nil {
        request = NewModifyBanStatusRequest()
    }
    response = NewModifyBanStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBanWhiteListRequest() (request *ModifyBanWhiteListRequest) {
    request = &ModifyBanWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyBanWhiteList")
    return
}

func NewModifyBanWhiteListResponse() (response *ModifyBanWhiteListResponse) {
    response = &ModifyBanWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改阻断白名单列表
func (c *Client) ModifyBanWhiteList(request *ModifyBanWhiteListRequest) (response *ModifyBanWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyBanWhiteListRequest()
    }
    response = NewModifyBanWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockStatusRequest() (request *ModifyBlockStatusRequest) {
    request = &ModifyBlockStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyBlockStatus")
    return
}

func NewModifyBlockStatusResponse() (response *ModifyBlockStatusResponse) {
    response = &ModifyBlockStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改阻断状态
func (c *Client) ModifyBlockStatus(request *ModifyBlockStatusRequest) (response *ModifyBlockStatusResponse, err error) {
    if request == nil {
        request = NewModifyBlockStatusRequest()
    }
    response = NewModifyBlockStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoginWhiteListRequest() (request *ModifyLoginWhiteListRequest) {
    request = &ModifyLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyLoginWhiteList")
    return
}

func NewModifyLoginWhiteListResponse() (response *ModifyLoginWhiteListResponse) {
    response = &ModifyLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 编辑白名单规则
func (c *Client) ModifyLoginWhiteList(request *ModifyLoginWhiteListRequest) (response *ModifyLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyLoginWhiteListRequest()
    }
    response = NewModifyLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProVersionRenewFlagRequest() (request *ModifyProVersionRenewFlagRequest) {
    request = &ModifyProVersionRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyProVersionRenewFlag")
    return
}

func NewModifyProVersionRenewFlagResponse() (response *ModifyProVersionRenewFlagResponse) {
    response = &ModifyProVersionRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ModifyProVersionRenewFlag) 用于修改专业版包年包月续费标识。
func (c *Client) ModifyProVersionRenewFlag(request *ModifyProVersionRenewFlagRequest) (response *ModifyProVersionRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyProVersionRenewFlagRequest()
    }
    response = NewModifyProVersionRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVerWithQuuidsRequest() (request *OpenProVerWithQuuidsRequest) {
    request = &OpenProVerWithQuuidsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "OpenProVerWithQuuids")
    return
}

func NewOpenProVerWithQuuidsResponse() (response *OpenProVerWithQuuidsResponse) {
    response = &OpenProVerWithQuuidsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (OpenProVersion) 用于通过Quuid开通专业版。
func (c *Client) OpenProVerWithQuuids(request *OpenProVerWithQuuidsRequest) (response *OpenProVerWithQuuidsResponse, err error) {
    if request == nil {
        request = NewOpenProVerWithQuuidsRequest()
    }
    response = NewOpenProVerWithQuuidsResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionRequest() (request *OpenProVersionRequest) {
    request = &OpenProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "OpenProVersion")
    return
}

func NewOpenProVersionResponse() (response *OpenProVersionResponse) {
    response = &OpenProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (OpenProVersion) 用于开通专业版。
func (c *Client) OpenProVersion(request *OpenProVersionRequest) (response *OpenProVersionResponse, err error) {
    if request == nil {
        request = NewOpenProVersionRequest()
    }
    response = NewOpenProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionPrepaidRequest() (request *OpenProVersionPrepaidRequest) {
    request = &OpenProVersionPrepaidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "OpenProVersionPrepaid")
    return
}

func NewOpenProVersionPrepaidResponse() (response *OpenProVersionPrepaidResponse) {
    response = &OpenProVersionPrepaidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (OpenProVersionPrepaid) 用于开通专业版(包年包月)。
func (c *Client) OpenProVersionPrepaid(request *OpenProVersionPrepaidRequest) (response *OpenProVersionPrepaidResponse, err error) {
    if request == nil {
        request = NewOpenProVersionPrepaidRequest()
    }
    response = NewOpenProVersionPrepaidResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionQuickRequest() (request *OpenProVersionQuickRequest) {
    request = &OpenProVersionQuickRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "OpenProVersionQuick")
    return
}

func NewOpenProVersionQuickResponse() (response *OpenProVersionQuickResponse) {
    response = &OpenProVersionQuickResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (OpenProVersionQuick) 用于快速开通所有主机为专业版。
func (c *Client) OpenProVersionQuick(request *OpenProVersionQuickRequest) (response *OpenProVersionQuickResponse, err error) {
    if request == nil {
        request = NewOpenProVersionQuickRequest()
    }
    response = NewOpenProVersionQuickResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverMalwaresRequest() (request *RecoverMalwaresRequest) {
    request = &RecoverMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "RecoverMalwares")
    return
}

func NewRecoverMalwaresResponse() (response *RecoverMalwaresResponse) {
    response = &RecoverMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RecoverMalwares）用于批量恢复已经被隔离的木马文件。
func (c *Client) RecoverMalwares(request *RecoverMalwaresRequest) (response *RecoverMalwaresResponse, err error) {
    if request == nil {
        request = NewRecoverMalwaresRequest()
    }
    response = NewRecoverMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewRenewProVersionRequest() (request *RenewProVersionRequest) {
    request = &RenewProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "RenewProVersion")
    return
}

func NewRenewProVersionResponse() (response *RenewProVersionResponse) {
    response = &RenewProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (RenewProVersion) 用于续费专业版(包年包月)。
func (c *Client) RenewProVersion(request *RenewProVersionRequest) (response *RenewProVersionResponse, err error) {
    if request == nil {
        request = NewRenewProVersionRequest()
    }
    response = NewRenewProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewRescanImpactedHostRequest() (request *RescanImpactedHostRequest) {
    request = &RescanImpactedHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "RescanImpactedHost")
    return
}

func NewRescanImpactedHostResponse() (response *RescanImpactedHostResponse) {
    response = &RescanImpactedHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (RescanImpactedHost) 用于漏洞重新检测。
func (c *Client) RescanImpactedHost(request *RescanImpactedHostRequest) (response *RescanImpactedHostResponse, err error) {
    if request == nil {
        request = NewRescanImpactedHostRequest()
    }
    response = NewRescanImpactedHostResponse()
    err = c.Send(request, response)
    return
}

func NewSeparateMalwaresRequest() (request *SeparateMalwaresRequest) {
    request = &SeparateMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "SeparateMalwares")
    return
}

func NewSeparateMalwaresResponse() (response *SeparateMalwaresResponse) {
    response = &SeparateMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SeparateMalwares）用于隔离木马。
func (c *Client) SeparateMalwares(request *SeparateMalwaresRequest) (response *SeparateMalwaresResponse, err error) {
    if request == nil {
        request = NewSeparateMalwaresRequest()
    }
    response = NewSeparateMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewSetAutoQuaraStatusRequest() (request *SetAutoQuaraStatusRequest) {
    request = &SetAutoQuaraStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "SetAutoQuaraStatus")
    return
}

func NewSetAutoQuaraStatusResponse() (response *SetAutoQuaraStatusResponse) {
    response = &SetAutoQuaraStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置自动隔离状态
func (c *Client) SetAutoQuaraStatus(request *SetAutoQuaraStatusRequest) (response *SetAutoQuaraStatusResponse, err error) {
    if request == nil {
        request = NewSetAutoQuaraStatusRequest()
    }
    response = NewSetAutoQuaraStatusResponse()
    err = c.Send(request, response)
    return
}

func NewSetBashEventsStatusRequest() (request *SetBashEventsStatusRequest) {
    request = &SetBashEventsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "SetBashEventsStatus")
    return
}

func NewSetBashEventsStatusResponse() (response *SetBashEventsStatusResponse) {
    response = &SetBashEventsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置高危命令事件状态
func (c *Client) SetBashEventsStatus(request *SetBashEventsStatusRequest) (response *SetBashEventsStatusResponse, err error) {
    if request == nil {
        request = NewSetBashEventsStatusRequest()
    }
    response = NewSetBashEventsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchBashRulesRequest() (request *SwitchBashRulesRequest) {
    request = &SwitchBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "SwitchBashRules")
    return
}

func NewSwitchBashRulesResponse() (response *SwitchBashRulesResponse) {
    response = &SwitchBashRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 切换高危命令规则状态
func (c *Client) SwitchBashRules(request *SwitchBashRulesRequest) (response *SwitchBashRulesResponse, err error) {
    if request == nil {
        request = NewSwitchBashRulesRequest()
    }
    response = NewSwitchBashRulesResponse()
    err = c.Send(request, response)
    return
}

func NewTrustMaliciousRequestRequest() (request *TrustMaliciousRequestRequest) {
    request = &TrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "TrustMaliciousRequest")
    return
}

func NewTrustMaliciousRequestResponse() (response *TrustMaliciousRequestResponse) {
    response = &TrustMaliciousRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (TrustMaliciousRequest) 用于恶意请求添加信任。
func (c *Client) TrustMaliciousRequest(request *TrustMaliciousRequestRequest) (response *TrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewTrustMaliciousRequestRequest()
    }
    response = NewTrustMaliciousRequestResponse()
    err = c.Send(request, response)
    return
}

func NewTrustMalwaresRequest() (request *TrustMalwaresRequest) {
    request = &TrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "TrustMalwares")
    return
}

func NewTrustMalwaresResponse() (response *TrustMalwaresResponse) {
    response = &TrustMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(TrustMalwares)将被识别木马文件设为信任。
func (c *Client) TrustMalwares(request *TrustMalwaresRequest) (response *TrustMalwaresResponse, err error) {
    if request == nil {
        request = NewTrustMalwaresRequest()
    }
    response = NewTrustMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMaliciousRequestRequest() (request *UntrustMaliciousRequestRequest) {
    request = &UntrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "UntrustMaliciousRequest")
    return
}

func NewUntrustMaliciousRequestResponse() (response *UntrustMaliciousRequestResponse) {
    response = &UntrustMaliciousRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (UntrustMaliciousRequest) 用于取消信任恶意请求。
func (c *Client) UntrustMaliciousRequest(request *UntrustMaliciousRequestRequest) (response *UntrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewUntrustMaliciousRequestRequest()
    }
    response = NewUntrustMaliciousRequestResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMalwaresRequest() (request *UntrustMalwaresRequest) {
    request = &UntrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "UntrustMalwares")
    return
}

func NewUntrustMalwaresResponse() (response *UntrustMalwaresResponse) {
    response = &UntrustMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UntrustMalwares）用于取消信任木马文件。
func (c *Client) UntrustMalwares(request *UntrustMalwaresRequest) (response *UntrustMalwaresResponse, err error) {
    if request == nil {
        request = NewUntrustMalwaresRequest()
    }
    response = NewUntrustMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
    request = &UpdatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "UpdatePolicy")
    return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
    response = &UpdatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新自定义策略
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
}
