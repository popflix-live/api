
# PopFlix API

This is the backend API for the PopFlix streaming service.

## Route Management

The API uses a centralized route configuration system to make it easy to add, modify, and maintain routes.

### How Routes Work

Routes are defined in the `src/lib/data` directory and registered automatically when the application starts. The system supports all standard HTTP methods (GET, POST, PUT, DELETE, PATCH).

### Adding a New Route

To add a new route to the API, follow these steps:

1. Create a handler function in the appropriate package
2. Add the route configuration to the routes list

#### Step 1: Create a Handler

Create a new handler in the appropriate directory structure. For example, if you're adding a user profile endpoint:

```go
// src/application/handlers/users/profile/get.go
package profile

import (
	"encoding/json"
	"net/http"
)

type UserProfile struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	profile := UserProfile{
		ID:        "user-123",
		Username:  "moviefan",
		Email:     "user@example.com",
		CreatedAt: "2023-01-15T00:00:00Z",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
```

#### Step 2: Add Route Configuration

Add your new route to the routes configuration in `src/lib/data/routes.go`:

```go
// Import your new handler
userProfile "github.com/popflix-live/api/src/application/handlers/users/profile"

// Add to the GetRoutes() function
{
    Method:      "GET",
    Path:        "/users/profile",
    Handler:     userProfile.GetHandler,
    Description: "Get user profile information",
},
```

### Route Configuration Options

Each route is defined with the following properties:

- **Method**: HTTP method (GET, POST, PUT, DELETE, PATCH)
- **Path**: URL path for the endpoint
- **Handler**: Function that handles the request
- **Description**: Human-readable description of what the endpoint does

### URL Parameters

To use URL parameters in your routes, include them in the path with a colon prefix:

```go
{
    Method:      "GET",
    Path:        "/movies/:id",
    Handler:     movieDetails.GetHandler,
    Description: "Get details for a specific movie",
},
```

In your handler, access the parameter using:

```go
id := chi.URLParam(r, "id")
```

### Query Parameters

To use query parameters (e.g., `/movies?genre=action`), access them in your handler:

```go
genre := r.URL.Query().Get("genre")
```

### Request Body

For POST/PUT requests with a JSON body, decode it in your handler:

```go
var req RequestStruct
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "Invalid request body", http.StatusBadRequest)
    return
}
```

### Multiple HTTP Methods for the Same Path

You can define multiple HTTP methods for the same path by adding separate route configurations:

```go
{
    Method:      "GET",
    Path:        "/auth/status",
    Handler:     authStatus.GetHandler,
    Description: "Get authentication status",
},
{
    Method:      "POST",
    Path:        "/auth/status",
    Handler:     authStatus.PostHandler,
    Description: "Update authentication status",
},
```

### Middleware

To add middleware to specific routes, you can extend the route configuration or apply middleware in the router setup.

## Running the API

Start the API server with:

```bash
go run src/main.go
```

The server will start on http://localhost:3001
