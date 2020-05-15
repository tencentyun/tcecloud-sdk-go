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

package v20200107

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type GetStatDayRequest struct {
	*tchttp.BaseRequest

	// 存储桶所在 COS 地域
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 存储桶名称
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 统计信息日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 存储类型 1:标准存储 2:低频存储
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *GetStatDayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetStatDayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetStatDayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStatDayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetStatDayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
