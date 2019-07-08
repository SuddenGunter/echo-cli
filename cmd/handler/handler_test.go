package handler

import (
	"testing"

	"github.com/spf13/cobra"

	"github.com/stretchr/testify/require"
)

func Test_Combine_WhenCalledWithTwoMethods_CallsThem(t *testing.T) {
	firstCalled := false
	first := func(cmd *cobra.Command, args []string) error {
		firstCalled = true
		return nil
	}

	secondCalled := false
	second := func(cmd *cobra.Command, args []string) error {
		secondCalled = true
		return nil
	}

	resultMethod := Combine(first, second)
	err := resultMethod(nil, nil)

	require.Nil(t, err)
	require.True(t, firstCalled)
	require.True(t, secondCalled)
}

// this method allows to test order in which Combine calls combined methods
func Test_Combine_WhenCalledWithTwoMethods_CallsThemInOrder(t *testing.T) {
	blocker := make(chan struct{})
	runner := make(chan struct{})
	firstCalled := false
	first := func(cmd *cobra.Command, args []string) error {
		<-runner
		firstCalled = true
		blocker <- struct{}{}
		return nil
	}

	secondCalled := false
	second := func(cmd *cobra.Command, args []string) error {
		<-runner
		secondCalled = true
		blocker <- struct{}{}
		return nil
	}

	resultMethod := Combine(first, second)
	go resultMethod(nil, nil)

	require.False(t, firstCalled)
	require.False(t, secondCalled)
	runner <- struct{}{}
	<-blocker
	require.True(t, firstCalled)
	require.False(t, secondCalled)
	runner <- struct{}{}
	<-blocker
	require.True(t, firstCalled)
	require.True(t, secondCalled)
}
