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

package v20190606

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收藏夹描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *AddFavoriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddFavoriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddFavoriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 收藏夹Id
		FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddFavoriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddFavoriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据最大保存时间，单位天
	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`

	// 数据最大存储容量，单位GB
	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`

	// 聚合周期
	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`

	// 维度列表
	Dimensions []*DimensionIn `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 指标列表
	Metrics []*MetricIn `json:"Metrics,omitempty" name:"Metrics" list`

	// namespace描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *AddNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 告警接收渠道
	ReceiveChannels []*string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels" list`

	// 触发条件
	PolicyTriggers []*PolicyTriggerIn `json:"PolicyTriggers,omitempty" name:"PolicyTriggers" list`

	// 告警策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in”
	FilterConditions []*string `json:"FilterConditions,omitempty" name:"FilterConditions" list`

	// 分组字段
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`

	// 告警接收用户的ID
	ReceiveUsers []*uint64 `json:"ReceiveUsers,omitempty" name:"ReceiveUsers" list`

	// 告警接收用户组的ID
	ReceiveGroups []*uint64 `json:"ReceiveGroups,omitempty" name:"ReceiveGroups" list`

	// 字段名为中文名，其他与FilterConditions相同
	FilterConditionCns []*string `json:"FilterConditionCns,omitempty" name:"FilterConditionCns" list`
}

func (r *AddPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 所属的收藏夹Id列表
	FavoriteIds []*uint64 `json:"FavoriteIds,omitempty" name:"FavoriteIds" list`

	// 所属的NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 视图名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视图参数
	Param *string `json:"Param,omitempty" name:"Param"`

	// 指标字段
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 视图描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *AddViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlertHistory struct {

	// 告警状态。未恢复、已恢复、已失效
	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 告警内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 告警发生时间
	OccurTime *uint64 `json:"OccurTime,omitempty" name:"OccurTime"`

	// 告警发送状态。已收敛、已屏蔽、已发送、无订阅
	Status *string `json:"Status,omitempty" name:"Status"`

	// 所属告警策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`

	// 所属namespace名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 告警持续时间，单位秒
	AlarmDuration *uint64 `json:"AlarmDuration,omitempty" name:"AlarmDuration"`

	// 告警接收渠道，多个逗号分隔
	ReceiveChannels *string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels"`

	// 所属告警策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警对象名称。只有tke告警历史才有值，argus暂未用到
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`

	// 最近告警发送时间
	RecentAlarmTime *uint64 `json:"RecentAlarmTime,omitempty" name:"RecentAlarmTime"`

	// 所属namespace的Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 告警等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertLevel *string `json:"AlertLevel,omitempty" name:"AlertLevel"`

	// 触发指标
	TriggerMetric *string `json:"TriggerMetric,omitempty" name:"TriggerMetric"`

	// 触发条件
	TriggerDimensions *string `json:"TriggerDimensions,omitempty" name:"TriggerDimensions"`
}

type ArgusReportRequest struct {
	*tchttp.BaseRequest

	// 指定上报的数据格式，其中lp代表Line Protocol，json代表Json数据协议，推荐使用Line Protocol。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当Type指定为lp时使用，数据格式说明请参考：http://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/ ，其中measurement部分填写上报的命名空间名称，时间戳精确到纳秒。
	LpData []*string `json:"LpData,omitempty" name:"LpData" list`

	// 当Type指定为json时使用，支持单个请求上报多个命名空间数据。
	JsonData []*ReportJson `json:"JsonData,omitempty" name:"JsonData" list`
}

func (r *ArgusReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArgusReportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ArgusReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回描述信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArgusReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArgusReportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收藏夹描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateFavoriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFavoriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFavoriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 收藏夹Id
		FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFavoriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFavoriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据最大保存时间，单位天
	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`

	// 数据最大存储容量，单位GB
	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`

	// 聚合周期
	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`

	// 维度列表
	Dimensions []*DimensionIn `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 指标列表
	Metrics []*MetricIn `json:"Metrics,omitempty" name:"Metrics" list`

	// namespace描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 告警接收渠道
	ReceiveChannels []*string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels" list`

	// 触发条件
	PolicyTriggers []*PolicyTriggerIn `json:"PolicyTriggers,omitempty" name:"PolicyTriggers" list`

	// 告警策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in”
	FilterConditions []*string `json:"FilterConditions,omitempty" name:"FilterConditions" list`

	// 分组字段
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`

	// 告警接收用户的ID
	ReceiveUsers []*uint64 `json:"ReceiveUsers,omitempty" name:"ReceiveUsers" list`

	// 告警接收用户组的ID
	ReceiveGroups []*uint64 `json:"ReceiveGroups,omitempty" name:"ReceiveGroups" list`

	// 告警等级
	AlertLevel *string `json:"AlertLevel,omitempty" name:"AlertLevel"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTkePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 告警数据所属命名空间。支持k8s_cluster、k8s_pod、k8s_node
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 过滤条件。其中按tke_cluster_instance_id进行过滤的条件为必填项。由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in”
	FilterConditions []*string `json:"FilterConditions,omitempty" name:"FilterConditions" list`

	// 告警接收渠道
	ReceiveChannels []*string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels" list`

	// 触发条件
	PolicyTriggers []*PolicyTriggerIn `json:"PolicyTriggers,omitempty" name:"PolicyTriggers" list`

	// 告警策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 告警接收用户的ID
	ReceiveUsers []*uint64 `json:"ReceiveUsers,omitempty" name:"ReceiveUsers" list`

	// 告警接收用户组的ID
	ReceiveGroups []*uint64 `json:"ReceiveGroups,omitempty" name:"ReceiveGroups" list`
}

