# TMUX Screen Saver

a tmux "screen saver".
i know the samurai logo doesn't look centered but it is ok, the chin is just not in the center.

anyways, you can use it like this

```sh
$ tmux list-sessions -F '#{session_last_attached}@#{session_name}' | sort -r | awk -F@ '{printf "%s@%s\n", $2, $1}' | samurai
```

This will display your last 10 sessions.



https://github.com/user-attachments/assets/fdac6520-7850-43b3-b8de-8f8e1a7acf71


```txt
Usage
samurai reads from stdin and accepts up to 10 lines. each line can have 2 components, one on the right and one on the left, they should be separated by -sep

$ tmux list-sessions -a -F '#{session_name}@#{session_created}' | samurai -sep @

  -noanimate
    	don't animate ascii art
  -rformat string
    	Treats right components as epoch values and formats them with the given format, an empty string disables this (default "3:04PM")
  -sep string
    	component separator (default "@")
  -title string
    	Sessions header (default " Recent sessions ")

```
