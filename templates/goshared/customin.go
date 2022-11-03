package goshared

const custominTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.CustomIn }}
		{{ $val := accessor . }}
		{{ $valContext := . }}
		{{- range $r.CustomIn }}
		if !validaterules.Validate("{{ . }}", {{ $val }}) {
			err := {{ err $valContext "value must be valid in " $r.CustomIn }}
			if !all { return err }
			errors = append(errors, err)
		}
		{{- end }}
	{{ end }}
`

/*
const custominTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.CustomIn }}
		{{ $val := accessor . }}
		{{ $valContext := . }}
		{{- range $r.CustomIn }}
		if val := {{ . }}({{ $val }}); val == nil {
			err := {{ err $valContext "value must be valid in " $r.CustomIn }}
			if !all { return err }
			errors = append(errors, err)
		}
		{{- end }}
	{{ end }}
`
*/
