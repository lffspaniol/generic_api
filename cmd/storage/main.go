package main

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/lffspaniol/generic_api/internal/storage"
)

func main() {
	t := time.Now()
	ts, _ := ptypes.TimestampProto(t)
	storage.Add(ts, 32.5)
	storage.Add(ts, 34.5)
	storage.Add(ts, 35.5)
	storage.Add(ts, 36.5)
	storage.Add(ts, 37.5)
	storage.Add(ts, 38.5)
	storage.Add(ts, 39.5)
	storage.Add(ts, 40.5)
	vet := storage.Get()
	for i, v := range vet {
		fmt.Println("abc: ", v, " i:", i)
	}
}
