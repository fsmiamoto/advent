import Data.List.Split
import Data.Maybe

data Choice = Rock | Paper | Scissors
    deriving (Eq, Show)

data Result = Win | Draw | Lose
    deriving (Eq, Show)

toChoice :: String -> Choice
toChoice "A" = Rock
toChoice "B" = Paper
toChoice "C" = Scissors

toResult :: String -> Result
toResult "X" = Lose
toResult "Y" = Draw
toResult "Z" = Win

scoreOfResult :: Result -> Int
scoreOfResult Lose = 0
scoreOfResult Draw = 3
scoreOfResult Win  = 6

scoreOfChoice :: Choice -> Int
scoreOfChoice Rock      = 1
scoreOfChoice Paper     = 2
scoreOfChoice Scissors  = 3

loseChoice :: Choice -> Choice
loseChoice Rock     = Scissors
loseChoice Paper    = Rock
loseChoice Scissors = Paper

winChoice :: Choice -> Choice
winChoice Rock     = Paper
winChoice Paper    = Scissors
winChoice Scissors = Rock

drawChoice :: Choice -> Choice
drawChoice x = x

scoreOfResultChoice :: Result -> Choice -> Int
scoreOfResultChoice result = scoreOfChoice . chooser result where  
    chooser Lose = loseChoice
    chooser Draw = drawChoice
    chooser Win  = winChoice

play :: (Choice, Result) -> Int
play (choice, result) = scoreOfResult result + scoreOfResultChoice result choice

splitOnNewLine :: String -> [String]
splitOnNewLine = splitOn "\n"

splitOnSpace :: String -> [String]
splitOnSpace = splitOn " "

toChoiceResult :: [String] -> Maybe (Choice,Result)
toChoiceResult [a, b] = Just (toChoice a ,toResult b)
toChoiceResult _ = Nothing

main :: IO ()
main = do
    input <- readFile "./input.txt"
    let result = sum $ map play $ mapMaybe toChoiceResult $ map splitOnSpace $ splitOnNewLine input
    print result
