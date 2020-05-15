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
