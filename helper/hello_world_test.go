package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Fitroh", func(t *testing.T) {
		result := HelloWorld("Fitroh")
		require.Equal(t, "Hello Fitroh",result, "Harusnya Hello Fitroh")
	})
	t.Run("Ikrom", func(t *testing.T) {
		result := HelloWorld("Ikrom")
		require.Equal(t, "Hello Ikrom",result, "Harusnya Hello Fitroh")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before test")
	
	m.Run()
	
    fmt.Println("After test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	} 
}

func TestHelloWorldFitroh(t *testing.T) {
	result := HelloWorld("Fitroh")
	if result != "Hello Fitroh" {
		t.Error("Harusnya Hello Fitroh")
	}
	fmt.Println("TestHelloWorldFitroh Done")
}

func TestHelloWorldDani(t *testing.T) {
	result := HelloWorld("Dani")
	if result != "Hello Dani" {
		t.Fatal("Harusnya Hello Dani")
	}
	fmt.Println("TestHelloWorldDani Done")
}
       
// Testify
func TestHelloWorldArkhan(t *testing.T) {
	result := HelloWorld("Arkhan")
	require.Equal(t, "Hello Arkhan",result, "Harusnya Hello Arkhan")
	fmt.Println("TestHelloWorldArkhan with Assert Done")
}
func TestHelloWorldIqbal(t *testing.T) {
	result := HelloWorld("Iqbal")
	assert.Equal(t, "Hello Iqbal", result, "Harusnya Hello Iqbal")
	fmt.Println("TestHelloWorldIqbal with Assert Done")
}

func TestHelloWorldAlamat(t *testing.T) {
	result := Alamat("Tangerang")
	assert.Equal(t, "Saya tinggal di Tangerang", result, "Harusnya Tangerang")
	fmt.Println("TestHelloWorldAlamat with Assert Done")
}

