package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/gogo/protobuf/jsonpb"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	goproto "google.golang.org/protobuf/proto"
	go_proto "request-response-generate-with-protobuf/generated_go"
	response "request-response-generate-with-protobuf/generated_gogo/response/gogo_proto"
	"testing"
)

func Test_ToGogoResponse_WithGogoProto(t *testing.T) {
	var p Product
	err := json.Unmarshal(product_json, &p)
	if err != nil {
		println(err.Error())
	}
	var expected response.ProductResponse
	err = gogoproto.Unmarshal(product_json, &expected)
	if err != nil {
		println(err.Error())
	}
	actual := p.ToGogoResponse()

	assert.NotNil(t, err)
	assert.NotEqual(t, expected, actual)
}

func Test_ToGogoResponse_WithGoProto(t *testing.T) {
	var p Product
	err := json.Unmarshal(product_json, &p)
	if err != nil {
		println(err.Error())
	}
	var expected go_proto.ProductResponse
	err = goproto.Unmarshal(product_json, &expected)
	if err != nil {
		println(err.Error())
	}
	actual := p.ToGogoResponse()

	assert.NotNil(t, err)
	assert.NotEqual(t, &expected, actual)
}

func Test_ToGogoResponse_ValidBody(t *testing.T) {
	var p Product
	err := json.Unmarshal(product_json, &p)
	if err != nil {
		println(err.Error())
	}
	var expected response.ProductResponse
	unmarshall := jsonpb.Unmarshaler{}
	buf := bytes.NewBuffer(product_json)
	err = unmarshall.Unmarshal(buf, &expected)
	if err != nil {
		println(err.Error())
	}
	actual := p.ToGogoResponse()

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ToGogoResponse_EmptyBody(t *testing.T) {
	p := response.ProductResponse{}
	marshaller := jsonpb.Marshaler{
		EmitDefaults: true,
	}

	expectedResponseString, err := marshaller.MarshalToString(&p)
	if err != nil {
		println(err.Error())
	}

	var expected response.ProductResponse
	unmarshall := jsonpb.Unmarshaler{}
	buf := bytes.NewBuffer([]byte(expectedResponseString))
	err = unmarshall.Unmarshal(buf, &expected)
	if err != nil {
		println(err.Error())
	}
	actualResponseString, err := marshaller.MarshalToString(&expected)
	if err != nil {
		println(err.Error())
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedResponseString, actualResponseString)
}
