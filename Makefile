# Default target: make start-day DAY=02
start-day:
	@if [ -z "$(DAY)" ]; then \
		echo "❌ Please pass: make start-day DAY=02"; \
		exit 1; \
	fi
	@echo "🚀 Creating Advent of Code day $(DAY)..."
	@go run generate.go $(DAY)

# Shortcut: make day3, make day12, etc.
day%:
	@DAY=$* $(MAKE) start-day

.PHONY: start-day day%
