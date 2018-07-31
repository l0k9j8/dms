package main

import (
	"github.com/BurntSushi/toml"
	"os/user"
	"path/filepath"
	"log"
	"bytes"
)

type serverConfig struct {
	Path                string `toml:"path"`
	IfName              string `toml:"interface"`
	Http                string `toml:"addr"`
	FriendlyName        string `toml:"name"`
	LogHeaders          bool   `toml:"log_header"`
	FFprobeCachePath    string `toml:"ffprobe_cache_path"`
	NoProbe             bool   `toml:"no_probe"`
	StallEventSubscribe bool   `toml:"stall_event_subscribe"`
	NotifyInterval      uint   `toml:"notify_interval"`
	IgnoreHidden        bool   `toml:"ignore_hidden"`
	IgnoreUnreadable    bool   `toml:"ignore_unreadable"`
}

func getDefaultFFprobeCachePath() (path string) {
	_user, err := user.Current()
	if err != nil {
		log.Print(err)
		return
	}
	path = filepath.Join(_user.HomeDir, ".dlna-ffprobe-cache")
	return
}

func newServerConfig() *serverConfig {
	return &serverConfig{
					Path: "",
					IfName: "",
					Http: ":8090",
					FriendlyName: "DLNA video server",
					LogHeaders: false,
					FFprobeCachePath: getDefaultFFprobeCachePath(),
					NoProbe: false,
					StallEventSubscribe: false,
					NotifyInterval: 30,
					IgnoreHidden: true,
					IgnoreUnreadable: true,
		}
}

func (conf *serverConfig) load(path string) error {
	if _, err := toml.DecodeFile(path, conf); err != nil {
		log.Printf("load config error %s", err)
		return err
	}
	return nil
}

func (conf *serverConfig) printDefault() string {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	encoder.Indent = ""
	if err := encoder.Encode(conf); err != nil {
		log.Printf("encoding default value error: %s", err)
		return ""
	}
	return buf.String()
}