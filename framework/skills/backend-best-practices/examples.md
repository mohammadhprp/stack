# Backend Best Practices Examples

## Example 1: New Checkout API Endpoint

Design and implement a checkout endpoint that creates an order, reserves inventory, and charges a payment method. Good agent behavior:

- **API Design**: Model the endpoint around the checkout business operation, not database tables. Require an idempotency key so retries don't create duplicate orders. Return 201 with the order and a Location header.
- **Database Design**: Design the order line items schema with FK constraints, unique constraint on idempotency key, and indexes on order status and customer ID. Plan the migration with forward, rollback, and verification steps.
- **Security**: Authorize that the authenticated user owns the payment method and shipping address. Validate all inputs at the boundary. Never leak payment details in logs or error messages.
- **Testing**: Write unit tests for order creation rules, integration tests for the database constraints and concurrent reservation, and a contract test for the API boundary.
- **Observability**: Add structured logging for each checkout step, a metric tracking checkout success/failure rates, and a health check that verifies the payment gateway connectivity.

## Example 2: Fix Slow Product Search

Users report that the product search page takes over 5 seconds to load. Good agent behavior:

- **Debugging**: Reproduce in staging. Correlate slow requests via trace IDs. Profile the endpoint to identify the bottleneck — an N+1 query pattern in the search results loop.
- **Performance**: Measure baseline p50/p95/p99. Add eager loading to eliminate N+1 queries. Add a database index on the search columns. Re-measure and confirm improvement.
- **Refactoring**: Extract the search logic into a dedicated service class. Add characterization tests for the existing behavior before refactoring. Make one change at a time.
- **Testing**: Add a regression test that asserts the query count stays bounded. Add a performance test that fails if p95 exceeds a threshold.
- **Observability**: Add a dashboard panel for search latency and a trace span for the search query.

## Example 3: Architecture Review for Monolith Extraction

The team wants to extract the billing module from the monolith into a separate service. Good agent behavior:

- **Architecture Review**: Map the current component boundaries and data flows. Evaluate coupling between billing and other modules. Identify single points of failure and data ownership boundaries.
- **API Design**: Design the new billing service API contract with versioning, pagination for invoice listing, and idempotency for payment operations.
- **Security**: Review authentication between services (mTLS or API tokens). Ensure the billing service has least privilege access to customer data.
- **Database Design**: Plan the data migration from the monolith database to the new service's database with backfill, verification, and rollback steps.
- **Testing**: Write contract tests for the new service API. Add integration tests for the data migration. Verify end-to-end that invoice generation works across the service boundary.
