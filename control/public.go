package control

import (
	"math/rand"
	"time"
)

// UserAgents 是不同种类的用户代理
var UserAgents = map[string]string{
	"Default": "R28tWXVKaWFuLzEuMCAoaHR0cHM6Ly9naXRodWIuY29tL0JveUNoYWkvR28tWXVKaWFuLmdpdCk=",
	"Google":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	"Edge":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
	"Firefox": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:125.0) Gecko/20100101 Firefox/125.0",
}

func (w *workCfg) getUserAgent(s []string) (string, string) {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	var candidates []string
	var types []string

	for _, item := range s {
		if agent, exists := UserAgents[item]; exists {
			candidates = append(candidates, agent)
			types = append(types, item)
		}
	}

	if len(candidates) == 0 {
		return "", ""
	}

	choice := rnd.Intn(len(candidates))
	return candidates[choice], types[choice]
}
