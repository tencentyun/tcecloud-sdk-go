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

package v20190305

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BeatTips struct {

	// 命中关键词
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 恶意类型，同type说明
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 恶意等级：0：无恶意 1~3：数字越小恶意等级越高
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type BspAudioRecognitionRequest struct {
	*tchttp.BaseRequest

	// 音频的URL地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 音频名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *BspAudioRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspAudioRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspAudioRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别信息
		Data *TextData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspAudioRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspAudioRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspCloseServiceRequest struct {
	*tchttp.BaseRequest

	// 内容安全服务ID标识
	ServicesId *int64 `json:"ServicesId,omitempty" name:"ServicesId"`

	// 客户uin
	UIN *uint64 `json:"UIN,omitempty" name:"UIN"`

	// 客户appid
	APPID *uint64 `json:"APPID,omitempty" name:"APPID"`
}

func (r *BspCloseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspCloseServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspCloseServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内容安全服务关闭信息
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspCloseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspCloseServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspImageRecognitionRequest struct {
	*tchttp.BaseRequest

	// 图片名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 图片的 URL 地址  与FileContent二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 对图片内容 Base64 位编码值 与FileUrl二选一
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

func (r *BspImageRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspImageRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspImageRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果详情
		Data *ImageData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspImageRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspImageRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspOpenAfterPayServiceRequest struct {
	*tchttp.BaseRequest

	// pid.
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// type.
	Type *string `json:"Type,omitempty" name:"Type"`

	// subtype.
	SubType *string `json:"SubType,omitempty" name:"SubType"`
}

func (r *BspOpenAfterPayServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspOpenAfterPayServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspOpenAfterPayServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspOpenAfterPayServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspOpenAfterPayServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspOpenServiceRequest struct {
	*tchttp.BaseRequest

	// 客户uin
	UIN *uint64 `json:"UIN,omitempty" name:"UIN"`

	// 客户appid
	APPID *uint64 `json:"APPID,omitempty" name:"APPID"`

	// 服务id
	ServicesId *int64 `json:"ServicesId,omitempty" name:"ServicesId"`
}

func (r *BspOpenServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspOpenServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspOpenServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务开通详情
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspOpenServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspOpenServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSearchAllServiceNumRequest struct {
	*tchttp.BaseRequest

	// 当天日期
	Today *string `json:"Today,omitempty" name:"Today"`

	// 昨日日期
	YesterDay *string `json:"YesterDay,omitempty" name:"YesterDay"`

	// 上周当天日期
	WeekFdate *string `json:"WeekFdate,omitempty" name:"WeekFdate"`
}

func (r *BspSearchAllServiceNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSearchAllServiceNumRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSearchAllServiceNumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回top50数据详情
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspSearchAllServiceNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSearchAllServiceNumResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSearchAllServiceOpensRequest struct {
	*tchttp.BaseRequest
}

func (r *BspSearchAllServiceOpensRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSearchAllServiceOpensRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSearchAllServiceOpensResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 客户服务开通详情信息
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspSearchAllServiceOpensResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSearchAllServiceOpensResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSearchServiceOpensRequest struct {
	*tchttp.BaseRequest

	// 内容安全服务ID标识
	ServicesId *int64 `json:"ServicesId,omitempty" name:"ServicesId"`
}

func (r *BspSearchServiceOpensRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSearchServiceOpensRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSearchServiceOpensResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务开通记录详情信息
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspSearchServiceOpensResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSearchServiceOpensResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSereachServiceAllRequest struct {
	*tchttp.BaseRequest

	// 客户uin
	UIN *uint64 `json:"UIN,omitempty" name:"UIN"`
}

func (r *BspSereachServiceAllRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSereachServiceAllRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSereachServiceAllResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户开通服务信息
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspSereachServiceAllResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSereachServiceAllResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspStatisticsContentRequest struct {
	*tchttp.BaseRequest

	// 开始日期
	StartFdate *string `json:"StartFdate,omitempty" name:"StartFdate"`

	// 结束日期
	EndFdate *string `json:"EndFdate,omitempty" name:"EndFdate"`

	// 内容安全服务id标识
	ServicesId *int64 `json:"ServicesId,omitempty" name:"ServicesId"`
}

func (r *BspStatisticsContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspStatisticsContentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspStatisticsContentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时间区域内的用量详情和恶意占比信息
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspStatisticsContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspStatisticsContentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspStatisticsNumDistributionRequest struct {
	*tchttp.BaseRequest

	// 开始日期
	StartFdate *string `json:"StartFdate,omitempty" name:"StartFdate"`

	// 结束日期
	EndFdate *string `json:"EndFdate,omitempty" name:"EndFdate"`
}

func (r *BspStatisticsNumDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspStatisticsNumDistributionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspStatisticsNumDistributionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回客户数据明细
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspStatisticsNumDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspStatisticsNumDistributionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSumUserRequestQuantityRequest struct {
	*tchttp.BaseRequest

	// 当前日期
	StartFdate *string `json:"StartFdate,omitempty" name:"StartFdate"`

	// 客户uin
	UIN *uint64 `json:"UIN,omitempty" name:"UIN"`

	// 结束日期
	EndFdate *string `json:"EndFdate,omitempty" name:"EndFdate"`
}

func (r *BspSumUserRequestQuantityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSumUserRequestQuantityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspSumUserRequestQuantityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求量明细数据
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspSumUserRequestQuantityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspSumUserRequestQuantityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspTextRecognitionRequest struct {
	*tchttp.BaseRequest

	// 识别的文本内容
	MessageContent *string `json:"MessageContent,omitempty" name:"MessageContent"`
}

func (r *BspTextRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspTextRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BspTextRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果
		Data *TextData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BspTextRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BspTextRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageData struct {

	// 公共错误码 0：表示成功 -1003表示图片处理超时；-1004表示文字处理超时；-1005表示内部处理异常；
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 恶意类型： 100:正常 20001:政治 20002:色情 20006:涉毒违法 24001:暴恐 21000:其他
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 图片识别结果详情
	Data []*ImageDetail `json:"Data,omitempty" name:"Data" list`
}

type ImageDetail struct {

	// 识别类型 PornDetect FaceRecog IllegalDetect TerroristDetect OCRDetect SimDetect
	Category *string `json:"Category,omitempty" name:"Category"`

	// 是否恶意：0否1是
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 0-100 分值越高越恶意，其中OCRDetect与SimDetect无此字段
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 特征中文描述，其中OCRDetect与SimDetect无此字段
	Label *string `json:"Label,omitempty" name:"Label"`

	// 打击原因，如命中的关键词(多个以分号分隔)  其中OCRDetect有此字段，其余识别类型无此字段
	BeatTips *string `json:"BeatTips,omitempty" name:"BeatTips"`

	// 恶意等级：0：无恶意 1~3：数字越小恶意等级越高 其中OCRDetect有此字段，其余识别类型无此字段
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type TextData struct {

	// 公共错误码 0：表示成功 -1003表示图片处理超时；-1004表示文字处理超时；-1005表示内部处理异常；
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 恶意类型： 100:正常 20001:政治 20002:色情 20006:涉毒违法 24001:暴恐 21000:其他
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 打击原因
	BeatTips []*BeatTips `json:"BeatTips,omitempty" name:"BeatTips" list`
}
