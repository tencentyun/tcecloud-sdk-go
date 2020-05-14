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

package v20181210

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessRuleInterfaceIDeleteRuleRequest struct {
	*tchttp.BaseRequest

	// 1
	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`

	// 1
	Data []*IdList `json:"Data,omitempty" name:"Data" list`

	// 1
	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`

	// 1
	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`

	// 1
	InfoValue *string `json:"InfoValue,omitempty" name:"InfoValue"`

	// 1
	Reboot *bool `json:"Reboot,omitempty" name:"Reboot"`
}

func (r *AccessRuleInterfaceIDeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AccessRuleInterfaceIDeleteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIDeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessRuleInterfaceIDeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AccessRuleInterfaceIDeleteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIUpdateRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 1
	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`

	// 1
	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`

	// 1
	InfoValue *string `json:"InfoValue,omitempty" name:"InfoValue"`

	// "data": [{
	//   "id" : "2",                # 唯一ID
	// }],
	Data []*IdList `json:"Data,omitempty" name:"Data" list`

	// 1
	Reboot *bool `json:"Reboot,omitempty" name:"Reboot"`

	// 1
	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`
}

func (r *AccessRuleInterfaceIUpdateRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AccessRuleInterfaceIUpdateRuleStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIUpdateRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessRuleInterfaceIUpdateRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AccessRuleInterfaceIUpdateRuleStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetSystemRawLogListRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 页数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 查询标志
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 时间
	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime" list`
}

func (r *AlarmInterfaceIGetSystemRawLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetSystemRawLogListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetSystemRawLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetSystemRawLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetSystemRawLogListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetSystemRawLogStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间
	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime" list`

	// 查询标识
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 当前页数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
}

