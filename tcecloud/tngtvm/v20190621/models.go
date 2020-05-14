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

package v20190621

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AssetHistoryTask struct {

	// 一次扫描任务唯一标识
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// 任务日志开始扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type AssetStatLog struct {

	// 资产值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ftarget *string `json:"Ftarget,omitempty" name:"Ftarget"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fstatus *uint64 `json:"Fstatus,omitempty" name:"Fstatus"`

	// 高风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	FhighRisk *uint64 `json:"FhighRisk,omitempty" name:"FhighRisk"`

	// 中风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	FmidRisk *uint64 `json:"FmidRisk,omitempty" name:"FmidRisk"`

	// 低风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowRisk *uint64 `json:"FlowRisk,omitempty" name:"FlowRisk"`

	// 提示风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinfoRisk *uint64 `json:"FinfoRisk,omitempty" name:"FinfoRisk"`

	// 限制端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	FportLimit *uint64 `json:"FportLimit,omitempty" name:"FportLimit"`

	// 内容风险1
	// 注意：此字段可能返回 null，表示取不到有效值。
	FyellowRisk *uint64 `json:"FyellowRisk,omitempty" name:"FyellowRisk"`

	// 内容风险2
	// 注意：此字段可能返回 null，表示取不到有效值。
	FpoliticalRisk *uint64 `json:"FpoliticalRisk,omitempty" name:"FpoliticalRisk"`

	// 内容风险3
	// 注意：此字段可能返回 null，表示取不到有效值。
	FdrugRisk *uint64 `json:"FdrugRisk,omitempty" name:"FdrugRisk"`

	// 内容风险4
	// 注意：此字段可能返回 null，表示取不到有效值。
	FhorrorRisk *uint64 `json:"FhorrorRisk,omitempty" name:"FhorrorRisk"`

	// 内容风险5
	// 注意：此字段可能返回 null，表示取不到有效值。
	FgambleRisk *uint64 `json:"FgambleRisk,omitempty" name:"FgambleRisk"`

	// 端口建议关闭数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuggestClose *uint64 `json:"SuggestClose,omitempty" name:"SuggestClose"`
}

type CreateAssetTaskRequest struct {
	*tchttp.BaseRequest

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 资产唯一标识符数组
	AssetIds []*uint64 `json:"AssetIds,omitempty" name:"AssetIds" list`

	// 扫描计划类型：instant-立即扫描，timed-定时扫描
	Type *string `json:"Type,omitempty" name:"Type"`

	// 扫描速度
	ScanSpeed *string `json:"ScanSpeed,omitempty" name:"ScanSpeed"`

	// 扫描类型
	DetectType *string `json:"DetectType,omitempty" name:"DetectType"`

	// 扫描资产类型0网站 1主机
	MonitorType *uint64 `json:"MonitorType,omitempty" name:"MonitorType"`

	// 定时任务扫描的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 扫描深度
	ScanDeep *string `json:"ScanDeep,omitempty" name:"ScanDeep"`

	// 页面创建来源：asset标识资产创建，task标识任务创建
	Page *string `json:"Page,omitempty" name:"Page"`
}

