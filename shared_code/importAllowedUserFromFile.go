package shared_code

import (
	"encoding/json"
	"log"
)

func ParseAllowedUsersFromEmbeddedFile(allowedUsers []byte) (err error) {

	err = json.Unmarshal(allowedUsers, &AllowedUsersLoadFromJsonFile)

	if err != nil {
		log.Fatalln("parsing json '[]byte' containing allowed users", err.Error())

		return err
	}

	return err

}
