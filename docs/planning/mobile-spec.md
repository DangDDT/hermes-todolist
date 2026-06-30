# Mobile App Specification (Roadmap-Near)

## 1. Goal
Prepare mobile architecture and scope so Flutter implementation can start without re-discovery.

## 2. Platform Scope
- iOS
- Android
- Shared codebase via Flutter

## 3. MVP-on-Mobile Scope (future near-term)
- login/register
- my tasks list
- task detail
- create/edit task
- update status quickly
- offline-not-required in first mobile cut

## 4. Architecture
- Flutter
- Riverpod recommended
- Dio/http client
- go_router for navigation
- feature-first folders

## 5. Screen Inventory
- splash / auth gate
- login
- register
- task list
- task detail
- task create/edit
- settings

## 6. Mobile-specific Considerations
- smaller forms
- touch target sizing
- optimistic refresh patterns
- pull-to-refresh
- token/session handling on app resume
- optional push notifications later

## 7. Risks
- duplicate effort if API contracts unstable
- notification scope creep
- offline expectations if users assume native app parity

## 8. Exit Criteria Before Build Starts
- API contract stable
- auth flow finalized
- task payloads stable
- design tokens reusable from web
