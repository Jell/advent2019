.DEFAULT_GOAL := all

MAIN_FILES := $(wildcard ./*.clj)

DIFFS := $(shell ls | grep -e "day.._part." | grep -o "day.." | uniq)

.PHONY: $(MAIN_FILES) $(DIFFS)

$(MAIN_FILES):
	clj -J-ss10000k $@

$(DIFFS):
	git diff --word-diff=color --color-moved-ws=ignore-all-space --diff-algorithm=minimal --no-index $@_part* | ansifilter -S > $@_diff.svg

all: $(MAIN_FILES) $(DIFFS)
