package eml

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Eml struct {
	ReturnPath  string
	From        string
	To          string
	Received    []string
	MimeVersion string
	Subject     string
	MessageId   string
	ContentType ContentType
	Date        string
	Body        []byte
	ReplyTo     string
	InReplyTo   string
	References  string
	XSender     string
	XReceiver   string
	Attachments []Attachment
	Message     *Eml
}

type Attachment struct {
	ContentType             ContentType
	ContentTransferEncoding string
	ContentDisposition      ContentDisposition
	Content                 string
}

type ContentDisposition struct {
	Value    string
	Filename string
}

type ContentType struct {
	ContentType string
	Boundary    string
	Name        string
}

func NewEml() *Eml {
	return &Eml{MimeVersion: "1.0"}
}

func (e *Eml) AddAttachment(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	bf, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	name := filepath.Base(filePath)
	a := Attachment{
		ContentTransferEncoding: "base64",
		ContentType: ContentType{
			ContentType: "application/octet-stream",
			Name:        name,
		},
		ContentDisposition: ContentDisposition{
			Value:    "attachment",
			Filename: name,
		},
	}
	a.Content = base64.StdEncoding.EncodeToString(bf)
	e.Attachments = append(e.Attachments, a)
	uuid := uuid.New()

	e.ContentType = ContentType{
		Boundary:    uuid.String(),
		ContentType: "multipart/related",
	}
	return nil
}
