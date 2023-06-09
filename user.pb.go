// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wx-miniprogram/user.proto

package wx_miniprogram // import "github.com/dev-openapi/wx-miniprogram"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPluginOpenPIdReq struct {
	AccessToken          string                    `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Body                 *GetPluginOpenPIdReq_Body `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetPluginOpenPIdReq) Reset()         { *m = GetPluginOpenPIdReq{} }
func (m *GetPluginOpenPIdReq) String() string { return proto.CompactTextString(m) }
func (*GetPluginOpenPIdReq) ProtoMessage()    {}
func (*GetPluginOpenPIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{0}
}
func (m *GetPluginOpenPIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginOpenPIdReq.Unmarshal(m, b)
}
func (m *GetPluginOpenPIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginOpenPIdReq.Marshal(b, m, deterministic)
}
func (dst *GetPluginOpenPIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginOpenPIdReq.Merge(dst, src)
}
func (m *GetPluginOpenPIdReq) XXX_Size() int {
	return xxx_messageInfo_GetPluginOpenPIdReq.Size(m)
}
func (m *GetPluginOpenPIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginOpenPIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginOpenPIdReq proto.InternalMessageInfo

func (m *GetPluginOpenPIdReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetPluginOpenPIdReq) GetBody() *GetPluginOpenPIdReq_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type GetPluginOpenPIdReq_Body struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginOpenPIdReq_Body) Reset()         { *m = GetPluginOpenPIdReq_Body{} }
func (m *GetPluginOpenPIdReq_Body) String() string { return proto.CompactTextString(m) }
func (*GetPluginOpenPIdReq_Body) ProtoMessage()    {}
func (*GetPluginOpenPIdReq_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{0, 0}
}
func (m *GetPluginOpenPIdReq_Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginOpenPIdReq_Body.Unmarshal(m, b)
}
func (m *GetPluginOpenPIdReq_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginOpenPIdReq_Body.Marshal(b, m, deterministic)
}
func (dst *GetPluginOpenPIdReq_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginOpenPIdReq_Body.Merge(dst, src)
}
func (m *GetPluginOpenPIdReq_Body) XXX_Size() int {
	return xxx_messageInfo_GetPluginOpenPIdReq_Body.Size(m)
}
func (m *GetPluginOpenPIdReq_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginOpenPIdReq_Body.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginOpenPIdReq_Body proto.InternalMessageInfo

func (m *GetPluginOpenPIdReq_Body) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type GetPluginOpenPidRes struct {
	Errcode              int32    `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	Openpid              string   `protobuf:"bytes,3,opt,name=openpid" json:"openpid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginOpenPidRes) Reset()         { *m = GetPluginOpenPidRes{} }
func (m *GetPluginOpenPidRes) String() string { return proto.CompactTextString(m) }
func (*GetPluginOpenPidRes) ProtoMessage()    {}
func (*GetPluginOpenPidRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{1}
}
func (m *GetPluginOpenPidRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginOpenPidRes.Unmarshal(m, b)
}
func (m *GetPluginOpenPidRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginOpenPidRes.Marshal(b, m, deterministic)
}
func (dst *GetPluginOpenPidRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginOpenPidRes.Merge(dst, src)
}
func (m *GetPluginOpenPidRes) XXX_Size() int {
	return xxx_messageInfo_GetPluginOpenPidRes.Size(m)
}
func (m *GetPluginOpenPidRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginOpenPidRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginOpenPidRes proto.InternalMessageInfo

func (m *GetPluginOpenPidRes) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *GetPluginOpenPidRes) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *GetPluginOpenPidRes) GetOpenpid() string {
	if m != nil {
		return m.Openpid
	}
	return ""
}

type CheckEncryptedDataReq struct {
	AccessToken          string                      `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Body                 *CheckEncryptedDataReq_Body `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CheckEncryptedDataReq) Reset()         { *m = CheckEncryptedDataReq{} }
func (m *CheckEncryptedDataReq) String() string { return proto.CompactTextString(m) }
func (*CheckEncryptedDataReq) ProtoMessage()    {}
func (*CheckEncryptedDataReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{2}
}
func (m *CheckEncryptedDataReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckEncryptedDataReq.Unmarshal(m, b)
}
func (m *CheckEncryptedDataReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckEncryptedDataReq.Marshal(b, m, deterministic)
}
func (dst *CheckEncryptedDataReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckEncryptedDataReq.Merge(dst, src)
}
func (m *CheckEncryptedDataReq) XXX_Size() int {
	return xxx_messageInfo_CheckEncryptedDataReq.Size(m)
}
func (m *CheckEncryptedDataReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckEncryptedDataReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckEncryptedDataReq proto.InternalMessageInfo

func (m *CheckEncryptedDataReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CheckEncryptedDataReq) GetBody() *CheckEncryptedDataReq_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type CheckEncryptedDataReq_Body struct {
	EncryptedMsgHash     string   `protobuf:"bytes,1,opt,name=encrypted_msg_hash,json=encryptedMsgHash" json:"encrypted_msg_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckEncryptedDataReq_Body) Reset()         { *m = CheckEncryptedDataReq_Body{} }
func (m *CheckEncryptedDataReq_Body) String() string { return proto.CompactTextString(m) }
func (*CheckEncryptedDataReq_Body) ProtoMessage()    {}
func (*CheckEncryptedDataReq_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{2, 0}
}
func (m *CheckEncryptedDataReq_Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckEncryptedDataReq_Body.Unmarshal(m, b)
}
func (m *CheckEncryptedDataReq_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckEncryptedDataReq_Body.Marshal(b, m, deterministic)
}
func (dst *CheckEncryptedDataReq_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckEncryptedDataReq_Body.Merge(dst, src)
}
func (m *CheckEncryptedDataReq_Body) XXX_Size() int {
	return xxx_messageInfo_CheckEncryptedDataReq_Body.Size(m)
}
func (m *CheckEncryptedDataReq_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckEncryptedDataReq_Body.DiscardUnknown(m)
}

var xxx_messageInfo_CheckEncryptedDataReq_Body proto.InternalMessageInfo

func (m *CheckEncryptedDataReq_Body) GetEncryptedMsgHash() string {
	if m != nil {
		return m.EncryptedMsgHash
	}
	return ""
}

type CheckEncryptedDataRes struct {
	Errcode              int32    `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	Valid                bool     `protobuf:"varint,3,opt,name=valid" json:"valid,omitempty"`
	CreateTime           int64    `protobuf:"varint,4,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckEncryptedDataRes) Reset()         { *m = CheckEncryptedDataRes{} }
func (m *CheckEncryptedDataRes) String() string { return proto.CompactTextString(m) }
func (*CheckEncryptedDataRes) ProtoMessage()    {}
func (*CheckEncryptedDataRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{3}
}
func (m *CheckEncryptedDataRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckEncryptedDataRes.Unmarshal(m, b)
}
func (m *CheckEncryptedDataRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckEncryptedDataRes.Marshal(b, m, deterministic)
}
func (dst *CheckEncryptedDataRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckEncryptedDataRes.Merge(dst, src)
}
func (m *CheckEncryptedDataRes) XXX_Size() int {
	return xxx_messageInfo_CheckEncryptedDataRes.Size(m)
}
func (m *CheckEncryptedDataRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckEncryptedDataRes.DiscardUnknown(m)
}

var xxx_messageInfo_CheckEncryptedDataRes proto.InternalMessageInfo

func (m *CheckEncryptedDataRes) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *CheckEncryptedDataRes) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *CheckEncryptedDataRes) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *CheckEncryptedDataRes) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type GetPaidUnionidReq struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Openid               string   `protobuf:"bytes,2,opt,name=openid" json:"openid,omitempty"`
	TransactionId        string   `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	MchId                string   `protobuf:"bytes,4,opt,name=mch_id,json=mchId" json:"mch_id,omitempty"`
	OutTradeNo           string   `protobuf:"bytes,5,opt,name=out_trade_no,json=outTradeNo" json:"out_trade_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPaidUnionidReq) Reset()         { *m = GetPaidUnionidReq{} }
func (m *GetPaidUnionidReq) String() string { return proto.CompactTextString(m) }
func (*GetPaidUnionidReq) ProtoMessage()    {}
func (*GetPaidUnionidReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{4}
}
func (m *GetPaidUnionidReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaidUnionidReq.Unmarshal(m, b)
}
func (m *GetPaidUnionidReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaidUnionidReq.Marshal(b, m, deterministic)
}
func (dst *GetPaidUnionidReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaidUnionidReq.Merge(dst, src)
}
func (m *GetPaidUnionidReq) XXX_Size() int {
	return xxx_messageInfo_GetPaidUnionidReq.Size(m)
}
func (m *GetPaidUnionidReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaidUnionidReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaidUnionidReq proto.InternalMessageInfo

func (m *GetPaidUnionidReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetPaidUnionidReq) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *GetPaidUnionidReq) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *GetPaidUnionidReq) GetMchId() string {
	if m != nil {
		return m.MchId
	}
	return ""
}

func (m *GetPaidUnionidReq) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

type GetPaidUnionidRes struct {
	Errcode              int32    `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	Unionid              string   `protobuf:"bytes,3,opt,name=unionid" json:"unionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPaidUnionidRes) Reset()         { *m = GetPaidUnionidRes{} }
func (m *GetPaidUnionidRes) String() string { return proto.CompactTextString(m) }
func (*GetPaidUnionidRes) ProtoMessage()    {}
func (*GetPaidUnionidRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{5}
}
func (m *GetPaidUnionidRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaidUnionidRes.Unmarshal(m, b)
}
func (m *GetPaidUnionidRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaidUnionidRes.Marshal(b, m, deterministic)
}
func (dst *GetPaidUnionidRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaidUnionidRes.Merge(dst, src)
}
func (m *GetPaidUnionidRes) XXX_Size() int {
	return xxx_messageInfo_GetPaidUnionidRes.Size(m)
}
func (m *GetPaidUnionidRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaidUnionidRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaidUnionidRes proto.InternalMessageInfo

func (m *GetPaidUnionidRes) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *GetPaidUnionidRes) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *GetPaidUnionidRes) GetUnionid() string {
	if m != nil {
		return m.Unionid
	}
	return ""
}

type GetUserEncryptKeyReq struct {
	AccessToken          string                     `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Body                 *GetUserEncryptKeyReq_Body `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetUserEncryptKeyReq) Reset()         { *m = GetUserEncryptKeyReq{} }
func (m *GetUserEncryptKeyReq) String() string { return proto.CompactTextString(m) }
func (*GetUserEncryptKeyReq) ProtoMessage()    {}
func (*GetUserEncryptKeyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{6}
}
func (m *GetUserEncryptKeyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserEncryptKeyReq.Unmarshal(m, b)
}
func (m *GetUserEncryptKeyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserEncryptKeyReq.Marshal(b, m, deterministic)
}
func (dst *GetUserEncryptKeyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserEncryptKeyReq.Merge(dst, src)
}
func (m *GetUserEncryptKeyReq) XXX_Size() int {
	return xxx_messageInfo_GetUserEncryptKeyReq.Size(m)
}
func (m *GetUserEncryptKeyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserEncryptKeyReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserEncryptKeyReq proto.InternalMessageInfo

func (m *GetUserEncryptKeyReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetUserEncryptKeyReq) GetBody() *GetUserEncryptKeyReq_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type GetUserEncryptKeyReq_Body struct {
	Openid               string   `protobuf:"bytes,1,opt,name=openid" json:"openid,omitempty"`
	Signature            string   `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	SigMethod            string   `protobuf:"bytes,3,opt,name=sig_method,json=sigMethod" json:"sig_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserEncryptKeyReq_Body) Reset()         { *m = GetUserEncryptKeyReq_Body{} }
func (m *GetUserEncryptKeyReq_Body) String() string { return proto.CompactTextString(m) }
func (*GetUserEncryptKeyReq_Body) ProtoMessage()    {}
func (*GetUserEncryptKeyReq_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{6, 0}
}
func (m *GetUserEncryptKeyReq_Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserEncryptKeyReq_Body.Unmarshal(m, b)
}
func (m *GetUserEncryptKeyReq_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserEncryptKeyReq_Body.Marshal(b, m, deterministic)
}
func (dst *GetUserEncryptKeyReq_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserEncryptKeyReq_Body.Merge(dst, src)
}
func (m *GetUserEncryptKeyReq_Body) XXX_Size() int {
	return xxx_messageInfo_GetUserEncryptKeyReq_Body.Size(m)
}
func (m *GetUserEncryptKeyReq_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserEncryptKeyReq_Body.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserEncryptKeyReq_Body proto.InternalMessageInfo

func (m *GetUserEncryptKeyReq_Body) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *GetUserEncryptKeyReq_Body) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *GetUserEncryptKeyReq_Body) GetSigMethod() string {
	if m != nil {
		return m.SigMethod
	}
	return ""
}

type GetUserEncryptKeyRes struct {
	Errcode              int32                           `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg               string                          `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	KeyInfoList          []*GetUserEncryptKeyRes_KeyInfo `protobuf:"bytes,3,rep,name=key_info_list,json=keyInfoList" json:"key_info_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetUserEncryptKeyRes) Reset()         { *m = GetUserEncryptKeyRes{} }
func (m *GetUserEncryptKeyRes) String() string { return proto.CompactTextString(m) }
func (*GetUserEncryptKeyRes) ProtoMessage()    {}
func (*GetUserEncryptKeyRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{7}
}
func (m *GetUserEncryptKeyRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserEncryptKeyRes.Unmarshal(m, b)
}
func (m *GetUserEncryptKeyRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserEncryptKeyRes.Marshal(b, m, deterministic)
}
func (dst *GetUserEncryptKeyRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserEncryptKeyRes.Merge(dst, src)
}
func (m *GetUserEncryptKeyRes) XXX_Size() int {
	return xxx_messageInfo_GetUserEncryptKeyRes.Size(m)
}
func (m *GetUserEncryptKeyRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserEncryptKeyRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserEncryptKeyRes proto.InternalMessageInfo

func (m *GetUserEncryptKeyRes) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *GetUserEncryptKeyRes) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *GetUserEncryptKeyRes) GetKeyInfoList() []*GetUserEncryptKeyRes_KeyInfo {
	if m != nil {
		return m.KeyInfoList
	}
	return nil
}

type GetUserEncryptKeyRes_KeyInfo struct {
	EncryptKey           string   `protobuf:"bytes,1,opt,name=encrypt_key,json=encryptKey" json:"encrypt_key,omitempty"`
	Version              int64    `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	ExpireIn             int64    `protobuf:"varint,3,opt,name=expire_in,json=expireIn" json:"expire_in,omitempty"`
	Iv                   string   `protobuf:"bytes,4,opt,name=iv" json:"iv,omitempty"`
	CreateTime           int64    `protobuf:"varint,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserEncryptKeyRes_KeyInfo) Reset()         { *m = GetUserEncryptKeyRes_KeyInfo{} }
func (m *GetUserEncryptKeyRes_KeyInfo) String() string { return proto.CompactTextString(m) }
func (*GetUserEncryptKeyRes_KeyInfo) ProtoMessage()    {}
func (*GetUserEncryptKeyRes_KeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{7, 0}
}
func (m *GetUserEncryptKeyRes_KeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserEncryptKeyRes_KeyInfo.Unmarshal(m, b)
}
func (m *GetUserEncryptKeyRes_KeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserEncryptKeyRes_KeyInfo.Marshal(b, m, deterministic)
}
func (dst *GetUserEncryptKeyRes_KeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserEncryptKeyRes_KeyInfo.Merge(dst, src)
}
func (m *GetUserEncryptKeyRes_KeyInfo) XXX_Size() int {
	return xxx_messageInfo_GetUserEncryptKeyRes_KeyInfo.Size(m)
}
func (m *GetUserEncryptKeyRes_KeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserEncryptKeyRes_KeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserEncryptKeyRes_KeyInfo proto.InternalMessageInfo

func (m *GetUserEncryptKeyRes_KeyInfo) GetEncryptKey() string {
	if m != nil {
		return m.EncryptKey
	}
	return ""
}

func (m *GetUserEncryptKeyRes_KeyInfo) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetUserEncryptKeyRes_KeyInfo) GetExpireIn() int64 {
	if m != nil {
		return m.ExpireIn
	}
	return 0
}

func (m *GetUserEncryptKeyRes_KeyInfo) GetIv() string {
	if m != nil {
		return m.Iv
	}
	return ""
}

func (m *GetUserEncryptKeyRes_KeyInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type GetPhoneNumberReq struct {
	AccessToken          string                  `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Body                 *GetPhoneNumberReq_Body `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetPhoneNumberReq) Reset()         { *m = GetPhoneNumberReq{} }
func (m *GetPhoneNumberReq) String() string { return proto.CompactTextString(m) }
func (*GetPhoneNumberReq) ProtoMessage()    {}
func (*GetPhoneNumberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{8}
}
func (m *GetPhoneNumberReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhoneNumberReq.Unmarshal(m, b)
}
func (m *GetPhoneNumberReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhoneNumberReq.Marshal(b, m, deterministic)
}
func (dst *GetPhoneNumberReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhoneNumberReq.Merge(dst, src)
}
func (m *GetPhoneNumberReq) XXX_Size() int {
	return xxx_messageInfo_GetPhoneNumberReq.Size(m)
}
func (m *GetPhoneNumberReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhoneNumberReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhoneNumberReq proto.InternalMessageInfo

func (m *GetPhoneNumberReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetPhoneNumberReq) GetBody() *GetPhoneNumberReq_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type GetPhoneNumberReq_Body struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPhoneNumberReq_Body) Reset()         { *m = GetPhoneNumberReq_Body{} }
func (m *GetPhoneNumberReq_Body) String() string { return proto.CompactTextString(m) }
func (*GetPhoneNumberReq_Body) ProtoMessage()    {}
func (*GetPhoneNumberReq_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{8, 0}
}
func (m *GetPhoneNumberReq_Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhoneNumberReq_Body.Unmarshal(m, b)
}
func (m *GetPhoneNumberReq_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhoneNumberReq_Body.Marshal(b, m, deterministic)
}
func (dst *GetPhoneNumberReq_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhoneNumberReq_Body.Merge(dst, src)
}
func (m *GetPhoneNumberReq_Body) XXX_Size() int {
	return xxx_messageInfo_GetPhoneNumberReq_Body.Size(m)
}
func (m *GetPhoneNumberReq_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhoneNumberReq_Body.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhoneNumberReq_Body proto.InternalMessageInfo

func (m *GetPhoneNumberReq_Body) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type GetPhoneNumberRes struct {
	Errcode              int32                        `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg               string                       `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	PhoneInfo            *GetPhoneNumberRes_PhoneInfo `protobuf:"bytes,3,opt,name=phone_info,json=phoneInfo" json:"phone_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetPhoneNumberRes) Reset()         { *m = GetPhoneNumberRes{} }
func (m *GetPhoneNumberRes) String() string { return proto.CompactTextString(m) }
func (*GetPhoneNumberRes) ProtoMessage()    {}
func (*GetPhoneNumberRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{9}
}
func (m *GetPhoneNumberRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhoneNumberRes.Unmarshal(m, b)
}
func (m *GetPhoneNumberRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhoneNumberRes.Marshal(b, m, deterministic)
}
func (dst *GetPhoneNumberRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhoneNumberRes.Merge(dst, src)
}
func (m *GetPhoneNumberRes) XXX_Size() int {
	return xxx_messageInfo_GetPhoneNumberRes.Size(m)
}
func (m *GetPhoneNumberRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhoneNumberRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhoneNumberRes proto.InternalMessageInfo

func (m *GetPhoneNumberRes) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *GetPhoneNumberRes) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *GetPhoneNumberRes) GetPhoneInfo() *GetPhoneNumberRes_PhoneInfo {
	if m != nil {
		return m.PhoneInfo
	}
	return nil
}

type GetPhoneNumberRes_PhoneInfo struct {
	PhoneNumber          string                       `protobuf:"bytes,1,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	PurePhoneNumber      string                       `protobuf:"bytes,2,opt,name=purePhoneNumber" json:"purePhoneNumber,omitempty"`
	CountryCode          string                       `protobuf:"bytes,3,opt,name=countryCode" json:"countryCode,omitempty"`
	Watermark            *GetPhoneNumberRes_Watermark `protobuf:"bytes,4,opt,name=watermark" json:"watermark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetPhoneNumberRes_PhoneInfo) Reset()         { *m = GetPhoneNumberRes_PhoneInfo{} }
func (m *GetPhoneNumberRes_PhoneInfo) String() string { return proto.CompactTextString(m) }
func (*GetPhoneNumberRes_PhoneInfo) ProtoMessage()    {}
func (*GetPhoneNumberRes_PhoneInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{9, 0}
}
func (m *GetPhoneNumberRes_PhoneInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhoneNumberRes_PhoneInfo.Unmarshal(m, b)
}
func (m *GetPhoneNumberRes_PhoneInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhoneNumberRes_PhoneInfo.Marshal(b, m, deterministic)
}
func (dst *GetPhoneNumberRes_PhoneInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhoneNumberRes_PhoneInfo.Merge(dst, src)
}
func (m *GetPhoneNumberRes_PhoneInfo) XXX_Size() int {
	return xxx_messageInfo_GetPhoneNumberRes_PhoneInfo.Size(m)
}
func (m *GetPhoneNumberRes_PhoneInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhoneNumberRes_PhoneInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhoneNumberRes_PhoneInfo proto.InternalMessageInfo

func (m *GetPhoneNumberRes_PhoneInfo) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *GetPhoneNumberRes_PhoneInfo) GetPurePhoneNumber() string {
	if m != nil {
		return m.PurePhoneNumber
	}
	return ""
}

func (m *GetPhoneNumberRes_PhoneInfo) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *GetPhoneNumberRes_PhoneInfo) GetWatermark() *GetPhoneNumberRes_Watermark {
	if m != nil {
		return m.Watermark
	}
	return nil
}

type GetPhoneNumberRes_Watermark struct {
	Appid                string   `protobuf:"bytes,1,opt,name=appid" json:"appid,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPhoneNumberRes_Watermark) Reset()         { *m = GetPhoneNumberRes_Watermark{} }
func (m *GetPhoneNumberRes_Watermark) String() string { return proto.CompactTextString(m) }
func (*GetPhoneNumberRes_Watermark) ProtoMessage()    {}
func (*GetPhoneNumberRes_Watermark) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3410a6c884bb1b2f, []int{9, 1}
}
func (m *GetPhoneNumberRes_Watermark) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhoneNumberRes_Watermark.Unmarshal(m, b)
}
func (m *GetPhoneNumberRes_Watermark) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhoneNumberRes_Watermark.Marshal(b, m, deterministic)
}
func (dst *GetPhoneNumberRes_Watermark) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhoneNumberRes_Watermark.Merge(dst, src)
}
func (m *GetPhoneNumberRes_Watermark) XXX_Size() int {
	return xxx_messageInfo_GetPhoneNumberRes_Watermark.Size(m)
}
func (m *GetPhoneNumberRes_Watermark) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhoneNumberRes_Watermark.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhoneNumberRes_Watermark proto.InternalMessageInfo

func (m *GetPhoneNumberRes_Watermark) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *GetPhoneNumberRes_Watermark) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPluginOpenPIdReq)(nil), "api.weixin.qq.com.GetPluginOpenPIdReq")
	proto.RegisterType((*GetPluginOpenPIdReq_Body)(nil), "api.weixin.qq.com.GetPluginOpenPIdReq.Body")
	proto.RegisterType((*GetPluginOpenPidRes)(nil), "api.weixin.qq.com.GetPluginOpenPidRes")
	proto.RegisterType((*CheckEncryptedDataReq)(nil), "api.weixin.qq.com.CheckEncryptedDataReq")
	proto.RegisterType((*CheckEncryptedDataReq_Body)(nil), "api.weixin.qq.com.CheckEncryptedDataReq.Body")
	proto.RegisterType((*CheckEncryptedDataRes)(nil), "api.weixin.qq.com.CheckEncryptedDataRes")
	proto.RegisterType((*GetPaidUnionidReq)(nil), "api.weixin.qq.com.GetPaidUnionidReq")
	proto.RegisterType((*GetPaidUnionidRes)(nil), "api.weixin.qq.com.GetPaidUnionidRes")
	proto.RegisterType((*GetUserEncryptKeyReq)(nil), "api.weixin.qq.com.GetUserEncryptKeyReq")
	proto.RegisterType((*GetUserEncryptKeyReq_Body)(nil), "api.weixin.qq.com.GetUserEncryptKeyReq.Body")
	proto.RegisterType((*GetUserEncryptKeyRes)(nil), "api.weixin.qq.com.GetUserEncryptKeyRes")
	proto.RegisterType((*GetUserEncryptKeyRes_KeyInfo)(nil), "api.weixin.qq.com.GetUserEncryptKeyRes.KeyInfo")
	proto.RegisterType((*GetPhoneNumberReq)(nil), "api.weixin.qq.com.GetPhoneNumberReq")
	proto.RegisterType((*GetPhoneNumberReq_Body)(nil), "api.weixin.qq.com.GetPhoneNumberReq.Body")
	proto.RegisterType((*GetPhoneNumberRes)(nil), "api.weixin.qq.com.GetPhoneNumberRes")
	proto.RegisterType((*GetPhoneNumberRes_PhoneInfo)(nil), "api.weixin.qq.com.GetPhoneNumberRes.PhoneInfo")
	proto.RegisterType((*GetPhoneNumberRes_Watermark)(nil), "api.weixin.qq.com.GetPhoneNumberRes.Watermark")
}

