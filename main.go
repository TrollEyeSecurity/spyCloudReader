package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"spyCloudReader/spycloud"
	"time"
)

func main() {
	spyCloudResponse := flag.String("spyCloudResponse", "", "Enter the path to the spy cloud json file.")
	months := flag.Int("months", 60, "Enter the number of months to look back.")
	userPass := flag.Bool("userPass", false, "Use to create username/password output.")
	dump2Mongo := flag.Bool("dump2Mongo", false, "Dump the output to mongodb")
	flag.Parse()
	content, contentErr := ioutil.ReadFile(*spyCloudResponse)
	if contentErr != nil {
		log.Fatal(contentErr)
	}
	if *dump2Mongo {
		responseResults := &spycloud.ResponseResultsMongoDb{}
		JsonUnmarshalErr := json.Unmarshal(content, responseResults)
		if JsonUnmarshalErr != nil {
			err := fmt.Errorf("breach record json-unmarshal error %v: %v", JsonUnmarshalErr, string(content))
			log.Fatal(err)
		}
		dumpToMongo(responseResults)
		return
	}
	responseResults := &spycloud.ResponseResults{}
	JsonUnmarshalErr := json.Unmarshal(content, responseResults)
	if JsonUnmarshalErr != nil {
		err := fmt.Errorf("breach record json-unmarshal error %v: %v", JsonUnmarshalErr, string(content))
		log.Fatal(err)
	}

	for _, breachRecord := range responseResults.Results {
		t, err := time.Parse(time.RFC3339, breachRecord.SpycloudPublishDate)
		if err != nil {
			panic(err)
		}
		count := monthsCountSince(t)
		if *months >= *count {
			if *userPass {
				if len(breachRecord.Username) > 0 && breachRecord.Password != "" {
					fmt.Println(breachRecord.Username + "," + breachRecord.Password)
				}
				if len(breachRecord.Email) > 0 && breachRecord.Password != "" {
					fmt.Println(breachRecord.Email + "," + breachRecord.Password)
				}
			} else {
				fmt.Println()
				if len(breachRecord.FirstName) > 0 {
					fmt.Println("First Name: " + breachRecord.FirstName)
					fmt.Println("Last Name: " + breachRecord.LastName)
				}
				if len(breachRecord.Username) > 0 {
					fmt.Println("Username: " + breachRecord.Username)
				}
				if len(breachRecord.Email) > 0 {
					fmt.Println("Email: " + breachRecord.Email)
				}
				if breachRecord.Password != "" {
					fmt.Println("Password: " + breachRecord.Password)
				}
				if breachRecord.PasswordPlaintext != "" {
					fmt.Println("Password Plaintext: " + breachRecord.PasswordPlaintext)
				}
				if len(breachRecord.PasswordType) > 0 {
					fmt.Println("Password Type: " + breachRecord.PasswordType)
				}
				if len(breachRecord.AccountSecret) > 0 {
					fmt.Println("Account Secret: " + breachRecord.AccountSecret)
				}
				if len(breachRecord.TargetUrl) > 0 {
					fmt.Println("Target Url: " + breachRecord.TargetUrl)
				}
				if len(breachRecord.TargetDomain) > 0 {
					fmt.Println("Target Domain: " + breachRecord.TargetDomain)
				}
				if len(breachRecord.InfectedMachineId) > 0 {
					fmt.Println("Infected Machine Id: " + breachRecord.InfectedMachineId)
				}
				if len(breachRecord.SourceFile) > 0 {
					fmt.Println("Source File: " + breachRecord.SourceFile)
				}
				if len(breachRecord.InfectedTime) > 0 {
					it, itErr := time.Parse(time.RFC3339, breachRecord.InfectedTime)
					if itErr != nil {
						panic(itErr)
					}
					fmt.Println("Infected Time: " + it.Format(time.ANSIC))
				}
				if len(breachRecord.RecordModificationDate) > 0 {
					md, mdErr := time.Parse(time.RFC3339, breachRecord.RecordModificationDate)
					if mdErr != nil {
						panic(mdErr)
					}
					fmt.Println("Record Modification Date: " + md.Format(time.ANSIC))
				}
				if len(breachRecord.RecordAdditionalDate) > 0 {
					ra, raErr := time.Parse(time.RFC3339, breachRecord.RecordAdditionalDate)
					if raErr != nil {
						panic(raErr)
					}
					fmt.Println("Record Additional Date: " + ra.Format(time.ANSIC))
				}
				fmt.Println("Publish Date: " + t.Format(time.ANSIC))
				fmt.Println()
				fmt.Println("---------------------------------------------------------------")
			}
		}
	}
}

func monthsCountSince(createdAtTime time.Time) *int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextMonth := createdAtTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}
	return &months
}

func dumpToMongo(responseResults *spycloud.ResponseResultsMongoDb) {
	for _, breachRecord := range responseResults.Results {
		spycloud.UpdateOrCreateBreachRecordFromString(&breachRecord)
	}
}
