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
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCheckClusterNameRequest() (request *CheckClusterNameRequest) {
    request = &CheckClusterNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "CheckClusterName")
    return
}

func NewCheckClusterNameResponse() (response *CheckClusterNameResponse) {
    response = &CheckClusterNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于检查要设置的ClusterName是否合法，是否重复等
func (c *Client) CheckClusterName(request *CheckClusterNameRequest) (response *CheckClusterNameResponse, err error) {
    if request == nil {
        request = NewCheckClusterNameRequest()
    }
    response = NewCheckClusterNameResponse()
    err = c.Send(request, response)
    return
}

func NewCheckForceRestartRequest() (request *CheckForceRestartRequest) {
    request = &CheckForceRestartRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "CheckForceRestart")
    return
}

func NewCheckForceRestartResponse() (response *CheckForceRestartResponse) {
    response = &CheckForceRestartResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查集群是否需要强制重启
func (c *Client) CheckForceRestart(request *CheckForceRestartRequest) (response *CheckForceRestartResponse, err error) {
    if request == nil {
        request = NewCheckForceRestartRequest()
    }
    response = NewCheckForceRestartResponse()
    err = c.Send(request, response)
    return
}

func NewCheckOperationRequest() (request *CheckOperationRequest) {
    request = &CheckOperationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "CheckOperation")
    return
}

