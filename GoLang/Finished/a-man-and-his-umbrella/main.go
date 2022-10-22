package main

import (
	"fmt"
)

func main() {
	// fmt.Println(MinUmbrellas([]string{"rainy", "rainy", "rainy", "thunderstorms"}))
	// fmt.Println(MinUmbrellas([]string{"cloudy", "thunderstorms", "thunderstorms", "thunderstorms", "thunderstorms", "thunderstorms", "clear", "thunderstorms", "sunny", "clear"}))
	fmt.Println(MinUmbrellas([]string{"rainy", "sunny", "rainy", "rainy", "sunny", "rainy", "rainy", "thunderstorms", "rainy", "cloudy"}))
}

func MinUmbrellas(weather []string) int {
	h, w, t, c := 0, 0, 0, 0
	for i, g := range weather {
		if i%2 == 0 && (g == "rainy" || g == "thunderstorms") {
			if h == 0 {
				w++
				c++
			} else {
				w++
				h--
			}
		} else if g == "rainy" || g == "thunderstorms" {
			if w == 0 {
				h++
				c++
			} else {
				h++
				w--
			}
		} else {
			t++
		}
	}

	if t == len(weather) {
		return 0
	} else if c == 0 {
		return 1
	} else {
		return c
	}
}
