Here's the corrected Go code from your screenshot. The main issues were:

Package-level variable declaration syntax (needs var or assignment inside main())

Incorrect channel usage (didn't close the channel, potentially causing deadlocks)

Better goroutine synchronization


https://www.youtube.com/watch?v=RGOqCzdI3wQ