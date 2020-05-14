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

package v20191021

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-21"

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


func NewCreateEnterpriseRequest() (request *CreateEnterpriseRequest) {
    request = &CreateEnterpriseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "CreateEnterprise")
    return
}

func NewCreateEnterpriseResponse() (response *CreateEnterpriseResponse) {
    response = &CreateEnterpriseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建企业。
func (c *Client) CreateEnterprise(request *CreateEnterpriseRequest) (response *CreateEnterpriseResponse, err error) {
    if request == nil {
        request = NewCreateEnterpriseRequest()
    }
    response = NewCreateEnterpriseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateJobRequest() (request *CreateJobRequest) {
    request = &CreateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "CreateJob")
    return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
    response = &CreateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建指定项目的 CI 任务
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
    if request == nil {
        request = NewCreateJobRequest()
    }
    response = NewCreateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "CreateProject")
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建项目。
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBuildRequest() (request *DeleteBuildRequest) {
    request = &DeleteBuildRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DeleteBuild")
    return
}

func NewDeleteBuildResponse() (response *DeleteBuildResponse) {
    response = &DeleteBuildResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除单个构建。
func (c *Client) DeleteBuild(request *DeleteBuildRequest) (response *DeleteBuildResponse, err error) {
    if request == nil {
        request = NewDeleteBuildRequest()
    }
    response = NewDeleteBuildResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobRequest() (request *DeleteJobRequest) {
    request = &DeleteJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DeleteJob")
    return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
    response = &DeleteJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除单个CI 任务。
func (c *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    if request == nil {
        request = NewDeleteJobRequest()
    }
    response = NewDeleteJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBuildRequest() (request *DescribeBuildRequest) {
    request = &DescribeBuildRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeBuild")
    return
}

func NewDescribeBuildResponse() (response *DescribeBuildResponse) {
    response = &DescribeBuildResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询单个构建。
func (c *Client) DescribeBuild(request *DescribeBuildRequest) (response *DescribeBuildResponse, err error) {
    if request == nil {
        request = NewDescribeBuildRequest()
    }
    response = NewDescribeBuildResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurrentUserRequest() (request *DescribeCurrentUserRequest) {
    request = &DescribeCurrentUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeCurrentUser")
    return
}

func NewDescribeCurrentUserResponse() (response *DescribeCurrentUserResponse) {
    response = &DescribeCurrentUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询当前用户信息
func (c *Client) DescribeCurrentUser(request *DescribeCurrentUserRequest) (response *DescribeCurrentUserResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentUserRequest()
    }
    response = NewDescribeCurrentUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeJobs")
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询单个项目下的所有CI 任务。
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBuildRequest() (request *DescribeListBuildRequest) {
    request = &DescribeListBuildRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeListBuild")
    return
}

func NewDescribeListBuildResponse() (response *DescribeListBuildResponse) {
    response = &DescribeListBuildResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询构建列表。
func (c *Client) DescribeListBuild(request *DescribeListBuildRequest) (response *DescribeListBuildResponse, err error) {
    if request == nil {
        request = NewDescribeListBuildRequest()
    }
    response = NewDescribeListBuildResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOneJobRequest() (request *DescribeOneJobRequest) {
    request = &DescribeOneJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeOneJob")
    return
}

func NewDescribeOneJobResponse() (response *DescribeOneJobResponse) {
    response = &DescribeOneJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询单个CI 任务。
func (c *Client) DescribeOneJob(request *DescribeOneJobRequest) (response *DescribeOneJobResponse, err error) {
    if request == nil {
        request = NewDescribeOneJobRequest()
    }
    response = NewDescribeOneJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackageListRequest() (request *DescribePackageListRequest) {
    request = &DescribePackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribePackageList")
    return
}

func NewDescribePackageListResponse() (response *DescribePackageListResponse) {
    response = &DescribePackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询包列表。
func (c *Client) DescribePackageList(request *DescribePackageListRequest) (response *DescribePackageListResponse, err error) {
    if request == nil {
        request = NewDescribePackageListRequest()
    }
    response = NewDescribePackageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
    request = &DescribeProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeProject")
    return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
    response = &DescribeProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询项目信息。
// 
// 可使用项目 ID 或项目名称来获取单个项目信息。
// 如果二者都填默认使用项目 ID 查询。
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    if request == nil {
        request = NewDescribeProjectRequest()
    }
    response = NewDescribeProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeProjects")
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeProjects) 查询项目列表。
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryListRequest() (request *DescribeRepositoryListRequest) {
    request = &DescribeRepositoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeRepositoryList")
    return
}

func NewDescribeRepositoryListResponse() (response *DescribeRepositoryListResponse) {
    response = &DescribeRepositoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询仓库列表。
func (c *Client) DescribeRepositoryList(request *DescribeRepositoryListRequest) (response *DescribeRepositoryListResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryListRequest()
    }
    response = NewDescribeRepositoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTeamRequest() (request *DescribeTeamRequest) {
    request = &DescribeTeamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeTeam")
    return
}

func NewDescribeTeamResponse() (response *DescribeTeamResponse) {
    response = &DescribeTeamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询团队信息。
// 
// 通过团队唯一标示或团队 ID 查询团队信息。
// 如果二者都填默认使用团队 ID 查询。
func (c *Client) DescribeTeam(request *DescribeTeamRequest) (response *DescribeTeamResponse, err error) {
    if request == nil {
        request = NewDescribeTeamRequest()
    }
    response = NewDescribeTeamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTestJobRequest() (request *DescribeTestJobRequest) {
    request = &DescribeTestJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeTestJob")
    return
}

func NewDescribeTestJobResponse() (response *DescribeTestJobResponse) {
    response = &DescribeTestJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 测试用 - 获取一个 Job 信息， 不可公开
func (c *Client) DescribeTestJob(request *DescribeTestJobRequest) (response *DescribeTestJobResponse, err error) {
    if request == nil {
        request = NewDescribeTestJobRequest()
    }
    response = NewDescribeTestJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVersionListRequest() (request *DescribeVersionListRequest) {
    request = &DescribeVersionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "DescribeVersionList")
    return
}

func NewDescribeVersionListResponse() (response *DescribeVersionListResponse) {
    response = &DescribeVersionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询制品版本列表。
func (c *Client) DescribeVersionList(request *DescribeVersionListRequest) (response *DescribeVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeVersionListRequest()
    }
    response = NewDescribeVersionListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyJobRequest() (request *ModifyJobRequest) {
    request = &ModifyJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "ModifyJob")
    return
}

func NewModifyJobResponse() (response *ModifyJobResponse) {
    response = &ModifyJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改任务。
func (c *Client) ModifyJob(request *ModifyJobRequest) (response *ModifyJobResponse, err error) {
    if request == nil {
        request = NewModifyJobRequest()
    }
    response = NewModifyJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopBuildRequest() (request *StopBuildRequest) {
    request = &StopBuildRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "StopBuild")
    return
}

func NewStopBuildResponse() (response *StopBuildResponse) {
    response = &StopBuildResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停止构建。
func (c *Client) StopBuild(request *StopBuildRequest) (response *StopBuildResponse, err error) {
    if request == nil {
        request = NewStopBuildRequest()
    }
    response = NewStopBuildResponse()
    err = c.Send(request, response)
    return
}

func NewTriggerBuildRequest() (request *TriggerBuildRequest) {
    request = &TriggerBuildRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("coding", APIVersion, "TriggerBuild")
    return
}

func NewTriggerBuildResponse() (response *TriggerBuildResponse) {
    response = &TriggerBuildResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 触发指定项目指定 CI 构建。
func (c *Client) TriggerBuild(request *TriggerBuildRequest) (response *TriggerBuildResponse, err error) {
    if request == nil {
        request = NewTriggerBuildRequest()
    }
    response = NewTriggerBuildResponse()
    err = c.Send(request, response)
    return
}
