package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

func GetTime() string {
	// now := time.Now()
	// year := now.Year()
	// month := now.Month()
	// day := now.Day()
	// hour := now.Hour()
	// minute := now.Minute()
	// second := now.Second()

	// curTimeDate, _ := fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	// fmt.Println(curTimeDate, "curTimeDate-curTimeDate")
	curTimeDate := time.Now().Format(" 2006-01-02 15:04:05")
	return curTimeDate
}

func HashPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}
