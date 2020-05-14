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

package v20190107

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-07"

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


func NewActiveIsolatedInstanceRequest() (request *ActiveIsolatedInstanceRequest) {
    request = &ActiveIsolatedInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "ActiveIsolatedInstance")
    return
}

func NewActiveIsolatedInstanceResponse() (response *ActiveIsolatedInstanceResponse) {
    response = &ActiveIsolatedInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ActiveIsolatedInstance) 用于恢复实例。
func (c *Client) ActiveIsolatedInstance(request *ActiveIsolatedInstanceRequest) (response *ActiveIsolatedInstanceResponse, err error) {
    if request == nil {
        request = NewActiveIsolatedInstanceRequest()
    }
    response = NewActiveIsolatedInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewActiveIsolatedPGInstanceRequest() (request *ActiveIsolatedPGInstanceRequest) {
    request = &ActiveIsolatedPGInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "ActiveIsolatedPGInstance")
    return
}

func NewActiveIsolatedPGInstanceResponse() (response *ActiveIsolatedPGInstanceResponse) {
    response = &ActiveIsolatedPGInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ActiveIsolatedPGInstance) 用于恢复Postgres实例。
func (c *Client) ActiveIsolatedPGInstance(request *ActiveIsolatedPGInstanceRequest) (response *ActiveIsolatedPGInstanceResponse, err error) {
    if request == nil {
        request = NewActiveIsolatedPGInstanceRequest()
    }
    response = NewActiveIsolatedPGInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceHourRequest() (request *CreateInstanceHourRequest) {
    request = &CreateInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "CreateInstanceHour")
    return
}

func NewCreateInstanceHourResponse() (response *CreateInstanceHourResponse) {
    response = &CreateInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (CreateInstanceHour) 用于创建实例。
func (c *Client) CreateInstanceHour(request *CreateInstanceHourRequest) (response *CreateInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateInstanceHourRequest()
    }
    response = NewCreateInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePGInstanceHourRequest() (request *CreatePGInstanceHourRequest) {
    request = &CreatePGInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "CreatePGInstanceHour")
    return
}

func NewCreatePGInstanceHourResponse() (response *CreatePGInstanceHourResponse) {
    response = &CreatePGInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (CreatePGInstanceHour) 用于创建Postgres实例。
func (c *Client) CreatePGInstanceHour(request *CreatePGInstanceHourRequest) (response *CreatePGInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreatePGInstanceHourRequest()
    }
    response = NewCreatePGInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePGReadOnlyVipRequest() (request *CreatePGReadOnlyVipRequest) {
    request = &CreatePGReadOnlyVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "CreatePGReadOnlyVip")
    return
}

