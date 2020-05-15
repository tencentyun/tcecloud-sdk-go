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

package v2018109

import (
    "encoding/json"

    tchttp "github.com/tencentyun/tcecloud-sdk-go/tcecloud/common/http"
)

type OperDocCategoryParams struct {

	// action
	Action *string `json:"Action,omitempty" name:"Action"`

	// docTypeId
	DocTypeId *int64 `json:"DocTypeId,omitempty" name:"DocTypeId"`

	// categoryId
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// titleEn
	TitleEn *string `json:"TitleEn,omitempty" name:"TitleEn"`

	// titleCn
	TitleCn *string `json:"TitleCn,omitempty" name:"TitleCn"`

	// weight
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// imageUrl
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

type OperDocCategoryRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *OperDocCategoryParams `json:"Params,omitempty" name:"Params"`
}

func (r *OperDocCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocCategoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocCategoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperDocCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocCategoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocMenuLockParams struct {

	// menuId
	MenuId *int64 `json:"MenuId,omitempty" name:"MenuId"`

	// reviewId
	ReviewId *int64 `json:"ReviewId,omitempty" name:"ReviewId"`

	// action
	Action *string `json:"Action,omitempty" name:"Action"`
}

type OperDocMenuLockRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *OperDocMenuLockParams `json:"Params,omitempty" name:"Params"`

	// ProductId
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *OperDocMenuLockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocMenuLockRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocMenuLockResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperDocMenuLockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocMenuLockResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocMenuParams struct {

	// action
	Action *string `json:"Action,omitempty" name:"Action"`

	// productId
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// pid
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// titleCn
	TitleCn *string `json:"TitleCn,omitempty" name:"TitleCn"`

	// weight
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// menuId
	MenuId *int64 `json:"MenuId,omitempty" name:"MenuId"`

	// reviewId
	ReviewId *int64 `json:"ReviewId,omitempty" name:"ReviewId"`

	// type
	Type *string `json:"Type,omitempty" name:"Type"`

	// contentCn
	ContentCn *string `json:"ContentCn,omitempty" name:"ContentCn"`

	// keywordCn
	KeywordCn []*string `json:"KeywordCn,omitempty" name:"KeywordCn" list`
}

type OperDocMenuRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *OperDocMenuParams `json:"Params,omitempty" name:"Params"`

	// ProductId
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *OperDocMenuRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocMenuRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocMenuResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperDocMenuResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocMenuResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocMenuStatusRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *OperDocMenuLockParams `json:"Params,omitempty" name:"Params"`
}

func (r *OperDocMenuStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocMenuStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocMenuStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperDocMenuStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocMenuStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocProductDefaultPageParams struct {

	// ProductId
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// PageType
	PageType *string `json:"PageType,omitempty" name:"PageType"`

	// PageContent
	PageContent *string `json:"PageContent,omitempty" name:"PageContent"`
}

type OperDocProductDefaultPageRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *OperDocProductDefaultPageParams `json:"Params,omitempty" name:"Params"`
}

