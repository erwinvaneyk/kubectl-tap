package version

import "encoding/json"

type Info struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"buildDate"`
}

func (i Info) String() string {
	bs, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return string(bs)
}