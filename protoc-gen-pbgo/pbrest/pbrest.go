// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbrest

import (
	"bytes"
	"log"
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"

	"github.com/chai2010/pbgo"
)

func init() {
	generator.RegisterPlugin(new(pbrestPlugin))
}

type pbrestPlugin struct{ *generator.Generator }

func (p *pbrestPlugin) Name() string                { return "pbrest" }
func (p *pbrestPlugin) Init(g *generator.Generator) { p.Generator = g }

func (p *pbrestPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

func (p *pbrestPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName       string
	InputTypeName    string
	OutputTypeName   string
	RestMethodOption *pbgo.RestMethodOption
}

func (p *pbrestPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P(`import "net/rpc"`)
}

func (p *pbrestPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}

	p.P(buf.String())
}

func (p *pbrestPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:       generator.CamelCase(m.GetName()),
			InputTypeName:    p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName:   p.TypeName(p.ObjectNamed(m.GetOutputType())),
			RestMethodOption: p.getServiceMethodOption(m),
		})
	}

	return spec
}

func (p *pbrestPlugin) getServiceMethodOption(m *descriptor.MethodDescriptorProto) *pbgo.RestMethodOption {
	if m.Options != nil && proto.HasExtension(m.Options, pbgo.E_RestMethodOption) {
		if ext, _ := proto.GetExtension(m.Options, pbgo.E_RestMethodOption); ext != nil {
			if x, _ := ext.(*pbgo.RestMethodOption); x != nil {
				return x
			}
		}
	}
	return nil
}

const tmplService = `
{{$root := .}}

type {{.ServiceName}}Interface interface {
	{{- range $_, $m := .MethodList}}
		{{$m.MethodName}}(in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}}) error
	{{- end}}
}

func Register{{.ServiceName}}(srv *rpc.Server, x {{.ServiceName}}Interface) error {
	if err := srv.RegisterName("{{.ServiceName}}", x); err != nil {
		return err
	}
	return nil
}

type {{.ServiceName}}Client struct {
	*rpc.Client
}

func Dial{{.ServiceName}}(network, address string) (*{{.ServiceName}}Client, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &{{.ServiceName}}Client{Client: c}, nil
}

{{range $_, $m := .MethodList}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(in *{{$m.InputTypeName}}) (*{{$m.OutputTypeName}}, error) {
	var out = new({{$m.OutputTypeName}})
	if err := p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out); err != nil {
		return nil, err
	}
	return out, nil
}
{{end}}
`
