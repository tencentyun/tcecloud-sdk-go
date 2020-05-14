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

package v20190305

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-05"

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


func NewBspAudioRecognitionRequest() (request *BspAudioRecognitionRequest) {
    request = &BspAudioRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspAudioRecognition")
    return
}

func NewBspAudioRecognitionResponse() (response *BspAudioRecognitionResponse) {
    response = &BspAudioRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 识别音频是否存在恶意信息
func (c *Client) BspAudioRecognition(request *BspAudioRecognitionRequest) (response *BspAudioRecognitionResponse, err error) {
    if request == nil {
        request = NewBspAudioRecognitionRequest()
    }
    response = NewBspAudioRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewBspCloseServiceRequest() (request *BspCloseServiceRequest) {
    request = &BspCloseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspCloseService")
    return
}

func NewBspCloseServiceResponse() (response *BspCloseServiceResponse) {
    response = &BspCloseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 单个服务关闭接口
func (c *Client) BspCloseService(request *BspCloseServiceRequest) (response *BspCloseServiceResponse, err error) {
    if request == nil {
        request = NewBspCloseServiceRequest()
    }
    response = NewBspCloseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewBspImageRecognitionRequest() (request *BspImageRecognitionRequest) {
    request = &BspImageRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspImageRecognition")
    return
}

func NewBspImageRecognitionResponse() (response *BspImageRecognitionResponse) {
    response = &BspImageRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 识别图片是否存在恶意信息
func (c *Client) BspImageRecognition(request *BspImageRecognitionRequest) (response *BspImageRecognitionResponse, err error) {
    if request == nil {
        request = NewBspImageRecognitionRequest()
    }
    response = NewBspImageRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewBspOpenAfterPayServiceRequest() (request *BspOpenAfterPayServiceRequest) {
    request = &BspOpenAfterPayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspOpenAfterPayService")
    return
}

func NewBspOpenAfterPayServiceResponse() (response *BspOpenAfterPayServiceResponse) {
    response = &BspOpenAfterPayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 客户点击开通服务接口
func (c *Client) BspOpenAfterPayService(request *BspOpenAfterPayServiceRequest) (response *BspOpenAfterPayServiceResponse, err error) {
    if request == nil {
        request = NewBspOpenAfterPayServiceRequest()
    }
    response = NewBspOpenAfterPayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewBspOpenServiceRequest() (request *BspOpenServiceRequest) {
    request = &BspOpenServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspOpenService")
    return
}

func NewBspOpenServiceResponse() (response *BspOpenServiceResponse) {
    response = &BspOpenServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开通内容安全服务
func (c *Client) BspOpenService(request *BspOpenServiceRequest) (response *BspOpenServiceResponse, err error) {
    if request == nil {
        request = NewBspOpenServiceRequest()
    }
    response = NewBspOpenServiceResponse()
    err = c.Send(request, response)
    return
}

func NewBspSearchAllServiceNumRequest() (request *BspSearchAllServiceNumRequest) {
    request = &BspSearchAllServiceNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspSearchAllServiceNum")
    return
}

func NewBspSearchAllServiceNumResponse() (response *BspSearchAllServiceNumResponse) {
    response = &BspSearchAllServiceNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Top50租户
func (c *Client) BspSearchAllServiceNum(request *BspSearchAllServiceNumRequest) (response *BspSearchAllServiceNumResponse, err error) {
    if request == nil {
        request = NewBspSearchAllServiceNumRequest()
    }
    response = NewBspSearchAllServiceNumResponse()
    err = c.Send(request, response)
    return
}

func NewBspSearchAllServiceOpensRequest() (request *BspSearchAllServiceOpensRequest) {
    request = &BspSearchAllServiceOpensRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspSearchAllServiceOpens")
    return
}

func NewBspSearchAllServiceOpensResponse() (response *BspSearchAllServiceOpensResponse) {
    response = &BspSearchAllServiceOpensResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 内容安全查询所有服务开通详情
func (c *Client) BspSearchAllServiceOpens(request *BspSearchAllServiceOpensRequest) (response *BspSearchAllServiceOpensResponse, err error) {
    if request == nil {
        request = NewBspSearchAllServiceOpensRequest()
    }
    response = NewBspSearchAllServiceOpensResponse()
    err = c.Send(request, response)
    return
}

func NewBspSearchServiceOpensRequest() (request *BspSearchServiceOpensRequest) {
    request = &BspSearchServiceOpensRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspSearchServiceOpens")
    return
}

func NewBspSearchServiceOpensResponse() (response *BspSearchServiceOpensResponse) {
    response = &BspSearchServiceOpensResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 内容安全查询单个服务是否开通
func (c *Client) BspSearchServiceOpens(request *BspSearchServiceOpensRequest) (response *BspSearchServiceOpensResponse, err error) {
    if request == nil {
        request = NewBspSearchServiceOpensRequest()
    }
    response = NewBspSearchServiceOpensResponse()
    err = c.Send(request, response)
    return
}

func NewBspSereachServiceAllRequest() (request *BspSereachServiceAllRequest) {
    request = &BspSereachServiceAllRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspSereachServiceAll")
    return
}

func NewBspSereachServiceAllResponse() (response *BspSereachServiceAllResponse) {
    response = &BspSereachServiceAllResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 租户服务开通信息
func (c *Client) BspSereachServiceAll(request *BspSereachServiceAllRequest) (response *BspSereachServiceAllResponse, err error) {
    if request == nil {
        request = NewBspSereachServiceAllRequest()
    }
    response = NewBspSereachServiceAllResponse()
    err = c.Send(request, response)
    return
}

func NewBspStatisticsContentRequest() (request *BspStatisticsContentRequest) {
    request = &BspStatisticsContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspStatisticsContent")
    return
}

func NewBspStatisticsContentResponse() (response *BspStatisticsContentResponse) {
    response = &BspStatisticsContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 内容安全用量详情和恶意占比
func (c *Client) BspStatisticsContent(request *BspStatisticsContentRequest) (response *BspStatisticsContentResponse, err error) {
    if request == nil {
        request = NewBspStatisticsContentRequest()
    }
    response = NewBspStatisticsContentResponse()
    err = c.Send(request, response)
    return
}

func NewBspStatisticsNumDistributionRequest() (request *BspStatisticsNumDistributionRequest) {
    request = &BspStatisticsNumDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspStatisticsNumDistribution")
    return
}

func NewBspStatisticsNumDistributionResponse() (response *BspStatisticsNumDistributionResponse) {
    response = &BspStatisticsNumDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 内容安全服务用量分布图
func (c *Client) BspStatisticsNumDistribution(request *BspStatisticsNumDistributionRequest) (response *BspStatisticsNumDistributionResponse, err error) {
    if request == nil {
        request = NewBspStatisticsNumDistributionRequest()
    }
    response = NewBspStatisticsNumDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewBspSumUserRequestQuantityRequest() (request *BspSumUserRequestQuantityRequest) {
    request = &BspSumUserRequestQuantityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspSumUserRequestQuantity")
    return
}

func NewBspSumUserRequestQuantityResponse() (response *BspSumUserRequestQuantityResponse) {
    response = &BspSumUserRequestQuantityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 请求数分布和明细记录
func (c *Client) BspSumUserRequestQuantity(request *BspSumUserRequestQuantityRequest) (response *BspSumUserRequestQuantityResponse, err error) {
    if request == nil {
        request = NewBspSumUserRequestQuantityRequest()
    }
    response = NewBspSumUserRequestQuantityResponse()
    err = c.Send(request, response)
    return
}

func NewBspTextRecognitionRequest() (request *BspTextRecognitionRequest) {
    request = &BspTextRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("contentsecurity", APIVersion, "BspTextRecognition")
    return
}

func NewBspTextRecognitionResponse() (response *BspTextRecognitionResponse) {
    response = &BspTextRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 识别文本是否存在恶意信息
func (c *Client) BspTextRecognition(request *BspTextRecognitionRequest) (response *BspTextRecognitionResponse, err error) {
    if request == nil {
        request = NewBspTextRecognitionRequest()
    }
    response = NewBspTextRecognitionResponse()
    err = c.Send(request, response)
    return
}
