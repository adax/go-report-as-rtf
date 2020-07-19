package main

func CreateReport(
	asDataFileName string,
	asPicFileName string,
	asRepFileName string,
) (err string) {
	var	rtfStrInp string
	var rtfStrOut string
	var szIO []byte
	var pos *byte
	var asDat string 
	var asParname string
	var asParvalue string
	var curlev int32
	var prevlev int32 = 0
	var kz int32 = 0
	var boolDOSdata bool
	var boolLev bool

	var asTmp string
	var ib, ib0, ib1, ib2, ilen, ilen2 int32
	var ch uint8

	if asPicFileName == asRepFileName {
		return "" 
	}
}