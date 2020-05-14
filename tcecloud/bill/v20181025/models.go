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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type  ResourceListData struct {

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 商品资源数量
	Num *uint64 `json:"Num,omitempty" name:"Num"`

	// 商品资源状态-1全部0正常1部分销毁2全部销毁3全部隔离
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ActionDetailItem struct {

	// 操作类型代码
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 操作类型名称
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type AddPartaMailingAddressRequest struct {
	*tchttp.BaseRequest

	// 收件人
	AdresseeName *string `json:"AdresseeName,omitempty" name:"AdresseeName"`

	// 收件人电话
	AdresseePhone *string `json:"AdresseePhone,omitempty" name:"AdresseePhone"`

	// 详细地址
	AdresseeDress *string `json:"AdresseeDress,omitempty" name:"AdresseeDress"`

	// 省份/直辖市
	AdresseeProvince *string `json:"AdresseeProvince,omitempty" name:"AdresseeProvince"`

	// 市/区
	AdresseeCity *string `json:"AdresseeCity,omitempty" name:"AdresseeCity"`

	// 是否为默认地址 1默认 0不默认
	DefaultState *uint64 `json:"DefaultState,omitempty" name:"DefaultState"`

	// 备注
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *AddPartaMailingAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddPartaMailingAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddPartaMailingAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddPartaMailingAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddPartaMailingAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddQuotaPara struct {

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 配额key
	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`

	// 配额值
	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`

	// uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type AddQuotaRequest struct {
	*tchttp.BaseRequest

	// "para": [{
	//                 "uin":"",
	//                 "productCode":"p_cvm",
	//                 "quotaKey":"p_cvm##v_cvm_cpu#",
	//                 "quotaValue":"1000"
	//         },
	//         {
	//                 "uin":"",
	//                 "productCode":"p_cvm",
	//                 "quotaKey":"p_cvm#sp_cvm_vself2##",
	//                 "quotaValue":"1000"
	//         },
	//         {
	//                 "uin":"",
	//                 "productCode":"p_cvm",
	//                 "quotaKey":"p_cvm###",
	//                 "quotaValue":"1000"
	//         }
	//         ]
	SpreadPara []*AddQuotaPara `json:"SpreadPara,omitempty" name:"SpreadPara" list`
}

func (r *AddQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BigDealListData struct {

	// 大订单号
	BigDealId []*int64 `json:"BigDealId,omitempty" name:"BigDealId" list`

	// 用户唯一编码
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 子产品名称
	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`

	// 交易动作
	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`

	// 动作名称
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 订单过期时间
	OverdueTime *string `json:"OverdueTime,omitempty" name:"OverdueTime"`

	// 付款时间
	PayEndTime *string `json:"PayEndTime,omitempty" name:"PayEndTime"`

	// 创建人
	Payer *string `json:"Payer,omitempty" name:"Payer"`

	// 状态 1待支付 2已支付 3发货中 4发货成功 5发货失败 6已退款 7取消 100删除
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// 金额总计
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 代金券金额
	VoucherDecline *string `json:"VoucherDecline,omitempty" name:"VoucherDecline"`

	// 实付金额总计
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 付款人
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 状态名称
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// 小订单列表
	DealList []*DealListData `json:"DealList,omitempty" name:"DealList" list`

	// 折扣金额（因折扣而减少的金额）
	DiscountCost *string `json:"DiscountCost,omitempty" name:"DiscountCost"`
}

type BusinessSummaryDetailItem struct {

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券花费
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 对比上月趋势
	Trend []*SummaryTrend `json:"Trend,omitempty" name:"Trend" list`
}

type BusinessSummaryOverviewItem struct {

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type BusinessSummaryTotal struct {

	// 总花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
}

type CancelInvoiceRequest struct {
	*tchttp.BaseRequest

	// 用户开票流水号
	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *CancelInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelInvoiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelInvoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelInvoiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConditionActionType struct {

	// 交易类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 交易类型名称
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`
}

type ConditionBusiness struct {

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
}

type ConditionComponent struct {

	// 组件码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 组件名称
	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`
}

type ConditionPayMode struct {

	// 付费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
}

type ConditionProduct struct {

	// 产品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 产品名称
	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
}

type ConditionRegion struct {

	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type Conditions struct {

	// 只支持6和12两个值
	TimeRange *uint64 `json:"TimeRange,omitempty" name:"TimeRange"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 付费模式，可选prePay和postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 资源关键字
	ResourceKeyword *string `json:"ResourceKeyword,omitempty" name:"ResourceKeyword"`

	// 产品编码
	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes" list`

	// 子产品编码
	SubProductCodes []*string `json:"SubProductCodes,omitempty" name:"SubProductCodes" list`

	// 地域ID
	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds" list`

	// 付费模式，可选prePay和postPay
	PayModes []*string `json:"PayModes,omitempty" name:"PayModes" list`

	// 交易类型
	ActionTypes []*string `json:"ActionTypes,omitempty" name:"ActionTypes" list`

	// 是否隐藏0元流水
	HideFreeCost *int64 `json:"HideFreeCost,omitempty" name:"HideFreeCost"`

	// 排序规则，可选desc和asc
	OrderByCost *string `json:"OrderByCost,omitempty" name:"OrderByCost"`

	// 交易ID
	BillIds []*string `json:"BillIds,omitempty" name:"BillIds" list`

	// 组件编码
	BillingItemCodes []*string `json:"BillingItemCodes,omitempty" name:"BillingItemCodes" list`

	// 文件ID
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds" list`

	// 文件类型
	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes" list`

	// 状态
	Status []*uint64 `json:"Status,omitempty" name:"Status" list`

	// 计费细项
	SubBillingItemCodes *string `json:"SubBillingItemCodes,omitempty" name:"SubBillingItemCodes"`

	// 导出字段
	ExportFields []*string `json:"ExportFields,omitempty" name:"ExportFields" list`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

type ContractListData struct {

	// 合同编号
	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`

	// 申请时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 合同类型ID号
	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`

	// 合同类型名称
	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`

	// 状态0审核中1已邮寄2审核拒绝3作废4取消
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 操作描述拒绝或通过原因
	OperateDesc *string `json:"OperateDesc,omitempty" name:"OperateDesc"`

	// 备注描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// uin
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 甲方名称
	PartAName *float64 `json:"PartAName,omitempty" name:"PartAName"`

	// 甲方地址
	PartAAddress *string `json:"PartAAddress,omitempty" name:"PartAAddress"`

	// 甲方联系人
	ConnectUser *string `json:"ConnectUser,omitempty" name:"ConnectUser"`

	// 甲方电话
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 甲方邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 收货人名称
	AdresseeName *string `json:"AdresseeName,omitempty" name:"AdresseeName"`

	// 收货人电话
	AdresseePhone *string `json:"AdresseePhone,omitempty" name:"AdresseePhone"`

	// 收货人地址
	AdresseeDress *string `json:"AdresseeDress,omitempty" name:"AdresseeDress"`

	// 模板链接
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

type ContractPartaData struct {

	// 甲方信息ID
	PartaInfoId *uint64 `json:"PartaInfoId,omitempty" name:"PartaInfoId"`

	// 甲方名称
	PartAName *string `json:"PartAName,omitempty" name:"PartAName"`

	// 甲方地址
	PartAAddress *string `json:"PartAAddress,omitempty" name:"PartAAddress"`

	// 甲方联系人
	ConnectUser *string `json:"ConnectUser,omitempty" name:"ConnectUser"`

	// 甲方电话
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 甲方邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 备注描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type ContractTypeListData struct {

	// 合同类型名称
	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`

	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 类型ID
	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`
}

type CouponInfo struct {

	// 持有者
	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`

	// 代金券名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 代金券codeid
	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`

	// 批次id
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 面值
	Value *uint64 `json:"Value,omitempty" name:"Value"`

	// 余额
	Balance *uint64 `json:"Balance,omitempty" name:"Balance"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 到期时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 付费类型
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 发放时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 适用产品
	ProductDefine *string `json:"ProductDefine,omitempty" name:"ProductDefine"`

	// 代金券id
	SpId *string `json:"SpId,omitempty" name:"SpId"`

	// 批次创建时间
	BatchCreateTime *string `json:"BatchCreateTime,omitempty" name:"BatchCreateTime"`

	// 发放者
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
}

type CouponInfoRequest struct {
	*tchttp.BaseRequest

	// 米大师appid
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 指定查询的每页最大记录数，默认10个
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定当前查询第几页，默认第1页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 指定查询状态
	SpStatus *string `json:"SpStatus,omitempty" name:"SpStatus"`

	// 指定产品
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 指定代金券编号
	SpId *string `json:"SpId,omitempty" name:"SpId"`

	// SortField
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// 指定升序降序：desc、asc
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *CouponInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CouponInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CouponInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代金券信息
		CouponsList []*CouponInfo `json:"CouponsList,omitempty" name:"CouponsList" list`

		// 记录总数
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CouponInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CouponInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CouponWater struct {

	// 抵扣订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 抵扣产品
	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 抵扣金额
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`

	// 抵扣时间
	TranTime *uint64 `json:"TranTime,omitempty" name:"TranTime"`
}

type CouponsList struct {

	// 代金券id
	SpId *string `json:"SpId,omitempty" name:"SpId"`

	// Codeid
	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`

	// 批次id
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type CouponsWaterRequest struct {
	*tchttp.BaseRequest

	// 米大师appid
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 代金券CodeId
	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`

	// 分页游标
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页展示记录数.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定代金券编号
	SpId *string `json:"SpId,omitempty" name:"SpId"`

	// 指定代金券批次
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *CouponsWaterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CouponsWaterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CouponsWaterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代金券信息
		CouponsList []*CouponWater `json:"CouponsList,omitempty" name:"CouponsList" list`

		// 记录总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CouponsWaterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CouponsWaterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateContractRequest struct {
	*tchttp.BaseRequest

	// 合同类型ID
	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`

	// 甲方名称
	PartAName *string `json:"PartAName,omitempty" name:"PartAName"`

	// 甲方地址
	PartAAddress *string `json:"PartAAddress,omitempty" name:"PartAAddress"`

	// 甲方联系人
	ConnectUser *string `json:"ConnectUser,omitempty" name:"ConnectUser"`

	// 甲方电话
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 甲方邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 收件人
	AdresseeName *string `json:"AdresseeName,omitempty" name:"AdresseeName"`

	// 收件人电话
	AdresseePhone *string `json:"AdresseePhone,omitempty" name:"AdresseePhone"`

	// 收件人地址
	AdresseeDress *string `json:"AdresseeDress,omitempty" name:"AdresseeDress"`

	// 省/市
	AdresseeProvince *string `json:"AdresseeProvince,omitempty" name:"AdresseeProvince"`

	// 市区
	AdresseeCity *string `json:"AdresseeCity,omitempty" name:"AdresseeCity"`

	// 模板链接
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 备注
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContractRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateContractResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContractResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 文件ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

func (r *CreateDownloadRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDownloadRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDownloadRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDownloadRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInvoiceRequest struct {
	*tchttp.BaseRequest

	// 地址id,用于提取用户地址
	ArrId *int64 `json:"ArrId,omitempty" name:"ArrId"`

	// 开票金额
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 开票账期,可以是多个账期,英文逗号分割,2019-09,2019-10
	AccountPeriods *string `json:"AccountPeriods,omitempty" name:"AccountPeriods"`

	// 发票备注,不传为默认空
	InvoiceRemark *string `json:"InvoiceRemark,omitempty" name:"InvoiceRemark"`
}

func (r *CreateInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInvoiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户开票流水号
		OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInvoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInvoiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePayOrderRequest struct {
	*tchttp.BaseRequest

	// 米大师的appid。由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 支付订单号，需要确保唯一
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 支付金额，单位：分
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`

	// ISO 货币代码，传：CNY
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 商品id。由计平提供
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 商品名称。支付过程会展示给用户看到
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 商品秒速
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`
}

func (r *CreatePayOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePayOrderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePayOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应答结构体
		Response *CreatePayOrderRsp `json:"Response,omitempty" name:"Response"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePayOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePayOrderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePayOrderRsp struct {

	// 计平订单号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 业务订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 支付信息，透传给前端JSAPI
	PayInfo *string `json:"PayInfo,omitempty" name:"PayInfo"`

	// 计平内部流水id
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRemitRecordRequest struct {
	*tchttp.BaseRequest

	// 汇款银行名称
	Bank *string `json:"Bank,omitempty" name:"Bank"`

	// 汇款银行账户
	AccountNum *string `json:"AccountNum,omitempty" name:"AccountNum"`

	// 汇款户名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 汇款金额
	Amount *string `json:"Amount,omitempty" name:"Amount"`

	// 联系人手机
	Tel *string `json:"Tel,omitempty" name:"Tel"`

	// 汇款时间
	RemitDate *string `json:"RemitDate,omitempty" name:"RemitDate"`

	// 备注
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateRemitRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRemitRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRemitRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRemitRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRemitRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWarningRequest struct {
	*tchttp.BaseRequest

	// 预警状态码 0:关，1:开
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 预警额度值
	BalanceThreshold *float64 `json:"BalanceThreshold,omitempty" name:"BalanceThreshold"`
}

func (r *CreateWarningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWarningRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWarningResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWarningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWarningResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DealListData struct {

	// 小订单号
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子商品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 商品码名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 子商品码名称
	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`

	// 类型，新购续费变配等
	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 提单人
	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`

	// 状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 订单详情，jason字符串
	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

	// 付费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 订单数量
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 业务返回资源ID详情，jason字符串
	TaskDetail *string `json:"TaskDetail,omitempty" name:"TaskDetail"`

	// 类型中文名称
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 状态名称
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// 用户uin
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 升配公式
	Formula *string `json:"Formula,omitempty" name:"Formula"`

	// 区域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用区ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 折扣，范围：0~1
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 折扣金额（因折扣减少的金额）
	DiscountCost *string `json:"DiscountCost,omitempty" name:"DiscountCost"`

	// 订单价格详情（包括折扣详情），json字符串
	// {
	//     "policyDetail":{
	//         "total":60,   //总的折扣，100表示100%不打折
	//         "user":60,  //用户个人折扣
	//         "common":100 //官网基础折扣
	//     },
	//     "voucherList":[
	//         {
	//             "id":"xxxx",
	//             "voucherAmount":600
	//         }
	//     ]
	// }
	GoodsPrice *string `json:"GoodsPrice,omitempty" name:"GoodsPrice"`
}

type DebtBillListData struct {

	// 账单号
	DebtBillNo *string `json:"DebtBillNo,omitempty" name:"DebtBillNo"`

	// 账单状态，数字 1-已还清 2-未还清
	DebtBillStatus *string `json:"DebtBillStatus,omitempty" name:"DebtBillStatus"`

	// 账单月份yyyymm
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 账单结束日期yyyymmdd（灵活出账周期模式）
	BillDateEnd *string `json:"BillDateEnd,omitempty" name:"BillDateEnd"`

	// 最后还款时间，UNIX时间戳
	DueDate *string `json:"DueDate,omitempty" name:"DueDate"`

	// 逾期利息
	DebtInterest *string `json:"DebtInterest,omitempty" name:"DebtInterest"`

	// 欠款总额
	TotalDebt *string `json:"TotalDebt,omitempty" name:"TotalDebt"`

	// 当前仍欠金额
	CurrentDebt *string `json:"CurrentDebt,omitempty" name:"CurrentDebt"`

	// 已还金额
	TotalRefund *string `json:"TotalRefund,omitempty" name:"TotalRefund"`

	// 还款记录，格式（time为UNIX时戳）： time1,refundbillno1,amt1;time2,refundbillno2,amt2;….
	RefundRecord *string `json:"RefundRecord,omitempty" name:"RefundRecord"`

	// 账期类型，1：按月，2：按天
	DuedelayType *string `json:"DuedelayType,omitempty" name:"DuedelayType"`

	// 账期，还款允许延期时间，当账期类型按月时，表示月份数，按天时，表示天数
	DueDelay *string `json:"DueDelay,omitempty" name:"DueDelay"`

	// 账单生成时间，UNIX时间戳
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 账单最后更新时间，UNIX时间戳
	LastChgTimes *string `json:"LastChgTimes,omitempty" name:"LastChgTimes"`

	// 账单备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DeductList struct {

	// 代金券id
	SpId *string `json:"SpId,omitempty" name:"SpId"`

	// codeid
	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`

	// 批次id
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 抵扣金额
	DeductAmount *uint64 `json:"DeductAmount,omitempty" name:"DeductAmount"`
}

type DeletePartaMailingAddressRequest struct {
	*tchttp.BaseRequest

	// 邮寄地址ID
	PartaMailingAddressId *uint64 `json:"PartaMailingAddressId,omitempty" name:"PartaMailingAddressId"`
}

func (r *DeletePartaMailingAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePartaMailingAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePartaMailingAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePartaMailingAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePartaMailingAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePostInfoRequest struct {
	*tchttp.BaseRequest

	// 地址的唯一识别号
	ArrId *int64 `json:"ArrId,omitempty" name:"ArrId"`
}

func (r *DeletePostInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePostInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePostInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePostInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePostInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteQuotaRequest struct {
	*tchttp.BaseRequest

	// uin
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`

	// 商品码，商品码和QuotaKey不能同时为空
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 配额key
	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
}

func (r *DeleteQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubUinQuotaRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 当前父账户Uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 当前子账户Uin
	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`

	// 商品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子商品代码（可为空）
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费项代码（可为空）
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费细项代码（可为空）
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
}

func (r *DeleteSubUinQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubUinQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubUinQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubUinQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubUinQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescriPartaMailingAddressListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescriPartaMailingAddressListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescriPartaMailingAddressListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescriPartaMailingAddressListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 合同列表
		List *MailingAddressDataList `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescriPartaMailingAddressListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescriPartaMailingAddressListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPeriodRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPeriodRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPeriodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账期
		AccountPeriod *string `json:"AccountPeriod,omitempty" name:"AccountPeriod"`

		// 此账期的消耗金额,单位分
		Amount *int64 `json:"Amount,omitempty" name:"Amount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPeriodResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountWaterRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识
	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`

	// 传后付费固定额度子账户标识: CREDIT_FIXED
	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`

	// 请求方设备，0表示PC，1表示手机
	Device *int64 `json:"Device,omitempty" name:"Device"`

	// UNIX时间戳（从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数）
	Ts *int64 `json:"Ts,omitempty" name:"Ts"`

	// 查询流水开始时间，数字，年月日20150617或年月日时分秒20150617111855
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询流水结束时间，数字，年月日20150617或年月日时分秒20150617111855
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// num，查询流水条数（最大支持20）
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// cursor，查询流水的游标
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 若传了， 则根据流水类型查询。目前支持：
	// RefundDebt ： 还款流水
	// WriteOff：销账流水
	WaterFilter *string `json:"WaterFilter,omitempty" name:"WaterFilter"`

	// 需要查询销账记录时，可传入对应还款订单号
	BillNo *string `json:"BillNo,omitempty" name:"BillNo"`
}

func (r *DescribeAccountWaterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountWaterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountWaterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 所有流水总数
		AllNum *string `json:"AllNum,omitempty" name:"AllNum"`

		// 此处返回的流水条数
		RetNum *string `json:"RetNum,omitempty" name:"RetNum"`

		// cursor，查询流水的当前游标值
		Offset *string `json:"Offset,omitempty" name:"Offset"`

		// 流水数据列表
		WaterSet []*WaterListData `json:"WaterSet,omitempty" name:"WaterSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountWaterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountWaterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDetailByResourceRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 付费模式，只能是prePay或者postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 预付费才需要传，为返回的BillId
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 预付费才需要传，为返回的Id
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBillDetailByResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDetailByResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDetailByResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 资源花费详情
		Total *ResourceSummaryDetailTotal `json:"Total,omitempty" name:"Total"`

		// 组件花费详情
		BillingItemDetail *ResourceSummaryDetailComponentDetailItem `json:"BillingItemDetail,omitempty" name:"BillingItemDetail"`

		// 资源历史月份花费趋势
		RealTotalCostTrend *ResourceSummaryDetailRealTotalCostTrend `json:"RealTotalCostTrend,omitempty" name:"RealTotalCostTrend"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDetailByResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDetailByResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadByResourceRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 需要导出相关字段
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillDownloadByResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadByResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadByResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务是否已完成（0未完成/1已完成）
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 文件名称，当任务已完成时返回该字段
		Name *string `json:"Name,omitempty" name:"Name"`

		// 下载链接，当任务已完成时返回该字段
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadByResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadByResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadCommonSummaryRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillDownloadCommonSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadCommonSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadCommonSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务是否已完成（0未完成/1已完成）
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 文件名称，当任务已完成时返回该字段
		Name *string `json:"Name,omitempty" name:"Name"`

		// 下载链接，当任务已完成时返回该字段
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadCommonSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadCommonSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadListRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 开始月份
	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`

	// 结束月份
	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`
}

func (r *DescribeBillDownloadListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 下载地址列表
		DownloadList *GetPkgDownloadListItem `json:"DownloadList,omitempty" name:"DownloadList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadRecordListRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否需要记录条数，0需要，1不需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 排序规则
	Sort *GetDownloadListSort `json:"Sort,omitempty" name:"Sort"`

	// 只支持FIleIds，FileTypes，Status三种过滤条件
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`

	// 是否需要条件取值（0不需要/1需要，默认为0）
	NeedConditionValue *string `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`
}

func (r *DescribeBillDownloadRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadRecordListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadRecordListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录数量，NeedRecordNum为0时返回null
		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

		// 下载记录列表
		Records *GetDownloadListRecordItem `json:"Records,omitempty" name:"Records"`

		// 过滤条件，NeedConditionValue为0是返回null
		ConditionValue *DetailConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadRecordListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadResourceDetailRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 需要导出相关字段
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillDownloadResourceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadResourceDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadResourceDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务是否已完成（0未完成/1已完成）
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 文件名称，当任务已完成时返回该字段
		Name *string `json:"Name,omitempty" name:"Name"`

		// 下载链接，当任务已完成时返回该字段
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadResourceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDownloadResourceDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByPayModeRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillFeeByPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillFeeByPayModeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByPayModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 计费模式花费详情
		SummaryDetail *PayModeSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillFeeByPayModeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByProductRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否返回最大页码，建议首次或变更PageSize后才需要，用于提升接口性
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeBillFeeByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillFeeByProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 记录数，不传NeedRecordNum或者传的值为0，该字段值为null
		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

		// 产品花费详情
		SummaryDetail *BusinessSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillFeeByProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByRegionRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否返回最大页码，建议首次或变更PageSize后才需要，用于提升接口性
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeBillFeeByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillFeeByRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 地域花费详情
		SummaryDetail *RegionSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`

		// 记录数，不传NeedRecordNum或者传的值为0，该字段值为null
		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillFeeByRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeTrendRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 只支持ProductCodes，RegionId，PayMode
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillFeeTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillFeeTrendRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 每月花费详情
		Detail *MonthCostDetailItem `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillFeeTrendResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByPayModeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 付费模式花费分布
		SummaryOverview *PayModeSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByPayModeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 总花费详情
		SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`

		// 各产品花费分布
		SummaryOverview *BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 各地域花费分布详情
		SummaryOverview *RegionSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByResourceRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否需要记录总数，0不需要，1需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 是否需要过滤条件，0需要，1不需要
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`

	// 只支持BusinessCodes，ProductCodes，PayModes, ProjectIds, RegionIds, PayModes, ActionTypes, HideFreeCost, ResourceKeyword, OrderByCost
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillSummaryByResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 资源汇总详情
		Data *SummaryByResourceDataItem `json:"Data,omitempty" name:"Data"`

		// 资源汇总花费详情
		Total []*SummaryByResourceTotal `json:"Total,omitempty" name:"Total" list`

		// 记录数量，NeedRecordNum为0时该值为null
		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

		// 过滤条件，NeedConditionValue为0时该值为null
		ConditionValue *SummaryByResourceConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillTrendByMonthRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *uint64 `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 只支持TimeRange
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillTrendByMonthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillTrendByMonthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillTrendByMonthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 每月花费详情
		Detail []*TrendByMonthDetail `json:"Detail,omitempty" name:"Detail" list`

		// 平均花费详情
		Stat []*TrendByMonthStat `json:"Stat,omitempty" name:"Stat" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillTrendByMonthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillTrendByMonthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConsolidatedBillDownloadUrlRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 子uin列表
	ChildUins []*string `json:"ChildUins,omitempty" name:"ChildUins" list`

	// 内部操作者
	InnerOperator *string `json:"InnerOperator,omitempty" name:"InnerOperator"`
}

func (r *DescribeConsolidatedBillDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConsolidatedBillDownloadUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConsolidatedBillDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务是否完成，0未完成，1完成
		TaskFinished *uint64 `json:"TaskFinished,omitempty" name:"TaskFinished"`

		// 文件下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 文件名称
		FileName *string `json:"FileName,omitempty" name:"FileName"`

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 任务完成时间
		FinishedTime *string `json:"FinishedTime,omitempty" name:"FinishedTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsolidatedBillDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConsolidatedBillDownloadUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContractListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 合同状态0审核中1已邮寄2审核拒绝3作废4取消
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 合同类型ID号
	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`

	// 申请时间排序 1正序 2倒序
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 合同ID
	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`
}

func (r *DescribeContractListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContractListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContractListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 合同列表
		List *ContractListData `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContractListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContractPartaInfoRequest struct {
	*tchttp.BaseRequest

	// 甲方信息ID
	PartaInfoId *uint64 `json:"PartaInfoId,omitempty" name:"PartaInfoId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContractPartaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContractPartaInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContractPartaInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 甲方合同信息
		List *ContractPartaData `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractPartaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContractPartaInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContractTypeListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeContractTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContractTypeListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContractTypeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 合同类型列表
		List *ContractTypeListData `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContractTypeListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCouponsDeductInfoRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 使用抵扣券列表
	CouponsList []*CouponsList `json:"CouponsList,omitempty" name:"CouponsList" list`

	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合。
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 商品列表列表
	ProductList []*ProductList `json:"ProductList,omitempty" name:"ProductList" list`

	// 升配时物品信息
	UpgradeProductInfo *UpgradeProductInfo `json:"UpgradeProductInfo,omitempty" name:"UpgradeProductInfo"`
}

func (r *DescribeCouponsDeductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCouponsDeductInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCouponsDeductInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总抵扣金额，tce场景下，单位为微元
		DeductTotalAmount *uint64 `json:"DeductTotalAmount,omitempty" name:"DeductTotalAmount"`

		// 抵扣信息。
		DeductList []*DeductList `json:"DeductList,omitempty" name:"DeductList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCouponsDeductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCouponsDeductInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDealListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 订单号
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 订单支付人uin
	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 订单类型
	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`

	// 状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 排序asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 大订单或小订单
	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`

	// 付费模式，0后付费/1预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 动作名称 隔离 销毁 返还 新购 续费 升配 降配
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 订单创建人 uin
	Creater *string `json:"Creater,omitempty" name:"Creater"`
}

func (r *DescribeDealListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDealListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDealListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 订单列表
		List *BigDealListData `json:"List,omitempty" name:"List"`

		// 服务器当前时间戳
		Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDealListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDealListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDealPriceRequest struct {
	*tchttp.BaseRequest

	// 商品询价参数
	ResInfo *string `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *DescribeDealPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDealPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDealPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 询价返回包
		List *PriceListData `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDealPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDealPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadVoucherListFileRequest struct {
	*tchttp.BaseRequest

	// 代金券筛选条件对象
	VoucherConditions *VoucherConditions `json:"VoucherConditions,omitempty" name:"VoucherConditions"`
}

func (r *DescribeDownloadVoucherListFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadVoucherListFileRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadVoucherListFileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 文件名称
		FileName *string `json:"FileName,omitempty" name:"FileName"`

		// 任务是否完成，0未完成，1完成
		TaskFinished *uint64 `json:"TaskFinished,omitempty" name:"TaskFinished"`

		// 任务id
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadVoucherListFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadVoucherListFileResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadWithAuthRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

func (r *DescribeDownloadWithAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadWithAuthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadWithAuthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadWithAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadWithAuthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGoodsPriceRequest struct {
	*tchttp.BaseRequest

	// 商品询价参数
	ResInfo []*Goods `json:"ResInfo,omitempty" name:"ResInfo" list`
}

func (r *DescribeGoodsPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGoodsPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGoodsPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGoodsPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGoodsPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceAmountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInvoiceAmountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvoiceAmountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceAmountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可开票金额,单位分
		AvailableOpenAmount *int64 `json:"AvailableOpenAmount,omitempty" name:"AvailableOpenAmount"`

		// 消耗金额,单位分
		Consume *int64 `json:"Consume,omitempty" name:"Consume"`

		// 已开票金额,单位分
		OpenedAmount *int64 `json:"OpenedAmount,omitempty" name:"OpenedAmount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvoiceAmountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvoiceAmountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceDetailRequest struct {
	*tchttp.BaseRequest

	// 用户开票流水号
	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *DescribeInvoiceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvoiceDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户开票流水号
		OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`

		// 发票金额
		Amount *int64 `json:"Amount,omitempty" name:"Amount"`

		// 申请开票时间
		ApplicationTime *string `json:"ApplicationTime,omitempty" name:"ApplicationTime"`

		// 发票类型,取值为(  2:公司普通发票  3:公司增值税发票 )
		UserType *int64 `json:"UserType,omitempty" name:"UserType"`

		// 发票抬头
		InvoiceHead *string `json:"InvoiceHead,omitempty" name:"InvoiceHead"`

		// 税务登记号
		TaxNumber *string `json:"TaxNumber,omitempty" name:"TaxNumber"`

		// 开户行
		BankUserName *string `json:"BankUserName,omitempty" name:"BankUserName"`

		// 银行账号
		AccountNumber *string `json:"AccountNumber,omitempty" name:"AccountNumber"`

		// 注册全地
		RegisterFullAddr *string `json:"RegisterFullAddr,omitempty" name:"RegisterFullAddr"`

		// 注册电话
		RegisterPhone *string `json:"RegisterPhone,omitempty" name:"RegisterPhone"`

		// 地址
		Addr *string `json:"Addr,omitempty" name:"Addr"`

		// 联系人
		Contact *string `json:"Contact,omitempty" name:"Contact"`

		// 手机号码
		Cellphone *string `json:"Cellphone,omitempty" name:"Cellphone"`

		// 发票状态,1:处理中,6:已取消 30:审核通过  31:审核不通过
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 发票备注
		InvoiceRemark *string `json:"InvoiceRemark,omitempty" name:"InvoiceRemark"`

		// 城市
		City *string `json:"City,omitempty" name:"City"`

		// 省份
		Province *string `json:"Province,omitempty" name:"Province"`

		// 账期
		AccountPeriods *string `json:"AccountPeriods,omitempty" name:"AccountPeriods"`

		// 审核操作人
		Auditor *string `json:"Auditor,omitempty" name:"Auditor"`

		// 审核人拒绝原因
		Reason *string `json:"Reason,omitempty" name:"Reason"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvoiceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvoiceDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceListRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页限制数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 开始时间 形如'Y-m-d H:m:s'
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间 形如'Y-m-d H:m:s'
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 根据申请时间排序,0:降序,1:升序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 搜索的状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeInvoiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvoiceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户开票流水号
		OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`

		// 发票类型
		UserType *int64 `json:"UserType,omitempty" name:"UserType"`

		// 金额
		Amount *int64 `json:"Amount,omitempty" name:"Amount"`

		// 抬头
		InvoiceHead *string `json:"InvoiceHead,omitempty" name:"InvoiceHead"`

		// 发票状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 申请时间时间戳
		ApplicationTime *string `json:"ApplicationTime,omitempty" name:"ApplicationTime"`

		// 备注
		InvoiceRemark *string `json:"InvoiceRemark,omitempty" name:"InvoiceRemark"`

		// 总和
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当前页数
		Cur *int64 `json:"Cur,omitempty" name:"Cur"`

		// 总页数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvoiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvoiceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenInvoiceInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOpenInvoiceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOpenInvoiceInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenInvoiceInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发票类型,取值为( 2:公司普通发票 3:公司增值税发票 )
		UserType *int64 `json:"UserType,omitempty" name:"UserType"`

		// 抬头
		InvoiceHead *string `json:"InvoiceHead,omitempty" name:"InvoiceHead"`

		// 税务登记号
		TaxNumber *string `json:"TaxNumber,omitempty" name:"TaxNumber"`

		// 开户行
		BankUserName *string `json:"BankUserName,omitempty" name:"BankUserName"`

		// 银行账号
		AccountNumber *string `json:"AccountNumber,omitempty" name:"AccountNumber"`

		// 注册全地址
		RegisterFullAddr *string `json:"RegisterFullAddr,omitempty" name:"RegisterFullAddr"`

		// 注册固定电话
		RegisterPhone *string `json:"RegisterPhone,omitempty" name:"RegisterPhone"`

		// 修改时间
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// 用户是否设置了发票信息，取值为（0：未设置，1：已设置）
		IsSetOpenInfo *int64 `json:"IsSetOpenInfo,omitempty" name:"IsSetOpenInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenInvoiceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOpenInvoiceInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayOrderRequest struct {
	*tchttp.BaseRequest

	// 米大师appid，由计平分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// type=by_order，根据订单号查订单
	// type=by_user，根据用户id查订单
	Type *string `json:"Type,omitempty" name:"Type"`

	// 精准查询，暂不支持批量。支付订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 查询的起始时间，unix时间戳,精确到秒。(订单创建时间)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的结束时间，unix时间戳,精确到秒。(订单创建时间)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 记录数偏移量，默认从0开始。根据用户号码查询订单列表时需要传。用于分页展示。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 支付渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 提现标记.0-可提现, 1-已提现
	RefundFlag *string `json:"RefundFlag,omitempty" name:"RefundFlag"`

	// 订单状态，用于过滤。 本场景固定传:["2","3","4"]
	OrderStateList []*string `json:"OrderStateList,omitempty" name:"OrderStateList" list`

	// 每页返回的记录数。根据用户号码查询订单列表时需要传。用于分页展示。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePayOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePayOrderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单数量
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 订单详情
		OrderList []*OrderInfo `json:"OrderList,omitempty" name:"OrderList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePayOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePayOrderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePostInfoListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePostInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePostInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePostInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地址的唯一识别号
		ArrId *int64 `json:"ArrId,omitempty" name:"ArrId"`

		// 联系人
		Contact *string `json:"Contact,omitempty" name:"Contact"`

		// 省份
		Province *string `json:"Province,omitempty" name:"Province"`

		// 城市
		City *string `json:"City,omitempty" name:"City"`

		// 详细地址
		Addr *string `json:"Addr,omitempty" name:"Addr"`

		// 手机号码
		Cellphone *string `json:"Cellphone,omitempty" name:"Cellphone"`

		// 修改时间
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// 0:不是默认地址,1:是默认地址
		IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePostInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePostInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRelationsRequest struct {
	*tchttp.BaseRequest

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 商品代码或名称
	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`

	// 子商品代码或名称
	SearchSubProduct *string `json:"SearchSubProduct,omitempty" name:"SearchSubProduct"`

	// 计费项代码或名称
	SearchBillingItem *string `json:"SearchBillingItem,omitempty" name:"SearchBillingItem"`

	// 计费细项代码或名称
	SearchSubBillingItem *string `json:"SearchSubBillingItem,omitempty" name:"SearchSubBillingItem"`

	// 四层定义状态，1-生效 0-失效
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 商品代码数组
	SearchProductCodes []*string `json:"SearchProductCodes,omitempty" name:"SearchProductCodes" list`

	// 子商品代码数组
	SearchSubProductCodes []*string `json:"SearchSubProductCodes,omitempty" name:"SearchSubProductCodes" list`

	// 计费项代码数组
	SearchBillingItemCodes []*string `json:"SearchBillingItemCodes,omitempty" name:"SearchBillingItemCodes" list`

	// 计费细项代码数组
	SearchSubBillingItemCodes []*string `json:"SearchSubBillingItemCodes,omitempty" name:"SearchSubBillingItemCodes" list`
}

func (r *DescribeProductRelationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductRelationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRelationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 产品定义列表
		List []*ProductDefineList `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductRelationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductRelationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest

	// 多个产品代码,例如['p_cvm','p_cvm2'...]
	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes" list`

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品列表
		List []*Product `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePropertiesRequest struct {
	*tchttp.BaseRequest

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// key码，空时返回所有
	KeyCode *string `json:"KeyCode,omitempty" name:"KeyCode"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 类型： 0（或者不传或者传空）是普通属性（默认）；1是自定义属性
	PropertyType *int64 `json:"PropertyType,omitempty" name:"PropertyType"`
}

func (r *DescribePropertiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePropertiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePropertiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总行数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 售卖属性列表
		List *ProductPropertyList `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePropertiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePropertiesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 游标，从第0开始取
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 总共取多少条
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额信息
		QuotaInfo []*QuotaInfo `json:"QuotaInfo,omitempty" name:"QuotaInfo" list`

		// 记录总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReconciliationListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 调账记录查询调账列表部分参数
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeReconciliationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReconciliationListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReconciliationListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 偏移量，与请求包一致
		Limit *string `json:"Limit,omitempty" name:"Limit"`

		// 偏移量，与请求包一致
		Offset *string `json:"Offset,omitempty" name:"Offset"`

		// 记录数，若调用时传的needRecordNum=0，则不返回该字段
		RecordNum *string `json:"RecordNum,omitempty" name:"RecordNum"`

		// 资源汇总详情
		Data []*SummaryByResourceDataItem `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReconciliationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReconciliationListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRefundRequest struct {
	*tchttp.BaseRequest

	// 米大师appid，由计平分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// type=by_refund_id，根据订单号查订单. type=by_user，根据用户Id查询订单.
	Type *string `json:"Type,omitempty" name:"Type"`

	// 查询的起始时间，unix时间戳（格林威治时间）,精确到秒
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的结束时间，unix时间戳（格林威治时间）,精确到秒
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每页返回的记录数。根据用户号码查询订单列表时需要传。用于分页展示
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 记录数偏移量，默认从0开始。根据用户号码查询订单列表时需要传。用于分页展示
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 指定退款id查询。
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 本场景固定传: unionpay
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 按状态过滤.提现状态.1-提现中;2-提现成功;3-提现失败. 
	// 不传默认查询全部
	State []*string `json:"State,omitempty" name:"State" list`
}

func (r *DescribeRefundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRefundRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRefundResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 提现(退款)明细
		RefundDetail []*RefundDetail `json:"RefundDetail,omitempty" name:"RefundDetail" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRefundResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRefundResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 付费类型.0:后付费,1:预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 类型.固定传:query_region_by_uin
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地域列表信息
		RegionDetail *RegionDetail `json:"RegionDetail,omitempty" name:"RegionDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordRequest struct {
	*tchttp.BaseRequest

	// 汇款户名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 汇款账号
	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 状态 0：处理中 1：成功 2：失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeRemitRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRemitRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 汇款记录列表
		List *RemitRecordData `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemitRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRemitRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRenewInfoRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 0:后付费,1:预付费。不传查询全部
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 自动续费模式： 
	// 1：自动续费; 
	// 0:手动续费,到期后一定间隔时间后回收; 
	// 2:到期关闭
	AutopayType *uint64 `json:"AutopayType,omitempty" name:"AutopayType"`

	// 状态。0：正常；1：销毁；2：隔离;不传查询全部
	Status *string `json:"Status,omitempty" name:"Status"`

	// 地域id
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 指定排序字段.取值:begin_time、end_time
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// 指定升序降序：desc、asc
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 游标
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 记录数目,最小值为1,最大值为10.超过范围则以最值代替
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 固定填:query_renew_info
	Type *string `json:"Type,omitempty" name:"Type"`

	// 开始时间
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" name:"EndTimeBegin"`

	// 结束时间
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" name:"EndTimeEnd"`
}

func (r *DescribeRenewInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRenewInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRenewInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 续费信息
		RenewDetail *RenewDetail `json:"RenewDetail,omitempty" name:"RenewDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRenewInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRenewInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDetailRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否需要记录总数，0不需要，1需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 是否需要过滤条件，0需要，1不需要
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`

	// 只支持BusinessCodes，ProductCodes，ComponentCodes，PayMode, ProjectIds, RegionIds, PayModes, ActionTypes, BillIds, HideFreeCost, ResourceKeyword
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeResourceBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceBillDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否展示数据，0不展示，1展示
		DisplayData *uint64 `json:"DisplayData,omitempty" name:"DisplayData"`

		// 记录数量，NeedRecordNum为0是返回null
		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

		// 过滤条件，NeedConditionValue为0是返回null
		ConditionValue *DetailConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`

		// 总花费详情
		Total *DetailTotal `json:"Total,omitempty" name:"Total"`

		// 组件花费详情
		Detail *DetailDetailItem `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceBillDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDetailListRequest struct {
	*tchttp.BaseRequest

	// 资源信息
	ResourcePara *string `json:"ResourcePara,omitempty" name:"ResourcePara"`
}

func (r *DescribeResourceDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceDetailListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDetailListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 续费参数信息
		List []*string `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceDetailListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 交易类型 0后付费 1预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *DescribeResourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 资源列表
		List []* ResourceListData `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubUinQuotasRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 商品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子商品代码（可为空）
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费项代码（可为空）
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费细项代码（可为空）
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 当前父账户Uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *DescribeSubUinQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubUinQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubUinQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 父账户可分配总配额
		AvailableTotalQuota *int64 `json:"AvailableTotalQuota,omitempty" name:"AvailableTotalQuota"`

		// 子账户配额信息列表
		SubUinQuotaArray []*SubUinQuotaArray `json:"SubUinQuotaArray,omitempty" name:"SubUinQuotaArray" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubUinQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubUinQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantQuotasRequest struct {
	*tchttp.BaseRequest

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 产品code。不传或者空就是查询全部
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 获取系统默认配额数据标识。1 获取系统默认配额数据；0或不传 获取用户Uin配额数据
	GetSystemDefaultData *int64 `json:"GetSystemDefaultData,omitempty" name:"GetSystemDefaultData"`
}

func (r *DescribeTenantQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTenantQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额数据列表
		List []*QuotaList `json:"List,omitempty" name:"List" list`

		// 数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTenantQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAccountRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识
	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`

	// 传后付费固定额度子账户标识: CREDIT_FIXED
	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`

	// 请求方设备，0表示PC，1表示手机
	Device *int64 `json:"Device,omitempty" name:"Device"`

	// UNIX时间戳（从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数）
	Ts *int64 `json:"Ts,omitempty" name:"Ts"`
}

func (r *DescribeUserAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账户状态，0为未开户或销户，1为正常，其他异常
	// （可根据账户状态来判断是否开户）
		Status *string `json:"Status,omitempty" name:"Status"`

		// 后付费余额（包含了cash）
		Balance *string `json:"Balance,omitempty" name:"Balance"`

		// 累计金额，后付费账户表示累积还款额（当前接口无需关心）
		AllIn *string `json:"AllIn,omitempty" name:"AllIn"`

		// 历史累计支出金额，后付费账户表示累积消耗额（当前接口无需关心）
		AllOut *string `json:"AllOut,omitempty" name:"AllOut"`

		// 过期时间，非过期类返回0，否则返回时间UNIX时间戳（当前接口无需关心）
		Exptime *string `json:"Exptime,omitempty" name:"Exptime"`

		// 过期订单信息（当前接口无需关心）
		ExpireinfoExtend *string `json:"ExpireinfoExtend,omitempty" name:"ExpireinfoExtend"`

		// 后付费额度上限
		Dplimit *string `json:"Dplimit,omitempty" name:"Dplimit"`

		// 后付费剩余可用额度（包含了cash）
		Dpbalance *string `json:"Dpbalance,omitempty" name:"Dpbalance"`

		// 当前额度使用量（包含转出和预授权等）
		Debt *string `json:"Debt,omitempty" name:"Debt"`

		// 后付费账户，存储现金/溢出金额，下次出账时抵消欠款
		Cash *string `json:"Cash,omitempty" name:"Cash"`

		// 出账日，1-28的数字，每个月第几天出账单
		BillingDate *string `json:"BillingDate,omitempty" name:"BillingDate"`

		// 还款日，1-28的数字，每个月第几天还款
		DueDate *string `json:"DueDate,omitempty" name:"DueDate"`

		// 商户类型，1：代理，2：子客，3：直客
		MerchantType *string `json:"MerchantType,omitempty" name:"MerchantType"`

		// 关联商户信息，当商户类型为代理或直客时，该字段为空；当为子客时，标识代理账户信息
		RelatedMerchant *string `json:"RelatedMerchant,omitempty" name:"RelatedMerchant"`

		// 账期类型，1：按月，2：按天
		DueDelayType *string `json:"DueDelayType,omitempty" name:"DueDelayType"`

		// 账期，还款允许延期时间，当账期类型按月时，表示月份数，按天时，表示天数
		DueDelay *string `json:"DueDelay,omitempty" name:"DueDelay"`

		// 下一个出账日将生效的新账期（无新账期则返回空）
		NextDueDelay *string `json:"NextDueDelay,omitempty" name:"NextDueDelay"`

		// 还款时账单未还清是否立即恢复额度
	// 0 (默认)不立即恢复 1 立即恢复
		RecoverImmediately *string `json:"RecoverImmediately,omitempty" name:"RecoverImmediately"`

		// 是否校验账单还款顺序
	// 0 (默认)按顺序  1 不按顺序
		RefundSequence *string `json:"RefundSequence,omitempty" name:"RefundSequence"`

		// 提前还款是否立即恢复额度
	// 0 (默认)不立即恢复，出账时抵消账单才恢复；  1 立即恢复
		CashRecoverImmediately *string `json:"CashRecoverImmediately,omitempty" name:"CashRecoverImmediately"`

		// 出账周期类型，1：按月。其他暂不支持
		BillingCycleType *string `json:"BillingCycleType,omitempty" name:"BillingCycleType"`

		// 出账周期，若按月则表示月份数，按天则表示天数
		BillingCycle *string `json:"BillingCycle,omitempty" name:"BillingCycle"`

		// 当天累计消耗
		DayOut *string `json:"DayOut,omitempty" name:"DayOut"`

		// 后付费二级额度。未设置则返回-1
		SubDplimit *string `json:"SubDplimit,omitempty" name:"SubDplimit"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDebtBillRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识
	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`

	// 传后付费固定额度子账户标识: CREDIT_FIXED
	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`

	// 请求方设备，0表示PC，1表示手机
	Device *int64 `json:"Device,omitempty" name:"Device"`

	// UNIX时间戳（从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数）
	Ts *int64 `json:"Ts,omitempty" name:"Ts"`

	// 账单起始年月yyyymm
	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`

	// 账单结束年月yyyymm
	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`

	// 查询的账单还款状态，数字 1-已还清 2-未还清
	DebtbillStatus *int64 `json:"DebtbillStatus,omitempty" name:"DebtbillStatus"`

	// 查询的账单是否逾期，数字 1-未逾期 2-逾期
	OverDue *int64 `json:"OverDue,omitempty" name:"OverDue"`

	// num，查询账单条数（最大支持20）
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// cursor，查询流水的游标
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUserDebtBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserDebtBillRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDebtBillResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账单详情列表
		DebtBillSet []*DebtBillListData `json:"DebtBillSet,omitempty" name:"DebtBillSet" list`

		// 账单总数
		AllNum *string `json:"AllNum,omitempty" name:"AllNum"`

		// 返回的账单数量
		RetNum *string `json:"RetNum,omitempty" name:"RetNum"`

		// 当前游标
		Offset *string `json:"Offset,omitempty" name:"Offset"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserDebtBillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserDebtBillResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWarningRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWarningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWarningRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWarningResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 租户uin
		Uin *string `json:"Uin,omitempty" name:"Uin"`

		// 预警功能状态
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 预警阈值
		BalanceThreshold *float64 `json:"BalanceThreshold,omitempty" name:"BalanceThreshold"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWarningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWarningResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetailConditionValue struct {

	// 产品列表
	Product *ConditionBusiness `json:"Product,omitempty" name:"Product"`

	// 子产品列表
	SubProduct *ConditionProduct `json:"SubProduct,omitempty" name:"SubProduct"`

	// 组件列表
	BillingItem *ConditionComponent `json:"BillingItem,omitempty" name:"BillingItem"`

	// 地域列表
	Region *ConditionRegion `json:"Region,omitempty" name:"Region"`

	// 付费模式列表
	PayMode *ConditionPayMode `json:"PayMode,omitempty" name:"PayMode"`

	// 交易类型列表
	ActionType *ConditionActionType `json:"ActionType,omitempty" name:"ActionType"`
}

type DetailDetailItem struct {

	// 资源所有者uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 支付者uin
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 操作者uin
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 付费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 子产品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品名称
	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`

	// 组件类型
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 组件类型名称
	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`

	// 组件码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 组件名称
	SubBillingItemCodeName *string `json:"SubBillingItemCodeName,omitempty" name:"SubBillingItemCodeName"`

	// 交易类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 交易类型名称
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域类型
	RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

	// 地域类型名称
	RegionTypeName *string `json:"RegionTypeName,omitempty" name:"RegionTypeName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 支付时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// 结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// 时间长度
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间单位名称
	TimeUnitName *string `json:"TimeUnitName,omitempty" name:"TimeUnitName"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 交易流水号
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 组件单价
	SinglePrice *string `json:"SinglePrice,omitempty" name:"SinglePrice"`

	// 组件价格单位
	SinglePriceUnit *string `json:"SinglePriceUnit,omitempty" name:"SinglePriceUnit"`

	// 组件用量
	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// 组件用量单位
	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`

	// 组件折前价
	Cost *string `json:"Cost,omitempty" name:"Cost"`

	// 组件折扣
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// 组件折扣价
	RealCost *string `json:"RealCost,omitempty" name:"RealCost"`

	// 组件应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 组件代金券支付金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 折扣
	TotalDiscount *string `json:"TotalDiscount,omitempty" name:"TotalDiscount"`
}

type DetailTotal struct {

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券花费
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
}

type EasyProduct struct {

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 商品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type ExportDealListRequest struct {
	*tchttp.BaseRequest

	// 订单号
	DealId *uint64 `json:"DealId,omitempty" name:"DealId"`

	// 大订单号
	BigDealId *uint64 `json:"BigDealId,omitempty" name:"BigDealId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 提单人
	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 订单类型
	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`

	// 状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 排序asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 大订单或小订单
	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`

	// 付费模式，0后付费/1预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 动作名称 隔离 销毁 返还 新购 续费 变配
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
}

func (r *ExportDealListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportDealListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportDealListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDealListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportDealListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateDealsRequest struct {
	*tchttp.BaseRequest

	// 商品信息
	Goods []*Goods `json:"Goods,omitempty" name:"Goods" list`
}

func (r *GenerateDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateDealsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateDealsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号列表
		OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds" list`

		// 大订单号
		BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateDealsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBillingProductCodeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBillingProductCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBillingProductCodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBillingProductCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品列表
		ProductList *EasyProduct `json:"ProductList,omitempty" name:"ProductList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBillingProductCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBillingProductCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDownloadListRecordItem struct {

	// 文件ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件类型名称
	FileTypeName *string `json:"FileTypeName,omitempty" name:"FileTypeName"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 生成时间
	GenerateTime *string `json:"GenerateTime,omitempty" name:"GenerateTime"`

	// 下载时间
	DownloadTime *string `json:"DownloadTime,omitempty" name:"DownloadTime"`

	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

type GetDownloadListSort struct {

	// 可选desc或者asc
	DownloadTime *string `json:"DownloadTime,omitempty" name:"DownloadTime"`

	// 可选desc或者asc
	GenerateTime *string `json:"GenerateTime,omitempty" name:"GenerateTime"`
}

type GetMineQuotaRequest struct {
	*tchttp.BaseRequest

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 商品配额key，由四层定义某几层以及"#"号拼接而成，如p_cvm#sp_cvm_vself2#v_cvm_cpu#sv_cvm_cpu_vself2
	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
}

func (r *GetMineQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMineQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMineQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMineQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMineQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMineResourceBillRequest struct {
	*tchttp.BaseRequest

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页大小(1~100)
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码(从0开始)
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`
}

func (r *GetMineResourceBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMineResourceBillRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMineResourceBillResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMineResourceBillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMineResourceBillResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPkgDownloadListItem struct {

	// 账单月份
	MonthName *string `json:"MonthName,omitempty" name:"MonthName"`

	// 当月开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 当月结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 状态(0未出账/1已出账但未生成账单/2已出账且账单生成中/4已出账且已生成账单)
	Status *string `json:"Status,omitempty" name:"Status"`

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券支付总价
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 下载链接，Status等于4该值才有效
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 当月周期类型
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// 出账日期
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`
}

type GetProductTreeRequest struct {
	*tchttp.BaseRequest

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子商品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费项
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费细项
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 页索引
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *GetProductTreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProductTreeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProductTreeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductTreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProductTreeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetQuotaLeftRequest struct {
	*tchttp.BaseRequest

	// 一层商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 商品配额key，由四层定义某几层以及"#"号拼接而成，如p_cvm#sp_cvm_vself2#v_cvm_cpu#sv_cvm_cpu_vself2
	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
}

func (r *GetQuotaLeftRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetQuotaLeftRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetQuotaLeftResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetQuotaLeftResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetQuotaLeftResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetQuotaRequest struct {
	*tchttp.BaseRequest

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 商品配额key，由四层定义某几层以及"#"号拼接而成，如p_cvm#sp_cvm_vself2#v_cvm_cpu#sv_cvm_cpu_vself2
	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`

	// uin
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`

	// ownerUin
	AccountOwnerUin *string `json:"AccountOwnerUin,omitempty" name:"AccountOwnerUin"`
}

func (r *GetQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账户额度详细
		List []*QuotaDetailList `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetResourceBillRequest struct {
	*tchttp.BaseRequest

	// 分页大小(1~100)
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码(从0开始)
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetResourceBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetResourceBillRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetResourceBillResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourceBillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetResourceBillResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Goods struct {

	// 区域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 新购时表示商品数量，续费与配置变更时固定传1
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 付费模式，0后付费/1预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 商品详情，timeSpan表示购买时间长度，timeUnit表示购买时间单位，pid表示定价公式ID
	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

	// 预付费操作方式，隔离isolate,销毁release,返还return,新购purchase,续费renew,变配降配downgrade，变配升配upgrade
	Action *string `json:"Action,omitempty" name:"Action"`

	// 订单类型ID，接入计费时由计费后台分配
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`

	// 可用区ID，没有可用区概念则传0
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 项目ID，没有项目概念则传0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 平台，0开放平台/1云平台
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
}

type GoodsDetail struct {

	// 计费项码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费细项码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 用量
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ListMineResourceRequest struct {
	*tchttp.BaseRequest

	// 页序号
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *ListMineResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListMineResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListMineResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMineResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListMineResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListResourceRequest struct {
	*tchttp.BaseRequest

	// 页序号，从0开始
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *ListResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MailingAddressDataList struct {

	// 用户uin
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// 邮寄地址ID
	PartaMailingAddressId *string `json:"PartaMailingAddressId,omitempty" name:"PartaMailingAddressId"`

	// 收件人
	AdresseeName *string `json:"AdresseeName,omitempty" name:"AdresseeName"`

	// 收件人电话
	AdresseePhone *string `json:"AdresseePhone,omitempty" name:"AdresseePhone"`

	// 详细地址
	AdresseeDress *string `json:"AdresseeDress,omitempty" name:"AdresseeDress"`

	// 是否为默认地址 1默认 0不默认
	DefaultState *uint64 `json:"DefaultState,omitempty" name:"DefaultState"`

	// 省份/直辖市
	AdresseeProvince *string `json:"AdresseeProvince,omitempty" name:"AdresseeProvince"`

	// 市/区
	AdresseeCity *string `json:"AdresseeCity,omitempty" name:"AdresseeCity"`
}

type ModifyAutoFlagRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 资源id列表，json格式见下.最多10个
	ResourceList []*ResourceList `json:"ResourceList,omitempty" name:"ResourceList" list`

	// 自动续费模式： 
	// 1：自动续费; 
	// 0:手动续费,到期后一定间隔时间后回收; 
	// 2:到期关闭
	AutopayType *uint64 `json:"AutopayType,omitempty" name:"AutopayType"`

	// 防重订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 付费模式,1:预付费。 目前只支持预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 类型.固定传: set_auto_flag
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyAutoFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContractStatusRequest struct {
	*tchttp.BaseRequest

	// 合同编号
	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`
}

func (r *ModifyContractStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContractStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContractStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContractStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContractStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDealStatusRequest struct {
	*tchttp.BaseRequest

	// 大订单号
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 修改动作  1取消订单 2删除订单
	ModifyStatus *uint64 `json:"ModifyStatus,omitempty" name:"ModifyStatus"`
}

func (r *ModifyDealStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDealStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDealStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDealStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDealStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDefaultMailingAddressRequest struct {
	*tchttp.BaseRequest

	// 邮寄地址ID
	PartaMailingAddressId *uint64 `json:"PartaMailingAddressId,omitempty" name:"PartaMailingAddressId"`

	// 是否为默认地址 1默认 0不默认
	DefaultState *uint64 `json:"DefaultState,omitempty" name:"DefaultState"`
}

func (r *ModifyDefaultMailingAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDefaultMailingAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDefaultMailingAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDefaultMailingAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDefaultMailingAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDefaultPostInfoRequest struct {
	*tchttp.BaseRequest

	// 地址的唯一识别号
	ArrId *int64 `json:"ArrId,omitempty" name:"ArrId"`
}

func (r *ModifyDefaultPostInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDefaultPostInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDefaultPostInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDefaultPostInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDefaultPostInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyOpenInvoiceInfoRequest struct {
	*tchttp.BaseRequest

	// 发票类型,取值为( 2:公司普通发票 3:公司增值税发票 )
	UserType *int64 `json:"UserType,omitempty" name:"UserType"`

	// 抬头
	InvoiceHead *string `json:"InvoiceHead,omitempty" name:"InvoiceHead"`

	// 税务登记号
	TaxNumber *string `json:"TaxNumber,omitempty" name:"TaxNumber"`

	// 开户行
	BankUserName *string `json:"BankUserName,omitempty" name:"BankUserName"`

	// 银行账号
	AccountNumber *string `json:"AccountNumber,omitempty" name:"AccountNumber"`

	// 注册全地址
	RegisterFullAddr *string `json:"RegisterFullAddr,omitempty" name:"RegisterFullAddr"`

	// 注册固定电话
	RegisterPhone *string `json:"RegisterPhone,omitempty" name:"RegisterPhone"`
}

func (r *ModifyOpenInvoiceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyOpenInvoiceInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyOpenInvoiceInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOpenInvoiceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyOpenInvoiceInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPartaMailingAddressRequest struct {
	*tchttp.BaseRequest

	// 邮寄地址ID
	PartaMailingAddressId *uint64 `json:"PartaMailingAddressId,omitempty" name:"PartaMailingAddressId"`

	// 收件人
	AdresseeName *string `json:"AdresseeName,omitempty" name:"AdresseeName"`

	// 收件人电话
	AdresseePhone *string `json:"AdresseePhone,omitempty" name:"AdresseePhone"`

	// 详细地址
	AdresseeDress *string `json:"AdresseeDress,omitempty" name:"AdresseeDress"`

	// 省份/直辖市
	AdresseeProvince *string `json:"AdresseeProvince,omitempty" name:"AdresseeProvince"`

	// 市/区
	AdresseeCity *string `json:"AdresseeCity,omitempty" name:"AdresseeCity"`

	// 备注
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *ModifyPartaMailingAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPartaMailingAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPartaMailingAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPartaMailingAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPartaMailingAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPostInfoRequest struct {
	*tchttp.BaseRequest

	// 地址的唯一识别号,新增则传:0,修改传服务器返回的id
	ArrId *int64 `json:"ArrId,omitempty" name:"ArrId"`

	// 联系人
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 省份
	Province *string `json:"Province,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitempty" name:"City"`

	// 详细地址
	Addr *string `json:"Addr,omitempty" name:"Addr"`

	// 手机号码 座机号和手机号必填一个
	Cellphone *string `json:"Cellphone,omitempty" name:"Cellphone"`
}

func (r *ModifyPostInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPostInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPostInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地址的唯一识别号
		ArrId *int64 `json:"ArrId,omitempty" name:"ArrId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPostInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPostInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyReconciliationRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 调账原因：计费错误-billing_error 产品故障-business_error 内部核销-internal_write_off 其它原因-other
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 调账类型 add-补偿/minus-扣费
	AdjustType *string `json:"AdjustType,omitempty" name:"AdjustType"`

	// 调账月份： "2019-05"
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// 调账金额：元
	Amount *string `json:"Amount,omitempty" name:"Amount"`

	// 调账订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子商品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费项
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费细项
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
}

func (r *ModifyReconciliationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyReconciliationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyReconciliationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReconciliationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyReconciliationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonthCostDetailItem struct {

	// 月份中文名，如2018年12月
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 应收花费
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 月份，如2018-12
	Month *string `json:"Month,omitempty" name:"Month"`
}

type OrderInfo struct {

	// 订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 米大师订单号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 用户id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 支付渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// ISO 货币代码，CNY
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 支付金额，单位：分。
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`

	// 当前订单的订单状态  0： 初始状态，获取Midas交易订单成功；   1： 拉起Midas支付页面成功，用户未支付；   2：用户支付成功，正在发货；   3：用户支付成功，发货失败；   4：用户支付成功，发货成功；   5：Midas支付页面正在失效中；   6：Midas支付页面已经失效；
	OrderState *string `json:"OrderState,omitempty" name:"OrderState"`

	// utc时间戳.支付时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 提现标记.0-可提现, 1-已提现
	RefundFlag *string `json:"RefundFlag,omitempty" name:"RefundFlag"`
}

type Paras struct {

	// 别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 别名对应的值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 三层定义
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 四层定义
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
}

type PayDealsRequest struct {
	*tchttp.BaseRequest

	// 支付订单号，可传多个，用引文“,”号隔开
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 使用抵扣券列表
	CouponsList []*CouponsList `json:"CouponsList,omitempty" name:"CouponsList" list`
}

func (r *PayDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PayDealsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PayDealsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PayDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PayDealsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PayModeSummaryDetailItem struct {

	// 计费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券花费
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 对比上月趋势
	Trend []*SummaryTrend `json:"Trend,omitempty" name:"Trend" list`
}

type PayModeSummaryOverviewItem struct {

	// 付费模式代码
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// 操作详情
	Detail *ActionDetailItem `json:"Detail,omitempty" name:"Detail"`
}

type PrepayDealRequest struct {
	*tchttp.BaseRequest

	// 支付订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 用户AppId
	UserAppId *int64 `json:"UserAppId,omitempty" name:"UserAppId"`
}

func (r *PrepayDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PrepayDealRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PrepayDealResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PrepayDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PrepayDealResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PriceListData struct {

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子商品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 单价
	Price *uint64 `json:"Price,omitempty" name:"Price"`

	// 实际支付价格
	RealTotalCost *uint64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 原价
	TotalCost *uint64 `json:"TotalCost,omitempty" name:"TotalCost"`

	// 折扣
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间大小
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 数量
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 各组件价格详情
	PartDetail *string `json:"PartDetail,omitempty" name:"PartDetail"`
}

type PriceRange struct {

	// 计费项的本区间单价 = 配置中的单价，单位：元后面8位
	UnitPrice *string `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 计费项的本区间的原价，单位：元后面8位。= unit_price * value
	OriginalAmount *string `json:"OriginalAmount,omitempty" name:"OriginalAmount"`

	// 本区间用量
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PriceRes struct {

	// 单价 * 量，单位：元后面8位
	UnitPrice *string `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 原价价格。一期 = amout
	OriginalAmount *string `json:"OriginalAmount,omitempty" name:"OriginalAmount"`

	// 实际价格，单位：元后面8位
	Amout *string `json:"Amout,omitempty" name:"Amout"`

	// 定价类型，当前固定为：linea
	PriceModel *string `json:"PriceModel,omitempty" name:"PriceModel"`

	// PriceRange
	PriceRange []*PriceRange `json:"PriceRange,omitempty" name:"PriceRange" list`
}

type Product struct {

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品中文名称
	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`

	// 产品英文名称
	EngName *string `json:"EngName,omitempty" name:"EngName"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ProductDefineList struct {

	// 产品代码
	ProductCode []*string `json:"ProductCode,omitempty" name:"ProductCode" list`

	// 产品中文名称
	ProductChnName *string `json:"ProductChnName,omitempty" name:"ProductChnName"`

	// 产品英文名称
	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品中文名称
	SubProductChnName *string `json:"SubProductChnName,omitempty" name:"SubProductChnName"`

	// 子产品英文名称
	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`

	// 计费项代码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费项中文名称
	BillingItemChnName *string `json:"BillingItemChnName,omitempty" name:"BillingItemChnName"`

	// 计费项英文名称
	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`

	// 计费细项代码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 计费细项中文名称
	SubBillingItemChnName *string `json:"SubBillingItemChnName,omitempty" name:"SubBillingItemChnName"`

	// 计费细项英文名称
	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`

	// 计费细项单位
	SubBillingItemsGoodNumUnit *string `json:"SubBillingItemsGoodNumUnit,omitempty" name:"SubBillingItemsGoodNumUnit"`

	// 计费细项单位英文
	SubBillingItemsGoodNumUnitEn *string `json:"SubBillingItemsGoodNumUnitEn,omitempty" name:"SubBillingItemsGoodNumUnitEn"`

	// 四层定义状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 最近操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ProductInfo struct {

	// GoodsDetail
	GoodsDetail *GoodsDetail `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

	// PriceRes
	PriceRes *PriceRes `json:"PriceRes,omitempty" name:"PriceRes"`
}

type ProductList struct {

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 付费模式,0:后付费,1:预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 别名信息
	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 数量
	Num *string `json:"Num,omitempty" name:"Num"`

	// 地域
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 区域
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 时长单位。1:年,2:月,3:日,4:小时,5:分钟,6:秒
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时长
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

type ProductPropertyFields struct {

	// 字段名
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 字段码
	FieldCode *string `json:"FieldCode,omitempty" name:"FieldCode"`

	// 字段标志： 0-普通 1-主标记
	FieldFlag *string `json:"FieldFlag,omitempty" name:"FieldFlag"`
}

type ProductPropertyList struct {

	// key中文名
	KeyNameCn *string `json:"KeyNameCn,omitempty" name:"KeyNameCn"`

	// key英文名
	KeyNameEn *string `json:"KeyNameEn,omitempty" name:"KeyNameEn"`

	// key码，后端定义标识规则
	KeyCode *string `json:"KeyCode,omitempty" name:"KeyCode"`

	// key的字段列表
	Fields *ProductPropertyFields `json:"Fields,omitempty" name:"Fields"`

	// JSON String 属性值列表  jsond ecode后是数组，数组元素的key不定  例如：[{
	Values *string `json:"Values,omitempty" name:"Values"`

	// 是否可被编辑。0是不可以；1是可以
	Modifiable *int64 `json:"Modifiable,omitempty" name:"Modifiable"`

	// 自定义属性关联的四层定义。　 
	// 字符串数组，例如： ["p_cvm##v_cvm_mem#sv_cvm_mem_s1","p_cvm##v_cvm_mem#sv_cvm_mem_s2"]
	CustomPropertyRelateProduct []*string `json:"CustomPropertyRelateProduct,omitempty" name:"CustomPropertyRelateProduct" list`

	// 自定义属性关联的模块。　 
	// 字符串数组，例如： ["price"]
	CustomPropertyRelateModels []*string `json:"CustomPropertyRelateModels,omitempty" name:"CustomPropertyRelateModels" list`
}

type QueryPublicAccountRequest struct {
	*tchttp.BaseRequest

	// 公共帐户ID默认写死1
	AccountId *uint64 `json:"AccountId,omitempty" name:"AccountId"`
}

func (r *QueryPublicAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryPublicAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryPublicAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 和入参ID保持一致（默认只有1）
		AccountId []*uint64 `json:"AccountId,omitempty" name:"AccountId" list`

		// 收款人户名
		AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

		// 收款卡号
		AccountNum []*uint64 `json:"AccountNum,omitempty" name:"AccountNum" list`

		// 银行
		Bank *string `json:"Bank,omitempty" name:"Bank"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPublicAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryPublicAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryQuotaRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 游标，从第0开始取
	CurPos *uint64 `json:"CurPos,omitempty" name:"CurPos"`

	// 总共取多少条
	MaxNum *uint64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

func (r *QueryQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额信息
		QuotaInfo []*QuotaInfo `json:"QuotaInfo,omitempty" name:"QuotaInfo" list`

		// 记录总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuotaDetailList struct {

	// 账户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 产品码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品配额标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`

	// 产品配额值
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 子产品码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 子计费码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 产品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品类名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`

	// 子产品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`

	// 计费名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`

	// 子计费名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`

	// 单位名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type QuotaInfo struct {

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费项码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费细项码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 当前使用配额
	Quota *uint64 `json:"Quota,omitempty" name:"Quota"`
}

type QuotaList struct {

	// 用户uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品配额值
	ProductQuotaValue *string `json:"ProductQuotaValue,omitempty" name:"ProductQuotaValue"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品名称
	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`

	// 子产品配额值
	SubProductQuotaValue *string `json:"SubProductQuotaValue,omitempty" name:"SubProductQuotaValue"`

	// 计费项代码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费项名称
	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`

	// 计费项配额值
	BillingItemQuotaValue *string `json:"BillingItemQuotaValue,omitempty" name:"BillingItemQuotaValue"`

	// 计费细项代码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 计费细项名称
	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`

	// 计费细项配额值
	SubBillingItemQuotaValue *string `json:"SubBillingItemQuotaValue,omitempty" name:"SubBillingItemQuotaValue"`

	// 产品英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`

	// 子产品英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`

	// 产品项英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`

	// 产品细项英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
}

type RefundDetail struct {

	// 提现id
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 外部订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 用户id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 提现金额,单位：分
	RefundAmt *string `json:"RefundAmt,omitempty" name:"RefundAmt"`

	// 原订单支付金额,单位分
	OrderPayAmt *string `json:"OrderPayAmt,omitempty" name:"OrderPayAmt"`

	// 支付渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 提现时间
	RefundTime *string `json:"RefundTime,omitempty" name:"RefundTime"`

	// 提现状态.1-提现中;2-提现成功;3-提现失败
	State *string `json:"State,omitempty" name:"State"`
}

type RefundRequest struct {
	*tchttp.BaseRequest

	// 米大师的应用ID
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 退款订单号
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 退款金额
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`

	// 充值订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
}

func (r *RefundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefundRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RefundResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefundResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefundResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionDetail struct {

	// 数量
	TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

	// 地域列表
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList" list`
}

type RegionSummaryDetailItem struct {

	// 该项目下产品消费明细
	Business *BusinessSummaryDetailItem `json:"Business,omitempty" name:"Business"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券花费
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 对比上月趋势
	Trend *SummaryTrend `json:"Trend,omitempty" name:"Trend"`
}

type RegionSummaryOverviewItem struct {

	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type RemitRecordData struct {

	// 用户uin
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 汇款银行名称
	Bank *string `json:"Bank,omitempty" name:"Bank"`

	// 汇款银行账户
	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`

	// 汇款户名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 汇款时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 联系人手机
	Tel *uint64 `json:"Tel,omitempty" name:"Tel"`

	// 状态 0:处理中 1:成功 2:失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type RenewDetail struct {

	// 数量
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// 续费列表
	RenewList []*RenewList `json:"RenewList,omitempty" name:"RenewList" list`
}

type RenewList struct {

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 状态。0：正常；1：隔离；2：销毁;不传查询全部
	Status *string `json:"Status,omitempty" name:"Status"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 地域id
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 区域id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 时长
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 时长单位:1:年,2:月,3:日,4:小时,5:分钟,6:秒
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 付费模式,0:后付费,1:预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 用户在业务侧的承载号码
	AppIdUsernum *string `json:"AppIdUsernum,omitempty" name:"AppIdUsernum"`

	// 三、四层定义
	Paras []*Paras `json:"Paras,omitempty" name:"Paras" list`

	// 自动续费模式： 0:手动续费,到期后一定间隔时间后回收;1：自动续费; 2:到期关闭
	AutopayType *string `json:"AutopayType,omitempty" name:"AutopayType"`

	// 预计隔离时间,格式为datetime
	ExpectIsolateTime *string `json:"ExpectIsolateTime,omitempty" name:"ExpectIsolateTime"`
}

type ResourceInfo struct {

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 别名信息
	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 地域id
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 区域id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type ResourceList struct {

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 地域
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type ResourceSummaryDetailComponentDetailItem struct {

	// 组件用量
	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// 组件码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 组件名称
	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`

	// 组件折后总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 组件折后总价占资源折后总价比率
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type ResourceSummaryDetailRealTotalCostTrend struct {

	// 组件历史月份花费详情
	Detail *ResourceSummaryDetailRealTotalCostTrendDetailItem `json:"Detail,omitempty" name:"Detail"`

	// 历史月份平均值
	Average *string `json:"Average,omitempty" name:"Average"`
}

type ResourceSummaryDetailRealTotalCostTrendDetailItem struct {

	// 月份，如2018-08
	Month *string `json:"Month,omitempty" name:"Month"`

	// 资源折后总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
}

type ResourceSummaryDetailTotal struct {

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券支付总价
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
}

type SetAutoFlagRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 资源id列表，json格式见下.最多10个
	ResourceList []*ResourceList `json:"ResourceList,omitempty" name:"ResourceList" list`

	// 自动续费模式： 
	// 1：自动续费; 
	// 0:手动续费,到期后一定间隔时间后回收; 
	// 2:到期关闭
	AutopayType *uint64 `json:"AutopayType,omitempty" name:"AutopayType"`

	// 防重订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 付费模式,1:预付费。 目前只支持预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 类型.固定传: set_auto_flag
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *SetAutoFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetAutoFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetSubUinQuotaRequest struct {
	*tchttp.BaseRequest

	// 子账户配额信息列表
	SubUinQuotaDataArray []*SubUinQuotaDataArray `json:"SubUinQuotaDataArray,omitempty" name:"SubUinQuotaDataArray" list`

	// 计费MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 商品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子商品代码（可为空）
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费项代码（可为空）
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费细项代码 （可为空）
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 当前父账户Uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *SetSubUinQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetSubUinQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetSubUinQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSubUinQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetSubUinQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubOutTradeNoAmount struct {

	// 子订单
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子订单单价，单位：分
	UnitPrice *string `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 子订单原价价格，单位：分。一期 = amout
	OriginalAmount *string `json:"OriginalAmount,omitempty" name:"OriginalAmount"`

	// 子订单实际价格，单位：分
	Amount *string `json:"Amount,omitempty" name:"Amount"`

	// 折扣
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// 三四层所有unit_price总合，单位为微元
	ActualPrice *string `json:"ActualPrice,omitempty" name:"ActualPrice"`

	// 三四层所有unit_price总合*数量*时长,单位微元，一期=actual_amount
	ActualOriAmount *string `json:"ActualOriAmount,omitempty" name:"ActualOriAmount"`

	// 三四层所有unit_price总合*数量*时长，单位为微元
	ActualAmount *string `json:"ActualAmount,omitempty" name:"ActualAmount"`

	// ProductInfo
	ProductInfo []*ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo" list`
}

type SubOutTradeNoList struct {

	// 子订单
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 付费模式,0:后付费,1:预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 1 : 同步发货.0:异步发货
	SyncProvide *int64 `json:"SyncProvide,omitempty" name:"SyncProvide"`

	// 别名信息
	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 数量
	Num *string `json:"Num,omitempty" name:"Num"`

	// 地域
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 区域
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 发货透传到云计费
	ProvideInfo *string `json:"ProvideInfo,omitempty" name:"ProvideInfo"`

	// 时长单位。1:年,2:月,3:日,4:小时,5:分钟,6:秒
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时长
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 自动续费类型. 0:手动续费,到期后一定间隔时间后回收;1：自动续费; 2:到期关闭
	AutopayType *string `json:"AutopayType,omitempty" name:"AutopayType"`
}

type SubUinQuotaArray struct {

	// 子账户
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 子账户设置的配额值
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`

	// 子账户已使用的配额值
	UsedQuota *int64 `json:"UsedQuota,omitempty" name:"UsedQuota"`
}

type SubUinQuotaDataArray struct {

	// 子账户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 要设置的子账户配额
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`
}

type SubscribeManualRenewRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 操作人
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户AppId,跟UserId对应
	UserAppId *string `json:"UserAppId,omitempty" name:"UserAppId"`

	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合。
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 子订单列表
	SubOutTradeNoList []*SubOutTradeNoList `json:"SubOutTradeNoList,omitempty" name:"SubOutTradeNoList" list`

	// 续费类型.固定传 manual 表示手工续费
	RenewType *string `json:"RenewType,omitempty" name:"RenewType"`

	// 订单有效时间.单位:秒
	OrderValidTime *uint64 `json:"OrderValidTime,omitempty" name:"OrderValidTime"`

	// 备注信息.透传到业务流水
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *SubscribeManualRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribeManualRenewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribeManualRenewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子单列表金额详情.
		SubOutTradeNoAmount []*SubOutTradeNoAmount `json:"SubOutTradeNoAmount,omitempty" name:"SubOutTradeNoAmount" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribeManualRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribeManualRenewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribeOrderRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 操作人
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 户AppId,跟UserId对应
	UserAppId *string `json:"UserAppId,omitempty" name:"UserAppId"`

	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合。
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 子单信息
	SubOutTradeNoList []*SubOutTradeNoList `json:"SubOutTradeNoList,omitempty" name:"SubOutTradeNoList" list`

	// 订单有效时间.单位:秒
	OrderValidTime *uint64 `json:"OrderValidTime,omitempty" name:"OrderValidTime"`

	// 备注信息.透传到业务流水
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *SubscribeOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribeOrderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribeOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子单批价信息.
		SubOutTradeNoAmount []*SubOutTradeNoAmount `json:"SubOutTradeNoAmount,omitempty" name:"SubOutTradeNoAmount" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribeOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribeOrderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribePayRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合。
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 子账户Id.固定传 CREDIT_FIXED
	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`

	// 使用抵扣券列表
	CouponsList []*CouponsList `json:"CouponsList,omitempty" name:"CouponsList" list`
}

func (r *SubscribePayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribePayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribePayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 支付相关信息
		ResourceList *ResourceList `json:"ResourceList,omitempty" name:"ResourceList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribePayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribePayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribeUpgradeRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 操作人
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户AppId,跟UserId对应
	UserAppId *string `json:"UserAppId,omitempty" name:"UserAppId"`

	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 老配置信息
	OldResourceInfo *ResourceInfo `json:"OldResourceInfo,omitempty" name:"OldResourceInfo"`

	// 新配置信息
	NewResourceInfo *ResourceInfo `json:"NewResourceInfo,omitempty" name:"NewResourceInfo"`

	// 内容为json //发货透传到云计费
	ProvideInfo *string `json:"ProvideInfo,omitempty" name:"ProvideInfo"`

	// 时长单位:1:年,2:月,3:日,4:小时,5:分钟,6:秒
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 付费类型，0:后付费,1:预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 取值:1 : 同步发货.0:异步发货
	SyncProvide *int64 `json:"SyncProvide,omitempty" name:"SyncProvide"`

	// 订单有效时间.单位:秒
	OrderValidTime *uint64 `json:"OrderValidTime,omitempty" name:"OrderValidTime"`

	// 0:手动续费,到期后一定间隔时间后回收;1：自动续费; 2:到期关闭
	AutopayType *uint64 `json:"AutopayType,omitempty" name:"AutopayType"`

	// 备注信息.透传到业务流水.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *SubscribeUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribeUpgradeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribeUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 升配批价信息
		UpgradePriceInfo *UpgradePriceInfo `json:"UpgradePriceInfo,omitempty" name:"UpgradePriceInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribeUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribeUpgradeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SummaryByResourceConditionValue struct {

	// 产品列表
	Product *ConditionBusiness `json:"Product,omitempty" name:"Product"`

	// 子产品列表
	SubProduct *ConditionProduct `json:"SubProduct,omitempty" name:"SubProduct"`

	// 地域列表
	Region *ConditionRegion `json:"Region,omitempty" name:"Region"`

	// 付费模式列表
	PayMode *ConditionPayMode `json:"PayMode,omitempty" name:"PayMode"`

	// 交易类型列表
	ActionType *ConditionActionType `json:"ActionType,omitempty" name:"ActionType"`
}

type SummaryByResourceDataItem struct {

	// 记录ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 支付者uin
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 子产品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品名称
	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 现金支付总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券支付总价
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 付费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 流水号
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 资源所有者uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 操作者uin
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 交易类型名称
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// 可用区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 支付时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// 结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// 组件配置
	ComponentConfig *string `json:"ComponentConfig,omitempty" name:"ComponentConfig"`
}

type SummaryByResourceTotal struct {

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券支付总价
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
}

type SummaryTrend struct {

	// upward上升/downward下降/none无
	Type *string `json:"Type,omitempty" name:"Type"`

	// 趋势白分值，保留两位小数
	Value *string `json:"Value,omitempty" name:"Value"`

	// 上月消耗
	PreMonthRealTotalCost *string `json:"PreMonthRealTotalCost,omitempty" name:"PreMonthRealTotalCost"`
}

type SyncContractPartaInfoRequest struct {
	*tchttp.BaseRequest

	// 甲方名称
	PartAName *string `json:"PartAName,omitempty" name:"PartAName"`

	// 甲方地址
	PartAAddress *string `json:"PartAAddress,omitempty" name:"PartAAddress"`

	// 甲方联系人
	ConnectUser *string `json:"ConnectUser,omitempty" name:"ConnectUser"`

	// 甲方电话
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 甲方邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 备注
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *SyncContractPartaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncContractPartaInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncContractPartaInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncContractPartaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncContractPartaInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TrendByMonthDetail struct {

	// 当前月份统计
	Current []*MonthCostDetailItem `json:"Current,omitempty" name:"Current" list`

	// 历史月份统计
	Previous []*MonthCostDetailItem `json:"Previous,omitempty" name:"Previous" list`

	// 未来月份统计
	Next []*MonthCostDetailItem `json:"Next,omitempty" name:"Next" list`
}

type TrendByMonthStat struct {

	// 平均花费详情
	Average []*TrendByMonthStatItem `json:"Average,omitempty" name:"Average" list`
}

type TrendByMonthStatItem struct {

	// 开始月份，如2018-05
	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`

	// 结束月份，如2018-11
	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`

	// 应付金额
	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`

	// 代金券花费
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
}

type UpgradePriceInfo struct {

	// 升配申请时间，格式为：YY-MM-DD HH::MM::SS
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 资源到期时间，格式为：YY-MM-DD HH::MM::SS
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时间单位，目前只能为月
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 升配时长
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 升配天数
	UpgradeDays *string `json:"UpgradeDays,omitempty" name:"UpgradeDays"`

	// 低配置按月价格，单位为: 元*10^8
	LowMonthPrice *string `json:"LowMonthPrice,omitempty" name:"LowMonthPrice"`

	// 高配置按月价格，单位为: 元*10^8
	HighMonthPrice *string `json:"HighMonthPrice,omitempty" name:"HighMonthPrice"`

	// 升配支付金额，单位为: 元*10^8
	UpgradePrice *string `json:"UpgradePrice,omitempty" name:"UpgradePrice"`

	// 目前没有折扣, 默认为：100
	Discount *string `json:"Discount,omitempty" name:"Discount"`
}

type UpgradePriceRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户AppId,跟UserId对应
	UserAppId *string `json:"UserAppId,omitempty" name:"UserAppId"`

	// 付费类型，0:后付费,1:预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 时长单位:1:年,2:月,3:日,4:小时,5:分钟,6:秒
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 老配置信息
	OldResourceInfo *ResourceInfo `json:"OldResourceInfo,omitempty" name:"OldResourceInfo"`

	// 新配置信息
	NewResourceInfo *ResourceInfo `json:"NewResourceInfo,omitempty" name:"NewResourceInfo"`

	// 固定传:upgrade_price
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *UpgradePriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradePriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradePriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 升配批价结果信息
		UpgradePriceInfo *UpgradePriceInfo `json:"UpgradePriceInfo,omitempty" name:"UpgradePriceInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradePriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradePriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeProductInfo struct {

	// 时长单位:1:年,2:月,3:日,4:小时,5:分钟,6:秒
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 付费类型，0:后付费,1:预付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 老配置信息
	OldResourceInfo *ResourceInfo `json:"OldResourceInfo,omitempty" name:"OldResourceInfo"`

	// 新配置信息
	NewResourceInfo *ResourceInfo `json:"NewResourceInfo,omitempty" name:"NewResourceInfo"`
}

type VoucherConditions struct {

	// 导出字段
	ExportFields []*string `json:"ExportFields,omitempty" name:"ExportFields" list`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 指定查询状态
	SpStatus *string `json:"SpStatus,omitempty" name:"SpStatus"`

	// 指定产品
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 指定代金券编号
	SpId *string `json:"SpId,omitempty" name:"SpId"`

	// 指定排序字段：begin_time、end_time。
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// 指定升序降序：desc、asc。
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

type WaterListData struct {

	// 流水类型
	WaterType *string `json:"WaterType,omitempty" name:"WaterType"`

	// 存取标志，1：支出；2：收入
	IoFlag *string `json:"IoFlag,omitempty" name:"IoFlag"`

	// 给外部提供的接口所对应的交易类型：  2-支付  10-转账  19-还款 20-销帐 21-预授权 22-预授权取消（恢复） 23-预授权确认（消耗）24-退款 25-产生账单 26-修改账单 27-临时账户新增扩展额度 28-撤销还款 29-预授权确认回退 30-预授权改单 31-预授权回退
	ExternTranType *string `json:"ExternTranType,omitempty" name:"ExternTranType"`

	// 对应到内部server接口类型： 2-支付  10-开户（开户接口触发） 14-转账扣款   16-转账充值  17-转账取消   42-注销   52-后付费上调额度   53-后付费下调额度   54-后付费更改账户信息（除调额外） 55-还款 56-销账 57-提前还款   58-抵消账单   59-预授权   60-预授权取消（恢复） 61-预授权工具取消（恢复） 62-预授权确认（消耗） 63-预授权订单方面 64-预授权取消订单方面 65-退款（给账户充） 66-应恢复额度流水 67-补扣/补量68-增加账单金额（出账工具）69-减少账单金额（出账工具）70-还原cash金额（出账工具）71-修改账单导致余额增加（出账工具）72-修改账单导致余额减少（出账工具）73-新增临时账户扩展额度   74-撤销还款   75-撤销销帐   76-撤销提前还款   82-预授权改单上调，账户减少   83-预授权改单下调，账户增加   84-预授权改单上调，预授权单增加   85-预授权改单下调，预授权单减少   86-预授权退款   87-预授权补扣
	BaseTranType *string `json:"BaseTranType,omitempty" name:"BaseTranType"`

	// 订单号
	BillNo *string `json:"BillNo,omitempty" name:"BillNo"`

	// 交易金额
	TranAmt *string `json:"TranAmt,omitempty" name:"TranAmt"`

	// 余额，后付费账户balance根据账户属性决定等于dpbalance还是dpbalance+cash
	Balance *string `json:"Balance,omitempty" name:"Balance"`

	// 流水产生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 渠道来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 客户端设备，0：PC；1：手机
	Device *string `json:"Device,omitempty" name:"Device"`

	// 用户IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 交易描述
	TranInfo *string `json:"TranInfo,omitempty" name:"TranInfo"`

	// 交易备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 保留字段1
	Reserve1 *string `json:"Reserve1,omitempty" name:"Reserve1"`

	// 保留字段2
	Reserve2 *string `json:"Reserve2,omitempty" name:"Reserve2"`

	// 由外部输入的扩展性保留字段1
	Extreserve1 *string `json:"Extreserve1,omitempty" name:"Extreserve1"`

	// 由外部输入的扩展性保留字段2
	Extreserve2 *string `json:"Extreserve2,omitempty" name:"Extreserve2"`

	// 关联id，如销账的账单号
	RelationId *string `json:"RelationId,omitempty" name:"RelationId"`

	// 后付费额度上限
	Dplimit *string `json:"Dplimit,omitempty" name:"Dplimit"`

	// 后付费剩余可用额度（根据账户属性决定是否包含cash）
	Dpbalance *string `json:"Dpbalance,omitempty" name:"Dpbalance"`

	// 后付费账户，固定账户存储现金余额
	Cash *string `json:"Cash,omitempty" name:"Cash"`

	// 商户类型，1：代理，2：子客，3：直客
	MerchantType *string `json:"MerchantType,omitempty" name:"MerchantType"`

	// 关联商户信息，当商户类型为代理或直客时，该字段为空；当为子客时，标识代理账户信息
	RelatedMerchant *string `json:"RelatedMerchant,omitempty" name:"RelatedMerchant"`

	// 保留字段3 （记录流水插入DB时间）
	Reserve3 *string `json:"Reserve3,omitempty" name:"Reserve3"`

	// 保留字段4（1. 更改账户的标记；2.还款流水记录销账的账单号；3.出账时，完全抵消账单金额标记AllOffSet）
	Reserve4 *string `json:"Reserve4,omitempty" name:"Reserve4"`

	// 保留字段5（交易时间微秒us）
	Reserve5 *string `json:"Reserve5,omitempty" name:"Reserve5"`

	// 保留字段6 （销帐记录：TotalRefund,1000;CurDebt,1110;DebtBillDate,201706; PayChannel,bank;）
	Reserve6 *string `json:"Reserve6,omitempty" name:"Reserve6"`

	// 保留字段7（销帐记录账单恢复金额：  Type,还款类型;Recover,恢复额度;ChgTimeUs,HOLD中的交易微妙; TranSeq,保存在账户中的交易序列号;）
	Reserve7 *string `json:"Reserve7,omitempty" name:"Reserve7"`
}
