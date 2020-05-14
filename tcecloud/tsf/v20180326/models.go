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

package v20180326

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联密钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`
}

func (r *AddClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加云主机的返回列表
		Result *AddInstanceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddInstanceRequest struct {
	*tchttp.BaseRequest

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 机器Id列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像Id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联秘钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`
}

func (r *AddInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加机器是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddInstanceResult struct {

	// 添加集群失败的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds" list`

	// 添加集群成功的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds" list`

	// 添加集群超时的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutInstanceIds []*string `json:"TimeoutInstanceIds,omitempty" name:"TimeoutInstanceIds" list`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联密钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`
}

func (r *AddInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加云主机是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AlarmOverviewContent struct {

	// 展示给前端的告警记录的指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 该告警指标24小时内，已恢复的告警条数
	RestoredNums *int64 `json:"RestoredNums,omitempty" name:"RestoredNums"`

	// 该告警指标24小时内，未恢复的告警条数
	UnRestoredNums *int64 `json:"UnRestoredNums,omitempty" name:"UnRestoredNums"`
}

type AlarmPolicyResult struct {

	// PolicyId
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// PolicyName
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// EventPolicies
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventPolicies []*EventPolicyResult `json:"EventPolicies,omitempty" name:"EventPolicies" list`

	// EnabledEmail
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnabledEmail *int64 `json:"EnabledEmail,omitempty" name:"EnabledEmail"`

	// EnabledSMS
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnabledSMS *int64 `json:"EnabledSMS,omitempty" name:"EnabledSMS"`

	// Enabled
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`

	// Receivers
	// 注意：此字段可能返回 null，表示取不到有效值。
	Receivers []*AlarmReceiverResult `json:"Receivers,omitempty" name:"Receivers" list`

	// AlarmSubId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmSubId *int64 `json:"AlarmSubId,omitempty" name:"AlarmSubId"`

	// UpdateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// EnabledWeChat
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnabledWeChat *string `json:"EnabledWeChat,omitempty" name:"EnabledWeChat"`

	// EnabledRtx
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnabledRtx *string `json:"EnabledRtx,omitempty" name:"EnabledRtx"`
}

type AlarmReceiverResult struct {

	// PolicyId
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// ReceiverId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverId *string `json:"ReceiverId,omitempty" name:"ReceiverId"`

	// Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// CellPhoneNumber
	// 注意：此字段可能返回 null，表示取不到有效值。
	CellPhoneNumber *string `json:"CellPhoneNumber,omitempty" name:"CellPhoneNumber"`

	// Email
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitempty" name:"Email"`
}

type AnalyzeLogSchemaRequest struct {
	*tchttp.BaseRequest

	// 解析规则类型
	SchemaType *uint64 `json:"SchemaType,omitempty" name:"SchemaType"`

	// 解析规则内容，包含解析规则时间格式和解析规则内容
	SchemaPatternLayout *string `json:"SchemaPatternLayout,omitempty" name:"SchemaPatternLayout"`

	// 模拟解析的日志内容
	SchemaLogMessage *string `json:"SchemaLogMessage,omitempty" name:"SchemaLogMessage"`
}

func (r *AnalyzeLogSchemaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AnalyzeLogSchemaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AnalyzeLogSchemaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 解析结果
		Result *BusinessLogPatternAnalysis `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AnalyzeLogSchemaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AnalyzeLogSchemaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApiDefinitionDescr struct {

	// 对象名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 对象属性列表
	Properties []*PropertyField `json:"Properties,omitempty" name:"Properties" list`
}

type ApiDetailInfo struct {

	// API ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// API 请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// Api 映射路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathMapping *string `json:"PathMapping,omitempty" name:"PathMapping"`

	// 请求方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 所属分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 是否禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// 发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`

	// 开启限流
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitStatus *string `json:"RateLimitStatus,omitempty" name:"RateLimitStatus"`

	// 是否开启mock
	// 注意：此字段可能返回 null，表示取不到有效值。
	MockStatus *string `json:"MockStatus,omitempty" name:"MockStatus"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleasedTime *string `json:"ReleasedTime,omitempty" name:"ReleasedTime"`

	// 所属分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type ApiDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API 请求参数
		Request []*ApiRequestDescr `json:"Request,omitempty" name:"Request" list`

		// API 响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Response []*ApiResponseDescr `json:"Response,omitempty" name:"Response" list`

		// API 复杂结构定义
		Definitions []*ApiDefinitionDescr `json:"Definitions,omitempty" name:"Definitions" list`

		// API 的 content type
	// 注意：此字段可能返回 null，表示取不到有效值。
		RequestContentType *string `json:"RequestContentType,omitempty" name:"RequestContentType"`

		// API  能否调试
	// 注意：此字段可能返回 null，表示取不到有效值。
		CanRun *bool `json:"CanRun,omitempty" name:"CanRun"`
	} `json:"Response"`
}

func (r *ApiDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApiDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApiGroupInfo struct {

	// Api Group Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api Group 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组上下文
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`

	// 鉴权类型。 secret： 秘钥鉴权； none:无鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 发布状态, drafted: 未发布。 released: 发布
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分组创建时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 分组更新时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// api分组已绑定的网关部署组
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindedGatewayDeployGroups []*GatewayDeployGroup `json:"BindedGatewayDeployGroups,omitempty" name:"BindedGatewayDeployGroups" list`

	// api 个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiCount *int64 `json:"ApiCount,omitempty" name:"ApiCount"`

	// 访问group的ACL类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclMode *string `json:"AclMode,omitempty" name:"AclMode"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ApiInfo struct {

	// 命空间Id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 服务Id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// API path
	Path *string `json:"Path,omitempty" name:"Path"`

	// Api 请求
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求映射
	PathMapping *string `json:"PathMapping,omitempty" name:"PathMapping"`
}

type ApiRateLimitRule struct {

	// rule Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// API ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 限流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 最大限流qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 生效/禁用, enabled/disabled
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// 规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleContent *string `json:"RuleContent,omitempty" name:"RuleContent"`

	// Tsf Rule ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRuleId *string `json:"TsfRuleId,omitempty" name:"TsfRuleId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}

type ApiRequestDescr struct {

	// 参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数位置
	In *string `json:"In,omitempty" name:"In"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数是否必须
	Required *bool `json:"Required,omitempty" name:"Required"`

	// 参数的默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
}

type ApiResponseDescr struct {

	// 参数描述
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ApiUseStatisticsEntity struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 次数
	Count *string `json:"Count,omitempty" name:"Count"`

	// 比率
	Ratio *string `json:"Ratio,omitempty" name:"Ratio"`
}

type ApiVersionArray struct {

	// App ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// App 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// App 包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`
}

type AppPkgInfo struct {

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用下的程序包数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgCount *int64 `json:"PkgCount,omitempty" name:"PkgCount"`
}

type AppPkgList struct {

	// 程序包总量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 应用的包列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*AppPkgInfo `json:"Content,omitempty" name:"Content" list`
}

type ApplicationApiAccess struct {

	// 发布状态: SUCCEEDED/FAILED
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 系统给该服务自动分配的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// ApiGateway 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 协议类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type ApplicationAttribute struct {

	// 总实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 运行实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 应用下部署组个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCount *int64 `json:"GroupCount,omitempty" name:"GroupCount"`
}

type ApplicationForPage struct {

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 编程语言
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgLang *string `json:"ProgLang,omitempty" name:"ProgLang"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 应用资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// 应用runtime类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`
}

type ApplicationServerLogContent struct {

	// 日志内容
	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 日志类型
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

type ApplicationServerResult struct {

	// ServerId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerId *string `json:"ServerId,omitempty" name:"ServerId"`

	// ClusterId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// GroupId
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Type
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// TaskStatus
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// AgentStatus
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentStatus *int64 `json:"AgentStatus,omitempty" name:"AgentStatus"`

	// MasterId
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterId *int64 `json:"MasterId,omitempty" name:"MasterId"`

	// GroupDesc
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// BasicEnvDesc
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicEnvDesc *string `json:"BasicEnvDesc,omitempty" name:"BasicEnvDesc"`

	// ApplicationType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// AgentVersion
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
}

type AssociateBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// TSF分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 日志配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList" list`
}

func (r *AssociateBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuthCondition struct {

	// 类型，meta 或者 tag；条件为服务名时用 meta，其他用 tag
	Type *string `json:"Type,omitempty" name:"Type"`

	// 条件为服务名时，key 为 sourceServiceId；条件是 tag 时，key 为用户自定义
	Key *string `json:"Key,omitempty" name:"Key"`

	// 操作符，从中选其一，equal, notEqual, in, notIn, regex
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 如果操作符不是 in / notIn， 值只有一个
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueList []*string `json:"ValueList,omitempty" name:"ValueList" list`

	// 如果条件是服务名，在回包中这个数组中包含主调服务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServiceList []*AuthMicroservice `json:"SourceServiceList,omitempty" name:"SourceServiceList" list`
}

type AuthConditionV2 struct {

	// 标签类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagType *string `json:"TagType,omitempty" name:"TagType"`

	// 标签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagField *string `json:"TagField,omitempty" name:"TagField"`

	// 标签运算符
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagOperator *string `json:"TagOperator,omitempty" name:"TagOperator"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *int64 `json:"TagId,omitempty" name:"TagId"`
}

type AuthMicroservice struct {

	// 服务 ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 服务名
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 命名空间 ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 服务是否存在
	IsNotExist *bool `json:"IsNotExist,omitempty" name:"IsNotExist"`

	// 帐号 appid
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 账号 uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 子账号 uin
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type AuthRule struct {

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 是否启用
	IsEnabled *string `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 标签列表
	Tags []*AuthConditionV2 `json:"Tags,omitempty" name:"Tags" list`
}

type AuthRuleGroup struct {

	// 规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*AuthRule `json:"Rules,omitempty" name:"Rules" list`

	// 规则列表计算逻辑
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleProgram *string `json:"RuleProgram,omitempty" name:"RuleProgram"`

	// 权限类型，D：未启用；B：黑名单模式；W：白名单模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 微服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

type AutoRetryTransactionRequest struct {
	*tchttp.BaseRequest

	// 自动重试事务ID列表
	TransactionId []*string `json:"TransactionId,omitempty" name:"TransactionId" list`

	// 是否开启自动重试
	AutoRetry *bool `json:"AutoRetry,omitempty" name:"AutoRetry"`
}

func (r *AutoRetryTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AutoRetryTransactionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AutoRetryTransactionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事务重试结果
		Result *TxRetry `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AutoRetryTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AutoRetryTransactionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindApiGroupRequest struct {
	*tchttp.BaseRequest

	// 分组绑定网关列表
	GroupGatewayList []*GatewayGroupIds `json:"GroupGatewayList,omitempty" name:"GroupGatewayList" list`
}

func (r *BindApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindApiGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果，成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindApiGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindPluginRequest struct {
	*tchttp.BaseRequest

	// 分组/API绑定插件列表
	PluginInstanceList []*GatewayPluginBoundParam `json:"PluginInstanceList,omitempty" name:"PluginInstanceList" list`
}

func (r *BindPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果，成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BusinesLogConfigAssociatedGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组所属应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组所属应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 部署组所属应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组所属命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组所属命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 部署组所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组所属集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 部署组所属集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 部署组关联日志配置时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedTime *string `json:"AssociatedTime,omitempty" name:"AssociatedTime"`
}

type BusinessLogConfig struct {

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项日志路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`

	// 配置项描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigDesc *string `json:"ConfigDesc,omitempty" name:"ConfigDesc"`

	// 配置项标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTags *string `json:"ConfigTags,omitempty" name:"ConfigTags"`

	// 配置项对应的ES管道
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigPipeline *string `json:"ConfigPipeline,omitempty" name:"ConfigPipeline"`

	// 配置项创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigCreateTime *string `json:"ConfigCreateTime,omitempty" name:"ConfigCreateTime"`

	// 配置项更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigUpdateTime *string `json:"ConfigUpdateTime,omitempty" name:"ConfigUpdateTime"`

	// 配置项解析规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigSchema *BusinessLogConfigSchema `json:"ConfigSchema,omitempty" name:"ConfigSchema"`

	// 配置项关联部署组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigAssociatedGroups []*BusinesLogConfigAssociatedGroup `json:"ConfigAssociatedGroups,omitempty" name:"ConfigAssociatedGroups" list`
}

type BusinessLogConfigSchema struct {

	// 解析规则类型
	SchemaType *int64 `json:"SchemaType,omitempty" name:"SchemaType"`

	// 解析规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaContent *string `json:"SchemaContent,omitempty" name:"SchemaContent"`

	// 解析规则时间格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaDateFormat *string `json:"SchemaDateFormat,omitempty" name:"SchemaDateFormat"`

	// 解析规则对应的多行匹配规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaMultilinePattern *string `json:"SchemaMultilinePattern,omitempty" name:"SchemaMultilinePattern"`

	// 解析规则创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaCreateTime *string `json:"SchemaCreateTime,omitempty" name:"SchemaCreateTime"`

	// 用户填写的解析规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaPatternLayout *string `json:"SchemaPatternLayout,omitempty" name:"SchemaPatternLayout"`
}

type BusinessLogPatternAnalysis struct {

	// 校验/解析是否成功
	Success *bool `json:"Success,omitempty" name:"Success"`

	// 校验/解析用户提示，包含校验不通过的原因，是否成功联动调用链等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hint *string `json:"Hint,omitempty" name:"Hint"`

	// 解析结果，复用Envs对象，本质是一个Map<String, String>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalyzeResult []*EnvsV2 `json:"AnalyzeResult,omitempty" name:"AnalyzeResult" list`
}

type BusinessLogV2 struct {

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 日志时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIp *string `json:"InstanceIp,omitempty" name:"InstanceIp"`

	// 日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *string `json:"LogId,omitempty" name:"LogId"`
}

type ChangeApiUsableStatusRequest struct {
	*tchttp.BaseRequest

	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 切换状态，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`
}

func (r *ChangeApiUsableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeApiUsableStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeApiUsableStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API 信息
		Result *ApiDetailInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeApiUsableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeApiUsableStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeContainerReplicasRequest struct {
	*tchttp.BaseRequest

	// groupId，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

func (r *ChangeContainerReplicasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeContainerReplicasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeContainerReplicasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeContainerReplicasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeContainerReplicasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeSecretStatusRequest struct {
	*tchttp.BaseRequest

	// 秘钥ID
	GwSecretId *string `json:"GwSecretId,omitempty" name:"GwSecretId"`

	// 启用: enabled/ 禁用: disabled
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ChangeSecretStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeSecretStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeSecretStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *SecretKeyInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeSecretStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeSecretStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckClusterCIDRRequest struct {
	*tchttp.BaseRequest

	// 集群CIDR
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 虚拟网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *CheckClusterCIDRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckClusterCIDRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckClusterCIDRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：可用；false：不可用
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterCIDRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckClusterCIDRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CircuitBreakerApi struct {

	// API PATH
	Path *string `json:"Path,omitempty" name:"Path"`

	// API Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 熔断策略ID
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

type CircuitBreakerRule struct {

	// 熔断规则微服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务所属命名空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 目标服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServiceName *string `json:"TargetServiceName,omitempty" name:"TargetServiceName"`

	// 目标服务所属命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// 熔断策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyList []*CircuitBreakerStrategy `json:"StrategyList,omitempty" name:"StrategyList" list`

	// 熔断级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolationLevel *string `json:"IsolationLevel,omitempty" name:"IsolationLevel"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 目标服务所属命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" name:"TargetNamespaceName"`

	// 熔断规则主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 是否开启此规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// APPID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// subAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type CircuitBreakerRules struct {

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 路由规则
	Content []*CircuitBreakerRule `json:"Content,omitempty" name:"Content" list`
}

type CircuitBreakerStrategy struct {

	// 最少请求数
	MinimumNumberOfCalls *int64 `json:"MinimumNumberOfCalls,omitempty" name:"MinimumNumberOfCalls"`

	// 滚动窗口统计时间
	SlidingWindowSize *int64 `json:"SlidingWindowSize,omitempty" name:"SlidingWindowSize"`

	// 熔断开启到半开间隔,单位s
	WaitDurationInOpenState *int64 `json:"WaitDurationInOpenState,omitempty" name:"WaitDurationInOpenState"`

	// 熔断策略ID
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 熔断策略所属熔断规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 失败请求比例
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureRateThreshold *int64 `json:"FailureRateThreshold,omitempty" name:"FailureRateThreshold"`

	// 最大熔断实例的比例
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxEjectionPercent *int64 `json:"MaxEjectionPercent,omitempty" name:"MaxEjectionPercent"`

	// 慢请求阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowCallDurationThreshold *int64 `json:"SlowCallDurationThreshold,omitempty" name:"SlowCallDurationThreshold"`

	// 慢请求比例
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowCallRateThreshold *int64 `json:"SlowCallRateThreshold,omitempty" name:"SlowCallRateThreshold"`

	// 熔断策略作用API
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiList []*CircuitBreakerApi `json:"ApiList,omitempty" name:"ApiList" list`
}

type ClearResourceDataRequest struct {
	*tchttp.BaseRequest

	// 资源类型，支持以下几个值，"ES"，"CTSDB"，"CONSUL"，"REDIS"
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 开始清理的时间点。0：清理全部，1：清理当天的数据，7：清理一周前到今天的数据。
	ClearedTime *uint64 `json:"ClearedTime,omitempty" name:"ClearedTime"`
}

func (r *ClearResourceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearResourceDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearResourceDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearResourceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearResourceDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloudMonitorMicroservice struct {

	// namespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// namespaceName
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// microserviceId
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// microserviceName
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`
}

type CloudMonitorMicroserviceResult struct {

	// TotalCount
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	Content []*CloudMonitorMicroservice `json:"Content,omitempty" name:"Content" list`
}

type CloudMonitorPolicies struct {

	// keywordsId
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeywordsId *string `json:"KeywordsId,omitempty" name:"KeywordsId"`

	// keyWords
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// groupId
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// groupName
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// namespaceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// namespaceName
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// applicationId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// applicationName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// clusterId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// clusterName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type CloudMonitorPolicyResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*CloudMonitorPolicies `json:"Content,omitempty" name:"Content" list`
}

type Cluster struct {

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群所属私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群总CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalCpu *float64 `json:"ClusterTotalCpu,omitempty" name:"ClusterTotalCpu"`

	// 集群总内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalMem *float64 `json:"ClusterTotalMem,omitempty" name:"ClusterTotalMem"`

	// 集群已使用CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedCpu *float64 `json:"ClusterUsedCpu,omitempty" name:"ClusterUsedCpu"`

	// 集群已使用内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedMem *float64 `json:"ClusterUsedMem,omitempty" name:"ClusterUsedMem"`

	// 集群机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 集群可用的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 集群正常状态的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalInstanceCount *int64 `json:"NormalInstanceCount,omitempty" name:"NormalInstanceCount"`

	// 删除标记：true：可以删除；false：不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群所属TSF地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionName *string `json:"TsfRegionName,omitempty" name:"TsfRegionName"`

	// 集群所属TSF可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 集群所属TSF可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneName *string `json:"TsfZoneName,omitempty" name:"TsfZoneName"`

	// 集群不可删除的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlagReason *string `json:"DeleteFlagReason,omitempty" name:"DeleteFlagReason"`

	// 集群最大CPU限制，单位：核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterLimitCpu *float64 `json:"ClusterLimitCpu,omitempty" name:"ClusterLimitCpu"`

	// 集群最大内存限制，单位：G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterLimitMem *float64 `json:"ClusterLimitMem,omitempty" name:"ClusterLimitMem"`

	// 集群可用的服务实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunServiceInstanceCount *int64 `json:"RunServiceInstanceCount,omitempty" name:"RunServiceInstanceCount"`
}

type ClusterLimitResourceV2 struct {

	// 最大分配 CPU 核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配内存 MiB 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 初始分配 CPU 核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配内存 MiB 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`
}

type ClusterV2 struct {

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群所属私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群总CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalCpu *float64 `json:"ClusterTotalCpu,omitempty" name:"ClusterTotalCpu"`

	// 集群总内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalMem *float64 `json:"ClusterTotalMem,omitempty" name:"ClusterTotalMem"`

	// 集群已使用CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedCpu *float64 `json:"ClusterUsedCpu,omitempty" name:"ClusterUsedCpu"`

	// 集群已使用内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedMem *float64 `json:"ClusterUsedMem,omitempty" name:"ClusterUsedMem"`

	// 集群机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 集群运行中的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 集群正常状态的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalInstanceCount *int64 `json:"NormalInstanceCount,omitempty" name:"NormalInstanceCount"`

	// 删除标记：true：可以删除；false：不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群所属TSF地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionName *string `json:"TsfRegionName,omitempty" name:"TsfRegionName"`

	// 集群所属TSF可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 集群所属TSF可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneName *string `json:"TsfZoneName,omitempty" name:"TsfZoneName"`

	// 集群不可删除的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlagReason *string `json:"DeleteFlagReason,omitempty" name:"DeleteFlagReason"`

	// 集群所属私有网络子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type CommonPkg struct {

	// 公共包名
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 包所属的应用ID，在微服务网关相关应用中使用
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 包的描述信息
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 回调接口
	CallbackAction *string `json:"CallbackAction,omitempty" name:"CallbackAction"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type Config struct {

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项版本描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 配置项类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 删除标识，true：可以删除；false：不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 配置项版本数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionCount *int64 `json:"ConfigVersionCount,omitempty" name:"ConfigVersionCount"`
}

type ConfigRelease struct {

	// 配置项发布ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

type ConfigReleaseLog struct {

	// 配置项发布日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`

	// 上次发布的配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigId *string `json:"LastConfigId,omitempty" name:"LastConfigId"`

	// 上次发布的配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigName *string `json:"LastConfigName,omitempty" name:"LastConfigName"`

	// 上次发布的配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigVersion *string `json:"LastConfigVersion,omitempty" name:"LastConfigVersion"`

	// 回滚标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollbackFlag *bool `json:"RollbackFlag,omitempty" name:"RollbackFlag"`
}

type ConfigTemplate struct {

	// 配置模板Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`

	// 配置模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateDesc *string `json:"ConfigTemplateDesc,omitempty" name:"ConfigTemplateDesc"`

	// 配置模板对应的微服务框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 配置模板数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateValue *string `json:"ConfigTemplateValue,omitempty" name:"ConfigTemplateValue"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ConfigTemplateResult struct {

	// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ConfigTemplate `json:"Content,omitempty" name:"Content" list`
}

type Consumer struct {

	// ConsumerControllerName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerControllerName *string `json:"ConsumerControllerName,omitempty" name:"ConsumerControllerName"`

	// ToProviderName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToProviderName *string `json:"ToProviderName,omitempty" name:"ToProviderName"`

	// ToProviderServiceName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToProviderServiceName *string `json:"ToProviderServiceName,omitempty" name:"ToProviderServiceName"`
}

type ContainGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`
}

type ContainGroupResult struct {

	// 部署组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ContainGroup `json:"Content,omitempty" name:"Content" list`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ContainerAdditionalResourceRequirement struct {

	// CPU 核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`

	// 内存 MiB 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mem *string `json:"Mem,omitempty" name:"Mem"`
}

type ContainerAdditionalResourceRequirementMap struct {

	// Mesh 应用部署时需要的额外资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	M *ContainerAdditionalResourceRequirement `json:"M,omitempty" name:"M"`

	// 普通应用部署时需要的额外资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	N *ContainerAdditionalResourceRequirement `json:"N,omitempty" name:"N"`
}

type ContainerGroupDeploy struct {

	// groupId
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 0:公网 1:集群内访问 2：NodePort
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// ProtocolPorts
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 更新方式：0:快速更新 1:滚动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// jvm参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`
}

type ContainerGroupDetail struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 负载均衡ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// Service ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// NodePort端口，只有公网和NodePort访问方式才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePort *int64 `json:"NodePort,omitempty" name:"NodePort"`

	// 最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 0:公网 1:集群内访问 2：NodePort
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *uint64 `json:"AccessType,omitempty" name:"AccessType"`

	// 更新方式：0:快速更新 1:滚动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 端口数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 环境变量数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envs []*Env `json:"Envs,omitempty" name:"Envs" list`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// pod错误信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`
}

type ContainerGroupOther struct {

	// 实例总数
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动实例总数
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 负载均衡ip
	LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

	// Service ip
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// 服务状态，请参考后面的的状态定义
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务状态，请参考后面的的状态定义
	Message *string `json:"Message,omitempty" name:"Message"`

	// 环境变量
	Envs []*Env `json:"Envs,omitempty" name:"Envs" list`
}

type ContainerGroupResourceConfig struct {

	// 不同类型的应用的容器部署组，部署时的额外资源要求
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdditionalResourceRequirement *ContainerAdditionalResourceRequirementMap `json:"AdditionalResourceRequirement,omitempty" name:"AdditionalResourceRequirement"`
}

type ContainerInstanceResourceConfig struct {

	// 实例导入方式，可多个，公有云为 ["R"]，独立版的取值有 "M" 脚本模式、"S" SSH 模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportMode []*string `json:"ImportMode,omitempty" name:"ImportMode" list`

	// SSH 模式时，前端应该限制用户填这个数量的 master 主机信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterNumLimit *int64 `json:"MasterNumLimit,omitempty" name:"MasterNumLimit"`

	// SSH 模式时，前端应该限制用户填的最高数量的 node 主机信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeNumLimitPerSetup *int64 `json:"NodeNumLimitPerSetup,omitempty" name:"NodeNumLimitPerSetup"`
}

type ContainerTasks struct {

	// taskId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTime *string `json:"TaskTime,omitempty" name:"TaskTime"`

	// 任务类型字段，0：没有任务（在此接口中，不用出现0），1：发布程序包；2.部署操作；3.扩容操作；4.启动操作；5.停止操作；6.缩容操作；7.发布日志配置,8.删除销毁操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 变更状态，0：成功 1:失败 2：执行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 分组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 镜像名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

	// 任务详情描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`

	// 任务总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 成功任务个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 失败任务个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
}

type Controller struct {

	// ControllerName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ControllerName *string `json:"ControllerName,omitempty" name:"ControllerName"`
}

type CosCredentials struct {

	// 会话Token
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// 临时应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpAppId *string `json:"TmpAppId,omitempty" name:"TmpAppId"`

	// 临时调用者身份ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 所在域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type CosDownloadInfo struct {

	// 桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 鉴权信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Credentials *CosCredentials `json:"Credentials,omitempty" name:"Credentials"`
}

type CosUploadInfo struct {

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 目标地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 鉴权信息
	Credentials *CosCredentials `json:"Credentials,omitempty" name:"Credentials"`
}

type CreateAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 告警订阅策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 告警订阅策略对应的事件
	EventPolicies []*EventPolicyResult `json:"EventPolicies,omitempty" name:"EventPolicies" list`

	// 是否开通邮件告警 0 : off ; 1 : on;
	EnabledEmail *int64 `json:"EnabledEmail,omitempty" name:"EnabledEmail"`

	// 是否开通电话告警 0 : off ; 1 : on;
	EnabledSMS *int64 `json:"EnabledSMS,omitempty" name:"EnabledSMS"`

	// 是否开通rtx告警 0 : off ; 1 : on;
	EnabledRtx *int64 `json:"EnabledRtx,omitempty" name:"EnabledRtx"`

	// 冗余字段
	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`

	// 是否开通wechat告警 0 : off ; 1 : on;
	EnabledWeChat *int64 `json:"EnabledWeChat,omitempty" name:"EnabledWeChat"`

	// 告警接收人员列表
	Persons []*Person `json:"Persons,omitempty" name:"Persons" list`
}

func (r *CreateAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmReceiverRequest struct {
	*tchttp.BaseRequest

	// PolicyId
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// CellPhoneNumber
	CellPhoneNumber *string `json:"CellPhoneNumber,omitempty" name:"CellPhoneNumber"`

	// Email
	Email *string `json:"Email,omitempty" name:"Email"`
}

func (r *CreateAlarmReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmReceiverRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmReceiverResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAlarmReceiverResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAndDownloadTemplateRequest struct {
	*tchttp.BaseRequest

	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// BasePackage
	BasePackage *string `json:"BasePackage,omitempty" name:"BasePackage"`

	// PomGroupId
	PomGroupId *string `json:"PomGroupId,omitempty" name:"PomGroupId"`

	// PomArtifactId
	PomArtifactId *string `json:"PomArtifactId,omitempty" name:"PomArtifactId"`

	// PomVersion
	PomVersion *string `json:"PomVersion,omitempty" name:"PomVersion"`

	// PomName
	PomName *string `json:"PomName,omitempty" name:"PomName"`

	// PomDesc
	PomDesc *string `json:"PomDesc,omitempty" name:"PomDesc"`

	// GetMethod
	GetMethod *string `json:"GetMethod,omitempty" name:"GetMethod"`

	// 微服务数组
	Ms []*Ms `json:"Ms,omitempty" name:"Ms" list`

	// ProjectId
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateAndDownloadTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAndDownloadTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAndDownloadTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAndDownloadTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAndDownloadTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiAccessRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateApiAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SubDomain, 系统给该服务自动分配的域名
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiGroupRequest struct {
	*tchttp.BaseRequest

	// 分组名称, 不能包含中文
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组上下文
	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`

	// 鉴权类型。secret： 秘钥鉴权； none:无鉴权
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API分组ID
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiRateLimitRuleRequest struct {
	*tchttp.BaseRequest

	// Api Id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// qps值
	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 开启/禁用，enabled/disabled, 不传默认开启
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`
}

func (r *CreateApiRateLimitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiRateLimitRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiRateLimitRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiRateLimitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiRateLimitRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型，V：虚拟机应用；C：容器应用；S：serverless应用
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用微服务类型，M：service mesh应用；N：普通应用；G：网关应用
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用描述
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用日志配置项，废弃参数
	ApplicationLogConfig *string `json:"ApplicationLogConfig,omitempty" name:"ApplicationLogConfig"`

	// 应用资源类型，废弃参数
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// 应用runtime类型
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`
}

func (r *CreateApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用ID
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 是否启用：1启用；0未启用
	IsEnabled *string `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 标签列表
	Tags *AuthConditionV2 `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuthorizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：创建成功；false：创建失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuthorizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// 日志配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 日志配置项路径
	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`

	// 日志配置项解析规则
	ConfigSchema *BusinessLogConfigSchema `json:"ConfigSchema,omitempty" name:"ConfigSchema"`

	// 日志配置项描述
	ConfigDesc *string `json:"ConfigDesc,omitempty" name:"ConfigDesc"`

	// 日志配置项标签
	ConfigTags *string `json:"ConfigTags,omitempty" name:"ConfigTags"`
}

func (r *CreateBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCircuitBreakerRuleRequest struct {
	*tchttp.BaseRequest

	// 熔断规则微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务所属命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 目标服务名
	TargetServiceName *string `json:"TargetServiceName,omitempty" name:"TargetServiceName"`

	// 目标服务所属命名空间
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// 熔断策略
	StrategyList []*CircuitBreakerStrategy `json:"StrategyList,omitempty" name:"StrategyList" list`

	// 熔断级别
	IsolationLevel *string `json:"IsolationLevel,omitempty" name:"IsolationLevel"`

	// 微服务所属命名空间
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" name:"TargetNamespaceName"`

	// 熔断规则主键
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 是否开启此规则
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

func (r *CreateCircuitBreakerRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCircuitBreakerRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCircuitBreakerRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则ID
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCircuitBreakerRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCircuitBreakerRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCkafkaRequest struct {
	*tchttp.BaseRequest

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ckafka instanceid
	CkafkaInstanceid *string `json:"CkafkaInstanceid,omitempty" name:"CkafkaInstanceid"`

	// ckafka topicid
	CkafkaTopicid *string `json:"CkafkaTopicid,omitempty" name:"CkafkaTopicid"`
}

func (r *CreateCkafkaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCkafkaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCkafkaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCkafkaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCkafkaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 分配给集群容器和服务IP的CIDR
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群备注
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群所属TSF地域
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF可用区
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建集群操作是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项值
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项值类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *CreateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：创建成功；false：创建失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 配置模板名称
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板对应的微服务框架
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 配置模板数据
	ConfigTemplateValue *string `json:"ConfigTemplateValue,omitempty" name:"ConfigTemplateValue"`

	// 配置模板描述
	ConfigTemplateDesc *string `json:"ConfigTemplateDesc,omitempty" name:"ConfigTemplateDesc"`
}

func (r *CreateConfigTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConfigTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConfigTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：创建成功；false：创建失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConfigTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateContainGroupRequest struct {
	*tchttp.BaseRequest

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 分组所属命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 分组名称字段，长度1~60，字母或下划线开头，可包含字母数字下划线
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 数组对象，见下方定义
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最大分配 CPU 核数，对应 K8S limit
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配内存 MiB 数，对应 K8S limit
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 分组备注字段，长度应不大于200字符
	GroupComment *string `json:"GroupComment,omitempty" name:"GroupComment"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 初始分配的 CPU 核数，对应 K8S request
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`
}

func (r *CreateContainGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContainGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateContainGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回创建成功的部署组ID，返回null表示失败
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateContainGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContainGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFileConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项文件名
	ConfigFileName *string `json:"ConfigFileName,omitempty" name:"ConfigFileName"`

	// 配置项文件内容
	ConfigFileValue *string `json:"ConfigFileValue,omitempty" name:"ConfigFileValue"`

	// 配置项关联应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 发布路径
	ConfigFilePath *string `json:"ConfigFilePath,omitempty" name:"ConfigFilePath"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项文件编码
	ConfigFileCode *string `json:"ConfigFileCode,omitempty" name:"ConfigFileCode"`

	// 后置命令
	ConfigPostCmd *string `json:"ConfigPostCmd,omitempty" name:"ConfigPostCmd"`
}

func (r *CreateFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFileConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：创建成功；false：创建失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFileConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayApiRequest struct {
	*tchttp.BaseRequest

	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api信息
	ApiList []*ApiInfo `json:"ApiList,omitempty" name:"ApiList" list`
}

func (r *CreateGatewayApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGatewayApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGatewayApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayJwtPluginRequest struct {
	*tchttp.BaseRequest

	// 插件类型 "Jwt"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 公钥对kid
	Kid *string `json:"Kid,omitempty" name:"Kid"`

	// 公钥对Json串
	PublicKeyJson *string `json:"PublicKeyJson,omitempty" name:"PublicKeyJson"`

	// 重定向地址
	TokenBaggagePosition *string `json:"TokenBaggagePosition,omitempty" name:"TokenBaggagePosition"`

	// token的key值
	TokenKeyName *string `json:"TokenKeyName,omitempty" name:"TokenKeyName"`

	// claim参数映射关系json串
	ClaimMappingJson *string `json:"ClaimMappingJson,omitempty" name:"ClaimMappingJson"`

	// 发布状态: drafted/released
	Status *string `json:"Status,omitempty" name:"Status"`

	// 插件描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// token携带位置query/header
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

func (r *CreateGatewayJwtPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGatewayJwtPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayJwtPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网关插件
		Result *GatewayPluginId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayJwtPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGatewayJwtPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayOAuthPluginRequest struct {
	*tchttp.BaseRequest

	// 插件类型 "OAuth"
	Name *string `json:"Name,omitempty" name:"Name"`

	// 插件名称
	Type *string `json:"Type,omitempty" name:"Type"`

	// 验证token路径
	TokenAuthUrl *string `json:"TokenAuthUrl,omitempty" name:"TokenAuthUrl"`

	// 验证token请求方法:get/post
	TokenAuthMethod *string `json:"TokenAuthMethod,omitempty" name:"TokenAuthMethod"`

	// 认证请求超时时间,单位:秒 范围:0~30
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// token携带位置，网关取token位置与发送认证请求时token位置一致,值:query/header
	TokenBaggagePosition *string `json:"TokenBaggagePosition,omitempty" name:"TokenBaggagePosition"`

	// token的key值
	TokenKeyName *string `json:"TokenKeyName,omitempty" name:"TokenKeyName"`

	// payload的映射参数名称
	PayloadMappingName *string `json:"PayloadMappingName,omitempty" name:"PayloadMappingName"`

	// payload映射到后端服务的携带位置,值:query/header
	PayloadMappingPosition *string `json:"PayloadMappingPosition,omitempty" name:"PayloadMappingPosition"`

	// 发布状态: drafted/released
	Status *string `json:"Status,omitempty" name:"Status"`

	// 插件描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 重定向地址
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

func (r *CreateGatewayOAuthPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGatewayOAuthPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayOAuthPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网关插件
		Result *GatewayPluginId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayOAuthPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGatewayOAuthPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayTagPluginRequest struct {
	*tchttp.BaseRequest

	// 插件名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 插件类型 "Tag"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数配置的Json串
	TagPluginInfoList *string `json:"TagPluginInfoList,omitempty" name:"TagPluginInfoList"`

	// 发布状态: drafted/released
	Status *string `json:"Status,omitempty" name:"Status"`

	// 插件描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateGatewayTagPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGatewayTagPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayTagPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网关插件
		Result *GatewayPluginId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayTagPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGatewayTagPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组所属的应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组所属命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组描述
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`
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

		// groupId， null表示创建失败
		Result *string `json:"Result,omitempty" name:"Result"`

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

type CreateGroupSecretRequest struct {
	*tchttp.BaseRequest

	// Api 分组 ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 秘钥名称
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *CreateGroupSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupSecretRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 秘钥ID
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupSecretResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMachinesRequest struct {
	*tchttp.BaseRequest

	// MachineList
	MachineList []*MachineList `json:"MachineList,omitempty" name:"MachineList" list`
}

func (r *CreateMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMachinesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMachinesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

func (r *CreateMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMicroserviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新增微服务是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMicroserviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMonitorStatisticsPolicyRequest struct {
	*tchttp.BaseRequest

	// 关键词;
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 部署组id列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateMonitorStatisticsPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMonitorStatisticsPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMonitorStatisticsPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 保存结果true：保存成功；false：保存失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMonitorStatisticsPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMonitorStatisticsPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间描述
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 命名空间资源类型(默认值为DEF)
	NamespaceResourceType *string `json:"NamespaceResourceType,omitempty" name:"NamespaceResourceType"`

	// 是否是全局命名空间(默认是DEF，表示普通命名空间；GLOBAL表示全局命名空间)
	NamespaceType *string `json:"NamespaceType,omitempty" name:"NamespaceType"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
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

		// 成功时为命名空间ID，失败为null
		Result *string `json:"Result,omitempty" name:"Result"`

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

type CreatePolicyDocumentRequest struct {
	*tchttp.BaseRequest

	// 策略项列表
	PolicyItemList []*PolicyItem `json:"PolicyItemList,omitempty" name:"PolicyItemList" list`
}

func (r *CreatePolicyDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyDocumentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyDocumentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略文档
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyDocumentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProgramRequest struct {
	*tchttp.BaseRequest

	// 数据集名称
	ProgramName *string `json:"ProgramName,omitempty" name:"ProgramName"`

	// 数据集描述
	ProgramDesc *string `json:"ProgramDesc,omitempty" name:"ProgramDesc"`

	// 数据项列表，传入null或空数组时不新增
	ProgramItemList []*ProgramItem `json:"ProgramItemList,omitempty" name:"ProgramItemList" list`
}

func (r *CreateProgramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProgramRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProgramResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true: 创建成功；false: 创建失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProgramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProgramResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项值，总是接收yaml格式的内容
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
}

func (r *CreatePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：创建成功；false：创建失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRatelimitRequest struct {
	*tchttp.BaseRequest

	// 名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 限流所作用的目标微服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Rule列表，Rule是一个Object
	Rules []*RatelimitRuleV2 `json:"Rules,omitempty" name:"Rules" list`
}

func (r *CreateRatelimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRatelimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRatelimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功与否
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRatelimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRatelimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRegionRequest struct {
	*tchttp.BaseRequest

	// TRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`

	// TRegionName
	TRegionName *string `json:"TRegionName,omitempty" name:"TRegionName"`

	// TRemark
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`
}

func (r *CreateRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色描述
	RoleDesc *string `json:"RoleDesc,omitempty" name:"RoleDesc"`

	// 角色拥有的权限组ID列表，不传入或为null时不创建
	PermCatIdList []*string `json:"PermCatIdList,omitempty" name:"PermCatIdList" list`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true: 创建成功；false: 创建失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRouteRequest struct {
	*tchttp.BaseRequest

	// 路由所属微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 路由名称
	RouteName *string `json:"RouteName,omitempty" name:"RouteName"`

	// 路由规则列表
	RuleList []*RouteRuleV2 `json:"RuleList,omitempty" name:"RuleList" list`

	// 路由描述信息
	RouteDesc *string `json:"RouteDesc,omitempty" name:"RouteDesc"`
}

func (r *CreateRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由ID
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRouteRuleRequest struct {
	*tchttp.BaseRequest

	// 路由规则所属微服务
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 路由规则名称
	RouteRuleName *string `json:"RouteRuleName,omitempty" name:"RouteRuleName"`

	// 路由规则类型，包括标签路由和权重标签，标签路由： T , 权重路由： W
	RouteRuleType *string `json:"RouteRuleType,omitempty" name:"RouteRuleType"`

	// 路由规则描述信息
	RouteRuleDesc *string `json:"RouteRuleDesc,omitempty" name:"RouteRuleDesc"`

	// 标签路由规则项
	TagRouteItemList []*TagRouteItemList `json:"TagRouteItemList,omitempty" name:"TagRouteItemList" list`

	// 权重路由规则项
	WeightRouteItemList []*WeightRouteItemList `json:"WeightRouteItemList,omitempty" name:"WeightRouteItemList" list`
}

func (r *CreateRouteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRouteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRouteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由规则Id
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRouteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScalableRuleRequest struct {
	*tchttp.BaseRequest

	// 弹性扩缩容规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 冷却时间, s为单位, 0-9999
	CoolTime *uint64 `json:"CoolTime,omitempty" name:"CoolTime"`

	// 包含缩容规则, 默认0, 0:否 1:是
	EnableShrink *uint64 `json:"EnableShrink,omitempty" name:"EnableShrink"`

	// 包含扩容规则，默认0, 0否 1是
	EnableExpand *uint64 `json:"EnableExpand,omitempty" name:"EnableExpand"`

	// 扩容规则持续时间,0-9999s
	ExpandPeriod *uint64 `json:"ExpandPeriod,omitempty" name:"ExpandPeriod"`

	// 缩容规则持续时间，0-9999s
	ShrinkPeriod *uint64 `json:"ShrinkPeriod,omitempty" name:"ShrinkPeriod"`

	// 单次扩容机器数量, 0-9999
	ExpandVmCount *uint64 `json:"ExpandVmCount,omitempty" name:"ExpandVmCount"`

	// 单次缩容机器数量, 0-9999
	ShrinkVmCount *uint64 `json:"ShrinkVmCount,omitempty" name:"ShrinkVmCount"`

	// 扩容之后，最大实例数目,0-9999
	ExpandVmCountLimit *uint64 `json:"ExpandVmCountLimit,omitempty" name:"ExpandVmCountLimit"`

	// 缩容之后，最小实例数目, 0-9999
	ShrinkVmCountLimit *uint64 `json:"ShrinkVmCountLimit,omitempty" name:"ShrinkVmCountLimit"`

	// 扩容规则数组
	ExpandSubruleList []*ScalableSubRule `json:"ExpandSubruleList,omitempty" name:"ExpandSubruleList" list`

	// 缩容规则数组
	ShrinkSubruleList []*ScalableSubRule `json:"ShrinkSubruleList,omitempty" name:"ShrinkSubruleList" list`
}

func (r *CreateScalableRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScalableRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScalableRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求结果,true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScalableRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScalableRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 分组名称字段，长度1~60，字母或下划线开头，可包含字母数字下划线
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组所属名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 分组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功的部署组ID，返回null表示失败
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceInstanceRequest struct {
	*tchttp.BaseRequest

	// ModuleId
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// IpList
	IpList []*IpListResult `json:"IpList,omitempty" name:"IpList" list`

	// Ports
	Ports []*PortsResult `json:"Ports,omitempty" name:"Ports" list`
}

func (r *CreateServiceInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateRequest struct {
	*tchttp.BaseRequest

	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// BasePackage
	BasePackage *string `json:"BasePackage,omitempty" name:"BasePackage"`

	// PomGroupId
	PomGroupId *string `json:"PomGroupId,omitempty" name:"PomGroupId"`

	// PomArtifactId
	PomArtifactId *string `json:"PomArtifactId,omitempty" name:"PomArtifactId"`

	// PomVersion
	PomVersion *string `json:"PomVersion,omitempty" name:"PomVersion"`

	// PomName
	PomName *string `json:"PomName,omitempty" name:"PomName"`

	// PomDesc
	PomDesc *string `json:"PomDesc,omitempty" name:"PomDesc"`

	// GetMethod
	GetMethod *string `json:"GetMethod,omitempty" name:"GetMethod"`

	// 新建项目为null，原有项目为原ProjectId
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 微服务数组
	Ms []*Ms `json:"Ms,omitempty" name:"Ms" list`
}

func (r *CreateTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTsfZoneRequest struct {
	*tchttp.BaseRequest

	// tZoneId
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`

	// tZoneName
	TZoneName *string `json:"TZoneName,omitempty" name:"TZoneName"`

	// tRemark
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`

	// tRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`
}

func (r *CreateTsfZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTsfZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTsfZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTsfZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTsfZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest

	// tZoneId
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`

	// tZoneName
	TZoneName *string `json:"TZoneName,omitempty" name:"TZoneName"`

	// tRemark
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`

	// tRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`
}

func (r *CreateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DailyUseStatistics struct {

	// 总调用数
	TotalRequestCount *DailyUseStatisticsEntity `json:"TotalRequestCount,omitempty" name:"TotalRequestCount"`

	// 平均错误率
	AvgErrorRate *DailyUseStatisticsEntity `json:"AvgErrorRate,omitempty" name:"AvgErrorRate"`

	// 平均响应耗时
	AvgResponseCost *DailyUseStatisticsEntity `json:"AvgResponseCost,omitempty" name:"AvgResponseCost"`
}

type DailyUseStatisticsEntity struct {

	// 日统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 日环比值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DailyRate *string `json:"DailyRate,omitempty" name:"DailyRate"`

	// 周同比值
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeekRate *string `json:"WeekRate,omitempty" name:"WeekRate"`
}

type DeleteAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// PolicyId
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmReceiversRequest struct {
	*tchttp.BaseRequest

	// ReceiverIds
	ReceiverIds []*string `json:"ReceiverIds,omitempty" name:"ReceiverIds" list`
}

func (r *DeleteAlarmReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmReceiversRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmReceiversResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmReceiversResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiAccessRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteApiAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiGroupRequest struct {
	*tchttp.BaseRequest

	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DeleteApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除应用操作是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 权限规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuthorizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除结果true：删除成功；false：删除失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuthorizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// 日志配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 处理结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCircuitBreakerRuleRequest struct {
	*tchttp.BaseRequest

	// 熔断规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteCircuitBreakerRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCircuitBreakerRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCircuitBreakerRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCircuitBreakerRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCircuitBreakerRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除集群操作是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：删除成功；false：删除失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 无
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
}

func (r *DeleteConfigTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConfigTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：删除成功；false：删除失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConfigTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除操作是否成功：
	// true：成功
	// false：失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFileConfigRequest struct {
	*tchttp.BaseRequest

	// 文件配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFileConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFileConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayApiRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api ID 数组
	ApiList []*string `json:"ApiList,omitempty" name:"ApiList" list`
}

func (r *DeleteGatewayApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGatewayApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGatewayApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGatewayApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayPluginRequest struct {
	*tchttp.BaseRequest

	// 网关插件id数组
	PluginIdList []*string `json:"PluginIdList,omitempty" name:"PluginIdList" list`
}

func (r *DeleteGatewayPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGatewayPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGatewayPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGatewayPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
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

		// 删除部署组操作是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

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

type DeleteGroupSecretRequest struct {
	*tchttp.BaseRequest

	// 秘钥ID
	GwSecretId *string `json:"GwSecretId,omitempty" name:"GwSecretId"`
}

func (r *DeleteGroupSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupSecretRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupSecretResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageTag struct {

	// 仓库名，如/tsf/nginx
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 版本号:如V1
	TagName *string `json:"TagName,omitempty" name:"TagName"`
}

type DeleteImageTagRequest struct {
	*tchttp.BaseRequest

	// 版本号:如V1
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// （优先使用）仓库名，如/tsf/nginx
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 旧版仓库名
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`
}

func (r *DeleteImageTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageTagsRequest struct {
	*tchttp.BaseRequest

	// 镜像版本数组
	ImageTags []*DeleteImageTag `json:"ImageTags,omitempty" name:"ImageTags" list`
}

func (r *DeleteImageTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量删除操作是否成功。
	// true：成功。
	// false：失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMachinesRequest struct {
	*tchttp.BaseRequest

	// 机器列表
	MachineList []*MachineList `json:"MachineList,omitempty" name:"MachineList" list`
}

func (r *DeleteMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMachinesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMachinesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *DeleteMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMicroserviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除微服务是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMicroserviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMonitorStatisticsPolicyRequest struct {
	*tchttp.BaseRequest

	// 监控统计策略id;
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteMonitorStatisticsPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMonitorStatisticsPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMonitorStatisticsPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除结果true：删除成功；false：删除失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMonitorStatisticsPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMonitorStatisticsPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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

		// 删除命名空间是否成功。
	// true：删除成功。
	// false：删除失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

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

type DeletePkgRequest struct {
	*tchttp.BaseRequest

	// 无
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 无
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
}

func (r *DeletePkgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePkgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePkgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePkgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePkgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePkgsRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 需要删除的程序包ID列表
	PkgIds []*string `json:"PkgIds,omitempty" name:"PkgIds" list`
}

func (r *DeletePkgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePkgsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePkgsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePkgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePkgsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePodRequest struct {
	*tchttp.BaseRequest

	// groupId，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例名称
	PodName *string `json:"PodName,omitempty" name:"PodName"`
}

func (r *DeletePodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePodRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePodResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProgramRequest struct {
	*tchttp.BaseRequest

	// 数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`
}

func (r *DeleteProgramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProgramRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProgramResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true: 删除成功；false: 删除失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProgramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProgramResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeletePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：删除成功；false：删除失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRatelimitRequest struct {
	*tchttp.BaseRequest

	// 名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 限流所作用的目标微服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Rule列表，仅需要填写Id
	Rules []*RatelimitRuleForUpdate `json:"Rules,omitempty" name:"Rules" list`
}

func (r *DeleteRatelimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRatelimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRatelimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功与否
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRatelimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRatelimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRegionRequest struct {
	*tchttp.BaseRequest

	// tRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`
}

func (r *DeleteRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// reult
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRelatedGroupRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteRelatedGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRelatedGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRelatedGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求结果,true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRelatedGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRelatedGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleRequest struct {
	*tchttp.BaseRequest

	// 角色ID
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true: 删除成功；false: 删除失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteRequest struct {
	*tchttp.BaseRequest

	// 删除的路由ID
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
}

func (r *DeleteRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除路由是否成功。
	// true： 删除路由成功。
	// false：删除路由失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteRuleRequest struct {
	*tchttp.BaseRequest

	// 路由规则id
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`
}

func (r *DeleteRouteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRouteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true / false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRouteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRouteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalableRuleRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteScalableRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalableRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalableRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求结果,true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScalableRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalableRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// groupId，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceInstancesRequest struct {
	*tchttp.BaseRequest

	// InstanceIds
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DeleteServiceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateRequest struct {
	*tchttp.BaseRequest

	// 工程id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTsfZoneRequest struct {
	*tchttp.BaseRequest

	// tZoneId
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`
}

func (r *DeleteTsfZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTsfZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTsfZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTsfZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTsfZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest

	// tZoneId
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`
}

func (r *DeleteZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployContainGroupRequest struct {
	*tchttp.BaseRequest

	// groupId，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 镜像server
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称,如v1
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 最大分配cpu 核数
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配内存M数
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// JVM参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`
}

func (r *DeployContainGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployContainGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployContainGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployContainGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployContainGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 镜像server
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像版本名称,如v1
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 旧版镜像名，如/tsf/nginx
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 最大的 CPU 核数，对应 K8S 的 limit；不填时默认为 request 的 2 倍
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大的内存 MiB 数，对应 K8S 的 limit；不填时默认为 request 的 2 倍
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// jvm参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 分配的 CPU 核数，对应 K8S 的 request
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 分配的内存 MiB 数，对应 K8S 的 request
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 是否不立即启动
	DoNotStart *bool `json:"DoNotStart,omitempty" name:"DoNotStart"`

	// （优先使用）新版镜像名，如/tsf/nginx
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`
}

func (r *DeployContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署容器应用是否成功。
	// true：成功。
	// false：失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployGroup struct {

	// 部署组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 集群Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 部署组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 部署组更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 部署组实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 部署组运行实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 部署组停止实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffInstanceCount *int64 `json:"OffInstanceCount,omitempty" name:"OffInstanceCount"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`

	// 部署组启动参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 部署组程序包Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 部署组程序包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 部署组程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 最大分配cpu 核数，如0.6
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配内存M数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 0:公网 1:集群内访问 2：NodePort
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 更新方式：0:快速更新 1:滚动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 端口数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// Service ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// 负载均衡ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

	// 环境变量数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envs []*Env `json:"Envs,omitempty" name:"Envs" list`

	// pod错误信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// NodePort端口，只有公网和NodePort访问方式才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePort *int64 `json:"NodePort,omitempty" name:"NodePort"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`
}

type DeployGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 部署组启动参数
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`
}

func (r *DeployGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployInstanceRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *DeployInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 所需实例内存大小，取值为 1Gi 2Gi 4Gi 8Gi 16Gi，缺省为 1Gi，不传表示维持原态
	Memory *string `json:"Memory,omitempty" name:"Memory"`

	// 要求最小实例数，取值范围 [1, 4]，缺省为 1，不传表示维持原态
	InstanceRequest *uint64 `json:"InstanceRequest,omitempty" name:"InstanceRequest"`

	// 部署组启动参数，不传表示维持原态
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`
}

func (r *DeployServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployTsfModulesRequest struct {
	*tchttp.BaseRequest

	// ModuleIds
	ModuleIds []*string `json:"ModuleIds,omitempty" name:"ModuleIds" list`
}

func (r *DeployTsfModulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployTsfModulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployTsfModulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployTsfModulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployTsfModulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddibleInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 可用区过滤字段
	InstanceZoneId *string `json:"InstanceZoneId,omitempty" name:"InstanceZoneId"`
}

func (r *DescribeAddibleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddibleInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddibleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddibleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddibleInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmOverviewListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmOverviewListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmOverviewListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmOverviewListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警历史信息列表
		Result []*AlarmOverviewContent `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmOverviewListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmOverviewListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// PolicyId
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *AlarmPolicyResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAccessRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeApiAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 输出结果
		Result *ApplicationApiAccess `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiDetailRequest struct {
	*tchttp.BaseRequest

	// 微服务id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApiDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API 详情
		Result *ApiDetailResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiGroupRequest struct {
	*tchttp.BaseRequest

	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API分组信息
		Result *ApiGroupInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeApiGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 翻页结构体
		Result *TsfPageApiGroupInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiRateLimitRulesRequest struct {
	*tchttp.BaseRequest

	// Api ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DescribeApiRateLimitRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiRateLimitRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiRateLimitRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 限流结果
		Result []*ApiRateLimitRule `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiRateLimitRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiRateLimitRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUseDetailRequest struct {
	*tchttp.BaseRequest

	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组Api ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeApiUseDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiUseDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUseDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日使用统计对象
		Result *GroupApiUseStatistics `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiUseDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiUseDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiVersionsRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// API 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *DescribeApiVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiVersionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiVersionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API版本列表
		Result []*ApiVersionArray `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiVersionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApisWithPluginRequest struct {
	*tchttp.BaseRequest

	// PluginId
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 绑定/未绑定: true / false
	Bound *string `json:"Bound,omitempty" name:"Bound"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// API 分组ID，若传入则根据分组过滤API
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeApisWithPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApisWithPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApisWithPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API分组信息列表
		Result *TsfPageApiDetailInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApisWithPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApisWithPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationAttributeRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用列表其它字段返回参数
		Result *ApplicationAttribute `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// TSF应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationInstancesRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeApplicationInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用信息
		Result *ApplicationForPage `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationsAttributeRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationsAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationsAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationsAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用列表其它字段返回参数
		Result *ApplicationAttribute `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationsAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationsAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationsOtherRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationsOtherRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationsOtherRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationsOtherResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用列表其它字段返回参数
		Result *ApplicationAttribute `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationsOtherResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationsOtherResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用的微服务类型
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用资源类型数组
	ApplicationResourceTypeList []*string `json:"ApplicationResourceTypeList,omitempty" name:"ApplicationResourceTypeList" list`
}

func (r *DescribeApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用分页列表信息
		Result *TsfPageApplication `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicatoinServerLogRequest struct {
	*tchttp.BaseRequest

	// 日志搜索的开始时间，格式："YYYY-MM-DD"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志搜索的结束时间，格式："YYYY-MM-DD"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 应用服务器的id
	ApplicationServerID *string `json:"ApplicationServerID,omitempty" name:"ApplicationServerID"`

	// 日志类型, "tsf-agent", "pilot-agent", "mesh-dns", "envoy"
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 用来过滤日志内容的关键词
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询起始偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeApplicatoinServerLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicatoinServerLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicatoinServerLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志内容
		Result *TsfPageApplicatoinServerLog `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicatoinServerLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicatoinServerLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthNamespacesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAuthNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthNamespacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间列表
		Result []*GatewayNamespace `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthNamespacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationInfoRequest struct {
	*tchttp.BaseRequest

	// 被调用服务 ID
	TargetServiceId *string `json:"TargetServiceId,omitempty" name:"TargetServiceId"`

	// 命名空间 ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DescribeAuthorizationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 鉴权规则列表
		Result *TsfPageAuthorization `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthorizationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationMicroservicesRequest struct {
	*tchttp.BaseRequest

	// 目标微服务ID
	TargetServiceId *string `json:"TargetServiceId,omitempty" name:"TargetServiceId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 查询字符串
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 查询数据偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAuthorizationMicroservicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationMicroservicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationMicroservicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当前页的微服务数据
		Content []*string `json:"Content,omitempty" name:"Content" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthorizationMicroservicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationMicroservicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 微服务权限规则
		Result *AuthRule `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationTypeRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *DescribeAuthorizationTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权限规则组
		Result *AuthRuleGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthorizationTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationsRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 规则ID，不传入时查询全量
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 是否生效，不传入时查询全量
	IsEnabled *string `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAuthorizationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthorizationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的微服务权限规则列表
		Result *TsfPageAuthRule `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthorizationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthorizationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志配置项
		Result *BusinessLogConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBusinessLogConfigsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量，取值范围[1, 50]，默认值为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊匹配关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeBusinessLogConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBusinessLogConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBusinessLogConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务日志配置列表
		Result *TsfPageBusinessLogConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBusinessLogConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBusinessLogConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCircuitBreakerRuleRequest struct {
	*tchttp.BaseRequest

	// 熔断规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeCircuitBreakerRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCircuitBreakerRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCircuitBreakerRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 熔断规则
		Result *CircuitBreakerRule `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCircuitBreakerRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCircuitBreakerRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCircuitBreakerRulesRequest struct {
	*tchttp.BaseRequest

	// 每页展示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 熔断规则微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务所属命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 熔断级别（用于过滤）
	IsolationLevel *string `json:"IsolationLevel,omitempty" name:"IsolationLevel"`
}

func (r *DescribeCircuitBreakerRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCircuitBreakerRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCircuitBreakerRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 熔断规则列表
		Result *CircuitBreakerRules `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCircuitBreakerRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCircuitBreakerRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudMonitorGroupsRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤GroupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeCloudMonitorGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudMonitorGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudMonitorGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署组分页信息
		Result *TsfPageMonitorGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudMonitorGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudMonitorGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceCountRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterInstanceCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstanceCountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群信息
		Result *Cluster `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstanceCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstanceCountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群机器实例分页信息
		Result *TsfPageInstance `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterLimitResourceRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterLimitResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterLimitResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterLimitResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群剩余资源
		Result *ClusterLimitResourceV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterLimitResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterLimitResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRequest struct {
	*tchttp.BaseRequest

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群
		Result *ClusterV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSchedulabilityRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组预分配的CPU资源（单位：核）
	// 用于新建、部署操作时必填；用于扩缩容时可不填
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 部署组预分配的内存资源（单位：M）
	// 用于新建、部署操作时必填；用于扩缩容时可不填
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`
}

func (r *DescribeClusterSchedulabilityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterSchedulabilityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSchedulabilityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群支持调度的实例数
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterSchedulabilityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterSchedulabilityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// 搜索词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClustersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster分页信息
		Result *TsfPageClusterV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClustersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCommonPkgRequest struct {
	*tchttp.BaseRequest

	// searchWord
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeCommonPkgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCommonPkgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCommonPkgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 公共包信息
		Result *TsfPageCommonPkg `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCommonPkgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCommonPkgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigReleaseLogsRequest struct {
	*tchttp.BaseRequest

	// 部署组ID，不传入时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeConfigReleaseLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigReleaseLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigReleaseLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的配置项发布历史列表
		Result *TsfPageConfigReleaseLog `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigReleaseLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigReleaseLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigReleasesRequest struct {
	*tchttp.BaseRequest

	// 配置项名称，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 部署组ID，不传入时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 配置ID，不传入时查询全量
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeConfigReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigReleasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigReleasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的配置发布信息
		Result *TsfPageConfigRelease `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigReleasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置项
		Result *Config `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigSummaryRequest struct {
	*tchttp.BaseRequest

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字，模糊查询：应用名称，配置项名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置项分页对象
		Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 无
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
}

func (r *DescribeConfigTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *ConfigTemplate `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigTemplatesRequest struct {
	*tchttp.BaseRequest

	// 配置模板名称
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板对应的微服务框架
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 距离第一条的偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 无
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeConfigTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置模板列表
		Result *ConfigTemplateResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 配置项ID，不传入时查询全量，高优先级
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项ID列表，不传入时查询全量，低优先级
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList" list`

	// 配置项名称，精确查询，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页后的配置项列表
		Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeContainerGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署组列表-其它字段
		Result *ContainerGroupOther `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupDeployInfoRequest struct {
	*tchttp.BaseRequest

	// 实例所属groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeContainerGroupDeployInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupDeployInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupDeployInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取部署组
		Result *ContainerGroupDeploy `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerGroupDeployInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupDeployInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupDetailRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeContainerGroupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 容器部署组详情
		Result *ContainerGroupDetail `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerGroupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段，模糊搜索groupName字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间 ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DescribeContainerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的权限数据对象
		Result *ContainGroupResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerMicroServiceListRequest struct {
	*tchttp.BaseRequest

	// 微服务Id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContainerMicroServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerMicroServiceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerMicroServiceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 列表信息
		Content []*string `json:"Content,omitempty" name:"Content" list`

		// 实例名称(对应到kubernetes的pod名称)
		PodName *string `json:"PodName,omitempty" name:"PodName"`

		// 实例ID(对应到kubernetes的pod id)
		PodId *string `json:"PodId,omitempty" name:"PodId"`

		// 实例状态，请参考后面的实例的状态定义:
	// Running     正常运行中
	// Waiting     等待运行中，例如正在下载镜像
	// Terminating  	实例有容器正在终止
	// Terminated    实例有容器已经终止
	// NotReady    	实例有容器处于未就绪状态，比如容器的健康检查失败
	// Error   其他未知异常
		Status *string `json:"Status,omitempty" name:"Status"`

		// 主机IP
		NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`

		// 实例ip
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// 镜像名，如/tsf/nginx
		Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

		// 镜像版本名称
		TagName *string `json:"TagName,omitempty" name:"TagName"`

		// 应用Id
		ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

		// 应用名称
		ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

		// 集群id
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 集群名称
		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

		// 命名空间id
		NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

		// 命名空间名称
		NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

		// 机器所属groupId
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 分组名称
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 机器状态
		InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

		// 实例可用状态
		InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

		// 服务实例状态
		ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerMicroServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerMicroServiceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerTasksRequest struct {
	*tchttp.BaseRequest

	// 操作列表所属applicationId
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 任务执行时间范围，格式："2017-10-01 10:03:27"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务执行时间范围，格式："2017-11-01 10:03:27"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 搜索字段，模糊搜索taskId字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段，默认为updateTime（修改时间字段）， 支持id， createTime， updateTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// v
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContainerTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询任务列表对象
		Result *DescribeContainerTasksResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerTasksResult struct {

	// 列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ContainerTasks `json:"Content,omitempty" name:"Content" list`
}

type DescribeDeployGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeDeployGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeployGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署组详情
		Result *DeployGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeployGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeployGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDeployGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeployGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeployGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeployGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerForUseRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDockerForUseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDockerForUseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerForUseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Docker使用指引
		Result *DockerForUse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDockerForUseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDockerForUseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
}

func (r *DescribeDownloadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// COS鉴权信息
		Result *CosDownloadInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigReleaseLogsRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeFileConfigReleaseLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigReleaseLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigReleaseLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置项发布历史分页
		Result *TsfPageFileConfigReleaseLog `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileConfigReleaseLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigReleaseLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigReleasesRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFileConfigReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigReleasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigReleasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置项发布信息列表
		Result *TsfPageFileConfigRelease `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileConfigReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigReleasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件配置项
		Result *FileConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigSummaryRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFileConfigSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置项分页对象
		Result *TsfPageFileConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileConfigSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigsRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList" list`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFileConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件配置项列表
		Result *TsfPageFileConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGWOverviewInvocationRequest struct {
	*tchttp.BaseRequest

	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 监控统计类型: "SumReqAmount", "AvgFailureRate", "AvgTimeCost"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 分组下的ApiID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 监控统计数据粒度
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 开始时间，默认当前时间, eg: 2019-11-15 14:28:27
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认当前时间 eg: 2019-11-15 14:28:27
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeGWOverviewInvocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGWOverviewInvocationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGWOverviewInvocationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控统计数据点
		Result []*MetricDataPoint `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGWOverviewInvocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGWOverviewInvocationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayAllGroupApisRequest struct {
	*tchttp.BaseRequest

	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`
}

func (r *DescribeGatewayAllGroupApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayAllGroupApisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayAllGroupApisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网关分组和API列表信息
		Result *GatewayVo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayAllGroupApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayAllGroupApisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayApisRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页的记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字，支持 API path
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeGatewayApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayApisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayApisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 翻页结构
		Result *TsfPageApiDetailInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayApisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayDailyUseStatisticsRequest struct {
	*tchttp.BaseRequest

	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 查询的日期,格式：yyyy-MM-dd
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 分组下的ApiID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DescribeGatewayDailyUseStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayDailyUseStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayDailyUseStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日使用统计对象
		Result *DailyUseStatistics `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayDailyUseStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayDailyUseStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayJwtPluginRequest struct {
	*tchttp.BaseRequest

	// 网关插件id
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeGatewayJwtPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayJwtPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayJwtPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 插件信息列表
		Result *GatewayJwtPlugin `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayJwtPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayJwtPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayMonitorOverviewRequest struct {
	*tchttp.BaseRequest

	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`
}

func (r *DescribeGatewayMonitorOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayMonitorOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayMonitorOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控概览对象
		Result *MonitorOverview `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayMonitorOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayMonitorOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayOAuthPluginRequest struct {
	*tchttp.BaseRequest

	// 网关插件id
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeGatewayOAuthPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayOAuthPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayOAuthPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 插件信息列表
		Result *GatewayOAuthPlugin `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayOAuthPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayOAuthPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayPluginTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGatewayPluginTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayPluginTypesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayPluginTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 插件类型列表
		Result []*GatewayPluginDefinition `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayPluginTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayPluginTypesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayPluginsRequest struct {
	*tchttp.BaseRequest

	// 翻页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页展示的条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 插件类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeGatewayPluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayPluginsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayPluginsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 插件信息列表
		Result *TsfPageGatewayPlugin `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayPluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayPluginsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayTagPluginRequest struct {
	*tchttp.BaseRequest

	// 网关插件id
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeGatewayTagPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayTagPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayTagPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 插件信息列表
		Result *GatewayTagPlugin `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayTagPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGatewayTagPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAddibleInstancesRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupAddibleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupAddibleInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAddibleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 虚拟机部署组可添加机器列表
		Result *TsfPageInstance `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupAddibleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupAddibleInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 部署组ID字段
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 虚拟机部署组信息
		Result *VmGroupOther `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupBindedGatewaysRequest struct {
	*tchttp.BaseRequest

	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeGroupBindedGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupBindedGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupBindedGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 翻页结构体
		Result *TsfPageGatewayDeployGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupBindedGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupBindedGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupBusinessLogConfigsRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupBusinessLogConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupBusinessLogConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupBusinessLogConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务日志配置列表
		Result *TsfPageBusinessLogConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupBusinessLogConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupBusinessLogConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupGatewaysRequest struct {
	*tchttp.BaseRequest

	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeGroupGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API分组信息
		Result *TsfPageApiGroupInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInstancesRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署组机器信息
		Result *TsfPageInstance `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupOtherRequest struct {
	*tchttp.BaseRequest

	// 部署组ID字段
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupOtherRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupOtherRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupOtherResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 虚拟机部署组信息
		Result *VmGroupOther `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupOtherResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupOtherResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
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

		// 虚拟机部署组详情
		Result *VmGroup `json:"Result,omitempty" name:"Result"`

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

type DescribeGroupSecretsRequest struct {
	*tchttp.BaseRequest

	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupSecretsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupSecretsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupSecretsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 秘钥返回结果
		Result []*SecretKeyInfo `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupSecretsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupSecretsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupUseDetailRequest struct {
	*tchttp.BaseRequest

	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定top的条数,默认为10
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

func (r *DescribeGroupUseDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupUseDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupUseDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日使用统计对象
		Result *GroupDailyUseStatistics `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupUseDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupUseDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组资源类型列表
	GroupResourceTypeList []*string `json:"GroupResourceTypeList,omitempty" name:"GroupResourceTypeList" list`
}

func (r *DescribeGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 虚拟机部署组分页信息
		Result *TsfPageVmGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsWithPluginRequest struct {
	*tchttp.BaseRequest

	// PluginId
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 绑定/未绑定: true / false
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeGroupsWithPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupsWithPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsWithPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API分组信息列表
		Result *TsfPageApiGroupInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupsWithPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupsWithPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRepositoryRequest struct {
	*tchttp.BaseRequest

	// 仓库名，搜索关键字,不带命名空间的
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageRepositoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的权限数据对象
		Result *ImageRepositoryResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageRepositoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageTagsRequest struct {
	*tchttp.BaseRequest

	// 应用Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 不填和0:查询 1:不查询
	QueryImageIdFlag *int64 `json:"QueryImageIdFlag,omitempty" name:"QueryImageIdFlag"`
}

func (r *DescribeImageTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的权限数据对象
		Result *ImageTagsResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageUserIsExistsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageUserIsExistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageUserIsExistsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageUserIsExistsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询到的信息
		Result *ImageUserIsExistsResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageUserIsExistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageUserIsExistsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
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

type DescribeInovcationIndicatorsRequest struct {
	*tchttp.BaseRequest

	// 维度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 调用方服务名
	CallerServiceName *string `json:"CallerServiceName,omitempty" name:"CallerServiceName"`

	// 被调方服务名
	CalleeServiceName *string `json:"CalleeServiceName,omitempty" name:"CalleeServiceName"`

	// 调用方接口名
	CallerInterfaceName *string `json:"CallerInterfaceName,omitempty" name:"CallerInterfaceName"`

	// 被调方接口名
	CalleeInterfaceName *string `json:"CalleeInterfaceName,omitempty" name:"CalleeInterfaceName"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInovcationIndicatorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInovcationIndicatorsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInovcationIndicatorsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务调用监控指标
		Result *InvocationIndicator `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInovcationIndicatorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInovcationIndicatorsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvocationStatisticsRatioRequest struct {
	*tchttp.BaseRequest

	// 维度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 目标开始时间
	TargetStartTime *string `json:"TargetStartTime,omitempty" name:"TargetStartTime"`

	// 目标结束时间
	TargetEndTime *string `json:"TargetEndTime,omitempty" name:"TargetEndTime"`

	// 服务角色
	Role *string `json:"Role,omitempty" name:"Role"`

	// 时间粒度
	Period *string `json:"Period,omitempty" name:"Period"`

	// 微服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 目标服务名，用于对比
	TargetServiceName *string `json:"TargetServiceName,omitempty" name:"TargetServiceName"`

	// 目标接口名，用于对比
	TargetInterfaceName *string `json:"TargetInterfaceName,omitempty" name:"TargetInterfaceName"`
}

func (r *DescribeInvocationStatisticsRatioRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvocationStatisticsRatioRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvocationStatisticsRatioResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调用统计数据对比
		Result *InvocationStatisticsRatio `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvocationStatisticsRatioResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvocationStatisticsRatioResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvocationStatisticsRequest struct {
	*tchttp.BaseRequest

	// 维度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务角色
	Role *string `json:"Role,omitempty" name:"Role"`

	// 时间粒度
	Period *string `json:"Period,omitempty" name:"Period"`

	// 微服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 目标服务名
	TargetServiceName *string `json:"TargetServiceName,omitempty" name:"TargetServiceName"`

	// 目标接口名
	TargetInterfaceName *string `json:"TargetInterfaceName,omitempty" name:"TargetInterfaceName"`
}

func (r *DescribeInvocationStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvocationStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInvocationStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调用统计数据
		Result *TsfPageInvocationStatisticsV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvocationStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInvocationStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseApplicationsRequest struct {
	*tchttp.BaseRequest

	// 列表起始偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页数量上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按被授予者名字搜索
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeLicenseApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLicenseApplicationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 申请总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Content []*LicenseApplicationRecord `json:"Content,omitempty" name:"Content" list`
	} `json:"Response"`
}

func (r *DescribeLicenseApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLicenseApplicationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroserviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 微服务详情实例列表
		Result *TsfPageMsInstance `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroserviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroservicesByAssociateApplicationRequest struct {
	*tchttp.BaseRequest

	// 服务所关联的应用 ID
	AssociateApplicationId *string `json:"AssociateApplicationId,omitempty" name:"AssociateApplicationId"`
}

func (r *DescribeMicroservicesByAssociateApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroservicesByAssociateApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroservicesByAssociateApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 微服务列表
		Result []*MicroserviceExtra `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroservicesByAssociateApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroservicesByAssociateApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroservicesRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMicroservicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroservicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroservicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 微服务分页列表信息
		Result *TsfPageMicroservice `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroservicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroservicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorStatisticsPolicyRequest struct {
	*tchttp.BaseRequest

	// 监控统计策略id;
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeMonitorStatisticsPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorStatisticsPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorStatisticsPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *MonitorStatisticsPolicy `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorStatisticsPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMonitorStatisticsPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsApiListRequest struct {
	*tchttp.BaseRequest

	// 微服务ID。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 搜索关键字。
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 没页的数量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMsApiListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsApiListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsApiListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 相应结果
		Result *TsfApiListResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsApiListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsApiListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsRouteFallbackRequest struct {
	*tchttp.BaseRequest

	// 路由所属微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *DescribeMsRouteFallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsRouteFallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsRouteFallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true： 保护策略开启状态； false：保护策略停用状态
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsRouteFallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsRouteFallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsRunningApplicationsRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *DescribeMsRunningApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsRunningApplicationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMsRunningApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 微服务运行态分页应用列表
		Result *TsfPageMsRunningApplicationV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsRunningApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMsRunningApplicationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceAffinitiesRequest struct {
	*tchttp.BaseRequest

	// 命名空间列表
	NamespaceIdList []*string `json:"NamespaceIdList,omitempty" name:"NamespaceIdList" list`
}

func (r *DescribeNamespaceAffinitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespaceAffinitiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceAffinitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 就近访问策略列表
		Result *TsfPageRouteAffinity `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespaceAffinitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespaceAffinitiesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceAffinityRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DescribeNamespaceAffinityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespaceAffinityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceAffinityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true： 就近访问策略开启状态； false：就近访问策略停用状态
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespaceAffinityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespaceAffinityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest

	// 搜索字段，查询命名空间ID或命名空间名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间名称，不传入时查询全量
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 表示要查询的 资源类型列表
	NamespaceResourceTypeList []*string `json:"NamespaceResourceTypeList,omitempty" name:"NamespaceResourceTypeList" list`
}

func (r *DescribeNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间分页信息
		Result *TsfPageNamespace `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewInvocationRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 监控统计类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 监控统计数据粒度
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOverviewInvocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviewInvocationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewInvocationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控统计数据列表
		Result []*MetricDataPoint `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewInvocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviewInvocationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewResourceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOverviewResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviewResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 概览页资源信息
		Result *OverviewResource `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviewResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviweServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOverviweServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviweServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviweServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 概览页微服务信息
		Result *OverviewMsResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviweServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOverviweServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionCategoriesRequest struct {
	*tchttp.BaseRequest

	// 角色ID，不传入时查询全量
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 产品编码，不传入时查询全量
	ServiceCode *string `json:"ServiceCode,omitempty" name:"ServiceCode"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePermissionCategoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePermissionCategoriesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionCategoriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权限组列表
		Result *PagedPermCat `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePermissionCategoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePermissionCategoriesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePkgsRequest struct {
	*tchttp.BaseRequest

	// 应用ID（只传入应用ID，返回该应用下所有软件包信息）
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字（支持根据包ID，包名，包版本号搜索）
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序关键字（默认为"UploadTime"：上传时间）
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序：0/降序：1（默认降序）
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// 查询起始偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePkgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePkgsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePkgsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询程序包信息列表
		Result *PkgList `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePkgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePkgsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePluginInstancesRequest struct {
	*tchttp.BaseRequest

	// 分组或者API的ID
	ScopeValue *string `json:"ScopeValue,omitempty" name:"ScopeValue"`

	// 绑定: true; 未绑定: false
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页展示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 插件类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribePluginInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePluginInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePluginInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 插件信息列表
		Result *TsfPageGatewayPlugin `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePluginInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePluginInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePodInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所属groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePodInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePodInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePodInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的权限数据对象
		Result *GroupPodResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePodInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePodInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProgramRequest struct {
	*tchttp.BaseRequest

	// 数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`
}

func (r *DescribeProgramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProgramRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProgramResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据集
		Result *Program `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProgramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProgramResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProgramsRequest struct {
	*tchttp.BaseRequest

	// 模糊查询数据集ID，数据集名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeProgramsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProgramsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProgramsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据集列表
		Result *PagedProgram `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProgramsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProgramsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigReleaseLogsRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePublicConfigReleaseLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigReleaseLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigReleaseLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页后的公共配置项发布历史列表
		Result *TsfPageConfigReleaseLog `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigReleaseLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigReleaseLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigReleasesRequest struct {
	*tchttp.BaseRequest

	// 配置项名称，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 配置项ID，不传入时查询全量
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribePublicConfigReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigReleasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigReleasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 公共配置发布信息
		Result *TsfPageConfigRelease `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigReleasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 需要查询的配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 全局配置
		Result *Config `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigSummaryRequest struct {
	*tchttp.BaseRequest

	// 查询关键字，模糊查询：配置项名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePublicConfigSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的全局配置统计信息列表
		Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigsRequest struct {
	*tchttp.BaseRequest

	// 配置项ID，不传入时查询全量，高优先级
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项ID列表，不传入时查询全量，低优先级
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList" list`

	// 配置项名称，精确查询，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
}

func (r *DescribePublicConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页后的全局配置项列表
		Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRatelimitCommitConfigRequest struct {
	*tchttp.BaseRequest

	// 名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 限流所作用的目标微服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeRatelimitCommitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRatelimitCommitConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRatelimitCommitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已有全局规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		GlobalRuleCount *uint64 `json:"GlobalRuleCount,omitempty" name:"GlobalRuleCount"`
	} `json:"Response"`
}

func (r *DescribeRatelimitCommitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRatelimitCommitConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRatelimitRequest struct {
	*tchttp.BaseRequest

	// 名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 限流所作用的目标微服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 列表偏移位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 列表最大长度
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 检索规则id或搜索规则名
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeRatelimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRatelimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRatelimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule列表
		Result *RatelimitListResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRatelimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRatelimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordCodesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRecordCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecordCodesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordCodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志术语返回结果
		Result *TsfRecordCodeResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecordCodesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordsRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 操作类型
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// 资源类型
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 搜索词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页发小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 操作者
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作记录分页列表
		Result *TsfPageRecordV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest

	// TRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`

	// TRegionName
	TRegionName *string `json:"TRegionName,omitempty" name:"TRegionName"`
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

		// Result
		Result *ListTsfRegionResult `json:"Result,omitempty" name:"Result"`

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

type DescribeReleasedConfigRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeReleasedConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReleasedConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReleasedConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已发布的配置内容
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReleasedConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReleasedConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceConfigLicense struct {

	// Function
	// 注意：此字段可能返回 null，表示取不到有效值。
	Function []*DescribeResourceConfigLicenseFunction `json:"Function,omitempty" name:"Function" list`

	// Resource
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource []*DescribeResourceConfigLicenseResource `json:"Resource,omitempty" name:"Resource" list`

	// utc时间 单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// utc时间 单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Countdown *uint64 `json:"Countdown,omitempty" name:"Countdown"`
}

type DescribeResourceConfigLicenseFunction struct {

	// name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// enable
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type DescribeResourceConfigLicenseResource struct {

	// Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Quota
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quota *uint64 `json:"Quota,omitempty" name:"Quota"`
}

type DescribeResourceConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeResourceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置详情
		Result *DescribeResourceConfigResultV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceConfigResultV2 struct {

	// Sts
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sts *DescribeResourceConfigSts `json:"Sts,omitempty" name:"Sts"`

	// License
	// 注意：此字段可能返回 null，表示取不到有效值。
	License *DescribeResourceConfigLicense `json:"License,omitempty" name:"License"`

	// 部署组相关的参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Group *GroupResourceConfig `json:"Group,omitempty" name:"Group"`

	// 实例相关的参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *InstanceResourceConfig `json:"Instance,omitempty" name:"Instance"`
}

type DescribeResourceConfigSts struct {

	// uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type DescribeResourceUsageRateRequest struct {
	*tchttp.BaseRequest

	// 资源类型，es，rediss，consul，ctsdb
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeResourceUsageRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceUsageRateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceUsageRateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源使用情况的返回结果
		Result *ResourceStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceUsageRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceUsageRateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest

	// 是否需要Action存在，Y：需要；N（默认）：不需要，为Y时，仅返回ResourceAction字段有值的数据。
	NeedAction *string `json:"NeedAction,omitempty" name:"NeedAction"`

	// 资源编码，不传入时查询全量
	ResourceCode *string `json:"ResourceCode,omitempty" name:"ResourceCode"`

	// 产品ID列表，不传入时查询全量
	ServiceIdList []*string `json:"ServiceIdList,omitempty" name:"ServiceIdList" list`

	// 产品编码列表，不传入时查询全量
	ServiceCodeList []*string `json:"ServiceCodeList,omitempty" name:"ServiceCodeList" list`

	// 模糊查询，资源ID/资源编码/资源描述 字段，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourcesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的资源列表
		Result *PagedResource `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourcesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleRequest struct {
	*tchttp.BaseRequest

	// 角色ID
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
}

func (r *DescribeRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色
		Result *Role `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesRequest struct {
	*tchttp.BaseRequest

	// 角色名称，不传入时查询全量
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 模糊查询，角色ID/角色名称/角色描述 字段，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRolesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的角色列表
		Result *PagedRole `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRolesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteReleaseHistoryRequest struct {
	*tchttp.BaseRequest

	// 微服务Id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRouteReleaseHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteReleaseHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteReleaseHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由规则启停记录
		Result *TsfPageRouteReleaseHistory `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteReleaseHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteReleaseHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteRequest struct {
	*tchttp.BaseRequest

	// 路由ID
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
}

func (r *DescribeRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由详情信息
		Result *RouteV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteRuleRequest struct {
	*tchttp.BaseRequest

	// 路由规则id
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`
}

func (r *DescribeRouteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由规则信息
		Result *RouteRule `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteRulesRequest struct {
	*tchttp.BaseRequest

	// 微服务id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRouteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由规则分页对象
		Result *TsfPageRouteRule `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoutesRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由分页列表
		Result *TsfPageRouteV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalableRuleAttributeByGroupRequest struct {
	*tchttp.BaseRequest

	// 应用id;
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeScalableRuleAttributeByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalableRuleAttributeByGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalableRuleAttributeByGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则信息
		Result *GroupsByScalableRuleId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScalableRuleAttributeByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalableRuleAttributeByGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalableRuleAttributeRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeScalableRuleAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalableRuleAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalableRuleAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *ScalableRuleAttribute `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScalableRuleAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalableRuleAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
		Result *ServerlessGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段，模糊搜索groupName字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，取值从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分组所属名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 分组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeServerlessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据列表对象
		Result *ServerlessGroupPage `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerlessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceStatisticsRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务ID列表
	ServiceIds []*string `json:"ServiceIds,omitempty" name:"ServiceIds" list`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeServiceStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计数据列表
		Result []*ServiceStatisticsV2 `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesRequest struct {
	*tchttp.BaseRequest

	// 产品编码，不填写时查询全量
	ServiceCode *string `json:"ServiceCode,omitempty" name:"ServiceCode"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品分页列表
		Result *PagedService `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleApplicationsRequest struct {
	*tchttp.BaseRequest

	// 应用ID列表
	ApplicationIdList []*string `json:"ApplicationIdList,omitempty" name:"ApplicationIdList" list`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 微服务类型
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 资源类型数组
	ApplicationResourceTypeList []*string `json:"ApplicationResourceTypeList,omitempty" name:"ApplicationResourceTypeList" list`

	// 通过id和name进行关键词过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeSimpleApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleApplicationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 简单应用分页对象
		Result *TsfPageSimpleApplication `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleApplicationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleClustersRequest struct {
	*tchttp.BaseRequest

	// 需要查询的集群ID列表，不填或不传入时查询所有内容
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList" list`

	// 需要查询的集群类型，不填或不传入时查询所有内容
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 查询偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 对id和name进行关键词过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeSimpleClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleClustersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TSF集群分页对象
		Result *TsfPageCluster `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleClustersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleGroupsRequest struct {
	*tchttp.BaseRequest

	// 部署组ID列表，不填写时查询全量
	GroupIdList []*string `json:"GroupIdList,omitempty" name:"GroupIdList" list`

	// 应用ID，不填写时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 集群ID，不填写时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间ID，不填写时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 部署组ID，不填写时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 模糊查询，部署组名称，不填写时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 部署组类型，精确过滤字段，M：service mesh, P：原生应用， M：网关应用
	AppMicroServiceType *string `json:"AppMicroServiceType,omitempty" name:"AppMicroServiceType"`
}

func (r *DescribeSimpleGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 简单部署组列表
		Result *TsfPageSimpleGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleInstancesRequest struct {
	*tchttp.BaseRequest

	// 机器实例ID列表，不传入时查询全量
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSimpleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器列表
		Result *TsfPageInstance `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleNamespacesRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID列表，不传入时查询全量
	NamespaceIdList []*string `json:"NamespaceIdList,omitempty" name:"NamespaceIdList" list`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 查询资源类型列表
	NamespaceResourceTypeList []*string `json:"NamespaceResourceTypeList,omitempty" name:"NamespaceResourceTypeList" list`

	// 通过id和name进行过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 查询的命名空间类型列表
	NamespaceTypeList []*string `json:"NamespaceTypeList,omitempty" name:"NamespaceTypeList" list`

	// 通过命名空间名精确过滤
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 通过是否是默认命名空间过滤，不传表示拉取全部命名空间。0：默认，命名空间。1：非默认命名空间
	IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`
}

func (r *DescribeSimpleNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleNamespacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间分页列表
		Result *TsfPageNamespace `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleNamespacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSingleContainerGroups struct {

	// groupId
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 分组创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 端口列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 更新类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新时间间隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 网络访问类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`
}

type DescribeSingleContainerGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，默认为0，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组资源类型数组
	GroupResourceTypeList []*string `json:"GroupResourceTypeList,omitempty" name:"GroupResourceTypeList" list`
}

func (r *DescribeSingleContainerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSingleContainerGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSingleContainerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务返回内容
		Result *DescribeSingleContainerGroupsResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSingleContainerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSingleContainerGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSingleContainerGroupsResult struct {

	// 列表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*DescribeSingleContainerGroups `json:"Content,omitempty" name:"Content" list`
}

type DescribeSubTasksRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量大小
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 虚拟机子任务信息
		Result *TsfPageVmSubTask `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubTransactionsRequest struct {
	*tchttp.BaseRequest

	// 事务id
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 搜索词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 事务状态
	TransactionState *int64 `json:"TransactionState,omitempty" name:"TransactionState"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubTransactionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubTransactionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubTransactionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubTransactionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubTransactionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest

	// 工程id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest

	// 搜索字段，模糊搜索ProjectName , BasePackage 字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段，默认为 LastTime 字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 查询偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *TemplateResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTransactionRequest struct {
	*tchttp.BaseRequest

	// 事务id
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
}

func (r *DescribeTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTransactionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTransactionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTransactionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTransactionsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 事务查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 事务查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 事务状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 事务状态
	TransactionState *int64 `json:"TransactionState,omitempty" name:"TransactionState"`
}

func (r *DescribeTransactionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTransactionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTransactionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询结果
		Result *TxList `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTransactionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTransactionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTsfRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTsfRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTsfRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTsfRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TSF地域列表
		Result *TsfPageRegion `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTsfRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTsfRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTsfZonesRequest struct {
	*tchttp.BaseRequest

	// TSF地域ID
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`
}

func (r *DescribeTsfZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTsfZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTsfZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *TsfPageZone `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTsfZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTsfZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTsfmanagerZonesRequest struct {
	*tchttp.BaseRequest

	// TZoneId
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`

	// TZoneName
	TZoneName *string `json:"TZoneName,omitempty" name:"TZoneName"`

	// TRemark
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`

	// TRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`
}

func (r *DescribeTsfmanagerZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTsfmanagerZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTsfmanagerZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *ListTsfZoneResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTsfmanagerZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTsfmanagerZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 程序包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 程序包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 程序包类型
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 程序包介绍
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`
}

func (r *DescribeUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUploadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// COS上传信息
		Result *CosUploadInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUploadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsableApisRequest struct {
	*tchttp.BaseRequest

	// 微服务id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 分组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeUsableApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsableApisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsableApisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Api信息列表
		Result *TsfPageApiDetailInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsableApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsableApisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsableGatewayGroupsRequest struct {
	*tchttp.BaseRequest

	// API分组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeUsableGatewayGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsableGatewayGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsableGatewayGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网关部署组信息列表
		Result *TsfPageGatewayDeployGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsableGatewayGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsableGatewayGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest

	// TZoneId
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`

	// TZoneName
	TZoneName *string `json:"TZoneName,omitempty" name:"TZoneName"`

	// TRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`
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

		// Result
		Result *ListTsfZoneResult `json:"Result,omitempty" name:"Result"`

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

type DisableCircuitBreakerRuleRequest struct {
	*tchttp.BaseRequest

	// 熔断规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DisableCircuitBreakerRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCircuitBreakerRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableCircuitBreakerRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCircuitBreakerRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCircuitBreakerRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableFallbackRouteRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *DisableFallbackRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableFallbackRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableFallbackRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 停用微服务路由保护策略操作是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableFallbackRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableFallbackRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableNamespaceAffinityRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DisableNamespaceAffinityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableNamespaceAffinityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableNamespaceAffinityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 停用命名空间就近访问策略操作是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableNamespaceAffinityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableNamespaceAffinityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableRouteRequest struct {
	*tchttp.BaseRequest

	// 停用的服务路由ID
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
}

func (r *DisableRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 停用服务路由操作是否成功。
	// true：停用操作成功。
	// false：停用操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableRouteRuleRequest struct {
	*tchttp.BaseRequest

	// 路由规则id
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`
}

func (r *DisableRouteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableRouteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableRouteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableRouteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableRouteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// 业务日志配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList" list`

	// TSF分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DisassociateBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DockerForUse struct {

	// 仓库中心地址
	Server *string `json:"Server,omitempty" name:"Server"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`
}

type DownloadMsApiRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 导出格式类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DownloadMsApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadMsApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadMsApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出内容，base64 编码的 zip 包
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadMsApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadMsApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadTemplateRequest struct {
	*tchttp.BaseRequest

	// 工程id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DownloadTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回base64
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DraftApiGroupRequest struct {
	*tchttp.BaseRequest

	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DraftApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DraftApiGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DraftApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true: 成功, false: 失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DraftApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DraftApiGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DraftPluginRequest struct {
	*tchttp.BaseRequest

	// 插件ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DraftPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DraftPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DraftPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DraftPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DraftPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DscribeTasksRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DscribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DscribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DscribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 变更记录分页列表
		Result *TsfPageVmTask `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DscribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DscribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DumpResult struct {

	// Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// State
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// Detail
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type EnableCircuitBreakerRuleRequest struct {
	*tchttp.BaseRequest

	// 熔断规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *EnableCircuitBreakerRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCircuitBreakerRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableCircuitBreakerRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true、false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCircuitBreakerRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCircuitBreakerRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableFallbackRouteRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *EnableFallbackRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableFallbackRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableFallbackRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 启用微服务路由保护策略操作是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableFallbackRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableFallbackRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableNamespaceAffinityRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *EnableNamespaceAffinityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableNamespaceAffinityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableNamespaceAffinityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 启用命名空间就近访问策略操作是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableNamespaceAffinityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableNamespaceAffinityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableRouteRequest struct {
	*tchttp.BaseRequest

	// 路由ID
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// 路由启用描述
	ReleaseLog *string `json:"ReleaseLog,omitempty" name:"ReleaseLog"`
}

func (r *EnableRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 启用路由操作是否成功。
	// true：启用路由成功。
	// false：启用路由失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableRouteRuleRequest struct {
	*tchttp.BaseRequest

	// 路由规则Id
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`

	// 路由规则启用描述
	ReleaseLog *string `json:"ReleaseLog,omitempty" name:"ReleaseLog"`
}

func (r *EnableRouteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableRouteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableRouteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableRouteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableRouteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Env struct {

	// 环境变量名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务端口
	Value *string `json:"Value,omitempty" name:"Value"`
}

type EnvsV2 struct {

	// 环境变量名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务端口
	Value *string `json:"Value,omitempty" name:"Value"`
}

type EventPolicyResult struct {

	// TriggerCondition
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerCondition *int64 `json:"TriggerCondition,omitempty" name:"TriggerCondition"`

	// Frequency
	// 注意：此字段可能返回 null，表示取不到有效值。
	Frequency *int64 `json:"Frequency,omitempty" name:"Frequency"`

	// EventPolicyId
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventPolicyId *string `json:"EventPolicyId,omitempty" name:"EventPolicyId"`

	// PolicyId
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

type ExpandGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 扩容的机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *ExpandGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExpandGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExpandGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpandGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExpandGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExpandNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 机器id列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`

	// 容器系统重装，系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// cvm系统重装，镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 秘钥
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组
	SgId *string `json:"SgId,omitempty" name:"SgId"`
}

func (r *ExpandNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExpandNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExpandNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpandNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExpandNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExpandSubrule struct {

	// Indicators
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicators *int64 `json:"Indicators,omitempty" name:"Indicators"`

	// IndicatorType
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndicatorType *int64 `json:"IndicatorType,omitempty" name:"IndicatorType"`

	// RuleType
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
}

type FileConfig struct {

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项版本描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFileName *string `json:"ConfigFileName,omitempty" name:"ConfigFileName"`

	// 配置项文件内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFileValue *string `json:"ConfigFileValue,omitempty" name:"ConfigFileValue"`

	// 配置项文件编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFileCode *string `json:"ConfigFileCode,omitempty" name:"ConfigFileCode"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 配置项归属应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 配置项版本数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionCount *int64 `json:"ConfigVersionCount,omitempty" name:"ConfigVersionCount"`

	// 配置项最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 发布路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFilePath *string `json:"ConfigFilePath,omitempty" name:"ConfigFilePath"`

	// 后置命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigPostCmd *string `json:"ConfigPostCmd,omitempty" name:"ConfigPostCmd"`
}

type FileConfigRelease struct {

	// 配置项发布ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type FileConfigReleaseLog struct {

	// 配置项发布日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 上次配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigId *string `json:"LastConfigId,omitempty" name:"LastConfigId"`

	// 上次配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigName *string `json:"LastConfigName,omitempty" name:"LastConfigName"`

	// 上次配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigVersion *string `json:"LastConfigVersion,omitempty" name:"LastConfigVersion"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 回滚标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollbackFlag *bool `json:"RollbackFlag,omitempty" name:"RollbackFlag"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type FindContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 无
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *FindContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 复杂对象
		Result *FindContainerGroupResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FindContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindContainerGroupResult struct {

	// groupId
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 已启动实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 负载均衡ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

	// Service ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// NodePort端口，只有公网和NodePort访问方式才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePort *string `json:"NodePort,omitempty" name:"NodePort"`

	// 最大分配cpu 核数，如0.6
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配内存M数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 0:公网 1:集群内访问 2：NodePort
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 更新方式：0:快速更新 1:滚动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 端口数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 环境变量数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envs []*Env `json:"Envs,omitempty" name:"Envs" list`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`
}

type FindContainerGroupsRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 关键字，按照分组名
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，默认为0，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *FindContainerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindContainerGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindContainerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FindContainerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindContainerGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindDeployModuleParamsRequest struct {
	*tchttp.BaseRequest

	// ParameterConfigType
	ParameterConfigType *string `json:"ParameterConfigType,omitempty" name:"ParameterConfigType"`

	// ParameterSubType
	ParameterSubType *string `json:"ParameterSubType,omitempty" name:"ParameterSubType"`
}

func (r *FindDeployModuleParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindDeployModuleParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindDeployModuleParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result []*ModuleParamResult `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FindDeployModuleParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindDeployModuleParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindMonitorObjectRequest struct {
	*tchttp.BaseRequest

	// 关键词id;
	KeywordsId *string `json:"KeywordsId,omitempty" name:"KeywordsId"`

	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *FindMonitorObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindMonitorObjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindMonitorObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
		Result *Result `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FindMonitorObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindMonitorObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindServiceMonitorObjectRequest struct {
	*tchttp.BaseRequest

	// NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// MicroserviceName
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`
}

func (r *FindServiceMonitorObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindServiceMonitorObjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FindServiceMonitorObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *Result `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FindServiceMonitorObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FindServiceMonitorObjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GatewayApiGroupVo struct {

	// 分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组下API个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupApiCount *uint64 `json:"GroupApiCount,omitempty" name:"GroupApiCount"`

	// 分组API列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupApis []*GatewayGroupApiVo `json:"GroupApis,omitempty" name:"GroupApis" list`
}

type GatewayDeployGroup struct {

	// 网关部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 网关部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用分类：V：虚拟机应用，C：容器应用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组应用状态,取值: Running、Waiting、Paused、Updating、RollingBack、Abnormal、Unknown
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`

	// 集群类型，C ：容器，V：虚拟机
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type GatewayGroupApiVo struct {

	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// API 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`
}

type GatewayGroupIds struct {

	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 分组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type GatewayJwtPlugin struct {

	// 网关插件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 插件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 插件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件创建时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 插件更新时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 公钥标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kid *string `json:"Kid,omitempty" name:"Kid"`

	// 公钥对
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicKeyJson *string `json:"PublicKeyJson,omitempty" name:"PublicKeyJson"`

	// 重定向地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// token携带位置，网关取token位置与发送认证请求时token位置一致,值:query/header
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenBaggagePosition *string `json:"TokenBaggagePosition,omitempty" name:"TokenBaggagePosition"`

	// token的key值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenKeyName *string `json:"TokenKeyName,omitempty" name:"TokenKeyName"`

	// claim参数映射关系json
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClaimMappingJson *string `json:"ClaimMappingJson,omitempty" name:"ClaimMappingJson"`

	// 发布状态: drafted/released
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type GatewayNamespace struct {

	// 命名空间主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceCode *string `json:"NamespaceCode,omitempty" name:"NamespaceCode"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceResourceType *string `json:"NamespaceResourceType,omitempty" name:"NamespaceResourceType"`
}

type GatewayOAuthPlugin struct {

	// 网关插件id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 插件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 插件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件创建时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 插件更新时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 验证token路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenAuthUrl *string `json:"TokenAuthUrl,omitempty" name:"TokenAuthUrl"`

	// 验证token请求方法:get/post
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenAuthMethod *string `json:"TokenAuthMethod,omitempty" name:"TokenAuthMethod"`

	// 认证请求超时时间,单位:秒 范围:0~30
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 重定向地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// token携带位置，网关取token位置与发送认证请求时token位置一致,值:query/header
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenBaggagePosition *string `json:"TokenBaggagePosition,omitempty" name:"TokenBaggagePosition"`

	// token的key值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenKeyName *string `json:"TokenKeyName,omitempty" name:"TokenKeyName"`

	// payload的映射参数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayloadMappingName *string `json:"PayloadMappingName,omitempty" name:"PayloadMappingName"`

	// payload映射到后端服务的携带位置,值:query/header
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayloadMappingPosition *string `json:"PayloadMappingPosition,omitempty" name:"PayloadMappingPosition"`

	// 发布状态: drafted/released
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type GatewayPlugin struct {

	// 网关插件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 插件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type GatewayPluginBoundParam struct {

	// 插件id
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 插件绑定到的对象类型:group/api
	ScopeType *string `json:"ScopeType,omitempty" name:"ScopeType"`

	// 插件绑定到的对象主键值，例如分组的ID/API的ID
	ScopeValue *string `json:"ScopeValue,omitempty" name:"ScopeValue"`
}

type GatewayPluginDefinition struct {

	// 网关插件定义id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 插件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 插件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 插件创建时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 插件更新时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 插件排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

type GatewayPluginId struct {

	// 网关插件ID
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
}

type GatewayTagPlugin struct {

	// 网关插件id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 插件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 插件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件创建时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 插件更新时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 发布状态: drafted/released
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 参数配置json串
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagPluginInfoList *string `json:"TagPluginInfoList,omitempty" name:"TagPluginInfoList"`
}

type GatewayVo struct {

	// 网关部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayDeployGroupName *string `json:"GatewayDeployGroupName,omitempty" name:"GatewayDeployGroupName"`

	// API 分组个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

	// API 分组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*GatewayApiGroupVo `json:"Groups,omitempty" name:"Groups" list`
}

type GetApmEsAuthInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetApmEsAuthInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetApmEsAuthInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetApmEsAuthInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApmEsAuthInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetApmEsAuthInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterLimitResourceRequest struct {
	*tchttp.BaseRequest

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetClusterLimitResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterLimitResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterLimitResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取集群limit资源返回参数
		Result *GetClusterLimitResourceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterLimitResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterLimitResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterLimitResourceResult struct {

	// 最大分配cpu 核数
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配mem内存数，单位M
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`
}

type GetConfigpropsRequest struct {
	*tchttp.BaseRequest

	// InstanceId
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// GroupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// EndpointName
	EndpointName *string `json:"EndpointName,omitempty" name:"EndpointName"`
}

func (r *GetConfigpropsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConfigpropsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConfigpropsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConfigpropsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConfigpropsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetContainGroupDeployInfo struct {

	// GroupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例总数
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动的实例数
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 镜像server
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 预分配cpu 核数，如0.2，
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 最大分配cpu 核数，如0.6
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 预分配内存M数
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 最大分配内存M数
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 数组
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// jvm参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 无用
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 无用
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 无用
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 无用
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 无用
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 无用
	KuberneteApiServer *string `json:"KuberneteApiServer,omitempty" name:"KuberneteApiServer"`

	// 无用
	KuberneteUser *string `json:"KuberneteUser,omitempty" name:"KuberneteUser"`

	// 无用
	KubernetePassword *string `json:"KubernetePassword,omitempty" name:"KubernetePassword"`

	// 无用
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 无用
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 无用
	GroupComment *string `json:"GroupComment,omitempty" name:"GroupComment"`

	// 无用
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// 无用
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 无用
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// 无用
	LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

	// 无用
	NodePort *string `json:"NodePort,omitempty" name:"NodePort"`

	// 无用
	IsStop *string `json:"IsStop,omitempty" name:"IsStop"`

	// 无用
	Status *string `json:"Status,omitempty" name:"Status"`

	// 无用
	Message *string `json:"Message,omitempty" name:"Message"`

	// 无用
	ChangType *int64 `json:"ChangType,omitempty" name:"ChangType"`

	// 无用
	ChangNum *int64 `json:"ChangNum,omitempty" name:"ChangNum"`

	// 无用
	IsFirst *int64 `json:"IsFirst,omitempty" name:"IsFirst"`

	// 无用
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 无用
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 无用
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 无用
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// 无用
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 无用
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 无用
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
}

type GetContainGroupDeployInfoRequest struct {
	*tchttp.BaseRequest

	// 实例所属groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetContainGroupDeployInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetContainGroupDeployInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetContainGroupDeployInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的复杂对象
		Result *GetContainGroupDeployInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetContainGroupDeployInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetContainGroupDeployInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetContainGroupOtherRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetContainGroupOtherRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetContainGroupOtherRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetContainGroupOtherResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取部署组其他字段-前端并发调用场景
		Result *GetContainGroupOtherResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetContainGroupOtherResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetContainGroupOtherResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetContainGroupOtherResult struct {

	// 实例总数
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动实例总数
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 负载均衡ip
	LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

	// Service ip
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// 服务状态，请参考后面的的状态定义
	Status *string `json:"Status,omitempty" name:"Status"`

	// 异常信息,服务状态为Abnormal时才有
	Message *string `json:"Message,omitempty" name:"Message"`
}

type GetDefaultValueRequest struct {
	*tchttp.BaseRequest

	// 解析规则类型
	SchemaType *uint64 `json:"SchemaType,omitempty" name:"SchemaType"`
}

func (r *GetDefaultValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDefaultValueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDefaultValueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 按照入参返回的默认值
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDefaultValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDefaultValueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDockerForUseRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDockerForUseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDockerForUseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDockerForUseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 仓库中心地址
		Server *string `json:"Server,omitempty" name:"Server"`

		// 用户名
		Username *string `json:"Username,omitempty" name:"Username"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDockerForUseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDockerForUseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDownloadInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
}

func (r *GetDownloadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDownloadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDownloadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Result *CosDownloadInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDownloadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDownloadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDumpRequest struct {
	*tchttp.BaseRequest

	// InstanceId
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// GroupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// EndpointName
	EndpointName *string `json:"EndpointName,omitempty" name:"EndpointName"`
}

func (r *GetDumpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDumpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDumpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result []*DumpResult `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDumpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDumpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEnvRequest struct {
	*tchttp.BaseRequest

	// InstanceId
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// GroupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// EndpointName
	EndpointName *string `json:"EndpointName,omitempty" name:"EndpointName"`
}

func (r *GetEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMetricsRequest struct {
	*tchttp.BaseRequest

	// InstanceId
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// GroupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// EndpointName
	EndpointName *string `json:"EndpointName,omitempty" name:"EndpointName"`
}

func (r *GetMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMetricsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOssTopologyGraphRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询开始时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetOssTopologyGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOssTopologyGraphRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOssTopologyGraphResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务（节点）集合
		Result []*GraphNodeV2 `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOssTopologyGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOssTopologyGraphResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOssTraceInterfacesRequest struct {
	*tchttp.BaseRequest

	// 接口所属服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 接口角色："caller"调用方，"callee"被调用方
	Role *string `json:"Role,omitempty" name:"Role"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 返回调用链接口偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回调用链接口数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetOssTraceInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOssTraceInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOssTraceInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务接口集合
		Result *TsfPageTraceInterface `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOssTraceInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOssTraceInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOssTraceServicesRequest struct {
	*tchttp.BaseRequest

	// 服务角色："caller"调用方，"callee"被调用方
	Role *string `json:"Role,omitempty" name:"Role"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 返回服务偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回服务数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetOssTraceServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOssTraceServicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOssTraceServicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务名集合
		Result *TsfPageTraceService `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOssTraceServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOssTraceServicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOssTraceSpansRequest struct {
	*tchttp.BaseRequest

	// 调用链ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`
}

func (r *GetOssTraceSpansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOssTraceSpansRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetOssTraceSpansResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调用链Span列表
		Result *TsfPageZipkinSpanInfoV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOssTraceSpansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetOssTraceSpansResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPkgInfoRequest struct {
	*tchttp.BaseRequest

	// 无
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 无
	PkgId []*string `json:"PkgId,omitempty" name:"PkgId" list`
}

func (r *GetPkgInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPkgInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPkgInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		PkgInfo []*string `json:"PkgInfo,omitempty" name:"PkgInfo" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPkgInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPkgInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatisticsRequest struct {
	*tchttp.BaseRequest

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetServiceStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetServiceStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务运行统计指标
		Result *ServiceStatistics `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetServiceStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopAvgTimeCostInterfacesRequest struct {
	*tchttp.BaseRequest

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 列表条目数量，取值范围[1, 50]，默认值10
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// TSF命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// TSF微服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *GetTopAvgTimeCostInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopAvgTimeCostInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopAvgTimeCostInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TopN条目列表
		Result []*StatisticsEntry `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopAvgTimeCostInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopAvgTimeCostInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopAvgTimeCostServicesRequest struct {
	*tchttp.BaseRequest

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 列表条目数量，取值范围[1, 50]，默认值10
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// TSF命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *GetTopAvgTimeCostServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopAvgTimeCostServicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopAvgTimeCostServicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TopN条目列表
		Result []*StatisticsEntry `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopAvgTimeCostServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopAvgTimeCostServicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopFailureRateInterfacesRequest struct {
	*tchttp.BaseRequest

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 列表条目数量，取值范围[1, 50]，默认值10
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// TSF命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// TSF微服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *GetTopFailureRateInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopFailureRateInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopFailureRateInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TopN条目列表
		Result []*StatisticsEntry `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopFailureRateInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopFailureRateInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopFailureRateServicesRequest struct {
	*tchttp.BaseRequest

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 列表条目数量，取值范围[1, 50]，默认值10
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// TSF命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *GetTopFailureRateServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopFailureRateServicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopFailureRateServicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TopN条目列表
		Result []*StatisticsEntry `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopFailureRateServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopFailureRateServicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopReqAmountInterfacesRequest struct {
	*tchttp.BaseRequest

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 列表条目数量，取值范围[1, 50]，默认值10
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// TSF命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// TSF微服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *GetTopReqAmountInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopReqAmountInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopReqAmountInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TopN条目列表
		Result []*StatisticsEntry `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopReqAmountInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopReqAmountInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopReqAmountServicesRequest struct {
	*tchttp.BaseRequest

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 列表条目数量，取值范围[1, 50]，默认值10
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// TSF命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *GetTopReqAmountServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopReqAmountServicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopReqAmountServicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TopN条目列表
		Result []*StatisticsEntry `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopReqAmountServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopReqAmountServicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopologyGraphRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 全局命名空间ID
	GlobalNamespaceId *string `json:"GlobalNamespaceId,omitempty" name:"GlobalNamespaceId"`
}

func (r *GetTopologyGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopologyGraphRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopologyGraphResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务（节点）集合
		Result []*GraphNodeV2 `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopologyGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopologyGraphResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTraceInterfacesRequest struct {
	*tchttp.BaseRequest

	// 接口所属服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 接口角色："caller"调用方，"callee"被调用方
	Role *string `json:"Role,omitempty" name:"Role"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 返回调用链接口偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回调用链接口数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetTraceInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTraceInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTraceInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务接口集合
		Result *TsfPageTraceInterface `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTraceInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTraceInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTraceRequest struct {
	*tchttp.BaseRequest

	// InstanceId
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// GroupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// EndpointName
	EndpointName *string `json:"EndpointName,omitempty" name:"EndpointName"`
}

func (r *GetTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTraceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTraceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTraceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTraceServicesRequest struct {
	*tchttp.BaseRequest

	// 服务角色："caller"调用方，"callee"被调用方
	Role *string `json:"Role,omitempty" name:"Role"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 返回服务偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回服务数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetTraceServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTraceServicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTraceServicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务名集合
		Result *TsfPageTraceService `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTraceServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTraceServicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTraceSpansRequest struct {
	*tchttp.BaseRequest

	// 调用链ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`
}

func (r *GetTraceSpansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTraceSpansRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTraceSpansResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调用链Span列表
		Result *TsfPageZipkinSpanInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTraceSpansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTraceSpansResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUploadInfoRequest struct {
	*tchttp.BaseRequest

	// 无
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 无
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 无
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 无
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 无
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`
}

func (r *GetUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUploadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Result *CosUploadInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUploadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GraphEdgeV2 struct {

	// 来源服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`

	// 目标服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`

	// 来源服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 目标服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 服务间总请求量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqTotalQty *uint64 `json:"ReqTotalQty,omitempty" name:"ReqTotalQty"`

	// 请求平均数，次每分钟，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqAvgQty *float64 `json:"ReqAvgQty,omitempty" name:"ReqAvgQty"`

	// 请求成功率，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqSuccessRate *float64 `json:"ReqSuccessRate,omitempty" name:"ReqSuccessRate"`

	// 请求平均耗时，单位毫秒，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqAvgDuration *float64 `json:"ReqAvgDuration,omitempty" name:"ReqAvgDuration"`

	// 健康度指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Apdex *float64 `json:"Apdex,omitempty" name:"Apdex"`

	// 来源节点标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 目标节点标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 来源命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceNamespaceId *string `json:"SourceNamespaceId,omitempty" name:"SourceNamespaceId"`

	// 目标命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`
}

type GraphNodeV2 struct {

	// 节点服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 节点服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 节点服务总请求量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqTotalQty *uint64 `json:"ReqTotalQty,omitempty" name:"ReqTotalQty"`

	// 节点请求平均数，次每分钟，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqAvgQty *float64 `json:"ReqAvgQty,omitempty" name:"ReqAvgQty"`

	// 节点请求成功率，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqSuccessRate *float64 `json:"ReqSuccessRate,omitempty" name:"ReqSuccessRate"`

	// 节点请求平均耗时，单位毫秒，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqAvgDuration *float64 `json:"ReqAvgDuration,omitempty" name:"ReqAvgDuration"`

	// 节点服务为目的的边列表（入度）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeList []*GraphEdgeV2 `json:"EdgeList,omitempty" name:"EdgeList" list`

	// 节点标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 节点命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

type GroupApiUseStatistics struct {

	// 总调用数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopStatusCode []*ApiUseStatisticsEntity `json:"TopStatusCode,omitempty" name:"TopStatusCode" list`

	// 平均错误率
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopTimeCost []*ApiUseStatisticsEntity `json:"TopTimeCost,omitempty" name:"TopTimeCost" list`

	// 分位值对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quantile *QuantileEntity `json:"Quantile,omitempty" name:"Quantile"`
}

type GroupDailyUseStatistics struct {

	// 总调用数
	TopReqAmount []*GroupUseStatisticsEntity `json:"TopReqAmount,omitempty" name:"TopReqAmount" list`

	// 平均错误率
	TopFailureRate []*GroupUseStatisticsEntity `json:"TopFailureRate,omitempty" name:"TopFailureRate" list`

	// 平均响应耗时
	TopAvgTimeCost []*GroupUseStatisticsEntity `json:"TopAvgTimeCost,omitempty" name:"TopAvgTimeCost" list`
}

type GroupPod struct {

	// 实例名称(对应到kubernetes的pod名称)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 实例ID(对应到kubernetes的pod id)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// 实例状态，请参考后面的实例以及容器的状态定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例处于当前状态的原因，例如容器下载镜像失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 实例中容器的重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// 实例中已就绪容器的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadyCount *int64 `json:"ReadyCount,omitempty" name:"ReadyCount"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 实例启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 服务实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 机器实例可使用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 机器实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
}

type GroupPodResult struct {

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*GroupPod `json:"Content,omitempty" name:"Content" list`
}

type GroupResourceConfig struct {

	// 容器部署组相关的参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *ContainerGroupResourceConfig `json:"Container,omitempty" name:"Container"`
}

type GroupUseStatisticsEntity struct {

	// API 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiPath *string `json:"ApiPath,omitempty" name:"ApiPath"`

	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type GroupsByScalableRuleId struct {

	// ApplicationId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// ApplicationName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// GroupId
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// GroupName
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Status
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// UpdateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 弹性伸缩规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 弹性伸缩规则名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type GroupsByScalableRuleIdList struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*GroupsByScalableRuleId `json:"Content,omitempty" name:"Content" list`
}

type ImageDeleteTagRequest struct {
	*tchttp.BaseRequest

	// 仓库名，如/tsf/nginx
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 版本号:如V1
	TagName *string `json:"TagName,omitempty" name:"TagName"`
}

func (r *ImageDeleteTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageDeleteTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageDeleteTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageDeleteTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageDeleteTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageGetRepositoryListRequest struct {
	*tchttp.BaseRequest

	// 仓库名，搜索关键字,不带命名空间的
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ImageGetRepositoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageGetRepositoryListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageGetRepositoryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 镜像服务器地址
		Server *int64 `json:"Server,omitempty" name:"Server"`

		// 列表信息
		Content []*string `json:"Content,omitempty" name:"Content" list`

		// 仓库名,含命名空间,如tsf/nginx
		Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

		// 命名空间,tsf
		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

		// 镜像版本数
		TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

		// 是否公共,1:公有,0:私有
		IsPublic *int64 `json:"IsPublic,omitempty" name:"IsPublic"`

		// 是否被用户收藏。true：是，false：否
		IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`

		// 是否是Tce官方仓库。 是否是Tce官方仓库。true：是，false：否
		IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`

		// 被所有用户收藏次数
		FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

		// 拉取次数
		PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

		// 描述内容
		Description *string `json:"Description,omitempty" name:"Description"`

		// 创建时间
		CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

		// 更新时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageGetRepositoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageGetRepositoryListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageGetTagListRequest struct {
	*tchttp.BaseRequest

	// 应用Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ImageGetTagListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageGetTagListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageGetTagListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 仓库名,含命名空间,如tsf/nginx
		Reponame *int64 `json:"Reponame,omitempty" name:"Reponame"`

		// 镜像服务器地址
		Server *int64 `json:"Server,omitempty" name:"Server"`

		// 列表信息(见下面字段)
		Content []*string `json:"Content,omitempty" name:"Content" list`

		// 仓库名，如nginx
		Repo_name *string `json:"Repo_name,omitempty" name:"Repo_name"`

		// 版本名称
		TagName *string `json:"TagName,omitempty" name:"TagName"`

		// 版本ID
		TagId *string `json:"TagId,omitempty" name:"TagId"`

		// 镜像ID
		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

		// 大小
		Size *string `json:"Size,omitempty" name:"Size"`

		// 创建时间
		CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

		// 更新时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 镜像制作者
		Author *string `json:"Author,omitempty" name:"Author"`

		// CPU架构
		Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

		// Docker客户端版本
		DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

		// 操作系统
		Os *string `json:"Os,omitempty" name:"Os"`

		// push时间
		PushTime *string `json:"PushTime,omitempty" name:"PushTime"`

		// 单位为字节
		SizeByte *string `json:"SizeByte,omitempty" name:"SizeByte"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageGetTagListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageGetTagListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageRegisterRequest struct {
	*tchttp.BaseRequest

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 确认密码
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`
}

func (r *ImageRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageRegisterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageRegisterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageRegisterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageRepository struct {

	// 仓库名,含命名空间,如tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Repotype *string `json:"Repotype,omitempty" name:"Repotype"`

	// 镜像版本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

	// 是否公共,1:公有,0:私有
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPublic *int64 `json:"IsPublic,omitempty" name:"IsPublic"`

	// 是否被用户收藏。true：是，false：否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`

	// 是否是Tce官方仓库。 是否是Tce官方仓库。true：是，false：否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`

	// 被所有用户收藏次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 拉取次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

	// 描述内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ImageRepositoryResult struct {

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像服务器地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ImageRepository `json:"Content,omitempty" name:"Content" list`
}

type ImageTag struct {

	// 仓库名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 版本名称
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 版本ID
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 大小
	Size *string `json:"Size,omitempty" name:"Size"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 镜像制作者
	Author *string `json:"Author,omitempty" name:"Author"`

	// CPU架构
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// Docker客户端版本
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// 操作系统
	Os *string `json:"Os,omitempty" name:"Os"`

	// push时间
	PushTime *string `json:"PushTime,omitempty" name:"PushTime"`

	// 单位为字节
	SizeByte *int64 `json:"SizeByte,omitempty" name:"SizeByte"`
}

type ImageTagsResult struct {

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库名,含命名空间,如tsf/ngin
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像服务器地址
	Server *string `json:"Server,omitempty" name:"Server"`

	// 列表信息
	Content []*ImageTag `json:"Content,omitempty" name:"Content" list`
}

type ImageUserIsExists struct {

	// 子账号是否存在,true：存在；false：不存在
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`

	// 主账号是否存在，true：存在；false：不存在
	MainIsExist *bool `json:"MainIsExist,omitempty" name:"MainIsExist"`
}

type ImageUserIsExistsRequest struct {
	*tchttp.BaseRequest
}

func (r *ImageUserIsExistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageUserIsExistsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageUserIsExistsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：存在；false：不存在
		IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageUserIsExistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageUserIsExistsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageUserIsExistsResult struct {

	// 用户是否存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ImageUserIsExists `json:"Data,omitempty" name:"Data"`
}

type IndicatorCoord struct {

	// 指标横坐标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoordX *string `json:"CoordX,omitempty" name:"CoordX"`

	// 指标纵坐标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoordY *string `json:"CoordY,omitempty" name:"CoordY"`

	// 指标标签，用于标识附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoordTag *string `json:"CoordTag,omitempty" name:"CoordTag"`
}

type InitializeApmRequest struct {
	*tchttp.BaseRequest
}

func (r *InitializeApmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitializeApmRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitializeApmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitializeApmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitializeApmResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstallAgentScriptsRequest struct {
	*tchttp.BaseRequest

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *InstallAgentScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstallAgentScriptsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstallAgentScriptsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安装agent脚本内容
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallAgentScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstallAgentScriptsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 机器实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 机器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 机器内网地址IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 机器外网地址IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 机器描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// VM的状态 虚机：虚机的状态 容器：Pod所在虚机的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// VM的可使用状态 虚机：虚机是否能够作为资源使用 容器：虚机是否能够作为资源部署POD
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 服务下的服务实例的状态 虚机：应用是否可用 + Agent状态 容器：Pod状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 标识此instance是否已添加在tsf中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountInTsf *int64 `json:"CountInTsf,omitempty" name:"CountInTsf"`

	// 机器所属部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器所属应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 机器所属应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 机器实例在CVM的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCreatedTime *string `json:"InstanceCreatedTime,omitempty" name:"InstanceCreatedTime"`

	// 机器实例在CVM的过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceExpiredTime *string `json:"InstanceExpiredTime,omitempty" name:"InstanceExpiredTime"`

	// 机器实例在CVM的计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 机器实例总CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTotalCpu *float64 `json:"InstanceTotalCpu,omitempty" name:"InstanceTotalCpu"`

	// 机器实例总内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTotalMem *float64 `json:"InstanceTotalMem,omitempty" name:"InstanceTotalMem"`

	// 机器实例使用的CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUsedCpu *float64 `json:"InstanceUsedCpu,omitempty" name:"InstanceUsedCpu"`

	// 机器实例使用的内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUsedMem *float64 `json:"InstanceUsedMem,omitempty" name:"InstanceUsedMem"`

	// 机器实例Limit CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLimitCpu *float64 `json:"InstanceLimitCpu,omitempty" name:"InstanceLimitCpu"`

	// 机器实例Limit 内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLimitMem *float64 `json:"InstanceLimitMem,omitempty" name:"InstanceLimitMem"`

	// 包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePkgVersion *string `json:"InstancePkgVersion,omitempty" name:"InstancePkgVersion"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 机器实例业务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例执行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationState *int64 `json:"OperationState,omitempty" name:"OperationState"`
}

type InstanceResourceConfig struct {

	// 容器实例相关的参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *ContainerInstanceResourceConfig `json:"Container,omitempty" name:"Container"`

	// 虚拟机实例相关的参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vm *VmInstanceResourceConfig `json:"Vm,omitempty" name:"Vm"`
}

type InstancesResult struct {

	// Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Status
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// Ports
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ports []*PortsResult `json:"Ports,omitempty" name:"Ports" list`
}

type InvocationIndicator struct {

	// 总请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationQuantity *int64 `json:"InvocationQuantity,omitempty" name:"InvocationQuantity"`

	// 请求成功率，百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationSuccessRate *float64 `json:"InvocationSuccessRate,omitempty" name:"InvocationSuccessRate"`

	// 请求平均耗时，单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationAvgDuration *float64 `json:"InvocationAvgDuration,omitempty" name:"InvocationAvgDuration"`

	// 成功请求数时间分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationSuccessDistribution []*IndicatorCoord `json:"InvocationSuccessDistribution,omitempty" name:"InvocationSuccessDistribution" list`

	// 失败请求数时间分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationFailedDistribution []*IndicatorCoord `json:"InvocationFailedDistribution,omitempty" name:"InvocationFailedDistribution" list`

	// 状态码分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationStatusDistribution []*IndicatorCoord `json:"InvocationStatusDistribution,omitempty" name:"InvocationStatusDistribution" list`

	// 时延分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationDurationDistribution []*IndicatorCoord `json:"InvocationDurationDistribution,omitempty" name:"InvocationDurationDistribution" list`

	// 并发请求次数时间分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationQuantityDistribution []*IndicatorCoord `json:"InvocationQuantityDistribution,omitempty" name:"InvocationQuantityDistribution" list`
}

type InvocationStatisticsRatio struct {

	// 当前统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationCurrentStatistics *InvocationStatisticsV2 `json:"InvocationCurrentStatistics,omitempty" name:"InvocationCurrentStatistics"`

	// 目标统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationTargetStatistics *InvocationStatisticsV2 `json:"InvocationTargetStatistics,omitempty" name:"InvocationTargetStatistics"`

	// 请求总数对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationSumQuantityRatio *float64 `json:"InvocationSumQuantityRatio,omitempty" name:"InvocationSumQuantityRatio"`

	// 请求平均耗时对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationAvgDurationRatio *float64 `json:"InvocationAvgDurationRatio,omitempty" name:"InvocationAvgDurationRatio"`

	// 2xx状态码响应请求数对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invocation2xxStatusQuantityRatio *float64 `json:"Invocation2xxStatusQuantityRatio,omitempty" name:"Invocation2xxStatusQuantityRatio"`

	// 4xx状态码响应请求数对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invocation4xxStatusQuantityRatio *float64 `json:"Invocation4xxStatusQuantityRatio,omitempty" name:"Invocation4xxStatusQuantityRatio"`

	// 5xx状态码响应请求数对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invocation5xxStatusQuantityRatio *float64 `json:"Invocation5xxStatusQuantityRatio,omitempty" name:"Invocation5xxStatusQuantityRatio"`

	// 其它状态码响应请求数对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationOtherStatusQuantityRatio *float64 `json:"InvocationOtherStatusQuantityRatio,omitempty" name:"InvocationOtherStatusQuantityRatio"`

	// 连接失败请求数对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationConnectFailedQuantityRatio *float64 `json:"InvocationConnectFailedQuantityRatio,omitempty" name:"InvocationConnectFailedQuantityRatio"`

	// 超时请求数对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationTimeoutQuantityRatio *float64 `json:"InvocationTimeoutQuantityRatio,omitempty" name:"InvocationTimeoutQuantityRatio"`

	// 服务不可用请求数对比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationUnavailableQuantityRatio *float64 `json:"InvocationUnavailableQuantityRatio,omitempty" name:"InvocationUnavailableQuantityRatio"`
}

type InvocationStatisticsV2 struct {

	// 调用统计名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationStatisticsName *string `json:"InvocationStatisticsName,omitempty" name:"InvocationStatisticsName"`

	// 请求总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationSumQuantity *int64 `json:"InvocationSumQuantity,omitempty" name:"InvocationSumQuantity"`

	// 请求平均耗时，单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationAvgDuration *float64 `json:"InvocationAvgDuration,omitempty" name:"InvocationAvgDuration"`

	// 2xx状态码响应请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invocation2xxStatusQuantity *int64 `json:"Invocation2xxStatusQuantity,omitempty" name:"Invocation2xxStatusQuantity"`

	// 4xx状态码响应请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invocation4xxStatusQuantity *int64 `json:"Invocation4xxStatusQuantity,omitempty" name:"Invocation4xxStatusQuantity"`

	// 5xx状态码响应请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invocation5xxStatusQuantity *int64 `json:"Invocation5xxStatusQuantity,omitempty" name:"Invocation5xxStatusQuantity"`

	// 其它状态码响应请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationOtherStatusQuantity *int64 `json:"InvocationOtherStatusQuantity,omitempty" name:"InvocationOtherStatusQuantity"`

	// 连接失败请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationConnectFailedQuantity *int64 `json:"InvocationConnectFailedQuantity,omitempty" name:"InvocationConnectFailedQuantity"`

	// 超时请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationTimeoutQuantity *int64 `json:"InvocationTimeoutQuantity,omitempty" name:"InvocationTimeoutQuantity"`

	// 服务不可用请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationUnavailableQuantity *int64 `json:"InvocationUnavailableQuantity,omitempty" name:"InvocationUnavailableQuantity"`

	// 请求数占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationQuantityProportion *float64 `json:"InvocationQuantityProportion,omitempty" name:"InvocationQuantityProportion"`
}

type IpListResult struct {

	// InstanceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// UserName
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Password
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// machineId
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineId *string `json:"MachineId,omitempty" name:"MachineId"`
}

type IssueLicenseRequest struct {
	*tchttp.BaseRequest

	// 授予对象
	Grantee *LicenseGrantee `json:"Grantee,omitempty" name:"Grantee"`

	// 申请时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

func (r *IssueLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IssueLicenseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IssueLicenseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书正文
	// 注意：此字段可能返回 null，表示取不到有效值。
		LicenseContent *string `json:"LicenseContent,omitempty" name:"LicenseContent"`
	} `json:"Response"`
}

func (r *IssueLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IssueLicenseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LicenseApplicationRecord struct {

	// 创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 授予对象
	Grantee *LicenseGrantee `json:"Grantee,omitempty" name:"Grantee"`

	// 所申请的产品信息列表
	Product []*LicenseProduct `json:"Product,omitempty" name:"Product" list`

	// 联系电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone []*string `json:"Phone,omitempty" name:"Phone" list`

	// 联系邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email []*string `json:"Email,omitempty" name:"Email" list`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 授予许可时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueTime *uint64 `json:"IssueTime,omitempty" name:"IssueTime"`

	// 有效期信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *LicenseDuration `json:"Duration,omitempty" name:"Duration"`
}

type LicenseDuration struct {

	// 开始时间，utc秒数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *uint64 `json:"ActiveTime,omitempty" name:"ActiveTime"`

	// 有效时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

type LicenseFunction struct {

	// 功能名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否启用
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type LicenseGrantee struct {

	// 被授予者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type LicenseProduct struct {

	// 产品名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 产品规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 选配功能列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Function []*LicenseFunction `json:"Function,omitempty" name:"Function" list`

	// 受限资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource []*LicenseResource `json:"Resource,omitempty" name:"Resource" list`
}

type LicenseResource struct {

	// 资源名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 配额
	Quota *uint64 `json:"Quota,omitempty" name:"Quota"`
}

type ListAlarmPoliciesRequest struct {
	*tchttp.BaseRequest

	// SearchWord
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListAlarmPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAlarmPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAlarmPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *ListAlarmPoliciesResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAlarmPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAlarmPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAlarmPoliciesResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*AlarmPolicyResult `json:"Content,omitempty" name:"Content" list`
}

type ListAlarmReceiverResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*AlarmReceiverResult `json:"Content,omitempty" name:"Content" list`
}

type ListAlarmReceiversRequest struct {
	*tchttp.BaseRequest

	// PolicyId
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *ListAlarmReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAlarmReceiversRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAlarmReceiversResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *ListAlarmReceiverResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAlarmReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAlarmReceiversResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAppPkgRequest struct {
	*tchttp.BaseRequest

	// 无
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 无
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListAppPkgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAppPkgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAppPkgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Result *AppPkgList `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppPkgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAppPkgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAppRequest struct {
	*tchttp.BaseRequest

	// 无
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAppRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAppResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAppResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListApplicationServerResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ApplicationServerResult `json:"Content,omitempty" name:"Content" list`
}

type ListApplicationServersRequest struct {
	*tchttp.BaseRequest

	// ServerId
	ServerId *string `json:"ServerId,omitempty" name:"ServerId"`

	// GroupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// AgentVersion
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
}

func (r *ListApplicationServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListApplicationServersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListApplicationServersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *ListApplicationServerResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListApplicationServersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListCloudMicroServiceFindPagedListRequest struct {
	*tchttp.BaseRequest

	// NamespaceName
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *ListCloudMicroServiceFindPagedListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCloudMicroServiceFindPagedListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListCloudMicroServiceFindPagedListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *CloudMonitorMicroserviceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudMicroServiceFindPagedListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCloudMicroServiceFindPagedListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListColudMonitorStatisticsPolicyRequest struct {
	*tchttp.BaseRequest

	// 关键字搜索;
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// KeywordsId
	KeywordsId *string `json:"KeywordsId,omitempty" name:"KeywordsId"`
}

func (r *ListColudMonitorStatisticsPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListColudMonitorStatisticsPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListColudMonitorStatisticsPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CloudMonitorPolicyresult
		Result *CloudMonitorPolicyResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListColudMonitorStatisticsPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListColudMonitorStatisticsPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListContainGroupRequest struct {
	*tchttp.BaseRequest

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 搜索字段，模糊搜索groupName字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListContainGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListContainGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListContainGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// groupId
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 分组名称
		GroupName []*string `json:"GroupName,omitempty" name:"GroupName" list`

		// 实例总数
		InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

		// 已启动实例总数
		CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 镜像server
		Server *string `json:"Server,omitempty" name:"Server"`

		// 镜像名，如/tsf/nginx
		Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

		// 镜像版本名称
		TagName *string `json:"TagName,omitempty" name:"TagName"`

		// 集群id
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 集群名称
		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

		// 命名空间id
		NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

		// 命名空间名称
		NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

		// 预分配cpu 核数，如0.2
		CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

		// 最大分配cpu 核数，如0.6
		CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

		// 预分配内存M数
		MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

		// 最大分配内存M数
		MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

		// 负载均衡ip
		LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

		// Service ip
		ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListContainGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListContainGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListContainerTaskRequest struct {
	*tchttp.BaseRequest

	// 操作列表所属applicationId
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 任务执行时间范围，格式："2017-10-01 10:03:27"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务执行时间范围，格式："2017-11-01 10:03:27"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 搜索字段，模糊搜索taskId字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段，默认为updateTime（修改时间字段）， 支持id， createTime， updateTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListContainerTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListContainerTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListContainerTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 列表信息
		Content []*string `json:"Content,omitempty" name:"Content" list`

		// taskId
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 任务开始时间
		TaskTime *string `json:"TaskTime,omitempty" name:"TaskTime"`

		// 任务类型字段，0：没有任务（在此接口中，不用出现0），1：发布程序包；2.部署操作；3.扩容操作；4.启动操作；5.停止操作；6.缩容操作；
		Type *int64 `json:"Type,omitempty" name:"Type"`

		// 变更状态，0：成功 1:失败 2：执行中
		Status *string `json:"Status,omitempty" name:"Status"`

		// 分组id
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 分组名称
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 镜像名称
		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

		// 镜像版本
		ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

		// 任务详情描述
		TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`

		// 任务总个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 成功任务个数
		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

		// 失败任务个数
		FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListContainerTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListContainerTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupPodRequest struct {
	*tchttp.BaseRequest

	// 实例所属groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListGroupPodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupPodRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupPodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例名称(对应到kubernetes的pod名称)
		PodName *string `json:"PodName,omitempty" name:"PodName"`

		// 实例ID(对应到kubernetes的pod id)
		PodId *string `json:"PodId,omitempty" name:"PodId"`

		// 实例状态：
	// Running ：正常运行中
	// Waiting：等待运行中，例如正在下载镜像
	// Terminating：实例有容器正在终止
	// Terminated：实例有容器已经终止
	// NotReady：实例有容器处于未就绪状态，比如容器的健康检查失败
	// Error：其他未知异常
		Status *string `json:"Status,omitempty" name:"Status"`

		// 实例处于当前状态的原因，例如容器下载镜像失败
		Reason *string `json:"Reason,omitempty" name:"Reason"`

		// 主机IP
		NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`

		// 实例IP
		Ip *string `json:"Ip,omitempty" name:"Ip"`

		// 实例中容器的重启次数
		RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

		// 实例中已就绪容器的个数
		ReadyCount *int64 `json:"ReadyCount,omitempty" name:"ReadyCount"`

		// 运行时长
		Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

		// 实例启动时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 容器数组
		Containers []*string `json:"Containers,omitempty" name:"Containers" list`

		// 容器名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 容器id(docker id)
		ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

		// 容器状态：
	// Running ：正常运行中
	// Waiting：等待运行中，例如正在下载镜像
	// Terminating：实例有容器正在终止
	// Terminated：实例有容器已经终止
	// NotReady：实例有容器处于未就绪状态，比如容器的健康检查失败
	// Error：其他未知异常
		Status *string `json:"Status,omitempty" name:"Status"`

		// 容器处于当前状态的原因，例如下载镜像失败
		Reason *string `json:"Reason,omitempty" name:"Reason"`

		// 容器的镜像
		Image *string `json:"Image,omitempty" name:"Image"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListGroupPodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupPodResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsByScalableRuleIdRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 偏移量，从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// limit量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// applicationId和groupId为参数查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *ListGroupsByScalableRuleIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsByScalableRuleIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsByScalableRuleIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *GroupsByScalableRuleIdList `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListGroupsByScalableRuleIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsByScalableRuleIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListMachinesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListMachinesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Result []*MachineList `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListMachinesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListMonitorStatisticsPolicyRequest struct {
	*tchttp.BaseRequest

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 关键字搜索
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// clusterId
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListMonitorStatisticsPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListMonitorStatisticsPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListMonitorStatisticsPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *MonitorStatisticsPolicyResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMonitorStatisticsPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListMonitorStatisticsPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPkgRequest struct {
	*tchttp.BaseRequest

	// 无
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 无
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 无
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 无
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListPkgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPkgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPkgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Result *PkgList `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPkgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPkgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListScalableRuleRequest struct {
	*tchttp.BaseRequest

	// 关键字搜索，按照name和id进行模糊匹配
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量,从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// limit量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListScalableRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListScalableRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListScalableRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果列表
		Result *ListScalableRuleResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListScalableRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListScalableRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListScalableRuleResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ScalableRule `json:"Content,omitempty" name:"Content" list`
}

type ListScalableTasksRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListScalableTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListScalableTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListScalableTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *ScalableTaskResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListScalableTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListScalableTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTsfModulesDetailRequest struct {
	*tchttp.BaseRequest

	// ModuleType
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// Status
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ListTsfModulesDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTsfModulesDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTsfModulesDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result []*ModulesResult `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTsfModulesDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTsfModulesDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTsfRegionResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*TsfRegionResult `json:"Content,omitempty" name:"Content" list`
}

type ListTsfZoneResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*TsfZoneResult `json:"Content,omitempty" name:"Content" list`
}

type MachineList struct {

	// 机器id
	MachineId *string `json:"MachineId,omitempty" name:"MachineId"`

	// ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// UserName
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// ModuleNames
	ModuleNames []*string `json:"ModuleNames,omitempty" name:"ModuleNames" list`
}

type MetricDataPoint struct {

	// 数据点键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 数据点值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 数据点标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type Microservice struct {

	// 微服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务的运行实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`
}

type MicroserviceExtra struct {

	// 微服务 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 关联的应用 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateApplicationId *string `json:"AssociateApplicationId,omitempty" name:"AssociateApplicationId"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceProtocol *string `json:"MicroserviceProtocol,omitempty" name:"MicroserviceProtocol"`

	// 监听端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceListeningPort *int64 `json:"MicroserviceListeningPort,omitempty" name:"MicroserviceListeningPort"`

	// 健康检查参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" name:"HealthCheckUrl"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceClustertype *string `json:"ServiceClustertype,omitempty" name:"ServiceClustertype"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 命名空间 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type ModifyAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// PolicyId
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// PolicyName
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// EventPolicies
	EventPolicies []*EventPolicyResult `json:"EventPolicies,omitempty" name:"EventPolicies" list`

	// EnabledEmail
	EnabledEmail *int64 `json:"EnabledEmail,omitempty" name:"EnabledEmail"`

	// EnabledSMS
	EnabledSMS *int64 `json:"EnabledSMS,omitempty" name:"EnabledSMS"`

	// Enabled
	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`

	// EnabledWeChat
	EnabledWeChat *int64 `json:"EnabledWeChat,omitempty" name:"EnabledWeChat"`

	// EnabledRtx
	EnabledRtx *int64 `json:"EnabledRtx,omitempty" name:"EnabledRtx"`
}

func (r *ModifyAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用备注
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`
}

func (r *ModifyApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthorizationInfoRequest struct {
	*tchttp.BaseRequest

	// 被调用服务 ID
	TargetServiceId *string `json:"TargetServiceId,omitempty" name:"TargetServiceId"`

	// 命名空间 ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 是否启用鉴权功能
	IsEnabled *bool `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 鉴权条件
	Conditions []*AuthCondition `json:"Conditions,omitempty" name:"Conditions" list`
}

func (r *ModifyAuthorizationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuthorizationInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthorizationInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAuthorizationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuthorizationInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthorizationRequest struct {
	*tchttp.BaseRequest

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称，为null时不更新
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 是否启用，为null时不更新
	IsEnabled *string `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 标签列表，为null时不更新
	Tags []*AuthConditionV2 `json:"Tags,omitempty" name:"Tags" list`
}

func (r *ModifyAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuthorizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：更新成功；false：更新失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuthorizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthorizationTypeRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 鉴权类型，D：未启用；B：黑名单模式；W：白名单模式
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyAuthorizationTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuthorizationTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthorizationTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：更新成功；false：更新失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAuthorizationTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAuthorizationTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
}

func (r *ModifyClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新集群详情操作是否成功。
	// true： 操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// ProtocolPorts数组
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`
}

func (r *ModifyContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新部署组是否成功。
	// true：成功。
	// false：失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContainerReplicasRequest struct {
	*tchttp.BaseRequest

	// 部署组ID，部署组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

func (r *ModifyContainerReplicasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContainerReplicasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContainerReplicasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContainerReplicasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContainerReplicasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组描述
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`
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

		// 更新部署组详情是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

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

type ModifyInstanceNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 机器id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyInstanceNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更换命名空间是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMachinesRequest struct {
	*tchttp.BaseRequest

	// 机器列表
	MachineList []*MachineList `json:"MachineList,omitempty" name:"MachineList" list`
}

func (r *ModifyMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMachinesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMachinesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 微服务 ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务备注信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

func (r *ModifyMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMicroserviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改微服务详情是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMicroserviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleConfRequest struct {
	*tchttp.BaseRequest

	// ModuleId
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// ExternalVip
	ExternalVip *string `json:"ExternalVip,omitempty" name:"ExternalVip"`

	// InternalVip
	InternalVip *string `json:"InternalVip,omitempty" name:"InternalVip"`

	// Vports
	Vports []*VportTypeResult `json:"Vports,omitempty" name:"Vports" list`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModifyModuleConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleConfRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleConfResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleConfResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMonitorStatisticsPolicyRequest struct {
	*tchttp.BaseRequest

	// 监控统计策略id;
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 关键词
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 部署组id列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`
}

func (r *ModifyMonitorStatisticsPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMonitorStatisticsPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMonitorStatisticsPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改结果true：修改成功；false：修改失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMonitorStatisticsPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMonitorStatisticsPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNamespaceCodeRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间编码
	NamespaceCode *string `json:"NamespaceCode,omitempty" name:"NamespaceCode"`
}

func (r *ModifyNamespaceCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNamespaceCodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNamespaceCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：更新成功；false：更新失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNamespaceCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNamespaceCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间备注
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`
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

type ModifyProgramRequest struct {
	*tchttp.BaseRequest

	// 数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// 数据集名称，不传入时不更新
	ProgramName *string `json:"ProgramName,omitempty" name:"ProgramName"`

	// 数据集描述，不传入时不更新
	ProgramDesc *string `json:"ProgramDesc,omitempty" name:"ProgramDesc"`

	// 数据项列表，传入null不更新，传入空数组全量删除
	ProgramItemList []*ProgramItem `json:"ProgramItemList,omitempty" name:"ProgramItemList" list`
}

func (r *ModifyProgramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProgramRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProgramResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true: 更新成功；false: 更新失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProgramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProgramResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRegionRequest struct {
	*tchttp.BaseRequest

	// TRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`

	// TRegionName
	TRegionName *string `json:"TRegionName,omitempty" name:"TRegionName"`

	// TRemark
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`
}

func (r *ModifyRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoleRequest struct {
	*tchttp.BaseRequest

	// 角色ID
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称，不传入或为null时不更新
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色描述，不传入或为null时不更新
	RoleDesc *string `json:"RoleDesc,omitempty" name:"RoleDesc"`

	// 角色拥有的权限组ID列表，不传入或为null时不更新
	PermCatIdList []*string `json:"PermCatIdList,omitempty" name:"PermCatIdList" list`
}

func (r *ModifyRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true: 更新成功；false: 更新失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteRequest struct {
	*tchttp.BaseRequest

	// 路由ID
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// 路由名称
	RouteName *string `json:"RouteName,omitempty" name:"RouteName"`

	// 路由描述
	RouteDesc *string `json:"RouteDesc,omitempty" name:"RouteDesc"`

	// 路由规则列表
	RuleList []*RouteRuleV2 `json:"RuleList,omitempty" name:"RuleList" list`
}

func (r *ModifyRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新路由详情操作是否成功。
	// true：更新路由操作成功。
	// false：更新路由操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteRuleRequest struct {
	*tchttp.BaseRequest

	// 路由规则Id
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`

	// 路由规则名称
	RouteRuleName *string `json:"RouteRuleName,omitempty" name:"RouteRuleName"`

	// 路由规则描述
	RouteRuleDesc *string `json:"RouteRuleDesc,omitempty" name:"RouteRuleDesc"`

	// 标签路由规则项
	TagRouteItemList []*TagRouteItemList `json:"TagRouteItemList,omitempty" name:"TagRouteItemList" list`

	// 权重路由规则项
	WeightRouteItemList []*WeightRouteItemList `json:"WeightRouteItemList,omitempty" name:"WeightRouteItemList" list`
}

func (r *ModifyRouteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true / false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRouteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScalableRuleRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 弹性伸缩规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 冷却时间, s为单位, 0-9999
	CoolTime *uint64 `json:"CoolTime,omitempty" name:"CoolTime"`

	// 包含缩容规则, 默认0, 0:否 1:是
	EnableShrink *uint64 `json:"EnableShrink,omitempty" name:"EnableShrink"`

	// 包含扩容规则，默认0, 0否 1是
	EnableExpand *uint64 `json:"EnableExpand,omitempty" name:"EnableExpand"`

	// 扩容规则持续时间,0-9999s
	ExpandPeriod *uint64 `json:"ExpandPeriod,omitempty" name:"ExpandPeriod"`

	// 缩容规则持续时间，0-9999s
	ShrinkPeriod *uint64 `json:"ShrinkPeriod,omitempty" name:"ShrinkPeriod"`

	// 单次扩容机器数量, 0-9999
	ExpandVmCount *uint64 `json:"ExpandVmCount,omitempty" name:"ExpandVmCount"`

	// 单次缩容机器数量, 0-9999
	ShrinkVmCount *uint64 `json:"ShrinkVmCount,omitempty" name:"ShrinkVmCount"`

	// 扩容之后，最大实例数目,0-9999
	ExpandVmCountLimit *uint64 `json:"ExpandVmCountLimit,omitempty" name:"ExpandVmCountLimit"`

	// 缩容之后，最小实例数目, 0-9999
	ShrinkVmCountLimit *uint64 `json:"ShrinkVmCountLimit,omitempty" name:"ShrinkVmCountLimit"`

	// 扩容规则
	ExpandSubruleList []*ScalableSubRule `json:"ExpandSubruleList,omitempty" name:"ExpandSubruleList" list`

	// 缩容规则
	ShrinkSubruleList []*ScalableSubRule `json:"ShrinkSubruleList,omitempty" name:"ShrinkSubruleList" list`
}

func (r *ModifyScalableRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScalableRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScalableRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求结果,true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyScalableRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScalableRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTsfZoneRequest struct {
	*tchttp.BaseRequest

	// TZoneId
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`

	// TZoneName
	TZoneName *string `json:"TZoneName,omitempty" name:"TZoneName"`

	// TRemark
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`

	// TRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`
}

func (r *ModifyTsfZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTsfZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTsfZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTsfZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTsfZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUploadInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 调用DescribeUploadInfo接口时返回的软件包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// COS返回上传结果（默认为0：成功，其他值表示失败）
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 程序包MD5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 程序包大小（单位字节）
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

func (r *ModifyUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUploadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUploadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest

	// TZoneId
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`

	// TZoneName
	TZoneName *string `json:"TZoneName,omitempty" name:"TZoneName"`

	// TRemark
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`

	// TRegionId
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`
}

func (r *ModifyZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModuleParamResult struct {

	// ParameterName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`

	// ParameterValue
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterValue *string `json:"ParameterValue,omitempty" name:"ParameterValue"`

	// ParameterType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterType *string `json:"ParameterType,omitempty" name:"ParameterType"`

	// ParameterEnums
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterEnums *string `json:"ParameterEnums,omitempty" name:"ParameterEnums"`

	// Desc
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Demo
	// 注意：此字段可能返回 null，表示取不到有效值。
	Demo *string `json:"Demo,omitempty" name:"Demo"`

	// ParameterId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterId *string `json:"ParameterId,omitempty" name:"ParameterId"`

	// Visible
	// 注意：此字段可能返回 null，表示取不到有效值。
	Visible *string `json:"Visible,omitempty" name:"Visible"`

	// ParameterConfigType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterConfigType *string `json:"ParameterConfigType,omitempty" name:"ParameterConfigType"`

	// ParameterSubType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterSubType *string `json:"ParameterSubType,omitempty" name:"ParameterSubType"`
}

type ModuleParameterResult struct {

	// ParameterConfigType
	ParameterConfigType *string `json:"ParameterConfigType,omitempty" name:"ParameterConfigType"`

	// ParameterSubType
	ParameterSubType *string `json:"ParameterSubType,omitempty" name:"ParameterSubType"`

	// ParameterName
	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`

	// ParameterValue
	ParameterValue *string `json:"ParameterValue,omitempty" name:"ParameterValue"`
}

type ModulesResult struct {

	// ModuleId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// ModuleName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// ModuleType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// ExternalVip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalVip *string `json:"ExternalVip,omitempty" name:"ExternalVip"`

	// InternalVip
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternalVip *string `json:"InternalVip,omitempty" name:"InternalVip"`

	// Vports
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vports []*VportTypeResult `json:"Vports,omitempty" name:"Vports" list`

	// Instances
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*InstancesResult `json:"Instances,omitempty" name:"Instances" list`

	// AlarmType
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`

	// ModuleCommon
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleCommon *int64 `json:"ModuleCommon,omitempty" name:"ModuleCommon"`

	// DeployOrder
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployOrder *int64 `json:"DeployOrder,omitempty" name:"DeployOrder"`

	// Visible
	// 注意：此字段可能返回 null，表示取不到有效值。
	Visible *string `json:"Visible,omitempty" name:"Visible"`

	// ViportVisible
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViportVisible *bool `json:"ViportVisible,omitempty" name:"ViportVisible"`

	// HasRole
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasRole *bool `json:"HasRole,omitempty" name:"HasRole"`
}

type MonitorGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 部署组描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`
}

type MonitorOverview struct {

	// 近24小时调用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationCountOfDay *string `json:"InvocationCountOfDay,omitempty" name:"InvocationCountOfDay"`

	// 总调用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationCount *string `json:"InvocationCount,omitempty" name:"InvocationCount"`

	// 近24小时调用错误数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCountOfDay *string `json:"ErrorCountOfDay,omitempty" name:"ErrorCountOfDay"`

	// 总调用错误数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCount *string `json:"ErrorCount,omitempty" name:"ErrorCount"`

	// 近24小时调用成功率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessRatioOfDay *string `json:"SuccessRatioOfDay,omitempty" name:"SuccessRatioOfDay"`

	// 总调用成功率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessRatio *string `json:"SuccessRatio,omitempty" name:"SuccessRatio"`
}

type MonitorStatisticsPolicy struct {

	// 监控统计策略id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 部署组列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupList []*MonitorStatisticsPolicyGroup `json:"GroupList,omitempty" name:"GroupList" list`

	// 命名空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type MonitorStatisticsPolicyGroup struct {

	// GroupId
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// GroupName
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// GroupDesc
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// NamespaceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// NamespaceName
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// ApplicationId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// ApplicationName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// ClusterId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ClusterName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type MonitorStatisticsPolicyResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*MonitorStatisticsPolicy `json:"Content,omitempty" name:"Content" list`
}

type Ms struct {

	// MsName
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsName *string `json:"MsName,omitempty" name:"MsName"`

	// MsPort
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsPort *string `json:"MsPort,omitempty" name:"MsPort"`

	// Controller
	// 注意：此字段可能返回 null，表示取不到有效值。
	Controller []*Controller `json:"Controller,omitempty" name:"Controller" list`

	// Provider
	// 注意：此字段可能返回 null，表示取不到有效值。
	Provider []*Provider `json:"Provider,omitempty" name:"Provider" list`

	// Consumer
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consumer []*Consumer `json:"Consumer,omitempty" name:"Consumer" list`
}

type MsApiArray struct {

	// API 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 方法描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type MsInstance struct {

	// 机器实例ID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 机器实例名称信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 服务运行的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 机器实例内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 机器实例外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 机器可用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 服务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 机器TSF可用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 健康检查URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" name:"HealthCheckUrl"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 应用程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationPackageVersion *string `json:"ApplicationPackageVersion,omitempty" name:"ApplicationPackageVersion"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`
}

type MsRunningApplicationV2 struct {

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
}

type Namespace struct {

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceCode *string `json:"NamespaceCode,omitempty" name:"NamespaceCode"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 默认命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`

	// 命名空间状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" name:"NamespaceStatus"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群数组，仅携带集群ID，集群名称，集群类型等基础信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterList []*Cluster `json:"ClusterList,omitempty" name:"ClusterList" list`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type OverviewMsResult struct {

	// 概览页微服务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunMicroserviceCount *int64 `json:"RunMicroserviceCount,omitempty" name:"RunMicroserviceCount"`
}

type OverviewResource struct {

	// 集群数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterCount *int64 `json:"ClusterCount,omitempty" name:"ClusterCount"`

	// 命名空间数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceCount *int64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`

	// 实例总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 运行中实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 停止实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffInstanceCount *int64 `json:"OffInstanceCount,omitempty" name:"OffInstanceCount"`

	// 异常实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnNormalInstanceCount *int64 `json:"UnNormalInstanceCount,omitempty" name:"UnNormalInstanceCount"`

	// 应用数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationCount *int64 `json:"ApplicationCount,omitempty" name:"ApplicationCount"`

	// 部署组数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCount *int64 `json:"GroupCount,omitempty" name:"GroupCount"`

	// 异常集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnNormalClusterCount *int64 `json:"UnNormalClusterCount,omitempty" name:"UnNormalClusterCount"`

	// 运行云主机数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunHostCount *int64 `json:"RunHostCount,omitempty" name:"RunHostCount"`

	// 云主机总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
}

type PagedPermCat struct {

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 产品列表
	Content []*PermCat `json:"Content,omitempty" name:"Content" list`
}

type PagedProgram struct {

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据集列表
	Content []*Program `json:"Content,omitempty" name:"Content" list`
}

type PagedResource struct {

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资源列表
	Content []*Resource `json:"Content,omitempty" name:"Content" list`
}

type PagedRole struct {

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 角色列表
	Content []*Role `json:"Content,omitempty" name:"Content" list`
}

type PagedService struct {

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 产品列表
	Content []*Service `json:"Content,omitempty" name:"Content" list`
}

type PermCat struct {

	// 权限组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermCatId *string `json:"PermCatId,omitempty" name:"PermCatId"`

	// 权限组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermCatName *string `json:"PermCatName,omitempty" name:"PermCatName"`

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCode *string `json:"ServiceCode,omitempty" name:"ServiceCode"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 权限组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermCatDesc *string `json:"PermCatDesc,omitempty" name:"PermCatDesc"`
}

type Person struct {

	// 接收成员电话
	CellPhoneNumber *string `json:"CellPhoneNumber,omitempty" name:"CellPhoneNumber"`

	// 接收成员email
	Email *string `json:"Email,omitempty" name:"Email"`

	// 接收成员的名字
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PkgInfo struct {

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 程序包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 程序包类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 程序包描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`

	// 上传时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadTime *string `json:"UploadTime,omitempty" name:"UploadTime"`

	// 程序包MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 程序包状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgPubStatus *int64 `json:"PkgPubStatus,omitempty" name:"PkgPubStatus"`
}

type PkgList struct {

	// 程序包总量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 程序包信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*PkgInfo `json:"Content,omitempty" name:"Content" list`
}

type PolicyItem struct {

	// 角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`
}

type PortsResult struct {

	// PortId
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortId *string `json:"PortId,omitempty" name:"PortId"`

	// PortType
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortType *string `json:"PortType,omitempty" name:"PortType"`

	// Port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// DefaultPort
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultPort *string `json:"DefaultPort,omitempty" name:"DefaultPort"`

	// ModuleId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type Program struct {

	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// 数据集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramName *string `json:"ProgramName,omitempty" name:"ProgramName"`

	// 数据集描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDesc *string `json:"ProgramDesc,omitempty" name:"ProgramDesc"`

	// 删除标识，true: 可以删除; false: 不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 数据项列表，无值时返回空数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramItemList []*ProgramItem `json:"ProgramItemList,omitempty" name:"ProgramItemList" list`
}

type ProgramItem struct {

	// 数据项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramItemId *string `json:"ProgramItemId,omitempty" name:"ProgramItemId"`

	// 资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`

	// 数据值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueList []*string `json:"ValueList,omitempty" name:"ValueList" list`

	// 全选标识，true: 全选；false: 非全选
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 删除标识，true: 可删除；false: 不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`
}

type PropertyField struct {

	// 属性名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

type ProtocolPort struct {

	// TCP UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 服务端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 容器端口
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`
}

type Provider struct {

	// ProviderControllerName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProviderControllerName *string `json:"ProviderControllerName,omitempty" name:"ProviderControllerName"`
}

type QuantileEntity struct {

	// 最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxValue *string `json:"MaxValue,omitempty" name:"MaxValue"`

	// 最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinValue *string `json:"MinValue,omitempty" name:"MinValue"`

	// 五分位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FifthPositionValue *string `json:"FifthPositionValue,omitempty" name:"FifthPositionValue"`

	// 九分位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NinthPositionValue *string `json:"NinthPositionValue,omitempty" name:"NinthPositionValue"`
}

type RatelimitDimension struct {

	// `S`:  系统标签/`U`: 自定义标签
	TagType *string `json:"TagType,omitempty" name:"TagType"`

	// 标签名
	TagField *string `json:"TagField,omitempty" name:"TagField"`

	// EQUAL/NOT_EQUAL/IN/NOT_IN/REGEX
	TagOperator *string `json:"TagOperator,omitempty" name:"TagOperator"`

	// 标签匹配值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type RatelimitListResult struct {

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*RatelimitRuleV2 `json:"Content,omitempty" name:"Content" list`
}

type RatelimitRuleForUpdate struct {

	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则名字，在一个微服务下唯一
	Name *string `json:"Name,omitempty" name:"Name"`

	// 状态 0表示启用 1表示停用
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 限流周期，单位秒
	DurationSecond *uint64 `json:"DurationSecond,omitempty" name:"DurationSecond"`

	// 每周期内的限流配额
	DurationQuota *uint64 `json:"DurationQuota,omitempty" name:"DurationQuota"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type RatelimitRuleForUpdateV2 struct {

	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则名字，在一个微服务下唯一
	Name *string `json:"Name,omitempty" name:"Name"`

	// 状态 0表示启用 1表示停用
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 限流周期，单位秒
	DurationSecond *uint64 `json:"DurationSecond,omitempty" name:"DurationSecond"`

	// 每周期内的限流配额
	DurationQuota *uint64 `json:"DurationQuota,omitempty" name:"DurationQuota"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 标签规则
	Dimensions []*RatelimitDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type RatelimitRuleV2 struct {

	// 规则名字，在一个微服务下唯一
	Name *string `json:"Name,omitempty" name:"Name"`

	// 限流周期，单位秒
	DurationSecond *uint64 `json:"DurationSecond,omitempty" name:"DurationSecond"`

	// 每周期内的限流配额
	DurationQuota *uint64 `json:"DurationQuota,omitempty" name:"DurationQuota"`

	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 状态 0表示启用 1表示停用
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 限流规则区分的来源微服务名
	SourceService *string `json:"SourceService,omitempty" name:"SourceService"`

	// 最近一次修改规则时间，UTC秒数
	ModifyTime *uint64 `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 限制条件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions []*RatelimitDimension `json:"Dimensions,omitempty" name:"Dimensions" list`
}

type ReRelateGroupToScalableRuleRequest struct {
	*tchttp.BaseRequest

	// 新规则id;
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 1:启用 2:关闭
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ReRelateGroupToScalableRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReRelateGroupToScalableRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReRelateGroupToScalableRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求结果,true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReRelateGroupToScalableRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReRelateGroupToScalableRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RealtimeLogSet struct {

	// 请求基准时间戳（用于下次调用）
	RealtimeTs *uint64 `json:"RealtimeTs,omitempty" name:"RealtimeTs"`

	// 业务日志集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessLogSet *TsfPageBusinessLogV2 `json:"BusinessLogSet,omitempty" name:"BusinessLogSet"`

	// 标准输出日志集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	StdoutLogSet *TsfPageStdoutLogV2 `json:"StdoutLogSet,omitempty" name:"StdoutLogSet"`
}

type ReassociateBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// 原关联日志配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 新关联日志配置ID
	NewConfigId *string `json:"NewConfigId,omitempty" name:"NewConfigId"`

	// TSF应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// TSF部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ReassociateBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReassociateBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReassociateBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReassociateBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReassociateBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecordV2 struct {

	// 账号AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 账号Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 账号SubAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// 操作记录Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 模块类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 操作资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationResource *string `json:"OperationResource,omitempty" name:"OperationResource"`

	// 操作记录描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationMsg *string `json:"OperationMsg,omitempty" name:"OperationMsg"`

	// 操作记录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`

	// 操作状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationStatus *string `json:"OperationStatus,omitempty" name:"OperationStatus"`

	// 操作人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type RegisterImageUserRequest struct {
	*tchttp.BaseRequest

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 确认密码
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`
}

func (r *RegisterImageUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterImageUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterImageUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterImageUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterImageUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RelateGroupToScalableRuleRequest struct {
	*tchttp.BaseRequest

	// 规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 1:启用 2:关闭
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *RelateGroupToScalableRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RelateGroupToScalableRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RelateGroupToScalableRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求结果,true/false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelateGroupToScalableRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RelateGroupToScalableRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseApiGroupRequest struct {
	*tchttp.BaseRequest

	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ReleaseApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseApiGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功/失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseApiGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// 日志配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 分组ID列表
	GroupIdList []*string `json:"GroupIdList,omitempty" name:"GroupIdList" list`
}

func (r *ReleaseBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发布结果，true：成功；false：失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseConfigRequest struct {
	*tchttp.BaseRequest

	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleaseConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：发布成功；false：发布失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseFileConfigRequest struct {
	*tchttp.BaseRequest

	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleaseFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseFileConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发布结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseFileConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleasePluginRequest struct {
	*tchttp.BaseRequest

	// 插件ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *ReleasePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleasePluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleasePluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleasePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleasePluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleasePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleasePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleasePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleasePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：发布成功；false：发布失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleasePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleasePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstanceRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 机器ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RemoveInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群移除机器是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机 ID 列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *RemoveInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群移除机器是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Resource struct {

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceCode *string `json:"ResourceCode,omitempty" name:"ResourceCode"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源所属产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCode *string `json:"ServiceCode,omitempty" name:"ServiceCode"`

	// 选取资源使用的Action
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceAction *string `json:"ResourceAction,omitempty" name:"ResourceAction"`

	// 资源数据查询的ID字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdField *string `json:"IdField,omitempty" name:"IdField"`

	// 资源数据查询的名称字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameField *string `json:"NameField,omitempty" name:"NameField"`

	// 资源数据查询的ID过滤字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelectIdsField *string `json:"SelectIdsField,omitempty" name:"SelectIdsField"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 资源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceDesc *string `json:"ResourceDesc,omitempty" name:"ResourceDesc"`

	// 是否可以选择全部
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanSelectAll *bool `json:"CanSelectAll,omitempty" name:"CanSelectAll"`

	// 资源数据查询的模糊查询字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchWordField *string `json:"SearchWordField,omitempty" name:"SearchWordField"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type ResourceNodeStatus struct {

	// 已使用的资源量
	Used *uint64 `json:"Used,omitempty" name:"Used"`

	// 总的资源量
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type ResourceStatus struct {

	// 该模块多个节点的总的资源使用量
	Used *uint64 `json:"Used,omitempty" name:"Used"`

	// 该模块的多个节点的总的资源量
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 各个节点的资源的情况
	NodesStatus []*ResourceNodeStatus `json:"NodesStatus,omitempty" name:"NodesStatus" list`
}

type Result struct {

	// KeyWordsName 和 GroupName
	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`
}

type RetryTransactionRequest struct {
	*tchttp.BaseRequest

	// 重试的事务ID列表
	TransactionId []*string `json:"TransactionId,omitempty" name:"TransactionId" list`
}

func (r *RetryTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RetryTransactionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RetryTransactionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 重试事务接口结果
		Result *TxRetry `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RetryTransactionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevocationConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

func (r *RevocationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevocationConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevocationConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：回滚成功；false：回滚失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevocationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevocationConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevocationPublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

func (r *RevocationPublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevocationPublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevocationPublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：撤销成功；false：撤销失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevocationPublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevocationPublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevokeFileConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

func (r *RevokeFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevokeFileConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevokeFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 撤回结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevokeFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevokeFileConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Role struct {

	// 角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleDesc *string `json:"RoleDesc,omitempty" name:"RoleDesc"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 角色权限集列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermCatList []*PermCat `json:"PermCatList,omitempty" name:"PermCatList" list`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type RollbackConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布历史ID
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 回滚描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *RollbackConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：回滚成功；false：回滚失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackFileConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布历史ID
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 回滚描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *RollbackFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackFileConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回滚结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackFileConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackPublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布历史ID
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 回滚描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *RollbackPublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackPublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackPublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：回滚成功；false：回滚失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackPublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackPublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RouteAffinity struct {

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 就近访问策略是否开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Affinity *bool `json:"Affinity,omitempty" name:"Affinity"`

	// 就近访问策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffinityId *string `json:"AffinityId,omitempty" name:"AffinityId"`

	// 就近访问策略更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTimestamp *string `json:"UpdateTimestamp,omitempty" name:"UpdateTimestamp"`
}

type RouteDestItemV2 struct {

	// 路由规则目标字段名称
	DestItemField *string `json:"DestItemField,omitempty" name:"DestItemField"`

	// 路由规则目标字段取值
	DestItemValue *string `json:"DestItemValue,omitempty" name:"DestItemValue"`

	// 路由规则路由目标匹配项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteDestItemId *string `json:"RouteDestItemId,omitempty" name:"RouteDestItemId"`

	// 所属路由规则路由目标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteDestId *string `json:"RouteDestId,omitempty" name:"RouteDestId"`
}

type RouteDestV2 struct {

	// 路由目标权重
	DestWeight *int64 `json:"DestWeight,omitempty" name:"DestWeight"`

	// 路由目标匹配条件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestItemList []*RouteDestItemV2 `json:"DestItemList,omitempty" name:"DestItemList" list`

	// 路由规则项路由目标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestId *string `json:"DestId,omitempty" name:"DestId"`

	// 服务路由目标所属路由规则项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`
}

type RouteReleaseHistory struct {

	// 路由规则发布记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteReleaseLogId *string `json:"RouteReleaseLogId,omitempty" name:"RouteReleaseLogId"`

	// 微服务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 路由规则Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`

	// 路由规则部署开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableTime *string `json:"EnableTime,omitempty" name:"EnableTime"`

	// 路由规则停止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisableTime *string `json:"DisableTime,omitempty" name:"DisableTime"`

	// 路由规则发布
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 路由规则详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRule *RouteRule `json:"RouteRule,omitempty" name:"RouteRule"`

	// 账号APPID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 账号Owner用户唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 账号用户唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type RouteRule struct {

	// 路由规则Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`

	// 路由规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRuleName *string `json:"RouteRuleName,omitempty" name:"RouteRuleName"`

	// 路由规则类型，包括标签路由和权重标签，标签路由： T , 权重路由： W
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRuleType *string `json:"RouteRuleType,omitempty" name:"RouteRuleType"`

	// 路由规则描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRuleDesc *string `json:"RouteRuleDesc,omitempty" name:"RouteRuleDesc"`

	// 微服务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// true/false，true：路由规则为启用状态， false：路由规则为停用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableStatus *bool `json:"EnableStatus,omitempty" name:"EnableStatus"`

	// 路由规则创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 路由规则更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 标签路由规则项信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagRouteItemList []*TagRouteItemList `json:"TagRouteItemList,omitempty" name:"TagRouteItemList" list`

	// 权重路由规则项信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeightRouteItemList []*WeightRouteItemList `json:"WeightRouteItemList,omitempty" name:"WeightRouteItemList" list`

	// 微服务所属命名空间Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`
}

type RouteRuleV2 struct {

	// 路由规则项包含目标列表
	DestList []*RouteDestV2 `json:"DestList,omitempty" name:"DestList" list`

	// 路由规则项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`

	// 路由规则项包含TAG列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*RouteTagV2 `json:"TagList,omitempty" name:"TagList" list`

	// 路由规则项所属路由规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
}

type RouteTagV2 struct {

	// 标签类型，S：系统标签， U：自定义标签
	TagType *string `json:"TagType,omitempty" name:"TagType"`

	// 标签字段名称
	TagField *string `json:"TagField,omitempty" name:"TagField"`

	// 标签匹配规则.
	// EQUAL：等于、NOT_EQUAL：不等于、IN：包含、NOT_IN：不包含、REGEX：正则
	TagOperator *string `json:"TagOperator,omitempty" name:"TagOperator"`

	// 标签取值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 路由规则项TAG匹配项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// TAG匹配项所述路由规则项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`
}

type RouteV2 struct {

	// 路由规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// 路由规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteName *string `json:"RouteName,omitempty" name:"RouteName"`

	// 路由规则描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteDesc *string `json:"RouteDesc,omitempty" name:"RouteDesc"`

	// 微服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 路由启用状态。
	// true：路由规则为启用状态，。
	// false：路由规则为停用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableStatus *bool `json:"EnableStatus,omitempty" name:"EnableStatus"`

	// 路由规则创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 路由规则更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 微服务所属命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 路由规则包含路由规则项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleList []*RouteRuleV2 `json:"RuleList,omitempty" name:"RuleList" list`
}

type RunMsApiRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// API 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// API 方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// API 的 ContentType
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// API 的 header 参数
	RequestHeader *string `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// API 的 query 参数
	RequestQuery *string `json:"RequestQuery,omitempty" name:"RequestQuery"`

	// API 的 path 参数
	RequestPath *string `json:"RequestPath,omitempty" name:"RequestPath"`

	// API 的 body 参数（前端拼装为 json array 形式）
	RequestBodyDict *string `json:"RequestBodyDict,omitempty" name:"RequestBodyDict"`

	// API 的 body 参数（用户直接填写）
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
}

func (r *RunMsApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunMsApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMsApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调试 API 的 http code
		ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 调试 API 的延迟
		Delay *string `json:"Delay,omitempty" name:"Delay"`

		// 调试 API 的 http header
		ReturnHeader *string `json:"ReturnHeader,omitempty" name:"ReturnHeader"`

		// 调试 API 的 http body
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnBody *string `json:"ReturnBody,omitempty" name:"ReturnBody"`
	} `json:"Response"`
}

func (r *RunMsApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunMsApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SaveDeployModuleParamsRequest struct {
	*tchttp.BaseRequest

	// ModuleParameters
	ModuleParameters []*ModuleParameterResult `json:"ModuleParameters,omitempty" name:"ModuleParameters" list`
}

func (r *SaveDeployModuleParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SaveDeployModuleParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SaveDeployModuleParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveDeployModuleParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SaveDeployModuleParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScalableRule struct {

	// RuleId
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// ExpandVmCountLimit
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpandVmCountLimit *int64 `json:"ExpandVmCountLimit,omitempty" name:"ExpandVmCountLimit"`

	// ShrinkVmCountLimit
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShrinkVmCountLimit *int64 `json:"ShrinkVmCountLimit,omitempty" name:"ShrinkVmCountLimit"`

	// GroupCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCount *int64 `json:"GroupCount,omitempty" name:"GroupCount"`
}

type ScalableRuleAttribute struct {

	// RuleId
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// EnableShrink
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableShrink *int64 `json:"EnableShrink,omitempty" name:"EnableShrink"`

	// EnableExpand
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableExpand *int64 `json:"EnableExpand,omitempty" name:"EnableExpand"`

	// ExpandPeriod
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpandPeriod *int64 `json:"ExpandPeriod,omitempty" name:"ExpandPeriod"`

	// ShrinkPeriod
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShrinkPeriod *int64 `json:"ShrinkPeriod,omitempty" name:"ShrinkPeriod"`

	// ExpandVmCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpandVmCount *int64 `json:"ExpandVmCount,omitempty" name:"ExpandVmCount"`

	// ShrinkVmCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShrinkVmCount *int64 `json:"ShrinkVmCount,omitempty" name:"ShrinkVmCount"`

	// CoolTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoolTime *int64 `json:"CoolTime,omitempty" name:"CoolTime"`

	// ExpandVmCountLimit
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpandVmCountLimit *int64 `json:"ExpandVmCountLimit,omitempty" name:"ExpandVmCountLimit"`

	// ShrinkVmCountLimit
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShrinkVmCountLimit *int64 `json:"ShrinkVmCountLimit,omitempty" name:"ShrinkVmCountLimit"`

	// ExpandSubruleList
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpandSubruleList []*ExpandSubrule `json:"ExpandSubruleList,omitempty" name:"ExpandSubruleList" list`

	// ShrinkSubruleList
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShrinkSubruleList []*ExpandSubrule `json:"ShrinkSubruleList,omitempty" name:"ShrinkSubruleList" list`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// UpdateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ScalableSubRule struct {

	// 监控指标的阈值。如指标为 CPU 利用率或内存利用率，则为百分比 0-100；如指标为响应时间或请求个数，则为 0-99999999
	Indicators *uint64 `json:"Indicators,omitempty" name:"Indicators"`

	// 规则指标类型: 1: CPU 利用率, 2: 内存利用率, 3: 请求响应时间, 4: 请求个数
	IndicatorType *uint64 `json:"IndicatorType,omitempty" name:"IndicatorType"`
}

type ScalableTask struct {

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Taskid *string `json:"Taskid,omitempty" name:"Taskid"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// EndTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// TaskInstanceList
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInstanceList []*TaskInstance `json:"TaskInstanceList,omitempty" name:"TaskInstanceList" list`
}

type ScalableTaskResult struct {

	// TotalCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ScalableTask `json:"Content,omitempty" name:"Content" list`
}

type SearchBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 日志配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 机器实例ID，不传表示全部实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 请求偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量，取值范围[1, 200]，默认值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序规则，默认值"time"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords" list`

	// 部署组ID列表，不传表示全部部署组
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// 检索类型，取值"LUCENE", "REGEXP", "NORMAL"
	SearchWordType *string `json:"SearchWordType,omitempty" name:"SearchWordType"`

	// 批量请求类型，取值"page"或"scroll"
	BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

	// 游标ID
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`
}

func (r *SearchBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务日志列表
		Result *TsfPageBusinessLogV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 实时日志基准时间戳
	RealtimeTs *uint64 `json:"RealtimeTs,omitempty" name:"RealtimeTs"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords" list`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 单页请求配置数量，取值范围[1, 200]，默认值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序规则，默认值"time"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 请求偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchOssBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务日志列表
		Result *TsfPageBusinessLogV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchOssBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssRealtimeBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 模块ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实时日志基准时间戳
	RealtimeTs *uint64 `json:"RealtimeTs,omitempty" name:"RealtimeTs"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords" list`
}

func (r *SearchOssRealtimeBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssRealtimeBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssRealtimeBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实时日志集合
		Result *RealtimeLogSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchOssRealtimeBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssRealtimeBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssSpanBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 调用链ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 调用链SpanID
	SpanId *string `json:"SpanId,omitempty" name:"SpanId"`

	// 这个spanId属于那个服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 请求偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchOssSpanBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssSpanBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssSpanBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务日志列表
		Result *TsfPageBusinessLogV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchOssSpanBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssSpanBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssStaticBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实时日志基准时间戳
	RealtimeTs *uint64 `json:"RealtimeTs,omitempty" name:"RealtimeTs"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords" list`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 单页请求配置数量，取值范围[1, 200]，默认值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序规则，默认值"time"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 请求偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchOssStaticBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssStaticBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssStaticBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务日志列表
		Result *TsfPageBusinessLogV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchOssStaticBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssStaticBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssSurroundBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 基准日志ID
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// 上下文日志基准时间戳
	SurroundTs *uint64 `json:"SurroundTs,omitempty" name:"SurroundTs"`

	// 上下文日志数量，取值范围[10, 100]，默认值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 上下文查询方向，B：双向，U：前向，D：向后
	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

func (r *SearchOssSurroundBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssSurroundBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssSurroundBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志列表
		Result *TsfPageBusinessLogV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchOssSurroundBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssSurroundBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssTraceRequest struct {
	*tchttp.BaseRequest

	// 搜索开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 搜索结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否仅展示错误调用链
	ErrorTrace *bool `json:"ErrorTrace,omitempty" name:"ErrorTrace"`

	// 错误码
	ResultStatus *string `json:"ResultStatus,omitempty" name:"ResultStatus"`

	// 错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 最小耗时
	MinDuration *uint64 `json:"MinDuration,omitempty" name:"MinDuration"`

	// 最长耗时
	MaxDuration *uint64 `json:"MaxDuration,omitempty" name:"MaxDuration"`

	// 排序方式，支持"timestamp"和"duration"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方向，支持"asc"和"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 返回trace数目，取值范围[0, 100]，默认值20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询trace偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 用户自定义标签
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 调用方服务名
	CallerServiceName *string `json:"CallerServiceName,omitempty" name:"CallerServiceName"`

	// 被调方服务名
	CalleeServiceName *string `json:"CalleeServiceName,omitempty" name:"CalleeServiceName"`

	// 调用方接口名
	CallerInterfaceName *string `json:"CallerInterfaceName,omitempty" name:"CallerInterfaceName"`

	// 被调方接口名
	CalleeInterfaceName *string `json:"CalleeInterfaceName,omitempty" name:"CalleeInterfaceName"`

	// 调用方IP
	CallerIp *string `json:"CallerIp,omitempty" name:"CallerIp"`

	// 被调方IP
	CalleeIp *string `json:"CalleeIp,omitempty" name:"CalleeIp"`
}

func (r *SearchOssTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssTraceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchOssTraceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调用链列表
		Result *TsfPageZipkinTraceInfoV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchOssTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchOssTraceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchRealtimeBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 日志配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组ID列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实时日志基准时间戳
	RealtimeTs *uint64 `json:"RealtimeTs,omitempty" name:"RealtimeTs"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords" list`
}

func (r *SearchRealtimeBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchRealtimeBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchRealtimeBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实时日志集合
		Result *RealtimeLogSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRealtimeBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchRealtimeBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchRealtimeStdoutLogRequest struct {
	*tchttp.BaseRequest

	// 机器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实时日志基准时间戳
	RealtimeTs *uint64 `json:"RealtimeTs,omitempty" name:"RealtimeTs"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords" list`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *SearchRealtimeStdoutLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchRealtimeStdoutLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchRealtimeStdoutLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实时日志集合
		Result *RealtimeLogSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRealtimeStdoutLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchRealtimeStdoutLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchSpanBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 调用链ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 调用链SpanID
	SpanId *string `json:"SpanId,omitempty" name:"SpanId"`

	// 请求偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchSpanBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchSpanBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchSpanBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务日志列表
		Result *TsfPageBusinessLogV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSpanBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchSpanBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStdoutLogRequest struct {
	*tchttp.BaseRequest

	// 机器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 单页请求配置数量，取值范围[1, 500]，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords" list`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 请求偏移量，取值范围大于等于0，默认值为
	// 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序规则，默认值"time"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，取值"asc"或"desc"，默认
	// 值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 检索类型，取值"LUCENE", "REGEXP",
	// "NORMAL"
	SearchWordType *string `json:"SearchWordType,omitempty" name:"SearchWordType"`

	// 批量请求类型，取值"page"或"scroll"，默认
	// 值"page"
	BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

	// 游标ID
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`
}

func (r *SearchStdoutLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStdoutLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStdoutLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 标准输出日志列表
		Result *TsfPageStdoutLogV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStdoutLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStdoutLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchSurroundBusinessLogRequest struct {
	*tchttp.BaseRequest

	// 日志配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 基准日志ID
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// 上下文日志基准时间戳
	SurroundTs *uint64 `json:"SurroundTs,omitempty" name:"SurroundTs"`

	// 部署组ID列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 上下文日志数量，取值范围[10, 100]，默认值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 上下文查询方向，B：双向，U：前向，D：向后
	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

func (r *SearchSurroundBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchSurroundBusinessLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchSurroundBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志列表
		Result *TsfPageBusinessLogV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSurroundBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchSurroundBusinessLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchTraceRequest struct {
	*tchttp.BaseRequest

	// 搜索开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 搜索结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调用方服务名
	CallerServiceName *string `json:"CallerServiceName,omitempty" name:"CallerServiceName"`

	// 被调方服务名
	CalleeServiceName *string `json:"CalleeServiceName,omitempty" name:"CalleeServiceName"`

	// 调用方接口名
	CallerInterfaceName *string `json:"CallerInterfaceName,omitempty" name:"CallerInterfaceName"`

	// 被调方接口名
	CalleeInterfaceName *string `json:"CalleeInterfaceName,omitempty" name:"CalleeInterfaceName"`

	// 错误码
	ResultStatus *string `json:"ResultStatus,omitempty" name:"ResultStatus"`

	// 错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 是否仅展示错误调用链
	ErrorTrace *bool `json:"ErrorTrace,omitempty" name:"ErrorTrace"`

	// 调用方IP
	CallerIp *string `json:"CallerIp,omitempty" name:"CallerIp"`

	// 被调方IP
	CalleeIp *string `json:"CalleeIp,omitempty" name:"CalleeIp"`

	// 最小耗时
	MinDuration *uint64 `json:"MinDuration,omitempty" name:"MinDuration"`

	// 最长耗时
	MaxDuration *uint64 `json:"MaxDuration,omitempty" name:"MaxDuration"`

	// 排序方式，支持"timestamp"和"duration"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方向，支持"asc"和"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 返回trace数目，取值范围[0, 100]，默认值20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// TSF集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// TSF命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 查询trace偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 用户自定义标签
	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

func (r *SearchTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchTraceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchTraceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调用链列表
		Result *TsfPageZipkinTraceInfoV2 `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchTraceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SecretKeyInfo struct {

	// secret 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// secret Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// secret key
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 创建时间，格式 yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// enabled: 启用状态。disabled: 禁用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 密钥ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// secret Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GwSecretId *string `json:"GwSecretId,omitempty" name:"GwSecretId"`

	// 过期时间，格式 yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 更新时间，格式 yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// API分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type ServerlessGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 服务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 程序包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// vpc ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 所需实例内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *string `json:"Memory,omitempty" name:"Memory"`

	// 要求最小实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceRequest *uint64 `json:"InstanceRequest,omitempty" name:"InstanceRequest"`

	// 部署组启动参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type ServerlessGroupPage struct {

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ServerlessGroup `json:"Content,omitempty" name:"Content" list`
}

type Service struct {

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCode *string `json:"ServiceCode,omitempty" name:"ServiceCode"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 删除标识，true: 可以删除；false: 不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`
}

type ServiceStatistics struct {

	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 总请求量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestAmount *string `json:"RequestAmount,omitempty" name:"RequestAmount"`

	// 请求成功量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessAmount *string `json:"SuccessAmount,omitempty" name:"SuccessAmount"`

	// 请求成功率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessRate *string `json:"SuccessRate,omitempty" name:"SuccessRate"`

	// 请求最大耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDuration *string `json:"MaxDuration,omitempty" name:"MaxDuration"`

	// 请求最小耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinDuration *string `json:"MinDuration,omitempty" name:"MinDuration"`

	// 请求平均耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgDuration *string `json:"AvgDuration,omitempty" name:"AvgDuration"`
}

type ServiceStatisticsV2 struct {

	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 总请求量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqTotalQty *uint64 `json:"ReqTotalQty,omitempty" name:"ReqTotalQty"`

	// 请求成功率
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqSuccessRate *float64 `json:"ReqSuccessRate,omitempty" name:"ReqSuccessRate"`

	// 请求平均耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqAvgDuration *float64 `json:"ReqAvgDuration,omitempty" name:"ReqAvgDuration"`
}

type SetCkafkaToContainerClusterRequest struct {
	*tchttp.BaseRequest

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ckafka instanceid
	CkafkaInstanceid *string `json:"CkafkaInstanceid,omitempty" name:"CkafkaInstanceid"`

	// ckafka topicid
	CkafkaTopicid *string `json:"CkafkaTopicid,omitempty" name:"CkafkaTopicid"`
}

func (r *SetCkafkaToContainerClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetCkafkaToContainerClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetCkafkaToContainerClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCkafkaToContainerClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetCkafkaToContainerClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetNamespaceCodeRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间编码
	NamespaceCode *string `json:"NamespaceCode,omitempty" name:"NamespaceCode"`
}

func (r *SetNamespaceCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetNamespaceCodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetNamespaceCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNamespaceCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetNamespaceCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShirkGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ShirkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShirkGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShirkGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShirkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShirkGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShirkInstanceRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 下线机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *ShirkInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShirkInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShirkInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShirkInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShirkInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShirkNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 机器id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ShirkNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShirkNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShirkNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 移除成功或失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShirkNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShirkNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ShrinkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShrinkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkInstanceRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 下线机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *ShrinkInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShrinkInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkInstancesRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 下线机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *ShrinkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShrinkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SimpleApplication struct {

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// ApplicationDesc
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// ProgLang
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgLang *string `json:"ProgLang,omitempty" name:"ProgLang"`

	// ApplicationResourceType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// UpdateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SimpleGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type StartContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StartContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 启动操作是否成功。
	// true：启动成功
	// false：启动失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StartGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartInstanceRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *StartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StartServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StatisticsEntry struct {

	// 主字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrimeKey *string `json:"PrimeKey,omitempty" name:"PrimeKey"`

	// 附字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubKey *string `json:"SubKey,omitempty" name:"SubKey"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type StdoutLogV2 struct {

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 日志时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIp *string `json:"InstanceIp,omitempty" name:"InstanceIp"`
}

type StopContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StopContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 停止操作是否成功。
	// true：停止成功
	// flase：停止失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StopGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstanceRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *StopInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StopServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SynchronizeContainerClusterRequest struct {
	*tchttp.BaseRequest

	// 容器平台的集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群描述
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群所属TSF地域
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF可用区
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`
}

func (r *SynchronizeContainerClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SynchronizeContainerClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SynchronizeContainerClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 同步是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SynchronizeContainerClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SynchronizeContainerClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TagRouteItemList struct {

	// 标签路由，标签类型，表示系统标签或自定义标签, 系统标签： S, 自定义标签： U
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 标签路由匹配源字段
	SourceField *string `json:"SourceField,omitempty" name:"SourceField"`

	// 标签路由匹配源规则, 等于: EQUAL， 不等于： NOT_EQUAL， 包含： IN， 不包含： NOT_IN， 正则表达式： REGEX
	SourceMatchRule *string `json:"SourceMatchRule,omitempty" name:"SourceMatchRule"`

	// 标签路由匹配源取值
	SourceValue *string `json:"SourceValue,omitempty" name:"SourceValue"`

	// 标签路由匹配目标字段
	TargetField *string `json:"TargetField,omitempty" name:"TargetField"`

	// 标签路由匹配目标取值
	TargetValue *string `json:"TargetValue,omitempty" name:"TargetValue"`

	// 标签路由规则项Id
	TagRouteId *string `json:"TagRouteId,omitempty" name:"TagRouteId"`

	// 标签路由所属路由Id
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`
}

type TaskId struct {

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type TaskInstance struct {

	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// status
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// mtime
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mtime *string `json:"Mtime,omitempty" name:"Mtime"`
}

type TemplateProject struct {

	// 工程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工程名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 包路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasePackage *string `json:"BasePackage,omitempty" name:"BasePackage"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *int64 `json:"LastTime,omitempty" name:"LastTime"`

	// Data
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// SubAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type TemplateResult struct {

	// TotalCount
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Content
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*TemplateProject `json:"Content,omitempty" name:"Content" list`
}

type TsfApiListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// API 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Content []*MsApiArray `json:"Content,omitempty" name:"Content" list`
	} `json:"Response"`
}

func (r *TsfApiListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TsfApiListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TsfPageApiDetailInfo struct {

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API 信息列表
	Content []*ApiDetailInfo `json:"Content,omitempty" name:"Content" list`
}

type TsfPageApiGroupInfo struct {

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API分组信息
	Content []*ApiGroupInfo `json:"Content,omitempty" name:"Content" list`
}

type TsfPageApplication struct {

	// 应用总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 应用信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ApplicationForPage `json:"Content,omitempty" name:"Content" list`
}

type TsfPageApplicatoinServerLog struct {

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 日志内容
	Content []*ApplicationServerLogContent `json:"Content,omitempty" name:"Content" list`
}

type TsfPageAuthRule struct {

	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 微服务权限规则列表
	Content []*AuthRule `json:"Content,omitempty" name:"Content" list`
}

type TsfPageAuthorization struct {

	// 鉴权规则是否被启用
	IsEnabled *bool `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 鉴权规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*AuthCondition `json:"Conditions,omitempty" name:"Conditions" list`
}

type TsfPageBusinessLogConfig struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 业务日志配置项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*BusinessLogConfig `json:"Content,omitempty" name:"Content" list`
}

type TsfPageBusinessLogV2 struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 业务日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*BusinessLogV2 `json:"Content,omitempty" name:"Content" list`

	// 游标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 查询状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type TsfPageCluster struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Cluster `json:"Content,omitempty" name:"Content" list`
}

type TsfPageClusterV2 struct {

	// 集群总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ClusterV2 `json:"Content,omitempty" name:"Content" list`
}

type TsfPageCommonPkg struct {

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 公共包信息
	Content []*CommonPkg `json:"Content,omitempty" name:"Content" list`
}

type TsfPageConfig struct {

	// TsfPageConfig
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项列表
	Content []*Config `json:"Content,omitempty" name:"Content" list`
}

type TsfPageConfigRelease struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项发布信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ConfigRelease `json:"Content,omitempty" name:"Content" list`
}

type TsfPageConfigReleaseLog struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项发布日志数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ConfigReleaseLog `json:"Content,omitempty" name:"Content" list`
}

type TsfPageFileConfig struct {

	// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 文件配置数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*FileConfig `json:"Content,omitempty" name:"Content" list`
}

type TsfPageFileConfigRelease struct {

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*FileConfigRelease `json:"Content,omitempty" name:"Content" list`
}

type TsfPageFileConfigReleaseLog struct {

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*FileConfigReleaseLog `json:"Content,omitempty" name:"Content" list`
}

type TsfPageGatewayDeployGroup struct {

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 记录实体列表
	Content []*GatewayDeployGroup `json:"Content,omitempty" name:"Content" list`
}

type TsfPageGatewayPlugin struct {

	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 记录实体列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*GatewayPlugin `json:"Content,omitempty" name:"Content" list`
}

type TsfPageInstance struct {

	// 机器实例总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 机器实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Instance `json:"Content,omitempty" name:"Content" list`
}

type TsfPageInvocationStatisticsV2 struct {

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 内容列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*InvocationStatisticsV2 `json:"Content,omitempty" name:"Content" list`
}

type TsfPageMicroservice struct {

	// 微服务总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 微服务列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Microservice `json:"Content,omitempty" name:"Content" list`
}

type TsfPageMonitorGroup struct {

	// 部署组总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 部署组列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*MonitorGroup `json:"Content,omitempty" name:"Content" list`
}

type TsfPageMsInstance struct {

	// 微服务实例总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 微服务实例列表内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*MsInstance `json:"Content,omitempty" name:"Content" list`
}

type TsfPageMsRunningApplicationV2 struct {

	// 运行微服务的应用个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 运行微服务的应用列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*MsRunningApplicationV2 `json:"Content,omitempty" name:"Content" list`
}

type TsfPageNamespace struct {

	// 命名空间总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 命名空间列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Namespace `json:"Content,omitempty" name:"Content" list`
}

type TsfPageRecordV2 struct {

	// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 操作记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*RecordV2 `json:"Content,omitempty" name:"Content" list`
}

type TsfPageRegion struct {

	// TSF地域总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// TSF地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*TsfRegion `json:"Content,omitempty" name:"Content" list`
}

type TsfPageRouteAffinity struct {

	// TSF就近策略总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// TSF就近策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*RouteAffinity `json:"Content,omitempty" name:"Content" list`
}

type TsfPageRouteReleaseHistory struct {

	// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 路由规则部署记录列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*RouteReleaseHistory `json:"Content,omitempty" name:"Content" list`
}

type TsfPageRouteRule struct {

	// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 路由规则数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*RouteRule `json:"Content,omitempty" name:"Content" list`
}

type TsfPageRouteV2 struct {

	// 服务路由总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务路由详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*RouteV2 `json:"Content,omitempty" name:"Content" list`
}

type TsfPageSimpleApplication struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 简单应用列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*SimpleApplication `json:"Content,omitempty" name:"Content" list`
}

type TsfPageSimpleGroup struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 简单部署组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*SimpleGroup `json:"Content,omitempty" name:"Content" list`
}

type TsfPageStdoutLogV2 struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 标准输出日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*StdoutLogV2 `json:"Content,omitempty" name:"Content" list`

	// 游标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 查询状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type TsfPageTraceInterface struct {

	// 符合条件的总调用链接口数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 调用链接口列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*string `json:"Content,omitempty" name:"Content" list`
}

type TsfPageTraceService struct {

	// 符合条件的总调用链服务名数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 调用链服务名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*string `json:"Content,omitempty" name:"Content" list`
}

type TsfPageVmGroup struct {

	// 虚拟机部署组总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 虚拟机部署组列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*VmGroupSimple `json:"Content,omitempty" name:"Content" list`
}

type TsfPageVmSubTask struct {

	// 子任务总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 成功子任务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 运行中子任务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCount *int64 `json:"RunCount,omitempty" name:"RunCount"`

	// 失败子任务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`

	// 子任务列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*VmSubTask `json:"Content,omitempty" name:"Content" list`
}

type TsfPageVmTask struct {

	// 任务总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*VmTask `json:"Content,omitempty" name:"Content" list`
}

type TsfPageZipkinSpanInfo struct {

	// 总条目数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 调用链Span列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ZipkinSpanInfo `json:"Content,omitempty" name:"Content" list`
}

type TsfPageZipkinSpanInfoV2 struct {

	// 总条目数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 调用链Span列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ZipkinSpanInfoV2 `json:"Content,omitempty" name:"Content" list`
}

type TsfPageZipkinTraceInfoV2 struct {

	// 总条目数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 调用链列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ZipkinTraceInfoV2 `json:"Content,omitempty" name:"Content" list`
}

type TsfPageZone struct {

	// TSF可用区总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// TSF可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*TsfZone `json:"Content,omitempty" name:"Content" list`
}

type TsfRecordCodeResult struct {

	// 操作类型术语
	Operations []*TypeCode `json:"Operations,omitempty" name:"Operations" list`

	// 模块类型术语
	Modules []*TypeCode `json:"Modules,omitempty" name:"Modules" list`
}

type TsfRegion struct {

	// TSF地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// TSF地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type TsfRegionResult struct {

	// tRegionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`

	// tRegionName
	// 注意：此字段可能返回 null，表示取不到有效值。
	TRegionName *string `json:"TRegionName,omitempty" name:"TRegionName"`

	// tRemark
	// 注意：此字段可能返回 null，表示取不到有效值。
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`
}

type TsfZone struct {

	// TSF可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// TSF可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// TSF可用区所属TSF地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`
}

type TsfZoneResult struct {

	// tZoneId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TZoneId *string `json:"TZoneId,omitempty" name:"TZoneId"`

	// tZoneName
	// 注意：此字段可能返回 null，表示取不到有效值。
	TZoneName *string `json:"TZoneName,omitempty" name:"TZoneName"`

	// tRemark
	// 注意：此字段可能返回 null，表示取不到有效值。
	TRemark *string `json:"TRemark,omitempty" name:"TRemark"`

	// tRegionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TRegionId *string `json:"TRegionId,omitempty" name:"TRegionId"`
}

type TxError struct {

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type TxList struct {

	// 返回的事务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 主事务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*TxMainTransaction `json:"Content,omitempty" name:"Content" list`

	// 请求异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *TxError `json:"Error,omitempty" name:"Error"`
}

type TxMainTransaction struct {

	// 事务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 服务器Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 方法名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MethodName *string `json:"MethodName,omitempty" name:"MethodName"`

	// 子事务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubCount *int64 `json:"SubCount,omitempty" name:"SubCount"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 事务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 超时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutMs *int64 `json:"TimeoutMs,omitempty" name:"TimeoutMs"`

	// 重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryTimes *int64 `json:"RetryTimes,omitempty" name:"RetryTimes"`

	// 是否开启自动重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRetry *bool `json:"AutoRetry,omitempty" name:"AutoRetry"`
}

type TxRetry struct {

	// 重试成功的事务ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success []*string `json:"Success,omitempty" name:"Success" list`

	// 重试失败的事务ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Missing []*string `json:"Missing,omitempty" name:"Missing" list`

	// 重试异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *TxError `json:"Error,omitempty" name:"Error"`
}

type TypeCode struct {

	// 术语英文描述
	Code *string `json:"Code,omitempty" name:"Code"`

	// 术语中文描述
	Name *string `json:"Name,omitempty" name:"Name"`
}

type UnbindApiGroupRequest struct {
	*tchttp.BaseRequest

	// 分组网关id列表
	GroupGatewayList []*GatewayGroupIds `json:"GroupGatewayList,omitempty" name:"GroupGatewayList" list`
}

func (r *UnbindApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindApiGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果，成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindApiGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindPluginRequest struct {
	*tchttp.BaseRequest

	// 分组/API解绑插件列表
	PluginInstanceList []*GatewayPluginBoundParam `json:"PluginInstanceList,omitempty" name:"PluginInstanceList" list`
}

func (r *UnbindPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果，成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UninstallAgentScriptsRequest struct {
	*tchttp.BaseRequest

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *UninstallAgentScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UninstallAgentScriptsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UninstallAgentScriptsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 卸载agent脚本内容
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallAgentScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UninstallAgentScriptsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiGroupRequest struct {
	*tchttp.BaseRequest

	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Api 分组描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 鉴权类型
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 分组上下文
	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`
}

func (r *UpdateApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果，true: 成功, false: 失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiRateLimitRuleRequest struct {
	*tchttp.BaseRequest

	// 限流规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 开启/禁用，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// qps值，开启限流规则时，必填
	MaxQps *int64 `json:"MaxQps,omitempty" name:"MaxQps"`
}

func (r *UpdateApiRateLimitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiRateLimitRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiRateLimitRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApiRateLimitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiRateLimitRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationAgentRequest struct {
	*tchttp.BaseRequest

	// agentIds
	AgentIds []*string `json:"AgentIds,omitempty" name:"AgentIds" list`

	// PackageSnip
	PackageSnip *string `json:"PackageSnip,omitempty" name:"PackageSnip"`

	// PackageName
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// PackageMd5v
	PackageMd5v *string `json:"PackageMd5v,omitempty" name:"PackageMd5v"`

	// PackageType
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// PackageSize
	PackageSize *int64 `json:"PackageSize,omitempty" name:"PackageSize"`

	// PackageAddrs
	PackageAddrs *string `json:"PackageAddrs,omitempty" name:"PackageAddrs"`

	// PackageId
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`
}

func (r *UpdateApplicationAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApplicationAgentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationAgentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApplicationAgentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateBusinessLogConfigRequest struct {
	*tchttp.BaseRequest

	// 业务日志配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 业务日志配置名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 业务日志配置路径
	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`

	// 业务日志配置描述
	ConfigDesc *string `json:"ConfigDesc,omitempty" name:"ConfigDesc"`

	// 业务日志配置标签
	ConfigTags *string `json:"ConfigTags,omitempty" name:"ConfigTags"`

	// 业务日志配置解析规则
	ConfigSchema *BusinessLogConfigSchema `json:"ConfigSchema,omitempty" name:"ConfigSchema"`
}

func (r *UpdateBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateBusinessLogConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新结果true：创建成功；false：创建失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateBusinessLogConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCircuitBreakerRuleRequest struct {
	*tchttp.BaseRequest

	// 熔断规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 熔断规则微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务所属命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 目标服务名
	TargetServiceName *string `json:"TargetServiceName,omitempty" name:"TargetServiceName"`

	// 目标服务所属命名空间
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// 熔断策略列表
	StrategyList []*CircuitBreakerStrategy `json:"StrategyList,omitempty" name:"StrategyList" list`

	// Enable
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 熔断隔离级别
	IsolationLevel *string `json:"IsolationLevel,omitempty" name:"IsolationLevel"`

	// 微服务所属命名空间
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" name:"TargetNamespaceName"`
}

func (r *UpdateCircuitBreakerRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCircuitBreakerRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCircuitBreakerRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true false
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCircuitBreakerRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCircuitBreakerRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 配置模板id
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`

	// 配置模板名称
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板对应的微服务框架
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 配置模板数据
	ConfigTemplateValue *string `json:"ConfigTemplateValue,omitempty" name:"ConfigTemplateValue"`

	// 配置模板描述
	ConfigTemplateDesc *string `json:"ConfigTemplateDesc,omitempty" name:"ConfigTemplateDesc"`
}

func (r *UpdateConfigTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateConfigTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateConfigTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateConfigTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateConfigTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 端口列表
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔,单位秒
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`
}

func (r *UpdateContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayJwtPluginRequest struct {
	*tchttp.BaseRequest

	// 网关插件id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件类型 "Jwt"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 公钥对kid
	Kid *string `json:"Kid,omitempty" name:"Kid"`

	// 公钥对Json串
	PublicKeyJson *string `json:"PublicKeyJson,omitempty" name:"PublicKeyJson"`

	// token携带位置query/header
	TokenBaggagePosition *string `json:"TokenBaggagePosition,omitempty" name:"TokenBaggagePosition"`

	// token的key值
	TokenKeyName *string `json:"TokenKeyName,omitempty" name:"TokenKeyName"`

	// claim参数映射关系json
	ClaimMappingJson *string `json:"ClaimMappingJson,omitempty" name:"ClaimMappingJson"`

	// 插件描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 重定向地址
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

func (r *UpdateGatewayJwtPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGatewayJwtPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayJwtPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果，成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGatewayJwtPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGatewayJwtPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayOAuthPluginRequest struct {
	*tchttp.BaseRequest

	// 网关插件id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件类型 "OAuth"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 验证token路径
	TokenAuthUrl *string `json:"TokenAuthUrl,omitempty" name:"TokenAuthUrl"`

	// 验证token请求方法:get/post
	TokenAuthMethod *string `json:"TokenAuthMethod,omitempty" name:"TokenAuthMethod"`

	// 认证请求超时时间,单位:秒 范围:0~30
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// token携带位置，网关取token位置与发送认证请求时token位置一致,值:cookie/query/header
	TokenBaggagePosition *string `json:"TokenBaggagePosition,omitempty" name:"TokenBaggagePosition"`

	// token的key值
	TokenKeyName *string `json:"TokenKeyName,omitempty" name:"TokenKeyName"`

	// payload的映射参数名称
	PayloadMappingName *string `json:"PayloadMappingName,omitempty" name:"PayloadMappingName"`

	// payload映射到后端服务的携带位置,值:cookie/query/header
	PayloadMappingPosition *string `json:"PayloadMappingPosition,omitempty" name:"PayloadMappingPosition"`

	// 插件描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 重定向地址
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`
}

func (r *UpdateGatewayOAuthPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGatewayOAuthPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayOAuthPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果，成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGatewayOAuthPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGatewayOAuthPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayTagPluginRequest struct {
	*tchttp.BaseRequest

	// 网关插件id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件类型 "tag"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数配置Jason串
	TagPluginInfoList *string `json:"TagPluginInfoList,omitempty" name:"TagPluginInfoList"`

	// 插件描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateGatewayTagPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGatewayTagPluginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayTagPluginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果，成功失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGatewayTagPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGatewayTagPluginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupSecretRequest struct {
	*tchttp.BaseRequest

	// 秘钥ID
	GwSecretId *string `json:"GwSecretId,omitempty" name:"GwSecretId"`
}

func (r *UpdateGroupSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGroupSecretRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *SecretKeyInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGroupSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGroupSecretResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRatelimitRequest struct {
	*tchttp.BaseRequest

	// 名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 限流所作用的目标微服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Rule列表，Id必填，其他字段仅修改时填写
	Rules []*RatelimitRuleForUpdateV2 `json:"Rules,omitempty" name:"Rules" list`
}

func (r *UpdateRatelimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRatelimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRatelimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功与否
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRatelimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRatelimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateUploadInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 调用GetUploadInfo接口时返回的软件包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// COS返回上传结果（默认为0：成功，其他值表示失败）
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 程序包MD5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 程序包大小
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

func (r *UpdateUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateUploadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateUploadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadLicenseApplicationRequest struct {
	*tchttp.BaseRequest

	// 授予对象
	Grantee *LicenseGrantee `json:"Grantee,omitempty" name:"Grantee"`

	// 申请正文
	ApplyContent *string `json:"ApplyContent,omitempty" name:"ApplyContent"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UploadLicenseApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadLicenseApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadLicenseApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功与否
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadLicenseApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadLicenseApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadPkgRequest struct {
	*tchttp.BaseRequest

	// 无
	File *string `json:"File,omitempty" name:"File"`

	// 无
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 无
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 无
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 无
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 无
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 无
	UploadId *string `json:"UploadId,omitempty" name:"UploadId"`

	// 无
	FileSegs *int64 `json:"FileSegs,omitempty" name:"FileSegs"`

	// 无
	FileSeg *int64 `json:"FileSeg,omitempty" name:"FileSeg"`

	// 无
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`
}

func (r *UploadPkgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadPkgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadPkgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码
		Code *int64 `json:"Code,omitempty" name:"Code"`

		// 错误说明
		Message *string `json:"Message,omitempty" name:"Message"`

		// 本次上传ID
		UploadId *string `json:"UploadId,omitempty" name:"UploadId"`

		// 上传包ID
		PkgId []*string `json:"PkgId,omitempty" name:"PkgId" list`

		// 请求ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadPkgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadPkgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateDeleteConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *ValidateDeleteConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateDeleteConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateDeleteConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：可以删除；false：不可删除
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateDeleteConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateDeleteConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateDeletePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *ValidateDeletePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateDeletePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateDeletePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateDeletePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateDeletePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateLogSchemaRequest struct {
	*tchttp.BaseRequest

	// 解析规则类型
	SchemaType *uint64 `json:"SchemaType,omitempty" name:"SchemaType"`

	// 解析规则内容，包含解析规则时间格式和解析规则内容
	SchemaPatternLayout *string `json:"SchemaPatternLayout,omitempty" name:"SchemaPatternLayout"`
}

func (r *ValidateLogSchemaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateLogSchemaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateLogSchemaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 校验结果
		Result *BusinessLogPatternAnalysis `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateLogSchemaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateLogSchemaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateNamespaceClusterVPCRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ValidateNamespaceClusterVPCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateNamespaceClusterVPCRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateNamespaceClusterVPCResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：一致；false：不一致
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateNamespaceClusterVPCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateNamespaceClusterVPCResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VmGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 程序包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 程序包版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 部署组机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 部署组运行中机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 部署组启动参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 部署组创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 部署组更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 部署组停止机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffInstanceCount *int64 `json:"OffInstanceCount,omitempty" name:"OffInstanceCount"`

	// 部署组描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`
}

type VmGroupOther struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 程序包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 部署组实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 部署组运行中实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 部署组中停止实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffInstanceCount *int64 `json:"OffInstanceCount,omitempty" name:"OffInstanceCount"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`
}

type VmGroupSimple struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组启动参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`
}

type VmInstanceResourceConfig struct {

	// 实例导入方式，可多个，公有云为 ["R", "M"]，独立版的取值仅有 "M" 脚本模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportMode []*string `json:"ImportMode,omitempty" name:"ImportMode" list`
}

type VmSubTask struct {

	// 机器Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 机器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 机器Agent状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentState *int64 `json:"AgentState,omitempty" name:"AgentState"`

	// 机器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *int64 `json:"InstanceState,omitempty" name:"InstanceState"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 任务记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRecord *string `json:"TaskRecord,omitempty" name:"TaskRecord"`
}

type VmTask struct {

	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 分组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 程序包Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 程序包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`

	// 子任务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 成功的子任务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 运行中的子任务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCount *int64 `json:"RunCount,omitempty" name:"RunCount"`

	// 失败的子任务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
}

type VportTypeResult struct {

	// VportId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VportId *string `json:"VportId,omitempty" name:"VportId"`

	// ModuleId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// VportType
	// 注意：此字段可能返回 null，表示取不到有效值。
	VportType *string `json:"VportType,omitempty" name:"VportType"`

	// DefaultVport
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultVport *string `json:"DefaultVport,omitempty" name:"DefaultVport"`

	// Vport
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *string `json:"Vport,omitempty" name:"Vport"`
}

type WeightRouteItemList struct {

	// 权重路由，百分比字段
	SourcePercent *int64 `json:"SourcePercent,omitempty" name:"SourcePercent"`

	// 权重路由规则匹配目标字段
	TargetField *string `json:"TargetField,omitempty" name:"TargetField"`

	// 权重路由规则匹配目标取值
	TargetValue *string `json:"TargetValue,omitempty" name:"TargetValue"`

	// 权重路由规则项Id
	WeightRouteId *string `json:"WeightRouteId,omitempty" name:"WeightRouteId"`

	// 权重路由规则所属路由Id
	RouteRuleId *string `json:"RouteRuleId,omitempty" name:"RouteRuleId"`
}

type ZipkinAnnotation struct {

	// 注解值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 注解时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 注解端点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *ZipkinEndPoint `json:"Endpoint,omitempty" name:"Endpoint"`
}

type ZipkinAnnotationV2 struct {

	// 注解值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 注解时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 注解端点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *ZipkinEndpointV2 `json:"Endpoint,omitempty" name:"Endpoint"`
}

type ZipkinBinaryAnnotation struct {

	// 注解键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 注解值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 注解端点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *ZipkinEndPoint `json:"Endpoint,omitempty" name:"Endpoint"`
}

type ZipkinBinaryAnnotationV2 struct {

	// 注解键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 注解值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 注解端点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *ZipkinEndpointV2 `json:"Endpoint,omitempty" name:"Endpoint"`
}

type ZipkinEndPoint struct {

	// 端点服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 端点IP地址（v4）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv4 *string `json:"Ipv4,omitempty" name:"Ipv4"`

	// 端点IP地址（v6）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *string `json:"Ipv6,omitempty" name:"Ipv6"`

	// 端点端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type ZipkinEndpointV2 struct {

	// 端点服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 端点IP地址（v4）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv4 *string `json:"Ipv4,omitempty" name:"Ipv4"`

	// 端点IP地址（v6）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *string `json:"Ipv6,omitempty" name:"Ipv6"`

	// 端点端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type ZipkinMetadata struct {

	// 元数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 元数据时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 元数据端点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *ZipkinEndPoint `json:"Endpoint,omitempty" name:"Endpoint"`
}

type ZipkinMetadataV2 struct {

	// 元数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 元数据时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 元数据端点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *ZipkinEndpointV2 `json:"Endpoint,omitempty" name:"Endpoint"`
}

type ZipkinSpanInfo struct {

	// Span ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 调用链ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 父Span ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`

	// Span名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Span时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Span耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Span结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultStatus *string `json:"ResultStatus,omitempty" name:"ResultStatus"`

	// Span服务名（前端展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Span注解列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationList []*ZipkinAnnotation `json:"AnnotationList,omitempty" name:"AnnotationList" list`

	// Span二进制注解列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BinaryAnnotationList []*ZipkinBinaryAnnotation `json:"BinaryAnnotationList,omitempty" name:"BinaryAnnotationList" list`

	// Span元数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetadataList []*ZipkinMetadata `json:"MetadataList,omitempty" name:"MetadataList" list`
}

type ZipkinSpanInfoV2 struct {

	// Span ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 调用链ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 父Span ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`

	// Span名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Span时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Span耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// Span结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultStatus *string `json:"ResultStatus,omitempty" name:"ResultStatus"`

	// Span服务名（前端展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Span注解列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationList []*ZipkinAnnotationV2 `json:"AnnotationList,omitempty" name:"AnnotationList" list`

	// Span二进制注解列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BinaryAnnotationList []*ZipkinBinaryAnnotationV2 `json:"BinaryAnnotationList,omitempty" name:"BinaryAnnotationList" list`

	// Span元数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetadataList []*ZipkinMetadataV2 `json:"MetadataList,omitempty" name:"MetadataList" list`

	// Span接口名（前端展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterfaceName *string `json:"InterfaceName,omitempty" name:"InterfaceName"`

	// Span本地端点
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalEndpoint *ZipkinEndpointV2 `json:"LocalEndpoint,omitempty" name:"LocalEndpoint"`

	// Span远程端点
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteEndpoint *ZipkinEndpointV2 `json:"RemoteEndpoint,omitempty" name:"RemoteEndpoint"`
}

type ZipkinTraceInfoV2 struct {

	// 调用链ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 调用链时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 调用链时耗
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 调用链结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultStatus *string `json:"ResultStatus,omitempty" name:"ResultStatus"`

	// 调用链错误标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultError *bool `json:"ResultError,omitempty" name:"ResultError"`
}
