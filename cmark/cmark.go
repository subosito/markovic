package cmark

/*
#include <stdio.h>
#include <stdlib.h>
#include "cmark.h"
*/
import "C"
import "unsafe"

type Format int

const (
	FORMAT_NONE Format = iota
	FORMAT_HTML
	FORMAT_XML
	FORMAT_MAN
	FORMAT_COMMONMARK
	FORMAT_LATEX
)

type Option int

const (
	OPTION_DEFAULT   Option = iota
	OPTION_SOURCEPOS        = 1 << iota
	OPTION_HARDBREAKS
	OPTION_SAFE
	OPTION_NOBREAKS
	OPTION_NORMALIZE = 1 << iota * 8
	OPTION_VALIDATE_UTF8
	OPTION_SMART
)

func HTML(s string, options Option) string {
	return render(s, FORMAT_HTML, options, 0)
}

func Version() string {
	return C.GoString(C.cmark_version_string())
}

func render(s string, format Format, options Option, width int) string {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))

	clen := C.size_t(len(s))
	node := C.cmark_parse_document(cstr, clen, C.int(options))
	defer C.cmark_node_free(node)

	var doc *C.char
	defer C.free(unsafe.Pointer(doc))

	switch format {
	case FORMAT_HTML:
		doc = html(node, options)
	case FORMAT_XML:
		doc = xml(node, options)
	case FORMAT_MAN:
		doc = man(node, options, width)
	case FORMAT_COMMONMARK:
		doc = commonmark(node, options, width)
	case FORMAT_LATEX:
		doc = latex(node, options, width)
	default:
		return ""
	}

	return C.GoString(doc)
}

func html(node *C.struct_cmark_node, options Option) *C.char {
	return C.cmark_render_html(node, C.int(options))
}

func xml(node *C.struct_cmark_node, options Option) *C.char {
	return C.cmark_render_xml(node, C.int(options))
}

func man(node *C.struct_cmark_node, options Option, width int) *C.char {
	return C.cmark_render_man(node, C.int(options), C.int(width))
}

func commonmark(node *C.struct_cmark_node, options Option, width int) *C.char {
	return C.cmark_render_commonmark(node, C.int(options), C.int(width))
}

func latex(node *C.struct_cmark_node, options Option, width int) *C.char {
	return C.cmark_render_latex(node, C.int(options), C.int(width))
}
