package markovic

import (
	"io"
	"io/ioutil"

	"github.com/subosito/markovic/cmark"
)

func HTML(r io.Reader) string {
	return render(r, "html")
}

func XML(r io.Reader) string {
	return render(r, "xml")
}

func Man(r io.Reader, width int) string {
	return renderWidth(r, "man", width)
}

func CommonMark(r io.Reader, width int) string {
	return renderWidth(r, "commonmark", width)
}

func Latex(r io.Reader, width int) string {
	return renderWidth(r, "latex", width)
}

func render(r io.Reader, format string) string {
	return renderWidth(r, format, 0)
}

func renderWidth(r io.Reader, format string, width int) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return ""
	}

	option := cmark.OPTION_DEFAULT

	switch format {
	case "html":
		return cmark.HTML(string(b), option)
	case "xml":
		return cmark.XML(string(b), option)
	case "man":
		return cmark.Man(string(b), option, width)
	case "commonmark":
		return cmark.CommonMark(string(b), option, width)
	case "latex":
		return cmark.Latex(string(b), option, width)
	}

	return ""
}
