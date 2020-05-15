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

type CreateModuleNewVersionCreatePkgRequest struct {
	*tchttp.BaseRequest

	// 临时目录ID
	TmpDirId *string `json:"TmpDirId,omitempty" name:"TmpDirId"`

	// 模块ID
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 模块版本号
	ModuleVersion *string `json:"ModuleVersion,omitempty" name:"ModuleVersion"`

	// 运行用户
	RunUser *string `json:"RunUser,omitempty" name:"RunUser"`

	// 状态检查脚本
	CheckShell *string `json:"CheckShell,omitempty" name:"CheckShell"`

	// 启动脚本
	StartShell *string `json:"StartShell,omitempty" name:"StartShell"`

	// 停止脚本
	StopShell *string `json:"StopShell,omitempty" name:"StopShell"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 版本描述
	ModuleVersionDesc *string `json:"ModuleVersionDesc,omitempty" name:"ModuleVersionDesc"`
}

func (r *CreateModuleNewVersionCreatePkgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleNewVersionCreatePkgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleNewVersionCreatePkgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateModuleNewVersionCreatePkgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleNewVersionCreatePkgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleNewVersionGetTmpdirRequest struct {
	*tchttp.BaseRequest

	// 模块ID
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *CreateModuleNewVersionGetTmpdirRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleNewVersionGetTmpdirRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleNewVersionGetTmpdirResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateModuleNewVersionGetTmpdirResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleNewVersionGetTmpdirResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 管理者
	Admin *string `json:"Admin,omitempty" name:"Admin"`

	// 模块英文名
	ModuleNameEn *string `json:"ModuleNameEn,omitempty" name:"ModuleNameEn"`

	// 模块说明
	ModuleDesc *string `json:"ModuleDesc,omitempty" name:"ModuleDesc"`
}

func (r *CreateModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModulesRequest struct {
	*tchttp.BaseRequest

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大限制数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 是否精准查询标志
	IsEqualSearch *bool `json:"IsEqualSearch,omitempty" name:"IsEqualSearch"`

	// 查询条件限制
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeModulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type GetEnumInfoRequest struct {
	*tchttp.BaseRequest

	// 对象ID
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 对象属性列表
	ObjAttrs []*string `json:"ObjAttrs,omitempty" name:"ObjAttrs" list`
}

func (r *GetEnumInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEnumInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEnumInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEnumInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEnumInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstallModulesRequest struct {
	*tchttp.BaseRequest

	// uuid列表
	Sn []*string `json:"Sn,omitempty" name:"Sn" list`

	// 模块列表
	ModuleList []*ModuleList `json:"ModuleList,omitempty" name:"ModuleList" list`
}

func (r *InstallModulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstallModulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstallModulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallModulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstallModulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleDescRequest struct {
	*tchttp.BaseRequest

	// 模块ID
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块描述
	ModuleDesc *string `json:"ModuleDesc,omitempty" name:"ModuleDesc"`
}

func (r *ModifyModuleDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleDescRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleDescResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleDescResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModuleList struct {

	// 模块id
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 版本号
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

type UninstallModulesRequest struct {
	*tchttp.BaseRequest

	// 主机的uuid
	UuidList []*string `json:"UuidList,omitempty" name:"UuidList" list`

	// 模块列表
	ModuleList []*int64 `json:"ModuleList,omitempty" name:"ModuleList" list`
}

func (r *UninstallModulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UninstallModulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UninstallModulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallModulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UninstallModulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
