package structures

type AttemptResponse struct {
    Model struct {
        ScoresTagsSplit map[string]string `json:"scores_tags_split"`
    } `json:"model"`
}