func (r *CreateTkePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTkePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTkePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTkePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTkePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 所属的收藏夹Id列表
	FavoriteIds []*uint64 `json:"FavoriteIds,omitempty" name:"FavoriteIds" list`

	// 所属的NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 视图名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视图参数
	Param *string `json:"Param,omitempty" name:"Param"`

	// 指标字段
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 视图描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteFavoriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFavoriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFavoriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFavoriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFavoriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace的名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略的ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeletePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTkePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略的ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteTkePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTkePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTkePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTkePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTkePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹Id。Id和Name必填其一，优先匹配Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 收藏夹名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescFavoriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescFavoriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescFavoriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的收藏夹
		Favorite *FavoriteOut `json:"Favorite,omitempty" name:"Favorite"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescFavoriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescFavoriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// namespace的Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的namespace
		Namespace *NamespaceOut `json:"Namespace,omitempty" name:"Namespace"`

		// 维度列表
		Dimensions []*DimensionOut `json:"Dimensions,omitempty" name:"Dimensions" list`

		// 指标列表
		Metrics []*MetricOut `json:"Metrics,omitempty" name:"Metrics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略的名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的告警策略
		Policy *PolicyOut `json:"Policy,omitempty" name:"Policy"`

		// 触发条件
		PolicyTriggers []*PolicyTriggerOut `json:"PolicyTriggers,omitempty" name:"PolicyTriggers" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图Id。Id和Name必填其一，优先匹配Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 视图名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收藏夹Id
	FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`
}

func (r *DescViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的视图
		View *ViewOut `json:"View,omitempty" name:"View"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertHistoryListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 告警策略Id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警接收用户组的ID
	ReceiveGroup *uint64 `json:"ReceiveGroup,omitempty" name:"ReceiveGroup"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 分页Offset，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，默认occurTime。另外还支持status、alarmStatus、recentAlarmTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc，默认desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 告警发送状态。已收敛、已屏蔽、已发送、无订阅
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// 告警状态。未恢复、已恢复、已失效
	AlarmStatus []*string `json:"AlarmStatus,omitempty" name:"AlarmStatus" list`

	// 告警接收用户的ID
	ReceiveUser *uint64 `json:"ReceiveUser,omitempty" name:"ReceiveUser"`

	// 告警等级。提醒、一般、严重、非常严重
	AlertLevel []*string `json:"AlertLevel,omitempty" name:"AlertLevel" list`
}

