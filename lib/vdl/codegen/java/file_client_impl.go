// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package java

import (
	"bytes"
	"fmt"
	"log"
	"path"

	"v.io/x/ref/lib/vdl/compile"
	"v.io/x/ref/lib/vdl/vdlutil"
)

const clientImplTmpl = header + `
// Source(s):  {{ .Source }}
package {{ .PackagePath }};

/**
 * Implementation of the {@link {{ .ServiceName }}Client} interface.
 */
final class {{ .ServiceName }}ClientImpl implements {{ .FullServiceName }}Client {
    private final io.v.v23.rpc.Client client;
    private final java.lang.String vName;

    {{/* Define fields to hold each of the embedded object impls*/}}
    {{ range $embed := .Embeds }}
    {{/* e.g. private final com.somepackage.ArithClient implArith; */}}
    private final {{ $embed.FullName }}Client impl{{ $embed.Name }};
    {{ end }}

    /**
     * Creates a new instance of {@link {{ .ServiceName }}ClientImpl}.
     *
     * @param client Vanadium client
     * @param vName  remote server name
     */
    public {{ .ServiceName }}ClientImpl(io.v.v23.rpc.Client client, java.lang.String vName) {
        this.client = client;
        this.vName = vName;
        {{/* Initialize the embeded impls */}}
        {{ range $embed := .Embeds }}
        {
            io.v.v23.Options opts = new io.v.v23.Options();
            opts.set(io.v.v23.OptionDefs.CLIENT, client);
            this.impl{{ $embed.Name }} = {{ $embed.FullName }}ClientFactory.get{{ $embed.Name }}Client(vName, opts);
        }
        {{ end }}
    }

    private io.v.v23.rpc.Client getClient(io.v.v23.context.VContext context) {
        return this.client != null ? client : io.v.v23.V.getClient(context);

    }

    // Methods from interface {{ .ServiceName }}Client.
{{/* Iterate over methods defined directly in the body of this service */}}
{{ range $method := .Methods }}
    {{/* The optionless overload simply calls the overload with options */}}
    @Override
    public {{ $method.RetType }} {{ $method.Name }}(io.v.v23.context.VContext context{{ $method.DeclarationArgs }}) throws io.v.v23.verror.VException {
        {{if $method.Returns }}return{{ end }} {{ $method.Name }}(context{{ $method.CallingArgsLeadingComma }}, (io.v.v23.Options) null);
    }
    {{/* The main client impl method body */}}
    @Override
    public {{ $method.RetType }} {{ $method.Name }}(io.v.v23.context.VContext context{{ $method.DeclarationArgs }}, io.v.v23.Options vOpts) throws io.v.v23.verror.VException {
        {{/* Start the vanadium call */}}
        // Start the call.
        java.lang.Object[] _args = new java.lang.Object[]{ {{ $method.CallingArgs }} };
        java.lang.reflect.Type[] _argTypes = new java.lang.reflect.Type[]{ {{ $method.CallingArgTypes }} };
        final io.v.v23.rpc.ClientCall _call = getClient(context).startCall(context, this.vName, "{{ $method.Name }}", _args, _argTypes, vOpts);

        // Finish the call.
        {{/* Now handle returning from the function. */}}
        {{ if $method.NotStreaming }}

        {{ if $method.IsVoid }}
        java.lang.reflect.Type[] _resultTypes = new java.lang.reflect.Type[]{};
        _call.finish(_resultTypes);
        {{ else }} {{/* else $method.IsVoid */}}
        java.lang.reflect.Type[] _resultTypes = new java.lang.reflect.Type[]{
            {{ range $outArg := $method.OutArgs }}
            new com.google.common.reflect.TypeToken<{{ $outArg.Type }}>() {}.getType(),
            {{ end }}
        };
        java.lang.Object[] _results = _call.finish(_resultTypes);
        {{ if $method.MultipleReturn }}
        {{ $method.DeclaredObjectRetType }} _ret = new {{ $method.DeclaredObjectRetType }}();
            {{ range $i, $outArg := $method.OutArgs }}
        _ret.{{ $outArg.FieldName }} = ({{ $outArg.Type }})_results[{{ $i }}];
            {{ end }} {{/* end range over outargs */}}
        return _ret;
        {{ else }} {{/* end if $method.MultipleReturn */}}
        return ({{ $method.DeclaredObjectRetType }})_results[0];
        {{ end }} {{/* end if $method.MultipleReturn */}}

        {{ end }} {{/* end if $method.IsVoid */}}

        {{else }} {{/* else $method.NotStreaming */}}
        return new io.v.v23.vdl.TypedClientStream<{{ $method.SendType }}, {{ $method.RecvType }}, {{ $method.DeclaredObjectRetType }}>() {
            @Override
            public void send({{ $method.SendType }} item) throws io.v.v23.verror.VException {
                java.lang.reflect.Type type = new com.google.common.reflect.TypeToken<{{ $method.SendType }}>() {}.getType();
                _call.send(item, type);
            }
            @Override
            public {{ $method.RecvType }} recv() throws java.io.EOFException, io.v.v23.verror.VException {
                java.lang.reflect.Type type = new com.google.common.reflect.TypeToken<{{ $method.RecvType }}>() {}.getType();
                java.lang.Object result = _call.recv(type);
                try {
                    return ({{ $method.RecvType }})result;
                } catch (java.lang.ClassCastException e) {
                    throw new io.v.v23.verror.VException("Unexpected result type: " + result.getClass().getCanonicalName());
                }
            }
            @Override
            public {{ $method.DeclaredObjectRetType }} finish() throws io.v.v23.verror.VException {
                {{ if $method.IsVoid }}
                java.lang.reflect.Type[] resultTypes = new java.lang.reflect.Type[]{};
                _call.finish(resultTypes);
                return null;
                {{ else }} {{/* else $method.IsVoid */}}
                java.lang.reflect.Type[] resultTypes = new java.lang.reflect.Type[]{
                    new com.google.common.reflect.TypeToken<{{ $method.DeclaredObjectRetType }}>() {}.getType()
                };
                return ({{ $method.DeclaredObjectRetType }})_call.finish(resultTypes)[0];
                {{ end }} {{/* end if $method.IsVoid */}}
            }
        };
        {{ end }}{{/* end if $method.NotStreaming */}}
    }
    @Override
    public void {{ $method.Name }}(io.v.v23.context.VContext context{{ $method.DeclarationArgs }}, io.v.v23.rpc.Callback<{{ $method.GenericRetType }}> callback) throws io.v.v23.verror.VException {
        {{ $method.Name }}(context{{ $method.CallingArgsLeadingComma }}, null, callback);
    }
    @Override
    public void {{ $method.Name }}(io.v.v23.context.VContext context{{ $method.DeclarationArgs }}, io.v.v23.Options vOpts, final io.v.v23.rpc.Callback<{{ $method.GenericRetType }}> callback) throws io.v.v23.verror.VException {
        {{/* Start the vanadium call */}}
        // Start the call.
        java.lang.Object[] _args = new java.lang.Object[]{ {{ $method.CallingArgs }} };
        java.lang.reflect.Type[] _argTypes = new java.lang.reflect.Type[]{ {{ $method.CallingArgTypes }} };
        io.v.v23.rpc.Callback<io.v.v23.rpc.ClientCall> clientCallback = new io.v.v23.rpc.Callback<io.v.v23.rpc.ClientCall>() {
            @Override
            public void onSuccess(final io.v.v23.rpc.ClientCall _call) {
                // Finish the call.
                {{ if $method.NotStreaming }}

                io.v.v23.rpc.Callback<Object[]> finishCallback = new io.v.v23.rpc.Callback<Object[]>() {
                    @Override
                    public void onSuccess(Object[] _results) {
                        {{ if $method.IsVoid }}
                            callback.onSuccess(null);
                        {{ else if $method.MultipleReturn }}
                        {{ $method.DeclaredObjectRetType }} _ret = new {{ $method.DeclaredObjectRetType }}();
                            {{ range $i, $outArg := $method.OutArgs }}
                        _ret.{{ $outArg.FieldName }} = ({{ $outArg.Type }})_results[{{ $i }}];
                            {{ end }} {{/* end range over outargs */}}
                        callback.onSuccess(_ret);
                        {{ else }} {{/* else if $method.MultipleReturn */}}
                        callback.onSuccess(({{ $method.DeclaredObjectRetType }})_results[0]);
                        {{ end }} {{/* end if $method.IsVoid */}}
                    }
                    @Override
                    public void onFailure(io.v.v23.verror.VException error) {
                        callback.onFailure(error);
                    }
                };

                {{ if $method.IsVoid }}
                java.lang.reflect.Type[] _resultTypes = new java.lang.reflect.Type[]{};

                {{ else }} {{/* else $method.IsVoid */}}
                java.lang.reflect.Type[] _resultTypes = new java.lang.reflect.Type[]{
                    {{ range $outArg := $method.OutArgs }}
                    new com.google.common.reflect.TypeToken<{{ $outArg.Type }}>() {}.getType(),
                    {{ end }}
                };
                {{ end }} {{/* end if $method.IsVoid */}}
                _call.finish(_resultTypes, finishCallback);

                {{else }} {{/* else $method.NotStreaming */}}
                callback.onSuccess(new io.v.v23.vdl.TypedClientStream<{{ $method.SendType }}, {{ $method.RecvType }}, {{ $method.DeclaredObjectRetType }}>() {
                    @Override
                    public void send({{ $method.SendType }} item) throws io.v.v23.verror.VException {
                        java.lang.reflect.Type type = new com.google.common.reflect.TypeToken<{{ $method.SendType }}>() {}.getType();
                        _call.send(item, type);
                    }
                    @Override
                    public {{ $method.RecvType }} recv() throws java.io.EOFException, io.v.v23.verror.VException {
                        java.lang.reflect.Type type = new com.google.common.reflect.TypeToken<{{ $method.RecvType }}>() {}.getType();
                        java.lang.Object result = _call.recv(type);
                        try {
                            return ({{ $method.RecvType }})result;
                        } catch (java.lang.ClassCastException e) {
                            throw new io.v.v23.verror.VException("Unexpected result type: " + result.getClass().getCanonicalName());
                        }
                    }
                    @Override
                    public {{ $method.DeclaredObjectRetType }} finish() throws io.v.v23.verror.VException {
                        {{ if $method.IsVoid }}
                        java.lang.reflect.Type[] resultTypes = new java.lang.reflect.Type[]{};
                        _call.finish(resultTypes);
                        return null;
                        {{ else }} {{/* else $method.IsVoid */}}
                        java.lang.reflect.Type[] resultTypes = new java.lang.reflect.Type[]{
                            new com.google.common.reflect.TypeToken<{{ $method.DeclaredObjectRetType }}>() {}.getType()
                        };
                        return ({{ $method.DeclaredObjectRetType }})_call.finish(resultTypes)[0];
                        {{ end }} {{/* end if $method.IsVoid */}}
                    }
                });
                {{ end }}{{/* end if $method.NotStreaming */}}

            }
            @Override
            public void onFailure(io.v.v23.verror.VException error) {
                callback.onFailure(error);
            }
        };

        getClient(context).startCall(context, this.vName, "{{ $method.Name }}", _args, _argTypes, vOpts, clientCallback);
    }
{{ end }}{{/* end range over methods */}}

{{/* Iterate over methods from embeded services and generate code to delegate the work */}}
{{ range $eMethod := .EmbedMethods }}
    @Override
    public {{ $eMethod.RetType }} {{ $eMethod.Name }}(io.v.v23.context.VContext context{{ $eMethod.DeclarationArgs }}) throws io.v.v23.verror.VException {
        {{/* e.g. return this.implArith.cosine(context, [args]) */}}
        {{ if $eMethod.Returns }}return{{ end }} this.impl{{ $eMethod.IfaceName }}.{{ $eMethod.Name }}(context{{ $eMethod.CallingArgsLeadingComma }});
    }
    @Override
    public {{ $eMethod.RetType }} {{ $eMethod.Name }}(io.v.v23.context.VContext context{{ $eMethod.DeclarationArgs }}, io.v.v23.Options vOpts) throws io.v.v23.verror.VException {
        {{/* e.g. return this.implArith.cosine(context, [args], options) */}}
        {{ if $eMethod.Returns }}return{{ end }}  this.impl{{ $eMethod.IfaceName }}.{{ $eMethod.Name }}(context{{ $eMethod.CallingArgsLeadingComma }}, vOpts);
    }
    @Override
    public void {{ $eMethod.Name }}(io.v.v23.context.VContext context{{ $eMethod.DeclarationArgs }}, io.v.v23.rpc.Callback<{{ $eMethod.GenericRetType }}> callback) throws io.v.v23.verror.VException {
        this.impl{{ $eMethod.IfaceName}}.{{ $eMethod.Name }}(context{{ $eMethod.CallingArgsLeadingComma }}, null, callback);
    }
    @Override
    public void {{ $eMethod.Name }}(io.v.v23.context.VContext context{{ $eMethod.DeclarationArgs }}, io.v.v23.Options vOpts, io.v.v23.rpc.Callback<{{ $eMethod.GenericRetType }}> callback) throws io.v.v23.verror.VException {
        this.impl{{ $eMethod.IfaceName}}.{{ $eMethod.Name }}(context{{ $eMethod.CallingArgsLeadingComma }}, vOpts, callback);
    }
{{ end }}

}
`

