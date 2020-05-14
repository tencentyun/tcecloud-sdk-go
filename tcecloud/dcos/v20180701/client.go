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

package v20180701

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-01"

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


func NewAddIdcLineRequest() (request *AddIdcLineRequest) {
    request = &AddIdcLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddIdcLine")
    return
}

func NewAddIdcLineResponse() (response *AddIdcLineResponse) {
    response = &AddIdcLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加idc专线信息
func (c *Client) AddIdcLine(request *AddIdcLineRequest) (response *AddIdcLineResponse, err error) {
    if request == nil {
        request = NewAddIdcLineRequest()
    }
    response = NewAddIdcLineResponse()
    err = c.Send(request, response)
    return
}

func NewAddIdcLineExportRequest() (request *AddIdcLineExportRequest) {
    request = &AddIdcLineExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddIdcLineExport")
    return
}

func NewAddIdcLineExportResponse() (response *AddIdcLineExportResponse) {
    response = &AddIdcLineExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加idc专线出口信息
func (c *Client) AddIdcLineExport(request *AddIdcLineExportRequest) (response *AddIdcLineExportResponse, err error) {
    if request == nil {
        request = NewAddIdcLineExportRequest()
    }
    response = NewAddIdcLineExportResponse()
    err = c.Send(request, response)
    return
}

func NewAddNetDeviceRequest() (request *AddNetDeviceRequest) {
    request = &AddNetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddNetDevice")
    return
}

func NewAddNetDeviceResponse() (response *AddNetDeviceResponse) {
    response = &AddNetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络设备配置信息新增
func (c *Client) AddNetDevice(request *AddNetDeviceRequest) (response *AddNetDeviceResponse, err error) {
    if request == nil {
        request = NewAddNetDeviceRequest()
    }
    response = NewAddNetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewAddNetPortRequest() (request *AddNetPortRequest) {
    request = &AddNetPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddNetPort")
    return
}

func NewAddNetPortResponse() (response *AddNetPortResponse) {
    response = &AddNetPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络设备端口信息新增
func (c *Client) AddNetPort(request *AddNetPortRequest) (response *AddNetPortResponse, err error) {
    if request == nil {
        request = NewAddNetPortRequest()
    }
    response = NewAddNetPortResponse()
    err = c.Send(request, response)
    return
}

func NewAddNetSegmentRequest() (request *AddNetSegmentRequest) {
    request = &AddNetSegmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddNetSegment")
    return
}

func NewAddNetSegmentResponse() (response *AddNetSegmentResponse) {
    response = &AddNetSegmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网段信息新增
func (c *Client) AddNetSegment(request *AddNetSegmentRequest) (response *AddNetSegmentResponse, err error) {
    if request == nil {
        request = NewAddNetSegmentRequest()
    }
    response = NewAddNetSegmentResponse()
    err = c.Send(request, response)
    return
}

func NewAddRackDescRequest() (request *AddRackDescRequest) {
    request = &AddRackDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddRackDesc")
    return
}

func NewAddRackDescResponse() (response *AddRackDescResponse) {
    response = &AddRackDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加idc机架机位配置信息
func (c *Client) AddRackDesc(request *AddRackDescRequest) (response *AddRackDescResponse, err error) {
    if request == nil {
        request = NewAddRackDescRequest()
    }
    response = NewAddRackDescResponse()
    err = c.Send(request, response)
    return
}

func NewAddRackInfoRequest() (request *AddRackInfoRequest) {
    request = &AddRackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddRackInfo")
    return
}

func NewAddRackInfoResponse() (response *AddRackInfoResponse) {
    response = &AddRackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加机架机位配置信息
func (c *Client) AddRackInfo(request *AddRackInfoRequest) (response *AddRackInfoResponse, err error) {
    if request == nil {
        request = NewAddRackInfoRequest()
    }
    response = NewAddRackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAddServerRequest() (request *AddServerRequest) {
    request = &AddServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddServer")
    return
}

func NewAddServerResponse() (response *AddServerResponse) {
    response = &AddServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增实体服务器配置信息
func (c *Client) AddServer(request *AddServerRequest) (response *AddServerResponse, err error) {
    if request == nil {
        request = NewAddServerRequest()
    }
    response = NewAddServerResponse()
    err = c.Send(request, response)
    return
}

func NewAddSnmpTemplateRequest() (request *AddSnmpTemplateRequest) {
    request = &AddSnmpTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AddSnmpTemplate")
    return
}

func NewAddSnmpTemplateResponse() (response *AddSnmpTemplateResponse) {
    response = &AddSnmpTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp导入模版
func (c *Client) AddSnmpTemplate(request *AddSnmpTemplateRequest) (response *AddSnmpTemplateResponse, err error) {
    if request == nil {
        request = NewAddSnmpTemplateRequest()
    }
    response = NewAddSnmpTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmReportCharRequest() (request *AlarmReportCharRequest) {
    request = &AlarmReportCharRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AlarmReportChar")
    return
}

func NewAlarmReportCharResponse() (response *AlarmReportCharResponse) {
    response = &AlarmReportCharResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警协议上报 – 字符型
func (c *Client) AlarmReportChar(request *AlarmReportCharRequest) (response *AlarmReportCharResponse, err error) {
    if request == nil {
        request = NewAlarmReportCharRequest()
    }
    response = NewAlarmReportCharResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmReportCharsRequest() (request *AlarmReportCharsRequest) {
    request = &AlarmReportCharsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AlarmReportChars")
    return
}

func NewAlarmReportCharsResponse() (response *AlarmReportCharsResponse) {
    response = &AlarmReportCharsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量告警协议上报 – 字符串型
func (c *Client) AlarmReportChars(request *AlarmReportCharsRequest) (response *AlarmReportCharsResponse, err error) {
    if request == nil {
        request = NewAlarmReportCharsRequest()
    }
    response = NewAlarmReportCharsResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmReportNumRequest() (request *AlarmReportNumRequest) {
    request = &AlarmReportNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AlarmReportNum")
    return
}

func NewAlarmReportNumResponse() (response *AlarmReportNumResponse) {
    response = &AlarmReportNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警协议上报 – 数值型
func (c *Client) AlarmReportNum(request *AlarmReportNumRequest) (response *AlarmReportNumResponse, err error) {
    if request == nil {
        request = NewAlarmReportNumRequest()
    }
    response = NewAlarmReportNumResponse()
    err = c.Send(request, response)
    return
}

func NewAlarmReportNumsRequest() (request *AlarmReportNumsRequest) {
    request = &AlarmReportNumsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AlarmReportNums")
    return
}

func NewAlarmReportNumsResponse() (response *AlarmReportNumsResponse) {
    response = &AlarmReportNumsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量告警协议上报 – 数值型
func (c *Client) AlarmReportNums(request *AlarmReportNumsRequest) (response *AlarmReportNumsResponse, err error) {
    if request == nil {
        request = NewAlarmReportNumsRequest()
    }
    response = NewAlarmReportNumsResponse()
    err = c.Send(request, response)
    return
}

func NewAllocateServerIpRequest() (request *AllocateServerIpRequest) {
    request = &AllocateServerIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "AllocateServerIp")
    return
}

func NewAllocateServerIpResponse() (response *AllocateServerIpResponse) {
    response = &AllocateServerIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务器内外网IP申请
func (c *Client) AllocateServerIp(request *AllocateServerIpRequest) (response *AllocateServerIpResponse, err error) {
    if request == nil {
        request = NewAllocateServerIpRequest()
    }
    response = NewAllocateServerIpResponse()
    err = c.Send(request, response)
    return
}

func NewBindSnmpTemplateRequest() (request *BindSnmpTemplateRequest) {
    request = &BindSnmpTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "BindSnmpTemplate")
    return
}

func NewBindSnmpTemplateResponse() (response *BindSnmpTemplateResponse) {
    response = &BindSnmpTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp绑定模版
func (c *Client) BindSnmpTemplate(request *BindSnmpTemplateRequest) (response *BindSnmpTemplateResponse, err error) {
    if request == nil {
        request = NewBindSnmpTemplateRequest()
    }
    response = NewBindSnmpTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmStrategyRequest() (request *CreateAlarmStrategyRequest) {
    request = &CreateAlarmStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "CreateAlarmStrategy")
    return
}

func NewCreateAlarmStrategyResponse() (response *CreateAlarmStrategyResponse) {
    response = &CreateAlarmStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建告警策略
func (c *Client) CreateAlarmStrategy(request *CreateAlarmStrategyRequest) (response *CreateAlarmStrategyResponse, err error) {
    if request == nil {
        request = NewCreateAlarmStrategyRequest()
    }
    response = NewCreateAlarmStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDelIdcLineRequest() (request *DelIdcLineRequest) {
    request = &DelIdcLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DelIdcLine")
    return
}

func NewDelIdcLineResponse() (response *DelIdcLineResponse) {
    response = &DelIdcLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除idc专线信息(t_cmdb_special_idc_line_info)
func (c *Client) DelIdcLine(request *DelIdcLineRequest) (response *DelIdcLineResponse, err error) {
    if request == nil {
        request = NewDelIdcLineRequest()
    }
    response = NewDelIdcLineResponse()
    err = c.Send(request, response)
    return
}

func NewDelIdcLineExportRequest() (request *DelIdcLineExportRequest) {
    request = &DelIdcLineExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DelIdcLineExport")
    return
}

func NewDelIdcLineExportResponse() (response *DelIdcLineExportResponse) {
    response = &DelIdcLineExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除idc专线出口信息(t_cmdb_idc_export_line_info)
func (c *Client) DelIdcLineExport(request *DelIdcLineExportRequest) (response *DelIdcLineExportResponse, err error) {
    if request == nil {
        request = NewDelIdcLineExportRequest()
    }
    response = NewDelIdcLineExportResponse()
    err = c.Send(request, response)
    return
}

func NewDelNetDeviceRequest() (request *DelNetDeviceRequest) {
    request = &DelNetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DelNetDevice")
    return
}

func NewDelNetDeviceResponse() (response *DelNetDeviceResponse) {
    response = &DelNetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络设备配置信息删除(t_cmdb_network_basic_info)
func (c *Client) DelNetDevice(request *DelNetDeviceRequest) (response *DelNetDeviceResponse, err error) {
    if request == nil {
        request = NewDelNetDeviceRequest()
    }
    response = NewDelNetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDelNetPortRequest() (request *DelNetPortRequest) {
    request = &DelNetPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DelNetPort")
    return
}

func NewDelNetPortResponse() (response *DelNetPortResponse) {
    response = &DelNetPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络设备端口信息删除(t_cmdb_network_port_info)
func (c *Client) DelNetPort(request *DelNetPortRequest) (response *DelNetPortResponse, err error) {
    if request == nil {
        request = NewDelNetPortRequest()
    }
    response = NewDelNetPortResponse()
    err = c.Send(request, response)
    return
}

func NewDelRackDescRequest() (request *DelRackDescRequest) {
    request = &DelRackDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DelRackDesc")
    return
}

func NewDelRackDescResponse() (response *DelRackDescResponse) {
    response = &DelRackDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除idc机架机位配置信息(t_cmdb_idc_rack_desc)
func (c *Client) DelRackDesc(request *DelRackDescRequest) (response *DelRackDescResponse, err error) {
    if request == nil {
        request = NewDelRackDescRequest()
    }
    response = NewDelRackDescResponse()
    err = c.Send(request, response)
    return
}

func NewDelRackInfoRequest() (request *DelRackInfoRequest) {
    request = &DelRackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DelRackInfo")
    return
}

func NewDelRackInfoResponse() (response *DelRackInfoResponse) {
    response = &DelRackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除机架机位配置信息(t_cmdb_rack_info)
func (c *Client) DelRackInfo(request *DelRackInfoRequest) (response *DelRackInfoResponse, err error) {
    if request == nil {
        request = NewDelRackInfoRequest()
    }
    response = NewDelRackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmStrategyRequest() (request *DeleteAlarmStrategyRequest) {
    request = &DeleteAlarmStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DeleteAlarmStrategy")
    return
}

func NewDeleteAlarmStrategyResponse() (response *DeleteAlarmStrategyResponse) {
    response = &DeleteAlarmStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除告警策略
func (c *Client) DeleteAlarmStrategy(request *DeleteAlarmStrategyRequest) (response *DeleteAlarmStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmStrategyRequest()
    }
    response = NewDeleteAlarmStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmStrategyMd5Request() (request *DeleteAlarmStrategyMd5Request) {
    request = &DeleteAlarmStrategyMd5Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DeleteAlarmStrategyMd5")
    return
}

func NewDeleteAlarmStrategyMd5Response() (response *DeleteAlarmStrategyMd5Response) {
    response = &DeleteAlarmStrategyMd5Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除告警策略
func (c *Client) DeleteAlarmStrategyMd5(request *DeleteAlarmStrategyMd5Request) (response *DeleteAlarmStrategyMd5Response, err error) {
    if request == nil {
        request = NewDeleteAlarmStrategyMd5Request()
    }
    response = NewDeleteAlarmStrategyMd5Response()
    err = c.Send(request, response)
    return
}

func NewDeleteServerRequest() (request *DeleteServerRequest) {
    request = &DeleteServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "DeleteServer")
    return
}

func NewDeleteServerResponse() (response *DeleteServerResponse) {
    response = &DeleteServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除实体服务器配置信息
func (c *Client) DeleteServer(request *DeleteServerRequest) (response *DeleteServerResponse, err error) {
    if request == nil {
        request = NewDeleteServerRequest()
    }
    response = NewDeleteServerResponse()
    err = c.Send(request, response)
    return
}

func NewFastReinstallServerRequest() (request *FastReinstallServerRequest) {
    request = &FastReinstallServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "FastReinstallServer")
    return
}

func NewFastReinstallServerResponse() (response *FastReinstallServerResponse) {
    response = &FastReinstallServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过接口对服务器进行快速重装
func (c *Client) FastReinstallServer(request *FastReinstallServerRequest) (response *FastReinstallServerResponse, err error) {
    if request == nil {
        request = NewFastReinstallServerRequest()
    }
    response = NewFastReinstallServerResponse()
    err = c.Send(request, response)
    return
}

func NewGetAggregationXflowDataRequest() (request *GetAggregationXflowDataRequest) {
    request = &GetAggregationXflowDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "GetAggregationXflowData")
    return
}

func NewGetAggregationXflowDataResponse() (response *GetAggregationXflowDataResponse) {
    response = &GetAggregationXflowDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询xflow汇总专线数据
func (c *Client) GetAggregationXflowData(request *GetAggregationXflowDataRequest) (response *GetAggregationXflowDataResponse, err error) {
    if request == nil {
        request = NewGetAggregationXflowDataRequest()
    }
    response = NewGetAggregationXflowDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetOriginalXflowDataRequest() (request *GetOriginalXflowDataRequest) {
    request = &GetOriginalXflowDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "GetOriginalXflowData")
    return
}

func NewGetOriginalXflowDataResponse() (response *GetOriginalXflowDataResponse) {
    response = &GetOriginalXflowDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询xflow原始专线数据
func (c *Client) GetOriginalXflowData(request *GetOriginalXflowDataRequest) (response *GetOriginalXflowDataResponse, err error) {
    if request == nil {
        request = NewGetOriginalXflowDataRequest()
    }
    response = NewGetOriginalXflowDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetPeriodAggregationXflowDataRequest() (request *GetPeriodAggregationXflowDataRequest) {
    request = &GetPeriodAggregationXflowDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "GetPeriodAggregationXflowData")
    return
}

func NewGetPeriodAggregationXflowDataResponse() (response *GetPeriodAggregationXflowDataResponse) {
    response = &GetPeriodAggregationXflowDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Xflow查询按周期（五分钟）汇总专线数据
func (c *Client) GetPeriodAggregationXflowData(request *GetPeriodAggregationXflowDataRequest) (response *GetPeriodAggregationXflowDataResponse, err error) {
    if request == nil {
        request = NewGetPeriodAggregationXflowDataRequest()
    }
    response = NewGetPeriodAggregationXflowDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopnXflowDataRequest() (request *GetTopnXflowDataRequest) {
    request = &GetTopnXflowDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "GetTopnXflowData")
    return
}

func NewGetTopnXflowDataResponse() (response *GetTopnXflowDataResponse) {
    response = &GetTopnXflowDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// XFLOW查询TopN汇总专线数据
func (c *Client) GetTopnXflowData(request *GetTopnXflowDataRequest) (response *GetTopnXflowDataResponse, err error) {
    if request == nil {
        request = NewGetTopnXflowDataRequest()
    }
    response = NewGetTopnXflowDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetXflowInfoRequest() (request *GetXflowInfoRequest) {
    request = &GetXflowInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "GetXflowInfo")
    return
}

func NewGetXflowInfoResponse() (response *GetXflowInfoResponse) {
    response = &GetXflowInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// XFLOW查询专线信息
func (c *Client) GetXflowInfo(request *GetXflowInfoRequest) (response *GetXflowInfoResponse, err error) {
    if request == nil {
        request = NewGetXflowInfoRequest()
    }
    response = NewGetXflowInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmStrategyRequest() (request *ModifyAlarmStrategyRequest) {
    request = &ModifyAlarmStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyAlarmStrategy")
    return
}

func NewModifyAlarmStrategyResponse() (response *ModifyAlarmStrategyResponse) {
    response = &ModifyAlarmStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新告警策略x
func (c *Client) ModifyAlarmStrategy(request *ModifyAlarmStrategyRequest) (response *ModifyAlarmStrategyResponse, err error) {
    if request == nil {
        request = NewModifyAlarmStrategyRequest()
    }
    response = NewModifyAlarmStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmStrategyMd5Request() (request *ModifyAlarmStrategyMd5Request) {
    request = &ModifyAlarmStrategyMd5Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyAlarmStrategyMd5")
    return
}

func NewModifyAlarmStrategyMd5Response() (response *ModifyAlarmStrategyMd5Response) {
    response = &ModifyAlarmStrategyMd5Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新告警策略
func (c *Client) ModifyAlarmStrategyMd5(request *ModifyAlarmStrategyMd5Request) (response *ModifyAlarmStrategyMd5Response, err error) {
    if request == nil {
        request = NewModifyAlarmStrategyMd5Request()
    }
    response = NewModifyAlarmStrategyMd5Response()
    err = c.Send(request, response)
    return
}

func NewModifyIdcLineRequest() (request *ModifyIdcLineRequest) {
    request = &ModifyIdcLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyIdcLine")
    return
}

func NewModifyIdcLineResponse() (response *ModifyIdcLineResponse) {
    response = &ModifyIdcLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新idc专线信息
func (c *Client) ModifyIdcLine(request *ModifyIdcLineRequest) (response *ModifyIdcLineResponse, err error) {
    if request == nil {
        request = NewModifyIdcLineRequest()
    }
    response = NewModifyIdcLineResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIdcLineExportRequest() (request *ModifyIdcLineExportRequest) {
    request = &ModifyIdcLineExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyIdcLineExport")
    return
}

func NewModifyIdcLineExportResponse() (response *ModifyIdcLineExportResponse) {
    response = &ModifyIdcLineExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新idc专线出口信息
func (c *Client) ModifyIdcLineExport(request *ModifyIdcLineExportRequest) (response *ModifyIdcLineExportResponse, err error) {
    if request == nil {
        request = NewModifyIdcLineExportRequest()
    }
    response = NewModifyIdcLineExportResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetDeviceRequest() (request *ModifyNetDeviceRequest) {
    request = &ModifyNetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyNetDevice")
    return
}

func NewModifyNetDeviceResponse() (response *ModifyNetDeviceResponse) {
    response = &ModifyNetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络设备配置信息修改(t_cmdb_network_basic_info)
func (c *Client) ModifyNetDevice(request *ModifyNetDeviceRequest) (response *ModifyNetDeviceResponse, err error) {
    if request == nil {
        request = NewModifyNetDeviceRequest()
    }
    response = NewModifyNetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetPortRequest() (request *ModifyNetPortRequest) {
    request = &ModifyNetPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyNetPort")
    return
}

func NewModifyNetPortResponse() (response *ModifyNetPortResponse) {
    response = &ModifyNetPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络设备端口信息修改(t_cmdb_network_port_info)
func (c *Client) ModifyNetPort(request *ModifyNetPortRequest) (response *ModifyNetPortResponse, err error) {
    if request == nil {
        request = NewModifyNetPortRequest()
    }
    response = NewModifyNetPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRackDescRequest() (request *ModifyRackDescRequest) {
    request = &ModifyRackDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyRackDesc")
    return
}

func NewModifyRackDescResponse() (response *ModifyRackDescResponse) {
    response = &ModifyRackDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改idc机架机位配置信息(t_cmdb_idc_rack_desc)
func (c *Client) ModifyRackDesc(request *ModifyRackDescRequest) (response *ModifyRackDescResponse, err error) {
    if request == nil {
        request = NewModifyRackDescRequest()
    }
    response = NewModifyRackDescResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRackInfoRequest() (request *ModifyRackInfoRequest) {
    request = &ModifyRackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyRackInfo")
    return
}

func NewModifyRackInfoResponse() (response *ModifyRackInfoResponse) {
    response = &ModifyRackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改机架机位配置信息(t_cmdb_rack_info)
func (c *Client) ModifyRackInfo(request *ModifyRackInfoRequest) (response *ModifyRackInfoResponse, err error) {
    if request == nil {
        request = NewModifyRackInfoRequest()
    }
    response = NewModifyRackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServerRequest() (request *ModifyServerRequest) {
    request = &ModifyServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ModifyServer")
    return
}

func NewModifyServerResponse() (response *ModifyServerResponse) {
    response = &ModifyServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改实体服务器配置信息
func (c *Client) ModifyServer(request *ModifyServerRequest) (response *ModifyServerResponse, err error) {
    if request == nil {
        request = NewModifyServerRequest()
    }
    response = NewModifyServerResponse()
    err = c.Send(request, response)
    return
}

func NewOperateServerRequest() (request *OperateServerRequest) {
    request = &OperateServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "OperateServer")
    return
}

func NewOperateServerResponse() (response *OperateServerResponse) {
    response = &OperateServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过接口对服务器进行重启、开机、关机、初始化、一致性检查、开机pxe操作
func (c *Client) OperateServer(request *OperateServerRequest) (response *OperateServerResponse, err error) {
    if request == nil {
        request = NewOperateServerRequest()
    }
    response = NewOperateServerResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAlarmRequest() (request *QueryAlarmRequest) {
    request = &QueryAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryAlarm")
    return
}

func NewQueryAlarmResponse() (response *QueryAlarmResponse) {
    response = &QueryAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 告警查询
func (c *Client) QueryAlarm(request *QueryAlarmRequest) (response *QueryAlarmResponse, err error) {
    if request == nil {
        request = NewQueryAlarmRequest()
    }
    response = NewQueryAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAlarmStrategysRequest() (request *QueryAlarmStrategysRequest) {
    request = &QueryAlarmStrategysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryAlarmStrategys")
    return
}

func NewQueryAlarmStrategysResponse() (response *QueryAlarmStrategysResponse) {
    response = &QueryAlarmStrategysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询告警策略
func (c *Client) QueryAlarmStrategys(request *QueryAlarmStrategysRequest) (response *QueryAlarmStrategysResponse, err error) {
    if request == nil {
        request = NewQueryAlarmStrategysRequest()
    }
    response = NewQueryAlarmStrategysResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDeployLogRequest() (request *QueryDeployLogRequest) {
    request = &QueryDeployLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryDeployLog")
    return
}

func NewQueryDeployLogResponse() (response *QueryDeployLogResponse) {
    response = &QueryDeployLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询部署日志
func (c *Client) QueryDeployLog(request *QueryDeployLogRequest) (response *QueryDeployLogResponse, err error) {
    if request == nil {
        request = NewQueryDeployLogRequest()
    }
    response = NewQueryDeployLogResponse()
    err = c.Send(request, response)
    return
}

func NewQueryIdcLineRequest() (request *QueryIdcLineRequest) {
    request = &QueryIdcLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryIdcLine")
    return
}

func NewQueryIdcLineResponse() (response *QueryIdcLineResponse) {
    response = &QueryIdcLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询idc专线信息
func (c *Client) QueryIdcLine(request *QueryIdcLineRequest) (response *QueryIdcLineResponse, err error) {
    if request == nil {
        request = NewQueryIdcLineRequest()
    }
    response = NewQueryIdcLineResponse()
    err = c.Send(request, response)
    return
}

func NewQueryIdcLineExportRequest() (request *QueryIdcLineExportRequest) {
    request = &QueryIdcLineExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryIdcLineExport")
    return
}

func NewQueryIdcLineExportResponse() (response *QueryIdcLineExportResponse) {
    response = &QueryIdcLineExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询idc专线出口信息
func (c *Client) QueryIdcLineExport(request *QueryIdcLineExportRequest) (response *QueryIdcLineExportResponse, err error) {
    if request == nil {
        request = NewQueryIdcLineExportRequest()
    }
    response = NewQueryIdcLineExportResponse()
    err = c.Send(request, response)
    return
}

func NewQueryIpRequest() (request *QueryIpRequest) {
    request = &QueryIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryIp")
    return
}

func NewQueryIpResponse() (response *QueryIpResponse) {
    response = &QueryIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IP详细信息查询
func (c *Client) QueryIp(request *QueryIpRequest) (response *QueryIpResponse, err error) {
    if request == nil {
        request = NewQueryIpRequest()
    }
    response = NewQueryIpResponse()
    err = c.Send(request, response)
    return
}

func NewQueryNetDeviceRequest() (request *QueryNetDeviceRequest) {
    request = &QueryNetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryNetDevice")
    return
}

func NewQueryNetDeviceResponse() (response *QueryNetDeviceResponse) {
    response = &QueryNetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络设备配置信息查询
func (c *Client) QueryNetDevice(request *QueryNetDeviceRequest) (response *QueryNetDeviceResponse, err error) {
    if request == nil {
        request = NewQueryNetDeviceRequest()
    }
    response = NewQueryNetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryNetPortRequest() (request *QueryNetPortRequest) {
    request = &QueryNetPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryNetPort")
    return
}

func NewQueryNetPortResponse() (response *QueryNetPortResponse) {
    response = &QueryNetPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网络设备端口信息查询
func (c *Client) QueryNetPort(request *QueryNetPortRequest) (response *QueryNetPortResponse, err error) {
    if request == nil {
        request = NewQueryNetPortRequest()
    }
    response = NewQueryNetPortResponse()
    err = c.Send(request, response)
    return
}

func NewQueryNetSegmentRequest() (request *QueryNetSegmentRequest) {
    request = &QueryNetSegmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryNetSegment")
    return
}

func NewQueryNetSegmentResponse() (response *QueryNetSegmentResponse) {
    response = &QueryNetSegmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 网段信息查询
func (c *Client) QueryNetSegment(request *QueryNetSegmentRequest) (response *QueryNetSegmentResponse, err error) {
    if request == nil {
        request = NewQueryNetSegmentRequest()
    }
    response = NewQueryNetSegmentResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOutbandInfoRequest() (request *QueryOutbandInfoRequest) {
    request = &QueryOutbandInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryOutbandInfo")
    return
}

func NewQueryOutbandInfoResponse() (response *QueryOutbandInfoResponse) {
    response = &QueryOutbandInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询带外用户信息
func (c *Client) QueryOutbandInfo(request *QueryOutbandInfoRequest) (response *QueryOutbandInfoResponse, err error) {
    if request == nil {
        request = NewQueryOutbandInfoRequest()
    }
    response = NewQueryOutbandInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOutbandTaskRequest() (request *QueryOutbandTaskRequest) {
    request = &QueryOutbandTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryOutbandTask")
    return
}

func NewQueryOutbandTaskResponse() (response *QueryOutbandTaskResponse) {
    response = &QueryOutbandTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询带外任务
func (c *Client) QueryOutbandTask(request *QueryOutbandTaskRequest) (response *QueryOutbandTaskResponse, err error) {
    if request == nil {
        request = NewQueryOutbandTaskRequest()
    }
    response = NewQueryOutbandTaskResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRackDescRequest() (request *QueryRackDescRequest) {
    request = &QueryRackDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryRackDesc")
    return
}

func NewQueryRackDescResponse() (response *QueryRackDescResponse) {
    response = &QueryRackDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询idc机架机位配置信息
func (c *Client) QueryRackDesc(request *QueryRackDescRequest) (response *QueryRackDescResponse, err error) {
    if request == nil {
        request = NewQueryRackDescRequest()
    }
    response = NewQueryRackDescResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRackInfoRequest() (request *QueryRackInfoRequest) {
    request = &QueryRackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryRackInfo")
    return
}

func NewQueryRackInfoResponse() (response *QueryRackInfoResponse) {
    response = &QueryRackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询idc机架机位配置信息
func (c *Client) QueryRackInfo(request *QueryRackInfoRequest) (response *QueryRackInfoResponse, err error) {
    if request == nil {
        request = NewQueryRackInfoRequest()
    }
    response = NewQueryRackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryServerRequest() (request *QueryServerRequest) {
    request = &QueryServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QueryServer")
    return
}

func NewQueryServerResponse() (response *QueryServerResponse) {
    response = &QueryServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实体服务器配置信息
func (c *Client) QueryServer(request *QueryServerRequest) (response *QueryServerResponse, err error) {
    if request == nil {
        request = NewQueryServerRequest()
    }
    response = NewQueryServerResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySnmpDataRequest() (request *QuerySnmpDataRequest) {
    request = &QuerySnmpDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QuerySnmpData")
    return
}

func NewQuerySnmpDataResponse() (response *QuerySnmpDataResponse) {
    response = &QuerySnmpDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp根据设备ID、指标名查询数据
func (c *Client) QuerySnmpData(request *QuerySnmpDataRequest) (response *QuerySnmpDataResponse, err error) {
    if request == nil {
        request = NewQuerySnmpDataRequest()
    }
    response = NewQuerySnmpDataResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySnmpDetailRequest() (request *QuerySnmpDetailRequest) {
    request = &QuerySnmpDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QuerySnmpDetail")
    return
}

func NewQuerySnmpDetailResponse() (response *QuerySnmpDetailResponse) {
    response = &QuerySnmpDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp根据设备ID、指标名、端口查询
func (c *Client) QuerySnmpDetail(request *QuerySnmpDetailRequest) (response *QuerySnmpDetailResponse, err error) {
    if request == nil {
        request = NewQuerySnmpDetailRequest()
    }
    response = NewQuerySnmpDetailResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySnmpDeviceByTemplateRequest() (request *QuerySnmpDeviceByTemplateRequest) {
    request = &QuerySnmpDeviceByTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QuerySnmpDeviceByTemplate")
    return
}

func NewQuerySnmpDeviceByTemplateResponse() (response *QuerySnmpDeviceByTemplateResponse) {
    response = &QuerySnmpDeviceByTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp查询模版对应绑定的设备类型
func (c *Client) QuerySnmpDeviceByTemplate(request *QuerySnmpDeviceByTemplateRequest) (response *QuerySnmpDeviceByTemplateResponse, err error) {
    if request == nil {
        request = NewQuerySnmpDeviceByTemplateRequest()
    }
    response = NewQuerySnmpDeviceByTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySnmpExportDataRequest() (request *QuerySnmpExportDataRequest) {
    request = &QuerySnmpExportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QuerySnmpExportData")
    return
}

func NewQuerySnmpExportDataResponse() (response *QuerySnmpExportDataResponse) {
    response = &QuerySnmpExportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp根据专线编号查询出口流量
func (c *Client) QuerySnmpExportData(request *QuerySnmpExportDataRequest) (response *QuerySnmpExportDataResponse, err error) {
    if request == nil {
        request = NewQuerySnmpExportDataRequest()
    }
    response = NewQuerySnmpExportDataResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySnmpLineDataRequest() (request *QuerySnmpLineDataRequest) {
    request = &QuerySnmpLineDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QuerySnmpLineData")
    return
}

func NewQuerySnmpLineDataResponse() (response *QuerySnmpLineDataResponse) {
    response = &QuerySnmpLineDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp根据专线编号查询专线流量
func (c *Client) QuerySnmpLineData(request *QuerySnmpLineDataRequest) (response *QuerySnmpLineDataResponse, err error) {
    if request == nil {
        request = NewQuerySnmpLineDataRequest()
    }
    response = NewQuerySnmpLineDataResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySnmpTemplateByDeviceRequest() (request *QuerySnmpTemplateByDeviceRequest) {
    request = &QuerySnmpTemplateByDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "QuerySnmpTemplateByDevice")
    return
}

func NewQuerySnmpTemplateByDeviceResponse() (response *QuerySnmpTemplateByDeviceResponse) {
    response = &QuerySnmpTemplateByDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp查询设备对应的模版
func (c *Client) QuerySnmpTemplateByDevice(request *QuerySnmpTemplateByDeviceRequest) (response *QuerySnmpTemplateByDeviceResponse, err error) {
    if request == nil {
        request = NewQuerySnmpTemplateByDeviceRequest()
    }
    response = NewQuerySnmpTemplateByDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewRecycleServerIpRequest() (request *RecycleServerIpRequest) {
    request = &RecycleServerIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "RecycleServerIp")
    return
}

func NewRecycleServerIpResponse() (response *RecycleServerIpResponse) {
    response = &RecycleServerIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务器内外网IP回收
func (c *Client) RecycleServerIp(request *RecycleServerIpRequest) (response *RecycleServerIpResponse, err error) {
    if request == nil {
        request = NewRecycleServerIpRequest()
    }
    response = NewRecycleServerIpResponse()
    err = c.Send(request, response)
    return
}

func NewReinstallServerRequest() (request *ReinstallServerRequest) {
    request = &ReinstallServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "ReinstallServer")
    return
}

func NewReinstallServerResponse() (response *ReinstallServerResponse) {
    response = &ReinstallServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过接口对服务器进行重装
func (c *Client) ReinstallServer(request *ReinstallServerRequest) (response *ReinstallServerResponse, err error) {
    if request == nil {
        request = NewReinstallServerRequest()
    }
    response = NewReinstallServerResponse()
    err = c.Send(request, response)
    return
}

func NewRelocateServerRequest() (request *RelocateServerRequest) {
    request = &RelocateServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "RelocateServer")
    return
}

func NewRelocateServerResponse() (response *RelocateServerResponse) {
    response = &RelocateServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务器搬迁
func (c *Client) RelocateServer(request *RelocateServerRequest) (response *RelocateServerResponse, err error) {
    if request == nil {
        request = NewRelocateServerRequest()
    }
    response = NewRelocateServerResponse()
    err = c.Send(request, response)
    return
}

func NewRelocateServerFinishRequest() (request *RelocateServerFinishRequest) {
    request = &RelocateServerFinishRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "RelocateServerFinish")
    return
}

func NewRelocateServerFinishResponse() (response *RelocateServerFinishResponse) {
    response = &RelocateServerFinishResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务器搬迁确认
func (c *Client) RelocateServerFinish(request *RelocateServerFinishRequest) (response *RelocateServerFinishResponse, err error) {
    if request == nil {
        request = NewRelocateServerFinishRequest()
    }
    response = NewRelocateServerFinishResponse()
    err = c.Send(request, response)
    return
}

func NewRelocateServerQueryRequest() (request *RelocateServerQueryRequest) {
    request = &RelocateServerQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "RelocateServerQuery")
    return
}

func NewRelocateServerQueryResponse() (response *RelocateServerQueryResponse) {
    response = &RelocateServerQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 服务器搬迁历史查询
func (c *Client) RelocateServerQuery(request *RelocateServerQueryRequest) (response *RelocateServerQueryResponse, err error) {
    if request == nil {
        request = NewRelocateServerQueryRequest()
    }
    response = NewRelocateServerQueryResponse()
    err = c.Send(request, response)
    return
}

func NewSyncSnmpDataRequest() (request *SyncSnmpDataRequest) {
    request = &SyncSnmpDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "SyncSnmpData")
    return
}

func NewSyncSnmpDataResponse() (response *SyncSnmpDataResponse) {
    response = &SyncSnmpDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Snmp同步全量数据
func (c *Client) SyncSnmpData(request *SyncSnmpDataRequest) (response *SyncSnmpDataResponse, err error) {
    if request == nil {
        request = NewSyncSnmpDataRequest()
    }
    response = NewSyncSnmpDataResponse()
    err = c.Send(request, response)
    return
}

func NewSyslogAddKeywordRequest() (request *SyslogAddKeywordRequest) {
    request = &SyslogAddKeywordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "SyslogAddKeyword")
    return
}

func NewSyslogAddKeywordResponse() (response *SyslogAddKeywordResponse) {
    response = &SyslogAddKeywordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Syslog新增关键字配置
func (c *Client) SyslogAddKeyword(request *SyslogAddKeywordRequest) (response *SyslogAddKeywordResponse, err error) {
    if request == nil {
        request = NewSyslogAddKeywordRequest()
    }
    response = NewSyslogAddKeywordResponse()
    err = c.Send(request, response)
    return
}

func NewSyslogDeleteKeywordRequest() (request *SyslogDeleteKeywordRequest) {
    request = &SyslogDeleteKeywordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "SyslogDeleteKeyword")
    return
}

func NewSyslogDeleteKeywordResponse() (response *SyslogDeleteKeywordResponse) {
    response = &SyslogDeleteKeywordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Syslog删除关键字
func (c *Client) SyslogDeleteKeyword(request *SyslogDeleteKeywordRequest) (response *SyslogDeleteKeywordResponse, err error) {
    if request == nil {
        request = NewSyslogDeleteKeywordRequest()
    }
    response = NewSyslogDeleteKeywordResponse()
    err = c.Send(request, response)
    return
}

func NewSyslogQueryRequest() (request *SyslogQueryRequest) {
    request = &SyslogQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "SyslogQuery")
    return
}

func NewSyslogQueryResponse() (response *SyslogQueryResponse) {
    response = &SyslogQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Syslog信息查询
func (c *Client) SyslogQuery(request *SyslogQueryRequest) (response *SyslogQueryResponse, err error) {
    if request == nil {
        request = NewSyslogQueryRequest()
    }
    response = NewSyslogQueryResponse()
    err = c.Send(request, response)
    return
}

func NewSyslogQueryKeywordRequest() (request *SyslogQueryKeywordRequest) {
    request = &SyslogQueryKeywordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "SyslogQueryKeyword")
    return
}

func NewSyslogQueryKeywordResponse() (response *SyslogQueryKeywordResponse) {
    response = &SyslogQueryKeywordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Syslog查询关键字
func (c *Client) SyslogQueryKeyword(request *SyslogQueryKeywordRequest) (response *SyslogQueryKeywordResponse, err error) {
    if request == nil {
        request = NewSyslogQueryKeywordRequest()
    }
    response = NewSyslogQueryKeywordResponse()
    err = c.Send(request, response)
    return
}

func NewSyslogUpdateKeywordRequest() (request *SyslogUpdateKeywordRequest) {
    request = &SyslogUpdateKeywordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcos", APIVersion, "SyslogUpdateKeyword")
    return
}

func NewSyslogUpdateKeywordResponse() (response *SyslogUpdateKeywordResponse) {
    response = &SyslogUpdateKeywordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Syslog修改关键字
func (c *Client) SyslogUpdateKeyword(request *SyslogUpdateKeywordRequest) (response *SyslogUpdateKeywordResponse, err error) {
    if request == nil {
        request = NewSyslogUpdateKeywordRequest()
    }
    response = NewSyslogUpdateKeywordResponse()
    err = c.Send(request, response)
    return
}
