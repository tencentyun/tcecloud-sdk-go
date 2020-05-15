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

package v20190924

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type AutoDelStrategyInfo struct {

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 仓库名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略值
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// Valid
	Valid *int64 `json:"Valid,omitempty" name:"Valid"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type AutoDelStrategyInfoResp struct {

	// 总数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 自动删除策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyInfo []*AutoDelStrategyInfo `json:"StrategyInfo,omitempty" name:"StrategyInfo" list`
}

type BatchDeleteImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag列表
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`
}

func (r *BatchDeleteImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDeleteImagePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDeleteImagePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称数组
	RepoNames []*string `json:"RepoNames,omitempty" name:"RepoNames" list`
}

func (r *BatchDeleteRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDeleteRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDeleteRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest

	// 触发器关联的镜像仓库，library/test格式
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发方式，"all"全部触发，"taglist"指定tag触发，"regex"正则触发
	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`

	// 应用所在TKE集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用所在TKE集群命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 应用所在TKE集群工作负载类型,支持Deployment、StatefulSet、DaemonSet、CronJob、Job。
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 应用所在TKE集群工作负载名称
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`

	// 应用所在TKE集群工作负载下容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 应用所在TKE集群地域
	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`

	// 触发方式对应的表达式
	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`
}

func (r *CreateApplicationTriggerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApplicationTriggerPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationTriggerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApplicationTriggerPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// keep_last_days:保留最近几天的数据;keep_last_nums:保留最近多少个
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略值
	Val *int64 `json:"Val,omitempty" name:"Val"`
}

func (r *CreateImageLifecyclePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageLifecyclePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageLifecyclePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageLifecyclePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 企业版实例名称
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`

	// 企业版实例类型
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 企业版实例Id
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceTokenRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *CreateInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 临时密码
		Token *string `json:"Token,omitempty" name:"Token"`

		// 临时密码有效期时间戳
		ExpTime *int64 `json:"ExpTime,omitempty" name:"ExpTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespacePersonalRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *CreateNamespacePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNamespacePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespacePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNamespacePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNamespacePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 是否公共,1:公共,0:私有
	Public *uint64 `json:"Public,omitempty" name:"Public"`

	// 仓库描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUserPersonalRequest struct {
	*tchttp.BaseRequest

	// 用户密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CreateUserPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUserPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUserPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUserPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
}

