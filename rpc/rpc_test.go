package rpc_test

import (
	"glsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Contest-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("Expected 15, got %d", contentLength)
	}
	if method != "hi" {
		t.Fatalf("Expected hi, got %s", method)
	}
}
