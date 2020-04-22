package swagger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func readWhitelist(f string, targ *[]WhitelistEntry) error{
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	err = json.NewDecoder(bytes.NewReader(file)).Decode(targ)
	if err != nil {
		return err
	}

	return nil
}

func writeWhitelist(f string, payl *[]WhitelistEntry) error{
	file, _ := json.MarshalIndent(payl, "", " ")
	err := ioutil.WriteFile(f, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
