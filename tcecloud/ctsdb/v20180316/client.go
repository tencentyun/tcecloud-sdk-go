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

package v20180316

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-16"

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


func NewCreateInstanceHourRequest() (request *CreateInstanceHourRequest) {
    request = &CreateInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "CreateInstanceHour")
    return
}

func NewCreateInstanceHourResponse() (response *CreateInstanceHourResponse) {
    response = &CreateInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateInstanceHour)用于创建实例。
func (c *Client) CreateInstanceHour(request *CreateInstanceHourRequest) (response *CreateInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateInstanceHourRequest()
    }
    response = NewCreateInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceMetricInfoRequest() (request *DescribeDBInstanceMetricInfoRequest) {
    request = &DescribeDBInstanceMetricInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "DescribeDBInstanceMetricInfo")
    return
}

func NewDescribeDBInstanceMetricInfoResponse() (response *DescribeDBInstanceMetricInfoResponse) {
    response = &DescribeDBInstanceMetricInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBInstanceMetricInfo）用于查询实例metric的详情信息。
func (c *Client) DescribeDBInstanceMetricInfo(request *DescribeDBInstanceMetricInfoRequest) (response *DescribeDBInstanceMetricInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceMetricInfoRequest()
    }
    response = NewDescribeDBInstanceMetricInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceMetricListRequest() (request *DescribeDBInstanceMetricListRequest) {
    request = &DescribeDBInstanceMetricListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "DescribeDBInstanceMetricList")
    return
}

func NewDescribeDBInstanceMetricListResponse() (response *DescribeDBInstanceMetricListResponse) {
    response = &DescribeDBInstanceMetricListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBInstanceMetricList）用于查询实例metric列表。
func (c *Client) DescribeDBInstanceMetricList(request *DescribeDBInstanceMetricListRequest) (response *DescribeDBInstanceMetricListResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceMetricListRequest()
    }
    response = NewDescribeDBInstanceMetricListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceMetricQueryRequest() (request *DescribeDBInstanceMetricQueryRequest) {
    request = &DescribeDBInstanceMetricQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "DescribeDBInstanceMetricQuery")
    return
}

func NewDescribeDBInstanceMetricQueryResponse() (response *DescribeDBInstanceMetricQueryResponse) {
    response = &DescribeDBInstanceMetricQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBInstanceMetricQuery）用于用户metric查询。
func (c *Client) DescribeDBInstanceMetricQuery(request *DescribeDBInstanceMetricQueryRequest) (response *DescribeDBInstanceMetricQueryResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceMetricQueryRequest()
    }
    response = NewDescribeDBInstanceMetricQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceNodesRequest() (request *DescribeDBInstanceNodesRequest) {
    request = &DescribeDBInstanceNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "DescribeDBInstanceNodes")
    return
}

func NewDescribeDBInstanceNodesResponse() (response *DescribeDBInstanceNodesResponse) {
    response = &DescribeDBInstanceNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBInstanceNodes）用于查询数据库实例节点列表。
// 使用建议：
// 1：不用每次连接CTSDB时都通过此接口查询实例节点列表进行连接；在第一次查询到实例节点列表后，业务侧可以缓存起来，之后每次连接CTSDB时使用缓存的节点列表来连接；当缓存的节点列表不可用时，再调用此接口拉取实例节点列表更新到缓存中再进行连接即可。
// 2：查询到的实例节点列表中的所有ip:port都可以使用，业务侧最好能较均衡地使用这些ip:port。
func (c *Client) DescribeDBInstanceNodes(request *DescribeDBInstanceNodesRequest) (response *DescribeDBInstanceNodesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceNodesRequest()
    }
    response = NewDescribeDBInstanceNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBInstances）用于查询数据库实例列表。
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesForIoTRequest() (request *DescribeInstancesForIoTRequest) {
    request = &DescribeInstancesForIoTRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "DescribeInstancesForIoT")
    return
}

func NewDescribeInstancesForIoTResponse() (response *DescribeInstancesForIoTResponse) {
    response = &DescribeInstancesForIoTResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// [IoT专用]查询实例列表详情
func (c *Client) DescribeInstancesForIoT(request *DescribeInstancesForIoTRequest) (response *DescribeInstancesForIoTResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesForIoTRequest()
    }
    response = NewDescribeInstancesForIoTResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSaleSpecRequest() (request *DescribeSaleSpecRequest) {
    request = &DescribeSaleSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "DescribeSaleSpec")
    return
}

