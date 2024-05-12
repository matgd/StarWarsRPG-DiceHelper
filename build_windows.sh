#!/bin/env bash

go install github.com/fyne-io/fyne-cross@latest
fyne-cross windows --pull
fyne-cross windows -arch=amd64 -output bin/swrpg_calc_windows_amd64.exe ./src