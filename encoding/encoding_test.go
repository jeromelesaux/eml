package encoding

import (
	"os"
	"testing"

	"github.com/jeromelesaux/eml"
	"github.com/jeromelesaux/eml/common"
)

func TestSimpleValueMatch(t *testing.T) {
	v := "From: foo@example.com"
	m := getSimpleValue(common.From, v)
	if m != "foo@example.com" {
		t.Fatal("expected matched.")
	}

	v1 := "To: blah@example.com"
	m = getSimpleValue(common.To, v1)
	if m != "blah@example.com" {
		t.Fatal("expected matched.")
	}
}

/*func TestMultineValueMatch(t *testing.T) {
	v0 := "Received: "
	v1 := "from cpc1-mort1-0-0-cust399.croy.cable.virginmedia.com ([82.44.61.144] helo=spike)\n" +
		"	by mail.example.com with esmtpsa (TLS1.0:RSA_AES_128_CBC_SHA1:16)\n" +
		"	(Exim 4.69)\n" +
		"	(envelope-from <from@example.com>)\n" +
		"	id 1NvUrP-0002Nw-6s" +
		"	for to@example.com; Sat, 27 Mar 2010 12:11:27 +0000\n"
	v3 := "From: \"Example\" <from@example.com>\n"
	v := v0 + v1 + v3

	m := getMultilineValue(common.Received, v)
	if m != v1 {
		t.Fatal("expected matched between [" + v1 + "] *** and *** [" + m + "]")
	}
}*/

func TestDecode(t *testing.T) {
	f, err := os.Open("../samples/attachment_with_encoded_name.eml.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	v := &eml.Eml{}
	if err = NewDecoder(f).Decode(v); err != nil {
		t.Fatal(err.Error())
	}
}

func TestAttachments(t *testing.T) {
	f1 := "/Users/jeromelesaux/Downloads/y2020.bas"
	f2 := "/Users/jeromelesaux/Downloads/crownland.dsk"
	f3 := "/Users/jeromelesaux/Downloads/y2020.dsk"
	e := eml.NewEml()
	e.AddAttachment(f1)
	e.AddAttachment(f2)
	e.AddAttachment(f3)
	e.From = "change@me.net"
	e.To = "change@me.net"
	e.XSender = "change@me.net"
	e.XReceiver = "change@me.net"
	mail, err := os.Create("test.eml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer mail.Close()
	if err := NewEncoder(mail).Encode(e); err != nil {
		t.Fatal(err.Error())
	}
}
