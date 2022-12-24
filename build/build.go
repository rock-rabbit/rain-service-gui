package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	os.MkdirAll("./out/darwin", os.ModePerm)
	os.MkdirAll("./out/windows", os.ModePerm)

	appname := "rain"

	// varsion := ""
	// f, err := os.Open("./FyneApp.toml")
	// if err == nil {
	// 	cfg, _ := io.ReadAll(f)
	// 	ds := regexp.MustCompile(`Version = "(.*?)"`).FindStringSubmatch(string(cfg))
	// 	if len(ds) > 1 {
	// 		varsion = ds[1]
	// 	}
	// }

	darwinName := fmt.Sprintf("./out/darwin/%s.app", appname)
	windowsName := fmt.Sprintf("./out/windows/%s.exe", appname)

	darwinService := filepath.Join(darwinName, "Contents/MacOS", "rain-service")
	windowsService := filepath.Join("./out/windows/", "rain-service.exe")

	if !FileNotExist("./rain.app") {
		if !FileNotExist(darwinName) {
			os.RemoveAll(darwinName)
		}
		os.Rename("./rain.app", darwinName)
	}

	if !FileNotExist("./rain.exe") {
		if !FileNotExist(windowsName) {
			os.RemoveAll(windowsName)
		}
		os.Rename("./rain.exe", windowsName)
	}

	if !FileNotExist("./build/rain-service") {
		if !FileNotExist(darwinService) {
			os.RemoveAll(darwinService)
		}
		Copy("./build/rain-service", darwinService)
	}

	if !FileNotExist("./build/rain-service.exe") {
		if !FileNotExist(windowsService) {
			os.RemoveAll(windowsService)
		}
		Copy("./build/rain-service.exe", windowsService)
	}

}

func Copy(f1, f2 string) {
	ff1, err := os.Open(f1)
	if err != nil {
		panic(err)
	}
	defer ff1.Close()
	ff2, err := os.OpenFile(f2, os.O_CREATE|os.O_RDWR, 0700)
	if err != nil {
		panic(err)
	}
	defer ff2.Close()
	_, err = io.Copy(ff2, ff1)
	if err != nil {
		panic(err)
	}
}

func FileNotExist(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsNotExist(err)
}
