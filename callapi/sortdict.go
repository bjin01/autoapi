package callapi

import (
	"fmt"
	"sort"
)

func sortdict(inputvars map[string]interface{}) []string {
	//var sorted_inputvars []interface{}
	var keylist []string

	for a, _ := range inputvars {
		keylist = append(keylist, a)
		sort.Strings(keylist)
		fmt.Printf("see the sorted keylist: %v\n", keylist)

		fmt.Println("")
		for i := 0; i < len(keylist); i++ {
			fmt.Printf("in for loop %v, %v %v\n", i, keylist[i], inputvars[keylist[i]])
		}
	}
	return keylist
}
