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

package v20190621

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-21"

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


func NewCreateAssetTaskRequest() (request *CreateAssetTaskRequest) {
    request = &CreateAssetTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "CreateAssetTask")
    return
}

func NewCreateAssetTaskResponse() (response *CreateAssetTaskResponse) {
    response = &CreateAssetTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从资产列表页，添加特定资产的立即扫描或定时扫描计划
func (c *Client) CreateAssetTask(request *CreateAssetTaskRequest) (response *CreateAssetTaskResponse, err error) {
    if request == nil {
        request = NewCreateAssetTaskRequest()
    }
    response = NewCreateAssetTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetsRequest() (request *CreateAssetsRequest) {
    request = &CreateAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "CreateAssets")
    return
}

func NewCreateAssetsResponse() (response *CreateAssetsResponse) {
    response = &CreateAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过资产类型和URL来创建资产
func (c *Client) CreateAssets(request *CreateAssetsRequest) (response *CreateAssetsResponse, err error) {
    if request == nil {
        request = NewCreateAssetsRequest()
    }
    response = NewCreateAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "CreateTask")
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建周期任务，定时任务等
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAssetRequest() (request *DeleteAssetRequest) {
    request = &DeleteAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DeleteAsset")
    return
}

func NewDeleteAssetResponse() (response *DeleteAssetResponse) {
    response = &DeleteAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过传递资产唯一标识符id进行资产删除
func (c *Client) DeleteAsset(request *DeleteAssetRequest) (response *DeleteAssetResponse, err error) {
    if request == nil {
        request = NewDeleteAssetRequest()
    }
    response = NewDeleteAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOperateLogsRequest() (request *DeleteOperateLogsRequest) {
    request = &DeleteOperateLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DeleteOperateLogs")
    return
}

func NewDeleteOperateLogsResponse() (response *DeleteOperateLogsResponse) {
    response = &DeleteOperateLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除特定的操作日志记录
func (c *Client) DeleteOperateLogs(request *DeleteOperateLogsRequest) (response *DeleteOperateLogsResponse, err error) {
    if request == nil {
        request = NewDeleteOperateLogsRequest()
    }
    response = NewDeleteOperateLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
    request = &DeleteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DeleteTask")
    return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
    response = &DeleteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户的各种任务
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskLogRequest() (request *DeleteTaskLogRequest) {
    request = &DeleteTaskLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DeleteTaskLog")
    return
}

func NewDeleteTaskLogResponse() (response *DeleteTaskLogResponse) {
    response = &DeleteTaskLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除任务日志，将不能再查看这次的任务运行日志
func (c *Client) DeleteTaskLog(request *DeleteTaskLogRequest) (response *DeleteTaskLogResponse, err error) {
    if request == nil {
        request = NewDeleteTaskLogRequest()
    }
    response = NewDeleteTaskLogResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWhiteUrlRequest() (request *DeleteWhiteUrlRequest) {
    request = &DeleteWhiteUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DeleteWhiteUrl")
    return
}

func NewDeleteWhiteUrlResponse() (response *DeleteWhiteUrlResponse) {
    response = &DeleteWhiteUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除URL白名单
func (c *Client) DeleteWhiteUrl(request *DeleteWhiteUrlRequest) (response *DeleteWhiteUrlResponse, err error) {
    if request == nil {
        request = NewDeleteWhiteUrlRequest()
    }
    response = NewDeleteWhiteUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetHistoryTasksRequest() (request *DescribeAssetHistoryTasksRequest) {
    request = &DescribeAssetHistoryTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeAssetHistoryTasks")
    return
}

func NewDescribeAssetHistoryTasksResponse() (response *DescribeAssetHistoryTasksResponse) {
    response = &DescribeAssetHistoryTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取资产历史任务日志的id和扫描开始时间
func (c *Client) DescribeAssetHistoryTasks(request *DescribeAssetHistoryTasksRequest) (response *DescribeAssetHistoryTasksResponse, err error) {
    if request == nil {
        request = NewDescribeAssetHistoryTasksRequest()
    }
    response = NewDescribeAssetHistoryTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetInfoByTaskIdRequest() (request *DescribeAssetInfoByTaskIdRequest) {
    request = &DescribeAssetInfoByTaskIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeAssetInfoByTaskId")
    return
}

func NewDescribeAssetInfoByTaskIdResponse() (response *DescribeAssetInfoByTaskIdResponse) {
    response = &DescribeAssetInfoByTaskIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取资产的特定任务日志的基本扫描信息
func (c *Client) DescribeAssetInfoByTaskId(request *DescribeAssetInfoByTaskIdRequest) (response *DescribeAssetInfoByTaskIdResponse, err error) {
    if request == nil {
        request = NewDescribeAssetInfoByTaskIdRequest()
    }
    response = NewDescribeAssetInfoByTaskIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetListRequest() (request *DescribeAssetListRequest) {
    request = &DescribeAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeAssetList")
    return
}

func NewDescribeAssetListResponse() (response *DescribeAssetListResponse) {
    response = &DescribeAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回资产列表和资产总数
func (c *Client) DescribeAssetList(request *DescribeAssetListRequest) (response *DescribeAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetListRequest()
    }
    response = NewDescribeAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetStatisRequest() (request *DescribeAssetStatisRequest) {
    request = &DescribeAssetStatisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeAssetStatis")
    return
}

func NewDescribeAssetStatisResponse() (response *DescribeAssetStatisResponse) {
    response = &DescribeAssetStatisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用日志ID查询资产的统计数据
func (c *Client) DescribeAssetStatis(request *DescribeAssetStatisRequest) (response *DescribeAssetStatisResponse, err error) {
    if request == nil {
        request = NewDescribeAssetStatisRequest()
    }
    response = NewDescribeAssetStatisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetStatusRequest() (request *DescribeAssetStatusRequest) {
    request = &DescribeAssetStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeAssetStatus")
    return
}

func NewDescribeAssetStatusResponse() (response *DescribeAssetStatusResponse) {
    response = &DescribeAssetStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询资产的认证状态
func (c *Client) DescribeAssetStatus(request *DescribeAssetStatusRequest) (response *DescribeAssetStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAssetStatusRequest()
    }
    response = NewDescribeAssetStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetStructureRequest() (request *DescribeAssetStructureRequest) {
    request = &DescribeAssetStructureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeAssetStructure")
    return
}

func NewDescribeAssetStructureResponse() (response *DescribeAssetStructureResponse) {
    response = &DescribeAssetStructureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取特定任务和资产下漏洞url的树状结构
func (c *Client) DescribeAssetStructure(request *DescribeAssetStructureRequest) (response *DescribeAssetStructureResponse, err error) {
    if request == nil {
        request = NewDescribeAssetStructureRequest()
    }
    response = NewDescribeAssetStructureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentVulDetailRequest() (request *DescribeContentVulDetailRequest) {
    request = &DescribeContentVulDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeContentVulDetail")
    return
}

func NewDescribeContentVulDetailResponse() (response *DescribeContentVulDetailResponse) {
    response = &DescribeContentVulDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据内容风险唯一标识以获取漏洞详情
func (c *Client) DescribeContentVulDetail(request *DescribeContentVulDetailRequest) (response *DescribeContentVulDetailResponse, err error) {
    if request == nil {
        request = NewDescribeContentVulDetailRequest()
    }
    response = NewDescribeContentVulDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentVulListRequest() (request *DescribeContentVulListRequest) {
    request = &DescribeContentVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeContentVulList")
    return
}

func NewDescribeContentVulListResponse() (response *DescribeContentVulListResponse) {
    response = &DescribeContentVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取风险管理的内容风险列表
func (c *Client) DescribeContentVulList(request *DescribeContentVulListRequest) (response *DescribeContentVulListResponse, err error) {
    if request == nil {
        request = NewDescribeContentVulListRequest()
    }
    response = NewDescribeContentVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentVulListNewRequest() (request *DescribeContentVulListNewRequest) {
    request = &DescribeContentVulListNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeContentVulListNew")
    return
}

func NewDescribeContentVulListNewResponse() (response *DescribeContentVulListNewResponse) {
    response = &DescribeContentVulListNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取风险管理的内容风险列表 支持status状态多选 支持
func (c *Client) DescribeContentVulListNew(request *DescribeContentVulListNewRequest) (response *DescribeContentVulListNewResponse, err error) {
    if request == nil {
        request = NewDescribeContentVulListNewRequest()
    }
    response = NewDescribeContentVulListNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurTaskRequest() (request *DescribeCurTaskRequest) {
    request = &DescribeCurTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeCurTask")
    return
}

func NewDescribeCurTaskResponse() (response *DescribeCurTaskResponse) {
    response = &DescribeCurTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取当前任务列表
func (c *Client) DescribeCurTask(request *DescribeCurTaskRequest) (response *DescribeCurTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCurTaskRequest()
    }
    response = NewDescribeCurTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogByTaskRequest() (request *DescribeLogByTaskRequest) {
    request = &DescribeLogByTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeLogByTask")
    return
}

func NewDescribeLogByTaskResponse() (response *DescribeLogByTaskResponse) {
    response = &DescribeLogByTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按任务查日志概况
func (c *Client) DescribeLogByTask(request *DescribeLogByTaskRequest) (response *DescribeLogByTaskResponse, err error) {
    if request == nil {
        request = NewDescribeLogByTaskRequest()
    }
    response = NewDescribeLogByTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperateLogsListRequest() (request *DescribeOperateLogsListRequest) {
    request = &DescribeOperateLogsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeOperateLogsList")
    return
}

func NewDescribeOperateLogsListResponse() (response *DescribeOperateLogsListResponse) {
    response = &DescribeOperateLogsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取操作记录列表
func (c *Client) DescribeOperateLogsList(request *DescribeOperateLogsListRequest) (response *DescribeOperateLogsListResponse, err error) {
    if request == nil {
        request = NewDescribeOperateLogsListRequest()
    }
    response = NewDescribeOperateLogsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewRequest() (request *DescribeOverviewRequest) {
    request = &DescribeOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeOverview")
    return
}

func NewDescribeOverviewResponse() (response *DescribeOverviewResponse) {
    response = &DescribeOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取总览页数据
func (c *Client) DescribeOverview(request *DescribeOverviewRequest) (response *DescribeOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewRequest()
    }
    response = NewDescribeOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePortListRequest() (request *DescribePortListRequest) {
    request = &DescribePortListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribePortList")
    return
}

func NewDescribePortListResponse() (response *DescribePortListResponse) {
    response = &DescribePortListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取风险管理的端口列表
func (c *Client) DescribePortList(request *DescribePortListRequest) (response *DescribePortListResponse, err error) {
    if request == nil {
        request = NewDescribePortListRequest()
    }
    response = NewDescribePortListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskTrendRequest() (request *DescribeRiskTrendRequest) {
    request = &DescribeRiskTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeRiskTrend")
    return
}

func NewDescribeRiskTrendResponse() (response *DescribeRiskTrendResponse) {
    response = &DescribeRiskTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 总览页查询7天/30天/90天的风险漏洞趋势
func (c *Client) DescribeRiskTrend(request *DescribeRiskTrendRequest) (response *DescribeRiskTrendResponse, err error) {
    if request == nil {
        request = NewDescribeRiskTrendRequest()
    }
    response = NewDescribeRiskTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSettingRequest() (request *DescribeSettingRequest) {
    request = &DescribeSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeSetting")
    return
}

func NewDescribeSettingResponse() (response *DescribeSettingResponse) {
    response = &DescribeSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回所有用户相关的配置
func (c *Client) DescribeSetting(request *DescribeSettingRequest) (response *DescribeSettingResponse, err error) {
    if request == nil {
        request = NewDescribeSettingRequest()
    }
    response = NewDescribeSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTCInstancesRequest() (request *DescribeTCInstancesRequest) {
    request = &DescribeTCInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeTCInstances")
    return
}

func NewDescribeTCInstancesResponse() (response *DescribeTCInstancesResponse) {
    response = &DescribeTCInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Tce的主机实例
func (c *Client) DescribeTCInstances(request *DescribeTCInstancesRequest) (response *DescribeTCInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTCInstancesRequest()
    }
    response = NewDescribeTCInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeTask")
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 显示用户的所有列表
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLogRequest() (request *DescribeTaskLogRequest) {
    request = &DescribeTaskLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeTaskLog")
    return
}

func NewDescribeTaskLogResponse() (response *DescribeTaskLogResponse) {
    response = &DescribeTaskLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回任务日志列表
func (c *Client) DescribeTaskLog(request *DescribeTaskLogRequest) (response *DescribeTaskLogResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogRequest()
    }
    response = NewDescribeTaskLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLogDetailRequest() (request *DescribeTaskLogDetailRequest) {
    request = &DescribeTaskLogDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeTaskLogDetail")
    return
}

func NewDescribeTaskLogDetailResponse() (response *DescribeTaskLogDetailResponse) {
    response = &DescribeTaskLogDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 关联任务的日志详情
func (c *Client) DescribeTaskLogDetail(request *DescribeTaskLogDetailRequest) (response *DescribeTaskLogDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogDetailRequest()
    }
    response = NewDescribeTaskLogDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskPdfRequest() (request *DescribeTaskPdfRequest) {
    request = &DescribeTaskPdfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeTaskPdf")
    return
}

func NewDescribeTaskPdfResponse() (response *DescribeTaskPdfResponse) {
    response = &DescribeTaskPdfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回当前任务日志的报表地址
func (c *Client) DescribeTaskPdf(request *DescribeTaskPdfRequest) (response *DescribeTaskPdfResponse, err error) {
    if request == nil {
        request = NewDescribeTaskPdfRequest()
    }
    response = NewDescribeTaskPdfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenRequest() (request *DescribeTokenRequest) {
    request = &DescribeTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeToken")
    return
}

func NewDescribeTokenResponse() (response *DescribeTokenResponse) {
    response = &DescribeTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 生成一个限时的token，用于下载文件或其它相关系统使用
func (c *Client) DescribeToken(request *DescribeTokenRequest) (response *DescribeTokenResponse, err error) {
    if request == nil {
        request = NewDescribeTokenRequest()
    }
    response = NewDescribeTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDetailRequest() (request *DescribeVulDetailRequest) {
    request = &DescribeVulDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeVulDetail")
    return
}

func NewDescribeVulDetailResponse() (response *DescribeVulDetailResponse) {
    response = &DescribeVulDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据漏洞唯一标识以获取漏洞详情
func (c *Client) DescribeVulDetail(request *DescribeVulDetailRequest) (response *DescribeVulDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVulDetailRequest()
    }
    response = NewDescribeVulDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulListRequest() (request *DescribeVulListRequest) {
    request = &DescribeVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeVulList")
    return
}

func NewDescribeVulListResponse() (response *DescribeVulListResponse) {
    response = &DescribeVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回风险管理的漏洞列表和漏洞总数
func (c *Client) DescribeVulList(request *DescribeVulListRequest) (response *DescribeVulListResponse, err error) {
    if request == nil {
        request = NewDescribeVulListRequest()
    }
    response = NewDescribeVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulListByTypeRequest() (request *DescribeVulListByTypeRequest) {
    request = &DescribeVulListByTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeVulListByType")
    return
}

func NewDescribeVulListByTypeResponse() (response *DescribeVulListByTypeResponse) {
    response = &DescribeVulListByTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据资产类别，返回风险管理的漏洞列表和漏洞总数
func (c *Client) DescribeVulListByType(request *DescribeVulListByTypeRequest) (response *DescribeVulListByTypeResponse, err error) {
    if request == nil {
        request = NewDescribeVulListByTypeRequest()
    }
    response = NewDescribeVulListByTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulStatusRequest() (request *DescribeVulStatusRequest) {
    request = &DescribeVulStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeVulStatus")
    return
}

func NewDescribeVulStatusResponse() (response *DescribeVulStatusResponse) {
    response = &DescribeVulStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询漏洞或内容风险的验证状态
func (c *Client) DescribeVulStatus(request *DescribeVulStatusRequest) (response *DescribeVulStatusResponse, err error) {
    if request == nil {
        request = NewDescribeVulStatusRequest()
    }
    response = NewDescribeVulStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteUrlRequest() (request *DescribeWhiteUrlRequest) {
    request = &DescribeWhiteUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "DescribeWhiteUrl")
    return
}

func NewDescribeWhiteUrlResponse() (response *DescribeWhiteUrlResponse) {
    response = &DescribeWhiteUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询白名单URL
func (c *Client) DescribeWhiteUrl(request *DescribeWhiteUrlRequest) (response *DescribeWhiteUrlResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteUrlRequest()
    }
    response = NewDescribeWhiteUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetIntelligenceDetailRequest() (request *GetIntelligenceDetailRequest) {
    request = &GetIntelligenceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "GetIntelligenceDetail")
    return
}

func NewGetIntelligenceDetailResponse() (response *GetIntelligenceDetailResponse) {
    response = &GetIntelligenceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单条威胁情报详情
func (c *Client) GetIntelligenceDetail(request *GetIntelligenceDetailRequest) (response *GetIntelligenceDetailResponse, err error) {
    if request == nil {
        request = NewGetIntelligenceDetailRequest()
    }
    response = NewGetIntelligenceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetIntelligenceListRequest() (request *GetIntelligenceListRequest) {
    request = &GetIntelligenceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "GetIntelligenceList")
    return
}

func NewGetIntelligenceListResponse() (response *GetIntelligenceListResponse) {
    response = &GetIntelligenceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获得威胁情报推送列表
func (c *Client) GetIntelligenceList(request *GetIntelligenceListRequest) (response *GetIntelligenceListResponse, err error) {
    if request == nil {
        request = NewGetIntelligenceListRequest()
    }
    response = NewGetIntelligenceListResponse()
    err = c.Send(request, response)
    return
}

func NewGetThreatIntelligencePushRequest() (request *GetThreatIntelligencePushRequest) {
    request = &GetThreatIntelligencePushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "GetThreatIntelligencePush")
    return
}

func NewGetThreatIntelligencePushResponse() (response *GetThreatIntelligencePushResponse) {
    response = &GetThreatIntelligencePushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取24小时内的威胁情报推送，默认只接收中危和高危的
func (c *Client) GetThreatIntelligencePush(request *GetThreatIntelligencePushRequest) (response *GetThreatIntelligencePushResponse, err error) {
    if request == nil {
        request = NewGetThreatIntelligencePushRequest()
    }
    response = NewGetThreatIntelligencePushResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetSimloginRequest() (request *ModifyAssetSimloginRequest) {
    request = &ModifyAssetSimloginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "ModifyAssetSimlogin")
    return
}

func NewModifyAssetSimloginResponse() (response *ModifyAssetSimloginResponse) {
    response = &ModifyAssetSimloginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 编辑模拟登录
func (c *Client) ModifyAssetSimlogin(request *ModifyAssetSimloginRequest) (response *ModifyAssetSimloginResponse, err error) {
    if request == nil {
        request = NewModifyAssetSimloginRequest()
    }
    response = NewModifyAssetSimloginResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetStatusRequest() (request *ModifyAssetStatusRequest) {
    request = &ModifyAssetStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "ModifyAssetStatus")
    return
}

func NewModifyAssetStatusResponse() (response *ModifyAssetStatusResponse) {
    response = &ModifyAssetStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量验证资产
func (c *Client) ModifyAssetStatus(request *ModifyAssetStatusRequest) (response *ModifyAssetStatusResponse, err error) {
    if request == nil {
        request = NewModifyAssetStatusRequest()
    }
    response = NewModifyAssetStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContentVulStatusRequest() (request *ModifyContentVulStatusRequest) {
    request = &ModifyContentVulStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "ModifyContentVulStatus")
    return
}

func NewModifyContentVulStatusResponse() (response *ModifyContentVulStatusResponse) {
    response = &ModifyContentVulStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改内容风险状态
func (c *Client) ModifyContentVulStatus(request *ModifyContentVulStatusRequest) (response *ModifyContentVulStatusResponse, err error) {
    if request == nil {
        request = NewModifyContentVulStatusRequest()
    }
    response = NewModifyContentVulStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySettingRequest() (request *ModifySettingRequest) {
    request = &ModifySettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "ModifySetting")
    return
}

func NewModifySettingResponse() (response *ModifySettingResponse) {
    response = &ModifySettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改用户配置
func (c *Client) ModifySetting(request *ModifySettingRequest) (response *ModifySettingResponse, err error) {
    if request == nil {
        request = NewModifySettingRequest()
    }
    response = NewModifySettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskRequest() (request *ModifyTaskRequest) {
    request = &ModifyTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "ModifyTask")
    return
}

func NewModifyTaskResponse() (response *ModifyTaskResponse) {
    response = &ModifyTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更改状态，标记删除
func (c *Client) ModifyTask(request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
    if request == nil {
        request = NewModifyTaskRequest()
    }
    response = NewModifyTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVulStatusRequest() (request *ModifyVulStatusRequest) {
    request = &ModifyVulStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "ModifyVulStatus")
    return
}

func NewModifyVulStatusResponse() (response *ModifyVulStatusResponse) {
    response = &ModifyVulStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改漏洞状态
func (c *Client) ModifyVulStatus(request *ModifyVulStatusRequest) (response *ModifyVulStatusResponse, err error) {
    if request == nil {
        request = NewModifyVulStatusRequest()
    }
    response = NewModifyVulStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWhiteUrlRequest() (request *ModifyWhiteUrlRequest) {
    request = &ModifyWhiteUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "ModifyWhiteUrl")
    return
}

func NewModifyWhiteUrlResponse() (response *ModifyWhiteUrlResponse) {
    response = &ModifyWhiteUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改白名单URL
func (c *Client) ModifyWhiteUrl(request *ModifyWhiteUrlRequest) (response *ModifyWhiteUrlResponse, err error) {
    if request == nil {
        request = NewModifyWhiteUrlRequest()
    }
    response = NewModifyWhiteUrlResponse()
    err = c.Send(request, response)
    return
}

func NewRestartAssetTaskRequest() (request *RestartAssetTaskRequest) {
    request = &RestartAssetTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "RestartAssetTask")
    return
}

func NewRestartAssetTaskResponse() (response *RestartAssetTaskResponse) {
    response = &RestartAssetTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重启曾被中断的扫描任务
func (c *Client) RestartAssetTask(request *RestartAssetTaskRequest) (response *RestartAssetTaskResponse, err error) {
    if request == nil {
        request = NewRestartAssetTaskRequest()
    }
    response = NewRestartAssetTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStartIntelligenceScanRequest() (request *StartIntelligenceScanRequest) {
    request = &StartIntelligenceScanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "StartIntelligenceScan")
    return
}

func NewStartIntelligenceScanResponse() (response *StartIntelligenceScanResponse) {
    response = &StartIntelligenceScanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发起针对特定威胁情报的扫描任务
func (c *Client) StartIntelligenceScan(request *StartIntelligenceScanRequest) (response *StartIntelligenceScanResponse, err error) {
    if request == nil {
        request = NewStartIntelligenceScanRequest()
    }
    response = NewStartIntelligenceScanResponse()
    err = c.Send(request, response)
    return
}

func NewStopRunningTaskRequest() (request *StopRunningTaskRequest) {
    request = &StopRunningTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tngtvm", APIVersion, "StopRunningTask")
    return
}

func NewStopRunningTaskResponse() (response *StopRunningTaskResponse) {
    response = &StopRunningTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止运行中的任务
func (c *Client) StopRunningTask(request *StopRunningTaskRequest) (response *StopRunningTaskResponse, err error) {
    if request == nil {
        request = NewStopRunningTaskRequest()
    }
    response = NewStopRunningTaskResponse()
    err = c.Send(request, response)
    return
}
