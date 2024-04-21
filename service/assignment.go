package services

import (
	"LavanderiaBackend/repository"
	"github.com/google/uuid"
	"log"
	"time"
)

type AssignmentService struct {
	Repo *repository.WashingMachineRepository
}

func NewAssignmentService(repo *repository.WashingMachineRepository) *AssignmentService {
	return &AssignmentService{Repo: repo}
}

func (as *AssignmentService) StartAssignmentProcess() {
	for {
		time.Sleep(6 * time.Minute) // Check every 10 minutes
		requests, err := as.Repo.FetchWashingRequests()
		if err != nil {
			log.Printf("Error fetching requests: %v", err)
			continue
		}
		for _, req := range requests {
			if !req.RequiresWashing() {
				continue
			}
			machine, err := as.Repo.GetAvailableMachine()
			if err != nil {
				log.Printf("No available machines: %v", err)
				continue
			}
			err = as.Repo.AssignMachineToRequest(*machine, req.Id)
			if err != nil {
				log.Printf("Failed to assign machine: %v", err)
				continue
			}
			log.Printf("Assigned machine %s to request %s", machine.Id, req.Id)
			go as.handleServiceCompletion(machine.Id)
		}
	}
}

func (as *AssignmentService) handleServiceCompletion(machineId uuid.UUID) {
	time.Sleep(5 * time.Minute) // Simulate service time
	err := as.Repo.SetMachineAvailable(machineId)
	if err != nil {
		log.Printf("Error setting machine to available: %v", err)
		return
	}
	log.Printf("Machine %s is now available", machineId)
}
