cmake_minimum_required(VERSION 3.10)
project(state_machine)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Add GoogleTest
include(FetchContent)
FetchContent_Declare(
    googletest
    GIT_REPOSITORY https://github.com/google/googletest.git
    GIT_TAG release-1.12.1
)
FetchContent_MakeAvailable(googletest)

# Include directories
include_directories(${PROJECT_SOURCE_DIR}/include)

# Add test executable
add_executable(state_machine_test test/state_machine_test.cpp)
target_link_libraries(state_machine_test gtest gtest_main)

# Enable testing
enable_testing()
add_test(NAME state_machine_test COMMAND state_machine_test)
