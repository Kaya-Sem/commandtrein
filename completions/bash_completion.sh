# Bash completion script for commandtrein, needs to be sourced to work

# $1 = name of command -> commandtrein
# $2 = current word being completed
# $3 = word before word being completed

_commandtrein(){
	# Use a cache that will update every week
	cache_dir="${XDG_CACHE_DIR:-$HOME/.cache}/commandtrein"
	file="${cache_dir}/$(date +'%m-%Y').txt"

	if ! [ -f "$file" ]; then 
		mkdir -p "${cache_dir}"
		# Remove older caches
		rm "${cache_dir}/*.txt" 2>/dev/null
		# Assumes that the binary is called commandtrein
		commandtrein search > "$file"
	fi
	mapfile -t COMPREPLY < <(grep "$2" "$file")
}

complete -F _commandtrein commandtrein
