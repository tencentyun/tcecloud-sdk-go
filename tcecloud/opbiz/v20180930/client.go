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

package v20180930

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-09-30"

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


func NewAllBizTopoRequest() (request *AllBizTopoRequest) {
    request = &AllBizTopoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "AllBizTopo")
    return
}

func NewAllBizTopoResponse() (response *AllBizTopoResponse) {
    response = &AllBizTopoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取所有的业务树
func (c *Client) AllBizTopo(request *AllBizTopoRequest) (response *AllBizTopoResponse, err error) {
    if request == nil {
        request = NewAllBizTopoRequest()
    }
    response = NewAllBizTopoResponse()
    err = c.Send(request, response)
    return
}

func NewBindPkgRequest() (request *BindPkgRequest) {
    request = &BindPkgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "BindPkg")
    return
}

func NewBindPkgResponse() (response *BindPkgResponse) {
    response = &BindPkgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑包
func (c *Client) BindPkg(request *BindPkgRequest) (response *BindPkgResponse, err error) {
    if request == nil {
        request = NewBindPkgRequest()
    }
    response = NewBindPkgResponse()
    err = c.Send(request, response)
    return
}

func NewBizTopoRequest() (request *BizTopoRequest) {
    request = &BizTopoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "BizTopo")
    return
}

