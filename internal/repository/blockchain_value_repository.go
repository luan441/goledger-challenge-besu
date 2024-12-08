package repository

import (
	"database/sql"

	"github.com/luan441/goledger-challenge-besu/internal/entity"
)

type BlockchainValueRepository struct {
	conn *sql.DB
}

func NewBlockchainValueRepository(conn *sql.DB) *BlockchainValueRepository {
	return &BlockchainValueRepository{
		conn: conn,
	}
}

func (bvr *BlockchainValueRepository) Insert(value int64) (id int, err error) {
	sql := "INSERT INTO blockchain_value (value) VALUES ($1) RETURNING id"

	err = bvr.conn.QueryRow(sql, value).Scan(&id)
	if err != nil {
		return 0, err
	}

	return
}

func (bvr *BlockchainValueRepository) GetLast() (bv entity.BlockchainValue, err error) {
	sql := "SELECT * FROM blockchain_value ORDER BY created_at DESC LIMIT 1"

	err = bvr.conn.QueryRow(sql).Scan(&bv.ID, &bv.Value, &bv.CreatedAt)
	if err != nil {
		return entity.BlockchainValue{}, err
	}

	return
}
