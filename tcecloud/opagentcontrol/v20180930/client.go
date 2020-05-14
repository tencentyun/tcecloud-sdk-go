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
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
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


func NewCreateModuleRequest() (request *CreateModuleRequest) {
    request = &CreateModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "CreateModule")
    return
}

func NewCreateModuleResponse() (response *CreateModuleResponse) {
    response = &CreateModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增模块
func (c *Client) CreateModule(request *CreateModuleRequest) (response *CreateModuleResponse, err error) {
    if request == nil {
        request = NewCreateModuleRequest()
    }
    response = NewCreateModuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModuleNewVersionCreatePkgRequest() (request *CreateModuleNewVersionCreatePkgRequest) {
    request = &CreateModuleNewVersionCreatePkgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "CreateModuleNewVersionCreatePkg")
    return
}

func NewCreateModuleNewVersionCreatePkgResponse() (response *CreateModuleNewVersionCreatePkgResponse) {
    response = &CreateModuleNewVersionCreatePkgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建新模块版本之创建
func (c *Client) CreateModuleNewVersionCreatePkg(request *CreateModuleNewVersionCreatePkgRequest) (response *CreateModuleNewVersionCreatePkgResponse, err error) {
    if request == nil {
        request = NewCreateModuleNewVersionCreatePkgRequest()
    }
    response = NewCreateModuleNewVersionCreatePkgResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModuleNewVersionGetTmpdirRequest() (request *CreateModuleNewVersionGetTmpdirRequest) {
    request = &CreateModuleNewVersionGetTmpdirRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "CreateModuleNewVersionGetTmpdir")
    return
}

func NewCreateModuleNewVersionGetTmpdirResponse() (response *CreateModuleNewVersionGetTmpdirResponse) {
    response = &CreateModuleNewVersionGetTmpdirResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建新模块版本之获取临时目录ID
func (c *Client) CreateModuleNewVersionGetTmpdir(request *CreateModuleNewVersionGetTmpdirRequest) (response *CreateModuleNewVersionGetTmpdirResponse, err error) {
    if request == nil {
        request = NewCreateModuleNewVersionGetTmpdirRequest()
    }
    response = NewCreateModuleNewVersionGetTmpdirResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModulesRequest() (request *DescribeModulesRequest) {
    request = &DescribeModulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "DescribeModules")
    return
}

func NewDescribeModulesResponse() (response *DescribeModulesResponse) {
    response = &DescribeModulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询模块列表
func (c *Client) DescribeModules(request *DescribeModulesRequest) (response *DescribeModulesResponse, err error) {
    if request == nil {
        request = NewDescribeModulesRequest()
    }
    response = NewDescribeModulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodesRequest() (request *DescribeNodesRequest) {
    request = &DescribeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "DescribeNodes")
    return
}

func NewDescribeNodesResponse() (response *DescribeNodesResponse) {
    response = &DescribeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取节点名称、os类型、ip、节点类型、地域、可用区、主控状态、部署模块（版本、模块状态）、最后心跳等
func (c *Client) DescribeNodes(request *DescribeNodesRequest) (response *DescribeNodesResponse, err error) {
    if request == nil {
        request = NewDescribeNodesRequest()
    }
    response = NewDescribeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewGetEnumInfoRequest() (request *GetEnumInfoRequest) {
    request = &GetEnumInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "GetEnumInfo")
    return
}

func NewGetEnumInfoResponse() (response *GetEnumInfoResponse) {
    response = &GetEnumInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取对象枚举信息
func (c *Client) GetEnumInfo(request *GetEnumInfoRequest) (response *GetEnumInfoResponse, err error) {
    if request == nil {
        request = NewGetEnumInfoRequest()
    }
    response = NewGetEnumInfoResponse()
    err = c.Send(request, response)
    return
}

func NewInstallModulesRequest() (request *InstallModulesRequest) {
    request = &InstallModulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "InstallModules")
    return
}

func NewInstallModulesResponse() (response *InstallModulesResponse) {
    response = &InstallModulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 模块安装
func (c *Client) InstallModules(request *InstallModulesRequest) (response *InstallModulesResponse, err error) {
    if request == nil {
        request = NewInstallModulesRequest()
    }
    response = NewInstallModulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleDescRequest() (request *ModifyModuleDescRequest) {
    request = &ModifyModuleDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "ModifyModuleDesc")
    return
}

func NewModifyModuleDescResponse() (response *ModifyModuleDescResponse) {
    response = &ModifyModuleDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改模块描述
func (c *Client) ModifyModuleDesc(request *ModifyModuleDescRequest) (response *ModifyModuleDescResponse, err error) {
    if request == nil {
        request = NewModifyModuleDescRequest()
    }
    response = NewModifyModuleDescResponse()
    err = c.Send(request, response)
    return
}

func NewSampleRequest() (request *SampleRequest) {
    request = &SampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "Sample")
    return
}

func NewSampleResponse() (response *SampleResponse) {
    response = &SampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 样例
func (c *Client) Sample(request *SampleRequest) (response *SampleResponse, err error) {
    if request == nil {
        request = NewSampleRequest()
    }
    response = NewSampleResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallMainAgentRequest() (request *UninstallMainAgentRequest) {
    request = &UninstallMainAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "UninstallMainAgent")
    return
}

func NewUninstallMainAgentResponse() (response *UninstallMainAgentResponse) {
    response = &UninstallMainAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 主控卸载
func (c *Client) UninstallMainAgent(request *UninstallMainAgentRequest) (response *UninstallMainAgentResponse, err error) {
    if request == nil {
        request = NewUninstallMainAgentRequest()
    }
    response = NewUninstallMainAgentResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallModulesRequest() (request *UninstallModulesRequest) {
    request = &UninstallModulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("opagentcontrol", APIVersion, "UninstallModules")
    return
}

func NewUninstallModulesResponse() (response *UninstallModulesResponse) {
    response = &UninstallModulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 卸载模块
func (c *Client) UninstallModules(request *UninstallModulesRequest) (response *UninstallModulesResponse, err error) {
    if request == nil {
        request = NewUninstallModulesRequest()
    }
    response = NewUninstallModulesResponse()
    err = c.Send(request, response)
    return
}
