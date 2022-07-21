# Go in-memory cache

## Installation

```
go get -u github.com/EmirShimshir/inMemoryCache
```

## Description
Go in-memory cache helps you store data of different types in a cache and access them by key.

```inMemoryCache.New()``` - constructor for new cache

```Set(key string, value interface{}) error``` - write the __value__ to the cache by the __key__. The __key__ must not be empty.

```Get(key string) (interface{}, error)``` - get the __value__ from the cache by the __key__. The __key__ must not be empty and the __value__ must exist.

```Delete(key) error``` - delete the __value__ from the cache by the __key__. The __key__ must not be empty and the __value__ must exist.

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/EmirShimshir/inMemoryCache"
)

func main() {
	cache := inMemoryCache.New()

	err := cache.Set("1", 42)
	if err != nil {
		log.Println(err.Error())
		return
	}

	value, err := cache.Get("1")
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Printf("value: %v\n", value)

	err = cache.Delete("1")
	if err != nil {
		log.Println(err.Error())
		return
	}
}

```
