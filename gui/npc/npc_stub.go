//go:build !cgo
// +build !cgo

package main

import "log"

func main() {
	log.Fatal("npc 图形界面仅支持在启用 CGO 且具备图形依赖的环境下构建")
}
