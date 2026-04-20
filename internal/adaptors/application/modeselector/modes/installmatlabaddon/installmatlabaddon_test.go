// Copyright 2026 The MathWorks, Inc.

package installmatlabaddon_test

import (
	"bytes"
	"testing"

	"github.com/matlab/matlab-mcp-core-server/internal/adaptors/application/modeselector/modes/installmatlabaddon"
	"github.com/matlab/matlab-mcp-core-server/internal/messages"
	"github.com/matlab/matlab-mcp-core-server/internal/testutils"
	installmatlabaddonmocks "github.com/matlab/matlab-mcp-core-server/mocks/adaptors/application/modeselector/modes/installmatlabaddon"
	entitiesmocks "github.com/matlab/matlab-mcp-core-server/mocks/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_HappyPath(t *testing.T) {
	// Arrange
	mockOSLayer := &installmatlabaddonmocks.MockOSLayer{}
	defer mockOSLayer.AssertExpectations(t)

	mockMessageCatalog := &installmatlabaddonmocks.MockMessageCatalog{}
	defer mockMessageCatalog.AssertExpectations(t)

	mockLoggerFactory := &installmatlabaddonmocks.MockLoggerFactory{}
	defer mockLoggerFactory.AssertExpectations(t)

	mockWatchdogClient := &installmatlabaddonmocks.MockWatchdogClient{}
	defer mockWatchdogClient.AssertExpectations(t)

	mockGlobalMATLAB := &installmatlabaddonmocks.MockGlobalMATLAB{}
	defer mockGlobalMATLAB.AssertExpectations(t)

	mockAddonManager := &installmatlabaddonmocks.MockAddonManager{}
	defer mockAddonManager.AssertExpectations(t)

	// Act
	mode := installmatlabaddon.New(mockOSLayer, mockMessageCatalog, mockLoggerFactory, mockWatchdogClient, mockGlobalMATLAB, mockAddonManager)

	// Assert
	assert.NotNil(t, mode)
}

func TestMode_StartAndWaitForCompletion_HappyPath(t *testing.T) {
	// Arrange
	mockOSLayer := &installmatlabaddonmocks.MockOSLayer{}
	defer mockOSLayer.AssertExpectations(t)

	mockMessageCatalog := &installmatlabaddonmocks.MockMessageCatalog{}
	defer mockMessageCatalog.AssertExpectations(t)

	mockLoggerFactory := &installmatlabaddonmocks.MockLoggerFactory{}
	defer mockLoggerFactory.AssertExpectations(t)

	mockWatchdogClient := &installmatlabaddonmocks.MockWatchdogClient{}
	defer mockWatchdogClient.AssertExpectations(t)

	mockGlobalMATLAB := &installmatlabaddonmocks.MockGlobalMATLAB{}
	defer mockGlobalMATLAB.AssertExpectations(t)

	mockAddonManager := &installmatlabaddonmocks.MockAddonManager{}
	defer mockAddonManager.AssertExpectations(t)

	mockClient := &entitiesmocks.MockMATLABSessionClient{}
	defer mockClient.AssertExpectations(t)

	mockLogger := testutils.NewInspectableLogger()

	expectedCtx := t.Context()
	successMessage := "Successfully installed MATLAB Add-On."

	mockLoggerFactory.EXPECT().
		GetGlobalLogger().
		Return(mockLogger, nil).
		Once()

	mockWatchdogClient.EXPECT().
		Start().
		Return(nil).
		Once()

	mockWatchdogClient.EXPECT().
		Stop().
		Return(nil).
		Once()

	mockGlobalMATLAB.EXPECT().
		Client(expectedCtx, mockLogger.AsMockArg()).
		Return(mockClient, nil).
		Once()

	mockAddonManager.EXPECT().
		Install(expectedCtx, mockLogger.AsMockArg(), mockClient).
		Return(nil).
		Once()

	mockMessageCatalog.EXPECT().
		Get(messages.CLIMessages_SuccessfullyInstalledMATLABAddOn).
		Return(successMessage).
		Once()

	mockOSLayer.EXPECT().
		Stdout().
		Return(&bytes.Buffer{}).
		Once()

	mode := installmatlabaddon.New(mockOSLayer, mockMessageCatalog, mockLoggerFactory, mockWatchdogClient, mockGlobalMATLAB, mockAddonManager)

	// Act
	err := mode.StartAndWaitForCompletion(expectedCtx)

	// Assert
	require.NoError(t, err)
}

