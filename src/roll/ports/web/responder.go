package web

import (
	"net/http"
	"strconv"

	resources "go-stack-yourself/src/roll/ports/resources"
)

func showResultAnonym(w http.ResponseWriter, r *http.Request, result int) error {
	return resources.RollResultAnonym(strconv.Itoa(result)).Render(r.Context(), w)
}

func showResultUser(w http.ResponseWriter, r *http.Request, p string, result int) error {
	return resources.RollResultUser(p, strconv.Itoa(result)).Render(r.Context(), w)
}

func showDiceRollHome(w http.ResponseWriter, r *http.Request) error {
	return resources.DiceRollHome().Render(r.Context(), w)
}
