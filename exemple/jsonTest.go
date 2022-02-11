package exemple

import (
	"encoding/json"
	"fmt"
)

func JsonTest() {
	t := taskJ{Name: "BLA", DisplayName: "BLA2"}
	b, err := json.Marshal(&t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

}

type taskJ struct {
	Name        string
	DisplayName string `json:"name2"`
}

func (t *task) MarshalJSON() ([]byte, error) {
	type JTask task
	t.Name = "TOTO"
	return json.Marshal((*JTask)(t))
}
