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

package v20180701

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateOfflineTaskRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：prm.createOfflineTask
	Op *string `json:"Op,omitempty" name:"Op"`

	// 事件ID
	ApiEventId *int64 `json:"ApiEventId,omitempty" name:"ApiEventId"`

	// User id
	ApiUserId *string `json:"ApiUserId,omitempty" name:"ApiUserId"`

	// API版本
	ApiVersion *int64 `json:"ApiVersion,omitempty" name:"ApiVersion"`

	// 下线机器IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *CreateOfflineTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOfflineTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOfflineTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 返回信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 接口版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 返回数据类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 下线结果
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOfflineTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOfflineTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOnlineTaskRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：prm.createOnlineTask
	Op *string `json:"Op,omitempty" name:"Op"`

	// 事件ID
	ApiEventId *int64 `json:"ApiEventId,omitempty" name:"ApiEventId"`

	// User id
	ApiUserId *string `json:"ApiUserId,omitempty" name:"ApiUserId"`

	// API版本
	ApiVersion *int64 `json:"ApiVersion,omitempty" name:"ApiVersion"`

	// 默认1
	AutoOnline *int64 `json:"AutoOnline,omitempty" name:"AutoOnline"`

	// 待投放的物理机ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 上线子步骤
	Steps *string `json:"Steps,omitempty" name:"Steps"`

	// 投放流程模板id
	TmplId *string `json:"TmplId,omitempty" name:"TmplId"`
}

func (r *CreateOnlineTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOnlineTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOnlineTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 返回信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 接口版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 返回数据类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 上线结果
		DataSet *bool `json:"DataSet,omitempty" name:"DataSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOnlineTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOnlineTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// u9700u8981u8fc7u6ee4u7684u5b57u6bb5
	Name *string `json:"Name,omitempty" name:"Name"`

	// u5b57u6bb5u7684u8fc7u6ee4u503c
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type GetTaskListRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：prm.getTaskList
	Op *string `json:"Op,omitempty" name:"Op"`

	// 事件ID
	ApiEventId *int64 `json:"ApiEventId,omitempty" name:"ApiEventId"`

	// User id
	ApiUserId *string `json:"ApiUserId,omitempty" name:"ApiUserId"`

	// API版本
	ApiVersion *int64 `json:"ApiVersion,omitempty" name:"ApiVersion"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询页
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 查询count
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 区间过滤
	RangeFilters []*RangeFilter `json:"RangeFilters,omitempty" name:"RangeFilters" list`

	// 排序
	OrderFilter *OrderFilter `json:"OrderFilter,omitempty" name:"OrderFilter"`
}

func (r *GetTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTaskListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTaskListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 返回信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 接口版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 返回数据类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 任务信息
		Data *TaskInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTaskListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTmplByIdRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：prm.getTmplById
	Op *string `json:"Op,omitempty" name:"Op"`

	// 模版id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 调用方的user id
	ApiUserId *string `json:"ApiUserId,omitempty" name:"ApiUserId"`
}

func (r *GetTmplByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTmplByIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTmplByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 返回信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 接口版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 返回数据类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 返回模版信息
		DataSet []*TmplInfo `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTmplByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTmplByIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTmplListRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：GetTmplList
	Op *string `json:"Op,omitempty" name:"Op"`

	// 调用方的user id
	ApiUserId *string `json:"ApiUserId,omitempty" name:"ApiUserId"`
}

func (r *GetTmplListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTmplListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTmplListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 返回信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 接口版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 返回数据类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 返回模版信息
		DataSet []*TmplInfo `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTmplListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTmplListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MarkFailRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：prm.markFail
	Op *string `json:"Op,omitempty" name:"Op"`

	// 事件ID
	ApiEventId *int64 `json:"ApiEventId,omitempty" name:"ApiEventId"`

	// User id
	ApiUserId *string `json:"ApiUserId,omitempty" name:"ApiUserId"`

	// API版本
	ApiVersion *int64 `json:"ApiVersion,omitempty" name:"ApiVersion"`

	// 任务ID列表
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`
}

func (r *MarkFailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MarkFailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MarkFailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 返回信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 接口版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 返回数据类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 设置结果
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MarkFailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MarkFailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OrderFilter struct {

	// u9700u8981u8fc7u6ee4u7684u5b57u6bb5
	OrderName *string `json:"OrderName,omitempty" name:"OrderName"`

	// u6392u5e8fu987au5e8fuff1a1u4e3au5347u5e8fuff0c2u4e3au964du5e8f
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

type RangeFilter struct {

	// u9700u8981u8fc7u6ee4u7684u5b57u6bb5
	Name *string `json:"Name,omitempty" name:"Name"`

	// u5b57u6bb5u7684u8d77u59cbu503c
	BeginValue *string `json:"BeginValue,omitempty" name:"BeginValue"`

	// u5b57u6bb5u7684u7ed3u675fu503c
	EndValue *string `json:"EndValue,omitempty" name:"EndValue"`
}

type RetryTaskRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：prm.retryTask
	Op *string `json:"Op,omitempty" name:"Op"`

	// 事件ID
	ApiEventId *int64 `json:"ApiEventId,omitempty" name:"ApiEventId"`

	// User id
	ApiUserId *string `json:"ApiUserId,omitempty" name:"ApiUserId"`

	// API版本
	ApiVersion *int64 `json:"ApiVersion,omitempty" name:"ApiVersion"`

	// 任务ID列表
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`
}

