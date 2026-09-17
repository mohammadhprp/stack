package model

type Skill struct {
	ID          string
	Name        string
	Description string
	Dir         string
	Files       []string
}

type MCPSpec struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Command     []string          `json:"command"`
	Env         map[string]string `json:"env"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
}

type MCP struct {
	Slug string
	MCPSpec
	Dir string
}

type Harness struct {
	ID             string
	Name           string
	SupportsSkills bool
	SkillsPath     string
	MCPConfigPath  string
}
