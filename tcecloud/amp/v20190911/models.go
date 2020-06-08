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

package v20190911

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type AddConvergence struct {

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 返回信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 告警收敛id
	ConvId *string `json:"ConvId,omitempty" name:"ConvId"`
}

type AddConvergenceRequest struct {
	*tchttp.BaseRequest

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 无
	Data *Conv `json:"Data,omitempty" name:"Data"`
}

func (r *AddConvergenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddConvergenceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddConvergenceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *AddConvergence `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddConvergenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddConvergenceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddHandle struct {

	// 告警处理id
	HandleId *string `json:"HandleId,omitempty" name:"HandleId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 返回信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

type AddHandleRequest struct {
	*tchttp.BaseRequest

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 无
	Data *Handle `json:"Data,omitempty" name:"Data"`
}

func (r *AddHandleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddHandleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddHandleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *AddHandle `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddHandleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddHandleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddShield struct {

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 返回信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 屏蔽id
	ShieldId *string `json:"ShieldId,omitempty" name:"ShieldId"`
}

type AddShieldRequest struct {
	*tchttp.BaseRequest

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 无
	Data *Shield `json:"Data,omitempty" name:"Data"`
}

func (r *AddShieldRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddShieldRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddShieldResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *AddShield `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddShieldResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddShieldResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSub struct {

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 告警订阅id
	SubsId *string `json:"SubsId,omitempty" name:"SubsId"`
}

type AddSubRequest struct {
	*tchttp.BaseRequest

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 订阅信息
	Data *Subs `json:"Data,omitempty" name:"Data"`
}

func (r *AddSubRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSubRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSubResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *AddSub `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSubResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSubResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTopic struct {

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 返回信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type AddTopicRequest struct {
	*tchttp.BaseRequest

	// 告警源信息
	Data *Topic `json:"Data,omitempty" name:"Data"`
}

func (r *AddTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *AddTopic `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlertField struct {

	// 自定义告警字段
	Key *string `json:"Key,omitempty" name:"Key"`

	// 自定义告警值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ConfigHistory struct {

	// id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 开发者账号
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 主账号uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 子账户uin
	SubUin *int64 `json:"SubUin,omitempty" name:"SubUin"`

	// 请求id
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 编辑时间戳
	EditTime *int64 `json:"EditTime,omitempty" name:"EditTime"`

	// 编辑类型
	EditType *string `json:"EditType,omitempty" name:"EditType"`

	// 编辑内容
	EditItem *string `json:"EditItem,omitempty" name:"EditItem"`

	// 编辑开始前
	BeforeEdit *string `json:"BeforeEdit,omitempty" name:"BeforeEdit"`

	// 编辑开始后
	AfterEdit *string `json:"AfterEdit,omitempty" name:"AfterEdit"`
}

type Conv struct {

	// 告警源ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 收敛规则名
	ConvName *string `json:"ConvName,omitempty" name:"ConvName"`

	// 收敛规则ID 更新必填
	ConvId *string `json:"ConvId,omitempty" name:"ConvId"`

	// 屏蔽条件
	ConvRules []*JudgeRules `json:"ConvRules,omitempty" name:"ConvRules" list`

	// 收敛时长
	ConvPeriod *int64 `json:"ConvPeriod,omitempty" name:"ConvPeriod"`

	// 最大收敛数
	ConvCounts *int64 `json:"ConvCounts,omitempty" name:"ConvCounts"`

	// 开发者帐号
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 告警订阅创建者uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 最后更新订阅的uin
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 最后更新时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 插入时间戳
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 收敛维度
	ConvRulesDimension *string `json:"ConvRulesDimension,omitempty" name:"ConvRulesDimension"`
}

type DeleteConvergencesRequest struct {
	*tchttp.BaseRequest

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 收敛规则id
	ConvIds []*string `json:"ConvIds,omitempty" name:"ConvIds" list`
}

func (r *DeleteConvergencesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConvergencesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConvergencesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConvergencesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConvergencesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteHandlesRequest struct {
	*tchttp.BaseRequest

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 无
	HandleIds []*string `json:"HandleIds,omitempty" name:"HandleIds" list`
}

func (r *DeleteHandlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteHandlesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteHandlesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHandlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteHandlesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteShieldsRequest struct {
	*tchttp.BaseRequest

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 屏蔽规则id
	ShieldIds []*string `json:"ShieldIds,omitempty" name:"ShieldIds" list`
}

func (r *DeleteShieldsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteShieldsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteShieldsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteShieldsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteShieldsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubsRequest struct {
	*tchttp.BaseRequest

	// 订阅规则id
	SubsIds []*string `json:"SubsIds,omitempty" name:"SubsIds" list`
}

func (r *DeleteSubsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicsRequest struct {
	*tchttp.BaseRequest

	// 告警源id数组
	TopicIds []*string `json:"TopicIds,omitempty" name:"TopicIds" list`
}

func (r *DeleteTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Field struct {

	// 参数字段
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段类型:string字符型  number数字型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数字段名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 默认值
	Default *string `json:"Default,omitempty" name:"Default"`
}

type FieldNew struct {

	// 参数字段
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段类型:string字符型  number数字型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数字段名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 是否为默认
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type GetAlertHistory struct {

	// 告警历史
	Histories []*History `json:"Histories,omitempty" name:"Histories" list`

	// 自定义字段
	Columns []*FieldNew `json:"Columns,omitempty" name:"Columns" list`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetAlertHistoryById struct {

	// 告警历史
	History *History `json:"History,omitempty" name:"History"`

	// 自定义告警字段
	Columns []*FieldNew `json:"Columns,omitempty" name:"Columns" list`
}

type GetAlertHistoryByIdRequest struct {
	*tchttp.BaseRequest

	// 告警id
	AlertId *string `json:"AlertId,omitempty" name:"AlertId"`
}

func (r *GetAlertHistoryByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAlertHistoryByIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAlertHistoryByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetAlertHistoryById `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertHistoryByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAlertHistoryByIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAlertHistoryRequest struct {
	*tchttp.BaseRequest

	// 分页需要的偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页中每页的大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索告警发生时间的开始时间戳
	OccurStartTimestamp *uint64 `json:"OccurStartTimestamp,omitempty" name:"OccurStartTimestamp"`

	// 搜索告警发生时间的结束时间戳
	OccurEndTimestamp *uint64 `json:"OccurEndTimestamp,omitempty" name:"OccurEndTimestamp"`

	// 搜索告警接收时间的开始时间戳
	ReceiveStartTimestamp *uint64 `json:"ReceiveStartTimestamp,omitempty" name:"ReceiveStartTimestamp"`

	// 搜索告警接收时间的结束时间戳
	ReceiveEndTimestamp *uint64 `json:"ReceiveEndTimestamp,omitempty" name:"ReceiveEndTimestamp"`

	// 搜索告警的状态限制 all 全部  continue 非连续点 convergence 已收敛 shield 已屏蔽 send 已发送 reserve 无订阅
	Status *string `json:"Status,omitempty" name:"Status"`

	// 要搜索告警的主题Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 全局搜索内容
	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`

	// 排序字段
	SortKey *string `json:"SortKey,omitempty" name:"SortKey"`

	// 排序字段顺序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 分字段进行告警搜索KV对 例如 [{\"key\":\"alarmStatus\",\"value\":\"告警中\"}]
	SearchKeyValuePairs *string `json:"SearchKeyValuePairs,omitempty" name:"SearchKeyValuePairs"`

	// 告警订阅Id
	SubId *string `json:"SubId,omitempty" name:"SubId"`
}

func (r *GetAlertHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAlertHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAlertHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetAlertHistory `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAlertHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConfigHistory struct {

	// 日志历史
	Histories []*ConfigHistory `json:"Histories,omitempty" name:"Histories" list`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetConfigHistoryRequest struct {
	*tchttp.BaseRequest

	// 分页需要的偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页中每页的大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 开始时间戳
	StartTimestamp *uint64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 结束时间戳
	EndTimestamp *uint64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`

	// 要搜索告警源Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 搜索类型 all create  update delete
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// 排序字段
	SortKey *string `json:"SortKey,omitempty" name:"SortKey"`

	// 排序字段顺序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *GetConfigHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConfigHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConfigHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetConfigHistory `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConfigHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConfigHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConvergence struct {

	// 收敛规则信息
	Convs *Conv `json:"Convs,omitempty" name:"Convs"`
}

type GetConvergenceRequest struct {
	*tchttp.BaseRequest

	// 无
	ConvId *string `json:"ConvId,omitempty" name:"ConvId"`
}

func (r *GetConvergenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConvergenceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConvergenceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Data *GetConvergence `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConvergenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConvergenceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConvergences struct {

	// 告警收敛规则
	Convs []*Conv `json:"Convs,omitempty" name:"Convs" list`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetConvergencesRequest struct {
	*tchttp.BaseRequest

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 无
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetConvergencesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConvergencesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConvergencesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetConvergences `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConvergencesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConvergencesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetHandle struct {

	// 处理规则
	Handle *Handle `json:"Handle,omitempty" name:"Handle"`
}

type GetHandleRequest struct {
	*tchttp.BaseRequest

	// 无
	HandleId *string `json:"HandleId,omitempty" name:"HandleId"`
}

func (r *GetHandleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetHandleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetHandleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetHandle `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHandleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetHandleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetHandles struct {

	// 处理规则信息
	Handles []*Handle `json:"Handles,omitempty" name:"Handles" list`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetHandlesRequest struct {
	*tchttp.BaseRequest

	// 无
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *GetHandlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetHandlesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetHandlesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetHandles `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHandlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetHandlesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetShield struct {

	// 屏蔽规则信息
	Shield *Shield `json:"Shield,omitempty" name:"Shield"`
}

type GetShieldRequest struct {
	*tchttp.BaseRequest

	// 无
	ShieldId *string `json:"ShieldId,omitempty" name:"ShieldId"`
}

func (r *GetShieldRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetShieldRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetShieldResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetShield `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetShieldResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetShieldResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetShields struct {

	// 屏蔽规则信息
	Shields []*Shield `json:"Shields,omitempty" name:"Shields" list`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetShieldsRequest struct {
	*tchttp.BaseRequest

	// 无
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *GetShieldsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetShieldsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetShieldsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetShields `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetShieldsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetShieldsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSub struct {

	// 告警订阅信息
	Subs *Subs `json:"Subs,omitempty" name:"Subs"`
}

type GetSubRequest struct {
	*tchttp.BaseRequest

	// 订阅规则id
	SubsId *string `json:"SubsId,omitempty" name:"SubsId"`
}

func (r *GetSubRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSubRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSubResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetSub `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSubResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSubs struct {

	// 告警订阅信息
	Subs []*Subs `json:"Subs,omitempty" name:"Subs" list`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetSubsRequest struct {
	*tchttp.BaseRequest

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 无
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索信息
	SearchSubsName *string `json:"SearchSubsName,omitempty" name:"SearchSubsName"`
}

func (r *GetSubsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSubsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSubsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetSubs `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSubsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopic struct {

	// 告警源信息
	Topic *Topic `json:"Topic,omitempty" name:"Topic"`
}

type GetTopicFieldTemplate struct {

	// 返回信息
	Columns []*FieldNew `json:"Columns,omitempty" name:"Columns" list`
}

type GetTopicFieldTemplateRequest struct {
	*tchttp.BaseRequest

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *GetTopicFieldTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicFieldTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicFieldTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetTopicFieldTemplate `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicFieldTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicFieldTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicRequest struct {
	*tchttp.BaseRequest

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *GetTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetTopic `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopics struct {

	// 告警源数据
	Topics []*Topic `json:"Topics,omitempty" name:"Topics" list`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetTopicsRequest struct {
	*tchttp.BaseRequest

	// 起始位置0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页限制长度
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊搜索告警源名称
	SearchTopicName *string `json:"SearchTopicName,omitempty" name:"SearchTopicName"`
}

func (r *GetTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *GetTopics `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Handle struct {

	// 告警源ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 告警处理名
	HandleName *string `json:"HandleName,omitempty" name:"HandleName"`

	// 判断条件类型 all / custom
	HandleJudgeType *string `json:"HandleJudgeType,omitempty" name:"HandleJudgeType"`

	// 处理方式 pre / parallel
	HandleActionType *string `json:"HandleActionType,omitempty" name:"HandleActionType"`

	// 回调url
	HandleActionURL *string `json:"HandleActionURL,omitempty" name:"HandleActionURL"`

	// 告警优先级
	HandleActionPriority *int64 `json:"HandleActionPriority,omitempty" name:"HandleActionPriority"`

	// 验证码
	HandleActionCode *string `json:"HandleActionCode,omitempty" name:"HandleActionCode"`

	// 告警处理ID 更新必填
	HandleId *string `json:"HandleId,omitempty" name:"HandleId"`

	// 告警处理条件
	HandleJudgeRules []*JudgeRules `json:"HandleJudgeRules,omitempty" name:"HandleJudgeRules" list`

	// 开发者帐号
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 屏蔽规则创建者uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 最后更新屏蔽规则的uin
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 最后更新时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 插入时间戳
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 是否验证
	HandleActionVerification *int64 `json:"HandleActionVerification,omitempty" name:"HandleActionVerification"`
}

type History struct {

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 告警源名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主账号uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 子账号uin
	SubUin *int64 `json:"SubUin,omitempty" name:"SubUin"`

	// 告警id
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 发生时间
	OccurTime *int64 `json:"OccurTime,omitempty" name:"OccurTime"`

	// 接收时间
	ReceiveTime *int64 `json:"ReceiveTime,omitempty" name:"ReceiveTime"`

	// 告警状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 接收组
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups" list`

	// 发送方式
	NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays" list`

	// 自定义字段
	AlertField []*AlertField `json:"AlertField,omitempty" name:"AlertField" list`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

type Inform struct {

	// 短信通道告警模板
	SMS *InformTmp `json:"SMS,omitempty" name:"SMS"`

	// 语音通知通道告警模板
	CALL *InformTmp `json:"CALL,omitempty" name:"CALL"`

	// 电子邮件通道告警模板
	EMAIL *InformTmp `json:"EMAIL,omitempty" name:"EMAIL"`

	// 微信通道告警模板
	WECHAT *InformTmp `json:"WECHAT,omitempty" name:"WECHAT"`
}

type InformTmp struct {

	// 是否启用该模板
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 告警模板标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 告警模板
	Template *string `json:"Template,omitempty" name:"Template"`

	// 告警模板匹配规则
	Rules []*JudgeRules `json:"Rules,omitempty" name:"Rules" list`
}

type JudgeRules struct {

	// 自定义参数
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段类型:string字符型  number数字型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 规则符  ==,!=,<=,>=
	Calc *string `json:"Calc,omitempty" name:"Calc"`

	// 规则值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SendAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 自定义参数
	Params *string `json:"Params,omitempty" name:"Params"`
}

func (r *SendAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendAlarmRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendAlarmResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Shield struct {

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 屏蔽规则名
	ShieldName *string `json:"ShieldName,omitempty" name:"ShieldName"`

	// 屏蔽条件 all / custom
	ShieldType *string `json:"ShieldType,omitempty" name:"ShieldType"`

	// 屏蔽条件 period / absolute
	ShieldTimeType *string `json:"ShieldTimeType,omitempty" name:"ShieldTimeType"`

	// 开始屏蔽时间戳
	ShieldTimeStart *int64 `json:"ShieldTimeStart,omitempty" name:"ShieldTimeStart"`

	// 结束屏蔽时间戳
	ShieldTimeEnd *int64 `json:"ShieldTimeEnd,omitempty" name:"ShieldTimeEnd"`

	// 屏蔽规则ID 更新必填
	ShieldId *string `json:"ShieldId,omitempty" name:"ShieldId"`

	// 自定义屏蔽条件
	ShieldRules []*JudgeRules `json:"ShieldRules,omitempty" name:"ShieldRules" list`

	// 周期屏蔽时的屏蔽设置星期几
	ShieldDayOfWeek []*int64 `json:"ShieldDayOfWeek,omitempty" name:"ShieldDayOfWeek" list`

	// 开发者帐号
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 屏蔽规则创建者uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 最后更新屏蔽规则的uin
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 最后更新时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 插入时间戳
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type Subs struct {

	// 告警源id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 订阅名称
	SubsName *string `json:"SubsName,omitempty" name:"SubsName"`

	// 订阅类型 all / custom
	SubsType *string `json:"SubsType,omitempty" name:"SubsType"`

	// 通知方式 SMS,EMAIL,WECHAT,CALL
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// 订阅id 更新必填
	SubsId *string `json:"SubsId,omitempty" name:"SubsId"`

	// 告警源名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅规则
	SubsRules []*JudgeRules `json:"SubsRules,omitempty" name:"SubsRules" list`

	// 接收用户组
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups" list`

	// 语音告警接收人顺序
	PhoneNotifyOrder []*int64 `json:"PhoneNotifyOrder,omitempty" name:"PhoneNotifyOrder" list`

	// 语音告警轮询次数
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitempty" name:"PhoneCircleTimes"`

	// 语音告警轮内间隔
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitempty" name:"PhoneInnerInterval"`

	// 语音告警次间间隔
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitempty" name:"PhoneCircleInterval"`

	// 语音告警触达通知
	PhoneArriveNotice *int64 `json:"PhoneArriveNotice,omitempty" name:"PhoneArriveNotice"`

	// 开发者帐号
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 告警订阅创建者uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 最后更新订阅的uin
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 最后更新时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 插入时间戳
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 订阅有效开始时间
	ValidStartTime *int64 `json:"ValidStartTime,omitempty" name:"ValidStartTime"`

	// 订阅有效结束时间
	ValidEndTime *int64 `json:"ValidEndTime,omitempty" name:"ValidEndTime"`

	// 接收人
	ReceiverUins []*int64 `json:"ReceiverUins,omitempty" name:"ReceiverUins" list`
}

type Topic struct {

	// 告警源名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 告警源类型(默认参数为manual)
	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`

	// 自定义字段信息
	Fields []*Field `json:"Fields,omitempty" name:"Fields" list`

	// 告警源id 更新必填
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Appid
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 创建人uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 最后编辑人uin
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 更新时间（时间戳）
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间（时间戳）
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 告警模板
	InformTemplate *Inform `json:"InformTemplate,omitempty" name:"InformTemplate"`
}

type UpdateConvergenceRequest struct {
	*tchttp.BaseRequest

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 无
	Data *Conv `json:"Data,omitempty" name:"Data"`
}

func (r *UpdateConvergenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateConvergenceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateConvergenceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateConvergenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateConvergenceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateHandleRequest struct {
	*tchttp.BaseRequest

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 无
	Data *Handle `json:"Data,omitempty" name:"Data"`
}

func (r *UpdateHandleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateHandleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateHandleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateHandleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateHandleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateOrDelReturn struct {

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

type UpdateSubRequest struct {
	*tchttp.BaseRequest

	// 无
	SubsId *string `json:"SubsId,omitempty" name:"SubsId"`

	// 无
	Data *Subs `json:"Data,omitempty" name:"Data"`

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *UpdateSubRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSubRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSubResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSubResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSubResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateTopicRequest struct {
	*tchttp.BaseRequest

	// 修改告警主题Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 修改告警内容
	Data *Topic `json:"Data,omitempty" name:"Data"`
}

func (r *UpdateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *UpdateOrDelReturn `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
