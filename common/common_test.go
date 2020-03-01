package common

import "testing"

func TestSimpleValueMatch(t *testing.T) {
	v := "From: foo@example.com"
	m := getSimpleValue(From, v)
	if m != "foo@example.com" {
		t.Fatal("expected matched.")
	}

	v1 := "To: blah@example.com"
	m = getSimpleValue(To, v1)
	if m != "blah@example.com" {
		t.Fatal("expected matched.")
	}
}

func TestMultineValueMatch(t *testing.T) {
	v0 := "Received: "
	v1 := "from cpc1-mort1-0-0-cust399.croy.cable.virginmedia.com ([82.44.61.144] helo=spike)\n" +
		"	by mail.example.com with esmtpsa (TLS1.0:RSA_AES_128_CBC_SHA1:16)\n" +
		"	(Exim 4.69)\n" +
		"	(envelope-from <from@example.com>)\n" +
		"	id 1NvUrP-0002Nw-6s" +
		"	for to@example.com; Sat, 27 Mar 2010 12:11:27 +0000\n"
	v3 := "From: \"Example\" <from@example.com>\n"
	v := v0 + v1 + v3

	m := getMutlineValue(Received, v)
	if m != v1 {
		t.Fatal("expected matched between [" + v1 + "] *** and *** [" + m + "]")
	}
}
