package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestPutAndSearch(t *testing.T) {
    tree := NewTree[string]()

    // Test inserting a single word
    tree.Put("hello", "world")
    value, exists := tree.Search("hello")
    assert.True(t, exists, "Expected 'hello' to exist")
    assert.Equal(t, "world", value, "Expected value for 'hello' to be 'world'")

    // Test inserting multiple words
    tree.Put("hi", "there")
    tree.Put("hey", "you")

    value, exists = tree.Search("hi")
    assert.True(t, exists, "Expected 'hi' to exist")
    assert.Equal(t, "there", value, "Expected value for 'hi' to be 'there'")

    value, exists = tree.Search("hey")
    assert.True(t, exists, "Expected 'hey' to exist")
    assert.Equal(t, "you", value, "Expected value for 'hey' to be 'you'")

    // Test searching for a non-existent word
    _, exists = tree.Search("unknown")
    assert.False(t, exists, "Expected 'unknown' not to exist")
}

func TestPutOverwrite(t *testing.T) {
    tree := NewTree[string]()

    // Test overwriting an existing word
    tree.Put("hello", "world")
    tree.Put("hello", "everyone")

    value, exists := tree.Search("hello")
    assert.True(t, exists, "Expected 'hello' to exist")
    assert.Equal(t, "everyone", value, "Expected value for 'hello' to be 'everyone'")
}

func TestDelete(t *testing.T) {
    tree := NewTree[string]()

    // Insert words into the tree
    tree.Put("hello", "world")
    tree.Put("hi", "there")
    tree.Put("hey", "you")

    // Delete a word and check if it is removed
    tree.Delete("hi")
    _, exists := tree.Search("hi")
    assert.False(t, exists, "Expected 'hi' to be deleted")

    // Ensure other words are still present
    value, exists := tree.Search("hello")
    assert.True(t, exists, "Expected 'hello' to still exist")
    assert.Equal(t, "world", value, "Expected value for 'hello' to be 'world'")

    value, exists = tree.Search("hey")
    assert.True(t, exists, "Expected 'hey' to still exist")
    assert.Equal(t, "you", value, "Expected value for 'hey' to be 'you'")

    // Delete a word that is a prefix of another word
    tree.Put("he", "prefix")
    tree.Delete("he")
    _, exists = tree.Search("he")
    assert.False(t, exists, "Expected 'he' to be deleted")

    // Ensure 'hello' and 'hey' are still present
    value, exists = tree.Search("hello")
    assert.True(t, exists, "Expected 'hello' to still exist")
    assert.Equal(t, "world", value, "Expected value for 'hello' to be 'world'")

    value, exists = tree.Search("hey")
    assert.True(t, exists, "Expected 'hey' to still exist")
    assert.Equal(t, "you", value, "Expected value for 'hey' to be 'you'")
}

func TestDeleteNonExistentWord(t *testing.T) {
    tree := NewTree[string]()

    // Attempt to delete a non-existent word
    tree.Delete("nonexistent")
    // Ensure no panic or error occurs and tree is still functional
    tree.Put("test", "value")
    value, exists := tree.Search("test")
    assert.True(t, exists, "Expected 'test' to exist")
    assert.Equal(t, "value", value, "Expected value for 'test' to be 'value'")
}

func TestGetWordsWithPrefix(t *testing.T) {
    tree := NewTree[string]()

    // Insert words into the tree
    tree.Put("hello", "world")
    tree.Put("hell", "fire")
    tree.Put("heaven", "peace")
    tree.Put("heavy", "metal")
    tree.Put("hero", "brave")

    // Test getting words with prefix "he"
    results := tree.GetWordsWithPrefix("he")
    expected := []Entry[string]{
        {Word: "hell", Value: "fire"},
        {Word: "hello", Value: "world"},
        {Word: "heaven", Value: "peace"},
        {Word: "heavy", Value: "metal"},
        {Word: "hero", Value: "brave"},
    }
    assert.ElementsMatch(t, expected, results, "Expected words with prefix 'he' to match")

    // Test getting words with prefix "hel"
    results = tree.GetWordsWithPrefix("hel")
    expected = []Entry[string]{
        {Word: "hell", Value: "fire"},
        {Word: "hello", Value: "world"},
    }
    assert.ElementsMatch(t, expected, results, "Expected words with prefix 'hel' to match")

    // Test getting words with prefix "hero"
    results = tree.GetWordsWithPrefix("hero")
    expected = []Entry[string]{
        {Word: "hero", Value: "brave"},
    }
    assert.ElementsMatch(t, expected, results, "Expected words with prefix 'hero' to match")

    // Test getting words with a non-existent prefix
    results = tree.GetWordsWithPrefix("nonexistent")
    assert.Empty(t, results, "Expected no words with prefix 'nonexistent'")
}
