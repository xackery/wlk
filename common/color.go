package common

type Color uint32

// RGB returns a Color with the specified red, green, blue values.
func RGB(r, g, b byte) Color {
	return Color(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

// R returns the red component of the Color.
func (c Color) R() byte {
	return byte(c & 0xff)
}

// G returns the green component of the Color.
func (c Color) G() byte {
	return byte((c >> 8) & 0xff)
}

// B returns the blue component of the Color.
func (c Color) B() byte {
	return byte((c >> 16) & 0xff)
}
