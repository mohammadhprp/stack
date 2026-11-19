# Diagnose hard failures

Use this inside an approved phase when the root cause or feasibility is unclear. Prefer read-only investigation and isolated, disposable reproductions. During solution design, do not edit the project to make a hypothesis easier to test.

- State what observation would establish the problem and what would show it is gone. Find the smallest reliable signal for the reported behavior, not just a nearby failure.
- Rank plausible explanations. Make a prediction for each before testing it. Change one variable at a time and inspect the actual outcome, not the confidence of a report.
- For hard-to-reproduce cases, record the conditions, frequency, and limits of the signal. If you cannot get a credible baseline, say what remains unknown.
- On a long investigation, record decisions and evidence in the task directory. Remove temporary instrumentation from the project after approved implementation; rerun the original check and document any gap.
