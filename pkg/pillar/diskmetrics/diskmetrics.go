// Copyright (c) 2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package diskmetrics

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// Matches the json output of qemu-img info
type ImgInfo struct {
	VirtualSize uint64 `json:"virtual-size"`
	Filename    string `json:"filename"`
	ClusterSize uint64 `json:"cluster-size"`
	Format      string `json:"format"`
	ActualSize  uint64 `json:"actual-size"`
	DirtyFlag   bool   `json:"dirty-flag"`
}

func GetImgInfo(diskfile string) (*ImgInfo, error) {
	var imgInfo ImgInfo

	if _, err := os.Stat(diskfile); err != nil {
		return nil, err
	}
	output, err := exec.Command("/usr/lib/xen/bin/qemu-img",
		"info", "-U", "--output=json", diskfile).CombinedOutput()
	if err != nil {
		errStr := fmt.Sprintf("qemu-img failed: %s, %s\n",
			err, output)
		return nil, errors.New(errStr)
	}
	if err := json.Unmarshal(output, &imgInfo); err != nil {
		return nil, err
	}
	return &imgInfo, nil
}

func ResizeImg(diskfile string, newsize uint64) error {

	if _, err := os.Stat(diskfile); err != nil {
		return err
	}
	output, err := exec.Command("/usr/lib/xen/bin/qemu-img",
		"resize", diskfile, fmt.Sprintf("%d", newsize)).CombinedOutput()
	if err != nil {
		errStr := fmt.Sprintf("qemu-img failed: %s, %s\n",
			err, output)
		return errors.New(errStr)
	}
	return nil
}
