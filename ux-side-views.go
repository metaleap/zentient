package z

type TreeItem struct {
	ID           string        `json:"id,omitempty"`
	Label        string        `json:"label,omitempty"`
	IconPath     string        `json:"iconPath,omitempty"`
	Tooltip      string        `json:"tooltip,omitempty"`
	Command      *EditorAction `json:"command,omitempty"`
	ContextValue string        `json:"contextValue,omitempty"`
}

type TreeView struct {
}
