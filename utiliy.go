package main

import (
	"image/color"
	"path/filepath"
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

func GetFilename(v string) string {
	_, t := filepath.Split(v)
	return t
}
