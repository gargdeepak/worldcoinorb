package main

type CPU struct {
	User float64 `json:"user"`
	System float64 `json:"system"`
	Idle float64 `json:"idle"`
}

type Memory struct {
	Total uint64 `json:"total"`
	Used uint64 `json:"used"`
	Cached uint64 `json:"cached"`
	Free uint64 `json:"free"`
}

type DiskUsage struct{
	Free uint64  `json:"free"`
	Available uint64 `json:"available"`
	Size uint64 `json:"size"`
	Used uint64 `json:"used"`
	Usage float32 `json:"usage"`
}

type OrbUsage struct{

}
