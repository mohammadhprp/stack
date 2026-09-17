# Web Performance Optimization Examples

## Slow landing page

User: "Our landing page takes forever to load. Make it faster."

Good agent behavior:

- Follow the measure-identify-prioritize-implement-verify cycle rather than guessing.
- Start with the highest-impact quick wins: LCP image loading with `fetchpriority="high"`, explicit dimensions, and modern formats.
- Fix font loading with `font-display: swap` and preconnect to the font origin.
- Verify against target metrics (LCP < 2.5s, INP < 200ms, CLS < 0.1) before calling it done.

## Reduce JavaScript bundle size

User: "The main bundle is over 1MB. How do we cut it down?"

Good agent behavior:

- Run a bundle analyzer (webpack or Vite) to find the heavy dependencies before proposing swaps.
- Replace large packages like `moment` with lighter alternatives when the API allows.
- Add route-level code splitting with lazy imports and vendor chunking.
- Defer third-party scripts to user interaction rather than loading them eagerly.
- Confirm the gzipped bundle is under 200KB afterward.

## Fix layout shift on product pages

User: "The images keep jumping around while the page loads."

Good agent behavior:

- Recognize CLS as the metric and add explicit `width`/`height` attributes to all images.
- Reserve space for below-fold content instead of letting it shift on load.
- Lazy load below-the-fold images while keeping the LCP image eager and high priority.
- Verify the CLS score drops below 0.1 after the change.
