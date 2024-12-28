package ColorMe

import (
	"fmt"
)

// Define constants for text colors and background colors
const (
	// Text colors (foreground)
	Black   = 30
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37

	// Background colors
	BlackBg   = 40
	RedBg     = 41
	GreenBg   = 42
	YellowBg  = 43
	BlueBg    = 44
	MagentaBg = 45
	CyanBg    = 46
	WhiteBg   = 47

	// Bold
	Bold = 1

	// Reset
	Reset = 0
)

// ColorMe struct holds the text and the applied styles
type ColorMe struct {
	text      string
	textColor int
	bgColor   int
	isBold    bool
}

// Constructor function to create a new ColorMe instance
func Me(text string) *ColorMe {
	return &ColorMe{
		text: text,
	}
}

// Apply color methods for each color

// Red applies the red text color
func (c *ColorMe) Red() *ColorMe {
	c.textColor = Red
	return c
}

// Green applies the green text color
func (c *ColorMe) Green() *ColorMe {
	c.textColor = Green
	return c
}

// Blue applies the blue text color
func (c *ColorMe) Blue() *ColorMe {
	c.textColor = Blue
	return c
}

// Yellow applies the yellow text color
func (c *ColorMe) Yellow() *ColorMe {
	c.textColor = Yellow
	return c
}

// Magenta applies the magenta text color
func (c *ColorMe) Magenta() *ColorMe {
	c.textColor = Magenta
	return c
}

// Cyan applies the cyan text color
func (c *ColorMe) Cyan() *ColorMe {
	c.textColor = Cyan
	return c
}

// White applies the white text color
func (c *ColorMe) White() *ColorMe {
	c.textColor = White
	return c
}

// Black applies the black text color
func (c *ColorMe) Black() *ColorMe {
	c.textColor = Black
	return c
}

// Apply background color methods

// RedBg applies a red background
func (c *ColorMe) RedBg() *ColorMe {
	c.bgColor = RedBg
	return c
}

// GreenBg applies a green background
func (c *ColorMe) GreenBg() *ColorMe {
	c.bgColor = GreenBg
	return c
}

// BlueBg applies a blue background
func (c *ColorMe) BlueBg() *ColorMe {
	c.bgColor = BlueBg
	return c
}

// YellowBg applies a yellow background
func (c *ColorMe) YellowBg() *ColorMe {
	c.bgColor = YellowBg
	return c
}

// MagentaBg applies a magenta background
func (c *ColorMe) MagentaBg() *ColorMe {
	c.bgColor = MagentaBg
	return c
}

// CyanBg applies a cyan background
func (c *ColorMe) CyanBg() *ColorMe {
	c.bgColor = CyanBg
	return c
}

// WhiteBg applies a white background
func (c *ColorMe) WhiteBg() *ColorMe {
	c.bgColor = WhiteBg
	return c
}

// BlackBg applies a black background
func (c *ColorMe) BlackBg() *ColorMe {
	c.bgColor = BlackBg
	return c
}

// Apply bold style

// Bold makes the text bold
func (c *ColorMe) Bold() *ColorMe {
	c.isBold = true
	return c
}

// Method to apply the formatting and return the styled text as a string
func (c *ColorMe) String() string {
	var escapeSeq string

	// Add bold if applicable
	if c.isBold {
		escapeSeq = fmt.Sprintf("\033[%d;", Bold)
	} else {
		escapeSeq = "\033["
	}

	// Add text color if applicable
	if c.textColor > 0 {
		escapeSeq += fmt.Sprintf("%dm", c.textColor)
	}

	// Add background color if applicable
	if c.bgColor > 0 {
		if c.textColor > 0 {
			escapeSeq = fmt.Sprintf("\033[%d;%dm", Bold, c.bgColor)
		} else {
			escapeSeq += fmt.Sprintf("%dm", c.bgColor)
		}
	}

	// Return the final formatted string with reset at the end
	return fmt.Sprintf("%s%s\033[0m", escapeSeq, c.text)
}

// Method to print directly to the console (with newline)
func (c *ColorMe) Print() {
	fmt.Print(c.String())
}

// Method to print directly to the console without newline
func (c *ColorMe) Println() {
	fmt.Println(c.String())
}
