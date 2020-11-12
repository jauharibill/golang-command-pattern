package main

type hello struct {
	isOn bool
}

func (h *hello) on()  {
	h.isOn = true
}

func (h *hello) off()  {
	h.isOn = true
}

func (h *hello) increaseVolume()  {
	h.isOn = true
}

func (h *hello) decreaseVolume()  {
	h.isOn = true
}

