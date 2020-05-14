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

package v20180411

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActiveDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID列表，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ActiveDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActiveDedicatedDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ActiveDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id，透传入参
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActiveDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActiveDedicatedDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ActiveHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dcdbt-87zif6ha
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ActiveHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActiveHourDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ActiveHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，形如dcdbt-87zif6ha
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActiveHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActiveHourDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddShardConfig struct {

	// 新增分片的数量
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
}

type AuditLog struct {

	// 实例类型
	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 应用id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 操作的数据库
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 操作用户
	User *string `json:"User,omitempty" name:"User"`

	// sql语句
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// 匹配规则
	MatchRule *string `json:"MatchRule,omitempty" name:"MatchRule"`

	// 客户端ip 端口
	ClientIpPort *string `json:"ClientIpPort,omitempty" name:"ClientIpPort"`

	// sql接收时间
	ReceiveTime *int64 `json:"ReceiveTime,omitempty" name:"ReceiveTime"`

	// sql接收时间字符串
	ReceiveTimeStr *string `json:"ReceiveTimeStr,omitempty" name:"ReceiveTimeStr"`

	// 花费时间
	CostTime *uint64 `json:"CostTime,omitempty" name:"CostTime"`

	// 影响行数
	RowsEffected *uint64 `json:"RowsEffected,omitempty" name:"RowsEffected"`

	// 执行结果
	ResultCode *int64 `json:"ResultCode,omitempty" name:"ResultCode"`

	// sql类型
	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`

	// 操作表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type AuditRule struct {

	// 规则Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 应用Id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// db类型
	DbType *int64 `json:"DbType,omitempty" name:"DbType"`

	// 审计规则
	Filters *string `json:"Filters,omitempty" name:"Filters"`

	// 规则状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type AuditRuleDetail struct {

	// 规则Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 应用Id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// db类型
	DbType *int64 `json:"DbType,omitempty" name:"DbType"`

	// 规则
	Filters *string `json:"Filters,omitempty" name:"Filters"`

	// 规则状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 规则列表
	Rules []*RuleItem `json:"Rules,omitempty" name:"Rules" list`
}

type AuditStrategy struct {

	// 策略Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 应用Id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 响应策略，0放行，1不放行
	AlterType *int64 `json:"AlterType,omitempty" name:"AlterType"`

	// 脱敏记录，0不脱敏，1脱敏
	ActionType *int64 `json:"ActionType,omitempty" name:"ActionType"`

	// 策略状态，0不启用，1启用
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 优先顺序
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AuthenticateCAMRequest struct {
	*tchttp.BaseRequest

	// 待鉴权的接口名，如 CreateDCDBInstance
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// 待鉴权资源信息，cam 六段式资源格式或 *
	Resources []*string `json:"Resources,omitempty" name:"Resources" list`
}

func (r *AuthenticateCAMRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuthenticateCAMRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuthenticateCAMResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前请求账号是否为根账号
		IsRoot *int64 `json:"IsRoot,omitempty" name:"IsRoot"`

		// 当前请求账号是否有全部权限
		IsAllPass *int64 `json:"IsAllPass,omitempty" name:"IsAllPass"`

		// 主账号的appId
		OwnerAppId *int64 `json:"OwnerAppId,omitempty" name:"OwnerAppId"`

		// 权限详情
		Permissions []*Permission `json:"Permissions,omitempty" name:"Permissions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuthenticateCAMResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuthenticateCAMResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckIpStatusRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：dcdbt-ow728lmc
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 目标子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 目标虚拟IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *CheckIpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckIpStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckIpStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 透传入参
		Vip *string `json:"Vip,omitempty" name:"Vip"`

		// Vip可用状态：0 - 可用，300 - 此ip已被子网内其他用户占用， 301 - 此ip不在子网内， 302 - ip格式错误， 303 - 该IP为实例当前IP。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckIpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckIpStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloneAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 源用户账户名
	SrcUser *string `json:"SrcUser,omitempty" name:"SrcUser"`

	// 源用户HOST
	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`

	// 目的用户账户名
	DstUser *string `json:"DstUser,omitempty" name:"DstUser"`

	// 目的用户HOST
	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`

	// 目的用户账户描述
	DstDesc *string `json:"DstDesc,omitempty" name:"DstDesc"`
}

