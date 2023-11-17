// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.14.0
// source: api/v1/common.proto

package v1

import (
	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The canonical error codes for gRPC APIs.
// https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// Sometimes multiple error codes may apply.  Services should return
// the most specific error code that applies.  For example, prefer
// `OUT_OF_RANGE` over `FAILED_PRECONDITION` if both codes apply.
// Similarly prefer `NOT_FOUND` or `ALREADY_EXISTS` over `FAILED_PRECONDITION`.
type Code int32

const (
	// Not an error; returned on success
	//
	// HTTP Mapping: 200 OK
	Code_OK Code = 0
	// The operation was cancelled, typically by the caller.
	//
	// HTTP Mapping: 499 Client Closed Request
	Code_CANCELLED Code = 1
	// Unknown error.  For example, this error may be returned when
	// a `Status` value received from another address space belongs to
	// an error space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_UNKNOWN Code = 2
	// The client specified an invalid argument.  Note that this differs
	// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	//
	// HTTP Mapping: 400 Bad Request
	Code_INVALID_ARGUMENT Code = 3
	// The deadline expired before the operation could complete. For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example, a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	//
	// HTTP Mapping: 504 Gateway Timeout
	Code_DEADLINE_EXCEEDED Code = 4
	// Some requested entity (e.g., file or directory) was not found.
	//
	// Note to server developers: if a request is denied for an entire class
	// of users, such as gradual feature rollout or undocumented whitelist,
	// `NOT_FOUND` may be used. If a request is denied for some users within
	// a class of users, such as user-based access control, `PERMISSION_DENIED`
	// must be used.
	//
	// HTTP Mapping: 404 Not Found
	Code_NOT_FOUND Code = 5
	// The entity that a client attempted to create (e.g., file or directory)
	// already exists.
	//
	// HTTP Mapping: 409 Conflict
	Code_ALREADY_EXISTS Code = 6
	// The caller does not have permission to execute the specified
	// operation. `PERMISSION_DENIED` must not be used for rejections
	// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
	// instead for those errors). `PERMISSION_DENIED` must not be
	// used if the caller can not be identified (use `UNAUTHENTICATED`
	// instead for those errors). This error code does not imply the
	// request is valid or the requested entity exists or satisfies
	// other pre-conditions.
	//
	// HTTP Mapping: 403 Forbidden
	Code_PERMISSION_DENIED Code = 7
	// The request does not have valid authentication credentials for the
	// operation.
	//
	// HTTP Mapping: 401 Unauthorized
	Code_UNAUTHENTICATED Code = 16
	// Some resource has been exhausted, perhaps a per-user quota, or
	// perhaps the entire file system is out of space.
	//
	// HTTP Mapping: 429 Too Many Requests
	Code_RESOURCE_EXHAUSTED Code = 8
	// The operation was rejected because the system is not in a state
	// required for the operation's execution.  For example, the directory
	// to be deleted is non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// Service implementors can use the following guidelines to decide
	// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
	//  (a) Use `UNAVAILABLE` if the client can retry just the failing call.
	//  (b) Use `ABORTED` if the client should retry at a higher level
	//      (e.g., when a client-specified test-and-set fails, indicating the
	//      client should restart a read-modify-write sequence).
	//  (c) Use `FAILED_PRECONDITION` if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, `FAILED_PRECONDITION`
	//      should be returned since the client should not retry unless
	//      the files are deleted from the directory.
	//
	// HTTP Mapping: 400 Bad Request
	Code_FAILED_PRECONDITION Code = 9
	// The operation was aborted, typically due to a concurrency issue such as
	// a sequencer check failure or transaction abort.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 409 Conflict
	Code_ABORTED Code = 10
	// The operation was attempted past the valid range.  E.g., seeking or
	// reading past end-of-file.
	//
	// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate `INVALID_ARGUMENT` if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// `OUT_OF_RANGE` if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between `FAILED_PRECONDITION` and
	// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an `OUT_OF_RANGE` error to detect when
	// they are done.
	//
	// HTTP Mapping: 400 Bad Request
	Code_OUT_OF_RANGE Code = 11
	// The operation is not implemented or is not supported/enabled in this
	// service.
	//
	// HTTP Mapping: 501 Not Implemented
	Code_UNIMPLEMENTED Code = 12
	// Internal errors.  This means that some invariants expected by the
	// underlying system have been broken.  This error code is reserved
	// for serious errors.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_INTERNAL Code = 13
	// The service is currently unavailable.  This is most likely a
	// transient condition, which can be corrected by retrying with
	// a backoff. Note that it is not always safe to retry
	// non-idempotent operations.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 503 Service Unavailable
	Code_UNAVAILABLE Code = 14
	// Unrecoverable data loss or corruption.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_DATA_LOSS Code = 15
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:  "OK",
		1:  "CANCELLED",
		2:  "UNKNOWN",
		3:  "INVALID_ARGUMENT",
		4:  "DEADLINE_EXCEEDED",
		5:  "NOT_FOUND",
		6:  "ALREADY_EXISTS",
		7:  "PERMISSION_DENIED",
		16: "UNAUTHENTICATED",
		8:  "RESOURCE_EXHAUSTED",
		9:  "FAILED_PRECONDITION",
		10: "ABORTED",
		11: "OUT_OF_RANGE",
		12: "UNIMPLEMENTED",
		13: "INTERNAL",
		14: "UNAVAILABLE",
		15: "DATA_LOSS",
	}
	Code_value = map[string]int32{
		"OK":                  0,
		"CANCELLED":           1,
		"UNKNOWN":             2,
		"INVALID_ARGUMENT":    3,
		"DEADLINE_EXCEEDED":   4,
		"NOT_FOUND":           5,
		"ALREADY_EXISTS":      6,
		"PERMISSION_DENIED":   7,
		"UNAUTHENTICATED":     16,
		"RESOURCE_EXHAUSTED":  8,
		"FAILED_PRECONDITION": 9,
		"ABORTED":             10,
		"OUT_OF_RANGE":        11,
		"UNIMPLEMENTED":       12,
		"INTERNAL":            13,
		"UNAVAILABLE":         14,
		"DATA_LOSS":           15,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_common_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_api_v1_common_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{0}
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type ExternalAuthn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host *Host `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *ExternalAuthn) Reset() {
	*x = ExternalAuthn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalAuthn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAuthn) ProtoMessage() {}

func (x *ExternalAuthn) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAuthn.ProtoReflect.Descriptor instead.
func (*ExternalAuthn) Descriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalAuthn) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

type DirectAuthn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     *Host  `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *DirectAuthn) Reset() {
	*x = DirectAuthn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectAuthn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectAuthn) ProtoMessage() {}

func (x *DirectAuthn) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectAuthn.ProtoReflect.Descriptor instead.
func (*DirectAuthn) Descriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *DirectAuthn) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *DirectAuthn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DirectAuthn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Authn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Authn:
	//	*Authn_DirectAuthn
	Authn isAuthn_Authn `protobuf_oneof:"authn"`
}

