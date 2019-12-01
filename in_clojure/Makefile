.DEFAULT_GOAL := all

MAIN_FILES := $(wildcard ./*.clj)

DIFFS := $(shell ls | grep -e "day.._part." | grep -o "day.." | sort | uniq)

.PHONY: $(MAIN_FILES) $(DIFFS) README.md

$(MAIN_FILES):
	clj -J-ss10000k $@

README.md:
	echo "# In Clojure\n" > README.md

$(DIFFS): README.md
	git diff --word-diff=color --color-moved-ws=ignore-all-space --diff-algorithm=minimal --no-index $@_part* | ansifilter -S > $@_diff.svg
	echo "## $@" >> README.md
	echo "<img src=\"./$@_diff.svg\">\n" >> README.md

all: $(MAIN_FILES) $(DIFFS) README.md
	echo "<br>" >> README.md
