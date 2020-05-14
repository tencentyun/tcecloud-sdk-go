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

package v20200326

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActionDescSample struct {

	// 输出参数
	OutputInfo *string `json:"OutputInfo,omitempty" name:"OutputInfo"`

	// 错误码
	Errorcodes *string `json:"Errorcodes,omitempty" name:"Errorcodes"`

	// 接口描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 接口文档分类
	Category *string `json:"Category,omitempty" name:"Category"`
}

type ActionInfo struct {

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 接口版本号
	Version *string `json:"Version,omitempty" name:"Version"`

	// 接口名
	Action *string `json:"Action,omitempty" name:"Action"`

	// 入参
	InputInfo *string `json:"InputInfo,omitempty" name:"InputInfo"`

	// 系统参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemParam *string `json:"SystemParam,omitempty" name:"SystemParam"`

	// attribute_flag
	Flag *int64 `json:"Flag,omitempty" name:"Flag"`

	// 后端地址单地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`

	// 后端地址多地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`

	// 接口中文名
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 发布状态
	RecordStatus *int64 `json:"RecordStatus,omitempty" name:"RecordStatus"`

	// CAM资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CamResource *string `json:"CamResource,omitempty" name:"CamResource"`

	// 接口对外是否可见
	CliVisibility *int64 `json:"CliVisibility,omitempty" name:"CliVisibility"`

	// CAM条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	CamCondition *string `json:"CamCondition,omitempty" name:"CamCondition"`

	// 限频
	SecRate *int64 `json:"SecRate,omitempty" name:"SecRate"`

	// 签名和鉴权
	IsAuth *int64 `json:"IsAuth,omitempty" name:"IsAuth"`

	// 后端地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`

	// 全程票据校验
	WholeProcessToken *int64 `json:"WholeProcessToken,omitempty" name:"WholeProcessToken"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 创建人
	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`

	// 限频白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLev *string `json:"RateLev,omitempty" name:"RateLev"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateAt *string `json:"UpdateAt,omitempty" name:"UpdateAt"`
}

type DataQueryApiInfo struct {

	// 总个数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 当前页号
	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`

	// 最后一页号
	LastPage *int64 `json:"LastPage,omitempty" name:"LastPage"`

	// 每页大小
	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`

	// 数据起始序号
	From *int64 `json:"From,omitempty" name:"From"`

	// 数据终止序号
	To *int64 `json:"To,omitempty" name:"To"`

	// 返回的API数据数组
	DataInfo []*ResApiInfo `json:"DataInfo,omitempty" name:"DataInfo" list`
}

type FilterApiInfoParam struct {

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 接口名
	Action *string `json:"Action,omitempty" name:"Action"`

	// 接口中文名
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 修改人名
	OperatorName *string `json:"OperatorName,omitempty" name:"OperatorName"`

	// 线上后端地址
	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`

	// 更新起始时间
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// 更新截止时间
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// 按某个字段排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// desc或asc
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type GetModuleNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetModuleNamesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetModuleNamesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetModuleNamesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleNamesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetModuleNamesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryApiInfoFilterRequest struct {
	*tchttp.BaseRequest

	// 过滤分页查询字段
	FilterParam *FilterApiInfoParam `json:"FilterParam,omitempty" name:"FilterParam"`

	// 分页每页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页当前页号
	Page *int64 `json:"Page,omitempty" name:"Page"`
}

func (r *QueryApiInfoFilterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryApiInfoFilterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryApiInfoFilterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的API数据
		Data *DataQueryApiInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApiInfoFilterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryApiInfoFilterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResApiInfo struct {

	// 接口详细info
	ActionInfo *ActionInfo `json:"ActionInfo,omitempty" name:"ActionInfo"`

	// 接口精简描述信息
	ActionDesc *ActionDescSample `json:"ActionDesc,omitempty" name:"ActionDesc"`

	// 白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhiteList *string `json:"WhiteList,omitempty" name:"WhiteList"`

	// 接口超时时间
	TimeoutInMillisec *int64 `json:"TimeoutInMillisec,omitempty" name:"TimeoutInMillisec"`
}
