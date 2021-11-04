package task

import (
	"strings"

	"github.com/rs/zerolog/log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal().Err(err)
	}
}

func getTag(title string) string {
	if strings.Contains(title, "!") {
		return HIGH.String()
	} else if strings.Contains(title, "#") {
		return LOW.String()
	}
	return MEDIUM.String()
}

func trimedTitle(title string) string {
  tag := getTag(title)
  if tag != MEDIUM.String() {
    return title[1:]
  }
  return title
}

func Search(length int, f func(index int) bool) int {
  for index := 0; index < length; index++ {
    if f(index) {
      return index
    }
  }
  return -1
}
