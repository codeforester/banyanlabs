# URL Shortener Service — Architecture Document

## Overview

A multi-user URL shortening service built in Go as a learning project within Banyan Labs. The service supports user authentication, system-generated and custom short codes, and full CRUD operations on shortened URLs. It is designed to run locally as a single instance initially, with a clear path to scaling.

---

## Goals

- Learn Go service development end-to-end
- Implement a real three-tier architecture (frontend, business logic, storage)
- Practice industry-standard patterns: server-rendered HTML, cookie-backed sessions, hash-based ID generation
- Keep the design professional and production-quality, not a toy project

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go |
| Web server | Go standard library `net/http` |
| Authentication | Cookie-backed server sessions for the HTML UI; JWT deferred for API clients |
| Storage | SQLite |
| Short code generation | SHA-256 hash + Base62 encoding |

---

## Three-Tier Architecture

### Tier 1 — Frontend
- Simple web interface for user signup, login, and URL management
- HTML forms for submitting long URLs and optional custom short codes
- Makes HTTP requests to the middle tier
- Displays shortened URLs and user's URL history

### Tier 2 — Business Logic (Middle Tier)
- Handles all API endpoints
- Validates inputs (URL format, short code format, sessions)
- Generates system short codes using SHA-256 + Base62 encoding
- Enforces deduplication logic
- Manages user authentication and authorization
- Talks to the storage layer for reads and writes

### Tier 3 — Storage (Backend)
- SQLite database running locally
- Stores user accounts and URL mappings
- Enforces uniqueness constraints at the database level
- Handles all persistence concerns

---

## Implementation Layers

The service should keep request handling, application behavior, and storage
separate.

```text
net/http handler
  -> parse request, session, and form data
  -> call application service/use case
  -> application layer validates and coordinates domain logic
  -> storage layer reads and writes SQLite
  -> application returns typed result or error
  -> handler maps result/error to HTML, redirect, or HTTP status
```

Layering rules:

- HTTP handlers must not contain SQL.
- Storage code must not render HTML or know about cookies.
- Application services own use cases such as signup, login, create URL,
  resolve code, and list user URLs.
- Domain helpers own URL normalization, short-code validation, and code
  generation.
- SQLite migrations define schema changes and run during local startup.
- Structured request logging belongs in HTTP middleware, with additional
  app/storage logs where useful.

---

## Database Schema

### Table: users

| Column | Type | Constraints | Notes |
|---|---|---|---|
| id | INTEGER | PRIMARY KEY, AUTOINCREMENT | Internal user ID |
| username | TEXT | NOT NULL, UNIQUE | Login username |
| email | TEXT | NOT NULL, UNIQUE | User email address |
| password_hash | TEXT | NOT NULL | Bcrypt hashed password |
| created_at | DATETIME | NOT NULL | Set on insert, never modified |
| modified_at | DATETIME | NOT NULL | Set on insert, updated on row change |

### Table: sessions

| Column | Type | Constraints | Notes |
|---|---|---|---|
| id | INTEGER | PRIMARY KEY, AUTOINCREMENT | Internal session ID |
| user_id | INTEGER | NOT NULL, FOREIGN KEY → users.id | Session owner |
| token_hash | TEXT | NOT NULL, UNIQUE | Hash of opaque session token |
| created_at | DATETIME | NOT NULL | Set on insert |
| expires_at | DATETIME | NOT NULL | Session expiration |
| last_seen_at | DATETIME | NULLABLE | Updated as the session is used |

### Table: shortened_urls

| Column | Type | Constraints | Notes |
|---|---|---|---|
| id | INTEGER | PRIMARY KEY, AUTOINCREMENT | Internal record ID |
| user_id | INTEGER | NOT NULL, FOREIGN KEY → users.id | Owner of this shortened URL |
| original_url | TEXT | NOT NULL | The original full URL used for display and redirect |
| normalized_url | TEXT | NOT NULL | Canonicalized URL used for per-user deduplication |
| created_at | DATETIME | NOT NULL | Set on insert, never modified |
| modified_at | DATETIME | NOT NULL | Set on insert, updated on row change |

### Table: short_codes

| Column | Type | Constraints | Notes |
|---|---|---|---|
| code | TEXT | PRIMARY KEY | Globally unique public short code |
| shortened_url_id | INTEGER | NOT NULL, FOREIGN KEY → shortened_urls.id | Target URL record |
| kind | TEXT | NOT NULL, CHECK (`system` or `custom`) | Code type |
| created_at | DATETIME | NOT NULL | Set on insert |

