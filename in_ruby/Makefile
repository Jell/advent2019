.DEFAULT_GOAL := all

MAIN_FILES := $(wildcard ./day*.rb)

DIFFS := $(shell ls | grep -e "day_.._part_." | grep -o "day_.." | uniq)

.PHONY: $(MAIN_FILES) $(DIFFS)

$(MAIN_FILES):
	ruby $@

$(DIFFS):
	git diff --word-diff=color --color-moved-ws=ignore-all-space --diff-algorithm=minimal --no-index $@_part_* | ansifilter -S > $@_diff.svg

all: $(MAIN_FILES) $(DIFFS)
