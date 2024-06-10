package text

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"git.sr.ht/~kvo/go-format/document"
	"git.sr.ht/~kvo/go-std/defs"
	"git.sr.ht/~kvo/go-std/errors"
)

type listConfig struct {
	Index *int
	Last bool
	Reversed bool
	Type string
}

func alphaConv(i uint) string {
	var s string
	symbols := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	for i -= 1; i/26 != 0; i -= 1 {
		s = symbols[i%26] + s
		i = i/26
	}
	s = symbols[i%26] + s
	return s
}

func toRoman(i uint) (string, error) {
	if i > 3999 {
		return strconv.Itoa(int(i)), errors.New(
			"integer is larger than largest Roman numeral", nil,
		)
	} else if i == 0 {
		return strconv.Itoa(int(i)), errors.New(
			"no Roman numeral for zero", nil,
		)
	}

	symbols := []struct {
		Value  uint
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, symbol := range symbols {
		for i >= symbol.Value {
			roman += symbol.Symbol
			i -= symbol.Value
		}
	}

	return roman, nil
}

// BUG: Spacing is oddly rendered.
func trimSpace(s string, last bool) (string, error) {
	// 1. Cull rogue carriage returns.
	s = strings.ReplaceAll(s, "\r", "")

	// 2. Remove all spaces and tabs immediately before and after a line break are ignored.

	re, err := regexp.Compile("( )+\n( )+")
	if err != nil {
		return "", errors.New(err.Error(), nil)
	}
	s = string(re.ReplaceAll([]byte(s), []byte("\n")))
	re, err = regexp.Compile("(\t)+\n(\t)+")
	if err != nil {
		return "", errors.New(err.Error(), nil)
	}
	s = string(re.ReplaceAll([]byte(s), []byte("\n")))

	// 3. Convert tabs to spaces (with tab stop of 4).

	tabstop := 4
	lines := strings.Split(s, "\n")
	s = ""
	for _, line := range lines {
		var detabbed string
		for _, c := range []rune(line) {
			if string(c) == "\t" {
				detabbed += " "
				for len(detabbed) % tabstop != 0 {
					detabbed += " "
				}
			} else {
				detabbed += string(c)
			}
		}
		s += detabbed + "\n"
	}

	// 4. Convert newlines to spaces.

	s = strings.ReplaceAll(s, "\n", " ")

	// 5. Remove duplicate whitespace.

	re, err = regexp.Compile("( )+")
	if err != nil {
		return "", errors.New(err.Error(), nil)
	}
	s = string(re.ReplaceAll([]byte(s), []byte(" ")))

	// 6. Remove all extra whitespace.

	if s == " " {
		return "", nil
	}

	/*if last {
		s = strings.TrimSuffix(s, " ")
	}*/

	return s, nil
}

func render(w io.Writer, n *document.Node, list listConfig) error {
	listIndex := 1

	newlist := listConfig{
		Index: &listIndex,
		Reversed: false,
		Type: "",
	}

	forbidden := []string{
		"area", "basefont", "canvas", "content", "dialog", "font", "head",
		"link", "map", "meta", "param", "script", "shadow", "slot", "style",
		"template", "title",
	}

	if n.Type != document.TextNode && defs.Has(forbidden, n.Data) {
		return nil
	}

	if n.Type == document.ElementNode {
		switch n.Data {
		case "abbr", "acronym", "dfn":
			var title string
			for _, a := range n.Attr {
				if a.Key == "title" {
					title = a.Val
					break
				}
			}
			fmt.Fprintf(w, "%s (", title)
		case "applet":
			var code string
			for _, a := range n.Attr {
				if a.Key == "code" {
					code = a.Val
					break
				}
			}
			fmt.Fprintf(w, "%s", code)
		case "audio", "bgsound", "embed", "frame", "iframe", "image", "img", "portal", "source", "video":
			var src string
			for _, a := range n.Attr {
				if a.Key == "src" {
					src = a.Val
					break
				}
			}
			fmt.Fprintf(w, "\n%s\n", src)
		case "button":
			fmt.Fprintf(w, "[ ")
		case "br":
			fmt.Fprintf(w, "\n")
		case "dd":
			fmt.Fprintf(w, "\t")
		case "h1", "h2", "h3", "h4", "h5", "h6":
			fmt.Fprintf(w, "\n")
		case "input":
			var inptype string
			for _, a := range n.Attr {
				if a.Key == "type" {
					inptype = a.Val
					break
				}
			}
			fmt.Fprintf(w, "\n[%s input field]\n", inptype)
		case "keygen":
			var keytype, challenge, keyparams string
			for _, a := range n.Attr {
				if a.Key == "keytype" {
					keytype = a.Val
				} else if a.Key == "challenge" {
					challenge = a.Val
				} else if a.Key == "keyparams" {
					keyparams = a.Val
				}
				if keytype != "" && challenge != "" && keyparams != "" {
					break
				}
			}
			fmt.Fprintf(
				w, `\n[ %s keygen | challenge: "%s" | key parameters: "%s" ]\n`,
				keytype, challenge, keyparams,
			)
		case "li":
			fmt.Fprintf(w, "\t")
			if n.Parent.Data == "menu" {
				fmt.Fprintf(w, "• ")
			} else if n.Parent.Data == "ul" {
				switch list.Type {
				case "circle":
					fmt.Fprintf(w, "◦ ")
				case "square":
					fmt.Fprintf(w, "▪ ")
				case "triangle":
					fmt.Fprintf(w, "▴ ")
				default:
					fmt.Fprintf(w, "• ")
				}
			} else if n.Parent.Data == "ol" {
				switch list.Type {
				case "a":
					fmt.Fprintf(w, "%s. ", alphaConv(uint(*list.Index)))
				case "A":
					fmt.Fprintf(w, "%s. ", strings.ToUpper(alphaConv(uint(*list.Index))))
				case "i":
					prefix, err := toRoman(uint(*list.Index))
					if err != nil {
						return errors.New("failed generating Roman numeral", err)
					}
					fmt.Fprintf(w, "%s. ", strings.ToLower(prefix))
				case "I":
					prefix, err := toRoman(uint(*list.Index))
					if err != nil {
						return errors.New("failed generating Roman numeral", err)
					}
					fmt.Fprintf(w, "%s. ", prefix)
				default:
					fmt.Fprintf(w, "%d. ", *list.Index)
				}
				if list.Reversed {
					*list.Index--
				} else {
					*list.Index++
				}
			}
		case "math":
			fmt.Fprintf(w, "\n[Mathematical formula]\n")
		case "ol":
			var reversed, start string
			for _, a := range n.Attr {
				if a.Key == "reversed" {
					reversed = a.Val
				} else if a.Key == "start" {
					start = a.Val
				} else if a.Key == "type" {
					newlist.Type = a.Val
				}
				if reversed != "" && start != "" && newlist.Type != "" {
					break
				}
			}
			startInt, err := strconv.Atoi(start)
			if start != "" && err == nil {
				*newlist.Index = startInt
			}
			if reversed == "true" {
				newlist.Reversed = true
			}
		case "option":
			fmt.Fprintf(w, "\n\t[ ] ")
		case "object":
			var data string
			for _, a := range n.Attr {
				if a.Key == "data" {
					data = a.Val
					break
				}
			}
			fmt.Fprintf(w, "\n%s\n", data)
		case "picture":
			fmt.Fprintf(w, "[Multi-layer image:]\n")
		case "pre", "textarea", "xmp":
			fmt.Fprintf(w, "\n————————————————————\n")
		case "progress":
			var value, max string
			for _, a := range n.Attr {
				if a.Key == "value" {
					value = a.Val
				} else if a.Key == "max" {
					max = a.Val
				}
				if value != "" && max != "" {
					break
				}
			}
			valfloat, err := strconv.ParseFloat(value, 64)
			if err != nil {
				valfloat = 0.0
			}
			maxfloat, err := strconv.ParseFloat(max, 64)
			if err != nil {
				maxfloat = 0.0
			}
			fmt.Fprintf(w, "%.f%% complete", valfloat/maxfloat*100)
		case "svg":
			fmt.Fprintf(w, "\n[SVG vector graphic]\n")
		case "sup":
			fmt.Fprintf(w, "^")
		case "track":
			var src string
			for _, a := range n.Attr {
				if a.Key == "src" {
					src = a.Val
					break
				}
			}
			fmt.Fprintf(w, "\n\t[Subtitles:] %s\n", src)
		}
	}

	if newlist.Reversed {
		*newlist.Index--
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == document.ElementNode && c.Data == "li" {
				*newlist.Index++
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		newlist.Last = false
		if c.NextSibling == nil {
			newlist.Last = true
		}

		if n.Type == document.ElementNode && n.Data == "ul" {
			newlist.Index = nil
			newlist.Reversed = false
		} else if n.Type != document.ElementNode && n.Data != "ol" {
			newlist.Index = nil
			newlist.Reversed = false
			newlist.Type = ""
		}

		err := render(w, c, newlist)
		if err != nil {
			return errors.New("error rendering node", err)
		}
	}

	if n.Type == document.TextNode {
		switch n.Parent.Data {
		case "applet", "math", "progress", "svg":
			return nil
		case "del", "s", "strike":
			for _, c := range []rune(n.Data) {
				fmt.Fprintf(w, string(c) + string(0x336))
			}
		case "pre", "xmp":
			fmt.Fprintf(w, n.Data)
		default:
			text, err := trimSpace(strings.TrimSpace(n.Data), list.Last)
			if err != nil {
				return errors.New("error trimming space from HTML text node", err)
			}
			fmt.Fprintf(w, text)
		}
	} else if n.Type == document.ElementNode {
		switch n.Data {
		case "a":
			var href string
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = a.Val
					break
				}
			}
			fmt.Fprintf(w, " (%s)", href)
		case "abbr", "acronym", "dfn":
			fmt.Fprintf(w, ")")
		case "button":
			fmt.Fprintf(w, " ]\n")
		case "caption", "dd", "dt", "figcaption", "label", "legend", "li", "optgroup", "summary", "option", "p":
			fmt.Fprintf(w, "\n")
		case "h1", "h2", "h3", "h4", "h5", "h6":
			fmt.Fprintf(w, "\n\n")
		case "pre", "textarea", "xmp":
			fmt.Fprintf(w, "\n————————————————————\n")
		}
	}
	return nil
}

// Render renders the document parse tree n to the given writer. Raises error if
// an error occurs.
func Render(w io.Writer, n *document.Node) error {
	newlist := listConfig{
		Index: nil,
		Last: false,
		Reversed: false,
		Type: "",
	}
	for n.NextSibling != nil {
		err := render(w, n, newlist)
		if err != nil {
			return errors.New("render HTML as text", err)
		}
		n = n.NextSibling
	}
	newlist.Last = true
	return render(w, n, newlist)
}
