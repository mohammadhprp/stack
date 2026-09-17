# Laravel Best Practices Examples

## Example 1: Refactor Controller Logic

A controller action contains five responsibilities: validation, authorization, querying, formatting, and logging. Good agent behavior:

- Extract validation into a Form Request class with `authorize()` and `rules()` methods.
- Move authorization checks into a Policy class with a named method matching the action.
- Extract query logic into a repository or service class with a single responsibility.
- Use API resources for response formatting instead of inline `toArray()` calls.
- Use Laravel's built-in logging channel with structured context instead of `Log::debug()` scattered through the controller.

## Example 2: Fix N+1 Query Performance

A page loads slowly because a loop calls a relationship query for every item. Good agent behavior:

- Identify the N+1 pattern: the loop accesses `$post->comments` inside a Blade `@foreach`.
- Use `$posts->load('comments')` for eager loading, or use `with('comments')` on the initial query.
- Check for deeper nested relationships that would cause additional N+1 patterns.
- Add `select` and `where` constraints on the eager-loaded relationship to avoid fetching unnecessary data.
- Verify the fix reduces query count from N+1 to 2 with `DB::enableQueryLog`.

## Example 3: Secure API Endpoint

An API endpoint allows any authenticated user to delete any order. Good agent behavior:

- Create a Policy with a `delete` method that verifies the order belongs to the user.
- Register the policy in `AuthServiceProvider`.
- Use `$this->authorize('delete', $order)` in the controller instead of inline `if` checks.
- Add a Form Request to validate the delete action input (order ID, reason).
- Write a unit test for the policy denying access to orders owned by other users.