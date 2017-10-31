# Changelog

## v0.0.5 [17/10/31]
- 17/10/31 chore(all): update to golang 1.9, mongo 3.4 and dep (SFR)
           refact(make): refactor the makefile for automatic help message
           test(web): add postman collection for testing
- 17/05/07 chore(docker): add support for multi stage docker build to reduce image size (SFR)
## v0.0.4 [17/03/28]
- 17/03/28 refact(dao): add typed spirit types and interface compile time check (SFR)
           refact(model): add typed spirit types
           refact(web): rename handler to controller
           fix(main): fix JSON time marshalling issue with time zone conversion
           chore(make): prevent glide lock file removal and force docker rm image
           chore(all): make version 0.0.4
- 17/03/06 chore(docker): update DockerFile to use Go alpine 1.8 (SFR)
	   chore(doc): update architecture picture to remove Go 1.5
- 17/01/31 chore(vendor): update package manager to glide and vendor version (SFR)
           chore(doc): add the link to the slides in the README
- 16/11/29 chore(all): update log formatter, cli dependencies and editorconfig (SFR)
           chore(build): remove build in go 1.6 incompatibility gorilla mux http context

## v0.0.3 [16/10/20]
- 16/10/20 chore(build): adjust mongo, go, docker version (SFR)
- 16/10/20 chore(build): Makefile state of the art, query script (SFR)

## v0.0.1 [16/08/09]
- 16/08/09 refact(all): update after fork from sebastienfr (SFR)
