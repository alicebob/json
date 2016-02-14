// Experimental JSON unmarshaler
package json

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// All dec* funcions eat their value + trailing whitespace. Can't have leading
// whitespace

var (
	ErrSyntax               = errors.New("unexpected char in json")
	ErrLeftoverBytes        = errors.New("leftover bytes")
	ErrNotAPointer          = errors.New("value is not a pointer")
	ErrUnsupportedType      = errors.New("unsupported type")
	ErrRawMessageNilPointer = errors.New("RawMessage: UnmarshalJSONs on nil pointer")
)

/*
type Unmarshaler interface {
	UnmarshalJSONs(string) error
}
*/

func Decode(s string, t interface{}) error {
	rv := reflect.ValueOf(t)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return ErrNotAPointer
	}
	s = skipWhitespace(s)
	rest, err := decValue(s, rv.Elem())
	if err != nil {
		return err
	}
	if len(rest) != 0 {
		return ErrLeftoverBytes
	}
	return nil
}

var rawMessageType = reflect.TypeOf(RawMessage("")) // Can this be done easier?

func decValue(s string, v reflect.Value) (string, error) {
	/*
		if v.Addr().Type().NumMethod() > 0 {
			if u, ok := v.Addr().Interface().(Unmarshaler); ok {
				return decUnmarshaler(s, u)
			}
		}
	*/
	switch v.Kind() {
	case reflect.String:
		if v.Type() == rawMessageType {
			return decRaw(s, v)
		}
		return decString(s, v)
	case reflect.Int:
		return decInt(s, v)
	case reflect.Float64:
		return decFloat(s, v)
	case reflect.Bool:
		return decBool(s, v)
	case reflect.Struct:
		return decStruct(s, v)
	case reflect.Map:
		return decMap(s, v.Addr())
	case reflect.Slice:
		return decSlice(s, v.Addr())
	case reflect.Ptr:
		if strings.HasPrefix(s, "null") {
			return skipWhitespace(s[4:]), nil
		}
		pv := reflect.New(v.Type().Elem())
		s, err := decValue(s, pv.Elem())
		if err != nil {
			return s, err
		}
		v.Set(pv)
		return s, nil
	default:
		return s, ErrUnsupportedType
	}
}

// space returns the number of leading whitespace bytes
func space(s string) int {
	for i, v := range s {
		switch v {
		case ' ', '\t', '\n': // TODO: unicode.IsSpace
		default:
			return i
		}
	}
	return len(s)
}

// skipWhitespace trims leading whitespace
func skipWhitespace(s string) string {
	// TrimLeftFunc is slower than space(), but more correct.
	return strings.TrimLeftFunc(s, unicode.IsSpace)
	// return s[space(s):]
}

// decString reads /"..."\s*/
func decString(s string, v reflect.Value) (string, error) {
	ss, l, err := nextString(s)
	if err != nil {
		return "", err
	}
	if l == 0 {
		return "", ErrSyntax
	}
	v.SetString(ss)
	return skipWhitespace(s[l:]), nil
}

// decInt reads /\d\s*/
func decInt(s string, v reflect.Value) (string, error) {
	n, l, err := nextInt(s)
	if err != nil {
		return "", err
	}
	if l == 0 {
		return "", ErrSyntax
	}
	v.SetInt(n)
	return skipWhitespace(s[l:]), nil
}

// decFloat reads /<float>\s*/
func decFloat(s string, v reflect.Value) (string, error) {
	n, l, err := nextNumber(s)
	if err != nil {
		return "", err
	}
	if l == 0 {
		return "", ErrSyntax
	}
	v.SetFloat(n)
	return skipWhitespace(s[l:]), nil
}

// decBool reads /(true|false)\s*/
func decBool(s string, v reflect.Value) (string, error) {
	t, l, err := nextBool(s)
	if err != nil {
		return "", err
	}
	v.SetBool(t)
	return skipWhitespace(s[l:]), nil
}

// decStruct reads /{...}\s*/
func decStruct(s string, v reflect.Value) (string, error) {
	if len(s) == 0 {
		return s, ErrSyntax
	}
	if strings.HasPrefix(s, "null") {
		return skipWhitespace(s[4:]), nil
	}
	tags := cachedTypeFields(v.Type())
	if s[0] != '{' {
		return s, ErrSyntax
	}
	s = skipWhitespace(s[1:])
	for {
		if len(s) == 0 {
			return s, ErrSyntax
		}
		if s[0] == '}' {
			return skipWhitespace(s[1:]), nil
		}

		field, l, err := nextString(s)
		if err != nil {
			return s, err
		}
		s = skipWhitespace(s[l:])
		if len(s) == 0 {
			return s, ErrSyntax
		}
		if s[0] != ':' {
			return s, ErrSyntax
		}
		s = skipWhitespace(s[1:])
		// find the tag
		found := false
		for _, t := range tags {
			if t.name == field {
				s, err = decValue(s, v.FieldByIndex(t.index))
				if err != nil {
					return s, err
				}
				found = true
				break
			}
		}
		if !found {
			_, s, err = decSkip(s)
			if err != nil {
				return s, err
			}
		}
		if len(s) == 0 {
			return s, ErrSyntax
		}
		if s[0] == ',' {
			s = skipWhitespace(s[1:])
			continue
		}
	}
	return "", nil
}

