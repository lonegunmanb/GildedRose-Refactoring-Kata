package main

type passCategory struct {
}

func (p passCategory) updateQuality(item *Item) {
	incrementQuality(item)
	if item.sellIn < 11 {
		incrementQuality(item)
	}
	if item.sellIn < 6 {
		incrementQuality(item)
	}
}

func (p passCategory) updateSellIn(item *Item) {
	decrementSellIn(item)
}

func (p passCategory) processExpire(item *Item) {

	item.quality = 0
}
