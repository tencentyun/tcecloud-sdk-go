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

package v20180525

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type AddClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// cvm创建透传参数，json化字符串格式
	CvmRunInstances *string `json:"CvmRunInstances,omitempty" name:"CvmRunInstances"`

	// 数据盘挂载点, 默认不挂载数据盘. 已格式化的 ext3，ext4，xfs 文件系统的数据盘将直接挂载，其他文件系统或未格式化的数据盘将自动格式化为ext4 并挂载，请注意备份数据! 无数据盘或有多块数据盘的云主机此设置不生效。
	MountTarget []*string `json:"MountTarget,omitempty" name:"MountTarget" list`

	// dockerd --graph 指定值, 默认为 /var/lib/docker
	DockerGraphPath []*string `json:"DockerGraphPath,omitempty" name:"DockerGraphPath" list`

	// base64 编码的用户脚本, 此脚本会在 k8s 组件运行后执行, 需要用户保证脚本的可重入及重试逻辑, 脚本及其生成的日志文件可在节点的 /data/ccs_userscript/ 路径查看, 如果要求节点需要在进行初始化完成后才可加入调度, 可配合 unschedulable 参数使用, 在 userScript 最后初始化完成后, 添加 kubectl uncordon nodename --kubeconfig=/root/.kube/config 命令使节点加入调度
	UserScript []*string `json:"UserScript,omitempty" name:"UserScript" list`

	// 系统名。centos7.2x86_64 或者 ubuntu16.04.1 LTSx86_64，仅当新建集群为空集群, 第一次向空集群添加节点时需要指定. 当集群系统确定后, 后续添加的节点都是集群系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 设置加入的节点是否参与调度，默认值为0，表示参与调度；非0表示不参与调度, 待节点初始化完成之后, 可执行kubectl uncordon nodename使node加入调度.
	Unschedulable *int64 `json:"Unschedulable,omitempty" name:"Unschedulable"`
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

		// 节点实例id
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

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

type AddExistedInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。（目前仅支持设置单个sgId）
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 重装系统时，可以指定修改实例的HostName(集群为HostName模式时，此参数必传，规则名称除不支持大写字符外与CVM创建实例接口HostName一致)
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *AddExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddExistedInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddExistedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 失败的节点ID
		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds" list`

		// 成功的节点ID
		SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds" list`

		// 超时未返回出来节点的ID(可能失败，也可能成功)
		TimeoutInstanceIds []*string `json:"TimeoutInstanceIds,omitempty" name:"TimeoutInstanceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddExistedInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AutoScalingGroupRange struct {

	// 伸缩组最小实例数
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// 伸缩组最大实例数
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type CheckClusterCIDRRequest struct {
	*tchttp.BaseRequest

	// 集群的vpc-id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群的CIDR
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
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

		// 是否存在CIDR冲突。
		IsConflict *bool `json:"IsConflict,omitempty" name:"IsConflict"`

		// CIDR冲突的类型("CIDR_CONFLICT_WITH_OTHER_CLUSTER" 同VPC其他集群CIDR存在冲突
	// "CIDR_CONFLICT_WITH_VPC_CIDR" 与VPC的CIDR存在冲突
	// "CIDR_CONFLICT_WITH_VPC_GLOBAL_ROUTE" 与同VPC的全局路由存在冲突)。
		ConflictType *string `json:"ConflictType,omitempty" name:"ConflictType"`

		// CIDR冲突描述信息。
		ConflictMsg *string `json:"ConflictMsg,omitempty" name:"ConflictMsg"`

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

type Cluster struct {

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述
	ClusterDescription *string `json:"ClusterDescription,omitempty" name:"ClusterDescription"`

	// 集群版本（默认值为1.10.5）
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 集群系统。centos7.2x86_64 或者 ubuntu16.04.1 LTSx86_64，默认取值为ubuntu16.04.1 LTSx86_64
	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`

	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群网络相关参数
	ClusterNetworkSettings *ClusterNetworkSettings `json:"ClusterNetworkSettings,omitempty" name:"ClusterNetworkSettings"`

	// 集群当前node数量
	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`

	// 集群所属的项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 标签描述列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification" list`

	// 集群状态 (Running 运行中  Creating 创建中 Abnormal 异常  )
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群属性(包括集群不同属性的MAP，属性字段包括NodeNameType (lan-ip模式和hostname 模式，默认无lan-ip模式))
	// 注意：此字段可能返回 null，表示取不到有效值。
	Property *string `json:"Property,omitempty" name:"Property"`

	// 集群当前master数量
	ClusterMaterNodeNum *uint64 `json:"ClusterMaterNodeNum,omitempty" name:"ClusterMaterNodeNum"`

	// 集群使用镜像id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// OsCustomizeType
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// 集群运行环境docker或container
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ClusterAdvancedSettings struct {

	// 是否启用IPVS
	IPVS *bool `json:"IPVS,omitempty" name:"IPVS"`

	// 是否启用集群节点自动扩缩容(创建集群流程不支持开启此功能)
	AsEnabled *bool `json:"AsEnabled,omitempty" name:"AsEnabled"`

	// 集群使用的runtime类型，包括"docker"和"containerd"两种类型，默认为"docker"
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// 集群中节点NodeName类型（包括 hostname,lan-ip两种形式，默认为lan-ip。如果开启了hostname模式，创建节点时需要设置HostName参数，并且InstanceName需要和HostName一致）
	NodeNameType *string `json:"NodeNameType,omitempty" name:"NodeNameType"`

	// 集群自定义参数
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// 集群网络类型（包括GR(全局路由)和VPC-CNI两种模式，默认为GR。
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 集群VPC-CNI模式是否为非固定IP，默认: FALSE 固定IP。
	IsNonStaticIpMode *bool `json:"IsNonStaticIpMode,omitempty" name:"IsNonStaticIpMode"`
}

