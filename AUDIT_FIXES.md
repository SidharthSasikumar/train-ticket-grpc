# Repository audit — 2026-09-20

## Changes

Validate nil requests and seat names; initialize maps; clone protobuf input/output to prevent caller mutation of stored receipts.

## Validation

Existing and new tests pass under go test -race ./....

## Limits

RPC service remains in-memory and unauthenticated. Exported maps still permit direct mutation by embedding callers.
