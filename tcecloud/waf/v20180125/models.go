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

package v20180125

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type AddCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies" list`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 动作类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`

	// "clb-waf"或者"sparta-waf"
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
}

func (r *AddCustomRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddCustomRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCustomRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddCustomRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAttackDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// 域名，所有域名填写all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 下载任务名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 拦截状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 自定义策略ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 攻击者IP
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// 攻击类型
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *CreateAttackDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAttackDownloadTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAttackDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Flow *string `json:"Flow,omitempty" name:"Flow"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAttackDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAttackDownloadTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 记录id
	Flow *string `json:"Flow,omitempty" name:"Flow"`
}

func (r *DeleteDownloadRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDownloadRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDownloadRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDownloadRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSessionRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// clb-waf 或者 sprta-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DeleteSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesPagingInfo struct {

	// 当前页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 当前页的最大数据条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCustomRulesRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 分页参数
	Paging *DescribeCustomRulesPagingInfo `json:"Paging,omitempty" name:"Paging"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 过滤参数：动作类型：0放行，1阻断，2人机识别，3观察，4重定向
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 过滤参数：规则名称过滤条件
	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeCustomRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则详情
		RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitempty" name:"RuleList" list`

		// 规则条数
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesRspRuleListItem struct {

	// 动作类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跳过的策略
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 重定向地址
	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`

	// 策略ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies" list`
}

type ModifyCustomRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *ModifyCustomRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomRuleStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomRuleStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Strategy struct {

	// 匹配字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 逻辑符号
	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`

	// 匹配内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 匹配参数
	Arg *string `json:"Arg,omitempty" name:"Arg"`
}
