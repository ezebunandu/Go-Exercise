package main

import "bookutil/author"

func main() {
	//Create an author instance
	chinuaAchebe := author.NewAuthor("Chinua Achebe", "achebe@example.com")
	//Write and review a chapter.
	chapterTitle := "The Education of a British protected child"
	chapterContent := "The British policy on education tried to force the natives to abandon their cultures and ways of life..."

	chinuaAchebe.WriteChapter(chapterTitle, chapterContent)
	chinuaAchebe.ReviewChapter(chapterTitle, chapterContent)
	chinuaAchebe.FinalizeChapter(chapterTitle)
}
