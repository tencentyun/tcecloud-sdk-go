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

package v20180416

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddGroupForChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 允许使用合约的组织ID列表
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList" list`
}

func (r *AddGroupForChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddGroupForChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddGroupForChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddGroupForChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddGroupForChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddGroupForChannelRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 操作所属的组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 加入通道的组织列表
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList" list`
}

func (r *AddGroupForChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddGroupForChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddGroupForChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddGroupForChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddGroupForChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateGrouptoMemberRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前组织ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 成员名称
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// 成员账号ID
	MemberUin *string `json:"MemberUin,omitempty" name:"MemberUin"`

	// 成员appid
	MemberAppid *string `json:"MemberAppid,omitempty" name:"MemberAppid"`

	// 待分配组织ID
	AllocateGroupId *string `json:"AllocateGroupId,omitempty" name:"AllocateGroupId"`
}

func (r *AllocateGrouptoMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateGrouptoMemberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateGrouptoMemberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateGrouptoMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateGrouptoMemberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyCertRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 证书类型-1:个人;2:企业
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// ca类型-0:“tbaas”;1：“cfca”
	CaType *uint64 `json:"CaType,omitempty" name:"CaType"`

	// 申请人
	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`

	// 用户标识列表
	UserIdentList []*string `json:"UserIdentList,omitempty" name:"UserIdentList" list`

	// 证件类型
	IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

	// 证件号码
	IdentityNum *string `json:"IdentityNum,omitempty" name:"IdentityNum"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 电话
	PhoneNo *string `json:"PhoneNo,omitempty" name:"PhoneNo"`

	// 地址信息
	Address *string `json:"Address,omitempty" name:"Address"`

	// 所属组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 签名算法
	SignAlg *string `json:"SignAlg,omitempty" name:"SignAlg"`

	// 密钥算法：0 ECC， 1 SM2
	KeyAlg *uint64 `json:"KeyAlg,omitempty" name:"KeyAlg"`

	// 密钥长度
	KeyLen *uint64 `json:"KeyLen,omitempty" name:"KeyLen"`

	// 有效时间
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// csr p10证书文件
	CsrData *string `json:"CsrData,omitempty" name:"CsrData"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

func (r *ApplyCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsynCheckChaincodeDevRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 异步事务ID
	TranId *string `json:"TranId,omitempty" name:"TranId"`
}

func (r *AsynCheckChaincodeDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsynCheckChaincodeDevRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsynCheckChaincodeDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步调用操作
		Action *string `json:"Action,omitempty" name:"Action"`

		// 返回码，0成功，1失败
		ResCode *uint64 `json:"ResCode,omitempty" name:"ResCode"`

		// 具体执行结果
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsynCheckChaincodeDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsynCheckChaincodeDevResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsynCompileChaincodeDevRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 智能合约工程名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *AsynCompileChaincodeDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsynCompileChaincodeDevRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsynCompileChaincodeDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步编译调用结果
		ResCode *uint64 `json:"ResCode,omitempty" name:"ResCode"`

		// 异步事务ID
		TranId *string `json:"TranId,omitempty" name:"TranId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsynCompileChaincodeDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsynCompileChaincodeDevResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuthData struct {

	// 成员类型:0个人，1企业
	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`

	// 成员名称:个人姓名或企业名称
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// 成员地域:0大陆、1港澳、2台湾、3外籍
	MemberArea *string `json:"MemberArea,omitempty" name:"MemberArea"`

	// 实名认证方式:0未知，1有效证件，2财富通，3银行卡，4微信，5手Q
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 企业认证类型:0营业执照号认证,1社会信用代码认证,2非盈利机构认证；type=1时才关注这个字段
	AuthenticateType *string `json:"AuthenticateType,omitempty" name:"AuthenticateType"`

	// 证件类型:身份证号/护照/回乡证/台胞证/营业执照号/社会信用代码
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 证件号码:身份证号/护照/回乡证/台胞证/营业执照号/社会信用代码
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 组织结构代码:0营业执照号认证,1社会信用代码认证,2非盈利机构认证；type=1时才关注这个字段
	OrgCode *string `json:"OrgCode,omitempty" name:"OrgCode"`

	// 企业联系人
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 实名认证通过时间
	PassTime *string `json:"PassTime,omitempty" name:"PassTime"`

	// 地址
	Address *string `json:"Address,omitempty" name:"Address"`
}

type Block struct {

	// 区块编号
	BlockNum *uint64 `json:"BlockNum,omitempty" name:"BlockNum"`

	// 区块Hash数值
	DataHash *string `json:"DataHash,omitempty" name:"DataHash"`

	// 区块ID，与区块编号一致
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 前一个区块Hash（未使用）,与区块Hash数值一致
	PreHash *string `json:"PreHash,omitempty" name:"PreHash"`

	// 区块内的交易数量
	TxCount *uint64 `json:"TxCount,omitempty" name:"TxCount"`
}

type CertFilters struct {

	// 证书类型列表
	CertType []*uint64 `json:"CertType,omitempty" name:"CertType" list`

	// 证书密钥加密算法列表
	KeyAlg []*string `json:"KeyAlg,omitempty" name:"KeyAlg" list`

	// 证书状态列表
	CertSta []*uint64 `json:"CertSta,omitempty" name:"CertSta" list`

	// 证书ID
	CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

	// 证书组织名称
	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`

	// 证书申请时间
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 证书终止时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type Certinfo struct {

	// 证书ID
	CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

	// 证书组织
	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`

	// 证书SN
	CertSn *string `json:"CertSn,omitempty" name:"CertSn"`

	// 证书DN
	CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

	// 证书类型
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 证书状态:0:正常;1:注销;2：冻结;3:申请失败
	CertSta *uint64 `json:"CertSta,omitempty" name:"CertSta"`

	// 申请时间
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 终止时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 密钥算法
	KeyAlg *string `json:"KeyAlg,omitempty" name:"KeyAlg"`

	// 用户证书标识，0：非用户证书；1：用户证书
	UserCert *uint64 `json:"UserCert,omitempty" name:"UserCert"`

	// CA类型：0，fabricCA;1,cfca
	CaType *uint64 `json:"CaType,omitempty" name:"CaType"`
}

type Chaincode struct {

	// 合约ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署通道数
	ChannelNumber *uint64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 合约名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 部署节点数
	PeerNumber *uint64 `json:"PeerNumber,omitempty" name:"PeerNumber"`

	// 当前合约版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 该合约版本列表
	ChaincodeVersionList []*ChaincodeVersion `json:"ChaincodeVersionList,omitempty" name:"ChaincodeVersionList" list`

	// 合约创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ChaincodeOfChannel struct {

	// 策略内容
	Endorsement *string `json:"Endorsement,omitempty" name:"Endorsement"`

	// 策略ID
	EndorsementId *string `json:"EndorsementId,omitempty" name:"EndorsementId"`

	// 策略用户
	EndorsementUser *string `json:"EndorsementUser,omitempty" name:"EndorsementUser"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 合约ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 合约名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合约版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 0:包含私有数据；1:不包含私有数据
	PrivateData *uint64 `json:"PrivateData,omitempty" name:"PrivateData"`
}

type ChaincodeOfEndorsement struct {

	// 合约ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 合约名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合约版本
	Version *string `json:"Version,omitempty" name:"Version"`
}

type ChaincodeVersion struct {

	// 合约ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署通道数
	ChannelNumber *uint64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`

	// 部署节点数
	PeerNumber *uint64 `json:"PeerNumber,omitempty" name:"PeerNumber"`

	// 版本号
	Version *string `json:"Version,omitempty" name:"Version"`
}

type Channel struct {

	// 策略ID
	EndorsementId *string `json:"EndorsementId,omitempty" name:"EndorsementId"`

	// 策略内容
	EndorsementUser *string `json:"EndorsementUser,omitempty" name:"EndorsementUser"`

	// 版本
	HasVersion *string `json:"HasVersion,omitempty" name:"HasVersion"`

	// ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 通道名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ChannelForCloudMonitor struct {

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 通道所属组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 通道所在网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type ChannelInviteeForEvent struct {

	// 组织成员ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组织成员名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 处理状态(0:等待接受;1:接受邀请;2:拒绝邀请;3:超时失效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下一阶段开始与否(0:未开始;1:进行中;2:已结束)
	NextStep *uint64 `json:"NextStep,omitempty" name:"NextStep"`

	// 操作时间
	OpTime *string `json:"OpTime,omitempty" name:"OpTime"`

	// 组织别名
	GroupAlias *string `json:"GroupAlias,omitempty" name:"GroupAlias"`

	// 成员账号
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 成员名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ChannelObj struct {

	// 组织ID
	GroupsId *string `json:"GroupsId,omitempty" name:"GroupsId"`

	// 组织名称
	GroupsName *string `json:"GroupsName,omitempty" name:"GroupsName"`

	// 高度
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 通道ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 通道备注
	Information *string `json:"Information,omitempty" name:"Information"`

	// 通道名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 通道准入率
	VoteRate *uint64 `json:"VoteRate,omitempty" name:"VoteRate"`
}

type ChannelOfChaincode struct {

	// 策略内容
	Endorsement *string `json:"Endorsement,omitempty" name:"Endorsement"`

	// 策略内容
	EndorsementUser *string `json:"EndorsementUser,omitempty" name:"EndorsementUser"`

	// 策略ID
	EndorsementsId *string `json:"EndorsementsId,omitempty" name:"EndorsementsId"`

	// 通道ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 通道备注
	Information *string `json:"Information,omitempty" name:"Information"`

	// 通道名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`
}

type ChannelVoterForEvent struct {

	// 组织成员ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组织成员名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 处理状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 操作时间
	OpTime *string `json:"OpTime,omitempty" name:"OpTime"`

	// 组织别名
	GroupAlias *string `json:"GroupAlias,omitempty" name:"GroupAlias"`

	// 成员账号
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 成员名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CheckChaincodeASTRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 合约类型：go，gozip，javazip
	ChaincodeFileType *string `json:"ChaincodeFileType,omitempty" name:"ChaincodeFileType"`

	// 合约是否上传cos：0，未上传；1：上传
	CosFlag *uint64 `json:"CosFlag,omitempty" name:"CosFlag"`

	// 合约内容
	ChaincodeContent *string `json:"ChaincodeContent,omitempty" name:"ChaincodeContent"`

	// 合约上传cos对应key
	CosKey *string `json:"CosKey,omitempty" name:"CosKey"`
}

func (r *CheckChaincodeASTRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckChaincodeASTRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckChaincodeASTResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合约AST检查结果：0：正确；1：出错
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 检查消息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 合约中的非法函数调用
		FunctionList []*string `json:"FunctionList,omitempty" name:"FunctionList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckChaincodeASTResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckChaincodeASTResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckChaincodeChannelRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CheckChaincodeChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckChaincodeChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckChaincodeChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否已初始化标识
		Inited *int64 `json:"Inited,omitempty" name:"Inited"`

		// 合约初始化结果
		Result *uint64 `json:"Result,omitempty" name:"Result"`

		// 初始化消息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckChaincodeChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckChaincodeChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckCreateChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 合约升级标志（0：创建；1：升级）
	UpGrade *uint64 `json:"UpGrade,omitempty" name:"UpGrade"`
}

func (r *CheckCreateChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckCreateChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckCreateChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合约检查结果（0：检查成功；1：检查失败）
		CheckResult *uint64 `json:"CheckResult,omitempty" name:"CheckResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCreateChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckCreateChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckEventConfigRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *CheckEventConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckEventConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckEventConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 错误；1：正确
		CheckResult *uint64 `json:"CheckResult,omitempty" name:"CheckResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckEventConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckEventConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckGroupClusterCreatorRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// cluster标识
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CheckGroupClusterCreatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckGroupClusterCreatorRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckGroupClusterCreatorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否创建者标识
		Creator *uint64 `json:"Creator,omitempty" name:"Creator"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckGroupClusterCreatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckGroupClusterCreatorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClusterFilter struct {

	// 筛选项名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 筛选项值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ClusterForCloudMonitor struct {

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 网络名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 网络中用户的组织数量
	GroupCount *uint64 `json:"GroupCount,omitempty" name:"GroupCount"`

	// 用户当前组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type ClusterItem struct {

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 网络名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 网络类型
	ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 联盟名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 成员数量
	MemberCount *uint64 `json:"MemberCount,omitempty" name:"MemberCount"`

	// 组织数量
	GroupCount *uint64 `json:"GroupCount,omitempty" name:"GroupCount"`

	// 节点数量
	PeerCount *uint64 `json:"PeerCount,omitempty" name:"PeerCount"`

	// 我的节点数量
	MyPeerCount *uint64 `json:"MyPeerCount,omitempty" name:"MyPeerCount"`

	// 集群状态
	ClusterStatus *uint64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 网络创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 网络创建标识
	Creator *uint64 `json:"Creator,omitempty" name:"Creator"`

	// 网络到期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 当前账号创建的组织数量
	MyGroupCount *uint64 `json:"MyGroupCount,omitempty" name:"MyGroupCount"`

	// 网络通道总数量
	ChannelCount *uint64 `json:"ChannelCount,omitempty" name:"ChannelCount"`

	// 当前账号加入的通道数量
	MyChannelCount *uint64 `json:"MyChannelCount,omitempty" name:"MyChannelCount"`
}

type ClusterMemberForEvent struct {

	// 组织成员ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组织成员名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 认证类型
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 加入时间
	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

	// 成员账号
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 成员昵称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ClusterNumberPerRegion struct {

	// 区域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 网络数量
	ClusterNumber *uint64 `json:"ClusterNumber,omitempty" name:"ClusterNumber"`
}

type CompileChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约内容（go语言）
	ChaincodeContent *string `json:"ChaincodeContent,omitempty" name:"ChaincodeContent"`
}

func (r *CompileChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompileChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompileChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 编译结果输出
		Output *string `json:"Output,omitempty" name:"Output"`

		// 编译状态码
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompileChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompileChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConsortiumClusterItem struct {

	// 网络资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 网络名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 区块链引擎类型（0：fabric，1：trustsql，2：bcos）
	BlockchainType *uint64 `json:"BlockchainType,omitempty" name:"BlockchainType"`

	// 网络所属地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 网络状态（1：运行中， 2：即将到期，3：已过期）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 网络组织数量
	ClusterGroupCount *uint64 `json:"ClusterGroupCount,omitempty" name:"ClusterGroupCount"`

	// 网络节点数量
	ClusterPeerCount *uint64 `json:"ClusterPeerCount,omitempty" name:"ClusterPeerCount"`

	// 我的节点数量
	MyPeerCount *uint64 `json:"MyPeerCount,omitempty" name:"MyPeerCount"`

	// 网络类型（0：私有链；1：联盟链）
	ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 网络创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 网络到期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type ConsortiumItem struct {

	// 联盟ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 联盟名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 联盟创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 联盟成员数量
	MemberCount *uint64 `json:"MemberCount,omitempty" name:"MemberCount"`

	// 联盟网络数量
	ClusterCount *uint64 `json:"ClusterCount,omitempty" name:"ClusterCount"`

	// 联盟创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 联盟加入时间
	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

	// 联盟创建标识
	CreatorFlag *uint64 `json:"CreatorFlag,omitempty" name:"CreatorFlag"`
}

type ConsortiumMember struct {

	// 联盟成员ID
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 联盟成员
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// 联盟成员
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 联盟成员认证类型，0:未认证;1:个人;2:企业;4:其他(如：企业和个人都可以，则typelimit = 1+2 = 3)
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 联盟成员加入时间
	OpTime *string `json:"OpTime,omitempty" name:"OpTime"`

	// 是否是操作者，标识
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

type ConsortiumMemberForEvent struct {

	// 联盟成员ID
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 联盟成员名称
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// 认证类型
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 加入时间
	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

	// 成员账号
	MemberUin *string `json:"MemberUin,omitempty" name:"MemberUin"`
}

type CreateChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 创建者的组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 合约版本号
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 合约内容（go语言）
	ChaincodeContent *string `json:"ChaincodeContent,omitempty" name:"ChaincodeContent"`

	// 允许使用合约的组织ID列表
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList" list`

	// go, gozip，javazip， nodezip
	ChaincodeFileType *string `json:"ChaincodeFileType,omitempty" name:"ChaincodeFileType"`

	// 合约上传至cos的key
	CosKey *string `json:"CosKey,omitempty" name:"CosKey"`

	// 合约是否存储在cos上，当CosFlag为1时CosKey的值有效
	CosFlag *uint64 `json:"CosFlag,omitempty" name:"CosFlag"`
}

func (r *CreateChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建合约的ID
		ChaincodeId *string `json:"ChaincodeId,omitempty" name:"ChaincodeId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateChannelForCBCRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 通道名
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 通道投票率
	VoteRate *uint64 `json:"VoteRate,omitempty" name:"VoteRate"`

	// 通道信息
	ChannelInfo *string `json:"ChannelInfo,omitempty" name:"ChannelInfo"`
}

func (r *CreateChannelForCBCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateChannelForCBCRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateChannelForCBCResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道ID
		ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateChannelForCBCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateChannelForCBCResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateChannelRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 创建者组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 通道备注
	ChannelInfo *string `json:"ChannelInfo,omitempty" name:"ChannelInfo"`

	// 关键的组织ID列表
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList" list`

	// 投票率
	VoteRate *string `json:"VoteRate,omitempty" name:"VoteRate"`
}

func (r *CreateChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建后的通道ID
		ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConsortiumRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 联盟名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 成员类型限制，//0:全部;1:个人;2:企业;4:其他(如：企业和个人都可以，则typelimit = 1+2 = 3)
	TypeLimit *uint64 `json:"TypeLimit,omitempty" name:"TypeLimit"`

	// 实名认证限制：0:不需要实名认证;1:需要实名认证
	AuthLimit *uint64 `json:"AuthLimit,omitempty" name:"AuthLimit"`

	// 联盟描述，选填
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateConsortiumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConsortiumRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConsortiumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConsortiumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConsortiumResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEndorsementRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 名称
	EndorsementName *string `json:"EndorsementName,omitempty" name:"EndorsementName"`

	// 背书策略内容（方便前端展示）
	EndorsementUser *string `json:"EndorsementUser,omitempty" name:"EndorsementUser"`

	// 背书策略内容
	Endorsement *string `json:"Endorsement,omitempty" name:"Endorsement"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 备注
	EndorsementInfo *string `json:"EndorsementInfo,omitempty" name:"EndorsementInfo"`
}

func (r *CreateEndorsementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEndorsementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEndorsementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的策略ID
		EndorsementId *string `json:"EndorsementId,omitempty" name:"EndorsementId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEndorsementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEndorsementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFabricClusterRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用区ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 商品详情
	GoodsDetail *GoodsDetailInfo `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
}

func (r *CreateFabricClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFabricClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFabricClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单列表
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFabricClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFabricClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateLinkRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 私有连接名称
	LinkName *string `json:"LinkName,omitempty" name:"LinkName"`

	// 客户端vpc统一ID
	ClientVpcId *string `json:"ClientVpcId,omitempty" name:"ClientVpcId"`

	// 客户端子网统一ID
	ClientSubNetId *string `json:"ClientSubNetId,omitempty" name:"ClientSubNetId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreatePrivateLinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePrivateLinkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateLinkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePrivateLinkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DealEventTaskRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务操作：1：同意/接受；2：反对/拒绝
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

func (r *DealEventTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DealEventTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DealEventTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DealEventTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DealEventTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateLinkRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 要删除的私有连接名称
	LinkId *uint64 `json:"LinkId,omitempty" name:"LinkId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeletePrivateLinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePrivateLinkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateLinkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePrivateLinkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestoryFabricClusterRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DestoryFabricClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestoryFabricClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestoryFabricClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestoryFabricClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestoryFabricClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCertRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 证书ID
	CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

	// 证书DN
	CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DownloadCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书名称
		CertName *string `json:"CertName,omitempty" name:"CertName"`

		// 证书内容
		CertCtx *string `json:"CertCtx,omitempty" name:"CertCtx"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadEventConfigRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 参数名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DownloadEventConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadEventConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadEventConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件配置文件名称
		EventConfigName *string `json:"EventConfigName,omitempty" name:"EventConfigName"`

		// 事件配置文件内容
		EventConfigContent *string `json:"EventConfigContent,omitempty" name:"EventConfigContent"`

		// 事件配置文件附加信息
		EventConfigInfo *string `json:"EventConfigInfo,omitempty" name:"EventConfigInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadEventConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadEventConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadTbaasIdentityRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *DownloadTbaasIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadTbaasIdentityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadTbaasIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户账号ID
		UserUin *string `json:"UserUin,omitempty" name:"UserUin"`

		// 用户APPID
		UserAppId *string `json:"UserAppId,omitempty" name:"UserAppId"`

		// 用户名称
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 用户tbaasId
		UserTbaasId *string `json:"UserTbaasId,omitempty" name:"UserTbaasId"`

		// 用户tbaasId认证信息
		TbaasIdentityContent *string `json:"TbaasIdentityContent,omitempty" name:"TbaasIdentityContent"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadTbaasIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadTbaasIdentityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Endorsement struct {

	// 应用合约数量
	ChaincodeNumber *int64 `json:"ChaincodeNumber,omitempty" name:"ChaincodeNumber"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 策略内容
	Endorsements *string `json:"Endorsements,omitempty" name:"Endorsements"`

	// 策略
	EndorsementUser *string `json:"EndorsementUser,omitempty" name:"EndorsementUser"`

	// 策略ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 策略信息
	Information *string `json:"Information,omitempty" name:"Information"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type EndorserGroup struct {

	// 背书组织名称
	EndorserGroupName *string `json:"EndorserGroupName,omitempty" name:"EndorserGroupName"`

	// 背书节点列表
	EndorserPeerList []*string `json:"EndorserPeerList,omitempty" name:"EndorserPeerList" list`
}

type EventConfigInfo struct {

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 事件配置文件名称
	EventConfigName *string `json:"EventConfigName,omitempty" name:"EventConfigName"`

	// 事件配置文件内容
	EventConfigContent *string `json:"EventConfigContent,omitempty" name:"EventConfigContent"`

	// 事件配置文件附加信息
	EventConfigInfo *string `json:"EventConfigInfo,omitempty" name:"EventConfigInfo"`
}

type EventCountItem struct {

	// 事件类型
	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`

	// 事件数量
	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
}

type EventItem struct {

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 事件名称
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 事件类型
	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`

	// 事件状态
	EventStatus *uint64 `json:"EventStatus,omitempty" name:"EventStatus"`

	// 事件所处步骤
	EventCurStep *uint64 `json:"EventCurStep,omitempty" name:"EventCurStep"`

	// 事件总步骤数
	EventTotalSteps *uint64 `json:"EventTotalSteps,omitempty" name:"EventTotalSteps"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 结束时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 0:我创建的;1:我代办的;2:我已办的
	DealFlag *uint64 `json:"DealFlag,omitempty" name:"DealFlag"`

	// 创建原因
	EventReason *string `json:"EventReason,omitempty" name:"EventReason"`

	// 事件发起人名称
	CreatorAppid *string `json:"CreatorAppid,omitempty" name:"CreatorAppid"`

	// 事件发起人账号
	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 事件发起人名称
	CreatorName *string `json:"CreatorName,omitempty" name:"CreatorName"`

	// 事件所属地域
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

type EventStepParam struct {

	// 事件的步骤
	Step *uint64 `json:"Step,omitempty" name:"Step"`

	// 该步骤状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 状态参数
	Params []*uint64 `json:"Params,omitempty" name:"Params" list`
}

type EventTarget struct {

	// 对象名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 对象判断标识，1:该对象是自己;0:该对象不是自己
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

type ExportChaincodeDevRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 导出类型：raw或者cos
	ExportType *string `json:"ExportType,omitempty" name:"ExportType"`
}

func (r *ExportChaincodeDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportChaincodeDevRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportChaincodeDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出内容
		Details []*ExportObj `json:"Details,omitempty" name:"Details" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportChaincodeDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportChaincodeDevResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportObj struct {

	// cos下载链接
	Context *string `json:"Context,omitempty" name:"Context"`

	// 导出方法
	ExportType *string `json:"ExportType,omitempty" name:"ExportType"`

	// 文件名
	File *string `json:"File,omitempty" name:"File"`
}

type FeeModifyPrice struct {

	// 价格
	Policy *int64 `json:"Policy,omitempty" name:"Policy"`

	// 价格
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 价格
	RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 价格
	TotalCost *int64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type FeePrice struct {

	// 价格
	Policy *int64 `json:"Policy,omitempty" name:"Policy"`

	// 价格
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 价格
	RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 价格
	TotalCost *int64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type GTask struct {

	// 待办任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 代办任务类型,1:邀请;2:投票
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type GetBlockDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 区块ID
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *GetBlockDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBlockDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBlockDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 区块信息详情
		BlockDetail *string `json:"BlockDetail,omitempty" name:"BlockDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBlockDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBlockDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBlockListRequest struct {
	*tchttp.BaseRequest

	// 模块名称，固定字段：block
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：block_list
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID，固定字段：0
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID，固定字段：0
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBlockListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBlockListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 区块数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 区块列表
		BlockList []*Block `json:"BlockList,omitempty" name:"BlockList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBlockListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBlockSpeedRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetBlockSpeedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBlockSpeedRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBlockSpeedResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 区块交易速率曲线数据
		SpeedList []*TransactionSpeed `json:"SpeedList,omitempty" name:"SpeedList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBlockSpeedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBlockSpeedResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBlockTransactionListRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 区块ID
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetBlockTransactionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBlockTransactionListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetBlockTransactionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 交易列表
		TransactionList []*TransactionItem `json:"TransactionList,omitempty" name:"TransactionList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBlockTransactionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetBlockTransactionListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCdhStatusRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 小区域ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 虚拟机型号
	VmType *string `json:"VmType,omitempty" name:"VmType"`
}

func (r *GetCdhStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCdhStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCdhStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// cdh售卖状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCdhStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCdhStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCertDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 证书ID
	CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

	// 证书DN
	CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetCertDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCertDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCertDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书ID
		CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

		// 证书类型
		CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

		// 证书SN
		CertSn *string `json:"CertSn,omitempty" name:"CertSn"`

		// 证书DN
		CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

		// 密钥算法
		KeyAlg *string `json:"KeyAlg,omitempty" name:"KeyAlg"`

		// 密钥长度
		KeyLen *uint64 `json:"KeyLen,omitempty" name:"KeyLen"`

		// 签名算法
		SignAlg *string `json:"SignAlg,omitempty" name:"SignAlg"`

		// 有效时间
		ValidityTime *uint64 `json:"ValidityTime,omitempty" name:"ValidityTime"`

		// 证书组织
		OrgName *string `json:"OrgName,omitempty" name:"OrgName"`

		// 证件类型
		IdentiType *string `json:"IdentiType,omitempty" name:"IdentiType"`

		// 证件号码
		IdentiNo *string `json:"IdentiNo,omitempty" name:"IdentiNo"`

		// 邮箱信息
		Email *string `json:"Email,omitempty" name:"Email"`

		// 电话信息
		PhoneNo *string `json:"PhoneNo,omitempty" name:"PhoneNo"`

		// 地址信息
		Address *string `json:"Address,omitempty" name:"Address"`

		// 用户标识
		UserLable *string `json:"UserLable,omitempty" name:"UserLable"`

		// 备注信息
		Notes *string `json:"Notes,omitempty" name:"Notes"`

		// 过期时间，和网络过期时间保持一致
		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCertDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCertDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCertListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 过滤器
	Filters *CertFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 最新前版本固定值为'v2.1'
	TbaasVersion *string `json:"TbaasVersion,omitempty" name:"TbaasVersion"`
}

func (r *GetCertListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCertListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCertListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 证书列表
		CertList []*Certinfo `json:"CertList,omitempty" name:"CertList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCertListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCertListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetChaincodeDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合约内容或合约在cos的文件路径
		Chaincode *string `json:"Chaincode,omitempty" name:"Chaincode"`

		// 合约文件类型（go或gozip）
		FileType *string `json:"FileType,omitempty" name:"FileType"`

		// 组织ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 组织名称
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 合约ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 合约名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 合约版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChaincodeDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeDevAccessAuthRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *GetChaincodeDevAccessAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeDevAccessAuthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeDevAccessAuthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0 : 不允许访问； 1：允许访问
		Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChaincodeDevAccessAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeDevAccessAuthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeListPerChannelRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查看关键词
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *GetChaincodeListPerChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeListPerChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeListPerChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合约列表
		Chaincodes []*ChaincodeOfChannel `json:"Chaincodes,omitempty" name:"Chaincodes" list`

		// 合约总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChaincodeListPerChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeListPerChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeListPerEndorsementRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 背书策略ID
	EndorsementId *string `json:"EndorsementId,omitempty" name:"EndorsementId"`
}

func (r *GetChaincodeListPerEndorsementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeListPerEndorsementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeListPerEndorsementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略详情中的应用合约列表
		Chaincodes []*ChaincodeOfEndorsement `json:"Chaincodes,omitempty" name:"Chaincodes" list`

		// 合约列表总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChaincodeListPerEndorsementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeListPerEndorsementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询关键词
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`
}

func (r *GetChaincodeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 后台返回合约列表数据
		Chaincodes []*Chaincode `json:"Chaincodes,omitempty" name:"Chaincodes" list`

		// 合约总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChaincodeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodePvtDataListRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 合约ID
	ChaincodeId *string `json:"ChaincodeId,omitempty" name:"ChaincodeId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 偏移项数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页内项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetChaincodePvtDataListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodePvtDataListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodePvtDataListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有数据集数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 私有数据集列表
		PrivateDataList []*PrivateData `json:"PrivateDataList,omitempty" name:"PrivateDataList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChaincodePvtDataListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodePvtDataListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
}

func (r *GetChaincodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模版内容
		Context *string `json:"Context,omitempty" name:"Context"`

		// 模版文件名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChaincodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChaincodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelDetailForEventRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *GetChannelDetailForEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelDetailForEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelDetailForEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道ID
		ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

		// 通道名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 通道描述
		Information *string `json:"Information,omitempty" name:"Information"`

		// 通道投票率
		VoteRate *uint64 `json:"VoteRate,omitempty" name:"VoteRate"`

		// 通道创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 事件发起人名称
		CreatorAppid *string `json:"CreatorAppid,omitempty" name:"CreatorAppid"`

		// 事件发起人账号
		CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

		// 事件发起人名称
		CreatorName *string `json:"CreatorName,omitempty" name:"CreatorName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChannelDetailForEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelDetailForEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetChannelDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道备注
		Information *string `json:"Information,omitempty" name:"Information"`

		// 通道名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 时间
		Time *string `json:"Time,omitempty" name:"Time"`

		// 通道准入率
		VoteRate *uint64 `json:"VoteRate,omitempty" name:"VoteRate"`

		// 当前组织名称
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 网络ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChannelDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelInviteesForEventRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetChannelInviteesForEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelInviteesForEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelInviteesForEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道被邀请组织总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 通道被邀请组织总列表
		InviteeList []*ChannelInviteeForEvent `json:"InviteeList,omitempty" name:"InviteeList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChannelInviteesForEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelInviteesForEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelListForCloudMonitorRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 偏移项数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页内项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 关键词
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *GetChannelListForCloudMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelListForCloudMonitorRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelListForCloudMonitorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 通道列表
		ChannelList []*ChannelForCloudMonitor `json:"ChannelList,omitempty" name:"ChannelList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChannelListForCloudMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelListForCloudMonitorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelListForInitRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetChannelListForInitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelListForInitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelListForInitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合约初始化下拉通道列表
		Channels []*Channel `json:"Channels,omitempty" name:"Channels" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChannelListForInitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelListForInitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelListPerChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetChannelListPerChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelListPerChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelListPerChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合约部署通道列表
		Channels []*ChannelOfChaincode `json:"Channels,omitempty" name:"Channels" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChannelListPerChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelListPerChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查看关键词（通道名称）
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *GetChannelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道列表
		Channels []*ChannelObj `json:"Channels,omitempty" name:"Channels" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChannelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelVotersForEventRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetChannelVotersForEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelVotersForEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetChannelVotersForEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道投票组织数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 通道投票组织列表
		VoterList []*ChannelVoterForEvent `json:"VoterList,omitempty" name:"VoterList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChannelVotersForEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetChannelVotersForEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusteDetailForEventRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *GetClusteDetailForEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusteDetailForEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusteDetailForEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 区块链网络ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 区块链网络名称
		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

		// 联盟名称
		ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

		// 组织数量
		GroupCount *uint64 `json:"GroupCount,omitempty" name:"GroupCount"`

		// 共识算法
		Consensus *uint64 `json:"Consensus,omitempty" name:"Consensus"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 事件发起人APPid
		CreatorAppid *string `json:"CreatorAppid,omitempty" name:"CreatorAppid"`

		// 事件发起人账号
		CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

		// 事件发起人名称
		CreatorName *string `json:"CreatorName,omitempty" name:"CreatorName"`

		// 网络状态
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 区块链网络根域
		ClusterDomain *string `json:"ClusterDomain,omitempty" name:"ClusterDomain"`

		// vpc id
		UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

		// vpc网段
		VpcCidr *string `json:"VpcCidr,omitempty" name:"VpcCidr"`

		// 联盟ID
		ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

		// 地域ID
		RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

		// 可用区ID
		ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

		// 网络CA类型:0，fabricCA;1,CFCA
		CaType *uint64 `json:"CaType,omitempty" name:"CaType"`

		// 网络证书加密算法类型：0,ecc；1：sm2
		KeyAlg *uint64 `json:"KeyAlg,omitempty" name:"KeyAlg"`

		// 网络peer节点总数量
		PeerCount *uint64 `json:"PeerCount,omitempty" name:"PeerCount"`

		// 世界状态数据库类型：0，leveldb；1，couchdb
		StateDBType *uint64 `json:"StateDBType,omitempty" name:"StateDBType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusteDetailForEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusteDetailForEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusteMemberForEventRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetClusteMemberForEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusteMemberForEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusteMemberForEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络成员总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 网络成员总列表
		MemberList []*ClusterMemberForEvent `json:"MemberList,omitempty" name:"MemberList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusteMemberForEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusteMemberForEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 区块链网络ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 资源ID
		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

		// 区块链网络名称
		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

		// 区块链类型: 0:私有链,1:公有链
		ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

		// 联盟ID
		Description *string `json:"Description,omitempty" name:"Description"`

		// 组织数量
		GroupCount *uint64 `json:"GroupCount,omitempty" name:"GroupCount"`

		// 总peer个数
		PeerCount *uint64 `json:"PeerCount,omitempty" name:"PeerCount"`

		// 共识算法
		Consensus *uint64 `json:"Consensus,omitempty" name:"Consensus"`

		// 集群状态
		ClusterStatus *uint64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 加入时间
		JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

		// 0:我创建的;1:我加入的
		CreateFlag *uint64 `json:"CreateFlag,omitempty" name:"CreateFlag"`

		// 联盟id
		ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

		// 联盟名称
		ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

		// 地域
		Region *uint64 `json:"Region,omitempty" name:"Region"`

		// 可用区
		ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

		// cvm数量
		CvmNumber *uint64 `json:"CvmNumber,omitempty" name:"CvmNumber"`

		// cvm型号
		CvmTypeCn *string `json:"CvmTypeCn,omitempty" name:"CvmTypeCn"`

		// peer的cbs大小
		TotalPeerCbs *string `json:"TotalPeerCbs,omitempty" name:"TotalPeerCbs"`

		// 0：leveldb；1：couchdb
		StateDBType *uint64 `json:"StateDBType,omitempty" name:"StateDBType"`

		// 资源标签
		TagList []*TagItem `json:"TagList,omitempty" name:"TagList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterListForCloudMonitorRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 偏移项数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页内项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 关键词
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *GetClusterListForCloudMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterListForCloudMonitorRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterListForCloudMonitorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 网络列表
		ClusterList []*ClusterForCloudMonitor `json:"ClusterList,omitempty" name:"ClusterList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterListForCloudMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterListForCloudMonitorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 关键字
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 集群状态
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList" list`

	// 网络列表筛选项
	Filters []*ClusterFilter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *GetClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 网络列表
		ClusterList []*ClusterItem `json:"ClusterList,omitempty" name:"ClusterList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterNoJoinMembersRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetClusterNoJoinMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterNoJoinMembersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterNoJoinMembersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 未加入网络的联盟成员列表
		NoJoinMembersList []*ConsortiumMember `json:"NoJoinMembersList,omitempty" name:"NoJoinMembersList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterNoJoinMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterNoJoinMembersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterNumberStatisticsRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 统计的网络状态
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList" list`

	// 联盟ID，可选参数，如果要统计联盟内的网络数量，需要传入该参数
	ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`
}

func (r *GetClusterNumberStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterNumberStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterNumberStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户网络总数
		TotalClusterNumber *uint64 `json:"TotalClusterNumber,omitempty" name:"TotalClusterNumber"`

		// 每个zone的网络数量
		ClusterList []*ClusterNumberPerRegion `json:"ClusterList,omitempty" name:"ClusterList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterNumberStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterNumberStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterResourceInfoRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetClusterResourceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterResourceInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterResourceInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 资源ID
		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

		// 用户AppID
		AppId *string `json:"AppId,omitempty" name:"AppId"`

		// 用户Uin
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// 资源数据库ID
		AutoId *uint64 `json:"AutoId,omitempty" name:"AutoId"`

		// 项目ID
		ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// 续费标识
		RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

		// 地域
		Region *uint64 `json:"Region,omitempty" name:"Region"`

		// 可用区ID
		ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

		// 资源状态
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 付费模式
		PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

		// 隔离时间
		IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 过期时间
		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// 产品详细信息
		GoodsDetail *GoodsDetailItem `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterResourceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterResourceInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterSummaryRequest struct {
	*tchttp.BaseRequest

	// 模块名称，固定字段：cluster_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：cluster_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织ID，固定字段：0
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *GetClusterSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络通道总数量
		TotalChannelCount *uint64 `json:"TotalChannelCount,omitempty" name:"TotalChannelCount"`

		// 当前组织创建的通道数量
		MyChannelCount *uint64 `json:"MyChannelCount,omitempty" name:"MyChannelCount"`

		// 当前组织加入的通道数量
		JoinChannelCount *uint64 `json:"JoinChannelCount,omitempty" name:"JoinChannelCount"`

		// 网络节点总数量
		TotalPeerCount *uint64 `json:"TotalPeerCount,omitempty" name:"TotalPeerCount"`

		// 当前组织创建的节点数量
		MyPeerCount *uint64 `json:"MyPeerCount,omitempty" name:"MyPeerCount"`

		// 其他组织创建的节点数量
		OrderCount *uint64 `json:"OrderCount,omitempty" name:"OrderCount"`

		// 网络组织总数量
		TotalGroupCount *uint64 `json:"TotalGroupCount,omitempty" name:"TotalGroupCount"`

		// 当前组织创建的组织数量
		MyGroupCount *uint64 `json:"MyGroupCount,omitempty" name:"MyGroupCount"`

		// 网络智能合约总数量
		TotalChaincodeCount *uint64 `json:"TotalChaincodeCount,omitempty" name:"TotalChaincodeCount"`

		// 最近7天发起的智能合约数量
		RecentChaincodeCount *uint64 `json:"RecentChaincodeCount,omitempty" name:"RecentChaincodeCount"`

		// 当前组织发起的智能合约数量
		MyChaincodeCount *uint64 `json:"MyChaincodeCount,omitempty" name:"MyChaincodeCount"`

		// 当前组织的证书总数量
		TotalCertCount *uint64 `json:"TotalCertCount,omitempty" name:"TotalCertCount"`

		// 颁发给当前组织的证书数量
		TlsCertCount *uint64 `json:"TlsCertCount,omitempty" name:"TlsCertCount"`

		// 网络背书节点证书数量
		PeerCertCount *uint64 `json:"PeerCertCount,omitempty" name:"PeerCertCount"`

		// 当前组织业务证书数量
		ClientCertCount *uint64 `json:"ClientCertCount,omitempty" name:"ClientCertCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterTypeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 最新前版本固定值为'v2.1'
	TbaasVersion *string `json:"TbaasVersion,omitempty" name:"TbaasVersion"`
}

func (r *GetClusterTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetClusterTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络类型：0，私有链;1，联盟链
		ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

		// CA类型：0，tbaas-ca；1：cfca
		CaType *uint64 `json:"CaType,omitempty" name:"CaType"`

		// 加密算法：0，ECC；1，SM2
		KeyAlg *uint64 `json:"KeyAlg,omitempty" name:"KeyAlg"`

		// fabric版本号
		Version *string `json:"Version,omitempty" name:"Version"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetClusterTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumClusterListSummaryRequest struct {
	*tchttp.BaseRequest

	// 模块名，默认参数：summary_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，默认参数：tbaas_cluster_list_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 关键字筛选（网络名称和资源ID）
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 筛选项：'status',//筛选网络状态 ；'blockchain_type',//筛选网络引擎类型
	Filters []*ClusterFilter `json:"Filters,omitempty" name:"Filters" list`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移项数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetConsortiumClusterListSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumClusterListSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumClusterListSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 网络列表
		ClusterList []*ConsortiumClusterItem `json:"ClusterList,omitempty" name:"ClusterList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsortiumClusterListSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumClusterListSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`
}

func (r *GetConsortiumDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 联盟ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 联盟名称
		ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

		// 联盟描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 联盟成员类型限制
		TypeLimit *uint64 `json:"TypeLimit,omitempty" name:"TypeLimit"`

		// 联盟成员实名认证限制
		AuthLimit *uint64 `json:"AuthLimit,omitempty" name:"AuthLimit"`

		// 联盟创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 联盟加入时间
		JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

		// 联盟创建标识// 0:我创建的;1:我加入的
		CreateFlag *uint64 `json:"CreateFlag,omitempty" name:"CreateFlag"`

		// 创建者APPID
		CreatorAppid *string `json:"CreatorAppid,omitempty" name:"CreatorAppid"`

		// 创建者账号
		CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

		// 创建者名称
		CreatorName *string `json:"CreatorName,omitempty" name:"CreatorName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsortiumDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumDtailForEventRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *GetConsortiumDtailForEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumDtailForEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumDtailForEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 联盟ID
		ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

		// 联盟名称
		ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

		// 联盟描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 成员类型
		TypeLimit *uint64 `json:"TypeLimit,omitempty" name:"TypeLimit"`

		// 认证限制
		AuthLimit *uint64 `json:"AuthLimit,omitempty" name:"AuthLimit"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 创建人名称
		CreatorAppid *string `json:"CreatorAppid,omitempty" name:"CreatorAppid"`

		// 创建人账号
		CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

		// 创建人名称
		CreatorName *string `json:"CreatorName,omitempty" name:"CreatorName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsortiumDtailForEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumDtailForEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 关键字
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方法
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *GetConsortiumListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 联盟总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 联盟列表
		ConsortiumList []*ConsortiumItem `json:"ConsortiumList,omitempty" name:"ConsortiumList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsortiumListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumMemberAuthDataRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 成员账号ID
	MemberUin *string `json:"MemberUin,omitempty" name:"MemberUin"`

	// 成员APPID
	MemberAppId *string `json:"MemberAppId,omitempty" name:"MemberAppId"`
}

func (r *GetConsortiumMemberAuthDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumMemberAuthDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumMemberAuthDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否实名认证：0，无；1，有
		IsAuthenticated *uint64 `json:"IsAuthenticated,omitempty" name:"IsAuthenticated"`

		// 用户认证信息，IsAuthenticated=1时有效
		AuthData *AuthData `json:"AuthData,omitempty" name:"AuthData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsortiumMemberAuthDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumMemberAuthDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumMemberForEventRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetConsortiumMemberForEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumMemberForEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumMemberForEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 联盟成员总量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 联盟成员列表
		MemberList []*ConsortiumMemberForEvent `json:"MemberList,omitempty" name:"MemberList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsortiumMemberForEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumMemberForEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumMembersRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`
}

func (r *GetConsortiumMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumMembersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumMembersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 联盟成员列表
		JoinMembersList []*ConsortiumMember `json:"JoinMembersList,omitempty" name:"JoinMembersList" list`

		// 带加入成员列表
		InviteMembersList []*ConsortiumMember `json:"InviteMembersList,omitempty" name:"InviteMembersList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsortiumMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumMembersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumSummaryRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *GetConsortiumSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetConsortiumSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 联盟总数量
		TotalConsortium *uint64 `json:"TotalConsortium,omitempty" name:"TotalConsortium"`

		// 我创建的联盟数量
		CreateConsortium *uint64 `json:"CreateConsortium,omitempty" name:"CreateConsortium"`

		// 我加入的联盟数量
		JoinConsortium *uint64 `json:"JoinConsortium,omitempty" name:"JoinConsortium"`

		// 网络数量
		TotalCluster *uint64 `json:"TotalCluster,omitempty" name:"TotalCluster"`

		// 我创建的网络数量
		CreateCluster *uint64 `json:"CreateCluster,omitempty" name:"CreateCluster"`

		// 我加入的网络数量
		JoinCluster *uint64 `json:"JoinCluster,omitempty" name:"JoinCluster"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConsortiumSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetConsortiumSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCosSignRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件名
	CosFile *string `json:"CosFile,omitempty" name:"CosFile"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *GetCosSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCosSignRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCosSignResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// cos签名
		CosSign *string `json:"CosSign,omitempty" name:"CosSign"`

		// cos关联appid
		CosAppid *string `json:"CosAppid,omitempty" name:"CosAppid"`

		// cos存储地域
		CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

		// cos存储桶名
		CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

		// cos访问域名
		CosDomain *string `json:"CosDomain,omitempty" name:"CosDomain"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCosSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCosSignResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCosUrlRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 获取cos链接的key（目录）
	CosKey *string `json:"CosKey,omitempty" name:"CosKey"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`
}

func (r *GetCosUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCosUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCosUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// cos文件下载链接
		CosFileUrl *string `json:"CosFileUrl,omitempty" name:"CosFileUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCosUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCosUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEndorsementDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 策略ID
	EndorsementId *string `json:"EndorsementId,omitempty" name:"EndorsementId"`
}

func (r *GetEndorsementDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEndorsementDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEndorsementDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道ID
		ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 策略内容
		Endorsement *string `json:"Endorsement,omitempty" name:"Endorsement"`

		// 策略
		EndorsementUser *string `json:"EndorsementUser,omitempty" name:"EndorsementUser"`

		// 策略ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 策略信息
		Information *string `json:"Information,omitempty" name:"Information"`

		// 策略名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEndorsementDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEndorsementDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEndorsementListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查看关键词（操作内容）
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *GetEndorsementListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEndorsementListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEndorsementListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略列表
		Endorsements []*Endorsement `json:"Endorsements,omitempty" name:"Endorsements" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEndorsementListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEndorsementListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEventDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *GetEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEventDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件ID
		EventId *string `json:"EventId,omitempty" name:"EventId"`

		// 事件名称
		EventName *string `json:"EventName,omitempty" name:"EventName"`

		// 事件类型
		EventType *uint64 `json:"EventType,omitempty" name:"EventType"`

		// 事件所处步骤
		EventCurStep *uint64 `json:"EventCurStep,omitempty" name:"EventCurStep"`

		// 事件总步骤数
		EventTotalSteps *uint64 `json:"EventTotalSteps,omitempty" name:"EventTotalSteps"`

		// 事件目标
		EventTarget []*EventTarget `json:"EventTarget,omitempty" name:"EventTarget" list`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 结束时间
		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// 事件状态,0:进行中;1:成功;2:失败
		EventStatus *uint64 `json:"EventStatus,omitempty" name:"EventStatus"`

		// 创建原因
		EventReason *string `json:"EventReason,omitempty" name:"EventReason"`

		// 0:我创建的;1:我代办的;2:我已办的
		DealFlag *uint64 `json:"DealFlag,omitempty" name:"DealFlag"`

		// 事件发起人名称
		CreatorAppid *string `json:"CreatorAppid,omitempty" name:"CreatorAppid"`

		// 事件发起人账号
		CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

		// 事件发起人名称
		CreatorName *string `json:"CreatorName,omitempty" name:"CreatorName"`

		// 事件所属地域
		RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEventDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEventListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 创建标识：0:创建;1:参与
	CreateFlag *uint64 `json:"CreateFlag,omitempty" name:"CreateFlag"`

	// 事件类型
	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`

	// 事件状态
	EventStatus []*uint64 `json:"EventStatus,omitempty" name:"EventStatus" list`

	// 关键词
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方式
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *GetEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEventListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 事件列表
		EventList []*EventItem `json:"EventList,omitempty" name:"EventList" list`

		// 待处理事件数量
		DealCount *uint64 `json:"DealCount,omitempty" name:"DealCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEventListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEventStepStatusRequest struct {
	*tchttp.BaseRequest

	// 区块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *GetEventStepStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEventStepStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEventStepStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前流程步骤
		TotalSteps *uint64 `json:"TotalSteps,omitempty" name:"TotalSteps"`

		// 当前流程步骤
		CurStep *uint64 `json:"CurStep,omitempty" name:"CurStep"`

		// 当前流程步骤状态
		StepParams []*EventStepParam `json:"StepParams,omitempty" name:"StepParams" list`

		// 到期时间
		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// 是否有待办事务：0，无；1：有
		ToDo *uint64 `json:"ToDo,omitempty" name:"ToDo"`

		// 待办事务，ToDo为1时有效
		GTask *GTask `json:"GTask,omitempty" name:"GTask"`

		// 事件状态
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEventStepStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEventStepStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEventSummaryRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：event_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：event_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *GetEventSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEventSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetEventSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的事件总数
		CreateEventCount *uint64 `json:"CreateEventCount,omitempty" name:"CreateEventCount"`

		// 待办任务总数
		ToDoEventCount *uint64 `json:"ToDoEventCount,omitempty" name:"ToDoEventCount"`

		// 已办任务总数
		DoneEventCount *uint64 `json:"DoneEventCount,omitempty" name:"DoneEventCount"`

		// 根据事件类型统计创建的事件总数
		EventCountList []*EventCountItem `json:"EventCountList,omitempty" name:"EventCountList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEventSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetEventSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListForCloudMonitorRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 偏移项数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页内项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 关键字
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *GetGroupListForCloudMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListForCloudMonitorRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListForCloudMonitorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组织总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 用户组织列表
		GroupList []*GroupForCloudMonitor `json:"GroupList,omitempty" name:"GroupList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListForCloudMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListForCloudMonitorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListNoChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetGroupListNoChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListNoChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListNoChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组织list
		Groups []*Group `json:"Groups,omitempty" name:"Groups" list`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListNoChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListNoChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListPerAppidRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *GetGroupListPerAppidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListPerAppidRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListPerAppidResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组织列表
		Groups []*Group `json:"Groups,omitempty" name:"Groups" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListPerAppidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListPerAppidResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListPerChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetGroupListPerChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListPerChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListPerChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合约详情中的组织列表
		Groups []*GroupOfChaincode `json:"Groups,omitempty" name:"Groups" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListPerChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListPerChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListPerChannelRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetGroupListPerChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListPerChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListPerChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组织列表
		Groups []*Group `json:"Groups,omitempty" name:"Groups" list`

		// 组织列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListPerChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListPerChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListPerClusterRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// cluster标识
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询关键词
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetGroupListPerClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListPerClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListPerClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组织列表
		Groups []*GroupOfCluster `json:"Groups,omitempty" name:"Groups" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListPerClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListPerClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInvokeTxRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：query_txid
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点名称，可以在通道详情中获取该通道上的节点名称极其所属组织名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 执行该查询交易的节点所属组织名称，可以在通道详情中获取该通道上的节点名称极其所属组织名称
	PeerGroup *string `json:"PeerGroup,omitempty" name:"PeerGroup"`

	// 交易ID
	TxId *string `json:"TxId,omitempty" name:"TxId"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *GetInvokeTxRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInvokeTxRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInvokeTxResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易执行状态码
		TxValidationCode *int64 `json:"TxValidationCode,omitempty" name:"TxValidationCode"`

		// 交易执行消息
		TxValidationMsg *string `json:"TxValidationMsg,omitempty" name:"TxValidationMsg"`

		// 交易所在区块ID
		BlockId *int64 `json:"BlockId,omitempty" name:"BlockId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInvokeTxResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInvokeTxResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLatesdTransactionListRequest struct {
	*tchttp.BaseRequest

	// 模块名称，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：latest_transaction_list
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID，固定字段：0
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID，固定字段：0
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 获取的最新交易的区块数量，取值范围1~5
	LatestBlockNumber *uint64 `json:"LatestBlockNumber,omitempty" name:"LatestBlockNumber"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetLatesdTransactionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLatesdTransactionListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLatesdTransactionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 交易列表
		TransactionList []*TransactionItem `json:"TransactionList,omitempty" name:"TransactionList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLatesdTransactionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLatesdTransactionListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLogDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 日志ID
	LogId *uint64 `json:"LogId,omitempty" name:"LogId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLogDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志ID
		Id *int64 `json:"Id,omitempty" name:"Id"`

		// 请求内容
		Req *string `json:"Req,omitempty" name:"Req"`

		// 返回内容
		Rsp *string `json:"Rsp,omitempty" name:"Rsp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLogDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLogListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查看关键词（操作内容）
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 排序方式；asc：正序，desc：倒序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *GetLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLogListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志列表
		Logs []*Log `json:"Logs,omitempty" name:"Logs" list`

		// 日志数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLogListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListForChannelRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *GetPeerListForChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListForChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListForChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点数组
		Peers []*Peer `json:"Peers,omitempty" name:"Peers" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPeerListForChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListForChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListForCloudMonitorRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 偏移项数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页内项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 关键字
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *GetPeerListForCloudMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListForCloudMonitorRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListForCloudMonitorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// peer节点总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// peer节点列表
		PeerList []*PeerForCloudMonitor `json:"PeerList,omitempty" name:"PeerList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPeerListForCloudMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListForCloudMonitorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListForInstallRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`
}

func (r *GetPeerListForInstallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListForInstallRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListForInstallResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点列表
		Peers []*Peer `json:"Peers,omitempty" name:"Peers" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPeerListForInstallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListForInstallResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListPerChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetPeerListPerChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListPerChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListPerChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合约部署节点列表
		Peers []*PeerOfChaincode `json:"Peers,omitempty" name:"Peers" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPeerListPerChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListPerChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListPerChannelRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 0:获取加入通道的所有节点,1:获取加入通道的当前组织的节点
	ByGroup *uint64 `json:"ByGroup,omitempty" name:"ByGroup"`
}

func (r *GetPeerListPerChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListPerChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListPerChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点列表
		Peers []*PeerOfChannel `json:"Peers,omitempty" name:"Peers" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPeerListPerChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListPerChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// cluster标识
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询关键词（组织名称）
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *GetPeerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPeerListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点列表
		Peers []*PeerOfChaincode `json:"Peers,omitempty" name:"Peers" list`

		// 列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPeerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPeerListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPrivateLinkDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 私有连接ID
	LinkId *uint64 `json:"LinkId,omitempty" name:"LinkId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetPrivateLinkDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPrivateLinkDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPrivateLinkDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组织节点名称
		PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

		// Nginx连接pem证书
		NginxPem *string `json:"NginxPem,omitempty" name:"NginxPem"`

		// ssl连接名称
		SSLName *string `json:"SSLName,omitempty" name:"SSLName"`

		// 客户端映射IP
		ClientVip *string `json:"ClientVip,omitempty" name:"ClientVip"`

		// 客户端映射端口
		ClientVport *uint64 `json:"ClientVport,omitempty" name:"ClientVport"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPrivateLinkDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPrivateLinkDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPrivateLinkListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移页数
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 每页项数
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetPrivateLinkListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPrivateLinkListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPrivateLinkListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有连接总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 私有连接列表
		PLinkList []*PLinkItem `json:"PLinkList,omitempty" name:"PLinkList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPrivateLinkListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPrivateLinkListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSrcPrefListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 节点ID
	Sid *string `json:"Sid,omitempty" name:"Sid"`

	// 维度
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetSrcPrefListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSrcPrefListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSrcPrefListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回内容
		SrvPerfs *SrvPerf `json:"SrvPerfs,omitempty" name:"SrvPerfs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSrcPrefListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSrcPrefListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSrvLogDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署名称
	DeployName *string `json:"DeployName,omitempty" name:"DeployName"`

	// 开始时间
	BegintTime *string `json:"BegintTime,omitempty" name:"BegintTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 最大显示条数
	MaxItems *uint64 `json:"MaxItems,omitempty" name:"MaxItems"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetSrvLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSrvLogDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSrvLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志内容
		Log *string `json:"Log,omitempty" name:"Log"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSrvLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSrvLogDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSrvStatusListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 服务类型
	ServiceType []*uint64 `json:"ServiceType,omitempty" name:"ServiceType" list`

	// 服务状态
	ServiceStatus []*uint64 `json:"ServiceStatus,omitempty" name:"ServiceStatus" list`

	// 查询关键词
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetSrvStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSrvStatusListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSrvStatusListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Service列表
		Services []*Service `json:"Services,omitempty" name:"Services" list`

		// Service总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSrvStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSrvStatusListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTbaasClusterListSummaryRequest struct {
	*tchttp.BaseRequest

	// 模块名，默认参数：summary_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，默认参数：tbaas_cluster_list_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 过滤项：'status',筛选网络状态 ；'blockchain_type',筛选网络引擎类型
	Filters []*ClusterFilter `json:"Filters,omitempty" name:"Filters" list`

	// 每页项数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 项数偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetTbaasClusterListSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTbaasClusterListSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTbaasClusterListSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 联盟列表
		ConsortiumList []*SummaryConsortiumItem `json:"ConsortiumList,omitempty" name:"ConsortiumList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTbaasClusterListSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTbaasClusterListSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTbaasClusterNumberSummaryRequest struct {
	*tchttp.BaseRequest

	// 模块名，默认参数：summary_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，默认参数：tbaas_cluster_number_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 引擎类型
	BlockChainTypeList []*uint64 `json:"BlockChainTypeList,omitempty" name:"BlockChainTypeList" list`

	// 地域列表
	RegionList []*uint64 `json:"RegionList,omitempty" name:"RegionList" list`

	// 网络状态筛选列表（0：部署中；1：运行中；2.隔离；3销毁）
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList" list`

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`
}

func (r *GetTbaasClusterNumberSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTbaasClusterNumberSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTbaasClusterNumberSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 运行中网络总数
		TotalClusterNumber *uint64 `json:"TotalClusterNumber,omitempty" name:"TotalClusterNumber"`

		// 各个地域网络数量
		ClusterNumberList []*ClusterNumberPerRegion `json:"ClusterNumberList,omitempty" name:"ClusterNumberList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTbaasClusterNumberSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTbaasClusterNumberSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTbaasClusterSummaryRequest struct {
	*tchttp.BaseRequest

	// 模块名，默认参数：summary_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，默认参数：tbaas_cluster_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *GetTbaasClusterSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTbaasClusterSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTbaasClusterSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络总数
		TotalClusterCount *uint64 `json:"TotalClusterCount,omitempty" name:"TotalClusterCount"`

		// 运行中网络总数（包括部署中和运行中）
		RunningClusterCount *uint64 `json:"RunningClusterCount,omitempty" name:"RunningClusterCount"`

		// 即将到期中网络总数
		ExpiringClusterCount *uint64 `json:"ExpiringClusterCount,omitempty" name:"ExpiringClusterCount"`

		// 过期网络总数（即隔离且未被销毁）
		ExpiredClusterCount *uint64 `json:"ExpiredClusterCount,omitempty" name:"ExpiredClusterCount"`

		// fabric网络总数
		FabricClusterCount *uint64 `json:"FabricClusterCount,omitempty" name:"FabricClusterCount"`

		// trustsql网络总数
		TrustsqlClusterCount *uint64 `json:"TrustsqlClusterCount,omitempty" name:"TrustsqlClusterCount"`

		// bcos网络总数
		BcosClusterCount *uint64 `json:"BcosClusterCount,omitempty" name:"BcosClusterCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTbaasClusterSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTbaasClusterSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTbaasKeySummaryRequest struct {
	*tchttp.BaseRequest

	// 模块名，默认参数：summary_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，默认参数：tbaas_key_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *GetTbaasKeySummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTbaasKeySummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTbaasKeySummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 加入联盟总数
		TotalConsortium *uint64 `json:"TotalConsortium,omitempty" name:"TotalConsortium"`

		// 待办事项总数
		ToDoEventCount *uint64 `json:"ToDoEventCount,omitempty" name:"ToDoEventCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTbaasKeySummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTbaasKeySummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTransactionDetailRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 区块ID
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 交易ID
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
}

func (r *GetTransactionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTransactionDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTransactionDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易ID
		TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

		// 交易hash
		TransactionHash *string `json:"TransactionHash,omitempty" name:"TransactionHash"`

		// 创建交易的组织名
		CreateOrgName *string `json:"CreateOrgName,omitempty" name:"CreateOrgName"`

		// 交易类型（普通交易和配置交易）
		TransactionType *string `json:"TransactionType,omitempty" name:"TransactionType"`

		// 交易创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 交易数据
		TransactionData *string `json:"TransactionData,omitempty" name:"TransactionData"`

		// 交易所在区块号
		BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

		// 交易所在区块哈希
		BlockHash *string `json:"BlockHash,omitempty" name:"BlockHash"`

		// 交易所在区块高度
		BlockHeight *uint64 `json:"BlockHeight,omitempty" name:"BlockHeight"`

		// 通道名称
		ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

		// 交易所在合约名称
		ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

		// 交易背书组织列表
		EndorserOrgList []*EndorserGroup `json:"EndorserOrgList,omitempty" name:"EndorserOrgList" list`

		// 交易状态
		TransactionStatus *string `json:"TransactionStatus,omitempty" name:"TransactionStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTransactionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTransactionDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTransactionSpeedRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetTransactionSpeedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTransactionSpeedRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTransactionSpeedResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易速度曲线数据
		Speeds []*TransactionSpeed `json:"Speeds,omitempty" name:"Speeds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTransactionSpeedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTransactionSpeedResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserAuthTypeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 用户ID
	UinId *uint64 `json:"UinId,omitempty" name:"UinId"`
}

func (r *GetUserAuthTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserAuthTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserAuthTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户认证类型
		AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserAuthTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserAuthTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserVpcListRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetUserVpcListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserVpcListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserVpcListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// vpc总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 用户vpc列表
		VpcList []*VpcItem `json:"VpcList,omitempty" name:"VpcList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserVpcListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserVpcListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GoodsDetailInfo struct {

	// 网络名称
	TbaasName *string `json:"TbaasName,omitempty" name:"TbaasName"`

	// 网络根域
	TbaasDomain *string `json:"TbaasDomain,omitempty" name:"TbaasDomain"`

	// 证书CA类型：0，fabric-ca；1，cfca
	CaType *uint64 `json:"CaType,omitempty" name:"CaType"`

	// 证书加密算法：0，'ECC'；1，'SM2'
	KeyAlg *uint64 `json:"KeyAlg,omitempty" name:"KeyAlg"`

	// 组织名称
	OrgNames *string `json:"OrgNames,omitempty" name:"OrgNames"`

	// 网络类型：0，私有链；1，联盟链
	TbaasType *uint64 `json:"TbaasType,omitempty" name:"TbaasType"`

	// 网络创建标识：1，创建网络；0，加入网络
	Creator *uint64 `json:"Creator,omitempty" name:"Creator"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 网络所属联盟名称
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 共识算法
	Consensus *string `json:"Consensus,omitempty" name:"Consensus"`

	// 网络描述
	ClusterInfo *string `json:"ClusterInfo,omitempty" name:"ClusterInfo"`

	// 网络任务详情
	TaskJson *string `json:"TaskJson,omitempty" name:"TaskJson"`

	// 商品ID
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 网络排序节点数量
	OrderNumber *uint64 `json:"OrderNumber,omitempty" name:"OrderNumber"`

	// 网络组织数量
	TbaasOrg *uint64 `json:"TbaasOrg,omitempty" name:"TbaasOrg"`

	// 网络节点数量
	TbaasPeer *uint64 `json:"TbaasPeer,omitempty" name:"TbaasPeer"`

	// 网络Cvm1数量
	TbaasCvm1 *uint64 `json:"TbaasCvm1,omitempty" name:"TbaasCvm1"`

	// 网络Cvm2数量
	TbaasCvm2 *uint64 `json:"TbaasCvm2,omitempty" name:"TbaasCvm2"`

	// 网络Cvm3数量
	TbaasCvm3 *uint64 `json:"TbaasCvm3,omitempty" name:"TbaasCvm3"`

	// 网络Cdh1数量
	TbaasCdh1 *uint64 `json:"TbaasCdh1,omitempty" name:"TbaasCdh1"`

	// 网络节点数据盘总量
	TbaasPeerCbs *uint64 `json:"TbaasPeerCbs,omitempty" name:"TbaasPeerCbs"`

	// 网络排序节点数据盘总量
	TbaasOrderCbs *uint64 `json:"TbaasOrderCbs,omitempty" name:"TbaasOrderCbs"`

	// 网络企业证书数量
	TbaasCompanyCert *uint64 `json:"TbaasCompanyCert,omitempty" name:"TbaasCompanyCert"`

	// 网络个人证书数量
	TbaasPersonCert *uint64 `json:"TbaasPersonCert,omitempty" name:"TbaasPersonCert"`

	// 购买时长
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 购买时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 商品展示信息
	ProductInfo []*ProductInfoItem `json:"ProductInfo,omitempty" name:"ProductInfo" list`
}

type GoodsDetailItem struct {

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 到期时间
	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`

	// 购买时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 购买时长
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 产品的PID
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// tbaas组织数量
	TbaasOrg *uint64 `json:"TbaasOrg,omitempty" name:"TbaasOrg"`

	// tbaas节点数量
	TbaasPeer *uint64 `json:"TbaasPeer,omitempty" name:"TbaasPeer"`

	// 排序节点数量
	OrderNumber *uint64 `json:"OrderNumber,omitempty" name:"OrderNumber"`

	// cvm_1数量
	TbaasCvm1 *uint64 `json:"TbaasCvm1,omitempty" name:"TbaasCvm1"`

	// cvm_2数量
	TbaasCvm2 *uint64 `json:"TbaasCvm2,omitempty" name:"TbaasCvm2"`

	// cvm_3数量
	TbaasCvm3 *uint64 `json:"TbaasCvm3,omitempty" name:"TbaasCvm3"`

	// cdh数量
	TbaasCdh1 *uint64 `json:"TbaasCdh1,omitempty" name:"TbaasCdh1"`

	// tbaas节点的cbs总量
	TbaasPeerCbs *uint64 `json:"TbaasPeerCbs,omitempty" name:"TbaasPeerCbs"`

	// tbaas排序节点的cbs总量
	TbaasOrderCbs *uint64 `json:"TbaasOrderCbs,omitempty" name:"TbaasOrderCbs"`

	// vpcID
	UniqVpcid *string `json:"UniqVpcid,omitempty" name:"UniqVpcid"`

	// vpc网段
	VpcCidr *string `json:"VpcCidr,omitempty" name:"VpcCidr"`

	// tbaas名称
	TbaasName *string `json:"TbaasName,omitempty" name:"TbaasName"`

	// tbaas根域
	TbaasDomain *string `json:"TbaasDomain,omitempty" name:"TbaasDomain"`

	// 组织名称
	OrgNames *string `json:"OrgNames,omitempty" name:"OrgNames"`

	// 区块链类型（0：私有链；1,：联盟链）
	TbaasType *uint64 `json:"TbaasType,omitempty" name:"TbaasType"`

	// 创建标识（0：加入；1：创建）
	Creator *uint64 `json:"Creator,omitempty" name:"Creator"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 联盟ID
	ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 共识算法
	Consensus *string `json:"Consensus,omitempty" name:"Consensus"`

	// 集群信息
	ClusterInfo *string `json:"ClusterInfo,omitempty" name:"ClusterInfo"`

	// CFCA企业证书数量
	TbaasCompanyCert *uint64 `json:"TbaasCompanyCert,omitempty" name:"TbaasCompanyCert"`

	// CFCA个人证书数量
	TbaasPersonCert *uint64 `json:"TbaasPersonCert,omitempty" name:"TbaasPersonCert"`

	// 网络CA类型：0，“tbaas”;1，“cfca”
	CaType *uint64 `json:"CaType,omitempty" name:"CaType"`

	// 网络签名算法：0 ECC， 1 SM2
	KeyAlg *uint64 `json:"KeyAlg,omitempty" name:"KeyAlg"`

	// cvm类型ID
	CvmTypeId *string `json:"CvmTypeId,omitempty" name:"CvmTypeId"`

	// cvm型号
	CvmTypeCn *string `json:"CvmTypeCn,omitempty" name:"CvmTypeCn"`

	// cvm数量
	CvmNumber *uint64 `json:"CvmNumber,omitempty" name:"CvmNumber"`

	// 每个peer的cbs大小
	CbsPerPeer *uint64 `json:"CbsPerPeer,omitempty" name:"CbsPerPeer"`

	// 产品信息列表
	ProductInfo []*ProductInfoItem `json:"ProductInfo,omitempty" name:"ProductInfo" list`
}

type Group struct {

	// 网络集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 长名称
	LongName *string `json:"LongName,omitempty" name:"LongName"`

	// mspID
	Mspid *string `json:"Mspid,omitempty" name:"Mspid"`

	// 短名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

	// 组织状态：1，正常；0，已销毁
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type GroupForCloudMonitor struct {

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 组织的peer节点数量
	PeerCount *uint64 `json:"PeerCount,omitempty" name:"PeerCount"`

	// 组织加入的通道数量
	ChannelCount *uint64 `json:"ChannelCount,omitempty" name:"ChannelCount"`
}

type GroupOfChaincode struct {

	// 网络集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 该组织是否已安装当前合约标识
	HasInstall *uint64 `json:"HasInstall,omitempty" name:"HasInstall"`

	// 组织ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 组名名称
	LongName *string `json:"LongName,omitempty" name:"LongName"`

	// 组名mspid
	Mspid *string `json:"Mspid,omitempty" name:"Mspid"`

	// 组名名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

	// 组织状态：1，正常；0，已销毁
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type GroupOfCluster struct {

	// 是否创建者标识
	IsCreator *bool `json:"IsCreator,omitempty" name:"IsCreator"`

	// 网络标识
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 组织名称
	LongName *string `json:"LongName,omitempty" name:"LongName"`

	// 组织Mspid
	Mspid *string `json:"Mspid,omitempty" name:"Mspid"`

	// 节点List
	PeerList *string `json:"PeerList,omitempty" name:"PeerList"`

	// 组织名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

	// 组织状态：1，正常；0，已销毁
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type InitializeChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 参数
	Args *string `json:"Args,omitempty" name:"Args"`

	// 背书
	Endorsement *string `json:"Endorsement,omitempty" name:"Endorsement"`

	// 背书ID
	EndorsementId *string `json:"EndorsementId,omitempty" name:"EndorsementId"`

	// 私有数据集列表
	PrivateDataList []*PrivateData `json:"PrivateDataList,omitempty" name:"PrivateDataList" list`
}

func (r *InitializeChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitializeChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitializeChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 超时为1，未超时为0
		TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

		// 任务ID，当Timeout为1时该值才有意义
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitializeChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitializeChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstallChaincodeRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 合约ID
	UserChaincodesId *string `json:"UserChaincodesId,omitempty" name:"UserChaincodesId"`

	// 节点ID
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`

	// 节点ID列表
	PeerList []*string `json:"PeerList,omitempty" name:"PeerList" list`
}

func (r *InstallChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstallChaincodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstallChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstallChaincodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InviteChannelForCBCRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 邀请组织列表
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList" list`

	// 默认是否同意
	Option *uint64 `json:"Option,omitempty" name:"Option"`

	// 邀请原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *InviteChannelForCBCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InviteChannelForCBCRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InviteChannelForCBCResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InviteChannelForCBCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InviteChannelForCBCResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InviteClusterMemberRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 邀请原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 邀请的联盟成员id列表
	ConsortiumMemberIds []*string `json:"ConsortiumMemberIds,omitempty" name:"ConsortiumMemberIds" list`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 邀请湖用户信息，非用户
	InviteeInfo *string `json:"InviteeInfo,omitempty" name:"InviteeInfo"`
}

func (r *InviteClusterMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InviteClusterMemberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InviteClusterMemberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 事件配置信息列表
		InviteEventList []*EventConfigInfo `json:"InviteEventList,omitempty" name:"InviteEventList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InviteClusterMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InviteClusterMemberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InviteConsortiumMemberRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 成员名称
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// 成员账号
	MemberUin *string `json:"MemberUin,omitempty" name:"MemberUin"`

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 邀请成员Appid
	MemberAppId *string `json:"MemberAppId,omitempty" name:"MemberAppId"`

	// 成员电话
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 成员邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 邀请原因
	InviteReason *string `json:"InviteReason,omitempty" name:"InviteReason"`
}

func (r *InviteConsortiumMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InviteConsortiumMemberRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InviteConsortiumMemberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InviteConsortiumMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InviteConsortiumMemberResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InvokeRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：invoke
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 对该笔交易进行背书的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称极其所属组织名称
	Peers []*PeerSet `json:"Peers,omitempty" name:"Peers" list`

	// 该笔交易需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 被调用的函数参数列表
	Args []*string `json:"Args,omitempty" name:"Args" list`

	// 同步调用标识，可选参数，值为0或者不传表示使用同步方法调用，调用后会等待交易执行后再返回执行结果；值为1时表示使用异步方式调用Invoke，执行后会立即返回交易对应的Txid，后续需要通过GetInvokeTx这个API查询该交易的执行结果。（对于逻辑较为简单的交易，可以使用同步模式；对于逻辑较为复杂的交易，建议使用异步模式，否则容易导致API因等待时间过长，返回等待超时）
	AsyncFlag *uint64 `json:"AsyncFlag,omitempty" name:"AsyncFlag"`
}

func (r *InvokeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InvokeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InvokeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易ID
		Txid *string `json:"Txid,omitempty" name:"Txid"`

		// 交易执行结果
		Events *string `json:"Events,omitempty" name:"Events"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InvokeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InvokeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type JoinChannelRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 节点ID
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`

	// 节点ID列表
	PeerList []*string `json:"PeerList,omitempty" name:"PeerList" list`
}

func (r *JoinChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinChannelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type JoinChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinChannelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Log struct {

	// 日志ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 操作类型
	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`

	// 操作结果
	OpResult *string `json:"OpResult,omitempty" name:"OpResult"`

	// 操作目标
	OpTarget *string `json:"OpTarget,omitempty" name:"OpTarget"`

	// 操作时间
	OpTime *string `json:"OpTime,omitempty" name:"OpTime"`

	// 操作Uin
	OpUin *int64 `json:"OpUin,omitempty" name:"OpUin"`

	// 操作内容
	OperationContent *string `json:"OperationContent,omitempty" name:"OperationContent"`
}

type ModifyChannelVoteRateForCBCRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 当前投票率
	CurrentRate *uint64 `json:"CurrentRate,omitempty" name:"CurrentRate"`

	// 修改后投票率
	NewRate *uint64 `json:"NewRate,omitempty" name:"NewRate"`

	// 默认是否投票同意
	Option *uint64 `json:"Option,omitempty" name:"Option"`

	// 修改原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyChannelVoteRateForCBCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyChannelVoteRateForCBCRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyChannelVoteRateForCBCResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyChannelVoteRateForCBCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyChannelVoteRateForCBCResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterDescriptionRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 网络描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyClusterDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFabricClusterRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用区ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 增配前配置
	OldConfig *GoodsDetailInfo `json:"OldConfig,omitempty" name:"OldConfig"`

	// 增配后配置
	NewConfig *GoodsDetailInfo `json:"NewConfig,omitempty" name:"NewConfig"`
}

func (r *ModifyFabricClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFabricClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFabricClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号列表
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFabricClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFabricClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFieldDescriptionRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 描述类型：0-联盟，1-网络，2-通道，3-策略
	FieldType *uint64 `json:"FieldType,omitempty" name:"FieldType"`

	// 描述类型ID，字符串类型
	FieldId *string `json:"FieldId,omitempty" name:"FieldId"`

	// 描述内容
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyFieldDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFieldDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFieldDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFieldDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFieldDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateLinkRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 私有连接ID
	LinkId *uint64 `json:"LinkId,omitempty" name:"LinkId"`

	// 修改名称
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyPrivateLinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPrivateLinkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateLinkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPrivateLinkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PLinkItem struct {

	// 私有连接ID
	PLinkId *uint64 `json:"PLinkId,omitempty" name:"PLinkId"`

	// 私有连接名称
	PLinkName *string `json:"PLinkName,omitempty" name:"PLinkName"`

	// 私有连接状态：1，已连接；2，不可用
	PLinkStatus *uint64 `json:"PLinkStatus,omitempty" name:"PLinkStatus"`

	// 私有连接所属区域
	VisitRegion *uint64 `json:"VisitRegion,omitempty" name:"VisitRegion"`

	// 客户端vpc统一ID
	ClientVpcId *string `json:"ClientVpcId,omitempty" name:"ClientVpcId"`

	// 客户端vpc中子网统一ID
	ClientSubNetId *string `json:"ClientSubNetId,omitempty" name:"ClientSubNetId"`

	// 客户端映射ip
	ClientVip *string `json:"ClientVip,omitempty" name:"ClientVip"`

	// 客户端映射port
	ClientVport *uint64 `json:"ClientVport,omitempty" name:"ClientVport"`

	// 用户vpc名称
	ClientVpcName *string `json:"ClientVpcName,omitempty" name:"ClientVpcName"`

	// 用户子网名称
	ClientSubNetName *string `json:"ClientSubNetName,omitempty" name:"ClientSubNetName"`
}

type Peer struct {

	// 节点ID
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`

	// 节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`
}

type PeerForCloudMonitor struct {

	// peer节点ID
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`

	// peer节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// peer节点所属组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// peer节点所在网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type PeerOfChaincode struct {

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组名名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 节点ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 节点名称
	LongName *string `json:"LongName,omitempty" name:"LongName"`

	// 节点名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`
}

type PeerOfChannel struct {

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组名名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 节点ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 节点名称
	LongName *string `json:"LongName,omitempty" name:"LongName"`

	// 节点Mspid
	Mspid *string `json:"Mspid,omitempty" name:"Mspid"`

	// 节点请求
	Requests *string `json:"Requests,omitempty" name:"Requests"`

	// 节点名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

	// 是否为该通道的anchor节点
	Anchor *uint64 `json:"Anchor,omitempty" name:"Anchor"`
}

type PeerSet struct {

	// 节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 组织名称
	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
}

type PreFeeGetModifyPriceRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 组织个数
	TbaasOrg *int64 `json:"TbaasOrg,omitempty" name:"TbaasOrg"`

	// 总的peer节点数
	TbaasPeer *int64 `json:"TbaasPeer,omitempty" name:"TbaasPeer"`

	// peer节点CBS总容量（GB）
	TbaasPeerCbs *int64 `json:"TbaasPeerCbs,omitempty" name:"TbaasPeerCbs"`

	// Order节点CBS总容量（GB）
	TbaasOrderCbs *int64 `json:"TbaasOrderCbs,omitempty" name:"TbaasOrderCbs"`

	// 型号1的cvm的个数
	TbaasCvm1 *uint64 `json:"TbaasCvm1,omitempty" name:"TbaasCvm1"`

	// 型号2的cvm的个数
	TbaasCvm2 *uint64 `json:"TbaasCvm2,omitempty" name:"TbaasCvm2"`

	// 型号3的cvm的个数
	TbaasCvm3 *uint64 `json:"TbaasCvm3,omitempty" name:"TbaasCvm3"`

	// 型号1的cdh的个数
	TbaasCdh1 *uint64 `json:"TbaasCdh1,omitempty" name:"TbaasCdh1"`

	// 企业证书数量
	TbaasCompanyCert *uint64 `json:"TbaasCompanyCert,omitempty" name:"TbaasCompanyCert"`

	// 个人证书数量
	TbaasPersonCert *uint64 `json:"TbaasPersonCert,omitempty" name:"TbaasPersonCert"`

	// 最新版本固定值为：'v2.1'
	TbaasVersion *string `json:"TbaasVersion,omitempty" name:"TbaasVersion"`
}

func (r *PreFeeGetModifyPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreFeeGetModifyPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PreFeeGetModifyPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 变配计费返回
		FeeModifyPrices []*FeeModifyPrice `json:"FeeModifyPrices,omitempty" name:"FeeModifyPrices" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreFeeGetModifyPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreFeeGetModifyPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PreFeeGetPriceRequest struct {
	*tchttp.BaseRequest

	// 组织个数
	TbaasOrg *uint64 `json:"TbaasOrg,omitempty" name:"TbaasOrg"`

	// 总的peer节点数
	TbaasPeer *uint64 `json:"TbaasPeer,omitempty" name:"TbaasPeer"`

	// peer节点CBS总容量（GB）
	TbaasPeerCbs *uint64 `json:"TbaasPeerCbs,omitempty" name:"TbaasPeerCbs"`

	// Order节点CBS总容量（GB）
	TbaasOrderCbs *uint64 `json:"TbaasOrderCbs,omitempty" name:"TbaasOrderCbs"`

	// 付费单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 付费时长
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 商品ID
	GoodsCategoryId *uint64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`

	// Order数量
	OrderNumber *uint64 `json:"OrderNumber,omitempty" name:"OrderNumber"`

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 型号1的cvm的个数
	TbaasCvm1 *uint64 `json:"TbaasCvm1,omitempty" name:"TbaasCvm1"`

	// 型号2的cvm的个数
	TbaasCvm2 *uint64 `json:"TbaasCvm2,omitempty" name:"TbaasCvm2"`

	// 型号3的cvm的个数
	TbaasCvm3 *uint64 `json:"TbaasCvm3,omitempty" name:"TbaasCvm3"`

	// 型号1的cdh的个数
	TbaasCdh1 *uint64 `json:"TbaasCdh1,omitempty" name:"TbaasCdh1"`

	// 企业证书数量
	TbaasCompanyCert *uint64 `json:"TbaasCompanyCert,omitempty" name:"TbaasCompanyCert"`

	// 个人证书数量
	TbaasPersonCert *uint64 `json:"TbaasPersonCert,omitempty" name:"TbaasPersonCert"`

	// 最新前版本固定值为'v2.1'
	TbaasVersion *string `json:"TbaasVersion,omitempty" name:"TbaasVersion"`
}

func (r *PreFeeGetPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreFeeGetPriceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PreFeeGetPriceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 价格
		PreFeePrices []*FeePrice `json:"PreFeePrices,omitempty" name:"PreFeePrices" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreFeeGetPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreFeeGetPriceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PrivateData struct {

	// 私有数据集ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 私有数据集名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 私有数据集授权组织名称列表
	Policy []*string `json:"Policy,omitempty" name:"Policy" list`

	// 最小分发数量
	RequiredPeerCount *uint64 `json:"RequiredPeerCount,omitempty" name:"RequiredPeerCount"`

	// 最大分发数量
	MaxPeerCount *uint64 `json:"MaxPeerCount,omitempty" name:"MaxPeerCount"`

	// 保存时间
	BlockToLive *uint64 `json:"BlockToLive,omitempty" name:"BlockToLive"`

	// 是否公开（0：不公开，1：公开）
	MemberOnlyRead *uint64 `json:"MemberOnlyRead,omitempty" name:"MemberOnlyRead"`

	// 描述备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ProductInfoItem struct {

	// 属性名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type QueryOnlineEditorAccessRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *QueryOnlineEditorAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryOnlineEditorAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryOnlineEditorAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否允许访问：0，不允许，1，允许。
		Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOnlineEditorAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryOnlineEditorAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryPlatFormResourceRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *QueryPlatFormResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryPlatFormResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryPlatFormResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// tbass平台类型列表
		PlatformList []*uint64 `json:"PlatformList,omitempty" name:"PlatformList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPlatFormResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryPlatFormResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：query
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称极其所属组织名称
	Peers []*PeerSet `json:"Peers,omitempty" name:"Peers" list`

	// 该笔交易查询需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 被调用的函数参数列表
	Args []*string `json:"Args,omitempty" name:"Args" list`
}

func (r *QueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询结果数据
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReissueCertRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 证书ID
	CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

	// 证书DN
	CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

	// csr文件
	CsrData *string `json:"CsrData,omitempty" name:"CsrData"`

	// 当前组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ReissueCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReissueCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReissueCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书名称
		CertName *string `json:"CertName,omitempty" name:"CertName"`

		// 证书内容
		CertCtx *string `json:"CertCtx,omitempty" name:"CertCtx"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReissueCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReissueCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunChaincodeDevRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 操作类型：invoke、query和init
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 函数参数（base64）
	Args *string `json:"Args,omitempty" name:"Args"`
}

func (r *RunChaincodeDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunChaincodeDevRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunChaincodeDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果状态码
		Rescode *uint64 `json:"Rescode,omitempty" name:"Rescode"`

		// 结果内容
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunChaincodeDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunChaincodeDevResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Service struct {

	// cbs空闲大小
	CbsFree *uint64 `json:"CbsFree,omitempty" name:"CbsFree"`

	// cbs总数大小
	CbsTotal *uint64 `json:"CbsTotal,omitempty" name:"CbsTotal"`

	// 部署名称
	DeployName *string `json:"DeployName,omitempty" name:"DeployName"`

	// ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Service类型
	SrvType *uint64 `json:"SrvType,omitempty" name:"SrvType"`

	// Service状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 当前组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type SetChannelAnchorPeerRequest struct {
	*tchttp.BaseRequest

	// 模块名称
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// peer节点ID列表
	PeerList []*string `json:"PeerList,omitempty" name:"PeerList" list`
}

func (r *SetChannelAnchorPeerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetChannelAnchorPeerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetChannelAnchorPeerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetChannelAnchorPeerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetChannelAnchorPeerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SrvPerf struct {

	// X轴坐标内容
	XAxis []*string `json:"XAxis,omitempty" name:"XAxis" list`

	// Y轴坐标内容
	YAxis []*YAxisObj `json:"YAxis,omitempty" name:"YAxis" list`
}

type SubNetItem struct {

	// vpc子网 ID
	SubNetId *uint64 `json:"SubNetId,omitempty" name:"SubNetId"`

	// vpc子网统一ID
	UniqSubNetId *string `json:"UniqSubNetId,omitempty" name:"UniqSubNetId"`

	// vpc子网名称
	SubNetName *string `json:"SubNetName,omitempty" name:"SubNetName"`
}

type SummaryClusterItem struct {

	// 网络名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 区块链引擎类型（0：fabric，1：trustsql，2：bcos）
	BlockchainType *uint64 `json:"BlockchainType,omitempty" name:"BlockchainType"`

	// 网络所属联盟ID
	ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 网络所属地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 网络状态（1：运行中， 2：即将到期，3：已过期）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 网络组织数量
	ClusterGroupCount *uint64 `json:"ClusterGroupCount,omitempty" name:"ClusterGroupCount"`

	// 网络节点数量
	ClusterPeerCount *uint64 `json:"ClusterPeerCount,omitempty" name:"ClusterPeerCount"`

	// 我的节点数量
	MyPeerCount *uint64 `json:"MyPeerCount,omitempty" name:"MyPeerCount"`
}

type SummaryConsortiumItem struct {

	// 联盟ID
	ConsortiumId *string `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 联盟名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 网络列表
	ClusterList []*SummaryClusterItem `json:"ClusterList,omitempty" name:"ClusterList" list`
}

type TagItem struct {

	// 标签的键名称
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签的键值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TransactionItem struct {

	// 交易ID
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 交易hash
	TransactionHash *string `json:"TransactionHash,omitempty" name:"TransactionHash"`

	// 创建交易的组织名
	CreateOrgName *string `json:"CreateOrgName,omitempty" name:"CreateOrgName"`

	// 交易所在区块号
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 交易类型（普通交易和配置交易）
	TransactionType *string `json:"TransactionType,omitempty" name:"TransactionType"`

	// 交易创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 交易所在区块高度
	BlockHeight *uint64 `json:"BlockHeight,omitempty" name:"BlockHeight"`

	// 交易状态
	TransactionStatus *string `json:"TransactionStatus,omitempty" name:"TransactionStatus"`
}

type TransactionSpeed struct {

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 速率
	Rate *float64 `json:"Rate,omitempty" name:"Rate"`
}

type UploadChaincodeDevRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 文件列表，里面为json格式
	Files []*UploadFileObj `json:"Files,omitempty" name:"Files" list`

	// 操作类型：init、del、add
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *UploadChaincodeDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadChaincodeDevRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadChaincodeDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadChaincodeDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadChaincodeDevResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadEventConfigRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 事件配置内容
	EventConfig *string `json:"EventConfig,omitempty" name:"EventConfig"`
}

func (r *UploadEventConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadEventConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadEventConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件ID
		EventId *string `json:"EventId,omitempty" name:"EventId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadEventConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadEventConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadFileObj struct {

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 是否使用cos上传
	CosFlag *uint64 `json:"CosFlag,omitempty" name:"CosFlag"`

	// 文件内容
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 文件hash
	FileHash *string `json:"FileHash,omitempty" name:"FileHash"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type VpcItem struct {

	// vpc ID
	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`

	// vpc统一ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// vpc子网列表
	SubNetList []*SubNetItem `json:"SubNetList,omitempty" name:"SubNetList" list`

	// vpc子网名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type YAxisObj struct {

	// 数值list
	Data []*float64 `json:"Data,omitempty" name:"Data" list`

	// 对象名称
	Name *string `json:"Name,omitempty" name:"Name"`
}
