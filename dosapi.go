package dosapi // import "github.com/SteelzZ/dosapi"

import (
	"github.com/nildev/lib/registry"
)

//go:generate nildev io --sourceDir github.com/SteelzZ/dosapi --tpl simple-handlers --org nildev --ver v0.1.0
//go:generate nildev r --services github.com/SteelzZ/dosapi --containerDir github.com/nildev/api-host --tpl simple-router --org nildev --ver v0.1.0

// RegisterBackend is used to register new backend for load balancer
//
// Sample request:
//
// curl -X POST -d'{"dest":"1.1.1.1", "nodeName":"nodeX"}'  http://192.168.99.100:8082/api/v1/nodes -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvLmJhcnRrdXNAZ21haWwuY29tIiwiZXhwIjoxNDU5MDg5MjQ3fQ.GnyZ3I569AsVjS4f-yDpov-TdcTfLXhtvF5Y8F-2yYM" -v
//
// The path of the endpoint will be: /nodes
// @method POST
// @path /nodes
// @protected
func RegisterBackend(user *registry.User, dest string, nodeName string) (result bool, err error) {

	return true, nil
}