func TestMode_StartAndWaitForCompletion_LoggerFactoryError(t *testing.T) {
	// Arrange
	mockOSLayer := &installmatlabaddonmocks.MockOSLayer{}
	defer mockOSLayer.AssertExpectations(t)

	mockMessageCatalog := &installmatlabaddonmocks.MockMessageCatalog{}
	defer mockMessageCatalog.AssertExpectations(t)

	mockLoggerFactory := &installmatlabaddonmocks.MockLoggerFactory{}
	defer mockLoggerFactory.AssertExpectations(t)

	mockWatchdogClient := &installmatlabaddonmocks.MockWatchdogClient{}
	defer mockWatchdogClient.AssertExpectations(t)

	mockGlobalMATLAB := &installmatlabaddonmocks.MockGlobalMATLAB{}
	defer mockGlobalMATLAB.AssertExpectations(t)

	mockAddonManager := &installmatlabaddonmocks.MockAddonManager{}
	defer mockAddonManager.AssertExpectations(t)

	mockLoggerFactory.EXPECT().
		GetGlobalLogger().
		Return(nil, messages.AnError).
		Once()

	mode := installmatlabaddon.New(mockOSLayer, mockMessageCatalog, mockLoggerFactory, mockWatchdogClient, mockGlobalMATLAB, mockAddonManager)

	// Act
	err := mode.StartAndWaitForCompletion(t.Context())

	// Assert
	require.ErrorIs(t, err, messages.AnError)
}

func TestMode_StartAndWaitForCompletion_WatchdogStartError(t *testing.T) {
	// Arrange
	mockOSLayer := &installmatlabaddonmocks.MockOSLayer{}
	defer mockOSLayer.AssertExpectations(t)

	mockMessageCatalog := &installmatlabaddonmocks.MockMessageCatalog{}
	defer mockMessageCatalog.AssertExpectations(t)

	mockLoggerFactory := &installmatlabaddonmocks.MockLoggerFactory{}
	defer mockLoggerFactory.AssertExpectations(t)

	mockWatchdogClient := &installmatlabaddonmocks.MockWatchdogClient{}
	defer mockWatchdogClient.AssertExpectations(t)

	mockGlobalMATLAB := &installmatlabaddonmocks.MockGlobalMATLAB{}
	defer mockGlobalMATLAB.AssertExpectations(t)

	mockAddonManager := &installmatlabaddonmocks.MockAddonManager{}
	defer mockAddonManager.AssertExpectations(t)

	mockLogger := testutils.NewInspectableLogger()

	mockLoggerFactory.EXPECT().
		GetGlobalLogger().
		Return(mockLogger, nil).
		Once()

	mockWatchdogClient.EXPECT().
		Start().
		Return(assert.AnError).
		Once()

	mode := installmatlabaddon.New(mockOSLayer, mockMessageCatalog, mockLoggerFactory, mockWatchdogClient, mockGlobalMATLAB, mockAddonManager)

	// Act
	err := mode.StartAndWaitForCompletion(t.Context())

	// Assert
	expectedError := messages.New_AddonManagerErrors_InstallFailed_Error()
	require.Equal(t, expectedError, err)
}

func TestMode_StartAndWaitForCompletion_MATLABClientError(t *testing.T) {
	// Arrange
	mockOSLayer := &installmatlabaddonmocks.MockOSLayer{}
	defer mockOSLayer.AssertExpectations(t)

	mockMessageCatalog := &installmatlabaddonmocks.MockMessageCatalog{}
	defer mockMessageCatalog.AssertExpectations(t)

	mockLoggerFactory := &installmatlabaddonmocks.MockLoggerFactory{}
	defer mockLoggerFactory.AssertExpectations(t)

	mockWatchdogClient := &installmatlabaddonmocks.MockWatchdogClient{}
	defer mockWatchdogClient.AssertExpectations(t)

	mockGlobalMATLAB := &installmatlabaddonmocks.MockGlobalMATLAB{}
	defer mockGlobalMATLAB.AssertExpectations(t)

	mockAddonManager := &installmatlabaddonmocks.MockAddonManager{}
	defer mockAddonManager.AssertExpectations(t)

	mockLogger := testutils.NewInspectableLogger()

	expectedCtx := t.Context()

	mockLoggerFactory.EXPECT().
		GetGlobalLogger().
		Return(mockLogger, nil).
		Once()

	mockWatchdogClient.EXPECT().
		Start().
		Return(nil).
		Once()

	mockWatchdogClient.EXPECT().
		Stop().
		Return(nil).
		Once()

	mockGlobalMATLAB.EXPECT().
		Client(expectedCtx, mockLogger.AsMockArg()).
		Return(nil, assert.AnError).
		Once()

	mode := installmatlabaddon.New(mockOSLayer, mockMessageCatalog, mockLoggerFactory, mockWatchdogClient, mockGlobalMATLAB, mockAddonManager)

	// Act
	err := mode.StartAndWaitForCompletion(expectedCtx)

	// Assert
	expectedError := messages.New_AddonManagerErrors_InstallFailed_Error()
	require.Equal(t, expectedError, err)
}

