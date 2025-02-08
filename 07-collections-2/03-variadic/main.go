package main

func DebugLog(args ...string) []string {
	result := []string{"[DEBUG]"}

	return append(result, args...)
}

func InfoLog(args ...string) []string {
	result := []string{"[INFO]"}

	return append(result, args...)
}

func ErrorLog(args ...string) []string {
	result := []string{"[ERROR]"}

	return append(result, args...)
}
