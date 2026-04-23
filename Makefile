all: aa

a: 
	././a.out

	# Находим все .out файлы рекурсивно
EXECUTABLES := $(shell find . -name "*.out" -type f)

.PHONY: run-all

aa:
	@for exec in $(EXECUTABLES); do \
		echo "Running $$exec..."; \
		chmod +x $$exec; \
		./$$exec; \
	done
