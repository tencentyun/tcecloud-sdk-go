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


func NewApprovalRequest() (request *ApprovalRequest) {
    request = &ApprovalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "Approval")
    return
}

func NewApprovalResponse() (response *ApprovalResponse) {
    response = &ApprovalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 资源交付审批
func (c *Client) Approval(request *ApprovalRequest) (response *ApprovalResponse, err error) {
    if request == nil {
        request = NewApprovalRequest()
    }
    response = NewApprovalResponse()
    err = c.Send(request, response)
    return
}

func NewModuleWorkflowRequest() (request *ModuleWorkflowRequest) {
    request = &ModuleWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "ModuleWorkflow")
    return
}

func NewModuleWorkflowResponse() (response *ModuleWorkflowResponse) {
    response = &ModuleWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 模块绑定的标准化流程
func (c *Client) ModuleWorkflow(request *ModuleWorkflowRequest) (response *ModuleWorkflowResponse, err error) {
    if request == nil {
        request = NewModuleWorkflowRequest()
    }
    response = NewModuleWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBizRequest() (request *QueryBizRequest) {
    request = &QueryBizRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "QueryBiz")
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

func NewQueryHostRequest() (request *QueryHostRequest) {
    request = &QueryHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "QueryHost")
    return
}

func NewQueryHostResponse() (response *QueryHostResponse) {
    response = &QueryHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询主机
func (c *Client) QueryHost(request *QueryHostRequest) (response *QueryHostResponse, err error) {
    if request == nil {
        request = NewQueryHostRequest()
    }
    response = NewQueryHostResponse()
    err = c.Send(request, response)
    return
}

func NewQueryModuleRequest() (request *QueryModuleRequest) {
    request = &QueryModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "QueryModule")
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

func NewQueryResDeliveryTaskRequest() (request *QueryResDeliveryTaskRequest) {
    request = &QueryResDeliveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "QueryResDeliveryTask")
    return
}

func NewQueryResDeliveryTaskResponse() (response *QueryResDeliveryTaskResponse) {
    response = &QueryResDeliveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询详情的任务列表
func (c *Client) QueryResDeliveryTask(request *QueryResDeliveryTaskRequest) (response *QueryResDeliveryTaskResponse, err error) {
    if request == nil {
        request = NewQueryResDeliveryTaskRequest()
    }
    response = NewQueryResDeliveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewQueryResDetailRequest() (request *QueryResDetailRequest) {
    request = &QueryResDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "QueryResDetail")
    return
}

func NewQueryResDetailResponse() (response *QueryResDetailResponse) {
    response = &QueryResDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 资源交付详情
func (c *Client) QueryResDetail(request *QueryResDetailRequest) (response *QueryResDetailResponse, err error) {
    if request == nil {
        request = NewQueryResDetailRequest()
    }
    response = NewQueryResDetailResponse()
    err = c.Send(request, response)
    return
}

func NewResDeliveryRequest() (request *ResDeliveryRequest) {
    request = &ResDeliveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "ResDelivery")
    return
}

func NewResDeliveryResponse() (response *ResDeliveryResponse) {
    response = &ResDeliveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 资源交付申请
func (c *Client) ResDelivery(request *ResDeliveryRequest) (response *ResDeliveryResponse, err error) {
    if request == nil {
        request = NewResDeliveryRequest()
    }
    response = NewResDeliveryResponse()
    err = c.Send(request, response)
    return
}

func NewResDeliveryHistoryRequest() (request *ResDeliveryHistoryRequest) {
    request = &ResDeliveryHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "ResDeliveryHistory")
    return
}

func NewResDeliveryHistoryResponse() (response *ResDeliveryHistoryResponse) {
    response = &ResDeliveryHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 资源交付历史
func (c *Client) ResDeliveryHistory(request *ResDeliveryHistoryRequest) (response *ResDeliveryHistoryResponse, err error) {
    if request == nil {
        request = NewResDeliveryHistoryRequest()
    }
    response = NewResDeliveryHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewSearchStandConfigureRequest() (request *SearchStandConfigureRequest) {
    request = &SearchStandConfigureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "SearchStandConfigure")
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

func NewVMOSListRequest() (request *VMOSListRequest) {
    request = &VMOSListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opresdelivery", APIVersion, "VMOSList")
    return
}

func NewVMOSListResponse() (response *VMOSListResponse) {
    response = &VMOSListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 虚拟机操作系统列表
func (c *Client) VMOSList(request *VMOSListRequest) (response *VMOSListResponse, err error) {
    if request == nil {
        request = NewVMOSListRequest()
    }
    response = NewVMOSListResponse()
    err = c.Send(request, response)
    return
}
