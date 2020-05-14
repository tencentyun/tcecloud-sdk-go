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

package v20170312

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeAlarmMetricsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmMetricsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAlarmMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest

	// 业务命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeBaseMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseMetricsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询得到的指标描述列表
		MetricSet []*MetricObject `json:"MetricSet,omitempty" name:"MetricSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaseMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseMetricsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 命名空间，每个云产品会有一个命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 实例对象的维度组合
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 监控统计周期。默认为取值为300，单位为s
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 起始时间，如 2018-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认为当前时间。 endTime不能小于startTime
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonitorDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控指标
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 数据点起始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 数据点结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 数据统计周期
		Period *uint64 `json:"Period,omitempty" name:"Period"`

		// 监控数据列表
		DataPoints []*PointsObject `json:"DataPoints,omitempty" name:"DataPoints" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetMonitorDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MetricObject struct {

	// 命名空间，每个云产品会有一个命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标使用的单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 指标支持的统计周期，单位是秒，如60、300
	Period []*int64 `json:"Period,omitempty" name:"Period" list`
}

type PointsObject struct {

	// 监控实例的维度组合
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions" list`

	// 监控数据点数组，每个点的时间跨度为一个Period值
	Points []*float64 `json:"Points,omitempty" name:"Points" list`
}