func (x *Authn) Reset() {
	*x = Authn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authn) ProtoMessage() {}

func (x *Authn) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authn.ProtoReflect.Descriptor instead.
func (*Authn) Descriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{3}
}

func (m *Authn) GetAuthn() isAuthn_Authn {
	if m != nil {
		return m.Authn
	}
	return nil
}

func (x *Authn) GetDirectAuthn() *DirectAuthn {
	if x, ok := x.GetAuthn().(*Authn_DirectAuthn); ok {
		return x.DirectAuthn
	}
	return nil
}

type isAuthn_Authn interface {
	isAuthn_Authn()
}

type Authn_DirectAuthn struct {
	DirectAuthn *DirectAuthn `protobuf:"bytes,1,opt,name=directAuthn,proto3,oneof"` //ExternalAuthn externalAuthn = 2;
}

func (*Authn_DirectAuthn) isAuthn_Authn() {}

type Vendor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Vendor) Reset() {
	*x = Vendor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vendor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vendor) ProtoMessage() {}

func (x *Vendor) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vendor.ProtoReflect.Descriptor instead.
func (*Vendor) Descriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *Vendor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_v1_common_proto protoreflect.FileDescriptor

var file_api_v1_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e,
	0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x40, 0x76, 0x30, 0x2e, 0x33, 0x2e, 0x32, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x0d, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x3b, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65,
	0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x3b, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e,
	0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x64,
	0x0a, 0x05, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x52, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72,
	0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x48, 0x00, 0x52, 0x0b,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x22, 0x1c, 0x0a, 0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x2a, 0xb7, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12,
	0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e,
	0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x41,
	0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x06, 0x12,
	0x15, 0x0a, 0x11, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45,
	0x4e, 0x49, 0x45, 0x44, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48,
	0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45,
	0x44, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x50, 0x52,
	0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x55, 0x54,
	0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x55,
	0x4e, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x0c,
	0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x0e, 0x12, 0x0d, 0x0a,
	0x09, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x0f, 0x42, 0x33, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2f, 0x70, 0x62, 0x6e, 0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0xea, 0x02, 0x0d, 0x50, 0x62, 0x6e, 0x6a, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_common_proto_rawDescOnce sync.Once
	file_api_v1_common_proto_rawDescData = file_api_v1_common_proto_rawDesc
)

func file_api_v1_common_proto_rawDescGZIP() []byte {
	file_api_v1_common_proto_rawDescOnce.Do(func() {
		file_api_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_common_proto_rawDescData)
	})
	return file_api_v1_common_proto_rawDescData
}

var file_api_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_common_proto_goTypes = []interface{}{
	(Code)(0),             // 0: github.com.tinkerbell.pbnj.api.v1.Code
	(*Host)(nil),          // 1: github.com.tinkerbell.pbnj.api.v1.Host
	(*ExternalAuthn)(nil), // 2: github.com.tinkerbell.pbnj.api.v1.ExternalAuthn
	(*DirectAuthn)(nil),   // 3: github.com.tinkerbell.pbnj.api.v1.DirectAuthn
	(*Authn)(nil),         // 4: github.com.tinkerbell.pbnj.api.v1.Authn
	(*Vendor)(nil),        // 5: github.com.tinkerbell.pbnj.api.v1.Vendor
}
var file_api_v1_common_proto_depIdxs = []int32{
	1, // 0: github.com.tinkerbell.pbnj.api.v1.ExternalAuthn.host:type_name -> github.com.tinkerbell.pbnj.api.v1.Host
	1, // 1: github.com.tinkerbell.pbnj.api.v1.DirectAuthn.host:type_name -> github.com.tinkerbell.pbnj.api.v1.Host
	3, // 2: github.com.tinkerbell.pbnj.api.v1.Authn.directAuthn:type_name -> github.com.tinkerbell.pbnj.api.v1.DirectAuthn
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_common_proto_init() }
func file_api_v1_common_proto_init() {
	if File_api_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalAuthn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectAuthn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vendor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_v1_common_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Authn_DirectAuthn)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_common_proto_goTypes,
		DependencyIndexes: file_api_v1_common_proto_depIdxs,
		EnumInfos:         file_api_v1_common_proto_enumTypes,
		MessageInfos:      file_api_v1_common_proto_msgTypes,
	}.Build()
	File_api_v1_common_proto = out.File
	file_api_v1_common_proto_rawDesc = nil
	file_api_v1_common_proto_goTypes = nil
	file_api_v1_common_proto_depIdxs = nil
}