func (r *DescribeAlertHistoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlertHistoryListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertHistoryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 匹配到的告警历史
		AlertHistories []*AlertHistory `json:"AlertHistories,omitempty" name:"AlertHistories" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertHistoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlertHistoryListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDimensionListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *DescribeDimensionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDimensionListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDimensionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Dimension下拉框数据的列表
		Dimensions []*DimensionValuesOut `json:"Dimensions,omitempty" name:"Dimensions" list`

		// 指标列表
		Metrics []*MetricOut `json:"Metrics,omitempty" name:"Metrics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDimensionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDimensionListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFavoriteListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段。"Id", "Name", "CreateTime", "UpdateTime"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeFavoriteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFavoriteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFavoriteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 匹配到的收藏夹列表
		Favorites []*FavoriteOut `json:"Favorites,omitempty" name:"Favorites" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFavoriteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFavoriteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹Id。Id和Name必填其一，优先匹配Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 收藏夹名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeFavoriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFavoriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFavoriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的收藏夹
		Favorite *FavoriteOut `json:"Favorite,omitempty" name:"Favorite"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFavoriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFavoriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNamespaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespaceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的namespace列表
		Namespaces []*NamespaceOut `json:"Namespaces,omitempty" name:"Namespaces" list`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespaceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// namespace的Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的namespace
		Namespace *NamespaceOut `json:"Namespace,omitempty" name:"Namespace"`

		// 维度列表
		Dimensions []*DimensionOut `json:"Dimensions,omitempty" name:"Dimensions" list`

		// 指标列表
		Metrics []*MetricOut `json:"Metrics,omitempty" name:"Metrics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 返回namespace个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// namespace的名称列表
	NamespaceNames []*string `json:"NamespaceNames,omitempty" name:"NamespaceNames" list`
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

		// 最近访问的namespace列表
		RecentNamespaces []*NamespaceOut `json:"RecentNamespaces,omitempty" name:"RecentNamespaces" list`

		// namespace总数
		TotalNamespaceCount *uint64 `json:"TotalNamespaceCount,omitempty" name:"TotalNamespaceCount"`

		// 容量预警的namespace个数
		WarningNamespaceCount *uint64 `json:"WarningNamespaceCount,omitempty" name:"WarningNamespaceCount"`

		// 告警策略总数
		TotalPolicyCount *uint64 `json:"TotalPolicyCount,omitempty" name:"TotalPolicyCount"`

		// 告警中策略个数
		AlertPolicyCount *uint64 `json:"AlertPolicyCount,omitempty" name:"AlertPolicyCount"`

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

type DescribeParamListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeParamListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeParamListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标单位
		Units []*ParamOut `json:"Units,omitempty" name:"Units" list`

		// 聚合方式
		AggTypes []*ParamOut `json:"AggTypes,omitempty" name:"AggTypes" list`

		// 聚合周期
		AggPeriods []*ParamOut `json:"AggPeriods,omitempty" name:"AggPeriods" list`

		// 告警接收渠道
		AlertChannels []*ParamOut `json:"AlertChannels,omitempty" name:"AlertChannels" list`

		// 连续触发阈值的告警检测次数
		AlertContinues []*ParamOut `json:"AlertContinues,omitempty" name:"AlertContinues" list`

		// 告警频率
		AlertFrequencies []*ParamOut `json:"AlertFrequencies,omitempty" name:"AlertFrequencies" list`

		// 指标比较符
		CompareSymbol []*ParamOut `json:"CompareSymbol,omitempty" name:"CompareSymbol" list`

		// namespace配置限制
		NamespaceRestrict []*ParamOut `json:"NamespaceRestrict,omitempty" name:"NamespaceRestrict" list`

		// 主地域
		MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

		// tke告警策略统计周期
		AlertTkeAggPeriods []*ParamOut `json:"AlertTkeAggPeriods,omitempty" name:"AlertTkeAggPeriods" list`

		// 告警等级
		AlertLevels []*ParamOut `json:"AlertLevels,omitempty" name:"AlertLevels" list`

		// 告警恢复状态
		AmpRecoverStatus []*ParamOut `json:"AmpRecoverStatus,omitempty" name:"AmpRecoverStatus" list`

		// 告警发送状态
		AmpSendStatus []*ParamOut `json:"AmpSendStatus,omitempty" name:"AmpSendStatus" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeParamListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段。"Id", "Name", "NamespaceId", "CreateTime", "UpdateTime"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 告警策略Id列表
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`

	// 创建人Uin
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DescribePolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的告警策略列表
		Policies []*PolicyOut `json:"Policies,omitempty" name:"Policies" list`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePolicyListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略的ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
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

		// 匹配到的告警策略
		Policy *PolicyOut `json:"Policy,omitempty" name:"Policy"`

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

type DescribeServiceTopologyJobRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 拉取数据的namespace列表
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces" list`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 服务网格Id
	MeshId *string `json:"MeshId,omitempty" name:"MeshId"`

	// 拓扑粒度。支持service、workload、pod
	TopoType *string `json:"TopoType,omitempty" name:"TopoType"`
}

func (r *DescribeServiceTopologyJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTopologyJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTopologyJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务Id
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTopologyJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTopologyJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTopologyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 拉取数据的namespace列表
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces" list`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 服务网格Id
	MeshId *string `json:"MeshId,omitempty" name:"MeshId"`

	// 拓扑粒度。支持service、workload、pod
	TopoType *string `json:"TopoType,omitempty" name:"TopoType"`
}

func (r *DescribeServiceTopologyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTopologyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTopologyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点。需要单独对Nodes部分进行一次JSON解码
		Nodes *string `json:"Nodes,omitempty" name:"Nodes"`

		// 边
		Edges []*MeshEdgeOut `json:"Edges,omitempty" name:"Edges" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTopologyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTopologyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTopologyResultRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 异步任务Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeServiceTopologyResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTopologyResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTopologyResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务状态：doing表示执行中，建议每隔一秒轮询一次；error表示执行出错；success表示执行成功
		Status *string `json:"Status,omitempty" name:"Status"`

		// 节点。需要单独对Nodes部分进行一次JSON解码
		Nodes *string `json:"Nodes,omitempty" name:"Nodes"`

		// 边
		Edges []*MeshEdgeOut `json:"Edges,omitempty" name:"Edges" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTopologyResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTopologyResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkeAlertHistoryListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 告警策略Id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警接收用户组的ID
	ReceiveGroup *uint64 `json:"ReceiveGroup,omitempty" name:"ReceiveGroup"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 分页Offset，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，默认occurTime。另外还支持status、alarmStatus、recentAlarmTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc，默认desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 告警发送状态。已收敛、已屏蔽、已发送、无订阅
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// 告警状态。未恢复、已恢复、已失效
	AlarmStatus []*string `json:"AlarmStatus,omitempty" name:"AlarmStatus" list`

	// 告警接收用户的ID
	ReceiveUser *uint64 `json:"ReceiveUser,omitempty" name:"ReceiveUser"`

	// 告警对象类型。tke_cluster、tke_node、tke_pod
	ObjType *string `json:"ObjType,omitempty" name:"ObjType"`
}

func (r *DescribeTkeAlertHistoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkeAlertHistoryListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkeAlertHistoryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 匹配到的告警历史
		AlertHistories []*AlertHistory `json:"AlertHistories,omitempty" name:"AlertHistories" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTkeAlertHistoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkeAlertHistoryListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkeAlertMetricListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeTkeAlertMetricListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkeAlertMetricListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkeAlertMetricListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警指标
		TkeAlertMetrics []*TkeAlertMetricOut `json:"TkeAlertMetrics,omitempty" name:"TkeAlertMetrics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTkeAlertMetricListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkeAlertMetricListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkeClusterListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// TKE集群Id列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`
}

