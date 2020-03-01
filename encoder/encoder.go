package encoder

type Eml struct {
	From        string
	To          string
	Received    []string
	MimeVersion string
	Subject     string
	MessageId   string
	ContentType string
	Body        []byte
	Message     *Eml
}
