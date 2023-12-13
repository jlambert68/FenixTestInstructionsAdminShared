package shared_code

import (
	"encoding/json"
	"log"
	"os"
)

func ImportAllowedUsersFromFile() (err error) {

	var allowedUsersFile *os.File
	allowedUsersFile, err = os.Open(RelativePathToAllowedUsersList)

	if err != nil {
		log.Fatalln("opening json file containing allowed users", err.Error())
		return err
	}

	jsonParser := json.NewDecoder(allowedUsersFile)
	if err = jsonParser.Decode(&AllowedUsersLoadFronJsonFile); err != nil {
		log.Fatalln("parsing json file containing allowed users", err.Error())

		return err
	}

	return err

}
