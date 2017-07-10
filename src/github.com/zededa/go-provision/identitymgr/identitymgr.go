// Copyright (c) 2017 Zededa, Inc.
// All rights reserved.

// Manage the allocation of EIDs for application instances based on input
// as EIDConfig structs in /var/tmp/identitymgr/config/*.json and report
// on status in the collection of EIDStatus structs in
// /var/run/identitymgr/status/*.json

package main

import (
	"encoding/json"
	"fmt"
	"github.com/zededa/go-provision/types"
	"github.com/zededa/go-provision/watch"
	"io/ioutil"
	"log"
	"os"
// XXX	"os/exec"
// XXX	"strconv"
	"strings"
)

func main() {
	// XXX make baseDirname and runDirname be arguments??
	// Keeping status in /var/run to be clean after a crash/reboot
	baseDirname := "/var/tmp/identitymgr"
	runDirname := "/var/run/identitymgr"
	configDirname := baseDirname + "/config"
	statusDirname := runDirname + "/status"
	
	if _, err := os.Stat(runDirname); err != nil {
		if err := os.Mkdir(runDirname, 0755); err != nil {
			log.Fatal("Mkdir ", runDirname, err)
		}
	}
	if _, err := os.Stat(statusDirname); err != nil {
		if err := os.Mkdir(statusDirname, 0755); err != nil {
			log.Fatal("Mkdir ", statusDirname, err)
		}
	}

	// XXX this is common code except for the types used with json
	
	fileChanges := make(chan string)
	go watch.WatchConfigStatus(configDirname, statusDirname, fileChanges)
	for {
		change := <-fileChanges
		parts := strings.Split(change, " ")
		operation := parts[0]
		fileName := parts[1]
		if !strings.HasSuffix(fileName, ".json") {
			log.Printf("Ignoring file <%s>\n", fileName)
			continue
		}
		if operation == "D" {
			statusFile := statusDirname + "/" + fileName
			if _, err := os.Stat(statusFile); err != nil {
				// File just vanished!
				log.Printf("File disappeared <%s>\n", fileName)
				continue
			}
			sb, err := ioutil.ReadFile(statusFile)
			if err != nil {
				log.Printf("%s for %s\n", err, statusFile)
				continue
			}
			status := types.EIDStatus{}
			if err := json.Unmarshal(sb, &status); err != nil {
				log.Printf("%s EIDStatus file: %s\n",
					err, statusFile)
				continue
			}
			uuid := status.UUIDandVersion.UUID
			if uuid.String()+".json" != fileName {
				log.Printf("Mismatch between filename and contained uuid: %s vs. %s\n",
					fileName, uuid.String())
				continue
			}
			statusName := statusDirname + "/" + fileName
			handleDelete(statusName, status)
			continue
		}
		if operation != "M" {
			log.Fatal("Unknown operation from Watcher: ", operation)
		}
		configFile := configDirname + "/" + fileName
		cb, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Printf("%s for %s\n", err, configFile)
			continue
		}
		config := types.EIDConfig{}
		if err := json.Unmarshal(cb, &config); err != nil {
			log.Printf("%s EIDConfig file: %s\n",
				err, configFile)
			continue
		}
		uuid := config.UUIDandVersion.UUID
		if uuid.String()+".json" != fileName {
			log.Printf("Mismatch between filename and contained uuid: %s vs. %s\n",
				fileName, uuid.String())
			continue
		}
		statusFile := statusDirname + "/" + fileName
		if _, err := os.Stat(statusFile); err != nil {
			// File does not exist in status hence new
			statusName := statusDirname + "/" + fileName
			handleCreate(statusName, config)
			continue
		}
		// Compare Version string
		sb, err := ioutil.ReadFile(statusFile)
		if err != nil {
			log.Printf("%s for %s\n", err, statusFile)
			continue
		}
		status := types.EIDStatus{}
		if err := json.Unmarshal(sb, &status); err != nil {
			log.Printf("%s EIDStatus file: %s\n",
				err, statusFile)
			continue
		}
		uuid = status.UUIDandVersion.UUID
		if uuid.String()+".json" != fileName {
			log.Printf("Mismatch between filename and contained uuid: %s vs. %s\n",
				fileName, uuid.String())
			continue
		}
		// Look for pending* in status and repeat that operation.
		// XXX After that do a full ReadDir to restart ...
		if status.PendingAdd {
			statusName := statusDirname + "/" + fileName
			handleCreate(statusName, config)
			// XXX set something to rescan?
			continue
		}
		if status.PendingDelete {
			statusName := statusDirname + "/" + fileName
			handleDelete(statusName, status)
			// XXX set something to rescan?
			continue
		}
		if status.PendingModify {
			statusName := statusDirname + "/" + fileName
			handleModify(statusName, config, status)
			// XXX set something to rescan?
			continue
		}
			
		if config.UUIDandVersion.Version ==
			status.UUIDandVersion.Version {
			fmt.Printf("Same version %s for %s\n",
				config.UUIDandVersion.Version,
				fileName)
			continue
		}
		statusName := statusDirname + "/" + fileName
		handleModify(statusName, config, status)
	}
}

func writeEIDStatus(status *types.EIDStatus,
	statusFilename string) {
	b, err := json.Marshal(status)
	if err != nil {
		log.Fatal(err, "json Marshal EIDStatus")
	}
	// We assume a /var/run path hence we don't need to worry about
	// partial writes/empty files due to a kernel crash.
	// XXX which permissions?
	err = ioutil.WriteFile(statusFilename, b, 0644)
	if err != nil {
		log.Fatal(err, statusFilename)
	}
}

func handleCreate(statusFilename string, config types.EIDConfig) {
	fmt.Printf("handleCreate(%v) for %s\n",
		config.UUIDandVersion, config.DisplayName)

	// Start by marking with PendingAdd
	status := types.EIDStatus{
		UUIDandVersion: config.UUIDandVersion,
		PendingAdd:     true,
		DisplayName:    config.DisplayName,
	}
	writeEIDStatus(&status, statusFilename)
	// XXX do work
	status.PendingAdd = false
	writeEIDStatus(&status, statusFilename)
}

// Need to compare what might have changed. If any content change
// then we need to reboot. Thus version by itself can change but nothing
// else. Such a version change would be e.g. due to an ACL change.
func handleModify(statusFilename string, config types.EIDConfig,
	status types.EIDStatus) {
	fmt.Printf("handleModify(%v) for %s\n",
		config.UUIDandVersion, config.DisplayName)

	status.PendingModify = true
	writeEIDStatus(&status, statusFilename)
	// XXX Any work?
	status.PendingModify = false
	writeEIDStatus(&status, statusFilename)
}

// Need the olNum and ulNum to delete and EID route to delete
func handleDelete(statusFilename string, status types.EIDStatus) {
	fmt.Printf("handleDelete(%v) for %s\n",
		status.UUIDandVersion, status.DisplayName)

	status.PendingDelete = true
	writeEIDStatus(&status, statusFilename)
	// XXX Do work?
	status.PendingDelete = false
	writeEIDStatus(&status, statusFilename)
}


