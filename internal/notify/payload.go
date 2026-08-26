package notify

import "encoding/json"

type Event struct {
	RecipientID string `json:"recipient_id"`
	Kind        string `json:"kind"`
	ObjectID    string `json:"object_id"`
	Message     string `json:"message"`
}

func Encode(e Event) (string, error) { b, err := json.Marshal(e); return string(b), err }
func Decode(payload string) (Event, error) {
	var e Event
	err := json.Unmarshal([]byte(payload), &e)
	return e, err
}
func TopicFor(kind string) string {
	switch kind {
	case "risk":
		return "practice.risk"
	case "attendance":
		return "practice.attendance"
	case "review":
		return "practice.review"
	default:
		return "practice.general"
	}
}
