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

package v20190606

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-06-06"

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


func NewArgusReportRequest() (request *ArgusReportRequest) {
    request = &ArgusReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ArgusReport")
    return
}

func NewArgusReportResponse() (response *ArgusReportResponse) {
    response = &ArgusReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//     （1）提供Line Protocol和Json两种数据格式上报。
//     （2）单用户上报请求数限制为100次/s，单个请求最大包体为256K，超过大小返回错误，最多上报150条，超过限制的数据自动丢弃。
//     （3）上报接口校验上报的时间戳参数，30天前和一天之后的数据会被丢弃。
//     （4）忽略上报的无效维度和指标，未上报的维度按空字符串进行存储。
//     （5）支持单个请求同时上报多个命名空间，上报的错误命名空间数据会进行丢弃，返回结果中显示上报成功数和失败数。
func (c *Client) ArgusReport(request *ArgusReportRequest) (response *ArgusReportResponse, err error) {
    if request == nil {
        request = NewArgusReportRequest()
    }
    response = NewArgusReportResponse()
    err = c.Send(request, response)
    return
}

func NewReportDataRequest() (request *ReportDataRequest) {
    request = &ReportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsm", APIVersion, "ReportData")
    return
}

func NewReportDataResponse() (response *ReportDataResponse) {
    response = &ReportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//     （1）提供Line Protocol和Json两种数据格式上报。
//     （2）单用户上报请求数限制为100次/s，单个请求最大包体为256K，超过大小返回错误，最多上报150条，超过限制的数据自动丢弃。
//     （3）上报接口校验上报的时间戳参数，30天前和一天之后的数据会被丢弃。
//     （4）忽略上报的无效维度和指标，未上报的维度按空字符串进行存储。
//     （5）支持单个请求同时上报多个命名空间，上报的错误命名空间数据会进行丢弃，返回结果中显示上报成功数和失败数。
func (c *Client) ReportData(request *ReportDataRequest) (response *ReportDataResponse, err error) {
    if request == nil {
        request = NewReportDataRequest()
    }
    response = NewReportDataResponse()
    err = c.Send(request, response)
    return
}
