package ColorMe

import (
  "fmt"
)

// Define constants for the text colors and background colors.
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

// Function to set background color and return text
func Background(text string, color int) string {
  return fmt.Sprintf("\033[%dm%s\033[0m", color, text)
}

// Function to set text color and return text
func Color(text string, color int) string {
  return fmt.Sprintf("\033[%dm%s\033[0m", color, text)
}

// Function to set text color and make it bold, then return text
func BoldColor(text string, textColor, bgColor int) string {
  return fmt.Sprintf("\033[%d;%dm%s\033[0m", Bold, textColor, text)
}

// Function to set background color and make it bold, then return text
func BoldBackground(text string, bgColor int) string {
  return fmt.Sprintf("\033[%d;%dm%s\033[0m", Bold, bgColor, text)
}

// Function to set background color and return text without newline
func BackgroundNoNewline(text string, color int) string {
  return fmt.Sprintf("\033[%dm%s\033[0m", color, text)
}

// Function to set text color and return text without newline
func ColorNoNewline(text string, color int) string {
  return fmt.Sprintf("\033[%dm%s\033[0m", color, text)
}

// Function to set text color and make it bold, then return text without newline
func BoldColorNoNewline(text string, textColor, bgColor int) string {
  return fmt.Sprintf("\033[%d;%dm%s\033[0m", Bold, textColor, text)
}

// Function to set background color and make it bold, then return text without newline
func BoldBackgroundNoNewline(text string, bgColor int) string {
  return fmt.Sprintf("\033[%d;%dm%s\033[0m", Bold, bgColor, text)
}

// ----------------------------------------
// Print functions that directly output to console
// ----------------------------------------

// Function to print background color and text with newline
func BackgroundPl(text string, color int) {
  fmt.Printf("\033[%dm%s\033[0m\n", color, text)
}

// Function to print text color and text with newline
func ColorPl(text string, color int) {
  fmt.Printf("\033[%dm%s\033[0m\n", color, text)
}

// Function to print bold text color and text with newline
func BoldColorPl(text string, textColor, bgColor int) {
  fmt.Printf("\033[%d;%dm%s\033[0m\n", Bold, textColor, text)
}

// Function to print bold background color and text with newline
func BoldBackgroundPl(text string, bgColor int) {
  fmt.Printf("\033[%d;%dm%s\033[0m\n", Bold, bgColor, text)
}

// Function to print background color and text without newline
func BackgroundP(text string, color int) {
  fmt.Printf("\033[%dm%s\033[0m", color, text)
}

// Function to print text color and text without newline
func ColorP(text string, color int) {
  fmt.Printf("\033[%dm%s\033[0m", color, text)
}

// Function to print bold text color and text without newline
func BoldColorP(text string, textColor, bgColor int) {
  fmt.Printf("\033[%d;%dm%s\033[0m", Bold, textColor, text)
}

// Function to print bold background color and text without newline
func BoldBackgroundP(text string, bgColor int) {
  fmt.Printf("\033[%d;%dm%s\033[0m", Bold, bgColor, text)
}