func (r *AlarmInterfaceIGetSystemRawLogStatisticInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetSystemRawLogStatisticInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetSystemRawLogStatisticInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetSystemRawLogStatisticInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetSystemRawLogStatisticInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogDetailRequest struct {
	*tchttp.BaseRequest

	// ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *AlarmInterfaceIGetThreatRawLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetThreatRawLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogHistogramRequest struct {
	*tchttp.BaseRequest

	// 根据ip，端口，cgi，host等信息搜索，精准匹配
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 威胁等级，低危，中危，高危，严重，致命
	ThreatLevel *string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`

	// 威胁类型
	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`

	// 层级
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 拦截，未拦截
	Status *string `json:"Status,omitempty" name:"Status"`

	// 柱状图的时间类型
	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`

	// 多余的参数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 多余的参数
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 多余的参数
	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime" list`

	// 多余的参数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *AlarmInterfaceIGetThreatRawLogHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogHistogramRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogHistogramResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetThreatRawLogHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogHistogramResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogListRequest struct {
	*tchttp.BaseRequest

	// 精准匹配
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 威胁等级
	ThreatLevel *string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`

	// 威胁类型
	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`

	// 应用层，传输层
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 当前页面
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 查询标志
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 多余的参数
	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime" list`
}

func (r *AlarmInterfaceIGetThreatRawLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetThreatRawLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 查询内容
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 威胁等级
	ThreatLevel *string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`

	// 威胁类型
	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`

	// 应用层，传输层
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 拦截，未拦截
	Status *string `json:"Status,omitempty" name:"Status"`

	// 多余的参数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 多余的参数
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 多余的参数
	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime" list`

	// 多余的参数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *AlarmInterfaceIGetThreatRawLogStatisticInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogStatisticInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogStatisticInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetThreatRawLogStatisticInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogStatisticInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BackendTaskInterfaceIGetDataRequest struct {
	*tchttp.BaseRequest

	// 1
	BackendTaskId *int64 `json:"BackendTaskId,omitempty" name:"BackendTaskId"`
}

func (r *BackendTaskInterfaceIGetDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BackendTaskInterfaceIGetDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BackendTaskInterfaceIGetDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BackendTaskInterfaceIGetDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BackendTaskInterfaceIGetDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIAddApiConfigRequest struct {
	*tchttp.BaseRequest

	// 授权
	SecretIds *string `json:"SecretIds,omitempty" name:"SecretIds"`

	// 密钥
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 平台名称
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 阻断持续时间，单位/分钟
	WindowTime *int64 `json:"WindowTime,omitempty" name:"WindowTime"`

	// 调用数量上限
	WindowCount *int64 `json:"WindowCount,omitempty" name:"WindowCount"`

	// ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 公司
	Company *string `json:"Company,omitempty" name:"Company"`
}

func (r *BlockConfigInterfaceIAddApiConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIAddApiConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIAddApiConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIAddApiConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIAddApiConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIChangeBlockStatusRequest struct {
	*tchttp.BaseRequest

	// "off/on"  off表示关，on表示开
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

func (r *BlockConfigInterfaceIChangeBlockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIChangeBlockStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIChangeBlockStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIChangeBlockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIChangeBlockStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIConfigBlockTimeRequest struct {
	*tchttp.BaseRequest

	// 分钟为单位
	BlackDuration *uint64 `json:"BlackDuration,omitempty" name:"BlackDuration"`
}

func (r *BlockConfigInterfaceIConfigBlockTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIConfigBlockTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIConfigBlockTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIConfigBlockTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIConfigBlockTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIConfigRedirectURLRequest struct {
	*tchttp.BaseRequest

	// 跳转地址
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`
}

func (r *BlockConfigInterfaceIConfigRedirectURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIConfigRedirectURLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIConfigRedirectURLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIConfigRedirectURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIConfigRedirectURLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIDeleteApiConfigRequest struct {
	*tchttp.BaseRequest

	// ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *BlockConfigInterfaceIDeleteApiConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIDeleteApiConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIDeleteApiConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIDeleteApiConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIDeleteApiConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIDeleteStopNetworkConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *BlockConfigInterfaceIDeleteStopNetworkConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIDeleteStopNetworkConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIDeleteStopNetworkConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIDeleteStopNetworkConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIDeleteStopNetworkConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIEditApiConfigRequest struct {
	*tchttp.BaseRequest

	// 一ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 阻断持续时间，单位/分钟
	WindowTime *int64 `json:"WindowTime,omitempty" name:"WindowTime"`

	// 调用数量上限
	WindowCount *int64 `json:"WindowCount,omitempty" name:"WindowCount"`

	// 1
	IsUse *int64 `json:"IsUse,omitempty" name:"IsUse"`

	// 1
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 1
	Company *string `json:"Company,omitempty" name:"Company"`

	// 1
	SecretIds *string `json:"SecretIds,omitempty" name:"SecretIds"`

	// 1
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

func (r *BlockConfigInterfaceIEditApiConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIEditApiConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIEditApiConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIEditApiConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIEditApiConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiBlockStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *BlockConfigInterfaceIGetApiBlockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetApiBlockStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiBlockStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetApiBlockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetApiBlockStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiBlockTrustThresholdRequest struct {
	*tchttp.BaseRequest
}

func (r *BlockConfigInterfaceIGetApiBlockTrustThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetApiBlockTrustThresholdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiBlockTrustThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetApiBlockTrustThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetApiBlockTrustThresholdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiConfigDetailRequest struct {
	*tchttp.BaseRequest

	// ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *BlockConfigInterfaceIGetApiConfigDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetApiConfigDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiConfigDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetApiConfigDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetApiConfigDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiConfigInfoListRequest struct {
	*tchttp.BaseRequest

	// 每页多少行
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// 第几页
	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
}

func (r *BlockConfigInterfaceIGetApiConfigInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetApiConfigInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiConfigInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetApiConfigInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetApiConfigInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetAvailableCountRequest struct {
	*tchttp.BaseRequest

	// id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *BlockConfigInterfaceIGetAvailableCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetAvailableCountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetAvailableCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetAvailableCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetAvailableCountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetBlockTimeRequest struct {
	*tchttp.BaseRequest
}

func (r *BlockConfigInterfaceIGetBlockTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetBlockTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetBlockTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetBlockTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetBlockTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetCurrentBlockStateRequest struct {
	*tchttp.BaseRequest
}

func (r *BlockConfigInterfaceIGetCurrentBlockStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetCurrentBlockStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetCurrentBlockStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetCurrentBlockStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetCurrentBlockStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetNetoffConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *BlockConfigInterfaceIGetNetoffConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetNetoffConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetNetoffConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetNetoffConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetNetoffConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetOrResetSecretKeyRequest struct {
	*tchttp.BaseRequest

	// id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *BlockConfigInterfaceIGetOrResetSecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetOrResetSecretKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetOrResetSecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetOrResetSecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetOrResetSecretKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetRebootAvailableCountRequest struct {
	*tchttp.BaseRequest

	// ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *BlockConfigInterfaceIGetRebootAvailableCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetRebootAvailableCountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetRebootAvailableCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetRebootAvailableCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetRebootAvailableCountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetRedirectURLRequest struct {
	*tchttp.BaseRequest
}

func (r *BlockConfigInterfaceIGetRedirectURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetRedirectURLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetRedirectURLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetRedirectURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetRedirectURLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetSecretInfoRequest struct {
	*tchttp.BaseRequest

	// time
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// user
	User *string `json:"User,omitempty" name:"User"`

	// passwd
	Passwd *string `json:"Passwd,omitempty" name:"Passwd"`

	// id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Key
	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *BlockConfigInterfaceIGetSecretInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetSecretInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetSecretInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetSecretInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIGetSecretInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIRebootApiRequest struct {
	*tchttp.BaseRequest

	// id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// add_count
	AddCount *int64 `json:"AddCount,omitempty" name:"AddCount"`
}

func (r *BlockConfigInterfaceIRebootApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIRebootApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIRebootApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIRebootApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIRebootApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceISetApiBlockStatusRequest struct {
	*tchttp.BaseRequest

	// 0为关闭，1为开启
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *BlockConfigInterfaceISetApiBlockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceISetApiBlockStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceISetApiBlockStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceISetApiBlockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceISetApiBlockStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceISetApiConfigStatusRequest struct {
	*tchttp.BaseRequest

	// Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Status
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *BlockConfigInterfaceISetApiConfigStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceISetApiConfigStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceISetApiConfigStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceISetApiConfigStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceISetApiConfigStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIStartNetworkRequest struct {
	*tchttp.BaseRequest
}

func (r *BlockConfigInterfaceIStartNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIStartNetworkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIStartNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIStartNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIStartNetworkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIStopNetworkRequest struct {
	*tchttp.BaseRequest

	// 断网原因
	BlockReason *string `json:"BlockReason,omitempty" name:"BlockReason"`

	// 断网原因html
	BlockReasonHtml *string `json:"BlockReasonHtml,omitempty" name:"BlockReasonHtml"`

	// 断网开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 断网结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *BlockConfigInterfaceIStopNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIStopNetworkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIStopNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIStopNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIStopNetworkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest struct {
	*tchttp.BaseRequest

	// 置信度阀值
	Threshold *float64 `json:"Threshold,omitempty" name:"Threshold"`
}

func (r *BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetAlarmLogListRequest struct {
	*tchttp.BaseRequest

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DayOverviewInterfaceIGetAlarmLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetAlarmLogListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetAlarmLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetAlarmLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetAlarmLogListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetAlarmReasonCntRequest struct {
	*tchttp.BaseRequest

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DayOverviewInterfaceIGetAlarmReasonCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetAlarmReasonCntRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetAlarmReasonCntResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetAlarmReasonCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetAlarmReasonCntResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetComplianceInfoRequest struct {
	*tchttp.BaseRequest

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DayOverviewInterfaceIGetComplianceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetComplianceInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetComplianceInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetComplianceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetComplianceInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetNetFlowRequest struct {
	*tchttp.BaseRequest

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DayOverviewInterfaceIGetNetFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetNetFlowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetNetFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetNetFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetNetFlowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetOverviewImportantInfoRequest struct {
	*tchttp.BaseRequest

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DayOverviewInterfaceIGetOverviewImportantInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetOverviewImportantInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetOverviewImportantInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetOverviewImportantInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetOverviewImportantInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetPlatformControlInfoRequest struct {
	*tchttp.BaseRequest

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DayOverviewInterfaceIGetPlatformControlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetPlatformControlInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetPlatformControlInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetPlatformControlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DayOverviewInterfaceIGetPlatformControlInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIAddCheckTaskRequest struct {
	*tchttp.BaseRequest

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`

	// 1
	Data []*GlobalImport `json:"Data,omitempty" name:"Data" list`
}

func (r *GlobalAccessInterfaceIAddCheckTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIAddCheckTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIAddCheckTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIAddCheckTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIAddCheckTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIAddRuleRequest struct {
	*tchttp.BaseRequest

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`

	// 1
	Name *string `json:"Name,omitempty" name:"Name"`

	// 1
	ValidDuration *string `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 1
	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`

	// 1
	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`

	// 1
	IsOverwrite *string `json:"IsOverwrite,omitempty" name:"IsOverwrite"`

	// 1
	IpType *string `json:"IpType,omitempty" name:"IpType"`

	// 1
	IpData []*GlobalImport `json:"IpData,omitempty" name:"IpData" list`

	// 1
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GlobalAccessInterfaceIAddRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIAddRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIAddRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIAddRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIAddRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceICommitCheckTaskRuleRequest struct {
	*tchttp.BaseRequest

	// 1
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 1
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
}

func (r *GlobalAccessInterfaceICommitCheckTaskRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceICommitCheckTaskRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceICommitCheckTaskRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceICommitCheckTaskRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceICommitCheckTaskRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIEditRuleRequest struct {
	*tchttp.BaseRequest

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`

	// 1
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 1
	Name *string `json:"Name,omitempty" name:"Name"`

	// 1
	ValidDuration *string `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 1
	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`

	// 1
	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`

	// 1
	IsOverwrite *string `json:"IsOverwrite,omitempty" name:"IsOverwrite"`

	// 1
	IpType *string `json:"IpType,omitempty" name:"IpType"`

	// 1
	IpData []*GlobalImport `json:"IpData,omitempty" name:"IpData" list`

	// 1
	Reboot *bool `json:"Reboot,omitempty" name:"Reboot"`
}

func (r *GlobalAccessInterfaceIEditRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIEditRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIEditRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIEditRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIEditRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetCheckTaskResultListRequest struct {
	*tchttp.BaseRequest

	// task_id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// check_status
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetCheckTaskResultListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetCheckTaskResultRequest struct {
	*tchttp.BaseRequest

	// task_id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetCheckTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetRuleListRequest struct {
	*tchttp.BaseRequest

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`

	// 1
	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`

	// 1
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 1
	Status *string `json:"Status,omitempty" name:"Status"`

	// 1
	ValidDuration *string `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 1
	HitHistory *string `json:"HitHistory,omitempty" name:"HitHistory"`

	// 1
	HitWeek *string `json:"HitWeek,omitempty" name:"HitWeek"`

	// 1
	HitDay *string `json:"HitDay,omitempty" name:"HitDay"`

	// 1
	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`

	// 1
	IpType *string `json:"IpType,omitempty" name:"IpType"`

	// 1
	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`

	// 1
	Order *OrderField `json:"Order,omitempty" name:"Order"`

	// 1
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 1
	AddTime []*string `json:"AddTime,omitempty" name:"AddTime" list`

	// 1
	ExcuteFlag *int64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 1
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
}

func (r *GlobalAccessInterfaceIGetRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIGetRuleListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetRuleListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIGetRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIGetRuleListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetStatisticsRequest struct {
	*tchttp.BaseRequest

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`

	// 1
	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`

	// 1
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 1
	Status *string `json:"Status,omitempty" name:"Status"`

	// 1
	ValidDuration *string `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 1
	HitHistory *string `json:"HitHistory,omitempty" name:"HitHistory"`

	// 1
	HitWeek *string `json:"HitWeek,omitempty" name:"HitWeek"`

	// 1
	HitDay *string `json:"HitDay,omitempty" name:"HitDay"`

	// 1
	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`

	// 1
	IpType *string `json:"IpType,omitempty" name:"IpType"`

	// 1
	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`

	// 1
	AddTime []*string `json:"AddTime,omitempty" name:"AddTime" list`

	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 1
	ExcuteFlag *int64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 1
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 1
	Order *OrderField `json:"Order,omitempty" name:"Order"`

	// 1
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *GlobalAccessInterfaceIGetStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIGetStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIGetStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GlobalAccessInterfaceIGetStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GlobalImport struct {

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源IP
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 目的IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 有效期
	ValidDuration *string `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 匹配动作
	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`

	// 来源平台
	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
}

type IdList struct {

	// ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type LabelManagerInterfaceIAddSourcePlatformRequest struct {
	*tchttp.BaseRequest

	// 1
	Label *string `json:"Label,omitempty" name:"Label"`
}

func (r *LabelManagerInterfaceIAddSourcePlatformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LabelManagerInterfaceIAddSourcePlatformRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIAddSourcePlatformResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LabelManagerInterfaceIAddSourcePlatformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LabelManagerInterfaceIAddSourcePlatformResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIDeleteSourcePlatformRequest struct {
	*tchttp.BaseRequest

	// 1
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *LabelManagerInterfaceIDeleteSourcePlatformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LabelManagerInterfaceIDeleteSourcePlatformRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIDeleteSourcePlatformResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LabelManagerInterfaceIDeleteSourcePlatformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LabelManagerInterfaceIDeleteSourcePlatformResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIGetSourcePlatformRequest struct {
	*tchttp.BaseRequest

	// 1
	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
}

func (r *LabelManagerInterfaceIGetSourcePlatformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LabelManagerInterfaceIGetSourcePlatformRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIGetSourcePlatformResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LabelManagerInterfaceIGetSourcePlatformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LabelManagerInterfaceIGetSourcePlatformResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceIGetMsgListRequest struct {
	*tchttp.BaseRequest

	// "已读" or "未读" or "all"
	Status *string `json:"Status,omitempty" name:"Status"`

	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 1
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// 1
	ExcuteFlag *int64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
}

func (r *MsgInterfaceIGetMsgListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MsgInterfaceIGetMsgListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceIGetMsgListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MsgInterfaceIGetMsgListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MsgInterfaceIGetMsgListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceIGetMsgStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 1
	Status *string `json:"Status,omitempty" name:"Status"`

	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 1
	ExcuteFlag *int64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 1
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *MsgInterfaceIGetMsgStatisticInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MsgInterfaceIGetMsgStatisticInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceIGetMsgStatisticInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MsgInterfaceIGetMsgStatisticInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MsgInterfaceIGetMsgStatisticInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceISetSingleUnreadMsgToReadRequest struct {
	*tchttp.BaseRequest

	// 1
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *MsgInterfaceISetSingleUnreadMsgToReadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MsgInterfaceISetSingleUnreadMsgToReadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceISetSingleUnreadMsgToReadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MsgInterfaceISetSingleUnreadMsgToReadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MsgInterfaceISetSingleUnreadMsgToReadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceISetUnreadMsgToReadRequest struct {
	*tchttp.BaseRequest
}

func (r *MsgInterfaceISetUnreadMsgToReadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MsgInterfaceISetUnreadMsgToReadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceISetUnreadMsgToReadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MsgInterfaceISetUnreadMsgToReadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MsgInterfaceISetUnreadMsgToReadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OrderField struct {

	// 要排序字
	Field *string `json:"Field,omitempty" name:"Field"`

	// 排序方向
	Direc *string `json:"Direc,omitempty" name:"Direc"`
}

type OverviewInterfaceIExportDataRequest struct {
	*tchttp.BaseRequest

	// 1
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// "data_list": [
	// "deal-index",
	// "warn-index",
	// "intercept-index",
	// "warn-analyse",
	// "rule-index",
	// "warn-distribution"
	// ], // 分别代表处置指标、告警指标、阻断指标、告警分析、规则指标、告警分布
	DataList []*string `json:"DataList,omitempty" name:"DataList" list`
}

func (r *OverviewInterfaceIExportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIExportDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIExportDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIExportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIExportDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmLogListRequest struct {
	*tchttp.BaseRequest

	// date
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *OverviewInterfaceIGetAlarmLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetAlarmLogListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetAlarmLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetAlarmLogListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmReasonCntRequest struct {
	*tchttp.BaseRequest

	// date
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *OverviewInterfaceIGetAlarmReasonCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetAlarmReasonCntRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmReasonCntResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetAlarmReasonCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetAlarmReasonCntResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmStatInfoRequest struct {
	*tchttp.BaseRequest

	// 1
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *OverviewInterfaceIGetAlarmStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetAlarmStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetAlarmStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetAlarmStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmTrendStatInfoRequest struct {
	*tchttp.BaseRequest

	// 1
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *OverviewInterfaceIGetAlarmTrendStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetAlarmTrendStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmTrendStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetAlarmTrendStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetAlarmTrendStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetBlockStatInfoRequest struct {
	*tchttp.BaseRequest

	// StartTime
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// EndTime
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *OverviewInterfaceIGetBlockStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetBlockStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetBlockStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetBlockStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetBlockStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetBlockTrendStatInfoRequest struct {
	*tchttp.BaseRequest

	// 1
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *OverviewInterfaceIGetBlockTrendStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetBlockTrendStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetBlockTrendStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetBlockTrendStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetBlockTrendStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetCheckDistribStatInfoRequest struct {
	*tchttp.BaseRequest

	// 1
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *OverviewInterfaceIGetCheckDistribStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetCheckDistribStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetCheckDistribStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetCheckDistribStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetCheckDistribStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetComplianceInfoRequest struct {
	*tchttp.BaseRequest

	// date
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *OverviewInterfaceIGetComplianceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetComplianceInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetComplianceInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetComplianceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetComplianceInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetNetFlowRequest struct {
	*tchttp.BaseRequest

	// date
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *OverviewInterfaceIGetNetFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetNetFlowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetNetFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetNetFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetNetFlowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetOverviewImportantInfoRequest struct {
	*tchttp.BaseRequest

	// date
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *OverviewInterfaceIGetOverviewImportantInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetOverviewImportantInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetOverviewImportantInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetOverviewImportantInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetOverviewImportantInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetPlatformControlInfoRequest struct {
	*tchttp.BaseRequest

	// date
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *OverviewInterfaceIGetPlatformControlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetPlatformControlInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetPlatformControlInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetPlatformControlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetPlatformControlInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetRuleStatInfoRequest struct {
	*tchttp.BaseRequest

	// 1
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *OverviewInterfaceIGetRuleStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetRuleStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetRuleStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetRuleStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetRuleStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetRuleTrendStatInfoRequest struct {
	*tchttp.BaseRequest

	// StartTime
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// EndTime
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *OverviewInterfaceIGetRuleTrendStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetRuleTrendStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetRuleTrendStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetRuleTrendStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetRuleTrendStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetTOPStatInfoRequest struct {
	*tchttp.BaseRequest

	// 1
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *OverviewInterfaceIGetTOPStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetTOPStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetTOPStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetTOPStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OverviewInterfaceIGetTOPStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RawlogSearchInterfaceIGetDataFieldsRequest struct {
	*tchttp.BaseRequest

	// 数据字段
	DataSrc *string `json:"DataSrc,omitempty" name:"DataSrc"`
}

func (r *RawlogSearchInterfaceIGetDataFieldsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RawlogSearchInterfaceIGetDataFieldsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RawlogSearchInterfaceIGetDataFieldsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RawlogSearchInterfaceIGetDataFieldsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RawlogSearchInterfaceIGetDataFieldsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RawlogSearchInterfaceIQueryRwLogRequest struct {
	*tchttp.BaseRequest

	// 组之间的操作符,JSON格式
	GroupBetweenCond *string `json:"GroupBetweenCond,omitempty" name:"GroupBetweenCond"`

	// 来源字段
	DataSrc *string `json:"DataSrc,omitempty" name:"DataSrc"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 当前页数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 柱状图按什么单位聚合
	AggreUnit *string `json:"AggreUnit,omitempty" name:"AggreUnit"`

	// JSON格式
	GroupVal *string `json:"GroupVal,omitempty" name:"GroupVal"`

	// 查询标识
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 多余参数
	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime" list`
}

func (r *RawlogSearchInterfaceIQueryRwLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RawlogSearchInterfaceIQueryRwLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RawlogSearchInterfaceIQueryRwLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RawlogSearchInterfaceIQueryRwLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RawlogSearchInterfaceIQueryRwLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RefreshConfigInterfaceIGetRefreshConfigRequest struct {
	*tchttp.BaseRequest

	// "page" : 1 // 1-安全总览，2-安全告警，3-日志审计，4-消息中心
	Page *int64 `json:"Page,omitempty" name:"Page"`
}

func (r *RefreshConfigInterfaceIGetRefreshConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefreshConfigInterfaceIGetRefreshConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RefreshConfigInterfaceIGetRefreshConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshConfigInterfaceIGetRefreshConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefreshConfigInterfaceIGetRefreshConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RefreshConfigInterfaceIUpdateRefreshConfigRequest struct {
	*tchttp.BaseRequest

	// 1
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 1
	AutoRefresh *int64 `json:"AutoRefresh,omitempty" name:"AutoRefresh"`

	// 1
	Frequency *int64 `json:"Frequency,omitempty" name:"Frequency"`
}

func (r *RefreshConfigInterfaceIUpdateRefreshConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefreshConfigInterfaceIUpdateRefreshConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RefreshConfigInterfaceIUpdateRefreshConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshConfigInterfaceIUpdateRefreshConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefreshConfigInterfaceIUpdateRefreshConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportIpFlowInterfaceIGetIpFlowOverviewRequest struct {
	*tchttp.BaseRequest

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *ReportIpFlowInterfaceIGetIpFlowOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportIpFlowInterfaceIGetIpFlowOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportIpFlowInterfaceIGetIpFlowOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportIpFlowInterfaceIGetIpFlowOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportIpFlowInterfaceIGetIpFlowOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportIpFlowInterfaceIGetIpFlowTableRequest struct {
	*tchttp.BaseRequest

	// 当前页数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 默认按in方向的流量峰值 大到小，JSON格式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询标记，多余参数
	ExcuteFlag *int64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 搜索内容
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
}

func (r *ReportIpFlowInterfaceIGetIpFlowTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportIpFlowInterfaceIGetIpFlowTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportIpFlowInterfaceIGetIpFlowTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportIpFlowInterfaceIGetIpFlowTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportIpFlowInterfaceIGetIpFlowTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportNetFlowInterfaceIGetNetFlowViewDetailRequest struct {
	*tchttp.BaseRequest

	// 趋势 trend, 对比 compare
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查看单日，可以为空或者与start_date相等，否则必须是有效的时间区间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *ReportNetFlowInterfaceIGetNetFlowViewDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportNetFlowInterfaceIGetNetFlowViewDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportNetFlowInterfaceIGetNetFlowViewDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportNetFlowInterfaceIGetNetFlowViewDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportNetFlowInterfaceIGetNetFlowViewDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest struct {
	*tchttp.BaseRequest

	// 规则列表
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchDeleteRuleRequest struct {
	*tchttp.BaseRequest

	// 删除的规则数组，JSON格式
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *RuleConfigManagerInterfaceIBatchDeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchDeleteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchDeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIBatchDeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchDeleteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchInsertRuleRequest struct {
	*tchttp.BaseRequest

	// 插入的规则，格式是JSON
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *RuleConfigManagerInterfaceIBatchInsertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchInsertRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchInsertRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIBatchInsertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchInsertRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIChangeWorkStatusRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 动作
	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`

	// 名单类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *RuleConfigManagerInterfaceIChangeWorkStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIChangeWorkStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIChangeWorkStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIChangeWorkStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIChangeWorkStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIDeleteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 动作
	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`

	// 名单类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *RuleConfigManagerInterfaceIDeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIDeleteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIDeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIDeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIDeleteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest struct {
	*tchttp.BaseRequest

	// 查询标识
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 名单类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 当前页数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时间
	AddTime []*string `json:"AddTime,omitempty" name:"AddTime" list`

	// 查询内容
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
}

func (r *RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIInsertRuleRequest struct {
	*tchttp.BaseRequest

	// 标题名称
	Name []*string `json:"Name,omitempty" name:"Name" list`

	// 白名单类型，ip...
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 白名单内容
	MatchDetail []*string `json:"MatchDetail,omitempty" name:"MatchDetail" list`

	// 有效期
	ValidDuration *float64 `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 黑白名单类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 操作
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 阻断方式
	MatchOperate *string `json:"MatchOperate,omitempty" name:"MatchOperate"`
}

func (r *RuleConfigManagerInterfaceIInsertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIInsertRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIInsertRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIInsertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIInsertRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIListRuleRequest struct {
	*tchttp.BaseRequest

	// 规则状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 黑白名单类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间
	AddTime []*string `json:"AddTime,omitempty" name:"AddTime" list`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 当前页数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 执行标志
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 输入内容
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *RuleConfigManagerInterfaceIListRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIListRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIListRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIListRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIListRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIUpdateRuleRequest struct {
	*tchttp.BaseRequest

	// 修改后的标题名称
	Name []*string `json:"Name,omitempty" name:"Name" list`

	// 需要修改得规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 白名单类型，ip...
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 白名单内容
	MatchDetail []*string `json:"MatchDetail,omitempty" name:"MatchDetail" list`

	// 有效期(整数)，有效时长，**注意永久用-1表示**
	ValidDuration *float64 `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 黑白名单类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 操作
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 添加时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 阻断方式
	MatchOperate *string `json:"MatchOperate,omitempty" name:"MatchOperate"`
}

func (r *RuleConfigManagerInterfaceIUpdateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIUpdateRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIUpdateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIUpdateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RuleConfigManagerInterfaceIUpdateRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleIdList struct {

	// 1
	Id *string `json:"Id,omitempty" name:"Id"`

	// 1
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 1
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
}

type SystemRuleInterfaceIExportDataRequest struct {
	*tchttp.BaseRequest
}

func (r *SystemRuleInterfaceIExportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SystemRuleInterfaceIExportDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIExportDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIExportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SystemRuleInterfaceIExportDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIGetRuleStatInfoRequest struct {
	*tchttp.BaseRequest

	// 1
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 1
	ValidDuration *string `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 1
	HitCount *string `json:"HitCount,omitempty" name:"HitCount"`

	// 1
	Source *string `json:"Source,omitempty" name:"Source"`

	// 1
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 1
	AddTime []*string `json:"AddTime,omitempty" name:"AddTime" list`

	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 1
	HitCount *string `json:"HitCount,omitempty" name:"HitCount"`

	// Order: {Field: "", Direc: ""}排序
	Order *OrderField `json:"Order,omitempty" name:"Order"`

	// 1
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// 1
	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`

	// 1
	Status *string `json:"Status,omitempty" name:"Status"`

	// 1
	ExcuteFlag *int64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
}

func (r *SystemRuleInterfaceIGetRuleStatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SystemRuleInterfaceIGetRuleStatInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIGetRuleStatInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIGetRuleStatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SystemRuleInterfaceIGetRuleStatInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIListRuleRequest struct {
	*tchttp.BaseRequest

	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 1
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// 1
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 1
	ValidDuration *string `json:"ValidDuration,omitempty" name:"ValidDuration"`

	// 1
	HitCount *string `json:"HitCount,omitempty" name:"HitCount"`

	// 1
	Source *string `json:"Source,omitempty" name:"Source"`

	// 11
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 11
	Order *OrderField `json:"Order,omitempty" name:"Order"`

	// 1
	AddTime []*string `json:"AddTime,omitempty" name:"AddTime" list`

	// 1
	ExcuteFlag *string `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// 1
	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`

	// 1
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *SystemRuleInterfaceIListRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SystemRuleInterfaceIListRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIListRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIListRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SystemRuleInterfaceIListRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIStopRulesRequest struct {
	*tchttp.BaseRequest

	// 1
	Rule []*RuleIdList `json:"Rule,omitempty" name:"Rule" list`
}

func (r *SystemRuleInterfaceIStopRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SystemRuleInterfaceIStopRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIStopRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIStopRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SystemRuleInterfaceIStopRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetFlowStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *TceManagerInterfaceIGetFlowStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetFlowStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetFlowStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetFlowStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetFlowStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetSystemRawLogListRequest struct {
	*tchttp.BaseRequest

	// 名单类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 当前页数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时间
	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime" list`

	// 查询标识
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
}

func (r *TceManagerInterfaceIGetSystemRawLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetSystemRawLogListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetSystemRawLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetSystemRawLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetSystemRawLogListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 名单类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 当前页数
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时间
	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime" list`

	// 查询标识
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
}

func (r *TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetThreatRawLogListRequest struct {
	*tchttp.BaseRequest

	// 拦截，未拦截
	Status *string `json:"Status,omitempty" name:"Status"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 当前页
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 查询内容
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询时间
	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime" list`

	// 应用层，传输层
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 威胁等级，低危，中危，高危，严重，致命
	ThreatLevel *string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`

	// webshell检测
	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`

	// 黑名单，黑客攻击
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 执行标志
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
}

func (r *TceManagerInterfaceIGetThreatRawLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetThreatRawLogListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetThreatRawLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetThreatRawLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetThreatRawLogListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 拦截，未拦截
	Status *string `json:"Status,omitempty" name:"Status"`

	// 每页条数
	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`

	// 当前页
	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 查询内容
	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询时间
	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime" list`

	// 应用层，传输层
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 威胁等级，低危，中危，高危，严重，致命
	ThreatLevel *string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`

	// 黑名单，黑客攻击
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 执行标志
	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`

	// webshell检测
	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`
}

func (r *TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIStopFlowRequest struct {
	*tchttp.BaseRequest

	// 操作
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

func (r *TceManagerInterfaceIStopFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIStopFlowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIStopFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIStopFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TceManagerInterfaceIStopFlowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
