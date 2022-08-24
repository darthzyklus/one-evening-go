package main

func DebugLog(args ...string) []string {
	levels := []string{"[DEBUG]"}
	return append(levels, args...)
}

func InfoLog(args ...string) []string {
	levels := []string{"[INFO]"}
	return append(levels, args...)
}

func ErrorLog(args ...string) []string {
	levels := []string{"[ERROR]"}
	return append(levels, args...)
}
