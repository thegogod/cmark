package emojis

import "slices"

func Get(alias string) (Emoji, bool) {
	i := slices.IndexFunc(emojis, func(e Emoji) bool {
		return slices.Contains(e.Aliases, alias)
	})

	if i < 0 {
		return Emoji{}, false
	}

	return emojis[i], true
}

func GetByTag(tag string) []Emoji {
	arr := []Emoji{}

	for _, emoji := range emojis {
		if slices.Contains(emoji.Tags, tag) {
			arr = append(arr, emoji)
		}
	}

	return arr
}

func GetByCategory(category string) []Emoji {
	arr := []Emoji{}

	for _, emoji := range emojis {
		if emoji.Category == category {
			arr = append(arr, emoji)
		}
	}

	return arr
}
