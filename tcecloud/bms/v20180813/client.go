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

package v20180813

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-08-13"

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


func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
    request = &DescribeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bms", APIVersion, "DescribeDisks")
    return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
    response = &DescribeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDisks）用于查询BMS硬盘列表。
// 
// * 可以根据BMS硬盘ID等信息来查询BMS硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的BMS硬盘列表。
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    if request == nil {
        request = NewDescribeDisksRequest()
    }
    response = NewDescribeDisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlavorsRequest() (request *DescribeFlavorsRequest) {
    request = &DescribeFlavorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bms", APIVersion, "DescribeFlavors")
    return
}

func NewDescribeFlavorsResponse() (response *DescribeFlavorsResponse) {
    response = &DescribeFlavorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 显示套餐列表详情。
func (c *Client) DescribeFlavors(request *DescribeFlavorsRequest) (response *DescribeFlavorsResponse, err error) {
    if request == nil {
        request = NewDescribeFlavorsRequest()
    }
    response = NewDescribeFlavorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bms", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeInstances) 用于查询一个或多个实例的详细信息。 
// 
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

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bms", APIVersion, "ModifyInstancesAttribute")
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ModifyInstancesAttribute) 用于修改实例的属性（目前只支持修改实例的名称）。
// 
// * “实例名称”仅为方便用户自己管理之用，Tce并不以此名称作为提交工单或是进行实例管理操作的依据。
// * 支持批量操作。每次请求批量实例的上限为100。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bms", APIVersion, "RebootInstances")
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (RebootInstances) 用于重启实例。
// 
// * 只有状态为`RUNNING`的实例才可以进行此操作。
// * 接口调用成功时，实例会进入`REBOOTING`状态；重启实例成功时，实例会进入`RUNNING`状态。
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
    request.Init().WithApiInfo("bms", APIVersion, "ResetInstance")
    return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
    response = &ResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ResetInstance) 用于重装指定实例上的操作系统。
// 
// * 如果指定了`OperatingSystem`参数，则使用指定的系统重装；否则按照当前实例使用的系统进行重装。
// * 系统盘将会被格式化，并重置；请确保系统盘中无重要文件。
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
    request = &RunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bms", APIVersion, "RunInstances")
    return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
    response = &RunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (RunInstances) 用于创建一个或多个指定配置的实例。
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
    request.Init().WithApiInfo("bms", APIVersion, "StartInstances")
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (StartInstances) 用于启动一个或多个实例。
// 
// * 只有状态为`STOPPED`的实例才可以进行此操作。
// * 接口调用成功时，实例会进入`STARTING`状态；启动实例成功时，实例会进入`RUNNING`状态。
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
    request.Init().WithApiInfo("bms", APIVersion, "StopInstances")
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (StopInstances) 用于关闭一个或多个实例。
// 
// * 只有状态为`RUNNING`的实例才可以进行此操作。
// * 接口调用成功时，实例会进入`STOPPING`状态；关闭实例成功时，实例会进入`STOPPED`状态。
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bms", APIVersion, "TerminateInstances")
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (TerminateInstances) 用于主动退还实例。
// 
// * 不再使用的实例，可通过本接口主动退还。
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
