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
// curl -X POST -d'{"dest":"1.1.1.1", "nodeName":"nodeX"}'  http://127.0.0.1:8080/api/v1/backends -H "Authorization: Bearer <token>" -v
//
// The path of the endpoint will be: /backends
// @method POST
// @path /backends
// @protected
func RegisterBackend(user *registry.User, dest string, nodeName string) (result bool, err error) {

	return true, nil
}