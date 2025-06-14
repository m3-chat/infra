package deploy

import (
	"fmt"
	"os"
	"os/exec"
)

func cloneAndInstall(repo, path string) error {
	fmt.Println("Cloning", repo)
	if err := exec.Command("git", "clone", repo, path).Run(); err != nil {
		return fmt.Errorf("failed to clone: %w", err)
	}

	fmt.Println("Installing dependencies")
	cmd := exec.Command("npm", "install")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Frontend() {
	fmt.Println("🚀 Deploying frontend...")
	if err := cloneAndInstall("https://github.com/m3-chat/frontend", "./_m3-frontend"); err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("ℹ️ Log into Vercel and deploy manually:")
	fmt.Println("cd _m3-frontend && vercel")
}

func Backend() {
	fmt.Println("🚀 Deploying backend...")
	if err := cloneAndInstall("https://github.com/m3-chat/backend", "./_m3-backend"); err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("ℹ️ Starting Cloudflare tunnel...")
	cmd := exec.Command("cloudflared", "tunnel", "--url", "http://localhost:2000", "--config", "./cloudflared.tunnel.config.yml")
	cmd.Dir = "./_m3-backend"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("❌ Error starting tunnel:", err)
	}
}
