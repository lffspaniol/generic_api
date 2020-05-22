package storage

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/lffspaniol/generic_api/gen/pb-go/gen"
)

var queue = []*gen.Temperature{}

func Add(time *timestamp.Timestamp, temp float32) {
	st := gen.Temperature{Data: time, Temp: temp}
	queue = append(queue, &st)
	if len(queue) > 5 {
		queue = remove(queue, 0)
	}
	return
}

func Get() []*gen.Temperature {
	return queue
}

func remove(slice []*gen.Temperature, s int) []*gen.Temperature {
	return append(slice[:s], slice[s+1:]...)
}
