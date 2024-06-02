package web

import (
	"embed"
	"io/fs"
)

//go:embed all:public
var public embed.FS

func PublicFS() (fs.FS, error) {
	return fs.Sub(public, "public")
}
