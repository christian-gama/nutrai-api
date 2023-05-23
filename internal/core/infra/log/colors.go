package log

import "github.com/fatih/color"

var (
	LoadingColor       = color.New(color.FgGreen).SprintfFunc()
	LoadingDetailColor = color.New(color.FgHiGreen, color.Bold, color.Underline).SprintfFunc()
	FatalColor         = color.New(color.FgRed).SprintFunc()
)