func NewBizTopoResponse() (response *BizTopoResponse) {
    response = &BizTopoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务拓扑
func (c *Client) BizTopo(request *BizTopoRequest) (response *BizTopoResponse, err error) {
    if request == nil {
        request = NewBizTopoRequest()
    }
    response = NewBizTopoResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStandConfigureRequest() (request *CreateStandConfigureRequest) {
    request = &CreateStandConfigureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "CreateStandConfigure")
    return
}

func NewCreateStandConfigureResponse() (response *CreateStandConfigureResponse) {
    response = &CreateStandConfigureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增标准化配置
func (c *Client) CreateStandConfigure(request *CreateStandConfigureRequest) (response *CreateStandConfigureResponse, err error) {
    if request == nil {
        request = NewCreateStandConfigureRequest()
    }
    response = NewCreateStandConfigureResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopoConfigRequest() (request *CreateTopoConfigRequest) {
    request = &CreateTopoConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "CreateTopoConfig")
    return
}

func NewCreateTopoConfigResponse() (response *CreateTopoConfigResponse) {
    response = &CreateTopoConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建配置
func (c *Client) CreateTopoConfig(request *CreateTopoConfigRequest) (response *CreateTopoConfigResponse, err error) {
    if request == nil {
        request = NewCreateTopoConfigRequest()
    }
    response = NewCreateTopoConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStandConfigureRequest() (request *DeleteStandConfigureRequest) {
    request = &DeleteStandConfigureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "DeleteStandConfigure")
    return
}

func NewDeleteStandConfigureResponse() (response *DeleteStandConfigureResponse) {
    response = &DeleteStandConfigureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除标准化配置
func (c *Client) DeleteStandConfigure(request *DeleteStandConfigureRequest) (response *DeleteStandConfigureResponse, err error) {
    if request == nil {
        request = NewDeleteStandConfigureRequest()
    }
    response = NewDeleteStandConfigureResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopoConfigRequest() (request *DeleteTopoConfigRequest) {
    request = &DeleteTopoConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "DeleteTopoConfig")
    return
}

func NewDeleteTopoConfigResponse() (response *DeleteTopoConfigResponse) {
    response = &DeleteTopoConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除配置
func (c *Client) DeleteTopoConfig(request *DeleteTopoConfigRequest) (response *DeleteTopoConfigResponse, err error) {
    if request == nil {
        request = NewDeleteTopoConfigRequest()
    }
    response = NewDeleteTopoConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeployOverviewRequest() (request *DeployOverviewRequest) {
    request = &DeployOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "DeployOverview")
    return
}

func NewDeployOverviewResponse() (response *DeployOverviewResponse) {
    response = &DeployOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 部署概览
func (c *Client) DeployOverview(request *DeployOverviewRequest) (response *DeployOverviewResponse, err error) {
    if request == nil {
        request = NewDeployOverviewRequest()
    }
    response = NewDeployOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewEnumInfoRequest() (request *EnumInfoRequest) {
    request = &EnumInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "EnumInfo")
    return
}

func NewEnumInfoResponse() (response *EnumInfoResponse) {
    response = &EnumInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对象枚举信息
func (c *Client) EnumInfo(request *EnumInfoRequest) (response *EnumInfoResponse, err error) {
    if request == nil {
        request = NewEnumInfoRequest()
    }
    response = NewEnumInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetModulePkgInfoRequest() (request *GetModulePkgInfoRequest) {
    request = &GetModulePkgInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "GetModulePkgInfo")
    return
}

func NewGetModulePkgInfoResponse() (response *GetModulePkgInfoResponse) {
    response = &GetModulePkgInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取模块绑定的文件包列表
func (c *Client) GetModulePkgInfo(request *GetModulePkgInfoRequest) (response *GetModulePkgInfoResponse, err error) {
    if request == nil {
        request = NewGetModulePkgInfoRequest()
    }
    response = NewGetModulePkgInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetModuleSchedulesListRequest() (request *GetModuleSchedulesListRequest) {
    request = &GetModuleSchedulesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "GetModuleSchedulesList")
    return
}

func NewGetModuleSchedulesListResponse() (response *GetModuleSchedulesListResponse) {
    response = &GetModuleSchedulesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 业务运维获取模块管理的定时任务列表
func (c *Client) GetModuleSchedulesList(request *GetModuleSchedulesListRequest) (response *GetModuleSchedulesListResponse, err error) {
    if request == nil {
        request = NewGetModuleSchedulesListRequest()
    }
    response = NewGetModuleSchedulesListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPkgInsInfoRequest() (request *GetPkgInsInfoRequest) {
    request = &GetPkgInsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "GetPkgInsInfo")
    return
}

func NewGetPkgInsInfoResponse() (response *GetPkgInsInfoResponse) {
    response = &GetPkgInsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取文件包实例列表
func (c *Client) GetPkgInsInfo(request *GetPkgInsInfoRequest) (response *GetPkgInsInfoResponse, err error) {
    if request == nil {
        request = NewGetPkgInsInfoRequest()
    }
    response = NewGetPkgInsInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetSchedulesTaskInfoRequest() (request *GetSchedulesTaskInfoRequest) {
    request = &GetSchedulesTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "GetSchedulesTaskInfo")
    return
}

func NewGetSchedulesTaskInfoResponse() (response *GetSchedulesTaskInfoResponse) {
    response = &GetSchedulesTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSchedulesTaskInfo
func (c *Client) GetSchedulesTaskInfo(request *GetSchedulesTaskInfoRequest) (response *GetSchedulesTaskInfoResponse, err error) {
    if request == nil {
        request = NewGetSchedulesTaskInfoRequest()
    }
    response = NewGetSchedulesTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskInfoRequest() (request *GetTaskInfoRequest) {
    request = &GetTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "GetTaskInfo")
    return
}

func NewGetTaskInfoResponse() (response *GetTaskInfoResponse) {
    response = &GetTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询作业编排任务列表
func (c *Client) GetTaskInfo(request *GetTaskInfoRequest) (response *GetTaskInfoResponse, err error) {
    if request == nil {
        request = NewGetTaskInfoRequest()
    }
    response = NewGetTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowTemplateDetailRequest() (request *GetWorkflowTemplateDetailRequest) {
    request = &GetWorkflowTemplateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "GetWorkflowTemplateDetail")
    return
}

func NewGetWorkflowTemplateDetailResponse() (response *GetWorkflowTemplateDetailResponse) {
    response = &GetWorkflowTemplateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取作业编排
func (c *Client) GetWorkflowTemplateDetail(request *GetWorkflowTemplateDetailRequest) (response *GetWorkflowTemplateDetailResponse, err error) {
    if request == nil {
        request = NewGetWorkflowTemplateDetailRequest()
    }
    response = NewGetWorkflowTemplateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowTemplatesRequest() (request *GetWorkflowTemplatesRequest) {
    request = &GetWorkflowTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "GetWorkflowTemplates")
    return
}

func NewGetWorkflowTemplatesResponse() (response *GetWorkflowTemplatesResponse) {
    response = &GetWorkflowTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取作业编排
func (c *Client) GetWorkflowTemplates(request *GetWorkflowTemplatesRequest) (response *GetWorkflowTemplatesResponse, err error) {
    if request == nil {
        request = NewGetWorkflowTemplatesRequest()
    }
    response = NewGetWorkflowTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStandConfigureRequest() (request *ModifyStandConfigureRequest) {
    request = &ModifyStandConfigureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "ModifyStandConfigure")
    return
}

func NewModifyStandConfigureResponse() (response *ModifyStandConfigureResponse) {
    response = &ModifyStandConfigureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更改标准化配置
func (c *Client) ModifyStandConfigure(request *ModifyStandConfigureRequest) (response *ModifyStandConfigureResponse, err error) {
    if request == nil {
        request = NewModifyStandConfigureRequest()
    }
    response = NewModifyStandConfigureResponse()
    err = c.Send(request, response)
    return
}

func NewNodeHostRequest() (request *NodeHostRequest) {
    request = &NodeHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "NodeHost")
    return
}

func NewNodeHostResponse() (response *NodeHostResponse) {
    response = &NodeHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 节点下所有主机
func (c *Client) NodeHost(request *NodeHostRequest) (response *NodeHostResponse, err error) {
    if request == nil {
        request = NewNodeHostRequest()
    }
    response = NewNodeHostResponse()
    err = c.Send(request, response)
    return
}

func NewNodeInfoRequest() (request *NodeInfoRequest) {
    request = &NodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "NodeInfo")
    return
}

func NewNodeInfoResponse() (response *NodeInfoResponse) {
    response = &NodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取节点基本信息
func (c *Client) NodeInfo(request *NodeInfoRequest) (response *NodeInfoResponse, err error) {
    if request == nil {
        request = NewNodeInfoRequest()
    }
    response = NewNodeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAuditRequest() (request *QueryAuditRequest) {
    request = &QueryAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "QueryAudit")
    return
}

func NewQueryAuditResponse() (response *QueryAuditResponse) {
    response = &QueryAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 审计日志
func (c *Client) QueryAudit(request *QueryAuditRequest) (response *QueryAuditResponse, err error) {
    if request == nil {
        request = NewQueryAuditRequest()
    }
    response = NewQueryAuditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBizRequest() (request *QueryBizRequest) {
    request = &QueryBizRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "QueryBiz")
    return
}

func NewQueryBizResponse() (response *QueryBizResponse) {
    response = &QueryBizResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询业务
func (c *Client) QueryBiz(request *QueryBizRequest) (response *QueryBizResponse, err error) {
    if request == nil {
        request = NewQueryBizRequest()
    }
    response = NewQueryBizResponse()
    err = c.Send(request, response)
    return
}

func NewQueryModuleRequest() (request *QueryModuleRequest) {
    request = &QueryModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "QueryModule")
    return
}

func NewQueryModuleResponse() (response *QueryModuleResponse) {
    response = &QueryModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询模块
func (c *Client) QueryModule(request *QueryModuleRequest) (response *QueryModuleResponse, err error) {
    if request == nil {
        request = NewQueryModuleRequest()
    }
    response = NewQueryModuleResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPkgListRequest() (request *QueryPkgListRequest) {
    request = &QueryPkgListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "QueryPkgList")
    return
}

func NewQueryPkgListResponse() (response *QueryPkgListResponse) {
    response = &QueryPkgListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取织云包列表
func (c *Client) QueryPkgList(request *QueryPkgListRequest) (response *QueryPkgListResponse, err error) {
    if request == nil {
        request = NewQueryPkgListRequest()
    }
    response = NewQueryPkgListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPkgVersionListRequest() (request *QueryPkgVersionListRequest) {
    request = &QueryPkgVersionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "QueryPkgVersionList")
    return
}

func NewQueryPkgVersionListResponse() (response *QueryPkgVersionListResponse) {
    response = &QueryPkgVersionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取织云包版本列表
func (c *Client) QueryPkgVersionList(request *QueryPkgVersionListRequest) (response *QueryPkgVersionListResponse, err error) {
    if request == nil {
        request = NewQueryPkgVersionListRequest()
    }
    response = NewQueryPkgVersionListResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySetRequest() (request *QuerySetRequest) {
    request = &QuerySetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "QuerySet")
    return
}

func NewQuerySetResponse() (response *QuerySetResponse) {
    response = &QuerySetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询蓝鲸CMDB集群
func (c *Client) QuerySet(request *QuerySetRequest) (response *QuerySetResponse, err error) {
    if request == nil {
        request = NewQuerySetRequest()
    }
    response = NewQuerySetResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTopoConfigRequest() (request *QueryTopoConfigRequest) {
    request = &QueryTopoConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "QueryTopoConfig")
    return
}

func NewQueryTopoConfigResponse() (response *QueryTopoConfigResponse) {
    response = &QueryTopoConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询配置
func (c *Client) QueryTopoConfig(request *QueryTopoConfigRequest) (response *QueryTopoConfigResponse, err error) {
    if request == nil {
        request = NewQueryTopoConfigRequest()
    }
    response = NewQueryTopoConfigResponse()
    err = c.Send(request, response)
    return
}

func NewQueryUuidPkgLastVersionRequest() (request *QueryUuidPkgLastVersionRequest) {
    request = &QueryUuidPkgLastVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "QueryUuidPkgLastVersion")
    return
}

func NewQueryUuidPkgLastVersionResponse() (response *QueryUuidPkgLastVersionResponse) {
    response = &QueryUuidPkgLastVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询机器的安装过某个包的上一版本
func (c *Client) QueryUuidPkgLastVersion(request *QueryUuidPkgLastVersionRequest) (response *QueryUuidPkgLastVersionResponse, err error) {
    if request == nil {
        request = NewQueryUuidPkgLastVersionRequest()
    }
    response = NewQueryUuidPkgLastVersionResponse()
    err = c.Send(request, response)
    return
}

func NewRunModuleWorkflowRequest() (request *RunModuleWorkflowRequest) {
    request = &RunModuleWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "RunModuleWorkflow")
    return
}

func NewRunModuleWorkflowResponse() (response *RunModuleWorkflowResponse) {
    response = &RunModuleWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 执行作业编排任务
func (c *Client) RunModuleWorkflow(request *RunModuleWorkflowRequest) (response *RunModuleWorkflowResponse, err error) {
    if request == nil {
        request = NewRunModuleWorkflowRequest()
    }
    response = NewRunModuleWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewRunStandardizationTaskRequest() (request *RunStandardizationTaskRequest) {
    request = &RunStandardizationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "RunStandardizationTask")
    return
}

func NewRunStandardizationTaskResponse() (response *RunStandardizationTaskResponse) {
    response = &RunStandardizationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 执行标准化任务
func (c *Client) RunStandardizationTask(request *RunStandardizationTaskRequest) (response *RunStandardizationTaskResponse, err error) {
    if request == nil {
        request = NewRunStandardizationTaskRequest()
    }
    response = NewRunStandardizationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSearchStandConfigureRequest() (request *SearchStandConfigureRequest) {
    request = &SearchStandConfigureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "SearchStandConfigure")
    return
}

func NewSearchStandConfigureResponse() (response *SearchStandConfigureResponse) {
    response = &SearchStandConfigureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询标准化配置
func (c *Client) SearchStandConfigure(request *SearchStandConfigureRequest) (response *SearchStandConfigureResponse, err error) {
    if request == nil {
        request = NewSearchStandConfigureRequest()
    }
    response = NewSearchStandConfigureResponse()
    err = c.Send(request, response)
    return
}

func NewSearchStandardizationHistoryRequest() (request *SearchStandardizationHistoryRequest) {
    request = &SearchStandardizationHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "SearchStandardizationHistory")
    return
}

func NewSearchStandardizationHistoryResponse() (response *SearchStandardizationHistoryResponse) {
    response = &SearchStandardizationHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 搜索标准化历史
func (c *Client) SearchStandardizationHistory(request *SearchStandardizationHistoryRequest) (response *SearchStandardizationHistoryResponse, err error) {
    if request == nil {
        request = NewSearchStandardizationHistoryRequest()
    }
    response = NewSearchStandardizationHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewSearchStandardizationTasksRequest() (request *SearchStandardizationTasksRequest) {
    request = &SearchStandardizationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "SearchStandardizationTasks")
    return
}

func NewSearchStandardizationTasksResponse() (response *SearchStandardizationTasksResponse) {
    response = &SearchStandardizationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询标准化任务
func (c *Client) SearchStandardizationTasks(request *SearchStandardizationTasksRequest) (response *SearchStandardizationTasksResponse, err error) {
    if request == nil {
        request = NewSearchStandardizationTasksRequest()
    }
    response = NewSearchStandardizationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindPkgRequest() (request *UnbindPkgRequest) {
    request = &UnbindPkgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "UnbindPkg")
    return
}

func NewUnbindPkgResponse() (response *UnbindPkgResponse) {
    response = &UnbindPkgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑包
func (c *Client) UnbindPkg(request *UnbindPkgRequest) (response *UnbindPkgResponse, err error) {
    if request == nil {
        request = NewUnbindPkgRequest()
    }
    response = NewUnbindPkgResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateHostRequest() (request *UpdateHostRequest) {
    request = &UpdateHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "UpdateHost")
    return
}

func NewUpdateHostResponse() (response *UpdateHostResponse) {
    response = &UpdateHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改主机
func (c *Client) UpdateHost(request *UpdateHostRequest) (response *UpdateHostResponse, err error) {
    if request == nil {
        request = NewUpdateHostRequest()
    }
    response = NewUpdateHostResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNodeRequest() (request *UpdateNodeRequest) {
    request = &UpdateNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "UpdateNode")
    return
}

func NewUpdateNodeResponse() (response *UpdateNodeResponse) {
    response = &UpdateNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改节点基本信息
func (c *Client) UpdateNode(request *UpdateNodeRequest) (response *UpdateNodeResponse, err error) {
    if request == nil {
        request = NewUpdateNodeRequest()
    }
    response = NewUpdateNodeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTopoConfigRequest() (request *UpdateTopoConfigRequest) {
    request = &UpdateTopoConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opbiz", APIVersion, "UpdateTopoConfig")
    return
}

func NewUpdateTopoConfigResponse() (response *UpdateTopoConfigResponse) {
    response = &UpdateTopoConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改配置
func (c *Client) UpdateTopoConfig(request *UpdateTopoConfigRequest) (response *UpdateTopoConfigResponse, err error) {
    if request == nil {
        request = NewUpdateTopoConfigRequest()
    }
    response = NewUpdateTopoConfigResponse()
    err = c.Send(request, response)
    return
}
