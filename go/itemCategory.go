package main

type itemCategory interface {
	updateQuality(item *Item)
	updateSellIn(item *Item)
	processExpire(item *Item)
}
