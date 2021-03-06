// Code generated by me. DO NOT EDIT.

package goIdlGiop

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "GIOP::MessageHeader_1_0", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type GiopMessageHeader10 struct {
	__goidl__.IdlObject
	Magic [4]byte `json:"Magic"`
	GiopVersion GiopVersion `json:"GiopVersion"`
	ByteOrder bool `json:"ByteOrder"`
	MessageType byte `json:"MessageType"`
	MessageSize uint32 `json:"MessageSize"`
}

//noinspection GoSnakeCaseUsage
const GiopMessageHeader10Id_Const = "IDL:omg.org/GIOP/MessageHeader_1_0:1.0"

func (self *GiopMessageHeader10) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *GiopMessageHeader10) GoString() string {
	return self.String()
}

func (self *GiopMessageHeader10) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(CharType)
	self.Magic, err = GiopMagicnHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.GiopVersion.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(BooleanType)
	self.ByteOrder, err = __goidl__.IdlBooleanHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlOctetKind)
	self.MessageType, err = __goidl__.IdlOctetHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.MessageSize, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *GiopMessageHeader10) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *GiopMessageHeader10) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(CharType)
	err = GiopMagicnHelper.Write(stream, self.Magic)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.GiopVersion.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(BooleanType)
	err = __goidl__.IdlBooleanHelper.Write(stream, self.ByteOrder)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlOctetKind)
	err = __goidl__.IdlOctetHelper.Write(stream, self.MessageType)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.MessageSize)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type GiopMessageHeader10_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var GiopMessageHeader10Helper = GiopMessageHeader10_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			GiopMessageHeader10Id_Const,
			__goidl__.StructType,
			"GIOP.idl",
			"xdl_GiopMessageHeader10.go",
			func() __goidl__.IIdlObject {
				return &GiopMessageHeader10{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &GiopMessageHeader10{}
			},
			__reflect__.TypeOf((*GiopMessageHeader10)(nil))))
}
