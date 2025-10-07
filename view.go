package main

func (m model) banner() string {
	if m.ready {
		return m.currentBanner[m.currentFrame]
	}

	return "Loading..."
}