type ClusterAsGroup struct {

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组状态(开启 enabled 开启中 enabling 关闭 disabled 关闭中 disabling 更新中 updating 删除中 deleting 开启缩容中 scaleDownEnabling 关闭缩容中 scaleDownDisabling)
	Status *string `json:"Status,omitempty" name:"Status"`

	// 节点是否设置成不可调度
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnschedulable *bool `json:"IsUnschedulable,omitempty" name:"IsUnschedulable"`

	// 伸缩组的label列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*Label `json:"Labels,omitempty" name:"Labels" list`
}

type ClusterAsGroupAttribute struct {

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 是否开启
	AutoScalingGroupEnabled *bool `json:"AutoScalingGroupEnabled,omitempty" name:"AutoScalingGroupEnabled"`

	// 伸缩组最大最小实例数
	AutoScalingGroupRange *AutoScalingGroupRange `json:"AutoScalingGroupRange,omitempty" name:"AutoScalingGroupRange"`
}

type ClusterAsGroupOption struct {

	// 是否开启缩容
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsScaleDownEnabled *bool `json:"IsScaleDownEnabled,omitempty" name:"IsScaleDownEnabled"`

	// 多伸缩组情况下扩容选择算法(random 随机选择，most-pods 最多类型的Pod least-waste 最少的资源浪费，默认为random)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expander *string `json:"Expander,omitempty" name:"Expander"`

	// 最大并发缩容数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxEmptyBulkDelete *int64 `json:"MaxEmptyBulkDelete,omitempty" name:"MaxEmptyBulkDelete"`

	// 集群扩容后多少分钟开始判断缩容（默认为10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleDownDelay *int64 `json:"ScaleDownDelay,omitempty" name:"ScaleDownDelay"`

	// 节点连续空闲多少分钟后被缩容（默认为 10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleDownUnneededTime *int64 `json:"ScaleDownUnneededTime,omitempty" name:"ScaleDownUnneededTime"`

	// 节点资源使用量低于多少(百分比)时认为空闲(默认: 50(百分比))
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleDownUtilizationThreshold *int64 `json:"ScaleDownUtilizationThreshold,omitempty" name:"ScaleDownUtilizationThreshold"`

	// 含有本地存储Pod的节点是否不缩容(默认： FALSE)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipNodesWithLocalStorage *bool `json:"SkipNodesWithLocalStorage,omitempty" name:"SkipNodesWithLocalStorage"`

	// 含有kube-system namespace下非DaemonSet管理的Pod的节点是否不缩容 (默认： FALSE)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipNodesWithSystemPods *bool `json:"SkipNodesWithSystemPods,omitempty" name:"SkipNodesWithSystemPods"`

	// 计算资源使用量时是否默认忽略DaemonSet的实例(默认值: False，不忽略)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreDaemonSetsUtilization *bool `json:"IgnoreDaemonSetsUtilization,omitempty" name:"IgnoreDaemonSetsUtilization"`
}