### Timestamp Behavior
- `created_at` is set automatically when a row is inserted and never changed again
- `modified_at` is set to the same value as `created_at` on insert, and updated whenever the row is modified

### Indexes
- Unique index on `users.username`
- Unique index on `users.email`
- Unique index on `sessions.token_hash`
- Unique constraint on `shortened_urls(user_id, normalized_url)` for per-user deduplication
- Index on `shortened_urls.user_id` for fast user URL listing
- Primary key on `short_codes.code` for global public-code uniqueness
- Index on `short_codes.shortened_url_id` for loading codes for a URL record

---

## Short Code Design

### System-Generated Short Codes

- **Algorithm:** SHA-256 hash of the long URL, first N bytes encoded in Base62
- **Length:** 8 characters
- **Character set:** Base62 — `[a-zA-Z0-9]` (62 possible characters per position)
- **Collision space:** 62^8 ≈ 218 trillion possible codes
- **Collision handling:** SQLite unique constraint on `short_codes.code`. If the first candidate collides, take the next N bytes of the hash and retry. Successive collisions are extremely rare.
- **Deterministic candidate:** Same normalized URL starts from the same first candidate, but collision handling can choose a later candidate when the global code namespace is already occupied.

### Custom Short Codes

- **Length:** 3 to 20 characters
- **Character set:** Alphanumeric plus dash and underscore — `[a-zA-Z0-9_-]`
- **Uniqueness:** Globally unique across all system and custom public codes
- **Optional:** User may provide a custom code at creation time or add/change it later via the update endpoint
- **Stored separately:** Custom codes are stored as `short_codes` rows with `kind = custom`

### Why Both Codes Are Stored

When user A shortens URL X with a custom code "mylink", and user B later shortens the same URL X without specifying a custom code, user B should receive the clean system-generated code — not user A's personal custom code. Storing both ensures:
- System-generated codes are always available as a neutral fallback
- Custom codes remain associated with the user who created them
- Deduplication returns the appropriate code based on context
- Public redirects can resolve a globally unique code without user context

---

## HTTP Endpoints

### Authentication Endpoints

#### POST /auth/signup
Create a new user account and start a browser session.

**Request body:**
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response (201 Created):**
```json
{
  "user_id": 1,
  "username": "johndoe"
}
```

The response sets an HTTP-only session cookie.

---

#### POST /auth/login
Authenticate an existing user and start a browser session.

**Request body:**
```json
{
  "username": "johndoe",
  "password": "securepassword"
}
```

**Response (200 OK):**
```json
{
  "user_id": 1,
  "username": "johndoe"
}
```

The response sets an HTTP-only session cookie.

---

#### POST /auth/logout
Delete the current session and clear the browser session cookie.

**Response (204 No Content)**

---

### URL Management Endpoints

All browser URL management endpoints require a valid session cookie. JWT-based
API authentication is deferred to a later API-client phase.

#### POST /shorten
Create a new shortened URL.

**Request body:**
```json
{
  "long_url": "https://www.example.com/some/very/long/path",
  "custom_short_code": "mylink"  // optional
}
```

**Business logic:**
1. Validate that `long_url` is a valid URL
2. Normalize `long_url` into `normalized_url`
3. If `custom_short_code` is provided, validate format (3–20 chars, `[a-zA-Z0-9_-]`) and check global public-code uniqueness
4. Query the database for the authenticated user's existing `normalized_url`
5. If this user already has the URL, return that user's existing codes
6. If it does not exist, create a `shortened_urls` row and generate a system short code using SHA-256 + Base62
7. Attempt to insert `short_codes` rows; handle global code uniqueness conflicts with retry for the system code
8. Return both codes

**Response (201 Created):**
```json
{
  "long_url": "https://www.example.com/some/very/long/path",
  "system_short_code": "aB3xYz9q",
  "custom_short_code": "mylink",
  "created_at": "2025-04-11T10:00:00Z"
}
```

---

#### PUT /shorten/:shortCode
Update the custom short code of an existing shortened URL. Only the owner of the URL can update it.

**Request body:**
```json
{
  "custom_short_code": "newcustomcode"
}
```

