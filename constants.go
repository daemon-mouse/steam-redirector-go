package main

import (
	"path/filepath"
)

var MO2PathFile = filepath.Join("modorganizer2", "instance_path.txt")

const (
	SchemeNXM         = "nxm://"
	SchemeMO2Shortcut = "moshortcut://"
	MO2ArgPick        = "--pick"
	EnvNoRedirect     = "STEAMREDIRECTOR_NO_REDIRECT"
)