type ClusterBasicSettings struct {

	// 集群系统。centos7.2x86_64 或者 ubuntu16.04.1 LTSx86_64，默认取值为ubuntu16.04.1 LTSx86_64
	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`

	// 集群版本,默认值为1.10.5
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述
	ClusterDescription *string `json:"ClusterDescription,omitempty" name:"ClusterDescription"`

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群内新增资源所属项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到集群实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification" list`

	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// 是否开启节点的默认安全组(默认: 否，Aphla特性)
	NeedWorkSecurityGroup *bool `json:"NeedWorkSecurityGroup,omitempty" name:"NeedWorkSecurityGroup"`
}

type ClusterCIDRSettings struct {

	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突。且网段范围必须在内网网段内，例如:10.1.0.0/14, 192.168.0.1/18,172.16.0.0/16。
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 是否忽略 ClusterCIDR 冲突错误, 默认不忽略
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitempty" name:"IgnoreClusterCIDRConflict"`

	// 集群中每个Node上最大的Pod数量。取值范围4～256。不为2的幂值时会向上取最接近的2的幂值。
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// 集群最大的service数量。取值范围32～32768，不为2的幂值时会向上取最接近的2的幂值。
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`

	// 用于分配集群服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突。且网段范围必须在内网网段内，例如:10.1.0.0/14, 192.168.0.1/18,172.16.0.0/16。
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`

	// VPC-CNI网络模式下，弹性网卡的子网Id。
	EniSubnetIds []*string `json:"EniSubnetIds,omitempty" name:"EniSubnetIds" list`

	// VPC-CNI网络模式下，弹性网卡IP的回收时间，取值范围[300,15768000)
	ClaimExpiredSeconds *int64 `json:"ClaimExpiredSeconds,omitempty" name:"ClaimExpiredSeconds"`
}

type ClusterExtraArgs struct {

	// kube-apiserver自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeAPIServer []*string `json:"KubeAPIServer,omitempty" name:"KubeAPIServer" list`

	// kube-controller-manager自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeControllerManager []*string `json:"KubeControllerManager,omitempty" name:"KubeControllerManager" list`

	// kube-scheduler自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeScheduler []*string `json:"KubeScheduler,omitempty" name:"KubeScheduler" list`
}

