package main

import "strings"

func CountCreatedEvents(events []string) int {
	result := 0

	for _, e := range events {
		if strings.HasSuffix(e, "_deleted") {
			break
		}

		if !strings.HasSuffix(e, "_created") {
			continue
		}

		result++

	}

	return result
}

func main() {
	events := []string{
		"product_created",
		"product_updated",
		"product_assigned",
		"order_created",
		"order_updated",
		"client_created",
		"client_updated",
		"client_refreshed",
		"client_deleted",
		"order_updated",
	}

	CountCreatedEvents(events)
}
