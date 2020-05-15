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

package v20180813

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type ActionTimer struct {

	// 扩展数据
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// 定时器
	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`

	// 执行时间
	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API简介的ids.N一节）。参数不支持同时指定`DiskIds`和`Filters`。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds" list`

	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：`disk-11112222`。<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区]过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// API简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的云硬盘数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 云硬盘的详细信息列表。
		DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlavorsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`flavor-11112222`,每次请求的实例的上限为100。参数不支持同时指定`FlavorIds`和`Filters`。
	FlavorIds []*string `json:"FlavorIds,omitempty" name:"FlavorIds" list`

	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> flavor-id - String - 是否必填：否 - （过滤条件）按照实例ID过滤。实例ID形如：flavor-11112222。</li>
	// <li> flavor-name - String - 是否必填：否 - （过滤条件）按照实例名称过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`FlavorIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> flavor-id - String - 是否必填：否 - （过滤条件）按照实例ID过滤。实例ID形如：flavor-11112222。</li>
	// <li> flavor-name - String - 是否必填：否 - （过滤条件）按照实例名称过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`FlavorIds`和`Filters`。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFlavorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlavorsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlavorsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 套餐详情列表
		FlavorSet []*Flavor `json:"FlavorSet,omitempty" name:"FlavorSet" list`

		// 套餐总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlavorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlavorsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`bms-11112222`。（此参数的具体格式可参考API简介的`id.N`一节）。每次请求的实例的上限为100。参数不支持同时指定`InstanceIds`和`Filters`。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> uuid- String - 是否必填：否 - （过滤条件）按照实例UUID过滤。</li>
	// <li> instance-id - String - 是否必填：否 - （过滤条件）按照实例ID过滤。实例ID形如：ins-11112222。</li>
	// <li> instance-name - String - 是否必填：否 - （过滤条件）按照实例名称过滤。</li>
	// <li> instance-state - String - 是否必填：否 -（过滤条件）按照实例的状态过滤 </li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照实例主网卡的内网IP过滤。</li>
	// <li> vpc-id - String - 是否必填：否 - （过滤条件）按照实例所属的vpcid进行过滤 </li>
	// <li> subnet-id - String - 是否必填：否 - （过滤条件）按照实例所属的子网id进行过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

		// 实例总数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例的详细信息列表
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

type Disk struct {

	// 硬盘ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 硬盘所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 硬盘名称。
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// 硬盘大小。
	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`

	// 硬盘挂载的云主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 硬盘的创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认
	SecurityService *bool `json:"SecurityService,omitempty" name:"SecurityService"`

	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *bool `json:"MonitorService,omitempty" name:"MonitorService"`
}

type Externals struct {
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Flavor struct {

	// 套餐ID。
	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`

	// 套餐名称。
	FlavorName *string `json:"FlavorName,omitempty" name:"FlavorName"`

	// 实例所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 支持的Raid类型。
	RaidType []*string `json:"RaidType,omitempty" name:"RaidType" list`

	// 支持的系统列表。
	OperatingSystem *OperatingSystem `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// cpu信息。
	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`

	// 内存信息。
	Memory *string `json:"Memory,omitempty" name:"Memory"`

	// 硬盘信息。
	SystemDisk *string `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 网卡信息
	NetSpeed *string `json:"NetSpeed,omitempty" name:"NetSpeed"`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type Instance struct {

	// 实例所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 实例`ID`。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Raid类型。
	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`

	// 操作系统类型。
	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`

	// 操作系统发行版本
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 实例主网卡的内网`IP`列表。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// 实例所属虚拟私有网络信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 套餐信息。
	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 实例状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type InternetAccessible struct {

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances` API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例名称。可任意命名，但不得超过60个字符。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
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

type OperatingSystem struct {

	// 支持的linux系统列表
	Linux []*string `json:"Linux,omitempty" name:"Linux" list`

	// 支持的windows系统列表
	Windows []*string `json:"Windows,omitempty" name:"Windows" list`
}

type Placement struct {

	// 实例所属的可用区ID。该参数也可以通过调用  DescribeZones 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过`DescribeInstances`接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
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

		// 返回任务的taskid,可通过taskid查询任务运行情况
		TaskId []*string `json:"TaskId,omitempty" name:"TaskId" list`

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

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID。可通过 [DescribeInstances]
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// raid类型
	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`

	// 操作系统类型
	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`

	// 操作系统发行版本
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 实例登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则关闭开启云监控、云安全服务。
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

		// 返回任务的taskid,可通过taskid查询任务运行情况
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

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

type RunInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，专用宿主机（对于独享母机付费模式的子机创建）等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 套餐id
	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`

	// 系统类型
	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`

	// 系统的发行版本号
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络ip，表示每个实例的主网卡ip，而且InstanceCount参数必须与私有网络ip的个数一致。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// raid类型
	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。<br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux 等）实例：字符长度为[2, 30]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 实例数量。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526
	// )和Linux启动时运行命令。
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云主机实例。
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
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

		// 返回改任务的taskid,可通过taskid查询任务运行情况
		TaskId []*string `json:"TaskId,omitempty" name:"TaskId" list`

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

		// 返回任务的taskid,可通过taskid查询任务运行情况
		TaskId []*string `json:"TaskId,omitempty" name:"TaskId" list`

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

		// 返回任务的taskid,可通过taskid查询任务运行情况
		TaskId []*string `json:"TaskId,omitempty" name:"TaskId" list`

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

type Tag struct {

	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagSpecification struct {

	// 标签绑定的资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 标签对列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
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

		// 返回任务的taskid,可通过taskid查询任务运行情况
		TaskId []*string `json:"TaskId,omitempty" name:"TaskId" list`

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

type VirtualPrivateCloud struct {

	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录控制台查询；也可以调用接口 DescribeVpcEx ，从接口返回中的`unVpcId`字段获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口  DescribeSubnetEx ，从接口返回中的`unSubnetId`字段获取。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`

	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`
}