type ClusterNetworkSettings struct {

	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 是否忽略 ClusterCIDR 冲突错误, 默认不忽略
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitempty" name:"IgnoreClusterCIDRConflict"`

	// 集群中每个Node上最大的Pod数量(默认为256)
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// 集群最大的service数量(默认为256)
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`

	// 是否启用IPVS(默认不开启)
	Ipvs *bool `json:"Ipvs,omitempty" name:"Ipvs"`

	// 集群的VPCID（如果创建空集群，为必传值，否则自动设置为和集群的节点保持一致）
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 网络插件是否启用CNI(默认开启)
	Cni *bool `json:"Cni,omitempty" name:"Cni"`
}

type CreateClusterAsGroupRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 伸缩组创建透传参数，json化字符串格式，详见伸缩组创建实例接口。LaunchConfigurationId由LaunchConfigurePara参数创建，不支持填写
	AutoScalingGroupPara *string `json:"AutoScalingGroupPara,omitempty" name:"AutoScalingGroupPara"`

	// 启动配置创建透传参数，json化字符串格式，详见创建启动配置接口。另外ImageId参数由于集群维度已经有的ImageId信息，这个字段不需要填写。UserData字段设置通过UserScript设置，这个字段不需要填写。
	LaunchConfigurePara *string `json:"LaunchConfigurePara,omitempty" name:"LaunchConfigurePara"`

	// 节点高级配置信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 节点Label数组
	Labels []*Label `json:"Labels,omitempty" name:"Labels" list`
}

func (r *CreateClusterAsGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterAsGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterAsGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 启动配置ID
		LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

		// 伸缩组ID
		AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterAsGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterAsGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// CVM创建透传参数，json化字符串格式，详见CVM创建实例接口。
	RunInstancePara *string `json:"RunInstancePara,omitempty" name:"RunInstancePara"`

	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
}

func (r *CreateClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点实例ID
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群容器网络配置信息
	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`

	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// CVM创建透传参数，json化字符串格式，详见CVM创建实例接口。总机型(包括地域)数量不超过10个，相同机型(地域)购买多台机器可以通过设置参数中RunInstances中InstanceCount来实现。
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode" list`

	// 集群的基本配置信息
	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`

	// 集群高级配置信息
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`

	// 节点高级配置信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 已存在实例的配置信息。所有实例必须在同一个VPC中，最大数量不超过100。
	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitempty" name:"ExistedInstancesForNode" list`

	// CVM类型和其对应的数据盘挂载配置信息
	InstanceDataDiskMountSettings []*InstanceDataDiskMountSetting `json:"InstanceDataDiskMountSettings,omitempty" name:"InstanceDataDiskMountSettings" list`
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

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

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

type DataDisk struct {

	// 云盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 文件系统(ext3/ext4/xfs)
	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`

	// 云盘大小(G）
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 是否自动化格式盘并挂载
	AutoFormatAndMount *bool `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`

	// 挂载目录
	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`
}

type DeleteClusterAsGroupsRequest struct {
	*tchttp.BaseRequest

	// 集群ID，通过DescribeClusters接口获取。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群伸缩组ID的列表
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds" list`

	// 是否保留伸缩组中的节点(默认值： false(不保留))
	KeepInstance *bool `json:"KeepInstance,omitempty" name:"KeepInstance"`
}

func (r *DeleteClusterAsGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterAsGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterAsGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterAsGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterAsGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 主机InstanceId列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 集群实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`

	// 是否强制删除(当节点在初始化时，可以指定参数为TRUE)
	ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *DeleteClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
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

type DescribeClusterAsGroupOptionRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterAsGroupOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterAsGroupOptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAsGroupOptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群弹性伸缩属性
		ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitempty" name:"ClusterAsGroupOption"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAsGroupOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterAsGroupOptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAsGroupsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 伸缩组ID列表，如果为空，表示拉取集群关联的所有伸缩组。
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds" list`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterAsGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterAsGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAsGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群关联的伸缩组总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群关联的伸缩组列表
		ClusterAsGroupSet *ClusterAsGroup `json:"ClusterAsGroupSet,omitempty" name:"ClusterAsGroupSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAsGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterAsGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceIdsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 节点角色, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, 默认为WORKER。默认为WORKER类型。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 根据节点的IP地址进行搜索，同时搜索内网IP和外网IP
	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`

	// 根据节点的名称进行模糊搜索
	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`

	// 根据节点的状态进行筛选
	InstanceStates []*string `json:"InstanceStates,omitempty" name:"InstanceStates" list`
}

func (r *DescribeClusterInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstanceIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点ID列表
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstanceIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 需要获取的节点实例Id列表。如果为空，表示拉取集群下所有节点实例。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 节点角色, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, 默认为WORKER。默认为WORKER类型。
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
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

		// 集群中实例总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群中实例列表
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

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

type DescribeClusterSecurityRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterSecurityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterSecurityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSecurityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群的账号名称
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 集群的访问密码
		Password *string `json:"Password,omitempty" name:"Password"`

		// 集群访问CA证书
		CertificationAuthority *string `json:"CertificationAuthority,omitempty" name:"CertificationAuthority"`

		// 集群访问的地址
		ClusterExternalEndpoint *string `json:"ClusterExternalEndpoint,omitempty" name:"ClusterExternalEndpoint"`

		// 集群访问的域名
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// 集群Endpoint地址
		PgwEndpoint *string `json:"PgwEndpoint,omitempty" name:"PgwEndpoint"`

		// 集群访问策略组
		SecurityPolicy []*string `json:"SecurityPolicy,omitempty" name:"SecurityPolicy" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterSecurityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterSecurityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServicesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 是否使用所有命名空间
	AllNamespace *uint64 `json:"AllNamespace,omitempty" name:"AllNamespace"`
}

func (r *DescribeClusterServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterServicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的Service总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Service详细信息
		Services []*SummaryService `json:"Services,omitempty" name:"Services" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterServicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件,当前只支持按照单个条件ClusterName进行过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
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

		// 集群总个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群信息列表
		Clusters []*Cluster `json:"Clusters,omitempty" name:"Clusters" list`

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

type DescribeExistedInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写查询集群列表 接口中返回的 ClusterId 字段（仅通过ClusterId获取需要过滤条件中的VPCID。节点状态比较时会使用该地域下所有集群中的节点进行比较。参数不支持同时指定InstanceIds和ClusterId。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 按照一个或者多个实例ID查询。实例ID形如：ins-xxxxxxxx。（此参数的具体格式可参考API简介的id.N一节）。每次请求的实例的上限为100。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 过滤条件,字段和详见CVM查询实例如果设置了ClusterId，会附加集群的VPCID作为查询字段，在此情况下如果在Filter中指定了"vpc-id"作为过滤字段，指定的VPCID必须与集群的VPCID相同。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 实例IP进行过滤(同时支持内网IP和外网IP)
	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`

	// 实例名称进行过滤
	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExistedInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExistedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已经存在的实例信息数组。
		ExistedInstanceSet []*ExistedInstance `json:"ExistedInstanceSet,omitempty" name:"ExistedInstanceSet" list`

		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExistedInstancesResponse) FromJsonString(s string) error {
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

		// 镜像数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 镜像信息列表
		ImageInstanceSet []*ImageInstance `json:"ImageInstanceSet,omitempty" name:"ImageInstanceSet" list`

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

type DescribeQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该账号当前地域支持的最大集群数量
		MaxClustersNum *uint64 `json:"MaxClustersNum,omitempty" name:"MaxClustersNum"`

		// 该账号当前地域单集群支持的最大节点数量
		MaxNodesNum *uint64 `json:"MaxNodesNum,omitempty" name:"MaxNodesNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaResponse) FromJsonString(s string) error {
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

		// 地域的数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 地域列表
		RegionInstanceSet []*RegionInstance `json:"RegionInstanceSet,omitempty" name:"RegionInstanceSet" list`

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

type DescribeVersionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVersionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 版本数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 版本列表
		VersionInstanceSet []*VersionInstance `json:"VersionInstanceSet,omitempty" name:"VersionInstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVersionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DrainClusterNodeRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DrainClusterNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DrainClusterNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DrainClusterNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DrainClusterNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DrainClusterNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type ExistedInstance struct {

	// 实例是否支持加入集群(TRUE 可以加入 FALSE 不能加入)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Usable *bool `json:"Usable,omitempty" name:"Usable"`

	// 实例不支持加入的原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnusableReason *string `json:"UnusableReason,omitempty" name:"UnusableReason"`

	// 实例已经所在的集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlreadyInCluster *string `json:"AlreadyInCluster,omitempty" name:"AlreadyInCluster"`

	// 实例ID形如：ins-xxxxxxxx。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例主网卡的内网IP列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// 实例主网卡的公网IP列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// 创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 实例计费模式。取值范围：
	// PREPAID：表示预付费，即包年包月
	// POSTPAID_BY_HOUR：表示后付费，即按量计费
	// CDHPAID：CDH付费，即只对CDH计费，不对CDH上的实例计费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例的CPU核数，单位：核。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// 实例内存容量，单位：GB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 操作系统名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 实例机型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type ExistedInstancesForNode struct {

	// 节点角色，取值:MASTER_ETCD, WORKER。MASTER_ETCD只有在创建 INDEPENDENT_CLUSTER 独立集群时需要指定。MASTER_ETCD节点数量为3～7，建议为奇数。MASTER_ETCD最小配置为4C8G。
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// 已存在实例的重装参数
	ExistedInstancesPara *ExistedInstancesPara `json:"ExistedInstancesPara,omitempty" name:"ExistedInstancesPara"`

	// 节点高级设置，会覆盖集群级别设置的InstanceAdvancedSettings（当前只对节点自定义参数ExtraArgs生效）
	InstanceAdvancedSettingsOverride *InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverride,omitempty" name:"InstanceAdvancedSettingsOverride"`
}

type ExistedInstancesPara struct {

	// 集群ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *string `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）
	LoginSettings *string `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 重装系统时，可以指定修改实例的HostName(集群为HostName模式时，此参数必传，规则名称除不支持大写字符外与CVM创建实例接口HostName一致)
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type Filter struct {

	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ForwardRequestRequest struct {
	*tchttp.BaseRequest

	// 请求tke-apiserver http请求对应的方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求tke-apiserver  http请求访问路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求tke-apiserver http头中Accept参数
	Accept *string `json:"Accept,omitempty" name:"Accept"`

	// 请求tke-apiserver http头中ContentType参数
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 请求tke-apiserver http请求body信息
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`

	// 请求tke-apiserver http头中X-TKE-ClusterName参数
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 是否编码请求body信息
	EncodedBody *bool `json:"EncodedBody,omitempty" name:"EncodedBody"`
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

		// 请求tke-apiserver http请求返回的body信息
		ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`

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

type ImageInstance struct {

	// 镜像别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 操作系统名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 镜像ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`
}

type Instance struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点角色, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, 默认为WORKER
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 实例异常(或者处于初始化中)的原因
	FailedReason *string `json:"FailedReason,omitempty" name:"FailedReason"`

	// 实例的状态（running 运行中，initializing 初始化中，failed 异常）
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// 实例是否封锁状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrainStatus *string `json:"DrainStatus,omitempty" name:"DrainStatus"`

	// 节点配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 添加时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type InstanceAdvancedSettings struct {

	// 数据盘挂载点, 默认不挂载数据盘. 已格式化的 ext3，ext4，xfs 文件系统的数据盘将直接挂载，其他文件系统或未格式化的数据盘将自动格式化为ext4 并挂载，请注意备份数据! 无数据盘或有多块数据盘的云主机此设置不生效。
	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`

	// dockerd --graph 指定值, 默认为 /var/lib/docker
	DockerGraphPath *string `json:"DockerGraphPath,omitempty" name:"DockerGraphPath"`

	// base64 编码的用户脚本, 此脚本会在 k8s 组件运行后执行, 需要用户保证脚本的可重入及重试逻辑, 脚本及其生成的日志文件可在节点的 /data/ccs_userscript/ 路径查看, 如果要求节点需要在进行初始化完成后才可加入调度, 可配合 unschedulable 参数使用, 在 userScript 最后初始化完成后, 添加 kubectl uncordon nodename --kubeconfig=/root/.kube/config 命令使节点加入调度
	UserScript *string `json:"UserScript,omitempty" name:"UserScript"`

	// 设置加入的节点是否参与调度，默认值为0，表示参与调度；非0表示不参与调度, 待节点初始化完成之后, 可执行kubectl uncordon nodename使node加入调度.
	Unschedulable *int64 `json:"Unschedulable,omitempty" name:"Unschedulable"`

	// 节点Label数组
	Labels []*Label `json:"Labels,omitempty" name:"Labels" list`

	// 数据盘相关信息
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 节点相关的自定义参数信息
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
}

type InstanceDataDiskMountSetting struct {

	// CVM实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 数据盘挂载信息
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// CVM实例所属可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type InstanceExtraArgs struct {

	// kubelet自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kubelet []*string `json:"Kubelet,omitempty" name:"Kubelet" list`
}

type Label struct {

	// map表中的Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// map表中的Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ModifyClusterAsGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群关联的伸缩组属性
	ClusterAsGroupAttribute *ClusterAsGroupAttribute `json:"ClusterAsGroupAttribute,omitempty" name:"ClusterAsGroupAttribute"`
}

func (r *ModifyClusterAsGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterAsGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAsGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterAsGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterAsGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAsGroupOptionAttributeRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群弹性伸缩属性
	ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitempty" name:"ClusterAsGroupOption"`
}

func (r *ModifyClusterAsGroupOptionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterAsGroupOptionAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAsGroupOptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterAsGroupOptionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterAsGroupOptionAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAttributeRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群所属项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
}

func (r *ModifyClusterAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群所属项目
		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// 集群名称
		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

		// 集群描述
		ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInstance struct {

	// 地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 地域特性开关(按照JSON的形式返回所有属性)
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeatureGates *string `json:"FeatureGates,omitempty" name:"FeatureGates"`

	// 地域简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 地域白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type RunInstancesForNode struct {

	// 节点角色，取值:MASTER_ETCD, WORKER。MASTER_ETCD只有在创建 INDEPENDENT_CLUSTER 独立集群时需要指定。MASTER_ETCD节点数量为3～7，建议为奇数。MASTER_ETCD节点最小配置为4C8G。
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// CVM创建透传参数，json化字符串格式，详见CVM创建实例接口，传入公共参数外的其他参数即可，其中ImageId会替换为TKE集群OS对应的镜像。
	RunInstancesPara []*string `json:"RunInstancesPara,omitempty" name:"RunInstancesPara" list`

	// 节点高级设置，该参数会覆盖集群级别设置的InstanceAdvancedSettings，和上边的RunInstancesPara按照顺序一一对应（当前只对节点自定义参数ExtraArgs生效）。
	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitempty" name:"InstanceAdvancedSettingsOverrides" list`
}

type RunMonitorServiceEnabled struct {

	// 是否开启云监控服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {

	// 是否开启云安全服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ServiceMeshForwardRequestRequest struct {
	*tchttp.BaseRequest

	// method
	Method *string `json:"Method,omitempty" name:"Method"`

	// path
	Path *string `json:"Path,omitempty" name:"Path"`

	// accept
	Accept *string `json:"Accept,omitempty" name:"Accept"`

	// ContentType
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// RequestBody
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
}

func (r *ServiceMeshForwardRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ServiceMeshForwardRequestRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ServiceMeshForwardRequestResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的内容
		ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServiceMeshForwardRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ServiceMeshForwardRequestResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SummaryService struct {

	// Service名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Service状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// Service IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceIp *string `json:"ServiceIp,omitempty" name:"ServiceIp"`

	// 外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalIp *string `json:"ExternalIp,omitempty" name:"ExternalIp"`

	// 负载均衡ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LbId *string `json:"LbId,omitempty" name:"LbId"`

	// 负载均衡状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	LbStatus *string `json:"LbStatus,omitempty" name:"LbStatus"`

	// Service访问类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`

	// 期望副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredReplicas *int64 `json:"DesiredReplicas,omitempty" name:"DesiredReplicas"`

	// 当前副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentReplicas *int64 `json:"CurrentReplicas,omitempty" name:"CurrentReplicas"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type Tag struct {

	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagSpecification struct {

	// 标签绑定的资源类型，当前支持类型："cluster"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 标签对列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type VersionInstance struct {

	// 版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// Remark
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}
