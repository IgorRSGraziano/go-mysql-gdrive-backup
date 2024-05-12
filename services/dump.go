package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Dump struct {
	GenerateCmd string
	FilePath    string
	FileName    *string
}

func NewDump(generateCmd, filePath string) *Dump {
	return &Dump{
		GenerateCmd: generateCmd,
		FilePath:    filePath,
	}
}

func (d *Dump) GetFullFilePath() string {
	return fmt.Sprintf("%s/%s", d.FilePath, *d.FileName)
}

func (d *Dump) Generate() error {
	preparedCmd, args := prepareCommand(d.GenerateCmd)
	out, err := exec.Command(preparedCmd, args...).CombinedOutput()
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s_backup.sql", time.Now().Format("2006010215"))
	d.FileName = &fileName
	err = os.WriteFile(d.GetFullFilePath(), out, 0644)
	return err
}

func prepareCommand(cmd string) (exec string, args []string) {
	splitCmd := strings.Fields(cmd)
	return splitCmd[0], splitCmd[1:]
}
