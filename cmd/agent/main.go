package main

import (
    "time"
    "devops-agent/internal/monitor"
    "devops-agent/internal/remediate"
)

func main() {
    for {
        state := monitor.Collect()
        remediate.Execute(state)
        time.Sleep(10 * time.Second)
    }
}
