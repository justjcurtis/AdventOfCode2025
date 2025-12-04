.PHONY: newday
.PHONY: update
.PHONY: year

newday:
	./scripts/newDay.sh
update:
	./scripts/updateReadme.sh
year:
	./scripts/updateYear.sh