// decMap reads /{...}\s*/
func decMap(s string, v reflect.Value) (string, error) {
	if len(s) == 0 {
		return s, ErrSyntax
	}
	if strings.HasPrefix(s, "null") {
		return skipWhitespace(s[4:]), nil
	}
	if s[0] != '{' {
		return s, ErrSyntax
	}
	s = skipWhitespace(s[1:])
	for {
		if len(s) == 0 {
			return s, ErrSyntax
		}
		if s[0] == '}' {
			return skipWhitespace(s[1:]), nil
		}

		field, l, err := nextString(s)
		if err != nil {
			return s, err
		}
		s = skipWhitespace(s[l:])
		if len(s) == 0 {
			return s, ErrSyntax
		}
		if s[0] != ':' {
			return s, ErrSyntax
		}
		s = skipWhitespace(s[1:])

		value := reflect.New(v.Type().Elem().Elem())
		s, err = decValue(s, value.Elem())
		if err != nil {
			return s, err
		}
		if v.Elem().IsNil() {
			v.Elem().Set(reflect.MakeMap(v.Elem().Type()))
		}
		key := reflect.New(reflect.TypeOf(""))
		key.Elem().SetString(field)
		v.Elem().SetMapIndex(key.Elem(), value.Elem())

		if len(s) == 0 {
			return s, ErrSyntax
		}
		if s[0] == ',' {
			s = skipWhitespace(s[1:])
			continue
		}
	}
	return "", nil
}

// decSlice reads /\[...,...,...\]\s*/
func decSlice(s string, v reflect.Value) (string, error) {
	// v needs to be pointer to a slice.
	if len(s) == 0 {
		return s, ErrSyntax
	}
	if strings.HasPrefix(s, "null") {
		return skipWhitespace(s[4:]), nil
	}
	if s[0] != '[' {
		return s, ErrSyntax
	}
	s = skipWhitespace(s[1:])
	if len(s) == 0 {
		return s, ErrSyntax
	}
	if s[0] == ']' {
		return skipWhitespace(s[1:]), nil
	}
	var (
		err error
		vp  = v.Elem()
	)
	for {
		vp, _, _ = grow(vp, 1)
		s, err = decValue(s, vp.Index(vp.Len()-1).Addr().Elem())
		if err != nil {
			return s, err
		}
		if len(s) == 0 {
			return s, ErrSyntax
		}
		if s[0] == ']' {
			v.Elem().Set(vp)
			return skipWhitespace(s[1:]), nil
		}
		if s[0] != ',' {
			return s, ErrSyntax
		}
		s = skipWhitespace(s[1:])
		if len(s) == 0 {
			return s, ErrSyntax
		}
	}
	return "", nil
}

/*
func decUnmarshaler(s string, u Unmarshaler) (string, error) {
	l, err := lenNext(s)
	if err := u.UnmarshalJSONs(s[:l]); err != nil {
		return s, err
	}
	return skipWhitespace(s[l:]), err
}
*/
func decRaw(s string, v reflect.Value) (string, error) {
	l, err := lenNext(s)
	v.SetString(s[:l])
	return skipWhitespace(s[l:]), err
}

// decSkip reads the next value, and returns the raw json. Any which type,
// recursively.
func decSkip(s string) (string, string, error) {
	l, err := lenNext(s)
	return s[:l], skipWhitespace(s[l:]), err
}

func lenNext(s string) (int, error) {
	if len(s) == 0 {
		return 0, ErrSyntax
	}
	v := s[0]
	switch {
	case (v >= '0' && v <= '9') || v == '-':
		_, l, err := nextNumber(s)
		return l, err
	case v == '"':
		_, l, err := nextString(s)
		return l, err
	case v == 't' || v == 'f':
		_, l, err := nextBool(s)
		return l, err
	case v == 'n':
		return lenNextNull(s)
	case v == '[':
		return lenNextArray(s)
	case v == '{':
		return lenNextObject(s)
	default:
		return len(s), ErrSyntax
	}
}

