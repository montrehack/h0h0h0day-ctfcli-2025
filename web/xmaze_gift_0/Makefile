default: zip

.PHONY: challenge

zip:
	rm -f backup/FLAG_1_c0ll1sion_f4ilure.zip
# Replace KONAMI code in ./static/js/const.js.old
# Replace FLAG and EASTER_EGG in ./src/flag.go
	zip backup/FLAG_1_c0ll1sion_f4ilure.zip -r ./* -x "./script/*" -x "./solution/*" -x "./challenge/*" -x "./backup/*" -x "./.vscode/*" -x "Makefile" 

challenge: 
	rm -f xmaze-gift.zip
	zip xmaze-gift.zip -r ./* -x "./solution/*"
