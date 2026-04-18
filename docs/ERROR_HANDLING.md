# Error Handling Guide

This document explains the error handling strategy used across the HackKU2026 backend and how developers should handle errors at each layer.

## Architecture Overview

The backend is organized into three layers, each with distinct error handling responsibilities:

```
┌─────────────────────────────────────────┐
│          Routes / HTTP Layer            │  ← Translate errors to JSON + status codes
├─────────────────────────────────────────┤
│      Service / Business Logic Layer     │  ← Apply security policies, transform errors
├─────────────────────────────────────────┤
│         Database / Data Layer           │  ← Return our custom error types
└─────────────────────────────────────────┘
```

## Error Types

### ClientError

A `ClientError` is any error safe to show to the client. It contains:
- **Type** (error tag): Machine-readable error identifier (e.g., `"invalid_credentials"`)
- **Message**: Human-readable error description shown to clients
- **Code**: HTTP status code for the response

All `ClientError` instances are pre-defined in the `errors` package (e.g., `errors.ErrUserNotFound`).

**Example:**
```go
var ErrInvalidUsernameOrPassword = New("invalid_credentials", "invalid username or password", http.StatusUnauthorized)
```

### Internal Errors

Any error from external libraries (MongoDB, bcrypt, etc.) that is NOT a `ClientError` is treated as internal. These errors are **never** exposed to clients—clients always receive `errors.ErrUnknown`.

## Layer Responsibilities

### 1. Database Layer (`database/`)

**Purpose:** Access data, translate Mongo errors to domain errors.

**Rules:**
- Convert database-specific errors to `ClientError` types
  - `mongo.ErrNoDocuments` → `ErrUserNotFound`
  - Duplicate key errors → `ErrUsernameAlreadyExists`
- **Never use `errors.Join()`** (see "Don't use errors.Join()" section)
- Return only `ClientError` or internal errors — don't mix them on purpose
- Internal errors (network failures, etc.) bubble up unchanged for logging

**Example:**
```go
func GetUser(id string) (*models.User, error) {
    var model models.User
    if err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&model); err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, errors.ErrUserNotFound  // Return ONLY the ClientError
        }
        return nil, err  // Internal error, will be logged and turned into 500
    }
    return &model, nil
}
```

### 2. Service Layer (`service/`)

**Purpose:** Business logic, enforce policies, transform errors for security.

**Rules:**
- Check errors from database layer
- If it's a `ClientError`, decide whether to:
  - Return it as-is (e.g., `ErrUserNotFound` when safe to reveal)
  - Transform it to a different `ClientError` (e.g., return `ErrInvalidUsernameOrPassword` instead of `ErrUserNotFound` to prevent enumeration)
  - Return a generic error (e.g., `ErrUnknown` if something truly unexpected happened)
- **Never expose internal details** to clients
- **Never use `errors.Join()` for hiding errors** (see below)

**Example (Security: Prevent Username Enumeration):**
```go
func Signin(request *models.SigninRequest) (*models.Session, error) {
    user, err := database.GetUserByUsername(request.Username)
    if err != nil {
        // Return the same generic error regardless of why lookup failed
        // (user not found, database error, etc.)
        // This prevents attackers from determining which usernames exist
        return nil, errors.ErrInvalidUsernameOrPassword
    }

    if err := bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {
        // Also return the same generic error for wrong password
        return nil, errors.ErrInvalidUsernameOrPassword
    }

    // ... create session and return
}
```

**Example (Safe Error Propagation):**
```go
func Signup(request *models.SignupRequest) (*models.User, error) {
    user := models.NewUser(request.Username, request.Email, hash)
    
    if err := database.CreateUser(user); err != nil {
        // If it's ErrUsernameAlreadyExists, the client should know
        // So return it as-is
        return nil, err
    }
    
    return user, nil
}
```

### 3. Routes / HTTP Layer (`routes/`)

**Purpose:** Serialize errors to JSON, set HTTP status codes.

