HOW TO RUN THE WORDLE GAME
1) Make sure you have an IDE downloaded on your device
2) Make sure the latest version of Golang is downloaded on your device
    - https://go.dev/dl/ - Make sure to install based on your corresponding operating system
3) Download the official Go extension on your IDE (if applicable)
    - Search up "Go" in extensions and download the first one
4) Navigate to your bash terminal
    - To ensure Go is installed, type in "go version", if the version pops up, you have successfully downloaded Go
    - Navigate to the directory of the unzipped wordle folder by using cd
        * e.g. cd /Users/username/Downloads/CSC372/wordle
    - Once you are in the directory, run the main.go file by typing out the command "go run main.go" and a window of the wordle game should pop up

HOW TO PLAY THE WORDLE GAME
1) There are 6 rows, 5 columns of input boxes
2) Each row represents a single guess.
3) The objective is to guess the random 5 letter word before your 6 guesses
4) Click each input box in the row to input each letter of your 5 letter guess
5) Every row is disabled except the one you are guessing on
6) After your guess, you will get hints at the bottom of the screen to help you with your guessing
7) To play a new game, exit out of the window and rerun the game (go run main.go)
8) RESTRICTIONS
    - Do not enter more than one character in each input box, if you do input more than one character in each input box, you won't be able to see it, but the game is taking in all those letters
9) MORE
    - The output in the terminal will tell you what word you are guessing towards, you current guess, and how many guesses you have left (all for testing purposes)
