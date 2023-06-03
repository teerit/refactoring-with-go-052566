package main

import (
	"github.com/kr/pretty"
	"testing"
)

func Test_Foo(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []*Item
		outputs []*Item
	}{
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", 10, 20}}, outputs: []*Item{{"+5 Dexterity Vest", 9, 19}}},
		{name: "#2", inputs: []*Item{{"Aged Brie", 2, 0}}, outputs: []*Item{{"Aged Brie", 1, 1}}},
		{name: "#3", inputs: []*Item{{"Elixir of the Mongoose", 5, 7}}, outputs: []*Item{{"Elixir of the Mongoose", 4, 6}}},
		{name: "#4", inputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
		{name: "#5", inputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
		{name: "#6", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 1}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 14, 2}}},
		{name: "#7", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 50}}},
		{name: "#8", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 50}}},
		// {name: "#1", inputs: []*Item{{"Conjured Mana Cake", 3, 6}}, outputs: []*Item{{"Conjured Mana Cake", 2, 5}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			UpdateQuality(c.inputs)

			if diff := pretty.Diff(c.inputs, c.outputs); len(diff) > 0 {
				t.Errorf("%s not match : \n%#v", c.name, pretty.Diff(c.inputs, c.outputs))
			}
		})
	}
}

func TestNormalItems(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []*Item
		outputs []*Item
	}{
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", 5, 15}}, outputs: []*Item{{"+5 Dexterity Vest", 4, 14}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", 4, 14}}, outputs: []*Item{{"+5 Dexterity Vest", 3, 13}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", 3, 13}}, outputs: []*Item{{"+5 Dexterity Vest", 2, 12}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", 2, 12}}, outputs: []*Item{{"+5 Dexterity Vest", 1, 11}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", 1, 11}}, outputs: []*Item{{"+5 Dexterity Vest", 0, 10}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", 0, 10}}, outputs: []*Item{{"+5 Dexterity Vest", -1, 8}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", -1, 8}}, outputs: []*Item{{"+5 Dexterity Vest", -2, 6}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", -2, 6}}, outputs: []*Item{{"+5 Dexterity Vest", -3, 4}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", -3, 4}}, outputs: []*Item{{"+5 Dexterity Vest", -4, 2}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", -4, 2}}, outputs: []*Item{{"+5 Dexterity Vest", -5, 0}}},
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", -5, 0}}, outputs: []*Item{{"+5 Dexterity Vest", -6, 0}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			UpdateQuality(c.inputs)

			if diff := pretty.Diff(c.inputs, c.outputs); len(diff) > 0 {
				t.Errorf("%s not match : \n%#v", c.name, pretty.Diff(c.inputs, c.outputs))
			}
		})
	}
}
