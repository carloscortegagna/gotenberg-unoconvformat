package main

import (
	gotenbergcmd "github.com/gotenberg/gotenberg/v7/cmd"

	// Gotenberg modules. You may also cherry-pick the standard modules.
	_ "github.com/carloscortegagna/gotenberg-doc/pkg/modules/unoconvdoc"
	_ "github.com/gotenberg/gotenberg/v7/pkg/standard"
)

func main() {
	gotenbergcmd.Run()
}
