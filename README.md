# Go thread-safe in-memory cache with element lifetime.

## Installation

```
go get -u github.com/EmirShimshir/inMemoryCache
```

## Description
Go in-memory cache is goroutine-safe, it helps you store data of different types in a cache and access them by key. When adding a new element to the cache, specify the time after which the element will be unavailable.

```inMemoryCache.New()``` - constructor for new cache

```Set(key string, value any, ttl time.Duration) error``` - write the __value__ to the cache by the __key__. The __key__ must not be empty. The __ttl__ is time to live for the new element, it must not be zero.

```Get(key string) (interface{}, error)``` - get the __value__ from the cache by the __key__. The __key__ must not be empty and the __value__ must exist.

```Delete(key) error``` - delete the __value__ from the cache by the __key__. The __key__ must not be empty and the __value__ must exist.

## Example

```go
package main

import (
	"fmt"
	"github.com/EmirShimshir/inMemoryCache"
	"log"
	"time"
)

func main() {
	cache := inMemoryCache.New()

	err := cache.Set("userId", 42, 5 * time.Second)
	if err != nil { // err == nil
		log.Fatal(err)
	}

	userId, err := cache.Get("userId")
	if err != nil { // err == nil
		log.Fatal(err)
	}
	fmt.Println(userId) // Output: 42

	time.Sleep(6 * time.Second)

	userId, err = cache.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // <--
	}
}

```

```
42
2022/07/27 14:43:08 no value for the key userId
exit status 1
```
