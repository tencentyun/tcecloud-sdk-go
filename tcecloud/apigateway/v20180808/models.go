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

package v20180808

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountLimitLogRuleClount struct {

	// 日志规则
	LogRuleClount *int64 `json:"LogRuleClount,omitempty" name:"LogRuleClount"`
}

type AccountLimitSecretIdCount struct {

	// 密钥计数
	SecretIdCount *int64 `json:"SecretIdCount,omitempty" name:"SecretIdCount"`

	// 密钥的使用计划
	UsagePlanCountForSecretId *int64 `json:"UsagePlanCountForSecretId,omitempty" name:"UsagePlanCountForSecretId"`
}

type AccountLimitServiceCount struct {

	// 服务数量
	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`

	// 服务的api数量
	ApiCountInService *int64 `json:"ApiCountInService,omitempty" name:"ApiCountInService"`

	// 服务的自定义域名数量
	DomainCountInService *int64 `json:"DomainCountInService,omitempty" name:"DomainCountInService"`
}

type AccountLimitUsagePlanCount struct {

	// 使用计划数量
	UsagePlanCount *int64 `json:"UsagePlanCount,omitempty" name:"UsagePlanCount"`

	// 密钥使用计划
	SecretIdCountInUsagePlan *int64 `json:"SecretIdCountInUsagePlan,omitempty" name:"SecretIdCountInUsagePlan"`

	// QPS使用计划
	MaxQPSInUsagePlan *int64 `json:"MaxQPSInUsagePlan,omitempty" name:"MaxQPSInUsagePlan"`

	// 阶段计数使用计划
	StageCountInUsagePlan *int64 `json:"StageCountInUsagePlan,omitempty" name:"StageCountInUsagePlan"`
}

type AccountSettings struct {

	// 用户配额
	AccountLimit *Limit `json:"AccountLimit,omitempty" name:"AccountLimit"`

	// 使用量
	AccountUsage *Usage `json:"AccountUsage,omitempty" name:"AccountUsage"`
}

type AccountUsageLogRuleClount struct {

	// 日志使用计数
	LogRuleClount *int64 `json:"LogRuleClount,omitempty" name:"LogRuleClount"`
}

type AccountUsageSecretIdCount struct {

	// 密钥使用计数
	SecretIdCount *int64 `json:"SecretIdCount,omitempty" name:"SecretIdCount"`
}

type AccountUsageServiceCount struct {

	// 服务数量
	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
}

type AccountUsageUsagePlanCount struct {

	// 使用计划计数
	UsagePlanCount *int64 `json:"UsagePlanCount,omitempty" name:"UsagePlanCount"`
}

type ApiEnvironmentApiKeys struct {

	// 密钥数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 密钥id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKeyIdList []*string `json:"AccessKeyIdList,omitempty" name:"AccessKeyIdList" list`
}

type ApiEnvironmentStrategy struct {

	// api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api名称
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 环境限流信息
	EnvironmentStrategySet []*EnvironmentStrategy `json:"EnvironmentStrategySet,omitempty" name:"EnvironmentStrategySet" list`
}

type ApiEnvironmentStrategyStataus struct {

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiEnvironmentStrategySet []*ApiEnvironmentStrategy `json:"ApiEnvironmentStrategySet,omitempty" name:"ApiEnvironmentStrategySet" list`
}

type ApiErrorRate struct {

	// 用户可读的api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// api 错误率
	ErrorRate *float64 `json:"ErrorRate,omitempty" name:"ErrorRate"`

	// 用户可读的服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api类型
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

type ApiIdStatus struct {

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API描述
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API PATH。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API METHOD。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 服务创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 服务修改时间。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// VPC唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API协议。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 是否买后调试。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 授权类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// API业务类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// 关联授权API唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// 授权API关联的业务API列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds" list`

	// oauth配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// oauth2.0API请求，token存放位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`
}

type ApiInfo struct {

	// 服务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// api id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// api名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// api类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 鉴权类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 业务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// 关联授权api
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// 鉴权配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 是否购买后调试
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 请求配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// 返回类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// 返回成功例子
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// 返回失败例子
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// 返回错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseErrorCodes []*ErrorCodes `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`

	// 请求参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// 超时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// 后端服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 服务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// 服务相关参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// 常量参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// mock 返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// 函数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// 函数命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// 函数限定符
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// 是否开启集成响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// 函数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// 函数命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// 函数限定符
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// 函数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// 函数命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// 函数限定符
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// 内部域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternalDomain *string `json:"InternalDomain,omitempty" name:"InternalDomain"`

	// 函数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// 函数命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// 函数限定符
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// 服务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroServices []*MicroService `json:"MicroServices,omitempty" name:"MicroServices" list`

	// 微服务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroServicesInfo []*int64 `json:"MicroServicesInfo,omitempty" name:"MicroServicesInfo" list`

	// tsf后端负载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// tsf健康检测信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// 是否开启CORS
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// api tag数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type ApiKey struct {

	// 密钥id
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 密钥key
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`

	// 密钥类型
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// 密钥名称
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 密钥状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ApiKeysStatus struct {

	// 密钥数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 密钥列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiKeySet []*ApiKey `json:"ApiKeySet,omitempty" name:"ApiKeySet" list`
}

type ApiReqCount struct {

	// 用户的api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api的路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求数
	ReqCount *uint64 `json:"ReqCount,omitempty" name:"ReqCount"`

	// 用户可读的服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api类型。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

type ApiRequestConfig struct {

	// path
	Path *string `json:"Path,omitempty" name:"Path"`

	// 方法
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ApiResponseTime struct {

	// api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 响应时间（毫秒）
	ResponseTime *uint64 `json:"ResponseTime,omitempty" name:"ResponseTime"`

	// 用户可读的服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api类型。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

type ApiTags struct {

	// key
	Key *string `json:"Key,omitempty" name:"Key"`

	// values
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ApiUsagePlan struct {

	// 服务唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// api路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// api方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 使用计划id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 使用计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 环境名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 已使用配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// 最大请求量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 最大qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type ApiUsagePlanSet struct {

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// api绑定使用计划列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiUsagePlanList []*ApiUsagePlan `json:"ApiUsagePlanList,omitempty" name:"ApiUsagePlanList" list`
}

type ApisStatus struct {

	// 总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// api描述集合
	ApiIdStatusSet []*DesApisStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet" list`
}

type BindEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 使用计划id列表
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds" list`

	// 绑定类型
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api id列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *BindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindEnvironmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindEnvironmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// ip策略id
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 环境
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 绑定的api id列表
	BindApiIds []*string `json:"BindApiIds,omitempty" name:"BindApiIds" list`
}

func (r *BindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 使用计划id
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 密钥列表
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds" list`
}

func (r *BindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务Id。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要绑定的自定义域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 服务支持协议，可选值为http、https、http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络类型，可选值为OUTER、INNER。
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 是否使用默认路径映射，即针对三个Environment：“test”、 ”prepub“、”release“，Path：“/test”、 ”/prepub“、”/release“。
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 内部二级域名。
	NetSubDomain *string `json:"NetSubDomain,omitempty" name:"NetSubDomain"`

	// 证书Id。针对Protocol 为https或http&https可以选择上传。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 自定义域名路径映射，最多输入三个Environment，并且只能分别取值“test”、 ”prepub“、”release“。
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`
}

func (r *BindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckCloneApisRequest struct {
	*tchttp.BaseRequest

	// 源 serviceid
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 目的 serviceid
	DestServiceId *string `json:"DestServiceId,omitempty" name:"DestServiceId"`

	// 目的地域  gz hk sh bj 等缩写
	DestRegion *string `json:"DestRegion,omitempty" name:"DestRegion"`

	// 源Apiid 列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *CheckCloneApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckCloneApisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckCloneApisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回格式校验
		Result *CheckSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCloneApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckCloneApisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckServiceConfigRequest struct {
	*tchttp.BaseRequest

	// 服务id。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务类型，如：HTTP、WEBSOCKET。若检测后端地址，ServiceType、BackendUrl必传，前端和后端不能同时检测。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 后端服务地址，若检测后端地址，ServiceType、BackendUrl必传，前端和后端不能同时检测。
	BackendUrl *string `json:"BackendUrl,omitempty" name:"BackendUrl"`

	// 前端路径配置，若检测前端路径，RequestConfig与ServiceType、BackendUrl不可同时存在。
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`
}

func (r *CheckServiceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckServiceConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckServiceConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// True:验证通过；False:验证不通过。
		Status *bool `json:"Status,omitempty" name:"Status"`

		// 验证结果的详细描述信息。
		VerifyMessage *string `json:"VerifyMessage,omitempty" name:"VerifyMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckServiceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckServiceConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckSet struct {

	// 返回码 0 ok 非0 失败
	CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 错误信息
	CheckMessageList []*string `json:"CheckMessageList,omitempty" name:"CheckMessageList" list`
}

type CloneApisRequest struct {
	*tchttp.BaseRequest

	// 源serviceid
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 目的serviceid
	DestServiceId *string `json:"DestServiceId,omitempty" name:"DestServiceId"`

	// 目的region  gz cq sh 等缩写
	DestRegion *string `json:"DestRegion,omitempty" name:"DestRegion"`

	// 源apiids 数组
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *CloneApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneApisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloneApisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 克隆结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneApisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConstantParameter struct {

	// 常量参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 参数所在位置
	Position *string `json:"Position,omitempty" name:"Position"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
}

type CreateApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥名称
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 密钥类型
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// 密钥id
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 密钥key
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *CreateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥信息
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 后端服务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 后端超时时间
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 请求配置
	RequestConfig *ApiRequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// api名称
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// api描述
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// api类型，默认为NORMAL
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// 鉴权类型
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 是否开启跨域
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// 常量参数
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// 请求参数
	RequestParameters []*RequestParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// api类型
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// mock类型返回
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// api绑定tst服务信息
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices" list`

	// tsf负载均衡配置
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// tsf健康检查配置
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// 后端资源信息，入参应该不需要？
	TargetServices []*TargetServicesReq `json:"TargetServices,omitempty" name:"TargetServices" list`

	// target负载均衡配置
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// target健康检查配置
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// scf 函数名称
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// scf websocket注册函数
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// scf websocket清理函数
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// scf websocket传输函数
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// scf 函数命名空间
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// scf函数版本
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// scf websocket注册函数命名空间
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// scf websocket传输函数版本
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// scf websocket传输函数命名空间
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// scf websocket传输函数版本
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// scf websocket清理函数命名空间
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// scf websocket清理函数版本
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// 是否开启响应集成
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// 开始调试后计费
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 是否删除返回错误码
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// 返回类型
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// 成功返回示例
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// 失败返回示例
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// 后端服务配置
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// 关联的授权api id
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// 服务参数
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// oauth配置
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 错误码信息
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`

	// tsf serverless ns id
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// 用户类型
	UserType *string `json:"UserType,omitempty" name:"UserType"`
}

func (r *CreateApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api信息
		Result *CreateApiRsp `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiRsp struct {

	// api id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// path
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// method
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CreateIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 策略类型
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// 策略内容
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *CreateIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略信息
		Result *IPStrategy `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLogRuleRequest struct {
	*tchttp.BaseRequest

	// 日志规则名称。
	LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 服务唯一ID，表示所需上报日志的服务ID，不传代表所有服务。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称，表示所需要上报日志的环境名称，不传代表所有环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
}

func (r *CreateLogRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLogRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLogRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志规则名称。
		LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`

		// 服务唯一ID，表示所需上报日志的服务ID，不传代表所有服务。
		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

		// 环境名称，表示所需要上报日志的环境名称，不传代表所有环境。
		EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

		// 日志集ID。
		LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

		// 日志主题ID。
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 日志规则唯一ID。
		LogRuleId *string `json:"LogRuleId,omitempty" name:"LogRuleId"`

		// 创建时间。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLogRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest

	// 服务名称。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 服务支持协议，可选值为http、https、http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 服务描述。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 独立集群名称，用于指定创建服务所在的独立集群。
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// 网络类型列表，用于指定支持的访问类型，INNER为内网访问，OUTER为外网访问。默认为OUTER。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

	// ip version
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// 集群名
	SetServerName *string `json:"SetServerName,omitempty" name:"SetServerName"`

	// 用户类型
	AppIdType *string `json:"AppIdType,omitempty" name:"AppIdType"`
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

		// 服务唯一ID。
		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

		// 服务名称。
		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

		// 服务描述。
		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

		// 外网子域名。
		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

		// 内网子域名。
		InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`

		// 服务创建时间。
		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

		// 网络类型列表，INNER为内网访问，OUTER为外网访问。
		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

		// ip version
		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

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

type CreateTcbScfApiRequest struct {
	*tchttp.BaseRequest

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// SCF方法名称。
	ScfFunctionName *string `json:"ScfFunctionName,omitempty" name:"ScfFunctionName"`

	// SCF方法命名空间。
	ScfFunctionNamespace *string `json:"ScfFunctionNamespace,omitempty" name:"ScfFunctionNamespace"`
}

func (r *CreateTcbScfApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTcbScfApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTcbScfApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境ID。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 请求PATH。
		Path *string `json:"Path,omitempty" name:"Path"`

		// 请求子域名。
		SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTcbScfApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTcbScfApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划名称
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划描述
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 最大调用次数
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// qps
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *CreateUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划详情
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥id
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DeleteApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DeleteApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApisRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 接口id数组
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *DeleteApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// ip策略id
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLogRulesRequest struct {
	*tchttp.BaseRequest

	// 日志规则ID列表。
	LogRuleIds []*string `json:"LogRuleIds,omitempty" name:"LogRuleIds" list`
}

func (r *DeleteLogRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLogRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLogRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLogRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceSubDomainMappingRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 自定义域名
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 环境，取值：“test“，”prepub“，”release“
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DeleteServiceSubDomainMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceSubDomainMappingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceSubDomainMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作成功返回true
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceSubDomainMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceSubDomainMappingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTcbScfApiRequest struct {
	*tchttp.BaseRequest

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// SCF方法名称。
	ScfFunctionName *string `json:"ScfFunctionName,omitempty" name:"ScfFunctionName"`

	// SCF方法命名空间。
	ScfFunctionNamespace *string `json:"ScfFunctionNamespace,omitempty" name:"ScfFunctionNamespace"`
}

func (r *DeleteTcbScfApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTcbScfApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTcbScfApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTcbScfApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTcbScfApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划id
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DeleteUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DemoteServiceUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划id
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DemoteServiceUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DemoteServiceUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DemoteServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DemoteServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DemoteServiceUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DesApisStatus struct {

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// 服务创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 服务修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// VPCID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// VPC唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API协议。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 是否买后调试。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 授权类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// API业务类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// 关联授权API唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// oauth配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// oauth2.0API请求，token存放位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`

	// 授权API关联的业务API列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds" list`

	// Tag数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`
}

type DescribeApiEnvironmentApiKeysRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 环境名
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
}

func (r *DescribeApiEnvironmentApiKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiEnvironmentApiKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiEnvironmentApiKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api绑定密钥列表
		Result *ApiEnvironmentApiKeys `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiEnvironmentApiKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiEnvironmentApiKeysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境列表
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames" list`

	// api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api绑定策略详情
		Result *ApiEnvironmentStrategyStataus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥id
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DescribeApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥详情
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeysStatusRequest struct {
	*tchttp.BaseRequest

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`
}

func (r *DescribeApiKeysStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeysStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeysStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥列表
		Result *ApiKeysStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiKeysStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeysStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DescribeApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api详情
		Result *ApiInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 服务唯一id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api绑定使用计划列表
		Result *ApiUsagePlanSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApisStatusRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// api过滤条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`

	// tag key
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys" list`

	// tag values
	TagValues []*string `json:"TagValues,omitempty" name:"TagValues" list`

	// tag标签
	Tags []*ApiTags `json:"Tags,omitempty" name:"Tags" list`
}

func (r *DescribeApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApisStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *ApisStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApisStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名状态
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveSetRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeExclusiveSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExclusiveSetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveSetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *ExclusiveSetList `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExclusiveSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExclusiveSetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyApisStatusRequest struct {
	*tchttp.BaseRequest

	// 服务唯一id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 策略唯一id
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 环境名
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`
}

func (r *DescribeIPStrategyApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategyApisStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境绑定api详情
		Result *IPStrategyApiStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategyApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategyApisStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// ip策略id
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 环境
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`
}

func (r *DescribeIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略详情
		Result *IPStrategy `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategysStatusRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 过滤条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`
}

func (r *DescribeIPStrategysStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategysStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategysStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略列表
		Result *IPStrategysStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategysStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategysStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogRulesRequest struct {
	*tchttp.BaseRequest

	// 日志规则唯一ID数组。
	LogRuleIds []*string `json:"LogRuleIds,omitempty" name:"LogRuleIds" list`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 日志规则名称，支持模糊匹配。
	LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`

	// 环境名称。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
}

func (r *DescribeLogRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志规则列表。
		LogRuleSet []*LogRule `json:"LogRuleSet,omitempty" name:"LogRuleSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchRequest struct {
	*tchttp.BaseRequest

	// 日志开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 精确查询，支持apiid/reqid搜索
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 根据上次返回的ConText，获取后续的内容，最多可获取10000条
	ConText *string `json:"ConText,omitempty" name:"ConText"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 模糊查询，根据关键字检索日志
	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeLogSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSearchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取更多检索结果的游标，值为""表示无后续结果
		ConText *string `json:"ConText,omitempty" name:"ConText"`

		// 由0或多条日志组成，每条日志格式如下：
	// '[$app_id][$env_name][$service_id][$http_host][$api_id][$uri][$scheme][rsp_st:$status][ups_st:$upstream_status]'
	// '[cip:$remote_addr][uip:$upstream_addr][vip:$server_addr][rsp_len:$bytes_sent][req_len:$request_length]'
	// '[req_t:$request_time][ups_rsp_t:$upstream_response_time][ups_conn_t:$upstream_connect_time][ups_head_t:$upstream_header_time]’
	// '[err_msg:$err_msg][tcp_rtt:$tcpinfo_rtt][$pid][$time_local][req_id:$request_id]';
	// 
	// 说明：
	// app_id： 用户 ID。
	// env_name：环境名称。
	// service_id： 服务 ID。
	// http_host： 域名。
	// api_id： API 的 ID。
	// uri：请求的路径。
	// scheme： HTTP/HTTPS 协议。
	// rsp_st： 请求响应状态码。
	// ups_st： 后端业务服务器的响应状态码（如果请求透传到后端，改变量不为空。如果请求在 APIGW 就被拦截了，那么该变量显示为 -）。
	// cip： 客户端 IP。
	// uip： 后端业务服务（upstream）的 IP。
	// vip： 请求访问的 VIP。
	// rsp_len： 响应长度。
	// req_len： 请求长度。
	// req_t： 请求响应的总时间。
	// ups_rsp_t： 后端响应的总时间（apigw 建立连接到接收到后端响应的时间）。
	// ups_conn_t: 与后端业务服务器连接建立成功时间。
	// ups_head_t：后端响应的头部到达时间。
	// err_msg： 错误信息。
	// tcp_rtt： 客户端 TCP 连接信息，RTT（Round Trip Time）由三部分组成：链路的传播时间（propagation delay）、末端系统的处理时间、路由器缓存中的排队和处理时间（queuing delay）。
	// req_id：请求id。
		LogSet []*string `json:"LogSet,omitempty" name:"LogSet" list`

		// 单次搜索返回的日志条数，TotalCount <= Limit
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSearchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentKeyMonitorUploadRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DescribeServiceEnvironmentKeyMonitorUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentKeyMonitorUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentKeyMonitorUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *EnvironmentList `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentKeyMonitorUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentKeyMonitorUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentListRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务绑定环境详情
		Result *ServiceEnvironmentSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentReleaseHistoryRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentReleaseHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务发布历史
		Result *ServiceReleaseHistory `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境列表
		Result *ServiceEnvironmentStrategyStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceReleaseVersionRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceReleaseVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceReleaseVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceReleaseVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务发布版本列表
		Result *ServiceReleaseVersion `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceReleaseVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceReleaseVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
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

		// 服务唯一ID。
		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

		// 服务 环境列表。
		AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments" list`

		// 服务名称。
		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

		// 服务描述。
		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

		// 服务支持协议，可选值为http、https、http&https。
		Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

		// 服务创建时间。
		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

		// 服务修改时间。
		ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

		// 独立集群名称。
		ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

		// 网络类型列表，INNER为内网访问，OUTER为外网访问。
		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

		// 内网访问子域名。
		InternalSubDomain *string `json:"InternalSubDomain,omitempty" name:"InternalSubDomain"`

		// 外网访问子域名。
		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

		// 内网访问http服务端口号。
		InnerHttpPort *int64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

		// 内网访问https端口号。
		InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

		// API总数。
		ApiTotalCount *int64 `json:"ApiTotalCount,omitempty" name:"ApiTotalCount"`

		// API列表。
		ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet" list`

		// 使用计划总数量。
		UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`

		// 使用计划数组。
		UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList" list`

		// IP版本
		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

		// 此服务的用户类型
		UserType *string `json:"UserType,omitempty" name:"UserType"`

		// 服务所在集群
		SetId *int64 `json:"SetId,omitempty" name:"SetId"`

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

type DescribeServiceSubDomainMappingsRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 自定义域名
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *DescribeServiceSubDomainMappingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainMappingsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainMappingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *ServiceSubDomainMappings `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceSubDomainMappingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainMappingsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainsRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceSubDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *DomainSets `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceSubDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 服务唯一id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务绑定使用计划列表
		Result *ServiceUsagePlanSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesStatusRequest struct {
	*tchttp.BaseRequest

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`
}

func (r *DescribeServicesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServicesStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务列表
		Result *ServicesStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServicesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServicesStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTcbScfApisRequest struct {
	*tchttp.BaseRequest

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// SCF方法名称。
	ScfFunctionName *string `json:"ScfFunctionName,omitempty" name:"ScfFunctionName"`

	// SCF方法命名空间。
	ScfFunctionNamespace *string `json:"ScfFunctionNamespace,omitempty" name:"ScfFunctionNamespace"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTcbScfApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTcbScfApisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTcbScfApisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足条件查询结果总数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// TCB-SCF-HTTP触发器API列表。
		TcbScfApiSet []*TcbScfApi `json:"TcbScfApiSet,omitempty" name:"TcbScfApiSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTcbScfApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTcbScfApisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 使用计划id
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 绑定类型
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanEnvironmentsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划绑定详情
		Result *UsagePlanEnvironmentStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanEnvironmentsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划id
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DescribeUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划详情
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 使用计划id
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划绑定密钥详情
		Result *UsagePlanBindSecretStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlansStatusRequest struct {
	*tchttp.BaseRequest

	// 使用计划过滤条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlansStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlansStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlansStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划详情
		Result *UsagePlansStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlansStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlansStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥id
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DisableApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DocumentSDK struct {

	// 文档地址
	DocumentURL *string `json:"DocumentURL,omitempty" name:"DocumentURL"`

	// sdk地址
	SdkURL *string `json:"SdkURL,omitempty" name:"SdkURL"`
}

type DomainSetList struct {

	// 域名
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 解析状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 证书Id
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 是否使用默认路径映射
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

type DomainSets struct {

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 自定义服务域名列表
	DomainSet []*DomainSetList `json:"DomainSet,omitempty" name:"DomainSet" list`
}

type EnableApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥id
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *EnableApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Environment struct {

	// 环境名
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 环境对应的url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 是否发布
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 版本号
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type EnvironmentList struct {

	// 总数
	TotolCount *int64 `json:"TotolCount,omitempty" name:"TotolCount"`

	// 服务的环境列表
	EnvironmentSet []*EnvironmentUpload `json:"EnvironmentSet,omitempty" name:"EnvironmentSet" list`
}

type EnvironmentStrategy struct {

	// 环境名
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 限流值
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`
}

type EnvironmentUpload struct {

	// 环境
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 是否上报
	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type ErrorCodes struct {

	// code
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Msg
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Desc
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// ConvertedCode
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`

	// NeedConvert
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type ExclusiveSet struct {

	// 集群名字
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// 集群状态
	ExclusiveStatus *string `json:"ExclusiveStatus,omitempty" name:"ExclusiveStatus"`
}

type ExclusiveSetList struct {

	// 返回总数
	TotolCount *int64 `json:"TotolCount,omitempty" name:"TotolCount"`

	// 独占集群列表
	ExclusiveSet []*ExclusiveSet `json:"ExclusiveSet,omitempty" name:"ExclusiveSet" list`
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type GenerateApiDocumentRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务环境
	GenEnvironment *string `json:"GenEnvironment,omitempty" name:"GenEnvironment"`

	// 待生成sdk的语言，目前只支持 "python", "javascript"
	GenLanguage *string `json:"GenLanguage,omitempty" name:"GenLanguage"`
}

func (r *GenerateApiDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateApiDocumentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateApiDocumentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api文档&sdk链接
		Result *DocumentSDK `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateApiDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateApiDocumentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAccountSettingsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAccountSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAccountSettingsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAccountSettingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据
		Result *AccountSettings `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAccountSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAccountSettingsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HealthCheckConf struct {

	// 是否健康检测
	IsHealthCheck *bool `json:"IsHealthCheck,omitempty" name:"IsHealthCheck"`

	// 阈值
	RequestVolumeThreshold *int64 `json:"RequestVolumeThreshold,omitempty" name:"RequestVolumeThreshold"`

	// 窗口大小
	SleepWindowInMilliseconds *int64 `json:"SleepWindowInMilliseconds,omitempty" name:"SleepWindowInMilliseconds"`

	// 阈值百分比
	ErrorThresholdPercentage *int64 `json:"ErrorThresholdPercentage,omitempty" name:"ErrorThresholdPercentage"`
}

type IPStrategy struct {

	// 策略id
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// 策略详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 服务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 绑定api总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindApiTotalCount *int64 `json:"BindApiTotalCount,omitempty" name:"BindApiTotalCount"`

	// 绑定api详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindApis []*DesApisStatus `json:"BindApis,omitempty" name:"BindApis" list`
}

type IPStrategyApi struct {

	// api唯一id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api名称
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// api类型
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 策略id
	OtherIPStrategyId *string `json:"OtherIPStrategyId,omitempty" name:"OtherIPStrategyId"`

	// 环境名
	OtherEnvironmentName *string `json:"OtherEnvironmentName,omitempty" name:"OtherEnvironmentName"`
}

type IPStrategyApiStatus struct {

	// 环境绑定api数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 环境绑定api详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiIdStatusSet []*IPStrategyApi `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet" list`
}

type IPStrategysStatus struct {

	// 策略数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategySet []*IPStrategy `json:"StrategySet,omitempty" name:"StrategySet" list`
}

type ImportOpenApiRequest struct {
	*tchttp.BaseRequest

	// API所在的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// openAPI正文内容。
	Content *string `json:"Content,omitempty" name:"Content"`

	// Content格式，只能是YAML或者JSON，默认是YAML，目前只支持YAML。
	EncodeType *string `json:"EncodeType,omitempty" name:"EncodeType"`

	// Content版本，只能是openAPI或者swagger，默认是openAPI，目前只支持openAPI。
	ContentVersion *string `json:"ContentVersion,omitempty" name:"ContentVersion"`
}

func (r *ImportOpenApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportOpenApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportOpenApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportOpenApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportOpenApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRequest struct {
	*tchttp.BaseRequest

	// 计费模式。其中0：按小时计费。
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *InquiryPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 询价数据列表
		Price *Price `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Limit struct {

	// 日志规则配额
	LogRuleClount *AccountLimitLogRuleClount `json:"LogRuleClount,omitempty" name:"LogRuleClount"`

	// 服务配额
	ServiceCount *AccountLimitServiceCount `json:"ServiceCount,omitempty" name:"ServiceCount"`

	// 使用计划配额
	UsagePlanCount *AccountLimitUsagePlanCount `json:"UsagePlanCount,omitempty" name:"UsagePlanCount"`

	// 密钥配额
	SecretIdCount *AccountLimitSecretIdCount `json:"SecretIdCount,omitempty" name:"SecretIdCount"`
}

type LogRule struct {

	// 日志规则唯一ID。
	LogRuleId *string `json:"LogRuleId,omitempty" name:"LogRuleId"`

	// 服务唯一ID，ALL表示全部服务。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称，ALL表示全部环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志规则名称
	LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 服务名称。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type MicroService struct {

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名
	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type MicroServiceReq struct {

	// tsf集群
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// tsf命名空间
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名
	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type ModifyApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 限速值
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 环境名
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// api列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *ModifyApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementRequest struct {
	*tchttp.BaseRequest

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 接口ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 需要修改的API auth类型(可选择OAUTH-授权API)
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// oauth接口需要修改的公钥值
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// oauth接口重定向地址
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

func (r *ModifyApiIncrementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiIncrementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiIncrementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiIncrementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 后端服务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 请求配置
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// api id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api名称
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// api描述
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// api类型，默认为NORMAL
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// 鉴权类型
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 是否需要鉴权
	AuthRequired *bool `json:"AuthRequired,omitempty" name:"AuthRequired"`

	// 后端超时时间
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 是否开启跨域
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// 常量参数
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// 请求参数
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// api类型
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// mock类型返回
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// api绑定tst服务信息
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices" list`

	// tsf负载均衡配置
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// tsf健康检查配置
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// target负载均衡配置
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// target健康检查配置
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// scf 函数名称
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// scf websocket注册函数
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// scf websocket清理函数
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// scf websocket传输函数
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// scf 函数命名空间
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// scf函数版本
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// scf websocket注册函数命名空间
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// scf websocket传输函数版本
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// scf websocket传输函数命名空间
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// scf websocket传输函数版本
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// scf websocket清理函数命名空间
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// scf websocket清理函数版本
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// 是否开启响应集成
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// 开始调试后计费
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 标签
	TagSpecifications *Tag `json:"TagSpecifications,omitempty" name:"TagSpecifications"`

	// 是否删除返回错误码
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// 返回类型
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// 成功返回示例
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// 失败返回示例
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// 后端服务配置
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// 关联的授权api id
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// 服务参数
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// oauth配置
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 错误码配置
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`
}

func (r *ModifyApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApisAuthRelationApiRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`

	// 关联授权api
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`
}

func (r *ModifyApisAuthRelationApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApisAuthRelationApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApisAuthRelationApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果，成功返回true
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApisAuthRelationApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApisAuthRelationApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 策略id
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 策略数据
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *ModifyIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLogRuleRequest struct {
	*tchttp.BaseRequest

	// 日志规则唯一ID。
	LogRuleId *string `json:"LogRuleId,omitempty" name:"LogRuleId"`

	// 服务唯一ID，表示所需上报日志的服务ID，ALL代表所有服务。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称，表示所需要上报日志的环境名称，ALL代表所有环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 日志集ID。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志规则名称。
	LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`
}

func (r *ModifyLogRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLogRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLogRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLogRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentKeyMonitorUploadRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 需要修改的服务环境列表
	EnvironmentList []*EnvironmentUpload `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

func (r *ModifyServiceEnvironmentKeyMonitorUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceEnvironmentKeyMonitorUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentKeyMonitorUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceEnvironmentKeyMonitorUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceEnvironmentKeyMonitorUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 限速值
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 环境列表
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames" list`
}

func (r *ModifyServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务名称。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 服务描述。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 服务支持协议，可选值为http、https、http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络类型列表，用于指定支持的访问类型，INNER为内网访问，OUTER为外网访问。默认为OUTER。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`
}

func (r *ModifyServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 自定义域名
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 是否使用默认路径映射（'TRUE' 或 'FALSE'）
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 证书
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 网络协议（http，https 或 http&https)
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 路径映射列表
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`

	// 网络类型 （'INNER' 或 'OUTER'）
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

func (r *ModifySubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作成功返回true
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划id
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 使用计划名称
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划描述
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 最大调用次数
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 最大qps
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *ModifyUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划详情
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonitorTopsRequest struct {
	*tchttp.BaseRequest

	// 环境标识
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *MonitorTopsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MonitorTopsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonitorTopsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// top 10 访问数的服务列表。
		Top10ReqService []*ServiceReqCount `json:"Top10ReqService,omitempty" name:"Top10ReqService" list`

		// top 10 访问数的api列表。
		Top10ReqApi []*ApiReqCount `json:"Top10ReqApi,omitempty" name:"Top10ReqApi" list`

		// top 10 响应时间的api列表。
		Top10ResponseTimeApi []*ApiResponseTime `json:"Top10ResponseTimeApi,omitempty" name:"Top10ResponseTimeApi" list`

		// top 10 错误率的api列表。
		Top10ErrorRateApi []*ApiErrorRate `json:"Top10ErrorRateApi,omitempty" name:"Top10ErrorRateApi" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MonitorTopsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MonitorTopsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonitorTrendRequest struct {
	*tchttp.BaseRequest

	// 日期格式如"2018-12-28 09:50:04"。
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 日期格式如"2018-12-28 09:50:04"。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 曲线粒度（可选60，3600，86400）
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *MonitorTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MonitorTrendRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonitorTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源利用率趋势图。
		ResourceUsageTrend []*float64 `json:"ResourceUsageTrend,omitempty" name:"ResourceUsageTrend" list`

		// 用户的独占集群数量。
		ExclusiveLdCount *uint64 `json:"ExclusiveLdCount,omitempty" name:"ExclusiveLdCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MonitorTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MonitorTrendResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OauthConfig struct {

	// 公钥，用于验证用户token。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// token传递位置。
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`

	// 重定向地址，用于引导用户登录操作。
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

type PathMapping struct {

	// 路径。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 发布环境，可选值为“test”、 ”prepub“、”release“。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type Price struct {

	// 请求数计费询价
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestPrice *PriceInfos `json:"RequestPrice,omitempty" name:"RequestPrice"`

	// 流量计费询价
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthPrice *PriceInfos `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type PriceInfos struct {

	// 计费单位价格（元）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *string `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 计费单位周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
}

type ReleaseService struct {

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
}

type ReleaseServiceRequest struct {
	*tchttp.BaseRequest

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 发布环境
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// apiId列表，默认全量api发布
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *ReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发布信息
		Result *ReleaseService `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReqParameter struct {

	// 参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数位置
	Position *string `json:"Position,omitempty" name:"Position"`

	// 参数类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 是否必须
	Required *bool `json:"Required,omitempty" name:"Required"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type RequestConfig struct {

	// 后端path
	Path *string `json:"Path,omitempty" name:"Path"`

	// 后端method
	Method *string `json:"Method,omitempty" name:"Method"`
}

type RequestParameter struct {

	// 请求参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 参数位置
	Position *string `json:"Position,omitempty" name:"Position"`

	// 参数类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 是否必须
	Required *bool `json:"Required,omitempty" name:"Required"`
}

type ResponseErrorCodeReq struct {

	// 错误码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 错误码描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 错误码转换
	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`

	// 是否需要转换
	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type RunApiForMarketRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 接口id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 请求的header
	RequestHeader *string `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// 请求的query
	RequestQuery *string `json:"RequestQuery,omitempty" name:"RequestQuery"`

	// 请求的path
	RequestPath *string `json:"RequestPath,omitempty" name:"RequestPath"`

	// 请求的method
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// 请求的body
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`

	// 请求的body（字典类型）
	RequestBodyDict *string `json:"RequestBodyDict,omitempty" name:"RequestBodyDict"`

	// 可选值（'application/json'， 'application/x-www-form-urlencoded'）
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 私钥id
	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *RunApiForMarketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunApiForMarketRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunApiForMarketResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的header。
		ReturnHeader *string `json:"ReturnHeader,omitempty" name:"ReturnHeader"`

		// 返回的Body。
		ReturnBody *string `json:"ReturnBody,omitempty" name:"ReturnBody"`

		// 返回码。
		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 延迟。
		Delay *int64 `json:"Delay,omitempty" name:"Delay"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunApiForMarketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunApiForMarketResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunApiRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// apiid
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 请求头
	RequestHeader *string `json:"RequestHeader,omitempty" name:"RequestHeader"`

	// 请求参数
	RequestQuery *string `json:"RequestQuery,omitempty" name:"RequestQuery"`

	// 路径参数
	RequestPath *string `json:"RequestPath,omitempty" name:"RequestPath"`

	// 请求方法
	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`

	// 请求body
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`

	// 编码类型
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 请求body
	RequestBodyDict *string `json:"RequestBodyDict,omitempty" name:"RequestBodyDict"`
}

func (r *RunApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *RunApiReturn `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunApiReturn struct {

	// 响应头部
	ReturnHeader *string `json:"ReturnHeader,omitempty" name:"ReturnHeader"`

	// 响应body
	ReturnBody *string `json:"ReturnBody,omitempty" name:"ReturnBody"`

	// 返回码
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 延时
	Delay *int64 `json:"Delay,omitempty" name:"Delay"`
}

type Service struct {

	// https端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

	// service描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 服务修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

	// 独占集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// 服务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// ip version
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// 环境列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments" list`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 外网域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 内网http端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpPort *uint64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

	// 内网域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`

	// 计费状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeIsolateStatus *int64 `json:"TradeIsolateStatus,omitempty" name:"TradeIsolateStatus"`
}

type ServiceConfig struct {

	// 后端类型
	Product *string `json:"Product,omitempty" name:"Product"`

	// vpc id
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 方法
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ServiceEnvironmentSet struct {

	// 环境总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 环境详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentList []*Environment `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

type ServiceEnvironmentStrategy struct {

	// 环境名
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 限速值
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
}

type ServiceEnvironmentStrategyStatus struct {

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 环境列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentList []*ServiceEnvironmentStrategy `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

type ServiceParameter struct {

	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *string `json:"Position,omitempty" name:"Position"`

	// 相关请求参数位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantRequestParameterPosition *string `json:"RelevantRequestParameterPosition,omitempty" name:"RelevantRequestParameterPosition"`

	// 相关请求参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantRequestParameterName *string `json:"RelevantRequestParameterName,omitempty" name:"RelevantRequestParameterName"`

	// 默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantRequestParameterDesc *string `json:"RelevantRequestParameterDesc,omitempty" name:"RelevantRequestParameterDesc"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantRequestParameterType *string `json:"RelevantRequestParameterType,omitempty" name:"RelevantRequestParameterType"`
}

type ServiceReleaseHistory struct {

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 发布历史
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList" list`
}

type ServiceReleaseHistoryInfo struct {

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 版本描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
}

type ServiceReleaseVersion struct {

	// 版本数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList" list`
}

type ServiceReqCount struct {

	// 用户服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务名字
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 请求数
	ReqCount *uint64 `json:"ReqCount,omitempty" name:"ReqCount"`
}

type ServiceSubDomainMappings struct {

	// 是否使用默认路径映射
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 路径映射列表
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`
}

type ServiceUsagePlanSet struct {

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceUsagePlanList []*ApiUsagePlan `json:"ServiceUsagePlanList,omitempty" name:"ServiceUsagePlanList" list`
}

type ServicesStatus struct {

	// 服务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceSet []*Service `json:"ServiceSet,omitempty" name:"ServiceSet" list`
}

type Tag struct {

	// tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TargetServicesReq struct {

	// vm ip
	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`

	// vpc id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vm port
	VmPort *int64 `json:"VmPort,omitempty" name:"VmPort"`

	// cvm所在宿主机ip
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// docker ip
	DockerIp *string `json:"DockerIp,omitempty" name:"DockerIp"`
}

type TcbScfApi struct {

	// 环境ID。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// SCF方法名称。
	ScfFunctionName *string `json:"ScfFunctionName,omitempty" name:"ScfFunctionName"`

	// SCF方法命名空间。
	ScfFunctionNamespace *string `json:"ScfFunctionNamespace,omitempty" name:"ScfFunctionNamespace"`

	// 自定义子域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 请求Path。
	Path *string `json:"Path,omitempty" name:"Path"`
}

type TsfLoadBalanceConfResp struct {

	// 是否开启负载均衡
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLoadBalance *bool `json:"IsLoadBalance,omitempty" name:"IsLoadBalance"`

	// 方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 是否会话保持
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionStickRequired *bool `json:"SessionStickRequired,omitempty" name:"SessionStickRequired"`

	// 会话保持超时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionStickTimeout *int64 `json:"SessionStickTimeout,omitempty" name:"SessionStickTimeout"`
}

type UnBindEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 绑定类型
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 使用计划id列表
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds" list`

	// 环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API id列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *UnBindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindEnvironmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindEnvironmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// ip策略
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 环境
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 解绑的api id列表
	UnBindApiIds []*string `json:"UnBindApiIds,omitempty" name:"UnBindApiIds" list`
}

func (r *UnBindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 使用计划id
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 解绑密钥列表
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds" list`
}

func (r *UnBindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 需解绑的自定义域名
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *UnBindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作成功返回true
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnReleaseServiceRequest struct {
	*tchttp.BaseRequest

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 需要下线的环境
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// apiId组，默认全量
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *UnReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnReleaseServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnReleaseServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥id
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 密钥key，自定义密钥更新必传
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *UpdateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥详情
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 版本Id
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 本次更新描述
	UpdateDesc *string `json:"UpdateDesc,omitempty" name:"UpdateDesc"`
}

func (r *UpdateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Usage struct {

	// 服务使用量
	ServiceCount *AccountUsageServiceCount `json:"ServiceCount,omitempty" name:"ServiceCount"`

	// 使用计划使用量
	UsagePlanCount *AccountUsageUsagePlanCount `json:"UsagePlanCount,omitempty" name:"UsagePlanCount"`

	// 密钥使用量
	SecretIdCount *AccountUsageSecretIdCount `json:"SecretIdCount,omitempty" name:"SecretIdCount"`

	// 日志使用量
	LogRuleClount *AccountUsageLogRuleClount `json:"LogRuleClount,omitempty" name:"LogRuleClount"`
}

type UsagePlan struct {

	// 环境名称。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 使用计划唯一ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 使用计划名称。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划描述。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 使用计划qps，-1表示没有限制。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 使用计划时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 使用计划修改时间。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type UsagePlanBindEnvironment struct {

	// 环境名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 服务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type UsagePlanBindSecret struct {

	// 密钥id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 密钥名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 密钥状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type UsagePlanBindSecretStatus struct {

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 密钥详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKeyList []*UsagePlanBindSecret `json:"AccessKeyList,omitempty" name:"AccessKeyList" list`
}

type UsagePlanEnvironment struct {

	// 服务唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// api名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// api路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// api方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 环境名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 已使用配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// 最大请求量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 最大qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type UsagePlanEnvironmentStatus struct {

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 使用计划绑定详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentList []*UsagePlanEnvironment `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

type UsagePlanInfo struct {

	// 使用计划id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 使用计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 初始化调用次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitQuota *int64 `json:"InitQuota,omitempty" name:"InitQuota"`

	// 最大qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 最大调用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 是否隐藏
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsHide *int64 `json:"IsHide,omitempty" name:"IsHide"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 绑定密钥数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindSecretIdTotalCount *int64 `json:"BindSecretIdTotalCount,omitempty" name:"BindSecretIdTotalCount"`

	// 绑定密钥详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindSecretIds []*string `json:"BindSecretIds,omitempty" name:"BindSecretIds" list`

	// 绑定环境数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindEnvironmentTotalCount *int64 `json:"BindEnvironmentTotalCount,omitempty" name:"BindEnvironmentTotalCount"`

	// 绑定环境详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindEnvironments []*UsagePlanBindEnvironment `json:"BindEnvironments,omitempty" name:"BindEnvironments" list`
}

type UsagePlanStatusInfo struct {

	// 使用计划id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 使用计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 最大qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 最大调用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type UsagePlansStatus struct {

	// 使用计划数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 使用计划详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanStatusSet []*UsagePlanStatusInfo `json:"UsagePlanStatusSet,omitempty" name:"UsagePlanStatusSet" list`
}
