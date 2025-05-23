Refs =>
https://go.dev/blog/maps

=> Concurrency¶

This statement declares a counter variable that is an anonymous struct containing a map and an embedded sync.RWMutex.

var counter = struct{
sync.RWMutex
m map[string]int
}{m: make(map[string]int)}
To read from the counter, take the read lock:

counter.RLock()
n := counter.m["some_key"]
counter.RUnlock()
fmt.Println("some_key:", n)
To write to the counter, take the write lock:

counter.Lock()
counter.m["some_key"]++
counter.Unlock()