func NewCreatePGReadOnlyVipResponse() (response *CreatePGReadOnlyVipResponse) {
    response = &CreatePGReadOnlyVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreatePGReadOnlyVip）用于给postgres实例添加只读VIP，同一实例只允许存在一个只读VIP，VIP可由用户自己指定或系统自行指定。
func (c *Client) CreatePGReadOnlyVip(request *CreatePGReadOnlyVipRequest) (response *CreatePGReadOnlyVipResponse, err error) {
    if request == nil {
        request = NewCreatePGReadOnlyVipRequest()
    }
    response = NewCreatePGReadOnlyVipResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePGReadOnlyVipRequest() (request *DeletePGReadOnlyVipRequest) {
    request = &DeletePGReadOnlyVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DeletePGReadOnlyVip")
    return
}

func NewDeletePGReadOnlyVipResponse() (response *DeletePGReadOnlyVipResponse) {
    response = &DeletePGReadOnlyVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于删除指定postgres实例的只读VIP。
func (c *Client) DeletePGReadOnlyVip(request *DeletePGReadOnlyVipRequest) (response *DeletePGReadOnlyVipResponse, err error) {
    if request == nil {
        request = NewDeletePGReadOnlyVipRequest()
    }
    response = NewDeletePGReadOnlyVipResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableCharsetRequest() (request *DescribeAvailableCharsetRequest) {
    request = &DescribeAvailableCharsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableCharset")
    return
}

func NewDescribeAvailableCharsetResponse() (response *DescribeAvailableCharsetResponse) {
    response = &DescribeAvailableCharsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeAvailableCharset) 用于查询创建实例时可选择的字符集。
func (c *Client) DescribeAvailableCharset(request *DescribeAvailableCharsetRequest) (response *DescribeAvailableCharsetResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableCharsetRequest()
    }
    response = NewDescribeAvailableCharsetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableEngineVersionsRequest() (request *DescribeAvailableEngineVersionsRequest) {
    request = &DescribeAvailableEngineVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableEngineVersions")
    return
}

func NewDescribeAvailableEngineVersionsResponse() (response *DescribeAvailableEngineVersionsResponse) {
    response = &DescribeAvailableEngineVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeAvailableEngineVersions) 用于查询创建实例时可选择的引擎版本。
func (c *Client) DescribeAvailableEngineVersions(request *DescribeAvailableEngineVersionsRequest) (response *DescribeAvailableEngineVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableEngineVersionsRequest()
    }
    response = NewDescribeAvailableEngineVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableZoneConfigRequest() (request *DescribeAvailableZoneConfigRequest) {
    request = &DescribeAvailableZoneConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableZoneConfig")
    return
}

func NewDescribeAvailableZoneConfigResponse() (response *DescribeAvailableZoneConfigResponse) {
    response = &DescribeAvailableZoneConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeAvailableZoneConfig)用于查询云数据库可售卖的地域和可用区信息。
func (c *Client) DescribeAvailableZoneConfig(request *DescribeAvailableZoneConfigRequest) (response *DescribeAvailableZoneConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableZoneConfigRequest()
    }
    response = NewDescribeAvailableZoneConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableZonesRequest() (request *DescribeAvailableZonesRequest) {
    request = &DescribeAvailableZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableZones")
    return
}

func NewDescribeAvailableZonesResponse() (response *DescribeAvailableZonesResponse) {
    response = &DescribeAvailableZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于获取指定资源池下的可用区信息
func (c *Client) DescribeAvailableZones(request *DescribeAvailableZonesRequest) (response *DescribeAvailableZonesResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableZonesRequest()
    }
    response = NewDescribeAvailableZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDetailsRequest() (request *DescribeInstanceDetailsRequest) {
    request = &DescribeInstanceDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribeInstanceDetails")
    return
}

func NewDescribeInstanceDetailsResponse() (response *DescribeInstanceDetailsResponse) {
    response = &DescribeInstanceDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeInstanceDetails) 用于查询单个实例的详情。
func (c *Client) DescribeInstanceDetails(request *DescribeInstanceDetailsRequest) (response *DescribeInstanceDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDetailsRequest()
    }
    response = NewDescribeInstanceDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTaskStatusRequest() (request *DescribeInstanceTaskStatusRequest) {
    request = &DescribeInstanceTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribeInstanceTaskStatus")
    return
}

func NewDescribeInstanceTaskStatusResponse() (response *DescribeInstanceTaskStatusResponse) {
    response = &DescribeInstanceTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于查询实力相关的异步流程任务的执行状态，可获知任务是否失败或成功
func (c *Client) DescribeInstanceTaskStatus(request *DescribeInstanceTaskStatusRequest) (response *DescribeInstanceTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTaskStatusRequest()
    }
    response = NewDescribeInstanceTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribeInstances")
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
// * 可以根据实例ID、实例名称或者实例计费模式等信息来查询实例的详细信息。过滤信息详细请见过滤器Filter。
// * 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的实例。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePGAvailableZoneConfigRequest() (request *DescribePGAvailableZoneConfigRequest) {
    request = &DescribePGAvailableZoneConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribePGAvailableZoneConfig")
    return
}

func NewDescribePGAvailableZoneConfigResponse() (response *DescribePGAvailableZoneConfigResponse) {
    response = &DescribePGAvailableZoneConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribePGAvailableZoneConfig)用于查询云数据库Postgres的可售卖地域和可用区信息。
func (c *Client) DescribePGAvailableZoneConfig(request *DescribePGAvailableZoneConfigRequest) (response *DescribePGAvailableZoneConfigResponse, err error) {
    if request == nil {
        request = NewDescribePGAvailableZoneConfigRequest()
    }
    response = NewDescribePGAvailableZoneConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePGInstanceDetailsRequest() (request *DescribePGInstanceDetailsRequest) {
    request = &DescribePGInstanceDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribePGInstanceDetails")
    return
}

func NewDescribePGInstanceDetailsResponse() (response *DescribePGInstanceDetailsResponse) {
    response = &DescribePGInstanceDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribePGInstanceDetails) 用于查询单个Postgres实例的详情。
func (c *Client) DescribePGInstanceDetails(request *DescribePGInstanceDetailsRequest) (response *DescribePGInstanceDetailsResponse, err error) {
    if request == nil {
        request = NewDescribePGInstanceDetailsRequest()
    }
    response = NewDescribePGInstanceDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePGInstancesRequest() (request *DescribePGInstancesRequest) {
    request = &DescribePGInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribePGInstances")
    return
}

func NewDescribePGInstancesResponse() (response *DescribePGInstancesResponse) {
    response = &DescribePGInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribePGInstances) 用于查询一个或多个Postgres实例的详细信息。
func (c *Client) DescribePGInstances(request *DescribePGInstancesRequest) (response *DescribePGInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePGInstancesRequest()
    }
    response = NewDescribePGInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcePoolsRequest() (request *DescribeResourcePoolsRequest) {
    request = &DescribeResourcePoolsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "DescribeResourcePools")
    return
}

func NewDescribeResourcePoolsResponse() (response *DescribeResourcePoolsResponse) {
    response = &DescribeResourcePoolsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于在postgres购买页，查询可部署实例的资源池，能看到的资源池个数与该用户的资源池查看权限有关，权限可在运营端修改
func (c *Client) DescribeResourcePools(request *DescribeResourcePoolsRequest) (response *DescribeResourcePoolsResponse, err error) {
    if request == nil {
        request = NewDescribeResourcePoolsRequest()
    }
    response = NewDescribeResourcePoolsResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateInstanceRequest() (request *IsolateInstanceRequest) {
    request = &IsolateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "IsolateInstance")
    return
}

func NewIsolateInstanceResponse() (response *IsolateInstanceResponse) {
    response = &IsolateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (IsolateInstance) 用于隔离实例。
func (c *Client) IsolateInstance(request *IsolateInstanceRequest) (response *IsolateInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateInstanceRequest()
    }
    response = NewIsolateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewIsolatePGInstanceRequest() (request *IsolatePGInstanceRequest) {
    request = &IsolatePGInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "IsolatePGInstance")
    return
}

func NewIsolatePGInstanceResponse() (response *IsolatePGInstanceResponse) {
    response = &IsolatePGInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (IsolatePGInstance) 用于隔离Postgres实例。
func (c *Client) IsolatePGInstance(request *IsolatePGInstanceRequest) (response *IsolatePGInstanceResponse, err error) {
    if request == nil {
        request = NewIsolatePGInstanceRequest()
    }
    response = NewIsolatePGInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
    request = &ModifyInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "ModifyInstanceName")
    return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
    response = &ModifyInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyInstanceName）用于修改云数据的名称。
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNameRequest()
    }
    response = NewModifyInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseIsolatedInstanceRequest() (request *ReleaseIsolatedInstanceRequest) {
    request = &ReleaseIsolatedInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "ReleaseIsolatedInstance")
    return
}

func NewReleaseIsolatedInstanceResponse() (response *ReleaseIsolatedInstanceResponse) {
    response = &ReleaseIsolatedInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ReleaseIsolatedInstance) 用于销毁单个已隔离实例。
func (c *Client) ReleaseIsolatedInstance(request *ReleaseIsolatedInstanceRequest) (response *ReleaseIsolatedInstanceResponse, err error) {
    if request == nil {
        request = NewReleaseIsolatedInstanceRequest()
    }
    response = NewReleaseIsolatedInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbase", APIVersion, "ResetAccountPassword")
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ResetAccountPassword）用于重置云数据库账号的密码。
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}
