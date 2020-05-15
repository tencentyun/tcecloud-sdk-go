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

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

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