func (r *DescribeTkeClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkeClusterListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkeClusterListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群详情列表
		Clusters []*TkeClusterOut `json:"Clusters,omitempty" name:"Clusters" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTkeClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkeClusterListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkePolicyListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段。"Id", "Name", "NamespaceId", "CreateTime", "UpdateTime"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 告警策略Id列表
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`

	// 创建人Uin
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 告警数据所属命名空间。支持k8s_cluster、k8s_pod、k8s_node
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *DescribeTkePolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkePolicyListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkePolicyListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的告警策略列表
		Policies []*PolicyOut `json:"Policies,omitempty" name:"Policies" list`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTkePolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkePolicyListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略的ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeTkePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTkePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的告警策略
		Policy *PolicyOut `json:"Policy,omitempty" name:"Policy"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTkePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTkePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeViewListRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹Id
	FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`

	// 视图名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段。"Id", "Name", "FavoriteId", "CreateTime", "UpdateTime"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 收藏夹名称
	FavoriteName *string `json:"FavoriteName,omitempty" name:"FavoriteName"`
}

func (r *DescribeViewListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeViewListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeViewListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 匹配到的收藏夹列表
		Views []*ViewOut `json:"Views,omitempty" name:"Views" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeViewListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeViewListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图Id。Id和Name必填其一，优先匹配Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 视图名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收藏夹Id
	FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`
}

func (r *DescribeViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的视图
		View *ViewOut `json:"View,omitempty" name:"View"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DimensionIn struct {

	// 英文名
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// 中文名
	CnName *string `json:"CnName,omitempty" name:"CnName"`

	// 操作类型。支持add,delete,update,noChange
	OpType *string `json:"OpType,omitempty" name:"OpType"`

	// Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DimensionOut struct {

	// Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// namespace的Id
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// 中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CnName *string `json:"CnName,omitempty" name:"CnName"`

	// 创建人
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 创建的时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新人
	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`

	// 更新的时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DimensionValuesOut struct {

	// 英文名
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// 值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitempty" name:"Values" list`

	// 中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CnName *string `json:"CnName,omitempty" name:"CnName"`
}