type clientImplMethodOutArg struct {
	FieldName string
	Type      string
}

type clientImplMethod struct {
	CallingArgs             string
	CallingArgTypes         string
	CallingArgsLeadingComma string
	DeclarationArgs         string
	DeclaredObjectRetType   string
	IsVoid                  bool
	MultipleReturn          bool
	Name                    string
	NotStreaming            bool
	OutArgs                 []clientImplMethodOutArg
	RecvType                string
	RetType                 string
	GenericRetType          string
	Returns                 bool
	SendType                string
	ServiceName             string
}

type clientImplEmbedMethod struct {
	CallingArgsLeadingComma string
	DeclarationArgs         string
	IfaceName               string
	Name                    string
	RetType                 string
	GenericRetType          string
	Returns                 bool
}

type clientImplEmbed struct {
	Name     string
	FullName string
}

func processClientImplMethod(iface *compile.Interface, method *compile.Method, env *compile.Env) clientImplMethod {
	outArgs := make([]clientImplMethodOutArg, len(method.OutArgs))
	for i := 0; i < len(method.OutArgs); i++ {
		if method.OutArgs[i].Name != "" {
			outArgs[i].FieldName = vdlutil.FirstRuneToLower(method.OutArgs[i].Name)
		} else {
			outArgs[i].FieldName = fmt.Sprintf("ret%d", i+1)
		}
		outArgs[i].Type = javaType(method.OutArgs[i].Type, true, env)
	}
	return clientImplMethod{
		CallingArgs:             javaCallingArgStr(method.InArgs, false),
		CallingArgTypes:         javaCallingArgTypeStr(method.InArgs, env),
		CallingArgsLeadingComma: javaCallingArgStr(method.InArgs, true),
		DeclarationArgs:         javaDeclarationArgStr(method.InArgs, env, true),
		DeclaredObjectRetType:   clientInterfaceNonStreamingOutArg(iface, method, true, env),
		IsVoid:                  len(method.OutArgs) < 1,
		MultipleReturn:          len(method.OutArgs) > 1,
		Name:                    vdlutil.FirstRuneToLower(method.Name),
		NotStreaming:            !isStreamingMethod(method),
		OutArgs:                 outArgs,
		RecvType:                javaType(method.OutStream, true, env),
		RetType:                 clientInterfaceOutArg(iface, method, env, false),
		GenericRetType:          clientInterfaceOutArg(iface, method, env, true),
		Returns:                 len(method.OutArgs) >= 1 || isStreamingMethod(method),
		SendType:                javaType(method.InStream, true, env),
		ServiceName:             vdlutil.FirstRuneToUpper(iface.Name),
	}
}

