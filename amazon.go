package main

import "fmt"

//Design and implement a data structure for Least Recently Used (LRU) cache.

//LRUCache cache = new LRUCache( 2 /* capacity */ );

//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.put(4, 4);    // evicts key 1
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4
//cache.put(5, 5)     //

type cache struct {
	l         int
	oldest    []int
	cacheSize int
	m         map[int]int
}

func main() {

	myCache := New(2)
	myCache.put(1, 1)
	myCache.put(2, 2)
	myCache.get(1)    // returns 1
	myCache.put(3, 3) // evicts key 2
	myCache.get(2)    // returns -1 (not found)
	myCache.put(4, 4) // evicts key 1
	myCache.get(1)    // returns -1 (not found)
	myCache.get(3)    // returns 3
	myCache.get(4)    // returns 4

}

func New(cacheSize int) *cache {

	return &cache{0, make([]int, 2), cacheSize, make(map[int]int)}

}

func (c *cache) check(key int, method string) {

	if method == "put" && c.m[key] != -1 {

		c.l++

	}

	if c.l > c.cacheSize {

		val := c.oldest[0]

		c.m[val] = -1

		c.l--

	}

	if method == "put" {

		if c.oldest[1] != key {
			c.oldest[0] = c.oldest[1]
		}

		c.oldest[1] = key

	} else if c.m[key] != -1 {

		if c.oldest[1] != key {
			c.oldest[0] = c.oldest[1]
		}

		c.oldest[1] = key

	}

}

func (c *cache) get(key int) int {

	fmt.Print("get ")
	fmt.Println(key)

	c.check(key, "get")

	fmt.Println(c.oldest)
	fmt.Println(c.m)
	fmt.Println(c.m[key])
	fmt.Println()

	return c.m[key]

}

func (c *cache) put(key int, value int) {

	fmt.Print("put ")
	fmt.Println(key)

	c.check(key, "put")

	c.m[key] = value

	fmt.Println(c.oldest)
	fmt.Println(c.m)
	fmt.Println()

}
