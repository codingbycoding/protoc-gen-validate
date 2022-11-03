package goshared

const custominfuncTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.CustomInFunc }}
		{{ $val := accessor . }}
		{{ $valContext := . }}
		{{- range $r.CustomInFunc }}
		if !{{ . }}({{ $val }}) {
			err := {{ err $valContext "value must be valid in func " $r.CustomInFunc }}
			if !all { return err }
			errors = append(errors, err)
		}
		{{- end }}
	{{ end }}
`
