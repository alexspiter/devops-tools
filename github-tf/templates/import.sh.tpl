#!/bin/bash
terraform import github_team.{{.Team.Slug}} {{$.Team.ID}}

{{- range $index, $user := .SortedUsers}} 
terraform import module.{{$.Team.Slug}}.github_team_membership.members[{{$index}}] {{$.Team.ID}}:{{$user}}
{{- end -}}