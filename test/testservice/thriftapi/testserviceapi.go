package main

import (
	"github.com/vaporz/turbo"
	"github.com/vaporz/turbo/test/testservice/gen"
	"github.com/vaporz/turbo/test/testservice/thriftapi/component"
)

func main() {
	s := turbo.NewThriftServer(turbo.GOPATH() + "/src/github.com/vaporz/turbo/test/testservice/service.yaml")
	component.RegisterComponents(s)
	s.Initializer = &component.ServiceInitializer{}
	s.StartThriftHTTPServer(component.ThriftClient, gen.ThriftSwitcher)
}