func TestMode_StartAndWaitForCompletion_AddonManagerInstallError(t *testing.T) {
	// Arrange
	mockOSLayer := &installmatlabaddonmocks.MockOSLayer{}
	defer mockOSLayer.AssertExpectations(t)

	mockMessageCatalog := &installmatlabaddonmocks.MockMessageCatalog{}
	defer mockMessageCatalog.AssertExpectations(t)

	mockLoggerFactory := &installmatlabaddonmocks.MockLoggerFactory{}
	defer mockLoggerFactory.AssertExpectations(t)

	mockWatchdogClient := &installmatlabaddonmocks.MockWatchdogClient{}
	defer mockWatchdogClient.AssertExpectations(t)

	mockGlobalMATLAB := &installmatlabaddonmocks.MockGlobalMATLAB{}
	defer mockGlobalMATLAB.AssertExpectations(t)

	mockAddonManager := &installmatlabaddonmocks.MockAddonManager{}
	defer mockAddonManager.AssertExpectations(t)

	mockClient := &entitiesmocks.MockMATLABSessionClient{}
	defer mockClient.AssertExpectations(t)

	mockLogger := testutils.NewInspectableLogger()

	expectedCtx := t.Context()

	mockLoggerFactory.EXPECT().
		GetGlobalLogger().
		Return(mockLogger, nil).
		Once()

	mockWatchdogClient.EXPECT().
		Start().
		Return(nil).
		Once()

	mockWatchdogClient.EXPECT().
		Stop().
		Return(nil).
		Once()

	mockGlobalMATLAB.EXPECT().
		Client(expectedCtx, mockLogger.AsMockArg()).
		Return(mockClient, nil).
		Once()

	mockAddonManager.EXPECT().
		Install(expectedCtx, mockLogger.AsMockArg(), mockClient).
		Return(messages.AnError).
		Once()

	mode := installmatlabaddon.New(mockOSLayer, mockMessageCatalog, mockLoggerFactory, mockWatchdogClient, mockGlobalMATLAB, mockAddonManager)

	// Act
	err := mode.StartAndWaitForCompletion(expectedCtx)

	// Assert
	require.ErrorIs(t, err, messages.AnError)
}

func TestMode_StartAndWaitForCompletion_WatchdogStopError(t *testing.T) {
	// Arrange
	mockOSLayer := &installmatlabaddonmocks.MockOSLayer{}
	defer mockOSLayer.AssertExpectations(t)

	mockMessageCatalog := &installmatlabaddonmocks.MockMessageCatalog{}
	defer mockMessageCatalog.AssertExpectations(t)

	mockLoggerFactory := &installmatlabaddonmocks.MockLoggerFactory{}
	defer mockLoggerFactory.AssertExpectations(t)

	mockWatchdogClient := &installmatlabaddonmocks.MockWatchdogClient{}
	defer mockWatchdogClient.AssertExpectations(t)

	mockGlobalMATLAB := &installmatlabaddonmocks.MockGlobalMATLAB{}
	defer mockGlobalMATLAB.AssertExpectations(t)

	mockAddonManager := &installmatlabaddonmocks.MockAddonManager{}
	defer mockAddonManager.AssertExpectations(t)

	mockClient := &entitiesmocks.MockMATLABSessionClient{}
	defer mockClient.AssertExpectations(t)

	mockLogger := testutils.NewInspectableLogger()

	expectedCtx := t.Context()
	successMessage := "Successfully installed MATLAB Add-On."

	mockLoggerFactory.EXPECT().
		GetGlobalLogger().
		Return(mockLogger, nil).
		Once()

	mockWatchdogClient.EXPECT().
		Start().
		Return(nil).
		Once()

	mockWatchdogClient.EXPECT().
		Stop().
		Return(assert.AnError).
		Once()

	mockGlobalMATLAB.EXPECT().
		Client(expectedCtx, mockLogger.AsMockArg()).
		Return(mockClient, nil).
		Once()

	mockAddonManager.EXPECT().
		Install(expectedCtx, mockLogger.AsMockArg(), mockClient).
		Return(nil).
		Once()

	mockMessageCatalog.EXPECT().
		Get(messages.CLIMessages_SuccessfullyInstalledMATLABAddOn).
		Return(successMessage).
		Once()

	mockOSLayer.EXPECT().
		Stdout().
		Return(&bytes.Buffer{}).
		Once()

	mode := installmatlabaddon.New(mockOSLayer, mockMessageCatalog, mockLoggerFactory, mockWatchdogClient, mockGlobalMATLAB, mockAddonManager)

	// Act
	err := mode.StartAndWaitForCompletion(expectedCtx)

	// Assert
	require.NoError(t, err)
}
