# battery-emoji
A tiny Go program that outputs your Macbook's battery state as emoji. For your shell prompt.

![what it looks like](https://cloud.githubusercontent.com/assets/213293/6364186/7e4478ca-bc70-11e4-8ea5-11b99f7ac833.png)

## Installation

Make sure you have Go installed (`brew install go`), then:

    make

It's that easy. Try it out with `./battery-emoji`.

To add it to your `zsh` theme, do something like:

	function battery_charge {
	    echo `~/dotfiles/batcharge` 2>/dev/null
	}

	RPROMPT='$(battery_charge)'
