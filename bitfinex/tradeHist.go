package main

type TradeHistDesc struct {
	MTS       int `gorm:"type:int unsigned;not null;default:0;unique_index:idx_mts" `
	OPEN      float64
	CLOSE     float64
	HIGH      float64
	LOW       float64
	VOLUME    float64
	tableName string
}
