package main

import "fmt"

type tv struct {
	isOn bool
	volume int
}

func (t *tv) on() {
	t.isOn = true
	fmt.Println("device is turning on")
}

func (t *tv) off() {
	t.isOn = false
	fmt.Println("device is turning off")
}

func (t *tv) increaseVolume() {
	if t.isOn {
		if t.volume > 0 && t.volume <= 100 {
			t.volume++
			fmt.Sprintf("volume is increased at level %d\n", t.volume)
		} else {
			fmt.Println("volume max")
		}
	} else {
		fmt.Println("device is off")
	}
}

func (t *tv) decreaseVolume() {
	if t.isOn {
		if t.volume > 0 && t.volume <= 100 {
			t.volume--
			fmt.Sprintf("volume is decreased at level %d\n", t.volume)
		} else {
			fmt.Println("volume is muted")
		}
	} else {
		fmt.Println("device is off")
	}
}
