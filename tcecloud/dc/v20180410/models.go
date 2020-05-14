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

package v20180410

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AcceptDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 专线ID
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 交换机端口ID，参数为空时，自动分配交换机端口
	PortId *string `json:"PortId,omitempty" name:"PortId"`
}

func (r *AcceptDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AcceptDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AcceptDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 物理专线拥有者接受共享专用通道申请
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *AcceptDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AcceptDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AccessPoint struct {

	// 接入点的名称。
	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`

	// 接入点唯一ID。
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 接入点的状态。AVAILABLE, UNAVAILABLE
	State *string `json:"State,omitempty" name:"State"`

	// 接入点所在机房地址
	Location *string `json:"Location,omitempty" name:"Location"`

	// 接入点支持的运营商列表。
	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator" list`

	// 接入点可用带宽百分比
	AvailableBandwidthRatio *uint64 `json:"AvailableBandwidthRatio,omitempty" name:"AvailableBandwidthRatio"`

	// 当前已使用带宽, 单位MBytes
	CurrentBandwidth *uint64 `json:"CurrentBandwidth,omitempty" name:"CurrentBandwidth"`

	// 带宽告警比例
	NotifyBandwidthRatio *uint64 `json:"NotifyBandwidthRatio,omitempty" name:"NotifyBandwidthRatio"`

	// 接入点总带宽, 单位 MBytes
	TotalBandwidth *uint64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`

	// IDC所在城市
	City *uint64 `json:"City,omitempty" name:"City"`
}

type BgpPeer struct {

	// 用户侧，BGP Asn
	Asn *int64 `json:"Asn,omitempty" name:"Asn"`

	// 用户侧BGP密钥
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`
}

type CreateAccessPointRequest struct {
	*tchttp.BaseRequest

	// 接入点名称
	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`

	// 接入点的状态。AVAILABLE, UNAVAILABLE
	State *string `json:"State,omitempty" name:"State"`

	// 接入点所在机房地址
	Location *string `json:"Location,omitempty" name:"Location"`

	// 接入点所在城市
	City *string `json:"City,omitempty" name:"City"`

	// 接入点支持的运营商列表，"中国电信", u"中国移动", "中国联通", "中国其他", "境外其他"
	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator" list`

	// 接入点可用带宽百分比
	AvailableBandwidthRatio *uint64 `json:"AvailableBandwidthRatio,omitempty" name:"AvailableBandwidthRatio"`

	// 带宽告警比例
	NotifyBandwidthRatio *uint64 `json:"NotifyBandwidthRatio,omitempty" name:"NotifyBandwidthRatio"`

	// 接入点总带宽, 单位 MBytes
	TotalBandwidth *uint64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`
}

func (r *CreateAccessPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccessPointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccessPointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前接入点列表
		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet" list`

		// 当前接入点数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccessPointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 物理专线的名称。
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// 物理专线所在的接入点。
	// 您可以通过调用 DescribeAccessPoints接口获取地域ID。所选择的接入点必须存在且处于可接入的状态。
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 本地数据中心的地理位置。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 物理专线接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。
	PortType *string `json:"PortType,omitempty" name:"PortType"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 物理专线接入接口带宽，单位为Mbps，默认值为1000，取值范围为 [2, 10240]。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`

	// 物理专线调试VLAN。默认开启VLAN，自动分配VLAN。
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。默认自动分配。
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。默认自动分配。
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`
}

func (r *CreateDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 物理专线的ID。
		DirectConnectIdSet []*string `json:"DirectConnectIdSet,omitempty" name:"DirectConnectIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 专线 ID，例如：dc-kd7d06of
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 物理专线 owner，缺省为当前客户（物理专线 owner）
	// 共享专线时这里需要填写共享专线的开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// 网络类型，分别为VPC、BMVPC，CCN，默认是VPC
	// VPC：私有网络
	// BMVPC：黑石网络
	// CCN：云联网
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 网络地域
	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 专线网关 ID，例如 dcg-d545ddf
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// 专线带宽，单位：Mbps
	// 默认是物理专线带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// BGP ：BGP路由
	// STATIC：静态
	// 默认为 BGP 路由
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 静态路由，用户IDC的网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes" list`

	// vlan，范围：0 ~ 3000
	// 0：不开启子接口
	// 默认值是非0
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// CustomerAddress，用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`
}

func (r *CreateDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专用通道ID
		DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet,omitempty" name:"DirectConnectTunnelIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePublicDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 专线 ID，例如：dc-kd7d06of
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 物理专线 owner，缺省为当前客户（物理专线 owner）
	// 共享专线时这里需要填写共享专线的开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// 专用通道带宽，单位：Mbps
	// 默认是物理专线带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// vlan，范围：0 ~ 3000
	// 0：不开启子接口
	// 默认值是非0
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// CustomerAddress，用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`
}

func (r *CreatePublicDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePublicDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePublicDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专线通道ID
		DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet,omitempty" name:"DirectConnectTunnelIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePublicDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePublicDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSwitchRequest struct {
	*tchttp.BaseRequest

	// 专线接入点ID，通过DescribeAccessPoints接口可获取接入点列表
	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 用户自定义交换机名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 交换机型号
	Model *string `json:"Model,omitempty" name:"Model"`

	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// 云控制器使用的账号名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 云控制器使用的密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 云控制器使用的管理IP
	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`

	// 云控制器使用的管理端口
	ManagementPort *uint64 `json:"ManagementPort,omitempty" name:"ManagementPort"`

	// 专线接入交换机vxlan隧道的VTEP ip
	VTEPIP *string `json:"VTEPIP,omitempty" name:"VTEPIP"`
}

func (r *CreateSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSwitchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSwitchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专线接入交换机列表
		SwitchSet []*DcSwitch `json:"SwitchSet,omitempty" name:"SwitchSet" list`

		// 专线接入交换机数目
		TotalCount []*uint64 `json:"TotalCount,omitempty" name:"TotalCount" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSwitchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DcSwitch struct {

	// 专线接入点ID，通过DescribeAccessPoints接口可获取接入点列表
	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 用户自定义交换机名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 交换机型号
	Model *string `json:"Model,omitempty" name:"Model"`

	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// 云控制器使用的管理IP
	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`

	// 云控制器使用的管理端口
	ManagementPort *string `json:"ManagementPort,omitempty" name:"ManagementPort"`

	// 专线接入交换机vxlan隧道的VTEP ip
	VTEPIP *string `json:"VTEPIP,omitempty" name:"VTEPIP"`

	// 专线接入交换机ID
	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
}

type DeleteDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 物理专线的ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
}

func (r *DeleteDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *DeleteDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSwitchRequest struct {
	*tchttp.BaseRequest

	// 交换机ID，通过DescribeSwitches可获取到交换机列表
	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`

	// 交换机管理IP
	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
}

func (r *DeleteSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSwitchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSwitchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSwitchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessPointsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessPointsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessPointsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接入点信息。
		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet" list`

		// 接入点数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessPointsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableSwitchesRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAvailableSwitchesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableSwitchesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableSwitchesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专线接入交换机列表
		SwitchSet []*SwitchNetdevice `json:"SwitchSet,omitempty" name:"SwitchSet" list`

		// 专线接入交换机数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableSwitchesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAvailableSwitchesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectTunnelsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectTunnelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectTunnelsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectTunnelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专用通道列表
		DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet,omitempty" name:"DirectConnectTunnelSet" list`

		// 符合条件的通道总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectTunnelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectTunnelsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件:
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 物理专线 ID数组
	DirectConnectIds []*string `json:"DirectConnectIds,omitempty" name:"DirectConnectIds" list`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 物理专线列表。
		DirectConnectSet []*DirectConnect `json:"DirectConnectSet,omitempty" name:"DirectConnectSet" list`

		// 符合物理专线列表数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchPortsRequest struct {
	*tchttp.BaseRequest

	// 交换机ID, 通过DescribeSwitches可获取交换机列表
	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`

	// Available, Unavailable
	State *string `json:"State,omitempty" name:"State"`
}

func (r *DescribeSwitchPortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSwitchPortsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchPortsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 端口列表
		SwitchPortSet []*SwitchPort `json:"SwitchPortSet,omitempty" name:"SwitchPortSet" list`

		// 端口数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchPortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSwitchPortsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchesRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSwitchesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSwitchesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专线接入交换机列表
		SwitchSet []*DcSwitch `json:"SwitchSet,omitempty" name:"SwitchSet" list`

		// 专线接入交换机数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSwitchesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DirectConnect struct {

	// 物理专线ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 物理专线的名称。
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// 物理专线的接入点ID。
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 物理专线的状态。
	// 申请中：PENDING 
	// 申请驳回：REJECTED   
	// 建设中：ALLOCATED   
	// 已开通：AVAILABLE  
	// 删除中 ：DELETING
	// 已删除：DELETED 。
	State *string `json:"State,omitempty" name:"State"`

	// 物理专线创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 物理专线的开通时间。
	EnabledTime *string `json:"EnabledTime,omitempty" name:"EnabledTime"`

	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 本地数据中心的地理位置。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 物理专线接入接口带宽，单位为Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 用户侧物理专线接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）
	IdcPortType *string `json:"IdcPortType,omitempty" name:"IdcPortType"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`

	// IDC所在城市
	IdcCity *string `json:"IdcCity,omitempty" name:"IdcCity"`

	// 云侧端口类型
	CloudPortType *string `json:"CloudPortType,omitempty" name:"CloudPortType"`

	// 接入点名称
	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`

	// 申请ID
	ApplyId *uint64 `json:"ApplyId,omitempty" name:"ApplyId"`

	// 交换机名称
	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`

	// 专线所有者AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 交换机ID
	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`

	// 交换机端口名称
	SwitchPort *string `json:"SwitchPort,omitempty" name:"SwitchPort"`

	// 专线下就绪的通道数目
	TunnelCount *uint64 `json:"TunnelCount,omitempty" name:"TunnelCount"`
}

type DirectConnectTunnel struct {

	// 专线通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 物理专线ID
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// "专线通道状态
	// 0:AVAILABLE:就绪或者已连接
	// 1: ALLOCATED：通道建设中
	// 2:ALLOCATING:配置中
	// 6:ALLOCATED:配置完成
	// 3:ALTERING:修改中
	// 4:DELETING:删除中
	// 5:DELETED:删除完成
	// 20: PENDING:待接受
	// 21:REJECTED:拒绝"
	State *string `json:"State,omitempty" name:"State"`

	// 专线的拥有者，开发商账号 ID
	DirectConnectOwnerAppId *string `json:"DirectConnectOwnerAppId,omitempty" name:"DirectConnectOwnerAppId"`

	// 网络类型，分别为VPC、BMVPC、CCN
	//  VPC：私有网络 ，BMVPC：黑石网络，CCN：云联网
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 专线网关 ID
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// BGP ：BGP路由 STATIC：静态
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// 用户侧网段地址
	IdcRoutes []*string `json:"IdcRoutes,omitempty" name:"IdcRoutes" list`

	// 专线通道的Vlan
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 云侧互联IP
	CloudAddress *string `json:"CloudAddress,omitempty" name:"CloudAddress"`

	// 用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 专线通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 专线通道创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 专线通道带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 关联的网络自定义探测ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`

	// 是否为Nat通道
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatType *int64 `json:"NatType,omitempty" name:"NatType"`

	// 通道连接的VPC所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`

	// 是否开启BFD
	// 注意：此字段可能返回 null，表示取不到有效值。
	BfdState *int64 `json:"BfdState,omitempty" name:"BfdState"`

	// 专线网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`

	// VPC名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// True: 开启, False: 不开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// 通道下支持的组播组
	// 注意：此字段可能返回 null，表示取不到有效值。
	MulticastGroups []*string `json:"MulticastGroups,omitempty" name:"MulticastGroups" list`

	// 专线通道的拥有者，开发上账号ID
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 通道负载均衡模式：
	// 0：None 1: LoadBalance 2: MasterSlave
	LoadMode []*string `json:"LoadMode,omitempty" name:"LoadMode" list`

	// 互联地址掩码
	ConnectSubnetMask *uint64 `json:"ConnectSubnetMask,omitempty" name:"ConnectSubnetMask"`

	// 专线所有者账号UIN
	DirectConnectOwnerUin *uint64 `json:"DirectConnectOwnerUin,omitempty" name:"DirectConnectOwnerUin"`

	// bfd协议中的interval字段值
	BfdInterval *uint64 `json:"BfdInterval,omitempty" name:"BfdInterval"`

	// 关联的通道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedDirectConnectTunnelId *string `json:"RelatedDirectConnectTunnelId,omitempty" name:"RelatedDirectConnectTunnelId"`
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ImplementDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 物理专线的ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`
}

func (r *ImplementDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImplementDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImplementDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImplementDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImplementDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsSameRegionRequest struct {
	*tchttp.BaseRequest

	// 本端地域
	LocalRegion *string `json:"LocalRegion,omitempty" name:"LocalRegion"`

	// 对端地域
	RemoteRegion *string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`

	// 比较的类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *IsSameRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsSameRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsSameRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// True是同地域，False不同地域
		SameRegion *bool `json:"SameRegion,omitempty" name:"SameRegion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsSameRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsSameRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessPointRequest struct {
	*tchttp.BaseRequest

	// 接入点ID
	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 接入点名称
	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`

	// 接入点的状态。AVAILABLE, UNAVAILABLE
	State *string `json:"State,omitempty" name:"State"`

	// 接入点所在机房地址
	Location *string `json:"Location,omitempty" name:"Location"`

	// 接入点所在城市
	City *string `json:"City,omitempty" name:"City"`

	// 接入点支持的运营商列表，"中国电信", u"中国移动", "中国联通", "中国其他", "境外其他"
	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator" list`

	// 接入点可用带宽百分比
	AvailableBandwidthRatio *uint64 `json:"AvailableBandwidthRatio,omitempty" name:"AvailableBandwidthRatio"`

	// 带宽告警比例
	NotifyBandwidthRatio *uint64 `json:"NotifyBandwidthRatio,omitempty" name:"NotifyBandwidthRatio"`

	// 接入点总带宽, 单位 MBytes
	TotalBandwidth *uint64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`
}

func (r *ModifyAccessPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccessPointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessPointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接入点信息。
		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet" list`

		// 接入点数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccessPointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectAttributeRequest struct {
	*tchttp.BaseRequest

	// 物理专线的ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`

	// 专线所在的端口
	PortId *string `json:"PortId,omitempty" name:"PortId"`
}

func (r *ModifyDirectConnectAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectTunnelAttributeRequest struct {
	*tchttp.BaseRequest

	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 用户侧BGP，包括Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 用户侧网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes" list`

	// 腾讯侧互联IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 用户侧互联IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值，单位为M。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

func (r *ModifyDirectConnectTunnelAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectTunnelAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectTunnelAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectTunnelAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectTunnelAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 专线ID
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
}

func (r *RejectDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 无
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *RejectDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RouteFilterPrefix struct {

	// 用户侧网段地址
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

type SwitchNetdevice struct {

	// 云控制器使用的管理IP
	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`

	// 云控制器使用的管理端口，为空时，业务使用默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagementPort *string `json:"ManagementPort,omitempty" name:"ManagementPort"`

	// 交换机名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// 交换机型号
	Model *string `json:"Model,omitempty" name:"Model"`
}

type SwitchPort struct {

	// 端口数字ID
	PortId *uint64 `json:"PortId,omitempty" name:"PortId"`

	// 端口名称
	PortName *string `json:"PortName,omitempty" name:"PortName"`
}

type UpdateVifNetDetectRequest struct {
	*tchttp.BaseRequest

	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 网络自定义探测ID
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
}

func (r *UpdateVifNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateVifNetDetectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateVifNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateVifNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateVifNetDetectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
