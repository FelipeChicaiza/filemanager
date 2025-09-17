# File Manager CLI (Case Study for DFS Project)

## 📖 Overview
This project is a **simple command-line file manager** built in Go.  
It serves as the **first step in my journey toward building a distributed file system (DFS)**.  
By implementing basic file operations locally, I am building the foundation for understanding how storage, file handling, and metadata management work before tackling networking, replication, and fault tolerance in a full DFS.  

## 🛠 Features
The CLI supports basic file operations:
- **Create**: Create a file with specified content.  
- **Read**: Read and display the contents of a file.  
- **Delete**: Remove a file from the filesystem.  
- **List**: List all files in the current directory.  

### Example Usage
```bash
# Create a file with content
go run main.go create test.txt "Hello World"

# Read a file
go run main.go read test.txt

# Delete a file
go run main.go delete test.txt

# List all files in directory
go run main.go list
