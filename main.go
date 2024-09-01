package main

import (
	"fmt"

	"github.com/Cache/lfu_cache"
	// "github.com/Cache/lru_cache"
)

func main() {
	// cache := lru_cache.Initialise_Cache(5)
	// cache.Get("www.google.com")
	// cache.Display()
	// cache.Get("www.yahoo.com")
	// cache.Display()
	// cache.Get("www.gmail.com")
	// cache.Display()
	// cache.Get("www.chatgpt.com")
	// cache.Display()
	// cache.Get("www.slackboat.com")
	// cache.Display()
	// cache.Get("www.yahoo.com")
	// cache.Display()
	// cache.Get("www.google.com")
	// cache.Display()
	// cache.Get("www.yahoo.com")
	// cache.Display()
	// cache.Get("www.youtube.com")
	// cache.Display()
	// cache.Get("www.google.com")
	// cache.Display()
	fmt.Println("Testing the lfu based on priority que cache: \n")
	cache := lfu_cache.Initialise_Cache(5)
	fmt.Println(cache.Get("you"))
	cache.PrintArr()
	cache.Set("you", "youtube.com")
	fmt.Println(cache.Get("you"))
	cache.PrintArr()
	cache.Set("bing", "bing.com")
	fmt.Println(cache.Get("bing"))
	cache.PrintArr()
	cache.Set("isc", "isckon.com")
	fmt.Println(cache.Get("isc"))
	cache.PrintArr()
	cache.Set("you", "youtube1.com")
	fmt.Println(cache.Get("you"))
	cache.PrintArr()
	cache.Set("cd", "www.cdn.com")
	fmt.Println(cache.Get("cd"))
	cache.PrintArr()
	cache.Set("shoe", "www.paragon.com")
	fmt.Println(cache.Get("shoe"))
	cache.PrintArr()
	cache.Set("watch", "www.watch.com")
	fmt.Println(cache.Get("watch"))
	cache.PrintArr()
	fmt.Println(cache.Get("cd"))
	cache.PrintArr()
}
