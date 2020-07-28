package unique

import (
	"fmt"
	"reflect"
)

// Ints returns a unique subset of the interfaces of slice provided.
func UniqueList(input []interface{}) []interface{} {

	u := make([]interface{}, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		switch reflect.TypeOf(val).Kind() {
		case reflect.Slice:
			fmt.Printf("input slice is %v\n", len(input))
			for i := 0; i < len(input); i++ {
				fmt.Printf("aaaaaaaa: %T\n", input[i])
				switch v := input[i].(type) {

				case int:
					fmt.Printf("int case %v\n", v)
					y := int(v)
					if _, ok := m[y]; !ok {
						m[v] = true
						u = append(u, v)
					}
				case []int:
					//fmt.Printf("[]int case %v\n", v)
					for z := 0; z < len(v); z++ {
						fmt.Printf("int case %v\n", v[z])
						/* y := int(v[z])
						if _, ok := m[y]; !ok {
							m[v[z]] = true
							u = append(u, v)
						} */
					}
				}

			}
		}
	}
	fmt.Printf("in uniquelist %v\n", u)
	return u
}
