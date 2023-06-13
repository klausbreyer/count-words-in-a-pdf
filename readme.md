# PDF Word Frequency Analyzer

This project is a simple tool written in Go that analyzes the frequency of words in a PDF document. The program counts and outputs the 100 most frequent words with at least 5 characters, sorted first by frequency and then alphabetically for words with the same frequency. It uses the dslipak/pdf library to extract the text from the PDF.

## Prerequisites

Before you begin, ensure you have met the following requirements:

    You have installed the latest version of Go.
    You have a <Windows/Linux/Mac> machine.

## Installing PDF Word Frequency Analyzer

To install PDF Word Frequency Analyzer, follow these steps:

    In your terminal, run the following command to clone this repository:
        git clone https://github.com/klausbreyer/count-words-in-a-pdf.git
    Then, navigate into the project directory and install the required dependencies:
        cd your-repo-name
        go get github.com/dslipak/pdf

## Using PDF Word Frequency Analyzer

To use PDF Word Frequency Analyzer, follow these steps:

    Run the main.go file with the path to your pdf as an argument:
        go run main.go /path/to/your/pdf

## Contributing to PDF Word Frequency Analyzer

To contribute to PDF Word Frequency Analyzer, follow these steps:

    Fork this repository.
    Create a branch: git checkout -b <branch_name>.
    Make your changes and commit them: git commit -m '<commit_message>'
    Push to the original branch: git push origin <project_name>/<location>
