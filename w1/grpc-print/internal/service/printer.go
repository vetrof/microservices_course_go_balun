package service

import (
	"fmt"
	"os"
)

type PrinterService struct {
	storagePath string
}

func NewPrinterService(path string) *PrinterService {
	return &PrinterService{storagePath: path}
}

func (ps *PrinterService) SaveMessage(message string) error {
	f, err := os.OpenFile(ps.storagePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	logLine := fmt.Sprintf("Saved: %s\n", message)
	_, err = f.WriteString(logLine)
	return err
}
