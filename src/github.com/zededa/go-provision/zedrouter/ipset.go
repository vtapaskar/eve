// Copyright (c) 2017 Zededa, Inc.
// All rights reserved.

// default eid ipset configlet for overlay interface towards domU

package main

import (
	"fmt"       
	"log"
	"os/exec"
	"github.com/zededa/go-provision/types"
)

// Create local IPv6 ipset called "local.ipv6".
func createDefaultIpset() {
	fmt.Printf("createDefaultIpset()\n")
	ipsetName := "local.ipv6"
	if !ipsetExists(ipsetName) {
		if err := ipsetCreate(ipsetName, "hash:net", 6); err != nil {
			log.Fatal("ipset create for ", ipsetName, err)
		}
	} else {
		if err := ipsetFlush(ipsetName); err != nil {
			log.Fatal("ipset flush for ", ipsetName, err)
		}
	}
	prefixes := []string{"fe80::/10", "ff02::/16"}
	for _, prefix := range prefixes {
		err := ipsetAdd(ipsetName, prefix)
		if err != nil {
			log.Println("ipset add ", ipsetName, prefix, err)
		}
	}
}

// Create an ipset called eids.<olIfname> with all the addresses from
// the nameToEidList.
// Would be more polite to return an error then to Fatal
func createEidIpsetConfiglet(olIfname string, nameToEidList []types.NameToEid) {
	fmt.Printf("createEidIpsetConfiglet: olifName %s nameToEidList %v\n",
		olIfname, nameToEidList)
	ipsetName := "eids." + olIfname	
	if !ipsetExists(ipsetName) {
		if err := ipsetCreate(ipsetName, "hash:ip", 6); err != nil {
			log.Fatal("ipset create for ", ipsetName, err)
		}
	} else {
		if err := ipsetFlush(ipsetName); err != nil {
			log.Fatal("ipset flush for ", ipsetName, err)
		}
	}
	for _, ne := range nameToEidList {
		for _, eid := range ne.EIDs {
			err := ipsetAdd(ipsetName, eid.String())
			if err != nil {
				log.Println("ipset add ", ipsetName,
					eid.String(), err)
			}
		}
	}
}

func updateEidIpsetConfiglet(olIfname string,
     oldNameToEidList []types.NameToEid, newNameToEidList []types.NameToEid) {
	fmt.Printf("updateEidIpsetConfiglet: olIfname %s old %v, new %v\n",
		olIfname, oldNameToEidList, newNameToEidList)
	ipsetName := "eids." + olIfname	
		
	// Look for EIDs which should be deleted
	for _, ne := range oldNameToEidList {
		for _, eid := range ne.EIDs {
			if !containsEID(newNameToEidList, eid) {
				err := ipsetDel(ipsetName, eid.String())
				if err != nil {
					log.Println("ipset del ", ipsetName,
						eid.String(), err)
				}
			}
		}
	}
		
	// Look for EIDs which should be added
	for _, ne := range newNameToEidList {
		for _, eid := range ne.EIDs {
			if !containsEID(oldNameToEidList, eid) {
				err := ipsetAdd(ipsetName, eid.String())
				if err != nil {
					log.Println("ipset add ", ipsetName,
						eid.String(), err)
				}
			}
		}
	}
}

func deleteEidIpsetConfiglet(olIfname string, printOnError bool) {
	fmt.Printf("deleteEidIpsetConfiglet: olIfname %s\n", olIfname)
	ipsetName := "eids." + olIfname	

	err := ipsetDestroy(ipsetName)
	if err != nil && printOnError {
		log.Println("ipset destroy ", ipsetName, err)
	}
}

// If doesn't exist create the ipv4/ipv6 pair of sets.
func ipsetCreatePair(ipsetName string) error {
	set4 := "ipv4." + ipsetName
	set6 := "ipv6." + ipsetName
	setType := "hash:ip"
	if !ipsetExists(set4) {
		if err := ipsetCreate(set4, setType, 4); err != nil {
			return err
		}
	}
	if !ipsetExists(set6) {
		if err := ipsetCreate(set6, setType, 6); err != nil {
			return err
		}
	}
	return nil
}

func ipsetCreate(ipsetName string, setType string, ipVer int) error {
	cmd := "ipset"
	family := ""
	if ipVer == 4 {
		family = "inet"
	} else if ipVer == 6 {
		family = "inet6"
	}
	args := []string{"create", ipsetName, setType, "family", family}
	if _, err := exec.Command(cmd, args...).Output(); err != nil {
		return err
	}
	return nil
}

func ipsetDestroy(ipsetName string) error {
	cmd := "ipset"
	args := []string{"destroy", ipsetName}
	if _, err := exec.Command(cmd, args...).Output(); err != nil {
		return err
	}
	return nil
}

func ipsetFlush(ipsetName string) error {
	cmd := "ipset"
	args := []string{"flush", ipsetName}
	if _, err := exec.Command(cmd, args...).Output(); err != nil {
		return err
	}
	return nil
}

func ipsetAdd(ipsetName string, member string) error {
	cmd := "ipset"
	args := []string{"add", ipsetName, member}
	if _, err := exec.Command(cmd, args...).Output(); err != nil {
		return err
	}
	return nil
}

func ipsetDel(ipsetName string, member string) error {
	cmd := "ipset"
	args := []string{"del", ipsetName, member}
	if _, err := exec.Command(cmd, args...).Output(); err != nil {
		return err
	}
	return nil
}

func ipsetExists(ipsetName string) bool {
	cmd := "ipset"
	args := []string{"list", ipsetName}
	if _, err := exec.Command(cmd, args...).Output(); err != nil {
		return false
	}
	return true
}