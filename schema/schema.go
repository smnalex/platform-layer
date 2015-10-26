package schema

import (
	"github.com/hailocab/protobuf/proto"

	"github.com/hailocab/platform-layer/errors"
	"github.com/hailocab/platform-layer/server"

	"github.com/hailocab/go-hailo-lib/schema"
	schemaProto "github.com/hailocab/platform-layer/proto/schema"
	service "github.com/hailocab/platform-layer/server"
)

func Endpoint(name string, configStruct interface{}) *service.Endpoint {
	handler := func(req *server.Request) (proto.Message, errors.Error) {
		return &schemaProto.Response{
			Schema: proto.String(schema.Of(configStruct).String()),
		}, nil
	}

	return &server.Endpoint{
		Name:       name,
		Mean:       200,
		Upper95:    400,
		Handler:    handler,
		Authoriser: service.OpenToTheWorldAuthoriser(),
	}
}
