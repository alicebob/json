package json_test

import (
	"fmt"

	"github.com/alicebob/json"
)

// Example_Unmarshal shows the basic unmarshal function.
func Example_Unmarshal() {
	var v = struct {
		Status string   `json:"status"`
		Code   int      `json:"code"`
		Items  []string `json:"items"`
	}{}
	if err := json.Decode(`{"status": "good", "code": 200, "items": ["red", "yellow", "blue"]}`, &v); err != nil {
		panic(err)
	}
	fmt.Printf("result: %#v\n", v)
	// Output: result: struct { Status string "json:\"status\""; Code int "json:\"code\""; Items []string "json:\"items\"" }{Status:"good", Code:200, Items:[]string{"red", "yellow", "blue"}}

}

// Example_Raw shows how to use RawMessage.
func Example_Raw() {
	var v = struct {
		Status string          `json:"status"`
		Code   int             `json:"code"`
		Items  json.RawMessage `json:"items"`
	}{}
	if err := json.Decode(`{"status": "good", "code": 200, "items": ["red", "yellow", "blue"]}`, &v); err != nil {
		panic(err)
	}
	fmt.Printf("result: %#v\n", v)
	// Output: result: struct { Status string "json:\"status\""; Code int "json:\"code\""; Items json.RawMessage "json:\"items\"" }{Status:"good", Code:200, Items:"[\"red\", \"yellow\", \"blue\"]"}
}
