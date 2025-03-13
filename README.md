# Linux VM Auto Shutdown

This project contains Go implementation for automatically shutting down Linux VMs when they are idle. Designed to be run as SystemD Service on boot.

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

The auto-shutdown app monitor the system load and user activity on a Linuxx VM. If the system remains idle for a specified period, the app initiate a shutdown sequence. This helps to save resources and reduce costs by turning off unused VMs.

## Supported Linux distributions

The script should work reliably on common Linux distributions where the `/proc` filesystem, `iproute2`, `util-linux`, and basic `GNU utilities` (coreutils) are available. This includes distributions like `Ubuntu`, `Debian`, `CentOS`, `Fedora`, and their derivatives.

For further compatibility, ensure the availability of these utilities and their correct paths if scripts use absolute paths for execution on specialized or minimal environments.

## Functionality

Core features:

1. **System Load Monitoring**: Checks the 5-minute load average.
2. **User Activity Monitoring**: Tracks SSH connections and logged-in users.
3. **Idle Time Tracking**: Counts consecutive idle intervals.
4. **Shutdown Sequence**: Initiates a shutdown after a specified idle period.

### Key Parameters

- `threshold`: The 5 minute system load average threshold below which the system is considered idle.
- `intervals`: The number of consecutive idle checks required before shutdown.
- `sleepTime`: The duration between each check (in seconds).

These can now be configured as system environment variables, with fallback to default settings if environment variables are not set:

**Environment variables:** 

- `SHUTDOWN_THRESHOLD`
- `SHUTDOWN_INTERVALS`
- `SHUTDOWN_SLEEP_TIME`

**Default values:**

- `threshold` = 0.15
- `intervals` = 15
- `sleepTime` = 30


## Universal implementation (for SystemD)

The Universal version is structured using a modular and type-safe approach:

1. **Package Structure**: The core functionality is implemented in the `universal` package located under `auto_shutdown/pkg/universal`.
2. **Main Function**: The `main.go` file in `auto_shutdown/cmd/universal` calls the `MonitorAndShutdown()` function from the `universal` package.
3. **Concurrency and System Calls**: Utilizes Go's concurrency features for efficient monitoring and shutdown operations.

### Building and Running

1. **Build the Binary**:

    ```bash
    task build-universal
    ```

2. **Run the Binary**:

    ```bash
    task run-universal
    ```

3. **Development**:

    - To run the source code directly:

        ```bash
        task dev-universal
        ```

## Deprecated GCP and Azure specific implementations

The Go implementations in `auto_shutdown/cmd/gcp` and `auto_shutdown/cmd/azure` are now deprecated in favor of the universal SystemD approach.

The implementations are provided as a reference. The tasks to build these implementations remain in the taskfile for now. Examine the `Taskfile.yaml` for details.

## Deprecated shell scripts

The shell script implementation in `scripts/auto_shutdown.sh` (which was the original implementation for GCP) remains for reference.

## Development (Go)

To develop and build the Go version:

1. Ensure you have `Go` (version 1.24.1 or later) installed on your system. You can download it from https://golang.org/dl/. Note: This is the Go version used in development, earlier versions may work perfectly fine.

2. Clone the repository.

5. To build the universal binary, run:

    ```bash
    task build-universal
    ```

    Binary can be found from `./build` directory.

6. For development, you can use the following commands:

   - To run the script without compiling:

        ```bash
        task dev-universal
        ```

   - To format your code:

        ```bash
        task fmt
        ```

### Using go-task for task automation

If you're using `go-task` for task automation, then you can use:

In addition to the `build` task, the [Taskfile.yaml](./Taskfile.yaml) includes:

1. `task build`: Compiles the executable binaries.
2. `task dev-universal`: Runs `auto_shutdown/cmd/universal/main.go` directly for quick testing.
3. `task run-universal`: Runs `build/auto_shutdown_universal` binary, if it exists.
4. `task clean`: Removes the built binaries from `./build` path.

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

Ensure that the script has the necessary permissions to execute system commands and initiate a shutdown. It's recommended to run these scripts with appropriate privileges VMs.
