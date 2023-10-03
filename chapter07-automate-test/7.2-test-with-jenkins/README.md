## Test with Jenkins

1. Create Project (or copy from build-tcir and modify Build-Step)
Name: test-tcir
Git: https://github.com/3dsinteractive/tcir-app.git
Branch: */main
Build-Step (shell) : ./test.sh

2. Modify ./test.sh to ./test-fail.sh to show if test is failed

3. Update project build-tcir to trigger only if test-tcir stable

4. Update test-tcir to poll SCM every minute