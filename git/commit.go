package git

import (
	"fmt"
	"os"
	"os/exec"
)

func MakePreCommitFile() error {
	file, err := os.OpenFile(".git/hooks/pre-commit", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	preCommitContent(file)
	err = exec.Command("chmod", "755", ".git/hooks/pre-commit").Run()
	if err != nil {
		return err
	}

	return nil
}

func preCommitContent(file *os.File) {
	fmt.Fprintln(file, "#!/bin/sh")
	fmt.Fprintln(file, "commit_cnt=$(git log --after=\"`date '+%Y-%m-%d'` 0:0\" --oneline --branches * | wc -l)")
	fmt.Fprintln(file, "echo \"Today's commit ${commit_cnt}")
	fmt.Fprintln(file, "if [ $commit_cnt -ge 0 ]; then")
	fmt.Fprintln(file, "negirai appreciate")
	fmt.Fprintln(file, "fi")
}
