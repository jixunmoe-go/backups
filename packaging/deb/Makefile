.PHONY: build copy-copyright

build: copy-copyright

copy-copyright:
	cp ../../LICENSE ./debian/copyright

install:
	install -m 0755 -d $(DESTDIR)$(bindir)/bin
	# TODO: support multi-arch packaging
	install -m 0755 ../../bin/backup-cli-linux-amd64 \
		$(DESTDIR)$(bindir)/bin/backup-cli
