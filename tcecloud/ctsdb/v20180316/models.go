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

package v20180316

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ChargeCommonGoodsDetailForSpec struct {

	// 规格码。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例版本。
	ArchVersion *string `json:"ArchVersion,omitempty" name:"ArchVersion"`

	// 节点个数。
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 磁盘大小。
	Disk *int64 `json:"Disk,omitempty" name:"Disk"`

	// 网络大小。
	NetworkSize *int64 `json:"NetworkSize,omitempty" name:"NetworkSize"`

	// 加个模型id。
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 内存大小。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
}

type ChargeCommonGoodsDetailForSpecNewConfig struct {

	// 规格码。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例版本。
	ArchVersion *string `json:"ArchVersion,omitempty" name:"ArchVersion"`

	// 节点个数。
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 磁盘大小。
	Disk *int64 `json:"Disk,omitempty" name:"Disk"`

	// 网络大小。
	NetworkSize *int64 `json:"NetworkSize,omitempty" name:"NetworkSize"`

	// 加个模型id。
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 内存大小。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 变配模式，横向扩容：ScaleOut，纵向扩容：ScaleUp。
	ModifyMode *int64 `json:"ModifyMode,omitempty" name:"ModifyMode"`

	// 期望执行时间。
	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`
}

type CreateInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 业务名称。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目id。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 规格码。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例版本。
	ArchVersion *string `json:"ArchVersion,omitempty" name:"ArchVersion"`

	// 节点数。
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 磁盘大小。
	Disk *int64 `json:"Disk,omitempty" name:"Disk"`

	// 网络大小，目前固定填100。
	NetworkSize *int64 `json:"NetworkSize,omitempty" name:"NetworkSize"`

	// 加个模型ID。
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 内存大小。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网络类型，固定填1：vpc网络。
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`

	// 私有网络id。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 购买的实例个数。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 时间单位，y：年，m：月，d：日
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间长度，与TimeUnit结合使用。
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 计费模式，0：按量计费。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
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

		// 冻结流水id。
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 创建的实例id列表。
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

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

type DescribeDBInstanceMetricInfoRequest struct {
	*tchttp.BaseRequest

	// 实例id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定的metric名。
	Metric *string `json:"Metric,omitempty" name:"Metric"`
}

func (r *DescribeDBInstanceMetricInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceMetricInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceMetricInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// metric的详细信息，以json格式的字符串返回。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceMetricInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceMetricInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceMetricListRequest struct {
	*tchttp.BaseRequest

	// 实例id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceMetricListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceMetricListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceMetricListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例metric信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *MetricsInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceMetricListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceMetricListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceMetricQueryRequest struct {
	*tchttp.BaseRequest

	// 实例id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 指定的metric名。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 时间区间。
	TimeRanges []*TimeRange `json:"TimeRanges,omitempty" name:"TimeRanges" list`

	// 指标列。
	Fields *string `json:"Fields,omitempty" name:"Fields"`

	// 过滤器。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 过滤器操作关系。and：与，or：或。
	FiltersOp *string `json:"FiltersOp,omitempty" name:"FiltersOp"`

	// 单一聚合。avg,count,sum,max,min。
	SingleAggr []*string `json:"SingleAggr,omitempty" name:"SingleAggr" list`

	// 百分位。
	Percentile []*float64 `json:"Percentile,omitempty" name:"Percentile" list`

	// 百分等级。
	PercentileRanks []*float64 `json:"PercentileRanks,omitempty" name:"PercentileRanks" list`

	// 基数标识。1开启，0关闭。
	CardinalityFlag *int64 `json:"CardinalityFlag,omitempty" name:"CardinalityFlag"`

	// 间隔。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 分组。
	Group *string `json:"Group,omitempty" name:"Group"`

	// 返回限制。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// timestamp字段名字
	TimestampName *string `json:"TimestampName,omitempty" name:"TimestampName"`
}

func (r *DescribeDBInstanceMetricQueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceMetricQueryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceMetricQueryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果信息。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceMetricQueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceMetricQueryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceNodesRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceNodesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点详细信息列表
		NodeSet []*Node `json:"NodeSet,omitempty" name:"NodeSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceNodesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 分页返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 可用区名称，如ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例状态列表，可选：creating,created,initializing,running,modifying,isolated,destroying,destroyed
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses" list`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例Vip列表
	Vips []*string `json:"Vips,omitempty" name:"Vips" list`

	// 实例名称列表
	Names []*string `json:"Names,omitempty" name:"Names" list`

	// 排序字段，可选对本接口返回的所有字段进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，可选：ASC,DESC，默认为ASC
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 付费模式，0表示按需计费/后付费，1表示预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例Vip，支持模糊查询
	VipLike *string `json:"VipLike,omitempty" name:"VipLike"`

	// 实例Name，支持模糊查询
	NameLike *string `json:"NameLike,omitempty" name:"NameLike"`

	// tagKey搜索列表
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitempty" name:"TagKeysForSearch" list`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例列表
		InstanceSet []*InstanceSet `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesForIoTRequest struct {
	*tchttp.BaseRequest

	// 实例状态，可选:-4:删除 -3:删除中 -2:隔离 1:流程中 2:运行中 3:未初始化 4:生产中
	InstanceStatus *int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 项目id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例描述
	Names []*string `json:"Names,omitempty" name:"Names" list`

	// vips
	Vips []*string `json:"Vips,omitempty" name:"Vips" list`

	// 分页每页记录数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页页数，从 1 开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，可选 instanceName,expiredTime,createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 可选 ASC DESC
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 查询实例IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeInstancesForIoTRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesForIoTRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesForIoTResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例列表
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesForIoTResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesForIoTResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleSpecRequest struct {
	*tchttp.BaseRequest

	// 可用区名称，如ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeSaleSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSaleSpecRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 售卖规格列表
		SaleSpecSet []*SaleSpecSet `json:"SaleSpecSet,omitempty" name:"SaleSpecSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSaleSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSaleSpecResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleZoneRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSaleZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSaleZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 售卖区域列表
		SaleZoneSet []*SaleZoneSet `json:"SaleZoneSet,omitempty" name:"SaleZoneSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSaleZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSaleZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 维度列。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 操作符。例：equ(==),lt(<),lte(<=),gt(>),gte(>=),in,not_in
	Op *string `json:"Op,omitempty" name:"Op"`

	// 值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type InitDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例内网端口，取值为：[1024,65535]
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 实例用户名，固定为：root
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例用户密码，长度为：[8,16]
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *InitDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// L5Cmdid
	L5Cmdid *int64 `json:"L5Cmdid,omitempty" name:"L5Cmdid"`

	// L5Modid
	L5Modid *int64 `json:"L5Modid,omitempty" name:"L5Modid"`

	// 实例创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例硬盘大小
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 实例cpu数量
	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 实例接入vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例接入port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type InstanceSet struct {

	// 实例ID，唯一标识一个CTSDB实例
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称，用户可修改
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例状态：creating,created,initializing,running,modifying,isolated,destroying,destroyed
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例所属应用ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 实例所属项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所在地域名称，如ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例所在可用区名称，如ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 网络类型：1 VPC网络
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`

	// 私有网络ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 内网IP地址
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 内网端口地址
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// L5ModeId
	L5ModId *int64 `json:"L5ModId,omitempty" name:"L5ModId"`

	// L5CmdId
	L5CmdId *int64 `json:"L5CmdId,omitempty" name:"L5CmdId"`

	// 是否被锁定：0 否，1是
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`

	// 创建时间，格式为 2006-01-02 15:04:05
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间，格式为 2006-01-02 15:04:05
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 开始时间，格式为 2006-01-02 15:04:05
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 过期时间，格式为 2006-01-02 15:04:05
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 隔离时间，格式为 2006-01-02 15:04:05
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// 销毁时间，格式为 2006-01-02 15:04:05
	DestroyTime *string `json:"DestroyTime,omitempty" name:"DestroyTime"`

	// 付费模式：0 后付费，1 预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 自动续费标记：0 默认状态(用户未设置，即初始状态)，1 自动续费，2 明确不自动续费(用户设置)
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 实例实现版本
	ArchVersion *string `json:"ArchVersion,omitempty" name:"ArchVersion"`

	// 磁盘大小，单位GB
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 节点个数
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// CPU核数
	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 内存大小，单位GB
	MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

	// 写入数每秒
	Wps *int64 `json:"Wps,omitempty" name:"Wps"`

	// 规格码
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// pid
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Vip6地址；空即没有Vip6地址
	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`

	// tag列表
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`
}

type MetricList struct {

	// 实例metric列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics" list`
}

type MetricsInfo struct {

	// 实例metric列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *MetricList `json:"Result,omitempty" name:"Result"`

	// 状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitempty" name:"Name"`
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

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstanceProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceUserPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例用户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例用户密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModifyDBInstanceUserPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceUserPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceUserPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceUserPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceUserPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGoodsDetail struct {

	// 实例原始规格配置。
	OldConfig []*ChargeCommonGoodsDetailForSpec `json:"OldConfig,omitempty" name:"OldConfig" list`

	// 实例希望变更后的规格配置。
	NewConfig *ChargeCommonGoodsDetailForSpecNewConfig `json:"NewConfig,omitempty" name:"NewConfig"`

	// 实例到期时间。
	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`

	// 资源实例个数。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
}

type ModifyInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 业务名称。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目id。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 商品变配详情。
	GoodsDetail *ModifyGoodsDetail `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

	// 计费模式，可选0：后付费。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 订单id。
	DealId *int64 `json:"DealId,omitempty" name:"DealId"`
}

func (r *ModifyInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceHourRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceHourResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Node struct {

	// 节点Ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 节点Port
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type RecycleDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RecycleDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecycleDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecycleDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecycleDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SaleSpecSet struct {

	// 地域名称，如ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区名称，如ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 规格码
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 实例实现版本
	ArchVersion *string `json:"ArchVersion,omitempty" name:"ArchVersion"`

	// 状态：0 无效，1 有效
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 更新时间，格式为 2006-01-02 15:04:05
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 售卖规格详情
	Spec *Spec `json:"Spec,omitempty" name:"Spec"`
}

type SaleZoneSet struct {

	// 可用区，如：ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 地域，如：ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区Id，如：100001
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区名称，如：广州一区
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 可用区状态。0未售卖，1售卖中，2售罄，3新区域
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type Spec struct {

	// 规格码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 机器类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// CPU核数
	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 内存大小，单位GB
	MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

	// 写入数每秒
	Wps *int64 `json:"Wps,omitempty" name:"Wps"`

	// 磁盘大小，单位GB，默认值
	DiskSizeDefault *int64 `json:"DiskSizeDefault,omitempty" name:"DiskSizeDefault"`

	// 磁盘大小，单位GB，下限值
	DiskSizeBegin *int64 `json:"DiskSizeBegin,omitempty" name:"DiskSizeBegin"`

	// 磁盘大小，单位GB，上限值
	DiskSizeEnd *int64 `json:"DiskSizeEnd,omitempty" name:"DiskSizeEnd"`

	// 网络大小，单位GB，下限值
	NetworkSizeBegin *int64 `json:"NetworkSizeBegin,omitempty" name:"NetworkSizeBegin"`

	// 网络大小，单位GB，上限值
	NetworkSizeEnd *int64 `json:"NetworkSizeEnd,omitempty" name:"NetworkSizeEnd"`

	// 节点个数，下限值
	NodeNumBegin *int64 `json:"NodeNumBegin,omitempty" name:"NodeNumBegin"`

	// 节点个数，上限值
	NodeNumEnd *int64 `json:"NodeNumEnd,omitempty" name:"NodeNumEnd"`

	// 预付费Pid
	PrePid *int64 `json:"PrePid,omitempty" name:"PrePid"`

	// 后付费Pid
	PostPid *int64 `json:"PostPid,omitempty" name:"PostPid"`

	// 适用信息
	SuitInfo *string `json:"SuitInfo,omitempty" name:"SuitInfo"`

	// 更新时间，格式为 2006-01-02 15:04:05
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TagInfo struct {

	// tag键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// tag值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TimeRange struct {

	// 开始时间。例：2016-01-01T02:53:57Z
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间。例：2016-01-01T03:53:57Z
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 时区设置。例：+08:00
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}
