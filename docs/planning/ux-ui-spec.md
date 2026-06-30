# UX/UI Specification

## 1. Design Principles
- simple first
- low cognitive load
- clear priority signaling
- fast scanability
- production-clean, not toy-like

## 2. Core Experience Goals
- user sees what matters in under 5 seconds
- adding a task feels lightweight
- changing status is obvious
- priority is visually clear without being noisy

## 3. Information Architecture
### Primary Navigation
- Tasks
- Settings

### Task Page Sections
- page header
- filters/search/sort
- task list content
- pagination controls
- create action

## 4. Visual Hierarchy
- title > status/priority > due date > assignee > tags
- destructive actions always secondary and confirmed
- empty states should suggest next action

## 5. Design Tokens Guidance
- status colors must be consistent across web/mobile/desktop
- priority colors must remain distinguishable in dark mode
- spacing scale and typography should match SDD baseline

## 6. Components Requiring Final Standardization
- task card
- task detail header
- status badge
- priority badge
- filter bar
- toast style
- delete confirmation dialog
- loading skeletons

## 7. Responsive Rules
- mobile: stacked filters and single-column content
- tablet: compact sidebar/drawer behavior
- desktop: persistent navigation and wider task content area

## 8. UX Risks
- too many visual accents reduce signal clarity
- dense filter controls may overwhelm small screens
- dark mode contrast failures can hurt readability
