package pkg

import "fmt"

func PrintMap(expectMap []map[string]any) {
	for _, m := range expectMap {

		// m is a map[string]interface.
		// loop over keys and values in the map.
		for k, v := range m {
			fmt.Println(k, "value is", v)
		}
	}
}
