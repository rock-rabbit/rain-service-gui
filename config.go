package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	RainService string
	RefreshTime int
	Outdir      string
}

var config = &Config{
	RainService: "http://127.0.0.1:7362/json",
	RefreshTime: 500,
	Outdir:      GetDownloadsDir(),
}

var configPath = filepath.Join(GetExecutable(), "config.json")

func init() {
	_, err := os.Stat(configPath)
	if err != nil && os.IsNotExist(err) {
		return
	}
	// 读取配置
	f, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer f.Close()
	cfgtxt, err := io.ReadAll(f)
	if err != nil {
		return
	}
	var cfg Config
	err = json.Unmarshal(cfgtxt, &cfg)
	if err != nil {
		return
	}
	config = &cfg
}

func SaveConfig(outdir, refresh, service string) error {
	// 解析
	ref, err := strconv.Atoi(refresh)
	if err != nil {
		return err
	}
	config.Outdir = outdir
	config.RainService = service
	config.RefreshTime = ref
	Rain.Host = service

	// 删除
	_, err = os.Stat(configPath)
	if !os.IsNotExist(err) {
		os.Remove(configPath)
	}
	// 写入
	f, err := os.OpenFile(configPath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	d, err := json.Marshal(&config)
	if err != nil {
		return err
	}
	f.Write(d)
	return nil
}
