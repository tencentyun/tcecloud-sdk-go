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

package v20200408

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type GetParams struct {

	// uin
	Uin *int64 `json:"Uin,omitempty" name:"Uin"`

	// tcloud/ocloud
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// language/timezone
	Key *string `json:"Key,omitempty" name:"Key"`
}

type GetRequest struct {
	*tchttp.BaseRequest

	// input
	Params *GetParams `json:"Params,omitempty" name:"Params"`
}

func (r *GetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutRequest struct {
	*tchttp.BaseRequest

	// input
	Params *SetParams `json:"Params,omitempty" name:"Params"`
}

func (r *PutRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetParams struct {

	// uin（子） uin+platform会作为存储维度确定一个账号
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// ocloud/tcloud
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// language/timezone
	Key *string `json:"Key,omitempty" name:"Key"`

	// value
	Value *string `json:"Value,omitempty" name:"Value"`
}
