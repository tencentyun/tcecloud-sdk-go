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

package v20200217

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-17"

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


func NewCompanyCertificateApplyRequest() (request *CompanyCertificateApplyRequest) {
    request = &CompanyCertificateApplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("certify", APIVersion, "CompanyCertificateApply")
    return
}

func NewCompanyCertificateApplyResponse() (response *CompanyCertificateApplyResponse) {
    response = &CompanyCertificateApplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// （企业）资质审核申请接口
func (c *Client) CompanyCertificateApply(request *CompanyCertificateApplyRequest) (response *CompanyCertificateApplyResponse, err error) {
    if request == nil {
        request = NewCompanyCertificateApplyRequest()
    }
    response = NewCompanyCertificateApplyResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsoleCertificateDetailRequest() (request *GetConsoleCertificateDetailRequest) {
    request = &GetConsoleCertificateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("certify", APIVersion, "GetConsoleCertificateDetail")
    return
}

func NewGetConsoleCertificateDetailResponse() (response *GetConsoleCertificateDetailResponse) {
    response = &GetConsoleCertificateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 租户端资质认证详情查询
func (c *Client) GetConsoleCertificateDetail(request *GetConsoleCertificateDetailRequest) (response *GetConsoleCertificateDetailResponse, err error) {
    if request == nil {
        request = NewGetConsoleCertificateDetailRequest()
    }
    response = NewGetConsoleCertificateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsoleCertificateStatusRequest() (request *GetConsoleCertificateStatusRequest) {
    request = &GetConsoleCertificateStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("certify", APIVersion, "GetConsoleCertificateStatus")
    return
}

func NewGetConsoleCertificateStatusResponse() (response *GetConsoleCertificateStatusResponse) {
    response = &GetConsoleCertificateStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 资质审核开启查询
func (c *Client) GetConsoleCertificateStatus(request *GetConsoleCertificateStatusRequest) (response *GetConsoleCertificateStatusResponse, err error) {
    if request == nil {
        request = NewGetConsoleCertificateStatusRequest()
    }
    response = NewGetConsoleCertificateStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsoleCurrentCertRequest() (request *GetConsoleCurrentCertRequest) {
    request = &GetConsoleCurrentCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("certify", APIVersion, "GetConsoleCurrentCert")
    return
}

func NewGetConsoleCurrentCertResponse() (response *GetConsoleCurrentCertResponse) {
    response = &GetConsoleCurrentCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 资质认证结果查询
func (c *Client) GetConsoleCurrentCert(request *GetConsoleCurrentCertRequest) (response *GetConsoleCurrentCertResponse, err error) {
    if request == nil {
        request = NewGetConsoleCurrentCertRequest()
    }
    response = NewGetConsoleCurrentCertResponse()
    err = c.Send(request, response)
    return
}

func NewPersonCertificateApplyRequest() (request *PersonCertificateApplyRequest) {
    request = &PersonCertificateApplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("certify", APIVersion, "PersonCertificateApply")
    return
}

func NewPersonCertificateApplyResponse() (response *PersonCertificateApplyResponse) {
    response = &PersonCertificateApplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// （个人）资质审核申请接口
func (c *Client) PersonCertificateApply(request *PersonCertificateApplyRequest) (response *PersonCertificateApplyResponse, err error) {
    if request == nil {
        request = NewPersonCertificateApplyRequest()
    }
    response = NewPersonCertificateApplyResponse()
    err = c.Send(request, response)
    return
}

func NewQueryProvinceAndCityInfoRequest() (request *QueryProvinceAndCityInfoRequest) {
    request = &QueryProvinceAndCityInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("certify", APIVersion, "QueryProvinceAndCityInfo")
    return
}

func NewQueryProvinceAndCityInfoResponse() (response *QueryProvinceAndCityInfoResponse) {
    response = &QueryProvinceAndCityInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取省份和城市信息
func (c *Client) QueryProvinceAndCityInfo(request *QueryProvinceAndCityInfoRequest) (response *QueryProvinceAndCityInfoResponse, err error) {
    if request == nil {
        request = NewQueryProvinceAndCityInfoRequest()
    }
    response = NewQueryProvinceAndCityInfoResponse()
    err = c.Send(request, response)
    return
}
