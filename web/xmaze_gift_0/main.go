package main

import (
	"log/slog"
	"net/http"
	xmazegift "xmaze-gift/src"
)

func main() {
	http.Handle("/", xmazegift.Static())
	http.HandleFunc("/api/move", xmazegift.MoveElf)
	http.HandleFunc("/api/newPlayer", xmazegift.RegisterNewElf)
	http.HandleFunc("/api/killPlayer", xmazegift.KillElf)
	http.HandleFunc("/api/ancestors", xmazegift.GetAncestors)
	http.HandleFunc("/api/explore", xmazegift.Explore)
	http.HandleFunc("/api/gift", xmazegift.Gift)
	http.HandleFunc("/api/hat", xmazegift.Hat)
	http.HandleFunc("/api/scroll", xmazegift.Scroll)
	http.HandleFunc("/api/b4ckup_d0wnl04D_P4tH", xmazegift.DownloadSrc)

	slog.Info("Starting server on port 25032")
	http.ListenAndServe(":25032", nil)
}
