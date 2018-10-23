package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Normal_Item_SellIn_And_Quality_Should_Decrement_By_1(t *testing.T) {
	assertQualityAndSellIn(t, "+5 Dexterity Vest", 1, 1, 0, 0)
}

func Test_Quality_Decrement_By_2_After_SellIn(t *testing.T) {
	assertQualityAndSellIn(t, "+5 Dexterity Vest", 0, 2, -1, 0)
}

func Test_Quality_Cannot_Be_Negative(t *testing.T) {
	assertQualityAndSellIn(t, "+5 Dexterity Vest", 1, 0, 0, 0)
}

func Test_Chess_Quality_Increment_By_1(t *testing.T) {
	assertQualityAndSellIn(t, "Aged Brie", 1, 1, 0, 2)
}

func Test_Quality_Never_More_Than_50(t *testing.T) {
	assertQualityAndSellIn(t, "Aged Brie", 1, 50, 0, 50)
}

func Test_Sulfruas_Never_Change_SellIn_And_Quality(t *testing.T) {
	assertQualityAndSellIn(t, "Sulfuras, Hand of Ragnaros", 0, 80, 0, 80)
}

func Test_Ticket_Quality_Should_Increment_By_1_When_SellIn_More_Than_10(t *testing.T) {
	assertQualityAndSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 11, 1, 10, 2)
}

func Test_Ticket_Quality_Should_Increment_By_2_When_SellIn_Between_10_To_6(t *testing.T) {
	assertQualityAndSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 10, 1, 9, 3)
	assertQualityAndSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 6, 1, 5, 3)
}

func Test_Ticket_Quality_Should_Increment_By_3_When_SellIn_Between_5_To_1(t *testing.T) {
	assertQualityAndSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 5, 1, 4, 4)
	assertQualityAndSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 1, 1, 0, 4)
}

func Test_Ticket_Quality_Should_Be_0_After_SellIn(t *testing.T) {
	assertQualityAndSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 0, 10, -1, 0)
}

func Test_Conjured_Quality_Should_Decrement_By_2_Before_SellIn(t *testing.T) {
	assertQualityAndSellIn(t, "Conjured Mana Cake", 1, 2, 0, 0)
	assertQualityAndSellIn(t, "Conjured Mana Cake", 1, 1, 0, 0)
}

func Test_Conjured_Quality_Should_Decrement_By_4_Before_SellIn(t *testing.T) {
	assertQualityAndSellIn(t, "Conjured Mana Cake", 0, 4, -1, 0)
	assertQualityAndSellIn(t, "Conjured Mana Cake", 0, 1, -1, 0)
}

func assertQualityAndSellIn(t *testing.T, name string, sellIn int, quality int, expectedSellIn int, expectedQuality int) {
	items := []Item{{name, sellIn, quality}}
	GildedRose(items)
	item := items[0]
	assert.Equal(t, expectedSellIn, item.sellIn)
	assert.Equal(t, expectedQuality, item.quality)
}
