package storage

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type storage struct {
	time *timestamp.Timestamp
	temp float32
}

func Add(time *timestamp.Timestamp, temp float32) error {
	return nil
}

func Get() ([]storage, error) {
	return nil, nil
}
