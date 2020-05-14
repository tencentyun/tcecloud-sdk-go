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

package v20190606

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-06"

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


func NewAddFavoriteRequest() (request *AddFavoriteRequest) {
    request = &AddFavoriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "AddFavorite")
    return
}

func NewAddFavoriteResponse() (response *AddFavoriteResponse) {
    response = &AddFavoriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus新增收藏夹
func (c *Client) AddFavorite(request *AddFavoriteRequest) (response *AddFavoriteResponse, err error) {
    if request == nil {
        request = NewAddFavoriteRequest()
    }
    response = NewAddFavoriteResponse()
    err = c.Send(request, response)
    return
}

func NewAddNamespaceRequest() (request *AddNamespaceRequest) {
    request = &AddNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "AddNamespace")
    return
}

func NewAddNamespaceResponse() (response *AddNamespaceResponse) {
    response = &AddNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus新增namespace
func (c *Client) AddNamespace(request *AddNamespaceRequest) (response *AddNamespaceResponse, err error) {
    if request == nil {
        request = NewAddNamespaceRequest()
    }
    response = NewAddNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewAddPolicyRequest() (request *AddPolicyRequest) {
    request = &AddPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "AddPolicy")
    return
}

func NewAddPolicyResponse() (response *AddPolicyResponse) {
    response = &AddPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus新增告警策略
func (c *Client) AddPolicy(request *AddPolicyRequest) (response *AddPolicyResponse, err error) {
    if request == nil {
        request = NewAddPolicyRequest()
    }
    response = NewAddPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewAddViewRequest() (request *AddViewRequest) {
    request = &AddViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "AddView")
    return
}

func NewAddViewResponse() (response *AddViewResponse) {
    response = &AddViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus新增视图
func (c *Client) AddView(request *AddViewRequest) (response *AddViewResponse, err error) {
    if request == nil {
        request = NewAddViewRequest()
    }
    response = NewAddViewResponse()
    err = c.Send(request, response)
    return
}

func NewArgusReportRequest() (request *ArgusReportRequest) {
    request = &ArgusReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ArgusReport")
    return
}

func NewArgusReportResponse() (response *ArgusReportResponse) {
    response = &ArgusReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//     （1）提供Line Protocol和Json两种数据格式上报。
//     （2）单用户上报请求数限制为100次/s，单个请求最大包体为256K，超过大小返回错误，最多上报150条，超过限制的数据自动丢弃。
//     （3）上报接口校验上报的时间戳参数，30天前和一天之后的数据会被丢弃。
//     （4）忽略上报的无效维度和指标，未上报的维度按空字符串进行存储。
//     （5）支持单个请求同时上报多个命名空间，上报的错误命名空间数据会进行丢弃，返回结果中显示上报成功数和失败数。
func (c *Client) ArgusReport(request *ArgusReportRequest) (response *ArgusReportResponse, err error) {
    if request == nil {
        request = NewArgusReportRequest()
    }
    response = NewArgusReportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFavoriteRequest() (request *CreateFavoriteRequest) {
    request = &CreateFavoriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "CreateFavorite")
    return
}

func NewCreateFavoriteResponse() (response *CreateFavoriteResponse) {
    response = &CreateFavoriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增收藏夹
func (c *Client) CreateFavorite(request *CreateFavoriteRequest) (response *CreateFavoriteResponse, err error) {
    if request == nil {
        request = NewCreateFavoriteRequest()
    }
    response = NewCreateFavoriteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "CreateNamespace")
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增namespace
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyRequest() (request *CreatePolicyRequest) {
    request = &CreatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "CreatePolicy")
    return
}

func NewCreatePolicyResponse() (response *CreatePolicyResponse) {
    response = &CreatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增告警策略
func (c *Client) CreatePolicy(request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    if request == nil {
        request = NewCreatePolicyRequest()
    }
    response = NewCreatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTkePolicyRequest() (request *CreateTkePolicyRequest) {
    request = &CreateTkePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "CreateTkePolicy")
    return
}

func NewCreateTkePolicyResponse() (response *CreateTkePolicyResponse) {
    response = &CreateTkePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增告警策略
func (c *Client) CreateTkePolicy(request *CreateTkePolicyRequest) (response *CreateTkePolicyResponse, err error) {
    if request == nil {
        request = NewCreateTkePolicyRequest()
    }
    response = NewCreateTkePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateViewRequest() (request *CreateViewRequest) {
    request = &CreateViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "CreateView")
    return
}

func NewCreateViewResponse() (response *CreateViewResponse) {
    response = &CreateViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增视图
func (c *Client) CreateView(request *CreateViewRequest) (response *CreateViewResponse, err error) {
    if request == nil {
        request = NewCreateViewRequest()
    }
    response = NewCreateViewResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFavoriteRequest() (request *DeleteFavoriteRequest) {
    request = &DeleteFavoriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DeleteFavorite")
    return
}

func NewDeleteFavoriteResponse() (response *DeleteFavoriteResponse) {
    response = &DeleteFavoriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除收藏夹
func (c *Client) DeleteFavorite(request *DeleteFavoriteRequest) (response *DeleteFavoriteResponse, err error) {
    if request == nil {
        request = NewDeleteFavoriteRequest()
    }
    response = NewDeleteFavoriteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DeleteNamespace")
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除namespace
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyRequest() (request *DeletePolicyRequest) {
    request = &DeletePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DeletePolicy")
    return
}

func NewDeletePolicyResponse() (response *DeletePolicyResponse) {
    response = &DeletePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除告警策略
func (c *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
    if request == nil {
        request = NewDeletePolicyRequest()
    }
    response = NewDeletePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTkePolicyRequest() (request *DeleteTkePolicyRequest) {
    request = &DeleteTkePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DeleteTkePolicy")
    return
}

func NewDeleteTkePolicyResponse() (response *DeleteTkePolicyResponse) {
    response = &DeleteTkePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除告警策略
func (c *Client) DeleteTkePolicy(request *DeleteTkePolicyRequest) (response *DeleteTkePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteTkePolicyRequest()
    }
    response = NewDeleteTkePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteViewRequest() (request *DeleteViewRequest) {
    request = &DeleteViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DeleteView")
    return
}

func NewDeleteViewResponse() (response *DeleteViewResponse) {
    response = &DeleteViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除视图
func (c *Client) DeleteView(request *DeleteViewRequest) (response *DeleteViewResponse, err error) {
    if request == nil {
        request = NewDeleteViewRequest()
    }
    response = NewDeleteViewResponse()
    err = c.Send(request, response)
    return
}

func NewDescFavoriteRequest() (request *DescFavoriteRequest) {
    request = &DescFavoriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescFavorite")
    return
}

func NewDescFavoriteResponse() (response *DescFavoriteResponse) {
    response = &DescFavoriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus获取收藏夹详情
func (c *Client) DescFavorite(request *DescFavoriteRequest) (response *DescFavoriteResponse, err error) {
    if request == nil {
        request = NewDescFavoriteRequest()
    }
    response = NewDescFavoriteResponse()
    err = c.Send(request, response)
    return
}

func NewDescNamespaceRequest() (request *DescNamespaceRequest) {
    request = &DescNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescNamespace")
    return
}

func NewDescNamespaceResponse() (response *DescNamespaceResponse) {
    response = &DescNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus获取namespace详情
func (c *Client) DescNamespace(request *DescNamespaceRequest) (response *DescNamespaceResponse, err error) {
    if request == nil {
        request = NewDescNamespaceRequest()
    }
    response = NewDescNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDescPolicyRequest() (request *DescPolicyRequest) {
    request = &DescPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescPolicy")
    return
}

func NewDescPolicyResponse() (response *DescPolicyResponse) {
    response = &DescPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus获取告警策略详情
func (c *Client) DescPolicy(request *DescPolicyRequest) (response *DescPolicyResponse, err error) {
    if request == nil {
        request = NewDescPolicyRequest()
    }
    response = NewDescPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescViewRequest() (request *DescViewRequest) {
    request = &DescViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescView")
    return
}

func NewDescViewResponse() (response *DescViewResponse) {
    response = &DescViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus获取视图详情
func (c *Client) DescView(request *DescViewRequest) (response *DescViewResponse, err error) {
    if request == nil {
        request = NewDescViewRequest()
    }
    response = NewDescViewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertHistoryListRequest() (request *DescribeAlertHistoryListRequest) {
    request = &DescribeAlertHistoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeAlertHistoryList")
    return
}

func NewDescribeAlertHistoryListResponse() (response *DescribeAlertHistoryListResponse) {
    response = &DescribeAlertHistoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询告警历史
func (c *Client) DescribeAlertHistoryList(request *DescribeAlertHistoryListRequest) (response *DescribeAlertHistoryListResponse, err error) {
    if request == nil {
        request = NewDescribeAlertHistoryListRequest()
    }
    response = NewDescribeAlertHistoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDimensionListRequest() (request *DescribeDimensionListRequest) {
    request = &DescribeDimensionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeDimensionList")
    return
}

func NewDescribeDimensionListResponse() (response *DescribeDimensionListResponse) {
    response = &DescribeDimensionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取维度下拉框数据
func (c *Client) DescribeDimensionList(request *DescribeDimensionListRequest) (response *DescribeDimensionListResponse, err error) {
    if request == nil {
        request = NewDescribeDimensionListRequest()
    }
    response = NewDescribeDimensionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFavoriteRequest() (request *DescribeFavoriteRequest) {
    request = &DescribeFavoriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeFavorite")
    return
}

func NewDescribeFavoriteResponse() (response *DescribeFavoriteResponse) {
    response = &DescribeFavoriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取收藏夹详情
func (c *Client) DescribeFavorite(request *DescribeFavoriteRequest) (response *DescribeFavoriteResponse, err error) {
    if request == nil {
        request = NewDescribeFavoriteRequest()
    }
    response = NewDescribeFavoriteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFavoriteListRequest() (request *DescribeFavoriteListRequest) {
    request = &DescribeFavoriteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeFavoriteList")
    return
}

func NewDescribeFavoriteListResponse() (response *DescribeFavoriteListResponse) {
    response = &DescribeFavoriteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询收藏夹
func (c *Client) DescribeFavoriteList(request *DescribeFavoriteListRequest) (response *DescribeFavoriteListResponse, err error) {
    if request == nil {
        request = NewDescribeFavoriteListRequest()
    }
    response = NewDescribeFavoriteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespaceRequest() (request *DescribeNamespaceRequest) {
    request = &DescribeNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeNamespace")
    return
}

func NewDescribeNamespaceResponse() (response *DescribeNamespaceResponse) {
    response = &DescribeNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取namespace详情
func (c *Client) DescribeNamespace(request *DescribeNamespaceRequest) (response *DescribeNamespaceResponse, err error) {
    if request == nil {
        request = NewDescribeNamespaceRequest()
    }
    response = NewDescribeNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespaceListRequest() (request *DescribeNamespaceListRequest) {
    request = &DescribeNamespaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeNamespaceList")
    return
}

func NewDescribeNamespaceListResponse() (response *DescribeNamespaceListResponse) {
    response = &DescribeNamespaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询namespace
func (c *Client) DescribeNamespaceList(request *DescribeNamespaceListRequest) (response *DescribeNamespaceListResponse, err error) {
    if request == nil {
        request = NewDescribeNamespaceListRequest()
    }
    response = NewDescribeNamespaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewRequest() (request *DescribeOverviewRequest) {
    request = &DescribeOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeOverview")
    return
}

func NewDescribeOverviewResponse() (response *DescribeOverviewResponse) {
    response = &DescribeOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询概览
func (c *Client) DescribeOverview(request *DescribeOverviewRequest) (response *DescribeOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewRequest()
    }
    response = NewDescribeOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamListRequest() (request *DescribeParamListRequest) {
    request = &DescribeParamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeParamList")
    return
}

func NewDescribeParamListResponse() (response *DescribeParamListResponse) {
    response = &DescribeParamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取下拉框参数
func (c *Client) DescribeParamList(request *DescribeParamListRequest) (response *DescribeParamListResponse, err error) {
    if request == nil {
        request = NewDescribeParamListRequest()
    }
    response = NewDescribeParamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyRequest() (request *DescribePolicyRequest) {
    request = &DescribePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribePolicy")
    return
}

func NewDescribePolicyResponse() (response *DescribePolicyResponse) {
    response = &DescribePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警策略详情
func (c *Client) DescribePolicy(request *DescribePolicyRequest) (response *DescribePolicyResponse, err error) {
    if request == nil {
        request = NewDescribePolicyRequest()
    }
    response = NewDescribePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyListRequest() (request *DescribePolicyListRequest) {
    request = &DescribePolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribePolicyList")
    return
}

func NewDescribePolicyListResponse() (response *DescribePolicyListResponse) {
    response = &DescribePolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询告警策略
func (c *Client) DescribePolicyList(request *DescribePolicyListRequest) (response *DescribePolicyListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyListRequest()
    }
    response = NewDescribePolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceTopologyRequest() (request *DescribeServiceTopologyRequest) {
    request = &DescribeServiceTopologyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeServiceTopology")
    return
}

func NewDescribeServiceTopologyResponse() (response *DescribeServiceTopologyResponse) {
    response = &DescribeServiceTopologyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务网格拓扑
func (c *Client) DescribeServiceTopology(request *DescribeServiceTopologyRequest) (response *DescribeServiceTopologyResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTopologyRequest()
    }
    response = NewDescribeServiceTopologyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceTopologyJobRequest() (request *DescribeServiceTopologyJobRequest) {
    request = &DescribeServiceTopologyJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeServiceTopologyJob")
    return
}

func NewDescribeServiceTopologyJobResponse() (response *DescribeServiceTopologyJobResponse) {
    response = &DescribeServiceTopologyJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 异步查询服务网格拓扑任务
func (c *Client) DescribeServiceTopologyJob(request *DescribeServiceTopologyJobRequest) (response *DescribeServiceTopologyJobResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTopologyJobRequest()
    }
    response = NewDescribeServiceTopologyJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceTopologyResultRequest() (request *DescribeServiceTopologyResultRequest) {
    request = &DescribeServiceTopologyResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeServiceTopologyResult")
    return
}

func NewDescribeServiceTopologyResultResponse() (response *DescribeServiceTopologyResultResponse) {
    response = &DescribeServiceTopologyResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 异步查询服务网格拓扑结果
func (c *Client) DescribeServiceTopologyResult(request *DescribeServiceTopologyResultRequest) (response *DescribeServiceTopologyResultResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTopologyResultRequest()
    }
    response = NewDescribeServiceTopologyResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTkeAlertHistoryListRequest() (request *DescribeTkeAlertHistoryListRequest) {
    request = &DescribeTkeAlertHistoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeTkeAlertHistoryList")
    return
}

func NewDescribeTkeAlertHistoryListResponse() (response *DescribeTkeAlertHistoryListResponse) {
    response = &DescribeTkeAlertHistoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询告警历史
func (c *Client) DescribeTkeAlertHistoryList(request *DescribeTkeAlertHistoryListRequest) (response *DescribeTkeAlertHistoryListResponse, err error) {
    if request == nil {
        request = NewDescribeTkeAlertHistoryListRequest()
    }
    response = NewDescribeTkeAlertHistoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTkeAlertMetricListRequest() (request *DescribeTkeAlertMetricListRequest) {
    request = &DescribeTkeAlertMetricListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeTkeAlertMetricList")
    return
}

func NewDescribeTkeAlertMetricListResponse() (response *DescribeTkeAlertMetricListResponse) {
    response = &DescribeTkeAlertMetricListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警指标
func (c *Client) DescribeTkeAlertMetricList(request *DescribeTkeAlertMetricListRequest) (response *DescribeTkeAlertMetricListResponse, err error) {
    if request == nil {
        request = NewDescribeTkeAlertMetricListRequest()
    }
    response = NewDescribeTkeAlertMetricListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTkeClusterListRequest() (request *DescribeTkeClusterListRequest) {
    request = &DescribeTkeClusterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeTkeClusterList")
    return
}

func NewDescribeTkeClusterListResponse() (response *DescribeTkeClusterListResponse) {
    response = &DescribeTkeClusterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TKE集群详情
func (c *Client) DescribeTkeClusterList(request *DescribeTkeClusterListRequest) (response *DescribeTkeClusterListResponse, err error) {
    if request == nil {
        request = NewDescribeTkeClusterListRequest()
    }
    response = NewDescribeTkeClusterListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTkePolicyRequest() (request *DescribeTkePolicyRequest) {
    request = &DescribeTkePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeTkePolicy")
    return
}

func NewDescribeTkePolicyResponse() (response *DescribeTkePolicyResponse) {
    response = &DescribeTkePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取告警策略详情
func (c *Client) DescribeTkePolicy(request *DescribeTkePolicyRequest) (response *DescribeTkePolicyResponse, err error) {
    if request == nil {
        request = NewDescribeTkePolicyRequest()
    }
    response = NewDescribeTkePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTkePolicyListRequest() (request *DescribeTkePolicyListRequest) {
    request = &DescribeTkePolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeTkePolicyList")
    return
}

func NewDescribeTkePolicyListResponse() (response *DescribeTkePolicyListResponse) {
    response = &DescribeTkePolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询告警策略
func (c *Client) DescribeTkePolicyList(request *DescribeTkePolicyListRequest) (response *DescribeTkePolicyListResponse, err error) {
    if request == nil {
        request = NewDescribeTkePolicyListRequest()
    }
    response = NewDescribeTkePolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeViewRequest() (request *DescribeViewRequest) {
    request = &DescribeViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeView")
    return
}

func NewDescribeViewResponse() (response *DescribeViewResponse) {
    response = &DescribeViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取视图详情
func (c *Client) DescribeView(request *DescribeViewRequest) (response *DescribeViewResponse, err error) {
    if request == nil {
        request = NewDescribeViewRequest()
    }
    response = NewDescribeViewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeViewListRequest() (request *DescribeViewListRequest) {
    request = &DescribeViewListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "DescribeViewList")
    return
}

func NewDescribeViewListResponse() (response *DescribeViewListResponse) {
    response = &DescribeViewListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询视图
func (c *Client) DescribeViewList(request *DescribeViewListRequest) (response *DescribeViewListResponse, err error) {
    if request == nil {
        request = NewDescribeViewListRequest()
    }
    response = NewDescribeViewListResponse()
    err = c.Send(request, response)
    return
}

func NewGetAlertHistoriesRequest() (request *GetAlertHistoriesRequest) {
    request = &GetAlertHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetAlertHistories")
    return
}

func NewGetAlertHistoriesResponse() (response *GetAlertHistoriesResponse) {
    response = &GetAlertHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus查询告警历史
func (c *Client) GetAlertHistories(request *GetAlertHistoriesRequest) (response *GetAlertHistoriesResponse, err error) {
    if request == nil {
        request = NewGetAlertHistoriesRequest()
    }
    response = NewGetAlertHistoriesResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataRequest() (request *GetDataRequest) {
    request = &GetDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetData")
    return
}

func NewGetDataResponse() (response *GetDataResponse) {
    response = &GetDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询数据
func (c *Client) GetData(request *GetDataRequest) (response *GetDataResponse, err error) {
    if request == nil {
        request = NewGetDataRequest()
    }
    response = NewGetDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataJobRequest() (request *GetDataJobRequest) {
    request = &GetDataJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetDataJob")
    return
}

func NewGetDataJobResponse() (response *GetDataJobResponse) {
    response = &GetDataJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 异步查询数据任务
func (c *Client) GetDataJob(request *GetDataJobRequest) (response *GetDataJobResponse, err error) {
    if request == nil {
        request = NewGetDataJobRequest()
    }
    response = NewGetDataJobResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataResultRequest() (request *GetDataResultRequest) {
    request = &GetDataResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetDataResult")
    return
}

func NewGetDataResultResponse() (response *GetDataResultResponse) {
    response = &GetDataResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 异步查询数据结果
func (c *Client) GetDataResult(request *GetDataResultRequest) (response *GetDataResultResponse, err error) {
    if request == nil {
        request = NewGetDataResultRequest()
    }
    response = NewGetDataResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetDimensionsRequest() (request *GetDimensionsRequest) {
    request = &GetDimensionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetDimensions")
    return
}

func NewGetDimensionsResponse() (response *GetDimensionsResponse) {
    response = &GetDimensionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus获取维度下拉框数据
func (c *Client) GetDimensions(request *GetDimensionsRequest) (response *GetDimensionsResponse, err error) {
    if request == nil {
        request = NewGetDimensionsRequest()
    }
    response = NewGetDimensionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetFavoritesRequest() (request *GetFavoritesRequest) {
    request = &GetFavoritesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetFavorites")
    return
}

func NewGetFavoritesResponse() (response *GetFavoritesResponse) {
    response = &GetFavoritesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus查询收藏夹
func (c *Client) GetFavorites(request *GetFavoritesRequest) (response *GetFavoritesResponse, err error) {
    if request == nil {
        request = NewGetFavoritesRequest()
    }
    response = NewGetFavoritesResponse()
    err = c.Send(request, response)
    return
}

func NewGetNamespacesRequest() (request *GetNamespacesRequest) {
    request = &GetNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetNamespaces")
    return
}

func NewGetNamespacesResponse() (response *GetNamespacesResponse) {
    response = &GetNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus查询namespace
func (c *Client) GetNamespaces(request *GetNamespacesRequest) (response *GetNamespacesResponse, err error) {
    if request == nil {
        request = NewGetNamespacesRequest()
    }
    response = NewGetNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewGetOverviewRequest() (request *GetOverviewRequest) {
    request = &GetOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetOverview")
    return
}

func NewGetOverviewResponse() (response *GetOverviewResponse) {
    response = &GetOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus查询概览
func (c *Client) GetOverview(request *GetOverviewRequest) (response *GetOverviewResponse, err error) {
    if request == nil {
        request = NewGetOverviewRequest()
    }
    response = NewGetOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewGetParamsRequest() (request *GetParamsRequest) {
    request = &GetParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetParams")
    return
}

func NewGetParamsResponse() (response *GetParamsResponse) {
    response = &GetParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus获取下拉框参数
func (c *Client) GetParams(request *GetParamsRequest) (response *GetParamsResponse, err error) {
    if request == nil {
        request = NewGetParamsRequest()
    }
    response = NewGetParamsResponse()
    err = c.Send(request, response)
    return
}

func NewGetPoliciesRequest() (request *GetPoliciesRequest) {
    request = &GetPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetPolicies")
    return
}

func NewGetPoliciesResponse() (response *GetPoliciesResponse) {
    response = &GetPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus查询告警策略
func (c *Client) GetPolicies(request *GetPoliciesRequest) (response *GetPoliciesResponse, err error) {
    if request == nil {
        request = NewGetPoliciesRequest()
    }
    response = NewGetPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTkeDataRequest() (request *GetTkeDataRequest) {
    request = &GetTkeDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetTkeData")
    return
}

func NewGetTkeDataResponse() (response *GetTkeDataResponse) {
    response = &GetTkeDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TKE数据
func (c *Client) GetTkeData(request *GetTkeDataRequest) (response *GetTkeDataResponse, err error) {
    if request == nil {
        request = NewGetTkeDataRequest()
    }
    response = NewGetTkeDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetTkeDataJobRequest() (request *GetTkeDataJobRequest) {
    request = &GetTkeDataJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetTkeDataJob")
    return
}

func NewGetTkeDataJobResponse() (response *GetTkeDataJobResponse) {
    response = &GetTkeDataJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 异步查询TKE数据任务
func (c *Client) GetTkeDataJob(request *GetTkeDataJobRequest) (response *GetTkeDataJobResponse, err error) {
    if request == nil {
        request = NewGetTkeDataJobRequest()
    }
    response = NewGetTkeDataJobResponse()
    err = c.Send(request, response)
    return
}

func NewGetTkeDataResultRequest() (request *GetTkeDataResultRequest) {
    request = &GetTkeDataResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetTkeDataResult")
    return
}

func NewGetTkeDataResultResponse() (response *GetTkeDataResultResponse) {
    response = &GetTkeDataResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 异步查询TKE数据结果
func (c *Client) GetTkeDataResult(request *GetTkeDataResultRequest) (response *GetTkeDataResultResponse, err error) {
    if request == nil {
        request = NewGetTkeDataResultRequest()
    }
    response = NewGetTkeDataResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetViewsRequest() (request *GetViewsRequest) {
    request = &GetViewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "GetViews")
    return
}

func NewGetViewsResponse() (response *GetViewsResponse) {
    response = &GetViewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus查询视图
func (c *Client) GetViews(request *GetViewsRequest) (response *GetViewsResponse, err error) {
    if request == nil {
        request = NewGetViewsRequest()
    }
    response = NewGetViewsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFavoriteRequest() (request *ModifyFavoriteRequest) {
    request = &ModifyFavoriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ModifyFavorite")
    return
}

func NewModifyFavoriteResponse() (response *ModifyFavoriteResponse) {
    response = &ModifyFavoriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新收藏夹
func (c *Client) ModifyFavorite(request *ModifyFavoriteRequest) (response *ModifyFavoriteResponse, err error) {
    if request == nil {
        request = NewModifyFavoriteRequest()
    }
    response = NewModifyFavoriteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNamespaceRequest() (request *ModifyNamespaceRequest) {
    request = &ModifyNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ModifyNamespace")
    return
}

func NewModifyNamespaceResponse() (response *ModifyNamespaceResponse) {
    response = &ModifyNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新namespace
func (c *Client) ModifyNamespace(request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceRequest()
    }
    response = NewModifyNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPolicyRequest() (request *ModifyPolicyRequest) {
    request = &ModifyPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ModifyPolicy")
    return
}

func NewModifyPolicyResponse() (response *ModifyPolicyResponse) {
    response = &ModifyPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新告警策略
func (c *Client) ModifyPolicy(request *ModifyPolicyRequest) (response *ModifyPolicyResponse, err error) {
    if request == nil {
        request = NewModifyPolicyRequest()
    }
    response = NewModifyPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTkePolicyRequest() (request *ModifyTkePolicyRequest) {
    request = &ModifyTkePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ModifyTkePolicy")
    return
}

func NewModifyTkePolicyResponse() (response *ModifyTkePolicyResponse) {
    response = &ModifyTkePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新告警策略
func (c *Client) ModifyTkePolicy(request *ModifyTkePolicyRequest) (response *ModifyTkePolicyResponse, err error) {
    if request == nil {
        request = NewModifyTkePolicyRequest()
    }
    response = NewModifyTkePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyViewRequest() (request *ModifyViewRequest) {
    request = &ModifyViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ModifyView")
    return
}

func NewModifyViewResponse() (response *ModifyViewResponse) {
    response = &ModifyViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新视图
func (c *Client) ModifyView(request *ModifyViewRequest) (response *ModifyViewResponse, err error) {
    if request == nil {
        request = NewModifyViewRequest()
    }
    response = NewModifyViewResponse()
    err = c.Send(request, response)
    return
}

func NewReportDataRequest() (request *ReportDataRequest) {
    request = &ReportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ReportData")
    return
}

func NewReportDataResponse() (response *ReportDataResponse) {
    response = &ReportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//     （1）提供Line Protocol和Json两种数据格式上报。
//     （2）单用户上报请求数限制为100次/s，单个请求最大包体为256K，超过大小返回错误，最多上报150条，超过限制的数据自动丢弃。
//     （3）上报接口校验上报的时间戳参数，30天前和一天之后的数据会被丢弃。
//     （4）忽略上报的无效维度和指标，未上报的维度按空字符串进行存储。
//     （5）支持单个请求同时上报多个命名空间，上报的错误命名空间数据会进行丢弃，返回结果中显示上报成功数和失败数。
func (c *Client) ReportData(request *ReportDataRequest) (response *ReportDataResponse, err error) {
    if request == nil {
        request = NewReportDataRequest()
    }
    response = NewReportDataResponse()
    err = c.Send(request, response)
    return
}

func NewServiceCallRelationRequest() (request *ServiceCallRelationRequest) {
    request = &ServiceCallRelationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ServiceCallRelation")
    return
}

func NewServiceCallRelationResponse() (response *ServiceCallRelationResponse) {
    response = &ServiceCallRelationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务网格拓扑
func (c *Client) ServiceCallRelation(request *ServiceCallRelationRequest) (response *ServiceCallRelationResponse, err error) {
    if request == nil {
        request = NewServiceCallRelationRequest()
    }
    response = NewServiceCallRelationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFavoriteRequest() (request *UpdateFavoriteRequest) {
    request = &UpdateFavoriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "UpdateFavorite")
    return
}

func NewUpdateFavoriteResponse() (response *UpdateFavoriteResponse) {
    response = &UpdateFavoriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus更新收藏夹
func (c *Client) UpdateFavorite(request *UpdateFavoriteRequest) (response *UpdateFavoriteResponse, err error) {
    if request == nil {
        request = NewUpdateFavoriteRequest()
    }
    response = NewUpdateFavoriteResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNamespaceRequest() (request *UpdateNamespaceRequest) {
    request = &UpdateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "UpdateNamespace")
    return
}

func NewUpdateNamespaceResponse() (response *UpdateNamespaceResponse) {
    response = &UpdateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus更新namespace
func (c *Client) UpdateNamespace(request *UpdateNamespaceRequest) (response *UpdateNamespaceResponse, err error) {
    if request == nil {
        request = NewUpdateNamespaceRequest()
    }
    response = NewUpdateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
    request = &UpdatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "UpdatePolicy")
    return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
    response = &UpdatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus更新告警策略
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateViewRequest() (request *UpdateViewRequest) {
    request = &UpdateViewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "UpdateView")
    return
}

func NewUpdateViewResponse() (response *UpdateViewResponse) {
    response = &UpdateViewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Argus更新视图
func (c *Client) UpdateView(request *UpdateViewRequest) (response *UpdateViewResponse, err error) {
    if request == nil {
        request = NewUpdateViewRequest()
    }
    response = NewUpdateViewResponse()
    err = c.Send(request, response)
    return
}