// nextString returns the string starting at s, and the number of bytes read.
func nextString(s string) (string, int, error) {
	if s == "" {
		return s, 0, ErrSyntax
	}
	if strings.HasPrefix(s, "null") {
		return "", 4, nil
	}
	if s[0] != '"' {
		return s, 0, ErrSyntax
	}
	hasEscapes := false
	for i := 1; i < len(s); i++ {
		if s[i] == '\\' {
			// just skip whatever comes next. Will skip any escaped `"`.
			i++
			hasEscapes = true
			continue
		}
		if s[i] == '"' {
			res := s[1:i]
			if hasEscapes {
				var ok bool
				res, ok = unquote([]byte(res))
				if !ok {
					return s, 0, ErrSyntax
				}
			}
			return res, i + 1, nil
		}
	}
	return "", 0, ErrSyntax
}

// nextNumber reads the number found in the beginning of s. Returns the number
// and the number of bytes read. If there is no number it'll return length 0.
func nextNumber(s string) (float64, int, error) {
	if strings.HasPrefix(s, "null") {
		return 0.0, 4, nil
	}
	var l = 0
	for i, v := range s {
		if v >= '0' && v <= '9' || v == '.' || v == '-' || v == '+' || v == 'e' || v == 'E' {
			l = i + 1
			continue
		}
		break
	}
	if l == 0 {
		return 0, 0, nil
	}
	f, err := strconv.ParseFloat(s[:l], 64)
	if err != nil {
		return 0, l, ErrSyntax
	}
	return f, l, nil
}

// nextInt reads the int found in the beginning of s. Returns the number
// and the number of bytes read. If there is no number it'll return length 0.
// This is a subset of what we allow for nextNumber.
func nextInt(s string) (int64, int, error) {
	if len(s) == 0 {
		return 0, 0, nil
	}
	if strings.HasPrefix(s, "null") {
		return 0, 4, nil
	}
	var (
		i      = 0
		negate = false
	)
	if s[i] == '-' {
		negate = true
		i++
		if len(s) == i {
			return 0, 0, ErrSyntax
		}
	}
	var n int64
	for ; i < len(s); i++ {
		v := s[i]
		if v >= '0' && v <= '9' {
			n = n*10 + int64(v-'0')
			continue
		}
		// Do we need to support sci notation?
		break
	}
	if negate {
		n = -n
	}
	return n, i, nil
}

func nextBool(s string) (bool, int, error) {
	if strings.HasPrefix(s, "true") {
		return true, 4, nil
	}
	if strings.HasPrefix(s, "false") {
		return false, 5, nil
	}
	if strings.HasPrefix(s, "null") {
		return false, 4, nil
	}
	return false, len(s), ErrSyntax
}

func lenNextNull(s string) (int, error) {
	if strings.HasPrefix(s, "null") {
		return 4, nil
	}
	return len(s), ErrSyntax
}

func lenNextArray(s string) (int, error) {
	i := 1
	i += space(s[i:])
	if i == len(s) {
		return i, ErrSyntax
	}
	if s[i] == ']' {
		return i + 1, nil
	}
	for i < len(s) {
		si, err := lenNext(s[i:])
		if err != nil {
			return 0, err
		}
		i += si
		i += space(s[i:])
		if i == len(s) {
			return i, ErrSyntax
		}
		if s[i] == ']' {
			return i + 1, nil
		}
		if s[i] != ',' {
			return i, ErrSyntax
		}
		i += 1 // the ,
		i += space(s[i:])
	}
	return len(s), ErrSyntax
}

func lenNextObject(s string) (int, error) {
	i := 1
	for i < len(s) {
		i += space(s[i:])
		if s[i] == '}' {
			return i + 1, nil
		}
		_, ssi, err := nextString(s[i:])
		if err != nil {
			return 0, err
		}
		i += ssi
		i += space(s[i:])
		if s[i] != ':' {
			return i, ErrSyntax
		}
		i += 1
		i += space(s[i:])
		si, err := lenNext(s[i:])
		if err != nil {
			return 0, err
		}
		i += si
		i += space(s[i:])
		if s[i] == '}' {
			return i + 1, nil
		}
		if s[i] != ',' {
			return i, ErrSyntax
		}
		i += 1 // the ,
	}
	return len(s), nil
}

type RawMessage string

/*
func (m *RawMessage) UnmarshalJSONs(s string) error {
	if m == nil {
		return ErrRawMessageNilPointer
	}
	*m = RawMessage(s)
	return nil
}
*/

// encoding.json Unmarshaler
func (m *RawMessage) UnmarshalJSON(b []byte) error {
	if m == nil {
		return ErrRawMessageNilPointer
	}
	*m = RawMessage(string(b))
	return nil
}

// encoding.json Marshaler
// MarshalJSON returns *m as the JSON encoding of m.
func (m *RawMessage) MarshalJSON() ([]byte, error) {
	return []byte(*m), nil
}

// var _ Unmarshaler = (*RawMessage)(nil)
