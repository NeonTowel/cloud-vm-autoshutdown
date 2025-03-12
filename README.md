# Cloud VM Auto Shutdown

This project contains scripts for automatically shutting down virtual machines when they are idle. Supported clouds are `GCP` and `Azure`. The functionality is implemented in both `Go` and `Bash`.

## Table of Contents
- [Overview](#overview)
- [Functionality](#functionality)
- [Azure: Go Implementation (auto_shutdown/cmd/azure)](#go-implementation-auto_shutdowncmdazure)
- [GCP: Go Implementation (auto_shutdown/cmd/gcp)](#go-implementation-auto_shutdowncmdgcp)
- [Legacy GCP Bash Implementation (auto_shutdown.sh)](#bash-implementation-auto_shutdownsh)
- [Usage](#usage)
- [Development (Go)](#development-go)
- [Auto Shutdown Usage in GCP](#auto-shutdown-usage-in-gcp)
- [Customization](#customization)
- [Note](#note)

## Overview

The auto-shutdown scripts monitor the system load and user activity on a GCE or Azure VM. If the system remains idle for a specified period, the scripts initiate a shutdown sequence. This helps to save resources and reduce costs by turning off unused VMs.

## Functionality

Core features:

1. **GCE or Azure VM Check**: Verifies that the script is running on a GCE or Azure VM.
2. **System Load Monitoring**: Checks the 5-minute load average.
3. **User Activity Monitoring**: Tracks SSH connections and logged-in users.
4. **Idle Time Tracking**: Counts consecutive idle intervals.
5. **Shutdown Sequence**: Initiates a shutdown after a specified idle period.

### Key Parameters

- `threshold`: The system load threshold below which the system is considered idle.
- `intervals`: The number of consecutive idle checks required before shutdown.
- `sleepTime`: The duration between each check (in seconds).

## Azure: Go Implementation (auto_shutdown/cmd/azure)

The Azure version is structured similarly to the GCP implementation, providing a modular and type-safe approach:

1. **Package Structure**: The core functionality is implemented in the `azure` package located under `auto_shutdown/pkg/azure`.
2. **Main Function**: The `main.go` file in `auto_shutdown/cmd/azure` calls the `MonitorAndShutdown()` function from the `azure` package.
3. **Concurrency and System Calls**: Utilizes Go's concurrency features for efficient monitoring and shutdown operations.

### Building and Running

1. **Build the Binary**:

    ```bash
    task build-azure
    ```

2. **Run the Binary**:

    ```bash
    task run-azure
    ```

3. **Development**:

    - To run the source code directly:

        ```bash
        task dev-azure
        ```


## GCP: Go Implementation (auto_shutdown/cmd/gcp)

The Go version is structured to provide a modular and type-safe implementation:

1. **Package Structure**: The core functionality is implemented in the `gcp` package located under `auto_shutdown/pkg/gcp`.
2. **Main Function**: The `main.go` file in `auto_shutdown/cmd/gcp` calls the `MonitorAndShutdown()` function from the `gcp` package.
3. **Concurrency and System Calls**: Utilizes Go's concurrency features for efficient monitoring and shutdown operations.

### Building and Running

1. **Build the Binary**:

    ```bash
    task build-gcp
    ```

2. **Run the Binary**:

    ```bash
    task run-gcp
    ```

3. **Development**:

    - To run the source code directly:

        ```bash
        task dev-gcp
        ```

## Legacy GCP Bash Implementation (auto_shutdown.sh)

The Bash script provides a lightweight solution:

1. Uses environment variables or default values for configuration.
2. Leverages system commands like `uptime`, `ss`, and `who` for monitoring.
3. Uses a while loop for continuous monitoring.
4. Implements simple arithmetic for counting idle intervals.

## Usage

### Go Version

1. Compile the Go script:

```bash
go build -o auto_shutdown auto_shutdown.go

# or if you have go-task installed:
task build
```

2. Run the compiled binary:

```bash
./auto_shutdown
```

### Bash Version

1. Make the script executable:

```bash
chmod +x auto_shutdown.sh
```

2. Run the script:

```bash
./auto_shutdown.sh
```

## Development (Go)

To develop and build the Go version:

1. Ensure you have `Go` (version 1.23.1 or later) installed on your system. You can download it from https://golang.org/dl/.

2. Clone the repository or navigate to the project directory:

    ```bash
    cd cloud/gcp/compute
    ```

5. To build the Go script, run:

    ```bash
    go build -o auto_shutdown auto_shutdown.go
    ```

6. For development, you can use the following commands:

   - To run the script without compiling:

        ```bash
        go run auto_shutdown.go
        ```

   - To format your code:

        ```bash
        go fmt auto_shutdown.go
        ```

   - To run tests (if you've written any):

        ```bash
        go test
        ```

7. To compile you can use:


    ```bash
    GOOS=linux GOARCH=amd64 go build -o auto_shutdown auto_shutdown.go
    ```

### Using go-task for task automation

If you're using `go-task` for task automation, then you can use:

In addition to the `build` task, the [Taskfile.yaml](./Taskfile.yaml) includes:

1. `build`: Compiles the `auto_shutdown.go` script to executable binary:

    ```bash
    task build
    ```

2. `dev`: Runs `auto_shutdown.go` directly for quick testing:

    ```bash
    task dev
    ```

3. `run`: Executes the compiled binary:

    ```bash
    task run
    ```

4. `clean`: Removes the built binary:

    ```bash
    task clean
    ```

These tasks streamline development, execution, and cleanup of the auto_shutdown project.
Remember to handle errors appropriately, implement logging for debugging, and consider adding unit tests for critical functions to ensure reliability.

#### How to install go-task

To download and install go-task:

   - For macOS (using Homebrew):

        ```bash
        brew install go-task/tap/go-task
        ```

   - For Linux:

        ```bash
        sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
        ```

   - For Windows (using Scoop):

        ```bash
        scoop bucket add extras
        scoop install task
        ```

   For other installation methods or more details, visit the official go-task GitHub repository:
   https://github.com/go-task/task#installation

## Auto Shutdown Usage in GCP

To set the binary or script to be executable when the GCP VM starts using VM metadata options:

1. Upload the compiled binary or Bash script to your VM (e.g., to `/opt/scripts/auto_shutdown` or `/opt/scripts/auto_shutdown.sh`).

2. To set the metadata when creating a new VM instance:

    ```bash
    gcloud compute instances create INSTANCE_NAME \
        --metadata-from-file startup-script=local_path_to_startup_binary_or_script.sh
    ```

3. To update the metadata for an existing VM instance:

    ```bash
   gcloud compute instances add-metadata INSTANCE_NAME \
     --metadata-from-file startup-script=local_path_to_startup_binary_or_script.sh
    ```

Replace `INSTANCE_NAME` with the name of your VM instance.

## Customization

Both versions allow customization of the threshold, intervals, and sleep time. Modify these values in the script to adjust the shutdown behavior according to your needs.

## Note

Ensure that the script has the necessary permissions to execute system commands and initiate a shutdown. It's recommended to run these scripts with appropriate privileges on your GCE VMs.
