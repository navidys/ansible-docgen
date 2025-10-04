## 📃 {{ .RoleName }}
{{if.RoleMetaInfo }}
{{ .RoleMetaInfo.GalaxyInfo.Description }}

|                   |     |
| ----------------- | --- |
| `Author`          | {{ .RoleMetaInfo.GalaxyInfo.Author }}            |
{{if .RoleMetaInfo.GalaxyInfo.Company -}}
| `Company`         | {{ .RoleMetaInfo.GalaxyInfo.Company }}           |
{{- end}}
| `License`         | {{ .RoleMetaInfo.GalaxyInfo.License }}           |
| `Ansible version` | {{ .RoleMetaInfo.GalaxyInfo.MinAnsibleVersion }} |

{{end}}
