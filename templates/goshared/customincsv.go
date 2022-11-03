package goshared

const customincsvTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.CustomInCsv }}
		{{ $val := accessor . }}
		{{ $valContext := . }}
		{{- range $r.CustomInCsv }}
		if !validaterules.Validate("{{ . }}", {{ $val }}) {
			err := {{ err $valContext "value must be valid in csv " $r.CustomInCsv }}
			if !all { return err }
			errors = append(errors, err)
		}
		{{- end }}
	{{ end }}
`
