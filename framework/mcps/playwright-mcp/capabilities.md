# Playwright MCP Capabilities

## What It Can Do

Tools are organized by capability category. Core tools are always available; others require opt-in via `--caps`.

### Core Automation

| Tool | Description | Read-only |
| --- | --- | --- |
| `browser_navigate` | Navigate to a URL | No |
| `browser_navigate_back` | Go back to previous page in history | No |
| `browser_click` | Perform click on a web page | No |
| `browser_hover` | Hover over element on page | No |
| `browser_type` | Type text into editable element | No |
| `browser_fill_form` | Fill multiple form fields | No |
| `browser_select_option` | Select an option in a dropdown | No |
| `browser_drag` | Perform drag and drop between two elements | No |
| `browser_drop` | Drop files or data onto an element | No |
| `browser_press_key` | Press a key on the keyboard | No |
| `browser_snapshot` | Capture accessibility snapshot of the current page | Yes |
| `browser_take_screenshot` | Take a screenshot of the page or element | Yes |
| `browser_evaluate` | Evaluate JavaScript expression on page or element | No |
| `browser_console_messages` | Returns all console messages | Yes |
| `browser_network_requests` | List network requests since page load | Yes |
| `browser_network_request` | Get full details of a specific network request | Yes |
| `browser_file_upload` | Upload one or multiple files | No |
| `browser_handle_dialog` | Accept or dismiss a browser dialog | No |
| `browser_wait_for` | Wait for text to appear/disappear or a set time | No |
| `browser_close` | Close the current page | No |
| `browser_resize` | Resize the browser window | No |
| `browser_run_code_unsafe` | Run arbitrary Playwright code (RCE-equivalent) | No |

### Tab Management

| Tool | Description | Read-only |
| --- | --- | --- |
| `browser_tabs` | List, create, close, or select a browser tab | No |

### Network (opt-in via `--caps=network`)

| Tool | Description | Read-only |
| --- | --- | --- |
| `browser_route` | Set up route to mock network requests | No |
| `browser_route_list` | List all active network routes | Yes |
| `browser_unroute` | Remove network routes | No |
| `browser_network_state_set` | Set browser online/offline | No |

### Storage (opt-in via `--caps=storage`)

| Tool | Description | Read-only |
| --- | --- | --- |
| `browser_cookie_list` | List all cookies | Yes |
| `browser_cookie_get` | Get a specific cookie by name | Yes |
| `browser_cookie_set` | Set a cookie with optional flags | No |
| `browser_cookie_delete` | Delete a specific cookie | No |
| `browser_cookie_clear` | Clear all cookies | No |
| `browser_localstorage_list` | List all localStorage key-value pairs | Yes |
| `browser_localstorage_get` | Get a localStorage item by key | Yes |
| `browser_localstorage_set` | Set a localStorage item | No |
| `browser_localstorage_delete` | Delete a localStorage item | No |
| `browser_localstorage_clear` | Clear all localStorage | No |
| `browser_sessionstorage_list` | List all sessionStorage key-value pairs | Yes |
| `browser_sessionstorage_get` | Get a sessionStorage item by key | Yes |
| `browser_sessionstorage_set` | Set a sessionStorage item | No |
| `browser_sessionstorage_delete` | Delete a sessionStorage item | No |
| `browser_sessionstorage_clear` | Clear all sessionStorage | No |
| `browser_storage_state` | Save storage state (cookies, localStorage) to file | Yes |
| `browser_set_storage_state` | Restore storage state from a file | No |

### DevTools (opt-in via `--caps=devtools`)

| Tool | Description | Read-only |
| --- | --- | --- |
| `browser_annotate` | Open Playwright Dashboard in annotation mode | Yes |
| `browser_hide_highlight` | Remove highlight overlay from an element | No |
| `browser_highlight` | Highlight an element on the page | No |
| `browser_inspect` | Inspect element's computed styles and attributes | Yes |
| `browser_measure` | Perform pixel measurement between elements | Yes |

### Vision (opt-in via `--caps=vision`)

| Tool | Description | Read-only |
| --- | --- | --- |
| `browser_click_position` | Click at specific x/y coordinates | No |
| `browser_drag_position` | Drag from one position to another | No |
| `browser_hover_position` | Hover at specific x/y coordinates | No |

### PDF (opt-in via `--caps=pdf`)

| Tool | Description | Read-only |
| --- | --- | --- |
| `browser_pdf` | Print current page to PDF | Yes |

### Configuration (opt-in via `--caps=config`)

| Tool | Description | Read-only |
| --- | --- | --- |
| `browser_get_config` | Get the final resolved config after merging CLI, env, and config file | Yes |

## What It Cannot Do

- It cannot work without Node.js 18+ and a supported browser engine.
- It does **not** provide a security boundary — the `browser_run_code_unsafe` tool is RCE-equivalent. Rely on client-level permissions for security.
- It cannot access local files outside the workspace unless `--allow-unrestricted-file-access` is set.
- Vision features require the `--caps=vision` flag.
- The Docker image only supports headless Chromium at this time.
- It cannot persist data across isolated sessions when `--isolated` is used.

## Best Practices

- Use `--headless` for CI and automated workflows.
- Use `--isolated` for concurrent agent sessions to avoid profile conflicts.
- Use `--caps` to enable only the tool categories you need — keeps the tool surface small.
- Use `--viewport-size` to set a consistent browser window size for reproducible snapshots.
- Use `--device` to emulate specific mobile devices.
- Use `--user-data-dir` with a persistent profile for logged-in sessions across restarts.
- Use the `--config` file for complex setups instead of long CLI arg lists.
- Do **not** use `browser_run_code_unsafe` unless the client has restricted tool permissions; it is equivalent to arbitrary code execution.

## Common Workflows

### Navigate and inspect a page

1. Call `browser_navigate` to open a URL.
2. Call `browser_snapshot` to capture the accessibility tree.
3. Call `browser_console_messages` to check for errors.
4. Call `browser_network_requests` to inspect network activity.

### Fill and submit a form

1. Call `browser_navigate` to open the page.
2. Call `browser_snapshot` to see the form structure.
3. Call `browser_fill_form` or `browser_type` to enter values.
4. Call `browser_click` to press the submit button.

### Debug a failing test

1. Call `browser_navigate` to the test URL.
2. Call `browser_snapshot` to verify page state.
3. Call `browser_evaluate` to run custom inspection scripts.
4. Call `browser_take_screenshot` to capture visual state.

### Mock API responses

1. Enable `--caps=network` in server args.
2. Call `browser_route` to intercept a URL pattern and return mock data.
3. Call `browser_navigate` to the page that uses the API.
4. Call `browser_network_requests` to verify the mock was used.
