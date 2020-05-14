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

package v20190107

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActiveIsolatedInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ActiveIsolatedInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActiveIsolatedInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ActiveIsolatedInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程Id
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActiveIsolatedInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActiveIsolatedInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ActiveIsolatedPGInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ActiveIsolatedPGInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActiveIsolatedPGInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ActiveIsolatedPGInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程Id
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActiveIsolatedPGInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActiveIsolatedPGInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 引擎类型，当前只支持TbaseV2和TbaseV3
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 引擎版本
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// CN数量
	CNCount *int64 `json:"CNCount,omitempty" name:"CNCount"`

	// CN 节点规格代码
	CNSpecCode *string `json:"CNSpecCode,omitempty" name:"CNSpecCode"`

	// 单个CN磁盘大小，单位：GB
	CNStorage *int64 `json:"CNStorage,omitempty" name:"CNStorage"`

	// DN数量
	DNCount *int64 `json:"DNCount,omitempty" name:"DNCount"`

	// DN 节点规格代码
	DNSpecCode *string `json:"DNSpecCode,omitempty" name:"DNSpecCode"`

	// 单组DN磁盘大小，单位：GB
	DNStorage *int64 `json:"DNStorage,omitempty" name:"DNStorage"`

	// 字符集，根据DescribeAvailableCharset查询
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 网络类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 管理员账号的密码
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 创建数量，(0,10]，不传默认创建1个
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *CreateInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceHourRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 冻结流水Id。
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 实例Id列表。
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 流程Id。
		FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceHourResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePGInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 网络类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 资源池id
	PoolId *int64 `json:"PoolId,omitempty" name:"PoolId"`

	// 引擎类型，当前只支持PostgreSQL
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 引擎版本
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 磁盘大小
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 节点规格代码
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 字符集，根据DescribeAvailableCharset查询
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 创建数量，(0,10]
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 实例类型，当前只支持master类型
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 主节点可用区
	MasterNodeZone *string `json:"MasterNodeZone,omitempty" name:"MasterNodeZone"`

	// 管理员账号的密码
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 从节点可用区
	FollowNodesZones []*string `json:"FollowNodesZones,omitempty" name:"FollowNodesZones" list`
}

