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
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
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
