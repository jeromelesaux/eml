package encoding

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/google/uuid"

	"github.com/jeromelesaux/eml"
	"github.com/jeromelesaux/eml/common"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(v *eml.Eml) error {
	_, err := e.w.Write([]byte(common.XSender + " " + v.XSender + "\n"))
	if err != nil {
		return err
	}
	_, err = e.w.Write([]byte(common.XReceiver + " " + v.XReceiver + "\n"))
	if err != nil {
		return err
	}
	_, err = e.w.Write([]byte(common.MimeVersion + " " + v.MimeVersion + "\n"))
	if err != nil {
		return err
	}
	_, err = e.w.Write([]byte(common.From + " " + v.From + "\n"))
	if err != nil {
		return err
	}
	_, err = e.w.Write([]byte(common.To + " " + v.To + "\n"))
	if err != nil {
		return err
	}
	_, err = e.w.Write([]byte(common.Date + " " + v.Date + "\n"))
	if err != nil {
		return err
	}
	if v.ContentType.ContentType != "" {
		_, err = e.w.Write([]byte(common.ContentType + " "))
		if err != nil {
			return err
		}
		_, err = e.w.Write([]byte(v.ContentType.ContentType + ";type=\"text/html\";boundary=\"" + v.ContentType.Boundary + "\"\n\n\n"))
		if err != nil {
			return err
		}
		e.w.Write([]byte("--" + v.ContentType.Boundary + "\n"))
		e.w.Write([]byte(common.ContentTransferEncoding + " 7bit\n" + common.ContentType + "text/html;\n\tcharset=us-ascii\n\n<html><body></body></html>\n\n"))

		if len(v.Attachments) > 0 {
			for _, a := range v.Attachments {
				e.w.Write([]byte("--" + v.ContentType.Boundary + "\n"))

				_, err = e.w.Write([]byte(common.ContentType))
				if err != nil {
					return err
				}
				_, err = e.w.Write([]byte(" " + a.ContentType.ContentType + "; name=\"" + a.ContentType.Name + "\"\n"))
				if err != nil {
					return err
				}
				_, err = e.w.Write([]byte(common.ContentTransferEncoding + " " + a.ContentTransferEncoding + "\n"))
				if err != nil {
					return err
				}
				_, err = e.w.Write([]byte(common.ContentDisposition + " " + a.ContentDisposition.Value + "; filename=" + a.ContentDisposition.Filename + "\n"))
				if err != nil {
					return err
				}
				contentId := uuid.New()
				_, err = e.w.Write([]byte(common.ContentId + " <" + contentId.String() + ">\n"))
				if err != nil {
					return err
				}
				e.w.Write([]byte("\n"))
				//	e.w.Write([]byte(a.Content))
				nbCycle := int(len(a.Content) / 77)
				for i := 0; i < nbCycle; i++ {
					index := i * 77
					e.w.Write([]byte(a.Content[index:index+77] + "\n"))
				}
				e.w.Write([]byte(a.Content[(nbCycle * 77):len(a.Content)]))
				e.w.Write([]byte("\n"))
			}
			e.w.Write([]byte("--" + v.ContentType.Boundary + "--\n"))
		}

	}
	return nil
}

type Decoder struct {
	r       io.Reader
	scanner *bufio.Scanner
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) Decode(v *eml.Eml) error {
	d.scanner = bufio.NewScanner(d.r)
	d.scanner.Split(bufio.ScanLines)
	for d.scanner.Scan() {
		b := d.scanner.Text()
		if val := getSimpleValue(common.MimeVersion, b); val != "" {
			v.MimeVersion = val
			continue
		}
		if val := getSimpleValue(common.ReturnPath, b); val != "" {
			v.ReturnPath = val
			continue
		}
		if val := getSimpleValue(common.MessageId, b); val != "" {
			v.MessageId = val
			continue
		}
		if val := getSimpleValue(common.Date, b); val != "" {
			v.Date = val
			continue
		}
		if val := getSimpleValue(common.From, b); val != "" {
			v.From = val
			continue
		}
		if val := getSimpleValue(common.ReplyTo, b); val != "" {
			v.ReplyTo = val
			continue
		}
		if val := getSimpleValue(common.To, b); val != "" {
			v.To = val
			continue
		}
		if val := getSimpleValue(common.Subject, b); val != "" {
			v.Subject = val
			continue
		}
		if val := getSimpleValue(common.InReplyTo, b); val != "" {
			v.InReplyTo = val
			continue
		}
		if val := getSimpleValue(common.InReplyTo, b); val != "" {
			v.InReplyTo = val
			continue
		}
		if val := getSimpleValue(common.References, b); val != "" {
			v.References = val
			continue
		}
		if val := getSimpleValue(common.ContentType, b); val != "" {
			v.ContentType.ContentType = val
			if val := getGroupValue(common.Boundary, b, 1); val != "" {
				v.ContentType.Boundary = val
			} else {
				if d.scanner.Scan() {
					if val = getGroupValue(common.Boundary, d.scanner.Text(), 1); val != "" {
						v.ContentType.Boundary = val
					}
				}
			}
			continue
		}
		fmt.Fprintf(os.Stdout, "text:%s\n", b)
	}
	return nil
}

func getGroupValue(pattern, value string, group int) string {
	r := regexp.MustCompile(pattern)
	matches := r.FindStringSubmatch(value)
	if len(matches) > group {
		return matches[group]
	}
	return ""
}

func getSimpleValue(key, str string) string {
	r := regexp.MustCompile(key + "\\s" + common.SimpleValueRegex)
	matches := r.FindStringSubmatch(str)
	if len(matches) == 0 {
		return ""
	}
	if len(matches) >= 0 {
		return matches[1]
	}
	return ""
}

func getMultilineValue(key string, scan *bufio.Scanner) string {
	value := ""
	r := regexp.MustCompile(key + "\\s" + "([^\\n]*)")
	str := scan.Text()
	matches := r.FindStringSubmatch(str)
	if len(matches) == 0 {
		return ""
	}

	if len(matches) >= 0 {
		value += matches[1] + "\n"
		r2 := regexp.MustCompile("^\\t([^\\n]*)")
		for {
			scan.Scan()
			t := scan.Text()
			match := r2.FindStringSubmatch(t)
			if match != nil && len(match) >= 0 {
				value += match[1] + "\n"
			} else {
				return value
			}
		}
	}
	return value
}
