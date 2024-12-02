CC = gcc
CFLAGS = -Wall -Wextra -I.
LDFLAGS =

LIB_SOURCES = state_machine.c
LIB_OBJECTS = $(LIB_SOURCES:.c=.o)
LIB = libstatemachine.a

TEST_SOURCES = test_state_machine.c
TEST_OBJECTS = $(TEST_SOURCES:.c=.o)
TEST = test_state_machine

.PHONY: all clean test

all: $(LIB) $(TEST)

$(LIB): $(LIB_OBJECTS)
	ar rcs $@ $^

$(TEST): $(TEST_OBJECTS) $(LIB)
	$(CC) $(TEST_OBJECTS) -L. -lstatemachine -o $@ $(LDFLAGS)

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

test: $(TEST)
	./$(TEST)

clean: header
	rm -f $(LIB_OBJECTS) $(TEST_OBJECTS) $(LIB) $(TEST)
