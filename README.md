# cache_with_ttl
```
go get github.com/dmytrodemianchuk/cache_with_ttl
```
```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dmytrodemianchuk/second"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second*5)
	userId, err := cache.Get("userId")
	if err != nil { // err == nil
		log.Fatal(err)
	}
	fmt.Println(userId) // Output: 42

	time.Sleep(time.Second * 6) // прошло 5 секунд

	userId = cache.Get("userId")
	userId, err = cache.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // сработает этот код
	}
}
```