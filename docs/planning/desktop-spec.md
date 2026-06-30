# Desktop App Specification (Roadmap-Near)

## 1. Goal
Prepare Electron-based desktop approach with maximum web reuse.

## 2. Strategy
- Reuse Next.js UI where practical
- Electron wrapper for desktop packaging
- Native desktop conveniences added incrementally

## 3. Desktop Scope (future near-term)
- authenticated shell
- task list/detail/edit
- desktop notifications for reminders
- tray icon
- quick-add task modal
- keyboard shortcuts

## 4. Architecture
- Electron main process
- preload bridge
- renderer using existing web frontend
- secure IPC boundaries

## 5. Native Concerns
- auto-update strategy
- notification permissions
- secure local storage for session-related data
- OS packaging for macOS/Windows/Linux

## 6. Why Desktop Exists
- faster access for power users
- better shortcut-driven workflow
- notification and tray ergonomics
- opens path for offline-lite later

## 7. Readiness Gates
- web flows stable
- auth session strategy desktop-safe
- packaging pipeline defined
