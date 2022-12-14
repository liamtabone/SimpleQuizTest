// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: simpleQuiz.proto

package proto

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

type Choice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChoiceId string `protobuf:"bytes,1,opt,name=ChoiceId,proto3" json:"ChoiceId,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *Choice) Reset() {
	*x = Choice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleQuiz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Choice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Choice) ProtoMessage() {}

func (x *Choice) ProtoReflect() protoreflect.Message {
	mi := &file_simpleQuiz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Choice.ProtoReflect.Descriptor instead.
func (*Choice) Descriptor() ([]byte, []int) {
	return file_simpleQuiz_proto_rawDescGZIP(), []int{0}
}

func (x *Choice) GetChoiceId() string {
	if x != nil {
		return x.ChoiceId
	}
	return ""
}

func (x *Choice) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type QuestionGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionGroupId int32       `protobuf:"varint,1,opt,name=QuestionGroupId,proto3" json:"QuestionGroupId,omitempty"`
	Question        []*Question `protobuf:"bytes,2,rep,name=Question,proto3" json:"Question,omitempty"`
}

func (x *QuestionGroup) Reset() {
	*x = QuestionGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleQuiz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionGroup) ProtoMessage() {}

func (x *QuestionGroup) ProtoReflect() protoreflect.Message {
	mi := &file_simpleQuiz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionGroup.ProtoReflect.Descriptor instead.
func (*QuestionGroup) Descriptor() ([]byte, []int) {
	return file_simpleQuiz_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionGroup) GetQuestionGroupId() int32 {
	if x != nil {
		return x.QuestionGroupId
	}
	return 0
}

func (x *QuestionGroup) GetQuestion() []*Question {
	if x != nil {
		return x.Question
	}
	return nil
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId int32     `protobuf:"varint,1,opt,name=QuestionId,proto3" json:"QuestionId,omitempty"`
	Text       string    `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	Choices    []*Choice `protobuf:"bytes,3,rep,name=Choices,proto3" json:"Choices,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleQuiz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_simpleQuiz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_simpleQuiz_proto_rawDescGZIP(), []int{2}
}

func (x *Question) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *Question) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Question) GetChoices() []*Choice {
	if x != nil {
		return x.Choices
	}
	return nil
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	QuestionId int32  `protobuf:"varint,2,opt,name=QuestionId,proto3" json:"QuestionId,omitempty"`
	AnswerId   string `protobuf:"bytes,3,opt,name=AnswerId,proto3" json:"AnswerId,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleQuiz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_simpleQuiz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_simpleQuiz_proto_rawDescGZIP(), []int{3}
}

func (x *Answer) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Answer) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *Answer) GetAnswerId() string {
	if x != nil {
		return x.AnswerId
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleQuiz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_simpleQuiz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_simpleQuiz_proto_rawDescGZIP(), []int{4}
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Percentile float32 `protobuf:"fixed32,1,opt,name=Percentile,proto3" json:"Percentile,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleQuiz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_simpleQuiz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_simpleQuiz_proto_rawDescGZIP(), []int{5}
}

func (x *Result) GetPercentile() float32 {
	if x != nil {
		return x.Percentile
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleQuiz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_simpleQuiz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_simpleQuiz_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_simpleQuiz_proto protoreflect.FileDescriptor

var file_simpleQuiz_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x38, 0x0a, 0x06, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x22, 0x65, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x08, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x60, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x22, 0x22, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x11, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x24, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x1a, 0x0b,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x27, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0a, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simpleQuiz_proto_rawDescOnce sync.Once
	file_simpleQuiz_proto_rawDescData = file_simpleQuiz_proto_rawDesc
)

func file_simpleQuiz_proto_rawDescGZIP() []byte {
	file_simpleQuiz_proto_rawDescOnce.Do(func() {
		file_simpleQuiz_proto_rawDescData = protoimpl.X.CompressGZIP(file_simpleQuiz_proto_rawDescData)
	})
	return file_simpleQuiz_proto_rawDescData
}

var file_simpleQuiz_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_simpleQuiz_proto_goTypes = []interface{}{
	(*Choice)(nil),        // 0: main.Choice
	(*QuestionGroup)(nil), // 1: main.QuestionGroup
	(*Question)(nil),      // 2: main.Question
	(*Answer)(nil),        // 3: main.Answer
	(*Empty)(nil),         // 4: main.Empty
	(*Result)(nil),        // 5: main.Result
	(*User)(nil),          // 6: main.User
}
var file_simpleQuiz_proto_depIdxs = []int32{
	2, // 0: main.QuestionGroup.Question:type_name -> main.Question
	0, // 1: main.Question.Choices:type_name -> main.Choice
	6, // 2: main.SimpleQuizService.newUser:input_type -> main.User
	6, // 3: main.SimpleQuizService.GetQuestion:input_type -> main.User
	3, // 4: main.SimpleQuizService.GetAnswer:input_type -> main.Answer
	6, // 5: main.SimpleQuizService.GetResult:input_type -> main.User
	4, // 6: main.SimpleQuizService.newUser:output_type -> main.Empty
	2, // 7: main.SimpleQuizService.GetQuestion:output_type -> main.Question
	4, // 8: main.SimpleQuizService.GetAnswer:output_type -> main.Empty
	5, // 9: main.SimpleQuizService.GetResult:output_type -> main.Result
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_simpleQuiz_proto_init() }
func file_simpleQuiz_proto_init() {
	if File_simpleQuiz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simpleQuiz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Choice); i {
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
		file_simpleQuiz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionGroup); i {
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
		file_simpleQuiz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_simpleQuiz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_simpleQuiz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_simpleQuiz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_simpleQuiz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_simpleQuiz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simpleQuiz_proto_goTypes,
		DependencyIndexes: file_simpleQuiz_proto_depIdxs,
		MessageInfos:      file_simpleQuiz_proto_msgTypes,
	}.Build()
	File_simpleQuiz_proto = out.File
	file_simpleQuiz_proto_rawDesc = nil
	file_simpleQuiz_proto_goTypes = nil
	file_simpleQuiz_proto_depIdxs = nil
}
