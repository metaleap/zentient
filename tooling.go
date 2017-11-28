package z

type Tool struct {
	Name      string
	Installed bool
	Website   string
}

type Tools []*Tool

func (me Tools) ByName(name string) *Tool {
	if name != "" {
		for _, tool := range me {
			if tool.Name == name {
				return tool
			}
		}
	}
	return nil
}
