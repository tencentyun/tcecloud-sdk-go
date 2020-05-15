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

package v20190606

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type ArgusReportRequest struct {
	*tchttp.BaseRequest

	// 指定上报的数据格式，其中lp代表Line Protocol，json代表Json数据协议，推荐使用Line Protocol。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当Type指定为lp时使用，数据格式说明请参考：http://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/ ，其中measurement部分填写上报的命名空间名称，时间戳精确到纳秒。
	LpData []*string `json:"LpData,omitempty" name:"LpData" list`

	// 当Type指定为json时使用，支持单个请求上报多个命名空间数据。
	JsonData []*ReportJson `json:"JsonData,omitempty" name:"JsonData" list`
}

func (r *ArgusReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArgusReportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ArgusReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回描述信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArgusReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArgusReportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportDataRequest struct {
	*tchttp.BaseRequest

	// 指定上报的数据格式，其中lp代表Line Protocol，json代表Json数据协议，推荐使用Line Protocol。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当Type指定为lp时使用，数据格式说明请参考：http://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/ ，其中measurement部分填写上报的命名空间名称，时间戳精确到纳秒。
	LpData []*string `json:"LpData,omitempty" name:"LpData" list`

	// 当Type指定为json时使用，支持单个请求上报多个命名空间数据。
	JsonData []*ReportJson `json:"JsonData,omitempty" name:"JsonData" list`
}

func (r *ReportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回描述信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReportDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReportDimension struct {

	// 维度名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ReportJson struct {

	// 命名空间名称
	Ns *string `json:"Ns,omitempty" name:"Ns"`

	// 维度组合信息
	Dimensions []*ReportDimension `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 指标信息
	Metrics []*ReportMetric `json:"Metrics,omitempty" name:"Metrics" list`

	// 时间戳（精确到毫秒）
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type ReportMetric struct {

	// 指标名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}
