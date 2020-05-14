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

package v20180416

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

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


func NewAddGroupForChaincodeRequest() (request *AddGroupForChaincodeRequest) {
    request = &AddGroupForChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "AddGroupForChaincode")
    return
}

func NewAddGroupForChaincodeResponse() (response *AddGroupForChaincodeResponse) {
    response = &AddGroupForChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智能合约新增组织
func (c *Client) AddGroupForChaincode(request *AddGroupForChaincodeRequest) (response *AddGroupForChaincodeResponse, err error) {
    if request == nil {
        request = NewAddGroupForChaincodeRequest()
    }
    response = NewAddGroupForChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewAddGroupForChannelRequest() (request *AddGroupForChannelRequest) {
    request = &AddGroupForChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "AddGroupForChannel")
    return
}

func NewAddGroupForChannelResponse() (response *AddGroupForChannelResponse) {
    response = &AddGroupForChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通道新增组织
func (c *Client) AddGroupForChannel(request *AddGroupForChannelRequest) (response *AddGroupForChannelResponse, err error) {
    if request == nil {
        request = NewAddGroupForChannelRequest()
    }
    response = NewAddGroupForChannelResponse()
    err = c.Send(request, response)
    return
}

func NewAllocateGrouptoMemberRequest() (request *AllocateGrouptoMemberRequest) {
    request = &AllocateGrouptoMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "AllocateGrouptoMember")
    return
}

func NewAllocateGrouptoMemberResponse() (response *AllocateGrouptoMemberResponse) {
    response = &AllocateGrouptoMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为联盟成员分配组织
func (c *Client) AllocateGrouptoMember(request *AllocateGrouptoMemberRequest) (response *AllocateGrouptoMemberResponse, err error) {
    if request == nil {
        request = NewAllocateGrouptoMemberRequest()
    }
    response = NewAllocateGrouptoMemberResponse()
    err = c.Send(request, response)
    return
}

func NewApplyCertRequest() (request *ApplyCertRequest) {
    request = &ApplyCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ApplyCert")
    return
}

func NewApplyCertResponse() (response *ApplyCertResponse) {
    response = &ApplyCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 申请证书
func (c *Client) ApplyCert(request *ApplyCertRequest) (response *ApplyCertResponse, err error) {
    if request == nil {
        request = NewApplyCertRequest()
    }
    response = NewApplyCertResponse()
    err = c.Send(request, response)
    return
}

func NewAsynCheckChaincodeDevRequest() (request *AsynCheckChaincodeDevRequest) {
    request = &AsynCheckChaincodeDevRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "AsynCheckChaincodeDev")
    return
}

func NewAsynCheckChaincodeDevResponse() (response *AsynCheckChaincodeDevResponse) {
    response = &AsynCheckChaincodeDevResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IDE智能合约异步操作结果查询
func (c *Client) AsynCheckChaincodeDev(request *AsynCheckChaincodeDevRequest) (response *AsynCheckChaincodeDevResponse, err error) {
    if request == nil {
        request = NewAsynCheckChaincodeDevRequest()
    }
    response = NewAsynCheckChaincodeDevResponse()
    err = c.Send(request, response)
    return
}

func NewAsynCompileChaincodeDevRequest() (request *AsynCompileChaincodeDevRequest) {
    request = &AsynCompileChaincodeDevRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "AsynCompileChaincodeDev")
    return
}

func NewAsynCompileChaincodeDevResponse() (response *AsynCompileChaincodeDevResponse) {
    response = &AsynCompileChaincodeDevResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IDE智能合约异步编译
func (c *Client) AsynCompileChaincodeDev(request *AsynCompileChaincodeDevRequest) (response *AsynCompileChaincodeDevResponse, err error) {
    if request == nil {
        request = NewAsynCompileChaincodeDevRequest()
    }
    response = NewAsynCompileChaincodeDevResponse()
    err = c.Send(request, response)
    return
}

func NewCheckChaincodeASTRequest() (request *CheckChaincodeASTRequest) {
    request = &CheckChaincodeASTRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CheckChaincodeAST")
    return
}

func NewCheckChaincodeASTResponse() (response *CheckChaincodeASTResponse) {
    response = &CheckChaincodeASTResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查合约的AST
func (c *Client) CheckChaincodeAST(request *CheckChaincodeASTRequest) (response *CheckChaincodeASTResponse, err error) {
    if request == nil {
        request = NewCheckChaincodeASTRequest()
    }
    response = NewCheckChaincodeASTResponse()
    err = c.Send(request, response)
    return
}

func NewCheckChaincodeChannelRequest() (request *CheckChaincodeChannelRequest) {
    request = &CheckChaincodeChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CheckChaincodeChannel")
    return
}

func NewCheckChaincodeChannelResponse() (response *CheckChaincodeChannelResponse) {
    response = &CheckChaincodeChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 校验合约在通道是否已初始化
func (c *Client) CheckChaincodeChannel(request *CheckChaincodeChannelRequest) (response *CheckChaincodeChannelResponse, err error) {
    if request == nil {
        request = NewCheckChaincodeChannelRequest()
    }
    response = NewCheckChaincodeChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCreateChaincodeRequest() (request *CheckCreateChaincodeRequest) {
    request = &CheckCreateChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CheckCreateChaincode")
    return
}

func NewCheckCreateChaincodeResponse() (response *CheckCreateChaincodeResponse) {
    response = &CheckCreateChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 合约创建检查
func (c *Client) CheckCreateChaincode(request *CheckCreateChaincodeRequest) (response *CheckCreateChaincodeResponse, err error) {
    if request == nil {
        request = NewCheckCreateChaincodeRequest()
    }
    response = NewCheckCreateChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewCheckEventConfigRequest() (request *CheckEventConfigRequest) {
    request = &CheckEventConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CheckEventConfig")
    return
}

func NewCheckEventConfigResponse() (response *CheckEventConfigResponse) {
    response = &CheckEventConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 事件配置校验
func (c *Client) CheckEventConfig(request *CheckEventConfigRequest) (response *CheckEventConfigResponse, err error) {
    if request == nil {
        request = NewCheckEventConfigRequest()
    }
    response = NewCheckEventConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCheckGroupClusterCreatorRequest() (request *CheckGroupClusterCreatorRequest) {
    request = &CheckGroupClusterCreatorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CheckGroupClusterCreator")
    return
}

func NewCheckGroupClusterCreatorResponse() (response *CheckGroupClusterCreatorResponse) {
    response = &CheckGroupClusterCreatorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询组织是否是链的创建者
func (c *Client) CheckGroupClusterCreator(request *CheckGroupClusterCreatorRequest) (response *CheckGroupClusterCreatorResponse, err error) {
    if request == nil {
        request = NewCheckGroupClusterCreatorRequest()
    }
    response = NewCheckGroupClusterCreatorResponse()
    err = c.Send(request, response)
    return
}

func NewCompileChaincodeRequest() (request *CompileChaincodeRequest) {
    request = &CompileChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CompileChaincode")
    return
}

func NewCompileChaincodeResponse() (response *CompileChaincodeResponse) {
    response = &CompileChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 验证智能合约是否合法
func (c *Client) CompileChaincode(request *CompileChaincodeRequest) (response *CompileChaincodeResponse, err error) {
    if request == nil {
        request = NewCompileChaincodeRequest()
    }
    response = NewCompileChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChaincodeRequest() (request *CreateChaincodeRequest) {
    request = &CreateChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CreateChaincode")
    return
}

func NewCreateChaincodeResponse() (response *CreateChaincodeResponse) {
    response = &CreateChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建合约
func (c *Client) CreateChaincode(request *CreateChaincodeRequest) (response *CreateChaincodeResponse, err error) {
    if request == nil {
        request = NewCreateChaincodeRequest()
    }
    response = NewCreateChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChannelRequest() (request *CreateChannelRequest) {
    request = &CreateChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CreateChannel")
    return
}

func NewCreateChannelResponse() (response *CreateChannelResponse) {
    response = &CreateChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建通道
func (c *Client) CreateChannel(request *CreateChannelRequest) (response *CreateChannelResponse, err error) {
    if request == nil {
        request = NewCreateChannelRequest()
    }
    response = NewCreateChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChannelForCBCRequest() (request *CreateChannelForCBCRequest) {
    request = &CreateChannelForCBCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CreateChannelForCBC")
    return
}

func NewCreateChannelForCBCResponse() (response *CreateChannelForCBCResponse) {
    response = &CreateChannelForCBCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建联盟链通道
func (c *Client) CreateChannelForCBC(request *CreateChannelForCBCRequest) (response *CreateChannelForCBCResponse, err error) {
    if request == nil {
        request = NewCreateChannelForCBCRequest()
    }
    response = NewCreateChannelForCBCResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsortiumRequest() (request *CreateConsortiumRequest) {
    request = &CreateConsortiumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CreateConsortium")
    return
}

func NewCreateConsortiumResponse() (response *CreateConsortiumResponse) {
    response = &CreateConsortiumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建联盟
func (c *Client) CreateConsortium(request *CreateConsortiumRequest) (response *CreateConsortiumResponse, err error) {
    if request == nil {
        request = NewCreateConsortiumRequest()
    }
    response = NewCreateConsortiumResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEndorsementRequest() (request *CreateEndorsementRequest) {
    request = &CreateEndorsementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CreateEndorsement")
    return
}

func NewCreateEndorsementResponse() (response *CreateEndorsementResponse) {
    response = &CreateEndorsementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建背书策略
func (c *Client) CreateEndorsement(request *CreateEndorsementRequest) (response *CreateEndorsementResponse, err error) {
    if request == nil {
        request = NewCreateEndorsementRequest()
    }
    response = NewCreateEndorsementResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFabricClusterRequest() (request *CreateFabricClusterRequest) {
    request = &CreateFabricClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CreateFabricCluster")
    return
}

func NewCreateFabricClusterResponse() (response *CreateFabricClusterResponse) {
    response = &CreateFabricClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建网络资源（后付费）
func (c *Client) CreateFabricCluster(request *CreateFabricClusterRequest) (response *CreateFabricClusterResponse, err error) {
    if request == nil {
        request = NewCreateFabricClusterRequest()
    }
    response = NewCreateFabricClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateLinkRequest() (request *CreatePrivateLinkRequest) {
    request = &CreatePrivateLinkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CreatePrivateLink")
    return
}

func NewCreatePrivateLinkResponse() (response *CreatePrivateLinkResponse) {
    response = &CreatePrivateLinkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建私有连接
func (c *Client) CreatePrivateLink(request *CreatePrivateLinkRequest) (response *CreatePrivateLinkResponse, err error) {
    if request == nil {
        request = NewCreatePrivateLinkRequest()
    }
    response = NewCreatePrivateLinkResponse()
    err = c.Send(request, response)
    return
}

func NewDealEventTaskRequest() (request *DealEventTaskRequest) {
    request = &DealEventTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DealEventTask")
    return
}

func NewDealEventTaskResponse() (response *DealEventTaskResponse) {
    response = &DealEventTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 处理代办任务
func (c *Client) DealEventTask(request *DealEventTaskRequest) (response *DealEventTaskResponse, err error) {
    if request == nil {
        request = NewDealEventTaskRequest()
    }
    response = NewDealEventTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateLinkRequest() (request *DeletePrivateLinkRequest) {
    request = &DeletePrivateLinkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DeletePrivateLink")
    return
}

func NewDeletePrivateLinkResponse() (response *DeletePrivateLinkResponse) {
    response = &DeletePrivateLinkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除私有连接
func (c *Client) DeletePrivateLink(request *DeletePrivateLinkRequest) (response *DeletePrivateLinkResponse, err error) {
    if request == nil {
        request = NewDeletePrivateLinkRequest()
    }
    response = NewDeletePrivateLinkResponse()
    err = c.Send(request, response)
    return
}

func NewDestoryFabricClusterRequest() (request *DestoryFabricClusterRequest) {
    request = &DestoryFabricClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DestoryFabricCluster")
    return
}

func NewDestoryFabricClusterResponse() (response *DestoryFabricClusterResponse) {
    response = &DestoryFabricClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 销毁网络资源（后付费）
func (c *Client) DestoryFabricCluster(request *DestoryFabricClusterRequest) (response *DestoryFabricClusterResponse, err error) {
    if request == nil {
        request = NewDestoryFabricClusterRequest()
    }
    response = NewDestoryFabricClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadCertRequest() (request *DownloadCertRequest) {
    request = &DownloadCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DownloadCert")
    return
}

func NewDownloadCertResponse() (response *DownloadCertResponse) {
    response = &DownloadCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 证书下载
func (c *Client) DownloadCert(request *DownloadCertRequest) (response *DownloadCertResponse, err error) {
    if request == nil {
        request = NewDownloadCertRequest()
    }
    response = NewDownloadCertResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadEventConfigRequest() (request *DownloadEventConfigRequest) {
    request = &DownloadEventConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DownloadEventConfig")
    return
}

func NewDownloadEventConfigResponse() (response *DownloadEventConfigResponse) {
    response = &DownloadEventConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载事件配置
func (c *Client) DownloadEventConfig(request *DownloadEventConfigRequest) (response *DownloadEventConfigResponse, err error) {
    if request == nil {
        request = NewDownloadEventConfigRequest()
    }
    response = NewDownloadEventConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadTbaasIdentityRequest() (request *DownloadTbaasIdentityRequest) {
    request = &DownloadTbaasIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DownloadTbaasIdentity")
    return
}

func NewDownloadTbaasIdentityResponse() (response *DownloadTbaasIdentityResponse) {
    response = &DownloadTbaasIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载用户TBaaSID身份信息
func (c *Client) DownloadTbaasIdentity(request *DownloadTbaasIdentityRequest) (response *DownloadTbaasIdentityResponse, err error) {
    if request == nil {
        request = NewDownloadTbaasIdentityRequest()
    }
    response = NewDownloadTbaasIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewExportChaincodeDevRequest() (request *ExportChaincodeDevRequest) {
    request = &ExportChaincodeDevRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ExportChaincodeDev")
    return
}

func NewExportChaincodeDevResponse() (response *ExportChaincodeDevResponse) {
    response = &ExportChaincodeDevResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IDE智能合约导出
func (c *Client) ExportChaincodeDev(request *ExportChaincodeDevRequest) (response *ExportChaincodeDevResponse, err error) {
    if request == nil {
        request = NewExportChaincodeDevRequest()
    }
    response = NewExportChaincodeDevResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockDetailRequest() (request *GetBlockDetailRequest) {
    request = &GetBlockDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockDetail")
    return
}

func NewGetBlockDetailResponse() (response *GetBlockDetailResponse) {
    response = &GetBlockDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看某个固定区块的详细信息
func (c *Client) GetBlockDetail(request *GetBlockDetailRequest) (response *GetBlockDetailResponse, err error) {
    if request == nil {
        request = NewGetBlockDetailRequest()
    }
    response = NewGetBlockDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockListRequest() (request *GetBlockListRequest) {
    request = &GetBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockList")
    return
}

func NewGetBlockListResponse() (response *GetBlockListResponse) {
    response = &GetBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看当前网络下的所有区块列表，分页展示
func (c *Client) GetBlockList(request *GetBlockListRequest) (response *GetBlockListResponse, err error) {
    if request == nil {
        request = NewGetBlockListRequest()
    }
    response = NewGetBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockSpeedRequest() (request *GetBlockSpeedRequest) {
    request = &GetBlockSpeedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockSpeed")
    return
}

func NewGetBlockSpeedResponse() (response *GetBlockSpeedResponse) {
    response = &GetBlockSpeedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询通道新增区块数趋势
func (c *Client) GetBlockSpeed(request *GetBlockSpeedRequest) (response *GetBlockSpeedResponse, err error) {
    if request == nil {
        request = NewGetBlockSpeedRequest()
    }
    response = NewGetBlockSpeedResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockTransactionListRequest() (request *GetBlockTransactionListRequest) {
    request = &GetBlockTransactionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockTransactionList")
    return
}

func NewGetBlockTransactionListResponse() (response *GetBlockTransactionListResponse) {
    response = &GetBlockTransactionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取区块交易列表
func (c *Client) GetBlockTransactionList(request *GetBlockTransactionListRequest) (response *GetBlockTransactionListResponse, err error) {
    if request == nil {
        request = NewGetBlockTransactionListRequest()
    }
    response = NewGetBlockTransactionListResponse()
    err = c.Send(request, response)
    return
}

func NewGetCdhStatusRequest() (request *GetCdhStatusRequest) {
    request = &GetCdhStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetCdhStatus")
    return
}

func NewGetCdhStatusResponse() (response *GetCdhStatusResponse) {
    response = &GetCdhStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询虚拟机的售卖状态
func (c *Client) GetCdhStatus(request *GetCdhStatusRequest) (response *GetCdhStatusResponse, err error) {
    if request == nil {
        request = NewGetCdhStatusRequest()
    }
    response = NewGetCdhStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetCertDetailRequest() (request *GetCertDetailRequest) {
    request = &GetCertDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetCertDetail")
    return
}

func NewGetCertDetailResponse() (response *GetCertDetailResponse) {
    response = &GetCertDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询证书详情
func (c *Client) GetCertDetail(request *GetCertDetailRequest) (response *GetCertDetailResponse, err error) {
    if request == nil {
        request = NewGetCertDetailRequest()
    }
    response = NewGetCertDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetCertListRequest() (request *GetCertListRequest) {
    request = &GetCertListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetCertList")
    return
}

func NewGetCertListResponse() (response *GetCertListResponse) {
    response = &GetCertListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询证书列表接口，用于筛选和查询证书
func (c *Client) GetCertList(request *GetCertListRequest) (response *GetCertListResponse, err error) {
    if request == nil {
        request = NewGetCertListRequest()
    }
    response = NewGetCertListResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeDetailRequest() (request *GetChaincodeDetailRequest) {
    request = &GetChaincodeDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeDetail")
    return
}

func NewGetChaincodeDetailResponse() (response *GetChaincodeDetailResponse) {
    response = &GetChaincodeDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询合约详情
func (c *Client) GetChaincodeDetail(request *GetChaincodeDetailRequest) (response *GetChaincodeDetailResponse, err error) {
    if request == nil {
        request = NewGetChaincodeDetailRequest()
    }
    response = NewGetChaincodeDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeDevAccessAuthRequest() (request *GetChaincodeDevAccessAuthRequest) {
    request = &GetChaincodeDevAccessAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeDevAccessAuth")
    return
}

func NewGetChaincodeDevAccessAuthResponse() (response *GetChaincodeDevAccessAuthResponse) {
    response = &GetChaincodeDevAccessAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在线编辑器访问权限查询
func (c *Client) GetChaincodeDevAccessAuth(request *GetChaincodeDevAccessAuthRequest) (response *GetChaincodeDevAccessAuthResponse, err error) {
    if request == nil {
        request = NewGetChaincodeDevAccessAuthRequest()
    }
    response = NewGetChaincodeDevAccessAuthResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeListRequest() (request *GetChaincodeListRequest) {
    request = &GetChaincodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeList")
    return
}

func NewGetChaincodeListResponse() (response *GetChaincodeListResponse) {
    response = &GetChaincodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询合约列表
func (c *Client) GetChaincodeList(request *GetChaincodeListRequest) (response *GetChaincodeListResponse, err error) {
    if request == nil {
        request = NewGetChaincodeListRequest()
    }
    response = NewGetChaincodeListResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeListPerChannelRequest() (request *GetChaincodeListPerChannelRequest) {
    request = &GetChaincodeListPerChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeListPerChannel")
    return
}

func NewGetChaincodeListPerChannelResponse() (response *GetChaincodeListPerChannelResponse) {
    response = &GetChaincodeListPerChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看通道的合约列表
func (c *Client) GetChaincodeListPerChannel(request *GetChaincodeListPerChannelRequest) (response *GetChaincodeListPerChannelResponse, err error) {
    if request == nil {
        request = NewGetChaincodeListPerChannelRequest()
    }
    response = NewGetChaincodeListPerChannelResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeListPerEndorsementRequest() (request *GetChaincodeListPerEndorsementRequest) {
    request = &GetChaincodeListPerEndorsementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeListPerEndorsement")
    return
}

func NewGetChaincodeListPerEndorsementResponse() (response *GetChaincodeListPerEndorsementResponse) {
    response = &GetChaincodeListPerEndorsementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询背书策略关联的合约列表
func (c *Client) GetChaincodeListPerEndorsement(request *GetChaincodeListPerEndorsementRequest) (response *GetChaincodeListPerEndorsementResponse, err error) {
    if request == nil {
        request = NewGetChaincodeListPerEndorsementRequest()
    }
    response = NewGetChaincodeListPerEndorsementResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodePvtDataListRequest() (request *GetChaincodePvtDataListRequest) {
    request = &GetChaincodePvtDataListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodePvtDataList")
    return
}

func NewGetChaincodePvtDataListResponse() (response *GetChaincodePvtDataListResponse) {
    response = &GetChaincodePvtDataListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取合约私有数据列表
func (c *Client) GetChaincodePvtDataList(request *GetChaincodePvtDataListRequest) (response *GetChaincodePvtDataListResponse, err error) {
    if request == nil {
        request = NewGetChaincodePvtDataListRequest()
    }
    response = NewGetChaincodePvtDataListResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeTemplateRequest() (request *GetChaincodeTemplateRequest) {
    request = &GetChaincodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeTemplate")
    return
}

func NewGetChaincodeTemplateResponse() (response *GetChaincodeTemplateResponse) {
    response = &GetChaincodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Chaincode模版
func (c *Client) GetChaincodeTemplate(request *GetChaincodeTemplateRequest) (response *GetChaincodeTemplateResponse, err error) {
    if request == nil {
        request = NewGetChaincodeTemplateRequest()
    }
    response = NewGetChaincodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelDetailRequest() (request *GetChannelDetailRequest) {
    request = &GetChannelDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelDetail")
    return
}

func NewGetChannelDetailResponse() (response *GetChannelDetailResponse) {
    response = &GetChannelDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询通道详情
func (c *Client) GetChannelDetail(request *GetChannelDetailRequest) (response *GetChannelDetailResponse, err error) {
    if request == nil {
        request = NewGetChannelDetailRequest()
    }
    response = NewGetChannelDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelDetailForEventRequest() (request *GetChannelDetailForEventRequest) {
    request = &GetChannelDetailForEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelDetailForEvent")
    return
}

func NewGetChannelDetailForEventResponse() (response *GetChannelDetailForEventResponse) {
    response = &GetChannelDetailForEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件中心通道概要
func (c *Client) GetChannelDetailForEvent(request *GetChannelDetailForEventRequest) (response *GetChannelDetailForEventResponse, err error) {
    if request == nil {
        request = NewGetChannelDetailForEventRequest()
    }
    response = NewGetChannelDetailForEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelInviteesForEventRequest() (request *GetChannelInviteesForEventRequest) {
    request = &GetChannelInviteesForEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelInviteesForEvent")
    return
}

func NewGetChannelInviteesForEventResponse() (response *GetChannelInviteesForEventResponse) {
    response = &GetChannelInviteesForEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取通道被邀请成员信息
func (c *Client) GetChannelInviteesForEvent(request *GetChannelInviteesForEventRequest) (response *GetChannelInviteesForEventResponse, err error) {
    if request == nil {
        request = NewGetChannelInviteesForEventRequest()
    }
    response = NewGetChannelInviteesForEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelListRequest() (request *GetChannelListRequest) {
    request = &GetChannelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelList")
    return
}

func NewGetChannelListResponse() (response *GetChannelListResponse) {
    response = &GetChannelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询通道列表
func (c *Client) GetChannelList(request *GetChannelListRequest) (response *GetChannelListResponse, err error) {
    if request == nil {
        request = NewGetChannelListRequest()
    }
    response = NewGetChannelListResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelListForCloudMonitorRequest() (request *GetChannelListForCloudMonitorRequest) {
    request = &GetChannelListForCloudMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelListForCloudMonitor")
    return
}

func NewGetChannelListForCloudMonitorResponse() (response *GetChannelListForCloudMonitorResponse) {
    response = &GetChannelListForCloudMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询组织下的通道列表（云监控实例）
func (c *Client) GetChannelListForCloudMonitor(request *GetChannelListForCloudMonitorRequest) (response *GetChannelListForCloudMonitorResponse, err error) {
    if request == nil {
        request = NewGetChannelListForCloudMonitorRequest()
    }
    response = NewGetChannelListForCloudMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelListForInitRequest() (request *GetChannelListForInitRequest) {
    request = &GetChannelListForInitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelListForInit")
    return
}

func NewGetChannelListForInitResponse() (response *GetChannelListForInitResponse) {
    response = &GetChannelListForInitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下拉通道列表（合约初始化）
func (c *Client) GetChannelListForInit(request *GetChannelListForInitRequest) (response *GetChannelListForInitResponse, err error) {
    if request == nil {
        request = NewGetChannelListForInitRequest()
    }
    response = NewGetChannelListForInitResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelListPerChaincodeRequest() (request *GetChannelListPerChaincodeRequest) {
    request = &GetChannelListPerChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelListPerChaincode")
    return
}

func NewGetChannelListPerChaincodeResponse() (response *GetChannelListPerChaincodeResponse) {
    response = &GetChannelListPerChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询合约关联的通道列表
func (c *Client) GetChannelListPerChaincode(request *GetChannelListPerChaincodeRequest) (response *GetChannelListPerChaincodeResponse, err error) {
    if request == nil {
        request = NewGetChannelListPerChaincodeRequest()
    }
    response = NewGetChannelListPerChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelVotersForEventRequest() (request *GetChannelVotersForEventRequest) {
    request = &GetChannelVotersForEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelVotersForEvent")
    return
}

func NewGetChannelVotersForEventResponse() (response *GetChannelVotersForEventResponse) {
    response = &GetChannelVotersForEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取通道投票参与者信息
func (c *Client) GetChannelVotersForEvent(request *GetChannelVotersForEventRequest) (response *GetChannelVotersForEventResponse, err error) {
    if request == nil {
        request = NewGetChannelVotersForEventRequest()
    }
    response = NewGetChannelVotersForEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusteDetailForEventRequest() (request *GetClusteDetailForEventRequest) {
    request = &GetClusteDetailForEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusteDetailForEvent")
    return
}

func NewGetClusteDetailForEventResponse() (response *GetClusteDetailForEventResponse) {
    response = &GetClusteDetailForEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件中心网络详情
func (c *Client) GetClusteDetailForEvent(request *GetClusteDetailForEventRequest) (response *GetClusteDetailForEventResponse, err error) {
    if request == nil {
        request = NewGetClusteDetailForEventRequest()
    }
    response = NewGetClusteDetailForEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusteMemberForEventRequest() (request *GetClusteMemberForEventRequest) {
    request = &GetClusteMemberForEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusteMemberForEvent")
    return
}

func NewGetClusteMemberForEventResponse() (response *GetClusteMemberForEventResponse) {
    response = &GetClusteMemberForEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件中心网络成员
func (c *Client) GetClusteMemberForEvent(request *GetClusteMemberForEventRequest) (response *GetClusteMemberForEventResponse, err error) {
    if request == nil {
        request = NewGetClusteMemberForEventRequest()
    }
    response = NewGetClusteMemberForEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterDetailRequest() (request *GetClusterDetailRequest) {
    request = &GetClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterDetail")
    return
}

func NewGetClusterDetailResponse() (response *GetClusterDetailResponse) {
    response = &GetClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取区块链网络详情
func (c *Client) GetClusterDetail(request *GetClusterDetailRequest) (response *GetClusterDetailResponse, err error) {
    if request == nil {
        request = NewGetClusterDetailRequest()
    }
    response = NewGetClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterListRequest() (request *GetClusterListRequest) {
    request = &GetClusterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterList")
    return
}

func NewGetClusterListResponse() (response *GetClusterListResponse) {
    response = &GetClusterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取区块链网络
func (c *Client) GetClusterList(request *GetClusterListRequest) (response *GetClusterListResponse, err error) {
    if request == nil {
        request = NewGetClusterListRequest()
    }
    response = NewGetClusterListResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterListForCloudMonitorRequest() (request *GetClusterListForCloudMonitorRequest) {
    request = &GetClusterListForCloudMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterListForCloudMonitor")
    return
}

func NewGetClusterListForCloudMonitorResponse() (response *GetClusterListForCloudMonitorResponse) {
    response = &GetClusterListForCloudMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络列表（云监控实例）
func (c *Client) GetClusterListForCloudMonitor(request *GetClusterListForCloudMonitorRequest) (response *GetClusterListForCloudMonitorResponse, err error) {
    if request == nil {
        request = NewGetClusterListForCloudMonitorRequest()
    }
    response = NewGetClusterListForCloudMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterNoJoinMembersRequest() (request *GetClusterNoJoinMembersRequest) {
    request = &GetClusterNoJoinMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterNoJoinMembers")
    return
}

func NewGetClusterNoJoinMembersResponse() (response *GetClusterNoJoinMembersResponse) {
    response = &GetClusterNoJoinMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取未加入网络的联盟成员信息
func (c *Client) GetClusterNoJoinMembers(request *GetClusterNoJoinMembersRequest) (response *GetClusterNoJoinMembersResponse, err error) {
    if request == nil {
        request = NewGetClusterNoJoinMembersRequest()
    }
    response = NewGetClusterNoJoinMembersResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterNumberStatisticsRequest() (request *GetClusterNumberStatisticsRequest) {
    request = &GetClusterNumberStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterNumberStatistics")
    return
}

func NewGetClusterNumberStatisticsResponse() (response *GetClusterNumberStatisticsResponse) {
    response = &GetClusterNumberStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取各个区域的网络数量
func (c *Client) GetClusterNumberStatistics(request *GetClusterNumberStatisticsRequest) (response *GetClusterNumberStatisticsResponse, err error) {
    if request == nil {
        request = NewGetClusterNumberStatisticsRequest()
    }
    response = NewGetClusterNumberStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterResourceInfoRequest() (request *GetClusterResourceInfoRequest) {
    request = &GetClusterResourceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterResourceInfo")
    return
}

func NewGetClusterResourceInfoResponse() (response *GetClusterResourceInfoResponse) {
    response = &GetClusterResourceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取网络对应资源的详情
func (c *Client) GetClusterResourceInfo(request *GetClusterResourceInfoRequest) (response *GetClusterResourceInfoResponse, err error) {
    if request == nil {
        request = NewGetClusterResourceInfoRequest()
    }
    response = NewGetClusterResourceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterSummaryRequest() (request *GetClusterSummaryRequest) {
    request = &GetClusterSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterSummary")
    return
}

func NewGetClusterSummaryResponse() (response *GetClusterSummaryResponse) {
    response = &GetClusterSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取区块链网络概要
func (c *Client) GetClusterSummary(request *GetClusterSummaryRequest) (response *GetClusterSummaryResponse, err error) {
    if request == nil {
        request = NewGetClusterSummaryRequest()
    }
    response = NewGetClusterSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterTypeRequest() (request *GetClusterTypeRequest) {
    request = &GetClusterTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterType")
    return
}

func NewGetClusterTypeResponse() (response *GetClusterTypeResponse) {
    response = &GetClusterTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取网络类型
func (c *Client) GetClusterType(request *GetClusterTypeRequest) (response *GetClusterTypeResponse, err error) {
    if request == nil {
        request = NewGetClusterTypeRequest()
    }
    response = NewGetClusterTypeResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumClusterListSummaryRequest() (request *GetConsortiumClusterListSummaryRequest) {
    request = &GetConsortiumClusterListSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetConsortiumClusterListSummary")
    return
}

func NewGetConsortiumClusterListSummaryResponse() (response *GetConsortiumClusterListSummaryResponse) {
    response = &GetConsortiumClusterListSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取联盟网络列表概览
func (c *Client) GetConsortiumClusterListSummary(request *GetConsortiumClusterListSummaryRequest) (response *GetConsortiumClusterListSummaryResponse, err error) {
    if request == nil {
        request = NewGetConsortiumClusterListSummaryRequest()
    }
    response = NewGetConsortiumClusterListSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumDetailRequest() (request *GetConsortiumDetailRequest) {
    request = &GetConsortiumDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetConsortiumDetail")
    return
}

func NewGetConsortiumDetailResponse() (response *GetConsortiumDetailResponse) {
    response = &GetConsortiumDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取联盟详情
func (c *Client) GetConsortiumDetail(request *GetConsortiumDetailRequest) (response *GetConsortiumDetailResponse, err error) {
    if request == nil {
        request = NewGetConsortiumDetailRequest()
    }
    response = NewGetConsortiumDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumDtailForEventRequest() (request *GetConsortiumDtailForEventRequest) {
    request = &GetConsortiumDtailForEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetConsortiumDtailForEvent")
    return
}

func NewGetConsortiumDtailForEventResponse() (response *GetConsortiumDtailForEventResponse) {
    response = &GetConsortiumDtailForEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件中心联盟详情
func (c *Client) GetConsortiumDtailForEvent(request *GetConsortiumDtailForEventRequest) (response *GetConsortiumDtailForEventResponse, err error) {
    if request == nil {
        request = NewGetConsortiumDtailForEventRequest()
    }
    response = NewGetConsortiumDtailForEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumListRequest() (request *GetConsortiumListRequest) {
    request = &GetConsortiumListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetConsortiumList")
    return
}

func NewGetConsortiumListResponse() (response *GetConsortiumListResponse) {
    response = &GetConsortiumListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取联盟列表
func (c *Client) GetConsortiumList(request *GetConsortiumListRequest) (response *GetConsortiumListResponse, err error) {
    if request == nil {
        request = NewGetConsortiumListRequest()
    }
    response = NewGetConsortiumListResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumMemberAuthDataRequest() (request *GetConsortiumMemberAuthDataRequest) {
    request = &GetConsortiumMemberAuthDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetConsortiumMemberAuthData")
    return
}

func NewGetConsortiumMemberAuthDataResponse() (response *GetConsortiumMemberAuthDataResponse) {
    response = &GetConsortiumMemberAuthDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取成员认证信息
func (c *Client) GetConsortiumMemberAuthData(request *GetConsortiumMemberAuthDataRequest) (response *GetConsortiumMemberAuthDataResponse, err error) {
    if request == nil {
        request = NewGetConsortiumMemberAuthDataRequest()
    }
    response = NewGetConsortiumMemberAuthDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumMemberForEventRequest() (request *GetConsortiumMemberForEventRequest) {
    request = &GetConsortiumMemberForEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetConsortiumMemberForEvent")
    return
}

func NewGetConsortiumMemberForEventResponse() (response *GetConsortiumMemberForEventResponse) {
    response = &GetConsortiumMemberForEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件中心联盟成员
func (c *Client) GetConsortiumMemberForEvent(request *GetConsortiumMemberForEventRequest) (response *GetConsortiumMemberForEventResponse, err error) {
    if request == nil {
        request = NewGetConsortiumMemberForEventRequest()
    }
    response = NewGetConsortiumMemberForEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumMembersRequest() (request *GetConsortiumMembersRequest) {
    request = &GetConsortiumMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetConsortiumMembers")
    return
}

func NewGetConsortiumMembersResponse() (response *GetConsortiumMembersResponse) {
    response = &GetConsortiumMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取联盟成员列表
func (c *Client) GetConsortiumMembers(request *GetConsortiumMembersRequest) (response *GetConsortiumMembersResponse, err error) {
    if request == nil {
        request = NewGetConsortiumMembersRequest()
    }
    response = NewGetConsortiumMembersResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumSummaryRequest() (request *GetConsortiumSummaryRequest) {
    request = &GetConsortiumSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetConsortiumSummary")
    return
}

func NewGetConsortiumSummaryResponse() (response *GetConsortiumSummaryResponse) {
    response = &GetConsortiumSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取联盟概况
func (c *Client) GetConsortiumSummary(request *GetConsortiumSummaryRequest) (response *GetConsortiumSummaryResponse, err error) {
    if request == nil {
        request = NewGetConsortiumSummaryRequest()
    }
    response = NewGetConsortiumSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetCosSignRequest() (request *GetCosSignRequest) {
    request = &GetCosSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetCosSign")
    return
}

func NewGetCosSignResponse() (response *GetCosSignResponse) {
    response = &GetCosSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Cos的Sign
func (c *Client) GetCosSign(request *GetCosSignRequest) (response *GetCosSignResponse, err error) {
    if request == nil {
        request = NewGetCosSignRequest()
    }
    response = NewGetCosSignResponse()
    err = c.Send(request, response)
    return
}

func NewGetCosUrlRequest() (request *GetCosUrlRequest) {
    request = &GetCosUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetCosUrl")
    return
}

func NewGetCosUrlResponse() (response *GetCosUrlResponse) {
    response = &GetCosUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Cos的下载链接
func (c *Client) GetCosUrl(request *GetCosUrlRequest) (response *GetCosUrlResponse, err error) {
    if request == nil {
        request = NewGetCosUrlRequest()
    }
    response = NewGetCosUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetEndorsementDetailRequest() (request *GetEndorsementDetailRequest) {
    request = &GetEndorsementDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetEndorsementDetail")
    return
}

func NewGetEndorsementDetailResponse() (response *GetEndorsementDetailResponse) {
    response = &GetEndorsementDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看背书策略详情
func (c *Client) GetEndorsementDetail(request *GetEndorsementDetailRequest) (response *GetEndorsementDetailResponse, err error) {
    if request == nil {
        request = NewGetEndorsementDetailRequest()
    }
    response = NewGetEndorsementDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetEndorsementListRequest() (request *GetEndorsementListRequest) {
    request = &GetEndorsementListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetEndorsementList")
    return
}

func NewGetEndorsementListResponse() (response *GetEndorsementListResponse) {
    response = &GetEndorsementListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看背书策略列表
func (c *Client) GetEndorsementList(request *GetEndorsementListRequest) (response *GetEndorsementListResponse, err error) {
    if request == nil {
        request = NewGetEndorsementListRequest()
    }
    response = NewGetEndorsementListResponse()
    err = c.Send(request, response)
    return
}

func NewGetEventDetailRequest() (request *GetEventDetailRequest) {
    request = &GetEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetEventDetail")
    return
}

func NewGetEventDetailResponse() (response *GetEventDetailResponse) {
    response = &GetEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件详情
func (c *Client) GetEventDetail(request *GetEventDetailRequest) (response *GetEventDetailResponse, err error) {
    if request == nil {
        request = NewGetEventDetailRequest()
    }
    response = NewGetEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetEventListRequest() (request *GetEventListRequest) {
    request = &GetEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetEventList")
    return
}

func NewGetEventListResponse() (response *GetEventListResponse) {
    response = &GetEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件列表
func (c *Client) GetEventList(request *GetEventListRequest) (response *GetEventListResponse, err error) {
    if request == nil {
        request = NewGetEventListRequest()
    }
    response = NewGetEventListResponse()
    err = c.Send(request, response)
    return
}

func NewGetEventStepStatusRequest() (request *GetEventStepStatusRequest) {
    request = &GetEventStepStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetEventStepStatus")
    return
}

func NewGetEventStepStatusResponse() (response *GetEventStepStatusResponse) {
    response = &GetEventStepStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件流程状态
func (c *Client) GetEventStepStatus(request *GetEventStepStatusRequest) (response *GetEventStepStatusResponse, err error) {
    if request == nil {
        request = NewGetEventStepStatusRequest()
    }
    response = NewGetEventStepStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetEventSummaryRequest() (request *GetEventSummaryRequest) {
    request = &GetEventSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetEventSummary")
    return
}

func NewGetEventSummaryResponse() (response *GetEventSummaryResponse) {
    response = &GetEventSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取事件中心概览
func (c *Client) GetEventSummary(request *GetEventSummaryRequest) (response *GetEventSummaryResponse, err error) {
    if request == nil {
        request = NewGetEventSummaryRequest()
    }
    response = NewGetEventSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListForCloudMonitorRequest() (request *GetGroupListForCloudMonitorRequest) {
    request = &GetGroupListForCloudMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetGroupListForCloudMonitor")
    return
}

func NewGetGroupListForCloudMonitorResponse() (response *GetGroupListForCloudMonitorResponse) {
    response = &GetGroupListForCloudMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户在区块链网络中的组织列表（云监控实例）
func (c *Client) GetGroupListForCloudMonitor(request *GetGroupListForCloudMonitorRequest) (response *GetGroupListForCloudMonitorResponse, err error) {
    if request == nil {
        request = NewGetGroupListForCloudMonitorRequest()
    }
    response = NewGetGroupListForCloudMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListNoChaincodeRequest() (request *GetGroupListNoChaincodeRequest) {
    request = &GetGroupListNoChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetGroupListNoChaincode")
    return
}

func NewGetGroupListNoChaincodeResponse() (response *GetGroupListNoChaincodeResponse) {
    response = &GetGroupListNoChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取还没有加入Chaincode的组织列表
func (c *Client) GetGroupListNoChaincode(request *GetGroupListNoChaincodeRequest) (response *GetGroupListNoChaincodeResponse, err error) {
    if request == nil {
        request = NewGetGroupListNoChaincodeRequest()
    }
    response = NewGetGroupListNoChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListPerAppidRequest() (request *GetGroupListPerAppidRequest) {
    request = &GetGroupListPerAppidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetGroupListPerAppid")
    return
}

func NewGetGroupListPerAppidResponse() (response *GetGroupListPerAppidResponse) {
    response = &GetGroupListPerAppidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询一个用户的组织列表
func (c *Client) GetGroupListPerAppid(request *GetGroupListPerAppidRequest) (response *GetGroupListPerAppidResponse, err error) {
    if request == nil {
        request = NewGetGroupListPerAppidRequest()
    }
    response = NewGetGroupListPerAppidResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListPerChaincodeRequest() (request *GetGroupListPerChaincodeRequest) {
    request = &GetGroupListPerChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetGroupListPerChaincode")
    return
}

func NewGetGroupListPerChaincodeResponse() (response *GetGroupListPerChaincodeResponse) {
    response = &GetGroupListPerChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询合约允许使用的组织列表
func (c *Client) GetGroupListPerChaincode(request *GetGroupListPerChaincodeRequest) (response *GetGroupListPerChaincodeResponse, err error) {
    if request == nil {
        request = NewGetGroupListPerChaincodeRequest()
    }
    response = NewGetGroupListPerChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListPerChannelRequest() (request *GetGroupListPerChannelRequest) {
    request = &GetGroupListPerChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetGroupListPerChannel")
    return
}

func NewGetGroupListPerChannelResponse() (response *GetGroupListPerChannelResponse) {
    response = &GetGroupListPerChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Channel允许使用的组织列表
func (c *Client) GetGroupListPerChannel(request *GetGroupListPerChannelRequest) (response *GetGroupListPerChannelResponse, err error) {
    if request == nil {
        request = NewGetGroupListPerChannelRequest()
    }
    response = NewGetGroupListPerChannelResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListPerClusterRequest() (request *GetGroupListPerClusterRequest) {
    request = &GetGroupListPerClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetGroupListPerCluster")
    return
}

func NewGetGroupListPerClusterResponse() (response *GetGroupListPerClusterResponse) {
    response = &GetGroupListPerClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询一个Cluster的组织列表
func (c *Client) GetGroupListPerCluster(request *GetGroupListPerClusterRequest) (response *GetGroupListPerClusterResponse, err error) {
    if request == nil {
        request = NewGetGroupListPerClusterRequest()
    }
    response = NewGetGroupListPerClusterResponse()
    err = c.Send(request, response)
    return
}

func NewGetInvokeTxRequest() (request *GetInvokeTxRequest) {
    request = &GetInvokeTxRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetInvokeTx")
    return
}

func NewGetInvokeTxResponse() (response *GetInvokeTxResponse) {
    response = &GetInvokeTxResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Invoke异步调用结果查询
func (c *Client) GetInvokeTx(request *GetInvokeTxRequest) (response *GetInvokeTxResponse, err error) {
    if request == nil {
        request = NewGetInvokeTxRequest()
    }
    response = NewGetInvokeTxResponse()
    err = c.Send(request, response)
    return
}

func NewGetLatesdTransactionListRequest() (request *GetLatesdTransactionListRequest) {
    request = &GetLatesdTransactionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetLatesdTransactionList")
    return
}

func NewGetLatesdTransactionListResponse() (response *GetLatesdTransactionListResponse) {
    response = &GetLatesdTransactionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取最新交易列表
func (c *Client) GetLatesdTransactionList(request *GetLatesdTransactionListRequest) (response *GetLatesdTransactionListResponse, err error) {
    if request == nil {
        request = NewGetLatesdTransactionListRequest()
    }
    response = NewGetLatesdTransactionListResponse()
    err = c.Send(request, response)
    return
}

func NewGetLogDetailRequest() (request *GetLogDetailRequest) {
    request = &GetLogDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetLogDetail")
    return
}

func NewGetLogDetailResponse() (response *GetLogDetailResponse) {
    response = &GetLogDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看日志详情
func (c *Client) GetLogDetail(request *GetLogDetailRequest) (response *GetLogDetailResponse, err error) {
    if request == nil {
        request = NewGetLogDetailRequest()
    }
    response = NewGetLogDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetLogListRequest() (request *GetLogListRequest) {
    request = &GetLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetLogList")
    return
}

func NewGetLogListResponse() (response *GetLogListResponse) {
    response = &GetLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看日志列表
func (c *Client) GetLogList(request *GetLogListRequest) (response *GetLogListResponse, err error) {
    if request == nil {
        request = NewGetLogListRequest()
    }
    response = NewGetLogListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPeerListRequest() (request *GetPeerListRequest) {
    request = &GetPeerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPeerList")
    return
}

func NewGetPeerListResponse() (response *GetPeerListResponse) {
    response = &GetPeerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Peer列表
func (c *Client) GetPeerList(request *GetPeerListRequest) (response *GetPeerListResponse, err error) {
    if request == nil {
        request = NewGetPeerListRequest()
    }
    response = NewGetPeerListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPeerListForChannelRequest() (request *GetPeerListForChannelRequest) {
    request = &GetPeerListForChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPeerListForChannel")
    return
}

func NewGetPeerListForChannelResponse() (response *GetPeerListForChannelResponse) {
    response = &GetPeerListForChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询可加入通道的的节点列表
func (c *Client) GetPeerListForChannel(request *GetPeerListForChannelRequest) (response *GetPeerListForChannelResponse, err error) {
    if request == nil {
        request = NewGetPeerListForChannelRequest()
    }
    response = NewGetPeerListForChannelResponse()
    err = c.Send(request, response)
    return
}

func NewGetPeerListForCloudMonitorRequest() (request *GetPeerListForCloudMonitorRequest) {
    request = &GetPeerListForCloudMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPeerListForCloudMonitor")
    return
}

func NewGetPeerListForCloudMonitorResponse() (response *GetPeerListForCloudMonitorResponse) {
    response = &GetPeerListForCloudMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询组织下的节点列表（云监控实例）
func (c *Client) GetPeerListForCloudMonitor(request *GetPeerListForCloudMonitorRequest) (response *GetPeerListForCloudMonitorResponse, err error) {
    if request == nil {
        request = NewGetPeerListForCloudMonitorRequest()
    }
    response = NewGetPeerListForCloudMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewGetPeerListForInstallRequest() (request *GetPeerListForInstallRequest) {
    request = &GetPeerListForInstallRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPeerListForInstall")
    return
}

func NewGetPeerListForInstallResponse() (response *GetPeerListForInstallResponse) {
    response = &GetPeerListForInstallResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下拉Peer列表（合约安装）
func (c *Client) GetPeerListForInstall(request *GetPeerListForInstallRequest) (response *GetPeerListForInstallResponse, err error) {
    if request == nil {
        request = NewGetPeerListForInstallRequest()
    }
    response = NewGetPeerListForInstallResponse()
    err = c.Send(request, response)
    return
}

func NewGetPeerListPerChaincodeRequest() (request *GetPeerListPerChaincodeRequest) {
    request = &GetPeerListPerChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPeerListPerChaincode")
    return
}

func NewGetPeerListPerChaincodeResponse() (response *GetPeerListPerChaincodeResponse) {
    response = &GetPeerListPerChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询合约关联的Peer列表
func (c *Client) GetPeerListPerChaincode(request *GetPeerListPerChaincodeRequest) (response *GetPeerListPerChaincodeResponse, err error) {
    if request == nil {
        request = NewGetPeerListPerChaincodeRequest()
    }
    response = NewGetPeerListPerChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetPeerListPerChannelRequest() (request *GetPeerListPerChannelRequest) {
    request = &GetPeerListPerChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPeerListPerChannel")
    return
}

func NewGetPeerListPerChannelResponse() (response *GetPeerListPerChannelResponse) {
    response = &GetPeerListPerChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Channel已加入的Peer列表
func (c *Client) GetPeerListPerChannel(request *GetPeerListPerChannelRequest) (response *GetPeerListPerChannelResponse, err error) {
    if request == nil {
        request = NewGetPeerListPerChannelRequest()
    }
    response = NewGetPeerListPerChannelResponse()
    err = c.Send(request, response)
    return
}

func NewGetPrivateLinkDetailRequest() (request *GetPrivateLinkDetailRequest) {
    request = &GetPrivateLinkDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPrivateLinkDetail")
    return
}

func NewGetPrivateLinkDetailResponse() (response *GetPrivateLinkDetailResponse) {
    response = &GetPrivateLinkDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取私有连接详情
func (c *Client) GetPrivateLinkDetail(request *GetPrivateLinkDetailRequest) (response *GetPrivateLinkDetailResponse, err error) {
    if request == nil {
        request = NewGetPrivateLinkDetailRequest()
    }
    response = NewGetPrivateLinkDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetPrivateLinkListRequest() (request *GetPrivateLinkListRequest) {
    request = &GetPrivateLinkListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPrivateLinkList")
    return
}

func NewGetPrivateLinkListResponse() (response *GetPrivateLinkListResponse) {
    response = &GetPrivateLinkListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取私有连接列表
func (c *Client) GetPrivateLinkList(request *GetPrivateLinkListRequest) (response *GetPrivateLinkListResponse, err error) {
    if request == nil {
        request = NewGetPrivateLinkListRequest()
    }
    response = NewGetPrivateLinkListResponse()
    err = c.Send(request, response)
    return
}

func NewGetSrcPrefListRequest() (request *GetSrcPrefListRequest) {
    request = &GetSrcPrefListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetSrcPrefList")
    return
}

func NewGetSrcPrefListResponse() (response *GetSrcPrefListResponse) {
    response = &GetSrcPrefListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群节点实时监控趋势
func (c *Client) GetSrcPrefList(request *GetSrcPrefListRequest) (response *GetSrcPrefListResponse, err error) {
    if request == nil {
        request = NewGetSrcPrefListRequest()
    }
    response = NewGetSrcPrefListResponse()
    err = c.Send(request, response)
    return
}

func NewGetSrvLogDetailRequest() (request *GetSrvLogDetailRequest) {
    request = &GetSrvLogDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetSrvLogDetail")
    return
}

func NewGetSrvLogDetailResponse() (response *GetSrvLogDetailResponse) {
    response = &GetSrvLogDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群容器日志
func (c *Client) GetSrvLogDetail(request *GetSrvLogDetailRequest) (response *GetSrvLogDetailResponse, err error) {
    if request == nil {
        request = NewGetSrvLogDetailRequest()
    }
    response = NewGetSrvLogDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetSrvStatusListRequest() (request *GetSrvStatusListRequest) {
    request = &GetSrvStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetSrvStatusList")
    return
}

func NewGetSrvStatusListResponse() (response *GetSrvStatusListResponse) {
    response = &GetSrvStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群服务监控列表
func (c *Client) GetSrvStatusList(request *GetSrvStatusListRequest) (response *GetSrvStatusListResponse, err error) {
    if request == nil {
        request = NewGetSrvStatusListRequest()
    }
    response = NewGetSrvStatusListResponse()
    err = c.Send(request, response)
    return
}

func NewGetTbaasClusterListSummaryRequest() (request *GetTbaasClusterListSummaryRequest) {
    request = &GetTbaasClusterListSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTbaasClusterListSummary")
    return
}

func NewGetTbaasClusterListSummaryResponse() (response *GetTbaasClusterListSummaryResponse) {
    response = &GetTbaasClusterListSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取tbaas网络列表概览
func (c *Client) GetTbaasClusterListSummary(request *GetTbaasClusterListSummaryRequest) (response *GetTbaasClusterListSummaryResponse, err error) {
    if request == nil {
        request = NewGetTbaasClusterListSummaryRequest()
    }
    response = NewGetTbaasClusterListSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetTbaasClusterNumberSummaryRequest() (request *GetTbaasClusterNumberSummaryRequest) {
    request = &GetTbaasClusterNumberSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTbaasClusterNumberSummary")
    return
}

func NewGetTbaasClusterNumberSummaryResponse() (response *GetTbaasClusterNumberSummaryResponse) {
    response = &GetTbaasClusterNumberSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取tbaas网络数量统计
func (c *Client) GetTbaasClusterNumberSummary(request *GetTbaasClusterNumberSummaryRequest) (response *GetTbaasClusterNumberSummaryResponse, err error) {
    if request == nil {
        request = NewGetTbaasClusterNumberSummaryRequest()
    }
    response = NewGetTbaasClusterNumberSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetTbaasClusterSummaryRequest() (request *GetTbaasClusterSummaryRequest) {
    request = &GetTbaasClusterSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTbaasClusterSummary")
    return
}

func NewGetTbaasClusterSummaryResponse() (response *GetTbaasClusterSummaryResponse) {
    response = &GetTbaasClusterSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取tbaas网络概览
func (c *Client) GetTbaasClusterSummary(request *GetTbaasClusterSummaryRequest) (response *GetTbaasClusterSummaryResponse, err error) {
    if request == nil {
        request = NewGetTbaasClusterSummaryRequest()
    }
    response = NewGetTbaasClusterSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetTbaasKeySummaryRequest() (request *GetTbaasKeySummaryRequest) {
    request = &GetTbaasKeySummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTbaasKeySummary")
    return
}

func NewGetTbaasKeySummaryResponse() (response *GetTbaasKeySummaryResponse) {
    response = &GetTbaasKeySummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取tbaas关键指标概览
func (c *Client) GetTbaasKeySummary(request *GetTbaasKeySummaryRequest) (response *GetTbaasKeySummaryResponse, err error) {
    if request == nil {
        request = NewGetTbaasKeySummaryRequest()
    }
    response = NewGetTbaasKeySummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransactionDetailRequest() (request *GetTransactionDetailRequest) {
    request = &GetTransactionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransactionDetail")
    return
}

func NewGetTransactionDetailResponse() (response *GetTransactionDetailResponse) {
    response = &GetTransactionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取交易详情
func (c *Client) GetTransactionDetail(request *GetTransactionDetailRequest) (response *GetTransactionDetailResponse, err error) {
    if request == nil {
        request = NewGetTransactionDetailRequest()
    }
    response = NewGetTransactionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransactionSpeedRequest() (request *GetTransactionSpeedRequest) {
    request = &GetTransactionSpeedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransactionSpeed")
    return
}

func NewGetTransactionSpeedResponse() (response *GetTransactionSpeedResponse) {
    response = &GetTransactionSpeedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询通道新增交易数趋势
func (c *Client) GetTransactionSpeed(request *GetTransactionSpeedRequest) (response *GetTransactionSpeedResponse, err error) {
    if request == nil {
        request = NewGetTransactionSpeedRequest()
    }
    response = NewGetTransactionSpeedResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserAuthTypeRequest() (request *GetUserAuthTypeRequest) {
    request = &GetUserAuthTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetUserAuthType")
    return
}

func NewGetUserAuthTypeResponse() (response *GetUserAuthTypeResponse) {
    response = &GetUserAuthTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户认证类型
func (c *Client) GetUserAuthType(request *GetUserAuthTypeRequest) (response *GetUserAuthTypeResponse, err error) {
    if request == nil {
        request = NewGetUserAuthTypeRequest()
    }
    response = NewGetUserAuthTypeResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserVpcListRequest() (request *GetUserVpcListRequest) {
    request = &GetUserVpcListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetUserVpcList")
    return
}

func NewGetUserVpcListResponse() (response *GetUserVpcListResponse) {
    response = &GetUserVpcListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户私有网络列表
func (c *Client) GetUserVpcList(request *GetUserVpcListRequest) (response *GetUserVpcListResponse, err error) {
    if request == nil {
        request = NewGetUserVpcListRequest()
    }
    response = NewGetUserVpcListResponse()
    err = c.Send(request, response)
    return
}

func NewInitializeChaincodeRequest() (request *InitializeChaincodeRequest) {
    request = &InitializeChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "InitializeChaincode")
    return
}

func NewInitializeChaincodeResponse() (response *InitializeChaincodeResponse) {
    response = &InitializeChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 初始化合约
func (c *Client) InitializeChaincode(request *InitializeChaincodeRequest) (response *InitializeChaincodeResponse, err error) {
    if request == nil {
        request = NewInitializeChaincodeRequest()
    }
    response = NewInitializeChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewInstallChaincodeRequest() (request *InstallChaincodeRequest) {
    request = &InstallChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "InstallChaincode")
    return
}

func NewInstallChaincodeResponse() (response *InstallChaincodeResponse) {
    response = &InstallChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安装合约
func (c *Client) InstallChaincode(request *InstallChaincodeRequest) (response *InstallChaincodeResponse, err error) {
    if request == nil {
        request = NewInstallChaincodeRequest()
    }
    response = NewInstallChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewInviteChannelForCBCRequest() (request *InviteChannelForCBCRequest) {
    request = &InviteChannelForCBCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "InviteChannelForCBC")
    return
}

func NewInviteChannelForCBCResponse() (response *InviteChannelForCBCResponse) {
    response = &InviteChannelForCBCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 邀请组织加入联盟通道
func (c *Client) InviteChannelForCBC(request *InviteChannelForCBCRequest) (response *InviteChannelForCBCResponse, err error) {
    if request == nil {
        request = NewInviteChannelForCBCRequest()
    }
    response = NewInviteChannelForCBCResponse()
    err = c.Send(request, response)
    return
}

func NewInviteClusterMemberRequest() (request *InviteClusterMemberRequest) {
    request = &InviteClusterMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "InviteClusterMember")
    return
}

func NewInviteClusterMemberResponse() (response *InviteClusterMemberResponse) {
    response = &InviteClusterMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 邀请联盟成员加入网络
func (c *Client) InviteClusterMember(request *InviteClusterMemberRequest) (response *InviteClusterMemberResponse, err error) {
    if request == nil {
        request = NewInviteClusterMemberRequest()
    }
    response = NewInviteClusterMemberResponse()
    err = c.Send(request, response)
    return
}

func NewInviteConsortiumMemberRequest() (request *InviteConsortiumMemberRequest) {
    request = &InviteConsortiumMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "InviteConsortiumMember")
    return
}

func NewInviteConsortiumMemberResponse() (response *InviteConsortiumMemberResponse) {
    response = &InviteConsortiumMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 邀请成员加入联盟
func (c *Client) InviteConsortiumMember(request *InviteConsortiumMemberRequest) (response *InviteConsortiumMemberResponse, err error) {
    if request == nil {
        request = NewInviteConsortiumMemberRequest()
    }
    response = NewInviteConsortiumMemberResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeRequest() (request *InvokeRequest) {
    request = &InvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "Invoke")
    return
}

func NewInvokeResponse() (response *InvokeResponse) {
    response = &InvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增交易
func (c *Client) Invoke(request *InvokeRequest) (response *InvokeResponse, err error) {
    if request == nil {
        request = NewInvokeRequest()
    }
    response = NewInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewJoinChannelRequest() (request *JoinChannelRequest) {
    request = &JoinChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "JoinChannel")
    return
}

func NewJoinChannelResponse() (response *JoinChannelResponse) {
    response = &JoinChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 节点加入通道
func (c *Client) JoinChannel(request *JoinChannelRequest) (response *JoinChannelResponse, err error) {
    if request == nil {
        request = NewJoinChannelRequest()
    }
    response = NewJoinChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyChannelVoteRateForCBCRequest() (request *ModifyChannelVoteRateForCBCRequest) {
    request = &ModifyChannelVoteRateForCBCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ModifyChannelVoteRateForCBC")
    return
}

func NewModifyChannelVoteRateForCBCResponse() (response *ModifyChannelVoteRateForCBCResponse) {
    response = &ModifyChannelVoteRateForCBCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改联盟链通道投票率
func (c *Client) ModifyChannelVoteRateForCBC(request *ModifyChannelVoteRateForCBCRequest) (response *ModifyChannelVoteRateForCBCResponse, err error) {
    if request == nil {
        request = NewModifyChannelVoteRateForCBCRequest()
    }
    response = NewModifyChannelVoteRateForCBCResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterDescriptionRequest() (request *ModifyClusterDescriptionRequest) {
    request = &ModifyClusterDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ModifyClusterDescription")
    return
}

func NewModifyClusterDescriptionResponse() (response *ModifyClusterDescriptionResponse) {
    response = &ModifyClusterDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改网络描述
func (c *Client) ModifyClusterDescription(request *ModifyClusterDescriptionRequest) (response *ModifyClusterDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyClusterDescriptionRequest()
    }
    response = NewModifyClusterDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFabricClusterRequest() (request *ModifyFabricClusterRequest) {
    request = &ModifyFabricClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ModifyFabricCluster")
    return
}

func NewModifyFabricClusterResponse() (response *ModifyFabricClusterResponse) {
    response = &ModifyFabricClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增配网络资源（后付费）
func (c *Client) ModifyFabricCluster(request *ModifyFabricClusterRequest) (response *ModifyFabricClusterResponse, err error) {
    if request == nil {
        request = NewModifyFabricClusterRequest()
    }
    response = NewModifyFabricClusterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFieldDescriptionRequest() (request *ModifyFieldDescriptionRequest) {
    request = &ModifyFieldDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ModifyFieldDescription")
    return
}

func NewModifyFieldDescriptionResponse() (response *ModifyFieldDescriptionResponse) {
    response = &ModifyFieldDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改联盟、网络、通道、策略等描述字段
func (c *Client) ModifyFieldDescription(request *ModifyFieldDescriptionRequest) (response *ModifyFieldDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyFieldDescriptionRequest()
    }
    response = NewModifyFieldDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateLinkRequest() (request *ModifyPrivateLinkRequest) {
    request = &ModifyPrivateLinkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ModifyPrivateLink")
    return
}

func NewModifyPrivateLinkResponse() (response *ModifyPrivateLinkResponse) {
    response = &ModifyPrivateLinkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改私有链接名称
func (c *Client) ModifyPrivateLink(request *ModifyPrivateLinkRequest) (response *ModifyPrivateLinkResponse, err error) {
    if request == nil {
        request = NewModifyPrivateLinkRequest()
    }
    response = NewModifyPrivateLinkResponse()
    err = c.Send(request, response)
    return
}

func NewPreFeeGetModifyPriceRequest() (request *PreFeeGetModifyPriceRequest) {
    request = &PreFeeGetModifyPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "PreFeeGetModifyPrice")
    return
}

func NewPreFeeGetModifyPriceResponse() (response *PreFeeGetModifyPriceResponse) {
    response = &PreFeeGetModifyPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 预付费变配询价
func (c *Client) PreFeeGetModifyPrice(request *PreFeeGetModifyPriceRequest) (response *PreFeeGetModifyPriceResponse, err error) {
    if request == nil {
        request = NewPreFeeGetModifyPriceRequest()
    }
    response = NewPreFeeGetModifyPriceResponse()
    err = c.Send(request, response)
    return
}

func NewPreFeeGetPriceRequest() (request *PreFeeGetPriceRequest) {
    request = &PreFeeGetPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "PreFeeGetPrice")
    return
}

func NewPreFeeGetPriceResponse() (response *PreFeeGetPriceResponse) {
    response = &PreFeeGetPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 预付费询价
func (c *Client) PreFeeGetPrice(request *PreFeeGetPriceRequest) (response *PreFeeGetPriceResponse, err error) {
    if request == nil {
        request = NewPreFeeGetPriceRequest()
    }
    response = NewPreFeeGetPriceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRequest() (request *QueryRequest) {
    request = &QueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "Query")
    return
}

func NewQueryResponse() (response *QueryResponse) {
    response = &QueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询交易
func (c *Client) Query(request *QueryRequest) (response *QueryResponse, err error) {
    if request == nil {
        request = NewQueryRequest()
    }
    response = NewQueryResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOnlineEditorAccessRequest() (request *QueryOnlineEditorAccessRequest) {
    request = &QueryOnlineEditorAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryOnlineEditorAccess")
    return
}

func NewQueryOnlineEditorAccessResponse() (response *QueryOnlineEditorAccessResponse) {
    response = &QueryOnlineEditorAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在线编辑器访问权限查询
func (c *Client) QueryOnlineEditorAccess(request *QueryOnlineEditorAccessRequest) (response *QueryOnlineEditorAccessResponse, err error) {
    if request == nil {
        request = NewQueryOnlineEditorAccessRequest()
    }
    response = NewQueryOnlineEditorAccessResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPlatFormResourceRequest() (request *QueryPlatFormResourceRequest) {
    request = &QueryPlatFormResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryPlatFormResource")
    return
}

func NewQueryPlatFormResourceResponse() (response *QueryPlatFormResourceResponse) {
    response = &QueryPlatFormResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// tbass平台类型查询
func (c *Client) QueryPlatFormResource(request *QueryPlatFormResourceRequest) (response *QueryPlatFormResourceResponse, err error) {
    if request == nil {
        request = NewQueryPlatFormResourceRequest()
    }
    response = NewQueryPlatFormResourceResponse()
    err = c.Send(request, response)
    return
}

func NewReissueCertRequest() (request *ReissueCertRequest) {
    request = &ReissueCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ReissueCert")
    return
}

func NewReissueCertResponse() (response *ReissueCertResponse) {
    response = &ReissueCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 证书重发
func (c *Client) ReissueCert(request *ReissueCertRequest) (response *ReissueCertResponse, err error) {
    if request == nil {
        request = NewReissueCertRequest()
    }
    response = NewReissueCertResponse()
    err = c.Send(request, response)
    return
}

func NewRunChaincodeDevRequest() (request *RunChaincodeDevRequest) {
    request = &RunChaincodeDevRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "RunChaincodeDev")
    return
}

func NewRunChaincodeDevResponse() (response *RunChaincodeDevResponse) {
    response = &RunChaincodeDevResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IDE智能合约执行
func (c *Client) RunChaincodeDev(request *RunChaincodeDevRequest) (response *RunChaincodeDevResponse, err error) {
    if request == nil {
        request = NewRunChaincodeDevRequest()
    }
    response = NewRunChaincodeDevResponse()
    err = c.Send(request, response)
    return
}

func NewSetChannelAnchorPeerRequest() (request *SetChannelAnchorPeerRequest) {
    request = &SetChannelAnchorPeerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "SetChannelAnchorPeer")
    return
}

func NewSetChannelAnchorPeerResponse() (response *SetChannelAnchorPeerResponse) {
    response = &SetChannelAnchorPeerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置通道的Anchor Peer
func (c *Client) SetChannelAnchorPeer(request *SetChannelAnchorPeerRequest) (response *SetChannelAnchorPeerResponse, err error) {
    if request == nil {
        request = NewSetChannelAnchorPeerRequest()
    }
    response = NewSetChannelAnchorPeerResponse()
    err = c.Send(request, response)
    return
}

func NewUploadChaincodeDevRequest() (request *UploadChaincodeDevRequest) {
    request = &UploadChaincodeDevRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "UploadChaincodeDev")
    return
}

func NewUploadChaincodeDevResponse() (response *UploadChaincodeDevResponse) {
    response = &UploadChaincodeDevResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IDE智能合约上传
func (c *Client) UploadChaincodeDev(request *UploadChaincodeDevRequest) (response *UploadChaincodeDevResponse, err error) {
    if request == nil {
        request = NewUploadChaincodeDevRequest()
    }
    response = NewUploadChaincodeDevResponse()
    err = c.Send(request, response)
    return
}

func NewUploadEventConfigRequest() (request *UploadEventConfigRequest) {
    request = &UploadEventConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "UploadEventConfig")
    return
}

func NewUploadEventConfigResponse() (response *UploadEventConfigResponse) {
    response = &UploadEventConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 上传事件配置
func (c *Client) UploadEventConfig(request *UploadEventConfigRequest) (response *UploadEventConfigResponse, err error) {
    if request == nil {
        request = NewUploadEventConfigRequest()
    }
    response = NewUploadEventConfigResponse()
    err = c.Send(request, response)
    return
}
