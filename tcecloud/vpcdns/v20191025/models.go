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

package v20191025

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type BindVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// VPC信息
	VpcInfos []*VpcInfos `json:"VpcInfos,omitempty" name:"VpcInfos" list`
}

func (r *BindVpcDnsDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindVpcDnsDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindVpcDnsDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindVpcDnsDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindVpcDnsDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsDomainRemarkRequest struct {
	*tchttp.BaseRequest

	// 域名ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateVpcDnsDomainRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcDnsDomainRemarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsDomainRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsDomainRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcDnsDomainRemarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 标签数组
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateVpcDnsDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcDnsDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcDnsDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录类型
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 记录值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 权重
	Weight *string `json:"Weight,omitempty" name:"Weight"`

	// MX优先级
	Mx *uint64 `json:"Mx,omitempty" name:"Mx"`
}

func (r *CreateVpcDnsRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcDnsRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcDnsRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名ID，以逗号分隔
	DomainIds *string `json:"DomainIds,omitempty" name:"DomainIds"`
}

func (r *DeleteVpcDnsDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcDnsDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcDnsDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcDnsDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 记录ID，逗号分隔
	RecordIds *string `json:"RecordIds,omitempty" name:"RecordIds"`
}

func (r *DeleteVpcDnsRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcDnsRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcDnsRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcDnsRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsDomainListRequest struct {
	*tchttp.BaseRequest

	// 长度
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤
	Filters []*DomainListFilters `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeVpcDnsDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcDnsDomainListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsDomainListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcDnsDomainListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsRecordListRequest struct {
	*tchttp.BaseRequest

	// 域名ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 长度
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤
	Filters []*RecordListFilters `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeVpcDnsRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcDnsRecordListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsRecordListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcDnsRecordListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DomainListFilters struct {

	// 过滤类型
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ModifyVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 记录ID
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录类型
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 记录值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 权重
	Weight *string `json:"Weight,omitempty" name:"Weight"`

	// MX优先级
	Mx *uint64 `json:"Mx,omitempty" name:"Mx"`
}

func (r *ModifyVpcDnsRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcDnsRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcDnsRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcDnsRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecordListFilters struct {

	// 过滤类型
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Tag struct {

	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type VpcInfos struct {

	// VpcId
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// RegionId
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// UnVpcId
	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
}
