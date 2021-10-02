package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	n := 0

	fmt.Scan(&n)
	data := []byte{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		data = append(data, scanner.Bytes()...)
	}

	obj := []interface{}{}

	json.Unmarshal(data, &obj)

	count := map[string]bool{}
	for _, o := range obj {
		if o == nil || o.(map[string]interface{}) == nil {
			continue
		}

		outer := o.(map[string]interface{})
		if v, found := outer["data"]; found {
			if v.(map[string]interface{}) == nil {
				continue
			}

			if key, found := v.(map[string]interface{})["key"]; found {
				count[key.(string)] = true
			}

		}
	}

	fmt.Println(len(count))
}
