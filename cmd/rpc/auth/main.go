package main

import (
	"my-im-server/pkg/common/cmd"
	"my-im-server/sdk/tools/system/program"
)

func main() {
	if err := cmd.NewRpcAuthCmd().Exec(); err != nil {
		program.ExitWithError(err)
	}
}
