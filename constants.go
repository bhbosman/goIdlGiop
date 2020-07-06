package goIdlGiop

import (
	__goidl__ "github.com/bhbosman/goidl"
)

var giopVersion10 = &GiopVersion{
	IdlObject: __goidl__.IdlObject{},
	Major:     1,
	Minor:     0,
}

var giopVersion11 = &GiopVersion{
	IdlObject: __goidl__.IdlObject{},
	Major:     1,
	Minor:     1,
}
var giopVersion12 = &GiopVersion{
	IdlObject: __goidl__.IdlObject{},
	Major:     1,
	Minor:     2,
}

func GiopVersion10() *GiopVersion {
	return giopVersion10
}
func GiopVersion11() *GiopVersion {
	return giopVersion11
}
func GiopVersion12() *GiopVersion {
	return giopVersion12
}
