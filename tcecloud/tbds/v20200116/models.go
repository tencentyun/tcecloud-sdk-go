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

package v20200116

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type ClusterExistsRequest struct {
	*tchttp.BaseRequest

	// 集群名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ClusterExistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClusterExistsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClusterExistsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClusterExistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClusterExistsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Docker镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例所属可用区ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// CVM示例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 要申请创建的CVM实例个数
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 本地磁盘的个数
	LocalStorageDiskCount *uint64 `json:"LocalStorageDiskCount,omitempty" name:"LocalStorageDiskCount"`

	// 创建人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 系统盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘大小，单位：GB。默认值为 50
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 数据盘
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 数据盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：CVM实例配置。默认值为0，表示不购买数据盘。更多限制详见产品文档
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘是否随子机销毁。TRUE：子机销毁时，销毁数据盘，只支持按小时后付费云盘。FALSE：子机销毁时，保留数据盘。默认取值：TRUE
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}

type GetClustersRequest struct {
	*tchttp.BaseRequest

	// 从0开始的页面index
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 关键字，支持|分隔多个关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *GetClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClustersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClustersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProgressRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProgressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProgressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
