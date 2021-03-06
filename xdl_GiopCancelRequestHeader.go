// Code generated by me. DO NOT EDIT.

package goIdlGiop

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "GIOP::CancelRequestHeader", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type GiopCancelRequestHeader struct {
	__goidl__.IdlObject
	RequestId uint32 `json:"RequestId"`
}

//noinspection GoSnakeCaseUsage
const GiopCancelRequestHeaderId_Const = "IDL:omg.org/GIOP/CancelRequestHeader:1.0"

func (self *GiopCancelRequestHeader) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *GiopCancelRequestHeader) GoString() string {
	return self.String()
}

func (self *GiopCancelRequestHeader) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.RequestId, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *GiopCancelRequestHeader) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *GiopCancelRequestHeader) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.RequestId)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type GiopCancelRequestHeader_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var GiopCancelRequestHeaderHelper = GiopCancelRequestHeader_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			GiopCancelRequestHeaderId_Const,
			__goidl__.StructType,
			"GIOP.idl",
			"xdl_GiopCancelRequestHeader.go",
			func() __goidl__.IIdlObject {
				return &GiopCancelRequestHeader{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &GiopCancelRequestHeader{}
			},
			__reflect__.TypeOf((*GiopCancelRequestHeader)(nil))))
}
