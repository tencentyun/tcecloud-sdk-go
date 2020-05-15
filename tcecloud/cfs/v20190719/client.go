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

package v20190719

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-07-19"

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


func NewDescribeAvailableZoneInfoRequest() (request *DescribeAvailableZoneInfoRequest) {
    request = &DescribeAvailableZoneInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeAvailableZoneInfo")
    return
}

func NewDescribeAvailableZoneInfoResponse() (response *DescribeAvailableZoneInfoResponse) {
    response = &DescribeAvailableZoneInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAvailableZoneInfo）用于查询区域的可用情况。
func (c *Client) DescribeAvailableZoneInfo(request *DescribeAvailableZoneInfoRequest) (response *DescribeAvailableZoneInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableZoneInfoRequest()
    }
    response = NewDescribeAvailableZoneInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsFileSystemClientsRequest() (request *DescribeCfsFileSystemClientsRequest) {
    request = &DescribeCfsFileSystemClientsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystemClients")
    return
}

func NewDescribeCfsFileSystemClientsResponse() (response *DescribeCfsFileSystemClientsResponse) {
    response = &DescribeCfsFileSystemClientsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询挂载该文件系统的客户端。此功能需要客户端安装CFS监控插件。
func (c *Client) DescribeCfsFileSystemClients(request *DescribeCfsFileSystemClientsRequest) (response *DescribeCfsFileSystemClientsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsFileSystemClientsRequest()
    }
    response = NewDescribeCfsFileSystemClientsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsFileSystemsRequest() (request *DescribeCfsFileSystemsRequest) {
    request = &DescribeCfsFileSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystems")
    return
}

func NewDescribeCfsFileSystemsResponse() (response *DescribeCfsFileSystemsResponse) {
    response = &DescribeCfsFileSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCfsFileSystems）用于查询文件系统
func (c *Client) DescribeCfsFileSystems(request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsFileSystemsRequest()
    }
    response = NewDescribeCfsFileSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsPGroupsRequest() (request *DescribeCfsPGroupsRequest) {
    request = &DescribeCfsPGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsPGroups")
    return
}

func NewDescribeCfsPGroupsResponse() (response *DescribeCfsPGroupsResponse) {
    response = &DescribeCfsPGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCfsPGroups）用于查询权限组列表。
func (c *Client) DescribeCfsPGroups(request *DescribeCfsPGroupsRequest) (response *DescribeCfsPGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsPGroupsRequest()
    }
    response = NewDescribeCfsPGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsRulesRequest() (request *DescribeCfsRulesRequest) {
    request = &DescribeCfsRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsRules")
    return
}

func NewDescribeCfsRulesResponse() (response *DescribeCfsRulesResponse) {
    response = &DescribeCfsRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCfsRules）用于查询权限组规则列表。
func (c *Client) DescribeCfsRules(request *DescribeCfsRulesRequest) (response *DescribeCfsRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCfsRulesRequest()
    }
    response = NewDescribeCfsRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsServiceStatusRequest() (request *DescribeCfsServiceStatusRequest) {
    request = &DescribeCfsServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsServiceStatus")
    return
}

func NewDescribeCfsServiceStatusResponse() (response *DescribeCfsServiceStatusResponse) {
    response = &DescribeCfsServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCfsServiceStatus）用于查询用户使用CFS的服务状态。
func (c *Client) DescribeCfsServiceStatus(request *DescribeCfsServiceStatusRequest) (response *DescribeCfsServiceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCfsServiceStatusRequest()
    }
    response = NewDescribeCfsServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
    request = &DescribeMountTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeMountTargets")
    return
}

func NewDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
    response = &DescribeMountTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeMountTargets）用于查询文件系统挂载点信息
func (c *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeMountTargetsRequest()
    }
    response = NewDescribeMountTargetsResponse()
    err = c.Send(request, response)
    return
}