func (r *CloneAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloneAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// 待关闭外网访问的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务Id，可通过 DescribeFlow 查询任务状态。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConstraintRange struct {

	// 约束类型为section时的最小值
	Min *string `json:"Min,omitempty" name:"Min"`

	// 约束类型为section时的最大值
	Max *string `json:"Max,omitempty" name:"Max"`
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 源用户名
	SrcUserName *string `json:"SrcUserName,omitempty" name:"SrcUserName"`

	// 源用户允许的访问 host
	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`

	// 目的用户名
	DstUserName *string `json:"DstUserName,omitempty" name:"DstUserName"`

	// 目的用户允许的访问 host
	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`

	// 源账号的 ReadOnly 属性
	SrcReadOnly *string `json:"SrcReadOnly,omitempty" name:"SrcReadOnly"`

	// 目的账号的 ReadOnly 属性
	DstReadOnly *string `json:"DstReadOnly,omitempty" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// AccountName
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 可以登录的主机，与mysql 账号的 host 格式一致，可以支持通配符，例如 %，10.%，10.20.%。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 账号密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否创建为只读账号，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败，3：只从备机读取。
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 账号备注，可以包含中文、英文字符、常见符号和数字，长度为0~256字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。
	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id，透传入参。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 用户名，透传入参。
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 允许访问的 host，透传入参。
		Host *string `json:"Host,omitempty" name:"Host"`

		// 透传入参。
		ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 分片节点可用区分布，最多可填两个可用区。当分片规格为一主两从时，其中两个节点在第一个可用区。
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`

	// 单个分片节点个数，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitempty" name:"ShardNodeCount"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 欲购买实例的数量，目前只支持购买1个实例
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 虚拟私有网络 ID，不传或传空表示创建为基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 虚拟私有网络子网 ID，VpcId不为空时必填
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 数据库引擎版本，当前可选：10.0.10，10.1.9，5.7.17。
	// 10.0.10 - Mariadb 10.0.10；
	// 10.1.9 - Mariadb 10.1.9；
	// 5.7.17 - Percona 5.7.17。
	// 如果不填的话，默认为10.1.9，表示Mariadb 10.1.9。
	DbVersionId *string `json:"DbVersionId,omitempty" name:"DbVersionId"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *CreateDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// VPC网络ID，形如vpc-qquscup9。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC网络子网ID，形如subnet-mfhkoln6。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 分片内存大小，单位为GB。
	ShardMemory *uint64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片磁盘大小，单位为GB。
	ShardStorage *uint64 `json:"ShardStorage,omitempty" name:"ShardStorage"`

	// 每个分片的节点个数，值为2或者3。2-一主一从；3-一主两从
	ShardNodeCount *uint64 `json:"ShardNodeCount,omitempty" name:"ShardNodeCount"`

	// 分片的个数。取值范围为[2, 64]，即最小分片个数为2，最大分片个数为64。
	ShardCount *uint64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 实例可用区信息，只能填一个可用区或者两个可用区。如果只填一个，那么主库和从库可用区一致；如果填两个，那么前一个表示主库可用区，后一个表示从库可用区。
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 数据库版本。取值为10.0.10，10.1.9，5.7.17或者5.6.39。不填默认为5.7.17。
	DbVersionId *string `json:"DbVersionId,omitempty" name:"DbVersionId"`

	// 购买数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 分片 Cpu，单位： 核
	ShardCpu *uint64 `json:"ShardCpu,omitempty" name:"ShardCpu"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 资源池 Id，TCE 专用
	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`
}

func (r *CreateHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHourDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的DCDB实例ID列表，形如dcdbt-87zif6ha
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 创建实例异步任务ID，以FlowId作为参数调用DescribeFlow API接口查询创建实例任务执行状态。
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHourDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTmpDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 回档实例的ID，如:dcdbt-fp0n6tul
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 回档时间点
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// 用户类型
	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
}

func (r *CreateTmpDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTmpDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTmpDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTmpDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTmpDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DBAccount struct {

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户可以从哪台主机登录（对应 MySQL 用户的 host 字段，UserName + Host 唯一标识一个用户，IP形式，IP段以%结尾；支持填入%；为空默认等于%）
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户备注信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 只读标记，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败。
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。
	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

type DBEngine struct {

	// 引擎类型，MariaDB 或 Percona
	Type *string `json:"Type,omitempty" name:"Type"`

	// 引擎版本号，创建实例时传递该值
	Version *string `json:"Version,omitempty" name:"Version"`

	// 名称，用于展示
	Name *string `json:"Name,omitempty" name:"Name"`

	// 引擎版本描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

type DBParamValue struct {

	// 参数名称
	Param *string `json:"Param,omitempty" name:"Param"`

	// 参数值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DCDBInstanceInfo struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// APPID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC数字ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet数字ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 状态中文描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 内网IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 内网端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 自动续费标志
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 分片个数
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 隔离时间
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`

	// UIN
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 分片详情
	ShardDetail []*ShardInfo `json:"ShardDetail,omitempty" name:"ShardDetail" list`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 临时实例标记，0 为非临时实例
	IsTmp *int64 `json:"IsTmp,omitempty" name:"IsTmp"`

	// 独享集群Id，为空表示非独享集群实例
	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`

	// 字符串型的私有网络Id
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 字符串型的私有网络子网Id
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// 数字实例Id（过时字段，请勿依赖该值）
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 外网访问的域名，公网可解析
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// 外网 IP 地址，公网可访问
	WanVip *string `json:"WanVip,omitempty" name:"WanVip"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// 产品类型 Id（过时字段，请勿依赖该值）
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 实例最后更新时间，格式为 2006-01-02 15:04:05
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 数据库引擎
	DbEngine *string `json:"DbEngine,omitempty" name:"DbEngine"`

	// 数据库引擎版本
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 付费模式
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`

	// 实例处于异步任务状态时，表示异步任务流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locker *int64 `json:"Locker,omitempty" name:"Locker"`

	// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中
	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

	// 该实例是否支持审计。1-支持；0-不支持
	IsAuditSupported *uint64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`
}

