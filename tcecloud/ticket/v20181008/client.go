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

package v20181008

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-10-08"

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


func NewCountTicketRequest() (request *CountTicketRequest) {
    request = &CountTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "CountTicket")
    return
}

func NewCountTicketResponse() (response *CountTicketResponse) {
    response = &CountTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 工单数量
func (c *Client) CountTicket(request *CountTicketRequest) (response *CountTicketResponse, err error) {
    if request == nil {
        request = NewCountTicketRequest()
    }
    response = NewCountTicketResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTicketRequest() (request *CreateTicketRequest) {
    request = &CreateTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "CreateTicket")
    return
}

func NewCreateTicketResponse() (response *CreateTicketResponse) {
    response = &CreateTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建工单
func (c *Client) CreateTicket(request *CreateTicketRequest) (response *CreateTicketResponse, err error) {
    if request == nil {
        request = NewCreateTicketRequest()
    }
    response = NewCreateTicketResponse()
    err = c.Send(request, response)
    return
}

func NewCustomerCancelTicketRequest() (request *CustomerCancelTicketRequest) {
    request = &CustomerCancelTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "CustomerCancelTicket")
    return
}

func NewCustomerCancelTicketResponse() (response *CustomerCancelTicketResponse) {
    response = &CustomerCancelTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 客户撤销工单
func (c *Client) CustomerCancelTicket(request *CustomerCancelTicketRequest) (response *CustomerCancelTicketResponse, err error) {
    if request == nil {
        request = NewCustomerCancelTicketRequest()
    }
    response = NewCustomerCancelTicketResponse()
    err = c.Send(request, response)
    return
}

func NewCustomerFinishTicketRequest() (request *CustomerFinishTicketRequest) {
    request = &CustomerFinishTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "CustomerFinishTicket")
    return
}

func NewCustomerFinishTicketResponse() (response *CustomerFinishTicketResponse) {
    response = &CustomerFinishTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 客户结单
func (c *Client) CustomerFinishTicket(request *CustomerFinishTicketRequest) (response *CustomerFinishTicketResponse, err error) {
    if request == nil {
        request = NewCustomerFinishTicketRequest()
    }
    response = NewCustomerFinishTicketResponse()
    err = c.Send(request, response)
    return
}

func NewCustomerRespondRequest() (request *CustomerRespondRequest) {
    request = &CustomerRespondRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "CustomerRespond")
    return
}

func NewCustomerRespondResponse() (response *CustomerRespondResponse) {
    response = &CustomerRespondResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 客户回复
func (c *Client) CustomerRespond(request *CustomerRespondRequest) (response *CustomerRespondResponse, err error) {
    if request == nil {
        request = NewCustomerRespondRequest()
    }
    response = NewCustomerRespondResponse()
    err = c.Send(request, response)
    return
}

func NewCustomerUrgeRequest() (request *CustomerUrgeRequest) {
    request = &CustomerUrgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "CustomerUrge")
    return
}

func NewCustomerUrgeResponse() (response *CustomerUrgeResponse) {
    response = &CustomerUrgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 客户催单
func (c *Client) CustomerUrge(request *CustomerUrgeRequest) (response *CustomerUrgeResponse, err error) {
    if request == nil {
        request = NewCustomerUrgeRequest()
    }
    response = NewCustomerUrgeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTicketRequest() (request *DescribeTicketRequest) {
    request = &DescribeTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "DescribeTicket")
    return
}

func NewDescribeTicketResponse() (response *DescribeTicketResponse) {
    response = &DescribeTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 工单详情 
func (c *Client) DescribeTicket(request *DescribeTicketRequest) (response *DescribeTicketResponse, err error) {
    if request == nil {
        request = NewDescribeTicketRequest()
    }
    response = NewDescribeTicketResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTicketsRequest() (request *DescribeTicketsRequest) {
    request = &DescribeTicketsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "DescribeTickets")
    return
}

func NewDescribeTicketsResponse() (response *DescribeTicketsResponse) {
    response = &DescribeTicketsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 工单列表
func (c *Client) DescribeTickets(request *DescribeTicketsRequest) (response *DescribeTicketsResponse, err error) {
    if request == nil {
        request = NewDescribeTicketsRequest()
    }
    response = NewDescribeTicketsResponse()
    err = c.Send(request, response)
    return
}

func NewGetCosBucketRequest() (request *GetCosBucketRequest) {
    request = &GetCosBucketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "GetCosBucket")
    return
}

func NewGetCosBucketResponse() (response *GetCosBucketResponse) {
    response = &GetCosBucketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCosBucket
func (c *Client) GetCosBucket(request *GetCosBucketRequest) (response *GetCosBucketResponse, err error) {
    if request == nil {
        request = NewGetCosBucketRequest()
    }
    response = NewGetCosBucketResponse()
    err = c.Send(request, response)
    return
}

func NewGetCosParamsByBucketRequest() (request *GetCosParamsByBucketRequest) {
    request = &GetCosParamsByBucketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "GetCosParamsByBucket")
    return
}

func NewGetCosParamsByBucketResponse() (response *GetCosParamsByBucketResponse) {
    response = &GetCosParamsByBucketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过bucket和区域获取COS参数
func (c *Client) GetCosParamsByBucket(request *GetCosParamsByBucketRequest) (response *GetCosParamsByBucketResponse, err error) {
    if request == nil {
        request = NewGetCosParamsByBucketRequest()
    }
    response = NewGetCosParamsByBucketResponse()
    err = c.Send(request, response)
    return
}

func NewGetTicketCategoryRequest() (request *GetTicketCategoryRequest) {
    request = &GetTicketCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "GetTicketCategory")
    return
}

func NewGetTicketCategoryResponse() (response *GetTicketCategoryResponse) {
    response = &GetTicketCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 租户端获取工单分类
func (c *Client) GetTicketCategory(request *GetTicketCategoryRequest) (response *GetTicketCategoryResponse, err error) {
    if request == nil {
        request = NewGetTicketCategoryRequest()
    }
    response = NewGetTicketCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewGetTicketCategoryByLevelRequest() (request *GetTicketCategoryByLevelRequest) {
    request = &GetTicketCategoryByLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ticket", APIVersion, "GetTicketCategoryByLevel")
    return
}

func NewGetTicketCategoryByLevelResponse() (response *GetTicketCategoryByLevelResponse) {
    response = &GetTicketCategoryByLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 租户端通过分类Id获取三级分类和对应的表单字段
func (c *Client) GetTicketCategoryByLevel(request *GetTicketCategoryByLevelRequest) (response *GetTicketCategoryByLevelResponse, err error) {
    if request == nil {
        request = NewGetTicketCategoryByLevelRequest()
    }
    response = NewGetTicketCategoryByLevelResponse()
    err = c.Send(request, response)
    return
}
