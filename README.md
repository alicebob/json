Experimental Go JSON parser


# status

don't use this

# features

Supported types to decode into:
    
    * int
    * float64
    * strings
    * structs made of those
    * slices made of those
    * pointers to all of those
    * RawMessage

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


