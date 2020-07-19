package main

import (
	"strings"
)

// DataPack ...
type DataPack struct{
	str string
} 

// Add ...
func (d DataPack) Add(inVariable string, inValue string) {
	var sb strings.Builder
	sb.WriteString(d.str)
	var val string = strings.Replace(inValue, "\r\n", "|EOL", -1)
	sb.WriteByte(13)
	sb.WriteByte(10)
	sb.WriteString("|0:")
	sb.WriteString(inVariable)
	sb.WriteString(":")
	sb.WriteString(val)
	d.str = sb.String()
}
