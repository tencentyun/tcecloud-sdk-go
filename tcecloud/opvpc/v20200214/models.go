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

package v20200214

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AclInfo struct {

	// VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// AclId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclId *string `json:"AclId,omitempty" name:"AclId"`

	// UniqVpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// SubnetNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetNum *string `json:"SubnetNum,omitempty" name:"SubnetNum"`

	// VpcSubnet
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcSubnet *string `json:"VpcSubnet,omitempty" name:"VpcSubnet"`

	// VpcMask
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcMask *string `json:"VpcMask,omitempty" name:"VpcMask"`

	// VpcName
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VpcBDefault
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcBDefault *string `json:"VpcBDefault,omitempty" name:"VpcBDefault"`

	// Owner
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// UniqAclId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqAclId *string `json:"UniqAclId,omitempty" name:"UniqAclId"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// SubNetSet
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubNetSet []*AclSubNetInfo `json:"SubNetSet,omitempty" name:"SubNetSet" list`
}

type AclSubNetInfo struct {

	// Subnet
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`

	// Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// BDefault
	// 注意：此字段可能返回 null，表示取不到有效值。
	BDefault *string `json:"BDefault,omitempty" name:"BDefault"`

	// Mask
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mask *string `json:"Mask,omitempty" name:"Mask"`

	// ZoneId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// IntMask
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntMask *string `json:"IntMask,omitempty" name:"IntMask"`

	// UniqSubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// SubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// DhcpFlag
	// 注意：此字段可能返回 null，表示取不到有效值。
	DhcpFlag *string `json:"DhcpFlag,omitempty" name:"DhcpFlag"`
}

type ActiveHealthCheckInfo struct {

	// DownTimes
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownTimes *string `json:"DownTimes,omitempty" name:"DownTimes"`

	// expect alive
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectAlive *string `json:"ExpectAlive,omitempty" name:"ExpectAlive"`

	// http send
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpSend *string `json:"HttpSend,omitempty" name:"HttpSend"`

	// interval
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// method
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// protocol
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// server name
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// up times
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpTimes *string `json:"UpTimes,omitempty" name:"UpTimes"`
}

type DescribeAclListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序类型
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 过滤器列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAclListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAclListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAclListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// acl总数
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// acl 列表
		AclSet []*AclInfo `json:"AclSet,omitempty" name:"AclSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAclListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAclListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmForSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 安全组id
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 业务id
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCvmForSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCvmForSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmForSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// vm 列表
		VmSet []*SgVmInfo `json:"VmSet,omitempty" name:"VmSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCvmForSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCvmForSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeL4RuleExRequest struct {
	*tchttp.BaseRequest

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 获取量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeL4RuleExRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeL4RuleExRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeL4RuleExResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 规则列表
		RuleSet []*Rule `json:"RuleSet,omitempty" name:"RuleSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeL4RuleExResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeL4RuleExResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeL7RuleExRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器数组
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeL7RuleExRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeL7RuleExRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeL7RuleExResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 规则列表
		RuleSet []*L7RuleInfo `json:"RuleSet,omitempty" name:"RuleSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeL7RuleExResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeL7RuleExResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLbLimitRequest struct {
	*tchttp.BaseRequest

	// appid 名字
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *DescribeLbLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLbLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLbLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// lb限制数组
		LbLimitSet []*LbLimit `json:"LbLimitSet,omitempty" name:"LbLimitSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLbLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLbLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewayRequest struct {
	*tchttp.BaseRequest

	// 排序类型
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// nat 网关数量
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// nat 网关列表
		NatSet []*NatInfo `json:"NatSet,omitempty" name:"NatSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupForCvmRequest struct {
	*tchttp.BaseRequest

	// 虚拟机id列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 虚拟机id
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeSecurityGroupForCvmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupForCvmRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupForCvmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组列表
		UsgSet []*VmSgInfoOuter `json:"UsgSet,omitempty" name:"UsgSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupForCvmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupForCvmResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 安全组id列表
	UsgIdSet []*string `json:"UsgIdSet,omitempty" name:"UsgIdSet" list`
}

