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

type ApprovalRequest struct {
	*tchttp.BaseRequest

	// 无
	ApplyNo *string `json:"ApplyNo,omitempty" name:"ApplyNo"`

	// 无
	Opinion *int64 `json:"Opinion,omitempty" name:"Opinion"`

	// 无
	Approver *string `json:"Approver,omitempty" name:"Approver"`
}

func (r *ApprovalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApprovalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApprovalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApprovalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApprovalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type HostInfo struct {

	// 无
	LanIP *string `json:"LanIP,omitempty" name:"LanIP"`

	// 无
	Sn *string `json:"Sn,omitempty" name:"Sn"`

	// 无
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 无
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 无
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 无
	HostId *uint64 `json:"HostId,omitempty" name:"HostId"`

	// 无
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// 无
	LanMask *string `json:"LanMask,omitempty" name:"LanMask"`

	// 无
	LanGateway *string `json:"LanGateway,omitempty" name:"LanGateway"`

	// 无
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 无
	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`

	// 无
	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

type ModuleWorkflowRequest struct {
	*tchttp.BaseRequest

	// 无
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 无
	SetId *uint64 `json:"SetId,omitempty" name:"SetId"`

	// 无
	ModuleId *uint64 `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *ModuleWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModuleWorkflowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModuleWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModuleWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModuleWorkflowResponse) FromJsonString(s string) error {
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

type QueryHostRequest struct {
	*tchttp.BaseRequest

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *QueryHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryHostRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryHostResponse) FromJsonString(s string) error {
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

type QueryResDeliveryTaskRequest struct {
	*tchttp.BaseRequest

	// 详情Id
	DetailId *uint64 `json:"DetailId,omitempty" name:"DetailId"`
}

func (r *QueryResDeliveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryResDeliveryTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryResDeliveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryResDeliveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryResDeliveryTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryResDetailRequest struct {
	*tchttp.BaseRequest

	// 无
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 无
	ApplyNo *string `json:"ApplyNo,omitempty" name:"ApplyNo"`

	// 无
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *QueryResDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryResDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryResDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryResDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryResDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResDeliveryHistoryRequest struct {
	*tchttp.BaseRequest

	// 无
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 无
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 无
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *ResDeliveryHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResDeliveryHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResDeliveryHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResDeliveryHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResDeliveryHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResDeliveryRequest struct {
	*tchttp.BaseRequest

	// 无
	IsReinstallOs *int64 `json:"IsReinstallOs,omitempty" name:"IsReinstallOs"`

	// 无
	IsStandardization *int64 `json:"IsStandardization,omitempty" name:"IsStandardization"`

	// 无
	HostList []*HostInfo `json:"HostList,omitempty" name:"HostList" list`

	// 无
	Level *string `json:"Level,omitempty" name:"Level"`

	// 无
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 无
	HostUsername *string `json:"HostUsername,omitempty" name:"HostUsername"`

	// 无
	HostPassword *string `json:"HostPassword,omitempty" name:"HostPassword"`

	// 无
	VmOsInfo *string `json:"VmOsInfo,omitempty" name:"VmOsInfo"`

	// 无
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 无
	BondInfo *string `json:"BondInfo,omitempty" name:"BondInfo"`

	// 无
	RaidInfo *string `json:"RaidInfo,omitempty" name:"RaidInfo"`
}

func (r *ResDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResDeliveryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResDeliveryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStandConfigureRequest struct {
	*tchttp.BaseRequest

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

type VMOSListRequest struct {
	*tchttp.BaseRequest
}

func (r *VMOSListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VMOSListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VMOSListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VMOSListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VMOSListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
