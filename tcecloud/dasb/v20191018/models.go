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

package v20191018

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type DescribeDasbImageIdsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDasbImageIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbImageIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbImageIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基础镜像ID
		BaseImageId *string `json:"BaseImageId,omitempty" name:"BaseImageId"`

		// AI镜像ID
		AiImageId *string `json:"AiImageId,omitempty" name:"AiImageId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbImageIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbImageIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
