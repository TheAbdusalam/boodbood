#!/bin/zsh

bd() {
	boodbood $1 $2
	source $HOME/.config/boodbood/actions/run.sh
	rm -f $HOME/.config/boodbood/actions/run.sh
}
