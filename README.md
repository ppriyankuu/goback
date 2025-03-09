# goBack
`goback` is a simple command-line backup tool written in Go. It scans files and directories, compresses backups to save space, and supports incremental backups by only storing changed files. You can easily restore files when needed, and the tool provides progress tracking to keep you updated. It’s lightweight, efficient, and does the job without any hassle.

## Features 
1. `Backup Basics` 
    - Scans files and folders, compresses backups (gzip, zip), supports incremental backups, restores   files when needed, and shows progress.
2. `Simple CLI`
    - Uses command-line arguments, provides help docs, supports config files, shows progress, and reports errors.
3. `File Handling` 
    - Goes through directories, tracks file changes, filters files, handles metadata, and keeps permissions intact.
4. `Storage Management`
    - Creates backup archives, tracks versions, stores metadata, and verifies backups.

## Folder Structure
```bash
goBack/
├── cmd/                               // Contains the main entry point of the application
│   └── main.go                        // Main entry point for the CLI application
├── internal/                          // Internal packages for various functionalities
│   ├── backup/                        // Backup-related functionalities
│   │   ├── backup.go                  // Core backup functionality
│   │   ├── incremental.go             // Incremental backup functionality
│   │   ├── restore.go                 // Restore functionality
│   │   └── version.go                 // Version tracking and reporting
│   ├── cli/                           // Command-line interface functionalities
│   │   ├── help.go                    // Help documentation
│   │   ├── config.go                  // Configuration file handling
│   │   └── progress.go                // Progress tracking and reporting
│   ├── fs/                            // File system operations
│   │   ├── traversal.go               // Directory traversal
│   │   ├── metadata.go                // File metadata handling
│   │   ├── change_detection.go        // Change detection
│   │   ├── filtering.go               // File filtering and exclusion
│   │   └── permissions.go             // Permission preservation
│   └── storage/                       // Storage management
│       ├── archive.go                 // Backup archive creation and extraction
│       ├── metadata.go                // Metadata storage and retrieval
│       ├── verification.go            // Backup verification
│       └── cleanup.go                 // Cleanup of old backups
├── config.yaml                        // Configuration file for the utility
└── go.mod                             // Go module configuration

```

## Installation 
To install the `goBack`, you need to have Go installed on your system.

1. Clone the repository
    ```bash
    git clone https://github.com/ppriyankuu/goback.git
    cd goback
    ```
2. Build the binary
    ```bash
    go build -o goback cmd/main.go
    ```

## Configuration 
The utility uses a YAML configuration file named `config.yaml`. You can customize the retention days for old backups in this file. 
### Example `config.yaml`
```bash
retention_days: 7
```

## Usage
#### Basic Commands
- Backup
    ```bash
    goback -s /path/to/source -d /path/to/destination
    ```
- Incremental Backup
    ```bash
    goback -s /path/to/source -d /path/to/destination -i
    ```
- Restore
    ```bash
    goback -s /path/to/source -d /path/to/destination -r
    ```
- Help
    ```bash
    goback -h
    ```

#### Options 
- `-c, --config <file>`: Path to the configuration file (default: config.yaml).
- `-s, --source <dir>`: Source directory to back up.
- `-d, --destination <dir>`: Destination directory for backups.
- `-i, --incremental`: Enable incremental backup.
- `-r, --restore`: Restore from backup.
- `-h, --help`: Show help documentation.
     
## Contributing 
Contributions are always welcome! If you find any bugs or have feature requests, just open an issue or submit a pull request.

1. Fork the repository .
2. Create a new branch :
    ```bash
    git checkout -b my-feature
    ```
3. Make your changes .
4. Commit your changes :
    ```bash
    git commit -m "Add my feature"
    ```
5. Push to the branch :
    ```bash
    git push origin my-feature
    ```
6. Open a pull request.

<hr />

```
Built with 90% AI, and 10% Brain (not skills :)
```