func (r *OperDocProductDefaultPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocProductDefaultPageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocProductDefaultPageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperDocProductDefaultPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocProductDefaultPageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocProductParams struct {

	// action
	Action *string `json:"Action,omitempty" name:"Action"`

	// productId
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// categoryId
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// titleEn
	TitleEn *string `json:"TitleEn,omitempty" name:"TitleEn"`

	// TitleCn
	TitleCn *string `json:"TitleCn,omitempty" name:"TitleCn"`

	// weight
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type OperDocProductRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *OperDocProductParams `json:"Params,omitempty" name:"Params"`
}

func (r *OperDocProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperDocProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocReviewParams struct {

	// reviewIdList
	ReviewIdList []*int64 `json:"ReviewIdList,omitempty" name:"ReviewIdList" list`

	// reviewPass
	ReviewPass *int64 `json:"ReviewPass,omitempty" name:"ReviewPass"`

	// reviewComment
	ReviewComment *string `json:"ReviewComment,omitempty" name:"ReviewComment"`
}

type OperDocReviewRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *OperDocReviewParams `json:"Params,omitempty" name:"Params"`
}

func (r *OperDocReviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocReviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocReviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperDocReviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocReviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocTypeParams struct {

	// action
	Action *string `json:"Action,omitempty" name:"Action"`

	// docTypeId
	DocTypeId *uint64 `json:"DocTypeId,omitempty" name:"DocTypeId"`

	// docType
	DocType *string `json:"DocType,omitempty" name:"DocType"`
}

type OperDocTypeRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *OperDocTypeParams `json:"Params,omitempty" name:"Params"`
}

func (r *OperDocTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperDocTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperDocTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OperDocTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocCategoryListRequest struct {
	*tchttp.BaseRequest

	// method
	Method *string `json:"Method,omitempty" name:"Method"`

	// userId
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *QueryDocCategoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocCategoryListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocCategoryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocCategoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocCategoryListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocCategoryParams struct {

	// filters
	Filters *QueryDocCategoryParamsFilters `json:"Filters,omitempty" name:"Filters"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type QueryDocCategoryParamsFilters struct {

	// docTypeId
	DocTypeId *int64 `json:"DocTypeId,omitempty" name:"DocTypeId"`

	// categoryId
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type QueryDocCategoryRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *QueryDocCategoryParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryDocCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocCategoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocCategoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocCategoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocMemuListParams struct {

	// productId
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// isAll
	IsAll *int64 `json:"IsAll,omitempty" name:"IsAll"`
}

type QueryDocMenuListRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *QueryDocMemuListParams `json:"Params,omitempty" name:"Params"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *QueryDocMenuListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocMenuListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocMenuListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocMenuListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocMenuListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocMenuParams struct {

	// MenuId
	MenuId *int64 `json:"MenuId,omitempty" name:"MenuId"`

	// ReviewId
	ReviewId *int64 `json:"ReviewId,omitempty" name:"ReviewId"`
}

type QueryDocMenuRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *QueryDocMenuParams `json:"Params,omitempty" name:"Params"`

	// ProductId
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *QueryDocMenuRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocMenuRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocMenuResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocMenuResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocMenuResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocProductDefaultPageParams struct {

	// ProductId
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`
}

type QueryDocProductDefaultPageRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *QueryDocProductDefaultPageParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryDocProductDefaultPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocProductDefaultPageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocProductDefaultPageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocProductDefaultPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocProductDefaultPageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocProductParams struct {

	// filters
	Filters *QueryDocProductParamsFilters `json:"Filters,omitempty" name:"Filters"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type QueryDocProductParamsFilters struct {

	// categoryId
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// productId
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type QueryDocProductRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *QueryDocProductParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryDocProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocReviewListParams struct {

	// filters
	Filters *QueryDocReviewListParamsFilters `json:"Filters,omitempty" name:"Filters"`

	// sortType
	SortType *string `json:"SortType,omitempty" name:"SortType"`

	// offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type QueryDocReviewListParamsFilters struct {

	// applyTime
	ApplyTime *QueryDocReviewListParamsFiltersApplyTime `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// titleCn
	TitleCn *string `json:"TitleCn,omitempty" name:"TitleCn"`

	// applier
	Applier *string `json:"Applier,omitempty" name:"Applier"`

	// reviewStatusId
	ReviewStatusId *int64 `json:"ReviewStatusId,omitempty" name:"ReviewStatusId"`
}

type QueryDocReviewListParamsFiltersApplyTime struct {

	// begin
	Begin *string `json:"Begin,omitempty" name:"Begin"`

	// end
	End *string `json:"End,omitempty" name:"End"`
}

type QueryDocReviewListRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *QueryDocReviewListParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryDocReviewListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocReviewListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocReviewListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocReviewListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocReviewListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocReviewParams struct {

	// reviewId
	ReviewId *int64 `json:"ReviewId,omitempty" name:"ReviewId"`
}

type QueryDocReviewRequest struct {
	*tchttp.BaseRequest

	// Method
	Method *string `json:"Method,omitempty" name:"Method"`

	// UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Params
	Params *QueryDocReviewParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryDocReviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocReviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocReviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocReviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocReviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocTypeParams struct {

	// Filters
	Filters []*QueryDocTypeParamsFilters `json:"Filters,omitempty" name:"Filters" list`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type QueryDocTypeParamsFilters struct {

	// docTypeId
	DocTypeId *uint64 `json:"DocTypeId,omitempty" name:"DocTypeId"`

	// keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type QueryDocTypeRequest struct {
	*tchttp.BaseRequest

	// method
	Method *string `json:"Method,omitempty" name:"Method"`

	// userId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// params
	Params *QueryDocTypeParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryDocTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryDocTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryDocTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