type DCDBShardInfo struct {

	// 所属实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分片SQL透传Id，用于将sql透传到指定分片执行
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`

	// 全局唯一的分片Id
	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`

	// 状态：0 创建中，1 流程处理中， 2 运行中，3 分片未初始化
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 状态中文描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 字符串格式的私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 字符串格式的私有网络子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 存储使用率，单位为 %
	StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`

	// 内存使用率，单位为 %
	MemoryUsage *float64 `json:"MemoryUsage,omitempty" name:"MemoryUsage"`

	// 数字分片Id（过时字段，请勿依赖该值）
	ShardId *int64 `json:"ShardId,omitempty" name:"ShardId"`

	// 产品ProductID
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Proxy版本
	ProxyVersion *string `json:"ProxyVersion,omitempty" name:"ProxyVersion"`
}

type Database struct {

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type DatabaseFunction struct {

	// 函数名称
	Func *string `json:"Func,omitempty" name:"Func"`
}

type DatabaseProcedure struct {

	// 存储过程名称
	Proc *string `json:"Proc,omitempty" name:"Proc"`
}

type DatabaseTable struct {

	// 表名
	Table *string `json:"Table,omitempty" name:"Table"`
}

type DatabaseView struct {

	// 视图名称
	View *string `json:"View,omitempty" name:"View"`
}

type Deal struct {

	// 订单号
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 所属账号
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 商品数量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 关联的流程 Id，可用于查询流程执行状态
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 只有创建实例的订单会填充该字段，表示该订单创建的实例的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 付费模式，0后付费/1预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTmpInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteTmpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTmpInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTmpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTmpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTmpInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示查询该数据库权限（即db.\*），此时忽略 Object 参数
	Type *string `json:"Type,omitempty" name:"Type"`

	// 具体的 Type 的名称，比如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitempty" name:"Object"`

	// 当 Type=table 时，ColName 为 \* 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限
	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 权限列表。
		Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

		// 数据库账号用户名
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 数据库账号Host
		Host *string `json:"Host,omitempty" name:"Host"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

		// 实例ID，透传入参。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例用户列表。
		Users []*DBAccount `json:"Users,omitempty" name:"Users" list`

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

type DescribeAuditLogsRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 客户端IP和端口列表
	ClientIpPorts []*string `json:"ClientIpPorts,omitempty" name:"ClientIpPorts" list`

	// sql操作者列表
	SqlUsers []*string `json:"SqlUsers,omitempty" name:"SqlUsers" list`

	// sql类型列表
	SqlTypes []*string `json:"SqlTypes,omitempty" name:"SqlTypes" list`

	// 审计策略名称列表
	StrategyNames []*string `json:"StrategyNames,omitempty" name:"StrategyNames" list`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 数据库名称列表
	DbNames []*string `json:"DbNames,omitempty" name:"DbNames" list`

	// 搜索关键字，sql可选
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 搜索关键字的条件值
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords" list`

	// 分页，页大小，默认150
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页，页偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，costtime,rowseffected,name,receivetime可选，默认receivetime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，desc，asc可选
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeAuditLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 审计日志列表
		AuditLogSet []*AuditLog `json:"AuditLogSet,omitempty" name:"AuditLogSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditRuleDetailRequest struct {
	*tchttp.BaseRequest

	// 规则Id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeAuditRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditRuleDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 审计规则
		AuditRuleDetail *AuditRuleDetail `json:"AuditRuleDetail,omitempty" name:"AuditRuleDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditRuleDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditRulesRequest struct {
	*tchttp.BaseRequest

	// 规则名称列表
	RuleNames []*string `json:"RuleNames,omitempty" name:"RuleNames" list`

	// 分页，页大小，默认150
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页，页偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，createTime,updateTime,name可选，默认createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，desc，asc可选
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeAuditRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 审计规则列表
		RuleSet []*AuditRule `json:"RuleSet,omitempty" name:"RuleSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditStrategiesRequest struct {
	*tchttp.BaseRequest

	// 实例Id，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略名称列表
	StrategyNames []*string `json:"StrategyNames,omitempty" name:"StrategyNames" list`

	// 规则名称列表
	RuleNames []*string `json:"RuleNames,omitempty" name:"RuleNames" list`

	// 响应策略，0放行，1不放行
	AlterTypes []*int64 `json:"AlterTypes,omitempty" name:"AlterTypes" list`

	// 脱敏记录，0不脱敏，1脱敏
	ActionTypes []*int64 `json:"ActionTypes,omitempty" name:"ActionTypes" list`

	// 策略状态，0不启用，1启用
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// 优先顺序
	Priorities []*int64 `json:"Priorities,omitempty" name:"Priorities" list`
}

func (r *DescribeAuditStrategiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditStrategiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditStrategiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 会话Id
		SessionId *int64 `json:"SessionId,omitempty" name:"SessionId"`

		// 审计策略列表
		StrategySet []*AuditStrategy `json:"StrategySet,omitempty" name:"StrategySet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditStrategiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditStrategiesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableExclusiveGroupsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableExclusiveGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableExclusiveGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableExclusiveGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 独享资源池数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 独享资源池信息
		Items []*FenceInfoItem `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableExclusiveGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableExclusiveGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchDCDBRenewalPriceRequest struct {
	*tchttp.BaseRequest

	// 待续费的实例ID数组。形如：["dcdbt-ow728lmc"]，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 续费时长，单位：月。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeBatchDCDBRenewalPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBatchDCDBRenewalPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchDCDBRenewalPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBatchDCDBRenewalPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBatchDCDBRenewalPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDetailMetricsRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：tdsql-ow728lmc
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 起始日期，格式yyyy-mm-dd
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束日期，格式yyyy-mm-dd
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 需要请求的指标名，当前支持的指标有：
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`

	// 监控统计周期。默认为取值为300，单位为s，当前可选：60，300
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 实例分片ID
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *DescribeDBDetailMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDetailMetricsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDetailMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 统计周期
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// 主节点数据点数组
		MasterDataPoints []*MetricDataPoint `json:"MasterDataPoints,omitempty" name:"MasterDataPoints" list`

		// 备机1数据点数组
		Slave1DataPoints []*MetricDataPoint `json:"Slave1DataPoints,omitempty" name:"Slave1DataPoints" list`

		// 备机2数据点数组
		Slave2DataPoints []*MetricDataPoint `json:"Slave2DataPoints,omitempty" name:"Slave2DataPoints" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBDetailMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDetailMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEnginesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBEnginesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBEnginesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEnginesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的引擎总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 支持的引擎列表
		Items []*DBEngine `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBEnginesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBEnginesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRsipRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceRsipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceRsipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRsipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 后端RS信息
		Rsips []*Rsip `json:"Rsips,omitempty" name:"Rsips" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceRsipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceRsipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分片 ID，形如：shard-7noic7tv
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID，形如：dcdbt-ow728lmc。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 请求日志类型。1-binlog，2-冷备，3-errlog，4-slowlog。
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// 请求日志总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 日志文件列表
		Files []*LogFileInfo `json:"Files,omitempty" name:"Files" list`

		// 如果是VPC网络的实例，做用本前缀加上URI为下载地址
		VpcPrefix *string `json:"VpcPrefix,omitempty" name:"VpcPrefix"`

		// 如果是普通网络的实例，做用本前缀加上URI为下载地址
		NormalPrefix *string `json:"NormalPrefix,omitempty" name:"NormalPrefix"`

		// 分片 ID，形如：shard-7noic7tv
		ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBMetricsRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：dcdbt-ow728lmc
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 起始日期，格式yyyy-mm-dd
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束日期，格式yyyy-mm-dd
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 需要请求的指标名，当前支持的指标有：
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`

	// 监控统计周期。默认为取值为300，单位为s，当前可选：60，300
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// DCDB分片ID，形如shard-15xl67ih。如果填充该字段，表示拉取的是该分片的监控信息
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *DescribeDBMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBMetricsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控统计周期
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// 监控指标数据点数组
		DataPoints []*MetricDataPoint `json:"DataPoints,omitempty" name:"DataPoints" list`

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID，形如：dcdbt-ow7t8lmc。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 请求DB的当前参数值
		Params []*ParamDesc `json:"Params,omitempty" name:"Params" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：mariadb,cdb,cynosdb,dcdb,redis,mongodb 等。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组规则
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogAnalysisRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如 dcdbt-hw0qj6m1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，形如 2006-01-02 15:04:05
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 要查询的慢查询语句对应的数据库名称
	Db *string `json:"Db,omitempty" name:"Db"`

	// 要查询的慢查询语句对应的用户名称
	User *string `json:"User,omitempty" name:"User"`

	// 要查询的慢查询语句的校验和，可以通过查询慢查询日志列表获得
	CheckSum *string `json:"CheckSum,omitempty" name:"CheckSum"`

	// 实例分片ID，形如 shard-nbchf68b
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

	// 结束时间，形如 2006-01-02 15:04:05，不填的话默认是当前时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否查询从机的慢查询。0-主机；1-从机，默认查询主机慢查询
	Slave *int64 `json:"Slave,omitempty" name:"Slave"`
}

func (r *DescribeDBSlowLogAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowLogAnalysisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 慢查询SQL出现的开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 慢查询SQL出现的结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 返回的慢查询详情数据。每个点代表一个时间段内慢查询SQL出现的记录次数，一个时间段的长度由请求的开始时间和结束时间的差值决定，小于1天是5分钟一段，大于1天小于7天是30分钟一段，大于7天是2个小时一段。
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// Data 数组中的数据点间隔，单位为秒
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowLogAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowLogAnalysisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-hw0qj6m1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 从结果的第几条数据开始返回
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的结果条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询的起始时间，形如2016-07-23 14:55:20
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 实例的分片ID，形如shard-53ima8ln
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

	// 查询的结束时间，形如2016-08-22 14:55:20。如果不填，那么查询结束时间就是当前时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 要查询的具体数据库名称
	Db *string `json:"Db,omitempty" name:"Db"`

	// 排序指标，取值为query_time_sum或者query_count。不填默认按照query_time_sum排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型，desc（降序）或者asc（升序）。不填默认desc排序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 是否查询从机的慢查询，0-主机; 1-从机。不填默认查询主机慢查询
	Slave *int64 `json:"Slave,omitempty" name:"Slave"`
}

func (r *DescribeDBSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 所有语句锁时间总和
		LockTimeSum *float64 `json:"LockTimeSum,omitempty" name:"LockTimeSum"`

		// 所有语句查询总次数
		QueryCount *int64 `json:"QueryCount,omitempty" name:"QueryCount"`

		// 总记录数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 所有语句查询时间总和
		QueryTimeSum *float64 `json:"QueryTimeSum,omitempty" name:"QueryTimeSum"`

		// 慢查询日志数据
		Data []*SlowLogData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSyncModeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 同步模式：0 异步，1 强同步， 2 强同步可退化
		SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`

		// 是否有修改流程在执行中：1 是， 0 否。
		IsModifying *int64 `json:"IsModifying,omitempty" name:"IsModifying"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSyncModeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBTmpInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBTmpInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBTmpInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBTmpInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 临时实例
		TmpInstances []*TmpInstance `json:"TmpInstances,omitempty" name:"TmpInstances" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBTmpInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBTmpInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBBinlogTimeRequest struct {
	*tchttp.BaseRequest

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户类型
	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
}

func (r *DescribeDCDBBinlogTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBBinlogTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBBinlogTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBBinlogTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBBinlogTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dcdbt-7oaxtcb7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDCDBInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBInstanceDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，形如dcdbt-7oaxtcb7
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例名称
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// 实例状态。0-实例创建中；1-异步任务处理中；2-运行中；3-实例未初始化；-1-实例已隔离
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 实例目前运行状态描述
		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

		// 实例内网IP地址
		Vip *string `json:"Vip,omitempty" name:"Vip"`

		// 实例内网端口
		Vport *int64 `json:"Vport,omitempty" name:"Vport"`

		// 实例节点数。值为2时表示一主一从，值为3时表示一主二从
		NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

		// 实例所在地域名称，形如ap-guangzhou
		Region *string `json:"Region,omitempty" name:"Region"`

		// 实例私有网络ID，形如vpc-r9jr0de3
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 实例私有网络子网ID，形如subnet-6rqs61o2
		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

		// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中
		WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

		// 外网访问的域名，公网可解析
		WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

		// 外网IP地址，公网可访问
		WanVip *string `json:"WanVip,omitempty" name:"WanVip"`

		// 外网访问端口
		WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

		// 实例所属项目ID
		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// 实例自动续费标志。0-正常续费；1-自动续费；2-到期不续费
		AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

		// 独享集群ID
		ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`

		// 付费模式。prepaid-预付费；postpaid-按量计费
		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

		// 实例创建时间，格式为 2006-01-02 15:04:05
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 实例到期时间，格式为 2006-01-02 15:04:05
		PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

		// 数据库版本信息
		DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

		// 实例是否支持审计。0-不支持；1-支持
		IsAuditSupported *int64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`

		// 实例是否支持数据加密。0-不支持；1-支持
		IsEncryptSupported *int64 `json:"IsEncryptSupported,omitempty" name:"IsEncryptSupported"`

		// 实例母机机器型号
		Machine *string `json:"Machine,omitempty" name:"Machine"`

		// 实例内存大小，单位 GB，各个分片的内存大小的和
		Memory *int64 `json:"Memory,omitempty" name:"Memory"`

		// 实例磁盘存储大小，单位 GB，各个分片的磁盘大小的和
		Storage *int64 `json:"Storage,omitempty" name:"Storage"`

		// 实例存储空间使用率，计算方式为：各个分片已经使用的磁盘大小的和/各个分片的磁盘大小的和。
		StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`

		// 日志存储空间大小，单位GB
		LogStorage *int64 `json:"LogStorage,omitempty" name:"LogStorage"`

		// 产品类型ID
		Pid *int64 `json:"Pid,omitempty" name:"Pid"`

		// 主DB可用区
		MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

		// 从DB可用区
		SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones" list`

		// 分片信息
		Shards []*ShardBriefInfo `json:"Shards,omitempty" name:"Shards" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBInstanceDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例 ID 查询。实例 ID 形如：dcdbt-2t4cf98d
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 搜索的字段名，当前支持的值有：instancename、vip、all。传 instancename 表示按实例名进行搜索；传 vip 表示按内网IP进行搜索；传 all 将会按实例ID、实例名和内网IP进行搜索。
	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`

	// 搜索的关键字，支持模糊搜索。多个关键字使用换行符（'\n'）分割。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 按项目 ID 查询
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 是否根据 VPC 网络来搜索
	IsFilterVpc *bool `json:"IsFilterVpc,omitempty" name:"IsFilterVpc"`

	// 私有网络 ID， IsFilterVpc 为 1 时有效
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络的子网 ID， IsFilterVpc 为 1 时有效
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 排序字段， projectId， createtime， instancename 三者之一
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 10，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 1非独享集群，2独享集群， 0全部
	ExclusterType *int64 `json:"ExclusterType,omitempty" name:"ExclusterType"`

	// 标识是否使用ExclusterType字段, false不使用，true使用
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitempty" name:"IsFilterExcluster"`

	// 独享集群ID
	ExclusterIds []*string `json:"ExclusterIds,omitempty" name:"ExclusterIds" list`
}

func (r *DescribeDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息列表
		Instances []*DCDBInstanceInfo `json:"Instances,omitempty" name:"Instances" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBPriceRequest struct {
	*tchttp.BaseRequest

	// 欲新购实例的可用区ID。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 欲购买实例的数量，目前只支持购买1个实例
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 单个分片节点个数大小，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitempty" name:"ShardNodeCount"`

	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`
}

func (r *DescribeDCDBPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBRenewalPriceRequest struct {
	*tchttp.BaseRequest

	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费时长，单位：月。不传则默认为1个月。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeDCDBRenewalPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBRenewalPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBRenewalPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBRenewalPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBRenewalPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBSaleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDCDBSaleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBSaleInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBSaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可售卖地域信息列表
		RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBSaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBSaleInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBShardsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分片Id列表。
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds" list`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段， 目前仅支持 createtime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeDCDBShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBShardsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBShardsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的分片数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 分片信息列表
		Shards []*DCDBShardInfo `json:"Shards,omitempty" name:"Shards" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBShardsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBUpgradePriceRequest struct {
	*tchttp.BaseRequest

	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitempty" name:"SplitShardConfig"`
}

func (r *DescribeDCDBUpgradePriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBUpgradePriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBUpgradePriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位：分
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际价格，单位：分。受折扣等影响，可能和原价不同。
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBUpgradePriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBUpgradePriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 透传入参。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 数据库名称。
		DbName *string `json:"DbName,omitempty" name:"DbName"`

		// 表列表。
		Tables []*DatabaseTable `json:"Tables,omitempty" name:"Tables" list`

		// 视图列表。
		Views []*DatabaseView `json:"Views,omitempty" name:"Views" list`

		// 存储过程列表。
		Procs []*DatabaseProcedure `json:"Procs,omitempty" name:"Procs" list`

		// 函数列表。
		Funcs []*DatabaseFunction `json:"Funcs,omitempty" name:"Funcs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseTableRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 表名称，通过 DescribeDatabaseObjects 接口获取。
	Table *string `json:"Table,omitempty" name:"Table"`
}

func (r *DescribeDatabaseTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例名称。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 数据库名称。
		DbName *string `json:"DbName,omitempty" name:"DbName"`

		// 表名称。
		Table *string `json:"Table,omitempty" name:"Table"`

		// 列信息。
		Cols []*TableColumn `json:"Cols,omitempty" name:"Cols" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该实例上的数据库列表。
		Databases []*Database `json:"Databases,omitempty" name:"Databases" list`

		// 透传入参。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFenceShardSpecRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFenceShardSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFenceShardSpecRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFenceShardSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 按机型分类的可售卖规格列表
		SpecConfig []*SpecConfig `json:"SpecConfig,omitempty" name:"SpecConfig" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFenceShardSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFenceShardSpecResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest

	// 异步请求接口返回的任务流程号。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程状态，0：成功，1：失败，2：运行中
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID的数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 每次最多拉取多少个实例
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 拉取实例列表时候的偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照项目ID进行过滤
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID，形如 vpc-nobbvxkf
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 子网ID，形如 subnet-dnn20os8
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例总的个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例列表
		Instances []*InstanceResource `json:"Instances,omitempty" name:"Instances" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestCloudDBAReportRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DCDB实例分片ID
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *DescribeLatestCloudDBAReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLatestCloudDBAReportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestCloudDBAReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次检测的有效分数
		Score *int64 `json:"Score,omitempty" name:"Score"`

		// 本次检测报告的名称
		ReportName *string `json:"ReportName,omitempty" name:"ReportName"`

		// 本次检测报告的详情
		ReportDetail *string `json:"ReportDetail,omitempty" name:"ReportDetail"`

		// 本次检测报告的完成时间
		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLatestCloudDBAReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLatestCloudDBAReportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeLogFileRetentionPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogFileRetentionPeriodRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID，形如：tdsql-ow728lmc。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 日志备份天数
		Days *uint64 `json:"Days,omitempty" name:"Days"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogFileRetentionPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogFileRetentionPeriodResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest

	// 待查询的长订单号列表，创建实例、续费实例、扩容实例接口返回。
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的订单数量。
		TotalCount []*int64 `json:"TotalCount,omitempty" name:"TotalCount" list`

		// 订单信息列表。
		Deals []*Deal `json:"Deals,omitempty" name:"Deals" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：mariadb,cdb,cynosdb,dcdb,redis,mongodb
	Product *string `json:"Product,omitempty" name:"Product"`

	// 项目Id
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组规则。
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目列表
		Projects []*Project `json:"Projects,omitempty" name:"Projects" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShardSpecRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeShardSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShardSpecRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShardSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 按机型分类的可售卖规格列表
		SpecConfig []*SpecConfig `json:"SpecConfig,omitempty" name:"SpecConfig" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShardSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShardSpecResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlLogsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// SQL日志偏移。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取数量（0-10000，为0时拉取总数信息）。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSqlLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSqlLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前消息队列中的sql日志条目数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 消息队列中的sql日志起始偏移。
		StartOffset *uint64 `json:"StartOffset,omitempty" name:"StartOffset"`

		// 消息队列中的sql日志结束偏移。
		EndOffset *uint64 `json:"EndOffset,omitempty" name:"EndOffset"`

		// 返回的第一条sql日志的偏移。
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// 返回的sql日志数量。
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// Sql日志列表。
		SqlItems []*SqlLogItem `json:"SqlItems,omitempty" name:"SqlItems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSqlLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSqlLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dcdbt-87zif6ha
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyHourDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，形如dcdbt-87zif6ha
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 删除DCDB任务异步流程ID，以FlowId作为参数调用API 接口DescribeFlow 查询删除任务执行状态。
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyHourDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExpandShardConfig struct {

	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds" list`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
}

type FenceInfoItem struct {

	// 独享资源池ID
	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 全局权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER，SHOW DATABASES 
	// 库权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER 
	// 表/视图权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE VIEW，SHOW VIEW，TRIGGER 
	// 存储过程/函数权限： ALTER ROUTINE，EXECUTE 
	// 字段权限： INSERT，REFERENCES，SELECT，UPDATE
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示设置该数据库权限（即db.\*），此时忽略 Object 参数
	Type *string `json:"Type,omitempty" name:"Type"`

	// 具体的 Type 的名称，比如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitempty" name:"Object"`

	// 当 Type=table 时，ColName 为 \* 表示对表授权，如果为具体字段名，表示对字段授权
	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Inbound struct {

	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type InitDCDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 待初始化的实例Id列表，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步）。
	Params []*DBParamValue `json:"Params,omitempty" name:"Params" list`
}

func (r *InitDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDCDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务Id，可通过 DescribeFlow 查询任务状态。
		FlowIds []*uint64 `json:"FlowIds,omitempty" name:"FlowIds" list`

		// 透传入参。
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDCDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceResource struct {

	// 数据库实例的数字ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 实例ID，形如 dcdbt-ifzc58fh或者tdsqlshard-ifzc58fh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 账户的应用ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 实例的原始SerialId
	OriginSerialId *string `json:"OriginSerialId,omitempty" name:"OriginSerialId"`

	// 实例的当前SerialId
	SerialId *string `json:"SerialId,omitempty" name:"SerialId"`

	// 实例所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例所在可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例所属VPC的数字ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 实例所属子网的数字ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例所属VPC的字符串唯一ID
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 实例所属子网的字符串唯一ID
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// 实例运行状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例Vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例Vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 实例所属Zk的名字
	ZkName *string `json:"ZkName,omitempty" name:"ZkName"`

	// 实例所属集群的名字
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 实例的节点个数
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 实例的数据库版本
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 分片信息
	ShardDetail []*InstanceShardInfo `json:"ShardDetail,omitempty" name:"ShardDetail" list`
}

type InstanceShardInfo struct {

	// 分片ID
	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`

	// 分片的SerialId
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`

	// 分片的运行状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type IsolateDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 Id，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateDedicatedDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolateDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateDedicatedDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolateHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dcdbt-87zif6ha。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateHourDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolateHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，形如dcdbt-87zif6ha。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateHourDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LogFileInfo struct {

	// Log最后修改时间
	Mtime *uint64 `json:"Mtime,omitempty" name:"Mtime"`

	// 文件长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`

	// 下载Log时用到的统一资源标识符
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type MetricDataPoint struct {

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 监控值数组，该数组和Timestamps一一对应
	Values []*float64 `json:"Values,omitempty" name:"Values" list`

	// 时间戳数组，表示那些时间点有数据，缺失的时间戳，没有数据点，可以理解为掉点了
	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps" list`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 新的账号备注，长度 0~256。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 实例 ID列表，形如：dcdbt-ow728lmc。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 自动续费标记。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *ModifyAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoRenewFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoRenewFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tdsql-hdaprz0v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：mariadb,cdb,cynosdb,dcdb,redis,mongodb 等。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// 待修改的实例ID列表。实例 ID 形如：dcdbt-ow728lmc。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 要分配的项目 ID，可以通过 DescribeProjects 查询项目列表接口获取。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 参数列表，每一个元素是Param和Value的组合
	Params []*DBParamValue `json:"Params,omitempty" name:"Params" list`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID，形如：dcdbt-ow728lmc。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 各参数修改结果
		Result []*ParamModifyResult `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 同步模式：0 异步，1 强同步， 2 强同步可退化
	SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`
}

func (r *ModifyDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBSyncModeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务Id，可通过 DescribeFlow 查询任务状态。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBSyncModeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNetworkRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 希望转到的VPC网络的VpcId
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 希望转到的VPC网络的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 如果需要指定VIP，填上该字段
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *ModifyInstanceNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceNetworkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID，根据此FlowId通过DescribeFlow接口查询任务进行状态
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceNetworkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRemarkRequest struct {
	*tchttp.BaseRequest

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyInstanceRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceRemarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 备注
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 是否修改
		IsModify *int64 `json:"IsModify,omitempty" name:"IsModify"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceRemarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVipRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例Vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *ModifyInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 保存的天数,不能超过30
	Days *uint64 `json:"Days,omitempty" name:"Days"`
}

func (r *ModifyLogFileRetentionPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLogFileRetentionPeriodRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例 ID，形如：tdsql-ow728lmc。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogFileRetentionPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLogFileRetentionPeriodResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// 待开放外网访问的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OpenDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务Id，可通过 DescribeFlow 查询任务状态。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ParamConstraint struct {

	// 约束类型,如枚举enum，区间section
	Type *string `json:"Type,omitempty" name:"Type"`

	// 约束类型为enum时的可选值列表
	Enum *string `json:"Enum,omitempty" name:"Enum"`

	// 约束类型为section时的范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *ConstraintRange `json:"Range,omitempty" name:"Range"`

	// 约束类型为string时的可选值列表
	String *string `json:"String,omitempty" name:"String"`
}

type ParamDesc struct {

	// 参数名字
	Param *string `json:"Param,omitempty" name:"Param"`

	// 当前参数值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 设置过的值，参数生效后，该值和value一样。未设置过就不返回该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetValue *string `json:"SetValue,omitempty" name:"SetValue"`

	// 系统默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数限制
	Constraint *ParamConstraint `json:"Constraint,omitempty" name:"Constraint"`
}

type ParamModifyResult struct {

	// 修改参数名字
	Param *string `json:"Param,omitempty" name:"Param"`

	// 参数修改结果。0表示修改成功；-1表示修改失败；-2表示该参数值非法
	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type Permission struct {

	// 资源对象
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 是否有权限，1 - 有权限， 0 - 无权限
	IsPermitted *int64 `json:"IsPermitted,omitempty" name:"IsPermitted"`
}

type Project struct {

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 资源拥有者（主账号）uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 应用Id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 项目名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建者uin
	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 来源平台
	SrcPlat *string `json:"SrcPlat,omitempty" name:"SrcPlat"`

	// 来源AppId
	SrcAppId *int64 `json:"SrcAppId,omitempty" name:"SrcAppId"`

	// 项目状态,0正常，-1关闭。默认项目为3
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否默认项目，1 是，0 不是
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 描述信息
	Info *string `json:"Info,omitempty" name:"Info"`
}

type RegionInfo struct {

	// 地域英文ID
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域数字ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域中文名
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区列表
	ZoneList []*ZonesInfo `json:"ZoneList,omitempty" name:"ZoneList" list`

	// 可选择的主可用区和从可用区
	AvailableChoice []*ShardZoneChooseInfo `json:"AvailableChoice,omitempty" name:"AvailableChoice" list`
}

type RenewDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费时长，单位：月。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *RenewDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Rsip struct {

	// Ip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 端口
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type RuleItem struct {

	// 规格类型
	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`

	// 操作类型
	OperationType *int64 `json:"OperationType,omitempty" name:"OperationType"`

	// 操作值
	OperationValue *string `json:"OperationValue,omitempty" name:"OperationValue"`
}

type SecurityGroup struct {

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称。
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注。
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`

	// 出站规则。
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound" list`

	// 入站规则。
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound" list`
}

type ShardBriefInfo struct {

	// 分片SerialId
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`

	// 分片ID，形如shard-7vg1o339
	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`

	// 分片运行状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 分片运行状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 分片创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 分片内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 分片磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 分片日志磁盘空间大小，单位GB
	LogDisk *int64 `json:"LogDisk,omitempty" name:"LogDisk"`

	// 分片节点个数
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 分片磁盘空间使用率
	StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`

	// 分片Proxy版本信息
	ProxyVersion *string `json:"ProxyVersion,omitempty" name:"ProxyVersion"`

	// 分片主DB可用区
	ShardMasterZone *string `json:"ShardMasterZone,omitempty" name:"ShardMasterZone"`

	// 分片从DB可用区
	ShardSlaveZones []*string `json:"ShardSlaveZones,omitempty" name:"ShardSlaveZones" list`
}

type ShardInfo struct {

	// 分片ID
	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`

	// 分片Set ID
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`

	// 状态：0 创建中，1 流程处理中， 2 运行中，3 分片未初始化，-2 分片已删除
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 分片数字ID
	ShardId *int64 `json:"ShardId,omitempty" name:"ShardId"`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 产品类型 Id（过时字段，请勿依赖该值）
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
}

type ShardZoneChooseInfo struct {

	// 主可用区
	MasterZone *ZonesInfo `json:"MasterZone,omitempty" name:"MasterZone"`

	// 可选的从可用区
	SlaveZones []*ZonesInfo `json:"SlaveZones,omitempty" name:"SlaveZones" list`
}

type SlowLogData struct {

	// 语句校验和，用于查询详情
	CheckSum *string `json:"CheckSum,omitempty" name:"CheckSum"`

	// 数据库名称
	Db *string `json:"Db,omitempty" name:"Db"`

	// 抽象的SQL语句
	FingerPrint *string `json:"FingerPrint,omitempty" name:"FingerPrint"`

	// 平均的锁时间
	LockTimeAvg *string `json:"LockTimeAvg,omitempty" name:"LockTimeAvg"`

	// 最大锁时间
	LockTimeMax *string `json:"LockTimeMax,omitempty" name:"LockTimeMax"`

	// 最小锁时间
	LockTimeMin *string `json:"LockTimeMin,omitempty" name:"LockTimeMin"`

	// 锁时间总和
	LockTimeSum *string `json:"LockTimeSum,omitempty" name:"LockTimeSum"`

	// 查询次数
	QueryCount *string `json:"QueryCount,omitempty" name:"QueryCount"`

	// 平均查询时间
	QueryTimeAvg *string `json:"QueryTimeAvg,omitempty" name:"QueryTimeAvg"`

	// 最大查询时间
	QueryTimeMax *string `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`

	// 最小查询时间
	QueryTimeMin *string `json:"QueryTimeMin,omitempty" name:"QueryTimeMin"`

	// 查询时间总和
	QueryTimeSum *string `json:"QueryTimeSum,omitempty" name:"QueryTimeSum"`

	// 扫描行数
	RowsExaminedSum *string `json:"RowsExaminedSum,omitempty" name:"RowsExaminedSum"`

	// 发送行数
	RowsSentSum *string `json:"RowsSentSum,omitempty" name:"RowsSentSum"`

	// 最后执行时间
	TsMax *string `json:"TsMax,omitempty" name:"TsMax"`

	// 首次执行时间
	TsMin *string `json:"TsMin,omitempty" name:"TsMin"`

	// 帐号
	User *string `json:"User,omitempty" name:"User"`
}

type SpecConfig struct {

	// 规格机型
	Machine *string `json:"Machine,omitempty" name:"Machine"`

	// 规格列表
	SpecConfigInfos []*SpecConfigInfo `json:"SpecConfigInfos,omitempty" name:"SpecConfigInfos" list`
}

type SpecConfigInfo struct {

	// 节点个数，2 表示一主一从，3 表示一主二从
	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 数据盘规格最小值，单位 GB
	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 数据盘规格最大值，单位 GB
	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// 推荐的使用场景
	SuitInfo *string `json:"SuitInfo,omitempty" name:"SuitInfo"`

	// 产品类型 Id
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 最大 Qps 值
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
}

type SplitShardConfig struct {

	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds" list`

	// 数据切分比例
	SplitRate *int64 `json:"SplitRate,omitempty" name:"SplitRate"`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitempty" name:"ShardStorage"`
}

type SqlLogItem struct {

	// 本条日志在消息队列中的偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 执行本条sql的用户。
	User *string `json:"User,omitempty" name:"User"`

	// 执行本条sql的客户端IP+端口。
	Client *string `json:"Client,omitempty" name:"Client"`

	// 数据库名称。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 执行的sql语句。
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// 返回的数据行数。
	SelectRowNum *uint64 `json:"SelectRowNum,omitempty" name:"SelectRowNum"`

	// 影响行数。
	AffectRowNum *uint64 `json:"AffectRowNum,omitempty" name:"AffectRowNum"`

	// Sql执行时间戳。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Sql耗时，单位为毫秒。
	TimeCostMs *uint64 `json:"TimeCostMs,omitempty" name:"TimeCostMs"`

	// Sql返回码，0为成功。
	ResultCode *uint64 `json:"ResultCode,omitempty" name:"ResultCode"`
}

type StartSmartDBARequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DCDB实例的分片ID
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *StartSmartDBARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartSmartDBARequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartSmartDBAResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartSmartDBAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartSmartDBAResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchRollbackInstanceRequest struct {
	*tchttp.BaseRequest

	// 源/旧实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标实例ID
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`
}

func (r *SwitchRollbackInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchRollbackInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchRollbackInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchRollbackInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchRollbackInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TableColumn struct {

	// 列名称
	Col *string `json:"Col,omitempty" name:"Col"`

	// 列类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

type TerminateDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 Id，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TerminateDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateDedicatedDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程Id
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateDedicatedDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TmpInstance struct {

	// 应用ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例备注
	InstanceRemark *string `json:"InstanceRemark,omitempty" name:"InstanceRemark"`

	// 0:非临时实例 ,1:无效临时实例, 2:回档成功的有效临时实例
	TempType *int64 `json:"TempType,omitempty" name:"TempType"`

	// 实例状态,0:待初始化,1:流程处理中,2:有效状态,-1:已隔离，-2：已下线
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例虚IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例虚端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 有效期结束时间
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// 源实例 ID，形如：tdsql-ow728lmc。
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// 实例状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 实例所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例所在可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type UpgradeDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitempty" name:"SplitShardConfig"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *UpgradeDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDedicatedDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 升级类型，取值为ADD，SPLIT和EXPAND。ADD-添加分片；SPLIT-切分某个分片；EXPAND-垂直扩容某个分片
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 实例ID，形如 dcdbt-mlfjm74h
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 当UpgradeType取值为ADD时，添加分片的配置参数
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitempty" name:"AddShardConfig"`

	// 当UpgradeType取值为EXPAND时，垂直扩容分片的配置参数
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitempty" name:"ExpandShardConfig"`

	// 当UpgradeType取值为SPLIT时，切分分片的配置参数
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitempty" name:"SplitShardConfig"`
}

func (r *UpgradeDedicatedDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDedicatedDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDedicatedDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDedicatedDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDedicatedDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dcdbt-87zif6ha
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级类型，取值为ADD，SPLIT或者EXPAND
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// UpgradeType为ADD类型时，需要填写该参数
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitempty" name:"AddShardConfig"`

	// UpgradeType为EXPAND类型时，需要填写该参数
	ExpandShardConfig []*ExpandShardConfig `json:"ExpandShardConfig,omitempty" name:"ExpandShardConfig" list`

	// UpgradeType为SPLIT类型时，需要填写该参数
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitempty" name:"SplitShardConfig"`

	// 切换开始时间，格式如: "2019-12-12 07:00:00"。开始时间必须在当前时间一个小时以后，3天以内。
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`

	// 切换结束时间,  格式如: "2019-12-12 07:15:00"，结束时间必须大于开始时间。
	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`

	// 是否自动重试。 0：不自动重试  1：自动重试
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitempty" name:"SwitchAutoRetry"`
}

func (r *UpgradeHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeHourDCDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，形如dcdbt-87zif6ha
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeHourDCDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZonesInfo struct {

	// 可用区英文ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区数字ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区中文名
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}
