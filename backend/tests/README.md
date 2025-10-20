# Running Tests

This guide explains how to run integration tests for the backend using the test database defined in `.env.test`.

---

## Start the Test Database

Start the PostgreSQL container for tests:

`docker-compose -f docker-compose.test.yml up -d`

This starts a separate test database on the port defined in `.env.test`.

---

## Importing Environment Variables

Tests require the environment variables from `.env.test`. You can skip manual imports and let `test-main.go` load them.

---

## Run All Tests

From the `backend/` folder:

go test ./tests/... -v

- `-v` enables verbose output.
- Tests will automatically connect to the test database.

---

## Clean Up Between Runs

Tests truncate tables before each run, so they remain isolated. After finishing tests, stop and remove the test database container:

`docker-compose -f docker-compose.test.yml down -v`

This removes the test container and associated volumes.

---

## Notes

- The `RunSchema` function automatically loads `db/schema.sql` into the test database before tests start.  
- All tests under `tests/` will use this test database; do not point them to development database.