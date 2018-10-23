package main

var categories = map[string]itemCategory{
	"+5 Dexterity Vest":      normalCategory{},
	"Elixir of the Mongoose": normalCategory{},
	"Aged Brie":              chessCategory{},
	"Backstage passes to a TAFKAL80ETC concert": passCategory{},
	"Sulfuras, Hand of Ragnaros":                legendaryCategory{},
	"Conjured Mana Cake":                        conjuredCategory{},
}
