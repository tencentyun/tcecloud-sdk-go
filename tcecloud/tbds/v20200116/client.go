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

package v20200116

import (
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common"
    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
    "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-01-16"

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


func NewClusterExistsRequest() (request *ClusterExistsRequest) {
    request = &ClusterExistsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbds", APIVersion, "ClusterExists")
    return
}

func NewClusterExistsResponse() (response *ClusterExistsResponse) {
    response = &ClusterExistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) ClusterExists(request *ClusterExistsRequest) (response *ClusterExistsResponse, err error) {
    if request == nil {
        request = NewClusterExistsRequest()
    }
    response = NewClusterExistsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbds", APIVersion, "CreateCluster")
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建TBDS集群，申请CVM机器，并在此机器上部署TBDS集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewGetClustersRequest() (request *GetClustersRequest) {
    request = &GetClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbds", APIVersion, "GetClusters")
    return
}

func NewGetClustersResponse() (response *GetClustersResponse) {
    response = &GetClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) GetClusters(request *GetClustersRequest) (response *GetClustersResponse, err error) {
    if request == nil {
        request = NewGetClustersRequest()
    }
    response = NewGetClustersResponse()
    err = c.Send(request, response)
    return
}

func NewGetProgressRequest() (request *GetProgressRequest) {
    request = &GetProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbds", APIVersion, "GetProgress")
    return
}

func NewGetProgressResponse() (response *GetProgressResponse) {
    response = &GetProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 无
func (c *Client) GetProgress(request *GetProgressRequest) (response *GetProgressResponse, err error) {
    if request == nil {
        request = NewGetProgressRequest()
    }
    response = NewGetProgressResponse()
    err = c.Send(request, response)
    return
}
