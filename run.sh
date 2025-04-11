echo "Running the script..."
files() {
    if [ -z "$1" ]; then
        path="src/application/handlers"
    else
        path="$1"
    fi
    
    for file in "$path"/*; do
        if [ -d "$file" ]; then
            # Recursively process directories without printing them
            files "$file"
        else
            # Only print actual files
            echo "$file"
        fi
    done
}

generate_route_registrar() {
    OUTPUT_FILE="src/application/router/autoRegister.go"
    echo "Generating route registrar at $OUTPUT_FILE..."
    
    # Create the file header
    cat > $OUTPUT_FILE << EOL
// This file is auto-generated. Do not edit manually.
package router

import (
    "fmt"

    "github.com/gin-gonic/gin"
EOL
    
    # Find all handler packages and add imports
    HANDLERS=$(find src/application/handlers -type d | grep -v "^src/application/handlers$")
    for dir in $HANDLERS; do
        # Convert directory path to Go import path
        IMPORT_PATH=$(echo $dir | sed 's/^src\//github.com\/popflix-live\/api\/src\//')
        
        # Extract package name from directory path
        PKG_NAME=$(basename $dir)
        
        # Check if directory contains get.go file
        if [ -f "$dir/get.go" ]; then
            echo "	$PKG_NAME \"$IMPORT_PATH\"" >> $OUTPUT_FILE
        fi
    done
    
    # Close imports and start function
    cat >> $OUTPUT_FILE << EOL
    "github.com/popflix-live/api/src/lib/models/router"
)

// AutoRegisterRoutes automatically registers all handler functions from the handlers directory
func AutoRegisterRoutes(r *gin.Engine) {
    routes := []router.RouteConfig{
EOL
    
    # Add route configurations for each handler
    for dir in $HANDLERS; do
        if [ -f "$dir/get.go" ]; then
            # Extract package name
            PKG_NAME=$(basename $dir)
            
            # Create path from directory structure
            PATH_PARTS=$(echo $dir | sed 's/^src\/application\/handlers//' | tr '/' ' ')
            API_PATH="/"
            for part in $PATH_PARTS; do
                if [ -n "$part" ]; then
                    API_PATH="${API_PATH}${part}/"
                fi
            done
            API_PATH=$(echo $API_PATH | sed 's/\/$//')
            
            # Add route configuration
            cat >> $OUTPUT_FILE << EOL
        {
            Method:      "GET",
            Path:        "$API_PATH",
            Handler:     $PKG_NAME.GetHandler,
            Description: "Auto-generated route for $API_PATH",
        },
EOL
        fi
    done
    
    # Complete the function
    cat >> $OUTPUT_FILE << EOL
    }
    
    for _, route := range routes {
        registerRoute(r, route)
    }
    
    fmt.Println("Auto-registered", len(routes), "routes")
}
func registerRoute(r *gin.Engine, route router.RouteConfig) {
    switch route.Method {
    case "GET":
        r.GET(route.Path, route.Handler)
    case "POST":
        r.POST(route.Path, route.Handler)
    case "PUT":
        r.PUT(route.Path, route.Handler)
    case "DELETE":
        r.DELETE(route.Path, route.Handler)
    default:
        fmt.Println("Unsupported method:", route.Method)
    }
}
EOL
    
    echo "Route registrar generated successfully with $(grep -c "Handler:" $OUTPUT_FILE) routes"
}
run_application() {
    echo "Starting the PopFlix API server..."
    go run src/main.go
}
stop_running_application() {
    echo "Checking for running application on port 8000..."
    
    # Use netstat to find PID of process using port 8080 (not 8000)
    netstat -ano | findstr ":8000" > temp_port_check.txt
    
    if [ -s temp_port_check.txt ]; then
        # Extract PID from the last column
        PID=$(cat temp_port_check.txt | awk '{print $NF}')
        echo "Found process with PID: $PID"
        
        echo "Stopping running application with PID $PID..."
        # Use cmd to execute taskkill properly
        cmd //c "taskkill /PID $PID /F"
        
        if [ $? -eq 0 ]; then
            echo "Process terminated successfully."
        else
            echo "Failed to terminate the process. You may need to stop it manually."
        fi
    else
        echo "No application is running on port 8000."
    fi
    
    # Clean up
    rm -f temp_port_check.txt
}


# You can call the original files function to list handlers
# files "src/application/handlers"

# Or generate the route registrar
stop_running_application
generate_route_registrar
run_application