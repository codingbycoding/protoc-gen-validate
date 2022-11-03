package goshared

const numTpl = `
	{{ if .Rules.GetIgnoreEmpty }}
		if {{ accessor . }} != 0 {
	{{ end }}

	{{ template "const" . }}
	{{ template "ltgt" . }}
	{{ template "in" . }}
	{{ template "custom_in" . }}
	{{ template "custom_in_csv" . }}
	{{ template "custom_in_func" . }}

	{{ if .Rules.GetIgnoreEmpty }}
		}
	{{ end }}

`
