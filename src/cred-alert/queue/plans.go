package queue

import "encoding/json"

type DiffScanPlan struct {
	Owner      string `json:"owner"`
	Repository string `json:"repository"`

	Start string `json:"start"`
	End   string `json:"end"`
}

func (p DiffScanPlan) Task() Task {
	payload, _ := json.Marshal(p)

	return basicTask{
		typee:   "diff-scan",
		payload: string(payload),
	}
}

type basicTask struct {
	typee   string
	payload string
}

func (t basicTask) Type() string {
	return t.typee
}

func (t basicTask) Payload() string {
	return t.payload
}