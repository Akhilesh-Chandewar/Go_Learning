package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapname map[keytype]valuetype
	// mapvaribale := make(map[keytype]valuetype)
	
	mp1 := make(map[string]int)
	fmt.Println(mp1)
	mp1["key1"] = 1
	mp1["key2"] = 1
	mp1["key1"] = 100
	mp1["key3"] = 3
	mp1["key4"] = 4
	mp1["key5"] = 5
	fmt.Println(mp1 , len(mp1))
	fmt.Println(mp1["key1"] , mp1["NE"])

	delete(mp1 , "key1")
	fmt.Println(mp1)

	for key , value := range mp1 {
		fmt.Println(key , value)
	}

	value , ok := mp1["key1"]
	fmt.Println(value , ok)

	value , ok = mp1["key2"]
	fmt.Println(value , ok)

	_, ok = mp1["key3"]
	fmt.Println(ok)

	clear(mp1)
	fmt.Println(mp1)

	mp2 := map[string]int{
		"a": 1,
		"b" : 2,
		"c" : 3,
	}
	fmt.Println(mp2)

	mp3 := map[string]int{
		"a": 1,
		"b" : 2,
		"c" : 3,
	}
	fmt.Println(mp3)

	if(maps.Equal(mp1 , mp2)) {
		fmt.Println("maps are equal")
	}

	if(maps.Equal(mp2 , mp3)) {
		fmt.Println("maps are equal")
	}

	var mp4 map[string]string
	fmt.Println(mp4)
	if mp4 == nil {
		fmt.Println("map is nil")
	} else {
		fmt.Println("map is not nil")
	}

	val := mp4["key"]
	fmt.Println(">>>", val)

	mp4 = make(map[string]string)
	mp4["key"] = "value"
	val = mp4["key"]
	fmt.Println(">>>", val)


	mp5 := make(map[string]map[string]string)
	mp5["map1"] = mp4
	fmt.Println(mp5)
}	