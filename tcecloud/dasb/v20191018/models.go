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

package v20191018

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateSCRTaskRequest struct {
	*tchttp.BaseRequest

	// 文件数据来源。0：文件 URL；1：文件数据（post body）
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 回调 URL，用户自行搭建的用于接收识别结果的服务器地址， 长度小于2048字节。如果用户使用回调方式获取识别结果，需提交该参数；如果用户使用轮询方式获取识别结果，则无需提交该参数。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 文件的URL地址，需要公网可下载。长度小于10240字节，当 SourceType 值为 0 时须填写该字段，为 1 时不需要填写。注意：请确保文件大小在10M之内，否则可能识别失败。请保证文件的下载速度，否则可能下载失败。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 文件数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。文件数据要小于10MB。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度，当 SourceType 值为1时必须填写，为0可不写（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *uint64 `json:"DataLen,omitempty" name:"DataLen"`

	// 识别规则列表，如果在调用时敏感内容识别时，不进行单独的规则设置，自动生效默认规则；如果进行了单独的规则设置，那么将生效单独设置的规则。
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 透传字段，透传简单信息。
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

func (r *CreateSCRTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSCRTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSCRTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSCRTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSCRTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DasbCvmConfig struct {

	// 主机名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 主机核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 外网带宽
	NetBand *uint64 `json:"NetBand,omitempty" name:"NetBand"`

	// 系统盘大小
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 数据盘大小
	DataDiskSize *uint64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// 购买月份
	MonthSpan *uint64 `json:"MonthSpan,omitempty" name:"MonthSpan"`

	// 所属商品模型ID
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名
	Region *string `json:"Region,omitempty" name:"Region"`

	// 使用的镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type DasbInstance struct {

	// 实例key
	PincKey *string `json:"PincKey,omitempty" name:"PincKey"`

	// 实例名称
	PsysName *string `json:"PsysName,omitempty" name:"PsysName"`

	// 公网ip
	MpublicIp *string `json:"MpublicIp,omitempty" name:"MpublicIp"`

	// 产品型号
	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`

	// 购买时间
	PaddTime *uint64 `json:"PaddTime,omitempty" name:"PaddTime"`

	// 内网IP
	MprivateIp *string `json:"MprivateIp,omitempty" name:"MprivateIp"`

	// 实例状态
	PsysStatus *uint64 `json:"PsysStatus,omitempty" name:"PsysStatus"`

	// 实例开始时间
	PstartTime *uint64 `json:"PstartTime,omitempty" name:"PstartTime"`

	// 实例结束时间
	PendTime *uint64 `json:"PendTime,omitempty" name:"PendTime"`

	// 实例所在地域ID
	Pregion *uint64 `json:"Pregion,omitempty" name:"Pregion"`

	// 实例所在地区ID
	PzoneId *uint64 `json:"PzoneId,omitempty" name:"PzoneId"`

	// CVM实例名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmInstanceNames []*string `json:"CvmInstanceNames,omitempty" name:"CvmInstanceNames" list`

	// CVM实例ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitempty" name:"CvmInstanceIds" list`
}

type DasbInstanceLog struct {

	// 日志ID
	LogId *uint64 `json:"LogId,omitempty" name:"LogId"`

	// 日志内容
	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`

	// 日志记录时间
	LogTime *uint64 `json:"LogTime,omitempty" name:"LogTime"`
}

type DescribeDasbCvmConfigRequest struct {
	*tchttp.BaseRequest

	// 商品定价ID
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

func (r *DescribeDasbCvmConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbCvmConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbCvmConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Dasb主机配置信息
		ConfigSet []*DasbCvmConfig `json:"ConfigSet,omitempty" name:"ConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbCvmConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbCvmConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbCvmInstanceRequest struct {
	*tchttp.BaseRequest

	// DASB产品实例唯一KEY
	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
}

func (r *DescribeDasbCvmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbCvmInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbCvmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Dasb主机信息
		InstanceSet []*DasbCvmConfig `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbCvmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbCvmInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbImageIdsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDasbImageIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbImageIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbImageIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基础镜像ID
		BaseImageId *string `json:"BaseImageId,omitempty" name:"BaseImageId"`

		// AI镜像ID
		AiImageId *string `json:"AiImageId,omitempty" name:"AiImageId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbImageIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbImageIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbInstanceAlarmRequest struct {
	*tchttp.BaseRequest

	// 实例唯一Key
	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
}

func (r *DescribeDasbInstanceAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbInstanceAlarmRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbInstanceAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志列表
		LogSet []*DasbInstanceLog `json:"LogSet,omitempty" name:"LogSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbInstanceAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbInstanceAlarmResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDasbServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例列表
		ServiceList []*DasbInstance `json:"ServiceList,omitempty" name:"ServiceList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDasbServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInitialPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInitialPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInitialPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInitialPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户初始化密码
		UserPwd *string `json:"UserPwd,omitempty" name:"UserPwd"`

		// 查看日志记录
		PwdLog []*PwdLog `json:"PwdLog,omitempty" name:"PwdLog" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInitialPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInitialPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDasbUserInfosRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDasbUserInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDasbUserInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDasbUserInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否已设置Secret，0：未设置；1：已设置
		IsSetSecret *int64 `json:"IsSetSecret,omitempty" name:"IsSetSecret"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDasbUserInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDasbUserInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDasbInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例唯一标识串
	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`

	// 配置内容
	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`
}

func (r *InitDasbInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDasbInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDasbInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDasbInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDasbInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PwdLog struct {

	// 用户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 订单关联标识
	IncKey *string `json:"IncKey,omitempty" name:"IncKey"`
}

type RunInstanceRequest struct {
	*tchttp.BaseRequest

	// 堡垒机基础配置信息
	BaseInfo *string `json:"BaseInfo,omitempty" name:"BaseInfo"`

	// CVM配置信息
	CvmInfo *string `json:"CvmInfo,omitempty" name:"CvmInfo"`

	// 数据盘配置信息
	CbsInfo *string `json:"CbsInfo,omitempty" name:"CbsInfo"`
}

func (r *RunInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetDasbSecretInfosRequest struct {
	*tchttp.BaseRequest

	// 密钥ID
	UserSecretId *string `json:"UserSecretId,omitempty" name:"UserSecretId"`

	// 密钥KEY
	UserSecretKey *string `json:"UserSecretKey,omitempty" name:"UserSecretKey"`
}

func (r *SetDasbSecretInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetDasbSecretInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetDasbSecretInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDasbSecretInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetDasbSecretInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
