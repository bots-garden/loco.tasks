#!/bin/bash
message="updated"

if [[ "$1" == "init" ]]; then
  message="🎉 initialize $2"
fi


if [[ "$1" == "wip" ]]; then
  message="🚧 WIP $2"
fi

if [[ "$1" == "save" ]]; then
  message="📝 save $2"
fi

if [[ "$1" == "gardening" ]]; then
  message="🌺 gardening $2"
fi

if [[ "$1" == "feature" ]]; then
  message="✨ feature $2"
fi

if [[ "$1" == "doc" ]]; then
  message="📝 update doc $2"
fi

if [[ "$1" == "fix" ]]; then
  message="🐛 fix $2"
fi

if [[ "$1" == "config" ]]; then
  message="🔧 config $2"
fi

if [[ "$1" == "docker" ]]; then
  message="🐳 docker $2"
fi

if [[ "$1" == "deploy" ]]; then
  message="🚀 deploy $2"
fi

git add .; git commit -m "${message}"; git push