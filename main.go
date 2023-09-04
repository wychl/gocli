/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package main

import (
	"github.com/wychl/gocli/cmd"
	_ "github.com/wychl/gocli/cmd/flate"
	_ "github.com/wychl/gocli/cmd/qrcode"
	_ "github.com/wychl/gocli/cmd/url"
)

func main() {
	cmd.Execute()
}
