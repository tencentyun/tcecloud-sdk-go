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

package v20180420

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 下单参数
	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例id
		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例列表
		InstanceList []*string `json:"InstanceList,omitempty" name:"InstanceList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductTypesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品规格列表
		ProductTypes []*ProductType `json:"ProductTypes,omitempty" name:"ProductTypes" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductTypesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSaleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSaleInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规格
		ProductTypes []*ProductType `json:"ProductTypes,omitempty" name:"ProductTypes" list`

		// 时长
		TimeSpan []*TimeSpan `json:"TimeSpan,omitempty" name:"TimeSpan" list`

		// 售卖方式,0:计量,1:预付费,2:后付费
		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

		// 镜像ID
		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

		// 应用状态
		AppStatus *string `json:"AppStatus,omitempty" name:"AppStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSaleInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProductType struct {

	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// PID
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 规格名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 规格代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// Cpu核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 系统盘大小
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 数据盘大小
	DataDiskSize *uint64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`
}

type TimeSpan struct {

	// 名称
	TimeName *string `json:"TimeName,omitempty" name:"TimeName"`

	// 月数
	TimeMonth *uint64 `json:"TimeMonth,omitempty" name:"TimeMonth"`
}
