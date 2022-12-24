Mingw32Version=10.0.0_3
CXX=/opt/homebrew/Cellar/mingw-w64/${Mingw32Version}/bin/x86_64-w64-mingw32-g++
CC=/opt/homebrew/Cellar/mingw-w64/${Mingw32Version}/bin/x86_64-w64-mingw32-gcc


build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 fyne package -os darwin -icon icons/darwin/Icon-MacOS-512x512@2x.png
	go run ./build/build.go

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 CC=${CC} CXX=${CXX} fyne package -os windows -icon icons/windows/Icon-Windows-512x512@2x.png
	go run ./build/build.go

build-all:build-mac build-windows