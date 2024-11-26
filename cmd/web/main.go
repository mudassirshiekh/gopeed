//go:build web
// +build web

package main

import (
	"embed"
	"github.com/GopeedLab/gopeed/cmd"
	"github.com/GopeedLab/gopeed/pkg/api/model"
	"github.com/GopeedLab/gopeed/pkg/base"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed dist/*
var dist embed.FS

func main() {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}

	args := parse()
	var webBasicAuth *model.WebBasicAuth
	if isNotBlank(args.Username) && isNotBlank(args.Password) {
		webBasicAuth = &model.WebBasicAuth{
			Username: *args.Username,
			Password: *args.Password,
		}
	}

	var dir string
	if args.StorageDir != nil && *args.StorageDir != "" {
		dir = *args.StorageDir
	} else {
		exe, err := os.Executable()
		if err != nil {
			panic(err)
		}
		dir = filepath.Dir(exe)
	}

	if args.DownloadConfig == nil {
		args.DownloadConfig = &base.DownloaderStoreConfig{}
	}
	args.DownloadConfig.Http = &base.DownloaderHttpConfig{
		Host:     *args.Host,
		Port:     *args.Port,
		ApiToken: *args.ApiToken,
	}

	cfg := &model.StartConfig{
		Storage:        model.StorageBolt,
		StorageDir:     filepath.Join(dir, "storage"),
		DownloadConfig: args.DownloadConfig,
		ProductionMode: true,
		WebEnable:      true,
		WebFS:          sub,
		WebBasicAuth:   webBasicAuth,
	}
	cmd.Start(cfg)
}

func isNotBlank(str *string) bool {
	return str != nil && *str != ""
}
