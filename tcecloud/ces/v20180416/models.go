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

package v20180416

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CheckClusterNameRequest struct {
	*tchttp.BaseRequest

	// 要检查的ClusterName
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 要修改ClusterName的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CheckClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckClusterNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ClusterName是否合法
		Valid *bool `json:"Valid,omitempty" name:"Valid"`

		// 检查结果信息
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckClusterNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckForceRestartRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CheckForceRestartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckForceRestartRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckForceRestartResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否需要强制重启
		NeedForceResult *bool `json:"NeedForceResult,omitempty" name:"NeedForceResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckForceRestartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckForceRestartResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckOperationRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 操作类型 <li>scalein：缩容</li><li>restart：重启</li>
	OpType *string `json:"OpType,omitempty" name:"OpType"`
}

func (r *CheckOperationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckOperationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckOperationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否检查通过
		CheckSucc *bool `json:"CheckSucc,omitempty" name:"CheckSucc"`

		// 是否需要强制重启
		NeedForceResult *bool `json:"NeedForceResult,omitempty" name:"NeedForceResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckOperationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckOperationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckScaleUpgradeRequest struct {
	*tchttp.BaseRequest

	// 需要检查的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CheckScaleUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckScaleUpgradeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckScaleUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否满足升级条件
		Valid *bool `json:"Valid,omitempty" name:"Valid"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckScaleUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckScaleUpgradeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckUpdateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 已废弃，请使用NodeInfoList
	// 变配后的节点个数（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 已废弃，请使用NodeInfoList
	// 变配后的节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 已废弃，请使用NodeInfoList
	// 变配后的磁盘大小（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 已废弃，请使用NodeInfoList
	// 变配后的专用主节点个数（仅支持3个和5个）
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitempty" name:"MasterNodeNum"`

	// 已废弃，请使用NodeInfoList
	// 变配后的专用主节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitempty" name:"MasterNodeType"`

	// 节点信息列表，可以只传递要更新的节点及其对应的规格信息。支持的操作包括<li>修改一种节点的个数</li><li>修改一种节点的节点规格及磁盘大小</li><li>增加一种节点类型（需要同时指定该节点的类型，个数，规格，磁盘等信息）</li>上述操作一次只能进行一种，且磁盘类型不支持修改
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`
}

func (r *CheckUpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckUpdateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckUpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否允许变配操作
		AllowUpdate *bool `json:"AllowUpdate,omitempty" name:"AllowUpdate"`

		// 不允许变配的原因
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckUpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckUpdateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClusterMetric struct {

	// 节点个数
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 数据节点个数
	DataNodeNum *uint64 `json:"DataNodeNum,omitempty" name:"DataNodeNum"`

	// 索引个数
	IndexNum *uint64 `json:"IndexNum,omitempty" name:"IndexNum"`

	// 文档个数
	DocNum *uint64 `json:"DocNum,omitempty" name:"DocNum"`

	// 磁盘已使用字节数
	DiskUsedInBytes *uint64 `json:"DiskUsedInBytes,omitempty" name:"DiskUsedInBytes"`

	// 分片个数
	ShardNum *uint64 `json:"ShardNum,omitempty" name:"ShardNum"`

	// 主分片个数
	PrimaryShardNum *uint64 `json:"PrimaryShardNum,omitempty" name:"PrimaryShardNum"`

	// 迁移中的分片个数
	RelocatingShardNum *uint64 `json:"RelocatingShardNum,omitempty" name:"RelocatingShardNum"`

	// 初始化中的分片个数
	InitializingShardNum *uint64 `json:"InitializingShardNum,omitempty" name:"InitializingShardNum"`

	// 未分配的分片个数
	UnassignedShardNum *uint64 `json:"UnassignedShardNum,omitempty" name:"UnassignedShardNum"`
}

type CosBackup struct {

	// 是否开启cos自动备份
	IsAutoBackup *bool `json:"IsAutoBackup,omitempty" name:"IsAutoBackup"`

	// 自动备份执行时间（精确到小时）, e.g. "22:00"
	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例版本（支持"5.6.4"、"6.4.3"、"6.8.2"、"7.5.1"）
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 访问密码（密码需8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号）
	Password *string `json:"Password,omitempty" name:"Password"`

	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 计费类型<li>POSTPAID_BY_HOUR：按小时后付费</li>
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 集群配置文件中的ClusterName（系统默认配置为实例ID，暂不支持自定义）
	ClusterNameInConf *string `json:"ClusterNameInConf,omitempty" name:"ClusterNameInConf"`

	// 集群部署方式<li>0：单可用区部署</li>
	DeployMode *uint64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// License类型<li>oss：开源版</li>
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// 节点信息列表， 用于描述集群各类节点的规格信息如节点类型，节点个数，节点规格，磁盘类型，磁盘大小等
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`
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

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

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

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

type DescribeInstanceLogsRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志类型，默认值为1
	// <li>1, 主日志</li>
	// <li>2, 搜索慢日志</li>
	// <li>3, 索引慢日志</li>
	// <li>4, GC日志</li>
	LogType *uint64 `json:"LogType,omitempty" name:"LogType"`

	// 搜索词，支持LUCENE语法，如 level:WARN、ip:1.1.1.1、message:test-index等
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 日志开始时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页起始值, 默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为100，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 时间排序方式，默认值为0
	// <li>0, 降序</li>
	// <li>1, 升序</li>
	OrderByType *uint64 `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeInstanceLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的日志条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志详细信息列表
		InstanceLogList []*InstanceLog `json:"InstanceLogList,omitempty" name:"InstanceLogList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 起始时间, e.g. "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间, e.g. "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页起始值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 操作记录
		Operations []*Operation `json:"Operations,omitempty" name:"Operations" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群实例所属可用区，不传则默认所有可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 集群实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 集群实例名称列表
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// 分页起始值, 默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段<li>1：实例ID</li><li>2：实例名称</li><li>3：可用区</li><li>4：创建时间</li>若orderKey未传递则按创建时间降序排序
	OrderByKey *uint64 `json:"OrderByKey,omitempty" name:"OrderByKey"`

	// 排序方式<li>0：升序</li><li>1：降序</li>若传递了orderByKey未传递orderByType, 则默认升序
	OrderByType *uint64 `json:"OrderByType,omitempty" name:"OrderByType"`

	// 节点标签信息列表
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`

	// 私有网络vip列表
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
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

		// 返回的实例个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息列表
		InstanceList []*InstanceInfo `json:"InstanceList,omitempty" name:"InstanceList" list`

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

type DescribeUpgradeRequest struct {
	*tchttp.BaseRequest

	// 需要升级的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUpgradeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可以升级到的大版本
		EsVersions []*string `json:"EsVersions,omitempty" name:"EsVersions" list`

		// 可以升级到的商业特性
		EsLicenses []*string `json:"EsLicenses,omitempty" name:"EsLicenses" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUpgradeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DictInfo struct {

	// 词典键值
	Key *string `json:"Key,omitempty" name:"Key"`

	// 词典名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 词典大小，单位B
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type DiskSizeRange struct {

	// 最小值
	Min *uint64 `json:"Min,omitempty" name:"Min"`

	// 小刻度值
	Sml *uint64 `json:"Sml,omitempty" name:"Sml"`

	// 中刻度值
	Med *uint64 `json:"Med,omitempty" name:"Med"`

	// 最大值
	Max *uint64 `json:"Max,omitempty" name:"Max"`
}

type EsAcl struct {

	// kibana访问黑名单
	BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList" list`

	// kibana访问白名单
	WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList" list`
}

type EsDictionaryInfo struct {

	// 启用词词典列表
	MainDict []*DictInfo `json:"MainDict,omitempty" name:"MainDict" list`

	// 停用词词典列表
	Stopwords []*DictInfo `json:"Stopwords,omitempty" name:"Stopwords" list`
}

type GetClusterMetricRequest struct {
	*tchttp.BaseRequest

	// 要获取指标的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetClusterMetricRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterMetricRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterMetricResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 集群监控指标
		ClusterMetric *ClusterMetric `json:"ClusterMetric,omitempty" name:"ClusterMetric"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterMetricResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterMetricResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetNodesMetricRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetNodesMetricRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetNodesMetricRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetNodesMetricResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 节点监控指标
		NodesMetric []*NodeMetric `json:"NodesMetric,omitempty" name:"NodesMetric" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNodesMetricResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetNodesMetricResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 用户ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 用户UIN
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 实例所属VPC的UID
	VpcUid *string `json:"VpcUid,omitempty" name:"VpcUid"`

	// 实例所属子网的UID
	SubnetUid *string `json:"SubnetUid,omitempty" name:"SubnetUid"`

	// 实例状态，0:处理中,1:正常,-1停止,-2:销毁中,-3:已销毁
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例计费模式。取值范围：  PREPAID：表示预付费，即包年包月  POSTPAID_BY_HOUR：表示后付费，即按量计费  CDHPAID：CDH付费，即只对CDH计费，不对CDH上的实例计费。
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 包年包月购买时长,单位:月
	ChargePeriod *uint64 `json:"ChargePeriod,omitempty" name:"ChargePeriod"`

	// 自动续费标识。取值范围：  NOTIFY_AND_AUTO_RENEW：通知过期且自动续费  NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费  DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费  默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点个数
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 节点CPU核数
	CpuNum *uint64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 节点内存大小，单位GB
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 节点磁盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 节点磁盘大小，单位GB
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ES域名
	EsDomain *string `json:"EsDomain,omitempty" name:"EsDomain"`

	// ES VIP
	EsVip *string `json:"EsVip,omitempty" name:"EsVip"`

	// ES端口
	EsPort *uint64 `json:"EsPort,omitempty" name:"EsPort"`

	// Kibana访问url
	KibanaUrl *string `json:"KibanaUrl,omitempty" name:"KibanaUrl"`

	// ES版本号
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// ES配置项
	EsConfig *string `json:"EsConfig,omitempty" name:"EsConfig"`

	// Kibana访问控制配置
	EsAcl *EsAcl `json:"EsAcl,omitempty" name:"EsAcl"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例最后修改操作时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例到期时间
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// 实例类型（实例类型标识，当前只有1,2两种）
	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Ik分词器配置
	IkConfig *EsDictionaryInfo `json:"IkConfig,omitempty" name:"IkConfig"`

	// 专用主节点配置
	MasterNodeInfo *MasterNodeInfo `json:"MasterNodeInfo,omitempty" name:"MasterNodeInfo"`

	// cos自动备份配置
	CosBackup *CosBackup `json:"CosBackup,omitempty" name:"CosBackup"`

	// 是否允许cos自动备份
	AllowCosBackup *bool `json:"AllowCosBackup,omitempty" name:"AllowCosBackup"`

	// 实例拥有的标签列表
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// 是否为冷热集群<li>true: 冷热集群</li><li>false: 非冷热集群</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableHotWarmMode *bool `json:"EnableHotWarmMode,omitempty" name:"EnableHotWarmMode"`

	// 冷节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmNodeType *string `json:"WarmNodeType,omitempty" name:"WarmNodeType"`

	// 冷节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmNodeNum *uint64 `json:"WarmNodeNum,omitempty" name:"WarmNodeNum"`

	// 冷节点CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmCpuNum *uint64 `json:"WarmCpuNum,omitempty" name:"WarmCpuNum"`

	// 冷节点内存内存大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmMemSize *uint64 `json:"WarmMemSize,omitempty" name:"WarmMemSize"`

	// 冷节点磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmDiskType *string `json:"WarmDiskType,omitempty" name:"WarmDiskType"`

	// 冷节点磁盘大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmDiskSize *uint64 `json:"WarmDiskSize,omitempty" name:"WarmDiskSize"`

	// 集群节点信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`

	// Es公网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsPublicUrl *string `json:"EsPublicUrl,omitempty" name:"EsPublicUrl"`

	// 多可用区网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitempty" name:"MultiZoneInfo" list`

	// 部署模式<li>0：单可用区</li><li>1：多可用区</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployMode *uint64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// ES公网访问状态<li>OPEN：开启</li><li>CLOSE：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccess *string `json:"PublicAccess,omitempty" name:"PublicAccess"`

	// ES公网访问控制配置
	EsPublicAcl *EsAcl `json:"EsPublicAcl,omitempty" name:"EsPublicAcl"`

	// Kibana内网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateUrl *string `json:"KibanaPrivateUrl,omitempty" name:"KibanaPrivateUrl"`

	// Kibana公网访问状态<li>OPEN：开启</li><li>CLOSE：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitempty" name:"KibanaPublicAccess"`

	// Kibana内网访问状态<li>OPEN：开启</li><li>CLOSE：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitempty" name:"KibanaPrivateAccess"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityType *uint64 `json:"SecurityType,omitempty" name:"SecurityType"`
}

type InstanceLog struct {

	// 日志时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 日志级别
	Level *string `json:"Level,omitempty" name:"Level"`

	// 集群节点ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 日志内容
	Message *string `json:"Message,omitempty" name:"Message"`
}

type KeyValue struct {

	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MasterNodeInfo struct {

	// 是否启用了专用主节点
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitempty" name:"EnableDedicatedMaster"`

	// 专用主节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitempty" name:"MasterNodeType"`

	// 专用主节点个数
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitempty" name:"MasterNodeNum"`

	// 专用主节点CPU核数
	MasterNodeCpuNum *uint64 `json:"MasterNodeCpuNum,omitempty" name:"MasterNodeCpuNum"`

	// 专用主节点内存大小，单位GB
	MasterNodeMemSize *uint64 `json:"MasterNodeMemSize,omitempty" name:"MasterNodeMemSize"`

	// 专用主节点磁盘大小，单位GB
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitempty" name:"MasterNodeDiskSize"`

	// 专用主节点磁盘类型
	MasterNodeDiskType *string `json:"MasterNodeDiskType,omitempty" name:"MasterNodeDiskType"`
}

type NodeInfo struct {

	// 节点数量
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点类型<li>hotData: 热数据节点</li>
	// <li>warmData: 冷数据节点</li>
	// <li>dedicatedMaster: 专用主节点</li>
	// 默认值为hotData
	Type *string `json:"Type,omitempty" name:"Type"`

	// 节点磁盘类型<li>CLOUD_SSD：SSD云硬盘</li><li>CLOUD_PREMIUM：高硬能云硬盘</li>默认值CLOUD_SSD
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 节点磁盘容量（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type NodeMetric struct {

	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 节点IP
	Host *string `json:"Host,omitempty" name:"Host"`

	// CPU使用率
	CpuUsage *float64 `json:"CpuUsage,omitempty" name:"CpuUsage"`

	// JVM内存使用率
	JvmMemUsage *float64 `json:"JvmMemUsage,omitempty" name:"JvmMemUsage"`

	// 磁盘可用空间字节数
	DiskAvailInBytes *uint64 `json:"DiskAvailInBytes,omitempty" name:"DiskAvailInBytes"`

	// 分片个数
	ShardNum *uint64 `json:"ShardNum,omitempty" name:"ShardNum"`
}

type NodeTypeInfo struct {

	// CPU核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位GB
	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`

	// 规格描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type NodeTypeResource struct {

	// 规格名称
	NodeTypeName *string `json:"NodeTypeName,omitempty" name:"NodeTypeName"`

	// 是否可售
	Available *bool `json:"Available,omitempty" name:"Available"`

	// 规格信息
	NodeTypeInfo *NodeTypeInfo `json:"NodeTypeInfo,omitempty" name:"NodeTypeInfo"`

	// SSD盘是否可售
	SsdAvailable *bool `json:"SsdAvailable,omitempty" name:"SsdAvailable"`

	// SSD磁盘取值范围
	SsdDiskSizeRange *DiskSizeRange `json:"SsdDiskSizeRange,omitempty" name:"SsdDiskSizeRange"`

	// SATA盘是否可售
	SataAvailable *bool `json:"SataAvailable,omitempty" name:"SataAvailable"`

	// SATA磁盘取值范围
	SataDiskSizeRange *DiskSizeRange `json:"SataDiskSizeRange,omitempty" name:"SataDiskSizeRange"`

	// SSD机型售罄原因
	SsdSoldOutReason *string `json:"SsdSoldOutReason,omitempty" name:"SsdSoldOutReason"`

	// SATA机型售罄原因
	SataSoldOutReason *string `json:"SataSoldOutReason,omitempty" name:"SataSoldOutReason"`
}

type Operation struct {

	// 操作唯一id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 操作开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 操作类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 操作详情
	Detail *OperationDetail `json:"Detail,omitempty" name:"Detail"`

	// 操作结果
	Result *string `json:"Result,omitempty" name:"Result"`

	// 流程任务信息
	Tasks []*TaskDetail `json:"Tasks,omitempty" name:"Tasks" list`

	// 操作进度
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
}

type OperationDetail struct {

	// 实例原始配置信息
	OldInfo []*KeyValue `json:"OldInfo,omitempty" name:"OldInfo" list`

	// 实例更新后配置信息
	NewInfo []*KeyValue `json:"NewInfo,omitempty" name:"NewInfo" list`
}

type QueryRegionZoneRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryRegionZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryRegionZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryRegionZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可售卖地域及可用区信息
		Regions []*RegionData `json:"Regions,omitempty" name:"Regions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryRegionZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryRegionZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryZoneResourceRequest struct {
	*tchttp.BaseRequest

	// 要检查的可用区
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// 操作类型（create数据节点创建，masterCreate专用主节点创建，scaleUp数据节点纵向扩容，masterAdd添加专用主节点，masterScaleUp专用主节点纵向扩容）
	OptType *string `json:"OptType,omitempty" name:"OptType"`

	// 实例ID(变配检查需要传递)
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群部署方式
	// <li>0, 单可用区部署</li>
	// <li>1, 多可用区部署</li>
	DeployMode *uint64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 计费类型<li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li>
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *QueryZoneResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryZoneResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryZoneResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用区资源描述列表
		ZoneResources []*ZoneResource `json:"ZoneResources,omitempty" name:"ZoneResources" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryZoneResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryZoneResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionData struct {

	// 地域名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 地域中文名称
	ZhName *string `json:"ZhName,omitempty" name:"ZhName"`

	// 地域是否售罄
	Available *bool `json:"Available,omitempty" name:"Available"`

	// 该地域可用区列表
	Child []*ZoneData `json:"Child,omitempty" name:"Child" list`

	// 地域英文名称
	EnName *string `json:"EnName,omitempty" name:"EnName"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否强制重启<li>true：强制重启</li><li>false：不强制重启</li>默认false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubTaskDetail struct {

	// 子任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子任务结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 子任务错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 子任务类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 子任务状态，0处理中 1成功 -1失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 升级检查失败的索引名
	FailedIndices []*string `json:"FailedIndices,omitempty" name:"FailedIndices" list`

	// 子任务结束时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 子任务等级，1警告 2失败
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type TagInfo struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TaskDetail struct {

	// 任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务进度
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`

	// 任务完成时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 子任务
	SubTasks []*SubTaskDetail `json:"SubTasks,omitempty" name:"SubTasks" list`
}

type UpdateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 配置项（JSON格式字符串）。当前仅支持以下配置项：<li>action.destructive_requires_name</li><li>indices.fielddata.cache.size</li><li>indices.query.bool.max_clause_count</li>
	EsConfig *string `json:"EsConfig,omitempty" name:"EsConfig"`

	// 默认用户elastic的密码（8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号）
	Password *string `json:"Password,omitempty" name:"Password"`

	// 更新配置时是否强制重启<li>true强制重启</li><li>false不强制重启</li>当前仅更新EsConfig时需要设置，默认值为false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`

	// COS自动备份信息
	CosBackup *CosBackup `json:"CosBackup,omitempty" name:"CosBackup"`

	// 节点信息列表，可以只传递要更新的节点及其对应的规格信息。支持的操作包括<li>修改一种节点的个数</li><li>修改一种节点的节点规格及磁盘大小</li><li>增加一种节点类型（需要同时指定该节点的类型，个数，规格，磁盘等信息）</li>上述操作一次只能进行一种，且磁盘类型不支持修改
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`

	// 访问控制列表
	EsAcl *EsAcl `json:"EsAcl,omitempty" name:"EsAcl"`

	// Kibana公网访问状态
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitempty" name:"KibanaPublicAccess"`

	// Kibana内网访问状态
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitempty" name:"KibanaPrivateAccess"`
}

func (r *UpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInternalRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要更新的Ik词库配置
	IkConfig *EsDictionaryInfo `json:"IkConfig,omitempty" name:"IkConfig"`

	// 更新类型（当前仅支持UpdateIkConfig）
	UpdateType *string `json:"UpdateType,omitempty" name:"UpdateType"`
}

func (r *UpdateInternalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInternalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInternalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInternalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInternalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标ES版本
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// 是否只做升级检查，默认值为false
	CheckOnly *bool `json:"CheckOnly,omitempty" name:"CheckOnly"`

	// 目标商业特性版本：<li>oss 开源版</li><li>basic 基础版</li>当前仅在5.6.4升级6.x版本时使用，默认值为basic
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`
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

type UpgradeLicenseRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`

	// 是否强制重启<li>true强制重启</li><li>false不强制重启</li> 默认值false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`
}

func (r *UpgradeLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLicenseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLicenseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLicenseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZoneData struct {

	// 可用区名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 可用区中文名称
	ZhName *string `json:"ZhName,omitempty" name:"ZhName"`

	// 可用区英文名称
	EnName *string `json:"EnName,omitempty" name:"EnName"`
}

type ZoneDetail struct {

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type ZoneResource struct {

	// 可用区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 是否可售
	Available *bool `json:"Available,omitempty" name:"Available"`

	// 节点规格资源列表
	NodeTypeList []*NodeTypeResource `json:"NodeTypeList,omitempty" name:"NodeTypeList" list`
}
