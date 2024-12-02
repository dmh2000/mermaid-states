CXX = g++
CXXFLAGS = -std=c++17 -Wall -Wextra -I.
LDFLAGS =

SOURCES = test_state_machine.cpp
OBJECTS = $(SOURCES:.cpp=.o)
TARGET = test_state_machine

.PHONY: all clean test

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CXX) $(OBJECTS) -o $(TARGET) $(LDFLAGS)

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@

test: $(TARGET)
	./$(TARGET)

clean: header
	rm -f $(OBJECTS) $(TARGET)