func (r *DescribeSecurityGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组策略列表
		UsgSet []*UsgPolicyInfo `json:"UsgSet,omitempty" name:"UsgSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序类型
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 过滤器列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数量
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// sg列表
		SgSet []*SgInfo `json:"SgSet,omitempty" name:"SgSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetExRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeSubnetExRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetExRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetExResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount []*string `json:"TotalCount,omitempty" name:"TotalCount" list`

		// 子网列表
		SubNetSet []*SubNetInfo `json:"SubNetSet,omitempty" name:"SubNetSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetExResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetExResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsExRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 简略返回
	Brief *string `json:"Brief,omitempty" name:"Brief"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 过滤器列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeVpcsExRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcsExRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsExResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 个数
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// vpc列表
		VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcsExResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcsExResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 字段名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ForwardStrategyInfo struct {

	// detail
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// type
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type HealthCheckInfo struct {

	// ActiveHealthCheck
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveHealthCheck *ActiveHealthCheckInfo `json:"ActiveHealthCheck,omitempty" name:"ActiveHealthCheck"`

	// PassiveHealthCheck
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassiveHealthCheck *PassiveHealthCheckInfo `json:"PassiveHealthCheck,omitempty" name:"PassiveHealthCheck"`

	// type
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type InOutBoundInfo struct {

	// Proto
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Action
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// Port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// Desc
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type L7RuleInfo struct {

	// location list
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationSet []*LocationInfo `json:"LocationSet,omitempty" name:"LocationSet" list`

	// virtual service
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualService *VirtualServiceInfo `json:"VirtualService,omitempty" name:"VirtualService"`
}

type LbLimit struct {

	// app id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 配额类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 当前配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 最大配额
	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
}

type LocationInfo struct {

	// CustomizedConf
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomizedConf *string `json:"CustomizedConf,omitempty" name:"CustomizedConf"`

	// ForwardStrategy
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardStrategy *ForwardStrategyInfo `json:"ForwardStrategy,omitempty" name:"ForwardStrategy"`

	// HealthCheck
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *HealthCheckInfo `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// KeepaliveConnNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepaliveConnNum *string `json:"KeepaliveConnNum,omitempty" name:"KeepaliveConnNum"`

	// KeepaliveEnable
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepaliveEnable *string `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// RsSet
	// 注意：此字段可能返回 null，表示取不到有效值。
	RsSet []*Rs `json:"RsSet,omitempty" name:"RsSet" list`

	// SessionPersistenceStrategy
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionPersistenceStrategy *SessionPersistenceStrategyInfo `json:"SessionPersistenceStrategy,omitempty" name:"SessionPersistenceStrategy"`

	// Type
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// Url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type ModifyLbLimitRequest struct {
	*tchttp.BaseRequest

	// 配额类型
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 当前配额
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 最大配额
	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`

	// 业务id
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *ModifyLbLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLbLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLbLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLbLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLbLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NatInfo struct {

	// WanInLimit
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanInLimit *string `json:"WanInLimit,omitempty" name:"WanInLimit"`

	// VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// UniqVpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Connections
	// 注意：此字段可能返回 null，表示取不到有效值。
	Connections *string `json:"Connections,omitempty" name:"Connections"`

	// Owner
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// UniqNatId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqNatId *string `json:"UniqNatId,omitempty" name:"UniqNatId"`

	// NatId
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// Zone
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ZoneId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// State
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// BDeleted
	// 注意：此字段可能返回 null，表示取不到有效值。
	BDeleted *string `json:"BDeleted,omitempty" name:"BDeleted"`

	// OwedWanOutLimit
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwedWanOutLimit *string `json:"OwedWanOutLimit,omitempty" name:"OwedWanOutLimit"`

	// GwId
	// 注意：此字段可能返回 null，表示取不到有效值。
	GwId *string `json:"GwId,omitempty" name:"GwId"`

	// HealthStatus
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthStatus *string `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// DetectGwIp
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectGwIp *string `json:"DetectGwIp,omitempty" name:"DetectGwIp"`

	// GwIp
	// 注意：此字段可能返回 null，表示取不到有效值。
	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// LastUpdateConnectionsTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateConnectionsTime *string `json:"LastUpdateConnectionsTime,omitempty" name:"LastUpdateConnectionsTime"`

	// OwedFlag
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwedFlag *string `json:"OwedFlag,omitempty" name:"OwedFlag"`

	// Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// LastMaxConnections
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastMaxConnections *string `json:"LastMaxConnections,omitempty" name:"LastMaxConnections"`

	// WanOutLimit
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanOutLimit *string `json:"WanOutLimit,omitempty" name:"WanOutLimit"`

	// Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip []*NatIpInfo `json:"Ip,omitempty" name:"Ip" list`
}

type NatIpInfo struct {

	// Eip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// LanIp
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// EipId
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// TsvIp
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsvIp *string `json:"TsvIp,omitempty" name:"TsvIp"`

	// BBlocked
	// 注意：此字段可能返回 null，表示取不到有效值。
	BBlocked *string `json:"BBlocked,omitempty" name:"BBlocked"`

	// BAutoApply
	// 注意：此字段可能返回 null，表示取不到有效值。
	BAutoApply *string `json:"BAutoApply,omitempty" name:"BAutoApply"`
}

type PassiveHealthCheckInfo struct {

	// 超时
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailTimeout *string `json:"FailTimeout,omitempty" name:"FailTimeout"`

	// 最大失败次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxFails *string `json:"MaxFails,omitempty" name:"MaxFails"`
}

type Rs struct {

	// rs ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`

	// rs port
	// 注意：此字段可能返回 null，表示取不到有效值。
	RsPort *int64 `json:"RsPort,omitempty" name:"RsPort"`

	// rs 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	RsWeight *int64 `json:"RsWeight,omitempty" name:"RsWeight"`

	// rs vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type Rule struct {

	// app name
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// biz admin
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizAdmin *string `json:"BizAdmin,omitempty" name:"BizAdmin"`

	// bizid
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizId *string `json:"BizId,omitempty" name:"BizId"`

	// biz type
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwdMode *string `json:"FwdMode,omitempty" name:"FwdMode"`

	// 机房名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// 运营商名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// keep time
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepTime *string `json:"KeepTime,omitempty" name:"KeepTime"`

	// mtu
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mtu *string `json:"Mtu,omitempty" name:"Mtu"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 允许quic协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuicEnable *string `json:"QuicEnable,omitempty" name:"QuicEnable"`

	// rule id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 轮询类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleType *string `json:"ScheduleType,omitempty" name:"ScheduleType"`

	// tsv vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsvVip *string `json:"TsvVip,omitempty" name:"TsvVip"`

	// vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// vport
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *string `json:"Vport,omitempty" name:"Vport"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// usage
	// 注意：此字段可能返回 null，表示取不到有效值。
	Usage *string `json:"Usage,omitempty" name:"Usage"`

	// set name
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetName *string `json:"SetName,omitempty" name:"SetName"`

	// rs 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RsNum *string `json:"RsNum,omitempty" name:"RsNum"`

	// rs 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RsSet []*Rs `json:"RsSet,omitempty" name:"RsSet" list`

	// alive interval
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliveInterval *string `json:"AliveInterval,omitempty" name:"AliveInterval"`

	// kick interval
	// 注意：此字段可能返回 null，表示取不到有效值。
	KickInterval *string `json:"KickInterval,omitempty" name:"KickInterval"`

	// probe interval
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProbeInterval *string `json:"ProbeInterval,omitempty" name:"ProbeInterval"`

	// probe timeout
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProbeTimeout *string `json:"ProbeTimeout,omitempty" name:"ProbeTimeout"`
}

type SessionPersistenceStrategyInfo struct {

	// cookie type
	// 注意：此字段可能返回 null，表示取不到有效值。
	CookieKey *string `json:"CookieKey,omitempty" name:"CookieKey"`

	// timeout
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`

	// type
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SgInfo struct {

	// usg id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsgId *string `json:"UsgId,omitempty" name:"UsgId"`

	// name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// owner
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// remark
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// counts
	// 注意：此字段可能返回 null，表示取不到有效值。
	Counts []*SgVmCount `json:"Counts,omitempty" name:"Counts" list`
}

type SgVmCount struct {

	// general
	// 注意：此字段可能返回 null，表示取不到有效值。
	General *string `json:"General,omitempty" name:"General"`

	// cvm
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cvm *string `json:"Cvm,omitempty" name:"Cvm"`
}

type SgVmInfo struct {

	// network interface id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// instance id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// alias
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc name
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// eni type
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniType *string `json:"EniType,omitempty" name:"EniType"`

	// eni name
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniName *string `json:"EniName,omitempty" name:"EniName"`
}

type SubNetInfo struct {

	// VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ZoneId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// UniqVpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VpcName
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Owner
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// SubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// UniqSubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// SubnetName
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// RtbId
	// 注意：此字段可能返回 null，表示取不到有效值。
	RtbId *string `json:"RtbId,omitempty" name:"RtbId"`

	// RtbName
	// 注意：此字段可能返回 null，表示取不到有效值。
	RtbName *string `json:"RtbName,omitempty" name:"RtbName"`

	// UniqRtbId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqRtbId *string `json:"UniqRtbId,omitempty" name:"UniqRtbId"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// CidrBlock
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// AvailableIpNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableIpNum *string `json:"AvailableIpNum,omitempty" name:"AvailableIpNum"`

	// UsedIpNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedIpNum *string `json:"UsedIpNum,omitempty" name:"UsedIpNum"`

	// TotalIpNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalIpNum *string `json:"TotalIpNum,omitempty" name:"TotalIpNum"`

	// Mask
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mask *string `json:"Mask,omitempty" name:"Mask"`
}

type UsgPolicyInfo struct {

	// ErrorCode
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// InBound
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBound []*InOutBoundInfo `json:"InBound,omitempty" name:"InBound" list`

	// ErrorInfo
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`

	// OutBound
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBound []*InOutBoundInfo `json:"OutBound,omitempty" name:"OutBound" list`

	// Version
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

type VipInfo struct {

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 运营商
	// 注意：此字段可能返回 null，表示取不到有效值。
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type VirtualServiceInfo struct {

	// app name
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// bg name
	// 注意：此字段可能返回 null，表示取不到有效值。
	BgName *string `json:"BgName,omitempty" name:"BgName"`

	// biz id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizId *string `json:"BizId,omitempty" name:"BizId"`

	// biz type
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// blocaked enable
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockedEnable *string `json:"BlockedEnable,omitempty" name:"BlockedEnable"`

	// cc guard enable
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcGuardEnable *string `json:"CcGuardEnable,omitempty" name:"CcGuardEnable"`

	// city
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// default server
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultServer *string `json:"DefaultServer,omitempty" name:"DefaultServer"`

	// domain
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// forward protocol
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// fwd mode
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwdMode *string `json:"FwdMode,omitempty" name:"FwdMode"`

	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// ld port
	// 注意：此字段可能返回 null，表示取不到有效值。
	LdPort *string `json:"LdPort,omitempty" name:"LdPort"`

	// 机房名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// set name
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetName *string `json:"SetName,omitempty" name:"SetName"`

	// vip数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vips []*VipInfo `json:"Vips,omitempty" name:"Vips" list`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 端口列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vports []*int64 `json:"Vports,omitempty" name:"Vports" list`
}

type VmSgInfo struct {

	// Sys
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sys *string `json:"Sys,omitempty" name:"Sys"`

	// UsgId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsgId *string `json:"UsgId,omitempty" name:"UsgId"`

	// UsgName
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsgName *string `json:"UsgName,omitempty" name:"UsgName"`

	// ProjectId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// Os
	// 注意：此字段可能返回 null，表示取不到有效值。
	Os *string `json:"Os,omitempty" name:"Os"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// UsgRemark
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsgRemark *string `json:"UsgRemark,omitempty" name:"UsgRemark"`
}

type VmSgInfoOuter struct {

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`

	// vm 管理的sg信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsgInfo []*VmSgInfo `json:"UsgInfo,omitempty" name:"UsgInfo" list`
}

type VpcInfo struct {

	// VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// UniqVpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VpcName
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Owner
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// CidrBlock
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// RtbNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	RtbNum *string `json:"RtbNum,omitempty" name:"RtbNum"`

	// SubnetNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetNum *string `json:"SubnetNum,omitempty" name:"SubnetNum"`

	// VpcPeerNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcPeerNum *string `json:"VpcPeerNum,omitempty" name:"VpcPeerNum"`

	// VpgNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpgNum *string `json:"VpgNum,omitempty" name:"VpgNum"`

	// VpngwNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpngwNum *string `json:"VpngwNum,omitempty" name:"VpngwNum"`

	// VmNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	VmNum *string `json:"VmNum,omitempty" name:"VmNum"`

	// NatNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatNum *string `json:"NatNum,omitempty" name:"NatNum"`

	// AclNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclNum *string `json:"AclNum,omitempty" name:"AclNum"`
}
