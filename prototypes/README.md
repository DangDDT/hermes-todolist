# UX/UI Prototypes

## Artifact
- `ux-ui-prototype-v1.html`

## Purpose
Artifact này dùng cho **Phase 2 UX/UI review** trước khi polish production UI sâu hơn.

## Review Order
1. System design
2. Components
3. Multi-device pages
4. Flows
5. Scenarios
6. Motion / animated posture

## What to Review
- hierarchy có rõ chưa?
- status / priority / actions có đủ scanable chưa?
- layout desktop / tablet / mobile có cùng logic chưa?
- create/edit/detail/list có đồng bộ một design language chưa?
- auth flow có đủ calm và low-friction chưa?
- motion có subtle và hữu ích chưa?

## Run locally
From repo root:

```bash
python3 -m http.server 8765
```

Open:
- `http://127.0.0.1:8765/prototypes/ux-ui-prototype-v1.html`

## Notes
Prototype này là **review artifact**, không phải production frontend cuối cùng. Nó đóng vai trò bridge giữa planning/system design và production polish / implementation sync.