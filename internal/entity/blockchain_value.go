package entity

import "time"

type BlockchainValue struct {
	ID        int
	Value     string
	CreatedAt time.Time
}
