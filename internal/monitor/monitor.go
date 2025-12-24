package monitor

import "os/exec"

func Collect() map[string]string {
    state := make(map[string]string)
    err := exec.Command("systemctl", "is-active", "docker").Run()
    if err != nil {
        state["docker"] = "down"
    } else {
        state["docker"] = "up"
    }
    return state
}
