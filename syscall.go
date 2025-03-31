//go:build windows

package winio

//go:generate go run github.com/observeinc/go-winio/tools/mkwinsyscall -output zsyscall_windows.go ./*.go