type FavoriteOut struct {

	// 收藏夹Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 收藏夹名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收藏夹下视图数量
	ViewNum *uint64 `json:"ViewNum,omitempty" name:"ViewNum"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 创建人
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 创建的时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新人
	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`

	// 更新的时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 收藏夹描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type GetAlertHistoriesRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 告警策略Id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警接收用户组的ID
	ReceiveGroup *uint64 `json:"ReceiveGroup,omitempty" name:"ReceiveGroup"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 分页Offset，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，默认occurTime。另外还支持status、alarmStatus
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc，默认desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 告警发送状态。已收敛、已屏蔽、已发送、无订阅
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// 告警状态。未恢复、已恢复、已失效
	AlarmStatus []*string `json:"AlarmStatus,omitempty" name:"AlarmStatus" list`
}

func (r *GetAlertHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAlertHistoriesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAlertHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 匹配到的告警历史
		AlertHistories []*AlertHistory `json:"AlertHistories,omitempty" name:"AlertHistories" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAlertHistoriesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDataJobRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// namespace的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 如果同时指定了groupBy参数，那么就能使用聚合函数修饰字段，当前支持sum()、avg()、min()、max()、count()。如果不填，则忽略所有其他参数，返回指定table的10条明细数据（一般用于测试）
	Fields []*string `json:"Fields,omitempty" name:"Fields" list`

	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in” 。暂不支持对聚合后的结果进行过滤
	Conditions []*string `json:"Conditions,omitempty" name:"Conditions" list`

	// 支持按时间粒度聚合，如：timestamp(5s)，粒度单位支持s，m，h，d。如果用了groupBy参数，则查询结果的字段名称为固定格式，比如sum(timeCost)返回的字段名称为timeCost_sum
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`

	// 排序字段，只支持一个字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 分页Offset，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// namespace的Id
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 支持60、300、3600三种粒度，单位秒。默认值60
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *GetDataJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDataJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDataJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务Id
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDataJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDataRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// namespace的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 如果同时指定了groupBy参数，那么就能使用聚合函数修饰字段，当前支持sum()、avg()、min()、max()、count()。如果不填，则忽略所有其他参数，返回指定table的10条明细数据（一般用于测试）
	Fields []*string `json:"Fields,omitempty" name:"Fields" list`

	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in” 。暂不支持对聚合后的结果进行过滤
	Conditions []*string `json:"Conditions,omitempty" name:"Conditions" list`

	// 支持按时间粒度聚合，如：timestamp(5s)，粒度单位支持s，m，h，d。如果用了groupBy参数，则查询结果的字段名称为固定格式，比如sum(timeCost)返回的字段名称为timeCost_sum
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`

	// 排序字段，只支持一个字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 分页Offset，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// namespace的Id
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 支持60、300、3600三种粒度，单位秒。默认值60
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *GetDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据的列名
		Columns []*string `json:"Columns,omitempty" name:"Columns" list`

		// 对应列名的数据。由于云网关限制，需要单独对Data进行JSON解码
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDataResultRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 异步任务Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *GetDataResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDataResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDataResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务状态：doing表示执行中，建议每隔一秒轮询一次；error表示执行出错；success表示执行成功
		Status *string `json:"Status,omitempty" name:"Status"`

		// 返回数据的列名
		Columns []*string `json:"Columns,omitempty" name:"Columns" list`

		// 对应列名的数据。由于云网关限制，需要单独对Data进行JSON解码
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDataResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDimensionsRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *GetDimensionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDimensionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDimensionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Dimension下拉框数据的列表
		Dimensions []*DimensionValuesOut `json:"Dimensions,omitempty" name:"Dimensions" list`

		// 指标列表
		Metrics []*MetricOut `json:"Metrics,omitempty" name:"Metrics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDimensionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDimensionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFavoritesRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段。"Id", "Name", "CreateTime", "UpdateTime"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *GetFavoritesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFavoritesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFavoritesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 匹配到的收藏夹列表
		Favorites []*FavoriteOut `json:"Favorites,omitempty" name:"Favorites" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFavoritesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFavoritesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetNamespacesRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// namespace名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *GetNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetNamespacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的namespace列表
		Namespaces []*NamespaceOut `json:"Namespaces,omitempty" name:"Namespaces" list`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetNamespacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOverviewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 返回namespace个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// namespace的名称列表
	NamespaceNames []*string `json:"NamespaceNames,omitempty" name:"NamespaceNames" list`
}

