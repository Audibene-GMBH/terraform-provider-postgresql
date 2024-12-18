package postgresql

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func getCommandOutput(ctx context.Context, command string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, command, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run %s %v stdout %s stderr %s %w", command, args, stdout.String(), stderr.String(), err)
	}
	return stdout.String(), nil
}
