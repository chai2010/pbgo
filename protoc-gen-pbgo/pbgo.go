// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"log"
	"sort"
	"strings"
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"

	"github.com/chai2010/pbgo"
)

const pbgoPluginName = "pbgo"

func init() {
	generator.RegisterPlugin(new(pbgoPlugin))
}

type pbgoPlugin struct{ *generator.Generator }

func (p *pbgoPlugin) Name() string                { return pbgoPluginName }
func (p *pbgoPlugin) Init(g *generator.Generator) { p.Generator = g }

func (p *pbgoPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) == 0 {
		return
	}
	p.genImportCode(file)
}

func (p *pbgoPlugin) Generate(file *generator.FileDescriptor) {
	if len(file.Service) == 0 {
		return
	}

	p.genReferenceImportCode(file)
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
	RestAPIs       []ServiceRestMethodSpec
}

type ServiceRestMethodSpec struct {
	Method      string
	Url         string
	ContentType string
	ContentBody string
}

func (p *pbgoPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P(`import "encoding/json"`)
	p.P(`import "net/rpc"`)
	p.P(`import "net/http"`)
	p.P(`import "regexp"`)
	p.P(`import "strings"`)
	p.P()
	p.P(`import "github.com/chai2010/pbgo"`)
	p.P(`import "github.com/julienschmidt/httprouter"`)
}

func (p *pbgoPlugin) genReferenceImportCode(file *generator.FileDescriptor) {
	p.P("// Reference imports to suppress errors if they are not otherwise used.")
	p.P("var _ = json.Marshal")
	p.P("var _ = http.ListenAndServe")
	p.P("var _ = regexp.Match")
	p.P("var _ = strings.Split")
	p.P("var _ = pbgo.PopulateFieldFromPath")
	p.P("var _ = httprouter.New")
	p.P()
}

func (p *pbgoPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}

	p.P(buf.String())
}

func (p *pbgoPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
			RestAPIs:       p.buildRestMethodSpec(m),
		})
	}

	return spec
}

func (p *pbgoPlugin) buildRestMethodSpec(m *descriptor.MethodDescriptorProto) []ServiceRestMethodSpec {
	var restapis []ServiceRestMethodSpec

	restSpec := p.getServiceMethodOption(m)
	if restSpec == nil {
		return nil
	}

	for _, v := range restSpec.AdditionalBindings {
		if v.Method != "" && v.Url != "" {
			restapis = append(restapis, ServiceRestMethodSpec{
				Method:      v.Method,
				Url:         v.Url,
				ContentType: v.ContentType,
				ContentBody: v.ContentBody,
			})
		}
	}

	for _, v := range []pbgo.CustomHttpRule{
		{Method: "GET", Url: restSpec.Get},
		{Method: "PUT", Url: restSpec.Put},
		{Method: "POST", Url: restSpec.Post},
		{Method: "DELETE", Url: restSpec.Delete},
		{Method: "PATCH", Url: restSpec.Patch},
	} {
		if v.Method != "" && v.Url != "" {
			restapis = append(restapis, ServiceRestMethodSpec{
				Method: v.Method,
				Url:    v.Url,
			})
		}
	}

	for i, v := range restapis {
		if strings.HasPrefix(v.ContentType, ":") {
			ss := strings.Split(strings.TrimLeft(v.ContentBody, ":*"), ".")
			for i := 0; i < len(ss); i++ {
				ss[i] = generator.CamelCase(ss[i])
			}
			restapis[i].ContentType = strings.Join(ss, ".")
		}
		if strings.HasPrefix(v.ContentBody, ":") {
			ss := strings.Split(strings.TrimLeft(v.ContentBody, ":*"), ".")
			for i := 0; i < len(ss); i++ {
				ss[i] = generator.CamelCase(ss[i])
			}
			restapis[i].ContentBody = strings.Join(ss, ".")
		}
	}

	sort.Slice(restapis, func(i, j int) bool {
		vi := restapis[i].Method + restapis[i].Url
		vj := restapis[j].Method + restapis[j].Url
		return vi < vj
	})

	return restapis
}

func (p *pbgoPlugin) getServiceMethodOption(m *descriptor.MethodDescriptorProto) *pbgo.HttpRule {
	if m.Options != nil && proto.HasExtension(m.Options, pbgo.E_RestApi) {
		if ext, _ := proto.GetExtension(m.Options, pbgo.E_RestApi); ext != nil {
			if x, _ := ext.(*pbgo.HttpRule); x != nil {
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

func {{.ServiceName}}Handler(svc {{.ServiceName}}Interface) http.Handler {
	var router = httprouter.New()

	var re = regexp.MustCompile("(\\*|\\:)(\\w|\\.)+")
	_ = re

	{{range $_, $m := .MethodList}}
		{{range $_, $rest := .RestAPIs}}
			router.Handle("{{$rest.Method}}", "{{$rest.Url}}",
				func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
					var (
						protoReq   {{$m.InputTypeName}}
						protoReply {{$m.OutputTypeName}}
					)

					for _, fieldPath := range re.FindAllString("{{$rest.Url}}", -1) {
						fieldPath := strings.TrimLeft(fieldPath, ":*")
						err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
						if err != nil {
							http.Error(w, err.Error(), http.StatusBadRequest)
							return
						}
					}

					if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					if err := svc.{{$m.MethodName}}(&protoReq, &protoReply); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					{{if $rest.ContentType}}
						w.Header().Set("Content-Type", "{{$rest.ContentType}}")
					{{else}}
						if strings.Contains(r.Header.Get("Accept"), "application/json") {
							w.Header().Set("Content-Type", "application/json")
						} else {
							w.Header().Set("Content-Type", "text/plain")
						}
					{{end}}

					{{if $rest.ContentBody}}
						if _, err := w.Write(protoReply.{{$rest.ContentBody}}); err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
					{{else}}
						if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
					{{end}}
				},
			)
		{{end}}
	{{end}}

	return router
}
`
