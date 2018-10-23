package main

type chessCategory struct {
}

func (c chessCategory) updateQuality(item *Item) {
	incrementQuality(item)
}

func (c chessCategory) updateSellIn(item *Item) {
	decrementSellIn(item)
}

func (c chessCategory) processExpire(item *Item) {
	incrementQuality(item)
}
