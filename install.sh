#!/bin/bash

if [ $TRAVIS_OS_NAME = 'osx' ]; then

  # Install some custom requirements on macOS
  # e.g. brew install pyenv-virtualenv
  xcode-select --install
else if [ $TRAVIS_OS_NAME = 'linux' ]; then
  # Install some custom requirements on Linux
  sudo apt-get update -y
  sudo apt-get install libgtk-3-dev libwebkit2gtk-4.0-dev -y
  sudo apt-get install build-essential libgtk-3-dev -y
  
else 
  # install dependencies on windows
  choco install -y mingw
fi