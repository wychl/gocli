package asset

import _ "embed"

//go:embed country.mmdb
var GeoIP []byte

//go:embed city.free.ipdb
var QQWRY []byte
