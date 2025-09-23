package schema

type Property struct {
    Type string `json:"type"`
}

type Mappings struct {
    Properties map[string]Property `json:"properties"`
}

type Schema struct {
    Mappings Mappings `json:"mappings"`
}