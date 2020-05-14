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

package v20190325

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountDetail struct {

	// 敏感操作标识
	ActionFlag *ActionLoginFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`

	// 是否允许控制台登录
	ConsoleLogin *string `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 登录保护
	LoginFlag *ActionLoginFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`

	// 是否需要重置密码
	NeedResetPassword *string `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`

	// 用户密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 使用Api
	UseApi *string `json:"UseApi,omitempty" name:"UseApi"`
}

type ActionLoginFlag struct {

	// 电话
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 软Token
	Stoken *string `json:"Stoken,omitempty" name:"Stoken"`

	// 硬Token
	Token *string `json:"Token,omitempty" name:"Token"`
}

type AddSubAccountRequest struct {
	*tchttp.BaseRequest

	// 用户信息
	UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo" list`

	// 语言
	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// 是否在白名单
	InWhiteList *bool `json:"InWhiteList,omitempty" name:"InWhiteList"`

	// 从Api创建
	FromAPI *uint64 `json:"FromAPI,omitempty" name:"FromAPI"`
}

func (r *AddSubAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSubAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddSubAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账号Uin
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// 子用户uin列表
		List []*uint64 `json:"List,omitempty" name:"List" list`

		// 子用户Uid
		Uids []*uint64 `json:"Uids,omitempty" name:"Uids" list`

		// 子用户详情
		SubAccounts []*SubAccounts `json:"SubAccounts,omitempty" name:"SubAccounts" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSubAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddSubAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeMailPasswordRequest struct {
	*tchttp.BaseRequest

	// 旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 新密码
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 语言类型，zh或en
	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *ChangeMailPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeMailPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeMailPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeMailPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeMailPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeSubAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 新密码
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`
}

func (r *ChangeSubAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeSubAccountPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeSubAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeSubAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeSubAccountPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMultiFactorParasRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMultiFactorParasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMultiFactorParasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMultiFactorParasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMultiFactorParasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMultiFactorParasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserInfoRequest struct {
	*tchttp.BaseRequest

	// 是否为主账号
	IsOwner *uint64 `json:"IsOwner,omitempty" name:"IsOwner"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 审核状态
	CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 邮箱是否审核通过
	MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`

	// 线下审核状态
	OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`

	// 首次购买带外网IP的cvm设备的时间
	WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`

	// 外网是否受限
	WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`

	// 返回的字段
	Fields *string `json:"Fields,omitempty" name:"Fields"`
}

func (r *GetUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// UIN
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// 用户名
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 昵称
		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

		// 主账号UIN
		OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

		// 是否为主账号
		IsOwner *int64 `json:"IsOwner,omitempty" name:"IsOwner"`

		// 资料是否审核通过
		CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`

		// 是否实名认证：0未认证，1已认证
		IsAuthenticate *int64 `json:"IsAuthenticate,omitempty" name:"IsAuthenticate"`

		// 邮箱是否审核通过
		MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`

		// 邮箱
		Mail *string `json:"Mail,omitempty" name:"Mail"`

		// 手机号码
		PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

		// 用户指引标识位
		GuideBit *int64 `json:"GuideBit,omitempty" name:"GuideBit"`

		// 用户首次购买带外网IP的cvm设备的时间
		WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`

		// 标识外网是否受限
		WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`

		// 创建时间
		AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

		// 修改时间
		ModTimestamp *string `json:"ModTimestamp,omitempty" name:"ModTimestamp"`

		// 线下审核状态
		OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`

		// 业务信息
		BizInfo *string `json:"BizInfo,omitempty" name:"BizInfo"`

		// 来源平台
		SrcPlatform *string `json:"SrcPlatform,omitempty" name:"SrcPlatform"`

		// 是否测试用户
		IsTestUser *int64 `json:"IsTestUser,omitempty" name:"IsTestUser"`

		// 客户来源
		ClientFrom *string `json:"ClientFrom,omitempty" name:"ClientFrom"`

		// register refer
		Referer *string `json:"Referer,omitempty" name:"Referer"`

		// 是否注册成功
		IsRegSucc *bool `json:"IsRegSucc,omitempty" name:"IsRegSucc"`

		// 用户属性集合
		Attributes *int64 `json:"Attributes,omitempty" name:"Attributes"`

		// 否导入了即时通白名单
		Isprotect *int64 `json:"Isprotect,omitempty" name:"Isprotect"`

		// 用户指引
		IsSeeGuidelines *int64 `json:"IsSeeGuidelines,omitempty" name:"IsSeeGuidelines"`

		// 是否接收推广信息
		IsAcceptProMsg *int64 `json:"IsAcceptProMsg,omitempty" name:"IsAcceptProMsg"`

		// 部署模块
		DeployName *string `json:"DeployName,omitempty" name:"DeployName"`

		// 账号列表
		AccountList *string `json:"AccountList,omitempty" name:"AccountList"`

		// 账号类型
		AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

		// 默认开发商
		DefaultOwner *int64 `json:"DefaultOwner,omitempty" name:"DefaultOwner"`

		// 国家代码
		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

		// 邮箱验证
		MailVerify *int64 `json:"MailVerify,omitempty" name:"MailVerify"`

		// 接收信息语言
		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`

		// 地域
		Area *string `json:"Area,omitempty" name:"Area"`

		// 是否需要完善信息
		Needinfo *int64 `json:"Needinfo,omitempty" name:"Needinfo"`

		// 绑定账号
		Account *string `json:"Account,omitempty" name:"Account"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubAccounts struct {

	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 秘钥Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// 秘钥Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type UserInfo struct {

	// 子账号类型
	CanLogin *string `json:"CanLogin,omitempty" name:"CanLogin"`

	// 区号
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 详情
	Detail *AccountDetail `json:"Detail,omitempty" name:"Detail"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 电话号码
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 系统类型
	SyStemType *string `json:"SyStemType,omitempty" name:"SyStemType"`
}
