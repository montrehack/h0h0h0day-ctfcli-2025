.PHONY: dist

dist:
	rm -rf dist
	mkdir dist

	# Copy all project files with redacted flag
	cp -r compose.yml web ./dist
	sed -i 's/h0h0h0{.*}/h0h0h0{redacted}/g' dist/web/Dockerfile

	# Tar it and ship it	
	cd dist && tar -czf ../polar-news.tar.gz *
	rm -rf dist