# Laravel Boost Capabilities

## What It Can Do

### Application information

Reads Laravel application details such as PHP version, Laravel version, database engine, installed ecosystem packages, package versions, and Eloquent models.

### Database schema inspection

Inspects database connections and schema so agents can reason about tables, columns, relationships, and constraints before changing code.

### Database queries

Executes database queries through the Laravel application environment. Use this carefully because the tool can access real project data when connected to a real database.

### Route inspection

Inspects application routes so agents can understand controllers, middleware, endpoints, and URL structure.

### Artisan command discovery

Lists and inspects available Artisan commands so agents can find framework and application commands without guessing.

### Tinker execution

Runs suggested PHP code through Laravel's runtime context. This is useful for focused inspection and debugging, but it should be treated as code execution with project permissions.

### Configuration access

Reads configuration keys and values to help agents generate code that matches the application environment.

### Documentation search

Queries Laravel's hosted documentation API using installed package context. This helps agents use current, version-aware Laravel ecosystem guidance.

### Logs and browser errors

Reads application logs, browser errors, and recent errors to support debugging workflows.

### AI guidelines and skills

Generates AI guidelines and optional agent skills for supported Laravel ecosystem packages detected in the project.

## What It Cannot Do

- It is not a generic MCP for non-Laravel projects.
- It does not replace human review for database writes, migrations, authorization changes, or production-impacting commands.
- It does not guarantee that every AI client supports every generated guideline or skill format.
- It cannot safely infer production intent; agents still need explicit instructions before changing data or running destructive commands.
- It cannot fix PHP, Composer, extension, container, or database connectivity issues outside the Laravel project environment.
- It cannot make an unsupported AI client support MCP; unsupported clients require manual configuration or a compatible extension.

## Best Practices

- Install Boost as a development dependency only.
- Run the MCP server from the Laravel application root.
- Use a local or disposable database for agent-assisted exploration whenever possible.
- Review database query and Tinker tool use before allowing write operations.
- Keep generated Boost resources current with `php artisan boost:update`.
- Prefer official Laravel documentation search through Boost when asking Laravel-version-specific questions.
- Keep client-specific config examples separate even when the JSON schema is the same.
- Avoid committing generated local agent configuration if it contains machine-specific paths or settings.

## Common Workflows

### Understand an existing Laravel application

1. Ask the agent to inspect application info.
2. Ask it to inspect routes and models relevant to the task.
3. Ask it to inspect database schema before proposing code changes.
4. Ask it to cite the files and Boost findings used in its plan.

### Debug a failing request

1. Reproduce the issue locally.
2. Ask the agent to inspect recent logs and browser errors.
3. Ask it to map the failing route to controller, middleware, model, and database dependencies.
4. Apply the smallest fix and run regression tests.

### Design a Laravel API change

1. Use Boost to inspect routes, validation patterns, models, and schema.
2. Apply this repository's `api-design`, `database-design`, and `testing` skills.
3. Ask the agent to produce a compatibility and test plan before editing code.
4. Verify with feature tests and targeted static checks.

### Update Laravel ecosystem code

1. Run `php artisan boost:update --discover` after dependency changes.
2. Ask the agent to use Boost documentation search for version-specific guidance.
3. Keep changes small and verify with the project's test suite.
