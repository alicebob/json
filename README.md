Alternative Go JSON parser

# features

Supported types to unmarshal into:
    
 * int
 * float64
 * string
 * bool
 * RawMessage
 * structs made of all of these
 * slices made of all of these
 * pointers to all of these
 * map[string]{all of these}

Other basic types would be easy to add; I have not needed them yet. Note that the empty interface is not supported.

# vs. encoding/json

 * this works on a string, not a []byte
 * encoding/json has nice error messages
 * encoding/json can decode to the empty interface
 * encoding/json supports the Unmarshaler interface

# performance

For the use case I'm mostly interested in (parsing openRTB requests) I get
twice the performance and half the allocs:
<pre>
BenchmarkRTB-4       	   50000	     29093 ns/op	    2536 B/op	      39 allocs/op
BenchmarkRTBThem-4   	   20000	     63858 ns/op	    3888 B/op	      87 allocs/op
</pre>

Similar for parsing large ElasticSearch bulk results:
<pre>
BenchmarkESbulk-4    	    3000	    473237 ns/op	  177979 B/op	     830 allocs/op
BenchmarkESbulkThem-4	    2000	   1055068 ns/op	  200884 B/op	    2283 allocs/op
</pre>

# status

Works fine for us.

[![Build Status](https://travis-ci.org/alicebob/json.svg?branch=master)](https://travis-ci.org/alicebob/json)

[![GoDoc](https://godoc.org/github.com/alicebob/json?status.svg)](https://godoc.org/github.com/alicebob/json)
