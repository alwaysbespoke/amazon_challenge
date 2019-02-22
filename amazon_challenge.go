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
	newest		int
	oldest		int
	l         	int
	lru    		[]int
	cacheSize 	int
	m         	map[int]int
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

	return &cache{0, 0, 0, make([]int, cacheSize), cacheSize, make(map[int]int)}

}

func (c *cache) check(key int, method string) {

	if method == "put" {

		if c.m[key] < 1 {

			c.l++

			if c.l > c.cacheSize {

				oldest := c.lru[c.oldest]

				c.m[oldest] = -1

				c.l--

			}

		}

		c.shuffle(key)

	} else if c.m[key] != -1 {

		c.shuffle(key)

	}

}

func (c *cache) shuffle(key int) {

	c.lru[c.newest] = key

	c.oldest = c.cacheSize + c.newest - 1

	if c.oldest >= c.cacheSize {

		c.oldest = c.oldest - c.cacheSize

	}
	/*
	fmt.Print("Newest_i ")
	fmt.Print(c.newest)
	fmt.Print(", Oldest_i ")
	fmt.Println(c.oldest)
	fmt.Print("Newest ")
	fmt.Print(c.lru[c.newest])
	fmt.Print(", Oldest ")
	fmt.Print(c.lru[c.oldest])
	fmt.Println()
	*/
	c.newest++

	if c.newest == c.cacheSize {

		c.newest = 0
	}

}

func (c *cache) get(key int) int {

	fmt.Print("get ")
	fmt.Print(key)

	c.check(key, "get")

	//fmt.Println(c.lru)
	//fmt.Println(c.m)
	fmt.Print(", returns -> ")
	fmt.Println(c.m[key])
	//fmt.Println()

	return c.m[key]

}

func (c *cache) put(key int, value int) {

	fmt.Print("put ")
	fmt.Println(key)

	c.check(key, "put")

	c.m[key] = value

	//fmt.Println(c.lru)
	//fmt.Println(c.m)
	//fmt.Println()

}
