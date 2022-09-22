package Basic

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func KillContainer(filename string) error {
	// Data comes in as byte format, need to convert into string type
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// If we can read the file, then after reading, at last delete the file
	defer os.Remove(filename)
	cid := strings.TrimSpace(string(data))
	if !isValid(cid) {
		return fmt.Errorf("bad container id: %q", cid)
	}
	log.Printf("stopping container %s", cid)
	// if err := exec.Command("docker", "rm", "-f", cid).Run(); err != nil {
	// 	return fmt.Errorf("failed to stop %s: %w", cid, err)
	// }
	return nil
}

func isValid(cid string) bool {
	return len(cid) == 12 || len(cid) == 64
}
