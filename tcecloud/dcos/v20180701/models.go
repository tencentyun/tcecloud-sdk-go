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

type AddIdcLineExportRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_export_line
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：add
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要配置的专线出口信息
	Datas []*IdcLineExport `json:"Datas,omitempty" name:"Datas" list`
}

func (r *AddIdcLineExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddIdcLineExportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddIdcLineExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet *AddInfoResult `json:"DataSet,omitempty" name:"DataSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddIdcLineExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddIdcLineExportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddIdcLineRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_special_line
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：add
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要配置的专线信息
	Datas []*IdcLine `json:"Datas,omitempty" name:"Datas" list`
}

func (r *AddIdcLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddIdcLineRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddIdcLineResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*AddInfoResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddIdcLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddIdcLineResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddInfoResult struct {

	// 结果，0成功，1失败
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 详情
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 添加的索引ID
	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
}

type AddNetDeviceRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：netdevice
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：add
	Op *string `json:"Op,omitempty" name:"Op"`

	// 新增数据
	Datas []*NetDeviceInfo `json:"Datas,omitempty" name:"Datas" list`

	// 开始索引地址
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *AddNetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddNetDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddNetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 新增返回
		DataSet []*AddInfoResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddNetDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddNetPortRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：netdevice_related
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：add
	Op *string `json:"Op,omitempty" name:"Op"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 新增数据
	Datas []*NetPortInfo `json:"Datas,omitempty" name:"Datas" list`
}

func (r *AddNetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddNetPortRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddNetPortResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 新增返回
		DataSet []*AddInfoResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNetPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddNetPortResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddNetSegmentRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：ipresource
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：add
	Op *string `json:"Op,omitempty" name:"Op"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 新增数据
	Datas []*NetSegmentInfo `json:"Datas,omitempty" name:"Datas" list`
}

func (r *AddNetSegmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddNetSegmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddNetSegmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 新增返回
		DataSet []*AddInfoResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNetSegmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddNetSegmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddRackDescRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_rack
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：add
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要配置的机架机位信息
	Datas []*IdcRackDesc `json:"Datas,omitempty" name:"Datas" list`
}

func (r *AddRackDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddRackDescRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddRackDescResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*AddInfoResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddRackDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddRackDescResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddRackInfoRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idcinfo
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：add
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要配置的机架机位信息
	Datas []*IdcRackInfo `json:"Datas,omitempty" name:"Datas" list`
}

func (r *AddRackInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddRackInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddRackInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*AddInfoResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddRackInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddRackInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddServerRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：server
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：add
	Op *string `json:"Op,omitempty" name:"Op"`

	// 操作服务器数据
	Data *ServerData `json:"Data,omitempty" name:"Data"`
}

func (r *AddServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet *string `json:"DataSet,omitempty" name:"DataSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSnmpTemplateRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：add_snmp_template
	Op *string `json:"Op,omitempty" name:"Op"`

	// 模版
	Template *string `json:"Template,omitempty" name:"Template"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *AddSnmpTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSnmpTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSnmpTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSnmpTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSnmpTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmReportCharRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：alarm_report_char
	Op *string `json:"Op,omitempty" name:"Op"`

	// 告警信息
	Info *AlarmReportInfo `json:"Info,omitempty" name:"Info"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *AlarmReportCharRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmReportCharRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmReportCharResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmReportCharResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmReportCharResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmReportCharsRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 报警类型：这里值为batch_alarms
	Op *string `json:"Op,omitempty" name:"Op"`

	// 告警信息
	DataInfos []*AlarmReportInfo `json:"DataInfos,omitempty" name:"DataInfos" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *AlarmReportCharsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmReportCharsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmReportCharsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmReportCharsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmReportCharsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmReportInfo struct {

	// 0->服务器，1->网络设备，2->网络专线，3->网络出口，4->自定义告警项，5->网络质量
	ItemType *int64 `json:"ItemType,omitempty" name:"ItemType"`

	// 告警类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 告警IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 固资号
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 告警系统
	AlarmSystem *string `json:"AlarmSystem,omitempty" name:"AlarmSystem"`

	// 告警类型
	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`

	// 对象类型，默认为host
	ObjType *string `json:"ObjType,omitempty" name:"ObjType"`

	// 告警索引ID
	AlarmIndexId *int64 `json:"AlarmIndexId,omitempty" name:"AlarmIndexId"`

	// 告警索引名称
	AlarmIndexName *string `json:"AlarmIndexName,omitempty" name:"AlarmIndexName"`

	// 告警内容 - 字符型
	Content *string `json:"Content,omitempty" name:"Content"`

	// 告警时间戳
	HappenTime *int64 `json:"HappenTime,omitempty" name:"HappenTime"`

	// 名字
	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`

	// 上报时间戳
	ReportTime *int64 `json:"ReportTime,omitempty" name:"ReportTime"`

	// 告警人
	NoticePeople *string `json:"NoticePeople,omitempty" name:"NoticePeople"`

	// 告警内容
	AlarmContent *string `json:"AlarmContent,omitempty" name:"AlarmContent"`

	// 告警等级
	AlarmLevel *string `json:"AlarmLevel,omitempty" name:"AlarmLevel"`

	// 告警状态
	AlarmStatus *int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 是否恢复
	HasRecovery *int64 `json:"HasRecovery,omitempty" name:"HasRecovery"`

	// To ecc
	ToEcc *int64 `json:"ToEcc,omitempty" name:"ToEcc"`

	// 二级CI
	SubItem *string `json:"SubItem,omitempty" name:"SubItem"`
}

type AlarmReportNum struct {

	// 告警信息
	Info *AlarmReportInfo `json:"Info,omitempty" name:"Info"`

	// 需要上报的数值
	Datas []*ValueArray `json:"Datas,omitempty" name:"Datas" list`
}

type AlarmReportNumRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：alarm_report_num
	Op *string `json:"Op,omitempty" name:"Op"`

	// 告警信息
	Info *AlarmReportInfo `json:"Info,omitempty" name:"Info"`

	// 需要上报的数值
	Datas []*ValueArray `json:"Datas,omitempty" name:"Datas" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *AlarmReportNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmReportNumRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmReportNumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmReportNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmReportNumResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmReportNumsRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 报警类型：这里值为batch_alarms
	Op *string `json:"Op,omitempty" name:"Op"`

	// 上报数据
	DataInfos []*AlarmReportNum `json:"DataInfos,omitempty" name:"DataInfos" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *AlarmReportNumsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmReportNumsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmReportNumsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmReportNumsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AlarmReportNumsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmStrategy struct {

	// 策略类型
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// CI项属性
	CiAttr *string `json:"CiAttr,omitempty" name:"CiAttr"`

	// 命令空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 告警指标类型
	AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警指标名
	AlarmName *string `json:"AlarmName,omitempty" name:"AlarmName"`

	// 数值判断方法
	JudgeMethod *int64 `json:"JudgeMethod,omitempty" name:"JudgeMethod"`

	// 告警恢复方式，0->主动恢复，1->被动恢复
	RecoverType *int64 `json:"RecoverType,omitempty" name:"RecoverType"`

	// 告警恢复时间
	RecoverTimes *int64 `json:"RecoverTimes,omitempty" name:"RecoverTimes"`

	// 数值判断类型
	JudgeValue *int64 `json:"JudgeValue,omitempty" name:"JudgeValue"`

	// 发生频率
	OccurrenceFreq *string `json:"OccurrenceFreq,omitempty" name:"OccurrenceFreq"`

	// 抑制类型
	InhibitType *int64 `json:"InhibitType,omitempty" name:"InhibitType"`

	// 告警通知人
	NoticePeople *string `json:"NoticePeople,omitempty" name:"NoticePeople"`

	// 告警级别
	AlarmLevel *string `json:"AlarmLevel,omitempty" name:"AlarmLevel"`

	// 无
	ToEcc *int64 `json:"ToEcc,omitempty" name:"ToEcc"`
}

type AlarmStrategyWithMd5 struct {

	// 告警策略MD5
	StrategyMd5 *string `json:"StrategyMd5,omitempty" name:"StrategyMd5"`

	// 告警策略
	Strategy *AlarmStrategy `json:"Strategy,omitempty" name:"Strategy"`
}

type AllocateServerIpRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：server or server_vm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：allocate_lan_ip或allocate_wan_ip
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Data *ServerIpData `json:"Data,omitempty" name:"Data"`
}

func (r *AllocateServerIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateServerIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateServerIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 查询返回
		Data *ServerDataWithResult `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateServerIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateServerIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSnmpTemplateRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：bind_snmp_template
	Op *string `json:"Op,omitempty" name:"Op"`

	// 模版配置
	Configs []*SnmpTempConfig `json:"Configs,omitempty" name:"Configs" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *BindSnmpTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSnmpTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSnmpTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 总数
		CountForAll *int64 `json:"CountForAll,omitempty" name:"CountForAll"`

		// 成功数
		CountForSucc *int64 `json:"CountForSucc,omitempty" name:"CountForSucc"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSnmpTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSnmpTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmStrategyRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：create_alarm_strategy
	Op *string `json:"Op,omitempty" name:"Op"`

	// 具体的告警策略
	Data *AlarmStrategy `json:"Data,omitempty" name:"Data"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *CreateAlarmStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 创建成功的策略MD5
		StrategyMd5 *string `json:"StrategyMd5,omitempty" name:"StrategyMd5"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelIdcLineExportItem struct {

	// 筛选idc专线
	Condition *IdcLineExport `json:"Condition,omitempty" name:"Condition"`
}

type DelIdcLineExportRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_export_line
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：delete
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要删除的idc专线出口信息
	Datas []*DelIdcLineExportItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *DelIdcLineExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelIdcLineExportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelIdcLineExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*IdcLineExportRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelIdcLineExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelIdcLineExportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelIdcLineItem struct {

	// 筛选idc专线
	Condition *IdcLine `json:"Condition,omitempty" name:"Condition"`
}

type DelIdcLineRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_special_line
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：delete
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要删除的idc专线信息
	Datas []*DelIdcLineItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *DelIdcLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelIdcLineRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelIdcLineResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*IdcLineRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelIdcLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelIdcLineResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelNetDeviceRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：netdevice
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：delete
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要删除网络设备信息
	Datas []*DelNetdeviceItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *DelNetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelNetDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelNetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*NetdeviceRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelNetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelNetDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelNetPortItem struct {

	// 筛选网络设备端口
	Condition *NetPortItem `json:"Condition,omitempty" name:"Condition"`
}

type DelNetPortRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：netdevice
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：delete
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要删除网络设备端口信息
	Datas []*DelNetPortItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *DelNetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelNetPortRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelNetPortResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result []*int64 `json:"Result,omitempty" name:"Result" list`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*NetPortRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelNetPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelNetPortResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelNetdeviceItem struct {

	// 筛选网络设备
	Condition *NetdeviceItem `json:"Condition,omitempty" name:"Condition"`
}

type DelRackDescItem struct {

	// 筛选idc机架机位配置
	Condition *RackDescItem `json:"Condition,omitempty" name:"Condition"`
}

type DelRackDescRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_rack
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：delete
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要删除的idc机架机位配置信息
	Datas []*DelRackDescItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *DelRackDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelRackDescRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelRackDescResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*RackDescRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelRackDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelRackDescResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelRackInfoRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idcinfo
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：delete
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改的机架机位信息
	Datas []*DelRackDescItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *DelRackInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelRackInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelRackInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*RackDescRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelRackInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelRackInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmStrategyMd5Request struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：delete_alarm_strategy
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改的策略MD5
	StrategyMd5 *string `json:"StrategyMd5,omitempty" name:"StrategyMd5"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *DeleteAlarmStrategyMd5Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmStrategyMd5Request) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmStrategyMd5Response struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmStrategyMd5Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmStrategyMd5Response) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmStrategyRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：delete_alarm_strategy
	Op *string `json:"Op,omitempty" name:"Op"`

	// 具体的告警策略
	Data *AlarmStrategy `json:"Data,omitempty" name:"Data"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *DeleteAlarmStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServerRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：server
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：delete
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Data *ServerData `json:"Data,omitempty" name:"Data"`
}

