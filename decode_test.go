package json

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	type testcase struct {
		raw  string
		want string
		left string
		err  error
	}
	for _, cas := range []testcase{
		{raw: `"foo"`, want: "foo"},
		{raw: "\"foo\nbar\"", want: "foo\nbar"},
		{raw: `""`, want: ""},
		{raw: ``, err: ErrSyntax},
		{raw: `"foo`, err: ErrSyntax},
		{raw: `"foo",`, want: "foo", left: `,`},
		{raw: `"foo"  `, want: "foo"},
		{raw: "\"foo\"\t\n ,", want: "foo", left: `,`},
		{raw: `"☃"`, want: "☃"},
		{raw: `"☃☃☃"`, want: "☃☃☃"},
		{raw: `"\\"`, want: `\`},
		{raw: `"\""`, want: `"`},
		{raw: `"\"\u2603\""`, want: `"☃"`},
		{raw: `"\"\u2603☃\""`, want: `"☃☃"`},
		{raw: `"\n\t\r"`, want: "\n\t\r"},
		{raw: `"foo\`, err: ErrSyntax},
		{raw: `"\"\u26""`, err: ErrSyntax},
		{raw: `null`, want: ""},
	} {
		dec := ""
		left, err := decString(cas.raw, reflect.ValueOf(&dec).Elem())
		if have, want := err, cas.err; have != want {
			t.Fatalf("%q: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := left, cas.left; have != want {
				t.Fatalf("have %q, want %q", have, want)
			}
			if have, want := dec, cas.want; have != want {
				t.Fatalf("have %q, want %q", have, want)
			}
		}
	}
}

func TestNextNumber(t *testing.T) {
	type testcase struct {
		raw  string
		len  int
		want float64
		err  error
	}
	for _, cas := range []testcase{
		{raw: ``, want: 0, len: 0},
		{raw: `123`, want: 123, len: 3},
		{raw: `0`, want: 0, len: 1},
		{raw: `foo`, len: 0},
		{raw: `123,`, want: 123, len: 3},
		{raw: `123.0`, want: 123.0, len: 5},
		{raw: `123.1`, want: 123.1, len: 5},
		{raw: `123.9`, want: 123.9, len: 5},
		{raw: `0.1`, want: 0.1, len: 3},
		{raw: `0.0`, want: 0, len: 3},
		{raw: `-1`, want: -1, len: 2},
		{raw: `-1.2`, want: -1.2, len: 4},
		// {raw: `+1.2`, len: 0},
		{raw: `123.123`, want: 123.123, len: 7},
		{raw: `123e10`, want: 123e10, len: 6},
		{raw: `0e10`, want: 0, len: 4},
		{raw: `123e`, err: ErrSyntax},
		{raw: `123e,`, err: ErrSyntax},
		{raw: `123E10`, want: 123E10, len: 6},
		{raw: `123f10`, want: 123, len: 3},
		{raw: `123e-10`, want: 123e-10, len: 7},
		{raw: `123e -10`, err: ErrSyntax},
		{raw: `123e+10`, want: 123e10, len: 7},
		{raw: `123.123e10`, want: 123.123e10, len: 10},
		{raw: `123 456`, want: 123, len: 3},
		{raw: `null`, want: 0.0, len: 4},
		// {raw: `00`, err: ErrSyntax},
	} {
		n, len, err := nextNumber(cas.raw)
		if have, want := err, cas.err; have != want {
			t.Fatalf("%q: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := len, cas.len; have != want {
				t.Fatalf("%q left: have %d, want %d", cas.raw, have, want)
			}
			if have, want := n, cas.want; have != want {
				t.Fatalf("%q want: have %d, want %d", cas.raw, have, want)
			}
		}
	}
}

func TestInt(t *testing.T) {
	type testcase struct {
		raw  string
		left string
		want int64
		err  error
	}
	for _, cas := range []testcase{
		{raw: `123`, want: 123},
		{raw: `0`, want: 0},
		{raw: `foo`, err: ErrSyntax},
		{raw: `123,`, want: 123, left: `,`},
		{raw: `123  `, want: 123},
		{raw: `123  ,`, want: 123, left: `,`},
		{raw: `123.9`, want: 123, left: `.9`},
		{raw: `0.1`, want: 0, left: `.1`},
		{raw: `-1`, want: -1},
		{raw: `-`, err: ErrSyntax},
		{raw: `123e10`, want: 123, left: `e10`},
		{raw: `123e,`, want: 123, left: `e,`},
		{raw: `null`, want: 0, left: ``},
	} {
		dec := int64(0)
		left, err := decInt(cas.raw, reflect.ValueOf(&dec).Elem())
		if have, want := err, cas.err; have != want {
			t.Fatalf("%q: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := left, cas.left; have != want {
				t.Fatalf("%q left: have %q, want %q", cas.raw, have, want)
			}
			if have, want := dec, cas.want; have != want {
				t.Fatalf("%q want: have %d, want %d", cas.raw, have, want)
			}
		}
	}
}

func TestFloat(t *testing.T) {
	type testcase struct {
		raw  string
		left string
		want float64
		err  error
	}
	for _, cas := range []testcase{
		{raw: `123`, want: 123},
		{raw: `0`, want: 0},
		{raw: `foo`, err: ErrSyntax},
		{raw: `123,`, want: 123, left: `,`},
		{raw: `123  `, want: 123},
		{raw: `123  ,`, want: 123, left: `,`},
		{raw: `123.0`, want: 123},
		{raw: `123.1`, want: 123.1},
		{raw: `123.9`, want: 123.9},
		{raw: `0.1`, want: 0.1},
		{raw: `-1`, want: -1},
		{raw: `123e10`, want: 123e10},
		{raw: `123e,`, err: ErrSyntax},
		{raw: `null`, want: 0.0},
	} {
		dec := float64(0)
		left, err := decFloat(cas.raw, reflect.ValueOf(&dec).Elem())
		if have, want := err, cas.err; have != want {
			t.Fatalf("%q: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := left, cas.left; have != want {
				t.Fatalf("%q left: have %q, want %q", cas.raw, have, want)
			}
			if have, want := dec, cas.want; have != want {
				t.Fatalf("%q want: have %f, want %f", cas.raw, have, want)
			}
		}
	}
}

func TestBool(t *testing.T) {
	type testcase struct {
		raw  string
		left string
		want bool
		err  error
	}
	for _, cas := range []testcase{
		{raw: `true`, want: true},
		{raw: `false`, want: false},
		{raw: `0`, err: ErrSyntax},
		{raw: `TRUE`, err: ErrSyntax},
		{raw: `FALSE`, err: ErrSyntax},
		{raw: `false?!?`, want: false, left: `?!?`},
		{raw: `null`, want: false},
	} {
		dec := false
		left, err := decBool(cas.raw, reflect.ValueOf(&dec).Elem())
		if have, want := err, cas.err; have != want {
			t.Fatalf("%q: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := left, cas.left; have != want {
				t.Fatalf("%q left: have %q, want %q", cas.raw, have, want)
			}
			if have, want := dec, cas.want; have != want {
				t.Fatalf("%q want: have %b, want %b", cas.raw, have, want)
			}
		}
	}
}

func TestStruct(t *testing.T) {
	type str struct {
		Foo string `json:"foo"`
	}
	type testcase struct {
		raw  string
		left string
		want str
		err  error
	}
	for _, cas := range []testcase{
		{raw: `{"foo":"bar"}`, want: str{Foo: "bar"}},
		{raw: `{}`, want: str{}},
		{raw: `{},`, want: str{}, left: ","},
		{raw: `{ },`, want: str{}, left: ","},
		{raw: `{ "foo" : "bar" }   ,`, want: str{Foo: "bar"}, left: ","},
		{raw: `{ "foo" : "bar", "unknown": 12 }   ,`, want: str{Foo: "bar"}, left: ","},
		{raw: `{"foo" "bar"}`, err: ErrSyntax},
		{raw: `{"foo" : "bar" 12}`, err: ErrSyntax},
		{raw: `{`, err: ErrSyntax},
		{raw: `{"foo"`, err: ErrSyntax},
		{raw: `{"foo" :`, err: ErrSyntax},
		{raw: `{"foo" : "123`, err: ErrSyntax},
		{raw: `{"foo" : "123"`, err: ErrSyntax},
		{raw: `{"foo" : "123" "123"`, err: ErrSyntax},
		{raw: `{"foo" "1": "123"}`, err: ErrSyntax},
		{raw: `{foo: "123"}`, err: ErrSyntax},
		{raw: `{1: "123"}`, err: ErrSyntax},
		{raw: `{[]: "123"}`, err: ErrSyntax},
		{raw: `{"foo": 123}`, err: ErrSyntax},
		{raw: `{"foo": [1,2,3]}`, err: ErrSyntax},
		{raw: `null`, want: str{}},
	} {
		v := &str{}
		left, err := decStruct(cas.raw, reflect.ValueOf(v).Elem())
		if have, want := err, cas.err; have != want {
			t.Fatalf("error in %s: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := left, cas.left; have != want {
				t.Fatalf("leftover in %s: have %q, want %q", cas.raw, have, want)
			}
			if have, want := *v, cas.want; have != want {
				t.Fatalf("dec in %s: have %v, want %v", cas.raw, have, want)
			}
		}
	}
}

func TestSlice(t *testing.T) {
	type sl []int
	type testcase struct {
		raw  string
		left string
		want sl
		err  error
	}
	for _, cas := range []testcase{
		{raw: `[1,2,3]`, want: sl{1, 2, 3}},
		{raw: `[1,2,3],`, want: sl{1, 2, 3}, left: ","},
		{raw: `[ 1 , 2 , 3 ]   ,`, want: sl{1, 2, 3}, left: ","},
		{raw: `[]`, want: sl{}},
		{raw: `[ ]`, want: sl{}},
		{raw: `[1 3]`, err: ErrSyntax},
		{raw: `[1,`, err: ErrSyntax},
		{raw: `[1`, err: ErrSyntax},
		{raw: `[1,]`, want: sl{1}}, // should be a ErrSyntax
		{raw: `null`, want: sl{}},
	} {
		v := &sl{}
		left, err := decSlice(cas.raw, reflect.ValueOf(v))
		if have, want := err, cas.err; have != want {
			t.Fatalf("error in %s: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := left, cas.left; have != want {
				t.Fatalf("leftover in %s: have %q, want %q", cas.raw, have, want)
			}
			if have, want := *v, cas.want; !reflect.DeepEqual(have, want) {
				t.Fatalf("dec in %s: have %#v/%T, want %#v/%T", cas.raw, have, have, want, want)
			}
		}
	}
}

func TestSkip(t *testing.T) {
	type testcase struct {
		raw  string
		left string
		want string
		err  error
	}
	for _, cas := range []testcase{
		{raw: `123`, want: `123`},
		{raw: `123,`, want: `123`, left: `,`},
		{raw: `123]`, want: `123`, left: `]`},
		{raw: `123}`, want: `123`, left: `}`},
		{raw: `123  ,`, want: `123`, left: `,`},
		{raw: `-123.0e-3  ,`, want: `-123.0e-3`, left: `,`},
		{raw: `"aap"`, want: `"aap"`},
		{raw: `"aap",`, want: `"aap"`, left: `,`},
		{raw: `[1,2,3]`, want: `[1,2,3]`},
		{raw: `[1,2,3] ,`, want: `[1,2,3]`, left: ","},
		{raw: `[  ]`, want: `[  ]`},
		{raw: `{"aap" : 1}`, want: `{"aap" : 1}`},
		{raw: `{"aap" : 1}, `, want: `{"aap" : 1}`, left: ", "},
		{raw: `{  },`, want: `{  }`, left: ","},
		{raw: `{"aap" : [1,2]}, `, want: `{"aap" : [1,2]}`, left: ", "},
		{raw: `{"aap" : true}`, want: `{"aap" : true}`},
		{raw: `{"aap" : false }`, want: `{"aap" : false }`},
		{raw: `null`, want: `null`},
	} {
		v, left, err := decSkip(cas.raw)
		if have, want := err, cas.err; have != want {
			t.Fatalf("error in %s: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := v, cas.want; have != want {
				t.Fatalf("value in %s: have %q, want %q", cas.raw, have, want)
			}
			if have, want := left, cas.left; have != want {
				t.Fatalf("leftover in %s: have %q, want %q", cas.raw, have, want)
			}
		}
	}
}

func TestRawMessage(t *testing.T) {
	type testcase struct {
		raw  string
		left string
		want RawMessage
		err  error
	}
	for _, cas := range []testcase{
		{raw: `"aap"`, want: RawMessage(`"aap"`)},
		{raw: `[1,2,3]`, want: RawMessage(`[1,2,3]`)},
		{raw: `null`, want: RawMessage(`null`)},
	} {
		v := RawMessage("")
		left, err := decRaw(cas.raw, reflect.ValueOf(&v).Elem())
		if have, want := err, cas.err; have != want {
			t.Fatalf("error in %s: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := left, cas.left; have != want {
				t.Fatalf("leftover in %s: have %q, want %q", cas.raw, have, want)
			}
			if have, want := v, cas.want; !reflect.DeepEqual(have, want) {
				t.Fatalf("dec in %s: have %#v/%T, want %#v/%T", cas.raw, have, have, want, want)
			}
		}
	}
}

func TestDecode(t *testing.T) {
	type Str struct {
		Foo string     `json:"foo"`
		Bar int        `json:"bar"`
		Baz string     `json:"baz"`
		Fl  float64    `json:"fl"`
		Bl  bool       `json:"bl"`
		Raw RawMessage `json:"raw"`
	}
	type Strsl []Str
	type Pointers struct {
		Structp *Str    `json:"structp"`
		Stringp *string `json:"stringp"`
		Intp    *int    `json:"intp"`
	}

	type testcase struct {
		raw  string
		ptr  interface{}
		want interface{}
		err  error
	}
	nineninenine := 999
	helloworld := "hello world"
	for _, cas := range []testcase{
		// strings
		{raw: `"foo"`, ptr: new(string), want: "foo"},
		{raw: `""`, ptr: new(string), want: ""},

		// ints
		{raw: `987`, ptr: new(int), want: 987},
		{raw: `0`, ptr: new(int), want: 0},

		// floats
		{raw: `987.6`, ptr: new(float64), want: 987.6},
		{raw: `-3.14e-100`, ptr: new(float64), want: -3.14e-100},

		// bools
		{raw: `true`, ptr: new(bool), want: true},

		// structs
		{raw: `{}`, ptr: new(Str), want: Str{}},
		{
			raw: `{"foo":"foovalue","bar":123,"baz":"bazvalue","fl":1.2,"bl":true,"raw":[1,2,3]}`,
			ptr: new(Str),
			want: Str{
				Foo: "foovalue",
				Bar: 123,
				Baz: "bazvalue",
				Fl:  1.2,
				Bl:  true,
				Raw: `[1,2,3]`,
			},
		},
		{raw: `{"bar":}`, ptr: new(Str), err: ErrSyntax},

		// slices
		// {raw: `[]`, ptr: new(Strsl), want: Strsl(nil)}, // encoding.json decodes this as a 0 length slice, not a nil
		{
			raw: `[{"foo":"foovalue","bar":123,"baz":"bazvalue"},{"foo":"foovalue2","bar":124,"baz":"bazvalue2"}]`,
			ptr: new(Strsl),
			want: Strsl{
				{
					Foo: "foovalue",
					Bar: 123,
					Baz: "bazvalue",
				},
				{
					Foo: "foovalue2",
					Bar: 124,
					Baz: "bazvalue2",
				},
			},
		},

		// pointers
		{
			raw: `{"structp":{"foo":"foovalue","bar":123,"baz":"bazvalue"},"intp":999,"stringp":"hello world"}`,
			ptr: new(Pointers),
			want: Pointers{
				Structp: &Str{
					Foo: "foovalue",
					Bar: 123,
					Baz: "bazvalue",
				},
				Intp:    &nineninenine,
				Stringp: &helloworld,
			},
		},

		// nulls
		{raw: `null`, ptr: new(string), want: ""},
		{raw: `null`, ptr: new(int), want: 0},
		{raw: `null`, ptr: new(float64), want: 0.0},
		{raw: `null`, ptr: new([]int), want: []int(nil)},
		{raw: `null`, ptr: new(Str), want: Str{}},
		{raw: `null`, ptr: new(*Str), want: (*Str)(nil)},
		{raw: `null`, ptr: new(*[]int), want: (*[]int)(nil)},

		// misc
		{raw: ` "foo"`, ptr: new(string), want: "foo"},
	} {
		v := reflect.New(reflect.TypeOf(cas.ptr).Elem())
		err := Decode(cas.raw, v.Interface())
		if have, want := err, cas.err; have != want {
			t.Fatalf("%q error: have %v, want %v", cas.raw, have, want)
		}
		if cas.err == nil {
			if have, want := v.Elem().Interface(), cas.want; !reflect.DeepEqual(have, want) {
				t.Fatalf("%q value: have %v, want %v", cas.raw, have, want)
			}
		}

		vthem := reflect.New(reflect.TypeOf(cas.ptr).Elem())
		err = json.Unmarshal([]byte(cas.raw), vthem.Interface())
		if cas.err != nil {
			if err == nil {
				t.Fatalf("expected an error: %q", cas.err)
			}
		}
		if cas.err == nil {
			if err != nil {
				t.Fatal(err)
			}
			if have, want := vthem.Elem().Interface(), cas.want; !reflect.DeepEqual(have, want) {
				t.Fatalf("error: have %v, want %v", have, want)
			}
		}
	}

}
