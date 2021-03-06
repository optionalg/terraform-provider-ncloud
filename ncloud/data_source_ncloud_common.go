package ncloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/terraform/helper/hashcode"
)

// Generates a hash for the set hash function used by the ID
func dataResourceIdHash(ids []string) string {
	var buf bytes.Buffer

	for _, id := range ids {
		buf.WriteString(fmt.Sprintf("%s-", id))
	}

	return fmt.Sprintf("%d", hashcode.String(buf.String()))
}

func writeToFile(filePath string, data interface{}) {
	log.Printf("[INFO] WriteToFile FilaPath: %s", filePath)

	os.Remove(filePath)

	if bs, err := json.MarshalIndent(data, "", "\t"); err == nil {
		str := string(bs)
		_ = ioutil.WriteFile(filePath, []byte(str), 777)
	}
}
