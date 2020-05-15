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

package v20180930

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type AllBizTopoRequest struct {
	*tchttp.BaseRequest
}

func (r *AllBizTopoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllBizTopoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllBizTopoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllBizTopoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllBizTopoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindPkgRequest struct {
	*tchttp.BaseRequest

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	InstId *uint64 `json:"InstId,omitempty" name:"InstId"`

	// 无
	PkgInfo []*PkgInfo `json:"PkgInfo,omitempty" name:"PkgInfo" list`
}

func (r *BindPkgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindPkgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindPkgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindPkgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindPkgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BizTopoRequest struct {
	*tchttp.BaseRequest

	// 无
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
}

func (r *BizTopoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BizTopoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BizTopoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BizTopoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BizTopoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStandConfigureRequest struct {
	*tchttp.BaseRequest

	// 标准化作业ID
	WorkflowTemplateId *uint64 `json:"WorkflowTemplateId,omitempty" name:"WorkflowTemplateId"`

	// 标准化作业名称
	WorkflowTemplateName *string `json:"WorkflowTemplateName,omitempty" name:"WorkflowTemplateName"`

	// 业务ID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 申请人
	UpdateUser *string `json:"UpdateUser,omitempty" name:"UpdateUser"`

	// 集群ID
	SetId *int64 `json:"SetId,omitempty" name:"SetId"`

	// 模块ID
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 对象类型
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`
}

func (r *CreateStandConfigureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStandConfigureRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStandConfigureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStandConfigureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStandConfigureResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopoConfigRequest struct {
	*tchttp.BaseRequest

	// 层级类型
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 层级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 配置字符串
	TopoConfigs *string `json:"TopoConfigs,omitempty" name:"TopoConfigs"`
}

func (r *CreateTopoConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopoConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopoConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopoConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopoConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteStandConfigureRequest struct {
	*tchttp.BaseRequest

	// 标准化配置ID
	StandardizationConfigId *int64 `json:"StandardizationConfigId,omitempty" name:"StandardizationConfigId"`

	// 对象类型
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`
}

func (r *DeleteStandConfigureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteStandConfigureRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteStandConfigureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStandConfigureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteStandConfigureResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopoConfigRequest struct {
	*tchttp.BaseRequest

	// 层级类型
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 层级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 删除数组
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids" list`
}

func (r *DeleteTopoConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopoConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopoConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopoConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopoConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployOverviewRequest struct {
	*tchttp.BaseRequest

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	InstId *int64 `json:"InstId,omitempty" name:"InstId"`

	// 无
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *DeployOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnumInfoRequest struct {
	*tchttp.BaseRequest

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`
}

func (r *EnumInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnumInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnumInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnumInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnumInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	Values []*string `json:"Values,omitempty" name:"Values" list`

	// 查询的操作类别，如==
	Op *string `json:"Op,omitempty" name:"Op"`
}

type GetModulePkgInfoRequest struct {
	*tchttp.BaseRequest

	// CMDB业务
	Level *string `json:"Level,omitempty" name:"Level"`

	// Level的最后一层级，本接口只支持module
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询返回数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetModulePkgInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetModulePkgInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetModulePkgInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModulePkgInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetModulePkgInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetModuleSchedulesListRequest struct {
	*tchttp.BaseRequest

	// CMDB业务
	Level *string `json:"Level,omitempty" name:"Level"`

	// Level的最后一层级，本接口只支持module
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询返回数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetModuleSchedulesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetModuleSchedulesListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetModuleSchedulesListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleSchedulesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetModuleSchedulesListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPkgInsInfoRequest struct {
	*tchttp.BaseRequest

	// CMDB业务
	Level *string `json:"Level,omitempty" name:"Level"`

	// Level的最后一层级，本接口只支持module
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询返回数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetPkgInsInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPkgInsInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPkgInsInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPkgInsInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPkgInsInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSchedulesTaskInfoRequest struct {
	*tchttp.BaseRequest

	// CMDB业务
	Level *string `json:"Level,omitempty" name:"Level"`

	// Level的最后一层级，本接口只支持module
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询返回数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetSchedulesTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSchedulesTaskInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSchedulesTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSchedulesTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSchedulesTaskInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTaskInfoRequest struct {
	*tchttp.BaseRequest

	// CMDB业务
	Level *string `json:"Level,omitempty" name:"Level"`

	// Level的最后一层级，本接口只支持module
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询返回数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTaskInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTaskInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowTemplateDetailRequest struct {
	*tchttp.BaseRequest

	// 作业编排ID
	WorkflowTemplateId *int64 `json:"WorkflowTemplateId,omitempty" name:"WorkflowTemplateId"`
}

func (r *GetWorkflowTemplateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetWorkflowTemplateDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowTemplateDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkflowTemplateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetWorkflowTemplateDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowTemplatesRequest struct {
	*tchttp.BaseRequest

	// 查询起始值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询返回数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetWorkflowTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetWorkflowTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkflowTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetWorkflowTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStandConfigureRequest struct {
	*tchttp.BaseRequest

	// CMDB业务
	Level *string `json:"Level,omitempty" name:"Level"`

	// 编排和包信息
	WorkflowPkg *WorkflowPkg `json:"WorkflowPkg,omitempty" name:"WorkflowPkg"`

	// 层级类型
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`
}

func (r *ModifyStandConfigureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStandConfigureRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStandConfigureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStandConfigureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStandConfigureResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NodeHostRequest struct {
	*tchttp.BaseRequest

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	InstId *int64 `json:"InstId,omitempty" name:"InstId"`

	// 无
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 无
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *NodeHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *NodeHostRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NodeHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NodeHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *NodeHostResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NodeInfoRequest struct {
	*tchttp.BaseRequest

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	InstId *int64 `json:"InstId,omitempty" name:"InstId"`
}

func (r *NodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *NodeInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *NodeInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PkgFile struct {

	// 包类型
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 包ID
	PkgId *int64 `json:"PkgId,omitempty" name:"PkgId"`

	// 包版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 这个包操作的机器id
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids" list`

	// 是否启动包
	IsAutoStart *bool `json:"IsAutoStart,omitempty" name:"IsAutoStart"`
}

type PkgInfo struct {

	// 无
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 无
	PkgId *uint64 `json:"PkgId,omitempty" name:"PkgId"`

	// 无
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 无
	Version *string `json:"Version,omitempty" name:"Version"`
}

type QueryAuditRequest struct {
	*tchttp.BaseRequest

	// 无
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 无
	InstId *uint64 `json:"InstId,omitempty" name:"InstId"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	LanIP *string `json:"LanIP,omitempty" name:"LanIP"`

	// 无
	OpType *int64 `json:"OpType,omitempty" name:"OpType"`

	// 无
	OpTarget *string `json:"OpTarget,omitempty" name:"OpTarget"`

	// 无
	OpDesc *string `json:"OpDesc,omitempty" name:"OpDesc"`

	// 无
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 无
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *QueryAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAuditRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAuditResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryBizRequest struct {
	*tchttp.BaseRequest

	// 无
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *QueryBizRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryBizRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryBizResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBizResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryBizResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryModuleRequest struct {
	*tchttp.BaseRequest

	// 无
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *QueryModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryPkgListRequest struct {
	*tchttp.BaseRequest

	// 查询起始值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询返回数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询包名，支持模糊
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 查询的包类型package or config or script
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`
}

func (r *QueryPkgListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryPkgListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryPkgListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPkgListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryPkgListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryPkgVersionListRequest struct {
	*tchttp.BaseRequest

	// 包ID
	PkgId *int64 `json:"PkgId,omitempty" name:"PkgId"`
}

func (r *QueryPkgVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryPkgVersionListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryPkgVersionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPkgVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryPkgVersionListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySetRequest struct {
	*tchttp.BaseRequest

	// 无
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *QuerySetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuerySetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryTopoConfigRequest struct {
	*tchttp.BaseRequest

	// 层级类型
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 层级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤值
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *QueryTopoConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryTopoConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryTopoConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTopoConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryTopoConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryUuidPkgLastVersionRequest struct {
	*tchttp.BaseRequest

	// 查询包ID
	PkgId *int64 `json:"PkgId,omitempty" name:"PkgId"`

	// 查询机器uuid
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids" list`
}

func (r *QueryUuidPkgLastVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryUuidPkgLastVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryUuidPkgLastVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryUuidPkgLastVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryUuidPkgLastVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunModuleWorkflowRequest struct {
	*tchttp.BaseRequest

	// CMDB业务
	Level *string `json:"Level,omitempty" name:"Level"`

	// 操作名
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// Level的最后一层级，本接口只支持module
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 包列表,优先于模块绑定的包
	Files []*PkgFile `json:"Files,omitempty" name:"Files" list`

	// 执行原因类别
	ReasonType *string `json:"ReasonType,omitempty" name:"ReasonType"`

	// 执行原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *RunModuleWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunModuleWorkflowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunModuleWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunModuleWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunModuleWorkflowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunStandardizationTaskRequest struct {
	*tchttp.BaseRequest

	// uuid/level列表
	UuidLevelList []*UuidLevel `json:"UuidLevelList,omitempty" name:"UuidLevelList" list`

	// 执行的类型 RunConf/RunSpecific/Pass
	OpType *string `json:"OpType,omitempty" name:"OpType"`

	// 编排和包信息
	WorkflowPkg *WorkflowPkg `json:"WorkflowPkg,omitempty" name:"WorkflowPkg"`
}

func (r *RunStandardizationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunStandardizationTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunStandardizationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunStandardizationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunStandardizationTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStandConfigureRequest struct {
	*tchttp.BaseRequest

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *SearchStandConfigureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStandConfigureRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStandConfigureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStandConfigureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStandConfigureResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStandardizationHistoryRequest struct {
	*tchttp.BaseRequest

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *SearchStandardizationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStandardizationHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStandardizationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStandardizationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStandardizationHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStandardizationTasksRequest struct {
	*tchttp.BaseRequest

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *SearchStandardizationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStandardizationTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStandardizationTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStandardizationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStandardizationTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindPkgRequest struct {
	*tchttp.BaseRequest

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	InstId *uint64 `json:"InstId,omitempty" name:"InstId"`

	// 无
	PkgInfo []*PkgInfo `json:"PkgInfo,omitempty" name:"PkgInfo" list`
}

func (r *UnbindPkgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindPkgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindPkgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindPkgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindPkgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateHostRequest struct {
	*tchttp.BaseRequest

	// 无
	HostId []*uint64 `json:"HostId,omitempty" name:"HostId" list`

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	BusinessStatus *string `json:"BusinessStatus,omitempty" name:"BusinessStatus"`

	// 无
	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`

	// 无
	Maintainer *string `json:"Maintainer,omitempty" name:"Maintainer"`

	// 无
	BakMaintainer *string `json:"BakMaintainer,omitempty" name:"BakMaintainer"`
}

func (r *UpdateHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateHostRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateHostResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNodeRequest struct {
	*tchttp.BaseRequest

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	InstId *int64 `json:"InstId,omitempty" name:"InstId"`

	// 无
	Name *string `json:"Name,omitempty" name:"Name"`

	// 无
	Maintainer *string `json:"Maintainer,omitempty" name:"Maintainer"`

	// 无
	BakMaintainer *string `json:"BakMaintainer,omitempty" name:"BakMaintainer"`

	// 无
	Developer *string `json:"Developer,omitempty" name:"Developer"`

	// 无
	Productor *string `json:"Productor,omitempty" name:"Productor"`

	// 无
	Tester *string `json:"Tester,omitempty" name:"Tester"`

	// 无
	FollowUser *string `json:"FollowUser,omitempty" name:"FollowUser"`

	// 无
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// 无
	ModuleLevel *string `json:"ModuleLevel,omitempty" name:"ModuleLevel"`
}

func (r *UpdateNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateTopoConfigRequest struct {
	*tchttp.BaseRequest

	// 层级类型
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 层级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 主键
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 修改值
	TopoConfigs *string `json:"TopoConfigs,omitempty" name:"TopoConfigs"`
}

func (r *UpdateTopoConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateTopoConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateTopoConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTopoConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateTopoConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UuidLevel struct {

	// level业务层级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 机器uuid/sn
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type WorkflowPkg struct {

	// 编排任务ID
	WorkflowId *int64 `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 编排任务code
	WtCode *string `json:"WtCode,omitempty" name:"WtCode"`

	// 包信息
	PkgInfoList []*PkgInfo `json:"PkgInfoList,omitempty" name:"PkgInfoList" list`
}
