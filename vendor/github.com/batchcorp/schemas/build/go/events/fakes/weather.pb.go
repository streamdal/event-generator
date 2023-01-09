// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weather.proto

package fakes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Weather struct {
	Coordinates          *WeatherCoordinates `protobuf:"bytes,1,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Overview             *WeatherOverview    `protobuf:"bytes,2,opt,name=overview,proto3" json:"overview,omitempty"`
	Base                 *WeatherBase        `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	Info                 *WeatherInfo        `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Visibility           int32               `protobuf:"varint,5,opt,name=visibility,proto3" json:"visibility,omitempty"`
	Wind                 *WeatherWind        `protobuf:"bytes,6,opt,name=wind,proto3" json:"wind,omitempty"`
	Rain                 *WeatherRain        `protobuf:"bytes,7,opt,name=rain,proto3" json:"rain,omitempty"`
	Clouds               *WeatherClouds      `protobuf:"bytes,8,opt,name=clouds,proto3" json:"clouds,omitempty"`
	Timezone             string              `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	UnixTimeNanoUtc      int64               `protobuf:"varint,10,opt,name=unix_time_nano_utc,json=unixTimeNanoUtc,proto3" json:"unix_time_nano_utc,omitempty"`
	System               *WeatherSystem      `protobuf:"bytes,11,opt,name=system,proto3" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Weather) Reset()         { *m = Weather{} }
func (m *Weather) String() string { return proto.CompactTextString(m) }
func (*Weather) ProtoMessage()    {}
func (*Weather) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{0}
}

func (m *Weather) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Weather.Unmarshal(m, b)
}
func (m *Weather) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Weather.Marshal(b, m, deterministic)
}
func (m *Weather) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Weather.Merge(m, src)
}
func (m *Weather) XXX_Size() int {
	return xxx_messageInfo_Weather.Size(m)
}
func (m *Weather) XXX_DiscardUnknown() {
	xxx_messageInfo_Weather.DiscardUnknown(m)
}

var xxx_messageInfo_Weather proto.InternalMessageInfo

func (m *Weather) GetCoordinates() *WeatherCoordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Weather) GetOverview() *WeatherOverview {
	if m != nil {
		return m.Overview
	}
	return nil
}

func (m *Weather) GetBase() *WeatherBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *Weather) GetInfo() *WeatherInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Weather) GetVisibility() int32 {
	if m != nil {
		return m.Visibility
	}
	return 0
}

func (m *Weather) GetWind() *WeatherWind {
	if m != nil {
		return m.Wind
	}
	return nil
}

func (m *Weather) GetRain() *WeatherRain {
	if m != nil {
		return m.Rain
	}
	return nil
}

func (m *Weather) GetClouds() *WeatherClouds {
	if m != nil {
		return m.Clouds
	}
	return nil
}

func (m *Weather) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Weather) GetUnixTimeNanoUtc() int64 {
	if m != nil {
		return m.UnixTimeNanoUtc
	}
	return 0
}

func (m *Weather) GetSystem() *WeatherSystem {
	if m != nil {
		return m.System
	}
	return nil
}

type WeatherCoordinates struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherCoordinates) Reset()         { *m = WeatherCoordinates{} }
func (m *WeatherCoordinates) String() string { return proto.CompactTextString(m) }
func (*WeatherCoordinates) ProtoMessage()    {}
func (*WeatherCoordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{1}
}

func (m *WeatherCoordinates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherCoordinates.Unmarshal(m, b)
}
func (m *WeatherCoordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherCoordinates.Marshal(b, m, deterministic)
}
func (m *WeatherCoordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherCoordinates.Merge(m, src)
}
func (m *WeatherCoordinates) XXX_Size() int {
	return xxx_messageInfo_WeatherCoordinates.Size(m)
}
func (m *WeatherCoordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherCoordinates.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherCoordinates proto.InternalMessageInfo

func (m *WeatherCoordinates) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *WeatherCoordinates) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type WeatherOverview struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	IconUrl              string   `protobuf:"bytes,3,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherOverview) Reset()         { *m = WeatherOverview{} }
func (m *WeatherOverview) String() string { return proto.CompactTextString(m) }
func (*WeatherOverview) ProtoMessage()    {}
func (*WeatherOverview) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{2}
}

