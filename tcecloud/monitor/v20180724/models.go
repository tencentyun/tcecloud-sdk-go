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

package v20180724

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AlarmBindingInstance struct {

	// 实例UUID
	UUID *string `json:"UUID,omitempty" name:"UUID"`

	// 实例维度组合信息
	Information *string `json:"Information,omitempty" name:"Information"`
}

type AmpFrontTunnelRequest struct {
	*tchttp.BaseRequest

	// HTTP请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// HTTP请求地址
	Path *string `json:"Path,omitempty" name:"Path"`

	// HTTP请求体
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
}

func (r *AmpFrontTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AmpFrontTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AmpFrontTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// HTTP请求返回结果
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AmpFrontTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AmpFrontTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ArgusFrontTunnelRequest struct {
	*tchttp.BaseRequest

	// HTTP请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// HTTP请求地址
	Path *string `json:"Path,omitempty" name:"Path"`

	// HTTP请求体
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
}

func (r *ArgusFrontTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArgusFrontTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ArgusFrontTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// HTTP请求返回结果
		ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArgusFrontTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArgusFrontTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttributeAggrValueInfoOutputData struct {

	// 返回聚合数据列表
	Data []*AttributeTimestampValueOutput `json:"Data,omitempty" name:"Data" list`

	// 返回聚合数据列表个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AttributeIdDeleteOutput struct {

	// 删除个数
	DeleteCount *int64 `json:"DeleteCount,omitempty" name:"DeleteCount"`

	// 属性ID
	AttributeId []*uint64 `json:"AttributeId,omitempty" name:"AttributeId" list`
}

type AttributeIdOutput struct {

	// 属性ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

type AttributeInfoInput struct {

	// 属性名称
	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`

	// 属性数据类型
	DataType *int64 `json:"DataType,omitempty" name:"DataType"`

	// 属性唯一名称。由字母、数字、横杠或下划线组成
	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`

	// 属性类型
	AttributeTypeId *uint64 `json:"AttributeTypeId,omitempty" name:"AttributeTypeId"`

	// 属性级别
	AttributeLevel *int64 `json:"AttributeLevel,omitempty" name:"AttributeLevel"`

	// 属性单位
	Unit *int64 `json:"Unit,omitempty" name:"Unit"`

	// 统计周期
	StatisticalPeriod *int64 `json:"StatisticalPeriod,omitempty" name:"StatisticalPeriod"`

	// 负责人列表ID。默认为登录用户sub_uin
	OwnerId []*uint64 `json:"OwnerId,omitempty" name:"OwnerId" list`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type AttributeInfoOutput struct {

	// 属性ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 属性名称
	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`

	// 属性级别
	AttributeLevel *int64 `json:"AttributeLevel,omitempty" name:"AttributeLevel"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 属性类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeType *string `json:"AttributeType,omitempty" name:"AttributeType"`

	// 属性数据类型
	DataType *int64 `json:"DataType,omitempty" name:"DataType"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 属性负责人
	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`

	// 单位ID
	UnitId *int64 `json:"UnitId,omitempty" name:"UnitId"`

	// 单位名称
	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`

	// 统计周期
	StatisticalPeriod *int64 `json:"StatisticalPeriod,omitempty" name:"StatisticalPeriod"`

	// 属性唯一字符串，字母、数字、横杠或下划线组成
	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`
}

type AttributeInfoOutputData struct {

	// 属性数据列表
	Data []*AttributeInfoOutput `json:"Data,omitempty" name:"Data" list`

	// 属性数据列表个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AttributeServerInfoOutput struct {

	// 服务器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`

	// 服务器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 服务器ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`

	// 地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type AttributeServerInfoOutputData struct {

	// 返回上报该属性的服务器列表
	Data []*AttributeServerInfoOutput `json:"Data,omitempty" name:"Data" list`

	// 返回上报该属性的服务器列表个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AttributeTimestampValueOutput struct {

	// 时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 属性值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type AttributeUnitInfoOutput struct {

	// 指标单位ID
	UnitId *uint64 `json:"UnitId,omitempty" name:"UnitId"`

	// 指标单位名称
	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type AttributeUnitInfoOutputData struct {

	// 属性单位数据列表
	Data []*AttributeUnitInfoOutput `json:"Data,omitempty" name:"Data" list`

	// 属性单位数据列表个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AttributeValueInfoOutput struct {

	// 属性值列表
	Values []*AttributeTimestampValueOutput `json:"Values,omitempty" name:"Values" list`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始位置
	StartPosition *int64 `json:"StartPosition,omitempty" name:"StartPosition"`

	// 结束位置
	EndPosition *int64 `json:"EndPosition,omitempty" name:"EndPosition"`
}

type AttributeValueInfoOutputData struct {

	// 属性上报数据列表
	Data []*AttributeValueInfoOutput `json:"Data,omitempty" name:"Data" list`

	// 属性数据列表个数个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type BindingPolicyObjectDimension struct {

	// 地域名
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 维度信息
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 事件维度信息
	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type BindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 策略分组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 必填。固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 实例分组ID
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 需要绑定的对象维度信息
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

func (r *BindingPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindingPolicyObjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindingPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindingPolicyObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CCMChartData struct {

	// 返回数据列表
	Data []*CCMChartEntry `json:"Data,omitempty" name:"Data" list`

	// 查询数据总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type CCMChartEntry struct {

	// 监控图表ID
	ChartId *int64 `json:"ChartId,omitempty" name:"ChartId"`

	// 分组视图ID
	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`

	// 监控图表名称
	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`

	// 监控图表类型。1 明细视图， 2 聚合视图
	ChartType *int64 `json:"ChartType,omitempty" name:"ChartType"`

	// 产品类型。1 基础监控， 2 自定义监控
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 指标ID
	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 聚合方式。0 SUM， 1 AVG， 2 MAX， 3 MIN
	Aggregation *int64 `json:"Aggregation,omitempty" name:"Aggregation"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建者ID
	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新者ID
	UpdaterId *uint64 `json:"UpdaterId,omitempty" name:"UpdaterId"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 指标名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`

	// 指标唯一英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`

	// 指标单位名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type CCMGroupViewEntry struct {

	// 分组视图ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`

	// 分组视图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 分组视图是否自动绑定指标。0：否，1：是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`
}

type CCMInstanceAttributeDataOutput struct {

	// 指标ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 服务器ID
	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`

	// 指标上报数据
	Values []*AttributeTimestampValueOutput `json:"Values,omitempty" name:"Values" list`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type CLMDescribeMetricSetsRequest struct {
	*tchttp.BaseRequest

	// 指标集CID
	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`

	// 指标集名称（模糊搜索）
	Name *string `json:"Name,omitempty" name:"Name"`

	// 日志集名称（模糊搜索）
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志主题名称（模糊搜索）
	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`

	// 偏移量，默认为0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数量，默认为20，最大100
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *CLMDescribeMetricSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CLMDescribeMetricSetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CLMDescribeMetricSetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标集列表结果数据
		Data *PCLMDescribeMetricSetsData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CLMDescribeMetricSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CLMDescribeMetricSetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CgrpGroupNode struct {

	// 分组ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 组深度
	High *int64 `json:"High,omitempty" name:"High"`

	// 是否是叶子节点
	IsLeaf *bool `json:"IsLeaf,omitempty" name:"IsLeaf"`

	// 当前在树中的层级
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 树节点
	Name *string `json:"Name,omitempty" name:"Name"`

	// 父节点Id
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// SubUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`

	// 允许叶子节点存放实例类型
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`

	// Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type CgrpInstance struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 叶子分组节点
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 类型
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type CgrpInstanceGroupNode struct {

	// 分组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CgrpInstanceNode struct {

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例类型
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 叶子分组Id
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 叶子分组名称
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 自增Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 一级分组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	L1Id *int64 `json:"L1Id,omitempty" name:"L1Id"`

	// 二级分组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	L2Id *int64 `json:"L2Id,omitempty" name:"L2Id"`

	// 三级分组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	L3Id *int64 `json:"L3Id,omitempty" name:"L3Id"`

	// appId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// SubUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`

	// 分组数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Group []*CgrpInstanceGroupNode `json:"Group,omitempty" name:"Group" list`
}

type CgrpModuleNode struct {

	// 叶子分组Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 父分组Id
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// 叶子分组允许添加的实例类型
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`

	// AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 分组树高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	High *int64 `json:"High,omitempty" name:"High"`

	// 第一层分组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	L1Id *int64 `json:"L1Id,omitempty" name:"L1Id"`

	// 第一层分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	L1Name *string `json:"L1Name,omitempty" name:"L1Name"`

	// 第二层分组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	L2Id *int64 `json:"L2Id,omitempty" name:"L2Id"`

	// 第二层分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	L2Name *string `json:"L2Name,omitempty" name:"L2Name"`

	// 第三层分组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	L3Id *int64 `json:"L3Id,omitempty" name:"L3Id"`

	// 第三层分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	L3Name *string `json:"L3Name,omitempty" name:"L3Name"`

	// 叶子分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备份负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerBack *string `json:"OwnerBack,omitempty" name:"OwnerBack"`

	// 主要负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerMain *string `json:"OwnerMain,omitempty" name:"OwnerMain"`

	// SubUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`

	// Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type ClmAlertFilterRule struct {

	// 关联的告警策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertPolicyId *int64 `json:"AlertPolicyId,omitempty" name:"AlertPolicyId"`

	// 过滤的关键字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 操作符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operating *string `json:"Operating,omitempty" name:"Operating"`

	// 筛选值（多个值逗号分隔）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 过滤条件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ClmAlertPolicy struct {

	// 策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标集ID(无效参数)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`

	// 状态 1=已开启 2=未开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 告警接收组ID 逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertGroupId *string `json:"AlertGroupId,omitempty" name:"AlertGroupId"`

	// 告警接收渠道  "SMS", "EMAIL", "WECHAT", "CALL" 多个逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertChannel *string `json:"AlertChannel,omitempty" name:"AlertChannel"`

	// 通知时段开始时间（从00:00:00开始计算的秒数）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticePeriodBegin *int64 `json:"NoticePeriodBegin,omitempty" name:"NoticePeriodBegin"`

	// 通知时段结束时间（从00:00:00开始计算的秒数）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticePeriodEnd *int64 `json:"NoticePeriodEnd,omitempty" name:"NoticePeriodEnd"`

	// 告警触发条件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertTriggerRules []*ClmAlertTriggerRule `json:"AlertTriggerRules,omitempty" name:"AlertTriggerRules" list`

	// 触发条件之间的关系 1与 2或
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerRelation *int64 `json:"TriggerRelation,omitempty" name:"TriggerRelation"`

	// 回调url的scheme
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlScheme *string `json:"UrlScheme,omitempty" name:"UrlScheme"`

	// 回调url 不包含scheme部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 告警过滤条件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertFilterRules []*ClmAlertFilterRule `json:"AlertFilterRules,omitempty" name:"AlertFilterRules" list`

	// 告警策略ID （出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 创建时间 （出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间 （出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 告警接收人ID 逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertUserId *string `json:"AlertUserId,omitempty" name:"AlertUserId"`

	// 最近告警时间（出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestAlertTime *string `json:"LatestAlertTime,omitempty" name:"LatestAlertTime"`

	// 过滤条件之间的关系 1与 2或
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterRelation *int64 `json:"FilterRelation,omitempty" name:"FilterRelation"`

	// 指标集CID
	MetricSetCID *float64 `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
}

type ClmAlertTriggerRule struct {

	// 关联的告警策略主键id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertPolicyId *int64 `json:"AlertPolicyId,omitempty" name:"AlertPolicyId"`

	// 指标名称（英文）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标类型 1=普通指标 2=复合指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricType *int64 `json:"MetricType,omitempty" name:"MetricType"`

	// 操作符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operating *string `json:"Operating,omitempty" name:"Operating"`

	// 指标阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitempty" name:"Value"`

	// 持续周期个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinuousCycleCount *int64 `json:"ContinuousCycleCount,omitempty" name:"ContinuousCycleCount"`

	// 通知频率（通知间隔秒数）
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeFrequencySec *int64 `json:"NoticeFrequencySec,omitempty" name:"NoticeFrequencySec"`

	// 触发条件ID (出参)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 创建时间 （出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 复合指标表达式 （基础指标传空）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricFormula *string `json:"MetricFormula,omitempty" name:"MetricFormula"`
}

type ClmAnalysisFilter struct {

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 运算符 目前只支持 = 
	// （当value数组为多个值时为in关系）
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 值
	Value []*string `json:"Value,omitempty" name:"Value" list`
}

type ClmDescribeAlertPoliciesData struct {

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表元素
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertPolicies []*ClmAlertPolicy `json:"AlertPolicies,omitempty" name:"AlertPolicies" list`
}

type ClmDescribeMetricSetsData struct {

	// 指标集列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricSets []*ClmMetricSet `json:"MetricSets,omitempty" name:"MetricSets" list`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ClmDimension struct {

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 条目ID（只作出参，入参不填）
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 创建时间（只作出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 指标集ID（只作出参，入参不填）
	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
}

type ClmLogFilterRule struct {

	// 关系 1AND 2OR
	Relation *int64 `json:"Relation,omitempty" name:"Relation"`

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 过滤操作符: 数字型字段支持：gt lt ge le eq ne in 字符型字段支持：eq ne in contains
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 过滤操作值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 条目ID（只作出参，入参不填）
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 指标集ID（只作出参，入参不填）
	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`

	// 创建时间（只作出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ClmLogProfileItem struct {

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 字段类型，可选：string、long、double
	Type *string `json:"Type,omitempty" name:"Type"`

	// id（只做出参，入参不填）
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 关联指标集ID（只做出参，入参不填）
	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`

	// 创建时间（只做出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只做出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ClmMetricAnalysisBasicMetric struct {

	// 操作符
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 指标名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ClmMetricAnalysisCustomMetric struct {

	// 运算表达式
	Formula *string `json:"Formula,omitempty" name:"Formula"`

	// 复合指标名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ClmMetricBasicItem struct {

	// 字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitempty" name:"Field"`

	// 指标统计方法：sum count min max
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标描述中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 指标过滤规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticFilterRules []*ClmStatisticFilterRule `json:"StatisticFilterRules,omitempty" name:"StatisticFilterRules" list`

	// 条目ID（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 指标集ID（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`

	// 创建时间（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ClmMetricCustomItem struct {

	// 表达式
	Formula *string `json:"Formula,omitempty" name:"Formula"`

	// 指标名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标描述中文名
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 条目ID（只作出参，入参不填）
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 指标集ID（只作出参，入参不填）
	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`

	// 创建时间（只作出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ClmMetricSet struct {

	// CLS日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// CLS日志集名
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// CLS日志主题ID
	LogtopicId *string `json:"LogtopicId,omitempty" name:"LogtopicId"`

	// CLS日志主题名
	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`

	// 指标集名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 统计周期 可选：60、300
	StatisticCycle *int64 `json:"StatisticCycle,omitempty" name:"StatisticCycle"`

	// 时间字段
	TimeField *string `json:"TimeField,omitempty" name:"TimeField"`

	// 日志过滤规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogFilterRules []*ClmLogFilterRule `json:"LogFilterRules,omitempty" name:"LogFilterRules" list`

	// 基础指标规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricBasicItems []*ClmMetricBasicItem `json:"MetricBasicItems,omitempty" name:"MetricBasicItems" list`

	// 复合指标规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricCustomItems []*ClmMetricCustomItem `json:"MetricCustomItems,omitempty" name:"MetricCustomItems" list`

	// 维度规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions []*ClmDimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 日志数据结构描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogProfileItems []*ClmLogProfileItem `json:"LogProfileItems,omitempty" name:"LogProfileItems" list`

	// 条目ID（只作出参，入参不填）
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 创建时间（只作出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// APPID（只作出参，入参不填）
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// UIN（只作出参，入参不填）
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 地域（只作出参，入参不填）
	Region *string `json:"Region,omitempty" name:"Region"`
}

type ClmStatisticFilterRule struct {

	// 关系 1AND 2OR
	Relation *int64 `json:"Relation,omitempty" name:"Relation"`

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 过滤操作符: 数字型字段支持：gt lt ge le eq ne in 字符型字段支持：eq ne in contains
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 条目ID（只做出参，入参不填）
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 关联指标条目ID（只做出参，入参不填）
	MetricItemId *int64 `json:"MetricItemId,omitempty" name:"MetricItemId"`

	// 创建时间（只做出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只做出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 过滤操作值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ClonePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 被复制的策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 是否复制回调方法，0不复制，1复制
	CopyAlarmMethod *int64 `json:"CopyAlarmMethod,omitempty" name:"CopyAlarmMethod"`
}

func (r *ClonePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClonePolicyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClonePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 复制成功的策略组Id
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClonePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClonePolicyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Condition struct {

	// 告警通知频率
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 重复通知策略预定义（0 - 只告警一次， 1 - 指数告警，2 - 连接告警）
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 检测方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcType *string `json:"CalcType,omitempty" name:"CalcType"`

	// 检测值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// 持续时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinueTime *string `json:"ContinueTime,omitempty" name:"ContinueTime"`

	// 指标ID
	MetricID *int64 `json:"MetricID,omitempty" name:"MetricID"`

	// 指标展示名称（对外）
	MetricDisplayName *string `json:"MetricDisplayName,omitempty" name:"MetricDisplayName"`

	// 周期
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 指标单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type Contact struct {

	// uid
	Uid *int64 `json:"Uid,omitempty" name:"Uid"`

	// uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 接收人名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否可登录,0表示不可登录,1表示可登录
	CanLogin *int64 `json:"CanLogin,omitempty" name:"CanLogin"`

	// 手机号是否验证通过
	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 邮箱是否验证通过
	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`

	// 邮箱号
	Email *string `json:"Email,omitempty" name:"Email"`

	// 是否是接收负责人
	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`

	// ownerUid
	OwnerUid *int64 `json:"OwnerUid,omitempty" name:"OwnerUid"`
}

type CopyConditionsTemplateRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 模板策略组ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 所属UIN
	OwnerUIN *string `json:"OwnerUIN,omitempty" name:"OwnerUIN"`
}

func (r *CopyConditionsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyConditionsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyConditionsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyConditionsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyConditionsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 实例组Id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
}

func (r *CopyInstanceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyInstanceGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyInstanceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyInstanceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyInstanceGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlertPolicyRequest struct {
	*tchttp.BaseRequest

	// 告警策略JSON对象
	AlertPolicy *ClmAlertPolicy `json:"AlertPolicy,omitempty" name:"AlertPolicy"`
}

func (r *CreateAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlertPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略ID
		Data *int64 `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlertPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAttributesRequest struct {
	*tchttp.BaseRequest

	// 属性信息
	Attribute []*AttributeInfoInput `json:"Attribute,omitempty" name:"Attribute" list`
}

func (r *CreateAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCCMChartRequest struct {
	*tchttp.BaseRequest

	// 监控图表名称
	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`

	// 监控图表类型。1 明细视图， 2 聚合视图
	ChartType *int64 `json:"ChartType,omitempty" name:"ChartType"`

	// 产品类型。1 基础监控， 2 自定义监控
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 指标ID
	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 聚合方式。0 SUM， 1 AVG， 2 MAX， 3 MIN
	Aggregation *int64 `json:"Aggregation,omitempty" name:"Aggregation"`

	// 分组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateCCMChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCCMChartRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCCMChartResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCCMChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCCMChartResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConditionsTemplateRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 组名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 是否为与关系
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 父ID
	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`

	// 是否屏蔽
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// 指标告警条件
	Conditions []*ModifyConditionsTemplateRequestCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// 事件告警条件
	EventConditions []*ModifyConditionsTemplateRequestEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`
}

func (r *CreateConditionsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConditionsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConditionsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板策略组ID
		GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConditionsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConditionsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 监控面板名
	DescName *string `json:"DescName,omitempty" name:"DescName"`

	// json描述元数据
	Meta *string `json:"Meta,omitempty" name:"Meta"`

	// 产品类型
	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`

	// CustomID
	CustomID *string `json:"CustomID,omitempty" name:"CustomID"`
}

func (r *CreateDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDashboardRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控面板ID
		DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDashboardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardViewRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 监控面板视图名
	DescName *string `json:"DescName,omitempty" name:"DescName"`

	// 名字空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 监控面板ID
	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`

	// 面板实例列表
	Instances []*string `json:"Instances,omitempty" name:"Instances" list`

	// 元数据，json串
	Meta *string `json:"Meta,omitempty" name:"Meta"`

	// 指标名列表
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`
}

func (r *CreateDashboardViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDashboardViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视图ID
		ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashboardViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDashboardViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 父分组ID
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// 分组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回创建的分组节点
		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 实例组名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 实例列表
	InstanceList []*BindingPolicyObjectDimension `json:"InstanceList,omitempty" name:"InstanceList" list`
}

func (r *CreateInstanceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例列表
	Instances []*CgrpInstance `json:"Instances,omitempty" name:"Instances" list`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例列表
		Data []*CgrpInstanceNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMetricSetRequest struct {
	*tchttp.BaseRequest

	// 指标集配置
	MetricSet *ClmMetricSet `json:"MetricSet,omitempty" name:"MetricSet"`
}

func (r *CreateMetricSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMetricSetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMetricSetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标集ID
		Data *int64 `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMetricSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMetricSetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleRequest struct {
	*tchttp.BaseRequest

	// 父分组Id
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// 叶子分组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分组允许添加的实例类型
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *CreateModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 叶子分组节点
		Data []*CgrpModuleNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMsgPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 接收用户组id列表
	ReceiverGroupIds []*int64 `json:"ReceiverGroupIds,omitempty" name:"ReceiverGroupIds" list`

	// 告警方式列表
	NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays" list`

	// 电话告警配置
	VoiceConfig *ModifyMsgPolicyVoiceConfig `json:"VoiceConfig,omitempty" name:"VoiceConfig"`
}

func (r *CreateMsgPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMsgPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMsgPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的策略id
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMsgPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMsgPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupCondition struct {

	// 指标Id
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等。如果指标有配置默认比较类型值可以不填。
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// 比较的值，如果指标不必须CalcValue可不填
	CalcValue *float64 `json:"CalcValue,omitempty" name:"CalcValue"`

	// 数据聚合周期(单位秒)，若指标有默认值可不填
	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`

	// 持续几个检测周期触发规则会告警
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// 如果通过模版创建，需要传入模版中该指标的对应RuleId
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupEventCondition struct {

	// 告警事件的Id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 如果通过模版创建，需要传入模版中该指标的对应RuleId
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 组策略名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组所属视图的名称，若通过模版创建，可不传入
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 策略组所属项目Id，会进行鉴权操作
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 模版策略组Id, 通过模版创建时才需要传
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`

	// 是否屏蔽策略组，0表示不屏蔽，1表示屏蔽。不填默认为0
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// 策略组的备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 插入时间，戳格式为Unix时间戳，不填则按后台处理时间填充
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 策略组中的阈值告警规则
	Conditions []*CreatePolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// 策略组中的时间告警规则
	EventConditions []*CreatePolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// 是否为后端调用。当且仅当值为1时，后台拉取策略模版中的规则填充入Conditions以及EventConditions字段
	BackEndCall *int64 `json:"BackEndCall,omitempty" name:"BackEndCall"`

	// 指标告警规则的且或关系，0表示或规则(满足任意规则就告警)，1表示且规则(满足所有规则才告警)
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

func (r *CreatePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功的策略组Id
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *ServiceInfo `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStrategyRequest struct {
	*tchttp.BaseRequest

	// 告警策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 告警对象类型
	MixType *int64 `json:"MixType,omitempty" name:"MixType"`

	// 告警有效起始时间
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" name:"EffectiveStartTime"`

	// 告警有效结束事件
	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" name:"EffectiveEndTime"`

	// 告警回调url
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 策略的告警规则列表
	Alarm []*StrategyEntryAlarm `json:"Alarm,omitempty" name:"Alarm" list`

	// 告警策略启停状态
	IsStart *int64 `json:"IsStart,omitempty" name:"IsStart"`

	// 告警类型
	StrategyAlarmType *int64 `json:"StrategyAlarmType,omitempty" name:"StrategyAlarmType"`

	// 当MixType为服务器时，标记服务器子类型
	MixSubType *int64 `json:"MixSubType,omitempty" name:"MixSubType"`

	// 告警对象类型ID
	MixId []*int64 `json:"MixId,omitempty" name:"MixId" list`

	// 告警接收类型
	ReceiverType *int64 `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 告警接收用户或用户组ID列表
	ReceiverId []*uint64 `json:"ReceiverId,omitempty" name:"ReceiverId" list`

	// 告警渠道列表
	AlarmChannel []*string `json:"AlarmChannel,omitempty" name:"AlarmChannel" list`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *StrategyIdInfo `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomAlarmList struct {

	// IP
	LocalIP *string `json:"LocalIP,omitempty" name:"LocalIP"`

	// 消息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 发生时间
	OccurTime *string `json:"OccurTime,omitempty" name:"OccurTime"`

	// 调用方
	Caller *string `json:"Caller,omitempty" name:"Caller"`
}

type CvmAgentStatus struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent是否安装
	AgentInstalled *bool `json:"AgentInstalled,omitempty" name:"AgentInstalled"`
}

type DataPoint struct {

	// 实例对象维度组合
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 时间戳数组，表示那些时间点有数据，缺失的时间戳，没有数据点，可以理解为掉点了
	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps" list`

	// 监控值数组，该数组和Timestamps一一对应
	Values []*float64 `json:"Values,omitempty" name:"Values" list`
}

type DeleteAlertPolicyRequest struct {
	*tchttp.BaseRequest

	// 要删除的策略ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlertPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功为true
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlertPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributeRequest struct {
	*tchttp.BaseRequest

	// 属性ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

func (r *DeleteAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *AttributeIdDeleteOutput `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributesRequest struct {
	*tchttp.BaseRequest

	// 指标ID列表
	AttributeId []*uint64 `json:"AttributeId,omitempty" name:"AttributeId" list`
}

func (r *DeleteAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *AttributeIdDeleteOutput `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCCMChartsRequest struct {
	*tchttp.BaseRequest

	// 监控图表ID列表。
	ChartId []*int64 `json:"ChartId,omitempty" name:"ChartId" list`
}

func (r *DeleteCCMChartsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCCMChartsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCCMChartsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCCMChartsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCCMChartsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConditionsTemplateRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 模板策略组ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 是否删除关联策略
	IsDeleteRelatedPolicy *int64 `json:"IsDeleteRelatedPolicy,omitempty" name:"IsDeleteRelatedPolicy"`

	// 所属用户UIN
	OwnerUIN *string `json:"OwnerUIN,omitempty" name:"OwnerUIN"`
}

func (r *DeleteConditionsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConditionsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConditionsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConditionsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConditionsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 监控面板ID
	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
}

func (r *DeleteDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDashboardRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDashboardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardViewRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 监控面板视图ID
	ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`
}

func (r *DeleteDashboardViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDashboardViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDashboardViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 分组Id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组节点
		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 实例分组id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 是否删除关联的告警策略组，1表示删除，其他数值表示不删除
	IsDelRelatedPolicy *int64 `json:"IsDelRelatedPolicy,omitempty" name:"IsDelRelatedPolicy"`
}

func (r *DeleteInstanceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesInInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 实例组Id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 实例的uniqueId列表
	UniqueIds []*string `json:"UniqueIds,omitempty" name:"UniqueIds" list`
}

func (r *DeleteInstancesInInstanceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstancesInInstanceGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesInInstanceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstancesInInstanceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstancesInInstanceGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesRequest struct {
	*tchttp.BaseRequest

	// 分组Id
	Instances []*CgrpInstance `json:"Instances,omitempty" name:"Instances" list`
}

func (r *DeleteInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组列表
		Data []*CgrpInstanceNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMetricSetRequest struct {
	*tchttp.BaseRequest

	// 指标集ID
	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`
}

func (r *DeleteMetricSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMetricSetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMetricSetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除成功失败布尔值
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMetricSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMetricSetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleRequest struct {
	*tchttp.BaseRequest

	// 分组Id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已删除的分组节点
		Data []*CgrpModuleNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMsgPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 自定义消息策略ID
	PolicyID *string `json:"PolicyID,omitempty" name:"PolicyID"`
}

func (r *DeleteMsgPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMsgPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMsgPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMsgPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMsgPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId []*int64 `json:"GroupId,omitempty" name:"GroupId" list`
}

func (r *DeletePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteStrategysRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID列表
	StrategyId []*uint64 `json:"StrategyId,omitempty" name:"StrategyId" list`
}

func (r *DeleteStrategysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteStrategysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteStrategysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStrategysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteStrategysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalObjectsEventObject struct {

	// 告警状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 告警内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维度组合，json字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 告警状态，“abnormal”表示告警未恢复，“warning”表示告警已恢复
	AbnormalStatus *string `json:"AbnormalStatus,omitempty" name:"AbnormalStatus"`

	// 首次发生时间
	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

	// 最后发生时间
	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

	// 维度组合md5
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// 策略组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品中文名称
	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`

	// 产品英文名称
	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`

	// 事件英文名称
	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`

	// 事件中文名称
	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`

	// 产品事件id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 产品维度组合md5
	EventUniqueId *string `json:"EventUniqueId,omitempty" name:"EventUniqueId"`
}

type DescribeAbnormalObjectsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警类型名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 起始时间，格式:"2019-01-01"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 项目id列表
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 告警状态，"all"表示所有实例，“abnormal”表示异常实例（告警未恢复），“warning”表示提醒实例（告警已恢复）
	AbnormalStatus *string `json:"AbnormalStatus,omitempty" name:"AbnormalStatus"`

	// 是否汇聚，不填表示不汇聚，"instance"表示按实例汇聚
	GroupBy *string `json:"GroupBy,omitempty" name:"GroupBy"`

	// 按维度过滤，json字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *DescribeAbnormalObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAbnormalObjectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalObjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标告警异常对象列表
		ThresholdObjects []*DescribeAbnormalObjectsThresholdObject `json:"ThresholdObjects,omitempty" name:"ThresholdObjects" list`

		// 产品事件告警异常对象列表
		EventObjects []*DescribeAbnormalObjectsEventObject `json:"EventObjects,omitempty" name:"EventObjects" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAbnormalObjectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalObjectsThresholdObject struct {

	// 告警id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 告警规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 告警状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 告警内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 首次发生时间
	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

	// 最后发生时间
	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

	// 维度组合，json字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 维度组合md5
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// 告警状态，“abnormal”表示告警未恢复，“warning”表示告警已恢复
	AbnormalStatus *string `json:"AbnormalStatus,omitempty" name:"AbnormalStatus"`

	// 对象id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeAccidentConfigAccident struct {

	// 业务id
	BusinessId *int64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 业务名称
	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`

	// 业务英文名称
	BusinessEnName *string `json:"BusinessEnName,omitempty" name:"BusinessEnName"`

	// 策略展示名称
	PolicyShowName *string `json:"PolicyShowName,omitempty" name:"PolicyShowName"`

	// 策略展示英文名称
	PolicyShowEnName *string `json:"PolicyShowEnName,omitempty" name:"PolicyShowEnName"`

	// 策略类型名称
	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`

	// 平台事件类型列表
	AccidentTypeInfo []*DescribeAccidentConfigAccidentTypeInfo `json:"AccidentTypeInfo,omitempty" name:"AccidentTypeInfo" list`
}

type DescribeAccidentConfigAccidentTypeInfo struct {

	// 平台事件id
	AccidentId *int64 `json:"AccidentId,omitempty" name:"AccidentId"`

	// 平台事件名称
	AccidentName *string `json:"AccidentName,omitempty" name:"AccidentName"`

	// 平台事件英文名称
	AccidentEnName *string `json:"AccidentEnName,omitempty" name:"AccidentEnName"`
}

type DescribeAccidentConfigRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAccidentConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccidentConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件配置列表
		AccidentConfigs []*DescribeAccidentConfigAccident `json:"AccidentConfigs,omitempty" name:"AccidentConfigs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccidentConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccidentConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentEventListAlarms struct {

	// 事件分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessTypeDesc *string `json:"BusinessTypeDesc,omitempty" name:"BusinessTypeDesc"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccidentTypeDesc *string `json:"AccidentTypeDesc,omitempty" name:"AccidentTypeDesc"`

	// 事件分类的ID，1表示服务问题，2表示其他订阅
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessID *int64 `json:"BusinessID,omitempty" name:"BusinessID"`

	// 事件状态的ID，0表示已恢复，1表示未恢复
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventStatus *int64 `json:"EventStatus,omitempty" name:"EventStatus"`

	// 影响的对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectResource *string `json:"AffectResource,omitempty" name:"AffectResource"`

	// 事件的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 事件发生的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OccurTime *string `json:"OccurTime,omitempty" name:"OccurTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAccidentEventListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前接口取值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 起始时间，默认一天前的时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认当前时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 根据UpdateTime排序的规则，取值asc或desc
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// 根据OccurTime排序的规则，取值asc或desc（优先根据UpdateTimeOrder排序）
	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`

	// 根据事件类型过滤，1表示服务问题，2表示其他订阅
	AccidentType []*int64 `json:"AccidentType,omitempty" name:"AccidentType" list`

	// 根据事件过滤，1表示云服务器存储问题，2表示云服务器网络连接问题，3表示云服务器运行异常，202表示运营商网络抖动
	AccidentEvent []*int64 `json:"AccidentEvent,omitempty" name:"AccidentEvent" list`

	// 根据事件状态过滤，0表示已恢复，1表示未恢复
	AccidentStatus []*int64 `json:"AccidentStatus,omitempty" name:"AccidentStatus" list`

	// 根据事件地域过滤，gz表示广州，sh表示上海等
	AccidentRegion []*string `json:"AccidentRegion,omitempty" name:"AccidentRegion" list`

	// 根据影响资源过滤，比如ins-19a06bka
	AffectResource *string `json:"AffectResource,omitempty" name:"AffectResource"`
}

func (r *DescribeAccidentEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccidentEventListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 平台事件列表
		Alarms []*DescribeAccidentEventListAlarms `json:"Alarms,omitempty" name:"Alarms" list`

		// 平台事件的总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccidentEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccidentEventListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentStatusHistory struct {

	// 时间戳
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAgentStatusHistoryData struct {

	// 实例ID
	UnInstanceID *string `json:"UnInstanceID,omitempty" name:"UnInstanceID"`

	// 子机状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 历史数据
	Histories []*DescribeAgentStatusHistory `json:"Histories,omitempty" name:"Histories" list`
}

type DescribeAgentStatusHistoryRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 实例ID列表
	UnInstanceIDs []*string `json:"UnInstanceIDs,omitempty" name:"UnInstanceIDs" list`

	// 查询终止时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAgentStatusHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentStatusHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentStatusHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子机状态信息列表
		Data []*DescribeAgentStatusHistoryData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentStatusHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentStatusHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmBindingInstanceListRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 绑定告警策略的AppID
	AppID *int64 `json:"AppID,omitempty" name:"AppID"`

	// 告警策略绑定的视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 查询限制，取值0 - 100，默认100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAlarmBindingInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmBindingInstanceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmBindingInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询结果数量
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// 实例列表
		InstanceList []*AlarmBindingInstance `json:"InstanceList,omitempty" name:"InstanceList" list`

		// 本次查询的限额
		Limit *int64 `json:"Limit,omitempty" name:"Limit"`

		// 本次查询的偏移量
		Offset *int64 `json:"Offset,omitempty" name:"Offset"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmBindingInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmBindingInstanceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmCallbackHistory struct {

	// 回调地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 回调是否验证
	ValidFlag *int64 `json:"ValidFlag,omitempty" name:"ValidFlag"`

	// 回调地址验证码
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

type DescribeAlarmCallbackHistoryRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmCallbackHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmCallbackHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmCallbackHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回调历史列表
		List []*DescribeAlarmCallbackHistory `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmCallbackHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmCallbackHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmCallbackVerifyCodeRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmCallbackVerifyCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmCallbackVerifyCodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmCallbackVerifyCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证码
		VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmCallbackVerifyCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmCallbackVerifyCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmHistoryByAlarmIdRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警Id
	AlarmId *int64 `json:"AlarmId,omitempty" name:"AlarmId"`

	// 起始时间，比如2020-02-07 15:47:24
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，比如2020-02-26 16:47:18
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAlarmHistoryByAlarmIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmHistoryByAlarmIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmHistoryByAlarmIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警列表
		Alarms []*HistoryAlarmInfo `json:"Alarms,omitempty" name:"Alarms" list`

		// 告警总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmHistoryByAlarmIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmHistoryByAlarmIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警Id
	AlarmId *int64 `json:"AlarmId,omitempty" name:"AlarmId"`
}

func (r *DescribeAlarmInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警Id
		Id *int64 `json:"Id,omitempty" name:"Id"`

		// 策略组Id
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// 策略组名
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 告警状态，OK、ALARM、NO_CONF、NO_DATA
		AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

		// 告警类型，0（指标告警），2（产品事件告警），3（平台事件告警）
		AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`

		// 阈值
		CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

		// 当前值
		CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

		// 告警内容
		Content *string `json:"Content,omitempty" name:"Content"`

		// 首次触发时间
		FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

		// 最后触发时间
		LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

		// 持续时长
		Duration *int64 `json:"Duration,omitempty" name:"Duration"`

		// 对象Id
		ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

		// 对象名
		ObjName *string `json:"ObjName,omitempty" name:"ObjName"`

		// 项目Id
		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// 项目名
		ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

		// 地域
		Region *string `json:"Region,omitempty" name:"Region"`

		// 取值 0，1，2，3，4，5，其中0表示未恢复，1表示已恢复，4表示已失效，2/3/5表示数据不足。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 视图名
		ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

		// 通知方式
		NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmSmsQuotaQuota struct {

	// 配额类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配额名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 免费配额剩余量
	FreeLeft *int64 `json:"FreeLeft,omitempty" name:"FreeLeft"`

	// 付费配额剩余量
	PurchaseLeft *int64 `json:"PurchaseLeft,omitempty" name:"PurchaseLeft"`

	// 已使用量
	Used *int64 `json:"Used,omitempty" name:"Used"`
}

type DescribeAlarmSmsQuotaRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmSmsQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmSmsQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmSmsQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 总使用量
		Used *int64 `json:"Used,omitempty" name:"Used"`

		// 短信配额信息列表
		QuotaList []*DescribeAlarmSmsQuotaQuota `json:"QuotaList,omitempty" name:"QuotaList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmSmsQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmSmsQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertPoliciesRequest struct {
	*tchttp.BaseRequest

	// 分页offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段和方式 （id desc）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 按ID精确搜索
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 按名字模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 按指标集id搜索
	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
}

func (r *DescribeAlertPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlertPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// data
		Data *ClmDescribeAlertPoliciesData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlertPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertPolicyRequest struct {
	*tchttp.BaseRequest

	// 告警策略id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlertPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略
		Data *ClmAlertPolicy `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlertPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppFlowConverterResponseDataPoint struct {

	// 数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DescribeAppFlowRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 名字空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 周期
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeAppFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppFlowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 周期
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// 指标名
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 数据点
		DataPoints []*DescribeAppFlowConverterResponseDataPoint `json:"DataPoints,omitempty" name:"DataPoints" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppFlowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeAggregateDataRequest struct {
	*tchttp.BaseRequest

	// 指标ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标聚合方式
	AttributeAggregation *int64 `json:"AttributeAggregation,omitempty" name:"AttributeAggregation"`

	// 地域列表
	IdcId []*int64 `json:"IdcId,omitempty" name:"IdcId" list`
}

func (r *DescribeAttributeAggregateDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttributeAggregateDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeAggregateDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *AttributeAggrValueInfoOutputData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttributeAggregateDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttributeAggregateDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeAllServerRequest struct {
	*tchttp.BaseRequest

	// 属性ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询关键字：服务器ID、服务器名称或服务器ip 搜索
	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
}

func (r *DescribeAttributeAllServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttributeAllServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeAllServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据列表
		Data *AttributeServerInfoOutputData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttributeAllServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttributeAllServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeUnitsRequest struct {
	*tchttp.BaseRequest

	// 指标单位ID列表
	UnitId []*uint64 `json:"UnitId,omitempty" name:"UnitId" list`

	// 指标单位名称列表
	UnitName []*string `json:"UnitName,omitempty" name:"UnitName" list`

	// 查询关键字：指标单位ID或指标单位名称搜索
	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
}

func (r *DescribeAttributeUnitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttributeUnitsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeUnitsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *AttributeUnitInfoOutputData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttributeUnitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttributeUnitsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributesRequest struct {
	*tchttp.BaseRequest

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指标id列表
	AttributeId []*uint64 `json:"AttributeId,omitempty" name:"AttributeId" list`

	// 指标名称列表
	AttributeName []*string `json:"AttributeName,omitempty" name:"AttributeName" list`

	// 指标类型名称列表
	AttributeType []*string `json:"AttributeType,omitempty" name:"AttributeType" list`

	// 负责人名称列表
	OwnerName []*string `json:"OwnerName,omitempty" name:"OwnerName" list`

	// 查询关键字
	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
}

func (r *DescribeAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据列表
		Data *AttributeInfoOutputData `json:"Data,omitempty" name:"Data"`

		// 请求处理时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsDataMeaning struct {

	// 语言
	Language *string `json:"Language,omitempty" name:"Language"`

	// 含义
	Meaning *string `json:"Meaning,omitempty" name:"Meaning"`
}

type DescribeBaseMetricsDataPeriod struct {

	// 周期
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatType []*string `json:"StatType,omitempty" name:"StatType" list`
}

type DescribeBaseMetricsForConsoleFontEndData struct {

	// 名字空间
	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标中文名
	MetricCNName *string `json:"MetricCNName,omitempty" name:"MetricCNName"`

	// 指标英文名
	MetricENName *string `json:"MetricENName,omitempty" name:"MetricENName"`

	// 维度
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 单位中午名
	UnitCNName *string `json:"UnitCNName,omitempty" name:"UnitCNName"`

	// 周期与统计方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodAndStatTypes []*DescribeBaseMetricsDataPeriod `json:"PeriodAndStatTypes,omitempty" name:"PeriodAndStatTypes" list`

	// 指标含义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Meaning []*DescribeBaseMetricsDataMeaning `json:"Meaning,omitempty" name:"Meaning" list`
}

type DescribeBaseMetricsForConsoleFontEndRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 名字空间
	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标名列表
	Batch *string `json:"Batch,omitempty" name:"Batch"`
}

func (r *DescribeBaseMetricsForConsoleFontEndRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseMetricsForConsoleFontEndRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsForConsoleFontEndResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标数据
		Data []*DescribeBaseMetricsForConsoleFontEndData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaseMetricsForConsoleFontEndResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseMetricsForConsoleFontEndResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest

	// 业务命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeBaseMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseMetricsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询得到的指标描述列表
		MetricSet []*MetricSet `json:"MetricSet,omitempty" name:"MetricSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaseMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAlarmListAlarms struct {

	// 该条告警的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 告警状态ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 告警状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 策略组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 发生时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

	// 持续时间，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

	// 告警内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 告警对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`

	// 告警对象ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// VPC，只有CVM有
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 指标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 告警类型，0表示指标告警，2表示产品事件告警，3表示平台事件告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 告警对象维度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 通知方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// 所属实例组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroup []*InstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup" list`
}

type DescribeBasicAlarmListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前取值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 起始时间，默认一天前的时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认当前时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 根据发生时间排序，取值ASC或DESC
	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`

	// 根据项目ID过滤
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 根据策略类型过滤
	ViewNames []*string `json:"ViewNames,omitempty" name:"ViewNames" list`

	// 根据告警状态过滤
	AlarmStatus []*int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus" list`

	// 根据告警对象过滤
	ObjLike *string `json:"ObjLike,omitempty" name:"ObjLike"`

	// 根据实例组ID过滤
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitempty" name:"InstanceGroupIds" list`
}

func (r *DescribeBasicAlarmListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBasicAlarmListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警列表
		Alarms []*DescribeBasicAlarmListAlarms `json:"Alarms,omitempty" name:"Alarms" list`

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicAlarmListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBasicAlarmListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListDimension struct {

	// 地域id
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域简称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 维度组合json字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 事件维度组合json字符串
	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type DescribeBindingPolicyObjectListInstance struct {

	// 对象唯一id
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// 表示对象实例的维度集合，jsonObj字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 对象是否被屏蔽，0表示未屏蔽，1表示被屏蔽
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// 对象所在的地域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeBindingPolicyObjectListInstanceGroup struct {

	// 实例分组id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 告警策略类型名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最后编辑uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 实例分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数量
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 实例所在的地域集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regions []*string `json:"Regions,omitempty" name:"Regions" list`
}

type DescribeBindingPolicyObjectListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 筛选对象的维度信息
	Dimensions []*DescribeBindingPolicyObjectListDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

func (r *DescribeBindingPolicyObjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBindingPolicyObjectListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 绑定的对象实例列表
		List []*DescribeBindingPolicyObjectListInstance `json:"List,omitempty" name:"List" list`

		// 绑定的对象实例总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 未屏蔽的对象实例数
		NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`

		// 绑定的实例分组信息，没有绑定实例分组则为空
		InstanceGroup *DescribeBindingPolicyObjectListInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindingPolicyObjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBindingPolicyObjectListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMChartsRequest struct {
	*tchttp.BaseRequest

	// 页码。默认为1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量。默认为全量数据总数。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 监控图表ID列表
	ChartId []*int64 `json:"ChartId,omitempty" name:"ChartId" list`

	// 监控图表名称
	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`

	// 监控图表类型。1 明细视图， 2 聚合视图
	ChartType *int64 `json:"ChartType,omitempty" name:"ChartType"`

	// 产品类型。1 基础监控， 2 自定义监控
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 指标ID
	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 分组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 排序方式。仅支持："asc","ASC","desc","DESC"
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段。仅支持："CreateTime","UpdateTime", "ChartName"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeCCMChartsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMChartsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMChartsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *CCMChartData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMChartsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMChartsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewAttributeRequest struct {
	*tchttp.BaseRequest

	// 实例分组 ID。GroupId或ViewId必选其一，若ViewId存在，则以ViewId为准
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 分组视图ID
	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`

	// 查询关键字
	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
}

func (r *DescribeCCMGroupViewAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMGroupViewAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*AttributeInfoOutput `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMGroupViewAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMGroupViewAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeCCMGroupViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMGroupViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *CCMGroupViewEntry `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMGroupViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMGroupViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewStrategyRequest struct {
	*tchttp.BaseRequest

	// 实例分组 ID。GroupId或ViewId必选其一，若ViewId存在，则以ViewId为准
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 分组视图ID
	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCCMGroupViewStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMGroupViewStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *StrategyData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMGroupViewStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMGroupViewStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMInstanceDatasRequest struct {
	*tchttp.BaseRequest

	// 实例名称列表
	InstanceName []*string `json:"InstanceName,omitempty" name:"InstanceName" list`

	// 指标ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 开始时间。例如："2019-09-11 00:00:00"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。例如："2019-09-11 00:00:00"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例类型。输入："cvm_device"或"SCF"
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *DescribeCCMInstanceDatasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMInstanceDatasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMInstanceDatasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*CCMInstanceAttributeDataOutput `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMInstanceDatasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCCMInstanceDatasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConditionsTemplateListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 模板策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 模板策略组ID
	GroupID *string `json:"GroupID,omitempty" name:"GroupID"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按更新时间排序方式，枚举值(asc, desc)
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`
}

func (r *DescribeConditionsTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConditionsTemplateListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConditionsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 模板列表
		TemplateGroupList []*TemplateGroup `json:"TemplateGroupList,omitempty" name:"TemplateGroupList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConditionsTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConditionsTemplateListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContactListRequest struct {
	*tchttp.BaseRequest

	// 固定值, "monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 偏移量, 从1开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数，每页返回的数量，取值1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContactListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContactListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContactListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接收人数量
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 接收人列表
		List []*Contact `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContactListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContactListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCurrentTimestampRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeCurrentTimestampRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCurrentTimestampRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCurrentTimestampResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器当前unix时间戳(秒数)
		Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCurrentTimestampResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCurrentTimestampResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomAlarmListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 自定义策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 起始时间，默认为一天前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认为当前系统时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按发生时间排序方式，asc或desc
	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`

	// 模糊搜索关键词
	Filter *string `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeCustomAlarmListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomAlarmListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 策略ID
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// 策略名
		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

		// 消息列表
		AlarmList []*CustomAlarmList `json:"AlarmList,omitempty" name:"AlarmList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomAlarmListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomAlarmListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmAgentStatusRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// CVM实例列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeCvmAgentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCvmAgentStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmAgentStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例的状态列表
		List []*CvmAgentStatus `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCvmAgentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCvmAgentStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardViewList struct {

	// 视图ID
	ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`

	// 描述名
	DescName *string `json:"DescName,omitempty" name:"DescName"`

	// 名字空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName []*string `json:"MetricName,omitempty" name:"MetricName" list`

	// 元数据json
	Meta *string `json:"Meta,omitempty" name:"Meta"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*string `json:"Instances,omitempty" name:"Instances" list`
}

type DescribeDashboardViewsRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 监控面板ID
	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
}

func (r *DescribeDashboardViewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDashboardViewsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardViewsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视图列表
		ViewList []*DescribeDashboardViewList `json:"ViewList,omitempty" name:"ViewList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardViewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDashboardViewsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardsList struct {

	// DashboardID
	DashboardID *string `json:"DashboardID,omitempty" name:"DashboardID"`

	// 描述名称
	DescName *string `json:"DescName,omitempty" name:"DescName"`

	// 元数据
	Meta *string `json:"Meta,omitempty" name:"Meta"`
}

type DescribeDashboardsRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 产品类型
	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`

	// CustomID
	CustomID *string `json:"CustomID,omitempty" name:"CustomID"`
}

func (r *DescribeDashboardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDashboardsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// dashboard列表
		DashboardList []*DescribeDashboardsList `json:"DashboardList,omitempty" name:"DashboardList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDashboardsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDimensionAnalysisDataRequest struct {
	*tchttp.BaseRequest

	// 指标集ID
	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`

	// 开始时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 结束时间戳
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`

	// 分析指标字段
	AnalyticMetric *string `json:"AnalyticMetric,omitempty" name:"AnalyticMetric"`

	// 分析的维度字段
	AnalyticDimensionKey *string `json:"AnalyticDimensionKey,omitempty" name:"AnalyticDimensionKey"`

	// 分析的维度值列表
	AnalyticDimensionValues []*string `json:"AnalyticDimensionValues,omitempty" name:"AnalyticDimensionValues" list`

	// 维度筛选条件
	Filters []*ClmAnalysisFilter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeDimensionAnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDimensionAnalysisDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDimensionAnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 维度数据结果
		Data *PCLMDimensionAnalysisResultData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDimensionAnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDimensionAnalysisDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEventPolicyConfig struct {

	// 事件ID
	EventID *string `json:"EventID,omitempty" name:"EventID"`

	// 产品名
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品英文名
	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`

	// 产品中文名
	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`

	// 事件名
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 事件英文名
	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`

	// 事件中文名
	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`

	// 维度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions []*DescribeEventPolicyConfigDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type DescribeEventPolicyConfigDimension struct {

	// 维度名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度key值
	Key *string `json:"Key,omitempty" name:"Key"`
}

type DescribeEventPolicyConfigRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 产品列表
	Products []*string `json:"Products,omitempty" name:"Products" list`
}

func (r *DescribeEventPolicyConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEventPolicyConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEventPolicyConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 事件策略配置列表项
		Configs []*DescribeEventPolicyConfig `json:"Configs,omitempty" name:"Configs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventPolicyConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEventPolicyConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGraphDataPartition struct {

	// 分区名
	Name []*string `json:"Name,omitempty" name:"Name" list`

	// 用量
	Usage *float64 `json:"Usage,omitempty" name:"Usage"`

	// 总量
	Total *float64 `json:"Total,omitempty" name:"Total"`
}

type DescribeGraphDataRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 名字空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例ID
	UnInstanceID *string `json:"UnInstanceID,omitempty" name:"UnInstanceID"`

	// 实例UUID
	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

func (r *DescribeGraphDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGraphDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGraphDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询起始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 查询结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 周期
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// 指标名
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 分区使用列表
		Partitions []*DescribeGraphDataPartition `json:"Partitions,omitempty" name:"Partitions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGraphDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGraphDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组树
		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcServerCountRequest struct {
	*tchttp.BaseRequest

	// 地域ID数组列表
	IdcId []*int64 `json:"IdcId,omitempty" name:"IdcId" list`
}

func (r *DescribeIdcServerCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdcServerCountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcServerCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *IdcInfoData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcServerCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdcServerCountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcServerRequest struct {
	*tchttp.BaseRequest

	// 地域ID列表
	IdcId []*int64 `json:"IdcId,omitempty" name:"IdcId" list`
}

func (r *DescribeIdcServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdcServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *AttributeServerInfoOutputData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdcServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcTreeRequest struct {
	*tchttp.BaseRequest

	// 0或者为空拉取一级业务树，传其他id表示二三四级业务树
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeIdcTreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdcTreeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcTreeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*IdcData `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcTreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdcTreeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceGroupListInstanceGroup struct {

	// 实例分组id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 实例分组策略类型名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最近更新人uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 实例分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组内的实例数
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 更新时间，unix时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间，unix时间戳
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 绑定的告警策略组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyGroups []*DescribeInstanceGroupListPolicyGroup `json:"PolicyGroups,omitempty" name:"PolicyGroups" list`
}

type DescribeInstanceGroupListPolicyGroup struct {

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type DescribeInstanceGroupListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按分组名称进行搜索
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例分组id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 告警策略类型名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 更新时间的排序, asc 或者 desc
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`
}

func (r *DescribeInstanceGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceGroupListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例分组列表
		InstanceGroupList []*DescribeInstanceGroupListInstanceGroup `json:"InstanceGroupList,omitempty" name:"InstanceGroupList" list`

		// 实例分组总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceGroupListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 实例组Id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 按实例过滤
	InstanceList []*BindingPolicyObjectDimension `json:"InstanceList,omitempty" name:"InstanceList" list`

	// 按地域过滤
	Regions []*string `json:"Regions,omitempty" name:"Regions" list`

	// 返回的实例列表的分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回的实例列表的分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例组基础信息
		InstanceGroupInfo *InstanceGroupInfo `json:"InstanceGroupInfo,omitempty" name:"InstanceGroupInfo"`

		// 实例列表
		InstanceList []*TransLogInstance `json:"InstanceList,omitempty" name:"InstanceList" list`

		// 实例组关联的策略
		PolicyGroups []*PolicyGroupItem `json:"PolicyGroups,omitempty" name:"PolicyGroups" list`

		// 实例所属的地域
		Regions []*string `json:"Regions,omitempty" name:"Regions" list`

		// 实例总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest

	// 分组Id，1-4级Id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例列表
		Data []*CgrpInstanceNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSetsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果数据json
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTopicIndexRequest struct {
	*tchttp.BaseRequest

	// 无
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeLogTopicIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogTopicIndexRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTopicIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果数据json
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogTopicIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogTopicIndexResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTopicsRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
}

func (r *DescribeLogTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogTopicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果数据json
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogTopicsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricAnalysisDataRequest struct {
	*tchttp.BaseRequest

	// 指标集id
	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`

	// 开始时间戳（秒）
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 结束时间戳（秒）
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`

	// 对比数据开始时间戳（秒）
	CStartTimestamp *int64 `json:"CStartTimestamp,omitempty" name:"CStartTimestamp"`

	// 对比数据结束时间戳（秒）
	CEndTimestamp *int64 `json:"CEndTimestamp,omitempty" name:"CEndTimestamp"`

	// 筛选条件
	Filters []*ClmAnalysisFilter `json:"Filters,omitempty" name:"Filters" list`

	// 基础指标列表
	BasicMetrics []*ClmMetricAnalysisBasicMetric `json:"BasicMetrics,omitempty" name:"BasicMetrics" list`

	// 复合指标列表
	CustomMetrics []*ClmMetricAnalysisCustomMetric `json:"CustomMetrics,omitempty" name:"CustomMetrics" list`
}

func (r *DescribeMetricAnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMetricAnalysisDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricAnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// data
		Data *PCLMMetricAnalysisData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricAnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMetricAnalysisDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSetRequest struct {
	*tchttp.BaseRequest

	// 指标集ID
	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`
}

func (r *DescribeMetricSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMetricSetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标集详情
		Data *ClmMetricSet `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMetricSetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSetsRequest struct {
	*tchttp.BaseRequest

	// 指标集ID
	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`

	// 指标集名称（模糊搜索）
	Name *string `json:"Name,omitempty" name:"Name"`

	// 日志集名称（模糊搜索）
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志主题名称（模糊搜索）
	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`

	// 偏移量，默认为0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数量，默认为20，最大100
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMetricSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMetricSetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标集列表结果数据
		Data *ClmDescribeMetricSetsData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMetricSetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleRequest struct {
	*tchttp.BaseRequest

	// 分组Id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组Id
		Data []*CgrpModuleNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorDataByAlarmIDDataPoint struct {

	// 是否为空
	IsNull *bool `json:"IsNull,omitempty" name:"IsNull"`

	// 数据值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DescribeMonitorDataByAlarmIDRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警ID
	AlarmID *string `json:"AlarmID,omitempty" name:"AlarmID"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeMonitorDataByAlarmIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorDataByAlarmIDRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorDataByAlarmIDResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标中文名
		MetricCName *string `json:"MetricCName,omitempty" name:"MetricCName"`

		// 指标单位
		MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`

		// 指标名
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 周期
		Period *string `json:"Period,omitempty" name:"Period"`

		// 数据点
		DataPoints []*DescribeMonitorDataByAlarmIDDataPoint `json:"DataPoints,omitempty" name:"DataPoints" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorDataByAlarmIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorDataByAlarmIDResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductByIdsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 产品Id列表
	ProductIds []*string `json:"ProductIds,omitempty" name:"ProductIds" list`
}

func (r *DescribeMonitorProductByIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorProductByIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductByIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品列表
		ProductList []*MonitorProductInfo `json:"ProductList,omitempty" name:"ProductList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorProductByIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorProductByIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeMonitorProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorProductsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云产品列表，json 字符串
		ProductList *string `json:"ProductList,omitempty" name:"ProductList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorProductsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略ID
	PolicyID *string `json:"PolicyID,omitempty" name:"PolicyID"`
}

func (r *DescribeMsgPolicyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsgPolicyInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略ID
		PolicyID *string `json:"PolicyID,omitempty" name:"PolicyID"`

		// 策略名
		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

		// 接收组ID
		ReceiverGroupIDs []*int64 `json:"ReceiverGroupIDs,omitempty" name:"ReceiverGroupIDs" list`

		// 通知方式
		NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays" list`

		// 是否为默认
		IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

		// 告警总数
		AlarmCount *int64 `json:"AlarmCount,omitempty" name:"AlarmCount"`

		// 语音配置
		VoiceConfig *DescribeMsgPolicyListVoiceConfig `json:"VoiceConfig,omitempty" name:"VoiceConfig"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsgPolicyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsgPolicyInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyListRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 过滤条件
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMsgPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsgPolicyListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 策略列表
		PolicyList []*MessagePolicy `json:"PolicyList,omitempty" name:"PolicyList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsgPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsgPolicyListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyListVoiceConfig struct {

	// 呼叫次数
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// UIN列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UIDList []*int64 `json:"UIDList,omitempty" name:"UIDList" list`

	// 每轮间隔
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// 没人间隔
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// 是否需要发送提醒
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`
}

type DescribePolicyConditionListCondition struct {

	// 策略视图名称
	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`

	// 事件告警条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventMetrics []*DescribePolicyConditionListEventMetric `json:"EventMetrics,omitempty" name:"EventMetrics" list`

	// 是否支持多地域
	IsSupportMultiRegion *bool `json:"IsSupportMultiRegion,omitempty" name:"IsSupportMultiRegion"`

	// 指标告警条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*DescribePolicyConditionListMetric `json:"Metrics,omitempty" name:"Metrics" list`

	// 策略类型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序id
	SortId *int64 `json:"SortId,omitempty" name:"SortId"`

	// 是否支持默认策略
	SupportDefault *bool `json:"SupportDefault,omitempty" name:"SupportDefault"`

	// 支持该策略类型的地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportRegions []*string `json:"SupportRegions,omitempty" name:"SupportRegions" list`
}

type DescribePolicyConditionListConfigManual struct {

	// 检测方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcType *DescribePolicyConditionListConfigManualCalcType `json:"CalcType,omitempty" name:"CalcType"`

	// 检测阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcValue *DescribePolicyConditionListConfigManualCalcValue `json:"CalcValue,omitempty" name:"CalcValue"`

	// 持续时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinueTime *DescribePolicyConditionListConfigManualContinueTime `json:"ContinueTime,omitempty" name:"ContinueTime"`

	// 数据周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *DescribePolicyConditionListConfigManualPeriod `json:"Period,omitempty" name:"Period"`

	// 持续周期个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodNum *DescribePolicyConditionListConfigManualPeriodNum `json:"PeriodNum,omitempty" name:"PeriodNum"`

	// 聚合方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatType *DescribePolicyConditionListConfigManualStatType `json:"StatType,omitempty" name:"StatType"`
}

type DescribePolicyConditionListConfigManualCalcType struct {

	// CalcType 取值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualCalcValue struct {

	// 默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *string `json:"Default,omitempty" name:"Default"`

	// 固定值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fixed *string `json:"Fixed,omitempty" name:"Fixed"`

	// 最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *string `json:"Max,omitempty" name:"Max"`

	// 最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *string `json:"Min,omitempty" name:"Min"`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualContinueTime struct {

	// 默认持续时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// 可选持续时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriod struct {

	// 默认周期，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// 可选周期，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriodNum struct {

	// 默认周期数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// 可选周期数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*int64 `json:"Keys,omitempty" name:"Keys" list`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualStatType struct {

	// 数据聚合方式，周期5秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	P5 *string `json:"P5,omitempty" name:"P5"`

	// 数据聚合方式，周期10秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	P10 *string `json:"P10,omitempty" name:"P10"`

	// 数据聚合方式，周期1分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	P60 *string `json:"P60,omitempty" name:"P60"`

	// 数据聚合方式，周期5分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	P300 *string `json:"P300,omitempty" name:"P300"`

	// 数据聚合方式，周期10分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	P600 *string `json:"P600,omitempty" name:"P600"`

	// 数据聚合方式，周期30分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	P1800 *string `json:"P1800,omitempty" name:"P1800"`

	// 数据聚合方式，周期1小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	P3600 *string `json:"P3600,omitempty" name:"P3600"`

	// 数据聚合方式，周期1天
	// 注意：此字段可能返回 null，表示取不到有效值。
	P86400 *string `json:"P86400,omitempty" name:"P86400"`
}

type DescribePolicyConditionListEventMetric struct {

	// 事件id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 事件名称
	EventShowName *string `json:"EventShowName,omitempty" name:"EventShowName"`

	// 是否需要恢复
	NeedRecovered *bool `json:"NeedRecovered,omitempty" name:"NeedRecovered"`

	// 事件类型，预留字段，当前固定取值为2
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribePolicyConditionListMetric struct {

	// 指标配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigManual *DescribePolicyConditionListConfigManual `json:"ConfigManual,omitempty" name:"ConfigManual"`

	// 指标id
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 指标名称
	MetricShowName *string `json:"MetricShowName,omitempty" name:"MetricShowName"`

	// 指标单位
	MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`
}

type DescribePolicyConditionListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribePolicyConditionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyConditionListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略条件列表
		Conditions []*DescribePolicyConditionListCondition `json:"Conditions,omitempty" name:"Conditions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyConditionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyConditionListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupCountRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 实例的维度列表
	Batch []*Instance `json:"Batch,omitempty" name:"Batch" list`
}

func (r *DescribePolicyGroupCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupCountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例绑定的告警策略数（json字符串）
		PolicyGroupCount *string `json:"PolicyGroupCount,omitempty" name:"PolicyGroupCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupCountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoCallback struct {

	// 用户回调接口地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 用户回调接口状态，0表示未验证，1表示已验证，2表示存在url但没有通过验证
	ValidFlag *int64 `json:"ValidFlag,omitempty" name:"ValidFlag"`

	// 用户回调接口验证码
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

type DescribePolicyGroupInfoCondition struct {

	// 指标名称
	MetricShowName *string `json:"MetricShowName,omitempty" name:"MetricShowName"`

	// 数据聚合周期(单位秒)
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 指标id
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 阈值规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 指标单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等，7表示日同比上涨，8表示日同比下降，9表示周同比上涨，10表示周同比下降，11表示周期环比上涨，12表示周期环比下降
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// 检测阈值
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// 持续多长时间触发规则会告警(单位秒)
	ContinueTime *int64 `json:"ContinueTime,omitempty" name:"ContinueTime"`
}

type DescribePolicyGroupInfoConditionTpl struct {

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 策略类型
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 策略组说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 最后编辑的用户uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 是否且规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupInfoEventCondition struct {

	// 事件id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 事件告警规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 事件名称
	EventShowName *string `json:"EventShowName,omitempty" name:"EventShowName"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
}

type DescribePolicyGroupInfoReceiverInfo struct {

	// 告警接收组id列表
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList" list`

	// 告警接收人id列表
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList" list`

	// 告警时间段开始时间。范围[0,86400)，作为unix时间戳转成北京时间后去掉日期，例如7200表示"10:0:0"
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 告警时间段结束时间。含义同StartTime
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 接收类型。“group”(接收组)或“user”(接收人)
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 告警通知方式。可选 "SMS","SITE","EMAIL","CALL","WECHAT"
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// 电话告警接收者uid
	// 注意：此字段可能返回 null，表示取不到有效值。
	UidList []*int64 `json:"UidList,omitempty" name:"UidList" list`

	// 电话告警轮数
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// 电话告警每轮间隔（秒）
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// 电话告警对个人间隔（秒）
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// 是否需要电话告警触达提示。0不需要，1需要
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor" list`

	// 恢复通知方式。可选"SMS"
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify" list`
}

type DescribePolicyGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribePolicyGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略组名称
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 策略组所属的项目id
		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// 是否为默认策略，0表示非默认策略，1表示默认策略
		IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

		// 策略类型
		ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

		// 策略说明
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 策略类型名称
		ShowName *string `json:"ShowName,omitempty" name:"ShowName"`

		// 最近编辑的用户uin
		LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

		// 最近编辑时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 该策略支持的地域
		Region []*string `json:"Region,omitempty" name:"Region" list`

		// 策略类型的维度列表
		DimensionGroup []*string `json:"DimensionGroup,omitempty" name:"DimensionGroup" list`

		// 阈值规则列表
		ConditionsConfig []*DescribePolicyGroupInfoCondition `json:"ConditionsConfig,omitempty" name:"ConditionsConfig" list`

		// 产品事件规则列表
		EventConfig []*DescribePolicyGroupInfoEventCondition `json:"EventConfig,omitempty" name:"EventConfig" list`

		// 用户接收人列表
		ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos" list`

		// 用户回调信息
		Callback *DescribePolicyGroupInfoCallback `json:"Callback,omitempty" name:"Callback"`

		// 模板策略组
		ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

		// 是否可以设置成默认策略
		CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

		// 是否且规则
		IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListGroup struct {

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 是否开启
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`

	// 策略视图名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最近编辑的用户uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 最后修改时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 策略组绑定的实例数
	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`

	// 策略组绑定的未屏蔽实例数
	NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`

	// 是否为默认策略，0表示非默认策略，1表示默认策略
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 是否可以设置成默认策略
	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// 父策略组id
	ParentGroupId *int64 `json:"ParentGroupId,omitempty" name:"ParentGroupId"`

	// 策略组备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 策略组所属项目id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 阈值规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*DescribePolicyGroupInfoCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// 产品事件规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventConditions []*DescribePolicyGroupInfoEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// 用户接收人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos" list`

	// 模板策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

	// 策略组绑定的实例组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroup *DescribePolicyGroupListGroupInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

	// 且或规则标识, 0表示或规则(任意一条规则满足阈值条件就告警), 1表示且规则(所有规则都满足阈值条件才告警)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupListGroupInstanceGroup struct {

	// 实例分组名称id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 策略类型视图名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最近编辑的用户uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 实例分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数量
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type DescribePolicyGroupListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 分页参数，每页返回的数量，取值1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按策略名搜索
	Like *string `json:"Like,omitempty" name:"Like"`

	// 实例分组id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 按更新时间排序, asc 或者 desc
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// 项目id列表
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 告警策略类型列表
	ViewNames []*string `json:"ViewNames,omitempty" name:"ViewNames" list`

	// 是否过滤无接收人策略组, 1表示过滤, 0表示不过滤
	FilterUnuseReceiver *int64 `json:"FilterUnuseReceiver,omitempty" name:"FilterUnuseReceiver"`

	// 过滤条件, 接收组列表
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers" list`

	// 过滤条件, 接收人列表
	ReceiverUserList []*string `json:"ReceiverUserList,omitempty" name:"ReceiverUserList" list`

	// 维度组合字段(json字符串), 例如[[{"name":"unInstanceId","value":"ins-6e4b2aaa"}]]
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 模板策略组id, 多个id用逗号分隔
	ConditionTempGroupId *string `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`

	// 过滤条件, 接收人或者接收组, user表示接收人, group表示接收组
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
}

func (r *DescribePolicyGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略组列表
		GroupList []*DescribePolicyGroupListGroup `json:"GroupList,omitempty" name:"GroupList" list`

		// 策略组总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyGroupListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyInfoByInstanceConditionsTemp struct {

	// 组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 组名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 最后编辑者UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 是否为与关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribePolicyInfoByInstanceGroupList struct {

	// 策略组ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 策略名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 是否开启
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最后编辑者UIN
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 是否为默认策略
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 是否能设为默认策略
	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// 父策略组ID
	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是否为与关系策略
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// 告警规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*Condition `json:"Conditions,omitempty" name:"Conditions" list`

	// 规则模板信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionsTemp *DescribePolicyInfoByInstanceConditionsTemp `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

	// 实例组
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroup *DescribePolicyInfoByInstanceInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

	// 未屏蔽实例数
	NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`

	// 实例总数
	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type DescribePolicyInfoByInstanceInstanceGroup struct {

	// 实例组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupID *int64 `json:"InstanceGroupID,omitempty" name:"InstanceGroupID"`

	// 视图
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最后编辑者UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 实例组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type DescribePolicyInfoByInstanceRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 维度
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePolicyInfoByInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyInfoByInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyInfoByInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// UniqueID
		UniqueID *string `json:"UniqueID,omitempty" name:"UniqueID"`

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 策略组信息
		GroupList []*DescribePolicyInfoByInstanceGroupList `json:"GroupList,omitempty" name:"GroupList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyInfoByInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyInfoByInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyObjectCountRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribePolicyObjectCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyObjectCountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyObjectCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否为多地域
		IsMultiRegion *bool `json:"IsMultiRegion,omitempty" name:"IsMultiRegion"`

		// 地域统计列表
		RegionList []*RegionPolicyObjectCount `json:"RegionList,omitempty" name:"RegionList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyObjectCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyObjectCountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyQuotaRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribePolicyQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额描述列表
		Data []*DescribePolicyQuotaV3ResponseData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyQuotaV3ResponseData struct {

	// 视图
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 是否为项目配额
	IsQuotaByProject *bool `json:"IsQuotaByProject,omitempty" name:"IsQuotaByProject"`

	// 配额列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaList []*DescribePolicyQuotaV3ResponseDataQuotaList `json:"QuotaList,omitempty" name:"QuotaList" list`

	// 配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quota *DescribePolicyQuotaV3ResponseDataQuota `json:"Quota,omitempty" name:"Quota"`
}

type DescribePolicyQuotaV3ResponseDataQuota struct {

	// 用户
	Used *int64 `json:"Used,omitempty" name:"Used"`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type DescribePolicyQuotaV3ResponseDataQuotaList struct {

	// 用户
	Used *int64 `json:"Used,omitempty" name:"Used"`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
}

type DescribePolicySituationAlarmSituation struct {

	// 告警规则数
	AlarmRuleCount *int64 `json:"AlarmRuleCount,omitempty" name:"AlarmRuleCount"`

	// 告警触达数
	AlarmTouchCount *int64 `json:"AlarmTouchCount,omitempty" name:"AlarmTouchCount"`
}

type DescribePolicySituationRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 项目id列表
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`
}

func (r *DescribePolicySituationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicySituationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicySituationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 未订阅告警策略组数量
		UnSubscribePolicyGroupCount *int64 `json:"UnSubscribePolicyGroupCount,omitempty" name:"UnSubscribePolicyGroupCount"`

		// 告警状态信息
		AlarmSituation *DescribePolicySituationAlarmSituation `json:"AlarmSituation,omitempty" name:"AlarmSituation"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicySituationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicySituationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyUseList struct {

	// 告警对象唯一id
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// 维度字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type DescribePolicyUseListRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 地域id
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 地域简称
	Area *string `json:"Area,omitempty" name:"Area"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键词
	Search *string `json:"Search,omitempty" name:"Search"`

	// 维度组合json字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *DescribePolicyUseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyUseListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyUseListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警对象数量
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 告警对象列表
		List []*DescribePolicyUseList `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyUseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyUseListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListDimensions struct {

	// 维度名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeProductEventListEvents struct {

	// 事件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 事件中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`

	// 事件英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`

	// 事件简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 产品中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`

	// 产品英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`

	// 产品简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否支持告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportAlarm *int64 `json:"SupportAlarm,omitempty" name:"SupportAlarm"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions []*DescribeProductEventListEventsDimensions `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 实例对象附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdditionMsg []*DescribeProductEventListEventsDimensions `json:"AdditionMsg,omitempty" name:"AdditionMsg" list`

	// 是否配置告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`

	// 策略信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupInfo []*DescribeProductEventListEventsGroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo" list`
}

type DescribeProductEventListEventsDimensions struct {

	// 维度名（英文）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 维度名（中文）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeProductEventListEventsGroupInfo struct {

	// 策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type DescribeProductEventListOverView struct {

	// 状态变更的事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusChangeAmount *int64 `json:"StatusChangeAmount,omitempty" name:"StatusChangeAmount"`

	// 告警状态未配置的事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnConfigAlarmAmount *int64 `json:"UnConfigAlarmAmount,omitempty" name:"UnConfigAlarmAmount"`

	// 异常事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnNormalEventAmount *int64 `json:"UnNormalEventAmount,omitempty" name:"UnNormalEventAmount"`

	// 未恢复的事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnRecoverAmount *int64 `json:"UnRecoverAmount,omitempty" name:"UnRecoverAmount"`
}

type DescribeProductEventListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 产品类型过滤，比如"cvm"表示云服务器
	ProductName []*string `json:"ProductName,omitempty" name:"ProductName" list`

	// 事件名称过滤，比如"guest_reboot"表示机器重启
	EventName []*string `json:"EventName,omitempty" name:"EventName" list`

	// 影响对象，比如ins-19708ino
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId" list`

	// 维度过滤，比如外网IP:10.0.0.1
	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 地域过滤，比如gz
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList" list`

	// 事件类型过滤，取值范围["status_change","abnormal"]，分别表示状态变更、异常事件
	Type []*string `json:"Type,omitempty" name:"Type" list`

	// 事件状态过滤，取值范围["recover","alarm","-"]，分别表示已恢复、未恢复、无状态
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// 项目ID过滤
	Project []*string `json:"Project,omitempty" name:"Project" list`

	// 告警状态配置过滤，1表示已配置，0表示未配置
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`

	// 按更新时间排序，ASC表示升序，DESC表示降序，默认DESC
	TimeOrder *string `json:"TimeOrder,omitempty" name:"TimeOrder"`

	// 起始时间，默认一天前的时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认当前时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 页偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页返回的数量，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductEventListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件列表
		Events []*DescribeProductEventListEvents `json:"Events,omitempty" name:"Events" list`

		// 事件统计
		OverView *DescribeProductEventListOverView `json:"OverView,omitempty" name:"OverView"`

		// 事件总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductEventListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductHealthStatusListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 按视图名过滤，比如"cvm_device"
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 1表示拉取所有地域，0相反
	AllRegions *int64 `json:"AllRegions,omitempty" name:"AllRegions"`

	// 根据项目id过滤
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`
}

func (r *DescribeProductHealthStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductHealthStatusListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductHealthStatusListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 各产品的健康状态列表
		List []*ProductHealthStatus `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductHealthStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductHealthStatusListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeProjectsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目列表
		Projects []*ProjectInfo `json:"Projects,omitempty" name:"Projects" list`

		// 项目总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecommendedTemplateRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

func (r *DescribeRecommendedTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecommendedTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecommendedTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板策略组总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 模板策略组
		GroupTemplateLists []*GroupTemplateList `json:"GroupTemplateLists,omitempty" name:"GroupTemplateLists" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecommendedTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecommendedTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerDataRequest struct {
	*tchttp.BaseRequest

	// 服务器ID
	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`

	// 指标ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 开始时间。例如：1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间。例如：1970-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeServerDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *AttributeValueInfoOutputData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerDatasRequest struct {
	*tchttp.BaseRequest

	// 服务器ID列表
	ServerId []*uint64 `json:"ServerId,omitempty" name:"ServerId" list`

	// 指标ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 开始时间。例如："2019-09-11 00:00:00"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。例如："2019-09-11 00:00:00"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeServerDatasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerDatasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerDatasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*ServerAttributeDataOutput `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerDatasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerDatasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServersRequest struct {
	*tchttp.BaseRequest

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 服务器内网IP列表
	Ip []*string `json:"Ip,omitempty" name:"Ip" list`

	// 服务器id列表
	ServerId []*uint64 `json:"ServerId,omitempty" name:"ServerId" list`

	// 查询关键字。服务器ID、服务器名称或服务器IP模糊搜索
	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`

	// 地域列表
	IdcId []*int64 `json:"IdcId,omitempty" name:"IdcId" list`

	// 排序方式。仅支持："asc","ASC","desc","DESC"
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段。仅支持："ServerId","Ip"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *ServerInfoData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *ServiceInfo `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSortObjectListDimension struct {

	// 维度名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeSortObjectListRequest struct {
	*tchttp.BaseRequest

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 待排序的指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 时间，精确到分钟，格式：“2019-01-01 05:23:00”
	Time *string `json:"Time,omitempty" name:"Time"`

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式，升序（“asc”）或者降序（“desc”）。默认“desc”
	Order *string `json:"Order,omitempty" name:"Order"`

	// 数据聚合周期，单位秒。不填则默认相应指标的默认周期
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 指定的维度组合
	Dimensions []*DescribeSortObjectListDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

func (r *DescribeSortObjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSortObjectListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSortObjectListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 结果对象列表，json字符串
		ObjList *string `json:"ObjList,omitempty" name:"ObjList"`

		// 耗时(秒)
		Cost *float64 `json:"Cost,omitempty" name:"Cost"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSortObjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSortObjectListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageDurationData struct {

	// 周期
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 持续时间
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
}

type DescribeStorageDurationRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeStorageDurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStorageDurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageDurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存储周期及持续时间列表
		Data []*DescribeStorageDurationData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStorageDurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStorageDurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStrategysRequest struct {
	*tchttp.BaseRequest

	// 页码。 默认为1
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量。默认为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 告警策略ID。
	StrategyId []*uint64 `json:"StrategyId,omitempty" name:"StrategyId" list`

	// 告警策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 告警对象类型。1：对象，2：视图
	MixType *uint64 `json:"MixType,omitempty" name:"MixType"`

	// 告警对象类型ID列表
	MixId []*int64 `json:"MixId,omitempty" name:"MixId" list`

	// 告警接收类型
	ReceiverType *uint64 `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 告警接收人ID列表
	ReceiverUserId []*uint64 `json:"ReceiverUserId,omitempty" name:"ReceiverUserId" list`

	// 告警接收用户组列表
	ReceiverGroupId []*uint64 `json:"ReceiverGroupId,omitempty" name:"ReceiverGroupId" list`

	// 默认降序desc。排序方式仅支持："asc","ASC","desc","DESC"
	Order *string `json:"Order,omitempty" name:"Order"`

	// 默认创建时间CreateTime。仅支持："CreateTime","UpdateTime"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 告警对象子类型ID。MixType=1，则1:所有服务器, 2: 指定服务器, 3:实例分组； MixType=2，则 1: 默认视图, 2: 非默认视图，指定视图ID；3：分组视图
	MixSubType *uint64 `json:"MixSubType,omitempty" name:"MixSubType"`
}

func (r *DescribeStrategysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStrategysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStrategysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *StrategyData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStrategysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStrategysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeSubscribeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscribeInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件告警列表
		Events []*SubscribeInfo `json:"Events,omitempty" name:"Events" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscribeInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTransLogRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 模块ID，1表示触发模板，2表示实例组，3表示告警策略
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 实例组ID
	DId *int64 `json:"DId,omitempty" name:"DId"`

	// 每页返回的数量，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTransLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTransLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTransLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 变更日志列表
		List []*TransLogItem `json:"List,omitempty" name:"List" list`

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTransLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTransLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 语种
		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeViewDataRequest struct {
	*tchttp.BaseRequest

	// 指标ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 开始时间。例如："2019-09-11 00:00:00"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。例如："2019-09-11 00:00:00"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 视图ID
	ViewId *uint64 `json:"ViewId,omitempty" name:"ViewId"`

	// 地域ID列表
	IdcId []*uint64 `json:"IdcId,omitempty" name:"IdcId" list`

	// 聚合方式。不传默认累加值（当前仅支持累加值聚合方式）。 (0: SUM累加值 )， (1: AVG平均值)， (2: MAX最大值)， (3: MIN最小值)
	AttributeAggregation *int64 `json:"AttributeAggregation,omitempty" name:"AttributeAggregation"`
}

func (r *DescribeViewDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeViewDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeViewDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *AttributeValueInfoOutputData `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeViewDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeViewDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {

	// 实例维度名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例维度值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DimensionsDesc struct {

	// 维度名数组
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type EventCondition struct {

	// 告警通知频率
	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 重复通知策略预定义（0 - 只告警一次， 1 - 指数告警，2 - 连接告警）
	AlarmNotifyType *string `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 事件ID
	EventID *string `json:"EventID,omitempty" name:"EventID"`

	// 事件展示名称（对外）
	EventDisplayName *string `json:"EventDisplayName,omitempty" name:"EventDisplayName"`

	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

type GetArgusMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 无
	ResourceId *uint64 `json:"ResourceId,omitempty" name:"ResourceId"`

	// 无
	Query *string `json:"Query,omitempty" name:"Query"`

	// 无
	Epoch *string `json:"Epoch,omitempty" name:"Epoch"`

	// 无
	IsEncrypt *uint64 `json:"IsEncrypt,omitempty" name:"IsEncrypt"`
}

func (r *GetArgusMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetArgusMonitorDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetArgusMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetArgusMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetArgusMonitorDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonitorDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计周期
		Period *uint64 `json:"Period,omitempty" name:"Period"`

		// 指标名
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 数据点数组
		DataPoints []*DataPoint `json:"DataPoints,omitempty" name:"DataPoints" list`

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonitorDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTkeDataRequest struct {
	*tchttp.BaseRequest

	// 命名空间。当前有tke_cluster、tke_component、tke_namespace、tke_node、tke_workload、tke_pod、tke_container
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 需要拉取的字段列表，包括维度字段和指标字段，不同Namespace的字段也各不相同
	Fields []*string `json:"Fields,omitempty" name:"Fields" list`

	// 开始时间，格式为Unix时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式为Unix时间戳。默认值当前时间。开始时间和结束时间的最大跨度为7天
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 返回数据条数，默认10，最大值65536
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件，只能按维度字段进行过滤，多个条件之间是“AND”的关系。需要选择合适的维度条件，避免单次查询返回数据条数超过最大值65536条
	Conditions []*Dimension `json:"Conditions,omitempty" name:"Conditions" list`
}

func (r *GetTkeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTkeDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTkeDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回字段列表
		Columns []*string `json:"Columns,omitempty" name:"Columns" list`

		// 数据。需要单独对Data部分进行一次JSON解码
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTkeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTkeDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupTemplateList struct {

	// 模板策略組ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 模板策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 是否为默认策略
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 告警策略启用状态
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 最后修改人UIN
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 未屏蔽的实例数
	NoShieldedInstanceCount *int64 `json:"NoShieldedInstanceCount,omitempty" name:"NoShieldedInstanceCount"`

	// 总绑定实例数
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" name:"TotalInstanceCount"`

	// 父策略组ID
	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 视图
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 指标告警触发条件
	Conditions []*Condition `json:"Conditions,omitempty" name:"Conditions" list`

	// 事件告警触发条件
	EventConditions []*EventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// 是否可设为默认告警策略
	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 修改时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type HistoryAlarmInfo struct {

	// 告警Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 告警类型，0（指标告警），2（产品事件告警），3（平台事件告警）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 当前值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 首次触发时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

	// 最后触发时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

	// 持续时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 对象Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 对象名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 状态，其中0表示未恢复，1表示已恢复，4表示已失效，2/3/5表示数据不足
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 视图名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

type IdcData struct {

	// idc id,在Tce上代表zoneId
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// idc id,在Tce上代表zone名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type IdcInfo struct {

	// 地域ID
	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`

	// 实例数
	ServerCount *int64 `json:"ServerCount,omitempty" name:"ServerCount"`

	// 地域名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type IdcInfoData struct {

	// 返回数据
	Data []*IdcInfo `json:"Data,omitempty" name:"Data" list`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type InitGroupRequest struct {
	*tchttp.BaseRequest

	// 分组树高度
	High *int64 `json:"High,omitempty" name:"High"`

	// 分组名称,数组大小和高度一致
	Name []*string `json:"Name,omitempty" name:"Name" list`

	// 叶子节点存放实例的类型，填SCF或cvm_device
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *InitGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回整棵分组树
		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例的维度组合
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type InstanceGroup struct {

	// 实例组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 实例组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" name:"InstanceGroupName"`
}

type InstanceGroupInfo struct {

	// 实例组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 视图名/策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 最后更新UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
}

type MessagePolicy struct {

	// 策略ID
	PolicyID *string `json:"PolicyID,omitempty" name:"PolicyID"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 接收组ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverGroupIds []*int64 `json:"ReceiverGroupIds,omitempty" name:"ReceiverGroupIds" list`

	// 通知方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// 语音配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceConfig *DescribeMsgPolicyListVoiceConfig `json:"VoiceConfig,omitempty" name:"VoiceConfig"`

	// 是否为默认策略
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 告警次数
	AlarmCount *int64 `json:"AlarmCount,omitempty" name:"AlarmCount"`
}

type MetricDatum struct {

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标的值
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type MetricObjectMeaning struct {

	// 指标英文解释
	En *string `json:"En,omitempty" name:"En"`

	// 指标中文解释
	Zh *string `json:"Zh,omitempty" name:"Zh"`
}

type MetricSet struct {

	// 命名空间，每个云产品会有一个命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标使用的单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 指标使用的单位
	UnitCname *string `json:"UnitCname,omitempty" name:"UnitCname"`

	// 指标支持的统计周期，单位是秒，如60、300
	Period []*int64 `json:"Period,omitempty" name:"Period" list`

	// 统计周期内指标方式
	Periods []*PeriodsSt `json:"Periods,omitempty" name:"Periods" list`

	// 统计指标含义解释
	Meaning *MetricObjectMeaning `json:"Meaning,omitempty" name:"Meaning"`

	// 维度描述信息
	Dimensions []*DimensionsDesc `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type ModifyAlarmCallbackRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 更新动作, "bind"表示绑定回调地址, "unbind"表示解绑回调地址
	UserAction *string `json:"UserAction,omitempty" name:"UserAction"`

	// 用户回调地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 验证码
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`

	// 是否为新增策略组, 1表示新增的策略组, 0表示已有的策略组
	NewPolicy *int64 `json:"NewPolicy,omitempty" name:"NewPolicy"`
}

func (r *ModifyAlarmCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmReceiversRequest struct {
	*tchttp.BaseRequest

	// 需要修改接收人的策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 必填。固定为“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 新接收人信息, 没有填写则删除所有接收人
	ReceiverInfos []*ReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos" list`
}

func (r *ModifyAlarmReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmReceiversRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmReceiversResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmReceiversResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertPolicyRequest struct {
	*tchttp.BaseRequest

	// 修改后的告警策略
	AlertPolicy *ClmAlertPolicy `json:"AlertPolicy,omitempty" name:"AlertPolicy"`

	// 触发条件或筛选条件是否有修改 1有修改 2无修改
	TriggerRuleUpdatedStatus *int64 `json:"TriggerRuleUpdatedStatus,omitempty" name:"TriggerRuleUpdatedStatus"`
}

func (r *ModifyAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlertPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功为true
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlertPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertPolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 告警策略状态 1开启 2关闭
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAlertPolicyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlertPolicyStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertPolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功为true
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlertPolicyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlertPolicyStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAttributeRequest struct {
	*tchttp.BaseRequest

	// 属性ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 属性数据类型
	DataType *int64 `json:"DataType,omitempty" name:"DataType"`

	// 属性名称
	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`

	// 属性类型ID
	AttributeTypeId *uint64 `json:"AttributeTypeId,omitempty" name:"AttributeTypeId"`

	// 属性级别
	AttributeLevel *int64 `json:"AttributeLevel,omitempty" name:"AttributeLevel"`

	// 单位
	Unit *int64 `json:"Unit,omitempty" name:"Unit"`

	// 统计周期
	StatisticalPeriod *int64 `json:"StatisticalPeriod,omitempty" name:"StatisticalPeriod"`

	// 负责人ID
	OwnerId []*uint64 `json:"OwnerId,omitempty" name:"OwnerId" list`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data *AttributeIdOutput `json:"Data,omitempty" name:"Data"`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMChartRequest struct {
	*tchttp.BaseRequest

	// 监控图表名称
	ChartId *int64 `json:"ChartId,omitempty" name:"ChartId"`

	// 监控图表类型。1 明细视图， 2 聚合视图
	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`

	// 产品类型。1 基础监控， 2 自定义监控
	ChartType *int64 `json:"ChartType,omitempty" name:"ChartType"`

	// 指标ID
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 聚合方式。0 SUM， 1 AVG， 2 MAX， 3 MIN
	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 描述
	Aggregation *int64 `json:"Aggregation,omitempty" name:"Aggregation"`

	// 分组ID
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyCCMChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCCMChartRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMChartResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCMChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCCMChartResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMGroupStrategyRequest struct {
	*tchttp.BaseRequest

	// 实例类型。仅支持："cvm_device","SCF"
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`

	// 实例名称列表。
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// 实例分组ID。
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyCCMGroupStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCCMGroupStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMGroupStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCMGroupStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCCMGroupStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMGroupViewRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 分组视图是否自动绑定指标。0：否，1：是
	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`
}

func (r *ModifyCCMGroupViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCCMGroupViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMGroupViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCMGroupViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCCMGroupViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyConditionsTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板策略ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 视图
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 模板名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 是否为与关系
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// 指标告警条件
	Conditions []*ModifyConditionsTemplateRequestCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// 事件告警条件
	EventConditions []*ModifyConditionsTemplateRequestEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 父策略组ID
	ParentGroupID *string `json:"ParentGroupID,omitempty" name:"ParentGroupID"`

	// 是否生效
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`
}

func (r *ModifyConditionsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyConditionsTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyConditionsTemplateRequestCondition struct {

	// 统计周期
	CalcPeriod *string `json:"CalcPeriod,omitempty" name:"CalcPeriod"`

	// 统计方式
	CalcType *string `json:"CalcType,omitempty" name:"CalcType"`

	// 统计值
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// 持续周期
	ContinuePeriod *string `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// 指标ID
	MetricID *int64 `json:"MetricID,omitempty" name:"MetricID"`

	// 告警通知周期
	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 告警通知方式
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

type ModifyConditionsTemplateRequestEventCondition struct {

	// 告警通知周期
	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 告警通知方式
	AlarmNotifyType *string `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 事件ID
	EventID *string `json:"EventID,omitempty" name:"EventID"`

	// 规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

type ModifyConditionsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略组ID
		GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConditionsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyConditionsTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 监控面板名
	DescName *string `json:"DescName,omitempty" name:"DescName"`

	// json描述元数据
	Meta *string `json:"Meta,omitempty" name:"Meta"`

	// 监控面板ID
	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
}

func (r *ModifyDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDashboardRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控面板ID
		DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDashboardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardViewRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图ID
	ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`

	// 名字空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 监控面板ID
	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`

	// 监控面板视图名
	DescName *string `json:"DescName,omitempty" name:"DescName"`

	// 面板实例列表，json字符串数据
	Instances []*string `json:"Instances,omitempty" name:"Instances" list`

	// 元数据，json串
	Meta *string `json:"Meta,omitempty" name:"Meta"`

	// 指标名列表
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames" list`
}

func (r *ModifyDashboardViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDashboardViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视图ID
		ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDashboardViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest

	// 分组Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组节点
		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 实例组名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 视图名
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 实例列表
	InstanceList []*BindingPolicyObjectDimension `json:"InstanceList,omitempty" name:"InstanceList" list`

	// 实例组Id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
}

func (r *ModifyInstanceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMetricSetRequest struct {
	*tchttp.BaseRequest

	// 指标集ID
	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`

	// 指标集详情
	MetricSet *ClmMetricSet `json:"MetricSet,omitempty" name:"MetricSet"`
}

func (r *ModifyMetricSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMetricSetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMetricSetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标集详情
		Data *ClmMetricSet `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMetricSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMetricSetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleRequest struct {
	*tchttp.BaseRequest

	// 分组Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分组名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分组节点允许添加的实例类型
	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *ModifyModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 叶子分组节点
		Data []*CgrpModuleNode `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMsgPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 接收用户组id列表
	ReceiverGroupIds []*int64 `json:"ReceiverGroupIds,omitempty" name:"ReceiverGroupIds" list`

	// 告警方式列表
	NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays" list`

	// 电话告警配置
	VoiceConfig *ModifyMsgPolicyVoiceConfig `json:"VoiceConfig,omitempty" name:"VoiceConfig"`
}

func (r *ModifyMsgPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMsgPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMsgPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略id
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMsgPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMsgPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMsgPolicyVoiceConfig struct {

	// uid列表
	UidList []*int64 `json:"UidList,omitempty" name:"UidList" list`

	// 告警轮数, 1到5之间
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// 每轮告警之间的时间间隔, 单位秒
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// 每个人之间的时间间隔, 单位秒
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// 是否需要通知, 1表示需要, 0表示不需要
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`
}

type ModifyNotifyBatchNotifyInfo struct {

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警方式列表
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`
}

type ModifyNotifyBatchRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警通知方式列表
	GroupNotifyInfos []*ModifyNotifyBatchNotifyInfo `json:"GroupNotifyInfos,omitempty" name:"GroupNotifyInfos" list`
}

func (r *ModifyNotifyBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNotifyBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNotifyBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNotifyBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNotifyBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupCondition struct {

	// 指标id
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// 检测阈值
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// 检测指标的数据周期
	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`

	// 持续周期个数
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 规则id，不填表示新增，填写了ruleId表示在已存在的规则基础上进行修改
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyPolicyGroupEventCondition struct {

	// 事件id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 规则id，不填表示新增，填写了ruleId表示在已存在的规则基础上进行修改
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyPolicyGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 组类型，1表示策略组，2表示实例组，3表示条件模板组
	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`

	// 修改的字段名，可选：groupName、remark，分别表示修改策略名、描述
	Key *string `json:"Key,omitempty" name:"Key"`

	// 新的字段值
	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyPolicyGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPolicyGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组Id
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolicyGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPolicyGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警类型
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 指标告警条件的且或关系，1表示且告警，所有指标告警条件都达到才告警，0表示或告警，任意指标告警条件达到都告警
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// 指标告警条件规则，不填表示删除已有的所有指标告警条件规则
	Conditions []*ModifyPolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions" list`

	// 事件告警条件，不填表示删除已有的事件告警条件
	EventConditions []*ModifyPolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// 模板策略组id
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`
}

func (r *ModifyPolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPolicyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略组id
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPolicyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRecoverNotifyBatchNotifyInfo struct {

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警恢复通知方式列表
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify" list`
}

type ModifyRecoverNotifyBatchRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 不同策略组的恢复通知方式
	GroupNotifyInfos []*ModifyRecoverNotifyBatchNotifyInfo `json:"GroupNotifyInfos,omitempty" name:"GroupNotifyInfos" list`
}

func (r *ModifyRecoverNotifyBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRecoverNotifyBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRecoverNotifyBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecoverNotifyBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRecoverNotifyBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyChannelsRequest struct {
	*tchttp.BaseRequest

	// 告警渠道列表
	Strategy []*StrategyChannelInfo `json:"Strategy,omitempty" name:"Strategy" list`
}

func (r *ModifyStrategyChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStrategyChannelsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyChannelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStrategyChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStrategyChannelsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID
	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`

	// 告警策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 告警对象类型
	MixType *int64 `json:"MixType,omitempty" name:"MixType"`

	// 告警有效起始时间
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" name:"EffectiveStartTime"`

	// 告警有效结束事件
	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" name:"EffectiveEndTime"`

	// 告警回调url
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 策略的告警规则列表
	Alarm []*StrategyEntryAlarm `json:"Alarm,omitempty" name:"Alarm" list`

	// 告警策略启停状态
	IsStart *int64 `json:"IsStart,omitempty" name:"IsStart"`

	// 告警类型
	StrategyAlarmType *int64 `json:"StrategyAlarmType,omitempty" name:"StrategyAlarmType"`

	// 当MixType为服务器时，标记服务器子类型
	MixSubType *int64 `json:"MixSubType,omitempty" name:"MixSubType"`

	// 告警对象类型ID
	MixId []*int64 `json:"MixId,omitempty" name:"MixId" list`

	// 告警接收类型
	ReceiverType *int64 `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 告警接收用户或用户组ID列表
	ReceiverId []*uint64 `json:"ReceiverId,omitempty" name:"ReceiverId" list`

	// 告警渠道列表
	AlarmChannel []*string `json:"AlarmChannel,omitempty" name:"AlarmChannel" list`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyStatesRequest struct {
	*tchttp.BaseRequest

	// 更新启停状态信息
	Strategy []*StrategyStatesInfo `json:"Strategy,omitempty" name:"Strategy" list`
}

func (r *ModifyStrategyStatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStrategyStatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyStatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Data []*int64 `json:"Data,omitempty" name:"Data" list`

		// 请求执行时间
		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStrategyStatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStrategyStatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonitorProductInfo struct {

	// 产品Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 产品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品视图名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`

	// 支持的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvaliableRegions []*string `json:"AvaliableRegions,omitempty" name:"AvaliableRegions" list`

	// Dashboard是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsShowInDashboard *bool `json:"IsShowInDashboard,omitempty" name:"IsShowInDashboard"`

	// 产品的详细信息，json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Meta *string `json:"Meta,omitempty" name:"Meta"`

	// 产品的namespace
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type PCLMDescribeMetricSetsData struct {

	// 指标集列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricSets []*PCLMMetricSet `json:"MetricSets,omitempty" name:"MetricSets" list`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type PCLMDimension struct {

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 创建时间（只作出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 指标集CID（只作出参，入参不填）
	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
}

type PCLMDimensionAnalysisResultData struct {

	// 维度数据
	Current *string `json:"Current,omitempty" name:"Current"`

	// 粒度
	Granularity *int64 `json:"Granularity,omitempty" name:"Granularity"`
}

type PCLMLogFilterRule struct {

	// 关系 1AND 2OR
	Relation *int64 `json:"Relation,omitempty" name:"Relation"`

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 过滤操作符: 数字型字段支持：gt lt ge le eq ne in 字符型字段支持：eq ne in contains
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 过滤操作值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 指标集CID（只作出参，入参不填）
	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`

	// 创建时间（只作出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PCLMLogProfileItem struct {

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 字段类型，可选：string、long、double
	Type *string `json:"Type,omitempty" name:"Type"`

	// 关联指标集CID（只做出参，入参不填）
	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`

	// 创建时间（只做出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只做出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PCLMMetricAnalysisData struct {

	// 对比数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compare []*PCLMMetricAnalysisPoint `json:"Compare,omitempty" name:"Compare" list`

	// 当前数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Current []*PCLMMetricAnalysisPoint `json:"Current,omitempty" name:"Current" list`

	// 时间粒度 （秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Granularity *int64 `json:"Granularity,omitempty" name:"Granularity"`
}

type PCLMMetricAnalysisPoint struct {

	// 时间戳（秒）
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitempty" name:"Value"`

	// 指标名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PCLMMetricBasicItem struct {

	// 字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitempty" name:"Field"`

	// 指标统计方法：sum count min max
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标描述中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 指标过滤规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticFilterRules []*PCLMStatisticFilterRule `json:"StatisticFilterRules,omitempty" name:"StatisticFilterRules" list`

	// 指标集CID（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`

	// 创建时间（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PCLMMetricCustomItem struct {

	// 表达式
	Formula *string `json:"Formula,omitempty" name:"Formula"`

	// 指标名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标描述中文名
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 指标集CID（只作出参，入参不填）
	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`

	// 创建时间（只作出参，入参不填）
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PCLMMetricSet struct {

	// CLS日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// CLS日志集名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// CLS日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogtopicId *string `json:"LogtopicId,omitempty" name:"LogtopicId"`

	// CLS日志主题名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`

	// 指标集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 统计周期 可选：60、300
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticCycle *int64 `json:"StatisticCycle,omitempty" name:"StatisticCycle"`

	// 时间字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeField *string `json:"TimeField,omitempty" name:"TimeField"`

	// 日志过滤规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogFilterRules []*PCLMLogFilterRule `json:"LogFilterRules,omitempty" name:"LogFilterRules" list`

	// 基础指标规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricBasicItems []*PCLMMetricBasicItem `json:"MetricBasicItems,omitempty" name:"MetricBasicItems" list`

	// 复合指标规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricCustomItems []*PCLMMetricCustomItem `json:"MetricCustomItems,omitempty" name:"MetricCustomItems" list`

	// 维度规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions []*PCLMDimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 日志数据结构描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogProfileItems []*PCLMLogProfileItem `json:"LogProfileItems,omitempty" name:"LogProfileItems" list`

	// CID（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CID *string `json:"CID,omitempty" name:"CID"`

	// 创建时间（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// APPID（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// UIN（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 地域（只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// SubAccountUin(只作出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type PCLMStatisticFilterRule struct {

	// 关系 1AND 2OR
	Relation *int64 `json:"Relation,omitempty" name:"Relation"`

	// 字段名
	Field *string `json:"Field,omitempty" name:"Field"`

	// 过滤操作符: 数字型字段支持：gt lt ge le eq ne in 字符型字段支持：eq ne in contains
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 过滤操作值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 关联指标条目ID（只做出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricItemId *int64 `json:"MetricItemId,omitempty" name:"MetricItemId"`

	// 创建时间（只做出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间（只做出参，入参不填）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PeriodsSt struct {

	// 周期
	Period *string `json:"Period,omitempty" name:"Period"`

	// 统计方式
	StatType []*string `json:"StatType,omitempty" name:"StatType" list`
}

type PolicyGroup struct {

	// 是否可设为默认告警策略
	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// 告警策略组ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 告警策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 是否为默认告警策略
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 告警策略启用状态
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 最后修改人UIN
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 未屏蔽的实例数
	NoShieldedInstanceCount *int64 `json:"NoShieldedInstanceCount,omitempty" name:"NoShieldedInstanceCount"`

	// 父策略组ID
	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`

	// 所属项目ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// 告警接收对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverInfos []*PolicyGroupReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos" list`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 修改时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 总绑定实例数
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" name:"TotalInstanceCount"`

	// 视图
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 是否为与关系规则
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type PolicyGroupItem struct {

	// 策略Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type PolicyGroupReceiverInfo struct {

	// 有效时段结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 是否需要发送通知
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// 告警接收渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// 电话告警对个人间隔（秒）
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// 消息接收组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList" list`

	// 接受者类型
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 接收人列表。通过平台接口查询到的接收人id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList" list`

	// 告警恢复通知方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify" list`

	// 电话告警每轮间隔（秒）
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// 电话告警轮数
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor" list`

	// 有效时段开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 电话告警接收者uid
	// 注意：此字段可能返回 null，表示取不到有效值。
	UIDList []*int64 `json:"UIDList,omitempty" name:"UIDList" list`
}

type ProductHealthStatus struct {

	// 视图名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 健康状态，1表示未恢复
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthStatus *int64 `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// 提醒的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarningCount *int64 `json:"WarningCount,omitempty" name:"WarningCount"`

	// 异常的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`

	// 告警的地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmRegions []*string `json:"AlarmRegions,omitempty" name:"AlarmRegions" list`
}

type ProjectInfo struct {

	// AppId
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 创建人Uin
	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 项目名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目信息
	Info *string `json:"Info,omitempty" name:"Info"`

	// 是否为默认
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 主账号Uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 源AppId
	SrcAppId *int64 `json:"SrcAppId,omitempty" name:"SrcAppId"`

	// 源平台
	SrcPlat *string `json:"SrcPlat,omitempty" name:"SrcPlat"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type PutMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 一组指标和数据
	Metrics []*MetricDatum `json:"Metrics,omitempty" name:"Metrics" list`

	// 上报时自行指定的 IP
	AnnounceIp *string `json:"AnnounceIp,omitempty" name:"AnnounceIp"`

	// 上报时自行指定的时间戳
	AnnounceTimestamp *uint64 `json:"AnnounceTimestamp,omitempty" name:"AnnounceTimestamp"`

	// 上报时自行指定的 IP 或 产品实例ID
	AnnounceInstance *string `json:"AnnounceInstance,omitempty" name:"AnnounceInstance"`
}

func (r *PutMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutMonitorDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutMonitorDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReceiverInfo struct {

	// 告警时间段开始时间。范围[0,86400)，作为unix时间戳转成北京时间后去掉日期，例如7200表示"10:0:0"
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 告警时间段结束时间。含义同StartTime
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 告警通知方式。可选 "SMS","SITE","EMAIL","CALL","WECHAT"
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// 接收人类型。“group” 或 “user”
	ReceiverType []*string `json:"ReceiverType,omitempty" name:"ReceiverType" list`

	// Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor" list`

	// 电话告警接收者uid
	UidList []*int64 `json:"UidList,omitempty" name:"UidList" list`

	// 电话告警轮数
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// 电话告警对个人间隔（秒）
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// 电话告警每轮间隔（秒）
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// 恢复通知方式。可选"SMS"
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify" list`

	// 是否需要电话告警触达提示。0不需要，1需要
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// 接收组列表。通过平台接口查询到的接收组id列表
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList" list`

	// 接收人列表。通过平台接口查询到的接收人id列表
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList" list`
}

type RegionPolicyObjectCount struct {

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 绑定的实例数量
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SearchDimensionValueRequest struct {
	*tchttp.BaseRequest

	// 指标集ID
	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`

	// 维度字段
	DimensionField *string `json:"DimensionField,omitempty" name:"DimensionField"`

	// 搜索内容
	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *SearchDimensionValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchDimensionValueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchDimensionValueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 维度值搜索结果列表。
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchDimensionValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchDimensionValueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchLogRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 主题ID列表，逗号分割
	TopicIds *string `json:"TopicIds,omitempty" name:"TopicIds"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 查询
	Query *string `json:"Query,omitempty" name:"Query"`

	// 加载更多使用，透传上次返回的 context 值
	Context *string `json:"Context,omitempty" name:"Context"`
}

func (r *SearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果数据json
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendCustomAlarmMsgRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前取值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 消息策略ID，在云监控自定义消息页面配置
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用户想要发送的自定义消息内容
	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

func (r *SendCustomAlarmMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendCustomAlarmMsgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendCustomAlarmMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCustomAlarmMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendCustomAlarmMsgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ServerAttributeDataOutput struct {

	// 指标ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 服务器ID
	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`

	// 指标上报值
	Values []*AttributeTimestampValueOutput `json:"Values,omitempty" name:"Values" list`
}

type ServerInfo struct {

	// 服务器ID
	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`

	// 服务器名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 服务器内网IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 地域ID
	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`

	// 地域名称
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type ServerInfoData struct {

	// 返回数据
	Data []*ServerInfo `json:"Data,omitempty" name:"Data" list`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ServiceInfo struct {

	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *int64 `json:"ServiceId,omitempty" name:"ServiceId"`

	// 是否开通服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`

	// 默认视图ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`
}

type SetDefaultPolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *SetDefaultPolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetDefaultPolicyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetDefaultPolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultPolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetDefaultPolicyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShieldPolicyAlarmRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警屏蔽状态，1表示屏蔽，0表示不屏蔽
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// 设置指定实例的屏蔽状态
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

func (r *ShieldPolicyAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShieldPolicyAlarmRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShieldPolicyAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShieldPolicyAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShieldPolicyAlarmResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StrategyChannelInfo struct {

	// 告警策略ID
	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`

	// 告警渠道列表
	AlarmChannel []*string `json:"AlarmChannel,omitempty" name:"AlarmChannel" list`
}

type StrategyData struct {

	// 返回数据
	Data []*StrategyEntry `json:"Data,omitempty" name:"Data" list`

	// 查询数据列表总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type StrategyEntry struct {

	// 告警策略ID
	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`

	// 告警策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 告警接受类型
	ReceiverType *int64 `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 告警接受用户或用户组ID列表
	ReceiverId []*uint64 `json:"ReceiverId,omitempty" name:"ReceiverId" list`

	// 有效开始时间
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" name:"EffectiveStartTime"`

	// 有效结束时间
	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" name:"EffectiveEndTime"`

	// 告警渠道列表。每个value可以为："SMS","EMAIL","WECHAT"或"CALL"
	AlarmChannel []*string `json:"AlarmChannel,omitempty" name:"AlarmChannel" list`

	// 告警回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建人ID
	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新人ID
	UpdaterId *uint64 `json:"UpdaterId,omitempty" name:"UpdaterId"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 是否启动。 0：停止中， 1：启动中
	IsStart *int64 `json:"IsStart,omitempty" name:"IsStart"`

	// 告警对象类型。 1：服务器，2：视图
	MixType *uint64 `json:"MixType,omitempty" name:"MixType"`

	// 告警对象子类型。MixType为2，则MixId为视图ID列表，MixId为空数组[]， 则后台会自动查询默认视图ID；MixType为1（服务器）时，若MixSubType为1，MixId可不填；若MixSubType为2，则MixId为服务器ID列表，可含有多个服务器ID
	MixSubType *int64 `json:"MixSubType,omitempty" name:"MixSubType"`

	// 告警规则列表
	Alarm []*StrategyEntryAlarm `json:"Alarm,omitempty" name:"Alarm" list`

	// 告警对象类型ID列表。
	MixId []*int64 `json:"MixId,omitempty" name:"MixId" list`

	// 若mixtype为服务器类型，且子类型为指定服务器id。则存在Ip列表
	Ip []*string `json:"Ip,omitempty" name:"Ip" list`
}

type StrategyEntryAlarm struct {

	// 策略的告警规则类型。100：最大值（即>）， 101：最大值（即>=），102：最小值告警（即 <）， 103：最小值告警（即 <=），104：波动值告警
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 策略的告警规则属性ID
	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`

	// 策略的告警规则阈值
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 策略的告警规则DetectPeriod分钟内告警次数
	DetectCount *int64 `json:"DetectCount,omitempty" name:"DetectCount"`

	// 策略的告警周期
	DetectPeriod *int64 `json:"DetectPeriod,omitempty" name:"DetectPeriod"`

	// 策略的告警规则告警间隔周期
	TriggerPeriod *int64 `json:"TriggerPeriod,omitempty" name:"TriggerPeriod"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 指标中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`

	// 指标英文描述，唯一名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`
}

type StrategyIdInfo struct {

	// 告警策略ID
	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

type StrategyStatesInfo struct {

	// 告警策略ID
	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`

	// 启停状态。默认为1。 0：停止中， 1：启动中
	IsStart *int64 `json:"IsStart,omitempty" name:"IsStart"`
}

type SubscribeEventRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 业务id
	BusinessId *int64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 事件id
	AccidentId *int64 `json:"AccidentId,omitempty" name:"AccidentId"`

	// 通知方式
	NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays" list`

	// 接收人(json字符串)
	Receivers *string `json:"Receivers,omitempty" name:"Receivers"`

	// 用户配置(json字符串)
	UserConfig *string `json:"UserConfig,omitempty" name:"UserConfig"`
}

func (r *SubscribeEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribeEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribeEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribeEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubscribeEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubscribeInfo struct {

	// 事件类型Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *int64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 事件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`

	// 问题类型Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccidentId *int64 `json:"AccidentId,omitempty" name:"AccidentId"`

	// 问题名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccidentName *string `json:"AccidentName,omitempty" name:"AccidentName"`

	// 通知方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*int64 `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// 接收人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Receivers []*SubscribeInfoReceiver `json:"Receivers,omitempty" name:"Receivers" list`

	// 提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 用户配置，json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserConfig *string `json:"UserConfig,omitempty" name:"UserConfig"`
}

type SubscribeInfoReceiver struct {

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 用户Uid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

type TemplateGroup struct {

	// 指标告警规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*Condition `json:"Conditions,omitempty" name:"Conditions" list`

	// 事件告警规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventConditions []*EventCondition `json:"EventConditions,omitempty" name:"EventConditions" list`

	// 关联告警策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyGroups []*PolicyGroup `json:"PolicyGroups,omitempty" name:"PolicyGroups" list`

	// 模板策略组ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 模板策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 最后修改人UIN
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 视图
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 是否为与关系
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type TransLogInstance struct {

	// 实例Dimension
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 实例UniqueID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// 屏蔽状态，1表示已屏蔽，0表示未屏蔽
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type TransLogItem struct {

	// 变更日志Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 模块Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 日志信息，json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogData *string `json:"LogData,omitempty" name:"LogData"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后修改UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
}

type UnBindingAllPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *UnBindingAllPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingAllPolicyObjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindingAllPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingAllPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingAllPolicyObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindingInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 实例组Id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 策略组Id
	PolicyGroupId *int64 `json:"PolicyGroupId,omitempty" name:"PolicyGroupId"`

	// 是否删除策略组，0表示否，1表示是
	IsDelRelatedPolicy *int64 `json:"IsDelRelatedPolicy,omitempty" name:"IsDelRelatedPolicy"`
}

func (r *UnBindingInstanceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingInstanceGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindingInstanceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingInstanceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingInstanceGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 待删除对象实例的唯一id列表
	UniqueId []*string `json:"UniqueId,omitempty" name:"UniqueId" list`

	// 实例分组id, 如果按实例分组删除的话UniqueId参数是无效的
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
}

func (r *UnBindingPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingPolicyObjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindingPolicyObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnsubscribeEventRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 业务id
	BusinessId *int64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 事件id
	AccidentId *int64 `json:"AccidentId,omitempty" name:"AccidentId"`
}

func (r *UnsubscribeEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnsubscribeEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnsubscribeEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnsubscribeEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnsubscribeEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyAlarmCallbackRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *VerifyAlarmCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyAlarmCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyAlarmCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证结果
		VerifyStatus *string `json:"VerifyStatus,omitempty" name:"VerifyStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyAlarmCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyAlarmCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
