all :
	go build -o play_icfp2015

test : all
	./play_icfp2015 -d=true -f p1.json

