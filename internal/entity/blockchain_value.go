package entity

import "time"

type BlockchainValue struct {
	ID        int
	Value     int64
	CreatedAt time.Time
}
