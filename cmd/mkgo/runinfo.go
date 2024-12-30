package main

import (
	"time"
)

const runInfoFile = "mkgo/run_info.json"

type runInfo struct {
	EntHash      string `json:"ent_hash"`
	CfgHash      string `json:"cfg_hash"`
	LastDepCheck time.Time
}

func (r *runInfo) hashCheck(entHash, cfgHash string) bool {
	return entHash == r.EntHash && cfgHash == r.CfgHash
}

func (r *runInfo) entHashCheck(entHash string) bool {
	return entHash == r.EntHash
}
