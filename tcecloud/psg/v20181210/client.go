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

package v20181210

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-12-10"

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


func NewAccessRuleInterfaceIDeleteRuleRequest() (request *AccessRuleInterfaceIDeleteRuleRequest) {
    request = &AccessRuleInterfaceIDeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "AccessRuleInterfaceIDeleteRule")
    return
}

func NewAccessRuleInterfaceIDeleteRuleResponse() (response *AccessRuleInterfaceIDeleteRuleResponse) {
    response = &AccessRuleInterfaceIDeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// {
// "access_type": "global",     # 管控类型，"global": "全局管控"; "precise": "精准管控"
// "type": "BLACK",             # 黑白名单类型。"WHITE"、"BLACK" 
// "data": [{
//   "id" : "2",                # 唯一ID
// }]
// }
// 
// 
// Data: [{Id: "11"}]
// NetSwitch: "delete"
// InfoType: ""
// InfoValue: ""
// 
func (c *Client) AccessRuleInterfaceIDeleteRule(request *AccessRuleInterfaceIDeleteRuleRequest) (response *AccessRuleInterfaceIDeleteRuleResponse, err error) {
    if request == nil {
        request = NewAccessRuleInterfaceIDeleteRuleRequest()
    }
    response = NewAccessRuleInterfaceIDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAccessRuleInterfaceIUpdateRuleStatusRequest() (request *AccessRuleInterfaceIUpdateRuleStatusRequest) {
    request = &AccessRuleInterfaceIUpdateRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "AccessRuleInterfaceIUpdateRuleStatus")
    return
}

func NewAccessRuleInterfaceIUpdateRuleStatusResponse() (response *AccessRuleInterfaceIUpdateRuleStatusResponse) {
    response = &AccessRuleInterfaceIUpdateRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// {
// "access_type": "global",     # 管控类型，"global": "全局管控"; "precise": "精准管控"
// "type": "BLACK",             # 黑白名单类型。"WHITE"、"BLACK" 
// "info_type": "status",       # 修改规则的信息类型。"status": 状态; "valid_duration": 有效期; "match_operation": 匹配动作[白名单没有该选项]。
// "info_value" : "1",          # info_type = "status"，取值为 "1"：生效；"0"：失效。
//                              # info_type = "valid_duration"，单位分钟，"-1": "永久性"。
//                              # info_type = "match_operation", 取值为 "1": "阻断"; "0": "告警"。[白名单没有该选项]
// "data": [{
//   "id" : "2",                # 唯一ID
// }],
// "reboot": true               # 是否重启规则。bool类型，true 表示重启; false 表示不重启。[info_value为"valid_duration"或者"match_operation"才有该选项]
// }
// 
// 
func (c *Client) AccessRuleInterfaceIUpdateRuleStatus(request *AccessRuleInterfaceIUpdateRuleStatusRequest) (response *AccessRuleInterfaceIUpdateRuleStatusResponse, err error) {
    if request == nil {
        request = NewAccessRuleInterfaceIUpdateRuleStatusRequest()
    }
    response = NewAccessRuleInterfaceIUpdateRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmInterfaceIGetSystemRawLogListRequest() (request *AlarmInterfaceIGetSystemRawLogListRequest) {
    request = &AlarmInterfaceIGetSystemRawLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "AlarmInterfaceIGetSystemRawLogList")
    return
}

func NewAlarmInterfaceIGetSystemRawLogListResponse() (response *AlarmInterfaceIGetSystemRawLogListResponse) {
    response = &AlarmInterfaceIGetSystemRawLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 故障监控-列表
func (c *Client) AlarmInterfaceIGetSystemRawLogList(request *AlarmInterfaceIGetSystemRawLogListRequest) (response *AlarmInterfaceIGetSystemRawLogListResponse, err error) {
    if request == nil {
        request = NewAlarmInterfaceIGetSystemRawLogListRequest()
    }
    response = NewAlarmInterfaceIGetSystemRawLogListResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmInterfaceIGetSystemRawLogStatisticInfoRequest() (request *AlarmInterfaceIGetSystemRawLogStatisticInfoRequest) {
    request = &AlarmInterfaceIGetSystemRawLogStatisticInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "AlarmInterfaceIGetSystemRawLogStatisticInfo")
    return
}

func NewAlarmInterfaceIGetSystemRawLogStatisticInfoResponse() (response *AlarmInterfaceIGetSystemRawLogStatisticInfoResponse) {
    response = &AlarmInterfaceIGetSystemRawLogStatisticInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 故障监控-筛选条件
func (c *Client) AlarmInterfaceIGetSystemRawLogStatisticInfo(request *AlarmInterfaceIGetSystemRawLogStatisticInfoRequest) (response *AlarmInterfaceIGetSystemRawLogStatisticInfoResponse, err error) {
    if request == nil {
        request = NewAlarmInterfaceIGetSystemRawLogStatisticInfoRequest()
    }
    response = NewAlarmInterfaceIGetSystemRawLogStatisticInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmInterfaceIGetThreatRawLogDetailRequest() (request *AlarmInterfaceIGetThreatRawLogDetailRequest) {
    request = &AlarmInterfaceIGetThreatRawLogDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "AlarmInterfaceIGetThreatRawLogDetail")
    return
}

func NewAlarmInterfaceIGetThreatRawLogDetailResponse() (response *AlarmInterfaceIGetThreatRawLogDetailResponse) {
    response = &AlarmInterfaceIGetThreatRawLogDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全告警-告警详情
func (c *Client) AlarmInterfaceIGetThreatRawLogDetail(request *AlarmInterfaceIGetThreatRawLogDetailRequest) (response *AlarmInterfaceIGetThreatRawLogDetailResponse, err error) {
    if request == nil {
        request = NewAlarmInterfaceIGetThreatRawLogDetailRequest()
    }
    response = NewAlarmInterfaceIGetThreatRawLogDetailResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmInterfaceIGetThreatRawLogHistogramRequest() (request *AlarmInterfaceIGetThreatRawLogHistogramRequest) {
    request = &AlarmInterfaceIGetThreatRawLogHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "AlarmInterfaceIGetThreatRawLogHistogram")
    return
}

func NewAlarmInterfaceIGetThreatRawLogHistogramResponse() (response *AlarmInterfaceIGetThreatRawLogHistogramResponse) {
    response = &AlarmInterfaceIGetThreatRawLogHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全告警-告警图形
func (c *Client) AlarmInterfaceIGetThreatRawLogHistogram(request *AlarmInterfaceIGetThreatRawLogHistogramRequest) (response *AlarmInterfaceIGetThreatRawLogHistogramResponse, err error) {
    if request == nil {
        request = NewAlarmInterfaceIGetThreatRawLogHistogramRequest()
    }
    response = NewAlarmInterfaceIGetThreatRawLogHistogramResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmInterfaceIGetThreatRawLogListRequest() (request *AlarmInterfaceIGetThreatRawLogListRequest) {
    request = &AlarmInterfaceIGetThreatRawLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "AlarmInterfaceIGetThreatRawLogList")
    return
}

func NewAlarmInterfaceIGetThreatRawLogListResponse() (response *AlarmInterfaceIGetThreatRawLogListResponse) {
    response = &AlarmInterfaceIGetThreatRawLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全告警-告警列表
func (c *Client) AlarmInterfaceIGetThreatRawLogList(request *AlarmInterfaceIGetThreatRawLogListRequest) (response *AlarmInterfaceIGetThreatRawLogListResponse, err error) {
    if request == nil {
        request = NewAlarmInterfaceIGetThreatRawLogListRequest()
    }
    response = NewAlarmInterfaceIGetThreatRawLogListResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmInterfaceIGetThreatRawLogStatisticInfoRequest() (request *AlarmInterfaceIGetThreatRawLogStatisticInfoRequest) {
    request = &AlarmInterfaceIGetThreatRawLogStatisticInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "AlarmInterfaceIGetThreatRawLogStatisticInfo")
    return
}

func NewAlarmInterfaceIGetThreatRawLogStatisticInfoResponse() (response *AlarmInterfaceIGetThreatRawLogStatisticInfoResponse) {
    response = &AlarmInterfaceIGetThreatRawLogStatisticInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全告警-筛选条件
func (c *Client) AlarmInterfaceIGetThreatRawLogStatisticInfo(request *AlarmInterfaceIGetThreatRawLogStatisticInfoRequest) (response *AlarmInterfaceIGetThreatRawLogStatisticInfoResponse, err error) {
    if request == nil {
        request = NewAlarmInterfaceIGetThreatRawLogStatisticInfoRequest()
    }
    response = NewAlarmInterfaceIGetThreatRawLogStatisticInfoResponse()
    err = c.Send(request, response)
    return
}

func NewBackendTaskInterfaceIGetDataRequest() (request *BackendTaskInterfaceIGetDataRequest) {
    request = &BackendTaskInterfaceIGetDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BackendTaskInterfaceIGetData")
    return
}

func NewBackendTaskInterfaceIGetDataResponse() (response *BackendTaskInterfaceIGetDataResponse) {
    response = &BackendTaskInterfaceIGetDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取后台任务数据
func (c *Client) BackendTaskInterfaceIGetData(request *BackendTaskInterfaceIGetDataRequest) (response *BackendTaskInterfaceIGetDataResponse, err error) {
    if request == nil {
        request = NewBackendTaskInterfaceIGetDataRequest()
    }
    response = NewBackendTaskInterfaceIGetDataResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIAddApiConfigRequest() (request *BlockConfigInterfaceIAddApiConfigRequest) {
    request = &BlockConfigInterfaceIAddApiConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIAddApiConfig")
    return
}

func NewBlockConfigInterfaceIAddApiConfigResponse() (response *BlockConfigInterfaceIAddApiConfigResponse) {
    response = &BlockConfigInterfaceIAddApiConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--添加厂商
func (c *Client) BlockConfigInterfaceIAddApiConfig(request *BlockConfigInterfaceIAddApiConfigRequest) (response *BlockConfigInterfaceIAddApiConfigResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIAddApiConfigRequest()
    }
    response = NewBlockConfigInterfaceIAddApiConfigResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIChangeBlockStatusRequest() (request *BlockConfigInterfaceIChangeBlockStatusRequest) {
    request = &BlockConfigInterfaceIChangeBlockStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIChangeBlockStatus")
    return
}

func NewBlockConfigInterfaceIChangeBlockStatusResponse() (response *BlockConfigInterfaceIChangeBlockStatusResponse) {
    response = &BlockConfigInterfaceIChangeBlockStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-关闭/开启阻断
func (c *Client) BlockConfigInterfaceIChangeBlockStatus(request *BlockConfigInterfaceIChangeBlockStatusRequest) (response *BlockConfigInterfaceIChangeBlockStatusResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIChangeBlockStatusRequest()
    }
    response = NewBlockConfigInterfaceIChangeBlockStatusResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIConfigBlockTimeRequest() (request *BlockConfigInterfaceIConfigBlockTimeRequest) {
    request = &BlockConfigInterfaceIConfigBlockTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIConfigBlockTime")
    return
}

func NewBlockConfigInterfaceIConfigBlockTimeResponse() (response *BlockConfigInterfaceIConfigBlockTimeResponse) {
    response = &BlockConfigInterfaceIConfigBlockTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-设置拉黑时长
func (c *Client) BlockConfigInterfaceIConfigBlockTime(request *BlockConfigInterfaceIConfigBlockTimeRequest) (response *BlockConfigInterfaceIConfigBlockTimeResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIConfigBlockTimeRequest()
    }
    response = NewBlockConfigInterfaceIConfigBlockTimeResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIConfigRedirectURLRequest() (request *BlockConfigInterfaceIConfigRedirectURLRequest) {
    request = &BlockConfigInterfaceIConfigRedirectURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIConfigRedirectURL")
    return
}

func NewBlockConfigInterfaceIConfigRedirectURLResponse() (response *BlockConfigInterfaceIConfigRedirectURLResponse) {
    response = &BlockConfigInterfaceIConfigRedirectURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-设置阻断URL
func (c *Client) BlockConfigInterfaceIConfigRedirectURL(request *BlockConfigInterfaceIConfigRedirectURLRequest) (response *BlockConfigInterfaceIConfigRedirectURLResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIConfigRedirectURLRequest()
    }
    response = NewBlockConfigInterfaceIConfigRedirectURLResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIDeleteApiConfigRequest() (request *BlockConfigInterfaceIDeleteApiConfigRequest) {
    request = &BlockConfigInterfaceIDeleteApiConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIDeleteApiConfig")
    return
}

func NewBlockConfigInterfaceIDeleteApiConfigResponse() (response *BlockConfigInterfaceIDeleteApiConfigResponse) {
    response = &BlockConfigInterfaceIDeleteApiConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--删除厂商
func (c *Client) BlockConfigInterfaceIDeleteApiConfig(request *BlockConfigInterfaceIDeleteApiConfigRequest) (response *BlockConfigInterfaceIDeleteApiConfigResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIDeleteApiConfigRequest()
    }
    response = NewBlockConfigInterfaceIDeleteApiConfigResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIDeleteStopNetworkConfigRequest() (request *BlockConfigInterfaceIDeleteStopNetworkConfigRequest) {
    request = &BlockConfigInterfaceIDeleteStopNetworkConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIDeleteStopNetworkConfig")
    return
}

func NewBlockConfigInterfaceIDeleteStopNetworkConfigResponse() (response *BlockConfigInterfaceIDeleteStopNetworkConfigResponse) {
    response = &BlockConfigInterfaceIDeleteStopNetworkConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑白名单配置模块--删除断网配置
func (c *Client) BlockConfigInterfaceIDeleteStopNetworkConfig(request *BlockConfigInterfaceIDeleteStopNetworkConfigRequest) (response *BlockConfigInterfaceIDeleteStopNetworkConfigResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIDeleteStopNetworkConfigRequest()
    }
    response = NewBlockConfigInterfaceIDeleteStopNetworkConfigResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIEditApiConfigRequest() (request *BlockConfigInterfaceIEditApiConfigRequest) {
    request = &BlockConfigInterfaceIEditApiConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIEditApiConfig")
    return
}

func NewBlockConfigInterfaceIEditApiConfigResponse() (response *BlockConfigInterfaceIEditApiConfigResponse) {
    response = &BlockConfigInterfaceIEditApiConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--编辑厂商
// 
// Id: 19
// WindowCount: "200"
// WindowTime: "30"
// IsUse: 1
// Platform: "未命名"
// Company: "未命名"
// Region: "chongqing"
// SecretId: "0********5"
// SecretIds: "0********5"
// SecretKey: "6********a"
// Version: "2018-12-10"
func (c *Client) BlockConfigInterfaceIEditApiConfig(request *BlockConfigInterfaceIEditApiConfigRequest) (response *BlockConfigInterfaceIEditApiConfigResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIEditApiConfigRequest()
    }
    response = NewBlockConfigInterfaceIEditApiConfigResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetApiBlockStatusRequest() (request *BlockConfigInterfaceIGetApiBlockStatusRequest) {
    request = &BlockConfigInterfaceIGetApiBlockStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetApiBlockStatus")
    return
}

func NewBlockConfigInterfaceIGetApiBlockStatusResponse() (response *BlockConfigInterfaceIGetApiBlockStatusResponse) {
    response = &BlockConfigInterfaceIGetApiBlockStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--获取API阻断开关状态
func (c *Client) BlockConfigInterfaceIGetApiBlockStatus(request *BlockConfigInterfaceIGetApiBlockStatusRequest) (response *BlockConfigInterfaceIGetApiBlockStatusResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetApiBlockStatusRequest()
    }
    response = NewBlockConfigInterfaceIGetApiBlockStatusResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetApiBlockTrustThresholdRequest() (request *BlockConfigInterfaceIGetApiBlockTrustThresholdRequest) {
    request = &BlockConfigInterfaceIGetApiBlockTrustThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetApiBlockTrustThreshold")
    return
}

func NewBlockConfigInterfaceIGetApiBlockTrustThresholdResponse() (response *BlockConfigInterfaceIGetApiBlockTrustThresholdResponse) {
    response = &BlockConfigInterfaceIGetApiBlockTrustThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-拉取置信度阈值
func (c *Client) BlockConfigInterfaceIGetApiBlockTrustThreshold(request *BlockConfigInterfaceIGetApiBlockTrustThresholdRequest) (response *BlockConfigInterfaceIGetApiBlockTrustThresholdResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetApiBlockTrustThresholdRequest()
    }
    response = NewBlockConfigInterfaceIGetApiBlockTrustThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetApiConfigDetailRequest() (request *BlockConfigInterfaceIGetApiConfigDetailRequest) {
    request = &BlockConfigInterfaceIGetApiConfigDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetApiConfigDetail")
    return
}

func NewBlockConfigInterfaceIGetApiConfigDetailResponse() (response *BlockConfigInterfaceIGetApiConfigDetailResponse) {
    response = &BlockConfigInterfaceIGetApiConfigDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--查看API阻断配置详情【2.1】
func (c *Client) BlockConfigInterfaceIGetApiConfigDetail(request *BlockConfigInterfaceIGetApiConfigDetailRequest) (response *BlockConfigInterfaceIGetApiConfigDetailResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetApiConfigDetailRequest()
    }
    response = NewBlockConfigInterfaceIGetApiConfigDetailResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetApiConfigInfoListRequest() (request *BlockConfigInterfaceIGetApiConfigInfoListRequest) {
    request = &BlockConfigInterfaceIGetApiConfigInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetApiConfigInfoList")
    return
}

func NewBlockConfigInterfaceIGetApiConfigInfoListResponse() (response *BlockConfigInterfaceIGetApiConfigInfoListResponse) {
    response = &BlockConfigInterfaceIGetApiConfigInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--获取API阻断配置信息列表
func (c *Client) BlockConfigInterfaceIGetApiConfigInfoList(request *BlockConfigInterfaceIGetApiConfigInfoListRequest) (response *BlockConfigInterfaceIGetApiConfigInfoListResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetApiConfigInfoListRequest()
    }
    response = NewBlockConfigInterfaceIGetApiConfigInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetAvailableCountRequest() (request *BlockConfigInterfaceIGetAvailableCountRequest) {
    request = &BlockConfigInterfaceIGetAvailableCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetAvailableCount")
    return
}

func NewBlockConfigInterfaceIGetAvailableCountResponse() (response *BlockConfigInterfaceIGetAvailableCountResponse) {
    response = &BlockConfigInterfaceIGetAvailableCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--获取剩余可用配额与用户调用数量上限
func (c *Client) BlockConfigInterfaceIGetAvailableCount(request *BlockConfigInterfaceIGetAvailableCountRequest) (response *BlockConfigInterfaceIGetAvailableCountResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetAvailableCountRequest()
    }
    response = NewBlockConfigInterfaceIGetAvailableCountResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetBlockTimeRequest() (request *BlockConfigInterfaceIGetBlockTimeRequest) {
    request = &BlockConfigInterfaceIGetBlockTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetBlockTime")
    return
}

func NewBlockConfigInterfaceIGetBlockTimeResponse() (response *BlockConfigInterfaceIGetBlockTimeResponse) {
    response = &BlockConfigInterfaceIGetBlockTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-获取拉黑时长
func (c *Client) BlockConfigInterfaceIGetBlockTime(request *BlockConfigInterfaceIGetBlockTimeRequest) (response *BlockConfigInterfaceIGetBlockTimeResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetBlockTimeRequest()
    }
    response = NewBlockConfigInterfaceIGetBlockTimeResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetCurrentBlockStateRequest() (request *BlockConfigInterfaceIGetCurrentBlockStateRequest) {
    request = &BlockConfigInterfaceIGetCurrentBlockStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetCurrentBlockState")
    return
}

func NewBlockConfigInterfaceIGetCurrentBlockStateResponse() (response *BlockConfigInterfaceIGetCurrentBlockStateResponse) {
    response = &BlockConfigInterfaceIGetCurrentBlockStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-获取当前阻断状态
func (c *Client) BlockConfigInterfaceIGetCurrentBlockState(request *BlockConfigInterfaceIGetCurrentBlockStateRequest) (response *BlockConfigInterfaceIGetCurrentBlockStateResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetCurrentBlockStateRequest()
    }
    response = NewBlockConfigInterfaceIGetCurrentBlockStateResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetNetoffConfigRequest() (request *BlockConfigInterfaceIGetNetoffConfigRequest) {
    request = &BlockConfigInterfaceIGetNetoffConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetNetoffConfig")
    return
}

func NewBlockConfigInterfaceIGetNetoffConfigResponse() (response *BlockConfigInterfaceIGetNetoffConfigResponse) {
    response = &BlockConfigInterfaceIGetNetoffConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-一键断网状态读取接口
func (c *Client) BlockConfigInterfaceIGetNetoffConfig(request *BlockConfigInterfaceIGetNetoffConfigRequest) (response *BlockConfigInterfaceIGetNetoffConfigResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetNetoffConfigRequest()
    }
    response = NewBlockConfigInterfaceIGetNetoffConfigResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetOrResetSecretKeyRequest() (request *BlockConfigInterfaceIGetOrResetSecretKeyRequest) {
    request = &BlockConfigInterfaceIGetOrResetSecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetOrResetSecretKey")
    return
}

func NewBlockConfigInterfaceIGetOrResetSecretKeyResponse() (response *BlockConfigInterfaceIGetOrResetSecretKeyResponse) {
    response = &BlockConfigInterfaceIGetOrResetSecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--获取新秘钥或重置密
func (c *Client) BlockConfigInterfaceIGetOrResetSecretKey(request *BlockConfigInterfaceIGetOrResetSecretKeyRequest) (response *BlockConfigInterfaceIGetOrResetSecretKeyResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetOrResetSecretKeyRequest()
    }
    response = NewBlockConfigInterfaceIGetOrResetSecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetRebootAvailableCountRequest() (request *BlockConfigInterfaceIGetRebootAvailableCountRequest) {
    request = &BlockConfigInterfaceIGetRebootAvailableCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetRebootAvailableCount")
    return
}

func NewBlockConfigInterfaceIGetRebootAvailableCountResponse() (response *BlockConfigInterfaceIGetRebootAvailableCountResponse) {
    response = &BlockConfigInterfaceIGetRebootAvailableCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--获取重启剩余可用配额与用户调用数量上限
func (c *Client) BlockConfigInterfaceIGetRebootAvailableCount(request *BlockConfigInterfaceIGetRebootAvailableCountRequest) (response *BlockConfigInterfaceIGetRebootAvailableCountResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetRebootAvailableCountRequest()
    }
    response = NewBlockConfigInterfaceIGetRebootAvailableCountResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetRedirectURLRequest() (request *BlockConfigInterfaceIGetRedirectURLRequest) {
    request = &BlockConfigInterfaceIGetRedirectURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetRedirectURL")
    return
}

func NewBlockConfigInterfaceIGetRedirectURLResponse() (response *BlockConfigInterfaceIGetRedirectURLResponse) {
    response = &BlockConfigInterfaceIGetRedirectURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-获取阻断URL
func (c *Client) BlockConfigInterfaceIGetRedirectURL(request *BlockConfigInterfaceIGetRedirectURLRequest) (response *BlockConfigInterfaceIGetRedirectURLResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetRedirectURLRequest()
    }
    response = NewBlockConfigInterfaceIGetRedirectURLResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIGetSecretInfoRequest() (request *BlockConfigInterfaceIGetSecretInfoRequest) {
    request = &BlockConfigInterfaceIGetSecretInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIGetSecretInfo")
    return
}

func NewBlockConfigInterfaceIGetSecretInfoResponse() (response *BlockConfigInterfaceIGetSecretInfoResponse) {
    response = &BlockConfigInterfaceIGetSecretInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--密码验证获取授权ID和秘钥
func (c *Client) BlockConfigInterfaceIGetSecretInfo(request *BlockConfigInterfaceIGetSecretInfoRequest) (response *BlockConfigInterfaceIGetSecretInfoResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIGetSecretInfoRequest()
    }
    response = NewBlockConfigInterfaceIGetSecretInfoResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIRebootApiRequest() (request *BlockConfigInterfaceIRebootApiRequest) {
    request = &BlockConfigInterfaceIRebootApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIRebootApi")
    return
}

func NewBlockConfigInterfaceIRebootApiResponse() (response *BlockConfigInterfaceIRebootApiResponse) {
    response = &BlockConfigInterfaceIRebootApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--重启第三方API
func (c *Client) BlockConfigInterfaceIRebootApi(request *BlockConfigInterfaceIRebootApiRequest) (response *BlockConfigInterfaceIRebootApiResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIRebootApiRequest()
    }
    response = NewBlockConfigInterfaceIRebootApiResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceISetApiBlockStatusRequest() (request *BlockConfigInterfaceISetApiBlockStatusRequest) {
    request = &BlockConfigInterfaceISetApiBlockStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceISetApiBlockStatus")
    return
}

func NewBlockConfigInterfaceISetApiBlockStatusResponse() (response *BlockConfigInterfaceISetApiBlockStatusResponse) {
    response = &BlockConfigInterfaceISetApiBlockStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--设置API阻断开关状态
func (c *Client) BlockConfigInterfaceISetApiBlockStatus(request *BlockConfigInterfaceISetApiBlockStatusRequest) (response *BlockConfigInterfaceISetApiBlockStatusResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceISetApiBlockStatusRequest()
    }
    response = NewBlockConfigInterfaceISetApiBlockStatusResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceISetApiConfigStatusRequest() (request *BlockConfigInterfaceISetApiConfigStatusRequest) {
    request = &BlockConfigInterfaceISetApiConfigStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceISetApiConfigStatus")
    return
}

func NewBlockConfigInterfaceISetApiConfigStatusResponse() (response *BlockConfigInterfaceISetApiConfigStatusResponse) {
    response = &BlockConfigInterfaceISetApiConfigStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 阻断模块--设置API阻断配置状态
func (c *Client) BlockConfigInterfaceISetApiConfigStatus(request *BlockConfigInterfaceISetApiConfigStatusRequest) (response *BlockConfigInterfaceISetApiConfigStatusResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceISetApiConfigStatusRequest()
    }
    response = NewBlockConfigInterfaceISetApiConfigStatusResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIStartNetworkRequest() (request *BlockConfigInterfaceIStartNetworkRequest) {
    request = &BlockConfigInterfaceIStartNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIStartNetwork")
    return
}

func NewBlockConfigInterfaceIStartNetworkResponse() (response *BlockConfigInterfaceIStartNetworkResponse) {
    response = &BlockConfigInterfaceIStartNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-恢复网络状态
func (c *Client) BlockConfigInterfaceIStartNetwork(request *BlockConfigInterfaceIStartNetworkRequest) (response *BlockConfigInterfaceIStartNetworkResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIStartNetworkRequest()
    }
    response = NewBlockConfigInterfaceIStartNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIStopNetworkRequest() (request *BlockConfigInterfaceIStopNetworkRequest) {
    request = &BlockConfigInterfaceIStopNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIStopNetwork")
    return
}

func NewBlockConfigInterfaceIStopNetworkResponse() (response *BlockConfigInterfaceIStopNetworkResponse) {
    response = &BlockConfigInterfaceIStopNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-设置断网状态
func (c *Client) BlockConfigInterfaceIStopNetwork(request *BlockConfigInterfaceIStopNetworkRequest) (response *BlockConfigInterfaceIStopNetworkResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIStopNetworkRequest()
    }
    response = NewBlockConfigInterfaceIStopNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewBlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest() (request *BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest) {
    request = &BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "BlockConfigInterfaceIUpdateApiBlockTrustThreshold")
    return
}

func NewBlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse() (response *BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse) {
    response = &BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 其他设置-设置置信度阈值
func (c *Client) BlockConfigInterfaceIUpdateApiBlockTrustThreshold(request *BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest) (response *BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse, err error) {
    if request == nil {
        request = NewBlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest()
    }
    response = NewBlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDayOverviewInterfaceIGetAlarmLogListRequest() (request *DayOverviewInterfaceIGetAlarmLogListRequest) {
    request = &DayOverviewInterfaceIGetAlarmLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "DayOverviewInterfaceIGetAlarmLogList")
    return
}

func NewDayOverviewInterfaceIGetAlarmLogListResponse() (response *DayOverviewInterfaceIGetAlarmLogListResponse) {
    response = &DayOverviewInterfaceIGetAlarmLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-告警日志
func (c *Client) DayOverviewInterfaceIGetAlarmLogList(request *DayOverviewInterfaceIGetAlarmLogListRequest) (response *DayOverviewInterfaceIGetAlarmLogListResponse, err error) {
    if request == nil {
        request = NewDayOverviewInterfaceIGetAlarmLogListRequest()
    }
    response = NewDayOverviewInterfaceIGetAlarmLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDayOverviewInterfaceIGetAlarmReasonCntRequest() (request *DayOverviewInterfaceIGetAlarmReasonCntRequest) {
    request = &DayOverviewInterfaceIGetAlarmReasonCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "DayOverviewInterfaceIGetAlarmReasonCnt")
    return
}

func NewDayOverviewInterfaceIGetAlarmReasonCntResponse() (response *DayOverviewInterfaceIGetAlarmReasonCntResponse) {
    response = &DayOverviewInterfaceIGetAlarmReasonCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-告警原因统计
func (c *Client) DayOverviewInterfaceIGetAlarmReasonCnt(request *DayOverviewInterfaceIGetAlarmReasonCntRequest) (response *DayOverviewInterfaceIGetAlarmReasonCntResponse, err error) {
    if request == nil {
        request = NewDayOverviewInterfaceIGetAlarmReasonCntRequest()
    }
    response = NewDayOverviewInterfaceIGetAlarmReasonCntResponse()
    err = c.Send(request, response)
    return
}

func NewDayOverviewInterfaceIGetComplianceInfoRequest() (request *DayOverviewInterfaceIGetComplianceInfoRequest) {
    request = &DayOverviewInterfaceIGetComplianceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "DayOverviewInterfaceIGetComplianceInfo")
    return
}

func NewDayOverviewInterfaceIGetComplianceInfoResponse() (response *DayOverviewInterfaceIGetComplianceInfoResponse) {
    response = &DayOverviewInterfaceIGetComplianceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-获取合规检测
func (c *Client) DayOverviewInterfaceIGetComplianceInfo(request *DayOverviewInterfaceIGetComplianceInfoRequest) (response *DayOverviewInterfaceIGetComplianceInfoResponse, err error) {
    if request == nil {
        request = NewDayOverviewInterfaceIGetComplianceInfoRequest()
    }
    response = NewDayOverviewInterfaceIGetComplianceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDayOverviewInterfaceIGetNetFlowRequest() (request *DayOverviewInterfaceIGetNetFlowRequest) {
    request = &DayOverviewInterfaceIGetNetFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "DayOverviewInterfaceIGetNetFlow")
    return
}

func NewDayOverviewInterfaceIGetNetFlowResponse() (response *DayOverviewInterfaceIGetNetFlowResponse) {
    response = &DayOverviewInterfaceIGetNetFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-网络流量
func (c *Client) DayOverviewInterfaceIGetNetFlow(request *DayOverviewInterfaceIGetNetFlowRequest) (response *DayOverviewInterfaceIGetNetFlowResponse, err error) {
    if request == nil {
        request = NewDayOverviewInterfaceIGetNetFlowRequest()
    }
    response = NewDayOverviewInterfaceIGetNetFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDayOverviewInterfaceIGetOverviewImportantInfoRequest() (request *DayOverviewInterfaceIGetOverviewImportantInfoRequest) {
    request = &DayOverviewInterfaceIGetOverviewImportantInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "DayOverviewInterfaceIGetOverviewImportantInfo")
    return
}

func NewDayOverviewInterfaceIGetOverviewImportantInfoResponse() (response *DayOverviewInterfaceIGetOverviewImportantInfoResponse) {
    response = &DayOverviewInterfaceIGetOverviewImportantInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-获取重点关注信息
func (c *Client) DayOverviewInterfaceIGetOverviewImportantInfo(request *DayOverviewInterfaceIGetOverviewImportantInfoRequest) (response *DayOverviewInterfaceIGetOverviewImportantInfoResponse, err error) {
    if request == nil {
        request = NewDayOverviewInterfaceIGetOverviewImportantInfoRequest()
    }
    response = NewDayOverviewInterfaceIGetOverviewImportantInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDayOverviewInterfaceIGetPlatformControlInfoRequest() (request *DayOverviewInterfaceIGetPlatformControlInfoRequest) {
    request = &DayOverviewInterfaceIGetPlatformControlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "DayOverviewInterfaceIGetPlatformControlInfo")
    return
}

func NewDayOverviewInterfaceIGetPlatformControlInfoResponse() (response *DayOverviewInterfaceIGetPlatformControlInfoResponse) {
    response = &DayOverviewInterfaceIGetPlatformControlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-平台管控
func (c *Client) DayOverviewInterfaceIGetPlatformControlInfo(request *DayOverviewInterfaceIGetPlatformControlInfoRequest) (response *DayOverviewInterfaceIGetPlatformControlInfoResponse, err error) {
    if request == nil {
        request = NewDayOverviewInterfaceIGetPlatformControlInfoRequest()
    }
    response = NewDayOverviewInterfaceIGetPlatformControlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGlobalAccessInterfaceIAddCheckTaskRequest() (request *GlobalAccessInterfaceIAddCheckTaskRequest) {
    request = &GlobalAccessInterfaceIAddCheckTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "GlobalAccessInterfaceIAddCheckTask")
    return
}

func NewGlobalAccessInterfaceIAddCheckTaskResponse() (response *GlobalAccessInterfaceIAddCheckTaskResponse) {
    response = &GlobalAccessInterfaceIAddCheckTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// {
// "type" : "WHITE",            # 黑白名单类型。"WHITE"、"BLACK"。
// "data": [{
//   "name": "规则1",            # 规则名称
//   "src_ip" : "192.128.1.1",   # 源IP
//   "dst_ip": "",               # 目的IP
//   "valid_duration": "30分钟", # 有效期
//   "match_operation": "阻断",  # 匹配动作。可选范围为“告警”、“阻断”   
//   "source_platform", "腾讯",  # 来源平台。
//   }]
// }
func (c *Client) GlobalAccessInterfaceIAddCheckTask(request *GlobalAccessInterfaceIAddCheckTaskRequest) (response *GlobalAccessInterfaceIAddCheckTaskResponse, err error) {
    if request == nil {
        request = NewGlobalAccessInterfaceIAddCheckTaskRequest()
    }
    response = NewGlobalAccessInterfaceIAddCheckTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGlobalAccessInterfaceIAddRuleRequest() (request *GlobalAccessInterfaceIAddRuleRequest) {
    request = &GlobalAccessInterfaceIAddRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "GlobalAccessInterfaceIAddRule")
    return
}

func NewGlobalAccessInterfaceIAddRuleResponse() (response *GlobalAccessInterfaceIAddRuleResponse) {
    response = &GlobalAccessInterfaceIAddRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 全局管控--页面添加规则（2.1）
// {
// "type" : "WHITE",             # 黑白名单类型。"WHITE"、"BLACK"。
// "name": "规则1",              # 规则名称
// "valid_duration": 30,        # 有效期，单位分钟。"30": "30分钟"; "60": "1小时"; "3600": "6小时"; "1440": "24小时"; "-1": "永久性"。
// "match_operation": "1",      # 匹配动作。"1": "阻断"; "0": "告警"; "2": "放行" 
// "source_platform", "腾讯",    # 来源平台。
// "is_overwrite": "1",          # 是否覆盖。"1": "是"; "0": "否"
// "ip_type": "source",          # IP匹配类型。"source":"仅源IP"；"target": "仅目的IP":； "source_target": "源IP+目的IP"
// "ip_data": [{
//   "src_ip": "192.128.1.1",    # 源IP
//   "dst_ip": "",               # 目的IP
// }]
// }
func (c *Client) GlobalAccessInterfaceIAddRule(request *GlobalAccessInterfaceIAddRuleRequest) (response *GlobalAccessInterfaceIAddRuleResponse, err error) {
    if request == nil {
        request = NewGlobalAccessInterfaceIAddRuleRequest()
    }
    response = NewGlobalAccessInterfaceIAddRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGlobalAccessInterfaceICommitCheckTaskRuleRequest() (request *GlobalAccessInterfaceICommitCheckTaskRuleRequest) {
    request = &GlobalAccessInterfaceICommitCheckTaskRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "GlobalAccessInterfaceICommitCheckTaskRule")
    return
}

func NewGlobalAccessInterfaceICommitCheckTaskRuleResponse() (response *GlobalAccessInterfaceICommitCheckTaskRuleResponse) {
    response = &GlobalAccessInterfaceICommitCheckTaskRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// "task_id": "1"                 # 任务ID
// "check_status": "success",     # 规则的检测状态类型，"success": 继续导入成功的规则；"warning": 继续导入冲突的规则。
func (c *Client) GlobalAccessInterfaceICommitCheckTaskRule(request *GlobalAccessInterfaceICommitCheckTaskRuleRequest) (response *GlobalAccessInterfaceICommitCheckTaskRuleResponse, err error) {
    if request == nil {
        request = NewGlobalAccessInterfaceICommitCheckTaskRuleRequest()
    }
    response = NewGlobalAccessInterfaceICommitCheckTaskRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGlobalAccessInterfaceIEditRuleRequest() (request *GlobalAccessInterfaceIEditRuleRequest) {
    request = &GlobalAccessInterfaceIEditRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "GlobalAccessInterfaceIEditRule")
    return
}

func NewGlobalAccessInterfaceIEditRuleResponse() (response *GlobalAccessInterfaceIEditRuleResponse) {
    response = &GlobalAccessInterfaceIEditRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// {
// "type" : "WHITE",             # 黑白名单类型。"WHITE"、"BLACK"。
// "id": 10,                     # 唯一标识ID
// "name": "规则1",              # 规则名称
// "valid_duration": 0.5,        # 有效期，单位分钟。"30": "30分钟"; "60": "1小时"; "3600": "6小时"; "1440": "24小时"; "-1": "永久性"。
// "match_operation": "1",       # 匹配动作。"1": "阻断"; "0": "告警"; "2": "放行"   
// "source_platform", "腾讯",    # 来源平台。
// "is_overwrite": "1",          # 是否覆盖。"1": "是"; "0": "否"
// "ip_type": "source",          # IP匹配类型。"source":"仅源IP"；"target": "仅目的IP":； "source_target": "源IP+目的IP"
// "ip_data": [{
//   "src_ip": "192.128.1.1",    # 源IP
//   "dst_ip": "",               # 目的IP
// }],
// "reboot": true,               # 是否重启规则。bool类型，true 表示重启; false 表示不重启
// }
func (c *Client) GlobalAccessInterfaceIEditRule(request *GlobalAccessInterfaceIEditRuleRequest) (response *GlobalAccessInterfaceIEditRuleResponse, err error) {
    if request == nil {
        request = NewGlobalAccessInterfaceIEditRuleRequest()
    }
    response = NewGlobalAccessInterfaceIEditRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGlobalAccessInterfaceIGetCheckTaskResultRequest() (request *GlobalAccessInterfaceIGetCheckTaskResultRequest) {
    request = &GlobalAccessInterfaceIGetCheckTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "GlobalAccessInterfaceIGetCheckTaskResult")
    return
}

func NewGlobalAccessInterfaceIGetCheckTaskResultResponse() (response *GlobalAccessInterfaceIGetCheckTaskResultResponse) {
    response = &GlobalAccessInterfaceIGetCheckTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// task_id
func (c *Client) GlobalAccessInterfaceIGetCheckTaskResult(request *GlobalAccessInterfaceIGetCheckTaskResultRequest) (response *GlobalAccessInterfaceIGetCheckTaskResultResponse, err error) {
    if request == nil {
        request = NewGlobalAccessInterfaceIGetCheckTaskResultRequest()
    }
    response = NewGlobalAccessInterfaceIGetCheckTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewGlobalAccessInterfaceIGetCheckTaskResultListRequest() (request *GlobalAccessInterfaceIGetCheckTaskResultListRequest) {
    request = &GlobalAccessInterfaceIGetCheckTaskResultListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "GlobalAccessInterfaceIGetCheckTaskResultList")
    return
}

func NewGlobalAccessInterfaceIGetCheckTaskResultListResponse() (response *GlobalAccessInterfaceIGetCheckTaskResultListResponse) {
    response = &GlobalAccessInterfaceIGetCheckTaskResultListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// {
// "task_id": "1",                # 任务ID
// "check_status": "failure",     # 规则的检测状态。 "failure": 检测失败；"warning": 检测冲突。
// }
func (c *Client) GlobalAccessInterfaceIGetCheckTaskResultList(request *GlobalAccessInterfaceIGetCheckTaskResultListRequest) (response *GlobalAccessInterfaceIGetCheckTaskResultListResponse, err error) {
    if request == nil {
        request = NewGlobalAccessInterfaceIGetCheckTaskResultListRequest()
    }
    response = NewGlobalAccessInterfaceIGetCheckTaskResultListResponse()
    err = c.Send(request, response)
    return
}

func NewGlobalAccessInterfaceIGetRuleListRequest() (request *GlobalAccessInterfaceIGetRuleListRequest) {
    request = &GlobalAccessInterfaceIGetRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "GlobalAccessInterfaceIGetRuleList")
    return
}

func NewGlobalAccessInterfaceIGetRuleListResponse() (response *GlobalAccessInterfaceIGetRuleListResponse) {
    response = &GlobalAccessInterfaceIGetRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// {
// "type" : "WHITE",                       # 黑白名单类型。"WHITE"、"BLACK"。
// "search_content": "192.182.2.2.2",      # 规则名称。不选时传空字符串""。
// "begin_time" : "2018-10-10 10:00:00",   # 开始时间。不选时传空字符串""。
// "end_time": "2018-10-10 10:00:00",      # 结束时间。不选时传空字符串""。
// "status": "1",                          # 状态，"1"：正在使用；"0"：已失效。不选时传空字符串""。
// "valid_duration": "30",                 # 有效期。"30": "30分钟"; "60": "1小时"; "3600": "6小时"; "1440": "24小时"; "-1": "永久性"。不选时传空字符串""。   
// "hit_history", "1",                     # 历史命中数。"0": "0次"，"1": "1次及以上"。不选时传空字符串""。【仅黑名单有该字段】
// "hit_week": "0",                        # 最近一周命中数。"0": "0次"，"1": "1次及以上"。不选时传空字符串""。【仅黑名单有该字段】
// "hit_day": "",                          # 最近一天命中数。"0": "0次"，"1": "1次及以上"。不选时传空字符串""。【仅黑名单有该字段】                       
// "match_operation": "1",                 # 匹配动作。"1": "阻断"; "0": "告警"; "2": "放行"。不选时传空字符串""。
// "ip_type": "source",                    # IP匹配类型。"source":"仅源IP"；"target": "仅目的IP":； "source_target": "源IP+目的IP"
// "source_platform": "腾讯",              # 来源平台。不选时传空字符串""。
// "order": {"field": "create_time", "direc": "asc"},   # 排序。field表示要排序字段；direc表示排序方向，"asc"表示升序，"desc"表示降序
// "per_page": "",                       
// "current_page": ""                      # 分页。从1开始。
// }
func (c *Client) GlobalAccessInterfaceIGetRuleList(request *GlobalAccessInterfaceIGetRuleListRequest) (response *GlobalAccessInterfaceIGetRuleListResponse, err error) {
    if request == nil {
        request = NewGlobalAccessInterfaceIGetRuleListRequest()
    }
    response = NewGlobalAccessInterfaceIGetRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewGlobalAccessInterfaceIGetStatisticsRequest() (request *GlobalAccessInterfaceIGetStatisticsRequest) {
    request = &GlobalAccessInterfaceIGetStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "GlobalAccessInterfaceIGetStatistics")
    return
}

func NewGlobalAccessInterfaceIGetStatisticsResponse() (response *GlobalAccessInterfaceIGetStatisticsResponse) {
    response = &GlobalAccessInterfaceIGetStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// request:
// {
// "Type" : "WHITE",                       # 黑白名单类型。"WHITE"、"BLACK"。
// "SearchContent": "192.182.2.2.2",      # 规则名称。不选时传空字符串""。
// "BeginTime" : "2018-10-10 10:00:00",   # 开始时间。不选时传空字符串""。
// "EndTime": "2018-10-10 10:00:00",      # 结束时间。不选时传空字符串""。
// "Status": "1",                          # 状态，"1"：正在使用；"0"：已失效。不选时传空字符串""。
// "ValidDuration": "30",                # 有效期。"30": "30分钟"; "60": "1小时"; "3600": "6小时"; "1440": "24小时"; "-1": "永久性"。不选时传空字符串""。   
// "HitHistory", "1",                     # 历史命中数。"0": "0次"，"1": "1次及以上"。不选时传空字符串""。【仅黑名单有该字段】
// "HitWeek": "0",                        # 最近一周命中数。"0": "0次"，"1": "1次及以上"。不选时传空字符串""。【仅黑名单有该字段】
// "HitDay": "",                          # 最近一天命中数。"0": "0次"，"1": "1次及以上"。不选时传空字符串""。【仅黑名单有该字段】                       
// "MatchOperation": "1",                 # 匹配动作。"1": "阻断"; "0": "告警"; "2": "放行"。不选时传空字符串""。
// "IpType":"source",                     # IP匹配类型。"source": 仅源IP; "target": 仅目的IP; "source_target": 源IP+目的IP
// "SourcePlatform": "腾讯",              # 来源平台。不选时传空字符串""。                       
// }
// AddTime: []
// BeginTime: ""
// CurrentPage: 1
// EndTime: ""
// ExcuteFlag: 0
// InputCond: ""
// Order: {Field: "", Direc: ""}
// PerPage: 20
// Region: "chongqing"
// SearchContent: ""
// Status: ""
// Type: "BLACK"
func (c *Client) GlobalAccessInterfaceIGetStatistics(request *GlobalAccessInterfaceIGetStatisticsRequest) (response *GlobalAccessInterfaceIGetStatisticsResponse, err error) {
    if request == nil {
        request = NewGlobalAccessInterfaceIGetStatisticsRequest()
    }
    response = NewGlobalAccessInterfaceIGetStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewLabelManagerInterfaceIAddSourcePlatformRequest() (request *LabelManagerInterfaceIAddSourcePlatformRequest) {
    request = &LabelManagerInterfaceIAddSourcePlatformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "LabelManagerInterfaceIAddSourcePlatform")
    return
}

func NewLabelManagerInterfaceIAddSourcePlatformResponse() (response *LabelManagerInterfaceIAddSourcePlatformResponse) {
    response = &LabelManagerInterfaceIAddSourcePlatformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 标签管理--增加来源平台标签（2.1）
func (c *Client) LabelManagerInterfaceIAddSourcePlatform(request *LabelManagerInterfaceIAddSourcePlatformRequest) (response *LabelManagerInterfaceIAddSourcePlatformResponse, err error) {
    if request == nil {
        request = NewLabelManagerInterfaceIAddSourcePlatformRequest()
    }
    response = NewLabelManagerInterfaceIAddSourcePlatformResponse()
    err = c.Send(request, response)
    return
}

func NewLabelManagerInterfaceIDeleteSourcePlatformRequest() (request *LabelManagerInterfaceIDeleteSourcePlatformRequest) {
    request = &LabelManagerInterfaceIDeleteSourcePlatformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "LabelManagerInterfaceIDeleteSourcePlatform")
    return
}

func NewLabelManagerInterfaceIDeleteSourcePlatformResponse() (response *LabelManagerInterfaceIDeleteSourcePlatformResponse) {
    response = &LabelManagerInterfaceIDeleteSourcePlatformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 标签管理--删除来源平台标签（2.1）
func (c *Client) LabelManagerInterfaceIDeleteSourcePlatform(request *LabelManagerInterfaceIDeleteSourcePlatformRequest) (response *LabelManagerInterfaceIDeleteSourcePlatformResponse, err error) {
    if request == nil {
        request = NewLabelManagerInterfaceIDeleteSourcePlatformRequest()
    }
    response = NewLabelManagerInterfaceIDeleteSourcePlatformResponse()
    err = c.Send(request, response)
    return
}

func NewLabelManagerInterfaceIGetSourcePlatformRequest() (request *LabelManagerInterfaceIGetSourcePlatformRequest) {
    request = &LabelManagerInterfaceIGetSourcePlatformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "LabelManagerInterfaceIGetSourcePlatform")
    return
}

func NewLabelManagerInterfaceIGetSourcePlatformResponse() (response *LabelManagerInterfaceIGetSourcePlatformResponse) {
    response = &LabelManagerInterfaceIGetSourcePlatformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 标签管理--获取平台标签（2.1）
func (c *Client) LabelManagerInterfaceIGetSourcePlatform(request *LabelManagerInterfaceIGetSourcePlatformRequest) (response *LabelManagerInterfaceIGetSourcePlatformResponse, err error) {
    if request == nil {
        request = NewLabelManagerInterfaceIGetSourcePlatformRequest()
    }
    response = NewLabelManagerInterfaceIGetSourcePlatformResponse()
    err = c.Send(request, response)
    return
}

func NewMsgInterfaceIGetMsgListRequest() (request *MsgInterfaceIGetMsgListRequest) {
    request = &MsgInterfaceIGetMsgListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "MsgInterfaceIGetMsgList")
    return
}

func NewMsgInterfaceIGetMsgListResponse() (response *MsgInterfaceIGetMsgListResponse) {
    response = &MsgInterfaceIGetMsgListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通知接口--获取消息列表
//  "status": "已读" or "未读" or "all"  
//     "current_page": // 当前第几页
//     "per_page": 0 // 每页多少行
// CurrentPage: 1
// ExcuteFlag: 1
// PerPage: 20
// Region: "chongqing"
// Status: "all"
func (c *Client) MsgInterfaceIGetMsgList(request *MsgInterfaceIGetMsgListRequest) (response *MsgInterfaceIGetMsgListResponse, err error) {
    if request == nil {
        request = NewMsgInterfaceIGetMsgListRequest()
    }
    response = NewMsgInterfaceIGetMsgListResponse()
    err = c.Send(request, response)
    return
}

func NewMsgInterfaceIGetMsgStatisticInfoRequest() (request *MsgInterfaceIGetMsgStatisticInfoRequest) {
    request = &MsgInterfaceIGetMsgStatisticInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "MsgInterfaceIGetMsgStatisticInfo")
    return
}

func NewMsgInterfaceIGetMsgStatisticInfoResponse() (response *MsgInterfaceIGetMsgStatisticInfoResponse) {
    response = &MsgInterfaceIGetMsgStatisticInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通知接口--消息统计
// CurrentPage: 1
// ExcuteFlag: 0
// PerPage: 20
// Status: "all"
func (c *Client) MsgInterfaceIGetMsgStatisticInfo(request *MsgInterfaceIGetMsgStatisticInfoRequest) (response *MsgInterfaceIGetMsgStatisticInfoResponse, err error) {
    if request == nil {
        request = NewMsgInterfaceIGetMsgStatisticInfoRequest()
    }
    response = NewMsgInterfaceIGetMsgStatisticInfoResponse()
    err = c.Send(request, response)
    return
}

func NewMsgInterfaceISetSingleUnreadMsgToReadRequest() (request *MsgInterfaceISetSingleUnreadMsgToReadRequest) {
    request = &MsgInterfaceISetSingleUnreadMsgToReadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "MsgInterfaceISetSingleUnreadMsgToRead")
    return
}

func NewMsgInterfaceISetSingleUnreadMsgToReadResponse() (response *MsgInterfaceISetSingleUnreadMsgToReadResponse) {
    response = &MsgInterfaceISetSingleUnreadMsgToReadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 消息通知接口--标记单条未读消息为已读
func (c *Client) MsgInterfaceISetSingleUnreadMsgToRead(request *MsgInterfaceISetSingleUnreadMsgToReadRequest) (response *MsgInterfaceISetSingleUnreadMsgToReadResponse, err error) {
    if request == nil {
        request = NewMsgInterfaceISetSingleUnreadMsgToReadRequest()
    }
    response = NewMsgInterfaceISetSingleUnreadMsgToReadResponse()
    err = c.Send(request, response)
    return
}

func NewMsgInterfaceISetUnreadMsgToReadRequest() (request *MsgInterfaceISetUnreadMsgToReadRequest) {
    request = &MsgInterfaceISetUnreadMsgToReadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "MsgInterfaceISetUnreadMsgToRead")
    return
}

func NewMsgInterfaceISetUnreadMsgToReadResponse() (response *MsgInterfaceISetUnreadMsgToReadResponse) {
    response = &MsgInterfaceISetUnreadMsgToReadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 一键标记所有未读消息为已读
func (c *Client) MsgInterfaceISetUnreadMsgToRead(request *MsgInterfaceISetUnreadMsgToReadRequest) (response *MsgInterfaceISetUnreadMsgToReadResponse, err error) {
    if request == nil {
        request = NewMsgInterfaceISetUnreadMsgToReadRequest()
    }
    response = NewMsgInterfaceISetUnreadMsgToReadResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIExportDataRequest() (request *OverviewInterfaceIExportDataRequest) {
    request = &OverviewInterfaceIExportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIExportData")
    return
}

func NewOverviewInterfaceIExportDataResponse() (response *OverviewInterfaceIExportDataResponse) {
    response = &OverviewInterfaceIExportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--导出数据报告
func (c *Client) OverviewInterfaceIExportData(request *OverviewInterfaceIExportDataRequest) (response *OverviewInterfaceIExportDataResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIExportDataRequest()
    }
    response = NewOverviewInterfaceIExportDataResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetAlarmLogListRequest() (request *OverviewInterfaceIGetAlarmLogListRequest) {
    request = &OverviewInterfaceIGetAlarmLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetAlarmLogList")
    return
}

func NewOverviewInterfaceIGetAlarmLogListResponse() (response *OverviewInterfaceIGetAlarmLogListResponse) {
    response = &OverviewInterfaceIGetAlarmLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-告警日志
func (c *Client) OverviewInterfaceIGetAlarmLogList(request *OverviewInterfaceIGetAlarmLogListRequest) (response *OverviewInterfaceIGetAlarmLogListResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetAlarmLogListRequest()
    }
    response = NewOverviewInterfaceIGetAlarmLogListResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetAlarmReasonCntRequest() (request *OverviewInterfaceIGetAlarmReasonCntRequest) {
    request = &OverviewInterfaceIGetAlarmReasonCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetAlarmReasonCnt")
    return
}

func NewOverviewInterfaceIGetAlarmReasonCntResponse() (response *OverviewInterfaceIGetAlarmReasonCntResponse) {
    response = &OverviewInterfaceIGetAlarmReasonCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-告警原因统计
func (c *Client) OverviewInterfaceIGetAlarmReasonCnt(request *OverviewInterfaceIGetAlarmReasonCntRequest) (response *OverviewInterfaceIGetAlarmReasonCntResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetAlarmReasonCntRequest()
    }
    response = NewOverviewInterfaceIGetAlarmReasonCntResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetAlarmStatInfoRequest() (request *OverviewInterfaceIGetAlarmStatInfoRequest) {
    request = &OverviewInterfaceIGetAlarmStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetAlarmStatInfo")
    return
}

func NewOverviewInterfaceIGetAlarmStatInfoResponse() (response *OverviewInterfaceIGetAlarmStatInfoResponse) {
    response = &OverviewInterfaceIGetAlarmStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--获取告警统计
func (c *Client) OverviewInterfaceIGetAlarmStatInfo(request *OverviewInterfaceIGetAlarmStatInfoRequest) (response *OverviewInterfaceIGetAlarmStatInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetAlarmStatInfoRequest()
    }
    response = NewOverviewInterfaceIGetAlarmStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetAlarmTrendStatInfoRequest() (request *OverviewInterfaceIGetAlarmTrendStatInfoRequest) {
    request = &OverviewInterfaceIGetAlarmTrendStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetAlarmTrendStatInfo")
    return
}

func NewOverviewInterfaceIGetAlarmTrendStatInfoResponse() (response *OverviewInterfaceIGetAlarmTrendStatInfoResponse) {
    response = &OverviewInterfaceIGetAlarmTrendStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--获取告警趋势
func (c *Client) OverviewInterfaceIGetAlarmTrendStatInfo(request *OverviewInterfaceIGetAlarmTrendStatInfoRequest) (response *OverviewInterfaceIGetAlarmTrendStatInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetAlarmTrendStatInfoRequest()
    }
    response = NewOverviewInterfaceIGetAlarmTrendStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetBlockStatInfoRequest() (request *OverviewInterfaceIGetBlockStatInfoRequest) {
    request = &OverviewInterfaceIGetBlockStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetBlockStatInfo")
    return
}

func NewOverviewInterfaceIGetBlockStatInfoResponse() (response *OverviewInterfaceIGetBlockStatInfoResponse) {
    response = &OverviewInterfaceIGetBlockStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--获取安全处置统计
func (c *Client) OverviewInterfaceIGetBlockStatInfo(request *OverviewInterfaceIGetBlockStatInfoRequest) (response *OverviewInterfaceIGetBlockStatInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetBlockStatInfoRequest()
    }
    response = NewOverviewInterfaceIGetBlockStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetBlockTrendStatInfoRequest() (request *OverviewInterfaceIGetBlockTrendStatInfoRequest) {
    request = &OverviewInterfaceIGetBlockTrendStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetBlockTrendStatInfo")
    return
}

func NewOverviewInterfaceIGetBlockTrendStatInfoResponse() (response *OverviewInterfaceIGetBlockTrendStatInfoResponse) {
    response = &OverviewInterfaceIGetBlockTrendStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--获取阻断趋势统计
func (c *Client) OverviewInterfaceIGetBlockTrendStatInfo(request *OverviewInterfaceIGetBlockTrendStatInfoRequest) (response *OverviewInterfaceIGetBlockTrendStatInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetBlockTrendStatInfoRequest()
    }
    response = NewOverviewInterfaceIGetBlockTrendStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetCheckDistribStatInfoRequest() (request *OverviewInterfaceIGetCheckDistribStatInfoRequest) {
    request = &OverviewInterfaceIGetCheckDistribStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetCheckDistribStatInfo")
    return
}

func NewOverviewInterfaceIGetCheckDistribStatInfoResponse() (response *OverviewInterfaceIGetCheckDistribStatInfoResponse) {
    response = &OverviewInterfaceIGetCheckDistribStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--获取安全检测分布图
func (c *Client) OverviewInterfaceIGetCheckDistribStatInfo(request *OverviewInterfaceIGetCheckDistribStatInfoRequest) (response *OverviewInterfaceIGetCheckDistribStatInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetCheckDistribStatInfoRequest()
    }
    response = NewOverviewInterfaceIGetCheckDistribStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetComplianceInfoRequest() (request *OverviewInterfaceIGetComplianceInfoRequest) {
    request = &OverviewInterfaceIGetComplianceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetComplianceInfo")
    return
}

func NewOverviewInterfaceIGetComplianceInfoResponse() (response *OverviewInterfaceIGetComplianceInfoResponse) {
    response = &OverviewInterfaceIGetComplianceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-获取合规检测
func (c *Client) OverviewInterfaceIGetComplianceInfo(request *OverviewInterfaceIGetComplianceInfoRequest) (response *OverviewInterfaceIGetComplianceInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetComplianceInfoRequest()
    }
    response = NewOverviewInterfaceIGetComplianceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetNetFlowRequest() (request *OverviewInterfaceIGetNetFlowRequest) {
    request = &OverviewInterfaceIGetNetFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetNetFlow")
    return
}

func NewOverviewInterfaceIGetNetFlowResponse() (response *OverviewInterfaceIGetNetFlowResponse) {
    response = &OverviewInterfaceIGetNetFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  今日总览-网络流量
func (c *Client) OverviewInterfaceIGetNetFlow(request *OverviewInterfaceIGetNetFlowRequest) (response *OverviewInterfaceIGetNetFlowResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetNetFlowRequest()
    }
    response = NewOverviewInterfaceIGetNetFlowResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetOverviewImportantInfoRequest() (request *OverviewInterfaceIGetOverviewImportantInfoRequest) {
    request = &OverviewInterfaceIGetOverviewImportantInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetOverviewImportantInfo")
    return
}

func NewOverviewInterfaceIGetOverviewImportantInfoResponse() (response *OverviewInterfaceIGetOverviewImportantInfoResponse) {
    response = &OverviewInterfaceIGetOverviewImportantInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-获取重点关注信息
func (c *Client) OverviewInterfaceIGetOverviewImportantInfo(request *OverviewInterfaceIGetOverviewImportantInfoRequest) (response *OverviewInterfaceIGetOverviewImportantInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetOverviewImportantInfoRequest()
    }
    response = NewOverviewInterfaceIGetOverviewImportantInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetPlatformControlInfoRequest() (request *OverviewInterfaceIGetPlatformControlInfoRequest) {
    request = &OverviewInterfaceIGetPlatformControlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetPlatformControlInfo")
    return
}

func NewOverviewInterfaceIGetPlatformControlInfoResponse() (response *OverviewInterfaceIGetPlatformControlInfoResponse) {
    response = &OverviewInterfaceIGetPlatformControlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 今日总览-平台管控
func (c *Client) OverviewInterfaceIGetPlatformControlInfo(request *OverviewInterfaceIGetPlatformControlInfoRequest) (response *OverviewInterfaceIGetPlatformControlInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetPlatformControlInfoRequest()
    }
    response = NewOverviewInterfaceIGetPlatformControlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetRuleStatInfoRequest() (request *OverviewInterfaceIGetRuleStatInfoRequest) {
    request = &OverviewInterfaceIGetRuleStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetRuleStatInfo")
    return
}

func NewOverviewInterfaceIGetRuleStatInfoResponse() (response *OverviewInterfaceIGetRuleStatInfoResponse) {
    response = &OverviewInterfaceIGetRuleStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--获取时间点规则统计
func (c *Client) OverviewInterfaceIGetRuleStatInfo(request *OverviewInterfaceIGetRuleStatInfoRequest) (response *OverviewInterfaceIGetRuleStatInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetRuleStatInfoRequest()
    }
    response = NewOverviewInterfaceIGetRuleStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetRuleTrendStatInfoRequest() (request *OverviewInterfaceIGetRuleTrendStatInfoRequest) {
    request = &OverviewInterfaceIGetRuleTrendStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetRuleTrendStatInfo")
    return
}

func NewOverviewInterfaceIGetRuleTrendStatInfoResponse() (response *OverviewInterfaceIGetRuleTrendStatInfoResponse) {
    response = &OverviewInterfaceIGetRuleTrendStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--获取规则趋势统计
func (c *Client) OverviewInterfaceIGetRuleTrendStatInfo(request *OverviewInterfaceIGetRuleTrendStatInfoRequest) (response *OverviewInterfaceIGetRuleTrendStatInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetRuleTrendStatInfoRequest()
    }
    response = NewOverviewInterfaceIGetRuleTrendStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOverviewInterfaceIGetTOPStatInfoRequest() (request *OverviewInterfaceIGetTOPStatInfoRequest) {
    request = &OverviewInterfaceIGetTOPStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "OverviewInterfaceIGetTOPStatInfo")
    return
}

func NewOverviewInterfaceIGetTOPStatInfoResponse() (response *OverviewInterfaceIGetTOPStatInfoResponse) {
    response = &OverviewInterfaceIGetTOPStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全总览--获取安全检测TOP统计
func (c *Client) OverviewInterfaceIGetTOPStatInfo(request *OverviewInterfaceIGetTOPStatInfoRequest) (response *OverviewInterfaceIGetTOPStatInfoResponse, err error) {
    if request == nil {
        request = NewOverviewInterfaceIGetTOPStatInfoRequest()
    }
    response = NewOverviewInterfaceIGetTOPStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRawlogSearchInterfaceIGetDataFieldsRequest() (request *RawlogSearchInterfaceIGetDataFieldsRequest) {
    request = &RawlogSearchInterfaceIGetDataFieldsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RawlogSearchInterfaceIGetDataFields")
    return
}

func NewRawlogSearchInterfaceIGetDataFieldsResponse() (response *RawlogSearchInterfaceIGetDataFieldsResponse) {
    response = &RawlogSearchInterfaceIGetDataFieldsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 日志审计-数据字段详情
func (c *Client) RawlogSearchInterfaceIGetDataFields(request *RawlogSearchInterfaceIGetDataFieldsRequest) (response *RawlogSearchInterfaceIGetDataFieldsResponse, err error) {
    if request == nil {
        request = NewRawlogSearchInterfaceIGetDataFieldsRequest()
    }
    response = NewRawlogSearchInterfaceIGetDataFieldsResponse()
    err = c.Send(request, response)
    return
}

func NewRawlogSearchInterfaceIQueryRwLogRequest() (request *RawlogSearchInterfaceIQueryRwLogRequest) {
    request = &RawlogSearchInterfaceIQueryRwLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RawlogSearchInterfaceIQueryRwLog")
    return
}

func NewRawlogSearchInterfaceIQueryRwLogResponse() (response *RawlogSearchInterfaceIQueryRwLogResponse) {
    response = &RawlogSearchInterfaceIQueryRwLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 日志审计-列表数据筛选及柱状图统计
func (c *Client) RawlogSearchInterfaceIQueryRwLog(request *RawlogSearchInterfaceIQueryRwLogRequest) (response *RawlogSearchInterfaceIQueryRwLogResponse, err error) {
    if request == nil {
        request = NewRawlogSearchInterfaceIQueryRwLogRequest()
    }
    response = NewRawlogSearchInterfaceIQueryRwLogResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshConfigInterfaceIGetRefreshConfigRequest() (request *RefreshConfigInterfaceIGetRefreshConfigRequest) {
    request = &RefreshConfigInterfaceIGetRefreshConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RefreshConfigInterfaceIGetRefreshConfig")
    return
}

func NewRefreshConfigInterfaceIGetRefreshConfigResponse() (response *RefreshConfigInterfaceIGetRefreshConfigResponse) {
    response = &RefreshConfigInterfaceIGetRefreshConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 刷新配置--获取配置信息
func (c *Client) RefreshConfigInterfaceIGetRefreshConfig(request *RefreshConfigInterfaceIGetRefreshConfigRequest) (response *RefreshConfigInterfaceIGetRefreshConfigResponse, err error) {
    if request == nil {
        request = NewRefreshConfigInterfaceIGetRefreshConfigRequest()
    }
    response = NewRefreshConfigInterfaceIGetRefreshConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshConfigInterfaceIUpdateRefreshConfigRequest() (request *RefreshConfigInterfaceIUpdateRefreshConfigRequest) {
    request = &RefreshConfigInterfaceIUpdateRefreshConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RefreshConfigInterfaceIUpdateRefreshConfig")
    return
}

func NewRefreshConfigInterfaceIUpdateRefreshConfigResponse() (response *RefreshConfigInterfaceIUpdateRefreshConfigResponse) {
    response = &RefreshConfigInterfaceIUpdateRefreshConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//     "page" : 1 // -1 应用到所有页面，1-安全总览，2-安全告警，3-日志审计，4-消息中心
//     "auto_refresh": 0 // 0-未开启自动刷新，1-已开启自动刷新
//     "frequency": 10 // 刷新间隔，单位秒，如3600表示一小时
func (c *Client) RefreshConfigInterfaceIUpdateRefreshConfig(request *RefreshConfigInterfaceIUpdateRefreshConfigRequest) (response *RefreshConfigInterfaceIUpdateRefreshConfigResponse, err error) {
    if request == nil {
        request = NewRefreshConfigInterfaceIUpdateRefreshConfigRequest()
    }
    response = NewRefreshConfigInterfaceIUpdateRefreshConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReportIpFlowInterfaceIGetIpFlowOverviewRequest() (request *ReportIpFlowInterfaceIGetIpFlowOverviewRequest) {
    request = &ReportIpFlowInterfaceIGetIpFlowOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "ReportIpFlowInterfaceIGetIpFlowOverview")
    return
}

func NewReportIpFlowInterfaceIGetIpFlowOverviewResponse() (response *ReportIpFlowInterfaceIGetIpFlowOverviewResponse) {
    response = &ReportIpFlowInterfaceIGetIpFlowOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 报表-流量带宽-出口IP流量-总览
func (c *Client) ReportIpFlowInterfaceIGetIpFlowOverview(request *ReportIpFlowInterfaceIGetIpFlowOverviewRequest) (response *ReportIpFlowInterfaceIGetIpFlowOverviewResponse, err error) {
    if request == nil {
        request = NewReportIpFlowInterfaceIGetIpFlowOverviewRequest()
    }
    response = NewReportIpFlowInterfaceIGetIpFlowOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewReportIpFlowInterfaceIGetIpFlowTableRequest() (request *ReportIpFlowInterfaceIGetIpFlowTableRequest) {
    request = &ReportIpFlowInterfaceIGetIpFlowTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "ReportIpFlowInterfaceIGetIpFlowTable")
    return
}

func NewReportIpFlowInterfaceIGetIpFlowTableResponse() (response *ReportIpFlowInterfaceIGetIpFlowTableResponse) {
    response = &ReportIpFlowInterfaceIGetIpFlowTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 报表-流量带宽-出口IP流量-表格
// 
// CurrentPage: 1
// Date: "2019-11-15"
// ExcuteFlag: 0
// InputCond: ""
// Order: {Field: "", Direc: ""}
// Direc: ""
// Field: ""
// PerPage: 20
func (c *Client) ReportIpFlowInterfaceIGetIpFlowTable(request *ReportIpFlowInterfaceIGetIpFlowTableRequest) (response *ReportIpFlowInterfaceIGetIpFlowTableResponse, err error) {
    if request == nil {
        request = NewReportIpFlowInterfaceIGetIpFlowTableRequest()
    }
    response = NewReportIpFlowInterfaceIGetIpFlowTableResponse()
    err = c.Send(request, response)
    return
}

func NewReportNetFlowInterfaceIGetNetFlowViewDetailRequest() (request *ReportNetFlowInterfaceIGetNetFlowViewDetailRequest) {
    request = &ReportNetFlowInterfaceIGetNetFlowViewDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "ReportNetFlowInterfaceIGetNetFlowViewDetail")
    return
}

func NewReportNetFlowInterfaceIGetNetFlowViewDetailResponse() (response *ReportNetFlowInterfaceIGetNetFlowViewDetailResponse) {
    response = &ReportNetFlowInterfaceIGetNetFlowViewDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 报表-流量带宽-整体视角
func (c *Client) ReportNetFlowInterfaceIGetNetFlowViewDetail(request *ReportNetFlowInterfaceIGetNetFlowViewDetailRequest) (response *ReportNetFlowInterfaceIGetNetFlowViewDetailResponse, err error) {
    if request == nil {
        request = NewReportNetFlowInterfaceIGetNetFlowViewDetailRequest()
    }
    response = NewReportNetFlowInterfaceIGetNetFlowViewDetailResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIBatchChangeWorkStatusRequest() (request *RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest) {
    request = &RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIBatchChangeWorkStatus")
    return
}

func NewRuleConfigManagerInterfaceIBatchChangeWorkStatusResponse() (response *RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse) {
    response = &RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑白名单配置模块--启动暂停规则
func (c *Client) RuleConfigManagerInterfaceIBatchChangeWorkStatus(request *RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest) (response *RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIBatchChangeWorkStatusRequest()
    }
    response = NewRuleConfigManagerInterfaceIBatchChangeWorkStatusResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIBatchDeleteRuleRequest() (request *RuleConfigManagerInterfaceIBatchDeleteRuleRequest) {
    request = &RuleConfigManagerInterfaceIBatchDeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIBatchDeleteRule")
    return
}

func NewRuleConfigManagerInterfaceIBatchDeleteRuleResponse() (response *RuleConfigManagerInterfaceIBatchDeleteRuleResponse) {
    response = &RuleConfigManagerInterfaceIBatchDeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑白名单配置模块--批量删除规则
func (c *Client) RuleConfigManagerInterfaceIBatchDeleteRule(request *RuleConfigManagerInterfaceIBatchDeleteRuleRequest) (response *RuleConfigManagerInterfaceIBatchDeleteRuleResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIBatchDeleteRuleRequest()
    }
    response = NewRuleConfigManagerInterfaceIBatchDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIBatchInsertRuleRequest() (request *RuleConfigManagerInterfaceIBatchInsertRuleRequest) {
    request = &RuleConfigManagerInterfaceIBatchInsertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIBatchInsertRule")
    return
}

func NewRuleConfigManagerInterfaceIBatchInsertRuleResponse() (response *RuleConfigManagerInterfaceIBatchInsertRuleResponse) {
    response = &RuleConfigManagerInterfaceIBatchInsertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  黑白名单配置模块--批量插入新规则
func (c *Client) RuleConfigManagerInterfaceIBatchInsertRule(request *RuleConfigManagerInterfaceIBatchInsertRuleRequest) (response *RuleConfigManagerInterfaceIBatchInsertRuleResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIBatchInsertRuleRequest()
    }
    response = NewRuleConfigManagerInterfaceIBatchInsertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIChangeWorkStatusRequest() (request *RuleConfigManagerInterfaceIChangeWorkStatusRequest) {
    request = &RuleConfigManagerInterfaceIChangeWorkStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIChangeWorkStatus")
    return
}

func NewRuleConfigManagerInterfaceIChangeWorkStatusResponse() (response *RuleConfigManagerInterfaceIChangeWorkStatusResponse) {
    response = &RuleConfigManagerInterfaceIChangeWorkStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑白名单配置模块--启动暂停规则
func (c *Client) RuleConfigManagerInterfaceIChangeWorkStatus(request *RuleConfigManagerInterfaceIChangeWorkStatusRequest) (response *RuleConfigManagerInterfaceIChangeWorkStatusResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIChangeWorkStatusRequest()
    }
    response = NewRuleConfigManagerInterfaceIChangeWorkStatusResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIDeleteRuleRequest() (request *RuleConfigManagerInterfaceIDeleteRuleRequest) {
    request = &RuleConfigManagerInterfaceIDeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIDeleteRule")
    return
}

func NewRuleConfigManagerInterfaceIDeleteRuleResponse() (response *RuleConfigManagerInterfaceIDeleteRuleResponse) {
    response = &RuleConfigManagerInterfaceIDeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  黑白名单配置模块--删除规则
func (c *Client) RuleConfigManagerInterfaceIDeleteRule(request *RuleConfigManagerInterfaceIDeleteRuleRequest) (response *RuleConfigManagerInterfaceIDeleteRuleResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIDeleteRuleRequest()
    }
    response = NewRuleConfigManagerInterfaceIDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIGetFilterItemsListInfoRequest() (request *RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest) {
    request = &RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIGetFilterItemsListInfo")
    return
}

func NewRuleConfigManagerInterfaceIGetFilterItemsListInfoResponse() (response *RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse) {
    response = &RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑白名单配置模块--筛选条件统计
func (c *Client) RuleConfigManagerInterfaceIGetFilterItemsListInfo(request *RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest) (response *RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIGetFilterItemsListInfoRequest()
    }
    response = NewRuleConfigManagerInterfaceIGetFilterItemsListInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIInsertRuleRequest() (request *RuleConfigManagerInterfaceIInsertRuleRequest) {
    request = &RuleConfigManagerInterfaceIInsertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIInsertRule")
    return
}

func NewRuleConfigManagerInterfaceIInsertRuleResponse() (response *RuleConfigManagerInterfaceIInsertRuleResponse) {
    response = &RuleConfigManagerInterfaceIInsertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑白名单配置模块--插入新规则
func (c *Client) RuleConfigManagerInterfaceIInsertRule(request *RuleConfigManagerInterfaceIInsertRuleRequest) (response *RuleConfigManagerInterfaceIInsertRuleResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIInsertRuleRequest()
    }
    response = NewRuleConfigManagerInterfaceIInsertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIListRuleRequest() (request *RuleConfigManagerInterfaceIListRuleRequest) {
    request = &RuleConfigManagerInterfaceIListRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIListRule")
    return
}

func NewRuleConfigManagerInterfaceIListRuleResponse() (response *RuleConfigManagerInterfaceIListRuleResponse) {
    response = &RuleConfigManagerInterfaceIListRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑白名单配置模块--列表
func (c *Client) RuleConfigManagerInterfaceIListRule(request *RuleConfigManagerInterfaceIListRuleRequest) (response *RuleConfigManagerInterfaceIListRuleResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIListRuleRequest()
    }
    response = NewRuleConfigManagerInterfaceIListRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRuleConfigManagerInterfaceIUpdateRuleRequest() (request *RuleConfigManagerInterfaceIUpdateRuleRequest) {
    request = &RuleConfigManagerInterfaceIUpdateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "RuleConfigManagerInterfaceIUpdateRule")
    return
}

func NewRuleConfigManagerInterfaceIUpdateRuleResponse() (response *RuleConfigManagerInterfaceIUpdateRuleResponse) {
    response = &RuleConfigManagerInterfaceIUpdateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑白名单配置模块--修改规则
func (c *Client) RuleConfigManagerInterfaceIUpdateRule(request *RuleConfigManagerInterfaceIUpdateRuleRequest) (response *RuleConfigManagerInterfaceIUpdateRuleResponse, err error) {
    if request == nil {
        request = NewRuleConfigManagerInterfaceIUpdateRuleRequest()
    }
    response = NewRuleConfigManagerInterfaceIUpdateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewSystemRuleInterfaceIExportDataRequest() (request *SystemRuleInterfaceIExportDataRequest) {
    request = &SystemRuleInterfaceIExportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "SystemRuleInterfaceIExportData")
    return
}

func NewSystemRuleInterfaceIExportDataResponse() (response *SystemRuleInterfaceIExportDataResponse) {
    response = &SystemRuleInterfaceIExportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 系统调用规则--全量导出
func (c *Client) SystemRuleInterfaceIExportData(request *SystemRuleInterfaceIExportDataRequest) (response *SystemRuleInterfaceIExportDataResponse, err error) {
    if request == nil {
        request = NewSystemRuleInterfaceIExportDataRequest()
    }
    response = NewSystemRuleInterfaceIExportDataResponse()
    err = c.Send(request, response)
    return
}

func NewSystemRuleInterfaceIGetRuleStatInfoRequest() (request *SystemRuleInterfaceIGetRuleStatInfoRequest) {
    request = &SystemRuleInterfaceIGetRuleStatInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "SystemRuleInterfaceIGetRuleStatInfo")
    return
}

func NewSystemRuleInterfaceIGetRuleStatInfoResponse() (response *SystemRuleInterfaceIGetRuleStatInfoResponse) {
    response = &SystemRuleInterfaceIGetRuleStatInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddTime: []
// BeginTime: ""
// CurrentPage: 1
// EndTime: ""
// ExcuteFlag: 0
// HitCount: ""
// InputCond: ""
// Order: {field: "", direc: ""}
// direc: ""
// field: ""
// PerPage: 20
// SearchContent: ""
// Source: ""
// Status: ""
// ValidDuration: ""
func (c *Client) SystemRuleInterfaceIGetRuleStatInfo(request *SystemRuleInterfaceIGetRuleStatInfoRequest) (response *SystemRuleInterfaceIGetRuleStatInfoResponse, err error) {
    if request == nil {
        request = NewSystemRuleInterfaceIGetRuleStatInfoRequest()
    }
    response = NewSystemRuleInterfaceIGetRuleStatInfoResponse()
    err = c.Send(request, response)
    return
}

func NewSystemRuleInterfaceIListRuleRequest() (request *SystemRuleInterfaceIListRuleRequest) {
    request = &SystemRuleInterfaceIListRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "SystemRuleInterfaceIListRule")
    return
}

func NewSystemRuleInterfaceIListRuleResponse() (response *SystemRuleInterfaceIListRuleResponse) {
    response = &SystemRuleInterfaceIListRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddTime: []
// BeginTime: ""
// CurrentPage: 1
// EndTime: ""
// ExcuteFlag: 1
// HitCount: ""
// InputCond: ""
// Order: {Field: "", Direc: ""}
// Direc: ""
// Field: ""
// PerPage: 20
// Region: "chongqing"
// SearchContent: ""
// Source: ""
// Status: ""
// ValidDuration: ""
func (c *Client) SystemRuleInterfaceIListRule(request *SystemRuleInterfaceIListRuleRequest) (response *SystemRuleInterfaceIListRuleResponse, err error) {
    if request == nil {
        request = NewSystemRuleInterfaceIListRuleRequest()
    }
    response = NewSystemRuleInterfaceIListRuleResponse()
    err = c.Send(request, response)
    return
}

func NewSystemRuleInterfaceIStopRulesRequest() (request *SystemRuleInterfaceIStopRulesRequest) {
    request = &SystemRuleInterfaceIStopRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "SystemRuleInterfaceIStopRules")
    return
}

func NewSystemRuleInterfaceIStopRulesResponse() (response *SystemRuleInterfaceIStopRulesResponse) {
    response = &SystemRuleInterfaceIStopRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Request
// {
// "rule": [{
// "id": "id",
// "source_id": "source_id"
// }]
// }
func (c *Client) SystemRuleInterfaceIStopRules(request *SystemRuleInterfaceIStopRulesRequest) (response *SystemRuleInterfaceIStopRulesResponse, err error) {
    if request == nil {
        request = NewSystemRuleInterfaceIStopRulesRequest()
    }
    response = NewSystemRuleInterfaceIStopRulesResponse()
    err = c.Send(request, response)
    return
}

func NewTceManagerInterfaceIGetFlowStatusRequest() (request *TceManagerInterfaceIGetFlowStatusRequest) {
    request = &TceManagerInterfaceIGetFlowStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "TceManagerInterfaceIGetFlowStatus")
    return
}

func NewTceManagerInterfaceIGetFlowStatusResponse() (response *TceManagerInterfaceIGetFlowStatusResponse) {
    response = &TceManagerInterfaceIGetFlowStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警日志模块--网络状态
func (c *Client) TceManagerInterfaceIGetFlowStatus(request *TceManagerInterfaceIGetFlowStatusRequest) (response *TceManagerInterfaceIGetFlowStatusResponse, err error) {
    if request == nil {
        request = NewTceManagerInterfaceIGetFlowStatusRequest()
    }
    response = NewTceManagerInterfaceIGetFlowStatusResponse()
    err = c.Send(request, response)
    return
}

func NewTceManagerInterfaceIGetSystemRawLogListRequest() (request *TceManagerInterfaceIGetSystemRawLogListRequest) {
    request = &TceManagerInterfaceIGetSystemRawLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "TceManagerInterfaceIGetSystemRawLogList")
    return
}

func NewTceManagerInterfaceIGetSystemRawLogListResponse() (response *TceManagerInterfaceIGetSystemRawLogListResponse) {
    response = &TceManagerInterfaceIGetSystemRawLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运维告警日志模块--列表查询接口
func (c *Client) TceManagerInterfaceIGetSystemRawLogList(request *TceManagerInterfaceIGetSystemRawLogListRequest) (response *TceManagerInterfaceIGetSystemRawLogListResponse, err error) {
    if request == nil {
        request = NewTceManagerInterfaceIGetSystemRawLogListRequest()
    }
    response = NewTceManagerInterfaceIGetSystemRawLogListResponse()
    err = c.Send(request, response)
    return
}

func NewTceManagerInterfaceIGetSystemRawLogStatisticInfoRequest() (request *TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest) {
    request = &TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "TceManagerInterfaceIGetSystemRawLogStatisticInfo")
    return
}

func NewTceManagerInterfaceIGetSystemRawLogStatisticInfoResponse() (response *TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse) {
    response = &TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运维告警日志模块--筛选条件统计
func (c *Client) TceManagerInterfaceIGetSystemRawLogStatisticInfo(request *TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest) (response *TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse, err error) {
    if request == nil {
        request = NewTceManagerInterfaceIGetSystemRawLogStatisticInfoRequest()
    }
    response = NewTceManagerInterfaceIGetSystemRawLogStatisticInfoResponse()
    err = c.Send(request, response)
    return
}

func NewTceManagerInterfaceIGetThreatRawLogListRequest() (request *TceManagerInterfaceIGetThreatRawLogListRequest) {
    request = &TceManagerInterfaceIGetThreatRawLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "TceManagerInterfaceIGetThreatRawLogList")
    return
}

func NewTceManagerInterfaceIGetThreatRawLogListResponse() (response *TceManagerInterfaceIGetThreatRawLogListResponse) {
    response = &TceManagerInterfaceIGetThreatRawLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警日志模块--列表查询接口
func (c *Client) TceManagerInterfaceIGetThreatRawLogList(request *TceManagerInterfaceIGetThreatRawLogListRequest) (response *TceManagerInterfaceIGetThreatRawLogListResponse, err error) {
    if request == nil {
        request = NewTceManagerInterfaceIGetThreatRawLogListRequest()
    }
    response = NewTceManagerInterfaceIGetThreatRawLogListResponse()
    err = c.Send(request, response)
    return
}

func NewTceManagerInterfaceIGetThreatRawLogStatisticInfoRequest() (request *TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest) {
    request = &TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "TceManagerInterfaceIGetThreatRawLogStatisticInfo")
    return
}

func NewTceManagerInterfaceIGetThreatRawLogStatisticInfoResponse() (response *TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse) {
    response = &TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警日志模块--获取告警详情接口
func (c *Client) TceManagerInterfaceIGetThreatRawLogStatisticInfo(request *TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest) (response *TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse, err error) {
    if request == nil {
        request = NewTceManagerInterfaceIGetThreatRawLogStatisticInfoRequest()
    }
    response = NewTceManagerInterfaceIGetThreatRawLogStatisticInfoResponse()
    err = c.Send(request, response)
    return
}

func NewTceManagerInterfaceIStopFlowRequest() (request *TceManagerInterfaceIStopFlowRequest) {
    request = &TceManagerInterfaceIStopFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("psg", APIVersion, "TceManagerInterfaceIStopFlow")
    return
}

func NewTceManagerInterfaceIStopFlowResponse() (response *TceManagerInterfaceIStopFlowResponse) {
    response = &TceManagerInterfaceIStopFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警日志模块--一键断网
func (c *Client) TceManagerInterfaceIStopFlow(request *TceManagerInterfaceIStopFlowRequest) (response *TceManagerInterfaceIStopFlowResponse, err error) {
    if request == nil {
        request = NewTceManagerInterfaceIStopFlowRequest()
    }
    response = NewTceManagerInterfaceIStopFlowResponse()
    err = c.Send(request, response)
    return
}
