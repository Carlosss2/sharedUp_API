package realtime

import (
	"encoding/json"
	"sharedup/app/src/posts/domain/entities"
)

type PostNotifier struct {
	hub *Hub
}

func NewPostNotifier(hub *Hub) *PostNotifier {
	return &PostNotifier{hub: hub}
}

func (p *PostNotifier) NotifyPostCreated(post entities.PostResponse) {

	event := map[string]interface{}{
		"type": "Nueva Publicacion",
		"data": post,
	}

	jsonData, _ := json.Marshal(event)

	p.hub.broadcast <- jsonData
}