func (r *CreatePGInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePGInstanceHourRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePGInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 冻结流水Id
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 实例Id列表
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 流程Id
		FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePGInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePGInstanceHourResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePGReadOnlyVipRequest struct {
	*tchttp.BaseRequest

	// 唯一标识的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 唯一标识的vpc子网Id
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 实例唯一标识Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 选填，可指定vip用于只读
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *CreatePGReadOnlyVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePGReadOnlyVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePGReadOnlyVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步生成只读vip的任务id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePGReadOnlyVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePGReadOnlyVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePGReadOnlyVipRequest struct {
	*tchttp.BaseRequest

	// 唯一标识的vpcId
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 唯一标识的vpc子网Id
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 要删除的VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 唯一标识的实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeletePGReadOnlyVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePGReadOnlyVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePGReadOnlyVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除只读VIP异步流程的任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePGReadOnlyVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePGReadOnlyVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableCharsetRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableCharsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableCharsetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableCharsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 支持的字符集列表
		AvailableCharsets []*string `json:"AvailableCharsets,omitempty" name:"AvailableCharsets" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableCharsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableCharsetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableEngineVersionsRequest struct {
	*tchttp.BaseRequest

	// 表示产品类型 (tbase, tpostgres)
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeAvailableEngineVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableEngineVersionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableEngineVersionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 引擎版本数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可用的引擎版本列表
		Items []*EngineVersion `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableEngineVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableEngineVersionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableZoneConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可售卖地域配置详情
		Items []*RegionSellConf `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableZoneConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZonesRequest struct {
	*tchttp.BaseRequest

	// 资源池名称
	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`

	// 资源池Id
	PoolId *int64 `json:"PoolId,omitempty" name:"PoolId"`
}

func (r *DescribeAvailableZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源池下的机器所在可用区
		Zones []*string `json:"Zones,omitempty" name:"Zones" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例名称
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// 管理员账号名
		AdminUserName *string `json:"AdminUserName,omitempty" name:"AdminUserName"`

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 状态描述
		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

		// 地域ID
		Region *string `json:"Region,omitempty" name:"Region"`

		// 地域名称
		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

		// 可用区ID
		Zone *string `json:"Zone,omitempty" name:"Zone"`

		// 可用区名称
		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

		// 私有网络Id
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 私有网络名称
		VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

		// 私有网络子网ID
		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

		// 私有网络子网名称
		SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

		// 创建时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 任务类型
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 任务描述
		TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`

		// 虚拟IP
		Vip *string `json:"Vip,omitempty" name:"Vip"`

		// 虚拟端口
		Vport *int64 `json:"Vport,omitempty" name:"Vport"`

		// 付费类型，postpaid - 后付费，prepaid - 预付费
		TradeType *string `json:"TradeType,omitempty" name:"TradeType"`

		// GTM节点列表
		GTMNodes []*InstanceNode `json:"GTMNodes,omitempty" name:"GTMNodes" list`

		// 计算节点列表
		CNNodes []*InstanceNode `json:"CNNodes,omitempty" name:"CNNodes" list`

		// 数据节点列表
		DNNodes []*InstanceNode `json:"DNNodes,omitempty" name:"DNNodes" list`

		// 备份及日志磁盘空间
		BackupStorage *int64 `json:"BackupStorage,omitempty" name:"BackupStorage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 流程任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeInstanceTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务的执行状态，有“success”，“running”，“failure”三种状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 按一个或多个实例Id进行查询，每次最多请求100个。指定该参数后，Filters参数失效。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 过滤条件。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 排序字段，支持 InstanceId, InstanceName, CreatedAt。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，支持ASC和DESC，默认为ASC。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 偏移量，默认为0。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数量，默认为20，最大值为100。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 实例状态信息，支持 creating, online, isolate
	Status *string `json:"Status,omitempty" name:"Status"`
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

		// 实例列表。
		Items []*Instance `json:"Items,omitempty" name:"Items" list`

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

type DescribePGAvailableZoneConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePGAvailableZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePGAvailableZoneConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePGAvailableZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可售卖地域配置详情
		Items []*PGRegionSellConf `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePGAvailableZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePGAvailableZoneConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePGInstanceDetailsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePGInstanceDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePGInstanceDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePGInstanceDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例名
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// 实例状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 实例状态描述
		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

		// 任务类型
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 任务描述
		TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`

		// 付费类型, postpaid-后付费, prepaid-预付费
		TradeType *string `json:"TradeType,omitempty" name:"TradeType"`

		// 地域信息
		Region *string `json:"Region,omitempty" name:"Region"`

		// 地域名称
		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

		// 读写vpc信息
		MainVpc *VpcInfo `json:"MainVpc,omitempty" name:"MainVpc"`

		// 只读vpc信息
		RoVpc *VpcInfo `json:"RoVpc,omitempty" name:"RoVpc"`

		// 字符集
		Charset *string `json:"Charset,omitempty" name:"Charset"`

		// 引擎版本
		EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

		// cpu核数
		Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

		// 内存大小
		Memory *int64 `json:"Memory,omitempty" name:"Memory"`

		// 磁盘大小
		Storage *int64 `json:"Storage,omitempty" name:"Storage"`

		// 节点信息
		Nodes []*Node `json:"Nodes,omitempty" name:"Nodes" list`

		// 实例创建时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePGInstanceDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePGInstanceDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePGInstancesRequest struct {
	*tchttp.BaseRequest

	// 按一个或多个实例Id进行查询，每次最多请求100个。指定该参数后，Filters参数失效。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 过滤条件。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 排序字段，支持 InstanceId, InstanceName, CreatedAt。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，支持ASC和DESC，默认为ASC。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 偏移量，默认为0。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数量，默认为20，最大值为100。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 实例状态信息，支持 creating, online, isolate。
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribePGInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePGInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePGInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Postgres实例列表。
		Items []*PGInstanceSet `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePGInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePGInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePoolsRequest struct {
	*tchttp.BaseRequest

	// 资源池名称，可用于模糊查询
	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`

	// 资源池Id，可用于精准查询
	PoolId *int64 `json:"PoolId,omitempty" name:"PoolId"`
}

func (r *DescribeResourcePoolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourcePoolsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePoolsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcePoolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourcePoolsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EngineVersion struct {

	// 引擎类型，例如 TBaseV2,TBaseV3
	Type *string `json:"Type,omitempty" name:"Type"`

	// 引擎版本，例如：1.0.1
	Version *string `json:"Version,omitempty" name:"Version"`
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Instance struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 地域ID
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 私有网络子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务描述
	TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`

	// 虚拟IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 虚拟端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 付费类型，postpaid - 后付费，prepaid - 预付费
	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`

	// 计算节点列表
	CNNodes []*InstanceNode `json:"CNNodes,omitempty" name:"CNNodes" list`

	// 数据节点列表
	DNNodes []*InstanceNode `json:"DNNodes,omitempty" name:"DNNodes" list`
}

type InstanceNode struct {

	// 节点类型
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点cpu
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 节点内存
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 节点数据盘大小
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 节点唯一ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 已用数据存储空间
	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程Id
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolatePGInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolatePGInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolatePGInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolatePGInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程Id
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolatePGInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolatePGInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 新的实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Node struct {

	// 节点角色
	Role *string `json:"Role,omitempty" name:"Role"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区名
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type PGInstanceSet struct {

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 私有网络子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 实例状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务描述信息
	TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`

	// 实例ip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 磁盘大小
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 节点数量
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 付费类型，(后付费-postpaid，预付费-prepaid)
	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`

	// 管理员账号名
	AdminUserName *string `json:"AdminUserName,omitempty" name:"AdminUserName"`
}

type PGRegionSellConf struct {

	// 地域英文名
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域中文名
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 所属大区
	Area *string `json:"Area,omitempty" name:"Area"`

	// 是否为默认地域
	IsDefaultRegion *bool `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`

	// 可用区售卖配置
	AvailableZones []*PGZoneSellConf `json:"AvailableZones,omitempty" name:"AvailableZones" list`
}

type PGSellSpec struct {

	// 规格码，唯一标识一个规格，创建实例时传该字段用于标识一个规格
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 节点类型: pg_dn（postgres实例存储节点）
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// cpu核心数，单位：核
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位：GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 磁盘最小规格，单位：GB
	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 磁盘最大规格，单位：GB
	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// 磁盘调整步长，单位：GB
	StorageStep *int64 `json:"StorageStep,omitempty" name:"StorageStep"`

	// 磁盘调整步长，单位：GB
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
}

type PGZoneSellConf struct {

	// 可用区英文名
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区中文名
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 是否为当前地域的默认可用区
	IsDefaultZone *bool `json:"IsDefaultZone,omitempty" name:"IsDefaultZone"`

	// 可选节点规格
	AvailableSpec []*PGSellSpec `json:"AvailableSpec,omitempty" name:"AvailableSpec" list`
}

type RegionSellConf struct {

	// 地域英文名
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域中文名
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 所属大区
	Area *string `json:"Area,omitempty" name:"Area"`

	// 是否为默认地域
	IsDefaultRegion *bool `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`

	// 可用区售卖配置
	AvailableZones []*ZoneSellConf `json:"AvailableZones,omitempty" name:"AvailableZones" list`
}

type ReleaseIsolatedInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ReleaseIsolatedInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseIsolatedInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程Id
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseIsolatedInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseIsolatedInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库账号名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为8~32位。
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
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

type SellSpec struct {

	// 规格码，唯一标识一个规格，创建实例时传该字段用于标识一个规格
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 节点类型: cn（计算节点）, dn（存储节点）
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// cpu核心数，单位：核
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位：GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 磁盘最小规格，单位：GB
	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 磁盘最大规格，单位：GB
	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// 磁盘调整步长，单位：GB
	StorageStep *int64 `json:"StorageStep,omitempty" name:"StorageStep"`

	// 磁盘调整步长，单位：GB
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
}

type VpcInfo struct {

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type ZoneSellConf struct {

	// 可用区英文名
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区中文名
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 是否为当前地域的默认可用区
	IsDefaultZone *bool `json:"IsDefaultZone,omitempty" name:"IsDefaultZone"`

	// 可选节点规格
	AvailableSpec []*SellSpec `json:"AvailableSpec,omitempty" name:"AvailableSpec" list`
}