func (r *DeleteServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 查询返回
		Data *ServerDataWithResult `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployLogCond struct {

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 需要查询日志的SN列表，可为一个或多个
	SnList []*string `json:"SnList,omitempty" name:"SnList" list`
}

type DeployLogDetail struct {

	// 启动时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 部署步骤
	Step *int64 `json:"Step,omitempty" name:"Step"`

	// 成功失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 若失败输出详细信息
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type DeployLogInfo struct {

	// SN
	Sn *string `json:"Sn,omitempty" name:"Sn"`

	// 任务执行详情
	Logs []*DeployLogDetail `json:"Logs,omitempty" name:"Logs" list`
}

type FastReinstallServerRequest struct {
	*tchttp.BaseRequest

	// 用户,默认取值sid_dcos
	User *string `json:"User,omitempty" name:"User"`

	// 命令,这里取值fask_deploy_task
	Command *string `json:"Command,omitempty" name:"Command"`

	// 操作类型111部署
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`

	// 带外接管,默认为0
	OutbandFlag *string `json:"OutbandFlag,omitempty" name:"OutbandFlag"`

	// 操作数据
	Data *ServerData `json:"Data,omitempty" name:"Data"`

	// dcos底层服务按idc部署,操作类请求需要指定idcid,路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *FastReinstallServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FastReinstallServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FastReinstallServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 操作返回任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FastReinstallServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FastReinstallServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FieldFilter struct {

	// 需要过滤的字段
	Name *string `json:"Name,omitempty" name:"Name"`
}

type Filter struct {

	// 无
	Name *string `json:"Name,omitempty" name:"Name"`

	// 无
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type GetAggregationXflowDataRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：xflow
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：get_aggregation_xflow_data
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Lines []*XflowLineInfo `json:"Lines,omitempty" name:"Lines" list`

	// 范围过滤器
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`
}

func (r *GetAggregationXflowDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAggregationXflowDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAggregationXflowDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*XflowLineInfo `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAggregationXflowDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAggregationXflowDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOriginalXflowDataRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：xflow
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：get_original_xflow_data
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Lines []*XflowLineInfo `json:"Lines,omitempty" name:"Lines" list`

	// 范围过滤器
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`
}

func (r *GetOriginalXflowDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOriginalXflowDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOriginalXflowDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*XflowLineResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOriginalXflowDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOriginalXflowDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeriodAggregationXflowDataRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：xflow
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：get_period_aggregation_xflow_data
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Lines []*XflowLineInfo `json:"Lines,omitempty" name:"Lines" list`

	// 范围过滤器，结束日期不应该小于开始日期，且查询日期间隔不超过3个月
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`
}

func (r *GetPeriodAggregationXflowDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeriodAggregationXflowDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeriodAggregationXflowDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*XflowLineResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPeriodAggregationXflowDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeriodAggregationXflowDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopnXflowDataRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：xflow
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：get_topn_xflow_data
	Op *string `json:"Op,omitempty" name:"Op"`

	// topN排序条件
	Orders []*string `json:"Orders,omitempty" name:"Orders" list`

	// 查询条件
	Line *XflowLineInfo `json:"Line,omitempty" name:"Line"`

	// 范围过滤器，结束日期不应该小于开始日期，且查询日期间隔不超过3个月
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`

	// 返回Top多少，默认为10
	TopN *int64 `json:"TopN,omitempty" name:"TopN"`
}

func (r *GetTopnXflowDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopnXflowDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopnXflowDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*XflowLineInfo `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopnXflowDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopnXflowDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetXflowInfoRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：xflow
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：get_line_info
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetXflowInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetXflowInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetXflowInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*XflowLineInfo `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetXflowInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetXflowInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IdcLine struct {

	// 专线编号
	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`

	// 专线用途 1,DCI专线 2,ECN专线 3,IT专线 4...
	SpeLineFunc *int64 `json:"SpeLineFunc,omitempty" name:"SpeLineFunc"`

	// 索引ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 专线名称
	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`

	// 专线类型 1,DDN 2,SDH 3,FIBER 4,MSTP 5,VPN,6...
	SpeLineType *int64 `json:"SpeLineType,omitempty" name:"SpeLineType"`

	// 专线状态 1,待启用 2,运营中 3,已裁撤 4...
	SpeLineStatus *int64 `json:"SpeLineStatus,omitempty" name:"SpeLineStatus"`

	// 专线速率
	SpeLineSpeed *string `json:"SpeLineSpeed,omitempty" name:"SpeLineSpeed"`

	// 专线所属业务
	LineBusiness *string `json:"LineBusiness,omitempty" name:"LineBusiness"`

	// 专线启用时间
	LineStartTime *string `json:"LineStartTime,omitempty" name:"LineStartTime"`

	// 专线裁撤时间
	LineEndTime *string `json:"LineEndTime,omitempty" name:"LineEndTime"`

	// 本端设备
	LocalDevice *string `json:"LocalDevice,omitempty" name:"LocalDevice"`

	// 本端设备端口
	LocalDevicePort *string `json:"LocalDevicePort,omitempty" name:"LocalDevicePort"`

	// 本端设备IP
	LocalDeviceIp *string `json:"LocalDeviceIp,omitempty" name:"LocalDeviceIp"`

	// 对端设备IP
	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`

	// 运营商
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 运营商接口人
	LocalOperatorOwner *string `json:"LocalOperatorOwner,omitempty" name:"LocalOperatorOwner"`

	// 接口人电话
	LocalOperatorOwnerPhone *string `json:"LocalOperatorOwnerPhone,omitempty" name:"LocalOperatorOwnerPhone"`

	// 运营商报障电话
	LocalOperatorWarningPhone *string `json:"LocalOperatorWarningPhone,omitempty" name:"LocalOperatorWarningPhone"`

	// 对端接口人
	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`

	// 对端接口人电话
	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`

	// 对端保障电话
	RemoteOperatorWarningPhone *string `json:"RemoteOperatorWarningPhone,omitempty" name:"RemoteOperatorWarningPhone"`

	// 合作单位
	CooperateUnit *string `json:"CooperateUnit,omitempty" name:"CooperateUnit"`

	// 建设单位
	ConstructUnit *string `json:"ConstructUnit,omitempty" name:"ConstructUnit"`

	// 专线负责人
	LineOwner *string `json:"LineOwner,omitempty" name:"LineOwner"`

	// 专线距离
	LineDistance *string `json:"LineDistance,omitempty" name:"LineDistance"`

	// 本端二层设备
	LocalSecondDevice *string `json:"LocalSecondDevice,omitempty" name:"LocalSecondDevice"`

	// 本端二层设备端口
	LocalSecondDevicePort *string `json:"LocalSecondDevicePort,omitempty" name:"LocalSecondDevicePort"`

	// IP SLA
	IpSla *string `json:"IpSla,omitempty" name:"IpSla"`

	// SLA序号
	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`

	// 所属VLAN信息
	VlanInfo *string `json:"VlanInfo,omitempty" name:"VlanInfo"`

	// 主备情况
	BackupStatus *string `json:"BackupStatus,omitempty" name:"BackupStatus"`

	// 切换能力
	SwitchAbility *string `json:"SwitchAbility,omitempty" name:"SwitchAbility"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type IdcLineExport struct {

	// 出口编号
	ExportLineId *string `json:"ExportLineId,omitempty" name:"ExportLineId"`

	// 索引ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 出口名称
	ExportLineName *string `json:"ExportLineName,omitempty" name:"ExportLineName"`

	// 出口状态 1,待启用 2,运营中 3,已裁撤 4...
	ExportLineStatus *int64 `json:"ExportLineStatus,omitempty" name:"ExportLineStatus"`

	// 运营商
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 技术可用带宽
	LineTechAvailableSpeed *string `json:"LineTechAvailableSpeed,omitempty" name:"LineTechAvailableSpeed"`

	// 商务可用带宽
	LineBusiAvailableSpeed *string `json:"LineBusiAvailableSpeed,omitempty" name:"LineBusiAvailableSpeed"`

	// 专线启用时间
	LineStartTime *string `json:"LineStartTime,omitempty" name:"LineStartTime"`

	// 专线裁撤时间
	LineEndTime *string `json:"LineEndTime,omitempty" name:"LineEndTime"`

	// 出口设备
	ExportDevice *string `json:"ExportDevice,omitempty" name:"ExportDevice"`

	// 出口设备端口
	ExportDevicePort *string `json:"ExportDevicePort,omitempty" name:"ExportDevicePort"`

	// 本端设备IP
	ExportDeviceIp *string `json:"ExportDeviceIp,omitempty" name:"ExportDeviceIp"`

	// 对端设备IP
	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`

	// 对端联系人
	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`

	// 对端联系人电话
	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`

	// 对端报障电话
	RemoteOperatorWarningPhone *string `json:"RemoteOperatorWarningPhone,omitempty" name:"RemoteOperatorWarningPhone"`

	// SLA 序号
	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type IdcLineExportRespItem struct {

	// 索引ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 出口编号
	ExportLineId *string `json:"ExportLineId,omitempty" name:"ExportLineId"`

	// 出口名称
	ExportLineName *string `json:"ExportLineName,omitempty" name:"ExportLineName"`

	// 出口状态 1,待启用 2,运营中 3,已裁撤 4..
	ExportLineStatus *int64 `json:"ExportLineStatus,omitempty" name:"ExportLineStatus"`

	// 运营商
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 技术可用带宽
	LineTechAvailableSpeed *string `json:"LineTechAvailableSpeed,omitempty" name:"LineTechAvailableSpeed"`

	// 商务可用带宽
	LineBusiAvailableSpeed *string `json:"LineBusiAvailableSpeed,omitempty" name:"LineBusiAvailableSpeed"`

	// 专线启用时间
	LineStartTime *string `json:"LineStartTime,omitempty" name:"LineStartTime"`

	// 专线裁撤时间
	LineEndTime *string `json:"LineEndTime,omitempty" name:"LineEndTime"`

	// 出口设备
	ExportDevice *string `json:"ExportDevice,omitempty" name:"ExportDevice"`

	// 出口设备端口
	ExportDevicePort *string `json:"ExportDevicePort,omitempty" name:"ExportDevicePort"`

	// 本端设备IP
	ExportDeviceIp *string `json:"ExportDeviceIp,omitempty" name:"ExportDeviceIp"`

	// 对端设备IP
	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`

	// 对端联系人
	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`

	// 对端联系人电话
	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`

	// 对端报障电话
	RemoteOperatorWarningPhone *string `json:"RemoteOperatorWarningPhone,omitempty" name:"RemoteOperatorWarningPhone"`

	// SLA 序号
	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 返回码
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 返回信息
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type IdcLineRespItem struct {

	// 索引ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 专线编号
	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`

	// 专线名称
	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`

	// 专线类型 1,DDN 2,SDH 3,FIBER 4,MSTP 5,VPN,6...
	SpeLineType *int64 `json:"SpeLineType,omitempty" name:"SpeLineType"`

	// 专线状态 1,待启用 2,运营中 3,已裁撤 4...
	SpeLineStatus *int64 `json:"SpeLineStatus,omitempty" name:"SpeLineStatus"`

	// 专线用途 1,DCI专线 2,ECN专线 3,IT专线 4...
	SpeLineFunc *int64 `json:"SpeLineFunc,omitempty" name:"SpeLineFunc"`

	// 专线速率
	SpeLineSpeed *string `json:"SpeLineSpeed,omitempty" name:"SpeLineSpeed"`

	// 专线所属业务
	LineBusiness *string `json:"LineBusiness,omitempty" name:"LineBusiness"`

	// 专线启用时间
	LineStartTime *string `json:"LineStartTime,omitempty" name:"LineStartTime"`

	// 专线裁撤时间
	LineEndTime *string `json:"LineEndTime,omitempty" name:"LineEndTime"`

	// 本端设备
	LocalDevice *string `json:"LocalDevice,omitempty" name:"LocalDevice"`

	// 本端设备端口
	LocalDevicePort *string `json:"LocalDevicePort,omitempty" name:"LocalDevicePort"`

	// 本端设备IP
	LocalDeviceIp *string `json:"LocalDeviceIp,omitempty" name:"LocalDeviceIp"`

	// 对端设备IP
	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`

	// 运营商
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 运营商接口人
	LocalOperatorOwner *string `json:"LocalOperatorOwner,omitempty" name:"LocalOperatorOwner"`

	// 接口人电话
	LocalOperatorOwnerPhone *string `json:"LocalOperatorOwnerPhone,omitempty" name:"LocalOperatorOwnerPhone"`

	// 运营商报障电话
	LocalOperatorWarningPhone *string `json:"LocalOperatorWarningPhone,omitempty" name:"LocalOperatorWarningPhone"`

	// 对端接口人
	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`

	// 对端接口人电话
	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`

	// 对端保障电话
	RemoteOperatorWarningPhone *string `json:"RemoteOperatorWarningPhone,omitempty" name:"RemoteOperatorWarningPhone"`

	// 合作单位
	CooperateUnit *string `json:"CooperateUnit,omitempty" name:"CooperateUnit"`

	// 建设单位
	ConstructUnit *string `json:"ConstructUnit,omitempty" name:"ConstructUnit"`

	// 专线负责人
	LineOwner *string `json:"LineOwner,omitempty" name:"LineOwner"`

	// 专线距离
	LineDistance *string `json:"LineDistance,omitempty" name:"LineDistance"`

	// 本端二层设备
	LocalSecondDevice *string `json:"LocalSecondDevice,omitempty" name:"LocalSecondDevice"`

	// 本端二层设备端口
	LocalSecondDevicePort *string `json:"LocalSecondDevicePort,omitempty" name:"LocalSecondDevicePort"`

	// IP SLA
	IpSla *string `json:"IpSla,omitempty" name:"IpSla"`

	// SLA序号
	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`

	// 所属VLAN信息
	VlanInfo *string `json:"VlanInfo,omitempty" name:"VlanInfo"`

	// 主备情况
	BackupStatus *string `json:"BackupStatus,omitempty" name:"BackupStatus"`

	// 切换能力
	SwitchAbility *string `json:"SwitchAbility,omitempty" name:"SwitchAbility"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 返回码
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 返回信息
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type IdcRackDesc struct {

	// 机房ID
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`

	// 无
	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`

	// IDC名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// 机架高度
	RackHeight *string `json:"RackHeight,omitempty" name:"RackHeight"`

	// 机架是否启用
	RackInUse *string `json:"RackInUse,omitempty" name:"RackInUse"`

	// 机架是否开电
	RackOpen *string `json:"RackOpen,omitempty" name:"RackOpen"`

	// 机架峰值功率
	RackMaxPower *string `json:"RackMaxPower,omitempty" name:"RackMaxPower"`

	// 机架开电时间
	RackPoweronTime *string `json:"RackPoweronTime,omitempty" name:"RackPoweronTime"`

	// 机架机位总数
	RackPosNum *string `json:"RackPosNum,omitempty" name:"RackPosNum"`

	// 机架名称
	RackName *string `json:"RackName,omitempty" name:"RackName"`

	// 机架启用时间
	RackUseTime *string `json:"RackUseTime,omitempty" name:"RackUseTime"`
}

type IdcRackInfo struct {

	// IDC名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// 存放机位ID
	DevicePosId *string `json:"DevicePosId,omitempty" name:"DevicePosId"`

	// 设备高度
	DeviceHeight *string `json:"DeviceHeight,omitempty" name:"DeviceHeight"`

	// 索引ID
	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`

	// 机架名称
	DeviceRackName *string `json:"DeviceRackName,omitempty" name:"DeviceRackName"`

	// 机位逻辑区域
	PosLogicArea *string `json:"PosLogicArea,omitempty" name:"PosLogicArea"`

	// 机位内网交换机端口
	PosInnerSwitchPortName *string `json:"PosInnerSwitchPortName,omitempty" name:"PosInnerSwitchPortName"`

	// 机位外网交换机端口
	PosOuterSwitchPortName *string `json:"PosOuterSwitchPortName,omitempty" name:"PosOuterSwitchPortName"`

	// 机位管理网交换机端口
	PosAdminSwitchPortName *string `json:"PosAdminSwitchPortName,omitempty" name:"PosAdminSwitchPortName"`

	// 上架设备类型 0无 1服务器 2网络设备
	OnlineDeviceType *string `json:"OnlineDeviceType,omitempty" name:"OnlineDeviceType"`

	// 上架设备固资
	OnlineDeviceAssetId *string `json:"OnlineDeviceAssetId,omitempty" name:"OnlineDeviceAssetId"`

	// IDCID
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

type IpData struct {

	// IP地址列表
	IpInfos []*IpInfo `json:"IpInfos,omitempty" name:"IpInfos" list`
}

type IpInfo struct {

	// ip信息
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// IP状态ID  0未分配 1已分配
	IpStatusId *int64 `json:"IpStatusId,omitempty" name:"IpStatusId"`

	// 网段名称
	NetSegmentName *string `json:"NetSegmentName,omitempty" name:"NetSegmentName"`

	// 网段掩码
	NetSegmentMask *string `json:"NetSegmentMask,omitempty" name:"NetSegmentMask"`

	// 网段网关
	NetSegmentGateway *string `json:"NetSegmentGateway,omitempty" name:"NetSegmentGateway"`

	// IP类型  0网络设备IP 1服务器IP
	IpType *int64 `json:"IpType,omitempty" name:"IpType"`

	// 关联固资编号
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 网段ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type KeyValue struct {

	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogEntry struct {

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 日志
	Raw *string `json:"Raw,omitempty" name:"Raw"`
}

type ModServerResultEntry struct {

	// 结果 0成功，1失败
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 详情
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 固资ID
	AssertId *string `json:"AssertId,omitempty" name:"AssertId"`
}

type ModifyAlarmStrategyMd5Request struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：modify_alarm_strategy
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改的策略MD5
	StrategyMd5 *string `json:"StrategyMd5,omitempty" name:"StrategyMd5"`

	// 具体的告警策略
	Data *AlarmStrategy `json:"Data,omitempty" name:"Data"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *ModifyAlarmStrategyMd5Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmStrategyMd5Request) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmStrategyMd5Response struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmStrategyMd5Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmStrategyMd5Response) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmStrategyRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：modify_alarm_strategy
	Op *string `json:"Op,omitempty" name:"Op"`

	// 具体的告警策略
	Data *AlarmStrategy `json:"Data,omitempty" name:"Data"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *ModifyAlarmStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIdcLineExportRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_export_line
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：modify
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改的专线信息
	Datas []*ModifyIdcLineReq `json:"Datas,omitempty" name:"Datas" list`
}

func (r *ModifyIdcLineExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIdcLineExportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIdcLineExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*ModifyIdcLineResp `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIdcLineExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIdcLineExportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIdcLineReq struct {

	// 专线ID
	LineId *string `json:"LineId,omitempty" name:"LineId"`

	// 修改字段
	Modifys []*KeyValue `json:"Modifys,omitempty" name:"Modifys" list`
}

type ModifyIdcLineRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_special_line
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：modify
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改的专线信息
	Datas []*ModifyIdcLineReq `json:"Datas,omitempty" name:"Datas" list`
}

func (r *ModifyIdcLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIdcLineRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIdcLineResp struct {

	// 结果，0成功，1失败
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 详情
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 专线ID
	LineId *string `json:"LineId,omitempty" name:"LineId"`
}

type ModifyIdcLineResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*ModifyIdcLineResp `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIdcLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIdcLineResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDeviceRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：netdevice
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：modify
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改网络设备信息
	Datas *DelRackDescItem `json:"Datas,omitempty" name:"Datas"`
}

func (r *ModifyNetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*NetdeviceRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetPortItem struct {

	// 筛选网络设备端口
	Condition *NetPortItem `json:"Condition,omitempty" name:"Condition"`

	// 修改网络设备端口
	Modify *NetPortItem `json:"Modify,omitempty" name:"Modify"`
}

type ModifyNetPortRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：netdevice_related
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：modify
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改网络设备端口信息
	Datas []*ModifyNetPortItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *ModifyNetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetPortRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetPortResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*NetPortRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetPortResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRackDescItem struct {

	// 筛选idc机架机位配置
	Condition *RackDescItem `json:"Condition,omitempty" name:"Condition"`

	// 修改idc机架机位配置
	Modify *RackDescItem `json:"Modify,omitempty" name:"Modify"`
}

type ModifyRackDescRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_rack
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：modify
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改的idc机架机位配置信息
	Datas []*ModifyRackDescItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *ModifyRackDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRackDescRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRackDescResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*RackDescRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRackDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRackDescResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRackInfoRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idcinfo
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 默认：sid_test
	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`

	// 操作名称，本接口取值：modify
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要修改的机架机位信息
	Datas []*DelRackDescItem `json:"Datas,omitempty" name:"Datas" list`
}

func (r *ModifyRackInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRackInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRackInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet []*RackDescRespItem `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRackInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRackInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServerRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：server
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：modify
	Op *string `json:"Op,omitempty" name:"Op"`

	// 操作服务器数据
	Datas []*ServerModifyData `json:"Datas,omitempty" name:"Datas" list`
}

func (r *ModifyServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		DataSet *ModifyServerResult `json:"DataSet,omitempty" name:"DataSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServerResult struct {

	// 服务信息列表
	Servers []*ModServerResultEntry `json:"Servers,omitempty" name:"Servers" list`
}

type NetDeviceData struct {

	// 网络设备列表
	NetDevices []*NetDeviceInfo `json:"NetDevices,omitempty" name:"NetDevices" list`
}

type NetDeviceInfo struct {

	// 固资ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// SN编号
	Sn *string `json:"Sn,omitempty" name:"Sn"`

	// 网络设备名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 设备类型(交换机,路由器,防火墙)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 设备用途(内网接入交换机,外网接入交换机,核心交换机,防火墙)
	Func *string `json:"Func,omitempty" name:"Func"`

	// 设备厂家
	Pro *string `json:"Pro,omitempty" name:"Pro"`

	// 设备型号
	Model *string `json:"Model,omitempty" name:"Model"`

	// 设备IDC
	Idc *string `json:"Idc,omitempty" name:"Idc"`

	// 机架名称
	RackName *string `json:"RackName,omitempty" name:"RackName"`

	// 管理ip
	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`

	// 其他ip(多个以”,”隔开)
	OtherIp *string `json:"OtherIp,omitempty" name:"OtherIp"`

	// 设备ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 网络设备引擎
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 设备OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// 设备IDC id(暂未使用该字段)
	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`

	// 设备当前状态(0待运营 1运营中 2已下架)
	CurrentStatus *int64 `json:"CurrentStatus,omitempty" name:"CurrentStatus"`

	// 物理机网段名称
	PmNetsegmentName *string `json:"PmNetsegmentName,omitempty" name:"PmNetsegmentName"`

	// 虚拟机网段名称
	VmNetSegmentName *string `json:"VmNetSegmentName,omitempty" name:"VmNetSegmentName"`

	// 导入时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
}

type NetPortData struct {

	// 网络端口列表
	NetDevices []*NetPortInfo `json:"NetDevices,omitempty" name:"NetDevices" list`
}

type NetPortInfo struct {

	// 网络设备固资编号
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 网络设备名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 端口名称
	PortName *string `json:"PortName,omitempty" name:"PortName"`

	// 端口速率
	PortSpeed *string `json:"PortSpeed,omitempty" name:"PortSpeed"`

	// IDC名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// 连接类型 0连接网络设备 1连接服务器 -1未连接
	ConnectionType *int64 `json:"ConnectionType,omitempty" name:"ConnectionType"`

	// 服务器ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// IDC ID 暂不使用该字段
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`

	// 机架名称
	DeviceRackName *string `json:"DeviceRackName,omitempty" name:"DeviceRackName"`

	// 存放机位ID
	DevicePosId *int64 `json:"DevicePosId,omitempty" name:"DevicePosId"`

	// 连接网络设备固资编号
	ConnectionAssertId *string `json:"ConnectionAssertId,omitempty" name:"ConnectionAssertId"`

	// 连接网络设备名称
	ConnectionName *string `json:"ConnectionName,omitempty" name:"ConnectionName"`

	// 连接网络设备端口名称
	ConnectionPort *string `json:"ConnectionPort,omitempty" name:"ConnectionPort"`

	// 连接网络设备端口ip
	ConnectionIp *string `json:"ConnectionIp,omitempty" name:"ConnectionIp"`

	// 连接服务器固资编号
	ConnectionServerAssertId *string `json:"ConnectionServerAssertId,omitempty" name:"ConnectionServerAssertId"`
}

type NetPortItem struct {

	// 索引ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 网络设备固资编号
	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`

	// 端口名称
	PortName *string `json:"PortName,omitempty" name:"PortName"`

	// 连接网络设备固资编号
	ConnectionNetdevAssetId *string `json:"ConnectionNetdevAssetId,omitempty" name:"ConnectionNetdevAssetId"`

	// 连接网络设备端口名称
	ConnectionNetdevPort *string `json:"ConnectionNetdevPort,omitempty" name:"ConnectionNetdevPort"`
}

type NetPortRespItem struct {

	// 索引ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 网络设备固资编号
	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`

	// 端口名称
	PortName *string `json:"PortName,omitempty" name:"PortName"`

	// 连接网络设备固资编号
	ConnectionNetdevAssetId *string `json:"ConnectionNetdevAssetId,omitempty" name:"ConnectionNetdevAssetId"`

	// 连接网络设备端口名称
	ConnectionNetdevPort *string `json:"ConnectionNetdevPort,omitempty" name:"ConnectionNetdevPort"`

	// 返回码
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 返回信息
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type NetSegmentData struct {

	// IP网段列表
	NetSegmentInfos []*NetSegmentInfo `json:"NetSegmentInfos,omitempty" name:"NetSegmentInfos" list`
}

type NetSegmentInfo struct {

	// 网段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网段掩码
	Mask *string `json:"Mask,omitempty" name:"Mask"`

	// 网段网关
	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`

	// IP资源池
	IpPool *string `json:"IpPool,omitempty" name:"IpPool"`

	// IDC名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// IDC ID(暂未使用)
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`

	// 网段分类 0物理机内网IP 1虚拟机内网IP 2TGW外网IP 3网络设备互联IP
	Type *string `json:"Type,omitempty" name:"Type"`

	// 网段ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 覆盖交换机
	RelatedNetDevices *string `json:"RelatedNetDevices,omitempty" name:"RelatedNetDevices"`

	// 逻辑区域
	LogicArea *string `json:"LogicArea,omitempty" name:"LogicArea"`

	// 网段分类ID
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`

	// vlan id
	VlanId *string `json:"VlanId,omitempty" name:"VlanId"`
}

type NetdeviceItem struct {

	// 固资号
	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`

	// SN号
	NetdevSn *string `json:"NetdevSn,omitempty" name:"NetdevSn"`

	// 设备引擎
	NetdevEngine *string `json:"NetdevEngine,omitempty" name:"NetdevEngine"`

	// 设备名
	NetdevName *string `json:"NetdevName,omitempty" name:"NetdevName"`

	// 设备类型(交换机,路由器,防火墙)
	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`

	// 设备用途(内网接入交换机,外网接入交换机,核心交换机,防火墙)
	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`

	// 设备厂家
	NetdevPro *string `json:"NetdevPro,omitempty" name:"NetdevPro"`

	// 设备型号
	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`

	// 设备OS
	NetdevOs *string `json:"NetdevOs,omitempty" name:"NetdevOs"`

	// 设备IDC
	NetdevIdc *string `json:"NetdevIdc,omitempty" name:"NetdevIdc"`

	// 设备IDC id
	NetdevIdcId *string `json:"NetdevIdcId,omitempty" name:"NetdevIdcId"`

	// 机架名称
	NetdevRackName *string `json:"NetdevRackName,omitempty" name:"NetdevRackName"`

	// 设备当前状态(0待运营 1运营中 2已下架)
	NetdevCurrentStatus *int64 `json:"NetdevCurrentStatus,omitempty" name:"NetdevCurrentStatus"`

	// 管理ip
	NetdevAdminIp *string `json:"NetdevAdminIp,omitempty" name:"NetdevAdminIp"`

	// 其他ip(多个)
	NetdevOtherIp *string `json:"NetdevOtherIp,omitempty" name:"NetdevOtherIp"`

	// 网络设备维护人
	NetdevMaintainer *string `json:"NetdevMaintainer,omitempty" name:"NetdevMaintainer"`

	// appid
	Appid *string `json:"Appid,omitempty" name:"Appid"`

	// 地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 网络设备id
	NetdevId *int64 `json:"NetdevId,omitempty" name:"NetdevId"`
}

type NetdeviceRespItem struct {

	// 索引ID
	NetdevId *int64 `json:"NetdevId,omitempty" name:"NetdevId"`

	// 固资号
	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`

	// SN号
	NetdevSn *string `json:"NetdevSn,omitempty" name:"NetdevSn"`

	// 设备引擎
	NetdevEngine *string `json:"NetdevEngine,omitempty" name:"NetdevEngine"`

	// 设备名
	NetdevName *string `json:"NetdevName,omitempty" name:"NetdevName"`

	// 设备类型(交换机,路由器,防火墙)
	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`

	// 设备用途(内网接入交换机,外网接入交换机,核心交换机,防火墙)
	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`

	// 设备厂家
	NetdevPro *string `json:"NetdevPro,omitempty" name:"NetdevPro"`

	// 设备型号
	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`

	// 设备OS
	NetdevOs *string `json:"NetdevOs,omitempty" name:"NetdevOs"`

	// 设备IDC
	NetdevIdc *string `json:"NetdevIdc,omitempty" name:"NetdevIdc"`

	// 设备IDC id
	NetdevIdcId *string `json:"NetdevIdcId,omitempty" name:"NetdevIdcId"`

	// 机架名称
	NetdevRackName *string `json:"NetdevRackName,omitempty" name:"NetdevRackName"`

	// 设备当前状态(0待运营 1运营中 2已下架)
	NetdevCurrentStatus *int64 `json:"NetdevCurrentStatus,omitempty" name:"NetdevCurrentStatus"`

	// 管理ip
	NetdevAdminIp *string `json:"NetdevAdminIp,omitempty" name:"NetdevAdminIp"`

	// 其他ip(多个)
	NetdevOtherIp *string `json:"NetdevOtherIp,omitempty" name:"NetdevOtherIp"`

	// 网络设备维护人
	NetdevMaintainer *string `json:"NetdevMaintainer,omitempty" name:"NetdevMaintainer"`

	// 网络设备信息更新时间
	NetdevUpdateTime *string `json:"NetdevUpdateTime,omitempty" name:"NetdevUpdateTime"`

	// appid
	Appid *string `json:"Appid,omitempty" name:"Appid"`

	// 地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 返回码
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 返回信息
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type NumKeyValue struct {

	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type OperateServerRequest struct {
	*tchttp.BaseRequest

	// 用户默认取值sid_dcos
	User *string `json:"User,omitempty" name:"User"`

	// 命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 操作类型4 一致性检查1 重启2 关机3开机16初始化密码修改19开机pxe
	OperateType *int64 `json:"OperateType,omitempty" name:"OperateType"`

	// 操作数据
	Data *ServerData `json:"Data,omitempty" name:"Data"`

	// dcos底层服务按idc部署,操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *OperateServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperateServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperateServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 操作返回任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperateServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperateServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OutbandTaskInfo struct {

	// SN
	Sn *string `json:"Sn,omitempty" name:"Sn"`

	// 服务器的执行状态。 //0 -> 未执行，1 -> 执行成功 //2 -> 执行失败，3 -> 执行中
	State *int64 `json:"State,omitempty" name:"State"`

	// 失败时输出详细信息
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type OutbandUserInfo struct {

	// SN固资编号
	Sn *string `json:"Sn,omitempty" name:"Sn"`

	// 带外dhcp_ip
	DhcpIp *string `json:"DhcpIp,omitempty" name:"DhcpIp"`

	// 带外用户名
	User *string `json:"User,omitempty" name:"User"`

	// 带外密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type QueryAlarmRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：query_alarm_record
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 范围过滤器
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QueryAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAlarmRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*AlarmReportInfo `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAlarmResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryAlarmStrategysRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：alarm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：query_alarm_strategy
	Op *string `json:"Op,omitempty" name:"Op"`

	// 需要查询的策略MD5
	StrategyMd5s []*string `json:"StrategyMd5s,omitempty" name:"StrategyMd5s" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QueryAlarmStrategysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAlarmStrategysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryAlarmStrategysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回策略列表
		DataSet []*AlarmStrategyWithMd5 `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAlarmStrategysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAlarmStrategysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDeployLogRequest struct {
	*tchttp.BaseRequest

	// 用户，默认取值sid_dcos
	User *string `json:"User,omitempty" name:"User"`

	// 命令，这里取值query_deploy_log
	Command *string `json:"Command,omitempty" name:"Command"`

	// 查询条件
	Data *DeployLogCond `json:"Data,omitempty" name:"Data"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QueryDeployLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDeployLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDeployLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 部署任务执行详情
		Data []*DeployLogInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDeployLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDeployLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryIdcLineExportRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_export_line
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryIdcLineExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryIdcLineExportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryIdcLineExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*QueryIdcLineExportResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryIdcLineExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryIdcLineExportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryIdcLineExportResult struct {

	// IDC专线出口列表
	LineExportInfos []*IdcLineExport `json:"LineExportInfos,omitempty" name:"LineExportInfos" list`
}

type QueryIdcLineRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_special_line
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryIdcLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryIdcLineRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryIdcLineResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*QueryIdcLineResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryIdcLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryIdcLineResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryIdcLineResult struct {

	// IDC专线列表
	LineInfos []*IdcLine `json:"LineInfos,omitempty" name:"LineInfos" list`
}

type QueryIpRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：ipdetail
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		Data *IpData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryNetDeviceRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：netdevice
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryNetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryNetDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryNetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		Data *NetDeviceData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryNetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryNetDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryNetPortRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：netdevice_related
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryNetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryNetPortRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryNetPortResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		Data *NetPortData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryNetPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryNetPortResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryNetSegmentRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：ipresource
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryNetSegmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryNetSegmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryNetSegmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		Data *NetSegmentData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryNetSegmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryNetSegmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryOutbandInfoRequest struct {
	*tchttp.BaseRequest

	// 用户，默认取值sid_dcos
	User *string `json:"User,omitempty" name:"User"`

	// 命令，这里取值query_outband_info
	Command *string `json:"Command,omitempty" name:"Command"`

	// 查询关键字，由固资编号sn组成的数组
	Datas []*string `json:"Datas,omitempty" name:"Datas" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QueryOutbandInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryOutbandInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryOutbandInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 详细用户密码信息
		DataSet []*OutbandUserInfo `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOutbandInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryOutbandInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryOutbandTaskRequest struct {
	*tchttp.BaseRequest

	// 用户
	User *string `json:"User,omitempty" name:"User"`

	// 命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 任务ID字典
	Data *NumKeyValue `json:"Data,omitempty" name:"Data"`
}

func (r *QueryOutbandTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryOutbandTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryOutbandTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 任务状态，1 -> 创建 2 -> 执行中 3 -> 完成
		State *int64 `json:"State,omitempty" name:"State"`

		// 执行详情
		Data *OutbandTaskInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOutbandTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryOutbandTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryRackDescRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idc_rack
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryRackDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryRackDescRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryRackDescResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*QueryRackDescResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryRackDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryRackDescResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryRackDescResult struct {

	// IDC信息列表
	RackDescs []*IdcRackDesc `json:"RackDescs,omitempty" name:"RackDescs" list`
}

type QueryRackInfoRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：idcinfo
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryRackInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryRackInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryRackInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		DataSet []*QueryRackInfoResult `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryRackInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryRackInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryRackInfoResult struct {

	// IDC信息列表
	IdcInfos []*IdcRackInfo `json:"IdcInfos,omitempty" name:"IdcInfos" list`
}

type QueryServerRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：server
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回字段过滤器
	FieldFilters []*FieldFilter `json:"FieldFilters,omitempty" name:"FieldFilters" list`
}

func (r *QueryServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回个数
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 符合当前查询条件的总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		Data *ServerData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpDataRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：query_snmp_data
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 是否需要总数
	NeedTotal *int64 `json:"NeedTotal,omitempty" name:"NeedTotal"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QuerySnmpDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 查询的日期
		Date *string `json:"Date,omitempty" name:"Date"`

		// 设备ID
		DevId *int64 `json:"DevId,omitempty" name:"DevId"`

		// 固资ID
		AssetId *int64 `json:"AssetId,omitempty" name:"AssetId"`

		// 0->查询成功且有数据，1->查询成功且无数据，2->查询失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 总数
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 分页索引
		StartIndex *int64 `json:"StartIndex,omitempty" name:"StartIndex"`

		// 每页返回的端口数
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// 数据
		DataSet []*SnmpQueryDataEntry `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySnmpDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpDetailRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：query_snmp_detail
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 范围查询
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QuerySnmpDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		Datas []*SnmpEntry `json:"Datas,omitempty" name:"Datas" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySnmpDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpDeviceByTemplateRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：query_devide_type_by_template
	Op *string `json:"Op,omitempty" name:"Op"`

	// 模版名列表
	Templates []*string `json:"Templates,omitempty" name:"Templates" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QuerySnmpDeviceByTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpDeviceByTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpDeviceByTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		TemplateSet []*SnmpTempConfig `json:"TemplateSet,omitempty" name:"TemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySnmpDeviceByTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpDeviceByTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpExportDataRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：query_export_data
	Op *string `json:"Op,omitempty" name:"Op"`

	// 范围条件
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QuerySnmpExportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpExportDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpExportDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 查询的日期
		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

		// 查询的日期
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// 专线ID
		LineId *string `json:"LineId,omitempty" name:"LineId"`

		// 返回结果
		DataSet []*SnmpEntry `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySnmpExportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpExportDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpLineDataRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：query_line_data
	Op *string `json:"Op,omitempty" name:"Op"`

	// 范围条件
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QuerySnmpLineDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpLineDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpLineDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 查询的日期
		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

		// 查询的日期
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// 专线ID
		LineId *string `json:"LineId,omitempty" name:"LineId"`

		// 返回结果
		DataSet []*SnmpEntry `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySnmpLineDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpLineDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpTemplateByDeviceRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：query_template_type_by_device_type
	Op *string `json:"Op,omitempty" name:"Op"`

	// 模版名列表
	Devices []*SnmpTempConfig `json:"Devices,omitempty" name:"Devices" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *QuerySnmpTemplateByDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpTemplateByDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySnmpTemplateByDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回结果
		TemplateSet []*SnmpTempConfig `json:"TemplateSet,omitempty" name:"TemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySnmpTemplateByDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySnmpTemplateByDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RackDescItem struct {

	// 索引ID
	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`

	// IDC名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// 机房ID
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`

	// 机架高度
	RackHeight *string `json:"RackHeight,omitempty" name:"RackHeight"`

	// 机架是否启用
	RackInUse *string `json:"RackInUse,omitempty" name:"RackInUse"`

	// 机架是否开电
	RackOpen *string `json:"RackOpen,omitempty" name:"RackOpen"`

	// 机架峰值功率
	RackMaxPower *string `json:"RackMaxPower,omitempty" name:"RackMaxPower"`

	// 机架开电时间
	RackPoweronTime *string `json:"RackPoweronTime,omitempty" name:"RackPoweronTime"`

	// 机架机位总数
	RackPosNum *string `json:"RackPosNum,omitempty" name:"RackPosNum"`

	// 机架名称
	RackName *string `json:"RackName,omitempty" name:"RackName"`

	// 机架启用时间
	RackUseTime *string `json:"RackUseTime,omitempty" name:"RackUseTime"`
}

type RackDescRespItem struct {

	// 索引ID
	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`

	// IDC名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// 机房ID
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`

	// 机架高度
	RackHeight *string `json:"RackHeight,omitempty" name:"RackHeight"`

	// 机架是否启用
	RackInUse *string `json:"RackInUse,omitempty" name:"RackInUse"`

	// 机架是否开电
	RackOpen *string `json:"RackOpen,omitempty" name:"RackOpen"`

	// 机架峰值功率
	RackMaxPower *string `json:"RackMaxPower,omitempty" name:"RackMaxPower"`

	// 机架开电时间
	RackPoweronTime *string `json:"RackPoweronTime,omitempty" name:"RackPoweronTime"`

	// 机架机位总数
	RackPosNum *string `json:"RackPosNum,omitempty" name:"RackPosNum"`

	// 机架名称
	RackName *string `json:"RackName,omitempty" name:"RackName"`

	// 机架启用时间
	RackUseTime *string `json:"RackUseTime,omitempty" name:"RackUseTime"`

	// 返回码
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 返回信息
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type RangeFilter struct {

	// 需要过滤的字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的起始值
	BeginValue *string `json:"BeginValue,omitempty" name:"BeginValue"`

	// 字段的结束值
	EndValue *string `json:"EndValue,omitempty" name:"EndValue"`
}

type RecycleServerIpRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：server or server_vm
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：recycle_lan_ip或recycle_wan_ip
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Data *ServerData `json:"Data,omitempty" name:"Data"`
}

func (r *RecycleServerIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecycleServerIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecycleServerIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 查询返回
		Data *ServerDataWithResult `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleServerIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecycleServerIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReinstallServerRequest struct {
	*tchttp.BaseRequest

	// 用户，默认取值sid_dcos
	User *string `json:"User,omitempty" name:"User"`

	// 命令，这里取值deploy_task
	Command *string `json:"Command,omitempty" name:"Command"`

	// 操作数据
	Data *ServerData `json:"Data,omitempty" name:"Data"`

	// 操作类型111部署
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`

	// 带外接管,默认为0
	OutbandFlag *string `json:"OutbandFlag,omitempty" name:"OutbandFlag"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *ReinstallServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReinstallServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReinstallServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 操作返回任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReinstallServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReinstallServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RelocateQueryResult struct {

	// 列表
	Servers *int64 `json:"Servers,omitempty" name:"Servers"`
}

type RelocateServerFinishRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：server
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：relocate_finish
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Data *ServerData `json:"Data,omitempty" name:"Data"`
}

func (r *RelocateServerFinishRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RelocateServerFinishRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RelocateServerFinishResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 查询返回
		Data *ServerDataWithResult `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerFinishResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RelocateServerFinishResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RelocateServerQueryRequest struct {
	*tchttp.BaseRequest

	// 公共参数,本接口取值server
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称,本接口取值relocate_query
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 根据需要填充需要返回的字段名,不同字段名使用,隔开
	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
}

func (r *RelocateServerQueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RelocateServerQueryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RelocateServerQueryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 返回数量
		ReturnNum *string `json:"ReturnNum,omitempty" name:"ReturnNum"`

		// 如果查询条件是固资，返回固资个数内的信息；否则返回查询条件下的server总数
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询返回
		Data *RelocateQueryResult `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerQueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RelocateServerQueryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RelocateServerRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：server
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 操作名称，本接口取值：relocate_start
	Op *string `json:"Op,omitempty" name:"Op"`

	// 查询条件
	Data *ServerData `json:"Data,omitempty" name:"Data"`
}

func (r *RelocateServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RelocateServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RelocateServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，1失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细说明
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 查询返回
		Data *ServerDataWithResult `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RelocateServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ServerAllocateIpInfo struct {

	// 服务器固资编号
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// Vlan号
	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`

	// 申请的IP
	IpWanted *string `json:"IpWanted,omitempty" name:"IpWanted"`
}

type ServerData struct {

	// 服务信息列表
	Servers []*ServerInfo `json:"Servers,omitempty" name:"Servers" list`
}

type ServerDataWithResult struct {

	// 服务信息列表
	Servers []*ServerInfoWithResult `json:"Servers,omitempty" name:"Servers" list`
}

type ServerInfo struct {

	// 服务器固资编号
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 服务器SN编号或虚拟机UUID编号
	Sn *string `json:"Sn,omitempty" name:"Sn"`

	// 设备型号
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// IDC名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// 机架名称
	RackName *string `json:"RackName,omitempty" name:"RackName"`

	// 存放机位
	PosId *int64 `json:"PosId,omitempty" name:"PosId"`

	// DeviceHeight
	DeviceHeight *string `json:"DeviceHeight,omitempty" name:"DeviceHeight"`

	// 服务器厂商
	Producer *string `json:"Producer,omitempty" name:"Producer"`

	// 服务器ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 设备别名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 服务器类型，0物理机，1虚拟机
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 服务器当前状态
	CurrentStatus *int64 `json:"CurrentStatus,omitempty" name:"CurrentStatus"`

	// 暂未使用
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`

	// 实体机当前最大虚拟机创建序号，初始值为0，虚拟机此字段无意义
	VmIndex *int64 `json:"VmIndex,omitempty" name:"VmIndex"`

	// 继承母机对应固资号
	SucceedAssetId *string `json:"SucceedAssetId,omitempty" name:"SucceedAssetId"`

	// 逻辑区域
	LogicErea *string `json:"LogicErea,omitempty" name:"LogicErea"`

	// 服务器硬件描述
	DeviceDescript *string `json:"DeviceDescript,omitempty" name:"DeviceDescript"`

	// 内网IP
	Lans []*ServerIpInfo `json:"Lans,omitempty" name:"Lans" list`

	// 外网IP
	Wans []*ServerIpInfo `json:"Wans,omitempty" name:"Wans" list`

	// 外网EIP
	WanEip *string `json:"WanEip,omitempty" name:"WanEip"`

	// RAID类型
	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`

	// 服务器操作系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 服务器agent版本
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`

	// 服务器所属业务
	Bussiness *string `json:"Bussiness,omitempty" name:"Bussiness"`

	// 服务器所属业务ID
	BussinessId *string `json:"BussinessId,omitempty" name:"BussinessId"`

	// 服务器启用时间
	FirstUseTime *string `json:"FirstUseTime,omitempty" name:"FirstUseTime"`

	// 服务器导入时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 服务器关联内网交换机固资(多台)，多台之间使用”,”隔开
	LanSwitchAssetId *string `json:"LanSwitchAssetId,omitempty" name:"LanSwitchAssetId"`

	// 服务器关联外网交换机固资(多台) ，多台之间使用”,”隔开
	WanSwitchAssetId *string `json:"WanSwitchAssetId,omitempty" name:"WanSwitchAssetId"`

	// 服务器关联管理网交换机固资(带外交换机)
	AdminSwitchAssetId *string `json:"AdminSwitchAssetId,omitempty" name:"AdminSwitchAssetId"`

	// 搬迁目的机房名称
	SvrRelocateIdcName *string `json:"SvrRelocateIdcName,omitempty" name:"SvrRelocateIdcName"`

	// 搬迁目的机房机架名称
	SvrRelocateRackName *string `json:"SvrRelocateRackName,omitempty" name:"SvrRelocateRackName"`

	// 搬迁目的机房机架机位
	SvrRelocatePosId *string `json:"SvrRelocatePosId,omitempty" name:"SvrRelocatePosId"`

	// 搬迁目的机房机架机位对应逻辑区域
	SvrRelocateLogicArea *string `json:"SvrRelocateLogicArea,omitempty" name:"SvrRelocateLogicArea"`

	// bonding flag
	BondingFlag *string `json:"BondingFlag,omitempty" name:"BondingFlag"`

	// vlan
	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
}

type ServerInfoWithResult struct {

	// 服务器信息
	Server *ServerInfo `json:"Server,omitempty" name:"Server"`

	// 结果 0成功，1失败
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 原因
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 分配IP
	AllocateIp *string `json:"AllocateIp,omitempty" name:"AllocateIp"`

	// 分配掩码
	AllocateMask *string `json:"AllocateMask,omitempty" name:"AllocateMask"`

	// 分配网关
	AllocateGateWay *string `json:"AllocateGateWay,omitempty" name:"AllocateGateWay"`

	// 搬迁后内网IP
	RelocateLanIp *string `json:"RelocateLanIp,omitempty" name:"RelocateLanIp"`

	// 搬迁后内网MASK
	RelocateLanMask *string `json:"RelocateLanMask,omitempty" name:"RelocateLanMask"`

	// 搬迁后内网GATEWAY
	RelocateLanGateway *string `json:"RelocateLanGateway,omitempty" name:"RelocateLanGateway"`

	// 固资
	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`

	// lanip
	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type ServerIpData struct {

	// 服务IP信息列表
	Servers []*ServerAllocateIpInfo `json:"Servers,omitempty" name:"Servers" list`
}

type ServerIpInfo struct {

	// 网络接口
	Nic *string `json:"Nic,omitempty" name:"Nic"`

	// IP地址
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 子网掩码
	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`

	// 网关
	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
}

type ServerModifyData struct {

	// 筛选服务器
	Condition *ServerInfo `json:"Condition,omitempty" name:"Condition"`

	// 修改数据
	Modify *ServerInfo `json:"Modify,omitempty" name:"Modify"`
}

type SnmpDevice struct {

	// 设备ID
	DevId *int64 `json:"DevId,omitempty" name:"DevId"`

	// 固资ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 设备具体信息
	InfoSet []*SnmpDeviceInfo `json:"InfoSet,omitempty" name:"InfoSet" list`
}

type SnmpDeviceInfo struct {

	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	Values []*NumKeyValue `json:"Values,omitempty" name:"Values" list`
}

type SnmpEntry struct {

	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type SnmpQueryDataEntry struct {

	// 设备端口
	DevPort *string `json:"DevPort,omitempty" name:"DevPort"`

	// 数据详情
	Details []*SnmpEntry `json:"Details,omitempty" name:"Details" list`
}

type SnmpTempConfig struct {

	// 设备类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 厂商
	Vendor *string `json:"Vendor,omitempty" name:"Vendor"`

	// 机型
	Model *string `json:"Model,omitempty" name:"Model"`

	// 模版名
	TempName *string `json:"TempName,omitempty" name:"TempName"`

	// 状态 0->有数据，1->无数据
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 字符串数组
	Datas []*string `json:"Datas,omitempty" name:"Datas" list`
}

type SyncSnmpDataRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：snmp
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：sync_snmp_data
	Op *string `json:"Op,omitempty" name:"Op"`

	// 这里填”sync_all”
	Type *string `json:"Type,omitempty" name:"Type"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *SyncSnmpDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncSnmpDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncSnmpDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 结果详情
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 本次数据上报时间
		ReportTime *string `json:"ReportTime,omitempty" name:"ReportTime"`

		// 分钟的索引，表示该天第几分钟的数据
		MinuteIdx *int64 `json:"MinuteIdx,omitempty" name:"MinuteIdx"`

		// 返回结果
		DeviceSet []*SnmpDevice `json:"DeviceSet,omitempty" name:"DeviceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncSnmpDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncSnmpDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogAddKeywordRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：syslog
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：add_syslog_keyword
	Op *string `json:"Op,omitempty" name:"Op"`

	// 关键字配置
	Keywords []*SyslogInfo `json:"Keywords,omitempty" name:"Keywords" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *SyslogAddKeywordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogAddKeywordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogAddKeywordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 详细信息
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyslogAddKeywordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogAddKeywordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogDeleteKeywordRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：syslog
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：delete_syslog_keyword
	Op *string `json:"Op,omitempty" name:"Op"`

	// 关键字配置
	Keywords []*SyslogInfo `json:"Keywords,omitempty" name:"Keywords" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *SyslogDeleteKeywordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogDeleteKeywordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogDeleteKeywordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 失败返回原因，成功为空
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyslogDeleteKeywordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogDeleteKeywordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogInfo struct {

	// 类型
	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`

	// 厂商
	NetdevVendor *string `json:"NetdevVendor,omitempty" name:"NetdevVendor"`

	// 型号
	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`

	// 功能
	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`

	// 告警关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 新关键字，修改的时候使用
	NewKeyword *string `json:"NewKeyword,omitempty" name:"NewKeyword"`

	// 所属告警类型
	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警关键字描述
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type SyslogQueryKeywordRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：syslog
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：select_syslog_keyword
	Op *string `json:"Op,omitempty" name:"Op"`

	// 关键字配置
	Keyword *SyslogInfo `json:"Keyword,omitempty" name:"Keyword"`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *SyslogQueryKeywordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogQueryKeywordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogQueryKeywordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 失败返回原因，成功为空
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 关键字配置
		DataSet []*ServerIpInfo `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyslogQueryKeywordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogQueryKeywordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogQueryRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：syslog
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：get_syslog_keyword
	Op *string `json:"Op,omitempty" name:"Op"`

	// 范围选择
	RangeFilter *RangeFilter `json:"RangeFilter,omitempty" name:"RangeFilter"`

	// 查询偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 1：需要总量，0：不需要总量
	NeedTotal *int64 `json:"NeedTotal,omitempty" name:"NeedTotal"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *SyslogQueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogQueryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogQueryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 失败返回原因，成功为空
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 如果NeedTotal为0，则忽略该项即可
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 日志数据
		DataSet []*LogEntry `json:"DataSet,omitempty" name:"DataSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyslogQueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogQueryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogUpdateKeywordRequest struct {
	*tchttp.BaseRequest

	// 公共参数，本接口取值：syslog
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 公共参数，本接口取值：update_syslog_keyword
	Op *string `json:"Op,omitempty" name:"Op"`

	// 关键字配置
	Keywords []*SyslogInfo `json:"Keywords,omitempty" name:"Keywords" list`

	// dcos底层服务按idc部署，操作类请求需要指定idcid，路由请求到对应idc的dcos处理
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *SyslogUpdateKeywordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogUpdateKeywordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyslogUpdateKeywordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果，0成功，其他失败
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 失败返回原因，成功为空
		Detail *string `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyslogUpdateKeywordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyslogUpdateKeywordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValueArray struct {

	// 数值数组
	Values []*int64 `json:"Values,omitempty" name:"Values" list`
}

type XflowLineDataInfo struct {

	// 对应的时间点
	Time *string `json:"Time,omitempty" name:"Time"`

	// 平均数据量，kbps
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// 总量，KB
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type XflowLineInfo struct {

	// 专线ID
	LineId *int64 `json:"LineId,omitempty" name:"LineId"`

	// 网络设备固资
	DeviceAsset *string `json:"DeviceAsset,omitempty" name:"DeviceAsset"`

	// 端口索引
	IfIndex *int64 `json:"IfIndex,omitempty" name:"IfIndex"`

	// 管理ip
	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`

	// 其他ip(多个以”,”隔开)
	OtherIp *string `json:"OtherIp,omitempty" name:"OtherIp"`

	// 源IP
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 目的IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 源端口
	SrcPort *int64 `json:"SrcPort,omitempty" name:"SrcPort"`

	// 目的端口
	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`

	// 协议，1->ICMP，6->TCP，17->UDP
	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`

	// 方向，1->入流量，2->出流量
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 值
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// 比例
	Proportion *float64 `json:"Proportion,omitempty" name:"Proportion"`
}

type XflowLineResult struct {

	// 专线信息
	LineInfo *XflowLineInfo `json:"LineInfo,omitempty" name:"LineInfo"`

	// 专线流量信息
	LineDataInfos []*XflowLineDataInfo `json:"LineDataInfos,omitempty" name:"LineDataInfos" list`
}
