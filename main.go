package main

import "fmt"

/*
#cgo pkg-config: libzstd
#include <zstd.h>
*/
import "C"

func zstdVersion() int {
    return int(C.ZSTD_versionNumber())
}

func main() {
	fmt.Println("libzstd version is", zstdVersion())
}
