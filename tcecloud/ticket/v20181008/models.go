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

package v20181008

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BasicFirstCategory struct {

	// 一级分类Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 一级分类信息
	Cell *FirstCategoryCell `json:"Cell,omitempty" name:"Cell"`

	// 子节点分类数组
	Children []*ChildrenCategory `json:"Children,omitempty" name:"Children" list`
}

type BasicThirdLevel struct {

	// 三级分类ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 三级分类信息
	Cell *ThirdCategoryCell `json:"Cell,omitempty" name:"Cell"`
}

type BasicTicket struct {

	// 工单ID
	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`

	// 工单一级分类
	TypeLevel1 *uint64 `json:"TypeLevel1,omitempty" name:"TypeLevel1"`

	// 工单二级分类
	TypeLevel2 *uint64 `json:"TypeLevel2,omitempty" name:"TypeLevel2"`

	// 工单三级分类
	TypeLevel3 *uint64 `json:"TypeLevel3,omitempty" name:"TypeLevel3"`

	// 租户端工单状态
	AskerStatus *uint64 `json:"AskerStatus,omitempty" name:"AskerStatus"`

	// 运营端工单状态
	AnswererStatus *uint64 `json:"AnswererStatus,omitempty" name:"AnswererStatus"`

	// 优先级
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

	// 工单内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 机密信息
	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`

	// 客户ID
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 工单当前处理客服
	CurrentStaff *string `json:"CurrentStaff,omitempty" name:"CurrentStaff"`

	// 结单人
	Finisher *string `json:"Finisher,omitempty" name:"Finisher"`

	// 工单创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 结单时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 撤单时间
	CancelTime *string `json:"CancelTime,omitempty" name:"CancelTime"`
}

type ChildrenCategory struct {

	// 分类Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 二级分类
	Cell *SecondCategoryCell `json:"Cell,omitempty" name:"Cell"`
}

type Code struct {

	// 状态码
	Code *uint64 `json:"Code,omitempty" name:"Code"`
}

