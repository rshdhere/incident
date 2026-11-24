# Coding

Put the v1 server, services, domain types and persistence code in `packages/api/v1`. Register authored incidents explicitly from `.aws` and keep worked responses separate from the initial exercise.

Start web features in `apps/web` and follow the route's existing imports. Move a component into `packages/ui` when sharing is part of the change and check its callers in both apps. Keep product data and incident-specific interactions close to their owning feature.

The shared TypeScript preset enables strict checking and checked indexed access. Model a missing scenario, evidence item or selection explicitly when implementing navigation. Use a client component at the interactive boundary for browser events or browser state. The web home renders a shared button whose click handler lives in a client component.

Style route-local elements in CSS modules. Keep global changes in the relevant app's global stylesheet. Check keyboard access and readable text when changing canvas interactions.
