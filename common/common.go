package common

import (
	"regexp"
	"strings"
)

var (
	MimeVersion             = "Mime-Version:"
	ContentType             = "Content-Type:" // multiline
	MessageId               = "Message-Id:"
	From                    = "From:"
	Subject                 = "Subject:"
	Date                    = "Date:"
	To                      = "To:"
	ContentTransferEncoding = "Content-Transfer-Encoding:"
	ReturnPath              = "Return-Path:"
	XOriginalTo             = "X-Original-To:"
	DeliveredTo             = "Delivered-To:"
	DomainKeySignature      = "DomainKey-Signature:"
	ContentDisposition      = "Content-Disposition:"
	References              = "References:"
	XMailer                 = "X-Mailer:"
	XSpamCheckerVersion     = "X-Spam-Checker-Version:"
	XSpamLevel              = "X-Spam-Level:"
	XSpamStatus             = "X-Spam-Status:"
	XMimeOle                = "X-MimeOLE:"
	ThreadIndex             = "Thread-Index:"
	Received                = "Received:" //multiline

	SimpleValueRegex = "([^\n]*)"
)

func getSimpleValue(key, str string) string {
	r := regexp.MustCompile(key + "\\s" + SimpleValueRegex)
	matches := r.FindStringSubmatch(str)
	if len(matches) >= 0 {
		return matches[1]
	}
	return ""
}

func getMutlineValue(key, str string) string {
	value := ""
	r := regexp.MustCompile(key + "\\s")
	matches := r.FindStringSubmatchIndex(str)

	if len(matches) >= 0 {
		substr := str[matches[len(matches)-1]:]
		res := strings.Split(substr, "\n")
		value += res[0] + "\n"
		for i := 1; i < len(res); i++ {
			v := res[i]
			if strings.HasPrefix(v, "\t") {
				value += v + "\n"
			} else {
				return value
			}
		}
	}
	return value
}
