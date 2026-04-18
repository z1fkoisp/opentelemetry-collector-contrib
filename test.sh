#!/bin/bash
COMPONENT='$(touch /tmp/pwned)'
get_codeowners() {
    echo "Running: grep \"^${1}\" .github/CODEOWNERS"
    grep "^${1}" .github/CODEOWNERS || true
}
get_codeowners "${COMPONENT}"
