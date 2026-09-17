# Laravel Boost Troubleshooting

## Installation failure

**Problem**

`composer require laravel/boost --dev` fails.

**Cause**

Composer cannot resolve dependencies, PHP is missing, PHP extensions are missing, or the command is not being run from a Laravel application root.

**Solution**

- Run `php -v` and `composer --version`.
- Run the command from the Laravel application root.
- Resolve Composer dependency conflicts shown in the error output.
- Install missing PHP extensions required by the Laravel application.

## Server won't start

**Problem**

The AI client reports that `laravel-boost` failed to start.

**Cause**

The client cannot run `php artisan boost:mcp`, the working directory is wrong, dependencies are not installed, or the Laravel app fails during bootstrap.

**Solution**

- Run `php artisan boost:mcp` manually from the Laravel application root.
- Fix the first Laravel, PHP, Composer, or environment error shown.
- Ensure `vendor/` exists by running `composer install`.
- Configure the client to start the server from the Laravel application root.

## Configuration errors

**Problem**

The MCP server is not listed, or the client ignores the config file.

**Cause**

The config file is in the wrong location, the JSON is invalid, or the client expects a different configuration wrapper.

**Solution**

- Validate the JSON with `python -m json.tool <file>` or another JSON parser.
- Confirm the client-specific config location in the client's documentation.
- Keep the server name as `laravel-boost`.
- Use the official command `php` and args `artisan`, `boost:mcp` unless your local environment requires a wrapper.

## Authentication issues

**Problem**

Boost tools cannot access application features that depend on authenticated users or external services.

**Cause**

MCP tools run in the local Laravel application context and do not automatically create browser sessions, API tokens, or third-party service credentials.

**Solution**

- Use local test accounts and seeded data when possible.
- Provide required `.env` values for local-only services.
- Avoid exposing production secrets to AI clients.
- Prefer tests or factories for authenticated workflows.

## Permission problems

**Problem**

Boost cannot read logs, write generated files, or access project paths.

**Cause**

Filesystem ownership, container mounts, or restricted directories prevent PHP or the AI client process from accessing required files.

**Solution**

- Check ownership and permissions for `storage/`, `bootstrap/cache/`, generated agent files, and client config files.
- Run the AI client and PHP process as a user with access to the project.
- In containerized environments, ensure the project is mounted and writable where needed.

## Version mismatch

**Problem**

Boost behavior or generated files do not match the documentation or the selected AI client.

**Cause**

The installed `laravel/boost` version, Laravel version, AI client version, or MCP support differs from the documented version.

**Solution**

- Check installed versions with Composer and the AI client's version command.
- Run `composer update laravel/boost --dev` when appropriate.
- Run `php artisan boost:update --discover` after dependency changes.
- Re-run `php artisan boost:install` if client support files need to be regenerated.

## Database query risk

**Problem**

An agent proposes running write queries or destructive Tinker code through Boost.

**Cause**

Boost exposes tools that can execute code or queries in the Laravel application context.

**Solution**

- Use a local or disposable database for agent work.
- Require human approval before write queries, migrations, seeders, or destructive commands.
- Ask the agent to show the exact query or code before execution.
- Back up important data before allowing inspection against non-disposable databases.
