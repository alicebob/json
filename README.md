Alternative Go JSON parser


# status

Works fine for us.

# features

Supported types to decode into:
    
 * int
 * float64
 * string
 * bool
 * RawMessage
 * structs made of all of these
 * slices made of all of these
 * pointers to all of these
 * map[string]<all of these>

Other basic types would be easy to add; I have not needed them yet.

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


