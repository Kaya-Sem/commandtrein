# Bash completion script for commandtrein, needs to be sourced to work

# $1 = name of command -> commandtrein
# $2 = current word being completed
# $3 = word before word being completed

_commandtrein() {
	# Use a cache that will update every week
	cache_dir="${XDG_CACHE_DIR:-$HOME/.cache}/commandtrein"
	file="${cache_dir}/$(date +'%m-%Y').txt"

	if ! [ -f "$file" ]; then
		mkdir -p "${cache_dir}"
		# Remove older caches
		find "${cache_dir}" \
			-maxdepth 1 \
			-type f \
			-name "[0-9][0-9]-2[0-9][0-9][0-9].txt" \
			-delete
		# Assumes that the binary is called commandtrein
		commandtrein search >"$file"
		commandtrein shortcut list -s >>"$file" # custom shortcuts
		echo "shortcut" >>"$file"
	fi

	# Handle subcommand "shortcut"
	if [[ "$3" == "shortcut" ]]; then
		mapfile -t COMPREPLY < <(compgen -W "add list" -- "$2")
		return
	fi

	mapfile -t COMPREPLY < <(grep -i "^$2" "$file")
}

complete -F _commandtrein commandtrein
