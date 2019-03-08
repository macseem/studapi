package studentcom

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"
)

func TestNewUri(t *testing.T) {
	uri := NewUri()
	if uri.Scheme != "https" || uri.Host != "student.com" {
		t.Fail()
	}
}

func TestGetListAcc(t *testing.T) {
	properties := GetListAccProperties{}
	resp, err := GetListAcc(properties)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(resp)
	if err != nil{
		t.Fatal(err)
	}
	expectedAnswer := mockHttpGetList(url.URL{})
	buffer := new(bytes.Buffer)
	json.Compact(buffer, expectedAnswer)



	if string(b) != buffer.String() {
		t.Fatalf("%s are not equal to %s", string(b), buffer.String())
	}
}
