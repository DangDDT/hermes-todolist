# Figma Adoption Plan

## 1. Purpose
Adopt **Figma** as the source of truth for all UX/UI work on Hermes TodoList from the next design milestone onward.

This plan defines how we move from planning artifacts and temporary review prototypes to a proper design workflow that is:
- cross-platform
- reusable
- traceable
- implementation-friendly
- consistent across desktop, tablet, and mobile

## 2. Why Figma
Figma is the right design hub for this project because it supports:
- a shared design system
- components and variants
- page-level composition
- flow/prototype linking
- comments and review cycles
- handoff to implementation

## 3. Design Workflow Order
All UX/UI work should follow this sequence:

1. **System design**
   - design principles
   - design tokens
   - spacing scale
   - typography scale
   - color semantics
   - state semantics

2. **Components**
   - task card
   - badges
   - buttons
   - form fields
   - dialogs
   - empty states
   - loading states

3. **Pages**
   - auth
   - task list
   - task detail
   - create/edit task
   - settings / future pages

4. **Flows**
   - login/register
   - browse/filter/sort tasks
   - create task
   - edit task
   - delete/confirm flow

5. **Scenarios**
   - empty state
   - happy path
   - error recovery
   - low-data / high-density usage
   - mobile-first usage

6. **Motion / animation**
   - page transitions
   - state changes
   - feedback moments
   - subtle interaction polish

## 4. Figma File Structure
Recommended workspace structure:

- **00 - Cover / Notes**
- **01 - System Design**
- **02 - Components**
- **03 - Desktop Screens**
- **04 - Tablet Screens**
- **05 - Mobile Screens**
- **06 - User Flows**
- **07 - Scenarios**
- **08 - Motion / Prototype**
- **09 - Handoff / Specs**

## 5. Cross-Platform Rules
The product must keep one unified design language across platforms:
- same semantic tokens
- same component behavior
- same color meaning for status/priority
- same spacing and hierarchy rules
- same form validation language
- same interaction logic, adapted to form factor

Platform differences should be structural, not stylistic.

## 6. Naming Conventions
Use stable names so the file stays navigable:

- `DS / Color`
- `DS / Type`
- `DS / Spacing`
- `CMP / Task Card`
- `CMP / Status Badge`
- `CMP / Priority Badge`
- `PGE / Task List`
- `PGE / Task Detail`
- `FLW / Create Task`
- `SCN / Empty State`
- `MOT / Transition Baseline`

## 7. Review Rules
Each review cycle should ask:
- Does the system hold together across all screens?
- Are the components reusable and consistent?
- Do the pages feel like one product?
- Does the flow minimize friction?
- Are scenarios handled gracefully?
- Does motion help clarity instead of adding noise?

## 8. Handoff Rules
Figma is the design source of truth, but implementation should still be traceable.

Handoff expectations:
- each screen is linked to a component basis
- component variants are documented
- design intent is clear for edge cases
- states are visible for loading, empty, error, success
- acceptance criteria are written for important screens

## 9. Milestone Plan
### Milestone A — Foundation
- set up Figma file structure
- define system tokens
- establish shared components

### Milestone B — Core Screens
- design auth
- task list
- task detail
- create/edit flow

### Milestone C — Prototype Polish
- wire flows
- add scenarios
- refine motion and transitions

### Milestone D — Implementation Sync
- align dev implementation to Figma system
- close gaps between design and code
- update docs if system rules change

## 10. Acceptance Criteria
The Figma adoption is considered successful when:
- the design system is defined in Figma
- components are reusable and consistent
- major pages exist in desktop/tablet/mobile variants
- flows are linked and reviewable
- scenarios are represented
- motion baseline exists
- implementation can follow the file without ambiguity

## 11. Temporary HTML Prototype Rule
Temporary HTML prototypes may still be used for quick review, but they are **not the design source of truth**.

Use them only when:
- a fast review artifact is needed
- Figma is not yet fully prepared for the milestone
- the goal is to validate a concept quickly

Once the Figma workflow is ready, Figma replaces HTML as the primary design artifact.

## 12. Outcome
This plan makes UX/UI work consistent, collaborative, and scalable while keeping the product experience aligned across desktop, tablet, and mobile.
