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

  // Background colors (for reference but not used in Background functions)
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

// Function to set foreground color and return text with newline
func Background(text string, color int) string {
  // Now changes foreground color, not background
  return fmt.Sprintf("\033[%dm%s\033[0m\n", color, text) // Set text color (foreground) and add newline
}

// Function to set foreground color and return text with newline
func Color(text string, color int) string {
  return fmt.Sprintf("\033[%dm%s\033[0m\n", color, text) // Set text color and add newline
}

// Function to set text color and make it bold, then return text with newline
func BoldColor(text string, textColor, bgColor int) string {
  return fmt.Sprintf("\033[%d;%dm%s\033[0m\n", Bold, textColor, text) // Bold text and color with newline
}

// Function to set background color and make it bold, then return text with newline
func BoldBackground(text string, bgColor int) string {
  return fmt.Sprintf("\033[%d;%dm%s\033[0m\n", Bold, bgColor, text) // Bold background color with newline
}

// Function to set foreground color and return text without newline
func BackgroundNoNewline(text string, color int) string {
  // Now changes foreground color, not background
  return fmt.Sprintf("\033[%dm%s\033[0m", color, text) // Set text color without newline
}

// Function to set foreground color and return text without newline
func ColorNoNewline(text string, color int) string {
  return fmt.Sprintf("\033[%dm%s\033[0m", color, text) // Set text color without newline
}

// Function to set text color and make it bold, then return text without newline
func BoldColorNoNewline(text string, textColor, bgColor int) string {
  return fmt.Sprintf("\033[%d;%dm%s\033[0m", Bold, textColor, text) // Bold text color without newline
}

// Function to set background color and make it bold, then return text without newline
func BoldBackgroundNoNewline(text string, bgColor int) string {
  return fmt.Sprintf("\033[%d;%dm%s\033[0m", Bold, bgColor, text) // Bold background color without newline
}

// ----------------------------------------
// Print functions that directly output to console
// ----------------------------------------

// Function to print foreground color and text with newline
func BackgroundPl(text string, color int) {
  // Now changes foreground color, not background
  fmt.Printf("\033[%dm%s\033[0m\n", color, text) // Set text color (foreground) and print with newline
}

// Function to print text color and text with newline
func ColorPl(text string, color int) {
  fmt.Printf("\033[%dm%s\033[0m\n", color, text) // Set text color and print with newline
}

// Function to print bold text color and text with newline
func BoldColorPl(text string, textColor, bgColor int) {
  fmt.Printf("\033[%d;%dm%s\033[0m\n", Bold, textColor, text) // Bold text color and print with newline
}

// Function to print bold background color and text with newline
func BoldBackgroundPl(text string, bgColor int) {
  fmt.Printf("\033[%d;%dm%s\033[0m\n", Bold, bgColor, text) // Bold background color and print with newline
}

// Function to print foreground color and text without newline
func BackgroundP(text string, color int) {
  // Now changes foreground color, not background
  fmt.Printf("\033[%dm%s\033[0m", color, text) // Print text color (foreground) without newline
}

// Function to print text color and text without newline
func ColorP(text string, color int) {
  fmt.Printf("\033[%dm%s\033[0m", color, text) // Print text color without newline
}

// Function to print bold text color and text without newline
func BoldColorP(text string, textColor, bgColor int) {
  fmt.Printf("\033[%d;%dm%s\033[0m", Bold, textColor, text) // Print bold text color without newline
}

// Function to print bold background color and text without newline
func BoldBackgroundP(text string, bgColor int) {
  fmt.Printf("\033[%d;%dm%s\033[0m", Bold, bgColor, text) // Print bold background color without newline
}