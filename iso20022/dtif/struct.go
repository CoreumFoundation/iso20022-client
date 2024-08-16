package dtif

import "time"

type Header struct {
	DTI             string `json:"DTI"`
	DTIType         int    `json:"DTIType"`
	TemplateVersion string `json:"templateVersion"`
	DLTType         int    `json:"DLTType,omitempty"`
}

type Record struct {
	Header Header `json:"Header"`
}

type JsonTime struct {
	time.Time
}

func (t *JsonTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}
