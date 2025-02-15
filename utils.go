package main

func switchWindow(name string) {
	for key := range isActive {
		if key == name {
			isActive[key] = true
		} else {
			isActive[key] = false
		}
	}
}
