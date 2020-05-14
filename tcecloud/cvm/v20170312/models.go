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

package v20170312

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActionTimer struct {

	// 扩展数据
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// 定时器
	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`

	// 执行时间
	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
}

type Address struct {

	// `EIP`的`ID`，是`EIP`的唯一标识。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// `EIP`名称。
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// `EIP`状态。
	AddressState *string `json:"AddressState,omitempty" name:"AddressState"`

	// 弹性外网IP
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`

	// 绑定的资源实例`ID`。可能是一个`CVM`，`NAT`，或是弹性网卡。
	BindedResourceId *string `json:"BindedResourceId,omitempty" name:"BindedResourceId"`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type AddressChargePrepaid struct {

	// 购买实例的时长
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标志
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressCount *int64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// 1
	IspId *int64 `json:"IspId,omitempty" name:"IspId"`

	// 1
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// 1
	VipSet []*string `json:"VipSet,omitempty" name:"VipSet" list`

	// 1
	TgwGroup *string `json:"TgwGroup,omitempty" name:"TgwGroup"`

	// 1
	ApplySilence *int64 `json:"ApplySilence,omitempty" name:"ApplySilence"`

	// 1
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 1
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 1
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`

	// 1
	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 申请到的 EIP 的唯一 ID 列表。
		AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateHostsRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 用于保证请求幂等性的字符串
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式）。
	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`

	// CDH实例机型
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// 购买CDH实例数量
	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`

	// 是否跳过实际执行逻辑
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 购买来源
	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *AllocateHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新创建云子机的实例id列表。
		HostIdSet []*string `json:"HostIdSet,omitempty" name:"HostIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 1
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。<br><br>可以通过以下方式获取可用的实例ID：<br><li>通过登录控制台查询实例ID。<br><li>通过调用接口 DescribeInstances ，取返回信息中的`InstanceId`获取实例ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 一个或多个待操作的密钥对ID，每次请求批量密钥对的上限为100。密钥对ID形如：`skey-11112222`。<br><br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录控制台查询密钥ID。<br><li>通过调用接口 DescribeKeyPairs ，取返回信息中的`KeyId`获取密钥对ID。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机。<br><li>FALSE：表示在正常关机失败后不进行强制关机。<br><br>默认取值：FALSE。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *AssociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateInstancesKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateInstancesKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 1
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuditMarketImageRequest struct {
	*tchttp.BaseRequest

	// 1
	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`

	// 1
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *AuditMarketImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuditMarketImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuditMarketImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditMarketImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuditMarketImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelAuditMarketImageRequest struct {
	*tchttp.BaseRequest

	// 1
	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`

	// 1
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *CancelAuditMarketImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelAuditMarketImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelAuditMarketImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelAuditMarketImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelAuditMarketImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type CheckInstancesConnectivityRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *CheckInstancesConnectivityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckInstancesConnectivityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckInstancesConnectivityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例连通性结果集合。
		InstanceConnectivitySet []*InstanceConnectivity `json:"InstanceConnectivitySet,omitempty" name:"InstanceConnectivitySet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckInstancesConnectivityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckInstancesConnectivityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClientSysInfo struct {

	// 操作系统
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 操作系统版本
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// 需要导入的系统盘数据盘信息
	DiskSize []*uint64 `json:"DiskSize,omitempty" name:"DiskSize" list`

	// 额外信息
	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
}

type ColdMigrateInstanceRequest struct {
	*tchttp.BaseRequest

	// 迁移的系统盘镜像COS链接
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 迁移目的实例ID。迁移完成后，该实例将运行迁入的操作系统。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务名称，用于在控制台展示。
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 用于测试参数合法性。默认为`false`
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *ColdMigrateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ColdMigrateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID，用于查询任务状态、进度。
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ColdMigrateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ColdMigrateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateRequest struct {
	*tchttp.BaseRequest

	// 待迁移的实例Id如："ins-38d920cd"
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标母机数组
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`

	// 要嵌入的售卖池，默认不变更
	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`

	// 需要跨可用迁移，可传入可用区列表
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// 最大迁移带宽
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 最大迁移超时
	MaxTimeout *int64 `json:"MaxTimeout,omitempty" name:"MaxTimeout"`

	// 是否保留hostname
	HostNameReserved *bool `json:"HostNameReserved,omitempty" name:"HostNameReserved"`
}

func (r *ColdMigrateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ColdMigrateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ColdMigrateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ColdMigrateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyInstanceDiskRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1
	DestinationDiskId *string `json:"DestinationDiskId,omitempty" name:"DestinationDiskId"`

	// 1
	SourceDiskId *string `json:"SourceDiskId,omitempty" name:"SourceDiskId"`
}

func (r *CopyInstanceDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyInstanceDiskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyInstanceDiskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyInstanceDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyInstanceDiskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分散置放群组类型，取值范围：<br><li>HOST：物理机<br><li>SW：交换机<br><li>RACK：机架
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 容灾组id。
		DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

		// 容灾组类型，与入参一致。
		Type *string `json:"Type,omitempty" name:"Type"`

		// 分散置放群组名称，长度1-60个字符，支持中、英文。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 置放群组内可容纳的云服务器数量。
		CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitempty" name:"CvmQuotaTotal"`

		// 置放群组内已有的云服务器数量。
		CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

		// 置放群组创建时间。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageRequest struct {
	*tchttp.BaseRequest

	// 需要制作镜像的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 软关机失败时是否执行强制关机以制作镜像
	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`

	// 创建Windows镜像时是否启用Sysprep
	Sysprep *string `json:"Sysprep,omitempty" name:"Sysprep"`

	// 实例处于运行中时，是否允许关机执行制作镜像任务。
	Reboot *string `json:"Reboot,omitempty" name:"Reboot"`
}

func (r *CreateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest

	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// 密钥对创建后所属的项目ID。<br><br>可以通过以下方式获取项目ID：<br><li>通过项目列表查询项目ID。<br><li>通过调用接口 DescribeProject ，取返回信息中的`projectId `获取项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateKeyPairRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥对信息。
		KeyPair *KeyPair `json:"KeyPair,omitempty" name:"KeyPair"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateKeyPairResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：CVM实例配置。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘类型。数据盘类型限制详见CVM实例配置。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。<br><br>该参数对`ResizeInstanceDisk`接口无效。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 数据盘是否随子机销毁。取值范围：
	// <li>TRUE：子机销毁时，销毁数据盘，只支持按小时后付费云盘
	// <li>FALSE：子机销毁时，保留数据盘<br>
	// 默认取值：TRUE<br>
	// 该参数目前仅用于 `RunInstances` 接口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// 数据盘指定的存储池。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
}

type DataDisks struct {

	// 数据盘镜像cos url
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 数据盘大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 数据盘对应设备名，目前是DiskId
	Device *string `json:"Device,omitempty" name:"Device"`
}

type DeleteDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 容灾组id列表，可通过DescribeDisasterRecoverGroups接口查询。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`
}

func (r *DeleteDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest

	// 准备删除的镜像Id列表
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds" list`
}

func (r *DeleteImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 1
	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds" list`
}

func (r *DeleteInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstancesActionTimerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstancesActionTimerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的密钥对ID。每次请求批量密钥对的上限为100。<br><br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录控制台查询密钥ID。<br><li>通过调用接口 DescribeKeyPairs ，取返回信息中的 `KeyId` 获取密钥对ID。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`
}

func (r *DeleteKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAttributesRequest struct {
	*tchttp.BaseRequest

	// 1
	AttributeName []*string `json:"AttributeName,omitempty" name:"AttributeName" list`
}

func (r *DescribeAccountAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressBandwidthConfigsRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAddressBandwidthConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressBandwidthConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressBandwidthConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressBandwidthConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressBandwidthConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账户 EIP 配额信息。
		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的 EIP 数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// EIP 详细信息列表。
		AddressSet []*Address `json:"AddressSet,omitempty" name:"AddressSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDisasterRecoverGroupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisasterRecoverGroupQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可创建置放群组数量的上限。
		GroupQuota *int64 `json:"GroupQuota,omitempty" name:"GroupQuota"`

		// 当前用户已经创建的置放群组数量。
		CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

		// 物理机类型容灾组内实例的配额数。
		CvmInHostGroupQuota *int64 `json:"CvmInHostGroupQuota,omitempty" name:"CvmInHostGroupQuota"`

		// 交换机类型容灾组内实例的配额数。
		CvmInSwGroupQuota *int64 `json:"CvmInSwGroupQuota,omitempty" name:"CvmInSwGroupQuota"`

		// 机架类型容灾组内实例的配额数。
		CvmInRackGroupQuota *int64 `json:"CvmInRackGroupQuota,omitempty" name:"CvmInRackGroupQuota"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisasterRecoverGroupQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest

	// 容灾组id列表。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`

	// 容灾组名称，支持模糊匹配。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDisasterRecoverGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisasterRecoverGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 容灾组信息列表
		DisasterRecoverGroupSet []*DisasterRecoverGroup `json:"DisasterRecoverGroupSet,omitempty" name:"DisasterRecoverGroupSet" list`

		// 用户置放群组总量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisasterRecoverGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的cdh实例总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// cdh实例详细信息列表
		HostSet []*HostItem `json:"HostSet,omitempty" name:"HostSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账户的镜像配额
		ImageNumQuota *int64 `json:"ImageNumQuota,omitempty" name:"ImageNumQuota"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSharePermissionRequest struct {
	*tchttp.BaseRequest

	// 需要共享的镜像Id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSharePermissionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像共享信息
		SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitempty" name:"SharePermissionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSharePermissionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSnapshotStatusRequest struct {
	*tchttp.BaseRequest

	// 需要查看的镜像Id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeImageSnapshotStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSnapshotStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSnapshotStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSnapshotStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSnapshotStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesAttributeRequest struct {
	*tchttp.BaseRequest

	// 1
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds" list`
}

func (r *DescribeImagesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// unImgId到deviceImageId的映射的数组
		ImageAttributeSet *int64 `json:"ImageAttributeSet,omitempty" name:"ImageAttributeSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest

	// 镜像ID列表 。镜像ID如：`img-gvbnzy6f`。array型参数的格式可以参考API简介。镜像ID可以通过如下方式获取：<br><li>通过DescribeImages接口返回的`ImageId`获取。<br><li>通过镜像控制台获取。
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds" list`

	// 过滤条件。其中Filters的上限为10，Filters.Values的上限为5。 注意：不可以同时指定ImageIds和Filters。可选值包括 `image-id`, `image-type`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于Offset详见API简介。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。关于Limit详见API简介。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 实例类型，如 `S1.SMALL1`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 内部参数
	IsInner *string `json:"IsInner,omitempty" name:"IsInner"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 一个关于镜像详细信息的结构体，主要包括镜像的主要状态与属性。
		ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet" list`

		// 符合要求的镜像数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImportImageOsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImportImageOsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImportImageOsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImportImageOsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 支持的导入镜像的操作系统类型
		ImportImageOsListSupported []*string `json:"ImportImageOsListSupported,omitempty" name:"ImportImageOsListSupported" list`

		// 支持的导入镜像的操作系统版本
		ImportImageOsVersionSupported []*string `json:"ImportImageOsVersionSupported,omitempty" name:"ImportImageOsVersionSupported" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportImageOsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImportImageOsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImportSnapshotTaskRequest struct {
	*tchttp.BaseRequest

	// 1
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeImportSnapshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImportSnapshotTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImportSnapshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportSnapshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImportSnapshotTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttachedDevicesRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceAttachedDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceAttachedDevicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttachedDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAttachedDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceAttachedDevicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceChargeTypeConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceChargeTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceChargeTypeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceChargeTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计费类型配置
		InstanceChargeTypeConfigSet []*InstanceChargeTypeConfig `json:"InstanceChargeTypeConfigSet,omitempty" name:"InstanceChargeTypeConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceChargeTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceChargeTypeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceConfigInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例静态配置信息列表。
		InstanceConfigInfos []*InstanceConfigInfoItem `json:"InstanceConfigInfos,omitempty" name:"InstanceConfigInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceConfigInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceFamilyConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceFamilyConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例机型组配置的列表信息
		InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet,omitempty" name:"InstanceFamilyConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceFamilyConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceFamilyConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceInternetBandwidthConfigsRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceInternetBandwidthConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceInternetBandwidthConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceInternetBandwidthConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 带宽配置信息列表。
		InternetBandwidthConfigSet []*InternetBandwidthConfig `json:"InternetBandwidthConfigSet,omitempty" name:"InternetBandwidthConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceInternetBandwidthConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceInternetBandwidthConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationLogsRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeInstanceOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatisticsRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeInstanceStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为1。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeInstanceTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例机型配置列表。
		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeNameConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceTypeNameConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeNameConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeNameConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例类型名称列表。
		InstanceTypeNameConfigSet []*InstanceTypeNameConfig `json:"InstanceTypeNameConfigSet,omitempty" name:"InstanceTypeNameConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeNameConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeNameConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeQuotaRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeInstanceTypeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例机型配额列表。
		InstanceTypeQuotaSet []*InstanceTypeQuota `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 定时器id数组
	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds" list`

	// 实例id数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 定时任务执行之间，格式如：2018-05-01 19:00:00，必须大于当前时间5分钟。
	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`

	// 执行时间的结束范围，用于条件筛选，格式如2018-05-01 19:00:00。
	EndActionTime *string `json:"EndActionTime,omitempty" name:"EndActionTime"`

	// 执行时间的开始范围，用于条件筛选，格式如2018-05-01 19:00:00。
	StartActionTime *string `json:"StartActionTime,omitempty" name:"StartActionTime"`

	// 定时器状态列表数组，可选值为：UNDO（未执行）， DOING（正在执行i），DONE（执行完成）。
	StatusList []*string `json:"StatusList,omitempty" name:"StatusList" list`
}

func (r *DescribeInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesActionTimerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 定时器数组
		ActionTimers []*ActionTimer `json:"ActionTimers,omitempty" name:"ActionTimers" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesActionTimerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例属性列表，标识实例id，实例订单Id。
		InstanceAttributeSet []*KeyPair `json:"InstanceAttributeSet,omitempty" name:"InstanceAttributeSet" list`

		// 拉取实例的个数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesCreateImageAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeInstancesCreateImageAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesCreateImageAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesCreateImageAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesCreateImageAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesCreateImageAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesOperationLimitRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *DescribeInstancesOperationLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesOperationLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesOperationLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesOperationLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesOperationLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRecentFailedOperationRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRecentFailedOperationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRecentFailedOperationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRecentFailedOperationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 最近失败操作详情列表。
		InstancesRecentFailedOperationSet []*InstancesRecentFailedOperationSet `json:"InstancesRecentFailedOperationSet,omitempty" name:"InstancesRecentFailedOperationSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesRecentFailedOperationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRecentFailedOperationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`ins-11112222`。此参数的具体格式可参考API简介的`id.N`一节）。每次请求的实例的上限为100。参数不支持同时指定`InstanceIds`和`Filters`。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 内部参数，数字型vpcId列表 。
	InnerVpcIds []*int64 `json:"InnerVpcIds,omitempty" name:"InnerVpcIds" list`

	// 内部参数，数字型subnetId列表。
	InnerSubnetIds []*int64 `json:"InnerSubnetIds,omitempty" name:"InnerSubnetIds" list`

	// 内部参数，精确内外网IP列表。
	IpAddresses []*string `json:"IpAddresses,omitempty" name:"IpAddresses" list`

	// 内部参数，模糊内外网IP。
	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`

	// 内部参数，模糊实例别名。
	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`
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

		// 符合条件的实例数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息列表。
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

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

type DescribeInstancesReturnableRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesReturnableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesReturnableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesReturnableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可退还实例详细信息列表。
		InstanceReturnableSet []*InstanceReturnable `json:"InstanceReturnableSet,omitempty" name:"InstanceReturnableSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesReturnableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesReturnableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesStatusRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`ins-11112222`。此参数的具体格式可参考API简介的`id.N`一节）。每次请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例状态数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例状态 列表。
		InstanceStatusSet []*InstanceStatus `json:"InstanceStatusSet,omitempty" name:"InstanceStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetChargeTypeConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInternetChargeTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInternetChargeTypeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetChargeTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络计费类型配置。
		InternetChargeTypeConfigSet []*InternetChargeTypeConfig `json:"InternetChargeTypeConfigSet,omitempty" name:"InternetChargeTypeConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternetChargeTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInternetChargeTypeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsAttributeRequest struct {
	*tchttp.BaseRequest

	// 1
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeKeyPairsAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyPairsAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyPairsAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyPairsAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 密钥对ID，密钥对ID形如：`skey-11112222`（此接口支持同时传入多个ID进行过滤。此参数的具体格式可参考 API 简介的 `id.N` 一节）。参数不支持同时指定 `KeyIds` 和 `Filters`。<br> 密钥对ID可以通过登录控制台查询。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 过滤条件，详见密钥对过滤条件表。参数不支持同时指定 `KeyIds` 和 `Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的密钥对数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 密钥对详细信息列表。
		KeyPairSet []*KeyPair `json:"KeyPairSet,omitempty" name:"KeyPairSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMarketImagesRequest struct {
	*tchttp.BaseRequest

	// 需要查询的市场镜像的镜像Id
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds" list`

	// 过滤条件。只支持按镜像Id过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移值，用于分页
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 返回镜像的数量。
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// 内部用户
	IsInner *string `json:"IsInner,omitempty" name:"IsInner"`
}

func (r *DescribeMarketImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMarketImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMarketImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMarketImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMarketImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 服务迁移任务taskId
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 控制台参数，用于返回额外信息
	IsInner *string `json:"IsInner,omitempty" name:"IsInner"`
}

func (r *DescribeMigrateTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 迁移进度。
		Progress *float64 `json:"Progress,omitempty" name:"Progress"`

		// 任务状态。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务名称。
		JobName *string `json:"JobName,omitempty" name:"JobName"`

		// 任务ID。
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 用户AppId。
		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

		// 账户ID。
		Uin *string `json:"Uin,omitempty" name:"Uin"`

		// CVM子机uuid。
		Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

		// 实例ID。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 镜像COS链接。
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// 镜像大小。
		DataSize *uint64 `json:"DataSize,omitempty" name:"DataSize"`

		// 地域信息。
		Region *string `json:"Region,omitempty" name:"Region"`

		// 任务创建时间。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 任务结束时间。
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 任务类型。
		Action *string `json:"Action,omitempty" name:"Action"`

		// 任务失败的错误原因
		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrateTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkSharingGroupsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNetworkSharingGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkSharingGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkSharingGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkSharingGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkSharingGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecomendedZonesRequest struct {
	*tchttp.BaseRequest

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的子机创建，具体取值可通过调用接口DescribeInstanceTypeConfigs来获得最新的规格表或参见实例类型描述。若不指定该参数，则默认机型为S1.SMALL1。<br><li>对于付费模式为CDHPAID的子机创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的子机，该参数应该为CDH_1C1G。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDHPAID：独享母机付费（基于专用宿主机创建，宿主机部分的资源不收费）<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *DescribeRecomendedZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecomendedZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecomendedZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 按照实例可购买配额，从高到低排序的按量计费可用区列表。
		PostPaidZoneSet []*string `json:"PostPaidZoneSet,omitempty" name:"PostPaidZoneSet" list`

		// 按照实例可购买配额，从高到低排序的包年包月可用区列表。
		PrePaidZoneSet []*string `json:"PrePaidZoneSet,omitempty" name:"PrePaidZoneSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecomendedZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecomendedZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecommendedZonesRequest struct {
	*tchttp.BaseRequest

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的子机创建，具体取值可通过调用接口DescribeInstanceTypeConfigs来获得最新的规格表或参见实例类型描述。若不指定该参数，则默认机型为S1.SMALL1。<br><li>对于付费模式为CDHPAID的子机创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的子机，该参数应该为CDH_1C1G。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDHPAID：独享母机付费（基于专用宿主机创建，宿主机部分的资源不收费）<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *DescribeRecommendedZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecommendedZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecommendedZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 按照实例可购买配额，从高到低排序的按量计费可用区列表。
		PostPaidZoneSet []*string `json:"PostPaidZoneSet,omitempty" name:"PostPaidZoneSet" list`

		// 按照实例可购买配额，从高到低排序的包年包月可用区列表。
		PrePaidZoneSet []*string `json:"PrePaidZoneSet,omitempty" name:"PrePaidZoneSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecommendedZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecommendedZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地域数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 地域列表信息
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeResourcesOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourcesOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourcesOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAvailableZonesRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeUserAvailableZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserAvailableZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAvailableZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAvailableZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserAvailableZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGlobalConfigsRequest struct {
	*tchttp.BaseRequest

	// 全局配置名称列表。
	Names []*string `json:"Names,omitempty" name:"Names" list`
}

func (r *DescribeUserGlobalConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserGlobalConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGlobalConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户全局配置列表。
		ConfigSet []*KvType `json:"ConfigSet,omitempty" name:"ConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGlobalConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserGlobalConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInstanceQuotaRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeUserInstanceQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserInstanceQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInstanceQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用于标识用户在可用区下包年包月、按量计费配额。
		DescribeUserInstanceQuota []*KeyPair `json:"DescribeUserInstanceQuota,omitempty" name:"DescribeUserInstanceQuota" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInstanceQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserInstanceQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserMigrateTasksRequest struct {
	*tchttp.BaseRequest

	// 查询的迁移任务类型，可选`ColdMigrateInstance`，`ImportCbs`和`OfflineMigrate`，其中`OfflineMigrate`代表两种离线迁移任务一起查询。
	MigrateTaskType *string `json:"MigrateTaskType,omitempty" name:"MigrateTaskType"`

	// 过滤条件，可选instance-id，job-name，job-id。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 任务数量限制，用于分页。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 任务起始数，用于分页。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUserMigrateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserMigrateTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserMigrateTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserMigrateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserMigrateTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserZoneStatusRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeUserZoneStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserZoneStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserZoneStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用区计费类型状态列表。
		UserZoneStatus []*UserZoneStatusItem `json:"UserZoneStatus,omitempty" name:"UserZoneStatus" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserZoneStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserZoneStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneCdhInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 可用区
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneCdhInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneCdhInstanceConfigInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneCdhInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专用宿主机机型配置信息列表
		HostTypeQuotaSet *string `json:"HostTypeQuotaSet,omitempty" name:"HostTypeQuotaSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneCdhInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneCdhInstanceConfigInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneCpuQuotaRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeZoneCpuQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneCpuQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneCpuQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用区 CPU 配额信息。
		ZoneCpuQuotaSet []*ZoneCpuQuota `json:"ZoneCpuQuotaSet,omitempty" name:"ZoneCpuQuotaSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneCpuQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneCpuQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneHostConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 可用区过滤条件
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneHostConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneHostConfigInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneHostConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专用宿主机机型配置信息列表
		HostTypeQuotaSet []*string `json:"HostTypeQuotaSet,omitempty" name:"HostTypeQuotaSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneHostConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneHostConfigInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneHostForSellStatusRequest struct {
	*tchttp.BaseRequest

	// 过滤条件
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneHostForSellStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneHostForSellStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneHostForSellStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专用宿主机可用区售罄状态
		HostForSellZoneStatus *string `json:"HostForSellZoneStatus,omitempty" name:"HostForSellZoneStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneHostForSellStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneHostForSellStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeZoneInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用区机型配置列表。
		InstanceTypeQuotaSet *InstanceTypeQuota `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用区数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可用区列表信息
		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 1
	KeepAddressIdBindWithEniPip *bool `json:"KeepAddressIdBindWithEniPip,omitempty" name:"KeepAddressIdBindWithEniPip"`

	// 1
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitempty" name:"ReallocateNormalPublicIp"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。<br><br>可以通过以下方式获取可用的实例ID：<br><li>通过登录控制台查询实例ID。<br><li>通过调用接口 DescribeInstances ，取返回信息中的 `InstanceId` 获取密钥对ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 密钥对ID列表，每次请求批量密钥对的上限为100。密钥对ID形如：`skey-11112222`。<br><br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录控制台查询密钥ID。<br><li>通过调用接口 DescribeKeyPairs ，取返回信息中的 `KeyId` 获取密钥对ID。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机。<br><li>FALSE：表示在正常关机失败后不进行强制关机。<br><br>默认取值：FALSE。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *DisassociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateInstancesKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateInstancesKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 1
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisasterRecoverGroup struct {

	// 分散置放群组id。
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// 分散置放群组名称，长度1-60个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分散置放群组类型，取值范围：<br><li>HOST：物理机<br><li>SW：交换机<br><li>RACK：机架
	Type *string `json:"Type,omitempty" name:"Type"`

	// 分散置放群组内最大容纳云服务器数量。
	CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitempty" name:"CvmQuotaTotal"`

	// 分散置放群组内云服务器当前数量。
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 分散置放群组内，云服务器id列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 分散置放群组创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type ExitLiveMigrateInstanceRequest struct {
	*tchttp.BaseRequest

	// 需要退出服务热迁移的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 服务迁移是否成功。
	MigrateResult *string `json:"MigrateResult,omitempty" name:"MigrateResult"`
}

func (r *ExitLiveMigrateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExitLiveMigrateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExitLiveMigrateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExitLiveMigrateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExitLiveMigrateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportImageRequest struct {
	*tchttp.BaseRequest

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// COS Bucket名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
}

func (r *ExportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Externals struct {

	// 释放地址
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`
}

type FailoverMigrateRequest struct {
	*tchttp.BaseRequest

	// 需要迁移的实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// lossy参数
	LossyLocal *bool `json:"LossyLocal,omitempty" name:"LossyLocal"`

	// 可指定母机迁移
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`
}

func (r *FailoverMigrateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FailoverMigrateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FailoverMigrateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FailoverMigrateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FailoverMigrateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type HostGoodsDetailItem struct {

	// 请求事务id
	transactionId *string `json:"transactionId,omitempty" name:"transactionId"`

	// 操作名称
	action *string `json:"action,omitempty" name:"action"`

	// 自动续费标记
	autoRenewFlag *uint64 `json:"autoRenewFlag,omitempty" name:"autoRenewFlag"`

	// pid
	pid *uint64 `json:"pid,omitempty" name:"pid"`

	// 购买或续费时长
	timeSpan *uint64 `json:"timeSpan,omitempty" name:"timeSpan"`

	// 时间单位
	timeUnit *string `json:"timeUnit,omitempty" name:"timeUnit"`

	// 资源id
	resourceId *string `json:"resourceId,omitempty" name:"resourceId"`

	// 当前到期时间
	curDeadline *string `json:"curDeadline,omitempty" name:"curDeadline"`

	// 数字签名
	signature *string `json:"signature,omitempty" name:"signature"`

	// 产品信息项列表
	productInfo []*ProductInfoItem `json:"productInfo,omitempty" name:"productInfo" list`
}

type HostGoodsItem struct {

	// goodsCategoryId
	goodsCategoryId *uint64 `json:"goodsCategoryId,omitempty" name:"goodsCategoryId"`

	// 实例付费模式
	payMode *uint64 `json:"payMode,omitempty" name:"payMode"`

	// 发货的实例个数
	goodsNum *uint64 `json:"goodsNum,omitempty" name:"goodsNum"`

	// 地域id
	regionId *uint64 `json:"regionId,omitempty" name:"regionId"`

	// 发起发货的用户uin
	uin *string `json:"uin,omitempty" name:"uin"`

	// 发起发货帐号的所有者uin
	ownerUin *string `json:"ownerUin,omitempty" name:"ownerUin"`

	// 发起发货帐号对应的appId
	appId *uint64 `json:"appId,omitempty" name:"appId"`

	// 项目id
	projectId *uint64 `json:"projectId,omitempty" name:"projectId"`

	// 可用区id
	zoneId *uint64 `json:"zoneId,omitempty" name:"zoneId"`

	// cdh实例详细信息
	goodsDetail *HostGoodsDetailItem `json:"goodsDetail,omitempty" name:"goodsDetail"`
}

type HostItem struct {

	// cdh实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// cdh实例id
	HostId *string `json:"HostId,omitempty" name:"HostId"`

	// cdh实例类型
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// cdh实例名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// cdh实例付费模式
	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`

	// cdh实例自动续费标记
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// cdh实例创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// cdh实例过期时间
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// cdh实例上已创建云子机的实例id列表
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// cdh实例状态
	HostState *string `json:"HostState,omitempty" name:"HostState"`

	// cdh实例ip
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// cdh实例资源信息
	HostResource *HostResource `json:"HostResource,omitempty" name:"HostResource"`
}

type HostOrder struct {

	// 订单所属帐号的应用id
	appId *uint64 `json:"appId,omitempty" name:"appId"`

	// 创建订单的帐号uin
	uin *string `json:"uin,omitempty" name:"uin"`

	// 订单所属帐号的所有者uin
	ownerUin *string `json:"ownerUin,omitempty" name:"ownerUin"`

	// 订单发货的资源信息列表
	goods []*HostGoodsItem `json:"goods,omitempty" name:"goods" list`
}

type HostPrice struct {

	// 描述了cdh实例相关的价格信息
	HostPrice *ItemPrice `json:"HostPrice,omitempty" name:"HostPrice"`
}

type HostResource struct {

	// cdh实例总cpu核数
	CpuTotal *uint64 `json:"CpuTotal,omitempty" name:"CpuTotal"`

	// cdh实例可用cpu核数
	CpuAvailable *uint64 `json:"CpuAvailable,omitempty" name:"CpuAvailable"`

	// cdh实例总内存大小（单位为:GiB）
	MemTotal *float64 `json:"MemTotal,omitempty" name:"MemTotal"`

	// cdh实例可用内存大小（单位为:GiB）
	MemAvailable *float64 `json:"MemAvailable,omitempty" name:"MemAvailable"`

	// cdh实例总磁盘大小（单位为:GiB）
	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`

	// cdh实例可用磁盘大小（单位为:GiB）
	DiskAvailable *uint64 `json:"DiskAvailable,omitempty" name:"DiskAvailable"`
}

type Image struct {

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像操作系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 镜像类型
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 镜像创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 镜像大小
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 镜像架构
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 镜像状态
	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`

	// 镜像来源平台
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 镜像创建者
	ImageCreator *string `json:"ImageCreator,omitempty" name:"ImageCreator"`

	// 镜像来源
	ImageSource *string `json:"ImageSource,omitempty" name:"ImageSource"`
}

type ImportCbsRequest struct {
	*tchttp.BaseRequest

	// 用户需要导入的数据盘镜像存放的COS链接。
	SnapshotUrl *string `json:"SnapshotUrl,omitempty" name:"SnapshotUrl"`

	// 导入目的云盘的ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 用于测试参数合法性。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 任务名称，用于控制台展示。
	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

func (r *ImportCbsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportCbsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportCbsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导入数据盘的任务ID，用于查询任务状态、进度。
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportCbsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportCbsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportFullCvmImageRequest struct {
	*tchttp.BaseRequest

	// 导入任务名称，不超过60个字符。
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 导入镜像的目标子机实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 系统盘镜像cos url，cos地域需要跟子机所在地域保持一致。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 镜像名称，IsCreate为`TRUE`则为必填，不超过20个字符。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像描述，不超过60个字符。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 数据盘镜像数组，最多4个数据盘镜像，传入的数据盘数量必须等于子机上挂载的盘的数量。
	DataDisks []*DataDisks `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 用于测试参数合法性，默认为`FALSE`。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 是否制作整机镜像，默认为`FALSE`。
	IsCreate *bool `json:"IsCreate,omitempty" name:"IsCreate"`
}

func (r *ImportFullCvmImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportFullCvmImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportFullCvmImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID，可利用该ID查询任务状态。
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportFullCvmImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportFullCvmImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportImageRequest struct {
	*tchttp.BaseRequest

	// 导入镜像的操作系统架构，`x86_64` 或 `i386`
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 导入镜像的操作系统类型，通过`DescribeImportImageOs`获取
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 导入镜像的操作系统版本，通过`DescribeImportImageOs`获取
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// 导入镜像存放的cos地址
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 只检查参数，不执行任务
	DryRun *string `json:"DryRun,omitempty" name:"DryRun"`

	// 是否强制导入，参考强制导入镜像
	Force *string `json:"Force,omitempty" name:"Force"`
}

func (r *ImportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`
}

func (r *ImportInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportInstancesActionTimerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 定时器id列表
		ActionTimerIds *string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportInstancesActionTimerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest

	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// 密钥对创建后所属的项目ID。<br><br>可以通过以下方式获取项目ID：<br><li>通过项目列表查询项目ID。<br><li>通过调用接口 DescribeProject，取返回信息中的 `projectId ` 获取项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 密钥对的公钥内容，`OpenSSH RSA` 格式。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

func (r *ImportKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportKeyPairRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥对ID。
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportKeyPairResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportSnapshotRequest struct {
	*tchttp.BaseRequest

	// 制作的快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 数据盘镜像COS链接
	SnapshotUrl *string `json:"SnapshotUrl,omitempty" name:"SnapshotUrl"`

	// 制作的快照大小
	SnapshotSize *int64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`

	// 制作的快照描述
	SnapshotDescription *string `json:"SnapshotDescription,omitempty" name:"SnapshotDescription"`

	// true为仅检查参数，默认为false
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *ImportSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportSnapshotRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID，可以用于查询任务状态
		TaskId []*string `json:"TaskId,omitempty" name:"TaskId" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportSnapshotResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateAddressesRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressCount *int64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// 1
	IspId *int64 `json:"IspId,omitempty" name:"IspId"`

	// 1
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// 1
	VipSet []*string `json:"VipSet,omitempty" name:"VipSet" list`

	// 1
	TgwGroup *string `json:"TgwGroup,omitempty" name:"TgwGroup"`

	// 1
	ApplySilence *int64 `json:"ApplySilence,omitempty" name:"ApplySilence"`

	// 1
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 1
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 1
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`

	// 1
	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *InquiryPriceAllocateAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceAllocateAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceAllocateAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceAllocateAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateHostsRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 用于保证请求幂等性的字符串。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式）。
	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`

	// CDH实例机型。
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// 购买CDH实例数量。
	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`

	// 是否跳过实际执行逻辑
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 购买来源
	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *InquiryPriceAllocateHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceAllocateHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CDH实例创建价格信息
		Price *HostPrice `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceAllocateHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceAllocateHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 1
	Name *string `json:"Name,omitempty" name:"Name"`

	// 1
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *InquiryPriceCreateDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceDeleteDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 1
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`
}

func (r *InquiryPriceDeleteDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceDeleteDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceDeleteDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceDeleteDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceDeleteDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`

	// 1
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 1
	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *InquiryPriceModifyAddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceModifyAddressesBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyAddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceModifyAddressesBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyInstanceInternetChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1
	InternetAccessible *InternetAccessibleModifyChargeType `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *InquiryPriceModifyInstanceInternetChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceModifyInstanceInternetChargeTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyInstanceInternetChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyInstanceInternetChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceModifyInstanceInternetChargeTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyInstancesChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 1
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 1
	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitempty" name:"ModifyPortableDataDisk"`
}

func (r *InquiryPriceModifyInstancesChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceModifyInstancesChargeTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyInstancesChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1
		Price *Price `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyInstancesChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceModifyInstancesChargeTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceQueryDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 1
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`
}

func (r *InquiryPriceQueryDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceQueryDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceQueryDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceQueryDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceQueryDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewAddressesRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressIds *string `json:"AddressIds,omitempty" name:"AddressIds"`

	// 1
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`

	// 1
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 1
	CurrentDeadline *string `json:"CurrentDeadline,omitempty" name:"CurrentDeadline"`
}

func (r *InquiryPriceRenewAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewHostsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds" list`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// 是否跳过实际执行逻辑。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *InquiryPriceRenewHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CDH实例续费价格信息
		Price *HostPrice `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 内部参数，公网带宽相关信息设置。
	InternetAccessible []*InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible" list`

	// 试运行。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 内部参数，续费弹性数据盘。
	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitempty" name:"RenewPortableDataDisk"`
}

func (r *InquiryPriceRenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该参数表示对应配置实例的价格。
		Price *Price `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID。可通过 DescribeInstances API返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定有效的镜像ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录控制台查询；`服务镜像市场`的镜像ID可通过云市场查询。</li><li>通过调用接口 DescribeImages ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *InquiryPriceResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该参数表示重装成对应配置实例的价格。
		Price *Price `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持`InternetMaxBandwidthOut`参数。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 带宽生效的终止时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过`DescribeInstances`接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该参数表示带宽调整为对应大小之后的价格。
		Price *Price `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表实例资源规格对照表，也可以调用查询实例资源规格列表接口获得最新的规格表。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *InquiryPriceResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该参数表示调整成对应机型实例的价格。
		Price *Price `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 待扩容的数据盘配置信息。只支持扩容随实例购买的数据盘，且数据盘类型为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。数据盘容量单位：GB。最小扩容步长：10G。关于数据盘类型的选择请参考硬盘产品简介。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *InquiryPriceResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResizeInstanceDisksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该参数表示磁盘扩容成对应配置的价格。
		Price *Price `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResizeInstanceDisksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRunInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 指定有效的镜像ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录控制台查询；`服务镜像市场`的镜像ID可通过云市场查询。</li><li>通过调用接口 DescribeImages ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口DescribeInstanceTypeConfigs来获得最新的规格表或参见CVM实例配置描述。若不指定该参数，则默认机型为S1.SMALL1。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，当前仅支持购买的时候指定一个数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络ip，那么InstanceCount参数只能为1。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 购买实例数量。取值范围：1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 实例显示名称。如果不指定则默认显示
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用DescribeSecurityGroups的返回值中的sgId字段来获取。若不指定该参数，则默认不绑定安全组<font style=
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 内部参数，购买来源。
	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *InquiryPriceRunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRunInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该参数表示对应配置实例的价格。
		Price *Price `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRunInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceTerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *InquiryPriceTerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceTerminateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceTerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceTerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceTerminateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 1
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// 1
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *InquiryPriceUpdateDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpdateDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpdateDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpdateDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 实例`ID`。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例的CPU核数，单位：核。
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 实例内存容量，单位：`GB`。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例业务状态。取值范围：<br><li>NORMAL：表示正常状态的实例<br><li>EXPIRED：表示过期的实例<br><li>PROTECTIVELY_ISOLATED：表示被安全隔离的实例。
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围：<br><li>`PREPAID`：表示预付费，即包年包月<br><li>`POSTPAID_BY_HOUR`：表示后付费，即按量计费<br><li>`CDHPAID`：`CDH`付费，即只对`CDH`计费，不对`CDH`上的实例计费。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例数据盘信息。只包含随实例购买的数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 实例主网卡的内网`IP`列表。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// 实例主网卡的公网`IP`列表。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 实例所属虚拟私有网络信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 生产实例所使用的镜像`ID`。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 自动续费标识。取值范围：<br><li>`NOTIFY_AND_MANUAL_RENEW`：表示通知即将过期，但不自动续费<br><li>`NOTIFY_AND_AUTO_RENEW`：表示通知即将过期，而且自动续费<br><li>`DISABLE_NOTIFY_AND_MANUAL_RENEW`：表示不通知即将过期，也不自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 到期时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type InstanceChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceChargeTypeConfig struct {

	// 实例计费类型
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例计费类型描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type InstanceConfigInfoItem struct {

	// 实例规格。
	type *string `json:"type,omitempty" name:"type"`

	// 实例规格名称。
	typeName *string `json:"typeName,omitempty" name:"typeName"`

	// 优先级。
	order *int64 `json:"order,omitempty" name:"order"`

	// 实例族信息列表。
	instanceFamilies []*InstanceFamilyItem `json:"instanceFamilies,omitempty" name:"instanceFamilies" list`
}

type InstanceConnectivity struct {

	// 实例`ID`
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 默认远程登录端口连通性状态
	DefaultLoginPortConnectivity *bool `json:"DefaultLoginPortConnectivity,omitempty" name:"DefaultLoginPortConnectivity"`

	// ping包是否可达
	ICMPConnectivity *bool `json:"ICMPConnectivity,omitempty" name:"ICMPConnectivity"`
}

type InstanceFamilyConfig struct {

	// 机型族名称的中文全称。
	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`

	// 机型族名称的英文简称。
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceFamilyItem struct {

	// 实例族。
	instanceFamily *string `json:"instanceFamily,omitempty" name:"instanceFamily"`

	// 优先级。
	order *int64 `json:"order,omitempty" name:"order"`

	// 实例类型信息列表。
	instanceTypes []*InstanceTypeItem `json:"instanceTypes,omitempty" name:"instanceTypes" list`
}

type InstanceOrder struct {
}

type InstanceReturnable struct {

	// 实例`ID`。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例是否可退还。
	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`

	// 实例退还失败错误码。
	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`

	// 实例退还失败错误信息。
	ReturnFailMessage *string `json:"ReturnFailMessage,omitempty" name:"ReturnFailMessage"`
}

type InstanceStatus struct {

	// 实例`ID`。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例状态。
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
}

type InstanceTypeConfig struct {

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// GPU核数，单位：核。
	GPU *int64 `json:"GPU,omitempty" name:"GPU"`

	// CPU核数，单位：核。
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 内存容量，单位：`GB`。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 是否支持云硬盘。取值范围：<br><li>`TRUE`：表示支持云硬盘；<br><li>`FALSE`：表示不支持云硬盘。
	CbsSupport *string `json:"CbsSupport,omitempty" name:"CbsSupport"`

	// 机型状态。取值范围：<br><li>`AVAILABLE`：表示机型可用；<br><li>`UNAVAILABLE`：表示机型不可用。
	InstanceTypeState *string `json:"InstanceTypeState,omitempty" name:"InstanceTypeState"`
}

type InstanceTypeItem struct {

	// 实例类型。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU核数。
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// GPU核数。
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// FPGA核数。
	Fpga *uint64 `json:"Fpga,omitempty" name:"Fpga"`

	// 存储块数。
	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`

	// 网卡数。
	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`

	// 最大带宽。
	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 主频。
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`

	// CPU型号名称。
	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`

	// 包转发率。
	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

	// 外部信息。
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// 备注信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type InstanceTypeNameConfig struct {

	// 是否显示到实例类型列表
	ShowInMenu *bool `json:"ShowInMenu,omitempty" name:"ShowInMenu"`

	// 实例类型中文名
	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`

	// 实例类型
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceTypeQuota struct {

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例机型配额。
	InstanceQuota *int64 `json:"InstanceQuota,omitempty" name:"InstanceQuota"`

	// 实例计费模式。取值范围： <br>*`PREPAID`：表示预付费，即包年包月 <br>* `POSTPAID_BY_HOUR`：表示后付费，即按量计费 * `CDHPAID`：CDH付费，即只对CDH计费，不对CDH上的实例计费。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例类型。
	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`

	// 实例的CPU核数，单位：核。
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 实例内存容量，单位：`GB`。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例的GPU核数，单位：核。
	GPU *int64 `json:"GPU,omitempty" name:"GPU"`
}

type InstancesRecentFailedOperationSet struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 操作事件类型
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 操作事件发生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type InternetAccessible struct {

	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算<br><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费<br><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费<br><li>BANDWIDTH_PACKAGE：带宽包用户<br>默认取值：TRAFFIC_POSTPAID_BY_HOUR。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

type InternetAccessibleModifyChargeType struct {

	// 网络付费模式
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 外网出带宽值
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type InternetBandwidthConfig struct {

	// 开始时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
}

type InternetChargeTypeConfig struct {

	// 网络计费模式。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 网络计费模式描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ItemPrice struct {

	// 后续单价，单位：元。
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 后续计价单元，可取值范围： <br><li>HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）：<br><li>GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`

	// 预支费用的原价，单位：元。
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 预支费用的折扣价，单位：元。
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type KeyPair struct {

	// 密钥对的`ID`，是密钥对的唯一标识。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 密钥对名称。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// 密钥对所属的项目`ID`。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 密钥对描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 密钥对的纯文本公钥。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 密钥对的纯文本私钥。Tce不会保管私钥，请用户自行妥善保存。
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 密钥关联的实例`ID`列表。
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds" list`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type KvType struct {

	// 键的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 键所对应的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LiveMigrateInstanceRequest struct {
	*tchttp.BaseRequest

	// 客户端版本。
	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`

	// 迁移目的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 迁移源实例系统信息。
	SysInfo *ClientSysInfo `json:"SysInfo,omitempty" name:"SysInfo"`

	// 是否忽略网络连通性检查
	IgnoreCheckNetworkConnectivity *bool `json:"IgnoreCheckNetworkConnectivity,omitempty" name:"IgnoreCheckNetworkConnectivity"`
}

func (r *LiveMigrateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveMigrateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 目的实例外网IP。
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// 目的子机密码。
		Password *string `json:"Password,omitempty" name:"Password"`

		// 目的实例内网IP。
		LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveMigrateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveMigrateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateRequest struct {
	*tchttp.BaseRequest

	// 待迁移实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定母机迁移
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`

	// 要迁入的售卖池， PLAIN, OVERSOLD等
	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`

	// 跨可用区迁移，要传入可用区表
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// 最大迁移带宽
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 最大迁移超时
	MaxTimeout *int64 `json:"MaxTimeout,omitempty" name:"MaxTimeout"`
}

func (r *LiveMigrateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveMigrateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveMigrateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveMigrateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 1
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// 1
	EipDirectConnection *string `json:"EipDirectConnection,omitempty" name:"EipDirectConnection"`
}

func (r *ModifyAddressAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`

	// 1
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 1
	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *ModifyAddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressesBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressesBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 容灾组id，可使用DescribeDisasterRecoverGroups接口查询。
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// 容灾组名称，最长128个字符。
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsAttributeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds" list`

	// CDH实例显示名称。可任意命名，但不得超过60个字符。
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 自动续费标识。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyHostsAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHostsAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHostsAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAttributeRequest struct {
	*tchttp.BaseRequest

	// 镜像ID，形如`img-gvbnzy6f`。镜像ID可以通过如下方式获取：<br><li>通过DescribeImages接口返回的`ImageId`获取。<br><li>通过镜像控制台获取。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 设置新的镜像名称；必须满足下列限制：<br> <li> 不得超过20个字符。<br> <li> 镜像名称不能与已有镜像重复。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 设置新的镜像描述；必须满足下列限制：<br> <li> 不得超过60个字符。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

func (r *ModifyImageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionRequest struct {
	*tchttp.BaseRequest

	// 镜像ID，形如`img-gvbnzy6f`。镜像Id可以通过如下方式获取：<br><li>通过DescribeImages接口返回的`ImageId`获取。<br><li>通过镜像控制台获取。 <br>镜像ID必须指定为状态为`NORMAL`的镜像。镜像状态请参考镜像数据表。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 接收分享镜像的账号Id列表，array型参数的格式可以参考API简介。帐号ID不同于QQ号，查询用户帐号ID请查看帐号信息中的帐号ID栏。
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds" list`

	// 操作，包括 `SHARE`，`CANCEL`。其中`SHARE`代表分享操作，`CANCEL`代表取消分享操作。
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}

func (r *ModifyImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSharePermissionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSharePermissionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceInternetChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1
	InternetAccessible *InternetAccessibleModifyChargeType `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *ModifyInstanceInternetChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceInternetChargeTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceInternetChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceInternetChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceInternetChargeTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 指定具体的定时器信息
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// 定时器id列表，可以通过DescribeInstancesActionTimer接口查询。
	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds" list`
}

func (r *ModifyInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesActionTimerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesActionTimerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances` API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例显示名称。可任意命名，但不得超过60个字符。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 内部参数，用户数据。
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 内部参数，安全组Id列表。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups" list`

	// 内部参数，未知。
	ResetNewCreationIdentify *bool `json:"ResetNewCreationIdentify,omitempty" name:"ResetNewCreationIdentify"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 1
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 1
	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitempty" name:"ModifyPortableDataDisk"`
}

func (r *ModifyInstancesChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesChargeTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesChargeTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances` API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 项目ID。项目可以使用AddProject接口创建。后续使用DescribeInstances接口查询实例时，项目ID可用于过滤结果。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances` API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyInstancesRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesRenewFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesRenewFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyPairAttributeRequest struct {
	*tchttp.BaseRequest

	// 密钥对ID，密钥对ID形如：`skey-11112222`。<br><br>可以通过以下方式获取可用的密钥 ID：<br><li>通过登录控制台查询密钥 ID。<br><li>通过调用接口 DescribeKeyPairs ，取返回信息中的 `KeyId` 获取密钥对 ID。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 修改后的密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// 修改后的密钥对描述信息。可任意命名，但不得超过60个字符。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyKeyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyKeyPairAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyKeyPairAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGlobalConfigsRequest struct {
	*tchttp.BaseRequest

	// 用户全局配置列表，目前配置仅支持StoppedMode一个key值，表示用户默认关机不收费配置，KEEP_CHARGING为关机收费，STOP_CHARGING为关机不收费，默认为KEEP_CHARGING。
	Configs []*KvType `json:"Configs,omitempty" name:"Configs" list`
}

func (r *ModifyUserGlobalConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUserGlobalConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGlobalConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserGlobalConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUserGlobalConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Placement struct {

	// 实例所属的可用区ID。该参数也可以通过调用  DescribeZones 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所属的专用宿主机ID列表。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。当前暂不支持。
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds" list`
}

type Price struct {

	// 描述了实例价格。
	InstancePrice *ItemPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// 描述了网络价格。
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type ProductInfoItem struct {

	// 信息项名称
	name *string `json:"name,omitempty" name:"name"`

	// 信息项对应的值
	value *string `json:"value,omitempty" name:"value"`
}

type QueryDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 1
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`
}

func (r *QueryDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryFlowLogsRequest struct {
	*tchttp.BaseRequest

	// 查询日志范围的起始时间；例：2019-07-05 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询日志范围的截止时间；例：2019-07-05 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryFlowLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志流水信息。
		FlowLogSet []*string `json:"FlowLogSet,omitempty" name:"FlowLogSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryFlowLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryHostsRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 租户端云服务器实例总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 租户端云服务器主机列表
		Hosts []*string `json:"Hosts,omitempty" name:"Hosts" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 1
	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds" list`

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`

	// 1
	EndActionTime *string `json:"EndActionTime,omitempty" name:"EndActionTime"`

	// 1
	StartActionTime *string `json:"StartActionTime,omitempty" name:"StartActionTime"`

	// 1
	StatusList []*string `json:"StatusList,omitempty" name:"StatusList" list`
}

func (r *QueryInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryInstancesActionTimerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryInstancesActionTimerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryInstancesRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 租户端云服务器实例总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 租户端云服务器实例列表
		Instances []*string `json:"Instances,omitempty" name:"Instances" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryMigrateTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *QueryMigrateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryMigrateTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryMigrateTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMigrateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryMigrateTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Quota struct {

	// 配额名称，取值范围：<br><li>`TOTAL_EIP_QUOTA`：用户当前地域下EIP的配额数；<br><li>`DAILY_EIP_APPLY`：用户当前地域下今日申购次数；<br><li>`DAILY_PUBLIC_IP_ASSIGN`：用户当前地域下，重新分配公网 IP次数。
	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`

	// 当前数量
	QuotaCurrent *int64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`

	// 配额数量
	QuotaLimit *int64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 是否在正常重启失败后选择强制重启实例。取值范围：<br><li>TRUE：表示在正常重启失败后进行强制重启<br><li>FALSE：表示在正常重启失败后不进行强制重启<br><br>默认取值：FALSE。
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebootInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebootInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RefreshInternalUserEnvironmentRequest struct {
	*tchttp.BaseRequest
}

func (r *RefreshInternalUserEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefreshInternalUserEnvironmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RefreshInternalUserEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshInternalUserEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefreshInternalUserEnvironmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域名称，例如，ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域描述，例如，华南地区(广州)
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域是否可用状态
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewAddressesRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`

	// 1
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`

	// 1
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 1
	CurrentDeadline *string `json:"CurrentDeadline,omitempty" name:"CurrentDeadline"`
}

func (r *RenewAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewHostsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds" list`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// 是否跳过实际执行逻辑。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *RenewHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 内部参数，公网带宽相关信息设置。
	InternetAccessible []*InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible" list`

	// 试运行。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 内部参数，续费弹性数据盘。
	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitempty" name:"RenewPortableDataDisk"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID。可通过 DescribeInstances API返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定有效的镜像ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录控制台查询；`服务镜像市场`的镜像ID可通过云市场查询。</li><li>通过调用接口 DescribeImages ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的 `InstanceId` 获取。 每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持 `InternetMaxBandwidthOut` 参数。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 带宽生效的终止时间。格式： `YYYY-MM-DD` ，例如：`2016-10-30` 。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过 `DescribeInstances`接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances` API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>`Linux`实例密码必须8到16位，至少包括两项`[a-z，A-Z]、[0-9]`和`[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]`中的符号。密码不允许以`/`符号开头。<br><li>`Windows`实例密码必须12到16位，至少包括三项`[a-z]，[A-Z]，[0-9]`和`[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]`中的符号。密码不允许以`/`符号开头。<br><li>如果实例即包含`Linux`实例又包含`Windows`实例，则密码复杂度限制按照`Windows`实例的限制。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 待重置密码的实例操作系统用户名。不得超过64个字符。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表实例资源规格对照表，也可以调用查询实例资源规格列表接口获得最新的规格表。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 待扩容的数据盘配置信息。只支持扩容随实例购买的数据盘，且数据盘类型为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。数据盘容量单位：GB。最小扩容步长：10G。关于数据盘类型的选择请参考硬盘产品简介。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResizeInstanceDisksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResizeInstanceDisksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeInstancesRequest struct {
	*tchttp.BaseRequest

	// 无
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *ResumeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReturnAddressesRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`
}

func (r *ReturnAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReturnAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReturnAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReturnAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReturnAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReturnNormalAddressesRequest struct {
	*tchttp.BaseRequest

	// 1
	AddressIps []*string `json:"AddressIps,omitempty" name:"AddressIps" list`
}

func (r *ReturnNormalAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReturnNormalAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReturnNormalAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReturnNormalAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReturnNormalAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，专用宿主机（对于独享母机付费模式的子机创建）等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 指定有效的镜像ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录控制台查询；`服务镜像市场`的镜像ID可通过云市场查询。</li><li>通过调用接口 DescribeImages ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDHPAID：独享母机付费（基于专用宿主机创建，宿主机部分的资源不收费）<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的子机创建，具体取值可通过调用接口DescribeInstanceTypeConfigs来获得最新的规格表或参见实例类型描述。若不指定该参数，则默认机型为S1.SMALL1。<br><li>对于付费模式为CDHPAID的子机创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的子机，该参数应该为CDH_1C1G。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，当前仅支持购买的时候指定一个数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络ip，那么InstanceCount参数只能为1。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 购买实例数量。取值范围：1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 实例显示名称。如果不指定则默认显示
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 用于指定价格生产，当前主要用于竞价实例
	SpotPrice *string `json:"SpotPrice,omitempty" name:"SpotPrice"`

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。<br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux 等）实例：字符长度为[2, 30]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *RunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例`ID`。返回实例`ID`列表并不代表实例创建成功，可根据 DescribeInstancesStatus 接口查询返回的InstancesSet中对应实例的`ID`的状态来判断创建是否完成；如果实例状态由“准备中”变为“正在运行”，则为创建成功。
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {

	// 是否开启云监控服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {

	// 是否开启云安全服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type SearchUserInstanceRequest struct {
	*tchttp.BaseRequest

	// 1
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchUserInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchUserInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchUserInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchUserInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchUserInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SharePermission struct {

	// 镜像分享时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像分享的账户ID
	Account *string `json:"Account,omitempty" name:"Account"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 是否在正常关闭失败后选择强制关闭实例。取值范围：<br><li>TRUE：表示在正常关闭失败后进行强制关闭<br><li>FALSE：表示在正常关闭失败后不进行强制关闭<br><br>默认取值：FALSE。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SuspendInstancesRequest struct {
	*tchttp.BaseRequest

	// 无
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 无
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
}

func (r *SuspendInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SuspendInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SuspendInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SuspendInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SuspendInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterAllocateHostsRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 用于保证请求幂等性的字符串。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式）。
	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`

	// CDH实例机型。
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// 购买CDH实例数量。
	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`

	// 是否跳过实际执行逻辑
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 购买来源
	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *SwitchParameterAllocateHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterAllocateHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterAllocateHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// cdh订单详细信息
		InstanceOrder *HostOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterAllocateHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterAllocateHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyInstanceInternetChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1
	InternetAccessible *InternetAccessibleModifyChargeType `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterModifyInstanceInternetChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterModifyInstanceInternetChargeTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyInstanceInternetChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例订单详情信息。
		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterModifyInstanceInternetChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterModifyInstanceInternetChargeTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyInstancesChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 1
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 1
	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitempty" name:"ModifyPortableDataDisk"`
}

func (r *SwitchParameterModifyInstancesChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterModifyInstancesChargeTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyInstancesChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例订单详情信息。
		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterModifyInstancesChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterModifyInstancesChargeTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewHostsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds" list`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// 是否跳过实际执行逻辑。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterRenewHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterRenewHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// cdh订单详细信息
		InstanceOrder *HostOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterRenewHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterRenewHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewInstancesRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 1
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 1
	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitempty" name:"RenewPortableDataDisk"`
}

func (r *SwitchParameterRenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterRenewInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例订单详情信息。
		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterRenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterRenewInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 1
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 1
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 1
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterResetInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例订单详情信息。
		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterResetInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 1
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 1
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例订单详情信息。
		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 1
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 1
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterResetInstancesTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例订单详情信息。
		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterResetInstancesTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 1
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterResizeInstanceDisksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例订单详情信息。
		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterResizeInstanceDisksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRunInstancesRequest struct {
	*tchttp.BaseRequest

	// 1
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 1
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 1
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 1
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 1
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 1
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 1
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 1
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 1
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 1
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 1
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 1
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 1
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 1
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 1
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 1
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 1
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 1
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 1
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`

	// 1
	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`

	// 1
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`

	// 1
	DesAction *string `json:"DesAction,omitempty" name:"DesAction"`

	// 1
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// 内部参数，购买来源
	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *SwitchParameterRunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterRunInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例订单详情信息。
		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterRunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchParameterRunInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncImagesRequest struct {
	*tchttp.BaseRequest

	// 镜像ID列表 ，镜像ID可以通过如下方式获取：<br><li>通过DescribeImages接口返回的`ImageId`获取。<br><li>通过镜像控制台获取。<br>镜像ID必须满足限制：<br><li>镜像ID对应的镜像状态必须为`NORMAL`。<br><li>镜像大小小于50GB。<br>镜像状态请参考镜像数据表。
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds" list`

	// 目的同步地域列表；必须满足限制：<br><li>不能为源地域，<br><li>必须是一个合法的Region。<br><li>暂不支持部分地域同步。<br>具体地域参数请参考Region。
	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions" list`
}

func (r *SyncImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {

	// 系统盘类型。系统盘类型限制详见CVM实例配置。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 系统盘大小，单位：GB。默认值为 50
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 系统盘指定的存储池。
	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 内部参数，释放弹性IP。
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`

	// 试运行。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransformAddressRequest struct {
	*tchttp.BaseRequest

	// 1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TransformAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransformAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransformAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransformAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransformAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 1
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// 1
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *UpdateDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceVpcConfigRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 私有网络相关信息配置。通过该参数指定私有网络的ID，子网ID，私有网络ip等信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 是否对运行中的实例选择强制关机。默认为TRUE。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *UpdateInstanceVpcConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceVpcConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceVpcConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstanceVpcConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceVpcConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 1
	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds" list`

	// 1
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`
}

func (r *UpdateInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstancesActionTimerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstancesActionTimerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UserZoneStatusItem struct {

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计费类型
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 售卖状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type VirtualPrivateCloud struct {

	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录控制台查询；也可以调用接口 DescribeVpcEx ，从接口返回中的`unVpcId`字段获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口  DescribeSubnetEx ，从接口返回中的`unSubnetId`字段获取。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`

	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// 为弹性网卡指定随机生成的 IPv6 地址数量。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

type ZoneCpuQuota struct {

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例计费模式 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费 | CDHPAID：表示CDH付费，即只对CDH计费，不对CDH上的实例计费。 )
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 可用CPU配额。
	cpuQuota *uint64 `json:"cpuQuota,omitempty" name:"cpuQuota"`
}

type ZoneInfo struct {

	// 可用区名称，例如，ap-guangzhou-3
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区描述，例如，广州三区
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区状态
	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}
