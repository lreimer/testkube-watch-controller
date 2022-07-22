package main

import "github.com/lreimer/testkube-watch-controller/cmd"

var version = ""
var commit = ""

func main() {
	cmd.SetVersion(version)
	cmd.SetCommit(commit)
	cmd.Execute()
}
