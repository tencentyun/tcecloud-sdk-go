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

package v20191021

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ArtifactPackage struct {

	// 包 ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 制品仓库 ID
	RepoId *uint64 `json:"RepoId,omitempty" name:"RepoId"`

	// 包名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 包描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 包下的版本数量
	VersionCount *uint64 `json:"VersionCount,omitempty" name:"VersionCount"`

	// 创建时间
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 删除时间
	DeletedAt *float64 `json:"DeletedAt,omitempty" name:"DeletedAt"`

	// 创建者 ID
	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`

	// 最新推送版本的版本号
	LatestVersionName *string `json:"LatestVersionName,omitempty" name:"LatestVersionName"`

	// 发布策略
	ReleaseStrategy *uint64 `json:"ReleaseStrategy,omitempty" name:"ReleaseStrategy"`

	// 最新推送版本的版本号 ID
	LatestVersionId *uint64 `json:"LatestVersionId,omitempty" name:"LatestVersionId"`

	// 最新推送版本的版本发布状态
	LatestVersionReleaseStatus *uint64 `json:"LatestVersionReleaseStatus,omitempty" name:"LatestVersionReleaseStatus"`
}

type ArtifactRepository struct {

	// 制品仓库 ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 仓库名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 团队 id
	TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`

	// 项目 id
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 仓库状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 仓库类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 头像地址
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// 仓库大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 仓库描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 权限范围
	AccessLevel *uint64 `json:"AccessLevel,omitempty" name:"AccessLevel"`

	// 创建时间
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 删除时间
	DeletedAt *uint64 `json:"DeletedAt,omitempty" name:"DeletedAt"`

	// 存储规则
	StorageRule *uint64 `json:"StorageRule,omitempty" name:"StorageRule"`

	// 创建者 id
	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`

	// 发布策略
	ReleaseStrategy *uint64 `json:"ReleaseStrategy,omitempty" name:"ReleaseStrategy"`
}

type ArtifactVersion struct {

	// 版本 ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 创建时间
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 制品包 ID
	PkgId *uint64 `json:"PkgId,omitempty" name:"PkgId"`

	// 版本号
	Version *string `json:"Version,omitempty" name:"Version"`

	// 版本哈希
	Hash *string `json:"Hash,omitempty" name:"Hash"`

	// 下载量
	DownloadCount *uint64 `json:"DownloadCount,omitempty" name:"DownloadCount"`

	// 上传者名称
	Uploader *string `json:"Uploader,omitempty" name:"Uploader"`

	// 上传者 id
	UploaderId *uint64 `json:"UploaderId,omitempty" name:"UploaderId"`

	// 版本描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 发布状态
	ReleaseStatus *uint64 `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`

	// 删除时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletedAt *uint64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
}

type CIBuild struct {

	// 构建 ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`

	// CCI ID
	CciId *string `json:"CciId,omitempty" name:"CciId"`

	// 构建编号
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 构建 commit sha
	CommitId *string `json:"CommitId,omitempty" name:"CommitId"`

	// 触发信息
	Cause *string `json:"Cause,omitempty" name:"Cause"`

	// 构建的分支
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 失败信息
	FailedMessage *string `json:"FailedMessage,omitempty" name:"FailedMessage"`

	// Jenkinsfile 内容
	JenkinsFileContent *string `json:"JenkinsFileContent,omitempty" name:"JenkinsFileContent"`

	// 构建耗时
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 构建状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 构建进行到了哪个 stage/node
	StatusNode *string `json:"StatusNode,omitempty" name:"StatusNode"`

	// 测试结果
	TestResult *CIBuildTestResult `json:"TestResult,omitempty" name:"TestResult"`

	// 创建时间
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type CIBuildTestResult struct {

	// 是否空
	Empty *bool `json:"Empty,omitempty" name:"Empty"`

	// 失败次数
	FailCount *uint64 `json:"FailCount,omitempty" name:"FailCount"`

	// 通过次数
	PassCount *uint64 `json:"PassCount,omitempty" name:"PassCount"`

	// 跳过次数
	SkipCount *uint64 `json:"SkipCount,omitempty" name:"SkipCount"`

	// 总次数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 时长
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type CIJob struct {

	// 任务 ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 项目 ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 仓库 ID
	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`

	// 仓库名称
	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 执行方式
	ExecuteIn *uint64 `json:"ExecuteIn,omitempty" name:"ExecuteIn"`

	// 触发机制
	TriggerMethods []*uint64 `json:"TriggerMethods,omitempty" name:"TriggerMethods" list`

	// 代码更新触发匹配规则
	HookType *uint64 `json:"HookType,omitempty" name:"HookType"`

	// 触发构建的分支
	BranchSelector *string `json:"BranchSelector,omitempty" name:"BranchSelector"`

	// 分支匹配正则
	BranchRegex *string `json:"BranchRegex,omitempty" name:"BranchRegex"`

	// Jenkinsfile 来源
	JenkinsFileFromType *uint64 `json:"JenkinsFileFromType,omitempty" name:"JenkinsFileFromType"`

	// Jenkinsfile 在仓库中的文件路径
	JenkinsFilePath *string `json:"JenkinsFilePath,omitempty" name:"JenkinsFilePath"`

	// Jenkinsfile 文件内容
	JenkinsFileStaticContent *string `json:"JenkinsFileStaticContent,omitempty" name:"JenkinsFileStaticContent"`

	// 自动取消相同版本
	AutoCancelSameRevision *bool `json:"AutoCancelSameRevision,omitempty" name:"AutoCancelSameRevision"`

	// 自动取消相同 MR
	AutoCancelSameMergeRequest *bool `json:"AutoCancelSameMergeRequest,omitempty" name:"AutoCancelSameMergeRequest"`

	// 构建结果通知触发者机制
	TriggerRemind *uint64 `json:"TriggerRemind,omitempty" name:"TriggerRemind"`

	// 任务缓存目录配置
	CachePaths []*CIJobCachePath `json:"CachePaths,omitempty" name:"CachePaths" list`

	// 环境变量配置
	Envs []*CIJobEnv `json:"Envs,omitempty" name:"Envs" list`

	// 针对 CRON triggerMethod 的 schedule 规则配置
	Schedules []*CIJobSchedule `json:"Schedules,omitempty" name:"Schedules" list`

	// 不管构建成功还是失败总是通知的用户
	AlwaysUserIds []*uint64 `json:"AlwaysUserIds,omitempty" name:"AlwaysUserIds" list`

	// 仅构建失败时要通知的用户
	BuildFailIds []*uint64 `json:"BuildFailIds,omitempty" name:"BuildFailIds" list`

	// 创建者
	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`

	// 创建时间
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 最后更新时间
	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type CIJobCachePath struct {

	// 绝对路径
	AbsolutePath *string `json:"AbsolutePath,omitempty" name:"AbsolutePath"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type CIJobEnv struct {

	// 环境变量名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 环境变量值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 是否保密
	Sensitive *bool `json:"Sensitive,omitempty" name:"Sensitive"`
}

type CIJobSchedule struct {

	// 代码无变化时是否触发
	RefChangeTrigger *bool `json:"RefChangeTrigger,omitempty" name:"RefChangeTrigger"`

	// 要触发的分支
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 星期几
	Weekend *string `json:"Weekend,omitempty" name:"Weekend"`

	// 是否周期触发（周期触发/单次触发）
	Repeat *bool `json:"Repeat,omitempty" name:"Repeat"`

	// 开始时间。如果是周期触发，精确到小时（ 8 ）如果是单次触发，精确到分钟数（ 8:20 ）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。如果是单次触发，结束时间为空
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 间隔
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type CIJobTest struct {

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 项目 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 仓库 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`

	// 仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 执行方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecuteIn *string `json:"ExecuteIn,omitempty" name:"ExecuteIn"`

	// 触发机制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerMethods []*string `json:"TriggerMethods,omitempty" name:"TriggerMethods" list`

	// 代码更新触发匹配规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	HookType *string `json:"HookType,omitempty" name:"HookType"`

	// 触发构建的分支
	// 注意：此字段可能返回 null，表示取不到有效值。
	BranchSelector *string `json:"BranchSelector,omitempty" name:"BranchSelector"`

	// 分支匹配正则
	// 注意：此字段可能返回 null，表示取不到有效值。
	BranchRegex *string `json:"BranchRegex,omitempty" name:"BranchRegex"`

	// Jenkinsfile 来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	JenkinsFileFromType *string `json:"JenkinsFileFromType,omitempty" name:"JenkinsFileFromType"`

	// Jenkinsfile 在仓库中的文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	JenkinsFilePath *string `json:"JenkinsFilePath,omitempty" name:"JenkinsFilePath"`

	// Jenkinsfile 文件内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	JenkinsFileStaticContent *string `json:"JenkinsFileStaticContent,omitempty" name:"JenkinsFileStaticContent"`

	// 自动取消相同版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoCancelSameRevision *bool `json:"AutoCancelSameRevision,omitempty" name:"AutoCancelSameRevision"`

	// 自动取消相同 MR
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoCancelSameMergeRequest *bool `json:"AutoCancelSameMergeRequest,omitempty" name:"AutoCancelSameMergeRequest"`

	// 构建结果通知触发者机制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerRemind *string `json:"TriggerRemind,omitempty" name:"TriggerRemind"`

	// 任务缓存目录配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CachePaths []*CIJobCachePath `json:"CachePaths,omitempty" name:"CachePaths" list`

	// 环境变量配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envs []*CIJobEnv `json:"Envs,omitempty" name:"Envs" list`

	// 针对 CRON triggerMethod 的 schedule 规则配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schedules []*CIJobSchedule `json:"Schedules,omitempty" name:"Schedules" list`

	// 不管构建成功还是失败总是通知的用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlwaysUserIds []*uint64 `json:"AlwaysUserIds,omitempty" name:"AlwaysUserIds" list`

	// 仅构建失败时要通知的用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildFailIds []*uint64 `json:"BuildFailIds,omitempty" name:"BuildFailIds" list`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type CreateEnterpriseRequest struct {
	*tchttp.BaseRequest

	// team domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 公司名 company name / 项目名
	Name *string `json:"Name,omitempty" name:"Name"`

	// Tce账户名
	QcloudName *string `json:"QcloudName,omitempty" name:"QcloudName"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户图像 url
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

func (r *CreateEnterpriseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEnterpriseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEnterpriseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnterpriseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEnterpriseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateJobRequest struct {
	*tchttp.BaseRequest

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`

	// 项目 ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 代码仓库名称
	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 0|1|2
	ExecuteIn *uint64 `json:"ExecuteIn,omitempty" name:"ExecuteIn"`

	// 0|1|2
	TriggerMethods []*uint64 `json:"TriggerMethods,omitempty" name:"TriggerMethods" list`

	// 0|1|2
	HookType *uint64 `json:"HookType,omitempty" name:"HookType"`

	// hookType 为 DEFAULT 时须指定
	BranchSelector *string `json:"BranchSelector,omitempty" name:"BranchSelector"`

	// hookType 为 CUSTOME 时须指定
	BranchRegex *string `json:"BranchRegex,omitempty" name:"BranchRegex"`

	// 0|1
	JenkinsFileFromType *uint64 `json:"JenkinsFileFromType,omitempty" name:"JenkinsFileFromType"`

	// jenkinsFileFromType 为 SCM 时须指定
	JenkinsFilePath *string `json:"JenkinsFilePath,omitempty" name:"JenkinsFilePath"`

	// jenkinsFileFromType 为 STATIC 时须指定
	JenkinsFileStaticContent *string `json:"JenkinsFileStaticContent,omitempty" name:"JenkinsFileStaticContent"`

	// 自动取消相同版本
	AutoCancelSameRevision *bool `json:"AutoCancelSameRevision,omitempty" name:"AutoCancelSameRevision"`

	// 自动取消相同 MR
	AutoCancelSameMergeRequest *bool `json:"AutoCancelSameMergeRequest,omitempty" name:"AutoCancelSameMergeRequest"`

	// 构建结果通知触发者机制
	// 0 -总是通知;
	// 1 -仅构建失败时通知;
	TriggerRemind *uint64 `json:"TriggerRemind,omitempty" name:"TriggerRemind"`

	// 任务缓存目录配置
	CachePaths []*CIJobCachePath `json:"CachePaths,omitempty" name:"CachePaths" list`

	// 环境变量配置
	Envs []*CIJobEnv `json:"Envs,omitempty" name:"Envs" list`

	// 针对 CRON triggerMethod 的 schedule 规则配置, 暂只用于添加
	Schedules []*CIJobSchedule `json:"Schedules,omitempty" name:"Schedules" list`

	// 不管构建成功还是失败总是通知的用户
	AlwaysUserIds []*uint64 `json:"AlwaysUserIds,omitempty" name:"AlwaysUserIds" list`

	// 仅构建失败时要通知的用户
	BuildFailIds []*uint64 `json:"BuildFailIds,omitempty" name:"BuildFailIds" list`
}

func (r *CreateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// team gk
	TeamGk *string `json:"TeamGk,omitempty" name:"TeamGk"`

	// user gk
	UserGk *string `json:"UserGk,omitempty" name:"UserGk"`

	// 项目名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目 display name
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 项目描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 启用：true，禁用：false
	GitEnabled *bool `json:"GitEnabled,omitempty" name:"GitEnabled"`

	// 非必填，git ignore文件类型
	GitIgnore *string `json:"GitIgnore,omitempty" name:"GitIgnore"`

	// "true"|"false"
	GitReadmeEnabled *string `json:"GitReadmeEnabled,omitempty" name:"GitReadmeEnabled"`

	// license
	GitLicense *string `json:"GitLicense,omitempty" name:"GitLicense"`

	// git｜svn｜hg
	VcsType *string `json:"VcsType,omitempty" name:"VcsType"`

	// "true"|"false"
	CreateSvnLayout *string `json:"CreateSvnLayout,omitempty" name:"CreateSvnLayout"`

	// 非必填，项目开始时间
	StartDate *uint64 `json:"StartDate,omitempty" name:"StartDate"`

	// 非必填，项目结束时间
	EndDate *uint64 `json:"EndDate,omitempty" name:"EndDate"`

	// 0： 不公开 1：公开源代码（公开后，任何人都可以访问代码仓库，请慎重考虑！
	Shared *uint64 `json:"Shared,omitempty" name:"Shared"`

	// 非必填，设置空字符串
	// 使用预置代码模板初始化仓库，java|ror|sinatra
	Template *string `json:"Template,omitempty" name:"Template"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBuildRequest struct {
	*tchttp.BaseRequest

	// 构建编号
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteBuildRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBuildRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBuildResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBuildResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBuildResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteJobRequest struct {
	*tchttp.BaseRequest

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBuildRequest struct {
	*tchttp.BaseRequest

	// 构建编号
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeBuildRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBuildRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBuildResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 构建信息
		Build *CIBuild `json:"Build,omitempty" name:"Build"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBuildResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBuildResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCurrentUserRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCurrentUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCurrentUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCurrentUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户 Id
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 邮箱
		Email *string `json:"Email,omitempty" name:"Email"`

		// 唯一标签
		GlobalKey *string `json:"GlobalKey,omitempty" name:"GlobalKey"`

		// 头像
		Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

		// 昵称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 手机
		Phone *string `json:"Phone,omitempty" name:"Phone"`

		// 团队 ID
		TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCurrentUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCurrentUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobsRequest struct {
	*tchttp.BaseRequest

	// 项目 ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CI 任务列表
		JobSet []*CIJob `json:"JobSet,omitempty" name:"JobSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeListBuildRequest struct {
	*tchttp.BaseRequest

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`

	// 请求页
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeListBuildRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeListBuildRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeListBuildResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求页
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// 分页大小
		Size *uint64 `json:"Size,omitempty" name:"Size"`

		// 总个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 构建列表
		InstanceSet []*CIBuild `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListBuildResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeListBuildResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOneJobRequest struct {
	*tchttp.BaseRequest

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeOneJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOneJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOneJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务
		Job *CIJob `json:"Job,omitempty" name:"Job"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOneJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOneJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePackageListRequest struct {
	*tchttp.BaseRequest

	// 团队 ID
	TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`

	// 项目 ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 请求页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 仓库名
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`
}

func (r *DescribePackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePackageListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePackageListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 制品包列表
		InstanceSet []*ArtifactPackage `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 请求页
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// 分页大小
		Size *uint64 `json:"Size,omitempty" name:"Size"`

		// 数据总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePackageListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest

	// 团队 ID
	TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`

	// 项目 ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 项目名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目 ID
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 创建时间
		CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 更新时间
		UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

		// 状态
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 类型
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// 名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 显示名称
		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

		// 描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 图标
		Icon *string `json:"Icon,omitempty" name:"Icon"`

		// 团队 ID
		TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`

		// 是否为模板项目
		IsDemo *bool `json:"IsDemo,omitempty" name:"IsDemo"`

		// 最大团员数
		MaxMember *uint64 `json:"MaxMember,omitempty" name:"MaxMember"`

		// 个人所有者 ID
		UserOwnerId *uint64 `json:"UserOwnerId,omitempty" name:"UserOwnerId"`

		// 是否压缩
		Archived *bool `json:"Archived,omitempty" name:"Archived"`

		// 项目开始时间
		StartDate *uint64 `json:"StartDate,omitempty" name:"StartDate"`

		// 团队所有者 ID
		TeamOwnerId *uint64 `json:"TeamOwnerId,omitempty" name:"TeamOwnerId"`

		// 项目结束时间
		EndDate *uint64 `json:"EndDate,omitempty" name:"EndDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 团队 ID
	TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 项目列表信息
		ProjectSet []*Project `json:"ProjectSet,omitempty" name:"ProjectSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryListRequest struct {
	*tchttp.BaseRequest

	// 团队 ID
	TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`

	// 项目 ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页展示数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 制品库类型
	// 0; // 未知
	// 1; // GENERIC 
	// 2; // DOCKER 
	// 3; // MAVEN 
	// 4; // NPM 
	// 5; // PYPI 
	// 6; // HELM 
	// 7; // COMPOSER 
	// 8; // NUGET
	// 9; // CONAN
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeRepositoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 仓库列表
		InstanceSet []*ArtifactRepository `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 请求页
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// 分页大小
		Size *uint64 `json:"Size,omitempty" name:"Size"`

		// 数据总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRepositoryListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamRequest struct {
	*tchttp.BaseRequest

	// 团队唯一标示
	TeamGlobalKey *string `json:"TeamGlobalKey,omitempty" name:"TeamGlobalKey"`

	// 团队 ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTeamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTeamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTeamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 团队 ID
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 创建时间
		CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 更新时间
		UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

		// 名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 名称拼音
		NamePinyin *string `json:"NamePinyin,omitempty" name:"NamePinyin"`

		// 介绍
		Introduction *string `json:"Introduction,omitempty" name:"Introduction"`

		// 图标
		Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

		// 链接
		HtmlLink *string `json:"HtmlLink,omitempty" name:"HtmlLink"`

		// 唯一标示
		GlobalKey *string `json:"GlobalKey,omitempty" name:"GlobalKey"`

		// 是否锁定
		Lock *bool `json:"Lock,omitempty" name:"Lock"`

		// 访问路径
		Path *string `json:"Path,omitempty" name:"Path"`

		// 团队所有者
		Owner *Owner `json:"Owner,omitempty" name:"Owner"`

		// 员工数量
		StaffCount *uint64 `json:"StaffCount,omitempty" name:"StaffCount"`

		// 人员数量
		MemberCount *uint64 `json:"MemberCount,omitempty" name:"MemberCount"`

		// 管理员数量
		ManagerCount *uint64 `json:"ManagerCount,omitempty" name:"ManagerCount"`

		// 管理员是否锁定
		AdminLocked *bool `json:"AdminLocked,omitempty" name:"AdminLocked"`

		// 管理员是否强制两步校验
		ManagerForceTwofaEnabled *bool `json:"ManagerForceTwofaEnabled,omitempty" name:"ManagerForceTwofaEnabled"`

		// 当前用户角色 ID
		CurrentUserRoleId *uint64 `json:"CurrentUserRoleId,omitempty" name:"CurrentUserRoleId"`

		// 是否强制两步校验
		ForceTwofaEnabled *bool `json:"ForceTwofaEnabled,omitempty" name:"ForceTwofaEnabled"`

		// 员工是否强制两步校验
		StaffForceTwofaEnabled *bool `json:"StaffForceTwofaEnabled,omitempty" name:"StaffForceTwofaEnabled"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTeamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTeamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTestJobRequest struct {
	*tchttp.BaseRequest

	// Job ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeTestJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTestJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTestJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Job 信息
		Job *CIJobTest `json:"Job,omitempty" name:"Job"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTestJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTestJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionListRequest struct {
	*tchttp.BaseRequest

	// 团队 ID
	TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`

	// 项目 ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页展示数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 制品库名
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 制品库包名
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
}

func (r *DescribeVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVersionListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 版本列表
		InstanceSet []*ArtifactVersion `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 请求页
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// 分页大小
		Size *uint64 `json:"Size,omitempty" name:"Size"`

		// 数据总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVersionListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyJobRequest struct {
	*tchttp.BaseRequest

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`

	// 项目 ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 代码仓库名称
	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 执行方式
	// 0; // standalone，已废弃
	// 1; // cvm, 默认用这个
	// 2; // used for private deploy
	ExecuteIn *uint64 `json:"ExecuteIn,omitempty" name:"ExecuteIn"`

	// 触发机制
	// 0; // 代码更新触发
	// 1; // 定时触发
	// 2; // MR变动触发
	TriggerMethods []*uint64 `json:"TriggerMethods,omitempty" name:"TriggerMethods" list`

	// 代码更新触发匹配规则
	// 0; // 选择指定分支更新时触发
	// 1; // 推送标签时触发
	// 2; // 推送分支时触发
	// 3; // 自定义触发条件正则表达式
	HookType *uint64 `json:"HookType,omitempty" name:"HookType"`

	// hookType 为 DEFAULT 时须指定
	BranchSelector *string `json:"BranchSelector,omitempty" name:"BranchSelector"`

	// hookType 为 CUSTOME 时须指定
	BranchRegex *string `json:"BranchRegex,omitempty" name:"BranchRegex"`

	// Jenkinsfile source 来源
	// 0; // 来源仓库
	// 1; // 静态输入的，需要指定字段
	JenkinsFileFromType *uint64 `json:"JenkinsFileFromType,omitempty" name:"JenkinsFileFromType"`

	// jenkinsFileFromType 为 SCM 时须指定
	JenkinsFilePath *string `json:"JenkinsFilePath,omitempty" name:"JenkinsFilePath"`

	// jenkinsFileFromType 为 STATIC 时须指定
	JenkinsFileStaticContent *string `json:"JenkinsFileStaticContent,omitempty" name:"JenkinsFileStaticContent"`

	// 自动取消相同版本
	AutoCancelSameRevision *bool `json:"AutoCancelSameRevision,omitempty" name:"AutoCancelSameRevision"`

	// 自动取消相同 MR
	AutoCancelSameMergeRequest *bool `json:"AutoCancelSameMergeRequest,omitempty" name:"AutoCancelSameMergeRequest"`

	// 构建结果通知触发者机制
	// 0; // 总是通知
	// 1; // 仅构建失败时通知
	TriggerRemind *uint64 `json:"TriggerRemind,omitempty" name:"TriggerRemind"`

	// 任务缓存目录配置
	CachePaths []*CIJobCachePath `json:"CachePaths,omitempty" name:"CachePaths" list`

	// 环境变量配置
	Envs []*CIJobEnv `json:"Envs,omitempty" name:"Envs" list`

	// 针对 CRON triggerMethod 的 schedule 规则配置, 暂只用于添加
	Schedules []*CIJobSchedule `json:"Schedules,omitempty" name:"Schedules" list`

	// 不管构建成功还是失败总是通知的用户
	AlwaysUserIds []*uint64 `json:"AlwaysUserIds,omitempty" name:"AlwaysUserIds" list`

	// 仅构建失败时要通知的用户
	BuildFailIds []*uint64 `json:"BuildFailIds,omitempty" name:"BuildFailIds" list`
}

func (r *ModifyJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Owner struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 地址
	Location *string `json:"Location,omitempty" name:"Location"`

	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 国家
	Country *string `json:"Country,omitempty" name:"Country"`

	// 创建时间
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 最后登录时间
	LastLoginedAt *uint64 `json:"LastLoginedAt,omitempty" name:"LastLoginedAt"`

	// 职位
	Job *uint64 `json:"Job,omitempty" name:"Job"`

	// 性别
	Sex *uint64 `json:"Sex,omitempty" name:"Sex"`

	// 生日
	Birthday *uint64 `json:"Birthday,omitempty" name:"Birthday"`

	// 全局唯一标志
	GlobalKey *string `json:"GlobalKey,omitempty" name:"GlobalKey"`

	// 状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 号码校验
	PhoneValidation *uint64 `json:"PhoneValidation,omitempty" name:"PhoneValidation"`

	// 邮箱校验
	EmailValidation *uint64 `json:"EmailValidation,omitempty" name:"EmailValidation"`

	// 电话国际区号
	PhoneCountryCode *string `json:"PhoneCountryCode,omitempty" name:"PhoneCountryCode"`

	// 名称拼音
	NamePinyin *string `json:"NamePinyin,omitempty" name:"NamePinyin"`

	// 介绍
	Introduction *string `json:"Introduction,omitempty" name:"Introduction"`

	// 头像
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 电话
	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

type Project struct {

	// 项目 ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 显示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitempty" name:"Icon"`

	// 团队 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeamId *uint64 `json:"TeamId,omitempty" name:"TeamId"`

	// 是否为模板项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDemo *bool `json:"IsDemo,omitempty" name:"IsDemo"`

	// 最大团员数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMember *uint64 `json:"MaxMember,omitempty" name:"MaxMember"`

	// 个人所有者 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserOwnerId *uint64 `json:"UserOwnerId,omitempty" name:"UserOwnerId"`

	// 是否压缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Archived *bool `json:"Archived,omitempty" name:"Archived"`

	// 项目开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartDate *uint64 `json:"StartDate,omitempty" name:"StartDate"`

	// 团队所有者 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeamOwnerId *uint64 `json:"TeamOwnerId,omitempty" name:"TeamOwnerId"`

	// 项目结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndDate *uint64 `json:"EndDate,omitempty" name:"EndDate"`
}

type StopBuildRequest struct {
	*tchttp.BaseRequest

	// 构建编号
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 任务 ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *StopBuildRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopBuildRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopBuildResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopBuildResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopBuildResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TriggerBuildRequest struct {
	*tchttp.BaseRequest

	// 代码版本
	Revision *string `json:"Revision,omitempty" name:"Revision"`

	// 启动参数
	Params []*CIJobEnv `json:"Params,omitempty" name:"Params" list`

	// Job ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *TriggerBuildRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TriggerBuildRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TriggerBuildResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TriggerBuildResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TriggerBuildResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
