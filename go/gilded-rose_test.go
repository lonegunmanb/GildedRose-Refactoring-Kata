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
欢迎来到Gilded Rose
Gilded Rose是魔兽世界中位于暴风城的一家客栈，他们同时也销售一些商品。
为了维护商品保质期等信息，客栈请了一位炸鸡勇者开发了一套库存系统，然后开发者就离开了客栈投入了他的冒险旅程。
你的任务是为这套库存系统添加一个新功能，使得客栈可以销售一种新的商品。

简单介绍一下当前的系统：
所有的商品都由SellIn属性，代表距离过期还有多少天（0代表过期前最后一天）
所有的商品都有Quality属性，代表商品的价值
库存系统每天会执行盘点，在盘点中扣减商品的SellIn与Quality

特殊规则：
- 一旦商品过期，商品价值以两倍速度扣减
- 商品的价值不可为负数
- "Aged Brie"的价值随着时间的流逝递增（每天增加1）
- 商品的价值最大为50
- "Sulfuras, Hand of Ragnaros"是一件传奇物品，是非卖品，保质期与价值不会变化
- "Backstage passes to a TAFKAL80ETC concert"类似"Aged Brie"，其价值随着时间的流逝递增（每天增加1）；
	从演出开始前10天开始，价值每天递增2
	从演出开始前5天开始，价值每天递增3
	演出结束后价值归零

客栈想要贩售的新商品是"Conjured Mana Cake"，类似于普通的商品，只是它的价值扣减的速度是普通商品的两倍

您可以对UpdateQuality进行您认为合适的修改，只要不要引入BUG就行。但是请不要修改Item类型极其属性的代码，因为一群残暴的
食人魔认定他们拥有Item代码的所有权。

最后补充说明的是，商品的价值最大为50，但是"Sulfuras, Hand of Ragnaros"这件非卖品不受此限制，它的价值可以是80。
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
	quality := 1
	t.Run("Not Expire", func(t *testing.T) {
		item, items := newItems("Aged Brie", 1, quality)
		UpdateQuality(items)
		assert.Equal(t, quality+1, item.quality)
	})
	t.Run("Expire", func(t *testing.T) {
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
	inputs := []struct {
		sellIn  int
		quality int
	}{
		{1, 1},
		{1, 0},
		{0, 2},
		{0, 1},
		{0, 0},
	}
	for _, name := range names {
		for _, input := range inputs {
			t.Run(fmt.Sprintf("%s sellIn:%d quality:%d", name, input.sellIn, input.quality), func(t *testing.T) {
				item, items := newItems(name, input.sellIn, input.quality)
				UpdateQuality(items)
				assert.Equal(t, 0, item.quality)
			})
		}
	}
}

func TestQualityShouldNotMoreThanFifty(t *testing.T) {
	names := []string{
		"Aged Brie",
		"Backstage passes to a TAFKAL80ETC concert",
	}
	for _, name := range names {
		sellIns := []int{11, 10, 6, 5, 1, 0}
		for _, sellIn := range sellIns {
			t.Run(fmt.Sprintf("%s sellIn:%d", name, sellIn), func(t *testing.T) {
				item, items := newItems(name, sellIn, 50)
				UpdateQuality(items)
				assert.True(t, item.quality <= 50)
			})
		}
	}
}

func TestSulfurasNeverChangeSellInOrQuality(t *testing.T) {
	inputs := []struct {
		sellIn          int
		expectedQuality int
	}{
		{10, 80},
		{-10, 80},
		{10, -1},
		{0, 80},
		{0, -1},
		{-1, -1},
	}

	for _, input := range inputs {
		quality := input.expectedQuality
		sellIn := input.sellIn
		t.Run(fmt.Sprintf("sellIn:%d quality:%d", sellIn, quality), func(t *testing.T) {
			item, items := newItems("Sulfuras, Hand of Ragnaros", sellIn, quality)
			UpdateQuality(items)
			assert.Equal(t, sellIn, item.sellIn)
			assert.Equal(t, quality, item.quality)
		})
	}
}

func TestBackstagePasses(t *testing.T) {
	quality := 10
	inputs := []struct {
		sellIn          int
		expectedQuality int
	}{
		{11, quality + 1},
		{10, quality + 2},
		{6, quality + 2},
		{5, quality + 3},
		{1, quality + 3},
		{0, 0},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("sellIn:%d", input.sellIn), func(t *testing.T) {
			item, items := newItems("Backstage passes to a TAFKAL80ETC concert", input.sellIn, quality)
			UpdateQuality(items)
			assert.Equal(t, input.expectedQuality, item.quality)
		})
	}
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
