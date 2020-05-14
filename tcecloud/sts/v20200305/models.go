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

package v20200305

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Credentials struct {

	// token
	Token *string `json:"Token,omitempty" name:"Token"`

	// 临时证书密钥ID
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时证书密钥Key
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
}

type GetConsoleSessionTokenRequest struct {
	*tchttp.BaseRequest

	// sessionId的md5值
	Sid *string `json:"Sid,omitempty" name:"Sid"`

	// 客户端UserAgent
	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`

	// 客户端设备(pc或者mobile)
	ClientDevice *string `json:"ClientDevice,omitempty" name:"ClientDevice"`

	// 分配给控制台的secretId
	ConSecretId *string `json:"ConSecretId,omitempty" name:"ConSecretId"`

	// 生成签名时使用
	Noncet *uint64 `json:"Noncet,omitempty" name:"Noncet"`

	// 根据签名的计算方法生成的签名
	Signaturet *string `json:"Signaturet,omitempty" name:"Signaturet"`

	// 过期时间
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 客户端IP
	ClientIPt *string `json:"ClientIPt,omitempty" name:"ClientIPt"`

	// 操作接口，获取临时证书后要操作的api名称
	Actiont *string `json:"Actiont,omitempty" name:"Actiont"`

	// 时间戳(计算签名时使用)
	Timestampt *uint64 `json:"Timestampt,omitempty" name:"Timestampt"`

	// Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 操作资源
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 控制台信息，cos v4 控制台填写"COSV4"，其他情况不需要传入该参数
	ConsoleInfo *string `json:"ConsoleInfo,omitempty" name:"ConsoleInfo"`
}

func (r *GetConsoleSessionTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsoleSessionTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsoleSessionTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// token
		Token *string `json:"Token,omitempty" name:"Token"`

		// 临时证书秘钥ID
		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

		// 临时证书秘钥Key
		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

		// 过期时间
		SecretExpire *uint64 `json:"SecretExpire,omitempty" name:"SecretExpire"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsoleSessionTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsoleSessionTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFederationTokenRequest struct {
	*tchttp.BaseRequest

	// 您可以自定义调用方英文名称，由字母组成。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 策略描述
	// 注意：
	// 1、policy 需要做 urlencode（如果通过 GET 方法请求云 API，发送请求前，所有参数都需要按照云 API 规范再 urlencode 一次）。
	// 2、策略语法参照 CAM 策略语法。
	// 3、策略中不能包含 principal 元素。
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// 指定临时证书的有效期，单位：秒，默认1800秒，最长可设定有效期为7200秒。
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`

	// 证书类型（0：长证书；1：短证书）。
	SecretType *uint64 `json:"SecretType,omitempty" name:"SecretType"`
}

func (r *GetFederationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFederationTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFederationTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 临时证书
		Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

		// 临时证书有效的时间，返回 Unix 时间戳，精确到秒
		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFederationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFederationTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
