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

package v20190116

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type AttachEntityOfPolicy struct {

	// 实体ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 实体名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实体Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 关联类型。1 用户关联 ； 2 用户组关联
	RelatedType *uint64 `json:"RelatedType,omitempty" name:"RelatedType"`
}

type AttachRolePoliciesRequest struct {
	*tchttp.BaseRequest

	// 角色ID(与角色名称必传一项)
	RoleId *uint64 `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称(与角色ID必传一项)
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 策略ID list(与策略名 list必传一项)
	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId" list`

	// 策略名 list(与策略ID list必传一项)
	PolicyName []*string `json:"PolicyName,omitempty" name:"PolicyName" list`
}

func (r *AttachRolePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachRolePoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachRolePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachRolePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachRolePoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID，入参PolicyId与PolicyName二选一
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 角色ID，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	AttachRoleId *string `json:"AttachRoleId,omitempty" name:"AttachRoleId"`

	// 角色名称，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	AttachRoleName *string `json:"AttachRoleName,omitempty" name:"AttachRoleName"`

	// 策略名，入参PolicyId与PolicyName二选一
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *AttachRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachRolePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachRolePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachedPolicyOfRole struct {

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 绑定时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 策略类型，User表示自定义策略，QCS表示预设策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 策略创建方式，1表示按产品功能或项目权限创建，其他表示按策略语法创建
	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
}

type CreateCASProviderRequest struct {
	*tchttp.BaseRequest

	// 身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份提供商url地址根目录
	CasRoot *string `json:"CasRoot,omitempty" name:"CasRoot"`

	// 身份提供商登录地址
	CasLoginUrl *string `json:"CasLoginUrl,omitempty" name:"CasLoginUrl"`

	// 身份提供商验证ticket地址
	CasValidateUrl *string `json:"CasValidateUrl,omitempty" name:"CasValidateUrl"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 身份提供商登出地址
	CasLogoutUrl *string `json:"CasLogoutUrl,omitempty" name:"CasLogoutUrl"`
}

func (r *CreateCASProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCASProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCASProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 身份提供商名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCASProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCASProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 角色描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否允许登录 1 为允许 0 为不允许
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 申请角色临时密钥的最长有效期限制(范围：0~43200)
	SessionDuration *uint64 `json:"SessionDuration,omitempty" name:"SessionDuration"`

	// 角色类型(system|user)
	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色ID
		RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleRequest struct {
	*tchttp.BaseRequest

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCasProviderRequest struct {
	*tchttp.BaseRequest

	// 每页展现多少条
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 从第几条开始查询
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 展现字段
	Fields *string `json:"Fields,omitempty" name:"Fields"`
}

func (r *DescribeCasProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCasProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCasProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 列表
		List []*ProviderList `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCasProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCasProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest

	// 页码，从1开始
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页行数，不能大于200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 按角色的服务账号载体过滤
	Service *string `json:"Service,omitempty" name:"Service"`

	// 按角色名或角色描述过滤
	Keyword []*string `json:"Keyword,omitempty" name:"Keyword" list`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoleListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色详情列表。
		List []*RoleInfo `json:"List,omitempty" name:"List" list`

		// 角色总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoleListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachGroupsPolicyRequest struct {
	*tchttp.BaseRequest

	// 用户组ID list
	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId" list`

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DetachGroupsPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachGroupsPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachGroupsPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachGroupsPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachGroupsPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachUsersPolicyRequest struct {
	*tchttp.BaseRequest

	// 目标用户ID list
	TargetUin []*uint64 `json:"TargetUin,omitempty" name:"TargetUin" list`

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DetachUsersPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachUsersPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachUsersPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachUsersPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachUsersPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略Id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *GetPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略名
		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

		// 策略描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 1 表示自定义策略，2 表示预设策略
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// 创建时间
		AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

		// 最近更新时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 策略文档
		PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

		// 备注
		PresetAlias *string `json:"PresetAlias,omitempty" name:"PresetAlias"`

		// 是否服务相关策略
		IsServiceLinkedRolePolicy *uint64 `json:"IsServiceLinkedRolePolicy,omitempty" name:"IsServiceLinkedRolePolicy"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRoleRequest struct {
	*tchttp.BaseRequest

	// 角色 ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *GetRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色详情
		RoleInfo *RoleInfo `json:"RoleInfo,omitempty" name:"RoleInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetServiceApiListRequest struct {
	*tchttp.BaseRequest

	// 按服务Id匹配，如cvm
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetServiceApiListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetServiceApiListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetServiceApiListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务以及API列表信息
		List []*ServiceApiInfo `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceApiListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetServiceApiListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesRequest struct {
	*tchttp.BaseRequest

	// 页码，从 1 开始
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页行数，不能大于200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 角色 ID。用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名。用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 按策略类型过滤，User表示仅查询自定义策略，QCS表示仅查询预设策略
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
}

func (r *ListAttachedRolePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedRolePoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色关联的策略列表
		List []*AttachedPolicyOfRole `json:"List,omitempty" name:"List" list`

		// 角色关联的策略总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedRolePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedRolePoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 可取值 'All'、'User'、'Group' 和 'Role'，'All' 表示获取所有实体类型，'User' 表示只获取子账号，'Group' 表示只获取用户组，'Role' 表示只获取角色，默认取 'All'
	EntityFilter *string `json:"EntityFilter,omitempty" name:"EntityFilter"`
}

func (r *ListEntitiesForPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEntitiesForPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实体总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 实体列表
		List []*AttachEntityOfPolicy `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEntitiesForPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEntitiesForPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesRequest struct {
	*tchttp.BaseRequest

	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 页码，默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 可取值 'All'、'QCS' 和 'Local'，'All' 获取所有策略，'QCS' 只获取预设策略，'Local' 只获取自定义策略，默认取 'All'
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// 按策略名匹配
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 按Uin匹配
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// 按组Id匹配
	TargetGroupId *uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 按角色Id匹配
	TargetRoleId *uint64 `json:"TargetRoleId,omitempty" name:"TargetRoleId"`

	// 按产品Id匹配，如cvm
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 调用来源，控制台为"console"
	Client *string `json:"Client,omitempty" name:"Client"`

	// 按Uin标记关联
	FlagUin *uint64 `json:"FlagUin,omitempty" name:"FlagUin"`

	// 按GroupId标记关联
	FlagGroupId *uint64 `json:"FlagGroupId,omitempty" name:"FlagGroupId"`

	// 按角色Id标记关联
	FlagRoleId *uint64 `json:"FlagRoleId,omitempty" name:"FlagRoleId"`

	// 是否展示服务相关策略
	ShowServiceLinkedRolePolicy *uint64 `json:"ShowServiceLinkedRolePolicy,omitempty" name:"ShowServiceLinkedRolePolicy"`

	// 创建类型：1.按产品功能或项目权限创建；2.按策略语法创建；3.按策略生成器创建；4.按标签授权创建；5.按权限边界规则创建
	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
}

func (r *ListPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProviderList struct {

	// cas Id号
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 创建者uin
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 主账号
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份提供商描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 身份提供商类型
	ProviderType *uint64 `json:"ProviderType,omitempty" name:"ProviderType"`

	// 身份提供商状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Cas属性
	Cas *string `json:"Cas,omitempty" name:"Cas"`
}

type RoleInfo struct {

	// 角色ID
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色的策略文档
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 角色描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 角色的创建时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 角色的最近一次时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ServiceApiInfo struct {

	// 服务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务ID
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 服务介绍文档链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArnDocument *string `json:"ArnDocument,omitempty" name:"ArnDocument"`

	// API信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiList []*ServiceApiListInfo `json:"ApiList,omitempty" name:"ApiList" list`

	// 条件规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionKeyList []*string `json:"ConditionKeyList,omitempty" name:"ConditionKeyList" list`
}

type ServiceApiListInfo struct {

	// API名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要关联对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNeedObject *string `json:"IsNeedObject,omitempty" name:"IsNeedObject"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 接口类别：0.读取，1.写入，2.标记，3.列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadWriteDetail *uint64 `json:"ReadWriteDetail,omitempty" name:"ReadWriteDetail"`

	// 授权粒度：0.接口级，1.资源级
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterfaceLevel *uint64 `json:"InterfaceLevel,omitempty" name:"InterfaceLevel"`

	// 资源六段式范例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceExample *string `json:"ResourceExample,omitempty" name:"ResourceExample"`
}

type UpdateAssumeRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *UpdateAssumeRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAssumeRolePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAssumeRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAssumeRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAssumeRolePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleDescriptionRequest struct {
	*tchttp.BaseRequest

	// 角色描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *UpdateRoleDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRoleDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRoleDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
