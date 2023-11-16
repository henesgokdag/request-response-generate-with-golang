package main

import (
	_ "embed"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"request-response-generate-with-protobuf/generated_go"
	"testing"
)

//go:embed test_data/product.json
var product_json []byte

func Test_ToGoResponse(t *testing.T) {
	var p Product
	err := json.Unmarshal(product_json, &p)
	if err != nil {
		println(err.Error())
	}
	expected := go_proto.ProductResponse{}
	err = proto.Unmarshal(product_json, &expected)
	expectedS := (*ProductGoResponse)(&expected)
	if err != nil {
		println(err.Error())
	}
	actual := p.TogoResponse()

	assert.NotNil(t, err)
	assert.NotEqual(t, expectedS, &actual)
}
