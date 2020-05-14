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

package v20200226

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CheckUserPolicyAttachmentRequest struct {
	*tchttp.BaseRequest

	// 子账号uin
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// 策略id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *CheckUserPolicyAttachmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckUserPolicyAttachmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckUserPolicyAttachmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0表示未关联，1表示关联
		ExistAttachment *uint64 `json:"ExistAttachment,omitempty" name:"ExistAttachment"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckUserPolicyAttachmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckUserPolicyAttachmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConfirmCASProviderRequest struct {
	*tchttp.BaseRequest

	// 身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份提供商主账号
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *ConfirmCASProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConfirmCASProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConfirmCASProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmCASProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConfirmCASProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type DeleteCASProviderRequest struct {
	*tchttp.BaseRequest

	// 身份提供商名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 主账号uin
	OwnerUin []*uint64 `json:"OwnerUin,omitempty" name:"OwnerUin" list`
}

func (r *DeleteCASProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCASProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCASProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCASProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCASProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubAccountRequest struct {
	*tchttp.BaseRequest

	// 用户信息
	UserInfo []*GroupUidUinInfo `json:"UserInfo,omitempty" name:"UserInfo" list`

	// 语言
	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// Api调用
	FromAPI *uint64 `json:"FromAPI,omitempty" name:"FromAPI"`
}

func (r *DeleteSubAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubAccountResponse) FromJsonString(s string) error {
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

type DisableCASProviderRequest struct {
	*tchttp.BaseRequest

	// 身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份提供商主账号
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *DisableCASProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCASProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableCASProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCASProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCASProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupUidUinInfo struct {

	// 子用户Uid
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 子用户Uin
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 用户组ID 如果没有任何组传递-1,传入指定组id表示将用户从组删除
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

type ListSubAccountsRequest struct {
	*tchttp.BaseRequest

	// 查询关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 子用户Uid
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 过滤类型
	IsFilter *int64 `json:"IsFilter,omitempty" name:"IsFilter"`

	// 过滤的用户组
	FilterGroups []*uint64 `json:"FilterGroups,omitempty" name:"FilterGroups" list`

	// 过滤的策略ID
	FilterPolicyId *uint64 `json:"FilterPolicyId,omitempty" name:"FilterPolicyId"`

	// 项目ID
	ProjectId []*uint64 `json:"ProjectId,omitempty" name:"ProjectId" list`

	// 用户类型
	UserType []*string `json:"UserType,omitempty" name:"UserType" list`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 起始条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排查的用户
	ExcludeUids []*uint64 `json:"ExcludeUids,omitempty" name:"ExcludeUids" list`
}

func (r *ListSubAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSubAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListSubAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 用户信息
		UserInfo []*SubAccountFilter `json:"UserInfo,omitempty" name:"UserInfo" list`

		// 主帐号信息
		OwnerInfo []*OwnerInfo `json:"OwnerInfo,omitempty" name:"OwnerInfo" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSubAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSubAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersForPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 页码
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页行数
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListUsersForPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersForPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersForPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 子账号列表
		List []*UserList `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersForPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersForPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OwnerInfo struct {

	// 主帐号Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 校验状态
	CheckStatus *uint64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
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

type SubAccountFilter struct {

	// 子用户Uid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 用户Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是都允许登录
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanLogin *uint64 `json:"CanLogin,omitempty" name:"CanLogin"`

	// 电话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 电话号码是否验证
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 邮箱是否验证
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`

	// 用户类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *int64 `json:"UserType,omitempty" name:"UserType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否消息接收人
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`

	// 是否需要重置密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`

	// 是否允许控制台登录
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 微信公众号关注状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxzsStatus *int64 `json:"WxzsStatus,omitempty" name:"WxzsStatus"`

	// 权限类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermType []*string `json:"PermType,omitempty" name:"PermType" list`
}

type UpdateCASProviderRequest struct {
	*tchttp.BaseRequest

	// 身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份提供商主账号
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 身份提供商url地址根目录
	CasRoot *string `json:"CasRoot,omitempty" name:"CasRoot"`

	// 身份提供商登录地址
	CasLoginUrl *string `json:"CasLoginUrl,omitempty" name:"CasLoginUrl"`

	// 身份提供商验证ticket地址
	CasValidateUrl *string `json:"CasValidateUrl,omitempty" name:"CasValidateUrl"`
}

func (r *UpdateCASProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCASProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCASProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCASProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCASProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UserList struct {

	// 子账号名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子账号uin
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}
