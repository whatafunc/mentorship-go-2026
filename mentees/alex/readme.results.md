a ` BenchmarkMapInsert-8   	       6	 272.498.913 ns/op	65.653.118 B/op	  999903 allocs/op
`

b ` BenchmarkMapInsert-8   	      12	  83.719.662 ns/op	40.236.384 B/op	      20 allocs/op
`

c ` BenchmarkMapInsert-8   	      15	  89.436.956 ns/op	40.236.382 B/op	      20 allocs/op
`


🧠 Key insight:
#1 benchmark measures string creation cost, not map bucket quality.
#2 Buckets are worse, but:
    CPU is still faster than allocating 1M strings
    Collisions don’t dominate yet
    ⚠️ This is a dangerous win — fast but structurally poor.
        Poor low-bit entropy
        More overflow buckets

#3 This is the best engineered result - balanced solution

Correct conclusions (this is important)

Let’s lock them in clearly:

❌ Wrong conclusion
“String keys are slow because of buckets”
✅ Correct conclusion
“String keys are slow here because of allocation, not buckets”

❌ Wrong conclusion
“i<<12 is good because it’s fastest”
✅ Correct conclusion
“i<<12 is fast but structurally unsafe”

✅ Best conclusion
“(i<<12)^i gives me almost the same speed with much better distribution”
👏 That’s the winner.

Final comparison table (bookmark this)
Algorithm	    Burst	Sustained control	Complexity	Brute-force fit
Token bucket	✅	    ✅	                Low     	🏆 Best
Sliding window	⚠️	     ✅	                 High	     Good
Leaky bucket	❌	    ✅	                Medium	    Poor
Fixed window	❌	    ❌	                Low	        Bad


┌─────────────────────────────┐
│ Brute-force algorithm       │  ← Token bucket (logic)
├─────────────────────────────┤
│ State model                 │  ← tokens, timestamps, counters
├─────────────────────────────┤
│ Storage                     │  ← map / sync.Map / Redis
├─────────────────────────────┤
│ Key design                  │  ← IP, user, IP+user ✅
├─────────────────────────────┤
│ Hashing & buckets           │  ← Go runtime internals✅
└─────────────────────────────┘
