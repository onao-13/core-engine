package scenario

// Region node infos

type PersonReplicaInfo struct {
	Person  string `json:"person"`
	Asset   string `json:"asset"`
	Replica string `json:"replica"`
}

type EnvironmentInfo struct {
	Background string `json:"background"`
	Music      string `json:"music"`
}

// End node infos

// Region node actions

type Action struct {
	ChangeEns *ActionChangeEns
}

type ActionChangeEns struct {
	NewValue string `json:"newValue"`
	Key      string `json:"type"`
	File     string `json:"ensFile"`
}

// End node actions

type Node struct {
	PersonInfo      *PersonReplicaInfo `json:"personInfo,omitempty"`
	EnvironmentInfo *EnvironmentInfo   `json:"environmentInfo,omitempty"`
	Action          *Action            `json:"action,omitempty"`
}
