# HoHoHo Blackmarket CTF Writeup

Goal: Find the hidden flag field stored in the search index.

## Approach
1. Verify you get results with simple queries:
   - `*:*` (match all documents)
   - `category:tools` or `name:Lockpick*`

2. Enumerate fields using query-string semantics:
   - Try `_exists__:name`, `_exists__:price`, etc.
   - Then test hidden candidates: `_exists__:flag`
   - If it matches, the field exists on those docs.

3. Force results to match only documents that have the hidden field:
   - `flag:*`
   - Or combine with known constraints, e.g. `category:tools AND flag:*`

4. Extract the flag:
    - Once you found the `flag` property, you can bruteforce the flag like `flag:FLAG*`

## Why this works
- The backend forwards `q` straight to the search engine’s `query_string` parser.
- Hidden field is indexed and searchable but only removed at render time in the web app.
- Querying `_exists_` and `field:*` lets you discover and target hidden fields.

## Flag
- Present in multiple seeded docs under `flag`.
