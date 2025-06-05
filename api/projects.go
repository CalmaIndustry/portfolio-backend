package api

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

func GetProject(w http.ResponseWriter, req *http.Request) {
	projects := []Project{
		{ID: 1, Title: "Kubernetes & Open Nebula", Description: "Designed and implemented a fully automated pipeline to provision a Kubernetes cluster (RKE2) on an OpenNebula private cloud using Terraform for infrastructure provisioning and Ansible for configuration management."},
		{ID: 2, Title: "Terraform Provider for OpenNebula - Contributor", Description: "Contributed to the development of a Terraform provider for OpenNebula by enhancing the Go codebase. Focused on extending functionality to support modification of physical devices and fixing critical VLAN configuration bugs."},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
