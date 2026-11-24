# Non-Overlapping Intervals

> NeetCode 150 · Intervals · Blind 75 — flip "minimum removals" into "maximum kept", which is activity selection: earliest end first, count the compatible. Go, stdlib only, tested.

## Problem

Remove as few intervals as possible so none of the rest overlap (sharing an endpoint doesn't count as overlap). `[1,9],[2,3],[4,5],[6,7]` removes just the long one.

## Approach

Minimizing removals is maximizing the kept set — and the maximum set of pairwise-compatible intervals is exactly activity selection (dsa-048): sort by end time, greedily keep whatever starts at or after the last kept end. The earliest-finisher exchange argument carries over unchanged; the answer is `n − kept`. The long-interval test shows why end-time (not start-time or length) is the right sort key: `[1,9]` finishes late, so the greedy never keeps it when shorter alternatives exist. One problem, zero new machinery — recognizing the reduction *is* the solution.

**Complexity:** O(n log n) time, O(n) copy.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
