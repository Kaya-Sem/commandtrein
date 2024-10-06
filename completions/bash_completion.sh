# Bash completion script for commandtrein, needs to be sourced to work

# $1 = name of command -> commandtrein
# $2 = current word being completed
# $3 = word before word being completed

_commandtrein(){
	# Use a cache that will update every week
	file="$HOME/.cache/commandtrein/$(date +'%m-%Y').txt"

	if ! [ -f "$file" ]; then 
		mkdir -p "$HOME/.cache/commandtrein/"
		# Remove older caches
		rm "$HOME/.cache/commandtrein/*.txt"
		# Assumes that the binary is called commandtrein
		commandtrein search > "$file"
	fi
	mapfile -t COMPREPLY < <(grep "$2" "$file")
}

complete -F _commandtrein commandtrein
