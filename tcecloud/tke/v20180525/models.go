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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AbnormalDetail struct {

	// kubernetes namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// kubernetes资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type AddAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 工作负载类型
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 告警策略设置
	AlarmPolicySettings *AlarmPolicySettings `json:"AlarmPolicySettings,omitempty" name:"AlarmPolicySettings"`

	// 告警通知设置
	NotifySettings *NotifySettings `json:"NotifySettings,omitempty" name:"NotifySettings"`

	// 告警屏蔽设置
	ShieldSettings *ShieldSettings `json:"ShieldSettings,omitempty" name:"ShieldSettings"`
}

func (r *AddAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddAlarmPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddAlarmPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterCIDRToCcnRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *AddClusterCIDRToCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterCIDRToCcnRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterCIDRToCcnResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云联网路由数组
		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterCIDRToCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterCIDRToCcnResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterCIDRToVbcRequest struct {
	*tchttp.BaseRequest

	// vpc唯一id
	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`

	// 集群cidr
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *AddClusterCIDRToVbcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterCIDRToVbcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterCIDRToVbcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回详情
		Detail []*CcnRoute `json:"Detail,omitempty" name:"Detail" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterCIDRToVbcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterCIDRToVbcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

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

	// 节点池选项
	NodePool *NodePoolOption `json:"NodePool,omitempty" name:"NodePool"`
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

type AlarmMetric struct {

	// 指标描述
	ArgusPolicyName *string `json:"ArgusPolicyName,omitempty" name:"ArgusPolicyName"`

	// 告警判断设置
	Evaluator *Evaluator `json:"Evaluator,omitempty" name:"Evaluator"`

	// 持续周期
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// 状态，true表示OK
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 告警指标ID
	MetricId *string `json:"MetricId,omitempty" name:"MetricId"`

	// ARGUS系统Measurement
	Measurement *string `json:"Measurement,omitempty" name:"Measurement"`

	// 统计周期
	StatisticsPeriod *int64 `json:"StatisticsPeriod,omitempty" name:"StatisticsPeriod"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type AlarmPolicy struct {

	// 告警策略ID
	AlarmPolicyId *string `json:"AlarmPolicyId,omitempty" name:"AlarmPolicyId"`

	// k8s集群ID
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`

	// k8s命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// k8s工作负载类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 告警策略settings
	AlarmPolicySettings *AlarmPolicySettings `json:"AlarmPolicySettings,omitempty" name:"AlarmPolicySettings"`

	// 告警通知settings
	NotifySettings *NotifySettings `json:"NotifySettings,omitempty" name:"NotifySettings"`

	// 告警屏蔽settings
	ShieldSettings *ShieldSettings `json:"ShieldSettings,omitempty" name:"ShieldSettings"`
}

type AlarmPolicyFilter struct {

	// 告警策略名称
	AlarmPolicyName *string `json:"AlarmPolicyName,omitempty" name:"AlarmPolicyName"`
}

type AlarmPolicySettings struct {

	// 告警策略名称
	AlarmPolicyName *string `json:"AlarmPolicyName,omitempty" name:"AlarmPolicyName"`

	// 告警类型，有cluster，node，pod三种类型。
	AlarmPolicyType *string `json:"AlarmPolicyType,omitempty" name:"AlarmPolicyType"`

	// 告警指标列表。
	AlarmMetrics []*AlarmMetric `json:"AlarmMetrics,omitempty" name:"AlarmMetrics" list`

	// 告警对象，多个用逗号分隔。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmObjects *string `json:"AlarmObjects,omitempty" name:"AlarmObjects"`

	// 告警策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmPolicyDescription *string `json:"AlarmPolicyDescription,omitempty" name:"AlarmPolicyDescription"`

	// 告警对象绑定类型。
	// all：全部对象绑定。
	// part：只有选择的部分对象绑定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmObjectsType *string `json:"AlarmObjectsType,omitempty" name:"AlarmObjectsType"`
}

type AutoScalingGroupRange struct {

	// 伸缩组最小实例数
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// 伸缩组最大实例数
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type AvailableExtraArgs struct {

	// kube-apiserver可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeAPIServer []*Flag `json:"KubeAPIServer,omitempty" name:"KubeAPIServer" list`

	// kube-controller-manager可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeControllerManager []*Flag `json:"KubeControllerManager,omitempty" name:"KubeControllerManager" list`

	// kube-scheduler可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeScheduler []*Flag `json:"KubeScheduler,omitempty" name:"KubeScheduler" list`

	// kubelet可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kubelet []*Flag `json:"Kubelet,omitempty" name:"Kubelet" list`
}

type CcnInstance struct {

	// 云联网实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 关联实例类型：
	// VPC：私有网络
	// DIRECTCONNECT：专线网关
	// BMVPC：黑石私有网络
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 关联实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 关联实例所属大区，例如：ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`

	// 关联实例所属UIN（根账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`

	// 关联实例CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidrs []*string `json:"Cidrs,omitempty" name:"Cidrs" list`

	// 关联实例状态：
	// PENDING：申请中
	// ACTIVE：已连接
	// EXPIRED：已过期
	// REJECTED：已拒绝
	// DELETED：已删除
	// FAILED：失败的（2小时后将异步强制解关联）
	// ATTACHING：关联中
	// DETACHING：解关联中
	// DETACHFAILED：解关联失败（2小时后将异步强制解关联）
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// 云联网所属UIN（根账号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
}

type CcnRoute struct {

	// 路由策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// 目的端
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// 下一跳类型（关联实例类型），所有类型：VPC、DIRECTCONNECT
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 下一跳（关联实例）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 下一跳所属地域（关联实例所属地域）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`

	// 关联实例所属UIN（根账号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
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

type CheckClusterHostNameRequest struct {
	*tchttp.BaseRequest

	// 节点主机名称列表
	HostNames []*HostNameValue `json:"HostNames,omitempty" name:"HostNames" list`

	// 集群ID(如果往已经存在的集群内加节点，则需要传对应的集群ID)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CheckClusterHostNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckClusterHostNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckClusterHostNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点主机名校验是否通过(true: 通过 false: 不通过 )
		BEnable *bool `json:"BEnable,omitempty" name:"BEnable"`

		// 节点主机名校验不通过的原因
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterHostNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckClusterHostNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckClusterImageRequest struct {
	*tchttp.BaseRequest

	// 指定有效的镜像ID，格式形如img-xxxx。可通过登录控制台查询，也可调用接口 DescribeImages，取返回信息中的ImageId字段。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *CheckClusterImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckClusterImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckClusterImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否支持设置为集群镜像
		Suitable *bool `json:"Suitable,omitempty" name:"Suitable"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckClusterImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckInstancesUpgradeAbleRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 节点列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 升级类型
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`
}

func (r *CheckInstancesUpgradeAbleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckInstancesUpgradeAbleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckInstancesUpgradeAbleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群master当前小版本
		ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

		// 集群master对应的大版本目前最新小版本
		LatestVersion *string `json:"LatestVersion,omitempty" name:"LatestVersion"`

		// 可升级节点列表
		UpgradeAbleInstances []*UpgradeAbleInstancesItem `json:"UpgradeAbleInstances,omitempty" name:"UpgradeAbleInstances" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckInstancesUpgradeAbleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckInstancesUpgradeAbleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckLogCollectorHostPathRequest struct {
	*tchttp.BaseRequest

	// 路径名
	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *CheckLogCollectorHostPathRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckLogCollectorHostPathRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckLogCollectorHostPathResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主机路径是否合法
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckLogCollectorHostPathResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckLogCollectorHostPathResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckLogCollectorNameRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 日志采集规则名称
	LogCollectorName *string `json:"LogCollectorName,omitempty" name:"LogCollectorName"`
}

func (r *CheckLogCollectorNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckLogCollectorNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckLogCollectorNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志采集规则是否存在（true：存在，false：不存在）
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckLogCollectorNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckLogCollectorNameResponse) FromJsonString(s string) error {
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

type ClusterCondition struct {

	// 集群创建过程类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 集群创建过程状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 最后一次探测到该状态的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastProbeTime *string `json:"LastProbeTime,omitempty" name:"LastProbeTime"`

	// 最后一次转换到该过程的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTransitionTime *string `json:"LastTransitionTime,omitempty" name:"LastTransitionTime"`

	// 转换到该过程的简明原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 转换到该过程的更多信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type ClusterCredential struct {

	// CA 根证书
	CACert *string `json:"CACert,omitempty" name:"CACert"`

	// 认证用的Token
	Token *string `json:"Token,omitempty" name:"Token"`
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

type ClusterHealthyPodsStatus struct {

	// pod总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// NotReady的pod总数
	NotReadyTotal *int64 `json:"NotReadyTotal,omitempty" name:"NotReadyTotal"`

	// NotReady的pod详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotReadyPods []*NotReadyPodsItem `json:"NotReadyPods,omitempty" name:"NotReadyPods" list`
}

type ClusterInspectionOverview struct {

	// 本次检查的id，用于与其他检测结果区分
	Id *string `json:"Id,omitempty" name:"Id"`

	// 检查的时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 如果检查异常终止，则该字段包含错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitempty" name:"Error"`

	// 正常项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodItem *int64 `json:"GoodItem,omitempty" name:"GoodItem"`

	// 警告项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnItem *int64 `json:"WarnItem,omitempty" name:"WarnItem"`

	// 高风险项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskItem *int64 `json:"RiskItem,omitempty" name:"RiskItem"`

	// 严重项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SeriousItem *int64 `json:"SeriousItem,omitempty" name:"SeriousItem"`

	// 检测失败项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedItem *int64 `json:"FailedItem,omitempty" name:"FailedItem"`
}

type ClusterInspectionProgress struct {

	// 当前步骤名称：
	// init_env: 初始化环境，安装agent。
	// init_k8s_resources：获取kubernetes资源
	// init_components: 获取核心组件参数
	// init_machines：获取节点系统参数
	// diagnostic： 正在检查集群
	Step *string `json:"Step,omitempty" name:"Step"`

	// 检查进度百分比
	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
}

type ClusterInstancesVersion struct {

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 节点版本统计信息
	InstanceVersions []*ClusterInstancesVersionItem `json:"InstanceVersions,omitempty" name:"InstanceVersions" list`

	// 出错信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitempty" name:"Error"`

	// 是否存在可升级节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeAble *bool `json:"UpgradeAble,omitempty" name:"UpgradeAble"`
}

type ClusterInstancesVersionItem struct {

	// 版本
	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`

	// 节点数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 是否可升级
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeAble *bool `json:"UpgradeAble,omitempty" name:"UpgradeAble"`
}

type ClusterInternalLB struct {

	// 是否开启内网访问LB
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 内网访问LB关联的子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
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

type ClusterPublicLB struct {

	// 是否开启公网访问LB
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 允许访问的来源CIDR列表
	AllowFromCidrs []*string `json:"AllowFromCidrs,omitempty" name:"AllowFromCidrs" list`
}

type ClusterStatus struct {

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群状态
	ClusterState *string `json:"ClusterState,omitempty" name:"ClusterState"`

	// 集群下机器实例的状态
	ClusterInstanceState *string `json:"ClusterInstanceState,omitempty" name:"ClusterInstanceState"`

	// 集群是否开启监控
	ClusterBMonitor *bool `json:"ClusterBMonitor,omitempty" name:"ClusterBMonitor"`

	// 集群创建中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	ClusterInitNodeNum *int64 `json:"ClusterInitNodeNum,omitempty" name:"ClusterInitNodeNum"`

	// 集群运行中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	ClusterRunningNodeNum *int64 `json:"ClusterRunningNodeNum,omitempty" name:"ClusterRunningNodeNum"`

	// 集群异常的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	ClusterFailedNodeNum *int64 `json:"ClusterFailedNodeNum,omitempty" name:"ClusterFailedNodeNum"`

	// 集群已关机的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterClosedNodeNum *int64 `json:"ClusterClosedNodeNum,omitempty" name:"ClusterClosedNodeNum"`

	// 集群关机中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterClosingNodeNum *int64 `json:"ClusterClosingNodeNum,omitempty" name:"ClusterClosingNodeNum"`
}

type ContainerStatus struct {

	// 容器名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 容器ID
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 容器状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 容器处于该状态的原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 容器镜像ID
	Image *string `json:"Image,omitempty" name:"Image"`
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

type CreateClusterEndpointRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否为外网访问（默认值: False 不为外网访问）
	BExtranet *bool `json:"BExtranet,omitempty" name:"BExtranet"`

	// 集群端口所在的子网ID  (仅在开启非外网访问时需要填，必须为集群所在VPC内的子网)
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *CreateClusterEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterEndpointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterEndpointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointVipRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)
	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies" list`
}

func (r *CreateClusterEndpointVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterEndpointVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求任务的FlowId
		RequestFlowId *int64 `json:"RequestFlowId,omitempty" name:"RequestFlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterEndpointVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterEndpointVipResponse) FromJsonString(s string) error {
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

type CreateClusterRouteRequest struct {
	*tchttp.BaseRequest

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 目的端CIDR。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// 下一跳地址。
	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
}

func (r *CreateClusterRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRouteTableRequest struct {
	*tchttp.BaseRequest

	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 路由表CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// 路由表绑定的VPC
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 是否忽略CIDR冲突
	IgnoreClusterCidrConflict *int64 `json:"IgnoreClusterCidrConflict,omitempty" name:"IgnoreClusterCidrConflict"`
}

func (r *CreateClusterRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRouteTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEKSClusterRequest struct {
	*tchttp.BaseRequest

	// k8s版本号。可为1.14.4, 1.12.8。
	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`

	// vpc 的Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 子网Id 列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// Serivce 所在子网Id
	ServiceSubnetId *string `json:"ServiceSubnetId,omitempty" name:"ServiceSubnetId"`
}

func (r *CreateEKSClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEKSClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEKSClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 弹性集群Id
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEKSClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEKSClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateIndependentClusterRequest struct {
	*tchttp.BaseRequest

	// CVM创建透传参数，json化字符串格式, 各种角色的节点配置信息
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode" list`

	// 集群容器网络配置信息
	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`

	// 集群的基本配置信息
	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`

	// 集群高级配置信息
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`

	// 节点高级配置信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
}

func (r *CreateIndependentClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIndependentClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateIndependentClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIndependentClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIndependentClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTKEClusterRequest struct {
	*tchttp.BaseRequest

	// 集群容器网络配置信息
	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`

	// 是否独立集群
	IndependentCluster *bool `json:"IndependentCluster,omitempty" name:"IndependentCluster"`

	// CVM创建透传参数，json化字符串格式, 各种角色的节点配置信息
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode" list`

	// 集群的基本配置信息
	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`

	// 集群高级配置信息
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`

	// 节点高级配置信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
}

func (r *CreateTKEClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTKEClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTKEClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTKEClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTKEClusterResponse) FromJsonString(s string) error {
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

type DeleteAlarmPoliciesRequest struct {
	*tchttp.BaseRequest

	// k8s集群ID
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`

	// 告警策略ID列表
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitempty" name:"AlarmPolicyIds" list`
}

func (r *DeleteAlarmPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAlarmPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type DeleteClusterCIDRFromCcnRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *DeleteClusterCIDRFromCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterCIDRFromCcnRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterCIDRFromCcnResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云联网路由数组
		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterCIDRFromCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterCIDRFromCcnResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterCIDRFromVbcRequest struct {
	*tchttp.BaseRequest

	// vpc唯一id
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 集群cidr
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *DeleteClusterCIDRFromVbcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterCIDRFromVbcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterCIDRFromVbcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 详情
		Detail []*CcnRoute `json:"Detail,omitempty" name:"Detail" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterCIDRFromVbcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterCIDRFromVbcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterEndpointRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否为外网访问（默认值: False 不为外网访问）
	BExtranet *bool `json:"BExtranet,omitempty" name:"BExtranet"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *DeleteClusterEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterEndpointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterEndpointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterEndpointVipRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterEndpointVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterEndpointVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterEndpointVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterEndpointVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterEndpointVipResponse) FromJsonString(s string) error {
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

type DeleteClusterRouteRequest struct {
	*tchttp.BaseRequest

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 下一跳地址。
	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`

	// 目的端CIDR。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
}

func (r *DeleteClusterRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRouteTableRequest struct {
	*tchttp.BaseRequest

	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *DeleteClusterRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRouteTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEKSClusterRequest struct {
	*tchttp.BaseRequest

	// 弹性集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteEKSClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEKSClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEKSClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEKSClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEKSClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPoliciesRequest struct {
	*tchttp.BaseRequest

	// k8s集群ID
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filter *AlarmPolicyFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeAlarmPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群告警策略总个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群告警列表
		AlarmPolicySet []*AlarmPolicy `json:"AlarmPolicySet,omitempty" name:"AlarmPolicySet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableClusterVersionRequest struct {
	*tchttp.BaseRequest

	// 集群 Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群 Id 列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`
}

func (r *DescribeAvailableClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableClusterVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableClusterVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 关联网络实例列表
		InstanceSet []*CcnInstance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *DescribeCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云联网路由数组
		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogSetsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClsLogSetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClsLogSetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogSetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志集数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志集集合
		Logsets []*LogSet `json:"Logsets,omitempty" name:"Logsets" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClsLogSetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClsLogSetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogTopicsRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
}

func (r *DescribeClsLogTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClsLogTopicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志主题的数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志主题集合
		Topics []*LogTopic `json:"Topics,omitempty" name:"Topics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClsLogTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClsLogTopicsResponse) FromJsonString(s string) error {
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

type DescribeClusterAvailableExtraArgsRequest struct {
	*tchttp.BaseRequest

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 集群类型(MANAGED_CLUSTER或INDEPENDENT_CLUSTER)
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

func (r *DescribeClusterAvailableExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterAvailableExtraArgsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAvailableExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群版本
		ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

		// 可用的自定义参数
		AvailableExtraArgs *AvailableExtraArgs `json:"AvailableExtraArgs,omitempty" name:"AvailableExtraArgs"`

		// 集群类型
		ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAvailableExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterAvailableExtraArgsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterCreateProgressRequest struct {
	*tchttp.BaseRequest

	// 集群 Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterCreateProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterCreateProgressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterCreateProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建进度
		Progress []*string `json:"Progress,omitempty" name:"Progress" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterCreateProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterCreateProgressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否为外网访问（默认值: False 不为外网访问）
	BExtranet *bool `json:"BExtranet,omitempty" name:"BExtranet"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *DescribeClusterEndpointStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterEndpointStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询集群访问端口状态（Created 开启成功，Creating 开启中中，NotFound 未开启）
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterEndpointStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointVipStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterEndpointVipStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterEndpointVipStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointVipStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 端口操作状态 (Creating 创建中  CreateFailed 创建失败 Created 创建完成 Deleting 删除中 DeletedFailed 删除失败 Deleted 已删除 NotFound 未发现操作 )
		Status *string `json:"Status,omitempty" name:"Status"`

		// 操作失败的原因
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointVipStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterEndpointVipStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterExtraArgsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterExtraArgsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群自定义参数
		ClusterExtraArgs *ClusterExtraArgs `json:"ClusterExtraArgs,omitempty" name:"ClusterExtraArgs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterExtraArgsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterHealthyStatusRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterHealthyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterHealthyStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterHealthyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群Pod健康状态
		PodsStatus *ClusterHealthyPodsStatus `json:"PodsStatus,omitempty" name:"PodsStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterHealthyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterHealthyStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionItem struct {

	// 集群实例id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 如果发生错误，则该字段包含错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitempty" name:"Error"`

	// 如果当前正在巡检，则该字段包含检测巡检
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *ClusterInspectionProgress `json:"Progress,omitempty" name:"Progress"`

	// 最近一次巡检结果概览
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastResult *ClusterInspectionOverview `json:"LastResult,omitempty" name:"LastResult"`

	// 自动巡检周期，crontab格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cron *string `json:"Cron,omitempty" name:"Cron"`
}

type DescribeClusterInspectionOverviewsRequest struct {
	*tchttp.BaseRequest

	// 集群实例id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最多返回多少条记录
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 从最近第几条记录开始返回
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeClusterInspectionOverviewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInspectionOverviewsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionOverviewsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群巡检报告列表
		OverviewSet []*ClusterInspectionOverview `json:"OverviewSet,omitempty" name:"OverviewSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInspectionOverviewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInspectionOverviewsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionReportRequest struct {
	*tchttp.BaseRequest

	// 集群实例id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 报告页id
	ReportId *string `json:"ReportId,omitempty" name:"ReportId"`
}

func (r *DescribeClusterInspectionReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInspectionReportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInspectionReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInspectionReportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionsRequest struct {
	*tchttp.BaseRequest

	// 目标集群实例id数组
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`
}

func (r *DescribeClusterInspectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInspectionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 各集群巡检概览
		ResultSet []*DescribeClusterInspectionItem `json:"ResultSet,omitempty" name:"ResultSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInspectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInspectionsResponse) FromJsonString(s string) error {
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

	// 过滤条件列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
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

type DescribeClusterPodsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Deployment名称
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 命名空间名称，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 过滤参数，从PodIP、NodeIP、PodName中过滤含有该参数的Pod实例
	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`

	// 最大输出条数，默认0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeClusterPodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterPodsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterPodsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的Pod总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Pod信息
		Pods []*PodInfo `json:"Pods,omitempty" name:"Pods" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterPodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterPodsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRouteTablesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClusterRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRouteTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群路由表对象。
		RouteTableSet []*RouteTableInfo `json:"RouteTableSet,omitempty" name:"RouteTableSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRouteTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *DescribeClusterRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群路由对象。
		RouteSet []*RouteInfo `json:"RouteSet,omitempty" name:"RouteSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSecurityRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否返回内网地址(默认: FALSE)
	JnsGwEndpointEnable *bool `json:"JnsGwEndpointEnable,omitempty" name:"JnsGwEndpointEnable"`
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

type DescribeClusterStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表，不传默认拉取所有集群
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`
}

func (r *DescribeClusterStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群状态列表
		ClusterStatusSet []*ClusterStatus `json:"ClusterStatusSet,omitempty" name:"ClusterStatusSet" list`

		// 集群个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterStatusResponse) FromJsonString(s string) error {
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

type DescribeClustersResourceStatusRequest struct {
	*tchttp.BaseRequest

	// 资源维度，可传列表
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 集群id，可传列表
	ClusterInstancesIds []*string `json:"ClusterInstancesIds,omitempty" name:"ClusterInstancesIds" list`

	// 获取资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeClustersResourceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClustersResourceStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResourceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群资源状态集合
		ResourceStatusSet []*ResourceStatus `json:"ResourceStatusSet,omitempty" name:"ResourceStatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResourceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClustersResourceStatusResponse) FromJsonString(s string) error {
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

type DescribeEKSClusterCredentialRequest struct {
	*tchttp.BaseRequest

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeEKSClusterCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEKSClusterCredentialRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEKSClusterCredentialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群的接入地址信息
		Addresses []*IPAddress `json:"Addresses,omitempty" name:"Addresses" list`

		// 集群的认证信息
		Credential *ClusterCredential `json:"Credential,omitempty" name:"Credential"`

		// 集群的公网访问信息
		PublicLB *ClusterPublicLB `json:"PublicLB,omitempty" name:"PublicLB"`

		// 集群的内网访问信息
		InternalLB *ClusterInternalLB `json:"InternalLB,omitempty" name:"InternalLB"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEKSClusterCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEKSClusterCredentialResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEKSClusterStatusRequest struct {
	*tchttp.BaseRequest

	// 弹性容器集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeEKSClusterStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEKSClusterStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEKSClusterStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群当前状态
		Phase *string `json:"Phase,omitempty" name:"Phase"`

		// 集群过程数组
		Conditions []*ClusterCondition `json:"Conditions,omitempty" name:"Conditions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEKSClusterStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEKSClusterStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEKSClustersRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件,当前只支持按照单个条件ClusterName进行过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeEKSClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEKSClustersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEKSClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群总个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群信息列表
		Clusters []*EksCluster `json:"Clusters,omitempty" name:"Clusters" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEKSClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEKSClustersResponse) FromJsonString(s string) error {
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

type DescribeFlowIdStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 开启集群外网访问端口的任务ID
	RequestFlowId *int64 `json:"RequestFlowId,omitempty" name:"RequestFlowId"`
}

func (r *DescribeFlowIdStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowIdStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowIdStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务错误信息
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowIdStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowIdStatusResponse) FromJsonString(s string) error {
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

type DescribeInstanceCreateProgressRequest struct {
	*tchttp.BaseRequest

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 节点实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceCreateProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceCreateProgressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceCreateProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建进度
		Progress []*string `json:"Progress,omitempty" name:"Progress" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceCreateProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceCreateProgressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesVersionRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表
	ClusterId []*string `json:"ClusterId,omitempty" name:"ClusterId" list`
}

func (r *DescribeInstancesVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 不同版本节点个数统计结果
		Clusters []*ClusterInstancesVersion `json:"Clusters,omitempty" name:"Clusters" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesVersionResponse) FromJsonString(s string) error {
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

type DescribeRouteTableConflictsRequest struct {
	*tchttp.BaseRequest

	// 路由表CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// 路由表绑定的VPC
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeRouteTableConflictsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTableConflictsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTableConflictsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由表是否冲突。
		HasConflict *bool `json:"HasConflict,omitempty" name:"HasConflict"`

		// 路由表冲突列表。
		RouteTableConflictSet []*RouteTableConflict `json:"RouteTableConflictSet,omitempty" name:"RouteTableConflictSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTableConflictsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTableConflictsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUpgradeClusterProgressRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeUpgradeClusterProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUpgradeClusterProgressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUpgradeClusterProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群升级整体进度
		Steps *TaskStepInfo `json:"Steps,omitempty" name:"Steps"`

		// 每个节点的进度
		Instances *TaskStepInfo `json:"Instances,omitempty" name:"Instances"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpgradeClusterProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUpgradeClusterProgressResponse) FromJsonString(s string) error {
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

	// 是否只是拉取列表
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
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

type EksCluster struct {

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// k8s 版本号
	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`

	// 集群状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Service 子网Id
	ServiceSubnetId *string `json:"ServiceSubnetId,omitempty" name:"ServiceSubnetId"`
}

type EnableVpcPeerClusterRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 对等连接实例ID
	VpcPeerId *string `json:"VpcPeerId,omitempty" name:"VpcPeerId"`
}

func (r *EnableVpcPeerClusterRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableVpcPeerClusterRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableVpcPeerClusterRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableVpcPeerClusterRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableVpcPeerClusterRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type Evaluator struct {

	// 告警判断类型，目前支持gt和lt两种
	Type *string `json:"Type,omitempty" name:"Type"`

	// 告警设置的阈值
	Value *float64 `json:"Value,omitempty" name:"Value"`
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

type Flag struct {

	// 参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数描述
	Usage *string `json:"Usage,omitempty" name:"Usage"`

	// 参数默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数可选范围（目前包含range和in两种，"[]"代表range，如"[1, 5]"表示参数必须>=1且 <=5, "()"代表in， 如"('aa', 'bb')"表示参数只能为字符串'aa'或者'bb'，该参数为空表示不校验）
	Constraint *string `json:"Constraint,omitempty" name:"Constraint"`
}

type ForwardClusterRequestRequest struct {
	*tchttp.BaseRequest

	// 请求集群资源的访问
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求集群资源的路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求集群资源后允许接收的数据格式
	Accept *string `json:"Accept,omitempty" name:"Accept"`

	// 请求集群资源的数据格式
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 请求集群资源的数据
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 是否编码请求内容
	EncodedBody *bool `json:"EncodedBody,omitempty" name:"EncodedBody"`
}

func (r *ForwardClusterRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForwardClusterRequestRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForwardClusterRequestResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求集群资源后返回的数据
		ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForwardClusterRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForwardClusterRequestResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type GetUpgradeClusterProgressRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetUpgradeClusterProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUpgradeClusterProgressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeClusterProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群升级整体进度
		Steps *TaskStepInfo `json:"Steps,omitempty" name:"Steps"`

		// 每个节点的进度
		Instances *TaskStepInfo `json:"Instances,omitempty" name:"Instances"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeClusterProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUpgradeClusterProgressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeInstanceProgressRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最多获取多少个节点的进度
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 从第几个节点开始获取进度
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetUpgradeInstanceProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUpgradeInstanceProgressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeInstanceProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 升级节点总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 已升级节点总数
		Done *int64 `json:"Done,omitempty" name:"Done"`

		// 升级任务生命周期
	// process 运行中
	// paused 已停止
	// pauing 正在停止
	// done  已完成
	// timeout 已超时
	// aborted 已取消
		LifeState *string `json:"LifeState,omitempty" name:"LifeState"`

		// 各节点升级进度详情
		Instances []*InstanceUpgradeProgressItem `json:"Instances,omitempty" name:"Instances" list`

		// 集群当前状态
		ClusterStatus *InstanceUpgradeClusterStatus `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeInstanceProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUpgradeInstanceProgressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetVbcInstanceRequest struct {
	*tchttp.BaseRequest

	// vpc唯一id
	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
}

func (r *GetVbcInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetVbcInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetVbcInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 详情
		Detail []*CcnRoute `json:"Detail,omitempty" name:"Detail" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetVbcInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetVbcInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetVbcRouteRequest struct {
	*tchttp.BaseRequest

	// vpc唯一id
	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`

	// 集群cidr
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *GetVbcRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetVbcRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetVbcRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 详情
		Detail []*CcnRoute `json:"Detail,omitempty" name:"Detail" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetVbcRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetVbcRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HostNameValue struct {

	// 主机名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 主机数量
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type IPAddress struct {

	// Ip 地址的类型。可为 advertise, public 等
	Type *string `json:"Type,omitempty" name:"Type"`

	// Ip 地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 网络端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`
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

type InstanceUpgradeClusterStatus struct {

	// pod总数
	PodTotal *int64 `json:"PodTotal,omitempty" name:"PodTotal"`

	// NotReady pod总数
	NotReadyPod *int64 `json:"NotReadyPod,omitempty" name:"NotReadyPod"`
}

type InstanceUpgradePreCheckResult struct {

	// 检查是否通过
	CheckPass *bool `json:"CheckPass,omitempty" name:"CheckPass"`

	// 检查项数组
	Items []*InstanceUpgradePreCheckResultItem `json:"Items,omitempty" name:"Items" list`

	// 本节点独立pod列表
	SinglePods []*string `json:"SinglePods,omitempty" name:"SinglePods" list`
}

type InstanceUpgradePreCheckResultItem struct {

	// 工作负载的命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 工作负载类型
	WorkLoadKind *string `json:"WorkLoadKind,omitempty" name:"WorkLoadKind"`

	// 工作负载名称
	WorkLoadName *string `json:"WorkLoadName,omitempty" name:"WorkLoadName"`

	// 驱逐节点前工作负载running的pod数目
	Before *uint64 `json:"Before,omitempty" name:"Before"`

	// 驱逐节点后工作负载running的pod数目
	After *uint64 `json:"After,omitempty" name:"After"`

	// 工作负载在本节点上的pod列表
	Pods []*string `json:"Pods,omitempty" name:"Pods" list`
}

type InstanceUpgradeProgressItem struct {

	// 节点instanceID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 任务生命周期
	// process 运行中
	// paused 已停止
	// pauing 正在停止
	// done  已完成
	// timeout 已超时
	// aborted 已取消
	// pending 还未开始
	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`

	// 升级开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// 升级结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`

	// 升级前检查结果
	CheckResult *InstanceUpgradePreCheckResult `json:"CheckResult,omitempty" name:"CheckResult"`

	// 升级步骤详情
	Detail []*TaskStepInfo `json:"Detail,omitempty" name:"Detail" list`
}

type Label struct {

	// map表中的Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// map表中的Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogSet struct {

	// 日志集ID
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志集名称
	LogSetName []*string `json:"LogSetName,omitempty" name:"LogSetName" list`

	// 周期
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type LogTopic struct {

	// 日志集ID
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 日志类型
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 是否采集
	Collection *bool `json:"Collection,omitempty" name:"Collection"`

	// 是否有索引
	Index *bool `json:"Index,omitempty" name:"Index"`
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ModifyAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// k8s集群ID
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`

	// 告警策略ID
	AlarmPolicyId *string `json:"AlarmPolicyId,omitempty" name:"AlarmPolicyId"`

	// k8s命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// k8s工作负载类型，有deplogyment，job，crontab
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 告警策略settings
	AlarmPolicySettings *AlarmPolicySettings `json:"AlarmPolicySettings,omitempty" name:"AlarmPolicySettings"`

	// 告警通知settings
	NotifySettings *NotifySettings `json:"NotifySettings,omitempty" name:"NotifySettings"`

	// 告警屏蔽settings
	ShieldSettings *ShieldSettings `json:"ShieldSettings,omitempty" name:"ShieldSettings"`
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

type ModifyClusterEndpointSPRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)
	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies" list`
}

func (r *ModifyClusterEndpointSPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterEndpointSPRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterEndpointSPResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterEndpointSPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterEndpointSPResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterImageRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 指定有效的镜像ID，格式形如img-xxxx。可通过登录控制台查询，也可调用接口 DescribeImages，取返回信息中的ImageId字段。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// "DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`
}

func (r *ModifyClusterImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterInspectionRequest struct {
	*tchttp.BaseRequest

	// 集群实例id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 自动巡检周期
	Cron *string `json:"Cron,omitempty" name:"Cron"`
}

func (r *ModifyClusterInspectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterInspectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterInspectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterInspectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterInspectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterUpgradingStateRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 操作类型：
	// pause 暂停集群升级
	// resume 继续集群升级
	// abort 取消集群升级
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 如果是取消集群升级，该参数设置为true表示强制取消
	ForceAbort *bool `json:"ForceAbort,omitempty" name:"ForceAbort"`
}

func (r *ModifyClusterUpgradingStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterUpgradingStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterUpgradingStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterUpgradingStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterUpgradingStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NodePoolOption struct {

	// 是否加入节点池
	AddToNodePool *bool `json:"AddToNodePool,omitempty" name:"AddToNodePool"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
}

type NotReadyPodsItem struct {

	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pods []*string `json:"Pods,omitempty" name:"Pods" list`
}

type NotifySettings struct {

	// 告警接收组（用户组）
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups" list`

	// 告警通知方式。目前有SMS、EMAIL、CALL、WECHAT方式。
	// 分别代表：短信、邮件、电话、微信
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay" list`

	// 电话告警顺序。
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNotifyOrder []*int64 `json:"PhoneNotifyOrder,omitempty" name:"PhoneNotifyOrder" list`

	// 电话告警次数。
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitempty" name:"PhoneCircleTimes"`

	// 电话告警轮内间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitempty" name:"PhoneInnerInterval"`

	// 电话告警轮外间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitempty" name:"PhoneCircleInterval"`

	// 电话告警触达通知
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneArriveNotice *int64 `json:"PhoneArriveNotice,omitempty" name:"PhoneArriveNotice"`
}

type OpUpgradeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 节点实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 操作类型
	// pause 暂停任务（任务必须处于运行中）
	// resume 恢复任务（任务必须处于已暂停状态）
	// abort 停止任务（任务必须处于已暂停状态）
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *OpUpgradeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpUpgradeClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpUpgradeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpUpgradeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpUpgradeClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PauseClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *PauseClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PauseClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PauseClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PauseClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PodInfo struct {

	// Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Pod当前状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 处于当前状态原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 在kubernetes中展示的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceReason *string `json:"SourceReason,omitempty" name:"SourceReason"`

	// Pod IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Pod重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// Pod就绪容器数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadyCount *int64 `json:"ReadyCount,omitempty" name:"ReadyCount"`

	// 所在节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 所在节点IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`

	// Pod启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Pod所含容器信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Containers []*ContainerStatus `json:"Containers,omitempty" name:"Containers" list`
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

type ResourceStatus struct {

	// 集群ID
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`

	// 各资源状态
	Status []*ResourceStatusItem `json:"Status,omitempty" name:"Status" list`
}

type ResourceStatusItem struct {

	// 维度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 总数
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// 异常数
	AbnormalNum *uint64 `json:"AbnormalNum,omitempty" name:"AbnormalNum"`

	// 异常详情
	AbnormalDetail []*AbnormalDetail `json:"AbnormalDetail,omitempty" name:"AbnormalDetail" list`
}

type ResumeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ResumeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RouteInfo struct {

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 目的端CIDR。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// 下一跳地址。
	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
}

type RouteTableConflict struct {

	// 路由表类型。
	RouteTableType *string `json:"RouteTableType,omitempty" name:"RouteTableType"`

	// 路由表CIDR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// 路由表名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 路由表ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type RouteTableInfo struct {

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 路由表CIDR。
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type RunClusterInspectionResponseItem struct {

	// 集群实例id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 如果请求未能正常被处理，则Error中将包含错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitempty" name:"Error"`
}

type RunClusterInspectionsRequest struct {
	*tchttp.BaseRequest

	// 目标集群列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`
}

func (r *RunClusterInspectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunClusterInspectionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunClusterInspectionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 每个集群的处理结果
		ResultSet []*RunClusterInspectionResponseItem `json:"ResultSet,omitempty" name:"ResultSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunClusterInspectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunClusterInspectionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type ShieldSettings struct {

	// 是否启动告警屏蔽功能
	EnableShield *bool `json:"EnableShield,omitempty" name:"EnableShield"`

	// 告警屏蔽开始时间，单位s，如8:00:00=> 8 * 60 * 60=28800
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShieldTimeStart *int64 `json:"ShieldTimeStart,omitempty" name:"ShieldTimeStart"`

	// 告警屏蔽结束时间，单位s，如10:00:00=> 10 * 60 * 60=36000
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShieldTimeEnd *int64 `json:"ShieldTimeEnd,omitempty" name:"ShieldTimeEnd"`
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

type TaskStepInfo struct {

	// 步骤名称
	Step *string `json:"Step,omitempty" name:"Step"`

	// 生命周期
	// pending : 步骤未开始
	// running: 步骤执行中
	// success: 步骤成功完成
	// failed: 步骤失败
	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`

	// 步骤开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// 步骤结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`

	// 若步骤生命周期为failed,则此字段显示错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedMsg *string `json:"FailedMsg,omitempty" name:"FailedMsg"`
}

type UpdateClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 需要升级的集群的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 升级的节点ID列表，请选择集群中需要升级的节点ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 描述了实例创建相关的一些高级设置。
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 描述了实例登录相关配置与信息。
	LoginSettings *string `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 描述了实例的增强服务启用情况与其设置，如云安全，云监控等实例 Agent。
	EnhancedService *string `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *UpdateClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功启动更新的节点Id列表
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterVersionRequest struct {
	*tchttp.BaseRequest

	// 集群 Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要升级到的版本
	DstVersion *string `json:"DstVersion,omitempty" name:"DstVersion"`

	// 集群自定义参数
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// 可容忍的最大不可用pod数目
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitempty" name:"MaxNotReadyPercent"`

	// 是否跳过预检查阶段
	SkipPreCheck *bool `json:"SkipPreCheck,omitempty" name:"SkipPreCheck"`
}

func (r *UpdateClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateClusterVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateClusterVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateEKSClusterRequest struct {
	*tchttp.BaseRequest

	// 弹性集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 弹性集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 弹性集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 子网Id 列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 弹性容器集群公网访问LB信息
	PublicLB *ClusterPublicLB `json:"PublicLB,omitempty" name:"PublicLB"`

	// 弹性容器集群内网访问LB信息
	InternalLB *ClusterInternalLB `json:"InternalLB,omitempty" name:"InternalLB"`

	// Service 子网Id
	ServiceSubnetId *string `json:"ServiceSubnetId,omitempty" name:"ServiceSubnetId"`
}

func (r *UpdateEKSClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateEKSClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateEKSClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateEKSClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateEKSClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeAbleInstancesItem struct {

	// 节点Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点的当前版本
	Version *string `json:"Version,omitempty" name:"Version"`
}

type UpgradeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// create 表示开始一次升级任务
	// pause 表示停止任务
	// resume表示继续任务
	// abort表示终止任务
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 升级类型，目前只支持reset
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// 需要升级的节点列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 当节点重新加入集群时候所使用的参数，参考添加已有节点接口
	ResetParam *UpgradeNodeResetParam `json:"ResetParam,omitempty" name:"ResetParam"`

	// 是否忽略节点升级前检查
	SkipPreCheck *bool `json:"SkipPreCheck,omitempty" name:"SkipPreCheck"`

	// 最大可容忍的不可用Pod比例
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitempty" name:"MaxNotReadyPercent"`
}

func (r *UpgradeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeNodeResetParam struct {

	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *string `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）
	LoginSettings *string `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。（目前仅支持设置单个sgId）
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
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
