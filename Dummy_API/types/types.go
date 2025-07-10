package types

type GithubProfile struct {
    Address  string   `json:"address"`
    Username string   `json:"username"`
    Email    string   `json:"email"`
    Repos    []string `json:"repos"`
    Badges   []string `json:"badges"`
}

type PSNProfile struct {
    Address   string   `json:"address"`
    Username  string   `json:"username"`
    Email     string   `json:"email"`
    Games     []string `json:"games"`
    Trophies  []string `json:"trophies"`
} 