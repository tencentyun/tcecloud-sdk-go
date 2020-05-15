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

package v20190411

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type AddTopicRequest struct {
	*tchttp.BaseRequest

	// 订阅名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// phone/email/web_cmgt/phone_email/phone_email_cmgtsite
	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`

	// user_id
	Users []*int64 `json:"Users,omitempty" name:"Users" list`

	// groupid
	Groups []*int64 `json:"Groups,omitempty" name:"Groups" list`

	// 描述
	TopicDesc *string `json:"TopicDesc,omitempty" name:"TopicDesc"`
}

func (r *AddTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelTopicRequest struct {
	*tchttp.BaseRequest

	// id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DelTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelUserCmgtSiteMsgRequest struct {
	*tchttp.BaseRequest

	// id数组
	Ids []*int64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DelUserCmgtSiteMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelUserCmgtSiteMsgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelUserCmgtSiteMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelUserCmgtSiteMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DelUserCmgtSiteMsgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filters struct {

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Values
	Values []*string `json:"Values,omitempty" name:"Values" list`

	// Op
	Op *string `json:"Op,omitempty" name:"Op"`
}

type GetTopicRequest struct {
	*tchttp.BaseRequest

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件
	Filters []*Filters `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserCmgtSiteMsgRequest struct {
	*tchttp.BaseRequest

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件
	Filters []*Filters `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetUserCmgtSiteMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserCmgtSiteMsgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserCmgtSiteMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserCmgtSiteMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserCmgtSiteMsgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReadUserCmgtSiteMsgRequest struct {
	*tchttp.BaseRequest

	// 用户Uin
	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`

	// 已读的id数组
	Ids []*int64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *ReadUserCmgtSiteMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReadUserCmgtSiteMsgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReadUserCmgtSiteMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReadUserCmgtSiteMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReadUserCmgtSiteMsgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateTopicRequest struct {
	*tchttp.BaseRequest

	// ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 订阅名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// phone/email/web_cmgt/phone_email/phone_email_cmgtsite
	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`

	// user_id
	Users []*int64 `json:"Users,omitempty" name:"Users" list`

	// groupid
	Groups []*int64 `json:"Groups,omitempty" name:"Groups" list`

	// 描述
	TopicDesc *string `json:"TopicDesc,omitempty" name:"TopicDesc"`
}

func (r *UpdateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
