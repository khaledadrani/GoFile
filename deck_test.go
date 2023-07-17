package main

import (
	"fmt"
	"os"
	"testing"
)

func TestDeckLength(t *testing.T) {
	d := createDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func TestElementsInDeck(t *testing.T) {
	d := createDeck()

	if d[0] != "Spades of Ace" {
		t.Errorf("First element must be Spaces of Ace!")
	}

	if d[15] != "Clubs of Three" {
		t.Errorf("Last element must be Clubs of Three!")

	}
}

func handleDeleteFile(fileName string, t *testing.T) {
	if _, err := os.Stat(fileName); err == nil {
		// File exists, delete it
		err := os.Remove(fileName)
		if err != nil {
			t.Errorf("Unable to delete file:\n %s", err)
			return
		}
		fmt.Println("File deleted successfully.")
	}
}

func TestFileCreation(t *testing.T) {
	fileName := "_deck_testing"

	handleDeleteFile(fileName, t)

	d := createDeck()

	err := d.SaveToFile(fileName)

	if err != nil {
		t.Errorf("Error while saving to file:\n %s", err)
		handleDeleteFile(fileName, t)
	}

	d, err = newDeckFromFile(fileName)

	if err != nil {
		t.Errorf("Error while loading from file:\n %s", err)
		handleDeleteFile(fileName, t)
	}

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	handleDeleteFile(fileName, t)

}
