#!/usr/bin/env bash

{ # this ensures the entire script is downloaded #

aish_has() {
  type "$1" > /dev/null 2>&1
}

aish_echo() {
  command printf %s\\n "$*" 2>/dev/null
}

if [ -z "${BASH_VERSION}" ] || [ -n "${ZSH_VERSION}" ]; then
  # shellcheck disable=SC2016
  aish_echo >&2 'Error: the install instructions explicitly say to pipe the install script to `bash`; please follow them'
  exit 1
fi

get_os() {
  if [[ "$(uname)" == "Linux" ]]; then
    echo "linux"
  elif [[ "$(uname)" == "Darwin" ]] && [[ "$(uname -m)" == "x86_64" ]]; then
    echo "darwin-amd64"
  elif [[ "$(uname)" == "Darwin" ]] && [[ "$(uname -m)" == "arm64" ]]; then
    echo "darwin-arm64"
  else
    echo "windows"
  fi
}

GITHUB_REPO="aiomni/aish"
RELEASE_API="https://api.github.com/repos/${GITHUB_REPO}/releases/latest"

echo "get the latest release information..."
RELEASE_INFO=$(curl -s ${RELEASE_API})

os=$(get_os)

TAG_NAME=$(echo ${RELEASE_INFO} | grep -o '"tag_name": *"[^"]*"' | awk -F'"' '{print $4}')
DOWNLOAD_URL="https://github.com/aiomni/aish/releases/download/${TAG_NAME}/aish-${os}"


echo "Download latest version: ${TAG_NAME}..."
sudo curl -o /usr/local/bin/aish -L ${DOWNLOAD_URL}
sudo chmod +x /usr/local/bin/aish
} # this ensures the entire script is downloaded #
