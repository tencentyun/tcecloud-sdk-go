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

package v2018109

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-10-9"

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


func NewOperDocCategoryRequest() (request *OperDocCategoryRequest) {
    request = &OperDocCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "OperDocCategory")
    return
}

func NewOperDocCategoryResponse() (response *OperDocCategoryResponse) {
    response = &OperDocCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperDocCategory
func (c *Client) OperDocCategory(request *OperDocCategoryRequest) (response *OperDocCategoryResponse, err error) {
    if request == nil {
        request = NewOperDocCategoryRequest()
    }
    response = NewOperDocCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewOperDocMenuRequest() (request *OperDocMenuRequest) {
    request = &OperDocMenuRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "OperDocMenu")
    return
}

func NewOperDocMenuResponse() (response *OperDocMenuResponse) {
    response = &OperDocMenuResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperDocMenu
func (c *Client) OperDocMenu(request *OperDocMenuRequest) (response *OperDocMenuResponse, err error) {
    if request == nil {
        request = NewOperDocMenuRequest()
    }
    response = NewOperDocMenuResponse()
    err = c.Send(request, response)
    return
}

func NewOperDocMenuLockRequest() (request *OperDocMenuLockRequest) {
    request = &OperDocMenuLockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "OperDocMenuLock")
    return
}

func NewOperDocMenuLockResponse() (response *OperDocMenuLockResponse) {
    response = &OperDocMenuLockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperDocMenuLock
func (c *Client) OperDocMenuLock(request *OperDocMenuLockRequest) (response *OperDocMenuLockResponse, err error) {
    if request == nil {
        request = NewOperDocMenuLockRequest()
    }
    response = NewOperDocMenuLockResponse()
    err = c.Send(request, response)
    return
}

func NewOperDocMenuStatusRequest() (request *OperDocMenuStatusRequest) {
    request = &OperDocMenuStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "OperDocMenuStatus")
    return
}

func NewOperDocMenuStatusResponse() (response *OperDocMenuStatusResponse) {
    response = &OperDocMenuStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperDocMenuStatus
func (c *Client) OperDocMenuStatus(request *OperDocMenuStatusRequest) (response *OperDocMenuStatusResponse, err error) {
    if request == nil {
        request = NewOperDocMenuStatusRequest()
    }
    response = NewOperDocMenuStatusResponse()
    err = c.Send(request, response)
    return
}

func NewOperDocProductRequest() (request *OperDocProductRequest) {
    request = &OperDocProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "OperDocProduct")
    return
}

func NewOperDocProductResponse() (response *OperDocProductResponse) {
    response = &OperDocProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperDocProduct
func (c *Client) OperDocProduct(request *OperDocProductRequest) (response *OperDocProductResponse, err error) {
    if request == nil {
        request = NewOperDocProductRequest()
    }
    response = NewOperDocProductResponse()
    err = c.Send(request, response)
    return
}

func NewOperDocProductDefaultPageRequest() (request *OperDocProductDefaultPageRequest) {
    request = &OperDocProductDefaultPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "OperDocProductDefaultPage")
    return
}

func NewOperDocProductDefaultPageResponse() (response *OperDocProductDefaultPageResponse) {
    response = &OperDocProductDefaultPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperDocProductDefaultPage
func (c *Client) OperDocProductDefaultPage(request *OperDocProductDefaultPageRequest) (response *OperDocProductDefaultPageResponse, err error) {
    if request == nil {
        request = NewOperDocProductDefaultPageRequest()
    }
    response = NewOperDocProductDefaultPageResponse()
    err = c.Send(request, response)
    return
}

func NewOperDocReviewRequest() (request *OperDocReviewRequest) {
    request = &OperDocReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "OperDocReview")
    return
}

func NewOperDocReviewResponse() (response *OperDocReviewResponse) {
    response = &OperDocReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperDocReview
func (c *Client) OperDocReview(request *OperDocReviewRequest) (response *OperDocReviewResponse, err error) {
    if request == nil {
        request = NewOperDocReviewRequest()
    }
    response = NewOperDocReviewResponse()
    err = c.Send(request, response)
    return
}

func NewOperDocTypeRequest() (request *OperDocTypeRequest) {
    request = &OperDocTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "OperDocType")
    return
}

func NewOperDocTypeResponse() (response *OperDocTypeResponse) {
    response = &OperDocTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperDocType 
func (c *Client) OperDocType(request *OperDocTypeRequest) (response *OperDocTypeResponse, err error) {
    if request == nil {
        request = NewOperDocTypeRequest()
    }
    response = NewOperDocTypeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocCategoryRequest() (request *QueryDocCategoryRequest) {
    request = &QueryDocCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocCategory")
    return
}

func NewQueryDocCategoryResponse() (response *QueryDocCategoryResponse) {
    response = &QueryDocCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDocCategory
func (c *Client) QueryDocCategory(request *QueryDocCategoryRequest) (response *QueryDocCategoryResponse, err error) {
    if request == nil {
        request = NewQueryDocCategoryRequest()
    }
    response = NewQueryDocCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocCategoryListRequest() (request *QueryDocCategoryListRequest) {
    request = &QueryDocCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocCategoryList")
    return
}

func NewQueryDocCategoryListResponse() (response *QueryDocCategoryListResponse) {
    response = &QueryDocCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDocCategoryList
func (c *Client) QueryDocCategoryList(request *QueryDocCategoryListRequest) (response *QueryDocCategoryListResponse, err error) {
    if request == nil {
        request = NewQueryDocCategoryListRequest()
    }
    response = NewQueryDocCategoryListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocMenuRequest() (request *QueryDocMenuRequest) {
    request = &QueryDocMenuRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocMenu")
    return
}

func NewQueryDocMenuResponse() (response *QueryDocMenuResponse) {
    response = &QueryDocMenuResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDocMenu
func (c *Client) QueryDocMenu(request *QueryDocMenuRequest) (response *QueryDocMenuResponse, err error) {
    if request == nil {
        request = NewQueryDocMenuRequest()
    }
    response = NewQueryDocMenuResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocMenuListRequest() (request *QueryDocMenuListRequest) {
    request = &QueryDocMenuListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocMenuList")
    return
}

func NewQueryDocMenuListResponse() (response *QueryDocMenuListResponse) {
    response = &QueryDocMenuListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDocMenuList
func (c *Client) QueryDocMenuList(request *QueryDocMenuListRequest) (response *QueryDocMenuListResponse, err error) {
    if request == nil {
        request = NewQueryDocMenuListRequest()
    }
    response = NewQueryDocMenuListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocProductRequest() (request *QueryDocProductRequest) {
    request = &QueryDocProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocProduct")
    return
}

func NewQueryDocProductResponse() (response *QueryDocProductResponse) {
    response = &QueryDocProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDocProduct
func (c *Client) QueryDocProduct(request *QueryDocProductRequest) (response *QueryDocProductResponse, err error) {
    if request == nil {
        request = NewQueryDocProductRequest()
    }
    response = NewQueryDocProductResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocProductDefaultPageRequest() (request *QueryDocProductDefaultPageRequest) {
    request = &QueryDocProductDefaultPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocProductDefaultPage")
    return
}

func NewQueryDocProductDefaultPageResponse() (response *QueryDocProductDefaultPageResponse) {
    response = &QueryDocProductDefaultPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDocProductDefaultPage
func (c *Client) QueryDocProductDefaultPage(request *QueryDocProductDefaultPageRequest) (response *QueryDocProductDefaultPageResponse, err error) {
    if request == nil {
        request = NewQueryDocProductDefaultPageRequest()
    }
    response = NewQueryDocProductDefaultPageResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocReviewRequest() (request *QueryDocReviewRequest) {
    request = &QueryDocReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocReview")
    return
}

func NewQueryDocReviewResponse() (response *QueryDocReviewResponse) {
    response = &QueryDocReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDocReview
func (c *Client) QueryDocReview(request *QueryDocReviewRequest) (response *QueryDocReviewResponse, err error) {
    if request == nil {
        request = NewQueryDocReviewRequest()
    }
    response = NewQueryDocReviewResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocReviewListRequest() (request *QueryDocReviewListRequest) {
    request = &QueryDocReviewListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocReviewList")
    return
}

func NewQueryDocReviewListResponse() (response *QueryDocReviewListResponse) {
    response = &QueryDocReviewListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDocReviewList
func (c *Client) QueryDocReviewList(request *QueryDocReviewListRequest) (response *QueryDocReviewListResponse, err error) {
    if request == nil {
        request = NewQueryDocReviewListRequest()
    }
    response = NewQueryDocReviewListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDocTypeRequest() (request *QueryDocTypeRequest) {
    request = &QueryDocTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("document", APIVersion, "QueryDocType")
    return
}

func NewQueryDocTypeResponse() (response *QueryDocTypeResponse) {
    response = &QueryDocTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询文档类型
func (c *Client) QueryDocType(request *QueryDocTypeRequest) (response *QueryDocTypeResponse, err error) {
    if request == nil {
        request = NewQueryDocTypeRequest()
    }
    response = NewQueryDocTypeResponse()
    err = c.Send(request, response)
    return
}
