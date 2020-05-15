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

package v20180808

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type DescribeLogSearchRequest struct {
	*tchttp.BaseRequest

	// 日志开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 精确查询，支持apiid/reqid搜索
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 根据上次返回的ConText，获取后续的内容，最多可获取10000条
	ConText *string `json:"ConText,omitempty" name:"ConText"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 模糊查询，根据关键字检索日志
	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeLogSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSearchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取更多检索结果的游标，值为""表示无后续结果
		ConText *string `json:"ConText,omitempty" name:"ConText"`

		// 由0或多条日志组成，每条日志格式如下：
	// '[$app_id][$env_name][$service_id][$http_host][$api_id][$uri][$scheme][rsp_st:$status][ups_st:$upstream_status]'
	// '[cip:$remote_addr][uip:$upstream_addr][vip:$server_addr][rsp_len:$bytes_sent][req_len:$request_length]'
	// '[req_t:$request_time][ups_rsp_t:$upstream_response_time][ups_conn_t:$upstream_connect_time][ups_head_t:$upstream_header_time]’
	// '[err_msg:$err_msg][tcp_rtt:$tcpinfo_rtt][$pid][$time_local][req_id:$request_id]';
	// 
	// 说明：
	// app_id： 用户 ID。
	// env_name：环境名称。
	// service_id： 服务 ID。
	// http_host： 域名。
	// api_id： API 的 ID。
	// uri：请求的路径。
	// scheme： HTTP/HTTPS 协议。
	// rsp_st： 请求响应状态码。
	// ups_st： 后端业务服务器的响应状态码（如果请求透传到后端，改变量不为空。如果请求在 APIGW 就被拦截了，那么该变量显示为 -）。
	// cip： 客户端 IP。
	// uip： 后端业务服务（upstream）的 IP。
	// vip： 请求访问的 VIP。
	// rsp_len： 响应长度。
	// req_len： 请求长度。
	// req_t： 请求响应的总时间。
	// ups_rsp_t： 后端响应的总时间（apigw 建立连接到接收到后端响应的时间）。
	// ups_conn_t: 与后端业务服务器连接建立成功时间。
	// ups_head_t：后端响应的头部到达时间。
	// err_msg： 错误信息。
	// tcp_rtt： 客户端 TCP 连接信息，RTT（Round Trip Time）由三部分组成：链路的传播时间（propagation delay）、末端系统的处理时间、路由器缓存中的排队和处理时间（queuing delay）。
	// req_id：请求id。
		LogSet []*string `json:"LogSet,omitempty" name:"LogSet" list`

		// 单次搜索返回的日志条数，TotalCount <= Limit
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSearchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ModifyApiIncrementRequest struct {
	*tchttp.BaseRequest

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 接口ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 需要修改的API auth类型(可选择OAUTH-授权API)
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// oauth接口需要修改的公钥值
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// oauth接口重定向地址
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

func (r *ModifyApiIncrementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiIncrementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiIncrementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiIncrementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
