# Skill Creator Examples

## Example 1: Create a New Skill from Scratch

Build a skill for Docker Compose management. Good agent behavior:

- Define the skill's purpose, triggers, and scope in the frontmatter `description` field.
- Write a clear SKILL.md with overview, when to use, step-by-step workflow, and examples.
- Draft 3 test prompts covering different aspects of Docker Compose workflows.
- Run the skill with each prompt and evaluate qualitative output quality.
- Iterate on the skill based on evaluation results before expanding the test set.

## Example 2: Improve an Existing Skill's Trigger Accuracy

A skill fires too often for unrelated tasks. Good agent behavior:

- Review the skill's `description` frontmatter field for overly broad triggers.
- Narrow the description to specific patterns that should activate the skill.
- Check the `name` field matches the skill's focused domain.
- Create evals that measure false-positive and false-negative rates.
- Compare trigger accuracy before and after the description change using the eval results.

## Example 3: Benchmark and Optimize Skill Performance

A skill works but produces inconsistent results across similar prompts. Good agent behavior:

- Create quantitative evals with clear pass/fail criteria for each expected behavior.
- Run the skill across the eval suite and collect pass rates per test case.
- Analyze variance: identify which test cases show inconsistent results.
- Rewrite the skill's workflow sections that correspond to failing or high-variance tests.
- Re-run the full eval suite and report the improvement in pass rates and variance reduction.