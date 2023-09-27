# GitHub Importer

GitHub Importer is a simple Go application that allows you to upload an XLSX file containing issue details and create GitHub issues. You can also assign users to the created issues and apply labels to them. To use this tool, follow the instructions below.

## Prerequisites

Before running the GitHub Importer, ensure you have the following prerequisites installed on your machine:

1. Go Programming Language: You can download and install Go from the [official website](https://golang.org/dl/).

## Getting Started


### Using Pre-built Binaries
If you don't have Go installed or prefer not to build the application yourself, you can download pre-built binaries for your platform from the [binaries](https://github.com/kishorevaishnav/github-importer/tree/main/bin).

1. Go to the [bin folder](https://github.com/kishorevaishnav/github-importer/tree/main/bin).

2. Download the binary for your operating system and architecture (e.g., `github-importer-linux-amd64` for Linux, `github-importer-darwin-amd64` for macOS, `github-importer-windows-amd64.exe` for Windows).
3. For Linux: Make the binary executable:
   ```shell
   chmod +x github-importer-linux-amd64
    ./github-importer-linux-amd64
   ```
   For MacOS: Make the binary executable:
   ```shell
   chmod +x github-importer-darwin-amd64
    ./github-importer-darwin-amd64
   ```
   For Windows: Just double click `github-importer-windows-amd64.exe`
4. Visit the following URL in your web browser: http://localhost:8080.

### Building from Source

1. Clone the repository to your local machine:
   ```shell
   git clone https://github.com/kishorevaishnav/github-importer.git
2. Change into the project directory:
    ```shell
    cd github-importer
3. Build and run the application:
    ```shell
    go run *.go
4. Visit the following URL in your web browser: http://localhost:8080.

## Usage
Once you have the application running, you can use it to import issues into a GitHub repository by providing the following parameters:

* **Access Token**: Your GitHub Personal Access Token. Make sure it has the necessary permissions to create issues, assign users, and apply labels to the repository.
* **Username**: Your GitHub username.
* **Repository Name**: The name of the target GitHub repository where you want to create issues.
* **XLSX File**: Upload a simple XLSX file (not XLS) containing issue details. The format should match the expected structure.

### Expected XLSX File Format
The XLSX file should have the following columns:

* **Title**: Issue title.
* **Description**: Issue description.
* **Test Steps**: Test Steps to reproduce.
* **Test Data**: Test data to be used for validating the test case.
* **Expected Results**: Expected results once the test case executed and validate.
* **Labels**: Comma-separated list of labels to apply to the issue (optional).
* **Assignee**: GitHub username to assign the issue (optional).

### Example XLSX File
You can find an example [XLSX](sample_test_case.xlsx) file here to see the expected format.

### Contributing
Contributions are welcome! If you find a bug or have an enhancement in mind, please open an issue or create a pull request.
