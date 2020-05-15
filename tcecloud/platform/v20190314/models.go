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

package v20190314

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type AccountList struct {

	// 是否激活
	Active *int64 `json:"Active,omitempty" name:"Active"`

	// 新增日期
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// APPID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 绑定状态
	BindStatus *int64 `json:"BindStatus,omitempty" name:"BindStatus"`

	// 上次登录时间
	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`

	// 邮件
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 修改日期
	ModTimestamp *string `json:"ModTimestamp,omitempty" name:"ModTimestamp"`

	// 昵称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 输主uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 标记
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 子账号数量
	SubCount *int64 `json:"SubCount,omitempty" name:"SubCount"`

	// uid
	Uid *int64 `json:"Uid,omitempty" name:"Uid"`

	// uin
	Uin *int64 `json:"Uin,omitempty" name:"Uin"`

	// uintype
	UinType *int64 `json:"UinType,omitempty" name:"UinType"`

	// username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 秘钥
	IdKeys []*IdKeys `json:"IdKeys,omitempty" name:"IdKeys" list`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type CancelWeChatNoticeRequest struct {
	*tchttp.BaseRequest

	// 用户Uin
	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`

	// 用户登录账号
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户昵称
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *CancelWeChatNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelWeChatNoticeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelWeChatNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelWeChatNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelWeChatNoticeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomAccountRequest struct {
	*tchttp.BaseRequest

	// 分页
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 是否获取加密key数据
	GetSecretField *bool `json:"GetSecretField,omitempty" name:"GetSecretField"`

	// 排序key
	SortKey *string `json:"SortKey,omitempty" name:"SortKey"`

	// 排序顺序
	SortTurn *int64 `json:"SortTurn,omitempty" name:"SortTurn"`
}

func (r *GetCustomAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户列表
		AccountList []*AccountList `json:"AccountList,omitempty" name:"AccountList" list`

		// 用户数量
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomSubAccountRequest struct {
	*tchttp.BaseRequest

	// 主账号uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 是否获取加密数据
	GetSecretField *bool `json:"GetSecretField,omitempty" name:"GetSecretField"`
}

func (r *GetCustomSubAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomSubAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomSubAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户列表
		SubAccountList []*AccountList `json:"SubAccountList,omitempty" name:"SubAccountList" list`

		// 用户数量
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomSubAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomSubAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSwitchInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSwitchInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSwitchInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSwitchInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询是否开启微信通道，值为False或True
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSwitchInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSwitchInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetVerifyQRCodeRequest struct {
	*tchttp.BaseRequest

	// 用户Uin
	User_Uin *uint64 `json:"User_Uin,omitempty" name:"User_Uin"`

	// 用户登录账号
	User_Id *string `json:"User_Id,omitempty" name:"User_Id"`

	// 用户昵称
	User_Name *string `json:"User_Name,omitempty" name:"User_Name"`
}

func (r *GetVerifyQRCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetVerifyQRCodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetVerifyQRCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证二维码url
		QRCodeUrl *string `json:"QRCodeUrl,omitempty" name:"QRCodeUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetVerifyQRCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetVerifyQRCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IdKeys struct {

	// createTime
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 秘钥id
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// 秘钥key
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// source
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 启用状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyWeChatRequest struct {
	*tchttp.BaseRequest

	// 用户Uin
	User_Uin *uint64 `json:"User_Uin,omitempty" name:"User_Uin"`

	// 用户登录账号
	User_Id *string `json:"User_Id,omitempty" name:"User_Id"`

	// 用户昵称
	User_Name *string `json:"User_Name,omitempty" name:"User_Name"`
}

func (r *ModifyWeChatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWeChatRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWeChatResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWeChatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWeChatResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NoAcceptNoticeRequest struct {
	*tchttp.BaseRequest

	// 用户Uin
	User_Uin *uint64 `json:"User_Uin,omitempty" name:"User_Uin"`
}

func (r *NoAcceptNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *NoAcceptNoticeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NoAcceptNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NoAcceptNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *NoAcceptNoticeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryCustomAccountRequest struct {
	*tchttp.BaseRequest

	// 分页
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 租户端登录名
	Account *string `json:"Account,omitempty" name:"Account"`

	// 租户手机号码
	PhoneNumber *int64 `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 租户昵称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 租户端appid
	CustomAppId *int64 `json:"CustomAppId,omitempty" name:"CustomAppId"`

	// 账号uin
	CustomUin *int64 `json:"CustomUin,omitempty" name:"CustomUin"`

	// 排序key
	SortKey *string `json:"SortKey,omitempty" name:"SortKey"`

	// 顺序或者逆序，0：顺序，1：逆序
	SortTurn *int64 `json:"SortTurn,omitempty" name:"SortTurn"`

	// 开始时间
	StartLastLoginTime *int64 `json:"StartLastLoginTime,omitempty" name:"StartLastLoginTime"`

	// 结束时间
	EndLastLoginTime *int64 `json:"EndLastLoginTime,omitempty" name:"EndLastLoginTime"`
}

func (r *QueryCustomAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryCustomAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryCustomAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户数量
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 用户列表
		AccountList []*AccountList `json:"AccountList,omitempty" name:"AccountList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustomAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryCustomAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendBindingEmailRequest struct {
	*tchttp.BaseRequest

	// 用户Uin
	User_Uin *uint64 `json:"User_Uin,omitempty" name:"User_Uin"`

	// 用户登录账号
	User_Id *string `json:"User_Id,omitempty" name:"User_Id"`

	// 用户昵称
	User_Name *string `json:"User_Name,omitempty" name:"User_Name"`
}

func (r *SendBindingEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendBindingEmailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendBindingEmailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendBindingEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendBindingEmailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
