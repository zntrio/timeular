// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: timeular/events/v1/events.proto

package eventsv1

import (
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

// EventType enumerates known device events.
type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED           EventType = 0
	EventType_EVENT_TYPE_DEVICE_CONNECTED      EventType = 1
	EventType_EVENT_TYPE_DEVICE_DISCONNECTED   EventType = 2
	EventType_EVENT_TYPE_BATTERY_LEVEL_CHANGED EventType = 3
	EventType_EVENT_TYPE_ORIENTATION_CHANGED   EventType = 4
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNSPECIFIED",
		1: "EVENT_TYPE_DEVICE_CONNECTED",
		2: "EVENT_TYPE_DEVICE_DISCONNECTED",
		3: "EVENT_TYPE_BATTERY_LEVEL_CHANGED",
		4: "EVENT_TYPE_ORIENTATION_CHANGED",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED":           0,
		"EVENT_TYPE_DEVICE_CONNECTED":      1,
		"EVENT_TYPE_DEVICE_DISCONNECTED":   2,
		"EVENT_TYPE_BATTERY_LEVEL_CHANGED": 3,
		"EVENT_TYPE_ORIENTATION_CHANGED":   4,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_timeular_events_v1_events_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_timeular_events_v1_events_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_timeular_events_v1_events_proto_rawDescGZIP(), []int{0}
}

