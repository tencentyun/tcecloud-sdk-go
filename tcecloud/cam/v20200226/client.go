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

package v20200226

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-02-26"

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


func NewCreateCASProviderRequest() (request *CreateCASProviderRequest) {
    request = &CreateCASProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreateCASProvider")
    return
}

func NewCreateCASProviderResponse() (response *CreateCASProviderResponse) {
    response = &CreateCASProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建CAS身份提供商
func (c *Client) CreateCASProvider(request *CreateCASProviderRequest) (response *CreateCASProviderResponse, err error) {
    if request == nil {
        request = NewCreateCASProviderRequest()
    }
    response = NewCreateCASProviderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasProviderRequest() (request *DescribeCasProviderRequest) {
    request = &DescribeCasProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DescribeCasProvider")
    return
}

func NewDescribeCasProviderResponse() (response *DescribeCasProviderResponse) {
    response = &DescribeCasProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取cas身份提供商信息
func (c *Client) DescribeCasProvider(request *DescribeCasProviderRequest) (response *DescribeCasProviderResponse, err error) {
    if request == nil {
        request = NewDescribeCasProviderRequest()
    }
    response = NewDescribeCasProviderResponse()
    err = c.Send(request, response)
    return
}