func (r *RetryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RetryTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RetryTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 返回信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 接口版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 返回数据类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 重试结果
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RetryTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Step struct {

	// u6b65u9aa4u987au5e8f
	StepNum *string `json:"StepNum,omitempty" name:"StepNum"`

	// u6b65u9aa4u63cfu8ff0
	StepDesc *string `json:"StepDesc,omitempty" name:"StepDesc"`
}

type TaskDetail struct {

	// u91cdu8bd5u6b21u6570
	Retry *string `json:"Retry,omitempty" name:"Retry"`

	// u6267u884cu8fc7u7684u6b65u9aa4
	ExecutedStep *string `json:"ExecutedStep,omitempty" name:"ExecutedStep"`

	// u672cu5730IPu5730u5740
	LocalIp *string `json:"LocalIp,omitempty" name:"LocalIp"`

	// u4efbu52a1u7ed3u675fu65f6u95f4
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// IPu5730u5740
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// u6b65u9aa4u5f00u59cbu65f6u95f4
	StepStartTime *string `json:"StepStartTime,omitempty" name:"StepStartTime"`

	// u7a0bu5e8fpid
	Pid *string `json:"Pid,omitempty" name:"Pid"`

	// u5f53u524du6b65u9aa4
	CurStep *string `json:"CurStep,omitempty" name:"CurStep"`

	// u53c2u6570
	Param *string `json:"Param,omitempty" name:"Param"`

	// u5b50u4efbu52a1id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// u6b65u9aa4u72b6u6001
	StepStatus *string `json:"StepStatus,omitempty" name:"StepStatus"`

	// u7b49u5f85u6b65u9aa4
	WaitStep []*string `json:"WaitStep,omitempty" name:"WaitStep" list`

	// u521bu5efau65f6u95f4
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// u4efbu52a1id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// u64cdu4f5cu4ebau5458
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// u8f93u51fa
	OutPut *string `json:"OutPut,omitempty" name:"OutPut"`

	// u4efbu52a1u72b6u6001
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// u65b9u6cd5
	Method *string `json:"Method,omitempty" name:"Method"`
}

type TaskInfo struct {

	// u4efbu52a1u5217u8868u4e2au6570
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// u603bu4efbu52a1u4e2au6570
	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// u6240u5728u9875u6570
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// u4efbu52a1u8be6u60c5
	Detail *TaskDetail `json:"Detail,omitempty" name:"Detail"`
}

type TmplInfo struct {

	// u5185u6838u7248u672c
	Kernel *string `json:"Kernel,omitempty" name:"Kernel"`

	// u9ed8u8ba4u201c1u201d
	AppMask *string `json:"AppMask,omitempty" name:"AppMask"`

	// u521bu5efau8005
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// [N][u817eu8bafu4e91u6bcdu673au6c60]
	SrcBsiOne *string `json:"SrcBsiOne,omitempty" name:"SrcBsiOne"`

	// [u8d44u6e90u6c60][u7ebfu4e0bu6d41u8f6c]
	SrcBsiTwo *string `json:"SrcBsiTwo,omitempty" name:"SrcBsiTwo"`

	// u672au5206u914du6a21u5757
	SrcBsiThree *string `json:"SrcBsiThree,omitempty" name:"SrcBsiThree"`

	// N][u817eu8bafu4e91u6bcdu673au6c60]
	DestBsiOne *string `json:"DestBsiOne,omitempty" name:"DestBsiOne"`

	// [u865au62dfu6bcdu673a][u6d4bu8bd5]
	DestBsiTwo *string `json:"DestBsiTwo,omitempty" name:"DestBsiTwo"`

	// [u72ecu7acbu5b9eu9a8cu5ba4][u5176u4ed6]
	DestBsiThree *string `json:"DestBsiThree,omitempty" name:"DestBsiThree"`

	// u865au62dfu5316u7c7bu578b
	Virtual *string `json:"Virtual,omitempty" name:"Virtual"`

	// u6a21u677fid
	Id *string `json:"Id,omitempty" name:"Id"`

	// u5185u5bb9
	Content *string `json:"Content,omitempty" name:"Content"`

	// u7c7bu578buff0cu4e0au7ebfu6216u4e0bu7ebf
	Type *string `json:"Type,omitempty" name:"Type"`

	// u6700u540eu66f4u65b0u65f6u95f4
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// [N][u817eu8bafu4e91u6bcdu673au6c60]
	TransBsiOne *string `json:"TransBsiOne,omitempty" name:"TransBsiOne"`

	// [u865au62dfu6bcdu673a][u6d4bu8bd5]
	TransBsiTwo *string `json:"TransBsiTwo,omitempty" name:"TransBsiTwo"`

	// [u4e34u65f6u4f7fu7528][u5176u4ed6]
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// u6bcdu673au6c60
	Pool *string `json:"Pool,omitempty" name:"Pool"`

	// u6a21u677fu540du79f0
	Name *string `json:"Name,omitempty" name:"Name"`

	// u6a21u677fu6b65u9aa4
	Steps *string `json:"Steps,omitempty" name:"Steps"`

	// u9ed8u8ba41
	AutoInit *string `json:"AutoInit,omitempty" name:"AutoInit"`

	// u6b65u9aa4u5217u8868
	StepList []*Step `json:"StepList,omitempty" name:"StepList" list`
}