func NewCheckOperationResponse() (response *CheckOperationResponse) {
    response = &CheckOperationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对涉及缩容或重启的操作，检查实例是否允许操作
func (c *Client) CheckOperation(request *CheckOperationRequest) (response *CheckOperationResponse, err error) {
    if request == nil {
        request = NewCheckOperationRequest()
    }
    response = NewCheckOperationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckScaleUpgradeRequest() (request *CheckScaleUpgradeRequest) {
    request = &CheckScaleUpgradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "CheckScaleUpgrade")
    return
}

func NewCheckScaleUpgradeResponse() (response *CheckScaleUpgradeResponse) {
    response = &CheckScaleUpgradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 扩缩容方式版本升级情况下，检查集群状态（节点个数、索引存储量等）是否满足升级条件
func (c *Client) CheckScaleUpgrade(request *CheckScaleUpgradeRequest) (response *CheckScaleUpgradeResponse, err error) {
    if request == nil {
        request = NewCheckScaleUpgradeRequest()
    }
    response = NewCheckScaleUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewCheckUpdateInstanceRequest() (request *CheckUpdateInstanceRequest) {
    request = &CheckUpdateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "CheckUpdateInstance")
    return
}

func NewCheckUpdateInstanceResponse() (response *CheckUpdateInstanceResponse) {
    response = &CheckUpdateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查实例变配操作是否可以发起
func (c *Client) CheckUpdateInstance(request *CheckUpdateInstanceRequest) (response *CheckUpdateInstanceResponse, err error) {
    if request == nil {
        request = NewCheckUpdateInstanceRequest()
    }
    response = NewCheckUpdateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "CreateInstance")
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建指定规格的ES集群实例
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DeleteInstance")
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 销毁集群实例 
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogsRequest() (request *DescribeInstanceLogsRequest) {
    request = &DescribeInstanceLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstanceLogs")
    return
}

func NewDescribeInstanceLogsResponse() (response *DescribeInstanceLogsResponse) {
    response = &DescribeInstanceLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户该地域下符合条件的ES集群的日志
func (c *Client) DescribeInstanceLogs(request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogsRequest()
    }
    response = NewDescribeInstanceLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceOperationsRequest() (request *DescribeInstanceOperationsRequest) {
    request = &DescribeInstanceOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstanceOperations")
    return
}

func NewDescribeInstanceOperationsResponse() (response *DescribeInstanceOperationsResponse) {
    response = &DescribeInstanceOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例指定条件下的操作记录
func (c *Client) DescribeInstanceOperations(request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceOperationsRequest()
    }
    response = NewDescribeInstanceOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户该地域下符合条件的所有实例
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpgradeRequest() (request *DescribeUpgradeRequest) {
    request = &DescribeUpgradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeUpgrade")
    return
}

func NewDescribeUpgradeResponse() (response *DescribeUpgradeResponse) {
    response = &DescribeUpgradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取实例可升级列表，包括可升级的大版本、商业特性
func (c *Client) DescribeUpgrade(request *DescribeUpgradeRequest) (response *DescribeUpgradeResponse, err error) {
    if request == nil {
        request = NewDescribeUpgradeRequest()
    }
    response = NewDescribeUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterMetricRequest() (request *GetClusterMetricRequest) {
    request = &GetClusterMetricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "GetClusterMetric")
    return
}

func NewGetClusterMetricResponse() (response *GetClusterMetricResponse) {
    response = &GetClusterMetricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据给定的实例ID，获取集群的实时监控指标，如节点个数，数据节点个数，索引数，文档数，已使用磁盘大小，分片数，主分片数，迁移中分片数，初始化分片数，未分配分片数等
func (c *Client) GetClusterMetric(request *GetClusterMetricRequest) (response *GetClusterMetricResponse, err error) {
    if request == nil {
        request = NewGetClusterMetricRequest()
    }
    response = NewGetClusterMetricResponse()
    err = c.Send(request, response)
    return
}

func NewGetNodesMetricRequest() (request *GetNodesMetricRequest) {
    request = &GetNodesMetricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "GetNodesMetric")
    return
}

func NewGetNodesMetricResponse() (response *GetNodesMetricResponse) {
    response = &GetNodesMetricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定实例节点维度的监控指标，如IP， CPU使用率，JVM内存使用率，可使用磁盘空间，分片数等
func (c *Client) GetNodesMetric(request *GetNodesMetricRequest) (response *GetNodesMetricResponse, err error) {
    if request == nil {
        request = NewGetNodesMetricRequest()
    }
    response = NewGetNodesMetricResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRegionZoneRequest() (request *QueryRegionZoneRequest) {
    request = &QueryRegionZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "QueryRegionZone")
    return
}

func NewQueryRegionZoneResponse() (response *QueryRegionZoneResponse) {
    response = &QueryRegionZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取可售卖地域及可用区
func (c *Client) QueryRegionZone(request *QueryRegionZoneRequest) (response *QueryRegionZoneResponse, err error) {
    if request == nil {
        request = NewQueryRegionZoneRequest()
    }
    response = NewQueryRegionZoneResponse()
    err = c.Send(request, response)
    return
}

func NewQueryZoneResourceRequest() (request *QueryZoneResourceRequest) {
    request = &QueryZoneResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "QueryZoneResource")
    return
}

func NewQueryZoneResourceResponse() (response *QueryZoneResourceResponse) {
    response = &QueryZoneResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定region下指定zone列表的资源情况
func (c *Client) QueryZoneResource(request *QueryZoneResourceRequest) (response *QueryZoneResourceResponse, err error) {
    if request == nil {
        request = NewQueryZoneResourceRequest()
    }
    response = NewQueryZoneResourceResponse()
    err = c.Send(request, response)
    return
}

func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
    request = &RestartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "RestartInstance")
    return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
    response = &RestartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重启ES集群实例(用于系统版本更新等操作) 
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    if request == nil {
        request = NewRestartInstanceRequest()
    }
    response = NewRestartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateInstanceRequest() (request *UpdateInstanceRequest) {
    request = &UpdateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpdateInstance")
    return
}

func NewUpdateInstanceResponse() (response *UpdateInstanceResponse) {
    response = &UpdateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对集群进行节点规格变更，修改实例名称，修改配置，重置密码， 添加Kibana黑白名单等操作。参数中InstanceId为必传参数，ForceRestart为选填参数，剩余参数传递组合及含义如下：
// - InstanceName：修改实例名称(仅用于标识实例)
// - NodeInfoList: 修改节点配置（节点横向扩缩容，纵向扩缩容，增加主节点，增加冷节点等）
// - EsConfig：修改集群配置
// - Password：修改默认用户elastic的密码
// - CosBackUp: 设置集群COS自动备份信息
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
func (c *Client) UpdateInstance(request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateInstanceRequest()
    }
    response = NewUpdateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateInternalRequest() (request *UpdateInternalRequest) {
    request = &UpdateInternalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpdateInternal")
    return
}

func NewUpdateInternalResponse() (response *UpdateInternalResponse) {
    response = &UpdateInternalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于ES内部使用的更新接口
func (c *Client) UpdateInternal(request *UpdateInternalRequest) (response *UpdateInternalResponse, err error) {
    if request == nil {
        request = NewUpdateInternalRequest()
    }
    response = NewUpdateInternalResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpgradeInstance")
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 升级ES集群版本
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeLicenseRequest() (request *UpgradeLicenseRequest) {
    request = &UpgradeLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpgradeLicense")
    return
}

func NewUpgradeLicenseResponse() (response *UpgradeLicenseResponse) {
    response = &UpgradeLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 升级ES商业特性
func (c *Client) UpgradeLicense(request *UpgradeLicenseRequest) (response *UpgradeLicenseResponse, err error) {
    if request == nil {
        request = NewUpgradeLicenseRequest()
    }
    response = NewUpgradeLicenseResponse()
    err = c.Send(request, response)
    return
}
