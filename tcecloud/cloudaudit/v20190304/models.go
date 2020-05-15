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

package v20190304

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type Attr struct {

	// 查询属性
	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`

	// 查询内容
	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

type LookupEventsRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 查询属性细节
	LookupAttributes []*Attr `json:"LookupAttributes,omitempty" name:"LookupAttributes" list`

	// 每次最大查询数量
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 查询分页
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 查询用户uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *LookupEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LookupEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LookupEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LookupEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LookupEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
