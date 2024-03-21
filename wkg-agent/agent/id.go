package agent

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func init() {
	if WorkingDirectory == "" {
		WorkingDirectory = "/var/run"
	}

	uuid := uuid.New()

	// 获取当前时间戳（以纳秒为单位）
	timestamp := time.Now().UnixNano()

	// 创建带时间戳的 UUID
	uuidWithTimestamp := uuid
	copy(uuidWithTimestamp[6:], timestampBytes(timestamp))

	ID = uuidWithTimestamp.String()

	fmt.Println("agentid:" + ID)
}

var (
	Context, Cancel            = context.WithCancel(context.Background())
	ID                         = ""
	WorkingDirectory, _        = os.Getwd()
	Product             string = "wkg-agent"
	// from linker
	Version string = "1.0.0"
)

func timestampBytes(ts int64) []byte {
	b := make([]byte, 8)
	for i := uint(0); i < 64; i += 8 {
		b[7-(i/8)] = byte(ts >> i)
	}
	return b
}

func fromUUIDFile(file string) (id uuid.UUID, err error) {
	var idBytes []byte
	idBytes, err = os.ReadFile(file)
	if err == nil {
		id, err = uuid.ParseBytes(bytes.TrimSpace(idBytes))
	}
	return
}

func fromIDFile(file string) (id []byte, err error) {
	id, err = os.ReadFile(file)
	if err == nil {
		if len(id) < 6 {
			err = errors.New("id too short")
			return
		}
		id = bytes.TrimSpace(id)
	}
	return
}
