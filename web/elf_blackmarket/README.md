# HoHoHo Blackmarket CTF

A simple vulnerable web app themed as Santa's elves' blackmarket for naughty hackers.

- Frontend: htmx
- Backend: FastAPI
- Search: OpenSearch

Intended vulnerability: naive Lucene query injection to reveal hidden fields.

## Running locally

Prereqs: Docker, Python 3.10+

1. docker compose up -d
2. python -m venv .venv && source .venv/bin/activate
3. pip install -r requirements.txt
4. uvicorn app.main:app --reload

Open http://localhost:8000

## Notes for CTF operators
- The hidden flag field is indexed but not returned in normal responses.
- The search endpoint passes the user-supplied query straight to OpenSearch `q` parameter.
- Players should use Lucene query tricks to enumerate fields and discover hidden ones.
- Do NOT expose the flag elsewhere.

## Disclaimers
This app is intentionally insecure. Do not deploy to the public internet.
