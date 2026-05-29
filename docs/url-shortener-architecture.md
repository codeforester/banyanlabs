# URL Shortener Service — Architecture Document

## Overview

A multi-user URL shortening service built in Go as a learning project within Banyan Labs. The service supports user authentication, system-generated and custom short codes, and full CRUD operations on shortened URLs. It is designed to run locally as a single instance initially, with a clear path to scaling.

---

## Goals

- Learn Go service development end-to-end
- Implement a real three-tier architecture (frontend, business logic, storage)
- Practice industry-standard patterns: REST APIs, JWT authentication, hash-based ID generation
- Keep the design professional and production-quality, not a toy project

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go |
| Web server | Go standard library `net/http` |
| Authentication | JWT (JSON Web Tokens) |
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
- Validates inputs (URL format, short code format, JWT tokens)
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

### Table: shortened_urls

| Column | Type | Constraints | Notes |
|---|---|---|---|
| id | INTEGER | PRIMARY KEY, AUTOINCREMENT | Internal record ID |
| user_id | INTEGER | NOT NULL, FOREIGN KEY → users.id | Owner of this shortened URL |
| long_url | TEXT | NOT NULL | The original full URL |
| system_short_code | TEXT | NOT NULL, UNIQUE | Always generated, hash-based, 8 chars |
| custom_short_code | TEXT | NULLABLE, UNIQUE | Optional user-defined code, 3–20 chars |
| created_at | DATETIME | NOT NULL | Set on insert, never modified |
| modified_at | DATETIME | NOT NULL | Set on insert, updated on row change |

### Timestamp Behavior
- `created_at` is set automatically when a row is inserted and never changed again
- `modified_at` is set to the same value as `created_at` on insert, and updated whenever the row is modified

### Indexes
- Unique index on `users.username`
- Unique index on `users.email`
- Unique index on `shortened_urls.system_short_code`
- Unique index on `shortened_urls.custom_short_code` (nullable — SQLite treats NULLs as distinct, so multiple NULLs do not violate uniqueness)
- Index on `shortened_urls.long_url` for fast deduplication lookups
- Index on `shortened_urls.user_id` for fast user URL listing

---

## Short Code Design

### System-Generated Short Codes

- **Algorithm:** SHA-256 hash of the long URL, first N bytes encoded in Base62
- **Length:** 8 characters
- **Character set:** Base62 — `[a-zA-Z0-9]` (62 possible characters per position)
- **Collision space:** 62^8 ≈ 218 trillion possible codes
- **Collision handling:** SQLite unique constraint on `system_short_code`. If insert fails due to collision, take the next N bytes of the hash and retry. Successive collisions are extremely rare.
- **Deterministic:** Same long URL always produces the same system short code

### Custom Short Codes

- **Length:** 3 to 20 characters
- **Character set:** Alphanumeric plus dash and underscore — `[a-zA-Z0-9_-]`
- **Uniqueness:** Globally unique across all users
- **Optional:** User may provide a custom code at creation time or add/change it later via the update endpoint
- **Stored separately:** Custom codes are stored in their own column alongside the system-generated code

### Why Both Codes Are Stored

When user A shortens URL X with a custom code "mylink", and user B later shortens the same URL X without specifying a custom code, user B should receive the clean system-generated code — not user A's personal custom code. Storing both ensures:
- System-generated codes are always available as a neutral fallback
- Custom codes remain associated with the user who created them
- Deduplication returns the appropriate code based on context

---

## API Endpoints

### Authentication Endpoints

#### POST /auth/signup
Create a new user account.

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
  "username": "johndoe",
  "token": "<JWT token>"
}
```

---

#### POST /auth/login
Authenticate an existing user and receive a JWT token.

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
  "username": "johndoe",
  "token": "<JWT token>"
}
```

---

### URL Management Endpoints

All URL management endpoints require authentication via the `Authorization: Bearer <token>` header.

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
2. If `custom_short_code` is provided, validate format (3–20 chars, `[a-zA-Z0-9_-]`) and check global uniqueness
3. Query the database to check if `long_url` already exists
4. If it exists, return the existing system short code (not the custom code of another user)
5. If it does not exist, generate a system short code using SHA-256 + Base62
6. Attempt to insert the new row; handle unique constraint violations with retry
7. Return both codes

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
1. Validate JWT and extract user ID
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
1. Validate JWT and extract user ID
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
1. Look up `:shortCode` in both `system_short_code` and `custom_short_code` columns
2. If found, return HTTP 301/302 redirect to the original long URL
3. If not found, return 404

**Response (302 Found):**
```
Location: https://www.example.com/some/very/long/path
```

---

## Authentication Design

### JWT Token Flow

1. User signs up or logs in via `/auth/signup` or `/auth/login`
2. Server validates credentials, generates a signed JWT token containing:
   - `user_id`
   - `username`
   - `issued_at` timestamp
   - `expires_at` timestamp (e.g. 24 hours)
3. Token is signed with a server-side secret key
4. Client includes the token in all authenticated requests: `Authorization: Bearer <token>`
5. Server validates the token signature and expiration on every authenticated request
6. If valid, extracts `user_id` and proceeds. If invalid or expired, returns 401 Unauthorized

### Limitations (Current Phase)
- Token revocation is not supported in this initial implementation — tokens remain valid until expiration
- Logout functionality will require a Redis-based token blacklist, deferred to a future phase

---

## Business Logic Flows

### Create Shortened URL Flow

```
1. Receive POST /shorten with long_url and optional custom_short_code
2. Validate JWT token → extract user_id
3. Validate long_url format
4. If custom_short_code provided:
   a. Validate format (3–20 chars, [a-zA-Z0-9_-])
   b. Check global uniqueness in database
   c. Return error if taken
5. Query database: SELECT * FROM shortened_urls WHERE long_url = ?
6. If found → return existing system_short_code (deduplication)
7. If not found:
   a. Compute SHA-256 hash of long_url
   b. Encode first 6 bytes of hash to Base62 → 8-char system_short_code
   c. Attempt INSERT into shortened_urls
   d. If unique constraint violation on system_short_code → take next 6 bytes, retry
   e. On success → return new record
```

### Redirect Flow

```
1. Receive GET /:shortCode
2. Query: SELECT long_url FROM shortened_urls
         WHERE system_short_code = ? OR custom_short_code = ?
3. If found → HTTP 302 redirect to long_url
4. If not found → HTTP 404
```

---

## Error Handling

| Scenario | HTTP Status | Response |
|---|---|---|
| Invalid JWT token | 401 | Unauthorized |
| Expired JWT token | 401 | Token expired |
| User does not own resource | 403 | Forbidden |
| Short code not found | 404 | Not found |
| Custom code already taken | 409 | Conflict |
| Invalid URL format | 422 | Unprocessable entity |
| Invalid custom code format | 422 | Unprocessable entity |
| Custom code too short or too long | 422 | Unprocessable entity |

---

## Future Considerations

- **Logout and token revocation:** Add Redis as a token blacklist store
- **Click analytics:** Track redirect counts per short code
- **URL expiration:** Add optional TTL on shortened URLs
- **Rate limiting:** Prevent abuse of the shorten endpoint
- **PostgreSQL migration:** Replace SQLite when scaling beyond single instance
- **Multi-instance deployment:** Requires shared database (PostgreSQL) and stateless JWT already supports this
- **HTTPS:** Required before any production deployment

---

*This document covers the design of the URL shortener service as the first learning service in Banyan Labs. Implementation language is Go using the standard library `net/http`.*
