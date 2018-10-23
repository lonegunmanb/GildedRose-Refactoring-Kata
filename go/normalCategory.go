package main

type normalCategory struct {
}

func (n normalCategory) updateQuality(item *Item) {
	decrementQuality(item)
}

func (n normalCategory) updateSellIn(item *Item) {
	decrementSellIn(item)
}

func (n normalCategory) processExpire(item *Item) {
	decrementQuality(item)
}
