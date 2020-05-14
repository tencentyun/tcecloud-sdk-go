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

type AccountListData struct {

	// 用户uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 欠款金额
	DebtAmount *string `json:"DebtAmount,omitempty" name:"DebtAmount"`

	// 帐户余额
	LeftAmount *string `json:"LeftAmount,omitempty" name:"LeftAmount"`
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

type AddQuotaPara struct {

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 配额key
	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`

	// 配额值
	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`

	// 用户uin。此参数不传时设置默认配额。
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type AddQuotaRequest struct {
	*tchttp.BaseRequest

	// 配额复杂类型
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

type BasicAccount struct {

	// 用户uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 用户appId
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 账户是否合法，1:合法，0：非法
	IsExist *int64 `json:"IsExist,omitempty" name:"IsExist"`
}

type BigDealListData struct {

	// 大订单号
	BigDealId *int64 `json:"BigDealId,omitempty" name:"BigDealId"`

	// 用户唯一标识
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

	// 创建人uin
	Payer *string `json:"Payer,omitempty" name:"Payer"`

	// 状态 1待支付 2已支付 3发货中 4发货成功 5发货失败 6已退款 7取消 100删除
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 金额总计
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 代金券金额
	VoucherDecline *string `json:"VoucherDecline,omitempty" name:"VoucherDecline"`

	// 实付金额总计
	RealTotalCost *bool `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 付款人
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 状态名称
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// 小订单列表
	DealList *DealListData `json:"DealList,omitempty" name:"DealList"`

	// 折扣金额（因折扣而减少的金额）
	DiscountCost *string `json:"DiscountCost,omitempty" name:"DiscountCost"`
}

type BillingItem struct {

	// 计费项代码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费项中文名
	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`

	// 计费项英文名
	EngName *string `json:"EngName,omitempty" name:"EngName"`

	// 计费细项状
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type BspResourceBill struct {

	// 用户uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 用户appId
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 计费单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 计费用时
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 数量
	UsedAmount *int64 `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// 单位
	UsedAmountUint *string `json:"UsedAmountUint,omitempty" name:"UsedAmountUint"`

	// 产品大类
	ProductClass *string `json:"ProductClass,omitempty" name:"ProductClass"`

	// 产品编码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 子产品编码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品名称
	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`

	// 地域id
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 计费项编码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费项名称
	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`

	// 子计费项编码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 子计费项名称
	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
}

type BusinessSummaryDetailItem struct {

	// 产品码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

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
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type BusinessSummaryTotal struct {

	// 总花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
}

type ConditionActionType struct {

	// 交易类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 交易类型名称
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`
}

type ConditionBusiness struct {

	// 产品码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`
}

type ConditionComponent struct {

	// 组件码
	ComponentCode *string `json:"ComponentCode,omitempty" name:"ComponentCode"`

	// 组件名称
	ComponentCodeName *string `json:"ComponentCodeName,omitempty" name:"ComponentCodeName"`
}

type ConditionPayMode struct {

	// 付费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
}

type ConditionProduct struct {

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
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

	// 导出字段
	ExportFields []*string `json:"ExportFields,omitempty" name:"ExportFields" list`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
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
	GoodsPrice *string `json:"GoodsPrice,omitempty" name:"GoodsPrice"`
}

type DebateListData struct {

	// 用户uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 欠费月份数量
	DebtNum *string `json:"DebtNum,omitempty" name:"DebtNum"`
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

	// Kv拼接的扩展字段，采用逗号和分号分割，示例如下：
	// Campid,xxxxx;Subclientid,xxxx;
	Reserve2 *string `json:"Reserve2,omitempty" name:"Reserve2"`
}

type DescribeAccountListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 管理端查询用户UIN，默认缺省查全部
	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *DescribeAccountListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 帐户信息列表
		List []*AccountListData `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountWaterRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识
	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`

	// 用户ID
	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`

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

		// 流水结果列表
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

type DescribeBasicAccountListRequest struct {
	*tchttp.BaseRequest

	// 用户uin数组
	Uins []*string `json:"Uins,omitempty" name:"Uins" list`
}

func (r *DescribeBasicAccountListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBasicAccountListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAccountListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基本账户信息List
		List []*BasicAccount `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicAccountListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBasicAccountListResponse) FromJsonString(s string) error {
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
		ComponentDetail *ResourceSummaryDetailComponentDetailItem `json:"ComponentDetail,omitempty" name:"ComponentDetail"`

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
	Conditions []*Conditions `json:"Conditions,omitempty" name:"Conditions" list`
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

type DescribeBillingItemsRequest struct {
	*tchttp.BaseRequest

	// midas appid
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 计费项code列表
	BillingItemCodes []*string `json:"BillingItemCodes,omitempty" name:"BillingItemCodes" list`
}

func (r *DescribeBillingItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillingItemsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingItemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计费项信息列表
		List *ProductBillingItemInfoList `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillingItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillingItemsResponse) FromJsonString(s string) error {
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

type DescribeDealListRequest struct {
	*tchttp.BaseRequest

	// 操作uin，查全量缺省
	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`

	// 大订单或小订单
	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`

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

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 付费模式，0后付费/1预付费
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 动作名称 隔离 销毁 返还 新购 续费 升配 降配
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 订单创建人uin
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

		// 大订单列表
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

type DescribeDebateListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 欠费月份，0就是查欠账月份大于等于0
	UnpayDebtNum *uint64 `json:"UnpayDebtNum,omitempty" name:"UnpayDebtNum"`

	// 针对DebtNum的比较条件， E: 等于， G: 大于，GE: 大于等于
	NumSort *string `json:"NumSort,omitempty" name:"NumSort"`

	// 管理端查询用户UIN，默认缺省查全部
	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *DescribeDebateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDebateListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDebateListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 账单欠费月份信息
		List *DebateListData `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDebateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDebateListResponse) FromJsonString(s string) error {
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

type DescribeLifecycleFlowListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// 要查询的UIN
	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`

	// 操作者UIN
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 操作
	OperateAction *string `json:"OperateAction,omitempty" name:"OperateAction"`
}

func (r *DescribeLifecycleFlowListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLifecycleFlowListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleFlowListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 流水列表
		List *int64 `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLifecycleFlowListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLifecycleFlowListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayOrderRequest struct {
	*tchttp.BaseRequest

	// 米大师appid，由计平分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// type=by_order，根据订单号查订单
	// type=by_user，根据用户id查订单
	Type *string `json:"Type,omitempty" name:"Type"`

	// 用户id
	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`

	// 查询的起始时间，unix时间戳,精确到秒。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的结束时间，unix时间戳,精确到秒。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 记录数偏移量，默认从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 支付订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 支付渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 提现标记.0-可提现, 1-已提现
	RefundFlag *string `json:"RefundFlag,omitempty" name:"RefundFlag"`

	// 订单状态，用于过滤。本场景固定传["2","3","4"]
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
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

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

type DescribeProductPricesRequest struct {
	*tchttp.BaseRequest

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeProductPricesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductPricesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductPricesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 价格策略列表
		List *ProductPriceList `json:"List,omitempty" name:"List"`

		// JSON string；是自定义属性keyCode的数组，包括了中英文名称  
	// 例如：{"system":{"chnName":"操作系统","engName":"systemEnglishName"},"brand":{"chnName":"品牌","engName":"brandEnglishName"}}
		CustomPropertyFields *string `json:"CustomPropertyFields,omitempty" name:"CustomPropertyFields"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductPricesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductPricesResponse) FromJsonString(s string) error {
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

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 多个产品代码,例如['p_cvm','p_cvm2'...]
	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes" list`
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

type DescribeQuotaAssignLogsListRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 例如：“2019-10-10 10:05:01”
	CreateStartTime *string `json:"CreateStartTime,omitempty" name:"CreateStartTime"`

	// 例如：“2019-10-12 10:05:02”
	CreateEndTime *string `json:"CreateEndTime,omitempty" name:"CreateEndTime"`

	// uin
	SearchUin *string `json:"SearchUin,omitempty" name:"SearchUin"`

	// 产品标签
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品标签
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 项目标签
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 子项目标签
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 创建人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 第几页，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页行数，默认30
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配额类型：sys或uin
	QuotaType *string `json:"QuotaType,omitempty" name:"QuotaType"`
}

func (r *DescribeQuotaAssignLogsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaAssignLogsListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaAssignLogsListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 账户配额转移流水列表
		QuotaAssignList []*QuotaAssignList `json:"QuotaAssignList,omitempty" name:"QuotaAssignList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaAssignLogsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaAssignLogsListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaTypeListRequest struct {
	*tchttp.BaseRequest

	// midas appid
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeQuotaTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaTypeListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaTypeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额类别列表
		QuoteTypeList *QuoteTypeList `json:"QuoteTypeList,omitempty" name:"QuoteTypeList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaTypeListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotasRequest struct {
	*tchttp.BaseRequest

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 传空字符串返回system默认配额， 
	// 不传就返回全量用户配额
	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`

	// 产品code。不传或者空就是查询全部
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotasResponse struct {
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

func (r *DescribeQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRefundRequest struct {
	*tchttp.BaseRequest

	// 米大师appid，由计平分配
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id
	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`

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

		// 数量
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

type DescribeResourceListRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户，只支持单个查询
	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 支付模式
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
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
		List []*ResourceListData `json:"List,omitempty" name:"List" list`

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

type DescribeSubBillingItemsRequest struct {
	*tchttp.BaseRequest

	// midas appid
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 计费细项code列表
	SubBillingItemCodes []*string `json:"SubBillingItemCodes,omitempty" name:"SubBillingItemCodes" list`
}

func (r *DescribeSubBillingItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubBillingItemsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubBillingItemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计费细项信息列表
		List *ProductSubBillingInfoList `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubBillingItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubBillingItemsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubProductsRequest struct {
	*tchttp.BaseRequest

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 子产品代码,例如['sp_cvm','sp_cvm_2']
	SubProductCodes []*string `json:"SubProductCodes,omitempty" name:"SubProductCodes" list`

	// 产品代码或名称
	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 状态，1-生效 0-失效
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeSubProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubProductsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 子产品信息
		List []*SubProductDefineList `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubProductsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAccountRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识
	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`

	// 用户ID
	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`

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

		// 当天累计消耗
		DayOut *string `json:"DayOut,omitempty" name:"DayOut"`

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

		// 二级额度（TCE的实际额度/固定额度），注意单位与后付费账户单位保持一致
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

	// 用户ID
	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`

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

		// 账单总数
		AllNum *string `json:"AllNum,omitempty" name:"AllNum"`

		// 返回的账单数量
		RetNum *string `json:"RetNum,omitempty" name:"RetNum"`

		// 当前游标
		Offset *string `json:"Offset,omitempty" name:"Offset"`

		// 账单详情列表
		DebtBillSet []*DebtBillListData `json:"DebtBillSet,omitempty" name:"DebtBillSet" list`

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

type DetailConditionValue struct {

	// 产品列表
	Business *ConditionBusiness `json:"Business,omitempty" name:"Business"`

	// 子产品列表
	Product *ConditionProduct `json:"Product,omitempty" name:"Product"`

	// 组件列表
	Component *ConditionComponent `json:"Component,omitempty" name:"Component"`

	// 地域列表
	Region *ConditionRegion `json:"Region,omitempty" name:"Region"`

	// 付费模式列表
	PayMode *ConditionPayMode `json:"PayMode,omitempty" name:"PayMode"`

	// 交易类型列表
	ActionType *ConditionActionType `json:"ActionType,omitempty" name:"ActionType"`
}

type EasyProduct struct {

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 商品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
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

	// 账户uin
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`

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

		// 剩余配额数组
		List []*QuotaLeft `json:"List,omitempty" name:"List" list`

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

	// AccountUin
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
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

	// uin
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`

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

		// 数量
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 计量推量明细数组
		ResourceBill []*BspResourceBill `json:"ResourceBill,omitempty" name:"ResourceBill" list`

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

type ListResourceRequest struct {
	*tchttp.BaseRequest

	// 页序号，从0开始
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// uin
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`

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

		// 查询结果总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 资源状态信息结果集
		List []*ResourceInfo `json:"List,omitempty" name:"List" list`

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

	// 1:已退款
	RefundFlag *string `json:"RefundFlag,omitempty" name:"RefundFlag"`

	// 支付时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`
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
	Trend *SummaryTrend `json:"Trend,omitempty" name:"Trend"`
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

type ProductBillingItemInfoList struct {

	// 产品code（第一层）
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品计费项code（第三层）
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 中文名
	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`

	// 英文名
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

type ProductPriceList struct {

	// 产品标签
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品标签
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费项标签
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费项名称
	BillingItemChnName *string `json:"BillingItemChnName,omitempty" name:"BillingItemChnName"`

	// 计费细项标签
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 计费细项名称
	SubBillingItemChnName *string `json:"SubBillingItemChnName,omitempty" name:"SubBillingItemChnName"`

	// priceID,价格策略id
	PriceID *string `json:"PriceID,omitempty" name:"PriceID"`

	// 付费类型
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 计费周期
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 计费模式
	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`

	// 定价类型
	PriceType *string `json:"PriceType,omitempty" name:"PriceType"`

	// 地域ID
	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`

	// 地域代码
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 价格,单位元
	Price *string `json:"Price,omitempty" name:"Price"`

	// 最近操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// JSON string；是自定义属性keyCode和value组成。  
	// 例如：{"system":"windows","brand":"tencent"}
	CustomPropertyValues *string `json:"CustomPropertyValues,omitempty" name:"CustomPropertyValues"`

	// 产品项英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`

	// 产品细项英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
}

type ProductSubBillingInfoList struct {

	// 产品code（第一层）
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品计费项code（第三层）
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 产品计费细项code（第四层）
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 中文名
	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`

	// 英文名
	EngName *string `json:"EngName,omitempty" name:"EngName"`

	// 单位
	GoodNumUnit *string `json:"GoodNumUnit,omitempty" name:"GoodNumUnit"`

	// 单位英文名
	GoodNumUnitEn *string `json:"GoodNumUnitEn,omitempty" name:"GoodNumUnitEn"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type QuotaAssignList struct {

	// 计费项变更前数据
	BillingItemAfterData *string `json:"BillingItemAfterData,omitempty" name:"BillingItemAfterData"`

	// 计费项变更后数据
	BillingItemBeforeData *string `json:"BillingItemBeforeData,omitempty" name:"BillingItemBeforeData"`

	// 计费项code
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费项中文名
	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`

	// 变更内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 变更时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 操作人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 产品变更前数据
	ProductAfterData *string `json:"ProductAfterData,omitempty" name:"ProductAfterData"`

	// 产品变更后数据
	ProductBeforeData *string `json:"ProductBeforeData,omitempty" name:"ProductBeforeData"`

	// 产品code
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品中文名
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 配额类型
	QuotaType *string `json:"QuotaType,omitempty" name:"QuotaType"`

	// 计费子项变更前数据
	SubBillingItemAfterData *string `json:"SubBillingItemAfterData,omitempty" name:"SubBillingItemAfterData"`

	// 计费子项变更后数据
	SubBillingItemBeforeData *string `json:"SubBillingItemBeforeData,omitempty" name:"SubBillingItemBeforeData"`

	// 计费子项code
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 计费子项中文名
	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`

	// 子产品变更前数据
	SubProductAfterData *string `json:"SubProductAfterData,omitempty" name:"SubProductAfterData"`

	// 子产品变更后数据
	SubProductBeforeData *string `json:"SubProductBeforeData,omitempty" name:"SubProductBeforeData"`

	// 子产品code
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品中文名
	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`

	// 被修改配额的用户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 产品英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`

	// 子产品英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`

	// 产品项英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`

	// 产品细项英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`

	// 英文内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngContent *string `json:"EngContent,omitempty" name:"EngContent"`
}

type QuotaDetailList struct {

	// 账户uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品配额标识
	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`

	// 产品配额值
	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 子产品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 子计费码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 产品名
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品类名
	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`

	// 子产品名
	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`

	// 计费名
	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`

	// 子计费名
	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`

	// 单位名
	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type QuotaLeft struct {

	// 配额key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 剩余配额
	Value *int64 `json:"Value,omitempty" name:"Value"`
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

type QuoteTypeList struct {

	// select的key
	QuotaTypeKey *string `json:"QuotaTypeKey,omitempty" name:"QuotaTypeKey"`

	// select的显示名称
	QuotaTypeValue *string `json:"QuotaTypeValue,omitempty" name:"QuotaTypeValue"`

	// select的显示英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaTypeValueEng *string `json:"QuotaTypeValueEng,omitempty" name:"QuotaTypeValueEng"`
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

type ResourceInfo struct {

	// uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// app ID
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 创建时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 资源状态
	BizStatus *string `json:"BizStatus,omitempty" name:"BizStatus"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 子产品码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品名称
	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`

	// 地区ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地区名
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 区域ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 销毁时间(新增字段，342前已销毁资源没有该数据）
	BizIsolatedTime *string `json:"BizIsolatedTime,omitempty" name:"BizIsolatedTime"`

	// 修改时间(也记录了状态变销毁的时间)
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 区域名
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 计费码
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费名称
	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`

	// 子计费码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 子计费名称
	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`

	// 资源数量
	Num *string `json:"Num,omitempty" name:"Num"`
}

type ResourceListData struct {

	// 用户uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 商品码名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 资源数量
	Num *uint64 `json:"Num,omitempty" name:"Num"`

	// 0正常1隔离中2隔离3恢复中4销毁中5销毁
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ResourceSummaryDetailComponentDetailItem struct {

	// 组件用量
	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// 组件码
	ComponentCode *string `json:"ComponentCode,omitempty" name:"ComponentCode"`

	// 组件名称
	ComponentCodeName *string `json:"ComponentCodeName,omitempty" name:"ComponentCodeName"`

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

type SaveQuotaDataArray struct {

	// 如果传了就是需要新增或编辑
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// ProductCode的数据项 如果传了就是需要新增或编辑
	ProductQuotaValue *int64 `json:"ProductQuotaValue,omitempty" name:"ProductQuotaValue"`

	// 如果传了就是需要新增或编辑
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// SubProductCode 的数据项 如果传了就是需要新增或编辑
	SubProductQuotaValue *int64 `json:"SubProductQuotaValue,omitempty" name:"SubProductQuotaValue"`

	// 如果传了就是需要新增或编辑
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// BillingItemCode 的数据项 如果传了就是需要新增或编辑
	BillingItemQuotaValue *int64 `json:"BillingItemQuotaValue,omitempty" name:"BillingItemQuotaValue"`

	// 如果传了就是需要新增或编辑
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// SubBillingItemCode 的数据项 如果传了就是需要新增或编辑
	SubBillingItemQuotaValue *int64 `json:"SubBillingItemQuotaValue,omitempty" name:"SubBillingItemQuotaValue"`
}

type SaveQuotaRequest struct {
	*tchttp.BaseRequest

	// 用户uin。传空字符串就新增或编辑system默认配额
	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 要新增或更新的数据数组
	SaveDataArray []*SaveQuotaDataArray `json:"SaveDataArray,omitempty" name:"SaveDataArray" list`
}

func (r *SaveQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SaveQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SaveQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SaveQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetProductInfoRequest struct {
	*tchttp.BaseRequest

	// 产品定义
	Product *Product `json:"Product,omitempty" name:"Product"`

	// 子产品定义
	SubProduct *SubProduct `json:"SubProduct,omitempty" name:"SubProduct"`

	// 计费项定义
	BillingItem *BillingItem `json:"BillingItem,omitempty" name:"BillingItem"`

	// 计费细项定义
	SubBillingItem *SubBillingItem `json:"SubBillingItem,omitempty" name:"SubBillingItem"`

	// 别名
	AliasCode *string `json:"AliasCode,omitempty" name:"AliasCode"`

	// 米大师AppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *SetProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetProductInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetProductInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubBillingItem struct {

	// 计费细项代码
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 计费细项中文名
	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`

	// 计费细项单位
	GoodNumUnit *string `json:"GoodNumUnit,omitempty" name:"GoodNumUnit"`

	// 计费细项英文名
	EngName *string `json:"EngName,omitempty" name:"EngName"`

	// 计费细项英文单位
	GoodNumUnitEn *string `json:"GoodNumUnitEn,omitempty" name:"GoodNumUnitEn"`

	// 计费细项状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SubProduct struct {

	// 子产品代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品中文名称
	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`

	// 子产品英文名称
	EngName *string `json:"EngName,omitempty" name:"EngName"`

	// 子产品状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SubProductDefineList struct {

	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品中文名称
	ProductChnName *string `json:"ProductChnName,omitempty" name:"ProductChnName"`

	// 产品英文名称
	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`

	// 子产品英文代码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 子产品中文名称
	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`

	// 子产品英文名称
	EngName *string `json:"EngName,omitempty" name:"EngName"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SummaryByResourceConditionValue struct {

	// 产品列表
	Business *ConditionBusiness `json:"Business,omitempty" name:"Business"`

	// 子产品列表
	Product *ConditionProduct `json:"Product,omitempty" name:"Product"`

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

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
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

type TrendByMonthDetail struct {

	// 当前月份统计
	Current *MonthCostDetailItem `json:"Current,omitempty" name:"Current"`

	// 历史月份统计
	Previous []*MonthCostDetailItem `json:"Previous,omitempty" name:"Previous" list`

	// 未来月份统计
	Next *MonthCostDetailItem `json:"Next,omitempty" name:"Next"`
}

type TrendByMonthStat struct {

	// 平均花费详情
	Average *TrendByMonthStatItem `json:"Average,omitempty" name:"Average"`
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

	// 保留字段8
	Reserve8 *string `json:"Reserve8,omitempty" name:"Reserve8"`

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
}