func (r *GetOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 最近访问的namespace列表
		RecentNamespaces []*NamespaceOut `json:"RecentNamespaces,omitempty" name:"RecentNamespaces" list`

		// namespace总数
		TotalNamespaceCount *uint64 `json:"TotalNamespaceCount,omitempty" name:"TotalNamespaceCount"`

		// 容量预警的namespace个数
		WarningNamespaceCount *uint64 `json:"WarningNamespaceCount,omitempty" name:"WarningNamespaceCount"`

		// 告警策略总数
		TotalPolicyCount *uint64 `json:"TotalPolicyCount,omitempty" name:"TotalPolicyCount"`

		// 告警中策略个数
		AlertPolicyCount *uint64 `json:"AlertPolicyCount,omitempty" name:"AlertPolicyCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetParamsRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *GetParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标单位
		Units []*ParamOut `json:"Units,omitempty" name:"Units" list`

		// 聚合方式
		AggTypes []*ParamOut `json:"AggTypes,omitempty" name:"AggTypes" list`

		// 聚合周期
		AggPeriods []*ParamOut `json:"AggPeriods,omitempty" name:"AggPeriods" list`

		// 告警接收渠道
		AlertChannels []*ParamOut `json:"AlertChannels,omitempty" name:"AlertChannels" list`

		// 连续触发阈值的周期次数
		AlertContinues []*ParamOut `json:"AlertContinues,omitempty" name:"AlertContinues" list`

		// 告警频率
		AlertFrequencies []*ParamOut `json:"AlertFrequencies,omitempty" name:"AlertFrequencies" list`

		// 指标比较符
		CompareSymbol []*ParamOut `json:"CompareSymbol,omitempty" name:"CompareSymbol" list`

		// namespace配置限制
		NamespaceRestrict []*ParamOut `json:"NamespaceRestrict,omitempty" name:"NamespaceRestrict" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPoliciesRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段。"Id", "Name", "NamespaceId", "ContinueCount", "AlertFrequency", "CreateTime", "UpdateTime"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 告警策略Id列表
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`

	// 创建人Uin
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *GetPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配到的告警策略列表
		Policies []*PolicyOut `json:"Policies,omitempty" name:"Policies" list`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTkeDataJobRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 命名空间。当前有k8s_cluster、k8s_component、k8s_namespace、k8s_node、k8s_workload、k8s_pod、k8s_container
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 需要拉取的字段列表，包括维度字段和指标字段，不同Namespace的字段也各不相同。如果同时指定了groupBy参数，那么就能使用聚合函数修饰字段，当前支持sum()、avg()、min()、max()、count()
	Fields []*string `json:"Fields,omitempty" name:"Fields" list`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 查询条件，每个条件需要单独进行JSON编码。多个条件之间是“AND”的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in” 。需要选择合适的维度条件，避免单次查询返回数据条数超过最大值65536条。暂不支持对聚合后的结果进行过滤
	Conditions []*string `json:"Conditions,omitempty" name:"Conditions" list`

	// 需要进行分组统计的字段，最多支持两个维度字段。支持按时间粒度聚合，如：timestamp(5s)，粒度单位支持s，m，h，d。如果用了groupBy参数，则查询结果的字段名称为固定格式，比如sum(timeCost)返回的字段名称为timeCost_sum
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`

	// 排序字段，只支持按一个字段进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 分页Offset，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetTkeDataJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTkeDataJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTkeDataJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务Id
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTkeDataJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTkeDataJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTkeDataRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 命名空间。当前有k8s_cluster、k8s_component、k8s_namespace、k8s_node、k8s_workload、k8s_pod、k8s_container
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 需要拉取的字段列表，包括维度字段和指标字段，不同Namespace的字段也各不相同。如果同时指定了groupBy参数，那么就能使用聚合函数修饰字段，当前支持sum()、avg()、min()、max()、count()
	Fields []*string `json:"Fields,omitempty" name:"Fields" list`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 查询条件，每个条件需要单独进行JSON编码。多个条件之间是“AND”的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in” 。需要选择合适的维度条件，避免单次查询返回数据条数超过最大值65536条。暂不支持对聚合后的结果进行过滤
	Conditions []*string `json:"Conditions,omitempty" name:"Conditions" list`

	// 需要进行分组统计的字段，最多支持两个维度字段。支持按时间粒度聚合，如：timestamp(5s)，粒度单位支持s，m，h，d。如果用了groupBy参数，则查询结果的字段名称为固定格式，比如sum(timeCost)返回的字段名称为timeCost_sum
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`

	// 排序字段，只支持按一个字段进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 分页Offset，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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

		// 对应字段的数据。需要单独对Data部分进行一次JSON解码
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

type GetTkeDataResultRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 异步任务Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *GetTkeDataResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTkeDataResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTkeDataResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务状态：doing表示执行中，建议每隔一秒轮询一次；error表示执行出错；success表示执行成功
		Status *string `json:"Status,omitempty" name:"Status"`

		// 返回字段列表
		Columns []*string `json:"Columns,omitempty" name:"Columns" list`

		// 对应字段的数据。需要单独对Data部分进行一次JSON解码
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTkeDataResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTkeDataResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetViewsRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹Id
	FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`

	// 视图名称，支持模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页Offset，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，只支持一个字段。"Id", "Name", "FavoriteId", "CreateTime", "UpdateTime"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序用asc，降序用desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 收藏夹名称
	FavoriteName *string `json:"FavoriteName,omitempty" name:"FavoriteName"`
}

func (r *GetViewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetViewsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetViewsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 匹配到的收藏夹列表
		Views []*ViewOut `json:"Views,omitempty" name:"Views" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetViewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetViewsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MeshEdgeOut struct {

	// 起始边的Id
	SId *string `json:"SId,omitempty" name:"SId"`

	// 目标边的Id
	DId *string `json:"DId,omitempty" name:"DId"`

	// 协议类型。http、tcp
	Type *string `json:"Type,omitempty" name:"Type"`
}

type MetricIn struct {

	// 英文名
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// 中文名
	CnName *string `json:"CnName,omitempty" name:"CnName"`

	// 聚合方式
	AggType *string `json:"AggType,omitempty" name:"AggType"`

	// 单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 操作类型。支持add,delete,update,noChange
	OpType *string `json:"OpType,omitempty" name:"OpType"`

	// Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 分类
	Category *string `json:"Category,omitempty" name:"Category"`
}

type MetricOut struct {

	// Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// namespace的Id
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// 中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CnName *string `json:"CnName,omitempty" name:"CnName"`

	// 聚合方式
	AggType *string `json:"AggType,omitempty" name:"AggType"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 创建人
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 创建的时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新人
	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`

	// 更新的时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type ModifyFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 收藏夹名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收藏夹描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyFavoriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFavoriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFavoriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFavoriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFavoriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 老的namespace名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// namespace描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据最大保存时间，单位天
	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`

	// 数据最大存储容量，单位GB
	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`

	// 聚合周期
	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`

	// 维度列表
	Dimensions []*DimensionIn `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 指标列表
	Metrics []*MetricIn `json:"Metrics,omitempty" name:"Metrics" list`

	// 新的namespace名称
	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

func (r *ModifyNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 告警策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 告警接收渠道
	ReceiveChannels []*string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels" list`

	// 触发条件
	PolicyTriggers []*PolicyTriggerIn `json:"PolicyTriggers,omitempty" name:"PolicyTriggers" list`

	// 告警策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in”
	FilterConditions []*string `json:"FilterConditions,omitempty" name:"FilterConditions" list`

	// 分组字段
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`

	// 告警接收用户的ID
	ReceiveUsers []*uint64 `json:"ReceiveUsers,omitempty" name:"ReceiveUsers" list`

	// 告警接收用户组的ID
	ReceiveGroups []*uint64 `json:"ReceiveGroups,omitempty" name:"ReceiveGroups" list`

	// 告警等级
	AlertLevel *string `json:"AlertLevel,omitempty" name:"AlertLevel"`
}

func (r *ModifyPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTkePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 告警策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警数据所属命名空间。支持k8s_cluster、k8s_pod、k8s_node
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 告警接收渠道
	ReceiveChannels []*string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels" list`

	// 触发条件
	PolicyTriggers []*PolicyTriggerIn `json:"PolicyTriggers,omitempty" name:"PolicyTriggers" list`

	// 告警策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in”
	FilterConditions []*string `json:"FilterConditions,omitempty" name:"FilterConditions" list`

	// 告警接收用户的ID
	ReceiveUsers []*uint64 `json:"ReceiveUsers,omitempty" name:"ReceiveUsers" list`

	// 告警接收用户组的ID
	ReceiveGroups []*uint64 `json:"ReceiveGroups,omitempty" name:"ReceiveGroups" list`
}

