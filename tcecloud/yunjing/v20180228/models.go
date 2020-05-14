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

package v20180228

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Account struct {

	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 帐号名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 帐号所属组。
	Groups *string `json:"Groups,omitempty" name:"Groups"`

	// 帐号类型。
	// <li>ORDINARY：普通帐号</li>
	// <li>SUPPER：超级管理员帐号</li>
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// 帐号创建时间。
	AccountCreateTime *string `json:"AccountCreateTime,omitempty" name:"AccountCreateTime"`

	// 帐号最后登录时间。
	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
}

type AccountStatistics struct {

	// 用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type AddLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单规则
	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLoginWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLoginWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddMachineTagRequest struct {
	*tchttp.BaseRequest

	// 云服务器ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 标签ID
	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`

	// 云服务器地区
	MRegion *string `json:"MRegion,omitempty" name:"MRegion"`

	// 云服务器类型(CVM|BM)
	MArea *string `json:"MArea,omitempty" name:"MArea"`
}

func (r *AddMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddMachineTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddMachineTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AgentVul struct {

	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 漏洞名称。
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞危害等级。
	// <li>HIGH：高危</li>
	// <li>MIDDLE：中危</li>
	// <li>LOW：低危</li>
	// <li>NOTICE：提示</li>
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// 最后扫描时间。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 漏洞描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 漏洞状态。
	// <li>UN_OPERATED : 待处理</li>
	// <li>FIXED : 已修复</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type BanWhiteList struct {

	// 白名单ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 白名单别名。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 阻断来源IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 修改白名单时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 创建白名单时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 白名单所属机器。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 白名单是否全局
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

type BashEvent struct {

	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 执行用户名
	User *string `json:"User,omitempty" name:"User"`

	// 平台类型
	Platform *uint64 `json:"Platform,omitempty" name:"Platform"`

	// 执行命令
	BashCmd *string `json:"BashCmd,omitempty" name:"BashCmd"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则等级
	RuleLevel *uint64 `json:"RuleLevel,omitempty" name:"RuleLevel"`

	// 处理状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
}

type BashRule struct {

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 危险等级(1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 正则表达式
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 规则描述
	Decription *string `json:"Decription,omitempty" name:"Decription"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态 (0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
}

type BruteAttack struct {

	// 事件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 破解事件状态
	// <li>BRUTEATTACK_FAIL_ACCOUNT： 暴力破解事件-失败(存在帐号)  </li>
	// <li>BRUTEATTACK_FAIL_NOACCOUNT：暴力破解事件-失败(帐号不存在)</li>
	// <li>BRUTEATTACK_SUCCESS：暴力破解事件-成功</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 用户名称。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 城市ID。
	City *uint64 `json:"City,omitempty" name:"City"`

	// 国家ID。
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 省份ID。
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// 来源IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 尝试破解次数。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 发生时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 是否专业版。
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// 阻断状态。
	BanStatus *string `json:"BanStatus,omitempty" name:"BanStatus"`

	// 机器UUID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type ChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type CloseProVersionRequest struct {
	*tchttp.BaseRequest

	// 主机唯一标识Uuid。
	// 黑石的InstanceId，CVM的Uuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *CloseProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseProVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseProVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseProVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Component struct {

	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 组件版本号。
	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`

	// 组件类型。
	// <li>SYSTEM：系统组件</li>
	// <li>WEB：Web组件</li>
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 组件名称。
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 组件检测更新时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type ComponentStatistics struct {

	// 组件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`

	// 组件名称。
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 组件类型。
	// <li>WEB：Web组件</li>
	// <li>SYSTEM：系统组件</li>
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 组件描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateBanRecordRequest struct {
	*tchttp.BaseRequest

	// 白名单信息
	Record *WhiteListRecord `json:"Record,omitempty" name:"Record"`
}

func (r *CreateBanRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBanRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBanRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBanRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBanRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBanWhiteListRequest struct {
	*tchttp.BaseRequest

	// 阻断规则
	Rules *BanWhiteList `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateBanWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBanWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBanWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBanWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBanWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOpenPortTaskRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateOpenPortTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOpenPortTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOpenPortTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOpenPortTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOpenPortTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProcessTaskRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateProcessTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProcessTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProcessTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProcessTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProcessTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端UUID数组。
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids" list`

	// 登录地域信息数组。
	Places []*Place `json:"Places,omitempty" name:"Places" list`
}

func (r *CreateUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsualLoginPlacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsualLoginPlacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DefendAttackLog struct {

	// 日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 来源IP
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 来源端口
	SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`

	// 攻击方式
	HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`

	// 攻击描述
	HttpCgi *string `json:"HttpCgi,omitempty" name:"HttpCgi"`

	// 攻击参数
	HttpParam *string `json:"HttpParam,omitempty" name:"HttpParam"`

	// 威胁类型
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 攻击时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 目标服务器IP
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 目标服务器名称
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 攻击内容
	HttpContent *string `json:"HttpContent,omitempty" name:"HttpContent"`
}

type DeleteAttackLogsRequest struct {
	*tchttp.BaseRequest

	// 日志ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAttackLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAttackLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBanListRequest struct {
	*tchttp.BaseRequest

	// 删除白名单的ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteBanListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBanListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBanListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBanListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBanListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBanWhiteListRequest struct {
	*tchttp.BaseRequest

	// 要删除的白名单ID列表
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteBanWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBanWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBanWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBanWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBanWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBashEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBashEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBashEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBashRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBashRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBashRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 暴力破解事件Id数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBruteAttacksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBruteAttacksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoginWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoginWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMachineRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMachineResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineTagRequest struct {
	*tchttp.BaseRequest

	// 关联的标签ID
	Rid *uint64 `json:"Rid,omitempty" name:"Rid"`
}

func (r *DeleteMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMachineTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMachineTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// 恶意请求记录ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMaliciousRequestsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMaliciousRequestsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马记录ID数组
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMalwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMalwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 异地登录事件ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNonlocalLoginPlacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNonlocalLoginPlacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeletePrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePrivilegeEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePrivilegeEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeletePrivilegeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePrivilegeRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivilegeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePrivilegeRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteReverseShellEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteReverseShellEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteReverseShellRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteReverseShellRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReverseShellRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteReverseShellRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTagsRequest struct {
	*tchttp.BaseRequest

	// 标签ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 已添加常用登录地城市ID数组
	CityIds []*uint64 `json:"CityIds,omitempty" name:"CityIds" list`
}

func (r *DeleteUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUsualLoginPlacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUsualLoginPlacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbTestStatusRequest struct {
	*tchttp.BaseRequest

	// 灰度ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeAbTestStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAbTestStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbTestStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该用户当前是否正在灰度。
		Status *bool `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbTestStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAbTestStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatisticsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号用户名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAccountStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 帐号统计列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 帐号统计列表。
		AccountStatistics []*AccountStatistics `json:"AccountStatistics,omitempty" name:"AccountStatistics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Uuid表示，查询该主机下列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Username表示，查询该用户名下列表信息。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	// <li>Privilege - String - 是否必填：否 - 帐号类型（ORDINARY: 普通帐号 | SUPPER: 超级管理员帐号）</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 帐号列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 帐号数据列表。
		Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentVulsRequest struct {
	*tchttp.BaseRequest

	// 漏洞类型。
	// <li>WEB: Web应用漏洞</li>
	// <li>SYSTEM：系统组件漏洞</li>
	// <li>BASELINE：安全基线</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 客户端UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAgentVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentVulsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentVulsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 主机漏洞信息
		AgentVuls []*AgentVul `json:"AgentVuls,omitempty" name:"AgentVuls" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentVulsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmAttributeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 防护软件离线告警状态：
	// <li>OPEN：告警已开启</li>
	// <li>CLOSE： 告警已关闭</li>
		Offline *string `json:"Offline,omitempty" name:"Offline"`

		// 发现木马告警状态：
	// <li>OPEN：告警已开启</li>
	// <li>CLOSE： 告警已关闭</li>
		Malware *string `json:"Malware,omitempty" name:"Malware"`

		// 发现异地登录告警状态：
	// <li>OPEN：告警已开启</li>
	// <li>CLOSE： 告警已关闭</li>
		NonlocalLogin *string `json:"NonlocalLogin,omitempty" name:"NonlocalLogin"`

		// 被暴力破解成功告警状态：
	// <li>OPEN：告警已开启</li>
	// <li>CLOSE： 告警已关闭</li>
		CrackSuccess *string `json:"CrackSuccess,omitempty" name:"CrackSuccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllMachinesRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。
	// 如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
}

func (r *DescribeAllMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllMachinesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 主机信息数组。
		Machines []*FullMachine `json:"Machines,omitempty" name:"Machines" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllMachinesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogInfoRequest struct {
	*tchttp.BaseRequest

	// 日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAttackLogInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttackLogInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志ID
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 主机ID
		Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

		// 攻击来源端口
		SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`

		// 攻击来源IP
		SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

		// 攻击目标端口
		DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

		// 攻击目标IP
		DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

		// 攻击方法
		HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`

		// 攻击目标主机
		HttpHost *string `json:"HttpHost,omitempty" name:"HttpHost"`

		// 攻击头信息
		HttpHead *string `json:"HttpHead,omitempty" name:"HttpHead"`

		// 攻击者浏览器标识
		HttpUserAgent *string `json:"HttpUserAgent,omitempty" name:"HttpUserAgent"`

		// 请求源
		HttpReferer *string `json:"HttpReferer,omitempty" name:"HttpReferer"`

		// 威胁类型
		VulType *string `json:"VulType,omitempty" name:"VulType"`

		// 攻击路径
		HttpCgi *string `json:"HttpCgi,omitempty" name:"HttpCgi"`

		// 攻击参数
		HttpParam *string `json:"HttpParam,omitempty" name:"HttpParam"`

		// 攻击时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 攻击内容
		HttpContent *string `json:"HttpContent,omitempty" name:"HttpContent"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackLogInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttackLogInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttackLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志列表
		AttackLogs []*DefendAttackLog `json:"AttackLogs,omitempty" name:"AttackLogs" list`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttackLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoQuaraInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAutoQuaraInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoQuaraInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoQuaraInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否开启
		AppidQuaraAll *uint64 `json:"AppidQuaraAll,omitempty" name:"AppidQuaraAll"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoQuaraInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoQuaraInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBanListRequest struct {
	*tchttp.BaseRequest

	// 筛选条件：源IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBanListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBanListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBanListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBanListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBanListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBanStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBanStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBanStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBanStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 阻断开关状态 0:关闭 1:开启
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBanStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBanStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBanWhiteListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeBanWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBanWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBanWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 白名单列表
		WhiteList []*BanWhiteList `json:"WhiteList,omitempty" name:"WhiteList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBanWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBanWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机内网IP)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBashEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 高危命令事件列表
		List []*BashEvent `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBashEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBashRulesRequest struct {
	*tchttp.BaseRequest

	// 0-系统规则; 1-用户规则
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(规则名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBashRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表内容
		List []*BashRule `json:"List,omitempty" name:"List" list`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBashRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBlockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBlockStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 开启：1，关闭：0
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 阻断功能是否可见
		IsShowBan *bool `json:"IsShowBan,omitempty" name:"IsShowBan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBlockStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 -  查询关键字</li>
	// <li>Status - String - 是否必填：否 -  查询状态（FAILED：破解失败 |SUCCESS：破解成功）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBruteAttacksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 暴力破解事件列表
		BruteAttacks []*BruteAttack `json:"BruteAttacks,omitempty" name:"BruteAttacks" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBruteAttacksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentInfoRequest struct {
	*tchttp.BaseRequest

	// 组件ID。
	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`
}

func (r *DescribeComponentInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComponentInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组件ID。
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 组件名称。
		ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

		// 组件类型。
	// <li>WEB：web组件</li>
	// <li>SYSTEM：系统组件</li>
		ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

		// 组件官网。
		Homepage *string `json:"Homepage,omitempty" name:"Homepage"`

		// 组件描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComponentInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentStatisticsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// ComponentName - String - 是否必填：否 - 组件名称
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeComponentStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComponentStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组件统计列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 组件统计列表数据数组。
		ComponentStatistics []*ComponentStatistics `json:"ComponentStatistics,omitempty" name:"ComponentStatistics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComponentStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentsRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。Uuid和ComponentId必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 组件ID。Uuid和ComponentId必填其一，使用ComponentId表示，查询该组件列表信息。
	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentVersion - String - 是否必填：否 - 组件版本号</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComponentsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组件列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 组件列表数据。
		Components []*Component `json:"Components,omitempty" name:"Components" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeComponentsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEventListRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>VulId- int - 是否必填：否 - 事件id</li>
	// <li>VulCategory- int - 是否必填：否 - 漏洞分类筛选 1: web应用漏洞 2:系统组件漏洞 3:安全基线 4: Linux系统漏洞 5: windows补丁</li>
	// <li>IfEmergency - String - 是否必填：否 - 是否为应急漏洞，查询应急漏洞传:yes</li>
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEventListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取事件列表
		EventList []*VulEventList `json:"EventList,omitempty" name:"EventList" list`

		// 获取事件总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEventListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEventlogRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>DateRange - String - 是否必填：否 - 日志时间范围 </li>
	// <li>SeverityClass - String  是否必填：否 - 威胁等级(10:严重 20:高危 30:中危 40:低危)</li>
	// <li>EventType - String  是否必填：否 - 事件类型</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeEventlogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEventlogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEventlogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 事件列表数据
		List []*Eventlog `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventlogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEventlogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryAccountsRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeHistoryAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHistoryAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 帐号变更历史列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 帐号变更历史数据数组。
		HistoryAccounts []*HistoryAccount `json:"HistoryAccounts,omitempty" name:"HistoryAccounts" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHistoryAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImpactedHostsRequest struct {
	*tchttp.BaseRequest

	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED：待处理 | FIXED：已修复）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImpactedHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 漏洞影响机器列表数组
		ImpactedHosts []*ImpactedHost `json:"ImpactedHosts,omitempty" name:"ImpactedHosts" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImpactedHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoginWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 异地登录白名单数组
		LoginWhiteLists []*LoginWhiteLists `json:"LoginWhiteLists,omitempty" name:"LoginWhiteLists" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoginWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineInfoRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMachineInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器ip。
		MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

		// 受云镜保护天数。
		ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`

		// 操作系统。
		MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`

		// 主机名称。
		MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

		// 在线状态。
	// <li>ONLINE： 在线</li>
	// <li>OFFLINE：离线</li>
		MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`

		// CVM或BM主机唯一标识。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 主机外网IP。
		MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

		// CVM或BM主机唯一Uuid。
		Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

		// 云镜客户端唯一Uuid。
		Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

		// 是否开通专业版。
	// <li>true：是</li>
	// <li>false：否</li>
		IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

		// 专业版开通时间。
		ProVersionOpenDate *string `json:"ProVersionOpenDate,omitempty" name:"ProVersionOpenDate"`

		// 云主机类型。
	// <li>CVM: 虚拟主机</li>
	// <li>BM: 黑石物理机</li>
		MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

		// 机器所属地域。如：ap-guangzhou，ap-shanghai
		MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

		// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>
		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

		// 免费木马剩余检测数量。
		FreeMalwaresLeft *uint64 `json:"FreeMalwaresLeft,omitempty" name:"FreeMalwaresLeft"`

		// 免费漏洞剩余检测数量。
		FreeVulsLeft *uint64 `json:"FreeVulsLeft,omitempty" name:"FreeVulsLeft"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMachineInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMachineRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMachineRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CVM地域列表
		CVM []*RegionInfo `json:"CVM,omitempty" name:"CVM" list`

		// BM地域列表
		BM []*RegionInfo `json:"BM,omitempty" name:"BM" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMachineRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMachinesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主机列表
		Machines []*Machine `json:"Machines,omitempty" name:"Machines" list`

		// 主机数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMachinesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | TRUSTED：已信任 | UN_TRUSTED：已取消信任）</li>
	// <li>Domain - String - 是否必填：否 - 恶意请求的域名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 云镜客户端唯一UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMaliciousRequestsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 恶意请求记录数组。
		MaliciousRequests []*MaliciousRequest `json:"MaliciousRequests,omitempty" name:"MaliciousRequests" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMaliciousRequestsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwaresRequest struct {
	*tchttp.BaseRequest

	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 木马状态（UN_OPERATED: 未处理 | SEGREGATED: 已隔离|TRUSTED：信任）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMalwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 木马总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Malware数组。
		Malwares []*Malware `json:"Malwares,omitempty" name:"Malwares" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMalwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 -  查询关键字</li>
	// <li>Status - String - 是否必填：否 -  登录状态（NON_LOCAL_LOGIN: 异地登录 | NORMAL_LOGIN : 正常登录）</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNonlocalLoginPlacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 异地登录信息数组。
		NonLocalLoginPlaces []*NonLocalLoginPlace `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNonlocalLoginPlacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortStatisticsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeOpenPortStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOpenPortStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 端口统计列表总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 端口统计数据列表
		OpenPortStatistics []*OpenPortStatistics `json:"OpenPortStatistics,omitempty" name:"OpenPortStatistics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOpenPortStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeOpenPortTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOpenPortTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态。
	// <li>COMPLETE：完成（此时可以调用DescribeOpenPorts接口获取实时进程列表）</li>
	// <li>AGENT_OFFLINE：云镜客户端离线</li>
	// <li>COLLECTING：端口获取中</li>
	// <li>FAILED：进程获取失败</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOpenPortTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortsRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。Port和Uuid必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 开放端口号。Port和Uuid必填其一，使用Port表示查询该端口的列表信息。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeOpenPortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOpenPortsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 端口列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 端口列表。
		OpenPorts []*OpenPort `json:"OpenPorts,omitempty" name:"OpenPorts" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOpenPortsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOverviewStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviewStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器在线数。
		OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`

		// 专业服务器数。
		ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`

		// 木马文件数。
		MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

		// 异地登录数。
		NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`

		// 暴力破解成功数。
		BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`

		// 漏洞数。
		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

		// 安全基线数。
		BaseLineNum *uint64 `json:"BaseLineNum,omitempty" name:"BaseLineNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviewStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否开启木马自动隔离(0:关闭 1:开启)
		MalwareAutoIsolate *int64 `json:"MalwareAutoIsolate,omitempty" name:"MalwareAutoIsolate"`

		// 是否开启木马定时扫描(0:关闭 1:开启)
		MalwareEnableDiskScan *int64 `json:"MalwareEnableDiskScan,omitempty" name:"MalwareEnableDiskScan"`

		// 设置木马定时扫描间隔(天数)
		MalwareDiskScanInterDay *int64 `json:"MalwareDiskScanInterDay,omitempty" name:"MalwareDiskScanInterDay"`

		// 设置木马定时扫描时间(0-23整点时间)
		MalwareDiskScanTime *int64 `json:"MalwareDiskScanTime,omitempty" name:"MalwareDiskScanTime"`

		// 设置木马扫描全盘扫描忽略路径(linux), 换行符分割多个
		MalwareIgnoreWhiteListLinux *string `json:"MalwareIgnoreWhiteListLinux,omitempty" name:"MalwareIgnoreWhiteListLinux"`

		// 设置木马扫描全盘扫描忽略路径(windows), 换行符分割多个
		MalwareIgnoreWhiteListWindows *string `json:"MalwareIgnoreWhiteListWindows,omitempty" name:"MalwareIgnoreWhiteListWindows"`

		// 是否开启漏洞定时扫描(0:关闭 1:开启)
		VulEnableVulnerDetect *int64 `json:"VulEnableVulnerDetect,omitempty" name:"VulEnableVulnerDetect"`

		// 设置漏洞定时扫描间隔(天数)
		VulVulnerDetectDay *int64 `json:"VulVulnerDetectDay,omitempty" name:"VulVulnerDetectDay"`

		// 设置漏洞定时扫描时间(0-23整点时间)
		VulVulnerDetectTime *int64 `json:"VulVulnerDetectTime,omitempty" name:"VulVulnerDetectTime"`

		// 设置木马扫描实时监控忽略路径(linux), 换行符分割多个
		MalwareMonitorIgnoreWhiteListLinux *string `json:"MalwareMonitorIgnoreWhiteListLinux,omitempty" name:"MalwareMonitorIgnoreWhiteListLinux"`

		// 设置木马扫描实时监控忽略路径(windows), 换行符分割多个
		MalwareMonitorIgnoreWhiteListWindows *string `json:"MalwareMonitorIgnoreWhiteListWindows,omitempty" name:"MalwareMonitorIgnoreWhiteListWindows"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机IP)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribePrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePrivilegeEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据列表
		List []*PrivilegeEscalationProcess `json:"List,omitempty" name:"List" list`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePrivilegeEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeRulesRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribePrivilegeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePrivilegeRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表内容
		List []*PrivilegeRule `json:"List,omitempty" name:"List" list`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivilegeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePrivilegeRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProVersionInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProVersionInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 后付费昨日扣费
		PostPayCost *uint64 `json:"PostPayCost,omitempty" name:"PostPayCost"`

		// 新增主机是否自动开通专业版
		IsAutoOpenProVersion *bool `json:"IsAutoOpenProVersion,omitempty" name:"IsAutoOpenProVersion"`

		// 开通专业版主机数
		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProVersionInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProVersionInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionStatusRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端UUID、填写"all"表示所有主机。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeProVersionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProVersionStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 开通状态。
	// <li>UNOPENED：未开通专业版</li>
	// <li>OPENED：已开通专业版</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProVersionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProVersionStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessStatisticsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcessStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 进程统计列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 进程统计列表数据数组。
		ProcessStatistics []*ProcessStatistics `json:"ProcessStatistics,omitempty" name:"ProcessStatistics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcessStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeProcessTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcessTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态。
	// <li>COMPLETE：完成（此时可以调用DescribeProcesses接口获取实时进程列表）</li>
	// <li>AGENT_OFFLINE：云镜客户端离线</li>
	// <li>COLLECTING：进程获取中</li>
	// <li>FAILED：进程获取失败</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcessTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessesRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。Uuid和ProcessName必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名。Uuid和ProcessName必填其一，使用ProcessName表示，查询该进程列表信息。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeProcessesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcessesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 进程列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 进程列表数据数组。
		Processes []*Process `json:"Processes,omitempty" name:"Processes" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcessesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(主机内网IP|进程名)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReverseShellEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表内容
		List []*ReverseShell `json:"List,omitempty" name:"List" list`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReverseShellEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellRulesRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeReverseShellRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReverseShellRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表内容
		List []*ReverseShellRule `json:"List,omitempty" name:"List" list`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReverseShellRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSecurityDynamicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityDynamicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全事件消息数组。
		SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics,omitempty" name:"SecurityDynamics" list`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityDynamicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityDynamicsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendsRequest struct {
	*tchttp.BaseRequest

	// 开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 结束时间。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeSecurityTrendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityTrendsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 木马事件统计数据数组。
		Malwares []*SecurityTrend `json:"Malwares,omitempty" name:"Malwares" list`

		// 异地登录事件统计数据数组。
		NonLocalLoginPlaces []*SecurityTrend `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces" list`

		// 密码破解事件统计数据数组。
		BruteAttacks []*SecurityTrend `json:"BruteAttacks,omitempty" name:"BruteAttacks" list`

		// 漏洞统计数据数组。
		Vuls []*SecurityTrend `json:"Vuls,omitempty" name:"Vuls" list`

		// 基线统计数据数组。
		BaseLines []*SecurityTrend `json:"BaseLines,omitempty" name:"BaseLines" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityTrendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityTrendsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStatRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStatRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStatResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器在线数。
		OnlineNum *uint64 `json:"OnlineNum,omitempty" name:"OnlineNum"`

		// 专业服务器数。
		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

		// 病毒库更新时间。
		VdbUpdateTime *string `json:"VdbUpdateTime,omitempty" name:"VdbUpdateTime"`

		// 漏洞库更新时间。
		PocUpdateTime *string `json:"PocUpdateTime,omitempty" name:"PocUpdateTime"`

		// 木马事件统计。
		MalwareStat *EventStat `json:"MalwareStat,omitempty" name:"MalwareStat"`

		// 异地事件统计。
		HostLoginStat *EventStat `json:"HostLoginStat,omitempty" name:"HostLoginStat"`

		// 爆破事件统计。
		BruteAttackStat *EventStat `json:"BruteAttackStat,omitempty" name:"BruteAttackStat"`

		// 恶意请求事件统计。
		MaliciousRequestStat *EventStat `json:"MaliciousRequestStat,omitempty" name:"MaliciousRequestStat"`

		// 本地提权事件统计。
		PrivilegeStat *EventStat `json:"PrivilegeStat,omitempty" name:"PrivilegeStat"`

		// 反弹Shell事件统计。
		ReverseShellStat *EventStat `json:"ReverseShellStat,omitempty" name:"ReverseShellStat"`

		// 高危指令事件统计。
		HighRiskBashStat *EventStat `json:"HighRiskBashStat,omitempty" name:"HighRiskBashStat"`

		// 高危漏洞事件统计。
		VulHighStat *EventStat `json:"VulHighStat,omitempty" name:"VulHighStat"`

		// 中危漏洞事件统计。
		VulNormalStat *EventStat `json:"VulNormalStat,omitempty" name:"VulNormalStat"`

		// 低危漏洞事件统计。
		VulLowStat *EventStat `json:"VulLowStat,omitempty" name:"VulLowStat"`

		// 高危基线漏洞事件统计。
		BaselineHighStat *EventStat `json:"BaselineHighStat,omitempty" name:"BaselineHighStat"`

		// 中危基线漏事件统计。
		BaselineNormalStat *EventStat `json:"BaselineNormalStat,omitempty" name:"BaselineNormalStat"`

		// 低危基线漏事件统计。
		BaselineLowStat *EventStat `json:"BaselineLowStat,omitempty" name:"BaselineLowStat"`

		// 网络攻击事件统计。
		AttackLogsStat *EventStat `json:"AttackLogsStat,omitempty" name:"AttackLogsStat"`

		// 有未处理安全事件的机器总数。
		MachineTotalAffectNum *uint64 `json:"MachineTotalAffectNum,omitempty" name:"MachineTotalAffectNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStatResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagMachinesRequest struct {
	*tchttp.BaseRequest

	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTagMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagMachinesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表数据
		List []*TagMachine `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagMachinesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表信息
		List []*Tag `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUndoVulCountsRequest struct {
	*tchttp.BaseRequest

	// 漏洞分类，最小值为1，最大值为5
	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`

	// 应急漏洞筛选
	IfEmergency *string `json:"IfEmergency,omitempty" name:"IfEmergency"`
}

func (r *DescribeUndoVulCountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUndoVulCountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUndoVulCountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取未处理的漏洞数
		UndoVulCount *uint64 `json:"UndoVulCount,omitempty" name:"UndoVulCount"`

		// 获取未处理的主机数
		UndoHostCount *int64 `json:"UndoHostCount,omitempty" name:"UndoHostCount"`

		// 获取普通版主机数
		NotProfessionCount *uint64 `json:"NotProfessionCount,omitempty" name:"NotProfessionCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUndoVulCountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUndoVulCountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端UUID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsualLoginPlacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 常用登录地数组
		UsualLoginPlaces []*UsualPlace `json:"UsualLoginPlaces,omitempty" name:"UsualLoginPlaces" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsualLoginPlacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVersionStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVersionStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基础版数量
		BasicVersionNum *uint64 `json:"BasicVersionNum,omitempty" name:"BasicVersionNum"`

		// 专业版数量
		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVersionStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVersionStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulCountByDatesRequest struct {
	*tchttp.BaseRequest

	// 需要查询最近几天的数据
	LastDays []*uint64 `json:"LastDays,omitempty" name:"LastDays" list`

	// 漏洞的分类，最小值为1最大值为5
	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`

	// 是否为应急漏洞筛选
	IfEmergency *string `json:"IfEmergency,omitempty" name:"IfEmergency"`
}

func (r *DescribeVulCountByDatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulCountByDatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulCountByDatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量获得对应天数的漏洞数量
		VulCount []*uint64 `json:"VulCount,omitempty" name:"VulCount" list`

		// 批量获得对应天数的主机数量
		HostCount []*uint64 `json:"HostCount,omitempty" name:"HostCount" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulCountByDatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulCountByDatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulInfoRequest struct {
	*tchttp.BaseRequest

	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeVulInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 漏洞种类ID。
		VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

		// 漏洞名称。
		VulName *string `json:"VulName,omitempty" name:"VulName"`

		// 漏洞等级。
		VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

		// 漏洞类型。
		VulType *string `json:"VulType,omitempty" name:"VulType"`

		// 漏洞描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 修复方案。
		RepairPlan *string `json:"RepairPlan,omitempty" name:"RepairPlan"`

		// 漏洞CVE。
		CveId *string `json:"CveId,omitempty" name:"CveId"`

		// 参考链接。
		Reference *string `json:"Reference,omitempty" name:"Reference"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulListRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>VulCategory - int - 是否必填：否 - 漏洞分类筛选 1: web应用漏洞 2:系统组件漏洞 3:安全基线</li>
	// <li>IfEmergency - String - 是否必填：否 - 是否为应急漏洞，查询应急漏洞传:yes</li>
	// <li>Status - String - 是否必填：是 - 漏洞状态筛选，0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态</li>
	// <li>Level - String - 是否必填：否 - 漏洞等级筛选 1:低 2:中 3:高 4:提示</li>
	// <li>VulName- String - 是否必填：否 - 漏洞名称搜索</li>
	Filters []*Filters `json:"Filters,omitempty" name:"Filters" list`
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

		// 漏洞列表
		VulInfoList []*VulInfoList `json:"VulInfoList,omitempty" name:"VulInfoList" list`

		// 漏洞总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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

type DescribeVulScanResultRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulScanResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulScanResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 漏洞数量。
		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

		// 专业版机器数。
		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

		// 受影响的专业版主机数。
		ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitempty" name:"ImpactedHostNum"`

		// 主机总数。
		HostNum *uint64 `json:"HostNum,omitempty" name:"HostNum"`

		// 基础版机器数。
		BasicVersionNum *uint64 `json:"BasicVersionNum,omitempty" name:"BasicVersionNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulScanResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulsRequest struct {
	*tchttp.BaseRequest

	// 漏洞类型。
	// <li>WEB：Web应用漏洞</li>
	// <li>SYSTEM：系统组件漏洞</li>
	// <li>BASELINE：安全基线</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）
	// 
	// Status过滤条件值只能取其一，不能是“或”逻辑。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVulsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 漏洞数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 漏洞列表数组。
		Vuls []*Vul `json:"Vuls,omitempty" name:"Vuls" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVulsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportBruteAttacksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专业周报密码破解数组。
		WeeklyReportBruteAttacks []*WeeklyReportBruteAttack `json:"WeeklyReportBruteAttacks,omitempty" name:"WeeklyReportBruteAttacks" list`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportBruteAttacksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportInfoRequest struct {
	*tchttp.BaseRequest

	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
}

func (r *DescribeWeeklyReportInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账号所属公司或个人名称。
		CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

		// 机器总数。
		MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`

		// 云镜客户端在线数。
		OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`

		// 云镜客户端离线数。
		OfflineMachineNum *uint64 `json:"OfflineMachineNum,omitempty" name:"OfflineMachineNum"`

		// 开通云镜专业版数量。
		ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`

		// 周报开始时间。
		BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

		// 周报结束时间。
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// 安全等级。
	// <li>HIGH：高</li>
	// <li>MIDDLE：中</li>
	// <li>LOW：低</li>
		Level *string `json:"Level,omitempty" name:"Level"`

		// 木马记录数。
		MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

		// 异地登录数。
		NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`

		// 密码破解成功数。
		BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`

		// 漏洞数。
		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

		// 导出文件下载地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportMalwaresRequest struct {
	*tchttp.BaseRequest

	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportMalwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专业周报木马数据。
		WeeklyReportMalwares []*WeeklyReportMalware `json:"WeeklyReportMalwares,omitempty" name:"WeeklyReportMalwares" list`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportMalwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 专业周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专业周报异地登录数据。
		WeeklyReportNonlocalLoginPlaces []*WeeklyReportNonlocalLoginPlace `json:"WeeklyReportNonlocalLoginPlaces,omitempty" name:"WeeklyReportNonlocalLoginPlaces" list`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportVulsRequest struct {
	*tchttp.BaseRequest

	// 专业版周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportVulsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportVulsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专业周报漏洞数据数组。
		WeeklyReportVuls []*WeeklyReportVul `json:"WeeklyReportVuls,omitempty" name:"WeeklyReportVuls" list`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportVulsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专业周报列表数组。
		WeeklyReports []*WeeklyReport `json:"WeeklyReports,omitempty" name:"WeeklyReports" list`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWeeklyReportsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditBashRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 危险等级(1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 正则表达式
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 规则ID(新增时不填)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

func (r *EditBashRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditBashRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditBashRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditBashRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditBashRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditPrivilegeRuleRequest struct {
	*tchttp.BaseRequest

	// 规则ID(新增时请留空)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否S权限进程
	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

func (r *EditPrivilegeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditPrivilegeRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditPrivilegeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditPrivilegeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditPrivilegeRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditReverseShellRuleRequest struct {
	*tchttp.BaseRequest

	// 规则ID(新增时请留空)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID(IsGlobal为1时，Uuid或Hostip必填一个)
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机IP(IsGlobal为1时，Uuid或Hostip必填一个)
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 目标IP
	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`

	// 目标端口
	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否全局规则(默认否)
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

func (r *EditReverseShellRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditReverseShellRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditReverseShellRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditReverseShellRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditReverseShellRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditTagsRequest struct {
	*tchttp.BaseRequest

	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CVM主机ID
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids" list`
}

func (r *EditTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EventStat struct {

	// 事件数
	EventsNum *uint64 `json:"EventsNum,omitempty" name:"EventsNum"`

	// 受影响的主机数
	MachineAffectNum *uint64 `json:"MachineAffectNum,omitempty" name:"MachineAffectNum"`
}

type Eventlog struct {

	// 事件ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 云镜ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 事件类型
	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 威胁等级
	SeverityClass *uint64 `json:"SeverityClass,omitempty" name:"SeverityClass"`

	// 详情原文
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 发生时间
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 关键信息
	Memo *string `json:"Memo,omitempty" name:"Memo"`
}

type ExportAttackLogsRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportAttackLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportAttackLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportBashEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportBashEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttacksRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportBruteAttacksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportBruteAttacksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportEventlogRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportEventlogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportEventlogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportEventlogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportEventlogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportEventlogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportMaliciousRequestsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportMaliciousRequestsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportMalwaresRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportMalwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportMalwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportPrivilegeEventsRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportPrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportPrivilegeEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportPrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportPrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportPrivilegeEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportReverseShellEventsRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportReverseShellEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportReverseShellEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Filters struct {

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type FullMachine struct {

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机外网IP。
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// 主机唯一标识Uuid。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 是否是专业版
	// * true：是
	// * false：否
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
}

type HistoryAccount struct {

	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 帐号名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 帐号变更类型。
	// <li>CREATE：表示新增帐号</li>
	// <li>MODIFY：表示修改帐号</li>
	// <li>DELETE：表示删除帐号</li>
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// 变更时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type IgnoreImpactedHostsRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *IgnoreImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IgnoreImpactedHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IgnoreImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IgnoreImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IgnoreImpactedHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImpactedHost struct {

	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 最后检测时间。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 漏洞状态。
	// <li>UN_OPERATED ：待处理</li>
	// <li>SCANING : 扫描中</li>
	// <li>FIXED : 已修复</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 漏洞描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 漏洞种类ID。
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 是否为专业版。
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
}

type InquiryPriceOpenProVersionPrepaidRequest struct {
	*tchttp.BaseRequest

	// 预付费模式(包年包月)参数设置。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 需要开通专业版机器列表数组。
	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines" list`
}

func (r *InquiryPriceOpenProVersionPrepaidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceOpenProVersionPrepaidRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceOpenProVersionPrepaidResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 预支费用的原价，单位：元。
		OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 预支费用的折扣价，单位：元。
		DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceOpenProVersionPrepaidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceOpenProVersionPrepaidResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LoginWhiteLists struct {

	// 记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 白名单地域
	Places []*Place `json:"Places,omitempty" name:"Places" list`

	// 白名单用户（多个用户逗号隔开）
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 白名单IP（多个IP逗号隔开）
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 是否为全局规则
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 创建白名单时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改白名单时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 机器IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
}

type LoginWhiteListsRule struct {

	// 加白地域
	Places []*Place `json:"Places,omitempty" name:"Places" list`

	// 加白源IP，支持网段，多个IP以逗号隔开
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 加白用户名，多个用户名以逗号隔开
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 是否对全局生效
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 白名单生效的机器
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 规则ID，用于更新规则
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type Machine struct {

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 主机系统。
	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`

	// 主机状态。
	// <li>OFFLINE: 离线  </li>
	// <li>ONLINE: 在线</li>
	// <li>MACHINE_STOPPED: 已关机</li>
	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`

	// 云镜客户端唯一Uuid，若客户端长时间不在线将返回空字符。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// CVM或BM机器唯一Uuid。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 漏洞数。
	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 是否是专业版。
	// <li>true： 是</li>
	// <li>false：否</li>
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// 主机外网IP。
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 木马数。
	MalwareNum *int64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

	// 标签信息
	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag" list`
}

type MachineTag struct {

	// 关联标签ID
	Rid *int64 `json:"Rid,omitempty" name:"Rid"`

	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type MaliciousRequest struct {
	*tchttp.BaseRequest

	// 记录ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 恶意请求域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 恶意请求数。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 记录状态。
	// <li>UN_OPERATED：待处理</li>
	// <li>TRUSTED：已信任</li>
	// <li>UN_TRUSTED：已取消信任</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 恶意请求域名描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参考地址。
	Reference *string `json:"Reference,omitempty" name:"Reference"`

	// 发现时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 记录合并时间。
	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`

	// 进程MD5
	// 值。
	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`

	// 执行命令行。
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 进程PID。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

func (r *MaliciousRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MaliciousRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Malware struct {

	// 事件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 当前木马状态。
	// <li>UN_OPERATED：未处理</li><li>SEGREGATED：已隔离</li><li>TRUSTED：已信任</li>
	// <li>SEPARATING：隔离中</li><li>RECOVERING：恢复中</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 木马所在的路径。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 木马描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 木马文件创建时间。
	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`

	// 木马文件修改时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type MisAlarmNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 异地登录事件Id数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *MisAlarmNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MisAlarmNonlocalLoginPlacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MisAlarmNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MisAlarmNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MisAlarmNonlocalLoginPlacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmAttributeRequest struct {
	*tchttp.BaseRequest

	// 告警项目。
	// <li>Offline：防护软件离线</li>
	// <li>Malware：发现木马文件</li>
	// <li>NonlocalLogin：发现异地登录行为</li>
	// <li>CrackSuccess：被暴力破解成功</li>
	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`

	// 告警项目属性。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyAlarmAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoOpenProVersionConfigRequest struct {
	*tchttp.BaseRequest

	// 设置自动开通状态。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAutoOpenProVersionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoOpenProVersionConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoOpenProVersionConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoOpenProVersionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoOpenProVersionConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBanRecordRequest struct {
	*tchttp.BaseRequest

	// 白名单信息
	Record *WhiteListRecord `json:"Record,omitempty" name:"Record"`
}

func (r *ModifyBanRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBanRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBanRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBanRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBanRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBanStatusRequest struct {
	*tchttp.BaseRequest

	// 阻断状态 0:关闭 1:开启
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyBanStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBanStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBanStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBanStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBanStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBanWhiteListRequest struct {
	*tchttp.BaseRequest

	// 修改白名单规则项
	Rules *BanWhiteList `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyBanWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBanWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBanWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBanWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBanWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockStatusRequest struct {
	*tchttp.BaseRequest

	// 0：关闭，1：开启
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyBlockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBlockStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBlockStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单规则
	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLoginWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLoginWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProVersionRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 主机唯一ID，对应CVM的uuid、BM的instanceId。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ModifyProVersionRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProVersionRenewFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProVersionRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProVersionRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProVersionRenewFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NonLocalLoginPlace struct {

	// 事件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 登录状态
	// <li>NON_LOCAL_LOGIN：异地登录</li>
	// <li>NORMAL_LOGIN：正常登录</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 城市ID。
	City *uint64 `json:"City,omitempty" name:"City"`

	// 国家ID。
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 省份ID。
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// 登录IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 机器名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 登录时间。
	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`

	// 云镜客户端唯一标识Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type OpenPort struct {

	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 开放端口号。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 端口对应进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 端口对应进程Pid。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 记录创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 记录更新时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type OpenPortStatistics struct {

	// 端口号
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 主机数量
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type OpenProVerWithQuuidsRequest struct {
	*tchttp.BaseRequest

	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM/ECM的Uuid
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids" list`
}

func (r *OpenProVerWithQuuidsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenProVerWithQuuidsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenProVerWithQuuidsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVerWithQuuidsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenProVerWithQuuidsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionPrepaidRequest struct {
	*tchttp.BaseRequest

	// 购买相关参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 需要开通专业版主机信息数组。
	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines" list`
}

func (r *OpenProVersionPrepaidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenProVersionPrepaidRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionPrepaidResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID列表。
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionPrepaidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenProVersionPrepaidResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionQuickRequest struct {
	*tchttp.BaseRequest

	// 活动ID
	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *OpenProVersionQuickRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenProVersionQuickRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionQuickResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionQuickResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenProVersionQuickResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。
	// 如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM的Uuid
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids" list`

	// 活动ID。
	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *OpenProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenProVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenProVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Place struct {

	// 城市 ID。
	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`

	// 省份 ID。
	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// 国家ID，暂只支持国内：1。
	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`
}

type PrivilegeEscalationProcess struct {

	// 数据ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// 执行命令
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户组
	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`

	// 进程文件权限
	ProcFilePrivilege *string `json:"ProcFilePrivilege,omitempty" name:"ProcFilePrivilege"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`

	// 父进程用户名
	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`

	// 进程树
	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`

	// 处理状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
}

type PrivilegeRule struct {

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否S权限
	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态(0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
}

type ProVersionMachine struct {

	// 主机类型。
	// <li>CVM: 虚拟主机</li>
	// <li>BM: 黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 主机所在地域。
	// 如：ap-guangzhou、ap-beijing
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机唯一标识Uuid。
	// 黑石的InstanceId，CVM的Uuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type Process struct {

	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 进程Pid。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 进程Ppid。
	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`

	// 进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 所属平台。
	// <li>WIN32：windows32位</li>
	// <li>WIN64：windows64位</li>
	// <li>LINUX32：Linux32位</li>
	// <li>LINUX64：Linux64位</li>
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 进程路径。
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ProcessStatistics struct {

	// 进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type RecoverMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马Id数组,单次最大删除不能超过200条
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *RecoverMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoverMalwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecoverMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 恢复成功id数组
		SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds" list`

		// 恢复失败id数组
		FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoverMalwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域标志，如 ap-guangzhou，ap-shanghai，ap-beijing
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域中文名，如华南地区（广州），华东地区（上海金融），华北地区（北京）
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域代码，如 gz，sh，bj
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
}

type RenewProVersionRequest struct {
	*tchttp.BaseRequest

	// 购买相关参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 主机唯一ID，对应CVM的uuid、BM的InstanceId。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *RenewProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewProVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewProVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewProVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RescanImpactedHostRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *RescanImpactedHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RescanImpactedHostRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RescanImpactedHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RescanImpactedHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RescanImpactedHostResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReverseShell struct {

	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜UUID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// 命令详情
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 执行用户
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 执行用户组
	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`

	// 父进程用户
	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`

	// 处理状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 产生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 进程树
	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`
}

type ReverseShellRule struct {

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 目标IP
	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`

	// 目标端口
	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态 (0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
}

type SecurityDynamic struct {

	// 云镜客户端UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 安全事件发生事件。
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// 安全事件类型。
	// <li>MALWARE：木马事件</li>
	// <li>NON_LOCAL_LOGIN：异地登录</li>
	// <li>BRUTEATTACK_SUCCESS：密码破解成功</li>
	// <li>VUL：漏洞</li>
	// <li>BASELINE：安全基线</li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 安全事件消息。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SecurityTrend struct {

	// 事件时间。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 事件数量。
	EventNum *uint64 `json:"EventNum,omitempty" name:"EventNum"`
}

type SeparateMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马事件ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *SeparateMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SeparateMalwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SeparateMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 隔离成功的id数组。
		SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds" list`

		// 隔离失败的id数组。
		FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeparateMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SeparateMalwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetAutoQuaraStatusRequest struct {
	*tchttp.BaseRequest

	// 是否开启自动隔离
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetAutoQuaraStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoQuaraStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetAutoQuaraStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoQuaraStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoQuaraStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetBashEventsStatusRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`

	// 新状态(0-待处理 1-高危 2-正常)
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetBashEventsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetBashEventsStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetBashEventsStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBashEventsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetBashEventsStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchBashRulesRequest struct {
	*tchttp.BaseRequest

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否禁用
	Disabled *uint64 `json:"Disabled,omitempty" name:"Disabled"`
}

func (r *SwitchBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchBashRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchBashRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务器数
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type TagMachine struct {

	// ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机名称
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// 主机区域
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机区域类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type TrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest

	// 恶意请求记录ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *TrustMaliciousRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TrustMaliciousRequestRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrustMaliciousRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TrustMaliciousRequestResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TrustMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *TrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TrustMalwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TrustMalwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UntrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest

	// 受信任记录ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UntrustMaliciousRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UntrustMaliciousRequestRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UntrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UntrustMaliciousRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UntrustMaliciousRequestResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UntrustMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马ID数组，单次最大处理不能超过200条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *UntrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UntrustMalwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UntrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UntrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UntrustMalwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest

	// 是否开启木马自动隔离(0:关闭 1:开启)
	MalwareAutoIsolate *int64 `json:"MalwareAutoIsolate,omitempty" name:"MalwareAutoIsolate"`

	// 是否开启木马定时扫描(0:关闭 1:开启)
	MalwareEnableDiskScan *int64 `json:"MalwareEnableDiskScan,omitempty" name:"MalwareEnableDiskScan"`

	// 设置木马定时扫描间隔(天数)
	MalwareDiskScanInterDay *int64 `json:"MalwareDiskScanInterDay,omitempty" name:"MalwareDiskScanInterDay"`

	// 设置木马定时扫描时间(0-23整点时间)
	MalwareDiskScanTime *int64 `json:"MalwareDiskScanTime,omitempty" name:"MalwareDiskScanTime"`

	// 设置木马扫描全盘扫描忽略路径(linux), 换行符分割多个
	MalwareIgnoreWhiteListLinux *string `json:"MalwareIgnoreWhiteListLinux,omitempty" name:"MalwareIgnoreWhiteListLinux"`

	// 设置木马扫描全盘扫描忽略路径(windows), 换行符分割多个
	MalwareIgnoreWhiteListWindows *string `json:"MalwareIgnoreWhiteListWindows,omitempty" name:"MalwareIgnoreWhiteListWindows"`

	// 是否开启漏洞定时扫描(0:关闭 1:开启)
	VulEnableVulnerDetect *int64 `json:"VulEnableVulnerDetect,omitempty" name:"VulEnableVulnerDetect"`

	// 设置漏洞定时扫描间隔(天数)
	VulVulnerDetectDay *int64 `json:"VulVulnerDetectDay,omitempty" name:"VulVulnerDetectDay"`

	// 设置漏洞定时扫描时间(0-23整点时间)
	VulVulnerDetectTime *int64 `json:"VulVulnerDetectTime,omitempty" name:"VulVulnerDetectTime"`

	// 设置木马扫描实时监控忽略路径(linux), 换行符分割多个
	MalwareMonitorIgnoreWhiteListLinux *string `json:"MalwareMonitorIgnoreWhiteListLinux,omitempty" name:"MalwareMonitorIgnoreWhiteListLinux"`

	// 设置木马扫描实时监控忽略路径(windows), 换行符分割多个
	MalwareMonitorIgnoreWhiteListWindows *string `json:"MalwareMonitorIgnoreWhiteListWindows,omitempty" name:"MalwareMonitorIgnoreWhiteListWindows"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UsualPlace struct {

	// ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 国家 ID。
	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`

	// 省份 ID。
	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// 城市 ID。
	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
}

type Vul struct {

	// 漏洞种类ID
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 漏洞名称
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞危害等级:
	// HIGH：高危
	// MIDDLE：中危
	// LOW：低危
	// NOTICE：提示
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// 最后扫描时间
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 受影响机器数量
	ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitempty" name:"ImpactedHostNum"`

	// 漏洞状态
	// * UN_OPERATED : 待处理
	// * FIXED : 已修复
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type VulEventList struct {

	// 漏洞id
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 漏洞事件id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 最后检测时间
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 事件状态:0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 漏洞披露时间
	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`

	// 漏洞名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 漏洞等级 1:低 2:中 3:高 4:提示
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 漏洞描述
	Descript *string `json:"Descript,omitempty" name:"Descript"`

	// 主机quuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// HostIp
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 是否为专业版防护
	IfProfession *bool `json:"IfProfession,omitempty" name:"IfProfession"`

	// 主机别名
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type VulInfoList struct {

	// 漏洞包含的事件id串，多个用“,”分割
	Ids *string `json:"Ids,omitempty" name:"Ids"`

	// 漏洞名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 0: 待处理 1:忽略  3:已修复  5:检测中， 控制台仅处理0,1,3,5四种状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 漏洞id
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 漏洞披露事件
	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`

	// 最后检测时间
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 影响主机数
	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`

	// 漏洞等级 1:低 2:中 3:高 4:提示
	Level *uint64 `json:"Level,omitempty" name:"Level"`
}

type WeeklyReport struct {

	// 周报开始时间。
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 周报结束时间。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type WeeklyReportBruteAttack struct {

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 被破解用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 源IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 尝试次数。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 攻击时间。
	AttackTime *string `json:"AttackTime,omitempty" name:"AttackTime"`
}

type WeeklyReportMalware struct {

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 木马文件路径。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 木马文件MD5值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 木马发现时间。
	FindTime *string `json:"FindTime,omitempty" name:"FindTime"`

	// 当前木马状态。
	// <li>UN_OPERATED：未处理</li>
	// <li>SEGREGATED：已隔离</li>
	// <li>TRUSTED：已信任</li>
	// <li>SEPARATING：隔离中</li>
	// <li>RECOVERING：恢复中</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type WeeklyReportNonlocalLoginPlace struct {

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 源IP。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 国家ID。
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 省份ID。
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// 城市ID。
	City *uint64 `json:"City,omitempty" name:"City"`

	// 登录时间。
	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
}

type WeeklyReportVul struct {

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 漏洞名称。
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞类型。
	// <li> WEB : Web漏洞</li>
	// <li> SYSTEM :系统组件漏洞</li>
	// <li> BASELINE : 安全基线</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 漏洞描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 漏洞状态。
	// <li> UN_OPERATED : 待处理</li>
	// <li> SCANING : 扫描中</li>
	// <li> FIXED : 已修复</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`

	// 最后扫描时间。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
}

type WhiteListRecord struct {

	// 源IP
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 白名单状态 0：关闭 1：启用
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// 白名单ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}