func (r *DeleteApplicationTriggerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApplicationTriggerPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationTriggerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApplicationTriggerPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteImageLifecycleGlobalPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageLifecycleGlobalPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteImageLifecyclePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageLifecyclePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageLifecyclePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageLifecyclePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DeleteImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImagePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImagePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespacePersonalRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteNamespacePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNamespacePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespacePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNamespacePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNamespacePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTriggerLogPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 升序或降序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 按某列排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeApplicationTriggerLogPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationTriggerLogPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTriggerLogPersonalResp struct {

	// 返回总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 触发日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogInfo []*TriggerLogResp `json:"LogInfo,omitempty" name:"LogInfo" list`
}

type DescribeApplicationTriggerLogPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 触发日志返回值
		Data *DescribeApplicationTriggerLogPersonalResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationTriggerLogPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationTriggerLogPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeApplicationTriggerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationTriggerPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTriggerPersonalResp struct {

	// 返回条目总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 触发器列表
	TriggerInfo []*TriggerResp `json:"TriggerInfo,omitempty" name:"TriggerInfo" list`
}

type DescribeApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 触发器列表返回值
		Data *DescribeApplicationTriggerPersonalResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationTriggerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationTriggerPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset用于分页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFavorRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFavorRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 个人收藏仓库列表返回信息
		Data *FavorResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFavorRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFavorRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageFilterPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeImageFilterPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageFilterPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageFilterPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// payload
		Data *SameImagesResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageFilterPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageFilterPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageLifecycleGlobalPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 全局自动删除策略信息
		Data *AutoDelStrategyInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageLifecycleGlobalPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeImageLifecyclePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageLifecyclePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自动删除策略信息
		Data *AutoDelStrategyInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageLifecyclePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageLifecyclePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// tag名称，可根据输入搜索
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像tag信息
		Data *TagInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatusRequest struct {
	*tchttp.BaseRequest

	// 实例ID的数组
	RegistryIds []*string `json:"RegistryIds,omitempty" name:"RegistryIds" list`
}

func (r *DescribeInstanceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例的状态列表
		RegistryStatusSet []*RegistryStatus `json:"RegistryStatusSet,omitempty" name:"RegistryStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表(为空时，
	// 表示获取账号下所有实例)
	Registryids []*string `json:"Registryids,omitempty" name:"Registryids" list`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 获取所有地域的实例，默认为False
	AllRegion *bool `json:"AllRegion,omitempty" name:"AllRegion"`
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

		// 总实例个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例信息列表
		Registries []*Registry `json:"Registries,omitempty" name:"Registries" list`

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

type DescribeNamespacePersonalRequest struct {
	*tchttp.BaseRequest

	// 命名空间，支持模糊查询
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNamespacePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespacePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户命名空间返回信息
		Data *NamespaceInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespacePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNamespacePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryFilterPersonalRequest struct {
	*tchttp.BaseRequest

	// 搜索镜像名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20，最大100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件：1表示public，0表示private
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeRepositoryFilterPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryFilterPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryFilterPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 仓库信息
		Data *SearchUserRepositoryResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryFilterPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryFilterPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryOwnerPersonalRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeRepositoryOwnerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryOwnerPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryOwnerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 仓库信息
		Data *RepoInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryOwnerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryOwnerPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名字
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 仓库信息
		Data *RepositoryInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotaPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserQuotaPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserQuotaPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotaPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额返回信息
		Data *RespLimit `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserQuotaPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserQuotaPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DupImageTagResp struct {

	// 镜像Digest值
	Digest *string `json:"Digest,omitempty" name:"Digest"`
}

type DuplicateImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 源镜像名称，不包含domain。例如： tencentyun/foo:v1
	SrcImage *string `json:"SrcImage,omitempty" name:"SrcImage"`

	// 目的镜像名称，不包含domain。例如： tencentyun/foo:latest
	DestImage *string `json:"DestImage,omitempty" name:"DestImage"`
}

func (r *DuplicateImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DuplicateImagePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DuplicateImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 复制镜像返回值
		Data *DupImageTagResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DuplicateImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DuplicateImagePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FavorResp struct {

	// 收藏仓库的总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoInfo []*Favors `json:"RepoInfo,omitempty" name:"RepoInfo" list`
}

type Favors struct {

	// 仓库名字
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// Pull总共的次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

	// 仓库收藏次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 仓库是否公开
	// 注意：此字段可能返回 null，表示取不到有效值。
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 是否为官方所有
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`

	// 仓库Tag的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

	// Logo
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitempty" name:"Logo"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域的Id
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type Filter struct {

	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Limit struct {

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 配额的类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配置的值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type ManageImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest

	// global_keep_last_days:全局保留最近几天的数据;global_keep_last_nums:全局保留最近多少个
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略值
	Val *int64 `json:"Val,omitempty" name:"Val"`
}

func (r *ManageImageLifecycleGlobalPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManageImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageImageLifecycleGlobalPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest

	// 触发器关联的镜像仓库，library/test格式
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发方式，"all"全部触发，"taglist"指定tag触发，"regex"正则触发
	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`

	// 触发方式对应的表达式
	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`

	// 应用所在TKE集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用所在TKE集群命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 应用所在TKE集群工作负载类型,支持Deployment、StatefulSet、DaemonSet、CronJob、Job。
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 应用所在TKE集群工作负载名称
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`

	// 应用所在TKE集群工作负载下容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 应用所在TKE集群地域数字ID，如1（广州）、16（成都）
	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`

	// 新触发器名称
	NewTriggerName *string `json:"NewTriggerName,omitempty" name:"NewTriggerName"`
}

func (r *ModifyApplicationTriggerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApplicationTriggerPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApplicationTriggerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApplicationTriggerPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRepositoryAccessPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 默认值为0
	Public *int64 `json:"Public,omitempty" name:"Public"`
}

func (r *ModifyRepositoryAccessPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRepositoryAccessPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRepositoryAccessPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRepositoryAccessPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRepositoryAccessPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRepositoryInfoPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyRepositoryInfoPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRepositoryInfoPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRepositoryInfoPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRepositoryInfoPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRepositoryInfoPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUserPasswordPersonalRequest struct {
	*tchttp.BaseRequest

	// 更新后的密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModifyUserPasswordPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUserPasswordPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUserPasswordPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserPasswordPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUserPasswordPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NamespaceInfo struct {

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 命名空间下仓库数量
	RepoCount *int64 `json:"RepoCount,omitempty" name:"RepoCount"`
}

type NamespaceInfoResp struct {

	// 命名空间数量
	NamespaceCount *int64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`

	// 命名空间信息
	NamespaceInfo []*NamespaceInfo `json:"NamespaceInfo,omitempty" name:"NamespaceInfo" list`
}

type NamespaceIsExistsResp struct {

	// 命名空间是否存在
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`

	// 是否为保留命名空间
	IsPreserved *bool `json:"IsPreserved,omitempty" name:"IsPreserved"`
}

type Registry struct {

	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 实例名称
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`

	// 实例规格
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 实例状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例的公共访问地址
	PublicDomain *string `json:"PublicDomain,omitempty" name:"PublicDomain"`

	// 实例创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域Id
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 是否支持匿名
	EnableAnonymous *bool `json:"EnableAnonymous,omitempty" name:"EnableAnonymous"`

	// Token有效时间
	TokenValidTime *uint64 `json:"TokenValidTime,omitempty" name:"TokenValidTime"`
}

type RegistryCondition struct {

	// 实例创建过程类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 实例创建过程状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 转换到该过程的简明原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type RegistryStatus struct {

	// 实例的Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 实例的状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 附加状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*RegistryCondition `json:"Conditions,omitempty" name:"Conditions" list`
}

type RepoInfo struct {

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// Tag数量
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

	// 是否为公开
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 是否为用户收藏
	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`

	// 是否为Tce官方仓库
	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`

	// 被收藏的个数
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 拉取的数量
	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 仓库创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 仓库更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type RepoInfoResp struct {

	// 仓库总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库信息列表
	RepoInfo []*RepoInfo `json:"RepoInfo,omitempty" name:"RepoInfo" list`

	// Server信息
	Server *string `json:"Server,omitempty" name:"Server"`
}

type RepoIsExistResp struct {

	// 仓库是否存在
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`
}

type RepositoryInfoResp struct {

	// 镜像仓库名字
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像仓库类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 镜像仓库服务地址
	Server *string `json:"Server,omitempty" name:"Server"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 镜像仓库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否为公有镜像
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 下载次数
	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

	// 收藏次数
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 是否为用户收藏
	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`

	// 是否为Tce官方镜像
	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`
}

type RespLimit struct {

	// 配额信息
	LimitInfo []*Limit `json:"LimitInfo,omitempty" name:"LimitInfo" list`
}

type SameImagesResp struct {

	// tag列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SameImages []*string `json:"SameImages,omitempty" name:"SameImages" list`
}

type SearchUserRepositoryResp struct {

	// 总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库列表
	RepoInfo []*RepoInfo `json:"RepoInfo,omitempty" name:"RepoInfo" list`

	// Server
	Server *string `json:"Server,omitempty" name:"Server"`

	// PrivilegeFiltered
	PrivilegeFiltered *bool `json:"PrivilegeFiltered,omitempty" name:"PrivilegeFiltered"`
}

type TagInfo struct {

	// Tag名称
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 镜像Id
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// docker image 可以看到的id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 大小
	Size *string `json:"Size,omitempty" name:"Size"`

	// 镜像的创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 镜像创建至今时间长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationDays *string `json:"DurationDays,omitempty" name:"DurationDays"`

	// 镜像的作者
	Author *string `json:"Author,omitempty" name:"Author"`

	// 次镜像建议运行的系统架构
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 创建此镜像的docker版本
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// 此镜像建议运行系统
	OS *string `json:"OS,omitempty" name:"OS"`

	// SizeByte
	SizeByte *int64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 数据更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 镜像更新时间
	PushTime *string `json:"PushTime,omitempty" name:"PushTime"`
}

type TagInfoResp struct {

	// Tag的总数
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

	// TagInfo列表
	TagInfo []*TagInfo `json:"TagInfo,omitempty" name:"TagInfo" list`

	// Server
	Server *string `json:"Server,omitempty" name:"Server"`

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type TriggerInvokeCondition struct {

	// 触发方式
	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`

	// 触发表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`
}

type TriggerInvokePara struct {

	// AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// TKE集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// TKE集群命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// TKE集群工作负载名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// TKE集群工作负载中容器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// TKE集群地域数字ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

type TriggerInvokeResult struct {

	// 请求TKE返回值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 请求TKE返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
}

type TriggerLogResp struct {

	// 仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 触发器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeSource *string `json:"InvokeSource,omitempty" name:"InvokeSource"`

	// 触发动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeAction *string `json:"InvokeAction,omitempty" name:"InvokeAction"`

	// 触发时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeTime *string `json:"InvokeTime,omitempty" name:"InvokeTime"`

	// 触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeCondition *TriggerInvokeCondition `json:"InvokeCondition,omitempty" name:"InvokeCondition"`

	// 触发参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokePara *TriggerInvokePara `json:"InvokePara,omitempty" name:"InvokePara"`

	// 触发结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeResult *TriggerInvokeResult `json:"InvokeResult,omitempty" name:"InvokeResult"`
}

type TriggerResp struct {

	// 触发器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeSource *string `json:"InvokeSource,omitempty" name:"InvokeSource"`

	// 触发动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeAction *string `json:"InvokeAction,omitempty" name:"InvokeAction"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeCondition *TriggerInvokeCondition `json:"InvokeCondition,omitempty" name:"InvokeCondition"`

	// 触发器参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokePara *TriggerInvokePara `json:"InvokePara,omitempty" name:"InvokePara"`
}

type ValidateNamespaceExistPersonalRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ValidateNamespaceExistPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateNamespaceExistPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateNamespaceExistPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间是否存在
		Data *NamespaceIsExistsResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateNamespaceExistPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateNamespaceExistPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateRepositoryExistPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *ValidateRepositoryExistPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateRepositoryExistPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateRepositoryExistPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 仓库是否存在
		Data *RepoIsExistResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateRepositoryExistPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateRepositoryExistPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
