package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
1. create a struct Monster with some attributes
2. Serialize Monster's data by declaring a function Store() and save into a file
3. Load from the file. Deserialize Monster's data by declaring a function ReStore()
4. Testing function Store() and ReStore() with a testing file store_test.go.
You can define testing function name: TestScore and Testrestore
*/

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// serialize
func (m *Monster) Store() bool {
	// serialize
	byte_data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Cannot do serialze, error=%v\n", err)
		return false
	}
	// save data in byte to a file
	save_path := "/home/xmzhao/mygit/GoLang/practice/example/test/monster/monster.ser"
	err = ioutil.WriteFile(save_path, byte_data, 0666)
	if err != nil {
		fmt.Println("write failed...")
		return false
	}
	return true
}

// deserialize
func (m *Monster) Restore() bool {
	load_path := "/home/xmzhao/mygit/GoLang/practice/example/test/monster/monster.ser"
	byte_data, err := ioutil.ReadFile(load_path)
	if err != nil {
		fmt.Printf("Cannot read, error=%v\n", err)
		return false
	}
	err = json.Unmarshal(byte_data, m)
	if err != nil {
		fmt.Printf("Cannot do deserialze, error=%v\n", err)
		return false
	}
	return true
}
