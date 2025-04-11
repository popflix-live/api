package external

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

func StartConsumetAPI() {
	if !isPortInUse(3000) {
		apiPath := "/Users/vedantmhapsekar/Documents/iself/popflix/api.consumet.org"

		if _, err := os.Stat(apiPath); !os.IsNotExist(err) {
			go startExternalAPI(apiPath)
			time.Sleep(3 * time.Second)
		} else {
			log.Println("Warning: api.consumet.org directory not found, skipping external API start")
		}
	} else {
		log.Println("External API already running on port 3000")
	}
}

func isPortInUse(port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", port), 500*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func startExternalAPI(apiPath string) {
	cmd := exec.Command("npm", "start")
	cmd.Dir = apiPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Starting external API in directory: %s\n", apiPath)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error starting external API: %v\n", err)
	}
}
