# Contributing to Video Compressor

We welcome contributions to the Video Compressor project! This document outlines the steps to get your development environment set up and how to build the application.

## Prerequisites

Before you begin, ensure you have the following installed:

*   **Go**: Version 1.23 or higher. You can download it from [golang.org/dl](https://golang.org/dl/).
*   **Wails CLI**: The Wails command-line interface. Install it using Go:
    ```bash
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```
    For more detailed installation instructions and platform-specific dependencies, refer to the [Wails documentation](https://wails.io/docs/gettingstarted/installation).

## Getting Started

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/hheimerd/video-compressor.git
    cd video-compressor
    ```

2.  **Install Dependencies**:
    Run the setup script to install Go dependencies and download required binaries (ffmpeg/ffprobe):
    ```bash
    ./scripts/install_go_dependencies.sh
    ```

3.  **Run in Development Mode**:
    To run the application in development mode with live-reloading:
    ```bash
    wails dev
    ```
    This command will open the application and automatically recompile and reload the frontend/backend when changes are detected.

4.  **Build for Production**:
    To build a production-ready executable:
    ```bash
    wails build
    ```
    The executable will be generated in the `build/bin` directory.

## Project Structure

*   `app.go`, `main.go`: Backend Go code for the application logic.
*   `frontend/`: Contains the frontend HTML, CSS, and JavaScript.
    *   `frontend/src/index.html`: The main frontend entry point.
*   `wails.json`: Wails project configuration.
*   `go.mod`: Go module definition and dependencies.

Feel free to explore the code, report issues, and submit pull requests!