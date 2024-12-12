#!/usr/bin/env bash

# Version of the Playwright container to use
playwright_version="v1.49.1-noble"

# Function to detect available container engine (Podman or Docker)
# It prioritizes Podman but falls back to Docker if Podman is unavailable.
detect_container_engine() {
  if command -v podman &>/dev/null; then
    echo "podman"
  elif command -v docker &>/dev/null; then
    echo "docker"
  else
    echo "Error: Neither Podman nor Docker is installed." >&2
    exit 1
  fi
}

# Assign the detected container engine to a variable
container_engine=$(detect_container_engine)

# Run Playwright tests in headless mode
test() {
  "$container_engine" run --rm -it \
    --net=host \
    -v "$(pwd)/browser/frontend/qa":/work \
    -w /work \
    mcr.microsoft.com/playwright:"${playwright_version}" \
    yarn playwright test
}

test_show_report() {
  # Allow local access to the X11 server
  #sudo xhost +local:

  "$container_engine" run --rm -it \
    --net=host \
    -v "$(pwd)/browser/frontend/qa":/work \
    -w /work \
    -e DISPLAY=$DISPLAY \
    -v /tmp/.X11-unix:/tmp/.X11-unix \
    mcr.microsoft.com/playwright:"${playwright_version}" \
    yarn playwright show-report

  # Remove local access to the X11 server after the container exits
  #sudo xhost -local:
}

# Run Playwright tests in UI mode, requiring access to X11 server
test_ui() {
  # Allow local access to the X11 server
  #sudo xhost +local:

  "$container_engine" run --rm -it \
    --net=host \
    -v "$(pwd)/browser/frontend/qa":/work \
    -w /work \
    -e DISPLAY=$DISPLAY \
    -v /tmp/.X11-unix:/tmp/.X11-unix \
    mcr.microsoft.com/playwright:"${playwright_version}" \
    yarn playwright test --ui

  # Remove local access to the X11 server after the container exits
  #sudo xhost -local:
}

# Disable X11 server local access explicitly (useful as a cleanup function)
close_xhost() {
  sudo xhost -local:
}

# Dynamic function invocation based on the first argument to the script
# Example usage: ./qa-container.sh test
"$@"
