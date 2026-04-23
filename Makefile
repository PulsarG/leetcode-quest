CC = gcc
CFLAGS = -std=c11 -Wall -Wextra -Werror

all: run

%: %.c
	$(CC) $(CFLAGS) $< -o a.out

run:
	@find . -name "*.out" -type f -exec echo "Running {}..." \; -exec chmod +x {} \; -exec ./{} \;

clean:
	rm -f a.out

cpp:
	cppcheck --enable=all --suppress=missingIncludeSystem .

st:
	clang-format -n *.c
	clang-format -i *.c

val:
	valgrind --tool=memcheck --leak-check=yes ./a.out