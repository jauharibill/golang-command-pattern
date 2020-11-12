package main

func main()  {
	tv := &tv {
		isOn: false,
		volume: 10,
	}

	hel := &hello{isOn: true}

	// Instantiate commands
	onCommand := &onCommand{
		device: hel,
	}
	offCommand := &offCommand{
		device: tv,
	}
	increaseVolumeCommand := &increaseVolumeCommand {
		device: tv,
	}
	decreaseVolumeCommand := &decreaseVolumeCommand {
		device: tv,
	}

	// Instantiate buttons
	onButton := &button{
		command: onCommand,
	}
	offButton := &button{
		command: offCommand,
	}
	increaseVolumeButton := &button {
		command: increaseVolumeCommand,
	}
	decreaseVolumeButton := &button {
		command: decreaseVolumeCommand,
	}

	// Execute
	increaseVolumeButton.press()
	onButton.press()
	increaseVolumeButton.press()
	decreaseVolumeButton.press()
	offButton.press()
}
