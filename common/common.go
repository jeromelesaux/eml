package common

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
	ReplyTo                 = "Reply-To:"
	InReplyTo               = "In-Reply-To:"
	Boundary                = "boundary=\"(.*)\""
	XSender                 = "X-Sender:"
	XReceiver               = "X-Receiver:"
	ContentId               = "Content-Id:"

	SimpleValueRegex = "([^\n]*)"
)
