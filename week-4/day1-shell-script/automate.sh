#!/bin/bash

dir_name="$1 at $(date)"
mkdir "$dir_name"

mkdir "$dir_name/about_me"
mkdir "$dir_name/about_me/personal"
touch "$dir_name/about_me/personal/facebook.txt"

mkdir "$dir_name/about_me/professional"
touch "$dir_name/about_me/professional/linkedin.txt"

mkdir "$dir_name/my_friends"
touch "$dir_name/my_friends/list_of_my_friends.txt"

mkdir "$dir_name/my_system_info"
touch "$dir_name/my_system_info/about_this_laptop.txt"
touch "$dir_name/my_system_info/internet_connection.txt"

# about this laptop
echo -e "My username: $1" >"$dir_name/my_system_info/about_this_laptop.txt"
echo -e "With host: $(uname -a)" >>"$dir_name/my_system_info/about_this_laptop.txt"

# facebook
echo "https://www.facebook.com/$2" >"$dir_name/about_me/personal/facebook.txt"

# linkedin
echo "https://www.linkedin.com/in/$3" >"$dir_name/about_me/professional/linkedin.txt"

# internet connection
ping_result=$(ping -c 3 google.com)
echo "${ping_result}" >"$dir_name/my_system_info/internet_connection.txt"

# list of my friends
URL="https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt"
echo "$(curl -s ${URL})" >"$dir_name/my_friends/list_of_my_friends.txt"
