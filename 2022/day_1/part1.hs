import System.IO
import Data.List.Split
import Text.Read
import Data.Maybe
import Data.List

splitByDoubleNewLine :: String -> [String]
splitByDoubleNewLine = splitOn "\n\n"

splitByNewLine :: String -> [String]
splitByNewLine = splitOn "\n"

readMaybeInt :: String -> Maybe Int
readMaybeInt = readMaybe

parseIntoInt :: [String] -> [Int]
parseIntoInt = mapMaybe readMaybeInt

main :: IO ()
main = do
    input <- readFile "./input.txt"
    let result = sum . take 3 . sortBy (flip compare) . map (sum.parseIntoInt.splitByNewLine) $ splitByDoubleNewLine input
    print result
