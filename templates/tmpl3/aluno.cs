using Stefanini.Domain.Entity;
using System;

namespace Stefanini.DemoApp.Domain.Entities
{
    public class {{ .Table.Name}} : BaseEntity
    {
    {{- range .Table.Columns }}    
        public virtual {{ .Type }} {{ .Name }} { get; set; }  
    {{- end}}
    }
}