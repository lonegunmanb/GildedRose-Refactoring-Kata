package main

import (
	"fmt"
)

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("OMGHAI!")
	// fmt.Print(items)
	GildedRose(items)
}

func GildedRose(items []Item) {
	for i := 0; i < len(items); i++ {
		item := &items[i]
		updateItem(item)
	}
}

func updateItem(item *Item) {
	category := categories[item.name]
	category.updateQuality(item)
	category.updateSellIn(item)
	if item.sellIn < 0 {
		category.processExpire(item)
	}
}
