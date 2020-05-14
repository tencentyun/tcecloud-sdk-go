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

package v20181025

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-10-25"

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


func NewAddQuotaRequest() (request *AddQuotaRequest) {
    request = &AddQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "AddQuota")
    return
}

func NewAddQuotaResponse() (response *AddQuotaResponse) {
    response = &AddQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增,修改quota配置  支持一次传入多个配额配置。
// 
// quotaKey需符合四层定义
func (c *Client) AddQuota(request *AddQuotaRequest) (response *AddQuotaResponse, err error) {
    if request == nil {
        request = NewAddQuotaRequest()
    }
    response = NewAddQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountListRequest() (request *DescribeAccountListRequest) {
    request = &DescribeAccountListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeAccountList")
    return
}

func NewDescribeAccountListResponse() (response *DescribeAccountListResponse) {
    response = &DescribeAccountListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理端查询帐户列表和用户代金券总金额
func (c *Client) DescribeAccountList(request *DescribeAccountListRequest) (response *DescribeAccountListResponse, err error) {
    if request == nil {
        request = NewDescribeAccountListRequest()
    }
    response = NewDescribeAccountListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountWaterRequest() (request *DescribeAccountWaterRequest) {
    request = &DescribeAccountWaterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeAccountWater")
    return
}

func NewDescribeAccountWaterResponse() (response *DescribeAccountWaterResponse) {
    response = &DescribeAccountWaterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理端，查询账户的还款或销账流水
func (c *Client) DescribeAccountWater(request *DescribeAccountWaterRequest) (response *DescribeAccountWaterResponse, err error) {
    if request == nil {
        request = NewDescribeAccountWaterRequest()
    }
    response = NewDescribeAccountWaterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicAccountListRequest() (request *DescribeBasicAccountListRequest) {
    request = &DescribeBasicAccountListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBasicAccountList")
    return
}

func NewDescribeBasicAccountListResponse() (response *DescribeBasicAccountListResponse) {
    response = &DescribeBasicAccountListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据uin批量获取账户基本信息
func (c *Client) DescribeBasicAccountList(request *DescribeBasicAccountListRequest) (response *DescribeBasicAccountListResponse, err error) {
    if request == nil {
        request = NewDescribeBasicAccountListRequest()
    }
    response = NewDescribeBasicAccountListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailByResourceRequest() (request *DescribeBillDetailByResourceRequest) {
    request = &DescribeBillDetailByResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDetailByResource")
    return
}

func NewDescribeBillDetailByResourceResponse() (response *DescribeBillDetailByResourceResponse) {
    response = &DescribeBillDetailByResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取资源花费详情
func (c *Client) DescribeBillDetailByResource(request *DescribeBillDetailByResourceRequest) (response *DescribeBillDetailByResourceResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailByResourceRequest()
    }
    response = NewDescribeBillDetailByResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDownloadByResourceRequest() (request *DescribeBillDownloadByResourceRequest) {
    request = &DescribeBillDownloadByResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadByResource")
    return
}

func NewDescribeBillDownloadByResourceResponse() (response *DescribeBillDownloadByResourceResponse) {
    response = &DescribeBillDownloadByResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载定制化账单按资源汇总
func (c *Client) DescribeBillDownloadByResource(request *DescribeBillDownloadByResourceRequest) (response *DescribeBillDownloadByResourceResponse, err error) {
    if request == nil {
        request = NewDescribeBillDownloadByResourceRequest()
    }
    response = NewDescribeBillDownloadByResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDownloadCommonSummaryRequest() (request *DescribeBillDownloadCommonSummaryRequest) {
    request = &DescribeBillDownloadCommonSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadCommonSummary")
    return
}

func NewDescribeBillDownloadCommonSummaryResponse() (response *DescribeBillDownloadCommonSummaryResponse) {
    response = &DescribeBillDownloadCommonSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载账单通用汇总（产品+地域）
func (c *Client) DescribeBillDownloadCommonSummary(request *DescribeBillDownloadCommonSummaryRequest) (response *DescribeBillDownloadCommonSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillDownloadCommonSummaryRequest()
    }
    response = NewDescribeBillDownloadCommonSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDownloadListRequest() (request *DescribeBillDownloadListRequest) {
    request = &DescribeBillDownloadListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadList")
    return
}

func NewDescribeBillDownloadListResponse() (response *DescribeBillDownloadListResponse) {
    response = &DescribeBillDownloadListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取账单下载地址列表
func (c *Client) DescribeBillDownloadList(request *DescribeBillDownloadListRequest) (response *DescribeBillDownloadListResponse, err error) {
    if request == nil {
        request = NewDescribeBillDownloadListRequest()
    }
    response = NewDescribeBillDownloadListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDownloadRecordListRequest() (request *DescribeBillDownloadRecordListRequest) {
    request = &DescribeBillDownloadRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadRecordList")
    return
}

func NewDescribeBillDownloadRecordListResponse() (response *DescribeBillDownloadRecordListResponse) {
    response = &DescribeBillDownloadRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取下载记录列表
func (c *Client) DescribeBillDownloadRecordList(request *DescribeBillDownloadRecordListRequest) (response *DescribeBillDownloadRecordListResponse, err error) {
    if request == nil {
        request = NewDescribeBillDownloadRecordListRequest()
    }
    response = NewDescribeBillDownloadRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillFeeByPayModeRequest() (request *DescribeBillFeeByPayModeRequest) {
    request = &DescribeBillFeeByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByPayMode")
    return
}

func NewDescribeBillFeeByPayModeResponse() (response *DescribeBillFeeByPayModeResponse) {
    response = &DescribeBillFeeByPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取按计费模式汇总明细费用
func (c *Client) DescribeBillFeeByPayMode(request *DescribeBillFeeByPayModeRequest) (response *DescribeBillFeeByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillFeeByPayModeRequest()
    }
    response = NewDescribeBillFeeByPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillFeeByProductRequest() (request *DescribeBillFeeByProductRequest) {
    request = &DescribeBillFeeByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByProduct")
    return
}

func NewDescribeBillFeeByProductResponse() (response *DescribeBillFeeByProductResponse) {
    response = &DescribeBillFeeByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取按产品汇总明细费用
func (c *Client) DescribeBillFeeByProduct(request *DescribeBillFeeByProductRequest) (response *DescribeBillFeeByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillFeeByProductRequest()
    }
    response = NewDescribeBillFeeByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillFeeByRegionRequest() (request *DescribeBillFeeByRegionRequest) {
    request = &DescribeBillFeeByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByRegion")
    return
}

func NewDescribeBillFeeByRegionResponse() (response *DescribeBillFeeByRegionResponse) {
    response = &DescribeBillFeeByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取按地域汇总明细费用
func (c *Client) DescribeBillFeeByRegion(request *DescribeBillFeeByRegionRequest) (response *DescribeBillFeeByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillFeeByRegionRequest()
    }
    response = NewDescribeBillFeeByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillFeeTrendRequest() (request *DescribeBillFeeTrendRequest) {
    request = &DescribeBillFeeTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeTrend")
    return
}

func NewDescribeBillFeeTrendResponse() (response *DescribeBillFeeTrendResponse) {
    response = &DescribeBillFeeTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取各类汇总明细费用趋势
func (c *Client) DescribeBillFeeTrend(request *DescribeBillFeeTrendRequest) (response *DescribeBillFeeTrendResponse, err error) {
    if request == nil {
        request = NewDescribeBillFeeTrendRequest()
    }
    response = NewDescribeBillFeeTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByPayMode")
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取计费模式汇总费用分布
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    response = NewDescribeBillSummaryByPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByProduct")
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取产品汇总费用分布
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    response = NewDescribeBillSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
    request = &DescribeBillSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByRegion")
    return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
    response = &DescribeBillSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取按地域汇总费用分布
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByRegionRequest()
    }
    response = NewDescribeBillSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByResourceRequest() (request *DescribeBillSummaryByResourceRequest) {
    request = &DescribeBillSummaryByResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByResource")
    return
}

func NewDescribeBillSummaryByResourceResponse() (response *DescribeBillSummaryByResourceResponse) {
    response = &DescribeBillSummaryByResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取按资源汇总数据
func (c *Client) DescribeBillSummaryByResource(request *DescribeBillSummaryByResourceRequest) (response *DescribeBillSummaryByResourceResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByResourceRequest()
    }
    response = NewDescribeBillSummaryByResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillTrendByMonthRequest() (request *DescribeBillTrendByMonthRequest) {
    request = &DescribeBillTrendByMonthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillTrendByMonth")
    return
}

func NewDescribeBillTrendByMonthResponse() (response *DescribeBillTrendByMonthResponse) {
    response = &DescribeBillTrendByMonthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询账单按月消费趋势
func (c *Client) DescribeBillTrendByMonth(request *DescribeBillTrendByMonthRequest) (response *DescribeBillTrendByMonthResponse, err error) {
    if request == nil {
        request = NewDescribeBillTrendByMonthRequest()
    }
    response = NewDescribeBillTrendByMonthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingItemsRequest() (request *DescribeBillingItemsRequest) {
    request = &DescribeBillingItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillingItems")
    return
}

func NewDescribeBillingItemsResponse() (response *DescribeBillingItemsResponse) {
    response = &DescribeBillingItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取计费项信息（第三层）
func (c *Client) DescribeBillingItems(request *DescribeBillingItemsRequest) (response *DescribeBillingItemsResponse, err error) {
    if request == nil {
        request = NewDescribeBillingItemsRequest()
    }
    response = NewDescribeBillingItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsolidatedBillDownloadUrlRequest() (request *DescribeConsolidatedBillDownloadUrlRequest) {
    request = &DescribeConsolidatedBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeConsolidatedBillDownloadUrl")
    return
}

func NewDescribeConsolidatedBillDownloadUrlResponse() (response *DescribeConsolidatedBillDownloadUrlResponse) {
    response = &DescribeConsolidatedBillDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取账单包下载地址
func (c *Client) DescribeConsolidatedBillDownloadUrl(request *DescribeConsolidatedBillDownloadUrlRequest) (response *DescribeConsolidatedBillDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeConsolidatedBillDownloadUrlRequest()
    }
    response = NewDescribeConsolidatedBillDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDealListRequest() (request *DescribeDealListRequest) {
    request = &DescribeDealListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeDealList")
    return
}

func NewDescribeDealListResponse() (response *DescribeDealListResponse) {
    response = &DescribeDealListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理端查询订单列表
func (c *Client) DescribeDealList(request *DescribeDealListRequest) (response *DescribeDealListResponse, err error) {
    if request == nil {
        request = NewDescribeDealListRequest()
    }
    response = NewDescribeDealListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDealPriceRequest() (request *DescribeDealPriceRequest) {
    request = &DescribeDealPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeDealPrice")
    return
}

func NewDescribeDealPriceResponse() (response *DescribeDealPriceResponse) {
    response = &DescribeDealPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询订单价格
func (c *Client) DescribeDealPrice(request *DescribeDealPriceRequest) (response *DescribeDealPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDealPriceRequest()
    }
    response = NewDescribeDealPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDebateListRequest() (request *DescribeDebateListRequest) {
    request = &DescribeDebateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeDebateList")
    return
}

func NewDescribeDebateListResponse() (response *DescribeDebateListResponse) {
    response = &DescribeDebateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理端查询用户账单是否欠费和欠费月份
func (c *Client) DescribeDebateList(request *DescribeDebateListRequest) (response *DescribeDebateListResponse, err error) {
    if request == nil {
        request = NewDescribeDebateListRequest()
    }
    response = NewDescribeDebateListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGoodsPriceRequest() (request *DescribeGoodsPriceRequest) {
    request = &DescribeGoodsPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeGoodsPrice")
    return
}

func NewDescribeGoodsPriceResponse() (response *DescribeGoodsPriceResponse) {
    response = &DescribeGoodsPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询商品价格
func (c *Client) DescribeGoodsPrice(request *DescribeGoodsPriceRequest) (response *DescribeGoodsPriceResponse, err error) {
    if request == nil {
        request = NewDescribeGoodsPriceRequest()
    }
    response = NewDescribeGoodsPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLifecycleFlowListRequest() (request *DescribeLifecycleFlowListRequest) {
    request = &DescribeLifecycleFlowListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeLifecycleFlowList")
    return
}

func NewDescribeLifecycleFlowListResponse() (response *DescribeLifecycleFlowListResponse) {
    response = &DescribeLifecycleFlowListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询资源操作流水
func (c *Client) DescribeLifecycleFlowList(request *DescribeLifecycleFlowListRequest) (response *DescribeLifecycleFlowListResponse, err error) {
    if request == nil {
        request = NewDescribeLifecycleFlowListRequest()
    }
    response = NewDescribeLifecycleFlowListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePayOrderRequest() (request *DescribePayOrderRequest) {
    request = &DescribePayOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribePayOrder")
    return
}

func NewDescribePayOrderResponse() (response *DescribePayOrderResponse) {
    response = &DescribePayOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询订单
func (c *Client) DescribePayOrder(request *DescribePayOrderRequest) (response *DescribePayOrderResponse, err error) {
    if request == nil {
        request = NewDescribePayOrderRequest()
    }
    response = NewDescribePayOrderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductPricesRequest() (request *DescribeProductPricesRequest) {
    request = &DescribeProductPricesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeProductPrices")
    return
}

func NewDescribeProductPricesResponse() (response *DescribeProductPricesResponse) {
    response = &DescribeProductPricesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取价格策略列表
func (c *Client) DescribeProductPrices(request *DescribeProductPricesRequest) (response *DescribeProductPricesResponse, err error) {
    if request == nil {
        request = NewDescribeProductPricesRequest()
    }
    response = NewDescribeProductPricesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductRelationsRequest() (request *DescribeProductRelationsRequest) {
    request = &DescribeProductRelationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeProductRelations")
    return
}

func NewDescribeProductRelationsResponse() (response *DescribeProductRelationsResponse) {
    response = &DescribeProductRelationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 产品定义列表
func (c *Client) DescribeProductRelations(request *DescribeProductRelationsRequest) (response *DescribeProductRelationsResponse, err error) {
    if request == nil {
        request = NewDescribeProductRelationsRequest()
    }
    response = NewDescribeProductRelationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductsRequest() (request *DescribeProductsRequest) {
    request = &DescribeProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeProducts")
    return
}

func NewDescribeProductsResponse() (response *DescribeProductsResponse) {
    response = &DescribeProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取商品信息(第一层)
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    if request == nil {
        request = NewDescribeProductsRequest()
    }
    response = NewDescribeProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaAssignLogsListRequest() (request *DescribeQuotaAssignLogsListRequest) {
    request = &DescribeQuotaAssignLogsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotaAssignLogsList")
    return
}

func NewDescribeQuotaAssignLogsListResponse() (response *DescribeQuotaAssignLogsListResponse) {
    response = &DescribeQuotaAssignLogsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询账户配额转移流水
func (c *Client) DescribeQuotaAssignLogsList(request *DescribeQuotaAssignLogsListRequest) (response *DescribeQuotaAssignLogsListResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaAssignLogsListRequest()
    }
    response = NewDescribeQuotaAssignLogsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaTypeListRequest() (request *DescribeQuotaTypeListRequest) {
    request = &DescribeQuotaTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotaTypeList")
    return
}

func NewDescribeQuotaTypeListResponse() (response *DescribeQuotaTypeListResponse) {
    response = &DescribeQuotaTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配额类别：uin、sys
func (c *Client) DescribeQuotaTypeList(request *DescribeQuotaTypeListRequest) (response *DescribeQuotaTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaTypeListRequest()
    }
    response = NewDescribeQuotaTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotasRequest() (request *DescribeQuotasRequest) {
    request = &DescribeQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotas")
    return
}

func NewDescribeQuotasResponse() (response *DescribeQuotasResponse) {
    response = &DescribeQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 产品配额列表查询
func (c *Client) DescribeQuotas(request *DescribeQuotasRequest) (response *DescribeQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeQuotasRequest()
    }
    response = NewDescribeQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRefundRequest() (request *DescribeRefundRequest) {
    request = &DescribeRefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeRefund")
    return
}

func NewDescribeRefundResponse() (response *DescribeRefundResponse) {
    response = &DescribeRefundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询提现(退款)信息
func (c *Client) DescribeRefund(request *DescribeRefundRequest) (response *DescribeRefundResponse, err error) {
    if request == nil {
        request = NewDescribeRefundRequest()
    }
    response = NewDescribeRefundResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceListRequest() (request *DescribeResourceListRequest) {
    request = &DescribeResourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceList")
    return
}

func NewDescribeResourceListResponse() (response *DescribeResourceListResponse) {
    response = &DescribeResourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理端查询用户产品列表，只支持到一层产品
func (c *Client) DescribeResourceList(request *DescribeResourceListRequest) (response *DescribeResourceListResponse, err error) {
    if request == nil {
        request = NewDescribeResourceListRequest()
    }
    response = NewDescribeResourceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubBillingItemsRequest() (request *DescribeSubBillingItemsRequest) {
    request = &DescribeSubBillingItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeSubBillingItems")
    return
}

func NewDescribeSubBillingItemsResponse() (response *DescribeSubBillingItemsResponse) {
    response = &DescribeSubBillingItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取计费细项信息（第四层）
func (c *Client) DescribeSubBillingItems(request *DescribeSubBillingItemsRequest) (response *DescribeSubBillingItemsResponse, err error) {
    if request == nil {
        request = NewDescribeSubBillingItemsRequest()
    }
    response = NewDescribeSubBillingItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubProductsRequest() (request *DescribeSubProductsRequest) {
    request = &DescribeSubProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeSubProducts")
    return
}

func NewDescribeSubProductsResponse() (response *DescribeSubProductsResponse) {
    response = &DescribeSubProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取子商品信息
func (c *Client) DescribeSubProducts(request *DescribeSubProductsRequest) (response *DescribeSubProductsResponse, err error) {
    if request == nil {
        request = NewDescribeSubProductsRequest()
    }
    response = NewDescribeSubProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserAccountRequest() (request *DescribeUserAccountRequest) {
    request = &DescribeUserAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeUserAccount")
    return
}

func NewDescribeUserAccountResponse() (response *DescribeUserAccountResponse) {
    response = &DescribeUserAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询单uin后付费账户信息，SubAcctid=CREDIT_FIXED 时才返回新账期NextDueDelay等详情
func (c *Client) DescribeUserAccount(request *DescribeUserAccountRequest) (response *DescribeUserAccountResponse, err error) {
    if request == nil {
        request = NewDescribeUserAccountRequest()
    }
    response = NewDescribeUserAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDebtBillRequest() (request *DescribeUserDebtBillRequest) {
    request = &DescribeUserDebtBillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "DescribeUserDebtBill")
    return
}

func NewDescribeUserDebtBillResponse() (response *DescribeUserDebtBillResponse) {
    response = &DescribeUserDebtBillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询后付费账单
func (c *Client) DescribeUserDebtBill(request *DescribeUserDebtBillRequest) (response *DescribeUserDebtBillResponse, err error) {
    if request == nil {
        request = NewDescribeUserDebtBillRequest()
    }
    response = NewDescribeUserDebtBillResponse()
    err = c.Send(request, response)
    return
}

func NewGetBillingProductCodeRequest() (request *GetBillingProductCodeRequest) {
    request = &GetBillingProductCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "GetBillingProductCode")
    return
}

func NewGetBillingProductCodeResponse() (response *GetBillingProductCodeResponse) {
    response = &GetBillingProductCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询由计量迁移至计费的产品码
func (c *Client) GetBillingProductCode(request *GetBillingProductCodeRequest) (response *GetBillingProductCodeResponse, err error) {
    if request == nil {
        request = NewGetBillingProductCodeRequest()
    }
    response = NewGetBillingProductCodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetProductTreeRequest() (request *GetProductTreeRequest) {
    request = &GetProductTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "GetProductTree")
    return
}

func NewGetProductTreeResponse() (response *GetProductTreeResponse) {
    response = &GetProductTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过商品码，子商品码，计费项，计费细项来查询产品定义，参数都可选，如果不传则返回全部数据
func (c *Client) GetProductTree(request *GetProductTreeRequest) (response *GetProductTreeResponse, err error) {
    if request == nil {
        request = NewGetProductTreeRequest()
    }
    response = NewGetProductTreeResponse()
    err = c.Send(request, response)
    return
}

func NewGetQuotaRequest() (request *GetQuotaRequest) {
    request = &GetQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "GetQuota")
    return
}

func NewGetQuotaResponse() (response *GetQuotaResponse) {
    response = &GetQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据商品码和配额key查询配额
func (c *Client) GetQuota(request *GetQuotaRequest) (response *GetQuotaResponse, err error) {
    if request == nil {
        request = NewGetQuotaRequest()
    }
    response = NewGetQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewGetQuotaLeftRequest() (request *GetQuotaLeftRequest) {
    request = &GetQuotaLeftRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "GetQuotaLeft")
    return
}

func NewGetQuotaLeftResponse() (response *GetQuotaLeftResponse) {
    response = &GetQuotaLeftResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取我的剩余配额
func (c *Client) GetQuotaLeft(request *GetQuotaLeftRequest) (response *GetQuotaLeftResponse, err error) {
    if request == nil {
        request = NewGetQuotaLeftRequest()
    }
    response = NewGetQuotaLeftResponse()
    err = c.Send(request, response)
    return
}

func NewGetResourceBillRequest() (request *GetResourceBillRequest) {
    request = &GetResourceBillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "GetResourceBill")
    return
}

func NewGetResourceBillResponse() (response *GetResourceBillResponse) {
    response = &GetResourceBillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户的账单信息
func (c *Client) GetResourceBill(request *GetResourceBillRequest) (response *GetResourceBillResponse, err error) {
    if request == nil {
        request = NewGetResourceBillRequest()
    }
    response = NewGetResourceBillResponse()
    err = c.Send(request, response)
    return
}

func NewListResourceRequest() (request *ListResourceRequest) {
    request = &ListResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "ListResource")
    return
}

func NewListResourceResponse() (response *ListResourceResponse) {
    response = &ListResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运营端根据Uin，ResourceId查询用户资源 
func (c *Client) ListResource(request *ListResourceRequest) (response *ListResourceResponse, err error) {
    if request == nil {
        request = NewListResourceRequest()
    }
    response = NewListResourceResponse()
    err = c.Send(request, response)
    return
}

func NewSaveQuotaRequest() (request *SaveQuotaRequest) {
    request = &SaveQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "SaveQuota")
    return
}

func NewSaveQuotaResponse() (response *SaveQuotaResponse) {
    response = &SaveQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 产品配额新增、编辑
func (c *Client) SaveQuota(request *SaveQuotaRequest) (response *SaveQuotaResponse, err error) {
    if request == nil {
        request = NewSaveQuotaRequest()
    }
    response = NewSaveQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewSetProductInfoRequest() (request *SetProductInfoRequest) {
    request = &SetProductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbill", APIVersion, "SetProductInfo")
    return
}

func NewSetProductInfoResponse() (response *SetProductInfoResponse) {
    response = &SetProductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 产品四层定义
func (c *Client) SetProductInfo(request *SetProductInfoRequest) (response *SetProductInfoResponse, err error) {
    if request == nil {
        request = NewSetProductInfoRequest()
    }
    response = NewSetProductInfoResponse()
    err = c.Send(request, response)
    return
}