func (r *CreateAssetTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAssetTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAssetTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAssetTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAssetsRequest struct {
	*tchttp.BaseRequest

	// 添加的资产类型：'1'-域名，'2'-ip
	Type *string `json:"Type,omitempty" name:"Type"`

	// 资产地址数组
	Url []*string `json:"Url,omitempty" name:"Url" list`
}

func (r *CreateAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAssetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAssetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资产对应的认证码
		AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`

		// 批量添加的资产唯一标识组成的数组
		Ids []*int64 `json:"Ids,omitempty" name:"Ids" list`

		// 认证文件名
		AuthFile *string `json:"AuthFile,omitempty" name:"AuthFile"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAssetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 频率
	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`

	// 资产列表
	AssetIds []*uint64 `json:"AssetIds,omitempty" name:"AssetIds" list`

	// 开始时间段
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间段
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 扫描类型
	DetectType []*string `json:"DetectType,omitempty" name:"DetectType" list`

	// 扫描速率 middle high low
	ScanSpeed *string `json:"ScanSpeed,omitempty" name:"ScanSpeed"`

	// 扫描内容0网站 1主机
	MonitorType *uint64 `json:"MonitorType,omitempty" name:"MonitorType"`

	// 页面创建来源：asset标识资产创建，task标识任务创建
	Page *string `json:"Page,omitempty" name:"Page"`

	// 扫描深度
	ScanDeep *string `json:"ScanDeep,omitempty" name:"ScanDeep"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataAsset struct {

	// 资产id唯一标识
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 用户唯一标识
	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`

	// 资产类型：1-域名，2-IP
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 资产名称
	Value *string `json:"Value,omitempty" name:"Value"`

	// 资产创建时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 验证状态：0-未验证，1-已验证，2-已失效
	AuthStatus *int64 `json:"AuthStatus,omitempty" name:"AuthStatus"`

	// 资产授权码
	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 高风险数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighRisk *uint64 `json:"HighRisk,omitempty" name:"HighRisk"`

	// 中风险数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MidRisk *uint64 `json:"MidRisk,omitempty" name:"MidRisk"`

	// 低风险数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowRisk *uint64 `json:"LowRisk,omitempty" name:"LowRisk"`

	// 提示风险数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfoRisk *uint64 `json:"InfoRisk,omitempty" name:"InfoRisk"`

	// 内容涉黄数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	YellowRisk *uint64 `json:"YellowRisk,omitempty" name:"YellowRisk"`

	// 内容涉政数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalRisk *uint64 `json:"PoliticalRisk,omitempty" name:"PoliticalRisk"`

	// 内容涉毒数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrugRisk *uint64 `json:"DrugRisk,omitempty" name:"DrugRisk"`

	// 内容涉恐数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HorrorRisk *uint64 `json:"HorrorRisk,omitempty" name:"HorrorRisk"`

	// 资产上次扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 资产扫描状态：0 未扫描  1正在扫描'
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 内容涉赌数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GambleRisk *uint64 `json:"GambleRisk,omitempty" name:"GambleRisk"`

	// 端口风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortLimit *uint64 `json:"PortLimit,omitempty" name:"PortLimit"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *uint64 `json:"AssetType,omitempty" name:"AssetType"`

	// cookie
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimloginCookie *string `json:"SimloginCookie,omitempty" name:"SimloginCookie"`

	// 进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *uint64 `json:"Process,omitempty" name:"Process"`

	// 验证文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthFile *string `json:"AuthFile,omitempty" name:"AuthFile"`
}

type DataContentVul struct {

	// 内容风险唯一标识
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Target *string `json:"Target,omitempty" name:"Target"`

	// 风险类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 风险描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// 风险状态
	VulStatus *uint64 `json:"VulStatus,omitempty" name:"VulStatus"`

	// 插入时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 检测类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectType *string `json:"DetectType,omitempty" name:"DetectType"`

	// 取证路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 来源页面
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceUrl *string `json:"SourceUrl,omitempty" name:"SourceUrl"`

	// 非法链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type DataOperateLog struct {

	// 日志唯一标识
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 用户uid
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 操作类型
	Operate *string `json:"Operate,omitempty" name:"Operate"`

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 页面操作路径
	Page *string `json:"Page,omitempty" name:"Page"`

	// 操作日志插入时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 登陆IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DataOverview struct {

	// 网站风险趋势分布，默认展示近一个月
	WebTrend []*RiskTrend `json:"WebTrend,omitempty" name:"WebTrend" list`

	// 套餐类型
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 到期时间
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// 域名总个数
	DomainTotal *uint64 `json:"DomainTotal,omitempty" name:"DomainTotal"`

	// 使用的域名数
	DomainUsedCount *uint64 `json:"DomainUsedCount,omitempty" name:"DomainUsedCount"`

	// Ip总个数
	IpTotal *uint64 `json:"IpTotal,omitempty" name:"IpTotal"`

	// 使用的Ip数
	IpUsedCount *uint64 `json:"IpUsedCount,omitempty" name:"IpUsedCount"`

	// Url总个数
	UrlTotal *uint64 `json:"UrlTotal,omitempty" name:"UrlTotal"`

	// 使用的Url数
	UrlUsedCount *uint64 `json:"UrlUsedCount,omitempty" name:"UrlUsedCount"`

	// 网站总个数
	WebTotal *uint64 `json:"WebTotal,omitempty" name:"WebTotal"`

	// 已认证网站总个数
	WebAuthTotal *uint64 `json:"WebAuthTotal,omitempty" name:"WebAuthTotal"`

	// 网站内容风险数
	ContentTotal *uint64 `json:"ContentTotal,omitempty" name:"ContentTotal"`

	// 网站漏洞高危个数
	WebHoleHigh *uint64 `json:"WebHoleHigh,omitempty" name:"WebHoleHigh"`

	// 网站漏洞中危个数
	WebHoleMid *uint64 `json:"WebHoleMid,omitempty" name:"WebHoleMid"`

	// 网站漏洞低危个数
	WebHoleLow *uint64 `json:"WebHoleLow,omitempty" name:"WebHoleLow"`

	// 网站漏洞提示个数
	WebHoleInfo *uint64 `json:"WebHoleInfo,omitempty" name:"WebHoleInfo"`

	// 主机总个数
	HostTotal *uint64 `json:"HostTotal,omitempty" name:"HostTotal"`

	// 已认证主机总个数
	HostAuthTotal *uint64 `json:"HostAuthTotal,omitempty" name:"HostAuthTotal"`

	// 端口风险数
	PortTotal *uint64 `json:"PortTotal,omitempty" name:"PortTotal"`

	// 主机漏洞高危个数
	HostHoleHigh *uint64 `json:"HostHoleHigh,omitempty" name:"HostHoleHigh"`

	// 主机漏洞低危个数
	HostHoleMid *uint64 `json:"HostHoleMid,omitempty" name:"HostHoleMid"`

	// 主机漏洞低危个数
	HostHoleLow *uint64 `json:"HostHoleLow,omitempty" name:"HostHoleLow"`

	// 主机漏洞提示个数
	HostHoleInfo *uint64 `json:"HostHoleInfo,omitempty" name:"HostHoleInfo"`

	// 主机风险趋势分布，默认展示近一个月
	HostTrend []*RiskTrend `json:"HostTrend,omitempty" name:"HostTrend" list`
}

type DataPort struct {

	// 任务日志唯一标识符
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`

	// 主机ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 端口服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 端口表唯一标识
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DataVul struct {

	// 用于标示漏洞ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 扫描计划ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`

	// 用户ID，用户标示
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *int64 `json:"UserId,omitempty" name:"UserId"`

	// 当前任务ID要扫描的目标(域名/IP)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Target *string `json:"Target,omitempty" name:"Target"`

	// 扫描结果来源  有webscan/pocscan/csscan/weakscan 四种
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 漏洞类型，如xss、sqli等，供四种扫描结果使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 存在漏洞的URL，供四种扫描结果使用，webscan格式为完整url，pocsuite、csscan为输入的target（若可获取完整URL，则输出完整URL），weakscan格式为 协议://IP:PORT
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulUrl *string `json:"VulUrl,omitempty" name:"VulUrl"`

	// 漏洞名，共pocsuite、csscan使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	PocName *string `json:"PocName,omitempty" name:"PocName"`

	// 漏洞详情唯一标示，用于给用户设定某个漏洞是否忽略、误报。由引擎生成，生成算法： 对webscan  MD5(fvul_url + fvul_type + fuser_id)  对pocscan MD5(ftarget+ fpoc_id+ fuser_id)  对weakscan MD5(ftarget + fpayload + fuser_id)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulId *string `json:"VulId,omitempty" name:"VulId"`

	// 漏洞状态，正常状态为0，已处理为1，待验证为2，验证中为3，忽略为4，误报为5。
	VulStatus *int64 `json:"VulStatus,omitempty" name:"VulStatus"`

	// 漏洞危害等级, 高为high/中为middle/低危为low/提示为info
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// 漏洞发现时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
}

type DeleteAssetRequest struct {
	*tchttp.BaseRequest

	// 资产唯一标识符
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAssetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAssetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateLogsRequest struct {
	*tchttp.BaseRequest

	// 操作日志唯一标识符
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteOperateLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOperateLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperateLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteOperateLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskLogRequest struct {
	*tchttp.BaseRequest

	// 任务日志的ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteTaskLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTaskLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTaskLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTaskLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWhiteUrlRequest struct {
	*tchttp.BaseRequest

	// ID数组
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteWhiteUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWhiteUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWhiteUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWhiteUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWhiteUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHistoryTasksRequest struct {
	*tchttp.BaseRequest

	// 资产唯一标识符
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetHistoryTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetHistoryTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHistoryTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取特定资产下的历史任务日志的subTaskid和扫描结束时间
		Data []*AssetHistoryTask `json:"Data,omitempty" name:"Data" list`

		// 资产url或ip
		Asset *string `json:"Asset,omitempty" name:"Asset"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetHistoryTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetHistoryTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetInfoByTaskIdRequest struct {
	*tchttp.BaseRequest

	// 每次扫描任务的唯一标识符
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资产名称
	Asset *string `json:"Asset,omitempty" name:"Asset"`
}

func (r *DescribeAssetInfoByTaskIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetInfoByTaskIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetInfoByTaskIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 高风险数量
		HighRisk *uint64 `json:"HighRisk,omitempty" name:"HighRisk"`

		// 中风险数量
		MidRisk *uint64 `json:"MidRisk,omitempty" name:"MidRisk"`

		// 低风险数量
		LowRisk *uint64 `json:"LowRisk,omitempty" name:"LowRisk"`

		// 提示风险数量
		InfoRisk *uint64 `json:"InfoRisk,omitempty" name:"InfoRisk"`

		// 内容涉黄数量
		YellowRisk *uint64 `json:"YellowRisk,omitempty" name:"YellowRisk"`

		// 内容涉政数量
		PoliticalRisk *uint64 `json:"PoliticalRisk,omitempty" name:"PoliticalRisk"`

		// 内容涉毒数量
		DrugRisk *uint64 `json:"DrugRisk,omitempty" name:"DrugRisk"`

		// 内容涉恐数量
		HorrorRisk *uint64 `json:"HorrorRisk,omitempty" name:"HorrorRisk"`

		// 端口建议限制访问
		PortLimit *uint64 `json:"PortLimit,omitempty" name:"PortLimit"`

		// 任务名称
		TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

		// 扫描耗时
		ConsumeTime *string `json:"ConsumeTime,omitempty" name:"ConsumeTime"`

		// 扫描计划
		ScanPlan *string `json:"ScanPlan,omitempty" name:"ScanPlan"`

		// 扫描速度
		ScanSpeed *string `json:"ScanSpeed,omitempty" name:"ScanSpeed"`

		// 检测类型
		DetectType *string `json:"DetectType,omitempty" name:"DetectType"`

		// 扫描开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 资产名称
		Asset *string `json:"Asset,omitempty" name:"Asset"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 扫描频率,单位天,1-每天,7-每周,30-月,0-扫描一次
		Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`

		// 扫描任务类型
		TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

		// 组件
		Component *string `json:"Component,omitempty" name:"Component"`

		// 内容涉赌数量
		GambleRisk *uint64 `json:"GambleRisk,omitempty" name:"GambleRisk"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetInfoByTaskIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetInfoByTaskIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetListRequest struct {
	*tchttp.BaseRequest

	// 限制数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 资产认证状态：'0‘-未验证,‘1‘-已验证,'2'-验证失效
	Status *string `json:"Status,omitempty" name:"Status"`

	// 资产类型：‘1’-域名，‘2’-ip
	Type *string `json:"Type,omitempty" name:"Type"`

	// 搜索过滤关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 排序列
	By *string `json:"By,omitempty" name:"By"`

	// 排序升降：‘desc’-降序排列，‘asc’-升序排列
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询资产类型
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
}

func (r *DescribeAssetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资产对象数组
		Data []*DataAsset `json:"Data,omitempty" name:"Data" list`

		// 资产总数统计
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetStatisRequest struct {
	*tchttp.BaseRequest

	// 日志ID，32位随机字符
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetStatisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetStatisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetStatisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资产
		Assets []*StatAsset `json:"Assets,omitempty" name:"Assets" list`

		// 列表
		List []*AssetStatLog `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetStatisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetStatisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetStatusRequest struct {
	*tchttp.BaseRequest

	// 资产的唯一标识数组
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DescribeAssetStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 判断资产认证过程是否结束
		Validated *bool `json:"Validated,omitempty" name:"Validated"`

		// 认证失败的资产唯一标识符数组
		Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`

		// 认证失败的资产地址
		Urls []*string `json:"Urls,omitempty" name:"Urls" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetStructureRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 当前任务ID要扫描的目标(域名/IP)
	Asset *string `json:"Asset,omitempty" name:"Asset"`
}

func (r *DescribeAssetStructureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetStructureRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetStructureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// json形式的字符串，展示特定任务和资产下漏洞url的树状结构
		UrlTree *string `json:"UrlTree,omitempty" name:"UrlTree"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetStructureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetStructureResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentVulDetailRequest struct {
	*tchttp.BaseRequest

	// 内容风险唯一标识
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeContentVulDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentVulDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentVulDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内容风险详情
		Vul *DataContentVul `json:"Vul,omitempty" name:"Vul"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContentVulDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentVulDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentVulListNewRequest struct {
	*tchttp.BaseRequest

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 漏洞状态：正常状态为0，已处理为1，待验证为2，验证中为3，忽略为4，误报为5，未修复为6. 空数组为全选；不传该参数默认为全选；
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// 风险类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 搜索过滤关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 任务日志ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 排序列
	By *string `json:"By,omitempty" name:"By"`

	// 排序列的升降
	Order *string `json:"Order,omitempty" name:"Order"`

	// 资产名称
	Asset *string `json:"Asset,omitempty" name:"Asset"`
}

func (r *DescribeContentVulListNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentVulListNewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentVulListNewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内容风险对象数组
		Data []*DataContentVul `json:"Data,omitempty" name:"Data" list`

		// 漏洞总数统计
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContentVulListNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentVulListNewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentVulListRequest struct {
	*tchttp.BaseRequest

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 漏洞状态：正常状态为0，已处理为1，待验证为2，验证中为3，忽略为4，误报为5
	Status *string `json:"Status,omitempty" name:"Status"`

	// 风险类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 搜索过滤关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 任务日志ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 排序列
	By *string `json:"By,omitempty" name:"By"`

	// 排序列的升降
	Order *string `json:"Order,omitempty" name:"Order"`

	// 资产名称
	Asset *string `json:"Asset,omitempty" name:"Asset"`
}

func (r *DescribeContentVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentVulListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentVulListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内容风险对象数组
		Data []*DataContentVul `json:"Data,omitempty" name:"Data" list`

		// 漏洞总数统计
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContentVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentVulListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCurTaskRequest struct {
	*tchttp.BaseRequest

	// 当前页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 返回条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCurTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCurTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCurTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表
		Data []*TaskObject `json:"Data,omitempty" name:"Data" list`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCurTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCurTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogByTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeLogByTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogByTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogByTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 漏洞总数
		Vul *uint64 `json:"Vul,omitempty" name:"Vul"`

		// 内容风险总数
		Content *uint64 `json:"Content,omitempty" name:"Content"`

		// 日志概况
		List []*TaskLogInfo `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogByTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogByTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateLogsListRequest struct {
	*tchttp.BaseRequest

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索过滤关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 排序列
	By *string `json:"By,omitempty" name:"By"`

	// 排序列的升降
	Order *string `json:"Order,omitempty" name:"Order"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOperateLogsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOperateLogsListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateLogsListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作记录数组
		Data []*DataOperateLog `json:"Data,omitempty" name:"Data" list`

		// 操作记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateLogsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOperateLogsListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总览页数据对象
		Overview *DataOverview `json:"Overview,omitempty" name:"Overview"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePortListRequest struct {
	*tchttp.BaseRequest

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 端口建议
	Advise *string `json:"Advise,omitempty" name:"Advise"`

	// 搜索过滤关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 任务日志ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 排序列
	By *string `json:"By,omitempty" name:"By"`

	// 排序列的升降
	Order *string `json:"Order,omitempty" name:"Order"`

	// 资产名称
	Asset *string `json:"Asset,omitempty" name:"Asset"`
}

func (r *DescribePortListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePortListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePortListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 端口对象数组
		Data []*DataPort `json:"Data,omitempty" name:"Data" list`

		// 端口总数统计
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePortListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePortListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskTrendRequest struct {
	*tchttp.BaseRequest

	// 查询多少天的风险趋势数据
	Day *int64 `json:"Day,omitempty" name:"Day"`
}

func (r *DescribeRiskTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRiskTrendRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 风险趋势分布
		RiskTrend []*RiskTrend `json:"RiskTrend,omitempty" name:"RiskTrend" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRiskTrendResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSettingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户数组
		User *UserObject `json:"User,omitempty" name:"User"`

		// 漏洞库版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 时间
		Time *uint64 `json:"Time,omitempty" name:"Time"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSettingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTCInstancesRequest struct {
	*tchttp.BaseRequest

	// 区域
	TCRegion *string `json:"TCRegion,omitempty" name:"TCRegion"`
}

func (r *DescribeTCInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTCInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTCInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例列表
		Instances []*TCInstanceObject `json:"Instances,omitempty" name:"Instances" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTCInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTCInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogDetailRequest struct {
	*tchttp.BaseRequest

	// 关联的子任务ID
	SubTaskId *string `json:"SubTaskId,omitempty" name:"SubTaskId"`

	// 条数限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTaskLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskLogDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 详细数据
		Data []*TaskLogDetailType `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskLogDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogRequest struct {
	*tchttp.BaseRequest

	// 单个Logid
	Id *string `json:"Id,omitempty" name:"Id"`

	// 页面数大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 当前页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 数据
		Data []*TaskLog `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskPdfRequest struct {
	*tchttp.BaseRequest

	// 任务日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTaskPdfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskPdfRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskPdfResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// pdf下载地址
		Url *string `json:"Url,omitempty" name:"Url"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskPdfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskPdfResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest

	// 模糊搜索关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 页面数大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 当前页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 任务当前状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务类型，0-周期任务,1-立即扫描,2-定时扫描
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 返回数据
		Data []*TaskObject `json:"Data,omitempty" name:"Data" list`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTokenRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 令牌
		Token *string `json:"Token,omitempty" name:"Token"`

		// 资源的下载地址
		Url *string `json:"Url,omitempty" name:"Url"`

		// 用户资源ID
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDetailRequest struct {
	*tchttp.BaseRequest

	// 漏洞唯一标识
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeVulDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 漏洞详情
		Vul *VulDetail `json:"Vul,omitempty" name:"Vul"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulListByTypeRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 漏洞状态：正常状态为0，已处理为1，待验证为2，验证中为3，忽略为4，误报为5，未修复为6. 空数组为全选；不传该参数默认为全选；
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// 漏洞危害等级：高为high/中为middle/低危为low/提示为info
	Level *string `json:"Level,omitempty" name:"Level"`

	// 搜索过滤关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 任务日志ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 存在漏洞的URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 排序列
	By *string `json:"By,omitempty" name:"By"`

	// 排序列的升降
	Order *string `json:"Order,omitempty" name:"Order"`

	// 资产名称
	Asset *string `json:"Asset,omitempty" name:"Asset"`

	// 资产类型，0为域名，1为主机IP；空数组或不传该参数则默认全选
	Type []*int64 `json:"Type,omitempty" name:"Type" list`
}

func (r *DescribeVulListByTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulListByTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulListByTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 漏洞对象数组
		Data []*DataVul `json:"Data,omitempty" name:"Data" list`

		// 漏洞总数统计
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulListByTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulListByTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulListRequest struct {
	*tchttp.BaseRequest

	// 限制数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 漏洞状态：正常状态为0，已处理为1，待验证为2，验证中为3，忽略为4，误报为5
	Status *string `json:"Status,omitempty" name:"Status"`

	// 漏洞危害等级：高为high/中为middle/低危为low/提示为info
	Level *string `json:"Level,omitempty" name:"Level"`

	// 搜索过滤关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 任务日志ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 存在漏洞的URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 排序列
	By *string `json:"By,omitempty" name:"By"`

	// 排序列的升降
	Order *string `json:"Order,omitempty" name:"Order"`

	// 资产名称
	Asset *string `json:"Asset,omitempty" name:"Asset"`
}

func (r *DescribeVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 漏洞对象数组
		Data []*DataVul `json:"Data,omitempty" name:"Data" list`

		// 漏洞总数统计
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulStatusRequest struct {
	*tchttp.BaseRequest

	// 漏洞或内容风险的唯一标识符数组
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`

	// 查询的是否是漏洞的验证状态
	IsPort *bool `json:"IsPort,omitempty" name:"IsPort"`
}

func (r *DescribeVulStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 判断漏洞或内容风险验证过程是否结束
		Validated *bool `json:"Validated,omitempty" name:"Validated"`

		// 通过验证的漏洞或内容风险验证的唯一标识符数组
		SucIds []*uint64 `json:"SucIds,omitempty" name:"SucIds" list`

		// 未通过验证的漏洞或内容风险验证的唯一标识符数组
		FailIds []*uint64 `json:"FailIds,omitempty" name:"FailIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteUrlRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWhiteUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWhiteUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 白名单列表
		Data []*WhiteUrl `json:"Data,omitempty" name:"Data" list`

		// 简单JSON格式 []
		Assets *string `json:"Assets,omitempty" name:"Assets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWhiteUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetIntelligenceDetailRequest struct {
	*tchttp.BaseRequest

	// 要查询的威胁情报ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *GetIntelligenceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetIntelligenceDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetIntelligenceDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 威胁情报推送详情
		Data *IntelligenceDetail `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetIntelligenceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetIntelligenceDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetIntelligenceListRequest struct {
	*tchttp.BaseRequest

	// 进行模糊匹配的关键字
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 每次请求的条数，默认为10
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 每次请求的偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 请求根据哪一个字段进行排序
	By *string `json:"By,omitempty" name:"By"`

	// 排序顺序，允许的值为desc/asc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 根据威胁情报威胁等级分类，1-4分别是高危-中危-低危-提示，传空值为默认全部
	Severity *string `json:"Severity,omitempty" name:"Severity"`

	// 根据威胁情报的发布时间月份进行筛选
	Month *string `json:"Month,omitempty" name:"Month"`
}

func (r *GetIntelligenceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetIntelligenceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetIntelligenceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据
		Data []*Intelligence `json:"Data,omitempty" name:"Data" list`

		// 符合条件的数据总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetIntelligenceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetIntelligenceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetThreatIntelligencePushRequest struct {
	*tchttp.BaseRequest
}

func (r *GetThreatIntelligencePushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetThreatIntelligencePushRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetThreatIntelligencePushResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 24小时内的威胁情报推送
		Data []*Intelligence `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetThreatIntelligencePushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetThreatIntelligencePushResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Intelligence struct {

	// 威胁情报标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitempty" name:"Title"`

	// 威胁情报ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 威胁等级：1-4依次为提示-低危-中危-高危
	// 注意：此字段可能返回 null，表示取不到有效值。
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PubTime *string `json:"PubTime,omitempty" name:"PubTime"`

	// 威胁情报来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 该威胁情报用于针对扫描可能需要的资产ID类别：-1为不扫描，0为网站，1为主机，2为网站+主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`

	// 该威胁情报用于针对扫描可能需要的poc_id值
	// 注意：此字段可能返回 null，表示取不到有效值。
	PocId *string `json:"PocId,omitempty" name:"PocId"`
}

type IntelligenceDetail struct {

	// 威胁情报标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitempty" name:"Title"`

	// 威胁情报ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 威胁等级：1-4依次为提示-低危-中危-高危
	// 注意：此字段可能返回 null，表示取不到有效值。
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PubTime *string `json:"PubTime,omitempty" name:"PubTime"`

	// 威胁情报来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 威胁情报内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 威胁情报摘要
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abstract *string `json:"Abstract,omitempty" name:"Abstract"`

	// 威胁情报作者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Author *string `json:"Author,omitempty" name:"Author"`

	// 威胁情报相关URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Urls *string `json:"Urls,omitempty" name:"Urls"`

	// 威胁情报相关平台
	// 注意：此字段可能返回 null，表示取不到有效值。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 威胁情报相关IOC
	// 注意：此字段可能返回 null，表示取不到有效值。
	Iocs *string `json:"Iocs,omitempty" name:"Iocs"`

	// 该威胁情报用于针对扫描可能需要的poc_id值
	// 注意：此字段可能返回 null，表示取不到有效值。
	PocId *string `json:"PocId,omitempty" name:"PocId"`

	// 该威胁情报用于针对扫描可能需要的资产ID类别：-1为不扫描，0为网站，1为主机，2为网站+主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
}

type ModifyAssetSimloginRequest struct {
	*tchttp.BaseRequest

	// 资产ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// cookie值
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
}

func (r *ModifyAssetSimloginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAssetSimloginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetSimloginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetSimloginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAssetSimloginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetStatusRequest struct {
	*tchttp.BaseRequest

	// 资产唯一标识数组
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *ModifyAssetStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAssetStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAssetStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContentVulStatusRequest struct {
	*tchttp.BaseRequest

	// 内容风险唯一标识
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 更改内容风险的状态，processed-设为已处理，ignore-设为忽略，misinfo-标记误报，validate-立即验证，ignoreCancel-取消忽略，misinfoCancel-取消标记误报
	Operate *string `json:"Operate,omitempty" name:"Operate"`

	// 接口操作页面来源
	Page *string `json:"Page,omitempty" name:"Page"`
}

func (r *ModifyContentVulStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContentVulStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContentVulStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContentVulStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContentVulStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySettingRequest struct {
	*tchttp.BaseRequest

	// 设置风险通知等级
	NotifyLevel []*string `json:"NotifyLevel,omitempty" name:"NotifyLevel" list`
}

func (r *ModifySettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySettingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySettingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 需要修改的数据
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVulStatusRequest struct {
	*tchttp.BaseRequest

	// 漏洞唯一标识
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 更改漏洞的状态，processed-设为已处理，ignore-设为忽略，misinfo-标记误报，validate-立即验证，ignoreCancel-取消忽略，misinfoCancel-取消标记误报
	Operate *string `json:"Operate,omitempty" name:"Operate"`

	// 接口操作页面来源
	Page *string `json:"Page,omitempty" name:"Page"`
}

func (r *ModifyVulStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVulStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVulStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVulStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWhiteUrlRequest struct {
	*tchttp.BaseRequest

	// 具体Url
	Url *string `json:"Url,omitempty" name:"Url"`

	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyWhiteUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWhiteUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWhiteUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWhiteUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWhiteUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartAssetTaskRequest struct {
	*tchttp.BaseRequest

	// 需要重新启动的GruopId
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *RestartAssetTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartAssetTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartAssetTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartAssetTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartAssetTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RiskTrend struct {

	// 日期
	Time *string `json:"Time,omitempty" name:"Time"`

	// 漏洞类型，取值：高危漏洞/中危漏洞/低危漏洞/提示漏洞
	Type *string `json:"Type,omitempty" name:"Type"`

	// 漏洞个数
	Val *int64 `json:"Val,omitempty" name:"Val"`
}

type StartIntelligenceScanRequest struct {
	*tchttp.BaseRequest

	// 针对该威胁情报扫描任务的PocId值
	PocId *string `json:"PocId,omitempty" name:"PocId"`

	// 该威胁情报扫描任务所扫描的资产类型，0为网站，1为主机
	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
}

func (r *StartIntelligenceScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartIntelligenceScanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartIntelligenceScanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartIntelligenceScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartIntelligenceScanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StatAsset struct {

	// Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 具体值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type StopRunningTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 若带上，就是停止某个log
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// 页面创建来源：overview标识首页，task标识任务
	Page *string `json:"Page,omitempty" name:"Page"`
}

func (r *StopRunningTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopRunningTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopRunningTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopRunningTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopRunningTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TCInstanceObject struct {

	// 外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitempty" name:"IP"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type TaskLog struct {

	// 日志id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlogId *string `json:"FlogId,omitempty" name:"FlogId"`

	// 资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	Target *string `json:"Target,omitempty" name:"Target"`

	// 扫描类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FscanType *string `json:"FscanType,omitempty" name:"FscanType"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fstatus *uint64 `json:"Fstatus,omitempty" name:"Fstatus"`

	// 频率
	// 注意：此字段可能返回 null，表示取不到有效值。
	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FstartTime *string `json:"FstartTime,omitempty" name:"FstartTime"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 高风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighRisk *uint64 `json:"HighRisk,omitempty" name:"HighRisk"`

	// 中风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MidRisk *uint64 `json:"MidRisk,omitempty" name:"MidRisk"`

	// 低风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowRisk *uint64 `json:"LowRisk,omitempty" name:"LowRisk"`

	// 提示风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfoRisk *uint64 `json:"InfoRisk,omitempty" name:"InfoRisk"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FendTime *string `json:"FendTime,omitempty" name:"FendTime"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FgroupId *uint64 `json:"FgroupId,omitempty" name:"FgroupId"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 监控类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorType *uint64 `json:"MonitorType,omitempty" name:"MonitorType"`

	// 内容风险1
	// 注意：此字段可能返回 null，表示取不到有效值。
	FyellowRisk *uint64 `json:"FyellowRisk,omitempty" name:"FyellowRisk"`

	// 内容风险2
	// 注意：此字段可能返回 null，表示取不到有效值。
	FpoliticalRisk *uint64 `json:"FpoliticalRisk,omitempty" name:"FpoliticalRisk"`

	// 内容风险3
	// 注意：此字段可能返回 null，表示取不到有效值。
	FdrugRisk *uint64 `json:"FdrugRisk,omitempty" name:"FdrugRisk"`

	// 内容风险4
	// 注意：此字段可能返回 null，表示取不到有效值。
	FhorrorRisk *uint64 `json:"FhorrorRisk,omitempty" name:"FhorrorRisk"`

	// 内容风险5
	// 注意：此字段可能返回 null，表示取不到有效值。
	FportLimit *uint64 `json:"FportLimit,omitempty" name:"FportLimit"`

	// 内容风险6
	// 注意：此字段可能返回 null，表示取不到有效值。
	FgambleRisk *uint64 `json:"FgambleRisk,omitempty" name:"FgambleRisk"`

	// 进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *uint64 `json:"Process,omitempty" name:"Process"`

	// 扫描速度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanSpeed *string `json:"ScanSpeed,omitempty" name:"ScanSpeed"`
}

type TaskLogDetailType struct {

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FpocName *string `json:"FpocName,omitempty" name:"FpocName"`

	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	FvulLevel *string `json:"FvulLevel,omitempty" name:"FvulLevel"`

	// 资产列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ftarget *string `json:"Ftarget,omitempty" name:"Ftarget"`

	// 扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinsertTime *string `json:"FinsertTime,omitempty" name:"FinsertTime"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	FvulStatus *uint64 `json:"FvulStatus,omitempty" name:"FvulStatus"`
}

type TaskLogInfo struct {

	// 具体的日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlogId *string `json:"FlogId,omitempty" name:"FlogId"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaddTime *string `json:"FaddTime,omitempty" name:"FaddTime"`
}

type TaskObject struct {

	// 名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Target *string `json:"Target,omitempty" name:"Target"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 频率
	// 注意：此字段可能返回 null，表示取不到有效值。
	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`

	// 上次扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastStartTime *string `json:"LastStartTime,omitempty" name:"LastStartTime"`

	// 开始时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 监控类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorType *uint64 `json:"MonitorType,omitempty" name:"MonitorType"`

	// 扫描速度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanSpeed *string `json:"ScanSpeed,omitempty" name:"ScanSpeed"`

	// 扫描内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectType *string `json:"DetectType,omitempty" name:"DetectType"`

	// 日志的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fstatus *uint64 `json:"Fstatus,omitempty" name:"Fstatus"`

	// 日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlogId *string `json:"FlogId,omitempty" name:"FlogId"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 扫描深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanDeep *string `json:"ScanDeep,omitempty" name:"ScanDeep"`
}

type UserObject struct {

	// 通知等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyLevel []*string `json:"NotifyLevel,omitempty" name:"NotifyLevel" list`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 1新用户 0付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isnew *uint64 `json:"Isnew,omitempty" name:"Isnew"`

	// 0为该用户未添加过资产；1为已添加过资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	HaveCreatedAsset *uint64 `json:"HaveCreatedAsset,omitempty" name:"HaveCreatedAsset"`
}

type VulDetail struct {

	// 当前任务ID要扫描的目标(域名/IP)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Target *string `json:"Target,omitempty" name:"Target"`

	// 扫描结果来源 有webscan/pocscan/csscan/weakscan 四种
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 存在漏洞的URL，供四种扫描结果使用，webscan格式为完整url，pocsuite、csscan为输入的target（若可获取完整URL，则输出完整URL），weakscan格式为 协议://IP:PORT
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulUrl *string `json:"VulUrl,omitempty" name:"VulUrl"`

	// 漏洞类型，如xss、sqli等，供四种扫描结果使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 漏洞状态，正常状态为0，已处理为1，待验证为2，验证中为3，忽略为4，误报为5
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulStatus *int64 `json:"VulStatus,omitempty" name:"VulStatus"`

	// 漏洞名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PocName *string `json:"PocName,omitempty" name:"PocName"`

	// 漏洞影响组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 漏洞影响版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 漏洞影响等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// CVE 编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cve *string `json:"Cve,omitempty" name:"Cve"`

	// 漏洞描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 参考链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	References *string `json:"References,omitempty" name:"References"`

	// 漏洞修复方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulRepair *string `json:"VulRepair,omitempty" name:"VulRepair"`

	// poc漏洞分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCategory *string `json:"VulCategory,omitempty" name:"VulCategory"`

	// 漏洞发现时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 漏洞中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 漏洞影响等级（新）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 漏洞描述（新）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 漏洞修复建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggest *string `json:"Suggest,omitempty" name:"Suggest"`

	// 漏洞影响内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Impact *string `json:"Impact,omitempty" name:"Impact"`

	// 漏洞证明内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

type WhiteUrl struct {

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fid *uint64 `json:"Fid,omitempty" name:"Fid"`

	// URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Furl *string `json:"Furl,omitempty" name:"Furl"`
}
