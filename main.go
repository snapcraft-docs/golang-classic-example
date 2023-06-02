package main

import "fmt"

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -Wl,-rpath=\$ORIGIN/lib:/snap/core22/current/lib/x86_64-linux-gnu -Wl,--disable-new-dtags -Wl,-dynamic-linker=/snap/core22/current/lib64/ld-linux-x86-64.so.2 -lzstd
#include <zstd.h>
*/
import "C"

func zstdVersion() int {
	return int(C.ZSTD_versionNumber())
}

func main() {
	fmt.Println("libzstd version is", zstdVersion())
}
