package main

import (
	"fmt"

	_m "github.com/elliotchance/orderedmap"
)

func main() {
	orderedMaps := _m.NewOrderedMap()

	orderedMaps.Set("Name", "Clinton")
	orderedMaps.Set("Stack", "Go")
	orderedMaps.Set("Location", "Benin City")

	for orderedMap := orderedMaps.Front(); orderedMap != nil; orderedMap = orderedMap.Next() {
		fmt.Printf("%s : %s \n", orderedMap.Key, orderedMap.Value)
	}

	orderedMaps.Set("Height", 1.27)
	orderedMaps.Set("Status", "Single")

	orderedMaps.Delete("Status")

	fmt.Println("................Reverse.......................")

	for orderedmap := orderedMaps.Back(); orderedmap != nil; orderedmap = orderedmap.Prev() {
		fmt.Printf("%s : %v \n", orderedmap.Key, orderedmap.Value)
	}
}
