package main

func decrementSellIn(item *Item) {
	item.sellIn = item.sellIn - 1
}

func decrementQuality(item *Item) {
	if item.quality > 0 {
		item.quality = item.quality - 1
	}
}

func incrementQuality(item *Item) {
	if item.quality < 50 {
		item.quality = item.quality + 1
	}
}