func NewDescribeSaleSpecResponse() (response *DescribeSaleSpecResponse) {
    response = &DescribeSaleSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSaleSpec）用于查询售卖规格。
func (c *Client) DescribeSaleSpec(request *DescribeSaleSpecRequest) (response *DescribeSaleSpecResponse, err error) {
    if request == nil {
        request = NewDescribeSaleSpecRequest()
    }
    response = NewDescribeSaleSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSaleZoneRequest() (request *DescribeSaleZoneRequest) {
    request = &DescribeSaleZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "DescribeSaleZone")
    return
}

func NewDescribeSaleZoneResponse() (response *DescribeSaleZoneResponse) {
    response = &DescribeSaleZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSaleZone）用于查询售卖区域。
func (c *Client) DescribeSaleZone(request *DescribeSaleZoneRequest) (response *DescribeSaleZoneResponse, err error) {
    if request == nil {
        request = NewDescribeSaleZoneRequest()
    }
    response = NewDescribeSaleZoneResponse()
    err = c.Send(request, response)
    return
}

func NewInitDBInstanceRequest() (request *InitDBInstanceRequest) {
    request = &InitDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "InitDBInstance")
    return
}

func NewInitDBInstanceResponse() (response *InitDBInstanceResponse) {
    response = &InitDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（InitDBInstance）用于初始化数据库实例。
func (c *Client) InitDBInstance(request *InitDBInstanceRequest) (response *InitDBInstanceResponse, err error) {
    if request == nil {
        request = NewInitDBInstanceRequest()
    }
    response = NewInitDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "ModifyDBInstanceName")
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstanceName）用于修改数据库实例名称。
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceProjectRequest() (request *ModifyDBInstanceProjectRequest) {
    request = &ModifyDBInstanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "ModifyDBInstanceProject")
    return
}

func NewModifyDBInstanceProjectResponse() (response *ModifyDBInstanceProjectResponse) {
    response = &ModifyDBInstanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstanceProject）用于修改数据库实例项目。
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceProjectRequest()
    }
    response = NewModifyDBInstanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceUserPasswordRequest() (request *ModifyDBInstanceUserPasswordRequest) {
    request = &ModifyDBInstanceUserPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "ModifyDBInstanceUserPassword")
    return
}

func NewModifyDBInstanceUserPasswordResponse() (response *ModifyDBInstanceUserPasswordResponse) {
    response = &ModifyDBInstanceUserPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstanceUserPassword）用于修改数据库实例用户密码。
func (c *Client) ModifyDBInstanceUserPassword(request *ModifyDBInstanceUserPasswordRequest) (response *ModifyDBInstanceUserPasswordResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceUserPasswordRequest()
    }
    response = NewModifyDBInstanceUserPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceHourRequest() (request *ModifyInstanceHourRequest) {
    request = &ModifyInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "ModifyInstanceHour")
    return
}

func NewModifyInstanceHourResponse() (response *ModifyInstanceHourResponse) {
    response = &ModifyInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ModifyInstanceHour) 用于修改实例配置。
func (c *Client) ModifyInstanceHour(request *ModifyInstanceHourRequest) (response *ModifyInstanceHourResponse, err error) {
    if request == nil {
        request = NewModifyInstanceHourRequest()
    }
    response = NewModifyInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewRecycleDBInstanceRequest() (request *RecycleDBInstanceRequest) {
    request = &RecycleDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ctsdb", APIVersion, "RecycleDBInstance")
    return
}

func NewRecycleDBInstanceResponse() (response *RecycleDBInstanceResponse) {
    response = &RecycleDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RecycleDBInstance）用于回收按量计费实例。
func (c *Client) RecycleDBInstance(request *RecycleDBInstanceRequest) (response *RecycleDBInstanceResponse, err error) {
    if request == nil {
        request = NewRecycleDBInstanceRequest()
    }
    response = NewRecycleDBInstanceResponse()
    err = c.Send(request, response)
    return
}
