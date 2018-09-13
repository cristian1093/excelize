/*
Package excelize providing a set of functions that allow you to write to
and read from XLSX files. Support reads and writes XLSX file generated by
Microsoft Excel™ 2007 and later. Support save file without losing original
charts of XLSX. This library needs Go version 1.8 or later.

Copyright 2016 - 2018 The excelize Authors. All rights reserved. Use of
this source code is governed by a BSD-style license that can be found in
the LICENSE file.
*/
package excelize

import "encoding/xml"

// xlsxTypes directly maps the types element of content types for relationship
// parts, it takes a Multipurpose Internet Mail Extension (MIME) media type as a
// value.
type xlsxTypes struct {
	XMLName   xml.Name       `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`
	Overrides []xlsxOverride `xml:"Override"`
	Defaults  []xlsxDefault  `xml:"Default"`
}

// xlsxOverride directly maps the override element in the namespace
// http://schemas.openxmlformats.org/package/2006/content-types
type xlsxOverride struct {
	PartName    string `xml:",attr"`
	ContentType string `xml:",attr"`
}

// xlsxDefault directly maps the default element in the namespace
// http://schemas.openxmlformats.org/package/2006/content-types
type xlsxDefault struct {
	Extension   string `xml:",attr"`
	ContentType string `xml:",attr"`
}
