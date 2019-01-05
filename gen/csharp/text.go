package csharp

// 报错行号+7
const TemplateText = `// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
using System;
using System.Collections.Generic;
using ProtoPlus;

namespace Proto
{
	{{range $a, $enumobj := .Enums}}
	public enum {{.Name}} 
	{
		{{range .Fields}}
		{{.Name}} = {{TagNumber $enumobj .}}, {{end}}
	} {{end}}
	{{range $a, $obj := .Structs}}
	{{ObjectLeadingComment .}}
	public partial class {{$obj.Name}} : {{$.StructBase}} 
	{
		{{range .Fields}}public {{CSTypeNameFull .}} {{.Name}};
		{{end}}
		#region Serialize Code
		public void Init( )
		{   {{range .Fields}}{{if IsPrimitiveSlice .}}
			{{.Name}} = new {{CSTypeNameFull .}}();	{{end}}{{end}}
 			{{range .Fields}}{{if IsStruct .}}
			{{.Name}} = ({{CSTypeNameFull .}}) InputStream.CreateStruct(typeof({{CSTypeNameFull .}})); {{end}} {{end}}
		}

		public void Marshal(OutputStream stream)
		{ {{range .Fields}} 
			stream.Write{{CodecName .}}({{TagNumber $obj .}}, {{.Name}} ); {{end}}
		}

		public int GetSize()
		{
			int size = 0; {{range .Fields}} 
			size += OutputStream.Size{{CodecName .}}({{TagNumber $obj .}}, {{.Name}}); {{end}}
			return size;
		}

 		public bool Unmarshal(InputStream stream, int fieldNumber, WireFormat.WireType wt)
		{
		 	switch (fieldNumber)
            { {{range .Fields}}
			case {{TagNumber $obj .}}:	
				stream.Read{{CodecName .}}(wt, ref {{.Name}});
                break; {{end}}
			default:
				return true;
            }

            return false;
		}
		#endregion
	}
{{end}}
}
`
