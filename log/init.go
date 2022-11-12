// @Title  
// @Description  
// @Author  MrXiang 2022/11/8 14:23
package log

import (
	"log"
	"os"
)

func New(prefix string) Entry {
	entry := &BaseLogger{
		Logger: log.New(os.Stderr, "", 0),
	}
	entry.WithField("prefix", prefix)
	return entry
}
