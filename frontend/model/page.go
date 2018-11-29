package model

import "crawler/engine"

type SearchResult struct {
	Hits int
	Start int
	Items []engine.Item
}