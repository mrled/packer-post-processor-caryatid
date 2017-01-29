package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonDecodingProvider(t *testing.T) {
	jstring := `{"name":"testname","url":"http://example.com/whatever","checksum_type":"dummy","checksum":"dummy"}`
	var prov Provider
	err := json.Unmarshal([]byte(jstring), &prov)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error unmarshalling JSON: %s", err))
	}
	if prov.Name != "testname" {
		t.Fatal(fmt.Sprintf("Decoded JSON object had bad Name property; should be 'testname' but was '%s'", prov.Name))
	}
}

func TestJsonDecodingCatalog(t *testing.T) {
	jstring := `{"name":"examplebox","description":"this is an example box","versions":[{"version":"12.34.56","providers":[{"name":"testname","url":"http://example.com/whatever","checksum_type":"dummy","checksum":"dummy"}]}]}`

	var cata Catalog
	err := json.Unmarshal([]byte(jstring), &cata)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error unmarshalling JSON: %s", err))
	}
	if cata.Name != "examplebox" {
		t.Fatal(fmt.Sprintf("Decoded JSON had bad Name property; should be 'examplebox' but was '%s'", cata.Name))
	}
	if len(cata.Versions) != 1 {
		t.Fatal(fmt.Sprintf("Expected decoded JSON to have %v elements in its Versions property, but actually had %v", 1, len(cata.Versions)))
	}
	vers := cata.Versions[0]
	if vers.Version != "12.34.56" {
		t.Fatal(fmt.Sprintf("Expected decoded JSON to have a Version with a version of '%s', but actually had a version of '%s'", "12.34.56", vers.Version))
	}
	if len(vers.Providers) != 1 {
		t.Fatal(fmt.Sprintf("Expected first Version to have %v elements in its Providers property, but actually had %v", 1, len(vers.Providers)))
	}
	prov := vers.Providers[0]
	if prov.Name != "testname" {
		t.Fatal(fmt.Sprintf("Expected first Provider to have a Name of '%s', but actually had '%s'", "testname", prov.Name))
	}
}

func TestJsonDecodingEmptyCatalog(t *testing.T) {
	var cata Catalog
	err := json.Unmarshal([]byte("{}"), &cata)
	if err != nil {
		t.Fatal("Failed to unmarshal empty catalog with error:", err)
	}
}