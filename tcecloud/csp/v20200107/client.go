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

package v20200107

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-01-07"

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


func NewDeleteMultipleObjectRequest() (request *DeleteMultipleObjectRequest) {
    request = &DeleteMultipleObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "DeleteMultipleObject")
    return
}

func NewDeleteMultipleObjectResponse() (response *DeleteMultipleObjectResponse) {
    response = &DeleteMultipleObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除多个对象
func (c *Client) DeleteMultipleObject(request *DeleteMultipleObjectRequest) (response *DeleteMultipleObjectResponse, err error) {
    if request == nil {
        request = NewDeleteMultipleObjectRequest()
    }
    response = NewDeleteMultipleObjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteObjectRequest() (request *DeleteObjectRequest) {
    request = &DeleteObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "DeleteObject")
    return
}

func NewDeleteObjectResponse() (response *DeleteObjectResponse) {
    response = &DeleteObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DELETE Object 接口请求可以在 COS 的存储桶中将一个对象（Object）删除。该操作需要请求者对存储桶有 WRITE 权限。
func (c *Client) DeleteObject(request *DeleteObjectRequest) (response *DeleteObjectResponse, err error) {
    if request == nil {
        request = NewDeleteObjectRequest()
    }
    response = NewDeleteObjectResponse()
    err = c.Send(request, response)
    return
}

func NewGetBucketRefererRequest() (request *GetBucketRefererRequest) {
    request = &GetBucketRefererRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetBucketReferer")
    return
}

func NewGetBucketRefererResponse() (response *GetBucketRefererResponse) {
    response = &GetBucketRefererResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取存储桶 Referer 白名单或者黑名单
func (c *Client) GetBucketReferer(request *GetBucketRefererRequest) (response *GetBucketRefererResponse, err error) {
    if request == nil {
        request = NewGetBucketRefererRequest()
    }
    response = NewGetBucketRefererResponse()
    err = c.Send(request, response)
    return
}

func NewGetOverviewRequest() (request *GetOverviewRequest) {
    request = &GetOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetOverview")
    return
}

func NewGetOverviewResponse() (response *GetOverviewResponse) {
    response = &GetOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户的概览信息
func (c *Client) GetOverview(request *GetOverviewRequest) (response *GetOverviewResponse, err error) {
    if request == nil {
        request = NewGetOverviewRequest()
    }
    response = NewGetOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewGetRegionListRequest() (request *GetRegionListRequest) {
    request = &GetRegionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetRegionList")
    return
}

func NewGetRegionListResponse() (response *GetRegionListResponse) {
    response = &GetRegionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户地域列表
func (c *Client) GetRegionList(request *GetRegionListRequest) (response *GetRegionListResponse, err error) {
    if request == nil {
        request = NewGetRegionListRequest()
    }
    response = NewGetRegionListResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceRequest() (request *GetServiceRequest) {
    request = &GetServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetService")
    return
}

func NewGetServiceResponse() (response *GetServiceResponse) {
    response = &GetServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取存储桶列表
func (c *Client) GetService(request *GetServiceRequest) (response *GetServiceResponse, err error) {
    if request == nil {
        request = NewGetServiceRequest()
    }
    response = NewGetServiceResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatDayRequest() (request *GetStatDayRequest) {
    request = &GetStatDayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetStatDay")
    return
}

func NewGetStatDayResponse() (response *GetStatDayResponse) {
    response = &GetStatDayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单日单存储桶统计信息
func (c *Client) GetStatDay(request *GetStatDayRequest) (response *GetStatDayResponse, err error) {
    if request == nil {
        request = NewGetStatDayRequest()
    }
    response = NewGetStatDayResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatDaysRequest() (request *GetStatDaysRequest) {
    request = &GetStatDaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetStatDays")
    return
}

func NewGetStatDaysResponse() (response *GetStatDaysResponse) {
    response = &GetStatDaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取多日单存储桶的统计信息
func (c *Client) GetStatDays(request *GetStatDaysRequest) (response *GetStatDaysResponse, err error) {
    if request == nil {
        request = NewGetStatDaysRequest()
    }
    response = NewGetStatDaysResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatHttpDayRequest() (request *GetStatHttpDayRequest) {
    request = &GetStatHttpDayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetStatHttpDay")
    return
}

func NewGetStatHttpDayResponse() (response *GetStatHttpDayResponse) {
    response = &GetStatHttpDayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单日单存储桶的http请求统计信息
func (c *Client) GetStatHttpDay(request *GetStatHttpDayRequest) (response *GetStatHttpDayResponse, err error) {
    if request == nil {
        request = NewGetStatHttpDayRequest()
    }
    response = NewGetStatHttpDayResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatHttpDaysRequest() (request *GetStatHttpDaysRequest) {
    request = &GetStatHttpDaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetStatHttpDays")
    return
}

func NewGetStatHttpDaysResponse() (response *GetStatHttpDaysResponse) {
    response = &GetStatHttpDaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取多日单存储桶的http请求统计信息
func (c *Client) GetStatHttpDays(request *GetStatHttpDaysRequest) (response *GetStatHttpDaysResponse, err error) {
    if request == nil {
        request = NewGetStatHttpDaysRequest()
    }
    response = NewGetStatHttpDaysResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserStatRequest() (request *GetUserStatRequest) {
    request = &GetUserStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "GetUserStat")
    return
}

func NewGetUserStatResponse() (response *GetUserStatResponse) {
    response = &GetUserStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取COS用户状态
func (c *Client) GetUserStat(request *GetUserStatRequest) (response *GetUserStatResponse, err error) {
    if request == nil {
        request = NewGetUserStatRequest()
    }
    response = NewGetUserStatResponse()
    err = c.Send(request, response)
    return
}

func NewHeadObjectRequest() (request *HeadObjectRequest) {
    request = &HeadObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "HeadObject")
    return
}

func NewHeadObjectResponse() (response *HeadObjectResponse) {
    response = &HeadObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取对象属性
func (c *Client) HeadObject(request *HeadObjectRequest) (response *HeadObjectResponse, err error) {
    if request == nil {
        request = NewHeadObjectRequest()
    }
    response = NewHeadObjectResponse()
    err = c.Send(request, response)
    return
}

func NewPutBucketRequest() (request *PutBucketRequest) {
    request = &PutBucketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "PutBucket")
    return
}

func NewPutBucketResponse() (response *PutBucketResponse) {
    response = &PutBucketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建存储桶
func (c *Client) PutBucket(request *PutBucketRequest) (response *PutBucketResponse, err error) {
    if request == nil {
        request = NewPutBucketRequest()
    }
    response = NewPutBucketResponse()
    err = c.Send(request, response)
    return
}

func NewPutBucketRefererRequest() (request *PutBucketRefererRequest) {
    request = &PutBucketRefererRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("csp", APIVersion, "PutBucketReferer")
    return
}

func NewPutBucketRefererResponse() (response *PutBucketRefererResponse) {
    response = &PutBucketRefererResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为存储桶设置 Referer 白名单或者黑名单
func (c *Client) PutBucketReferer(request *PutBucketRefererRequest) (response *PutBucketRefererResponse, err error) {
    if request == nil {
        request = NewPutBucketRefererRequest()
    }
    response = NewPutBucketRefererResponse()
    err = c.Send(request, response)
    return
}
