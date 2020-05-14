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

package v20190625

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ColdMigrateInstancesRequest struct {
	*tchttp.BaseRequest

	// 要迁移的一个或者多个实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 迁移超时时间
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 指定带宽
	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 是否强制光机
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// 指定要迁移的目标母鸡列表
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`

	// 是否保留ip
	IpReserved *bool `json:"IpReserved,omitempty" name:"IpReserved"`
}

func (r *ColdMigrateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ColdMigrateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ColdMigrateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ColdMigrateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateTenantInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例的uuid数组，如["5182c761-90d7-4cd9-bbf8-4c19c4896dc8"]
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

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

func (r *ColdMigrateTenantInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ColdMigrateTenantInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateTenantInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务数组
		TaskSet []*TenantTask `json:"TaskSet,omitempty" name:"TaskSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ColdMigrateTenantInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ColdMigrateTenantInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHostTypesRequest struct {
	*tchttp.BaseRequest

	// 宿主机类型名称，如："Y0-GI51-25G"
	HostTypeName *string `json:"HostTypeName,omitempty" name:"HostTypeName"`

	// 该宿主机类型可售卖的核数，单位个
	CPUTotal *uint64 `json:"CPUTotal,omitempty" name:"CPUTotal"`

	// 该宿主机类型可售卖的内存大小，单位GB
	MemTotal *uint64 `json:"MemTotal,omitempty" name:"MemTotal"`

	// 该宿主机类型可售卖的本地磁盘大小，单位GB
	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`

	// 该宿主机类型可装箱的子机类型，如["S3", "S2"]
	InstanceFamilySet []*string `json:"InstanceFamilySet,omitempty" name:"InstanceFamilySet" list`

	// 该宿主机定制化参数，磁盘直通场景等需要填写
	CheckInfo *string `json:"CheckInfo,omitempty" name:"CheckInfo"`

	// 该宿主机定制化参数，和QOS服务相关
	HostExtraSpec *string `json:"HostExtraSpec,omitempty" name:"HostExtraSpec"`
}

func (r *CreateHostTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHostTypesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHostTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总共导入的宿主机类型个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 成功的宿主机类型名列表
		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList" list`

		// 失败的宿主机类型名列表
		FailList []*string `json:"FailList,omitempty" name:"FailList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHostTypesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageRequest struct {
	*tchttp.BaseRequest

	// 要制作自定义镜像的uuid
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 自定义镜像的名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义镜像的备注说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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

		// 异步任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

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

type CreateInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 实例大类，比如“S”，“I”
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`

	// 实例族“S3”，“S2”
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// 实例名，如“S2.2XLARGE2”
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU数量，单位个
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 最大带宽大小，如1.5，单位MB/s
	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 是否售卖
	// 1 : 售卖
	// 0：暂停售卖
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateInstanceTypeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceTypeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 尝试添加的总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 失败的config的id列表，可以用来排错
		FailList []*string `json:"FailList,omitempty" name:"FailList" list`

		// 成功的config的id列表
		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceTypeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceTypeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePublicImageRequest struct {
	*tchttp.BaseRequest

	// 镜像描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义镜像名称
	SourceImageId *string `json:"SourceImageId,omitempty" name:"SourceImageId"`

	// 自定义镜像名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// OS类别
	OsClass *string `json:"OsClass,omitempty" name:"OsClass"`

	// OS类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// OS架构
	Arch *string `json:"Arch,omitempty" name:"Arch"`

	// 是否可见
	Visible *int64 `json:"Visible,omitempty" name:"Visible"`

	// 是否reset base
	ResetBase *string `json:"ResetBase,omitempty" name:"ResetBase"`

	// 是否检查入参
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 镜像默认登录账户
	DefaultUser *string `json:"DefaultUser,omitempty" name:"DefaultUser"`
}

func (r *CreatePublicImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePublicImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePublicImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePublicImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePublicImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateZoneInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 要导入的可用区的zoneid
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 机型大类
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`

	// 实例族
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// 实例类型，可用区
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 机型核数
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 机型内存大小，单位为GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 机型的最大带宽
	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 售罄策略：
	// 0：资源决定
	// 1：强制售罄
	SoldOutStrategy *int64 `json:"SoldOutStrategy,omitempty" name:"SoldOutStrategy"`

	// 是否可售卖：
	// 1：开放售卖
	// 0：停止售卖
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 本地盘类型
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList" list`
}

func (r *CreateZoneInstanceTypeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateZoneInstanceTypeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateZoneInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 尝试创建的配置个数
		TotalCountT *uint64 `json:"TotalCountT,omitempty" name:"TotalCountT"`

		// 创建失败的配置ID列表
		FailList []*string `json:"FailList,omitempty" name:"FailList" list`

		// 创建成功的配置ID列表
		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateZoneInstanceTypeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateZoneInstanceTypeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 数据盘大小
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 数据盘ID
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 数据盘的序列号
	DiskSerial *string `json:"DiskSerial,omitempty" name:"DiskSerial"`
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest

	// 要操作的镜像数组
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

type DescribeFlowLogsRequest struct {
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

func (r *DescribeFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志流水信息。
		FlowLogSet []*FlowLog `json:"FlowLogSet,omitempty" name:"FlowLogSet" list`

		// 日志总条目数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostTypesRequest struct {
	*tchttp.BaseRequest

	// 可支持过滤key: 
	// host-type-list：按照HostTypeName进行过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeHostTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHostTypesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 宿主机详细信息列表
		HostTypeSet []*HostType `json:"HostTypeSet,omitempty" name:"HostTypeSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHostTypesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest

	// zone-id-list:要查的可用区的id，
	// net-version-list: 过滤要查询的宿主机的网络版本，
	// network-device-id-list：要过滤的母机的网络设备版本
	// host-asset-list: 固资号
	// host-ip-list: 母机ip
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 每页的数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始查找位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

		// 宿主机资源信息
		HostItem []*HostItem `json:"HostItem,omitempty" name:"HostItem" list`

		// 分页所用信息
		Total *string `json:"Total,omitempty" name:"Total"`

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

type DescribeImageRecordRequest struct {
	*tchttp.BaseRequest

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 每一页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选任务创建时间的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 筛选任务创建时间的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序类型
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeImageRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合条件的记录列表
		ImageRecordSet []*RecordList `json:"ImageRecordSet,omitempty" name:"ImageRecordSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest

	// image-uuid-list: 传入镜像的uuid列表，获取镜像的详细信息
	// image-id-list: 传入镜像的短id列表
	// appid: 传入镜像所有者的appId
	// exclude-appid-list: 不显示列表中的appid拥有的镜像
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// PUBLIC_IMAGE - 公共镜像
	// PRIVATE_IMAGE -私有镜像
	// 默认为公共镜像
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 开始查找的位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

		// 总共的镜像个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 查询的镜像的信息详情
		ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet" list`

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

type DescribeInstanceClassConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceClassConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceClassConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceClassConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 所有的实例系列，如“S”，“D”，“I”
		InstanceClassConfigSet []*InstanceClassConfig `json:"InstanceClassConfigSet,omitempty" name:"InstanceClassConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceClassConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceClassConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 接受三个数据参数：
	// instance-class-list,
	// instance-family-list,
	// instance-type-list
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

		// 全局机器类型详细列表
		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet" list`

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

type DescribeInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest

	// 获取的可用区ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例族名
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
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

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，详见下表：实例过滤条件表。
	// instance-name: 实例名字筛选（模糊查找）
	// host-ip-list: 根据母机ip筛选
	// network-device-id-list：根据上联网络设备号筛选
	// status-list：根据子机状态列表筛选
	// zone-id-list： 根据可用区筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，常常和Limit结合
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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

		// 符合条件的实例数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例相信信息列表
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

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 要查询的过滤器：
	// uuid-list：要查询的密钥对所属的子机列表（先在只支持填一个）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
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

		// 密钥对信息
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

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 要查询的任务的id数组
	TaskIds []*int64 `json:"TaskIds,omitempty" name:"TaskIds" list`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务详细信息
		TaskSet []*Task `json:"TaskSet,omitempty" name:"TaskSet" list`

		// 任务数量
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例的uuid数组，如["5182c761-90d7-4cd9-bbf8-4c19c4896dc8"]
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeTenantInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTenantInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTenantInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVncRequest struct {
	*tchttp.BaseRequest

	// 要查看的VNC的实例的uuid
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeVncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVncRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVncResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// VNC ip地址
		VNCIP *string `json:"VNCIP,omitempty" name:"VNCIP"`

		// VNC 端口号
		VNCPort *string `json:"VNCPort,omitempty" name:"VNCPort"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVncResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVsChildTaskIdByParentRequest struct {
	*tchttp.BaseRequest

	// CVM_VS 父任务ID列表
	VSParentTaskIds []*string `json:"VSParentTaskIds,omitempty" name:"VSParentTaskIds" list`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVsChildTaskIdByParentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVsChildTaskIdByParentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVsChildTaskIdByParentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 父任务到子任务的映射列表。
		ParentToChildSet []*ParentToChild `json:"ParentToChildSet,omitempty" name:"ParentToChildSet" list`

		// 总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVsChildTaskIdByParentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVsChildTaskIdByParentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVsChildTasksRequest struct {
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

	// 是否包含MQ详情，默认为False。
	IncludeDetails *bool `json:"IncludeDetails,omitempty" name:"IncludeDetails"`
}

func (r *DescribeVsChildTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVsChildTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVsChildTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志详情。
		TaskInfoSet []*VsChildTask `json:"TaskInfoSet,omitempty" name:"TaskInfoSet" list`

		// 日志总量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVsChildTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVsChildTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneImportInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// zone-id:可用区ID
	// instance-family-list: 机型族列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeZoneImportInstanceTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneImportInstanceTypeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneImportInstanceTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可导入的子机详细信息
		InstanceTypeConfigImportSet []*InstanceTypeConfigImport `json:"InstanceTypeConfigImportSet,omitempty" name:"InstanceTypeConfigImportSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneImportInstanceTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneImportInstanceTypeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// zone-id - Integer - 是否必填：否 -（过滤条件）按照可用区过滤。
	// 
	// instance-family-list  是否必填：否 -（过滤条件）按照机型系列过滤。按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等
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

		// 可用区内机型配置详细列表
		InstanceTypeQuotaSet []*InstanceTypeQuota `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet" list`

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

type FailTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 失败信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

func (r *FailTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FailTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FailTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FailTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FailTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FailoverMigrateInstancesRequest struct {
	*tchttp.BaseRequest

	// 要进行迁移的虚拟机
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 迁移到指定的母机ip列表
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`

	// 如果虚拟机是本地盘，这个字段必传。1表示数据允许丢失
	Lossy *uint64 `json:"Lossy,omitempty" name:"Lossy"`

	// 如果传了lossy字段并且为1，那么该字段必传。
	// 用该镜像ID来创建虚拟机
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *FailoverMigrateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FailoverMigrateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FailoverMigrateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FailoverMigrateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FailoverMigrateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type FlowLog struct {

	// 执行该操作的启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 执行该操作的结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 执行该请求的来源方
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`

	// 执行的操作方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 操作者AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateAppId *string `json:"OperateAppId,omitempty" name:"OperateAppId"`

	// 操作者Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 操作者SubAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateSubAccountUin *string `json:"OperateSubAccountUin,omitempty" name:"OperateSubAccountUin"`

	// 请求的唯一RequestId
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 调用周边组件的方法名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterfaceName *string `json:"InterfaceName,omitempty" name:"InterfaceName"`

	// 具体执行的组件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Component *string `json:"Component,omitempty" name:"Component"`

	// 日志类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 调用方
	// 注意：此字段可能返回 null，表示取不到有效值。
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被调用方
	// 注意：此字段可能返回 null，表示取不到有效值。
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 调用的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// des 组件任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DESTaskId *string `json:"DESTaskId,omitempty" name:"DESTaskId"`

	// vs组件父任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VSParentTaskId *string `json:"VSParentTaskId,omitempty" name:"VSParentTaskId"`

	// 请求的入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputData *string `json:"InputData,omitempty" name:"InputData"`

	// 请求的返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputData *string `json:"OutputData,omitempty" name:"OutputData"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type HostItem struct {

	// 宿主机内网IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 宿主机资源信息
	HostResource *HostResource `json:"HostResource,omitempty" name:"HostResource"`

	// 宿主机标记
	HostTag *string `json:"HostTag,omitempty" name:"HostTag"`

	// 宿主机固资号
	HostAsset *string `json:"HostAsset,omitempty" name:"HostAsset"`
}

type HostResource struct {

	// 宿主机总cpu核数
	CpuTotal *uint64 `json:"CpuTotal,omitempty" name:"CpuTotal"`

	// 宿主机可用cpu核数
	CpuAvailable *uint64 `json:"CpuAvailable,omitempty" name:"CpuAvailable"`

	// 宿主机总内存大小
	MemTotal *float64 `json:"MemTotal,omitempty" name:"MemTotal"`

	// 宿主机可用内存大小
	MemAvailable *float64 `json:"MemAvailable,omitempty" name:"MemAvailable"`

	// 宿主机总硬盘大小
	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`

	// 宿主机可用硬盘大小
	DiskAvailable *uint64 `json:"DiskAvailable,omitempty" name:"DiskAvailable"`
}

type HostType struct {

	// 类型名称
	HostTypeName *string `json:"HostTypeName,omitempty" name:"HostTypeName"`

	// 该宿主机可支持的虚拟子机类型
	InstanceFamilySet []*string `json:"InstanceFamilySet,omitempty" name:"InstanceFamilySet" list`

	// 可售卖CPU核数，单位个
	CPUTotal *uint64 `json:"CPUTotal,omitempty" name:"CPUTotal"`

	// 可售卖内存大小，单位GB
	MemTotal *uint64 `json:"MemTotal,omitempty" name:"MemTotal"`

	// 可售卖磁盘大小，单位GB
	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`

	// 宿主机定制化参数，为json格式字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostExtraSpec *string `json:"HostExtraSpec,omitempty" name:"HostExtraSpec"`

	// 宿主机定制化参数，磁盘直通功能相关，为json格式字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckInfo *string `json:"CheckInfo,omitempty" name:"CheckInfo"`
}

type Image struct {

	// 镜像id如：img-qy9cl0e3
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像操作系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 镜像大小
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 镜像操作系统
	OsKey *string `json:"OsKey,omitempty" name:"OsKey"`

	// PUBLIC_IMAGE: 公共镜像 PRIVATE_IMAGE:私有镜像
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 镜像创建时间
	CreatedTime []*string `json:"CreatedTime,omitempty" name:"CreatedTime" list`

	// 镜像当前状态
	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`

	// EXTERNAL_IMPORT, CREATE_IMAGE, OFFICIAL
	ImageSource *string `json:"ImageSource,omitempty" name:"ImageSource"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 操作系统平台
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 操作系统架构
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 镜像全局id如：img2018053006451178
	ImageUuid *string `json:"ImageUuid,omitempty" name:"ImageUuid"`
}

type Instance struct {

	// CPU核数
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例名称(alias)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例运行状态(realStatus)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 内网IP地址
	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 所在母机Ip地址
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 模块Id
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 上联网络设备Id(netdev_asset_id)
	NetworkDeviceId *string `json:"NetworkDeviceId,omitempty" name:"NetworkDeviceId"`

	// 对应asset
	Asset *string `json:"Asset,omitempty" name:"Asset"`

	// 可用区Id
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 操作系统名字
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 原uuid
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 拥有者名字
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 完成时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// vpcId
	VirtualPrivateCloudId *int64 `json:"VirtualPrivateCloudId,omitempty" name:"VirtualPrivateCloudId"`

	// 设备Id
	DeviceId *int64 `json:"DeviceId,omitempty" name:"DeviceId"`

	// 数据卷
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 系统卷
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
}

type InstanceClassConfig struct {

	// 实例大类
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`

	// 实例大类显示名称
	InstanceClassName *string `json:"InstanceClassName,omitempty" name:"InstanceClassName"`

	// 优先级
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 实例族信息
	InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet,omitempty" name:"InstanceFamilyConfigSet" list`
}

type InstanceFamilyConfig struct {

	// 实例族名称
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// 优先级
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 显示名称
	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`

	// 这个实例族最大可配置CPU数量（单位个）
	CPUTotal *int64 `json:"CPUTotal,omitempty" name:"CPUTotal"`

	// 这个实例族最大可配置内存大小（单位G）
	MemTotal *int64 `json:"MemTotal,omitempty" name:"MemTotal"`

	// 网卡数量（单位个）
	NetworkCardTotal *int64 `json:"NetworkCardTotal,omitempty" name:"NetworkCardTotal"`
}

type InstanceTypeConfig struct {

	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU个数
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// GPU核数
	GPU *uint64 `json:"GPU,omitempty" name:"GPU"`

	// FPGA核数
	FPGA *uint64 `json:"FPGA,omitempty" name:"FPGA"`

	// 存储块数
	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`

	// 网卡数
	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`

	// 最大带宽
	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 配置Id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 售卖状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceTypeConfigImport struct {

	// 机型类型名称，如“S3.SMALL!”
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU个数
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 售卖状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 评论
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 最大网络带宽G/s
	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 本地盘类型
	DiskType *int64 `json:"DiskType,omitempty" name:"DiskType"`
}

type InstanceTypeQuota struct {

	// 机型类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU大小
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// GPU大小
	GPU *uint64 `json:"GPU,omitempty" name:"GPU"`

	// FPGA核数
	FPGA *uint64 `json:"FPGA,omitempty" name:"FPGA"`

	// 存储快熟
	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`

	// 网卡数
	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`

	// 最大带宽
	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 配置Id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 售卖状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 售罄规则。0：按资源决定，1：强制售罄
	SoldOutStrategy *int64 `json:"SoldOutStrategy,omitempty" name:"SoldOutStrategy"`

	// 本地盘信息
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList" list`
}

type KeyPair struct {

	// 唯一id
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 密钥名称
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
}

type LiveMigrateInstancesRequest struct {
	*tchttp.BaseRequest

	// 要做热迁移的实例id列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 指定迁移的母鸡
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`

	// 超时时间，单位为秒
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 虚拟机热迁移pause阶段最大暂停时间，单位毫秒。
	// VStation默认100毫秒，建议30ms~30000ms
	PauseDuration *uint64 `json:"PauseDuration,omitempty" name:"PauseDuration"`

	// 指定带宽，单位为MiB/s
	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
}

func (r *LiveMigrateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveMigrateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveMigrateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveMigrateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateTenantInstancesRequest struct {
	*tchttp.BaseRequest

	// 待迁移实例uuid列表:["5182c761-90d7-4cd9-bbf8-4c19c4896dc8"]
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

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

func (r *LiveMigrateTenantInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveMigrateTenantInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateTenantInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务明细
		TaskSet []*TenantTask `json:"TaskSet,omitempty" name:"TaskSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveMigrateTenantInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveMigrateTenantInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LocalDiskType struct {

	// 本地磁盘类型：“LOCAL_BASIC”、“LOCAL_SSD”
	Type *string `json:"Type,omitempty" name:"Type"`

	// 本地磁盘属性。“ROOT”、“DATA”
	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`

	// 本地磁盘最小值
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// 本地磁盘最大值
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type LoginSettings struct {

	// 实例登录密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	Username *string `json:"Username,omitempty" name:"Username"`
}

type ModifyHostsTagRequest struct {
	*tchttp.BaseRequest

	// 要修改的母机ip列表
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`

	// 要增加的Tag:
	// 如HOST_FAULT
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 操作原因。当删除标签的时候，需要使用这个信息。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 本次的操作人。用于追溯信息使用。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 标签的描述。详细的描述一下标签的说明。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyHostsTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHostsTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHostsTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 1 - 售卖
	// 2 - 停止售卖
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 需要修改的机型配置名
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyInstanceTypeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceTypeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 尝试修改的配置个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 修改失败的配置ID列表
		FailList []*string `json:"FailList,omitempty" name:"FailList" list`

		// 修改成功的配置ID列表
		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceTypeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceTypeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 一个或者多个实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例显示名字
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

		// 总共操作的机器个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 成功的机器的uuid列表
		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList" list`

		// 失败的机器的uuid列表
		FailList []*string `json:"FailList,omitempty" name:"FailList" list`

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

type ModifyZoneInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 可用区的zone id
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 要修改的实例机型config id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 售卖状态：
	// 0：停止售卖
	// 1：开启售卖
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 售罄策略：
	// 0：按资源决定
	// 1：强制售罄
	SoldOutStrategy *int64 `json:"SoldOutStrategy,omitempty" name:"SoldOutStrategy"`

	// 本地盘类型
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList" list`
}

func (r *ModifyZoneInstanceTypeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyZoneInstanceTypeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 尝试修改的个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 修改失败的配置的ID列表
		FailList *string `json:"FailList,omitempty" name:"FailList"`

		// 修改成功的配置的ID列表
		SuccessList *string `json:"SuccessList,omitempty" name:"SuccessList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneInstanceTypeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyZoneInstanceTypeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperationRecords struct {

	// 操作，包括自定义镜像转公共镜像和公共镜像跨地域复制等操作
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 状态：已转为公共镜像，转换失败，转化中，已复制到其它地域，复制中，复制失败等
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作时间
	OperateTime *string `json:"OperateTime,omitempty" name:"OperateTime"`
}

type ParentToChild struct {

	// 父任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentTaskId *string `json:"ParentTaskId,omitempty" name:"ParentTaskId"`

	// 子任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildTaskId *string `json:"ChildTaskId,omitempty" name:"ChildTaskId"`
}

type Placement struct {

	// 实例所属可用区的ID。该参数也可以通过调用的返回值中的Zone字段来获取。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 所属项目Id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 宿主机ip列表
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 要中期的实例id
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 是否强制重启
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

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

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

type RecordList struct {

	// 原镜像ID
	SrcImageId *string `json:"SrcImageId,omitempty" name:"SrcImageId"`

	// 转化后的镜像ID
	DstImageId *string `json:"DstImageId,omitempty" name:"DstImageId"`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作记录
	OperationRecords []*OperationRecords `json:"OperationRecords,omitempty" name:"OperationRecords" list`

	// 具体操作
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 状态
	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`

	// 目的地域
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// 操作时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 错误码
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 操作记录
	Para *string `json:"Para,omitempty" name:"Para"`
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 虚拟机新的镜像ID
	ImageUuid *string `json:"ImageUuid,omitempty" name:"ImageUuid"`

	// 重装虚拟机的ROOT盘lv的大小。
	ResetSize *uint64 `json:"ResetSize,omitempty" name:"ResetSize"`

	// 重装系统后的登录信息
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
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

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

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

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest

	// 一个或多个实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例登录密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 待重置密码的实例操作系统用户名。不得超过64个字符。
	Username *string `json:"Username,omitempty" name:"Username"`
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

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

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

type RunInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，专用宿主机（对于独享母机付费模式的子机创建）等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 有效镜像Id
	ImageUuid *string `json:"ImageUuid,omitempty" name:"ImageUuid"`

	// 系统盘配置信息
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// CPU核数
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小，单位为MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 数据盘配置信息
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 购买数目
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 实例显示名字
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例登录设置
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 母鸡类型，如“M10”
	HostType *string `json:"HostType,omitempty" name:"HostType"`
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

		// 返回当前创建任务的task_id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

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

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

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

	// 关闭一个或者多个实例
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

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

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

type SyncHostTypesRequest struct {
	*tchttp.BaseRequest

	// 需要同步过来的地域名称，如“ap-guangzhou”
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *SyncHostTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncHostTypesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncHostTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总共同步的宿主机类型个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 同步成功的宿主机类型列表
		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList" list`

		// 同步失败的宿主机类型列表
		FailList []*string `json:"FailList,omitempty" name:"FailList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncHostTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncHostTypesResponse) FromJsonString(s string) error {
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

type SyncInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest

	// 需要同步到的可用区
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 需要同步的实例类型
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

func (r *SyncInstanceFamilyConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncInstanceFamilyConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 尝试同步的配置个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 同步失败的配置ID列表
		FailList []*string `json:"FailList,omitempty" name:"FailList" list`

		// 同步成功的配置ID列表
		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncInstanceFamilyConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncInstanceFamilyConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {

	// 系统盘大小
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 系统盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘Id
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 系统盘序列号
	DiskSerial *string `json:"DiskSerial,omitempty" name:"DiskSerial"`

	// 系统盘的镜像ID
	ImageUuid *string `json:"ImageUuid,omitempty" name:"ImageUuid"`
}

type Task struct {

	// 任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务详细信息
	TaskInfoSet []*TaskInfo `json:"TaskInfoSet,omitempty" name:"TaskInfoSet" list`
}

type TaskInfo struct {

	// 任务操作的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务状态: WAITING: 任务等待中。RUNNING：任务正在执行中。SUCCESS：任务已经完成。FAIL：任务失败。UNKNOWN：未知
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type TenantTask struct {

	// 租户端任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 租户端资源操作的Instance id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务信息（报错）
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 是否强制退还
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
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

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

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

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级或者降级到的CPU核数
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 升级或降级的内存大小
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 升级或者降级磁盘大小
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VsChildTask struct {

	// 子任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 父任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentTaskId *string `json:"ParentTaskId,omitempty" name:"ParentTaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 虚拟机uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`

	// 虚拟机内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIp *string `json:"InstanceIp,omitempty" name:"InstanceIp"`

	// 虚拟机所在的宿主机ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 虚拟机使用的镜像id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 虚拟机使用的镜像osName
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 虚拟机所属的用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 日志详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogDetail *string `json:"LogDetail,omitempty" name:"LogDetail"`

	// 操作启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 操作更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}
