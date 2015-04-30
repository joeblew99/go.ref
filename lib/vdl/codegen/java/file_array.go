// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package java

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"v.io/x/ref/lib/vdl/compile"
	"v.io/x/ref/lib/vdl/vdlutil"
)

const arrayTmpl = header + `
// Source: {{.SourceFile}}

package {{.Package}};

/**
 * type {{.Name}} {{.VdlTypeString}} {{.Doc}}
 **/
@io.v.v23.vdl.GeneratedFromVdl(name = "{{.VdlTypeName}}")
@io.v.v23.vdl.ArrayLength({{.Length}})
{{ .AccessModifier }} class {{.Name}} extends io.v.v23.vdl.VdlArray<{{.ElemType}}> {
    private static final long serialVersionUID = 1L;

    public static final io.v.v23.vdl.VdlType VDL_TYPE =
            io.v.v23.vdl.Types.getVdlTypeFromReflect({{.Name}}.class);

    public {{.Name}}({{.ElemType}}[] arr) {
        super(VDL_TYPE, arr);
    }

    public {{.Name}}() {
        this({{.ZeroValue}});
    }

    {{ if .ElemIsPrimitive }}
    public {{.Name}}({{ .ElemPrimitiveType }}[] arr) {
        super(VDL_TYPE, convert(arr));
    }

    private static {{ .ElemType }}[] convert({{ .ElemPrimitiveType }}[] arr) {
        final {{ .ElemType }}[] ret = new {{ .ElemType }}[arr.length];
        for (int i = 0; i < arr.length; ++i) {
            ret[i] = arr[i];
        }
        return ret;
    }
    {{ end }}
}
`

// genJavaArrayFile generates the Java class file for the provided named array type.
func genJavaArrayFile(tdef *compile.TypeDef, env *compile.Env) JavaFileInfo {
	javaTypeName := vdlutil.FirstRuneToUpper(tdef.Name)
	elemType := javaType(tdef.Type.Elem(), true, env)
	elems := strings.TrimSuffix(strings.Repeat(javaZeroValue(tdef.Type.Elem(), env)+", ", tdef.Type.Len()), ", ")
	zeroValue := fmt.Sprintf("new %s[] {%s}", elemType, elems)
	data := struct {
		FileDoc           string
		AccessModifier    string
		Doc               string
		ElemType          string
		ElemIsPrimitive   bool
		ElemPrimitiveType string
		Length            int
		Name              string
		Package           string
		SourceFile        string
		VdlTypeName       string
		VdlTypeString     string
		ZeroValue         string
	}{
		FileDoc:           tdef.File.Package.FileDoc,
		AccessModifier:    accessModifierForName(tdef.Name),
		Doc:               javaDocInComment(tdef.Doc),
		ElemType:          elemType,
		ElemIsPrimitive:   !isClass(tdef.Type.Elem(), env),
		ElemPrimitiveType: javaType(tdef.Type.Elem(), false, env),
		Length:            tdef.Type.Len(),
		Name:              javaTypeName,
		Package:           javaPath(javaGenPkgPath(tdef.File.Package.GenPath)),
		SourceFile:        tdef.File.BaseName,
		VdlTypeName:       tdef.Type.Name(),
		VdlTypeString:     tdef.Type.String(),
		ZeroValue:         zeroValue,
	}
	var buf bytes.Buffer
	err := parseTmpl("array", arrayTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute array template: %v", err)
	}
	return JavaFileInfo{
		Name: javaTypeName + ".java",
		Data: buf.Bytes(),
	}
}
