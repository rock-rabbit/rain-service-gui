package main

import (
	"fmt"
	"image/color"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func HexToNRGBA(hex string) color.NRGBA {
	c := color.NRGBA{255, 255, 255, 255}

	lens := []int{4, 7, 9}
	if !IsInt(lens, len(hex)) {
		return c
	}

	// 将所有 hex 的出现方式进行处理
	// #fff #ffffff #ffffff
	var (
		strR, strG, strB, strA string
	)
	switch len(hex) {
	case 4:
		strR = strings.Repeat(hex[1:2], 2)
		strG = strings.Repeat(hex[2:3], 2)
		strB = strings.Repeat(hex[3:], 2)
		strA = "ff"
	case 7:
		strR = hex[1:3]
		strG = hex[3:5]
		strB = hex[5:]
		strA = "ff"
	case 9:
		strR = hex[1:3]
		strG = hex[3:5]
		strB = hex[5:7]
		strA = hex[7:]
	}

	// 解析 16 进制
	r, _ := strconv.ParseInt(strR, 16, 16)
	g, _ := strconv.ParseInt(strG, 16, 16)
	b, _ := strconv.ParseInt(strB, 16, 16)
	a, _ := strconv.ParseInt(strA, 16, 16)
	c.R = uint8(r)
	c.G = uint8(g)
	c.B = uint8(b)
	c.A = uint8(a)

	return c
}

// IsInt 查询 int 值是否在切片内
func IsInt(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

// GetFilename 获取文件名
func GetFilename(v string) string {
	_, t := filepath.Split(v)
	if len(t) > 26 {
		t = fmt.Sprintf("%s...%s", t[:16], t[len(t)-10:])
	}
	return t
}

// GetDownloadsDir 获取系统默认下载地址
func GetDownloadsDir() string {
	if runtime.GOOS == "windows" {
		return path.Join(os.Getenv("HOMEPATH"), "Downloads")
	}
	return path.Join(os.Getenv("HOME"), "Downloads")
}

// GetLogsDir
func GetLogsDir() string {
	if runtime.GOOS == "windows" {
		return GetExecutable()
	}
	return GetExecutable()
}

// GetExecutable 获取执行文件目录
func GetExecutable() string {
	p, err := os.Executable()
	if err != nil {
		p, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		return p
	}
	return filepath.Dir(p)
}

func FileNotExist(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsNotExist(err)
}
