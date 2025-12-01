GO = go
BIN = bin/aoc

.PHONY: build test clean

build:
	@mkdir -p $(dir $(BIN))
	@$(GO) build -o $(BIN) ./cmd/aoc

test:
	@$(GO) test ./internal/...

clean:
	@rm -f $(BIN)

%:
	@DAY_ARG=$@; \
	if echo $$DAY_ARG | grep -qE '^[0-9]{4}D[0-9]+$$'; then \
		YEAR=$$(echo $$DAY_ARG | sed 's/D.*//'); \
		DAY_NUM=$$(echo $$DAY_ARG | sed 's/.*D//'); \
		DAY_NAME=$$(case $$DAY_NUM in \
			1) echo "One" ;; \
			2) echo "Two" ;; \
			3) echo "Three" ;; \
			4) echo "Four" ;; \
			5) echo "Five" ;; \
			6) echo "Six" ;; \
			7) echo "Seven" ;; \
			8) echo "Eight" ;; \
			9) echo "Nine" ;; \
			10) echo "Ten" ;; \
			11) echo "Eleven" ;; \
			12) echo "Twelve" ;; \
			13) echo "Thirteen" ;; \
			14) echo "Fourteen" ;; \
			15) echo "Fifteen" ;; \
			16) echo "Sixteen" ;; \
			17) echo "Seventeen" ;; \
			18) echo "Eighteen" ;; \
			19) echo "Nineteen" ;; \
			20) echo "Twenty" ;; \
			21) echo "TwentyOne" ;; \
			22) echo "TwentyTwo" ;; \
			23) echo "TwentyThree" ;; \
			24) echo "TwentyFour" ;; \
			25) echo "TwentyFive" ;; \
			*) echo "Unknown" ;; \
		esac); \
		mkdir -p internal/$$YEAR/Day$$DAY_NAME; \
		mkdir -p internal/data/$$YEAR/Day$$DAY_NAME; \
		touch internal/$$YEAR/Day$$DAY_NAME/code.go; \
		touch internal/$$YEAR/Day$$DAY_NAME/code_test.go; \
		if [ -n "$$ADVENTOFCODE_SESSION" ]; then \
			curl -s -H "Cookie: session=$$ADVENTOFCODE_SESSION" \
				https://adventofcode.com/$$YEAR/day/$$DAY_NUM/input \
				| perl -pe 'chomp if eof' > internal/data/$$YEAR/Day$$DAY_NAME/input.txt; \
		else \
			touch internal/data/$$YEAR/Day$$DAY_NAME/input.txt; \
		fi; \
		echo "$$DAY_ARG ready to be solved."; \
	fi
