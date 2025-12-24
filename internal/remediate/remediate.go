package remediate

import "os/exec"

func Execute(state map[string]string) {
    if state["docker"] == "down" {
        exec.Command("systemctl", "restart", "docker").Run()
    }
}
