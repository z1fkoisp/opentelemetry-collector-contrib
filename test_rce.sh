#!/bin/bash
OWNER='@$(touch /tmp/pwned)'
REVIEWERS+=$(echo "${OWNER}" | sed -E 's/@(.+)/"\1"/')
echo "Reviewers: ${REVIEWERS}"
ls /tmp/pwned
