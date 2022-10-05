// Code generated by goctl-resty-discover. DO NOT EDIT.
package client
import(
    "context"
    "github.com/mengdj/goctl-rest-discover/conf"
    "github.com/mengdj/goctl-rest-discover/factory"
)
type(
    {{range .Type}}
    {{range .Docs -}}
    {{.}}
    {{- end}}
    {{.RawName}} struct{
    {{if .Members}}
    {{if .Docs}} {{range .Docs}}
    {{.}}
    {{- end}}{{- end}}
    {{range .Members -}}
        {{if .IsInline}}
            {{.Type.RawName}} {{if .Comment}}{{.Comment}}{{- end}}
        {{else}}
            {{.Name}} {{.Type.RawName}} {{if .Tag}}{{.Tag}}{{- end}} {{if .Comment}}{{.Comment}}{{- end}}
        {{- end}}
    {{- end}}
    {{- end}}
    }
    {{- end}}

    // Client
    Client interface{
        {{range .Route -}}
        {{range .Comment -}}
        // {{.}}
        {{- end}}
        // {{.Handler}}
        {{.Handler}}(context.Context{{if .RequestName}},*{{.RequestName}}{{- end}})(*{{.ResponseName}},error)
        {{- end}}
        Invoke(context.Context,string,string,interface{},interface{}) error
    }
    clientFactory struct{
        *factory.RestDiscoverFactory
    }
)

// MustClient
func MustClient(c conf.DiscoverConf) Client{
    return &clientFactory{
        RestDiscoverFactory:factory.NewRestDiscoverFactory(c),
    }
}

func (cf *clientFactory) Invoke(ctx context.Context,method string,path string,entity interface{},resp interface{}) error{
    return cf.RestDiscoverFactory.Invoke(ctx,method,path,entity,resp)
}

{{range .Route}}
{{range .Comment}}
// {{.}}
{{- end}}
// {{.Handler}} {{.Text}}
func (cf *clientFactory) {{.Handler}}(ctx context.Context{{if .RequestName}},entity *{{.RequestName}}{{- end}})(resp *{{.ResponseName}},err error){
    resp=new({{.ResponseName}})
    err=cf.Invoke(ctx,"{{.Method}}","{{.Path}}",{{if .RequestName}}entity{{else}}nil{{- end}},resp)
    if nil!=err{
        return nil,err
    }
    return resp,nil
}
{{- end}}