func (r *ModifyTkePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTkePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTkePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTkePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTkePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 视图名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视图描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NamespaceOut struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据最大保存时间，单位天
	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`

	// 数据最大存储容量，单位GB
	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`

	// 聚合周期
	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 创建人
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 创建的时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新人
	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`

	// 更新的时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否有数据
	HasData *bool `json:"HasData,omitempty" name:"HasData"`

	// 磁盘已使用容量，单位MB
	DataDiskUsage *uint64 `json:"DataDiskUsage,omitempty" name:"DataDiskUsage"`
}

type ParamOut struct {

	// 选项值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type PolicyOut struct {

	// Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 告警策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 过滤条件，逗号分隔
	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`

	// 分组字段，逗号分隔
	GroupBy *string `json:"GroupBy,omitempty" name:"GroupBy"`

	// 告警接收用户，逗号分隔
	ReceiveUsers *string `json:"ReceiveUsers,omitempty" name:"ReceiveUsers"`

	// 告警接收用户组，逗号分隔
	ReceiveGroups *string `json:"ReceiveGroups,omitempty" name:"ReceiveGroups"`

	// 告警接收渠道，逗号分隔
	ReceiveChannels *string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 创建人
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 创建的时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新人
	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`

	// 更新的时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// namespace名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 状态。正常ok，异常error，失效disable
	Status *string `json:"Status,omitempty" name:"Status"`

	// 主账号Uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 告警延迟等待时间，单位秒
	Delay *uint64 `json:"Delay,omitempty" name:"Delay"`

	// 需要附加到告警记录的标签
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 数据源类型，当前有argus、argus_tke
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// namespace在amp中的productName
	AmpProductName *string `json:"AmpProductName,omitempty" name:"AmpProductName"`

	// 告警策略触发条件
	PolicyTriggers []*PolicyTriggerOut `json:"PolicyTriggers,omitempty" name:"PolicyTriggers" list`

	// tke告警策略的集群Id，argus暂未用到
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceObj *string `json:"DataSourceObj,omitempty" name:"DataSourceObj"`

	// 告警等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertLevel *string `json:"AlertLevel,omitempty" name:"AlertLevel"`
}

type PolicyTriggerIn struct {

	// 检测类型。threshold表示阈值检测
	Type *string `json:"Type,omitempty" name:"Type"`

	// 指标名称
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 比较符
	CompareSymbol *string `json:"CompareSymbol,omitempty" name:"CompareSymbol"`

	// 阈值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 连续触发阈值的周期次数
	ContinueCount *uint64 `json:"ContinueCount,omitempty" name:"ContinueCount"`

	// 告警频率，单位分钟
	AlertFrequency *uint64 `json:"AlertFrequency,omitempty" name:"AlertFrequency"`

	// 触发条件的Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 操作类型。支持add,delete,update,noChange
	OpType *string `json:"OpType,omitempty" name:"OpType"`

	// 告警统计周期
	AggPeriodNum *uint64 `json:"AggPeriodNum,omitempty" name:"AggPeriodNum"`
}

type PolicyTriggerOut struct {

	// Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 告警策略Id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 检测类型。threshold表示阈值检测
	Type *string `json:"Type,omitempty" name:"Type"`

	// 指标名称
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 比较符
	CompareSymbol *string `json:"CompareSymbol,omitempty" name:"CompareSymbol"`

	// 阈值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 创建人
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 创建的时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新人
	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`

	// 更新的时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 连续触发阈值的周期次数
	ContinueCount *uint64 `json:"ContinueCount,omitempty" name:"ContinueCount"`

	// 告警频率，单位分钟
	AlertFrequency *uint64 `json:"AlertFrequency,omitempty" name:"AlertFrequency"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 主账号Uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 告警统计周期
	AggPeriodNum *uint64 `json:"AggPeriodNum,omitempty" name:"AggPeriodNum"`

	// 指标中文名
	MetricCnName *string `json:"MetricCnName,omitempty" name:"MetricCnName"`

	// 指标单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type ReportDataRequest struct {
	*tchttp.BaseRequest

	// 指定上报的数据格式，其中lp代表Line Protocol，json代表Json数据协议，推荐使用Line Protocol。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当Type指定为lp时使用，数据格式说明请参考：http://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/ ，其中measurement部分填写上报的命名空间名称，时间戳精确到纳秒。
	LpData []*string `json:"LpData,omitempty" name:"LpData" list`

	// 当Type指定为json时使用，支持单个请求上报多个命名空间数据。
	JsonData []*ReportJson `json:"JsonData,omitempty" name:"JsonData" list`
}

func (r *ReportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回描述信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportDimension struct {

	// 维度名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ReportJson struct {

	// 命名空间名称
	Ns *string `json:"Ns,omitempty" name:"Ns"`

	// 维度组合信息
	Dimensions []*ReportDimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 指标信息
	Metrics []*ReportMetric `json:"Metrics,omitempty" name:"Metrics" list`

	// 时间戳（精确到毫秒）
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type ReportMetric struct {

	// 指标名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type ServiceCallRelationRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 拉取数据的namespace列表
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces" list`

	// 开始时间，毫秒时间戳。默认值1小时前
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，毫秒时间戳。默认值当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 服务网格Id
	MeshId *string `json:"MeshId,omitempty" name:"MeshId"`

	// 拓扑粒度。支持service、workload、pod
	TopoType *string `json:"TopoType,omitempty" name:"TopoType"`
}