**Rules:**
- Call business logic and error handlers
- Use centralized error response function `writeError()`
- **Never manually set status codes or write error text**
- `writeError()` handles all JSON serialization

**Example:**
```go
func signup(w http.ResponseWriter, r *http.Request) {
    var request models.SignupRequest
    if err := readBody(r, &request); err != nil {
        // readBody returns ErrInvalidBody if JSON is malformed
        writeError(w, err)
        return
    }

    user, err := service.Signup(&request)
    if err != nil {
        writeError(w, err)  // Centralized JSON response
        return
    }

    writeJSON(w, http.StatusOK, user)
}
```

### 4. Middleware (`middleware.go`)

**Purpose:** Enforce authn/authz, protect routes.

**Rules:**
- Check session validity
- Return `ClientError` types (e.g., `ErrInvalidSessionID`, `ErrSessionExpired`)
- Use `writeErrorJSON()` helper for JSON responses
- **Never use hardcoded `http.Error()`**

**Example:**
```go
func RequireSession(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("session_id")
        if err != nil {
            writeErrorJSON(w, errors.ErrInvalidSessionID)
            return
        }
        // ... validate session, attach to context
    })
}
```

## Don't Use `errors.Join()` for Hiding Errors

### ❌ Wrong
```go
// This exposes the internal mongo error or bcrypt error
return nil, errors.Join(err, errors.ErrInvalidUsernameOrPassword)
```

**Why?** `errors.Join()` combines errors. A caller can later do:
```go
if errors.Is(combinedErr, mongo.ErrNoDocuments) {
    // This will be TRUE, revealing that the user wasn't found!
}
```

This defeats security measures like preventing username enumeration.

### ✅ Correct
```go
// This returns ONLY the public error, hiding the true reason
return nil, errors.ErrInvalidUsernameOrPassword
```

**When is `errors.Join()` OK?** Only when you genuinely want to combine multiple independent errors for the client (rare). Example: validation errors on multiple fields.

## Adding New Errors

1. **Define in the appropriate errors file** (`errors/account.go`, `errors/user.go`, `errors/session.go`):
   ```go
   var ErrNewError = New("error_type", "human readable message", http.StatusXXX)
   ```

2. **Use in database layer** to convert library errors:
   ```go
   if someCondition {
       return nil, errors.ErrNewError
   }
   ```

3. **Decide in service layer** whether to:
   - Return it as-is
   - Transform it to a different error
   - Hide it with a generic error

4. **Routes automatically handle it** via `writeError()`.

## Testing Error Flows

When testing, verify:
- ✅ Database layer returns correct `ClientError` type
- ✅ Service layer transforms sensitive errors appropriately
- ✅ Routes return correct HTTP status code and JSON shape
- ✅ Internal errors become `ErrUnknown` (500) to the client

Example:
```go
func TestSignin_InvalidUsername(t *testing.T) {
    session, err := service.Signin(&models.SigninRequest{
        Username: "nonexistent",
        Password: "anypassword",
    })
    
    // Should return ErrInvalidUsernameOrPassword, NOT ErrUserNotFound
    if !errors.Is(err, errors.ErrInvalidUsernameOrPassword) {
        t.Errorf("expected ErrInvalidUsernameOrPassword, got %v", err)
    }
}
```

## Summary

| Layer | Input | Logic | Output |
|-------|-------|-------|--------|
| **Database** | Mongo errors, raw data | Convert Mongo errors to `ClientError` | `ClientError` or bubble internal errors |
| **Service** | `ClientError` from DB | Apply business rules, transform for security | `ClientError` or internal errors |
| **Routes** | Error from service | Use `writeError()` | JSON response with correct status |
| **Middleware** | Session state | Check validity | `ErrSessionExpired` etc. or pass through |

**Golden Rule:** A `ClientError` is the contract between backend and client—only these errors are intentionally exposed. Everything else becomes `ErrUnknown`.
