package write

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func randomInt() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()
}

func SaveData2(path string, data []byte) error {
    tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
    fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
    if err != nil {
        return err
    }
    defer func() { 
        fp.Close() // not expected to fail
        if err != nil {
            os.Remove(tmp)
        }
    }()

    if _, err = fp.Write(data); err != nil { // 1. save to the temporary file
        return err
    }
    if err = fp.Sync(); err != nil { // 2. fsync
        return err
    }
    err = os.Rename(tmp, path) // 3. replace the target

    if err = syncDir(filepath.Dir(path)); err != nil {
		return fmt.Errorf("syncDir: %w", err)
	}
    return nil
}

func syncDir(dir string) error {
    dfp,err := os.Open(dir)
    if err != nil {
        return err
    }
    defer dfp.Close()
    return dfp.Sync()
}