#!/bin/env bash

# go install github.com/fyne-io/fyne-cross@latest
# fyne-cross windows --pull
fyne-cross windows -arch=amd64 -output swrpg_calc_windows_amd64.exe --app-id matgd.github.com.starwarsrpg.dicehelper ./src