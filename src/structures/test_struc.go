package structures

type Tests struct {
    ID        string   `json:"id"`
    UniqueID  string   `json:"unique_id"`
    Name      string   `json:"name"`
    Duration  int      `json:"duration"`
    Questions []string   `json:"questions"`
    State     string   `json:"state"`
    Locked    bool     `json:"locked"`
    Tags      []string `json:"tags"`
}