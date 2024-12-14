package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	friendsMap := make(map[string]struct{})
	for _, friend := range friendships[username] {
		friendsMap[friend] = struct{}{}
	}

	suggestions := make(map[string]struct{})
	for friend := range friendsMap {
		for _, mutual := range friendships[friend] {
			if _, isDirectFriend := friendsMap[mutual]; !isDirectFriend && mutual != username {
				suggestions[mutual] = struct{}{}
			}
		}
	}

	var recommended []string
	for suggestion := range suggestions {
		recommended = append(recommended, suggestion)
	}

	return recommended
}
