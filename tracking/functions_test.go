package tracking

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/denifrahman/shipper-go"
	"github.com/joho/godotenv"
)

func init() {
	var err = godotenv.Load("../.env")

	if err != nil {
		panic("Error loading .env file.")
	}

	productionMode, errParse := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))

	if errParse != nil {
		productionMode = false
	}

	shipper.Conf.SetAPIKey(os.Getenv("API_KEY")).SetProductionMode(productionMode)
}

func TestGetAllStatus(t *testing.T) {
	allStatus, err := GetAllStatus()

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(allStatus, "", "\t")

	fmt.Println("All Tracking Status:")
	fmt.Println(string(s))
}
