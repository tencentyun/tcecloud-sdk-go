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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessVpc struct {

	// Vpc的Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 内网接入状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 内网接入Ip
	AccessIp *string `json:"AccessIp,omitempty" name:"AccessIp"`
}

type AuthUser struct {

	// 用户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

type AuthUserInfoResp struct {

	// 构建用户信息
	User *AuthUser `json:"User,omitempty" name:"User"`
}

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

type BatchDeleteFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 收藏仓库的列表
	Favors []*RequestFavor `json:"Favors,omitempty" name:"Favors" list`
}

func (r *BatchDeleteFavorRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDeleteFavorRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteFavorRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDeleteFavorRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type BuildBranchResp struct {

	// 构建的分支信息
	Branches []*string `json:"Branches,omitempty" name:"Branches" list`
}

type BuildHistoryResp struct {

	// 构建信息
	BuildHistory *BuildInfo `json:"BuildHistory,omitempty" name:"BuildHistory"`
}

type BuildInfo struct {

	// Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 构建类型
	BuildType *string `json:"BuildType,omitempty" name:"BuildType"`

	// 是否手动构建
	BuildManually *int64 `json:"BuildManually,omitempty" name:"BuildManually"`

	// 构建工作目录
	BuildWorkDir *string `json:"BuildWorkDir,omitempty" name:"BuildWorkDir"`

	// 构建参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Args *string `json:"Args,omitempty" name:"Args"`

	// 构建状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 构建开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 构建结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 构建仓库地址
	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`

	// repo所在的group
	Group *string `json:"Group,omitempty" name:"Group"`

	// 构建所在的repo
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// Repo url地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoUrl *string `json:"RepoUrl,omitempty" name:"RepoUrl"`

	// 用户在git服务器上的用户名
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 构建的分支
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// dockerfile在仓库中的路径
	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

	// registry中的namespace
	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`

	// 用户在registry中的用户名
	RegistryUsername *string `json:"RegistryUsername,omitempty" name:"RegistryUsername"`

	// 镜像名称，不包含tag
	Image *string `json:"Image,omitempty" name:"Image"`

	// 镜像名
	ForceImage *string `json:"ForceImage,omitempty" name:"ForceImage"`

	// 提交的SHA
	CommitSHA *string `json:"CommitSHA,omitempty" name:"CommitSHA"`

	// 提交的作者
	CommitAuthor *string `json:"CommitAuthor,omitempty" name:"CommitAuthor"`

	// 提交的信息
	CommitMessage *string `json:"CommitMessage,omitempty" name:"CommitMessage"`

	// 提交的时间
	CommitTime *string `json:"CommitTime,omitempty" name:"CommitTime"`

	// 构建日志
	BuildLog *string `json:"BuildLog,omitempty" name:"BuildLog"`
}

type BuildInfoResp struct {

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 构建信息列表
	BuildList []*BuildInfo `json:"BuildList,omitempty" name:"BuildList" list`
}

type BuildRepo struct {

	// 源码所在的git服务
	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`

	// repo所在的group
	Group *string `json:"Group,omitempty" name:"Group"`

	// 源码在git服务器上的仓库名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库Id
	RepoId *int64 `json:"RepoId,omitempty" name:"RepoId"`

	// 仓库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 是否为私有
	// 注意：此字段可能返回 null，表示取不到有效值。
	Private *bool `json:"Private,omitempty" name:"Private"`

	// WebUrl
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebUrl *string `json:"WebUrl,omitempty" name:"WebUrl"`
}

type BuildReposList struct {

	// 仓库信息列表
	Repos []*BuildRepo `json:"Repos,omitempty" name:"Repos" list`
}

type BuildRule struct {

	// registry中的namespace
	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`

	// 用户在registry中的用户名
	RegistryUsername *string `json:"RegistryUsername,omitempty" name:"RegistryUsername"`

	// 镜像名称，不包含tag
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像tag的格式
	ImageTagFormat *string `json:"ImageTagFormat,omitempty" name:"ImageTagFormat"`

	// 源码所在的git服务
	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`

	// repo所在的group
	Group *string `json:"Group,omitempty" name:"Group"`

	// 源码在git服务器上的仓库名
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// 用户在git服务器上的用户名
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 分支
	// 注意：此字段可能返回 null，表示取不到有效值。
	Branches []*string `json:"Branches,omitempty" name:"Branches" list`

	// Tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *int64 `json:"Tag,omitempty" name:"Tag"`

	// dockerfile在仓库中的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

	// 工作目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildWorkDir *string `json:"BuildWorkDir,omitempty" name:"BuildWorkDir"`

	// 构建出的镜像覆盖该Tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceTag *string `json:"ForceTag,omitempty" name:"ForceTag"`

	// Args
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildArgs *string `json:"BuildArgs,omitempty" name:"BuildArgs"`
}

type BuildRuleResp struct {

	// BuildRule信息
	BuildRule *BuildRule `json:"BuildRule,omitempty" name:"BuildRule"`
}

type CreateApplicationTokenPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateApplicationTokenPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApplicationTokenPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationTokenPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建第三方应用访问凭证执行结果
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationTokenPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApplicationTokenPersonalResponse) FromJsonString(s string) error {
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

type CreateFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库的名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库的类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
}

func (r *CreateFavorRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFavorRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFavorRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFavorRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildPersonalRequest struct {
	*tchttp.BaseRequest

	// registry中的namespace
	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`

	// 用户在registry中的用户名
	RegistryUsername *string `json:"RegistryUsername,omitempty" name:"RegistryUsername"`

	// 镜像名称，不包含tag
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像tag的格式
	ImageTagFormat *string `json:"ImageTagFormat,omitempty" name:"ImageTagFormat"`

	// 源码所在的git服务
	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`

	// repo所在的group
	Group *string `json:"Group,omitempty" name:"Group"`

	// 源码在git服务器上的仓库名
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// 用户在git服务器上的用户名
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// Trigger信息
	Trigger *Trigger `json:"Trigger,omitempty" name:"Trigger"`

	// dockerfile在仓库中的路径
	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

	// 工作目录
	WorkDir *string `json:"WorkDir,omitempty" name:"WorkDir"`

	// 构建出的镜像覆盖该Tag
	ForceTag *string `json:"ForceTag,omitempty" name:"ForceTag"`

	// Args
	Args []*string `json:"Args,omitempty" name:"Args" list`
}

func (r *CreateImageBuildPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageBuildPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageBuildPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageBuildPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildTaskDockerPersonalRequest struct {
	*tchttp.BaseRequest

	// dockerfile在仓库中的路径
	Dockerfile *string `json:"Dockerfile,omitempty" name:"Dockerfile"`

	// 用户在registry中的用户名
	RegistryUserName *string `json:"RegistryUserName,omitempty" name:"RegistryUserName"`

	// registry中的namespace
	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`

	// 镜像名称
	Image *string `json:"Image,omitempty" name:"Image"`
}

func (r *CreateImageBuildTaskDockerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageBuildTaskDockerPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildTaskDockerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 构建Id
		Data *int64 `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageBuildTaskDockerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageBuildTaskDockerPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildTaskManuallyPersonalRequest struct {
	*tchttp.BaseRequest

	// registry所在的namespace
	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// Tag的格式
	TagFormat *string `json:"TagFormat,omitempty" name:"TagFormat"`

	// 仓库类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 仓库分支或者是commit
	BranchOrCommit *string `json:"BranchOrCommit,omitempty" name:"BranchOrCommit"`
}

func (r *CreateImageBuildTaskManuallyPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageBuildTaskManuallyPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildTaskManuallyPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageBuildTaskManuallyPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageBuildTaskManuallyPersonalResponse) FromJsonString(s string) error {
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

type CreateRepoRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 是否公共,1:公共,0:私有
	Public *uint64 `json:"Public,omitempty" name:"Public"`

	// 仓库描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRepoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRepoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRepoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRepoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRepoResponse) FromJsonString(s string) error {
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

type CreateSecurityPoliciesRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSourceCodeAuthPersonalRequest struct {
	*tchttp.BaseRequest

	// GitServer地址
	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`

	// 用户信息
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// git 秘钥
	GitToken *string `json:"GitToken,omitempty" name:"GitToken"`
}

func (r *CreateSourceCodeAuthPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSourceCodeAuthPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSourceCodeAuthPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSourceCodeAuthPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSourceCodeAuthPersonalResponse) FromJsonString(s string) error {
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

type DeleteFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
}

func (r *DeleteFavorRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFavorRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFavorRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFavorRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFavorRepositoryRegionPersonalRequest struct {
	*tchttp.BaseRequest

	// 被收藏的仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteFavorRepositoryRegionPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFavorRepositoryRegionPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFavorRepositoryRegionPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFavorRepositoryRegionPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFavorRepositoryRegionPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageBuildPersonalRequest struct {
	*tchttp.BaseRequest

	// 镜像列表
	Images []*string `json:"Images,omitempty" name:"Images" list`
}

func (r *DeleteImageBuildPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageBuildPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageBuildPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageBuildPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageBuildPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageBuildTaskPersonalRequest struct {
	*tchttp.BaseRequest

	// 构建任务ID
	BuildId *int64 `json:"BuildId,omitempty" name:"BuildId"`
}

func (r *DeleteImageBuildTaskPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageBuildTaskPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageBuildTaskPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageBuildTaskPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageBuildTaskPersonalResponse) FromJsonString(s string) error {
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

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
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

type DeleteSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 白名单Id
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 白名单版本
	PolicyVersion *string `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

func (r *DeleteSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSourceCodeAuthPersonalRequest struct {
	*tchttp.BaseRequest

	// Server 地址（GitHub或者GitLab）
	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
}

func (r *DeleteSourceCodeAuthPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSourceCodeAuthPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSourceCodeAuthPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSourceCodeAuthPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSourceCodeAuthPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DesGitAuthRsp struct {

	// 用户GitAuth信息
	AuthMap *string `json:"AuthMap,omitempty" name:"AuthMap"`
}

type DescribeApplicationTokenPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApplicationTokenPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationTokenPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTokenPersonalResp struct {

	// 镜像仓库凭证
	ApplicationToken *string `json:"ApplicationToken,omitempty" name:"ApplicationToken"`
}

type DescribeApplicationTokenPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取镜像仓库个人版凭证
		Data *DescribeApplicationTokenPersonalResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationTokenPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationTokenPersonalResponse) FromJsonString(s string) error {
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

type DescribeDockerHubImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeDockerHubImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDockerHubImagePersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询dockerhub仓库镜像列表的返回值
		Data *DockerHubTagList `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDockerHubImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDockerHubImagePersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubRepositoryInfoPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeDockerHubRepositoryInfoPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDockerHubRepositoryInfoPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubRepositoryInfoPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// DockerHub仓库信息
		Data *DockerHubRepoinfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDockerHubRepositoryInfoPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDockerHubRepositoryInfoPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Limit用于分页
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量用于分页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDockerHubRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDockerHubRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// dockerhub仓库列表
		Data *RespDockerHubRepoList `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDockerHubRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDockerHubRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalEndpointStatusRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeExternalEndpointStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExternalEndpointStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalEndpointStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 开启公网访问状态，包括开启中，开启成功以及关闭和更新失败等
		Status *string `json:"Status,omitempty" name:"Status"`

		// 原因
		Reason *string `json:"Reason,omitempty" name:"Reason"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExternalEndpointStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExternalEndpointStatusResponse) FromJsonString(s string) error {
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

type DescribeImageBuildPersonalRequest struct {
	*tchttp.BaseRequest

	// registry中的namespace
	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`

	// Image名称
	Image *string `json:"Image,omitempty" name:"Image"`
}

func (r *DescribeImageBuildPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageBuildPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// BuildRule信息返回值
		Data *BuildRuleResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageBuildPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageBuildPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildTaskLogInfoPersonalRequest struct {
	*tchttp.BaseRequest

	// 构建日志ID
	BuildId *int64 `json:"BuildId,omitempty" name:"BuildId"`
}

func (r *DescribeImageBuildTaskLogInfoPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageBuildTaskLogInfoPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildTaskLogInfoPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 构建历史信息
		Data *BuildHistoryResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageBuildTaskLogInfoPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageBuildTaskLogInfoPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildTaskLogPersonalRequest struct {
	*tchttp.BaseRequest

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// registry在仓库中的namespace
	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`

	// Offset用于分页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit用于分页
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageBuildTaskLogPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageBuildTaskLogPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildTaskLogPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 构建日志
		Data *BuildInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageBuildTaskLogPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageBuildTaskLogPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageConfigPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名称
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeImageConfigPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageConfigPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageConfigPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// payload
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageConfigPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageConfigPersonalResponse) FromJsonString(s string) error {
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

type DescribeInternalEndpointsRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeInternalEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInternalEndpointsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内网接入信息的列表
		AccessVpcSet []*AccessVpc `json:"AccessVpcSet,omitempty" name:"AccessVpcSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternalEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInternalEndpointsResponse) FromJsonString(s string) error {
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

		// 返回的总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 地域信息列表
		Regions []*Region `json:"Regions,omitempty" name:"Regions" list`

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

type DescribeRepositoryAllPersonalRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 筛选条件，支持pullCount和official两个值
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序还是降序，默认是desc，还可以选择asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRepositoryAllPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryAllPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryAllPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像仓库列表
		Data *RepoInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryAllPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryAllPersonalResponse) FromJsonString(s string) error {
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

type DescribeSecurityPoliciesRequest struct {
	*tchttp.BaseRequest

	// 实例的Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeSecurityPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例安全策略组
		SecurityPolicySet []*SecurityPolicy `json:"SecurityPolicySet,omitempty" name:"SecurityPolicySet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeAuthPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSourceCodeAuthPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSourceCodeAuthPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeAuthPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询信息
		Data *DesGitAuthRsp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceCodeAuthPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSourceCodeAuthPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeAuthUserInfoPersonalRequest struct {
	*tchttp.BaseRequest

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSourceCodeAuthUserInfoPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSourceCodeAuthUserInfoPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeAuthUserInfoPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 源代码授权用户信息
		Data *AuthUserInfoResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceCodeAuthUserInfoPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSourceCodeAuthUserInfoPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeRepositoryBranchPersonalRequest struct {
	*tchttp.BaseRequest

	// 代码仓库类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// repo所在的group
	Group *string `json:"Group,omitempty" name:"Group"`

	// 源码在git服务器上的仓库名
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// Page
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// PerPage
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *DescribeSourceCodeRepositoryBranchPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSourceCodeRepositoryBranchPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeRepositoryBranchPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 构建仓库分支信息
		Data *BuildBranchResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceCodeRepositoryBranchPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSourceCodeRepositoryBranchPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 代码仓库类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// Page
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// PerPage
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *DescribeSourceCodeRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSourceCodeRepositoryPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 构建仓库列表
		Data *BuildReposList `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceCodeRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSourceCodeRepositoryPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的用户信息
		Data *UserInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserPersonalResponse) FromJsonString(s string) error {
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

type DockerHubRepoinfo struct {

	// 仓库名称
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 仓库类型
	Repotype *string `json:"Repotype,omitempty" name:"Repotype"`

	// 仓库Logo
	Logo *string `json:"Logo,omitempty" name:"Logo"`

	// 简述
	SimpleDesc *string `json:"SimpleDesc,omitempty" name:"SimpleDesc"`

	// 详述
	DetailDesc *string `json:"DetailDesc,omitempty" name:"DetailDesc"`

	// 收藏次数
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 是否用户的收藏
	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`
}

type DockerHubTagList struct {

	// DockerHub的仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// Tag的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*string `json:"TagList,omitempty" name:"TagList" list`
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

type ForwardRequestRequest struct {
	*tchttp.BaseRequest

	// 请求tcr对应的方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求tcr对应的路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求的实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 请求tcr中的HTTP头中Accept参数
	Accept *string `json:"Accept,omitempty" name:"Accept"`

	// 请求tcr http头中ContentType参数
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 请求tcr中的body信息
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
}

func (r *ForwardRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForwardRequestRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForwardRequestResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代理请求响应body
		ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`

		// 代理请求响应的header
		ResponseHeaders []*ResponseHeader `json:"ResponseHeaders,omitempty" name:"ResponseHeaders" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForwardRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForwardRequestResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HubSimpleInfo struct {

	// 仓库名称
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 仓库类型
	Repotype *string `json:"Repotype,omitempty" name:"Repotype"`

	// 仓库Logo
	Logo *string `json:"Logo,omitempty" name:"Logo"`

	// 仓库简述
	SimpleDesc *string `json:"SimpleDesc,omitempty" name:"SimpleDesc"`

	// 是否为收藏
	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`

	// 收藏数量
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`
}

type Limit struct {

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 配额的类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配置的值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type ManageExternalEndpointRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 操作（Create/Delete）
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *ManageExternalEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageExternalEndpointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManageExternalEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageExternalEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageExternalEndpointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type ManageInternalEndpointRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Create/Delete
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 需要接入的用户vpcid
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 需要接入的用户子网id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *ManageInternalEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageInternalEndpointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManageInternalEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageInternalEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageInternalEndpointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManageReplicationRequest struct {
	*tchttp.BaseRequest

	// 复制源实例ID
	SourceRegistryId *string `json:"SourceRegistryId,omitempty" name:"SourceRegistryId"`

	// 复制目标实例ID
	DestinationRegistryId *string `json:"DestinationRegistryId,omitempty" name:"DestinationRegistryId"`

	// 同步规则
	Rule *ReplicationRule `json:"Rule,omitempty" name:"Rule"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ManageReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageReplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManageReplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManageReplicationResponse) FromJsonString(s string) error {
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

type ModifyImageBuildPersonalRequest struct {
	*tchttp.BaseRequest

	// registry中的namespace
	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`

	// 用户在registry中的用户名
	RegistryUsername *string `json:"RegistryUsername,omitempty" name:"RegistryUsername"`

	// 镜像名称，不包含tag
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像tag的格式
	ImageTagFormat *string `json:"ImageTagFormat,omitempty" name:"ImageTagFormat"`

	// 源码所在的git服务
	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`

	// repo所在的group
	Group *string `json:"Group,omitempty" name:"Group"`

	// 源码在git服务器上的仓库名
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// 用户在git服务器上的用户名
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// Trigger
	Trigger *Trigger `json:"Trigger,omitempty" name:"Trigger"`

	// dockerfile在仓库中的路径
	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

	// 工作目录
	WorkDir *string `json:"WorkDir,omitempty" name:"WorkDir"`

	// 构建出的镜像覆盖该Tag
	ForceTag *string `json:"ForceTag,omitempty" name:"ForceTag"`

	// Args
	Args []*string `json:"Args,omitempty" name:"Args" list`
}

func (r *ModifyImageBuildPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageBuildPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageBuildPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageBuildPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageBuildPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 实例的规格
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceResponse) FromJsonString(s string) error {
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

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 实例的Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// PolicyId
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 192.168.0.0/24 白名单Ip
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifySecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例Id
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityPolicyResponse) FromJsonString(s string) error {
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

type Region struct {

	// gz
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 1
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// ap-guangzhou
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// alluser
	Status *string `json:"Status,omitempty" name:"Status"`

	// remark
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// id
	Id *int64 `json:"Id,omitempty" name:"Id"`
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

type ReplicationFilter struct {

	// 类型（name、tag和resource）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 默认为空
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ReplicationRule struct {

	// 同步规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目标命名空间
	DestNamespace *string `json:"DestNamespace,omitempty" name:"DestNamespace"`

	// 是否覆盖
	Override *bool `json:"Override,omitempty" name:"Override"`

	// 同步过滤条件
	Filters []*ReplicationFilter `json:"Filters,omitempty" name:"Filters" list`
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

type RequestFavor struct {

	// 收藏的镜像仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 收藏的镜像仓库类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 地域Id
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type RespDockerHubRepoList struct {

	// 仓库总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库信息列表
	RepoInfo []*HubSimpleInfo `json:"RepoInfo,omitempty" name:"RepoInfo" list`
}

type RespLimit struct {

	// 配额信息
	LimitInfo []*Limit `json:"LimitInfo,omitempty" name:"LimitInfo" list`
}

type ResponseHeader struct {

	// Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
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

type SecurityPolicy struct {

	// 策略索引
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// 192.168.1.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 安全策略的版本
	PolicyVersion *string `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
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

type Trigger struct {

	// 分支
	Branches []*string `json:"Branches,omitempty" name:"Branches" list`

	// Tag
	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
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

type UserInfo struct {

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// AppId
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 密码
	ApplicationToken *string `json:"ApplicationToken,omitempty" name:"ApplicationToken"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type UserIsExistsResp struct {

	// 用户是否存在
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`

	// 主账号是否存在
	MainIsExist *bool `json:"MainIsExist,omitempty" name:"MainIsExist"`
}

type ValidateApplicationTokenPersonalRequest struct {
	*tchttp.BaseRequest

	// 用户凭证
	ApplicationToken *string `json:"ApplicationToken,omitempty" name:"ApplicationToken"`
}

func (r *ValidateApplicationTokenPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateApplicationTokenPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateApplicationTokenPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证结果
		Data *bool `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateApplicationTokenPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateApplicationTokenPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateGitHubAuthPersonalRequest struct {
	*tchttp.BaseRequest

	// Code
	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *ValidateGitHubAuthPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateGitHubAuthPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateGitHubAuthPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateGitHubAuthPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateGitHubAuthPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type ValidateUserPersonalRequest struct {
	*tchttp.BaseRequest

	// 自定义用户名
	Username *string `json:"Username,omitempty" name:"Username"`
}

func (r *ValidateUserPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateUserPersonalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ValidateUserPersonalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 判断用户存在的返回值
		Data *UserIsExistsResp `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateUserPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ValidateUserPersonalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
