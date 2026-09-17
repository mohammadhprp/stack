# Taste Skill — Example Outputs

Realistic examples of how the `taste` skill routes a brief to the right reference and enforces its standards.

## Example 1: Landing Page (routes to `design-taste-frontend`)

**User:** "Build a landing page for a developer analytics platform. Clean, premium, not AI-sloppy."

**Design read:** B2B SaaS landing for technical buyers, with a Linear-style minimalist language, leaning toward Tailwind utilities + Geist + restrained motion.

**What the agent does (following `design-taste-frontend.md`):**

1. Sets the dials from the read: VARIANCE 5, MOTION 4, DENSITY 3.
2. States the design read in one line before writing any code.
3. Chooses a non-default hero: centered over a wide headline with two high-contrast CTAs and a full-bleed background image with a dark radial wash — not a narrow 6-line text wall.
4. Bans Inter and Lucide; uses Geist + Phosphor. No AI-purple gradient, no three equal feature cards — uses an asymmetric bento with `grid-flow-dense`.
5. Applies `min-h-[100dvh]` (never `h-screen`), animates only `transform`/`opacity`, generous `py-32 md:py-48` section spacing.

**Result:** A distinctive, premium landing that does not read as a template.

## Example 2: Brand Kit (routes to `brandkit`)

**User:** "Generate a brand-kit board for 'Aster', a security platform for cloud-native startups."

**What the agent does (following `brandkit.md`):**

1. Infers brand strategy first: category security, audience CTOs/security engineers, core metaphor "vigilance + protection".
2. Picks the Dark Security / Threat Intelligence visual mode: black/navy canvas, shield forms, radar lines, red/blue alert chips.
3. Chooses a `3 × 3` layout with strong gutters and sparse typography: logo cover, logo construction grid, browser mockup, tagline, color system, typography, physical application, image direction, system detail.
4. Designs one strong mark — an eye formed by crossing shield shapes — repeated consistently across panels. No generic lightning bolts, no fake luxury crests.
5. Writes a short, specific tagline ("Your network. Our watch."), one URL, and minimal labels. No lorem ipsum, no dense body copy.

**Result:** A premium, presentation-ready identity-system board that feels like a serious studio deck.

## Example 3: Mobile App Screens (routes to `imagegen-frontend-mobile`)

**User:** "Show me the onboarding flow for a habit-tracking app — three screens."

**What the agent does (following `imagegen-frontend-mobile.md`):**

1. Generates three separate app-native screens (welcome, permissions, home) inside a subtle premium phone mockup with a visible frame.
2. Keeps hierarchy clean and text comfortably readable — no tiny fake labels, no dense dashboards.
3. Uses one controlled palette with textured surfaces and image-led composition.
4. Keeps the three screens visually consistent (shared color, type scale, iconography).
5. Produces images only — no code.

**Result:** Three cohesive, premium, app-native concepts a designer can hand to an engineer.

## Example 4: Redesign (routes to `redesign-existing-projects`)

**User:** "The dashboard we built last month looks generic. Make it premium without breaking it."

**What the agent does (following `redesign-existing-projects.md`):**

1. **Scan** — identifies the stack (React + Tailwind v4) and current patterns.
2. **Diagnose** — audits: Inter everywhere, three equal card columns, pure black background, purple-blue gradient hero, no hover/active states, no loading or empty states, generic Lucide icons.
3. **Fix in priority order** — font swap first (Geist), then palette cleanup (off-black + one desaturated accent), then hover/active states, layout, and finally component and icon swaps.
4. Works with the existing stack — no framework migration, no rewrites. Tests after each change.

**Result:** The same app, elevated — no broken functionality, no templated look.

## Example 5: Full Output (routes to `full-output-enforcement`)

**User:** "Write the complete `auth.ts` service — every function, no shortcuts."

**What the agent does (following `full-output-enforcement.md`):**

1. Counts the deliverables (all auth functions + exports) and locks the scope.
2. Emits the full file — no `// rest of code`, no `// TODO`, no `...` placeholders, no "the rest follows the same pattern."
3. If the response approaches the token limit, stops at a clean breakpoint and marks it: `[PAUSED — 2 of 4 complete. Send "continue" to resume from: token refresh]`.
4. Cross-checks output against the original request before finalizing.

**Result:** A complete, runnable file with zero placeholder patterns.