// Event represents device events broadcasted by the device.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique event identifier.
	EventId []byte `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Event type to determine the payload type.
	EventType EventType `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3,enum=timeular.events.v1.EventType" json:"event_type,omitempty"`
	// Event timestamp.
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Event payload matching the event type.
	//
	// Types that are assignable to Payload:
	//
	//	*Event_DeviceConnected
	//	*Event_DeviceDisconnected
	//	*Event_BatteryLevelChanged
	//	*Event_OrientationChanged
	Payload isEvent_Payload `protobuf_oneof:"payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeular_events_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_timeular_events_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_timeular_events_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEventId() []byte {
	if x != nil {
		return x.EventId
	}
	return nil
}

func (x *Event) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_EVENT_TYPE_UNSPECIFIED
}

func (x *Event) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetDeviceConnected() *DeviceConnectedPayload {
	if x, ok := x.GetPayload().(*Event_DeviceConnected); ok {
		return x.DeviceConnected
	}
	return nil
}

func (x *Event) GetDeviceDisconnected() *DeviceDisconnectedPayload {
	if x, ok := x.GetPayload().(*Event_DeviceDisconnected); ok {
		return x.DeviceDisconnected
	}
	return nil
}

func (x *Event) GetBatteryLevelChanged() *BatteryLevelChangedPayload {
	if x, ok := x.GetPayload().(*Event_BatteryLevelChanged); ok {
		return x.BatteryLevelChanged
	}
	return nil
}

func (x *Event) GetOrientationChanged() *OrientationChangedPayload {
	if x, ok := x.GetPayload().(*Event_OrientationChanged); ok {
		return x.OrientationChanged
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_DeviceConnected struct {
	DeviceConnected *DeviceConnectedPayload `protobuf:"bytes,10,opt,name=device_connected,json=deviceConnected,proto3,oneof"`
}

type Event_DeviceDisconnected struct {
	DeviceDisconnected *DeviceDisconnectedPayload `protobuf:"bytes,11,opt,name=device_disconnected,json=deviceDisconnected,proto3,oneof"`
}

type Event_BatteryLevelChanged struct {
	BatteryLevelChanged *BatteryLevelChangedPayload `protobuf:"bytes,12,opt,name=battery_level_changed,json=batteryLevelChanged,proto3,oneof"`
}

type Event_OrientationChanged struct {
	OrientationChanged *OrientationChangedPayload `protobuf:"bytes,13,opt,name=orientation_changed,json=orientationChanged,proto3,oneof"`
}

func (*Event_DeviceConnected) isEvent_Payload() {}

func (*Event_DeviceDisconnected) isEvent_Payload() {}

func (*Event_BatteryLevelChanged) isEvent_Payload() {}

func (*Event_OrientationChanged) isEvent_Payload() {}

// DeviceConnectedPayload holds the properties of a device connection event.
type DeviceConnectedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeviceConnectedPayload) Reset() {
	*x = DeviceConnectedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeular_events_v1_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceConnectedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceConnectedPayload) ProtoMessage() {}

func (x *DeviceConnectedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_timeular_events_v1_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceConnectedPayload.ProtoReflect.Descriptor instead.
func (*DeviceConnectedPayload) Descriptor() ([]byte, []int) {
	return file_timeular_events_v1_events_proto_rawDescGZIP(), []int{1}
}

// DeviceDisconnectedPayload holds the properties of a device disconnection
// event.
type DeviceDisconnectedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeviceDisconnectedPayload) Reset() {
	*x = DeviceDisconnectedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeular_events_v1_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDisconnectedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDisconnectedPayload) ProtoMessage() {}

func (x *DeviceDisconnectedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_timeular_events_v1_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDisconnectedPayload.ProtoReflect.Descriptor instead.
func (*DeviceDisconnectedPayload) Descriptor() ([]byte, []int) {
	return file_timeular_events_v1_events_proto_rawDescGZIP(), []int{2}
}

// BatteryLevelChangedPayload holds the properties of a device battery level
// modification event.
type BatteryLevelChangedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level uint32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *BatteryLevelChangedPayload) Reset() {
	*x = BatteryLevelChangedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeular_events_v1_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatteryLevelChangedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatteryLevelChangedPayload) ProtoMessage() {}

func (x *BatteryLevelChangedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_timeular_events_v1_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatteryLevelChangedPayload.ProtoReflect.Descriptor instead.
func (*BatteryLevelChangedPayload) Descriptor() ([]byte, []int) {
	return file_timeular_events_v1_events_proto_rawDescGZIP(), []int{3}
}

func (x *BatteryLevelChangedPayload) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

// OrientationChangedPayload holds the properties of a device orientation
// modification.
type OrientationChangedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceId uint32 `protobuf:"varint,1,opt,name=face_id,json=faceId,proto3" json:"face_id,omitempty"`
}

func (x *OrientationChangedPayload) Reset() {
	*x = OrientationChangedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeular_events_v1_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrientationChangedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrientationChangedPayload) ProtoMessage() {}

func (x *OrientationChangedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_timeular_events_v1_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrientationChangedPayload.ProtoReflect.Descriptor instead.
func (*OrientationChangedPayload) Descriptor() ([]byte, []int) {
	return file_timeular_events_v1_events_proto_rawDescGZIP(), []int{4}
}

func (x *OrientationChangedPayload) GetFaceId() uint32 {
	if x != nil {
		return x.FaceId
	}
	return 0
}

var File_timeular_events_v1_events_proto protoreflect.FileDescriptor

var file_timeular_events_v1_events_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x8c, 0x04, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x57, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x0f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12,
	0x60, 0x0a, 0x13, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x12, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x12, 0x64, 0x0a, 0x15, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x48, 0x00, 0x52, 0x13, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x60, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x12, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x1b,
	0x0a, 0x19, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x32, 0x0a, 0x1a, 0x42,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22,
	0x34, 0x0a, 0x19, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x66,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x2a, 0xb6, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1f, 0x0a, 0x1b, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x22, 0x0a, 0x1e, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x45, 0x52, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52, 0x49, 0x45, 0x4e, 0x54, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x04, 0x42, 0xc1,
	0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x7a, 0x6e, 0x74, 0x72, 0x2e, 0x69,
	0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x45, 0x58,
	0xaa, 0x02, 0x12, 0x54, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x54, 0x69, 0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72,
	0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x54, 0x69, 0x6d,
	0x65, 0x75, 0x6c, 0x61, 0x72, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x54, 0x69,
	0x6d, 0x65, 0x75, 0x6c, 0x61, 0x72, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timeular_events_v1_events_proto_rawDescOnce sync.Once
	file_timeular_events_v1_events_proto_rawDescData = file_timeular_events_v1_events_proto_rawDesc
)

func file_timeular_events_v1_events_proto_rawDescGZIP() []byte {
	file_timeular_events_v1_events_proto_rawDescOnce.Do(func() {
		file_timeular_events_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_timeular_events_v1_events_proto_rawDescData)
	})
	return file_timeular_events_v1_events_proto_rawDescData
}

var file_timeular_events_v1_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_timeular_events_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_timeular_events_v1_events_proto_goTypes = []interface{}{
	(EventType)(0),                     // 0: timeular.events.v1.EventType
	(*Event)(nil),                      // 1: timeular.events.v1.Event
	(*DeviceConnectedPayload)(nil),     // 2: timeular.events.v1.DeviceConnectedPayload
	(*DeviceDisconnectedPayload)(nil),  // 3: timeular.events.v1.DeviceDisconnectedPayload
	(*BatteryLevelChangedPayload)(nil), // 4: timeular.events.v1.BatteryLevelChangedPayload
	(*OrientationChangedPayload)(nil),  // 5: timeular.events.v1.OrientationChangedPayload
}
var file_timeular_events_v1_events_proto_depIdxs = []int32{
	0, // 0: timeular.events.v1.Event.event_type:type_name -> timeular.events.v1.EventType
	2, // 1: timeular.events.v1.Event.device_connected:type_name -> timeular.events.v1.DeviceConnectedPayload
	3, // 2: timeular.events.v1.Event.device_disconnected:type_name -> timeular.events.v1.DeviceDisconnectedPayload
	4, // 3: timeular.events.v1.Event.battery_level_changed:type_name -> timeular.events.v1.BatteryLevelChangedPayload
	5, // 4: timeular.events.v1.Event.orientation_changed:type_name -> timeular.events.v1.OrientationChangedPayload
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_timeular_events_v1_events_proto_init() }
func file_timeular_events_v1_events_proto_init() {
	if File_timeular_events_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timeular_events_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_timeular_events_v1_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceConnectedPayload); i {
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
		file_timeular_events_v1_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDisconnectedPayload); i {
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
		file_timeular_events_v1_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatteryLevelChangedPayload); i {
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
		file_timeular_events_v1_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrientationChangedPayload); i {
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
	file_timeular_events_v1_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_DeviceConnected)(nil),
		(*Event_DeviceDisconnected)(nil),
		(*Event_BatteryLevelChanged)(nil),
		(*Event_OrientationChanged)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_timeular_events_v1_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timeular_events_v1_events_proto_goTypes,
		DependencyIndexes: file_timeular_events_v1_events_proto_depIdxs,
		EnumInfos:         file_timeular_events_v1_events_proto_enumTypes,
		MessageInfos:      file_timeular_events_v1_events_proto_msgTypes,
	}.Build()
	File_timeular_events_v1_events_proto = out.File
	file_timeular_events_v1_events_proto_rawDesc = nil
	file_timeular_events_v1_events_proto_goTypes = nil
	file_timeular_events_v1_events_proto_depIdxs = nil
}
