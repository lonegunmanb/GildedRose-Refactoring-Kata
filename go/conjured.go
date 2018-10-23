package main

type conjuredCategory struct {
}

func (c conjuredCategory) updateQuality(item *Item) {
	decrementQuality(item)
	decrementQuality(item)
}

func (c conjuredCategory) updateSellIn(item *Item) {
	decrementSellIn(item)
}

func (c conjuredCategory) processExpire(item *Item) {
	decrementQuality(item)
	decrementQuality(item)
}
