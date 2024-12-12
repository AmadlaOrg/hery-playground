# Dev | Docs | Hery - Playground
## Backend
Simple [Golang](https://go.dev/) HTTP web server.

## Frontend
It is time to put an end to over complicated front end's.

This is why the only two libraries used are:
- [HTMX](https://htmx.org/)
- [Alpine.js](https://alpinejs.dev/)
- [Monaco Editor](https://microsoft.github.io/monaco-editor/)
- [sql.js](https://github.com/sql-js/sql.js/)

## QA
In the `.script/qa-container.sh` there are functions that can be called with `make` command to work with [Playwright](https://playwright.dev/).

Here is a list of the commands:
- `make qa-test` - Run [Playwright](https://playwright.dev/) test
- `make qa-test-show-report` - To get the report
- `make qa-test-ui` - Run [Playwright](https://playwright.dev/) test UI

### Functional Testing
This project uses [Playwright](https://playwright.dev/) to do the functional testing. This is to ensure that the all the pieces work together.

The test can all be run locally if [Playwright](https://playwright.dev/) is installed. Or the examples below show how to use it using a container:

```bash
podman run --rm -it --net=host \
  -v $(pwd):/work \
  -w /work \
  -e DISPLAY=$DISPLAY \
  -v /tmp/.X11-unix:/tmp/.X11-unix \
  mcr.microsoft.com/playwright:v1.49.1-noble \
  yarn playwright test --ui
```

> [!NOTE]
> Podman is used in this example, but it could also be used with Docker.

### Running tests
```bash
podman run --rm -it --net=host \
  -v $(pwd):/work \
  -w /work \
  mcr.microsoft.com/playwright:v1.49.1-noble \
  yarn playwright test
```

In the case of just running the test, there is no need to integrate with the host's X11 server.
