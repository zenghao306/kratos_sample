// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.0
// source: conf.proto

package conf

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Env           *Env                   `protobuf:"bytes,1,opt,name=env,proto3" json:"env,omitempty"`
	Service       *Service               `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Trace         *Trace                 `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
	Server        *Server                `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`
	Data          *Data                  `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Etcd          *Etcd                  `protobuf:"bytes,6,opt,name=etcd,proto3" json:"etcd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetEnv() *Env {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Bootstrap) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Bootstrap) GetTrace() *Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetEtcd() *Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gateway       *Service_Endpoint      `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway,omitempty"`
	User          *Service_Endpoint      `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Device        *Service_Endpoint      `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Service) Reset() {
	*x = Service{}
	mi := &file_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Service) GetGateway() *Service_Endpoint {
	if x != nil {
		return x.Gateway
	}
	return nil
}

func (x *Service) GetUser() *Service_Endpoint {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Service) GetDevice() *Service_Endpoint {
	if x != nil {
		return x.Device
	}
	return nil
}

type Env struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mode          string                 `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Env) Reset() {
	*x = Env{}
	mi := &file_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Env) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Env) ProtoMessage() {}

func (x *Env) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Env.ProtoReflect.Descriptor instead.
func (*Env) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Env) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GatewayHttp   *Server_HTTP           `protobuf:"bytes,1,opt,name=gateway_http,json=gatewayHttp,proto3" json:"gateway_http,omitempty"`
	DeviceGrpc    *Server_GRPC           `protobuf:"bytes,2,opt,name=device_grpc,json=deviceGrpc,proto3" json:"device_grpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Server) GetGatewayHttp() *Server_HTTP {
	if x != nil {
		return x.GatewayHttp
	}
	return nil
}

func (x *Server) GetDeviceGrpc() *Server_GRPC {
	if x != nil {
		return x.DeviceGrpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Database      *Data_Database         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis         *Data_Redis            `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_conf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type Etcd struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Etcd) Reset() {
	*x = Etcd{}
	mi := &file_conf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etcd) ProtoMessage() {}

func (x *Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etcd.ProtoReflect.Descriptor instead.
func (*Etcd) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{5}
}

func (x *Etcd) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Trace struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trace) Reset() {
	*x = Trace{}
	mi := &file_conf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{6}
}

func (x *Trace) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Service_Endpoint struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Service_Endpoint) Reset() {
	*x = Service_Endpoint{}
	mi := &file_conf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service_Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service_Endpoint) ProtoMessage() {}

