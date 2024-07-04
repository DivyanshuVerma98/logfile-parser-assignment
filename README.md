![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) &nbsp;
# Log Parser

## Overview

This project is a log parsing utility written in Go. It reads logs from an input file, processes each log entry in parallel, and categorizes them into either a normal log file or an error log file based on the content of the log entry. If a log entry contains the word "ERROR", it is classified as an error log; otherwise, it is classified as a normal log.

## Project Structure

- `main.go`: The main file containing the code for reading, processing, and writing log entries.

## How to Use

1. Clone the repository:
   ```sh
   git clone https://github.com/DivyanshuVerma98/logfile-parser-assignment.git
   ```
2. Run the project:
    ```sh
    go run main.go
    ```
    
## Code Explanation
### Main Functions
1. **LogPraser**:
- Opens the input file for reading logs.
- Creates and opens normal and error log files for writing.
- Spawns goroutines to handle writing to these files.
- Reads the input file line by line, spawning a new goroutine for each log entry to determine if it is an error or normal log.
- Ensures that the channels and wait groups are properly managed to avoid deadlocks.

2. **WriteReport**:
- Continuously reads from a channel and writes the log entries to the specified file.
- Uses a sync.WaitGroup to ensure all goroutines complete their work before the program exits.

## Concurrency and Synchronization
The challenging part of this assignment is handling concurrency and synchronization:
- **Channels**: Used to pass log entries between the main function and the WriteReport goroutines.
- **WaitGroups**: Used to wait for all goroutines to finish their tasks before closing the channels and exiting the program.
- **Buffered Channels**: Channels with a buffer size of 10 are used to avoid blocking the sending goroutines immediately.

By using these synchronization mechanisms, we ensure that all log entries are processed and written to the correct files while maintaining the order of log entries.

## Conclusion
This project demonstrates an efficient way to process and categorize log entries using Go's concurrency features. The primary challenge was ensuring that the order of log entries is maintained while processing them in parallel, which was successfully achieved using channels and wait groups.