func (r *ServiceCallRelationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ServiceCallRelationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ServiceCallRelationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点。需要单独对Nodes部分进行一次JSON解码
		Nodes *string `json:"Nodes,omitempty" name:"Nodes"`

		// 边
		Edges []*MeshEdgeOut `json:"Edges,omitempty" name:"Edges" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServiceCallRelationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ServiceCallRelationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TkeAlertMetricOut struct {

	// TKE数据命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标英文名
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// 指标中文名
	CnName *string `json:"CnName,omitempty" name:"CnName"`

	// 指标描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 指标单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 是否默认选中。0否，1是
	Selected *string `json:"Selected,omitempty" name:"Selected"`
}

type TkeClusterOut struct {

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 告警策略数量
	PolicyCount *uint64 `json:"PolicyCount,omitempty" name:"PolicyCount"`

	// 集群健康状态。0正常，1异常
	Health *uint64 `json:"Health,omitempty" name:"Health"`
}

type UpdateFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 收藏夹Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 收藏夹名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收藏夹描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateFavoriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFavoriteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFavoriteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFavoriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFavoriteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 老的namespace名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// namespace描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据最大保存时间，单位天
	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`

	// 数据最大存储容量，单位GB
	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`

	// 聚合周期
	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`

	// 维度列表
	Dimensions []*DimensionIn `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 指标列表
	Metrics []*MetricIn `json:"Metrics,omitempty" name:"Metrics" list`

	// 新的namespace名称
	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

func (r *UpdateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 老的告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 新的告警策略名称
	NewName *string `json:"NewName,omitempty" name:"NewName"`

	// 告警策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// NamespaceId
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 告警接收渠道
	ReceiveChannels []*string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels" list`

	// 触发条件
	PolicyTriggers []*PolicyTriggerIn `json:"PolicyTriggers,omitempty" name:"PolicyTriggers" list`

	// 告警策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[["name","=","zhangsan"],["city","in",["shenzhen","beijing"]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in”
	FilterConditions []*string `json:"FilterConditions,omitempty" name:"FilterConditions" list`

	// 分组字段
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys" list`

	// 告警接收用户的ID
	ReceiveUsers []*uint64 `json:"ReceiveUsers,omitempty" name:"ReceiveUsers" list`

	// 告警接收用户组的ID
	ReceiveGroups []*uint64 `json:"ReceiveGroups,omitempty" name:"ReceiveGroups" list`

	// 字段名为中文名，其他与FilterConditions相同
	FilterConditionCns []*string `json:"FilterConditionCns,omitempty" name:"FilterConditionCns" list`
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

type UpdateViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 视图名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视图描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ViewOut struct {

	// 视图Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 所属收藏夹Id
	FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`

	// 所属命名空间Id
	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 视图名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视图参数
	Param *string `json:"Param,omitempty" name:"Param"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 创建人
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 创建的时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新人
	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`

	// 更新的时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 视图描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 所属命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 指标字段
	Metric *string `json:"Metric,omitempty" name:"Metric"`
}