func init() { proto.RegisterFile("wx-miniprogram/user.proto", fileDescriptor_user_3410a6c884bb1b2f) }

var fileDescriptor_user_3410a6c884bb1b2f = []byte{
	// 1007 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x8f, 0xdb, 0x44,
	0x14, 0x97, 0x37, 0x9b, 0xb4, 0x99, 0xb4, 0x4b, 0x3b, 0xb4, 0x55, 0x30, 0x45, 0x04, 0xab, 0xa5,
	0x41, 0x74, 0x1d, 0x75, 0x41, 0x42, 0xfc, 0xa9, 0x16, 0xba, 0x45, 0x25, 0xb4, 0x5b, 0x2a, 0x77,
	0x2b, 0x24, 0x38, 0x58, 0x13, 0xfb, 0xd5, 0x19, 0x65, 0x3d, 0xe3, 0x9d, 0x19, 0x67, 0x37, 0x42,
	0x48, 0x88, 0x6f, 0x00, 0xa8, 0x67, 0x0e, 0x5c, 0x38, 0x21, 0x44, 0x3f, 0x04, 0x1f, 0x80, 0x0b,
	0x37, 0x2e, 0x7c, 0x10, 0x34, 0xe3, 0x71, 0xfe, 0x6c, 0xb2, 0xd4, 0xda, 0x9b, 0xdf, 0x9b, 0x37,
	0xcf, 0xbf, 0xf7, 0x9b, 0xdf, 0xbc, 0x37, 0xe8, 0x95, 0xc3, 0xa3, 0xcd, 0x94, 0x32, 0x9a, 0x09,
	0x9e, 0x08, 0x92, 0xf6, 0x72, 0x09, 0xc2, 0xcf, 0x04, 0x57, 0x1c, 0x5f, 0x24, 0x19, 0xf5, 0x0f,
	0x81, 0x1e, 0x51, 0xe6, 0x1f, 0x1c, 0xf8, 0x11, 0x4f, 0xdd, 0xab, 0x09, 0xe7, 0xc9, 0x3e, 0xf4,
	0x48, 0x46, 0x7b, 0x84, 0x31, 0xae, 0x88, 0xa2, 0x9c, 0xc9, 0x62, 0x83, 0xf7, 0xcc, 0x41, 0x2f,
	0xdf, 0x03, 0xf5, 0x68, 0x3f, 0x4f, 0x28, 0xfb, 0x22, 0x03, 0xf6, 0xa8, 0x1f, 0x07, 0x70, 0x80,
	0xdf, 0x40, 0xe7, 0x48, 0x14, 0x81, 0x94, 0xa1, 0xe2, 0x23, 0x60, 0x6d, 0xa7, 0xe3, 0x74, 0x9b,
	0x41, 0xab, 0xf0, 0xed, 0x69, 0x17, 0xde, 0x46, 0xeb, 0x03, 0x1e, 0x4f, 0xda, 0x6b, 0x1d, 0xa7,
	0xdb, 0xda, 0x7a, 0xdb, 0x5f, 0xfa, 0xb5, 0xbf, 0x22, 0xb1, 0x7f, 0x87, 0xc7, 0x93, 0xc0, 0x6c,
	0x74, 0x5d, 0xb4, 0xae, 0x2d, 0x8c, 0xd1, 0x7a, 0xc4, 0x63, 0xb0, 0xff, 0x30, 0xdf, 0x1e, 0x39,
	0x0e, 0x8b, 0xc6, 0x01, 0x48, 0xdc, 0x46, 0x67, 0x40, 0x88, 0x69, 0x74, 0x3d, 0x28, 0x4d, 0x7c,
	0x05, 0x35, 0x40, 0x88, 0x54, 0x26, 0x06, 0x4f, 0x33, 0xb0, 0x96, 0xde, 0xc1, 0x33, 0x60, 0x19,
	0x8d, 0xdb, 0x35, 0xb3, 0x50, 0x9a, 0xde, 0x73, 0x07, 0x5d, 0xde, 0x19, 0x42, 0x34, 0xfa, 0x94,
	0x45, 0x62, 0x92, 0x29, 0x88, 0xef, 0x12, 0x45, 0x2a, 0x16, 0xff, 0xc9, 0x42, 0xf1, 0x9b, 0x2b,
	0x8a, 0x5f, 0x99, 0x7a, 0xbe, 0xfc, 0x77, 0x6d, 0xf9, 0x37, 0x11, 0x86, 0x32, 0x2c, 0x4c, 0x65,
	0x12, 0x0e, 0x89, 0x1c, 0xda, 0x7f, 0x5e, 0x98, 0xae, 0xec, 0xca, 0xe4, 0x33, 0x22, 0x87, 0xde,
	0x77, 0x27, 0xa0, 0x3e, 0x0d, 0x37, 0x97, 0x50, 0x7d, 0x4c, 0xf6, 0x2d, 0x33, 0x67, 0x83, 0xc2,
	0xc0, 0xaf, 0xa3, 0x56, 0x24, 0x80, 0x28, 0x08, 0x15, 0x4d, 0xa1, 0xbd, 0xde, 0x71, 0xba, 0xb5,
	0x00, 0x15, 0xae, 0x3d, 0x9a, 0x82, 0xf7, 0x9b, 0x83, 0x2e, 0xea, 0xc3, 0x21, 0x34, 0x7e, 0xc2,
	0x28, 0x67, 0xb4, 0xaa, 0x62, 0xae, 0xa0, 0x86, 0x26, 0x9f, 0xc6, 0x25, 0x8e, 0xc2, 0xc2, 0xd7,
	0xd1, 0x86, 0x12, 0x84, 0x49, 0x12, 0x69, 0x69, 0x86, 0xd3, 0xa3, 0x3a, 0x3f, 0xe7, 0xed, 0xc7,
	0xf8, 0x32, 0x6a, 0xa4, 0xd1, 0x50, 0x2f, 0xaf, 0x9b, 0xe5, 0x7a, 0x1a, 0x0d, 0xfb, 0x31, 0xee,
	0xa0, 0x73, 0x3c, 0x57, 0xa1, 0x12, 0x24, 0x86, 0x90, 0xf1, 0x76, 0xdd, 0x2c, 0x22, 0x9e, 0xab,
	0x3d, 0xed, 0x7a, 0xc8, 0xbd, 0x70, 0x19, 0xef, 0x29, 0xa5, 0x94, 0x17, 0xfb, 0x4b, 0x29, 0x59,
	0xd3, 0xfb, 0xdb, 0x41, 0x97, 0xee, 0x81, 0x7a, 0x22, 0x41, 0xd8, 0x63, 0xb9, 0x0f, 0x93, 0x8a,
	0xa4, 0x7c, 0xbc, 0xa0, 0xa4, 0x9b, 0xab, 0xaf, 0xd1, 0x52, 0xe6, 0x79, 0x21, 0x7d, 0x6d, 0x85,
	0x34, 0xa3, 0xd7, 0x59, 0xa0, 0xf7, 0x2a, 0x6a, 0x4a, 0x9a, 0x30, 0xa2, 0x72, 0x01, 0xb6, 0xa4,
	0x99, 0x03, 0xbf, 0x86, 0x90, 0xa4, 0x49, 0x98, 0x82, 0x1a, 0xf2, 0xb2, 0x30, 0xbd, 0xbc, 0x6b,
	0x1c, 0xde, 0x1f, 0x6b, 0x2b, 0x4b, 0x3b, 0x0d, 0x7f, 0x8f, 0xd1, 0xf9, 0x11, 0x4c, 0x42, 0xca,
	0x9e, 0xf2, 0x70, 0x9f, 0x4a, 0xd5, 0xae, 0x75, 0x6a, 0xdd, 0xd6, 0x56, 0xaf, 0x5a, 0xc9, 0xd2,
	0xbf, 0x0f, 0x93, 0x3e, 0x7b, 0xca, 0x83, 0xd6, 0xa8, 0xf8, 0x78, 0x40, 0xa5, 0x72, 0x7f, 0x74,
	0xd0, 0x19, 0xbb, 0xa0, 0x95, 0x6b, 0xef, 0x4b, 0x38, 0x82, 0x89, 0x65, 0x01, 0xc1, 0x34, 0x89,
	0xc6, 0x3c, 0x06, 0x21, 0x29, 0x67, 0x06, 0x5a, 0x2d, 0x28, 0x4d, 0xfc, 0x2a, 0x6a, 0xc2, 0x51,
	0x46, 0x05, 0x84, 0x94, 0x19, 0x12, 0x6a, 0xc1, 0xd9, 0xc2, 0xd1, 0x67, 0x78, 0x03, 0xad, 0xd1,
	0xb1, 0x15, 0xdd, 0x1a, 0x1d, 0x1f, 0xbf, 0x21, 0xf5, 0xa5, 0x1b, 0xf2, 0x83, 0xbd, 0x21, 0x43,
	0xce, 0xe0, 0x61, 0x9e, 0x0e, 0x40, 0x54, 0x14, 0xc3, 0xed, 0x05, 0x31, 0xbc, 0x75, 0x42, 0x4f,
	0x5d, 0x48, 0x5b, 0xb5, 0xa3, 0xfe, 0x5c, 0x5b, 0xc6, 0x74, 0x9a, 0x53, 0xdc, 0x45, 0x28, 0xd3,
	0x39, 0xcc, 0x39, 0x1a, 0xaa, 0x5a, 0x5b, 0x7e, 0x05, 0xa0, 0xd2, 0x37, 0xa6, 0x39, 0xc1, 0x66,
	0x56, 0x7e, 0xba, 0x7f, 0x3a, 0xa8, 0x39, 0x5d, 0xc0, 0x1d, 0xd4, 0xca, 0x66, 0x9b, 0x4a, 0x86,
	0xe6, 0x5c, 0xb8, 0x8b, 0x5e, 0xca, 0x72, 0x01, 0x73, 0xa9, 0x2d, 0xbe, 0xe3, 0x6e, 0x9d, 0x2b,
	0xe2, 0x39, 0x53, 0x62, 0xb2, 0xa3, 0xcb, 0x2b, 0x94, 0x3d, 0xef, 0xc2, 0x0f, 0x50, 0xf3, 0x90,
	0x28, 0x10, 0x29, 0x11, 0x23, 0x73, 0xbc, 0x55, 0x2b, 0xf9, 0xb2, 0xdc, 0x15, 0xcc, 0x12, 0xb8,
	0xdb, 0xa8, 0x39, 0xf5, 0xeb, 0xd6, 0x4a, 0xb2, 0x6c, 0x7a, 0x15, 0x0b, 0x43, 0xdf, 0x44, 0xad,
	0x18, 0xa9, 0x48, 0x9a, 0x59, 0x05, 0xce, 0x1c, 0x5b, 0xff, 0x34, 0x50, 0x4b, 0xab, 0xfe, 0x31,
	0x88, 0x31, 0x8d, 0x00, 0xff, 0xe2, 0xa0, 0x0b, 0xc7, 0x47, 0x28, 0x7e, 0xb3, 0xda, 0x9c, 0x75,
	0x5f, 0x1c, 0x67, 0xda, 0xa0, 0xb7, 0xfd, 0xfd, 0x5f, 0xff, 0xfe, 0xb4, 0xf6, 0xbe, 0x77, 0xab,
	0x77, 0x78, 0x44, 0x7a, 0x09, 0xa8, 0xcc, 0x44, 0xd8, 0x21, 0xb9, 0x3d, 0x2f, 0xd8, 0xdb, 0xdf,
	0xcc, 0x5b, 0xdf, 0x7e, 0x60, 0x34, 0x87, 0x9f, 0x3b, 0x08, 0x2f, 0x0f, 0x24, 0xdc, 0xad, 0x3a,
	0x12, 0xdd, 0xaa, 0x91, 0xd2, 0xeb, 0x1b, 0xac, 0x3b, 0xde, 0x87, 0x06, 0xeb, 0x20, 0x97, 0x94,
	0x81, 0x94, 0xbd, 0x48, 0x07, 0x4f, 0x07, 0x65, 0x2a, 0x93, 0x0a, 0xa8, 0x9f, 0x39, 0x68, 0x63,
	0x71, 0x26, 0xe0, 0x6b, 0x27, 0x30, 0xb6, 0x30, 0xe6, 0xdc, 0x2a, 0x51, 0xd2, 0x7b, 0xcf, 0x20,
	0xbd, 0x85, 0x7b, 0x53, 0x56, 0x09, 0x8d, 0xed, 0xb4, 0xf8, 0x3f, 0x74, 0xf8, 0xf7, 0xa2, 0x73,
	0x2c, 0x36, 0x3f, 0x7c, 0xa3, 0xe2, 0x54, 0x70, 0x2b, 0x06, 0x9e, 0x48, 0x65, 0x02, 0x4a, 0xbf,
	0x26, 0x2d, 0x99, 0x23, 0x98, 0x54, 0xa0, 0xf2, 0x57, 0x4b, 0xe5, 0xdc, 0xd5, 0xbb, 0x56, 0xa5,
	0x71, 0xb9, 0x55, 0xa2, 0xa4, 0xf7, 0xb9, 0x41, 0x7a, 0xd7, 0xfb, 0x68, 0x25, 0x52, 0xd3, 0x1a,
	0x98, 0x09, 0x7e, 0x31, 0xd4, 0x3b, 0x37, 0xbe, 0xba, 0x9e, 0x50, 0x35, 0xcc, 0x07, 0xfa, 0x5f,
	0xbd, 0x18, 0xc6, 0x9b, 0x5a, 0xe7, 0xfa, 0x65, 0xbc, 0xf8, 0xa4, 0x1e, 0x34, 0xcc, 0xeb, 0xf8,
	0x9d, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0x1c, 0x72, 0xb0, 0x6b, 0x0b, 0x00, 0x00,
}