func (m *WeatherOverview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherOverview.Unmarshal(m, b)
}
func (m *WeatherOverview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherOverview.Marshal(b, m, deterministic)
}
func (m *WeatherOverview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherOverview.Merge(m, src)
}
func (m *WeatherOverview) XXX_Size() int {
	return xxx_messageInfo_WeatherOverview.Size(m)
}
func (m *WeatherOverview) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherOverview.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherOverview proto.InternalMessageInfo

func (m *WeatherOverview) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WeatherOverview) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WeatherOverview) GetIconUrl() string {
	if m != nil {
		return m.IconUrl
	}
	return ""
}

type WeatherBase struct {
	StationId            string   `protobuf:"bytes,1,opt,name=station_id,json=stationId,proto3" json:"station_id,omitempty"`
	StationName          string   `protobuf:"bytes,2,opt,name=station_name,json=stationName,proto3" json:"station_name,omitempty"`
	StationCountry       string   `protobuf:"bytes,3,opt,name=station_country,json=stationCountry,proto3" json:"station_country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherBase) Reset()         { *m = WeatherBase{} }
func (m *WeatherBase) String() string { return proto.CompactTextString(m) }
func (*WeatherBase) ProtoMessage()    {}
func (*WeatherBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{3}
}

func (m *WeatherBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherBase.Unmarshal(m, b)
}
func (m *WeatherBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherBase.Marshal(b, m, deterministic)
}
func (m *WeatherBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherBase.Merge(m, src)
}
func (m *WeatherBase) XXX_Size() int {
	return xxx_messageInfo_WeatherBase.Size(m)
}
func (m *WeatherBase) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherBase.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherBase proto.InternalMessageInfo

func (m *WeatherBase) GetStationId() string {
	if m != nil {
		return m.StationId
	}
	return ""
}

func (m *WeatherBase) GetStationName() string {
	if m != nil {
		return m.StationName
	}
	return ""
}

func (m *WeatherBase) GetStationCountry() string {
	if m != nil {
		return m.StationCountry
	}
	return ""
}

type WeatherInfo struct {
	TemperatureC         float32  `protobuf:"fixed32,1,opt,name=temperature_c,json=temperatureC,proto3" json:"temperature_c,omitempty"`
	FeelsLikeC           float32  `protobuf:"fixed32,2,opt,name=feels_like_c,json=feelsLikeC,proto3" json:"feels_like_c,omitempty"`
	TempMin              float32  `protobuf:"fixed32,3,opt,name=temp_min,json=tempMin,proto3" json:"temp_min,omitempty"`
	TempMax              float32  `protobuf:"fixed32,4,opt,name=temp_max,json=tempMax,proto3" json:"temp_max,omitempty"`
	Pressure             int32    `protobuf:"varint,5,opt,name=pressure,proto3" json:"pressure,omitempty"`
	Humidity             int32    `protobuf:"varint,6,opt,name=humidity,proto3" json:"humidity,omitempty"`
	SeaLevel             int32    `protobuf:"varint,7,opt,name=sea_level,json=seaLevel,proto3" json:"sea_level,omitempty"`
	GroundLevel          int32    `protobuf:"varint,8,opt,name=ground_level,json=groundLevel,proto3" json:"ground_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherInfo) Reset()         { *m = WeatherInfo{} }
func (m *WeatherInfo) String() string { return proto.CompactTextString(m) }
func (*WeatherInfo) ProtoMessage()    {}
func (*WeatherInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{4}
}

func (m *WeatherInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherInfo.Unmarshal(m, b)
}
func (m *WeatherInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherInfo.Marshal(b, m, deterministic)
}
func (m *WeatherInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherInfo.Merge(m, src)
}
func (m *WeatherInfo) XXX_Size() int {
	return xxx_messageInfo_WeatherInfo.Size(m)
}
func (m *WeatherInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherInfo proto.InternalMessageInfo

func (m *WeatherInfo) GetTemperatureC() float32 {
	if m != nil {
		return m.TemperatureC
	}
	return 0
}

func (m *WeatherInfo) GetFeelsLikeC() float32 {
	if m != nil {
		return m.FeelsLikeC
	}
	return 0
}

func (m *WeatherInfo) GetTempMin() float32 {
	if m != nil {
		return m.TempMin
	}
	return 0
}

func (m *WeatherInfo) GetTempMax() float32 {
	if m != nil {
		return m.TempMax
	}
	return 0
}

func (m *WeatherInfo) GetPressure() int32 {
	if m != nil {
		return m.Pressure
	}
	return 0
}

func (m *WeatherInfo) GetHumidity() int32 {
	if m != nil {
		return m.Humidity
	}
	return 0
}

func (m *WeatherInfo) GetSeaLevel() int32 {
	if m != nil {
		return m.SeaLevel
	}
	return 0
}

func (m *WeatherInfo) GetGroundLevel() int32 {
	if m != nil {
		return m.GroundLevel
	}
	return 0
}

type WeatherWind struct {
	Speed                float32  `protobuf:"fixed32,1,opt,name=speed,proto3" json:"speed,omitempty"`
	Deg                  int32    `protobuf:"varint,2,opt,name=deg,proto3" json:"deg,omitempty"`
	Gust                 float32  `protobuf:"fixed32,3,opt,name=gust,proto3" json:"gust,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherWind) Reset()         { *m = WeatherWind{} }
func (m *WeatherWind) String() string { return proto.CompactTextString(m) }
func (*WeatherWind) ProtoMessage()    {}
func (*WeatherWind) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{5}
}

func (m *WeatherWind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherWind.Unmarshal(m, b)
}
func (m *WeatherWind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherWind.Marshal(b, m, deterministic)
}
func (m *WeatherWind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherWind.Merge(m, src)
}
func (m *WeatherWind) XXX_Size() int {
	return xxx_messageInfo_WeatherWind.Size(m)
}
func (m *WeatherWind) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherWind.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherWind proto.InternalMessageInfo

func (m *WeatherWind) GetSpeed() float32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *WeatherWind) GetDeg() int32 {
	if m != nil {
		return m.Deg
	}
	return 0
}

func (m *WeatherWind) GetGust() float32 {
	if m != nil {
		return m.Gust
	}
	return 0
}

type WeatherRain struct {
	OneHour              float32  `protobuf:"fixed32,1,opt,name=one_hour,json=oneHour,proto3" json:"one_hour,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherRain) Reset()         { *m = WeatherRain{} }
func (m *WeatherRain) String() string { return proto.CompactTextString(m) }
func (*WeatherRain) ProtoMessage()    {}
func (*WeatherRain) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{6}
}

func (m *WeatherRain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherRain.Unmarshal(m, b)
}
func (m *WeatherRain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherRain.Marshal(b, m, deterministic)
}
func (m *WeatherRain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherRain.Merge(m, src)
}
func (m *WeatherRain) XXX_Size() int {
	return xxx_messageInfo_WeatherRain.Size(m)
}
func (m *WeatherRain) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherRain.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherRain proto.InternalMessageInfo

func (m *WeatherRain) GetOneHour() float32 {
	if m != nil {
		return m.OneHour
	}
	return 0
}

type WeatherClouds struct {
	All                  int32    `protobuf:"varint,1,opt,name=all,proto3" json:"all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherClouds) Reset()         { *m = WeatherClouds{} }
func (m *WeatherClouds) String() string { return proto.CompactTextString(m) }
func (*WeatherClouds) ProtoMessage()    {}
func (*WeatherClouds) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{7}
}

func (m *WeatherClouds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherClouds.Unmarshal(m, b)
}
func (m *WeatherClouds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherClouds.Marshal(b, m, deterministic)
}
func (m *WeatherClouds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherClouds.Merge(m, src)
}
func (m *WeatherClouds) XXX_Size() int {
	return xxx_messageInfo_WeatherClouds.Size(m)
}
func (m *WeatherClouds) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherClouds.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherClouds proto.InternalMessageInfo

func (m *WeatherClouds) GetAll() int32 {
	if m != nil {
		return m.All
	}
	return 0
}

type WeatherSystem struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Message              float32  `protobuf:"fixed32,3,opt,name=message,proto3" json:"message,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Sunrise              int32    `protobuf:"varint,5,opt,name=sunrise,proto3" json:"sunrise,omitempty"`
	Sunset               int32    `protobuf:"varint,6,opt,name=sunset,proto3" json:"sunset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherSystem) Reset()         { *m = WeatherSystem{} }
func (m *WeatherSystem) String() string { return proto.CompactTextString(m) }
func (*WeatherSystem) ProtoMessage()    {}
func (*WeatherSystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{8}
}

func (m *WeatherSystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherSystem.Unmarshal(m, b)
}
func (m *WeatherSystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherSystem.Marshal(b, m, deterministic)
}
func (m *WeatherSystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherSystem.Merge(m, src)
}
func (m *WeatherSystem) XXX_Size() int {
	return xxx_messageInfo_WeatherSystem.Size(m)
}
func (m *WeatherSystem) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherSystem.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherSystem proto.InternalMessageInfo

func (m *WeatherSystem) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *WeatherSystem) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WeatherSystem) GetMessage() float32 {
	if m != nil {
		return m.Message
	}
	return 0
}

func (m *WeatherSystem) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *WeatherSystem) GetSunrise() int32 {
	if m != nil {
		return m.Sunrise
	}
	return 0
}

func (m *WeatherSystem) GetSunset() int32 {
	if m != nil {
		return m.Sunset
	}
	return 0
}

func init() {
	proto.RegisterType((*Weather)(nil), "fakes.Weather")
	proto.RegisterType((*WeatherCoordinates)(nil), "fakes.WeatherCoordinates")
	proto.RegisterType((*WeatherOverview)(nil), "fakes.WeatherOverview")
	proto.RegisterType((*WeatherBase)(nil), "fakes.WeatherBase")
	proto.RegisterType((*WeatherInfo)(nil), "fakes.WeatherInfo")
	proto.RegisterType((*WeatherWind)(nil), "fakes.WeatherWind")
	proto.RegisterType((*WeatherRain)(nil), "fakes.WeatherRain")
	proto.RegisterType((*WeatherClouds)(nil), "fakes.WeatherClouds")
	proto.RegisterType((*WeatherSystem)(nil), "fakes.WeatherSystem")
}

func init() { proto.RegisterFile("weather.proto", fileDescriptor_231dcd72b885f4be) }

var fileDescriptor_231dcd72b885f4be = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4b, 0x6b, 0x2b, 0x37,
	0x14, 0xc6, 0x6f, 0xcf, 0x71, 0x1e, 0x45, 0x84, 0x30, 0xe9, 0x0b, 0x67, 0x0a, 0xad, 0xa1, 0xc5,
	0x86, 0xb4, 0xbb, 0xee, 0xe2, 0x4d, 0x03, 0x69, 0x0a, 0x6a, 0x43, 0xa0, 0x8b, 0x0e, 0xf2, 0xcc,
	0xb1, 0x2d, 0x32, 0x23, 0x19, 0x3d, 0x1c, 0xbb, 0xcb, 0xfb, 0x2f, 0xee, 0xaf, 0xbb, 0x7f, 0xe5,
	0x22, 0x8d, 0x66, 0xe2, 0xe4, 0x7a, 0xa7, 0xef, 0xa1, 0x73, 0x46, 0x47, 0x9f, 0x06, 0x4e, 0x5f,
	0x90, 0x99, 0x35, 0xaa, 0xe9, 0x46, 0x49, 0x23, 0x49, 0x6f, 0xc9, 0x9e, 0x51, 0x27, 0x9f, 0x3a,
	0x30, 0x78, 0xaa, 0x04, 0xf2, 0x3b, 0x8c, 0x32, 0x29, 0x55, 0xce, 0x05, 0x33, 0xa8, 0xe3, 0xd6,
	0xb8, 0x35, 0x19, 0xdd, 0x5c, 0x4d, 0xbd, 0x71, 0x1a, 0x4c, 0xf3, 0x57, 0x03, 0x3d, 0x74, 0x93,
	0x1b, 0x18, 0xca, 0x2d, 0xaa, 0x2d, 0xc7, 0x97, 0xb8, 0xed, 0x77, 0x5e, 0xbe, 0xdd, 0xf9, 0x57,
	0x50, 0x69, 0xe3, 0x23, 0x3f, 0x42, 0x77, 0xc1, 0x34, 0xc6, 0x1d, 0xef, 0x27, 0x6f, 0xfd, 0xb7,
	0x4c, 0x23, 0xf5, 0xba, 0xf3, 0x71, 0xb1, 0x94, 0x71, 0xf7, 0x98, 0xef, 0x4e, 0x2c, 0x25, 0xf5,
	0x3a, 0xf9, 0x1e, 0x60, 0xcb, 0x35, 0x5f, 0xf0, 0x82, 0x9b, 0x7d, 0xdc, 0x1b, 0xb7, 0x26, 0x3d,
	0x7a, 0xc0, 0xb8, 0x3a, 0x2f, 0x5c, 0xe4, 0x71, 0xff, 0x58, 0x9d, 0x27, 0x2e, 0x72, 0xea, 0x75,
	0xe7, 0x53, 0x8c, 0x8b, 0x78, 0x70, 0xcc, 0x47, 0x19, 0x17, 0xd4, 0xeb, 0xe4, 0x17, 0xe8, 0x67,
	0x85, 0xb4, 0xb9, 0x8e, 0x87, 0xde, 0x79, 0xf1, 0x6e, 0x56, 0x5e, 0xa3, 0xc1, 0x43, 0xbe, 0x86,
	0xa1, 0xe1, 0x25, 0xfe, 0x2f, 0x05, 0xc6, 0xd1, 0xb8, 0x35, 0x89, 0x68, 0x83, 0xc9, 0xcf, 0x40,
	0xac, 0xe0, 0xbb, 0xd4, 0x11, 0xa9, 0x60, 0x42, 0xa6, 0xd6, 0x64, 0x31, 0x8c, 0x5b, 0x93, 0x0e,
	0x3d, 0x77, 0xca, 0x3f, 0xbc, 0xc4, 0x07, 0x26, 0xe4, 0xa3, 0xc9, 0x5c, 0x5b, 0xbd, 0xd7, 0x06,
	0xcb, 0x78, 0x74, 0xac, 0xed, 0xdf, 0x5e, 0xa3, 0xc1, 0x93, 0x3c, 0x00, 0xf9, 0xf2, 0xee, 0xdc,
	0xc7, 0x14, 0xcc, 0x70, 0x63, 0x73, 0xf4, 0x17, 0xdd, 0xa6, 0x0d, 0x26, 0xdf, 0x42, 0x54, 0x48,
	0xb1, 0xaa, 0xc4, 0xb6, 0x17, 0x5f, 0x89, 0xe4, 0x3f, 0x38, 0x7f, 0x77, 0xa3, 0xe4, 0x0c, 0xda,
	0x3c, 0xf7, 0x65, 0x22, 0xda, 0xe6, 0x39, 0x19, 0xc3, 0x28, 0x47, 0x9d, 0x29, 0xbe, 0x31, 0x5c,
	0x0a, 0x5f, 0x22, 0xa2, 0x87, 0x14, 0xb9, 0x82, 0x21, 0xcf, 0xa4, 0x48, 0xad, 0x2a, 0xfc, 0xed,
	0x47, 0x74, 0xe0, 0xf0, 0xa3, 0x2a, 0x92, 0x1d, 0x8c, 0x0e, 0x12, 0x40, 0xbe, 0x03, 0xd0, 0x86,
	0xb9, 0x4d, 0x69, 0xd3, 0x23, 0x0a, 0xcc, 0x5d, 0x4e, 0xae, 0xe1, 0xa4, 0x96, 0x05, 0x2b, 0xb1,
	0xee, 0x15, 0xb8, 0x07, 0x56, 0x22, 0xf9, 0x09, 0xce, 0x6b, 0x4b, 0x26, 0xad, 0x30, 0x6a, 0x1f,
	0x5a, 0x9e, 0x05, 0x7a, 0x5e, 0xb1, 0xc9, 0x87, 0x76, 0xd3, 0xda, 0x85, 0x8a, 0xfc, 0x00, 0xa7,
	0x06, 0xcb, 0x0d, 0x2a, 0x66, 0xac, 0xc2, 0x34, 0x0b, 0x83, 0x3a, 0x39, 0x20, 0xe7, 0x64, 0x0c,
	0x27, 0x4b, 0xc4, 0x42, 0xa7, 0x05, 0x7f, 0x76, 0x9e, 0x6a, 0x5e, 0xe0, 0xb9, 0x7b, 0xfe, 0x8c,
	0x73, 0x77, 0x56, 0xb7, 0x23, 0x2d, 0xb9, 0xf0, 0x8d, 0xdb, 0x74, 0xe0, 0xf0, 0x9f, 0x5c, 0xbc,
	0x4a, 0x6c, 0xe7, 0xc3, 0x5d, 0x4b, 0x6c, 0xe7, 0x2e, 0x68, 0xa3, 0x50, 0x6b, 0xab, 0x30, 0x24,
	0xb9, 0xc1, 0x4e, 0x5b, 0xdb, 0x92, 0xe7, 0x2e, 0xe5, 0xfd, 0x4a, 0xab, 0x31, 0xf9, 0x06, 0x22,
	0x8d, 0x2c, 0x2d, 0x70, 0x8b, 0x85, 0x0f, 0x70, 0x8f, 0x0e, 0x35, 0xb2, 0x7b, 0x87, 0xdd, 0xb4,
	0x56, 0x4a, 0x5a, 0x91, 0x07, 0x7d, 0xe8, 0xf5, 0x51, 0xc5, 0x79, 0x4b, 0x72, 0xd7, 0xcc, 0xc0,
	0x3d, 0x08, 0x72, 0x01, 0x3d, 0xbd, 0x41, 0xcc, 0xc3, 0xd9, 0x2b, 0x40, 0xbe, 0x82, 0x4e, 0x8e,
	0x2b, 0x7f, 0xd6, 0x1e, 0x75, 0x4b, 0x42, 0xa0, 0xbb, 0xb2, 0xda, 0x84, 0x03, 0xfa, 0x75, 0x32,
	0x69, 0x4a, 0xb9, 0x37, 0xe3, 0x0e, 0x2b, 0x05, 0xa6, 0x6b, 0x69, 0x55, 0xa8, 0x36, 0x90, 0x02,
	0xff, 0x90, 0x56, 0x25, 0xd7, 0x70, 0xfa, 0xe6, 0xcd, 0xb8, 0x06, 0xac, 0x28, 0xbc, 0xad, 0x47,
	0xdd, 0x32, 0xf9, 0xd8, 0x6a, 0x3c, 0x55, 0xc0, 0x5d, 0x4b, 0xb3, 0xdf, 0x60, 0x30, 0xf9, 0x75,
	0x48, 0x62, 0xf5, 0x5d, 0x2e, 0x89, 0x31, 0x0c, 0x4a, 0xd4, 0x9a, 0xad, 0xb0, 0x1e, 0x7d, 0x80,
	0x4e, 0xa9, 0xd3, 0xd0, 0xad, 0x02, 0x18, 0xa0, 0x53, 0xb4, 0x15, 0x8a, 0xeb, 0x7a, 0xf0, 0x35,
	0x24, 0x97, 0xd0, 0xd7, 0x56, 0x68, 0x34, 0x61, 0xea, 0x01, 0xdd, 0xfe, 0xf6, 0xef, 0xcd, 0x8a,
	0x9b, 0xb5, 0x5d, 0x4c, 0x33, 0x59, 0xce, 0x16, 0xcc, 0x64, 0xeb, 0x4c, 0xaa, 0xcd, 0x4c, 0x67,
	0x6b, 0x2c, 0x99, 0x9e, 0x2d, 0x2c, 0x2f, 0xf2, 0xd9, 0x4a, 0xce, 0x70, 0x8b, 0xc2, 0xe8, 0x99,
	0x7f, 0xae, 0x8b, 0xbe, 0xff, 0x11, 0xff, 0xfa, 0x39, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xf8, 0x3f,
	0xdf, 0x99, 0x05, 0x00, 0x00,
}