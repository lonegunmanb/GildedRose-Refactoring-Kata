package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
======================================
Gilded Rose Requirements Specification
======================================

Hi and welcome to team Gilded Rose. As you know, we are a small inn with a prime location in a
prominent city ran by a friendly innkeeper named Allison. We also buy and sell only the finest goods.
Unfortunately, our goods are constantly degrading in quality as they approach their sell by date. We
have a system in place that updates our inventory for us. It was developed by a no-nonsense type named
Leeroy, who has moved on to new adventures. Your task is to add the new feature to our system so that
we can begin selling a new category of items. First an introduction to our system:
	- All items have a SellIn value which denotes the number of days we have to sell the item
	- All items have a Quality value which denotes how valuable the item is
	- At the end of each day our system lowers both values for every item

Pretty simple, right? Well this is where it gets interesting:

	- Once the sell by date has passed, Quality degrades twice as fast
	- The Quality of an item is never negative
	- "Aged Brie" actually increases in Quality the older it gets
	- The Quality of an item is never more than 50
	- "Sulfuras", being a legendary item, never has to be sold or decreases in Quality
	- "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
	Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
	Quality drops to 0 after the concert

We have recently signed a supplier of conjured items. This requires an update to our system:

	- "Conjured" items degrade in Quality twice as fast as normal items

Feel free to make any changes to the UpdateQuality method and add any new code as long as everything
still works correctly. However, do not alter the Item class or Items property as those belong to the
goblin in the corner who will insta-rage and one-shot you as he doesn't believe in shared code
ownership (you can make the UpdateQuality method and Items property static if you like, we'll cover
for you).

Just for clarification, an item can never have its Quality increase above 50, however "Sulfuras" is a
legendary item and as such its Quality is 80 and it never alters.
*/

func TestShouldDecreaseSellIn(t *testing.T) {
	names := []string{
		"+5 Dexterity Vest",
		"Aged Brie",
		"Elixir of the Mongoose",
		"Backstage passes to a TAFKAL80ETC concert",
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			sellIn := 10
			item, items := newItems(name, sellIn, 10)
			UpdateQuality(items)
			assert.Equal(t, sellIn-1, item.sellIn)
		})
	}
}

func TestShouldDecreaseQuality(t *testing.T) {
	names := []string{
		"+5 Dexterity Vest",
		"Elixir of the Mongoose",
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			quality := 10
			item, items := newItems(name, 10, quality)
			UpdateQuality(items)
			assert.Equal(t, quality-1, item.quality)
		})
	}
}

func TestShouldDecreaseQualityTwiceAfterExpire(t *testing.T) {
	names := []string{
		"+5 Dexterity Vest",
		"Elixir of the Mongoose",
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			quality := 10
			item, items := newItems(name, 0, quality)
			UpdateQuality(items)
			assert.Equal(t, quality-2, item.quality)
		})
	}
}

func TestBrieShouldIncreaseQuality(t *testing.T) {
	t.Run("Not Expire", func(t *testing.T) {
		quality := 1
		item, items := newItems("Aged Brie", 1, quality)
		UpdateQuality(items)
		assert.Equal(t, quality+1, item.quality)
	})
	t.Run("Expire", func(t *testing.T) {
		quality := 1
		item, items := newItems("Aged Brie", 0, quality)
		UpdateQuality(items)
		assert.Equal(t, quality+2, item.quality)
	})
}

func TestQualityShouldNeverBeNegative(t *testing.T) {
	names := []string{
		"+5 Dexterity Vest",
		"Elixir of the Mongoose",
	}

	for _, name := range names {
		f := func(sellIn int, quality int) {
			t.Run(fmt.Sprintf("%s sellIn:%d quality:%d", name, sellIn, quality), func(t *testing.T) {
				item, items := newItems(name, sellIn, quality)
				UpdateQuality(items)
				assert.Equal(t, 0, item.quality)
			})
		}

		f(1, 1)
		f(1, 0)
		f(0, 2)
		f(0, 1)
		f(0, 0)
	}
}

func TestQualityShouldNotMoreThanFifty(t *testing.T) {
	names := []string{
		"Aged Brie",
		"Backstage passes to a TAFKAL80ETC concert",
	}
	for _, name := range names {
		f := func(sellIn int) {
			t.Run(fmt.Sprintf("%s sellIn:%d", name, sellIn), func(t *testing.T) {
				item, items := newItems(name, sellIn, 50)
				UpdateQuality(items)
				assert.True(t, item.quality <= 50)
			})
		}
		f(11)
		f(10)
		f(6)
		f(5)
		f(1)
		f(0)
	}
}

func TestSulfurasNeverChangeSellInOrQuality(t *testing.T) {
	f := func(sellIn int, quality int) {
		t.Run(fmt.Sprintf("sellIn:%d quality:%d", sellIn, quality), func(t *testing.T) {
			item, items := newItems("Sulfuras, Hand of Ragnaros", sellIn, quality)
			UpdateQuality(items)
			assert.Equal(t, sellIn, item.sellIn)
			assert.Equal(t, quality, item.quality)
		})
	}

	f(10, 80)
	f(-10, 80)
	f(10, -1)
	f(0, 80)
	f(0, -1)
	f(-1, -1)
}

func TestBackstagePasses(t *testing.T) {
	quality := 10
	f := func(sellIn int, expectedQuality int) {
		t.Run(fmt.Sprintf("sellIn:%d", sellIn), func(t *testing.T) {
			item, items := newItems("Backstage passes to a TAFKAL80ETC concert", sellIn, quality)
			UpdateQuality(items)
			assert.Equal(t, expectedQuality, item.quality)
		})
	}
	f(11, quality+1)
	f(10, quality+2)
	f(6, quality+2)
	f(5, quality+3)
	f(1, quality+3)
	f(0, 0)
}

func newItems(name string, sellIn int, quality int) (*Item, []*Item) {
	item := &Item{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
	items := []*Item{
		item,
	}
	return item, items
}
