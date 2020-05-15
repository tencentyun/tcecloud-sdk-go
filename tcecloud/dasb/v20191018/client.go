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

package v20191018

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-10-18"

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


func NewDescribeDasbImageIdsRequest() (request *DescribeDasbImageIdsRequest) {
    request = &DescribeDasbImageIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbImageIds")
    return
}

func NewDescribeDasbImageIdsResponse() (response *DescribeDasbImageIdsResponse) {
    response = &DescribeDasbImageIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取镜像列表
func (c *Client) DescribeDasbImageIds(request *DescribeDasbImageIdsRequest) (response *DescribeDasbImageIdsResponse, err error) {
    if request == nil {
        request = NewDescribeDasbImageIdsRequest()
    }
    response = NewDescribeDasbImageIdsResponse()
    err = c.Send(request, response)
    return
}
