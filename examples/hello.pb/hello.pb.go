// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/pbgo"

import "encoding/json"
import "net/rpc"
import "net/http"
import "regexp"
import "strings"

import "github.com/chai2010/pbgo"
import "github.com/julienschmidt/httprouter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_d80cc3bf83ccd227, []int{0}
}
func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (dst *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(dst, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StaticFile struct {
	ContentType          string   `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	ContentBody          []byte   `protobuf:"bytes,2,opt,name=content_body,json=contentBody,proto3" json:"content_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticFile) Reset()         { *m = StaticFile{} }
func (m *StaticFile) String() string { return proto.CompactTextString(m) }
func (*StaticFile) ProtoMessage()    {}
func (*StaticFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_d80cc3bf83ccd227, []int{1}
}
func (m *StaticFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticFile.Unmarshal(m, b)
}
func (m *StaticFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticFile.Marshal(b, m, deterministic)
}
func (dst *StaticFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticFile.Merge(dst, src)
}
func (m *StaticFile) XXX_Size() int {
	return xxx_messageInfo_StaticFile.Size(m)
}
func (m *StaticFile) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticFile.DiscardUnknown(m)
}

var xxx_messageInfo_StaticFile proto.InternalMessageInfo

func (m *StaticFile) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *StaticFile) GetContentBody() []byte {
	if m != nil {
		return m.ContentBody
	}
	return nil
}

type Message struct {
	Value                string            `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Array                []int32           `protobuf:"varint,2,rep,packed,name=array,proto3" json:"array,omitempty"`
	Dict                 map[string]string `protobuf:"bytes,3,rep,name=dict,proto3" json:"dict,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Subfiled             *String           `protobuf:"bytes,4,opt,name=subfiled,proto3" json:"subfiled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_d80cc3bf83ccd227, []int{2}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Message) GetArray() []int32 {
	if m != nil {
		return m.Array
	}
	return nil
}

func (m *Message) GetDict() map[string]string {
	if m != nil {
		return m.Dict
	}
	return nil
}

func (m *Message) GetSubfiled() *String {
	if m != nil {
		return m.Subfiled
	}
	return nil
}

func init() {
	proto.RegisterType((*String)(nil), "hello_pb.String")
	proto.RegisterType((*StaticFile)(nil), "hello_pb.StaticFile")
	proto.RegisterType((*Message)(nil), "hello_pb.Message")
	proto.RegisterMapType((map[string]string)(nil), "hello_pb.Message.DictEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ = json.Marshal
var _ = http.ListenAndServe
var _ = regexp.Match
var _ = strings.Split
var _ = pbgo.PopulateFieldFromPath
var _ = httprouter.New

type HelloServiceInterface interface {
	Hello(in *String, out *String) error
	Echo(in *Message, out *Message) error
	Static(in *String, out *StaticFile) error
}

func RegisterHelloService(srv *rpc.Server, x HelloServiceInterface) error {
	if err := srv.RegisterName("HelloService", x); err != nil {
		return err
	}
	return nil
}

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(in *String) (*String, error) {
	var out = new(String)
	if err := p.Client.Call("HelloService.Hello", in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (p *HelloServiceClient) Echo(in *Message) (*Message, error) {
	var out = new(Message)
	if err := p.Client.Call("HelloService.Echo", in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (p *HelloServiceClient) Static(in *String) (*StaticFile, error) {
	var out = new(StaticFile)
	if err := p.Client.Call("HelloService.Static", in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func HelloServiceHandler(svc HelloServiceInterface) http.Handler {
	var router = httprouter.New()

	var re = regexp.MustCompile("(\\*|\\:)(\\w|\\.)+")
	_ = re

	router.Handle("DELETE", "/hello",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   String
				protoReply String
			)

			for _, fieldPath := range re.FindAllString("/hello", -1) {
				fieldPath := strings.TrimLeft(fieldPath, ":*")
				err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.Hello(&protoReq, &protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if strings.Contains(r.Header.Get("Accept"), "application/json") {
				w.Header().Set("Content-Type", "application/json")
			} else {
				w.Header().Set("Content-Type", "text/plain")
			}

			if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)

	router.Handle("GET", "/hello/:value",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   String
				protoReply String
			)

			for _, fieldPath := range re.FindAllString("/hello/:value", -1) {
				fieldPath := strings.TrimLeft(fieldPath, ":*")
				err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.Hello(&protoReq, &protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if strings.Contains(r.Header.Get("Accept"), "application/json") {
				w.Header().Set("Content-Type", "application/json")
			} else {
				w.Header().Set("Content-Type", "text/plain")
			}

			if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)

	router.Handle("PATCH", "/hello",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   String
				protoReply String
			)

			for _, fieldPath := range re.FindAllString("/hello", -1) {
				fieldPath := strings.TrimLeft(fieldPath, ":*")
				err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := json.NewDecoder(r.Body).Decode(&protoReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.Hello(&protoReq, &protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if strings.Contains(r.Header.Get("Accept"), "application/json") {
				w.Header().Set("Content-Type", "application/json")
			} else {
				w.Header().Set("Content-Type", "text/plain")
			}

			if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)

	router.Handle("POST", "/hello",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   String
				protoReply String
			)

			for _, fieldPath := range re.FindAllString("/hello", -1) {
				fieldPath := strings.TrimLeft(fieldPath, ":*")
				err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := json.NewDecoder(r.Body).Decode(&protoReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.Hello(&protoReq, &protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if strings.Contains(r.Header.Get("Accept"), "application/json") {
				w.Header().Set("Content-Type", "application/json")
			} else {
				w.Header().Set("Content-Type", "text/plain")
			}

			if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)

	router.Handle("GET", "/echo/:subfiled.value",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   Message
				protoReply Message
			)

			for _, fieldPath := range re.FindAllString("/echo/:subfiled.value", -1) {
				fieldPath := strings.TrimLeft(fieldPath, ":*")
				err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.Echo(&protoReq, &protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if strings.Contains(r.Header.Get("Accept"), "application/json") {
				w.Header().Set("Content-Type", "application/json")
			} else {
				w.Header().Set("Content-Type", "text/plain")
			}

			if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)

	router.Handle("GET", "/static/:value",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   String
				protoReply StaticFile
			)

			for _, fieldPath := range re.FindAllString("/static/:value", -1) {
				fieldPath := strings.TrimLeft(fieldPath, ":*")
				err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.Static(&protoReq, &protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "ContentType")

			if _, err := w.Write(protoReply.ContentBody); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)

	return router
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_d80cc3bf83ccd227) }

var fileDescriptor_hello_d80cc3bf83ccd227 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xdd, 0x6e, 0x94, 0x40,
	0x18, 0x0d, 0xcb, 0x2e, 0xb6, 0x1f, 0x5b, 0x5d, 0x27, 0x35, 0x12, 0x34, 0x06, 0xf1, 0x86, 0x0b,
	0xc3, 0x54, 0x1a, 0xa3, 0xe1, 0xc6, 0xa8, 0x45, 0x1b, 0x53, 0x13, 0xc3, 0x72, 0xdf, 0xc0, 0x30,
	0xc2, 0x44, 0x64, 0x08, 0xcc, 0x36, 0xe1, 0xc5, 0x7c, 0x0b, 0x5f, 0xc0, 0x87, 0xf0, 0x0d, 0x4c,
	0xcc, 0x0c, 0xcb, 0xee, 0x36, 0xf4, 0x86, 0x70, 0xce, 0x9c, 0x39, 0xdf, 0xcf, 0x1c, 0x30, 0x4b,
	0x5a, 0x55, 0xdc, 0x6f, 0x5a, 0x2e, 0x38, 0x3a, 0x52, 0xe0, 0xba, 0xc9, 0xec, 0x17, 0x05, 0x13,
	0xe5, 0x26, 0xf3, 0x09, 0xff, 0x89, 0x49, 0x99, 0xb2, 0xe0, 0xec, 0xd5, 0x19, 0x6e, 0xb2, 0x82,
	0xab, 0xcf, 0x20, 0x77, 0x9f, 0x81, 0xb1, 0x16, 0x2d, 0xab, 0x0b, 0x74, 0x0a, 0x8b, 0x9b, 0xb4,
	0xda, 0x50, 0x4b, 0x73, 0x34, 0xef, 0x38, 0x1e, 0x80, 0x1b, 0x03, 0xac, 0x45, 0x2a, 0x18, 0xf9,
	0xc4, 0x2a, 0x8a, 0x9e, 0xc3, 0x92, 0xf0, 0x5a, 0xd0, 0x5a, 0x5c, 0x8b, 0xbe, 0x19, 0xa5, 0xe6,
	0x96, 0x4b, 0xfa, 0xe6, 0x96, 0x24, 0xe3, 0x79, 0x6f, 0xcd, 0x1c, 0xcd, 0x5b, 0xee, 0x24, 0x1f,
	0x78, 0xde, 0xbb, 0xbf, 0x35, 0xb8, 0xf7, 0x95, 0x76, 0x5d, 0x5a, 0xd0, 0xbb, 0xab, 0x4a, 0x36,
	0x6d, 0xdb, 0x54, 0xde, 0xd6, 0xbd, 0x45, 0x3c, 0x00, 0x84, 0x61, 0x9e, 0x33, 0x22, 0x2c, 0xdd,
	0xd1, 0x3d, 0x33, 0x78, 0xe2, 0x8f, 0x93, 0xfa, 0x5b, 0x33, 0xff, 0x82, 0x11, 0x11, 0xd5, 0xa2,
	0xed, 0x63, 0x25, 0x44, 0x2f, 0xe1, 0xa8, 0xdb, 0x64, 0xdf, 0x59, 0x45, 0x73, 0x6b, 0xee, 0x68,
	0x9e, 0x19, 0xac, 0xf6, 0x97, 0x86, 0xb1, 0xe3, 0x9d, 0xc2, 0x7e, 0x03, 0xc7, 0x3b, 0x03, 0xb4,
	0x02, 0xfd, 0x07, 0xed, 0xb7, 0x5d, 0xc9, 0xdf, 0x7d, 0xa7, 0xb3, 0x83, 0x4e, 0xc3, 0xd9, 0x5b,
	0x2d, 0xf8, 0x35, 0x83, 0xe5, 0xa5, 0xb4, 0x5d, 0xd3, 0xf6, 0x86, 0x11, 0x8a, 0x2a, 0x58, 0x28,
	0x8c, 0x26, 0xe5, 0xec, 0x09, 0xe3, 0xbe, 0xfb, 0xf3, 0xf7, 0xdf, 0x55, 0x08, 0x27, 0x58, 0x1d,
	0xe0, 0x50, 0xf9, 0xda, 0xc6, 0x00, 0x83, 0x15, 0x18, 0x17, 0xd1, 0x55, 0x94, 0x44, 0x68, 0x64,
	0x1e, 0xc0, 0xe2, 0xdb, 0xfb, 0xe4, 0xe3, 0xe5, 0x48, 0xa0, 0x2f, 0x30, 0x8f, 0x48, 0xc9, 0xd1,
	0xc3, 0xc9, 0x42, 0xec, 0x29, 0xe5, 0x3e, 0x95, 0xe5, 0x1e, 0xc3, 0x23, 0x4c, 0x49, 0xc9, 0x71,
	0x38, 0x4e, 0xef, 0x0f, 0x8b, 0xaf, 0x64, 0x1c, 0xe4, 0x73, 0xdf, 0xd1, 0xfa, 0xe9, 0x21, 0x33,
	0x46, 0xc2, 0x0d, 0xa5, 0xdf, 0xeb, 0xe0, 0x1c, 0xf4, 0xcf, 0x51, 0x82, 0xee, 0xe3, 0x4e, 0x1d,
	0x8d, 0x63, 0x9c, 0x84, 0x87, 0x79, 0x71, 0xf7, 0x50, 0x66, 0x23, 0x33, 0x54, 0x06, 0xcf, 0xff,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x16, 0xdb, 0xb6, 0x60, 0xc1, 0x02, 0x00, 0x00,
}
