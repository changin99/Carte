package daemon

import (
    "log"
    "os"
    "os/exec"
)

// 컨테이너 초기 세팅
func InitContainer(command string) {
    cmd := exec.Command("sh", "-c", command)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        log.Fatal(err)
    }
}