**Business logic:**
1. Validate session and extract user ID
2. Look up the shortened URL by `:shortCode` (system or custom)
3. Verify the requesting user owns this record
4. Validate the new custom short code format
5. Check global uniqueness of the new custom short code
6. Update the record and set `modified_at`

**Response (200 OK):**
```json
{
  "long_url": "https://www.example.com/some/very/long/path",
  "system_short_code": "aB3xYz9q",
  "custom_short_code": "newcustomcode",
  "modified_at": "2025-04-11T11:00:00Z"
}
```

---

#### DELETE /shorten/:shortCode
Delete a shortened URL. Only the owner can delete it.

**Business logic:**
1. Validate session and extract user ID
2. Look up the record by `:shortCode`
3. Verify ownership
4. Delete the record

**Response (204 No Content)**

---

#### GET /user/urls
Retrieve all shortened URLs belonging to the authenticated user.

**Response (200 OK):**
```json
[
  {
    "long_url": "https://www.example.com/path",
    "system_short_code": "aB3xYz9q",
    "custom_short_code": "mylink",
    "created_at": "2025-04-11T10:00:00Z",
    "modified_at": "2025-04-11T11:00:00Z"
  }
]
```

---

### Public Redirect Endpoint

#### GET /:shortCode
Redirect a short code to its original long URL. No authentication required.

**Business logic:**
1. Look up `:shortCode` in the global `short_codes` table
2. If found, return HTTP 301/302 redirect to the original long URL
3. If not found, return 404

**Response (302 Found):**
```
Location: https://www.example.com/some/very/long/path
```

---

## Authentication Design

### Cookie Session Flow

1. User signs up or logs in through the HTML UI
2. Server validates credentials and generates an opaque random session token
3. Server stores only a hash of the session token in SQLite
4. Server returns the raw session token in an HTTP-only cookie
5. Browser includes the cookie on authenticated requests
6. Server hashes the cookie token, finds the session row, verifies expiration,
   and extracts `user_id`
7. Logout deletes the session row and clears the cookie

### Limitations (Current Phase)
- Sessions are stored in SQLite, so the first implementation is single-instance
- JWT authentication for non-browser API clients is deferred to a future phase
- Distributed session storage can be revisited when Banyan Labs adds multi-instance deployment

---

## Business Logic Flows

### Create Shortened URL Flow

```
1. Receive POST /shorten with long_url and optional custom_short_code
2. Validate session cookie → extract user_id
3. Validate long_url format
4. Normalize long_url
5. If custom_short_code provided:
   a. Validate format (3–20 chars, [a-zA-Z0-9_-])
   b. Check global public-code uniqueness in database
   c. Return error if taken
6. Query database: SELECT * FROM shortened_urls WHERE user_id = ? AND normalized_url = ?
7. If found → return this user's existing codes
8. If not found:
   a. Compute SHA-256 hash of long_url
   b. Encode first 6 bytes of hash to Base62 → 8-char system_short_code
   c. Attempt INSERT into shortened_urls and short_codes
   d. If unique constraint violation on short_codes.code → take next 6 bytes, retry
   e. On success → return new record
```

### Redirect Flow

```
1. Receive GET /:shortCode
2. Query short_codes joined to shortened_urls by code
3. If found → HTTP 302 redirect to long_url
4. If not found → HTTP 404
```

---

## Error Handling

| Scenario | HTTP Status | Response |
|---|---|---|
| Missing or invalid session | 401 | Unauthorized |
| Expired session | 401 | Session expired |
| User does not own resource | 403 | Forbidden |
| Short code not found | 404 | Not found |
| Custom code already taken | 409 | Conflict |
| Invalid URL format | 422 | Unprocessable entity |
| Invalid custom code format | 422 | Unprocessable entity |
| Custom code too short or too long | 422 | Unprocessable entity |

---

## Future Considerations

- **JWT authentication:** Add JWT support for non-browser API clients
- **Distributed sessions:** Add Redis or another shared session store when multi-instance deployment requires it
- **Click analytics:** Track redirect counts per short code
- **URL expiration:** Add optional TTL on shortened URLs
- **Rate limiting:** Prevent abuse of the shorten endpoint
- **PostgreSQL migration:** Replace SQLite when scaling beyond single instance
- **Multi-instance deployment:** Requires shared database (PostgreSQL) and shared session storage or stateless API credentials
- **HTTPS:** Required before any production deployment

---

*This document covers the design of the URL shortener service as the first learning service in Banyan Labs. Implementation language is Go using the standard library `net/http`.*
