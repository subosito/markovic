package cmark

/*
#include <stdio.h>
#include <stdlib.h>
#include "cmark.h"
*/
import "C"
import "unsafe"

func MarkdownHTML(s string, flags int) string {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	cLen := C.size_t(len(s))
	cOut := C.cmark_markdown_to_html(cStr, cLen, C.int(flags))
	defer C.free(unsafe.Pointer(cOut))

	return C.GoString(cOut)
}

func Version() string {
	return C.GoString(C.cmark_version_string())
}
