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

package v20190625

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-25"

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


func NewColdMigrateInstancesRequest() (request *ColdMigrateInstancesRequest) {
    request = &ColdMigrateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "ColdMigrateInstances")
    return
}

func NewColdMigrateInstancesResponse() (response *ColdMigrateInstancesResponse) {
    response = &ColdMigrateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将子机迁移到指定的宿主机上面
func (c *Client) ColdMigrateInstances(request *ColdMigrateInstancesRequest) (response *ColdMigrateInstancesResponse, err error) {
    if request == nil {
        request = NewColdMigrateInstancesRequest()
    }
    response = NewColdMigrateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewColdMigrateTenantInstancesRequest() (request *ColdMigrateTenantInstancesRequest) {
    request = &ColdMigrateTenantInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "ColdMigrateTenantInstances")
    return
}

func NewColdMigrateTenantInstancesResponse() (response *ColdMigrateTenantInstancesResponse) {
    response = &ColdMigrateTenantInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运营端调用，去冷迁移租户端的子机
func (c *Client) ColdMigrateTenantInstances(request *ColdMigrateTenantInstancesRequest) (response *ColdMigrateTenantInstancesResponse, err error) {
    if request == nil {
        request = NewColdMigrateTenantInstancesRequest()
    }
    response = NewColdMigrateTenantInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostTypesRequest() (request *CreateHostTypesRequest) {
    request = &CreateHostTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "CreateHostTypes")
    return
}

func NewCreateHostTypesResponse() (response *CreateHostTypesResponse) {
    response = &CreateHostTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于创建新的宿主机售卖类型
func (c *Client) CreateHostTypes(request *CreateHostTypesRequest) (response *CreateHostTypesResponse, err error) {
    if request == nil {
        request = NewCreateHostTypesRequest()
    }
    response = NewCreateHostTypesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
    request = &CreateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "CreateImage")
    return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
    response = &CreateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建自定义镜像接口
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceTypeConfigRequest() (request *CreateInstanceTypeConfigRequest) {
    request = &CreateInstanceTypeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "CreateInstanceTypeConfig")
    return
}

func NewCreateInstanceTypeConfigResponse() (response *CreateInstanceTypeConfigResponse) {
    response = &CreateInstanceTypeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 供运营端调用，创建全部规格配置
func (c *Client) CreateInstanceTypeConfig(request *CreateInstanceTypeConfigRequest) (response *CreateInstanceTypeConfigResponse, err error) {
    if request == nil {
        request = NewCreateInstanceTypeConfigRequest()
    }
    response = NewCreateInstanceTypeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublicImageRequest() (request *CreatePublicImageRequest) {
    request = &CreatePublicImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "CreatePublicImage")
    return
}

func NewCreatePublicImageResponse() (response *CreatePublicImageResponse) {
    response = &CreatePublicImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePublicImage将自定义镜像转为公共镜像
func (c *Client) CreatePublicImage(request *CreatePublicImageRequest) (response *CreatePublicImageResponse, err error) {
    if request == nil {
        request = NewCreatePublicImageRequest()
    }
    response = NewCreatePublicImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneInstanceTypeConfigRequest() (request *CreateZoneInstanceTypeConfigRequest) {
    request = &CreateZoneInstanceTypeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "CreateZoneInstanceTypeConfig")
    return
}

func NewCreateZoneInstanceTypeConfigResponse() (response *CreateZoneInstanceTypeConfigResponse) {
    response = &CreateZoneInstanceTypeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 基于全局机型配置创建可用区机型
func (c *Client) CreateZoneInstanceTypeConfig(request *CreateZoneInstanceTypeConfigRequest) (response *CreateZoneInstanceTypeConfigResponse, err error) {
    if request == nil {
        request = NewCreateZoneInstanceTypeConfigRequest()
    }
    response = NewCreateZoneInstanceTypeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
    request = &DeleteImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DeleteImages")
    return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
    response = &DeleteImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteImages）用于删除一个或多个镜像。
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    if request == nil {
        request = NewDeleteImagesRequest()
    }
    response = NewDeleteImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowLogsRequest() (request *DescribeFlowLogsRequest) {
    request = &DescribeFlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeFlowLogs")
    return
}

func NewDescribeFlowLogsResponse() (response *DescribeFlowLogsResponse) {
    response = &DescribeFlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeFlowLogs) 用于查询 CVM 后端业务模块的日志流水。
// 
// * 支持根据 AppId、Uin、SubAccountUin 查询该用户在指定业务模块下操作记录。
// * 支持根据 RequestId 查询该请求在 CVM_API 这个模块下的操作日志信息。
// * 支持根据 RequestId 查询该请求调用 CVM_DES 这个模块返回的的任务 ID。
// * 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的日志信息。
func (c *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (response *DescribeFlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLogsRequest()
    }
    response = NewDescribeFlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostTypesRequest() (request *DescribeHostTypesRequest) {
    request = &DescribeHostTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeHostTypes")
    return
}

func NewDescribeHostTypesResponse() (response *DescribeHostTypesResponse) {
    response = &DescribeHostTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询当前系统支持的宿主机类型列表
func (c *Client) DescribeHostTypes(request *DescribeHostTypesRequest) (response *DescribeHostTypesResponse, err error) {
    if request == nil {
        request = NewDescribeHostTypesRequest()
    }
    response = NewDescribeHostTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
    request = &DescribeHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeHosts")
    return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
    response = &DescribeHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询宿主机信息
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
    if request == nil {
        request = NewDescribeHostsRequest()
    }
    response = NewDescribeHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRecordRequest() (request *DescribeImageRecordRequest) {
    request = &DescribeImageRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeImageRecord")
    return
}

func NewDescribeImageRecordResponse() (response *DescribeImageRecordResponse) {
    response = &DescribeImageRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 包括自定义镜像转公共镜像及公共镜像同步到多region的记录
func (c *Client) DescribeImageRecord(request *DescribeImageRecordRequest) (response *DescribeImageRecordResponse, err error) {
    if request == nil {
        request = NewDescribeImageRecordRequest()
    }
    response = NewDescribeImageRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeImages")
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 统一查询镜像的接口
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceClassConfigsRequest() (request *DescribeInstanceClassConfigsRequest) {
    request = &DescribeInstanceClassConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceClassConfigs")
    return
}

func NewDescribeInstanceClassConfigsResponse() (response *DescribeInstanceClassConfigsResponse) {
    response = &DescribeInstanceClassConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询当前支持的实例系列的相信信息，比如“S”
func (c *Client) DescribeInstanceClassConfigs(request *DescribeInstanceClassConfigsRequest) (response *DescribeInstanceClassConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceClassConfigsRequest()
    }
    response = NewDescribeInstanceClassConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceConfigInfosRequest() (request *DescribeInstanceConfigInfosRequest) {
    request = &DescribeInstanceConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceConfigInfos")
    return
}

func NewDescribeInstanceConfigInfosResponse() (response *DescribeInstanceConfigInfosResponse) {
    response = &DescribeInstanceConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询全局机器信息
func (c *Client) DescribeInstanceConfigInfos(request *DescribeInstanceConfigInfosRequest) (response *DescribeInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceConfigInfosRequest()
    }
    response = NewDescribeInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceFamilyConfigsRequest() (request *DescribeInstanceFamilyConfigsRequest) {
    request = &DescribeInstanceFamilyConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceFamilyConfigs")
    return
}

func NewDescribeInstanceFamilyConfigsResponse() (response *DescribeInstanceFamilyConfigsResponse) {
    response = &DescribeInstanceFamilyConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstanceFamilyConfigs）查询当前用户和地域所支持的机型族列表信息
func (c *Client) DescribeInstanceFamilyConfigs(request *DescribeInstanceFamilyConfigsRequest) (response *DescribeInstanceFamilyConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceFamilyConfigsRequest()
    }
    response = NewDescribeInstanceFamilyConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypeConfigsRequest() (request *DescribeInstanceTypeConfigsRequest) {
    request = &DescribeInstanceTypeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceTypeConfigs")
    return
}

func NewDescribeInstanceTypeConfigsResponse() (response *DescribeInstanceTypeConfigsResponse) {
    response = &DescribeInstanceTypeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeInstanceTypeConfigs) 用于查询实例机型配置。
func (c *Client) DescribeInstanceTypeConfigs(request *DescribeInstanceTypeConfigsRequest) (response *DescribeInstanceTypeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypeConfigsRequest()
    }
    response = NewDescribeInstanceTypeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeInstances) 用于查询一个或多个实例的详细信息。
// * 可以根据实例`ID`、实例名称或者实例计费模式等信息来查询实例的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的实例。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
    request = &DescribeKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeKeyPairs")
    return
}

func NewDescribeKeyPairsResponse() (response *DescribeKeyPairsResponse) {
    response = &DescribeKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询密钥对
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    if request == nil {
        request = NewDescribeKeyPairsRequest()
    }
    response = NewDescribeKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询不同任务的状态
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTenantInstancesRequest() (request *DescribeTenantInstancesRequest) {
    request = &DescribeTenantInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeTenantInstances")
    return
}

func NewDescribeTenantInstancesResponse() (response *DescribeTenantInstancesResponse) {
    response = &DescribeTenantInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运营端调用
func (c *Client) DescribeTenantInstances(request *DescribeTenantInstancesRequest) (response *DescribeTenantInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTenantInstancesRequest()
    }
    response = NewDescribeTenantInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVncRequest() (request *DescribeVncRequest) {
    request = &DescribeVncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeVnc")
    return
}

func NewDescribeVncResponse() (response *DescribeVncResponse) {
    response = &DescribeVncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看VNC信息
func (c *Client) DescribeVnc(request *DescribeVncRequest) (response *DescribeVncResponse, err error) {
    if request == nil {
        request = NewDescribeVncRequest()
    }
    response = NewDescribeVncResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVsChildTaskIdByParentRequest() (request *DescribeVsChildTaskIdByParentRequest) {
    request = &DescribeVsChildTaskIdByParentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeVsChildTaskIdByParent")
    return
}

func NewDescribeVsChildTaskIdByParentResponse() (response *DescribeVsChildTaskIdByParentResponse) {
    response = &DescribeVsChildTaskIdByParentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeVsChildTaskIdByParent) 用于根据父 CVM_VS 任务ID查询子任务。
func (c *Client) DescribeVsChildTaskIdByParent(request *DescribeVsChildTaskIdByParentRequest) (response *DescribeVsChildTaskIdByParentResponse, err error) {
    if request == nil {
        request = NewDescribeVsChildTaskIdByParentRequest()
    }
    response = NewDescribeVsChildTaskIdByParentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVsChildTasksRequest() (request *DescribeVsChildTasksRequest) {
    request = &DescribeVsChildTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeVsChildTasks")
    return
}

func NewDescribeVsChildTasksResponse() (response *DescribeVsChildTasksResponse) {
    response = &DescribeVsChildTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeVsChildTasks) 用于查询 CVM_VS 子任务详情。
func (c *Client) DescribeVsChildTasks(request *DescribeVsChildTasksRequest) (response *DescribeVsChildTasksResponse, err error) {
    if request == nil {
        request = NewDescribeVsChildTasksRequest()
    }
    response = NewDescribeVsChildTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneImportInstanceTypeConfigsRequest() (request *DescribeZoneImportInstanceTypeConfigsRequest) {
    request = &DescribeZoneImportInstanceTypeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeZoneImportInstanceTypeConfigs")
    return
}

func NewDescribeZoneImportInstanceTypeConfigsResponse() (response *DescribeZoneImportInstanceTypeConfigsResponse) {
    response = &DescribeZoneImportInstanceTypeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看指定可用区可导入的机型配置
func (c *Client) DescribeZoneImportInstanceTypeConfigs(request *DescribeZoneImportInstanceTypeConfigsRequest) (response *DescribeZoneImportInstanceTypeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeZoneImportInstanceTypeConfigsRequest()
    }
    response = NewDescribeZoneImportInstanceTypeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneInstanceConfigInfosRequest() (request *DescribeZoneInstanceConfigInfosRequest) {
    request = &DescribeZoneInstanceConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "DescribeZoneInstanceConfigInfos")
    return
}

func NewDescribeZoneInstanceConfigInfosResponse() (response *DescribeZoneInstanceConfigInfosResponse) {
    response = &DescribeZoneInstanceConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeZoneInstanceConfigInfos) 获取可用区的机型信息
func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeZoneInstanceConfigInfosRequest()
    }
    response = NewDescribeZoneInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewFailTaskRequest() (request *FailTaskRequest) {
    request = &FailTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "FailTask")
    return
}

func NewFailTaskResponse() (response *FailTaskResponse) {
    response = &FailTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置一个vs任务为失败
func (c *Client) FailTask(request *FailTaskRequest) (response *FailTaskResponse, err error) {
    if request == nil {
        request = NewFailTaskRequest()
    }
    response = NewFailTaskResponse()
    err = c.Send(request, response)
    return
}

func NewFailoverMigrateInstancesRequest() (request *FailoverMigrateInstancesRequest) {
    request = &FailoverMigrateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "FailoverMigrateInstances")
    return
}

func NewFailoverMigrateInstancesResponse() (response *FailoverMigrateInstancesResponse) {
    response = &FailoverMigrateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 当母机产生故障的时候迁移虚拟机
func (c *Client) FailoverMigrateInstances(request *FailoverMigrateInstancesRequest) (response *FailoverMigrateInstancesResponse, err error) {
    if request == nil {
        request = NewFailoverMigrateInstancesRequest()
    }
    response = NewFailoverMigrateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewLiveMigrateInstancesRequest() (request *LiveMigrateInstancesRequest) {
    request = &LiveMigrateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "LiveMigrateInstances")
    return
}

func NewLiveMigrateInstancesResponse() (response *LiveMigrateInstancesResponse) {
    response = &LiveMigrateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 不关机热迁移实例
func (c *Client) LiveMigrateInstances(request *LiveMigrateInstancesRequest) (response *LiveMigrateInstancesResponse, err error) {
    if request == nil {
        request = NewLiveMigrateInstancesRequest()
    }
    response = NewLiveMigrateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewLiveMigrateTenantInstancesRequest() (request *LiveMigrateTenantInstancesRequest) {
    request = &LiveMigrateTenantInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "LiveMigrateTenantInstances")
    return
}

func NewLiveMigrateTenantInstancesResponse() (response *LiveMigrateTenantInstancesResponse) {
    response = &LiveMigrateTenantInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 热迁移租户端实例
func (c *Client) LiveMigrateTenantInstances(request *LiveMigrateTenantInstancesRequest) (response *LiveMigrateTenantInstancesResponse, err error) {
    if request == nil {
        request = NewLiveMigrateTenantInstancesRequest()
    }
    response = NewLiveMigrateTenantInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostsTagRequest() (request *ModifyHostsTagRequest) {
    request = &ModifyHostsTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "ModifyHostsTag")
    return
}

func NewModifyHostsTagResponse() (response *ModifyHostsTagResponse) {
    response = &ModifyHostsTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加宿主机标签
func (c *Client) ModifyHostsTag(request *ModifyHostsTagRequest) (response *ModifyHostsTagResponse, err error) {
    if request == nil {
        request = NewModifyHostsTagRequest()
    }
    response = NewModifyHostsTagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceTypeConfigRequest() (request *ModifyInstanceTypeConfigRequest) {
    request = &ModifyInstanceTypeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "ModifyInstanceTypeConfig")
    return
}

func NewModifyInstanceTypeConfigResponse() (response *ModifyInstanceTypeConfigResponse) {
    response = &ModifyInstanceTypeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改全局配置的机器
func (c *Client) ModifyInstanceTypeConfig(request *ModifyInstanceTypeConfigRequest) (response *ModifyInstanceTypeConfigResponse, err error) {
    if request == nil {
        request = NewModifyInstanceTypeConfigRequest()
    }
    response = NewModifyInstanceTypeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "ModifyInstancesAttribute")
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ModifyInstancesAttribute) 用于修改实例的属性（目前只支持修改实例的名称）。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneInstanceTypeConfigRequest() (request *ModifyZoneInstanceTypeConfigRequest) {
    request = &ModifyZoneInstanceTypeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "ModifyZoneInstanceTypeConfig")
    return
}

func NewModifyZoneInstanceTypeConfigResponse() (response *ModifyZoneInstanceTypeConfigResponse) {
    response = &ModifyZoneInstanceTypeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改可用区内的实例机型
func (c *Client) ModifyZoneInstanceTypeConfig(request *ModifyZoneInstanceTypeConfigRequest) (response *ModifyZoneInstanceTypeConfigResponse, err error) {
    if request == nil {
        request = NewModifyZoneInstanceTypeConfigRequest()
    }
    response = NewModifyZoneInstanceTypeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "RebootInstances")
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (RebootInstances) 用于重启实例。
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
    request = &ResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "ResetInstance")
    return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
    response = &ResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ResetInstance) 用于重装指定实例上的操作系统。
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "ResetInstancesPassword")
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ResetInstancesPassword) 用于将实例操作系统的密码重置为用户指定的密码。
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
    request = &RunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "RunInstances")
    return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
    response = &RunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运营端创建实例，以运营端的身份创建实例，和租户资源分开
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "StartInstances")
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (StartInstances) 用于启动一个或多个实例
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "StopInstances")
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (StopInstances) 用于关闭一个或多个实例。
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewSyncHostTypesRequest() (request *SyncHostTypesRequest) {
    request = &SyncHostTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "SyncHostTypes")
    return
}

func NewSyncHostTypesResponse() (response *SyncHostTypesResponse) {
    response = &SyncHostTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将指定地域的宿主机类型配置同步到当前地域
// 1、已存在的宿主机类型配置不会被覆盖
// 2、可装箱的虚拟机类型如果在当前地域不支持，则同步过来的宿主机类型也不支持该虚拟机类型
func (c *Client) SyncHostTypes(request *SyncHostTypesRequest) (response *SyncHostTypesResponse, err error) {
    if request == nil {
        request = NewSyncHostTypesRequest()
    }
    response = NewSyncHostTypesResponse()
    err = c.Send(request, response)
    return
}

func NewSyncImagesRequest() (request *SyncImagesRequest) {
    request = &SyncImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "SyncImages")
    return
}

func NewSyncImagesResponse() (response *SyncImagesResponse) {
    response = &SyncImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SyncImages）用于将自定义镜像同步到其它地区。
// 
// * 该接口每次调用只支持同步一个镜像。
// * 该接口支持多个同步地域。
// * 单个帐号在每个地域最多支持存在10个自定义镜像。
func (c *Client) SyncImages(request *SyncImagesRequest) (response *SyncImagesResponse, err error) {
    if request == nil {
        request = NewSyncImagesRequest()
    }
    response = NewSyncImagesResponse()
    err = c.Send(request, response)
    return
}

func NewSyncInstanceFamilyConfigsRequest() (request *SyncInstanceFamilyConfigsRequest) {
    request = &SyncInstanceFamilyConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "SyncInstanceFamilyConfigs")
    return
}

func NewSyncInstanceFamilyConfigsResponse() (response *SyncInstanceFamilyConfigsResponse) {
    response = &SyncInstanceFamilyConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将实例族信息同步到指定的可用区
func (c *Client) SyncInstanceFamilyConfigs(request *SyncInstanceFamilyConfigsRequest) (response *SyncInstanceFamilyConfigsResponse, err error) {
    if request == nil {
        request = NewSyncInstanceFamilyConfigsRequest()
    }
    response = NewSyncInstanceFamilyConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "TerminateInstances")
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (TerminateInstances) 用于主动退还实例。
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opcvm", APIVersion, "UpgradeInstance")
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 自由调整实例配置
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}
