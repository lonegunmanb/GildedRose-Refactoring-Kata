package main

//func (BackstagePasses) UpdateItemSellIn(item *Item) {
//	item.sellIn--
//}
//
//func (BackstagePasses) DealItemExpire(item *Item) {
//	item.quality = item.quality - item.quality
//}
//
//func (AgedBrie) UpdateItemSellIn(item *Item) {
//	item.sellIn--
//}
//
//func (AgedBrie) DealItemExpire(item *Item) {
//	if item.quality < 50 {
//		item.quality = item.quality + 1
//	}
//}
//
//func (Normal) UpdateItemSellIn(item *Item) {
//	item.sellIn--
//}
//
//func (Normal) DealItemExpire(item *Item) {
//	if item.quality > 0 {
//		item.quality = item.quality - 1
//	}
//}
//
//func (Sulfuras) UpdateItemSellIn(item *Item) {
//
//}
//
//func (Sulfuras) DealItemExpire(item *Item) {
//
//}
//
//func dealExpire(item *Item) {
//	if item.name == "Aged Brie" {
//		a := AgedBrie{}
//		a.DealItemExpire(item)
//	} else if item.name == "Backstage passes to a TAFKAL80ETC concert" {
//		b := BackstagePasses{}
//		b.DealItemExpire(item)
//
//	} else if item.name == "Sulfuras, Hand of Ragnaros" {
//		s := Sulfuras{}
//		s.DealItemExpire(item)
//	} else {
//		n := Normal{}
//		n.DealItemExpire(item)
//	}
//}
//
//func updateSellIn(item *Item) {
//	if item.name == "Backstage passes to a TAFKAL80ETC concert" {
//		b := BackstagePasses{}
//		b.UpdateItemSellIn(item)
//	} else if item.name == "Aged Brie" {
//		a := AgedBrie{}
//		a.UpdateItemSellIn(item)
//	} else if item.name == "Sulfuras, Hand of Ragnaros" {
//		s := Sulfuras{}
//		s.UpdateItemSellIn(item)
//	} else {
//		n := Normal{}
//		n.UpdateItemSellIn(item)
//	}
//}
