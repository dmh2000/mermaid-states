CC = gcc
CFLAGS = -Wall -Wno-unused-parameter -Wextra -I./include
LDFLAGS = 

SRC_DIR = src
TEST_DIR = test
OBJ_DIR = obj

SRCS = $(wildcard $(SRC_DIR)/*.c)
TEST_SRCS = $(wildcard $(TEST_DIR)/*.c)
OBJS = $(SRCS:$(SRC_DIR)/%.c=$(OBJ_DIR)/%.o)
TEST_OBJS = $(TEST_SRCS:$(TEST_DIR)/%.c=$(OBJ_DIR)/%.o)

TEST_BINARY = test_state_machine

.PHONY: all clean test

all: $(TEST_BINARY)

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c
	@mkdir -p $(OBJ_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(OBJ_DIR)/%.o: $(TEST_DIR)/%.c
	@mkdir -p $(OBJ_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(TEST_BINARY): $(OBJS) $(TEST_OBJS)
	$(CC) $(OBJS) $(TEST_OBJS) $(LDFLAGS) -o $@

test: $(TEST_BINARY)
	./$(TEST_BINARY)

clean:
	rm -rf $(OBJ_DIR) $(TEST_BINARY)
