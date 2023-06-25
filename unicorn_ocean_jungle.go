package main

import (
	"fmt"
)

// A craft brewery that creates unique and flavorful micro-brews using locally sourced ingredients and artisanal techniques.
type Brewery struct {
	Name       string
	Ingredients []string
	Style      string
	Process    string
}

// Create a new brewery
func NewBrewery(name string, ingredients []string, style string, process string) *Brewery {
	return &Brewery{
		Name:       name,
		Ingredients: ingredients,
		Style:      style,
		Process:    process,
	}
}

// Describe a brewery
func (b *Brewery) Describe() {
	fmt.Printf("%s is a craft brewery that creates unique and flavorful micro-brews using locally sourced ingredients and artisanal techniques.\n", b.Name)
	fmt.Printf("Ingredients include: %s\n", b.Ingredients)
	fmt.Printf("Brewery specializes in %s style beer\n", b.Style)
	fmt.Printf("Brewing process is %s\n", b.Process)
}

func main() {
	ingredients1 := []string{"Hop", "Malt", "Yeast", "Water"}
	brew1 := NewBrewery("Brew1", ingredients1, "IPA", "Traditional")
	brew1.Describe()

	ingredients2 := []string{"Hop", "Malt", "Yeast", "Water", "Flavor"}
	brew2 := NewBrewery("Brew2", ingredients2, "Pale Ale", "Unique")
	brew2.Describe()
}