type CosParams struct {

	// SecretId
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// SecretKey
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type CountTicket struct {

	// 待确认结单工单数
	WaitingFinishCount *uint64 `json:"WaitingFinishCount,omitempty" name:"WaitingFinishCount"`

	// 待补充工单数
	WaitingSupplyCount *uint64 `json:"WaitingSupplyCount,omitempty" name:"WaitingSupplyCount"`

	// 处理中工单数
	ProcessingCount *uint64 `json:"ProcessingCount,omitempty" name:"ProcessingCount"`

	// 工单总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type CountTicketRequest struct {
	*tchttp.BaseRequest

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *CountTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CountTicketRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CountTicketResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 工单数量
		Data *CountTicket `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CountTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CountTicketResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTicketRequest struct {
	*tchttp.BaseRequest

	// 工单一级分类
	TypeLevel1 *uint64 `json:"TypeLevel1,omitempty" name:"TypeLevel1"`

	// 工单二级分类
	TypeLevel2 *uint64 `json:"TypeLevel2,omitempty" name:"TypeLevel2"`

	// 工单三级分类
	TypeLevel3 *uint64 `json:"TypeLevel3,omitempty" name:"TypeLevel3"`

	// 工单内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`

	// 工单机密信息
	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`

	// 自定义字段
	CustomFields *string `json:"CustomFields,omitempty" name:"CustomFields"`
}

func (r *CreateTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTicketRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTicketResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 工单ID
		Data *ObjectId `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTicketResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomField struct {

	// 字段ID
	Fields *uint64 `json:"Fields,omitempty" name:"Fields"`

	// 归属的三级分类ID
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 表单字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 表单字段类型
	Types *string `json:"Types,omitempty" name:"Types"`

	// 可选的键值
	FieldKv *string `json:"FieldKv,omitempty" name:"FieldKv"`

	// 字段文案
	Copywriter *string `json:"Copywriter,omitempty" name:"Copywriter"`

	// 排序权重
	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`

	// 是否必填
	Required *uint64 `json:"Required,omitempty" name:"Required"`
}

type CustomerCancelTicketRequest struct {
	*tchttp.BaseRequest

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`

	// 工单ID
	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
}

func (r *CustomerCancelTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CustomerCancelTicketRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerCancelTicketResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态码
		Data *Code `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CustomerCancelTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CustomerCancelTicketResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerFinishTicketRequest struct {
	*tchttp.BaseRequest

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`

	// 工单ID
	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
}

func (r *CustomerFinishTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CustomerFinishTicketRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerFinishTicketResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态码
		Data *Code `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CustomerFinishTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CustomerFinishTicketResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerRespondRequest struct {
	*tchttp.BaseRequest

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`

	// 回复内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 工单ID
	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`

	// 机密信息
	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
}

func (r *CustomerRespondRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CustomerRespondRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerRespondResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态码
		Data *Code `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CustomerRespondResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CustomerRespondResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerUrgeRequest struct {
	*tchttp.BaseRequest

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`

	// 工单ID
	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
}

func (r *CustomerUrgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CustomerUrgeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerUrgeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态码
		Data *Code `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CustomerUrgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CustomerUrgeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCosBuckets struct {

	// appid
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 桶名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeFirstCategorysData struct {

	// 状态码
	Code *uint64 `json:"Code,omitempty" name:"Code"`

	// 提示信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 一级分类信息数组
	Data []*BasicFirstCategory `json:"Data,omitempty" name:"Data" list`
}

type DescribeThirdCategorysData struct {

	// 状态码
	Code *uint64 `json:"Code,omitempty" name:"Code"`

	// 提示信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 三级分类信息数组
	Data []*BasicThirdLevel `json:"Data,omitempty" name:"Data" list`
}

type DescribeTicketRequest struct {
	*tchttp.BaseRequest

	// 工单ID
	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *DescribeTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTicketRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 工单详情
		Data *Ticket `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTicketResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketsData struct {

	// 工单列表
	Data []*BasicTicket `json:"Data,omitempty" name:"Data" list`

	// 工单总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DescribeTicketsRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`

	// 工单过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeTicketsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTicketsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 工单列表
		Data []*DescribeTicketsData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTicketsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTicketsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 工单一级分类
	TypeLevel1 []*uint64 `json:"TypeLevel1,omitempty" name:"TypeLevel1" list`

	// 工单二级分类
	TypeLevel2 []*uint64 `json:"TypeLevel2,omitempty" name:"TypeLevel2" list`

	// 工单三级分类
	TypeLevel3 []*uint64 `json:"TypeLevel3,omitempty" name:"TypeLevel3" list`

	// 工单优先级
	Priority []*uint64 `json:"Priority,omitempty" name:"Priority" list`

	// 客户ID
	CustomerId []*string `json:"CustomerId,omitempty" name:"CustomerId" list`

	// 当前处理客服
	CurrentStaff []*string `json:"CurrentStaff,omitempty" name:"CurrentStaff" list`

	// 结单人
	Finisher *string `json:"Finisher,omitempty" name:"Finisher"`

	// 创建时间（起点）
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" name:"CreateTimeStart"`

	// 结单时间（起点）
	FinishTimeStart *string `json:"FinishTimeStart,omitempty" name:"FinishTimeStart"`

	// 撤单时间（起点）
	CancelTimeStart *string `json:"CancelTimeStart,omitempty" name:"CancelTimeStart"`

	// 创建时间（终点）
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" name:"CreateTimeEnd"`

	// 结单时间（终点）
	FinishTimeEnd *string `json:"FinishTimeEnd,omitempty" name:"FinishTimeEnd"`

	// 撤单时间（终点）
	CancelTimeEnd *string `json:"CancelTimeEnd,omitempty" name:"CancelTimeEnd"`

	// 按创建时间排序,asc或desc
	OrderByCreateTime *string `json:"OrderByCreateTime,omitempty" name:"OrderByCreateTime"`

	// 按结单时间排序,asc或desc
	OrderByFinishTime *string `json:"OrderByFinishTime,omitempty" name:"OrderByFinishTime"`

	// 按撤单时间排序,asc或desc
	OrderByCancelTime *string `json:"OrderByCancelTime,omitempty" name:"OrderByCancelTime"`

	// 工单状态（兼容租户端和运营端）
	Status []*uint64 `json:"Status,omitempty" name:"Status" list`

	// 工单ID
	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`

	// 按更新时间排序,asc或desc
	OrderByUpdateTime *string `json:"OrderByUpdateTime,omitempty" name:"OrderByUpdateTime"`
}

type FirstCategoryCell struct {

	// 一级分类名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 显示排序权重
	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
}

type GetCosBucketRequest struct {
	*tchttp.BaseRequest

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetCosBucketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCosBucketRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCosBucketResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 公用桶参数
		Public []*DescribeCosBuckets `json:"Public,omitempty" name:"Public" list`

		// 私有桶参数
		Private []*DescribeCosBuckets `json:"Private,omitempty" name:"Private" list`

		// COS域名
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCosBucketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCosBucketResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCosParamsByBucketRequest struct {
	*tchttp.BaseRequest

	// 桶类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetCosParamsByBucketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCosParamsByBucketRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCosParamsByBucketResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// COS密钥
		Data *CosParams `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCosParamsByBucketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCosParamsByBucketResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTicketCategoryByLevelRequest struct {
	*tchttp.BaseRequest

	// 分类ID
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetTicketCategoryByLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTicketCategoryByLevelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTicketCategoryByLevelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 三级分类和分类对应的表单字段
		Data *DescribeThirdCategorysData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTicketCategoryByLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTicketCategoryByLevelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTicketCategoryRequest struct {
	*tchttp.BaseRequest

	// 请求来源
	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetTicketCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTicketCategoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTicketCategoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询一级分类返回的数据结构
		Data *DescribeFirstCategorysData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTicketCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTicketCategoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ObjectId struct {

	// 对象ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type Operation struct {

	// 操作人角色
	Role *string `json:"Role,omitempty" name:"Role"`

	// 内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 是否存在机密信息
	SecretContent *bool `json:"SecretContent,omitempty" name:"SecretContent"`

	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`
}

type SecondCategoryCell struct {

	// 分类名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分类描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 分类是否常用
	Common *uint64 `json:"Common,omitempty" name:"Common"`

	// 分类显示权重
	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`

	// 分类图标
	ImgUrl *string `json:"ImgUrl,omitempty" name:"ImgUrl"`
}

type ThirdCategoryCell struct {

	// 三级分类名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 显示排序权重
	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`

	// 表单字段数组
	CustomFields []*CustomField `json:"CustomFields,omitempty" name:"CustomFields" list`
}

type Ticket struct {

	// 工单ID
	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`

	// 工单一级分类
	TypeLevel1 *uint64 `json:"TypeLevel1,omitempty" name:"TypeLevel1"`

	// 工单二级分类
	TypeLevel2 *uint64 `json:"TypeLevel2,omitempty" name:"TypeLevel2"`

	// 工单三级分类
	TypeLevel3 *uint64 `json:"TypeLevel3,omitempty" name:"TypeLevel3"`

	// 优先级
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

	// 工单内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 是否存在机密信息
	SecretContent *bool `json:"SecretContent,omitempty" name:"SecretContent"`

	// 客户ID
	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`

	// 工单当前处理客服
	CurrentStaff *string `json:"CurrentStaff,omitempty" name:"CurrentStaff"`

	// 结单人
	Finisher *string `json:"Finisher,omitempty" name:"Finisher"`

	// 工单创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 结单时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 撤单时间
	CancelTime *string `json:"CancelTime,omitempty" name:"CancelTime"`

	// 工单状态（租户端）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 工单状态名称
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// 操作记录
	Operations []*Operation `json:"Operations,omitempty" name:"Operations" list`
}
