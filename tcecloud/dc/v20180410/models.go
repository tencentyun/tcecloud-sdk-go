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

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

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
	City *string `json:"City,omitempty" name:"City"`

	// 物理专线用户侧接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。
	CloudPortType []*string `json:"CloudPortType,omitempty" name:"CloudPortType" list`

	// 物理专线IDC侧接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。
	IdcPortType []*string `json:"IdcPortType,omitempty" name:"IdcPortType" list`
}

type ApproveDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 专线通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 专线通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 审批结果, True: 同意， False: 不同意
	Approved *bool `json:"Approved,omitempty" name:"Approved"`

	// 备注
	Comments *string `json:"Comments,omitempty" name:"Comments"`
}

func (r *ApproveDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApproveDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApproveDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApproveDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApproveDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BgpPeer struct {

	// 用户侧，BGP Asn
	Asn *int64 `json:"Asn,omitempty" name:"Asn"`

	// 用户侧BGP密钥
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`
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
	CloudPortType *string `json:"CloudPortType,omitempty" name:"CloudPortType"`

	// 物理专线接入接口带宽，单位为Mbps，默认值为1000，取值范围为 [2, 10240]。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// 物理专线接入IDC侧端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。
	IdcPortType *string `json:"IdcPortType,omitempty" name:"IdcPortType"`

	// 本地数据中心所在城市
	IdcCity *string `json:"IdcCity,omitempty" name:"IdcCity"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`
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

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 专线网关 ID，例如 dcg-d545ddf
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// 专线带宽，单位：Mbps
	// 默认是物理专线带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// BGP ：BGP路由
	// STATIC：静态
	// 默认为 BGP 路由
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// 静态路由，用户IDC的网段地址
	IdcRoutes *string `json:"IdcRoutes,omitempty" name:"IdcRoutes"`

	// vlan，范围：0 ~ 3000
	// 0：不开启子接口
	// 默认值是非0
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP
	CloudAddress *string `json:"CloudAddress,omitempty" name:"CloudAddress"`

	// CustomerAddress，用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 通道负载均衡模式：
	// None 非冗余模式
	// LoadBalance：负载均衡
	// MasterSlave：主备
	LoadMode *string `json:"LoadMode,omitempty" name:"LoadMode"`

	// vpc数字ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 是否开启BFD
	EnableBfd *bool `json:"EnableBfd,omitempty" name:"EnableBfd"`

	// 互联地址掩码
	ConnectSubnetMask *uint64 `json:"ConnectSubnetMask,omitempty" name:"ConnectSubnetMask"`

	// vpc所在地域
	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`

	// 物理专线 owner，缺省为当前客户（物理专线 owner）
	// 共享专线时这里需要填写共享专线的开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 关联的冗余通道ID
	RelatedDirectConnectTunnelId *string `json:"RelatedDirectConnectTunnelId,omitempty" name:"RelatedDirectConnectTunnelId"`

	// BFD协议interval配置
	BfdInterval *int64 `json:"BfdInterval,omitempty" name:"BfdInterval"`
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
	LoadMode *string `json:"LoadMode,omitempty" name:"LoadMode"`

	// 互联地址掩码
	ConnectSubnetMask *uint64 `json:"ConnectSubnetMask,omitempty" name:"ConnectSubnetMask"`

	// 专线所有者账号UIN
	DirectConnectOwnerUin *uint64 `json:"DirectConnectOwnerUin,omitempty" name:"DirectConnectOwnerUin"`

	// bfd协议中的interval字段值
	BfdInterval *uint64 `json:"BfdInterval,omitempty" name:"BfdInterval"`

	// 关联的通道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedDirectConnectTunnelId *string `json:"RelatedDirectConnectTunnelId,omitempty" name:"RelatedDirectConnectTunnelId"`

	// 通道备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comments *string `json:"Comments,omitempty" name:"Comments"`
}

type ModifyDirectConnectTunnelAttributeRequest struct {
	*tchttp.BaseRequest

	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 用户侧网段地址
	IdcRoutes *string `json:"IdcRoutes,omitempty" name:"IdcRoutes"`

	// 专用通道带宽值，单位为M。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 通道备注
	Comments *string `json:"Comments,omitempty" name:"Comments"`

	// 是否开启BFD
	EnableBfd *bool `json:"EnableBfd,omitempty" name:"EnableBfd"`

	// BFD报文interval时间
	BfdInterval *int64 `json:"BfdInterval,omitempty" name:"BfdInterval"`

	// 是否开启组播
	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// 通道组播组地址
	MulticastGroups *string `json:"MulticastGroups,omitempty" name:"MulticastGroups"`
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

type UpdateVifAssociatedRequest struct {
	*tchttp.BaseRequest

	// 专线通道ID, 比如dcx-nsurd338
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 通道负载均衡模式：
	// None 非冗余模式
	// LoadBalance：负载均衡
	// MasterSlave：主备
	LoadMode *string `json:"LoadMode,omitempty" name:"LoadMode"`

	// 关联专线通道ID，比如dcx-nsurd338
	RelatedDirectConnectTunnelId *string `json:"RelatedDirectConnectTunnelId,omitempty" name:"RelatedDirectConnectTunnelId"`
}

func (r *UpdateVifAssociatedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateVifAssociatedRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateVifAssociatedResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateVifAssociatedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateVifAssociatedResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