func (x *Service_Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service_Endpoint.ProtoReflect.Descriptor instead.
func (*Service_Endpoint) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Service_Endpoint) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Server_HTTP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *duration.Duration     `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	mi := &file_conf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *duration.Duration     `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	mi := &file_conf_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source        string                 `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	mi := &file_conf_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Db            int32                  `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`
	DialTimeout   *duration.Duration     `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`
	ReadTimeout   *duration.Duration     `protobuf:"bytes,6,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout  *duration.Duration     `protobuf:"bytes,7,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	TvPlayerTopic string                 `protobuf:"bytes,8,opt,name=tv_player_topic,json=tvPlayerTopic,proto3" json:"tv_player_topic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	mi := &file_conf_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{4, 1}
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Data_Redis) GetDialTimeout() *duration.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Data_Redis) GetReadTimeout() *duration.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *duration.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Data_Redis) GetTvPlayerTopic() string {
	if x != nil {
		return x.TvPlayerTopic
	}
	return ""
}

var File_conf_proto protoreflect.FileDescriptor

const file_conf_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"conf.proto\x12\n" +
	"kratos.api\x1a\x1egoogle/protobuf/duration.proto\"\xfe\x01\n" +
	"\tBootstrap\x12!\n" +
	"\x03env\x18\x01 \x01(\v2\x0f.kratos.api.EnvR\x03env\x12-\n" +
	"\aservice\x18\x02 \x01(\v2\x13.kratos.api.ServiceR\aservice\x12'\n" +
	"\x05trace\x18\x03 \x01(\v2\x11.kratos.api.TraceR\x05trace\x12*\n" +
	"\x06server\x18\x04 \x01(\v2\x12.kratos.api.ServerR\x06server\x12$\n" +
	"\x04data\x18\x05 \x01(\v2\x10.kratos.api.DataR\x04data\x12$\n" +
	"\x04etcd\x18\x06 \x01(\v2\x10.kratos.api.EtcdR\x04etcd\"\xd1\x01\n" +
	"\aService\x126\n" +
	"\agateway\x18\x01 \x01(\v2\x1c.kratos.api.Service.EndpointR\agateway\x120\n" +
	"\x04user\x18\x02 \x01(\v2\x1c.kratos.api.Service.EndpointR\x04user\x124\n" +
	"\x06device\x18\x03 \x01(\v2\x1c.kratos.api.Service.EndpointR\x06device\x1a&\n" +
	"\bEndpoint\x12\x1a\n" +
	"\bendpoint\x18\x01 \x01(\tR\bendpoint\"\x19\n" +
	"\x03Env\x12\x12\n" +
	"\x04mode\x18\x01 \x01(\tR\x04mode\"\xd4\x02\n" +
	"\x06Server\x12:\n" +
	"\fgateway_http\x18\x01 \x01(\v2\x17.kratos.api.Server.HTTPR\vgatewayHttp\x128\n" +
	"\vdevice_grpc\x18\x02 \x01(\v2\x17.kratos.api.Server.GRPCR\n" +
	"deviceGrpc\x1ai\n" +
	"\x04HTTP\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x123\n" +
	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\x1ai\n" +
	"\x04GRPC\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x123\n" +
	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\"\xd5\x03\n" +
	"\x04Data\x125\n" +
	"\bdatabase\x18\x01 \x01(\v2\x19.kratos.api.Data.DatabaseR\bdatabase\x12,\n" +
	"\x05redis\x18\x02 \x01(\v2\x16.kratos.api.Data.RedisR\x05redis\x1a:\n" +
	"\bDatabase\x12\x16\n" +
	"\x06driver\x18\x01 \x01(\tR\x06driver\x12\x16\n" +
	"\x06source\x18\x02 \x01(\tR\x06source\x1a\xab\x02\n" +
	"\x05Redis\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x0e\n" +
	"\x02db\x18\x04 \x01(\x05R\x02db\x12<\n" +
	"\fdial_timeout\x18\x05 \x01(\v2\x19.google.protobuf.DurationR\vdialTimeout\x12<\n" +
	"\fread_timeout\x18\x06 \x01(\v2\x19.google.protobuf.DurationR\vreadTimeout\x12>\n" +
	"\rwrite_timeout\x18\a \x01(\v2\x19.google.protobuf.DurationR\fwriteTimeout\x12&\n" +
	"\x0ftv_player_topic\x18\b \x01(\tR\rtvPlayerTopic\" \n" +
	"\x04Etcd\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\"#\n" +
	"\x05Trace\x12\x1a\n" +
	"\bendpoint\x18\x01 \x01(\tR\bendpointB(Z&dskratos/app/device/internal/conf;confb\x06proto3"

var (
	file_conf_proto_rawDescOnce sync.Once
	file_conf_proto_rawDescData []byte
)

func file_conf_proto_rawDescGZIP() []byte {
	file_conf_proto_rawDescOnce.Do(func() {
		file_conf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_conf_proto_rawDesc), len(file_conf_proto_rawDesc)))
	})
	return file_conf_proto_rawDescData
}

var file_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_conf_proto_goTypes = []any{
	(*Bootstrap)(nil),         // 0: kratos.api.Bootstrap
	(*Service)(nil),           // 1: kratos.api.Service
	(*Env)(nil),               // 2: kratos.api.Env
	(*Server)(nil),            // 3: kratos.api.Server
	(*Data)(nil),              // 4: kratos.api.Data
	(*Etcd)(nil),              // 5: kratos.api.Etcd
	(*Trace)(nil),             // 6: kratos.api.Trace
	(*Service_Endpoint)(nil),  // 7: kratos.api.Service.Endpoint
	(*Server_HTTP)(nil),       // 8: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),       // 9: kratos.api.Server.GRPC
	(*Data_Database)(nil),     // 10: kratos.api.Data.Database
	(*Data_Redis)(nil),        // 11: kratos.api.Data.Redis
	(*duration.Duration)(nil), // 12: google.protobuf.Duration
}
var file_conf_proto_depIdxs = []int32{
	2,  // 0: kratos.api.Bootstrap.env:type_name -> kratos.api.Env
	1,  // 1: kratos.api.Bootstrap.service:type_name -> kratos.api.Service
	6,  // 2: kratos.api.Bootstrap.trace:type_name -> kratos.api.Trace
	3,  // 3: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	4,  // 4: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	5,  // 5: kratos.api.Bootstrap.etcd:type_name -> kratos.api.Etcd
	7,  // 6: kratos.api.Service.gateway:type_name -> kratos.api.Service.Endpoint
	7,  // 7: kratos.api.Service.user:type_name -> kratos.api.Service.Endpoint
	7,  // 8: kratos.api.Service.device:type_name -> kratos.api.Service.Endpoint
	8,  // 9: kratos.api.Server.gateway_http:type_name -> kratos.api.Server.HTTP
	9,  // 10: kratos.api.Server.device_grpc:type_name -> kratos.api.Server.GRPC
	10, // 11: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	11, // 12: kratos.api.Data.redis:type_name -> kratos.api.Data.Redis
	12, // 13: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	12, // 14: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	12, // 15: kratos.api.Data.Redis.dial_timeout:type_name -> google.protobuf.Duration
	12, // 16: kratos.api.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	12, // 17: kratos.api.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_conf_proto_init() }
func file_conf_proto_init() {
	if File_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_proto_rawDesc), len(file_conf_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_proto_goTypes,
		DependencyIndexes: file_conf_proto_depIdxs,
		MessageInfos:      file_conf_proto_msgTypes,
	}.Build()
	File_conf_proto = out.File
	file_conf_proto_goTypes = nil
	file_conf_proto_depIdxs = nil
}
