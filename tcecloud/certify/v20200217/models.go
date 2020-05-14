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

package v20200217

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CompanyCertificateApplyParams struct {

	// 企业全称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 真实姓名
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 联系电话
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 省份
	AddrProvince *string `json:"AddrProvince,omitempty" name:"AddrProvince"`

	// 城市
	AddrCity *string `json:"AddrCity,omitempty" name:"AddrCity"`

	// 地址信息
	AddrDetail *string `json:"AddrDetail,omitempty" name:"AddrDetail"`

	// 统一社会信用代码
	CreditCode *string `json:"CreditCode,omitempty" name:"CreditCode"`

	// 营业执照照片的存储路径
	LicensePhoto *string `json:"LicensePhoto,omitempty" name:"LicensePhoto"`
}

type CompanyCertificateApplyRequest struct {
	*tchttp.BaseRequest

	// input
	Params *CompanyCertificateApplyParams `json:"Params,omitempty" name:"Params"`
}

func (r *CompanyCertificateApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompanyCertificateApplyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompanyCertificateApplyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompanyCertificateApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompanyCertificateApplyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsoleCertificateDetailParams struct {

	// 审核记录related_uin
	Account *string `json:"Account,omitempty" name:"Account"`

	// 审核记录related_appid
	Appid *string `json:"Appid,omitempty" name:"Appid"`
}

type GetConsoleCertificateDetailRequest struct {
	*tchttp.BaseRequest

	// input
	Params *GetConsoleCertificateDetailParams `json:"Params,omitempty" name:"Params"`
}

func (r *GetConsoleCertificateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsoleCertificateDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsoleCertificateDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsoleCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsoleCertificateDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsoleCertificateStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetConsoleCertificateStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsoleCertificateStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsoleCertificateStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsoleCertificateStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsoleCertificateStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsoleCurrentCertRequest struct {
	*tchttp.BaseRequest
}

func (r *GetConsoleCurrentCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsoleCurrentCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsoleCurrentCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsoleCurrentCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsoleCurrentCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PersonCertificateApplyParams struct {

	// 身份证号码
	CardId *string `json:"CardId,omitempty" name:"CardId"`

	// 真实姓名
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 联系电话
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 省份
	AddrProvince *string `json:"AddrProvince,omitempty" name:"AddrProvince"`

	// 城市
	AddrCity *string `json:"AddrCity,omitempty" name:"AddrCity"`

	// 地址信息
	AddrDetail *string `json:"AddrDetail,omitempty" name:"AddrDetail"`

	// 身份证正面照片的存储路径
	CardPhoto1 *string `json:"CardPhoto1,omitempty" name:"CardPhoto1"`

	// 身份证反面照片的存储路径
	CardPhoto2 *string `json:"CardPhoto2,omitempty" name:"CardPhoto2"`
}

type PersonCertificateApplyRequest struct {
	*tchttp.BaseRequest

	// input
	Params *PersonCertificateApplyParams `json:"Params,omitempty" name:"Params"`
}

func (r *PersonCertificateApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PersonCertificateApplyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PersonCertificateApplyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PersonCertificateApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PersonCertificateApplyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryProvinceAndCityInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryProvinceAndCityInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryProvinceAndCityInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryProvinceAndCityInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProvinceAndCityInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryProvinceAndCityInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