func processClientImplEmbedMethod(iface *compile.Interface, embedMethod *compile.Method, env *compile.Env) clientImplEmbedMethod {
	return clientImplEmbedMethod{
		CallingArgsLeadingComma: javaCallingArgStr(embedMethod.InArgs, true),
		DeclarationArgs:         javaDeclarationArgStr(embedMethod.InArgs, env, true),
		IfaceName:               vdlutil.FirstRuneToUpper(iface.Name),
		Name:                    vdlutil.FirstRuneToLower(embedMethod.Name),
		RetType:                 clientInterfaceOutArg(iface, embedMethod, env, false),
		GenericRetType:          clientInterfaceOutArg(iface, embedMethod, env, true),
		Returns:                 len(embedMethod.OutArgs) >= 1 || isStreamingMethod(embedMethod),
	}
}

// genJavaClientImplFile generates a client impl for the specified interface.
func genJavaClientImplFile(iface *compile.Interface, env *compile.Env) JavaFileInfo {
	embeds := []clientImplEmbed{}
	for _, embed := range allEmbeddedIfaces(iface) {
		embeds = append(embeds, clientImplEmbed{
			Name:     vdlutil.FirstRuneToUpper(embed.Name),
			FullName: javaPath(javaGenPkgPath(path.Join(embed.File.Package.GenPath, vdlutil.FirstRuneToUpper(embed.Name)))),
		})
	}
	embedMethods := []clientImplEmbedMethod{}
	for _, embedMao := range dedupedEmbeddedMethodAndOrigins(iface) {
		embedMethods = append(embedMethods, processClientImplEmbedMethod(embedMao.Origin, embedMao.Method, env))
	}
	methods := make([]clientImplMethod, len(iface.Methods))
	for i, method := range iface.Methods {
		methods[i] = processClientImplMethod(iface, method, env)
	}
	javaServiceName := vdlutil.FirstRuneToUpper(iface.Name)
	data := struct {
		FileDoc         string
		EmbedMethods    []clientImplEmbedMethod
		Embeds          []clientImplEmbed
		FullServiceName string
		Methods         []clientImplMethod
		PackagePath     string
		ServiceName     string
		Source          string
	}{
		FileDoc:         iface.File.Package.FileDoc,
		EmbedMethods:    embedMethods,
		Embeds:          embeds,
		FullServiceName: javaPath(interfaceFullyQualifiedName(iface)),
		Methods:         methods,
		PackagePath:     javaPath(javaGenPkgPath(iface.File.Package.GenPath)),
		ServiceName:     javaServiceName,
		Source:          iface.File.BaseName,
	}
	var buf bytes.Buffer
	err := parseTmpl("client impl", clientImplTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute client impl template: %v", err)
	}
	return JavaFileInfo{
		Name: javaServiceName + "ClientImpl.java",
		Data: buf.Bytes(),
	}
}
