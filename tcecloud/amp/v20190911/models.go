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

type AlertField struct {

	// 自定义告警字段
	Key *string `json:"Key,omitempty" name:"Key"`

	// 自定义告警值
	Value *string `json:"Value,omitempty" name:"Value"`
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
