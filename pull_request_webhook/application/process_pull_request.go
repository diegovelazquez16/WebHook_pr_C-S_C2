package application

import (
	"encoding/json"
	"log"

	"github/pull_request_webhook/domain/value_objects"
)

func ProcessPullRequestEvent(rawData []byte) int {
	var eventPayload value_objects.PullRequestEvent

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Printf("Error al deserializar el payload: %v", err)
		return 403
	}

	log.Printf("Evento pull request recibido con acción: %s", eventPayload.Action)

	if eventPayload.Action == "closed" {
		log.Println("El Pull Request ha sido cerrado.")
	}

	log.Printf("El Pull Request apunta hacia la rama: %s", eventPayload.PullRequest.Base.Ref)

	log.Printf("El Pull Request proviene de la rama: %s", eventPayload.PullRequest.Head.Ref)

	log.Printf("El Pull Request fue creado por: %s", eventPayload.PullRequest.User.UserName)

	log.Printf("El Pull Request se realizó en el repositorio: %s", eventPayload.Repository.Name)

	return 200
}
