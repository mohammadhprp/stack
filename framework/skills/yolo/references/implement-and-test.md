# Implement and test

Start only after solution approval. Check that the working tree and applicable instructions still match the assumptions in the plan. Preserve user changes. A material deviation sends the task back to the relevant checkpoint; do not quietly widen the scope.

For each slice:

1. Use the nearest stable public boundary that exposes the required behavior. Derive expected values from the task or an independent example, not the implementation.
2. When practical, run a focused test before changing implementation and confirm it fails for the intended reason. If the only test would be misleading or impractical, report why and choose a credible scenario, script, or manual check.
3. Make the smallest complete change and run the same check again. Test behavior, not private calls or the current layout of the code. Do not weaken existing assertions to get a green run.
4. Clean up temporary instrumentation and run the focused check after any cleanup or refactor. Finish each slice in a verifiable state before moving on.

At the checkpoint, show the real changed paths or diff, evidence from focused tests, deviations, and anything that the final verification still needs to prove. Do not commit or publish in this phase.
