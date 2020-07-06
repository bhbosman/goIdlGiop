// Code generated by me. DO NOT EDIT.

package goIdlGiop

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "GIOP::MessageHeader_1_2", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type GiopMessageHeader12 struct {
	GiopMessageHeader11
}

//noinspection GoSnakeCaseUsage
const GiopMessageHeader12Id_Const = "IDL:GIOP/MessageHeader_1_2:1.0"

func (self *GiopMessageHeader12) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *GiopMessageHeader12) GoString() string {
	return self.String()
}

func (self *GiopMessageHeader12) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.GiopMessageHeader11.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *GiopMessageHeader12) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *GiopMessageHeader12) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.GiopMessageHeader11.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type GiopMessageHeader12_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var GiopMessageHeader12Helper = GiopMessageHeader12_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			GiopMessageHeader12Id_Const,
			__goidl__.StructType,
			"GIOP.idl",
			"xdl_GiopMessageHeader12.go",
			func() __goidl__.IIdlObject {
				return &GiopMessageHeader12{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &GiopMessageHeader12{}
			},
			__reflect__.TypeOf((*GiopMessageHeader12)(nil))))
}
