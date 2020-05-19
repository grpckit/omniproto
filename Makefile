.PHONY: googleapis
googleapis:
	git clone https://github.com/googleapis/googleapis.git /tmp/googleapis
	cp -R /tmp/googleapis/google protorepo-example/
