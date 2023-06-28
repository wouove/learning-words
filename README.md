# learning-words

## Introduction
This tool helps in converting saved words on a Kobo ereader to a file format that can be used to study the words in the
[ANKI-app](https://apps.ankiweb.net/). During reading books in English, I quite often come across words that I do not
know the translation of. This set up helps me in studying those words.

1. The tool reads the .sqlite file in which the saved words on the Kobo ereader are stored.
2. Those words are translated using the [DeepL API](https://www.deepl.com/en/docs-api). Only the words that are not yet in the output file will be translated.
3. The combination of words and translations are then stored to a CSV file which can be imported to ANKI.

## How to use
### Environment variables
The following environment variables need to be set:
1. `DEEPL_API_KEY` - The key that you get when creating a DeepL account, [link](https://www.deepl.com/docs-api/api-access/).
2. `INPUT_PATH` - Path to the .sqlite file from the Kobo ereader
3. `OUTPUT_PATH` - File path to store the output CSV

### Running
By setting the `StringFormatter` in `main.go` one can choose between saving the output in the 'classic' card format or in the multiple choice format.
Using the make file, run `make run`