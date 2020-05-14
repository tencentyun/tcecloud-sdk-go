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


func NewAddPartaMailingAddressRequest() (request *AddPartaMailingAddressRequest) {
    request = &AddPartaMailingAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "AddPartaMailingAddress")
    return
}

func NewAddPartaMailingAddressResponse() (response *AddPartaMailingAddressResponse) {
    response = &AddPartaMailingAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加邮寄地址
func (c *Client) AddPartaMailingAddress(request *AddPartaMailingAddressRequest) (response *AddPartaMailingAddressResponse, err error) {
    if request == nil {
        request = NewAddPartaMailingAddressRequest()
    }
    response = NewAddPartaMailingAddressResponse()
    err = c.Send(request, response)
    return
}

func NewAddQuotaRequest() (request *AddQuotaRequest) {
    request = &AddQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "AddQuota")
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

func NewCancelInvoiceRequest() (request *CancelInvoiceRequest) {
    request = &CancelInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CancelInvoice")
    return
}

func NewCancelInvoiceResponse() (response *CancelInvoiceResponse) {
    response = &CancelInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 取消开票申请
func (c *Client) CancelInvoice(request *CancelInvoiceRequest) (response *CancelInvoiceResponse, err error) {
    if request == nil {
        request = NewCancelInvoiceRequest()
    }
    response = NewCancelInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewCouponInfoRequest() (request *CouponInfoRequest) {
    request = &CouponInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CouponInfo")
    return
}

func NewCouponInfoResponse() (response *CouponInfoResponse) {
    response = &CouponInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询代金券信息
func (c *Client) CouponInfo(request *CouponInfoRequest) (response *CouponInfoResponse, err error) {
    if request == nil {
        request = NewCouponInfoRequest()
    }
    response = NewCouponInfoResponse()
    err = c.Send(request, response)
    return
}

func NewCouponsWaterRequest() (request *CouponsWaterRequest) {
    request = &CouponsWaterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CouponsWater")
    return
}

func NewCouponsWaterResponse() (response *CouponsWaterResponse) {
    response = &CouponsWaterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 代金券流水
func (c *Client) CouponsWater(request *CouponsWaterRequest) (response *CouponsWaterResponse, err error) {
    if request == nil {
        request = NewCouponsWaterRequest()
    }
    response = NewCouponsWaterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContractRequest() (request *CreateContractRequest) {
    request = &CreateContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CreateContract")
    return
}

func NewCreateContractResponse() (response *CreateContractResponse) {
    response = &CreateContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建合同
func (c *Client) CreateContract(request *CreateContractRequest) (response *CreateContractResponse, err error) {
    if request == nil {
        request = NewCreateContractRequest()
    }
    response = NewCreateContractResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDownloadRecordRequest() (request *CreateDownloadRecordRequest) {
    request = &CreateDownloadRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CreateDownloadRecord")
    return
}

func NewCreateDownloadRecordResponse() (response *CreateDownloadRecordResponse) {
    response = &CreateDownloadRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户点击下载时通知后台
func (c *Client) CreateDownloadRecord(request *CreateDownloadRecordRequest) (response *CreateDownloadRecordResponse, err error) {
    if request == nil {
        request = NewCreateDownloadRecordRequest()
    }
    response = NewCreateDownloadRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInvoiceRequest() (request *CreateInvoiceRequest) {
    request = &CreateInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CreateInvoice")
    return
}

func NewCreateInvoiceResponse() (response *CreateInvoiceResponse) {
    response = &CreateInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户申请开票
func (c *Client) CreateInvoice(request *CreateInvoiceRequest) (response *CreateInvoiceResponse, err error) {
    if request == nil {
        request = NewCreateInvoiceRequest()
    }
    response = NewCreateInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePayOrderRequest() (request *CreatePayOrderRequest) {
    request = &CreatePayOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CreatePayOrder")
    return
}

func NewCreatePayOrderResponse() (response *CreatePayOrderResponse) {
    response = &CreatePayOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建现金支付订单
func (c *Client) CreatePayOrder(request *CreatePayOrderRequest) (response *CreatePayOrderResponse, err error) {
    if request == nil {
        request = NewCreatePayOrderRequest()
    }
    response = NewCreatePayOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRemitRecordRequest() (request *CreateRemitRecordRequest) {
    request = &CreateRemitRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CreateRemitRecord")
    return
}

func NewCreateRemitRecordResponse() (response *CreateRemitRecordResponse) {
    response = &CreateRemitRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户新增汇款记录接口
func (c *Client) CreateRemitRecord(request *CreateRemitRecordRequest) (response *CreateRemitRecordResponse, err error) {
    if request == nil {
        request = NewCreateRemitRecordRequest()
    }
    response = NewCreateRemitRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWarningRequest() (request *CreateWarningRequest) {
    request = &CreateWarningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "CreateWarning")
    return
}

func NewCreateWarningResponse() (response *CreateWarningResponse) {
    response = &CreateWarningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加或更新预警信息
func (c *Client) CreateWarning(request *CreateWarningRequest) (response *CreateWarningResponse, err error) {
    if request == nil {
        request = NewCreateWarningRequest()
    }
    response = NewCreateWarningResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePartaMailingAddressRequest() (request *DeletePartaMailingAddressRequest) {
    request = &DeletePartaMailingAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DeletePartaMailingAddress")
    return
}

func NewDeletePartaMailingAddressResponse() (response *DeletePartaMailingAddressResponse) {
    response = &DeletePartaMailingAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除邮寄地址
func (c *Client) DeletePartaMailingAddress(request *DeletePartaMailingAddressRequest) (response *DeletePartaMailingAddressResponse, err error) {
    if request == nil {
        request = NewDeletePartaMailingAddressRequest()
    }
    response = NewDeletePartaMailingAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePostInfoRequest() (request *DeletePostInfoRequest) {
    request = &DeletePostInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DeletePostInfo")
    return
}

func NewDeletePostInfoResponse() (response *DeletePostInfoResponse) {
    response = &DeletePostInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户的邮寄信息
func (c *Client) DeletePostInfo(request *DeletePostInfoRequest) (response *DeletePostInfoResponse, err error) {
    if request == nil {
        request = NewDeletePostInfoRequest()
    }
    response = NewDeletePostInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQuotaRequest() (request *DeleteQuotaRequest) {
    request = &DeleteQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DeleteQuota")
    return
}

func NewDeleteQuotaResponse() (response *DeleteQuotaResponse) {
    response = &DeleteQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 接口描述
func (c *Client) DeleteQuota(request *DeleteQuotaRequest) (response *DeleteQuotaResponse, err error) {
    if request == nil {
        request = NewDeleteQuotaRequest()
    }
    response = NewDeleteQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubUinQuotaRequest() (request *DeleteSubUinQuotaRequest) {
    request = &DeleteSubUinQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DeleteSubUinQuota")
    return
}

func NewDeleteSubUinQuotaResponse() (response *DeleteSubUinQuotaResponse) {
    response = &DeleteSubUinQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除父账户下子账户的配额数据
func (c *Client) DeleteSubUinQuota(request *DeleteSubUinQuotaRequest) (response *DeleteSubUinQuotaResponse, err error) {
    if request == nil {
        request = NewDeleteSubUinQuotaRequest()
    }
    response = NewDeleteSubUinQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescriPartaMailingAddressListRequest() (request *DescriPartaMailingAddressListRequest) {
    request = &DescriPartaMailingAddressListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescriPartaMailingAddressList")
    return
}

func NewDescriPartaMailingAddressListResponse() (response *DescriPartaMailingAddressListResponse) {
    response = &DescriPartaMailingAddressListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询邮寄地址列表
func (c *Client) DescriPartaMailingAddressList(request *DescriPartaMailingAddressListRequest) (response *DescriPartaMailingAddressListResponse, err error) {
    if request == nil {
        request = NewDescriPartaMailingAddressListRequest()
    }
    response = NewDescriPartaMailingAddressListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPeriodRequest() (request *DescribeAccountPeriodRequest) {
    request = &DescribeAccountPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeAccountPeriod")
    return
}

func NewDescribeAccountPeriodResponse() (response *DescribeAccountPeriodResponse) {
    response = &DescribeAccountPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取账期信息
func (c *Client) DescribeAccountPeriod(request *DescribeAccountPeriodRequest) (response *DescribeAccountPeriodResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPeriodRequest()
    }
    response = NewDescribeAccountPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountWaterRequest() (request *DescribeAccountWaterRequest) {
    request = &DescribeAccountWaterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeAccountWater")
    return
}

func NewDescribeAccountWaterResponse() (response *DescribeAccountWaterResponse) {
    response = &DescribeAccountWaterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询账户的还款或销账流水
func (c *Client) DescribeAccountWater(request *DescribeAccountWaterRequest) (response *DescribeAccountWaterResponse, err error) {
    if request == nil {
        request = NewDescribeAccountWaterRequest()
    }
    response = NewDescribeAccountWaterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailByResourceRequest() (request *DescribeBillDetailByResourceRequest) {
    request = &DescribeBillDetailByResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDetailByResource")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadByResource")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadCommonSummary")
    return
}

func NewDescribeBillDownloadCommonSummaryResponse() (response *DescribeBillDownloadCommonSummaryResponse) {
    response = &DescribeBillDownloadCommonSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载账单通用汇总（产品+项目+地域）
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadList")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadRecordList")
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

func NewDescribeBillDownloadResourceDetailRequest() (request *DescribeBillDownloadResourceDetailRequest) {
    request = &DescribeBillDownloadResourceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadResourceDetail")
    return
}

func NewDescribeBillDownloadResourceDetailResponse() (response *DescribeBillDownloadResourceDetailResponse) {
    response = &DescribeBillDownloadResourceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载定制化账单明细
func (c *Client) DescribeBillDownloadResourceDetail(request *DescribeBillDownloadResourceDetailRequest) (response *DescribeBillDownloadResourceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillDownloadResourceDetailRequest()
    }
    response = NewDescribeBillDownloadResourceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillFeeByPayModeRequest() (request *DescribeBillFeeByPayModeRequest) {
    request = &DescribeBillFeeByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByPayMode")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByProduct")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByRegion")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeTrend")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByPayMode")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByProduct")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByRegion")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByResource")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeBillTrendByMonth")
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

func NewDescribeConsolidatedBillDownloadUrlRequest() (request *DescribeConsolidatedBillDownloadUrlRequest) {
    request = &DescribeConsolidatedBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeConsolidatedBillDownloadUrl")
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

func NewDescribeContractListRequest() (request *DescribeContractListRequest) {
    request = &DescribeContractListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeContractList")
    return
}

func NewDescribeContractListResponse() (response *DescribeContractListResponse) {
    response = &DescribeContractListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询合同列表接口
func (c *Client) DescribeContractList(request *DescribeContractListRequest) (response *DescribeContractListResponse, err error) {
    if request == nil {
        request = NewDescribeContractListRequest()
    }
    response = NewDescribeContractListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContractPartaInfoRequest() (request *DescribeContractPartaInfoRequest) {
    request = &DescribeContractPartaInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeContractPartaInfo")
    return
}

func NewDescribeContractPartaInfoResponse() (response *DescribeContractPartaInfoResponse) {
    response = &DescribeContractPartaInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询合同甲方信息
func (c *Client) DescribeContractPartaInfo(request *DescribeContractPartaInfoRequest) (response *DescribeContractPartaInfoResponse, err error) {
    if request == nil {
        request = NewDescribeContractPartaInfoRequest()
    }
    response = NewDescribeContractPartaInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContractTypeListRequest() (request *DescribeContractTypeListRequest) {
    request = &DescribeContractTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeContractTypeList")
    return
}

func NewDescribeContractTypeListResponse() (response *DescribeContractTypeListResponse) {
    response = &DescribeContractTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取合同类型接口
func (c *Client) DescribeContractTypeList(request *DescribeContractTypeListRequest) (response *DescribeContractTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeContractTypeListRequest()
    }
    response = NewDescribeContractTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCouponsDeductInfoRequest() (request *DescribeCouponsDeductInfoRequest) {
    request = &DescribeCouponsDeductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeCouponsDeductInfo")
    return
}

func NewDescribeCouponsDeductInfoResponse() (response *DescribeCouponsDeductInfoResponse) {
    response = &DescribeCouponsDeductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询抵扣券作用在物品时，抵扣金额信息
func (c *Client) DescribeCouponsDeductInfo(request *DescribeCouponsDeductInfoRequest) (response *DescribeCouponsDeductInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCouponsDeductInfoRequest()
    }
    response = NewDescribeCouponsDeductInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDealListRequest() (request *DescribeDealListRequest) {
    request = &DescribeDealListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeDealList")
    return
}

func NewDescribeDealListResponse() (response *DescribeDealListResponse) {
    response = &DescribeDealListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询订单列表
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeDealPrice")
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

func NewDescribeDownloadVoucherListFileRequest() (request *DescribeDownloadVoucherListFileRequest) {
    request = &DescribeDownloadVoucherListFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadVoucherListFile")
    return
}

func NewDescribeDownloadVoucherListFileResponse() (response *DescribeDownloadVoucherListFileResponse) {
    response = &DescribeDownloadVoucherListFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载代金券信息excel
func (c *Client) DescribeDownloadVoucherListFile(request *DescribeDownloadVoucherListFileRequest) (response *DescribeDownloadVoucherListFileResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadVoucherListFileRequest()
    }
    response = NewDescribeDownloadVoucherListFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDownloadWithAuthRequest() (request *DescribeDownloadWithAuthRequest) {
    request = &DescribeDownloadWithAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadWithAuth")
    return
}

func NewDescribeDownloadWithAuthResponse() (response *DescribeDownloadWithAuthResponse) {
    response = &DescribeDownloadWithAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeDownloadWithAuth(request *DescribeDownloadWithAuthRequest) (response *DescribeDownloadWithAuthResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadWithAuthRequest()
    }
    response = NewDescribeDownloadWithAuthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGoodsPriceRequest() (request *DescribeGoodsPriceRequest) {
    request = &DescribeGoodsPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeGoodsPrice")
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

func NewDescribeInvoiceAmountRequest() (request *DescribeInvoiceAmountRequest) {
    request = &DescribeInvoiceAmountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceAmount")
    return
}

func NewDescribeInvoiceAmountResponse() (response *DescribeInvoiceAmountResponse) {
    response = &DescribeInvoiceAmountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取可开票金额
func (c *Client) DescribeInvoiceAmount(request *DescribeInvoiceAmountRequest) (response *DescribeInvoiceAmountResponse, err error) {
    if request == nil {
        request = NewDescribeInvoiceAmountRequest()
    }
    response = NewDescribeInvoiceAmountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvoiceDetailRequest() (request *DescribeInvoiceDetailRequest) {
    request = &DescribeInvoiceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceDetail")
    return
}

func NewDescribeInvoiceDetailResponse() (response *DescribeInvoiceDetailResponse) {
    response = &DescribeInvoiceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理端、租户端-获取发票的开票详情
func (c *Client) DescribeInvoiceDetail(request *DescribeInvoiceDetailRequest) (response *DescribeInvoiceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInvoiceDetailRequest()
    }
    response = NewDescribeInvoiceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvoiceListRequest() (request *DescribeInvoiceListRequest) {
    request = &DescribeInvoiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceList")
    return
}

func NewDescribeInvoiceListResponse() (response *DescribeInvoiceListResponse) {
    response = &DescribeInvoiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理端、租户端-获取用户开票记录
func (c *Client) DescribeInvoiceList(request *DescribeInvoiceListRequest) (response *DescribeInvoiceListResponse, err error) {
    if request == nil {
        request = NewDescribeInvoiceListRequest()
    }
    response = NewDescribeInvoiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenInvoiceInfoRequest() (request *DescribeOpenInvoiceInfoRequest) {
    request = &DescribeOpenInvoiceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeOpenInvoiceInfo")
    return
}

func NewDescribeOpenInvoiceInfoResponse() (response *DescribeOpenInvoiceInfoResponse) {
    response = &DescribeOpenInvoiceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户开票信息
func (c *Client) DescribeOpenInvoiceInfo(request *DescribeOpenInvoiceInfoRequest) (response *DescribeOpenInvoiceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeOpenInvoiceInfoRequest()
    }
    response = NewDescribeOpenInvoiceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePayOrderRequest() (request *DescribePayOrderRequest) {
    request = &DescribePayOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribePayOrder")
    return
}

func NewDescribePayOrderResponse() (response *DescribePayOrderResponse) {
    response = &DescribePayOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询支付订单。
func (c *Client) DescribePayOrder(request *DescribePayOrderRequest) (response *DescribePayOrderResponse, err error) {
    if request == nil {
        request = NewDescribePayOrderRequest()
    }
    response = NewDescribePayOrderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePostInfoListRequest() (request *DescribePostInfoListRequest) {
    request = &DescribePostInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribePostInfoList")
    return
}

func NewDescribePostInfoListResponse() (response *DescribePostInfoListResponse) {
    response = &DescribePostInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户的邮寄信息
func (c *Client) DescribePostInfoList(request *DescribePostInfoListRequest) (response *DescribePostInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePostInfoListRequest()
    }
    response = NewDescribePostInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductRelationsRequest() (request *DescribeProductRelationsRequest) {
    request = &DescribeProductRelationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeProductRelations")
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeProducts")
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

func NewDescribePropertiesRequest() (request *DescribePropertiesRequest) {
    request = &DescribePropertiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeProperties")
    return
}

func NewDescribePropertiesResponse() (response *DescribePropertiesResponse) {
    response = &DescribePropertiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询售卖属性列表
func (c *Client) DescribeProperties(request *DescribePropertiesRequest) (response *DescribePropertiesResponse, err error) {
    if request == nil {
        request = NewDescribePropertiesRequest()
    }
    response = NewDescribePropertiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaRequest() (request *DescribeQuotaRequest) {
    request = &DescribeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeQuota")
    return
}

func NewDescribeQuotaResponse() (response *DescribeQuotaResponse) {
    response = &DescribeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户当前使用配额
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaRequest()
    }
    response = NewDescribeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReconciliationListRequest() (request *DescribeReconciliationListRequest) {
    request = &DescribeReconciliationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeReconciliationList")
    return
}

func NewDescribeReconciliationListResponse() (response *DescribeReconciliationListResponse) {
    response = &DescribeReconciliationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 调账记录查询调账列表
func (c *Client) DescribeReconciliationList(request *DescribeReconciliationListRequest) (response *DescribeReconciliationListResponse, err error) {
    if request == nil {
        request = NewDescribeReconciliationListRequest()
    }
    response = NewDescribeReconciliationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRefundRequest() (request *DescribeRefundRequest) {
    request = &DescribeRefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeRefund")
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

func NewDescribeRegionRequest() (request *DescribeRegionRequest) {
    request = &DescribeRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeRegion")
    return
}

func NewDescribeRegionResponse() (response *DescribeRegionResponse) {
    response = &DescribeRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询地域.  用户id通过公共参数UIN传入
func (c *Client) DescribeRegion(request *DescribeRegionRequest) (response *DescribeRegionResponse, err error) {
    if request == nil {
        request = NewDescribeRegionRequest()
    }
    response = NewDescribeRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRemitRecordRequest() (request *DescribeRemitRecordRequest) {
    request = &DescribeRemitRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeRemitRecord")
    return
}

func NewDescribeRemitRecordResponse() (response *DescribeRemitRecordResponse) {
    response = &DescribeRemitRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户查询汇款记录
func (c *Client) DescribeRemitRecord(request *DescribeRemitRecordRequest) (response *DescribeRemitRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRemitRecordRequest()
    }
    response = NewDescribeRemitRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRenewInfoRequest() (request *DescribeRenewInfoRequest) {
    request = &DescribeRenewInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeRenewInfo")
    return
}

func NewDescribeRenewInfoResponse() (response *DescribeRenewInfoResponse) {
    response = &DescribeRenewInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询续费信息。用户id通过云api公共字段转入
func (c *Client) DescribeRenewInfo(request *DescribeRenewInfoRequest) (response *DescribeRenewInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRenewInfoRequest()
    }
    response = NewDescribeRenewInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceBillDetailRequest() (request *DescribeResourceBillDetailRequest) {
    request = &DescribeResourceBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceBillDetail")
    return
}

func NewDescribeResourceBillDetailResponse() (response *DescribeResourceBillDetailResponse) {
    response = &DescribeResourceBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取组件级别明细账单
func (c *Client) DescribeResourceBillDetail(request *DescribeResourceBillDetailRequest) (response *DescribeResourceBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeResourceBillDetailRequest()
    }
    response = NewDescribeResourceBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceDetailListRequest() (request *DescribeResourceDetailListRequest) {
    request = &DescribeResourceDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceDetailList")
    return
}

func NewDescribeResourceDetailListResponse() (response *DescribeResourceDetailListResponse) {
    response = &DescribeResourceDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户续费参数
func (c *Client) DescribeResourceDetailList(request *DescribeResourceDetailListRequest) (response *DescribeResourceDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeResourceDetailListRequest()
    }
    response = NewDescribeResourceDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceListRequest() (request *DescribeResourceListRequest) {
    request = &DescribeResourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceList")
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

func NewDescribeSubUinQuotasRequest() (request *DescribeSubUinQuotasRequest) {
    request = &DescribeSubUinQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeSubUinQuotas")
    return
}

func NewDescribeSubUinQuotasResponse() (response *DescribeSubUinQuotasResponse) {
    response = &DescribeSubUinQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取子账号配额
func (c *Client) DescribeSubUinQuotas(request *DescribeSubUinQuotasRequest) (response *DescribeSubUinQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeSubUinQuotasRequest()
    }
    response = NewDescribeSubUinQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTenantQuotasRequest() (request *DescribeTenantQuotasRequest) {
    request = &DescribeTenantQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeTenantQuotas")
    return
}

func NewDescribeTenantQuotasResponse() (response *DescribeTenantQuotasResponse) {
    response = &DescribeTenantQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 租户端产品配额列表查询
func (c *Client) DescribeTenantQuotas(request *DescribeTenantQuotasRequest) (response *DescribeTenantQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeTenantQuotasRequest()
    }
    response = NewDescribeTenantQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserAccountRequest() (request *DescribeUserAccountRequest) {
    request = &DescribeUserAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeUserAccount")
    return
}

func NewDescribeUserAccountResponse() (response *DescribeUserAccountResponse) {
    response = &DescribeUserAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询单uin后付费账户信息，SubAcctid=CREDIT_FIXED 时才返回新账期NextDueDelay
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
    request.Init().WithApiInfo("bill", APIVersion, "DescribeUserDebtBill")
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

func NewDescribeWarningRequest() (request *DescribeWarningRequest) {
    request = &DescribeWarningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "DescribeWarning")
    return
}

func NewDescribeWarningResponse() (response *DescribeWarningResponse) {
    response = &DescribeWarningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 租户查询预警状态和阈值
func (c *Client) DescribeWarning(request *DescribeWarningRequest) (response *DescribeWarningResponse, err error) {
    if request == nil {
        request = NewDescribeWarningRequest()
    }
    response = NewDescribeWarningResponse()
    err = c.Send(request, response)
    return
}

func NewExportDealListRequest() (request *ExportDealListRequest) {
    request = &ExportDealListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ExportDealList")
    return
}

func NewExportDealListResponse() (response *ExportDealListResponse) {
    response = &ExportDealListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导入订单列表信息
func (c *Client) ExportDealList(request *ExportDealListRequest) (response *ExportDealListResponse, err error) {
    if request == nil {
        request = NewExportDealListRequest()
    }
    response = NewExportDealListResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateDealsRequest() (request *GenerateDealsRequest) {
    request = &GenerateDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GenerateDeals")
    return
}

func NewGenerateDealsResponse() (response *GenerateDealsResponse) {
    response = &GenerateDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建订单 
func (c *Client) GenerateDeals(request *GenerateDealsRequest) (response *GenerateDealsResponse, err error) {
    if request == nil {
        request = NewGenerateDealsRequest()
    }
    response = NewGenerateDealsResponse()
    err = c.Send(request, response)
    return
}

func NewGetBillingProductCodeRequest() (request *GetBillingProductCodeRequest) {
    request = &GetBillingProductCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetBillingProductCode")
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

func NewGetMineQuotaRequest() (request *GetMineQuotaRequest) {
    request = &GetMineQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetMineQuota")
    return
}

func NewGetMineQuotaResponse() (response *GetMineQuotaResponse) {
    response = &GetMineQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据商品码和配额key查询用户自身的配额
func (c *Client) GetMineQuota(request *GetMineQuotaRequest) (response *GetMineQuotaResponse, err error) {
    if request == nil {
        request = NewGetMineQuotaRequest()
    }
    response = NewGetMineQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewGetMineResourceBillRequest() (request *GetMineResourceBillRequest) {
    request = &GetMineResourceBillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetMineResourceBill")
    return
}

func NewGetMineResourceBillResponse() (response *GetMineResourceBillResponse) {
    response = &GetMineResourceBillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户自身的账单信息
func (c *Client) GetMineResourceBill(request *GetMineResourceBillRequest) (response *GetMineResourceBillResponse, err error) {
    if request == nil {
        request = NewGetMineResourceBillRequest()
    }
    response = NewGetMineResourceBillResponse()
    err = c.Send(request, response)
    return
}

func NewGetProductTreeRequest() (request *GetProductTreeRequest) {
    request = &GetProductTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "GetProductTree")
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
    request.Init().WithApiInfo("bill", APIVersion, "GetQuota")
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
    request.Init().WithApiInfo("bill", APIVersion, "GetQuotaLeft")
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
    request.Init().WithApiInfo("bill", APIVersion, "GetResourceBill")
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

func NewListMineResourceRequest() (request *ListMineResourceRequest) {
    request = &ListMineResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ListMineResource")
    return
}

func NewListMineResourceResponse() (response *ListMineResourceResponse) {
    response = &ListMineResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 租户段用户查询自己拥有的资源
func (c *Client) ListMineResource(request *ListMineResourceRequest) (response *ListMineResourceResponse, err error) {
    if request == nil {
        request = NewListMineResourceRequest()
    }
    response = NewListMineResourceResponse()
    err = c.Send(request, response)
    return
}

func NewListResourceRequest() (request *ListResourceRequest) {
    request = &ListResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ListResource")
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

func NewModifyAutoFlagRequest() (request *ModifyAutoFlagRequest) {
    request = &ModifyAutoFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyAutoFlag")
    return
}

func NewModifyAutoFlagResponse() (response *ModifyAutoFlagResponse) {
    response = &ModifyAutoFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置续费类型。用户id通过云api公共参数传入
func (c *Client) ModifyAutoFlag(request *ModifyAutoFlagRequest) (response *ModifyAutoFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoFlagRequest()
    }
    response = NewModifyAutoFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContractStatusRequest() (request *ModifyContractStatusRequest) {
    request = &ModifyContractStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyContractStatus")
    return
}

func NewModifyContractStatusResponse() (response *ModifyContractStatusResponse) {
    response = &ModifyContractStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改合同状态
func (c *Client) ModifyContractStatus(request *ModifyContractStatusRequest) (response *ModifyContractStatusResponse, err error) {
    if request == nil {
        request = NewModifyContractStatusRequest()
    }
    response = NewModifyContractStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDealStatusRequest() (request *ModifyDealStatusRequest) {
    request = &ModifyDealStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyDealStatus")
    return
}

func NewModifyDealStatusResponse() (response *ModifyDealStatusResponse) {
    response = &ModifyDealStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据大订单号修改大订单状态
// 
// 当订单状态为待支付时调用此接口可将订单状态改为取消
// 
// 当订单状态为取消时调用此接口可将订单状态改为删除
func (c *Client) ModifyDealStatus(request *ModifyDealStatusRequest) (response *ModifyDealStatusResponse, err error) {
    if request == nil {
        request = NewModifyDealStatusRequest()
    }
    response = NewModifyDealStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDefaultMailingAddressRequest() (request *ModifyDefaultMailingAddressRequest) {
    request = &ModifyDefaultMailingAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyDefaultMailingAddress")
    return
}

func NewModifyDefaultMailingAddressResponse() (response *ModifyDefaultMailingAddressResponse) {
    response = &ModifyDefaultMailingAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改默认邮寄地址
func (c *Client) ModifyDefaultMailingAddress(request *ModifyDefaultMailingAddressRequest) (response *ModifyDefaultMailingAddressResponse, err error) {
    if request == nil {
        request = NewModifyDefaultMailingAddressRequest()
    }
    response = NewModifyDefaultMailingAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDefaultPostInfoRequest() (request *ModifyDefaultPostInfoRequest) {
    request = &ModifyDefaultPostInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyDefaultPostInfo")
    return
}

func NewModifyDefaultPostInfoResponse() (response *ModifyDefaultPostInfoResponse) {
    response = &ModifyDefaultPostInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置默认用户的邮寄信息
func (c *Client) ModifyDefaultPostInfo(request *ModifyDefaultPostInfoRequest) (response *ModifyDefaultPostInfoResponse, err error) {
    if request == nil {
        request = NewModifyDefaultPostInfoRequest()
    }
    response = NewModifyDefaultPostInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOpenInvoiceInfoRequest() (request *ModifyOpenInvoiceInfoRequest) {
    request = &ModifyOpenInvoiceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyOpenInvoiceInfo")
    return
}

func NewModifyOpenInvoiceInfoResponse() (response *ModifyOpenInvoiceInfoResponse) {
    response = &ModifyOpenInvoiceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置开票信息
func (c *Client) ModifyOpenInvoiceInfo(request *ModifyOpenInvoiceInfoRequest) (response *ModifyOpenInvoiceInfoResponse, err error) {
    if request == nil {
        request = NewModifyOpenInvoiceInfoRequest()
    }
    response = NewModifyOpenInvoiceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPartaMailingAddressRequest() (request *ModifyPartaMailingAddressRequest) {
    request = &ModifyPartaMailingAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyPartaMailingAddress")
    return
}

func NewModifyPartaMailingAddressResponse() (response *ModifyPartaMailingAddressResponse) {
    response = &ModifyPartaMailingAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改邮寄地址
func (c *Client) ModifyPartaMailingAddress(request *ModifyPartaMailingAddressRequest) (response *ModifyPartaMailingAddressResponse, err error) {
    if request == nil {
        request = NewModifyPartaMailingAddressRequest()
    }
    response = NewModifyPartaMailingAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPostInfoRequest() (request *ModifyPostInfoRequest) {
    request = &ModifyPostInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyPostInfo")
    return
}

func NewModifyPostInfoResponse() (response *ModifyPostInfoResponse) {
    response = &ModifyPostInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改新增用户的邮寄信息
func (c *Client) ModifyPostInfo(request *ModifyPostInfoRequest) (response *ModifyPostInfoResponse, err error) {
    if request == nil {
        request = NewModifyPostInfoRequest()
    }
    response = NewModifyPostInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReconciliationRequest() (request *ModifyReconciliationRequest) {
    request = &ModifyReconciliationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "ModifyReconciliation")
    return
}

func NewModifyReconciliationResponse() (response *ModifyReconciliationResponse) {
    response = &ModifyReconciliationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提交调账
func (c *Client) ModifyReconciliation(request *ModifyReconciliationRequest) (response *ModifyReconciliationResponse, err error) {
    if request == nil {
        request = NewModifyReconciliationRequest()
    }
    response = NewModifyReconciliationResponse()
    err = c.Send(request, response)
    return
}

func NewPayDealsRequest() (request *PayDealsRequest) {
    request = &PayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "PayDeals")
    return
}

func NewPayDealsResponse() (response *PayDealsResponse) {
    response = &PayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 预付费资源支付
func (c *Client) PayDeals(request *PayDealsRequest) (response *PayDealsResponse, err error) {
    if request == nil {
        request = NewPayDealsRequest()
    }
    response = NewPayDealsResponse()
    err = c.Send(request, response)
    return
}

func NewPrepayDealRequest() (request *PrepayDealRequest) {
    request = &PrepayDealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "PrepayDeal")
    return
}

func NewPrepayDealResponse() (response *PrepayDealResponse) {
    response = &PrepayDealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 预付费资源支付
func (c *Client) PrepayDeal(request *PrepayDealRequest) (response *PrepayDealResponse, err error) {
    if request == nil {
        request = NewPrepayDealRequest()
    }
    response = NewPrepayDealResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPublicAccountRequest() (request *QueryPublicAccountRequest) {
    request = &QueryPublicAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "QueryPublicAccount")
    return
}

func NewQueryPublicAccountResponse() (response *QueryPublicAccountResponse) {
    response = &QueryPublicAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询公共帐户信息
func (c *Client) QueryPublicAccount(request *QueryPublicAccountRequest) (response *QueryPublicAccountResponse, err error) {
    if request == nil {
        request = NewQueryPublicAccountRequest()
    }
    response = NewQueryPublicAccountResponse()
    err = c.Send(request, response)
    return
}

func NewQueryQuotaRequest() (request *QueryQuotaRequest) {
    request = &QueryQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "QueryQuota")
    return
}

func NewQueryQuotaResponse() (response *QueryQuotaResponse) {
    response = &QueryQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户当前使用配额
func (c *Client) QueryQuota(request *QueryQuotaRequest) (response *QueryQuotaResponse, err error) {
    if request == nil {
        request = NewQueryQuotaRequest()
    }
    response = NewQueryQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewRefundRequest() (request *RefundRequest) {
    request = &RefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "Refund")
    return
}

func NewRefundResponse() (response *RefundResponse) {
    response = &RefundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提现（退款）
func (c *Client) Refund(request *RefundRequest) (response *RefundResponse, err error) {
    if request == nil {
        request = NewRefundRequest()
    }
    response = NewRefundResponse()
    err = c.Send(request, response)
    return
}

func NewSetAutoFlagRequest() (request *SetAutoFlagRequest) {
    request = &SetAutoFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "SetAutoFlag")
    return
}

func NewSetAutoFlagResponse() (response *SetAutoFlagResponse) {
    response = &SetAutoFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置续费类型。用户id通过云api公共参数传入
func (c *Client) SetAutoFlag(request *SetAutoFlagRequest) (response *SetAutoFlagResponse, err error) {
    if request == nil {
        request = NewSetAutoFlagRequest()
    }
    response = NewSetAutoFlagResponse()
    err = c.Send(request, response)
    return
}

func NewSetSubUinQuotaRequest() (request *SetSubUinQuotaRequest) {
    request = &SetSubUinQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "SetSubUinQuota")
    return
}

func NewSetSubUinQuotaResponse() (response *SetSubUinQuotaResponse) {
    response = &SetSubUinQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置子账号配额
func (c *Client) SetSubUinQuota(request *SetSubUinQuotaRequest) (response *SetSubUinQuotaResponse, err error) {
    if request == nil {
        request = NewSetSubUinQuotaRequest()
    }
    response = NewSetSubUinQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewSubscribeManualRenewRequest() (request *SubscribeManualRenewRequest) {
    request = &SubscribeManualRenewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "SubscribeManualRenew")
    return
}

func NewSubscribeManualRenewResponse() (response *SubscribeManualRenewResponse) {
    response = &SubscribeManualRenewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 预付费手工续费下单
func (c *Client) SubscribeManualRenew(request *SubscribeManualRenewRequest) (response *SubscribeManualRenewResponse, err error) {
    if request == nil {
        request = NewSubscribeManualRenewRequest()
    }
    response = NewSubscribeManualRenewResponse()
    err = c.Send(request, response)
    return
}

func NewSubscribeOrderRequest() (request *SubscribeOrderRequest) {
    request = &SubscribeOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "SubscribeOrder")
    return
}

func NewSubscribeOrderResponse() (response *SubscribeOrderResponse) {
    response = &SubscribeOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 普通预付费下单  
func (c *Client) SubscribeOrder(request *SubscribeOrderRequest) (response *SubscribeOrderResponse, err error) {
    if request == nil {
        request = NewSubscribeOrderRequest()
    }
    response = NewSubscribeOrderResponse()
    err = c.Send(request, response)
    return
}

func NewSubscribePayRequest() (request *SubscribePayRequest) {
    request = &SubscribePayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "SubscribePay")
    return
}

func NewSubscribePayResponse() (response *SubscribePayResponse) {
    response = &SubscribePayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 普通预付费支付
func (c *Client) SubscribePay(request *SubscribePayRequest) (response *SubscribePayResponse, err error) {
    if request == nil {
        request = NewSubscribePayRequest()
    }
    response = NewSubscribePayResponse()
    err = c.Send(request, response)
    return
}

func NewSubscribeUpgradeRequest() (request *SubscribeUpgradeRequest) {
    request = &SubscribeUpgradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "SubscribeUpgrade")
    return
}

func NewSubscribeUpgradeResponse() (response *SubscribeUpgradeResponse) {
    response = &SubscribeUpgradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 预付费升配下单。
func (c *Client) SubscribeUpgrade(request *SubscribeUpgradeRequest) (response *SubscribeUpgradeResponse, err error) {
    if request == nil {
        request = NewSubscribeUpgradeRequest()
    }
    response = NewSubscribeUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewSyncContractPartaInfoRequest() (request *SyncContractPartaInfoRequest) {
    request = &SyncContractPartaInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "SyncContractPartaInfo")
    return
}

func NewSyncContractPartaInfoResponse() (response *SyncContractPartaInfoResponse) {
    response = &SyncContractPartaInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 同步合同甲方信息
func (c *Client) SyncContractPartaInfo(request *SyncContractPartaInfoRequest) (response *SyncContractPartaInfoResponse, err error) {
    if request == nil {
        request = NewSyncContractPartaInfoRequest()
    }
    response = NewSyncContractPartaInfoResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradePriceRequest() (request *UpgradePriceRequest) {
    request = &UpgradePriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill", APIVersion, "UpgradePrice")
    return
}

func NewUpgradePriceResponse() (response *UpgradePriceResponse) {
    response = &UpgradePriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口无效，请勿调用，管理员请及时下线
func (c *Client) UpgradePrice(request *UpgradePriceRequest) (response *UpgradePriceResponse, err error) {
    if request == nil {
        request = NewUpgradePriceRequest()
    }
    response = NewUpgradePriceResponse()
    err = c.Send(request, response)
    return
}
