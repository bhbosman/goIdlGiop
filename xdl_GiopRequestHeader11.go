// Code generated by me. DO NOT EDIT.

package goIdlGiop

import __goIdlIop__ "github.com/bhbosman/goIdlIop"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "GIOP::RequestHeader_1_1", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type GiopRequestHeader11 struct {
	__goidl__.IdlObject
	ServiceContext __goIdlIop__.IopServiceContextList `json:"ServiceContext"`
	RequestId uint32 `json:"RequestId"`
	ResponseExpected bool `json:"ResponseExpected"`
	Reserved [3]byte `json:"Reserved"`
	ObjectKey __goIdlIop__.IopObjectKey `json:"ObjectKey"`
	Operation string `json:"Operation"`
	RequestingPrincipal GiopPrincipal `json:"RequestingPrincipal"`
}

//noinspection GoSnakeCaseUsage
const GiopRequestHeader11Id_Const = "IDL:omg.org/GIOP/RequestHeader_1_1:1.0"

func (self *GiopRequestHeader11) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *GiopRequestHeader11) GoString() string {
	return self.String()
}

func (self *GiopRequestHeader11) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.ServiceContext.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.RequestId, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(BooleanType)
	self.ResponseExpected, err = __goidl__.IdlBooleanHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlOctetKind)
	self.Reserved, err = GiopRequestReservedHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.ObjectKey.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Operation, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.RequestingPrincipal.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *GiopRequestHeader11) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *GiopRequestHeader11) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.ServiceContext.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.RequestId)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(BooleanType)
	err = __goidl__.IdlBooleanHelper.Write(stream, self.ResponseExpected)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlOctetKind)
	err = GiopRequestReservedHelper.Write(stream, self.Reserved)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.ObjectKey.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Operation)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.RequestingPrincipal.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type GiopRequestHeader11_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var GiopRequestHeader11Helper = GiopRequestHeader11_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			GiopRequestHeader11Id_Const,
			__goidl__.StructType,
			"GIOP.idl",
			"xdl_GiopRequestHeader11.go",
			func() __goidl__.IIdlObject {
				return &GiopRequestHeader11{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &GiopRequestHeader11{}
			},
			__reflect__.TypeOf((*GiopRequestHeader11)(nil))))
}