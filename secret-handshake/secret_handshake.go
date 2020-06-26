package secret

// Handshake convert given code to the appropriate
// sequence of events for a secret handshake
func Handshake(code uint) []string {
	var h []string

	if code == 0 {
		return h
	}
	if code&1 != 0 {
		h = append(h, "wink")
	}
	if code&2 != 0 {
		h = append(h, "double blink")
	}
	if code&4 != 0 {
		h = append(h, "close your eyes")
	}
	if code&8 != 0 {
		h = append(h, "jump")
	}
	if code&16 != 0 {
		for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
			h[i], h[j] = h[j], h[i]
		}
	}

	return h
}
