Works nice when file to count words in is single. 
But since every feeded files run in separate goroutines, all writes to one object - MAP(words.found)
Fix commented in code.