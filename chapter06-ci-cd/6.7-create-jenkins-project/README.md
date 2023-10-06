## Create Jenkins Project

1. Create Project
Name: build-tcir
Git: https://github.com/3dsinteractive/tcir-app.git
Branch: */main
Build-Step (Execute shell) : ./deploy.sh

2. Set docker user and docker pass 
Username and password (separated)
DOCKER_USER
DOCKER_PASS