package nex

import (
	"reflect"
	"strconv"
)

var errorMask = 1 << 31

type ErrorsStruct struct {
	Core struct {
		Unknown               uint32
		NotImplemented        uint32
		InvalidPointer        uint32
		OperationAborted      uint32
		Exception             uint32
		AccessDenied          uint32
		InvalidHandle         uint32
		InvalidIndex          uint32
		OutOfMemory           uint32
		InvalidArgument       uint32
		Timeout               uint32
		InitializationFailure uint32
		CallInitiationFailure uint32
		RegistrationError     uint32
		BufferOverflow        uint32
		InvalidLockState      uint32
		InvalidSequence       uint32
		SystemError           uint32
		Cancelled             uint32
	}

	DDL struct {
		InvalidSignature uint32
		IncorrectVersion uint32
	}

	RendezVous struct {
		ConnectionFailure                        uint32
		NotAuthenticated                         uint32
		InvalidUsername                          uint32
		InvalidPassword                          uint32
		UsernameAlreadyExists                    uint32
		AccountDisabled                          uint32
		AccountExpired                           uint32
		ConcurrentLoginDenied                    uint32
		EncryptionFailure                        uint32
		InvalidPID                               uint32
		MaxConnectionsReached                    uint32
		InvalidGID                               uint32
		InvalidControlScriptID                   uint32
		InvalidOperationInLiveEnvironment        uint32
		DuplicateEntry                           uint32
		ControlScriptFailure                     uint32
		ClassNotFound                            uint32
		SessionVoid                              uint32
		DDLMismatch                              uint32
		InvalidConfiguration                     uint32
		SessionFull                              uint32
		InvalidGatheringPassword                 uint32
		WithoutParticipationPeriod               uint32
		PersistentGatheringCreationMax           uint32
		PersistentGatheringParticipationMax      uint32
		DeniedByParticipants                     uint32
		ParticipantInBlackList                   uint32
		GameServerMaintenance                    uint32
		OperationPostpone                        uint32
		OutOfRatingRange                         uint32
		ConnectionDisconnected                   uint32
		InvalidOperation                         uint32
		NotParticipatedGathering                 uint32
		MatchmakeSessionUserPasswordUnmatch      uint32
		MatchmakeSessionSystemPasswordUnmatch    uint32
		UserIsOffline                            uint32
		AlreadyParticipatedGathering             uint32
		PermissionDenied                         uint32
		NotFriend                                uint32
		SessionClosed                            uint32
		DatabaseTemporarilyUnavailable           uint32
		InvalidUniqueId                          uint32
		MatchmakingWithdrawn                     uint32
		LimitExceeded                            uint32
		AccountTemporarilyDisabled               uint32
		PartiallyServiceClosed                   uint32
		ConnectionDisconnectedForConcurrentLogin uint32
	}

	PythonCore struct {
		Exception        uint32
		TypeError        uint32
		IndexError       uint32
		InvalidReference uint32
		CallFailure      uint32
		MemoryError      uint32
		KeyError         uint32
		OperationError   uint32
		ConversionError  uint32
		ValidationError  uint32
	}

	Transport struct {
		Unknown                       uint32
		ConnectionFailure             uint32
		InvalidUrl                    uint32
		InvalidKey                    uint32
		InvalidURLType                uint32
		DuplicateEndpoint             uint32
		IOError                       uint32
		Timeout                       uint32
		ConnectionReset               uint32
		IncorrectRemoteAuthentication uint32
		ServerRequestError            uint32
		DecompressionFailure          uint32
		ReliableSendBufferFullFatal   uint32
		UPnPCannotInit                uint32
		UPnPCannotAddMapping          uint32
		NatPMPCannotInit              uint32
		NatPMPCannotAddMapping        uint32
		UnsupportedNAT                uint32
		DnsError                      uint32
		ProxyError                    uint32
		DataRemaining                 uint32
		NoBuffer                      uint32
		NotFound                      uint32
		TemporaryServerError          uint32
		PermanentServerError          uint32
		ServiceUnavailable            uint32
		ReliableSendBufferFull        uint32
		InvalidStation                uint32
		InvalidSubStreamID            uint32
		PacketBufferFull              uint32
		NatTraversalError             uint32
		NatCheckError                 uint32
	}

	DOCore struct {
		StationNotReached             uint32
		TargetStationDisconnect       uint32
		LocalStationLeaving           uint32
		ObjectNotFound                uint32
		InvalidRole                   uint32
		CallTimeout                   uint32
		RMCDispatchFailed             uint32
		MigrationInProgress           uint32
		NoAuthority                   uint32
		NoTargetStationSpecified      uint32
		JoinFailed                    uint32
		JoinDenied                    uint32
		ConnectivityTestFailed        uint32
		Unknown                       uint32
		UnfreedReferences             uint32
		JobTerminationFailed          uint32
		InvalidState                  uint32
		FaultRecoveryFatal            uint32
		FaultRecoveryJobProcessFailed uint32
		StationInconsitency           uint32
		AbnormalMasterState           uint32
		VersionMismatch               uint32
	}

	FPD struct {
		NotInitialized               uint32
		AlreadyInitialized           uint32
		NotConnected                 uint32
		Connected                    uint32
		InitializationFailure        uint32
		OutOfMemory                  uint32
		RmcFailed                    uint32
		InvalidArgument              uint32
		InvalidLocalAccountID        uint32
		InvalidPrincipalID           uint32
		InvalidLocalFriendCode       uint32
		LocalAccountNotExists        uint32
		LocalAccountNotLoaded        uint32
		LocalAccountAlreadyLoaded    uint32
		FriendAlreadyExists          uint32
		FriendNotExists              uint32
		FriendNumMax                 uint32
		NotFriend                    uint32
		FileIO                       uint32
		P2PInternetProhibited        uint32
		Unknown                      uint32
		InvalidState                 uint32
		AddFriendProhibited          uint32
		InvalidAccount               uint32
		BlacklistedByMe              uint32
		FriendAlreadyAdded           uint32
		MyFriendListLimitExceed      uint32
		RequestLimitExceed           uint32
		InvalidMessageID             uint32
		MessageIsNotMine             uint32
		MessageIsNotForMe            uint32
		FriendRequestBlocked         uint32
		NotInMyFriendList            uint32
		FriendListedByMe             uint32
		NotInMyBlacklist             uint32
		IncompatibleAccount          uint32
		BlockSettingChangeNotAllowed uint32
		SizeLimitExceeded            uint32
		OperationNotAllowed          uint32
		NotNetworkAccount            uint32
		NotificationNotFound         uint32
		PreferenceNotInitialized     uint32
		FriendRequestNotAllowed      uint32
	}

	Ranking struct {
		NotInitialized    uint32
		InvalidArgument   uint32
		RegistrationError uint32
		NotFound          uint32
		InvalidScore      uint32
		InvalidDataSize   uint32
		PermissionDenied  uint32
		Unknown           uint32
		NotImplemented    uint32
	}

	Authentication struct {
		NASAuthenticateError             uint32
		TokenParseError                  uint32
		HttpConnectionError              uint32
		HttpDNSError                     uint32
		HttpGetProxySetting              uint32
		TokenExpired                     uint32
		ValidationFailed                 uint32
		InvalidParam                     uint32
		PrincipalIdUnmatched             uint32
		MoveCountUnmatch                 uint32
		UnderMaintenance                 uint32
		UnsupportedVersion               uint32
		ServerVersionIsOld               uint32
		Unknown                          uint32
		ClientVersionIsOld               uint32
		AccountLibraryError              uint32
		ServiceNoLongerAvailable         uint32
		UnknownApplication               uint32
		ApplicationVersionIsOld          uint32
		OutOfService                     uint32
		NetworkServiceLicenseRequired    uint32
		NetworkServiceLicenseSystemError uint32
		NetworkServiceLicenseError3      uint32
		NetworkServiceLicenseError4      uint32
	}

	DataStore struct {
		Unknown             uint32
		InvalidArgument     uint32
		PermissionDenied    uint32
		NotFound            uint32
		AlreadyLocked       uint32
		UnderReviewing      uint32
		Expired             uint32
		InvalidCheckToken   uint32
		SystemFileError     uint32
		OverCapacity        uint32
		OperationNotAllowed uint32
		InvalidPassword     uint32
		ValueNotEqual       uint32
	}

	ServiceItem struct {
		Unknown                  uint32
		InvalidArgument          uint32
		EShopUnknownHttpError    uint32
		EShopResponseParseError  uint32
		NotOwned                 uint32
		InvalidLimitationType    uint32
		ConsumptionRightShortage uint32
	}

	MatchmakeReferee struct {
		Unknown                  uint32
		InvalidArgument          uint32
		AlreadyExists            uint32
		NotParticipatedGathering uint32
		NotParticipatedRound     uint32
		StatsNotFound            uint32
		RoundNotFound            uint32
		RoundArbitrated          uint32
		RoundNotArbitrated       uint32
	}

	Subscriber struct {
		Unknown          uint32
		InvalidArgument  uint32
		OverLimit        uint32
		PermissionDenied uint32
	}

	Ranking2 struct {
		Unknown         uint32
		InvalidArgument uint32
		InvalidScore    uint32
	}

	SmartDeviceVoiceChat struct {
		Unknown                       uint32
		InvalidArgument               uint32
		InvalidResponse               uint32
		InvalidAccessToken            uint32
		Unauthorized                  uint32
		AccessError                   uint32
		UserNotFound                  uint32
		RoomNotFound                  uint32
		RoomNotActivated              uint32
		ApplicationNotSupported       uint32
		InternalServerError           uint32
		ServiceUnavailable            uint32
		UnexpectedError               uint32
		UnderMaintenance              uint32
		ServiceNoLongerAvailable      uint32
		AccountTemporarilyDisabled    uint32
		PermissionDenied              uint32
		NetworkServiceLicenseRequired uint32
		AccountLibraryError           uint32
		GameModeNotFound              uint32
	}

	Screening struct {
		Unknown         uint32
		InvalidArgument uint32
		NotFound        uint32
	}

	Custom struct {
		Unknown uint32
	}

	Ess struct {
		Unknown                uint32
		GameSessionError       uint32
		GameSessionMaintenance uint32
	}
}

var ErrorNames = map[uint32]string{}

var Errors ErrorsStruct

func InitErrorsData() {
	Errors.Core.Unknown = 0x00010001
	Errors.Core.NotImplemented = 0x00010002
	Errors.Core.InvalidPointer = 0x00010003
	Errors.Core.OperationAborted = 0x00010004
	Errors.Core.Exception = 0x00010005
	Errors.Core.AccessDenied = 0x00010006
	Errors.Core.InvalidHandle = 0x00010007
	Errors.Core.InvalidIndex = 0x00010008
	Errors.Core.OutOfMemory = 0x00010009
	Errors.Core.InvalidArgument = 0x0001000A
	Errors.Core.Timeout = 0x0001000B
	Errors.Core.InitializationFailure = 0x0001000C
	Errors.Core.CallInitiationFailure = 0x0001000D
	Errors.Core.RegistrationError = 0x0001000E
	Errors.Core.BufferOverflow = 0x0001000F
	Errors.Core.InvalidLockState = 0x00010010
	Errors.Core.InvalidSequence = 0x00010011
	Errors.Core.SystemError = 0x00010012
	Errors.Core.Cancelled = 0x00010013

	Errors.DDL.InvalidSignature = 0x00020001
	Errors.DDL.IncorrectVersion = 0x00020002

	Errors.RendezVous.ConnectionFailure = 0x00030001
	Errors.RendezVous.NotAuthenticated = 0x00030002
	Errors.RendezVous.InvalidUsername = 0x00030064
	Errors.RendezVous.InvalidPassword = 0x00030065
	Errors.RendezVous.UsernameAlreadyExists = 0x00030066
	Errors.RendezVous.AccountDisabled = 0x00030067
	Errors.RendezVous.AccountExpired = 0x00030068
	Errors.RendezVous.ConcurrentLoginDenied = 0x00030069
	Errors.RendezVous.EncryptionFailure = 0x0003006A
	Errors.RendezVous.InvalidPID = 0x0003006B
	Errors.RendezVous.MaxConnectionsReached = 0x0003006C
	Errors.RendezVous.InvalidGID = 0x0003006D
	Errors.RendezVous.InvalidControlScriptID = 0x0003006E
	Errors.RendezVous.InvalidOperationInLiveEnvironment = 0x0003006F
	Errors.RendezVous.DuplicateEntry = 0x00030070
	Errors.RendezVous.ControlScriptFailure = 0x00030071
	Errors.RendezVous.ClassNotFound = 0x00030072
	Errors.RendezVous.SessionVoid = 0x00030073
	Errors.RendezVous.DDLMismatch = 0x00030075
	Errors.RendezVous.InvalidConfiguration = 0x00030076
	Errors.RendezVous.SessionFull = 0x000300C8
	Errors.RendezVous.InvalidGatheringPassword = 0x000300C9
	Errors.RendezVous.WithoutParticipationPeriod = 0x000300CA
	Errors.RendezVous.PersistentGatheringCreationMax = 0x000300CB
	Errors.RendezVous.PersistentGatheringParticipationMax = 0x000300CC
	Errors.RendezVous.DeniedByParticipants = 0x000300CD
	Errors.RendezVous.ParticipantInBlackList = 0x000300CE
	Errors.RendezVous.GameServerMaintenance = 0x000300CF
	Errors.RendezVous.OperationPostpone = 0x000300D0
	Errors.RendezVous.OutOfRatingRange = 0x000300D1
	Errors.RendezVous.ConnectionDisconnected = 0x000300D2
	Errors.RendezVous.InvalidOperation = 0x000300D3
	Errors.RendezVous.NotParticipatedGathering = 0x000300D4
	Errors.RendezVous.MatchmakeSessionUserPasswordUnmatch = 0x000300D5
	Errors.RendezVous.MatchmakeSessionSystemPasswordUnmatch = 0x000300D6
	Errors.RendezVous.UserIsOffline = 0x000300D7
	Errors.RendezVous.AlreadyParticipatedGathering = 0x000300D8
	Errors.RendezVous.PermissionDenied = 0x000300D9
	Errors.RendezVous.NotFriend = 0x000300DA
	Errors.RendezVous.SessionClosed = 0x000300DB
	Errors.RendezVous.DatabaseTemporarilyUnavailable = 0x000300DC
	Errors.RendezVous.InvalidUniqueId = 0x000300DD
	Errors.RendezVous.MatchmakingWithdrawn = 0x000300DE
	Errors.RendezVous.LimitExceeded = 0x000300DF
	Errors.RendezVous.AccountTemporarilyDisabled = 0x000300E0
	Errors.RendezVous.PartiallyServiceClosed = 0x000300E1
	Errors.RendezVous.ConnectionDisconnectedForConcurrentLogin = 0x000300E2

	Errors.PythonCore.Exception = 0x00040001
	Errors.PythonCore.TypeError = 0x00040002
	Errors.PythonCore.IndexError = 0x00040003
	Errors.PythonCore.InvalidReference = 0x00040004
	Errors.PythonCore.CallFailure = 0x00040005
	Errors.PythonCore.MemoryError = 0x00040006
	Errors.PythonCore.KeyError = 0x00040007
	Errors.PythonCore.OperationError = 0x00040008
	Errors.PythonCore.ConversionError = 0x00040009
	Errors.PythonCore.ValidationError = 0x0004000A

	Errors.Transport.Unknown = 0x00050001
	Errors.Transport.ConnectionFailure = 0x00050002
	Errors.Transport.InvalidUrl = 0x00050003
	Errors.Transport.InvalidKey = 0x00050004
	Errors.Transport.InvalidURLType = 0x00050005
	Errors.Transport.DuplicateEndpoint = 0x00050006
	Errors.Transport.IOError = 0x00050007
	Errors.Transport.Timeout = 0x00050008
	Errors.Transport.ConnectionReset = 0x00050009
	Errors.Transport.IncorrectRemoteAuthentication = 0x0005000A
	Errors.Transport.ServerRequestError = 0x0005000B
	Errors.Transport.DecompressionFailure = 0x0005000C
	Errors.Transport.ReliableSendBufferFullFatal = 0x0005000D
	Errors.Transport.UPnPCannotInit = 0x0005000E
	Errors.Transport.UPnPCannotAddMapping = 0x0005000F
	Errors.Transport.NatPMPCannotInit = 0x00050010
	Errors.Transport.NatPMPCannotAddMapping = 0x00050011
	Errors.Transport.UnsupportedNAT = 0x00050013
	Errors.Transport.DnsError = 0x00050014
	Errors.Transport.ProxyError = 0x00050015
	Errors.Transport.DataRemaining = 0x00050016
	Errors.Transport.NoBuffer = 0x00050017
	Errors.Transport.NotFound = 0x00050018
	Errors.Transport.TemporaryServerError = 0x00050019
	Errors.Transport.PermanentServerError = 0x0005001A
	Errors.Transport.ServiceUnavailable = 0x0005001B
	Errors.Transport.ReliableSendBufferFull = 0x0005001C
	Errors.Transport.InvalidStation = 0x0005001D
	Errors.Transport.InvalidSubStreamID = 0x0005001E
	Errors.Transport.PacketBufferFull = 0x0005001F
	Errors.Transport.NatTraversalError = 0x00050020
	Errors.Transport.NatCheckError = 0x00050021

	Errors.DOCore.StationNotReached = 0x00060001
	Errors.DOCore.TargetStationDisconnect = 0x00060002
	Errors.DOCore.LocalStationLeaving = 0x00060003
	Errors.DOCore.ObjectNotFound = 0x00060004
	Errors.DOCore.InvalidRole = 0x00060005
	Errors.DOCore.CallTimeout = 0x00060006
	Errors.DOCore.RMCDispatchFailed = 0x00060007
	Errors.DOCore.MigrationInProgress = 0x00060008
	Errors.DOCore.NoAuthority = 0x00060009
	Errors.DOCore.NoTargetStationSpecified = 0x0006000A
	Errors.DOCore.JoinFailed = 0x0006000B
	Errors.DOCore.JoinDenied = 0x0006000C
	Errors.DOCore.ConnectivityTestFailed = 0x0006000D
	Errors.DOCore.Unknown = 0x0006000E
	Errors.DOCore.UnfreedReferences = 0x0006000F
	Errors.DOCore.JobTerminationFailed = 0x00060010
	Errors.DOCore.InvalidState = 0x00060011
	Errors.DOCore.FaultRecoveryFatal = 0x00060012
	Errors.DOCore.FaultRecoveryJobProcessFailed = 0x00060013
	Errors.DOCore.StationInconsitency = 0x00060014
	Errors.DOCore.AbnormalMasterState = 0x00060015
	Errors.DOCore.VersionMismatch = 0x00060016

	Errors.FPD.NotInitialized = 0x00650000
	Errors.FPD.AlreadyInitialized = 0x00650001
	Errors.FPD.NotConnected = 0x00650002
	Errors.FPD.Connected = 0x00650003
	Errors.FPD.InitializationFailure = 0x00650004
	Errors.FPD.OutOfMemory = 0x00650005
	Errors.FPD.RmcFailed = 0x00650006
	Errors.FPD.InvalidArgument = 0x00650007
	Errors.FPD.InvalidLocalAccountID = 0x00650008
	Errors.FPD.InvalidPrincipalID = 0x00650009
	Errors.FPD.InvalidLocalFriendCode = 0x0065000A
	Errors.FPD.LocalAccountNotExists = 0x0065000B
	Errors.FPD.LocalAccountNotLoaded = 0x0065000C
	Errors.FPD.LocalAccountAlreadyLoaded = 0x0065000D
	Errors.FPD.FriendAlreadyExists = 0x0065000E
	Errors.FPD.FriendNotExists = 0x0065000F
	Errors.FPD.FriendNumMax = 0x00650010
	Errors.FPD.NotFriend = 0x00650011
	Errors.FPD.FileIO = 0x00650012
	Errors.FPD.P2PInternetProhibited = 0x00650013
	Errors.FPD.Unknown = 0x00650014
	Errors.FPD.InvalidState = 0x00650015
	Errors.FPD.AddFriendProhibited = 0x00650017
	Errors.FPD.InvalidAccount = 0x00650019
	Errors.FPD.BlacklistedByMe = 0x0065001A
	Errors.FPD.FriendAlreadyAdded = 0x0065001C
	Errors.FPD.MyFriendListLimitExceed = 0x0065001D
	Errors.FPD.RequestLimitExceed = 0x0065001E
	Errors.FPD.InvalidMessageID = 0x0065001F
	Errors.FPD.MessageIsNotMine = 0x00650020
	Errors.FPD.MessageIsNotForMe = 0x00650021
	Errors.FPD.FriendRequestBlocked = 0x00650022
	Errors.FPD.NotInMyFriendList = 0x00650023
	Errors.FPD.FriendListedByMe = 0x00650024
	Errors.FPD.NotInMyBlacklist = 0x00650025
	Errors.FPD.IncompatibleAccount = 0x00650026
	Errors.FPD.BlockSettingChangeNotAllowed = 0x00650027
	Errors.FPD.SizeLimitExceeded = 0x00650028
	Errors.FPD.OperationNotAllowed = 0x00650029
	Errors.FPD.NotNetworkAccount = 0x0065002A
	Errors.FPD.NotificationNotFound = 0x0065002B
	Errors.FPD.PreferenceNotInitialized = 0x0065002C
	Errors.FPD.FriendRequestNotAllowed = 0x0065002D

	Errors.Ranking.NotInitialized = 0x00670001
	Errors.Ranking.InvalidArgument = 0x00670002
	Errors.Ranking.RegistrationError = 0x00670003
	Errors.Ranking.NotFound = 0x00670005
	Errors.Ranking.InvalidScore = 0x00670006
	Errors.Ranking.InvalidDataSize = 0x00670007
	Errors.Ranking.PermissionDenied = 0x00670009
	Errors.Ranking.Unknown = 0x0067000A
	Errors.Ranking.NotImplemented = 0x0067000B

	Errors.Authentication.NASAuthenticateError = 0x00680001
	Errors.Authentication.TokenParseError = 0x00680002
	Errors.Authentication.HttpConnectionError = 0x00680003
	Errors.Authentication.HttpDNSError = 0x00680004
	Errors.Authentication.HttpGetProxySetting = 0x00680005
	Errors.Authentication.TokenExpired = 0x00680006
	Errors.Authentication.ValidationFailed = 0x00680007
	Errors.Authentication.InvalidParam = 0x00680008
	Errors.Authentication.PrincipalIdUnmatched = 0x00680009
	Errors.Authentication.MoveCountUnmatch = 0x0068000A
	Errors.Authentication.UnderMaintenance = 0x0068000B
	Errors.Authentication.UnsupportedVersion = 0x0068000C
	Errors.Authentication.ServerVersionIsOld = 0x0068000D
	Errors.Authentication.Unknown = 0x0068000E
	Errors.Authentication.ClientVersionIsOld = 0x0068000F
	Errors.Authentication.AccountLibraryError = 0x00680010
	Errors.Authentication.ServiceNoLongerAvailable = 0x00680011
	Errors.Authentication.UnknownApplication = 0x00680012
	Errors.Authentication.ApplicationVersionIsOld = 0x00680013
	Errors.Authentication.OutOfService = 0x00680014
	Errors.Authentication.NetworkServiceLicenseRequired = 0x00680015
	Errors.Authentication.NetworkServiceLicenseSystemError = 0x00680016
	Errors.Authentication.NetworkServiceLicenseError3 = 0x00680017
	Errors.Authentication.NetworkServiceLicenseError4 = 0x00680018

	Errors.DataStore.Unknown = 0x00690001
	Errors.DataStore.InvalidArgument = 0x00690002
	Errors.DataStore.PermissionDenied = 0x00690003
	Errors.DataStore.NotFound = 0x00690004
	Errors.DataStore.AlreadyLocked = 0x00690005
	Errors.DataStore.UnderReviewing = 0x00690006
	Errors.DataStore.Expired = 0x00690007
	Errors.DataStore.InvalidCheckToken = 0x00690008
	Errors.DataStore.SystemFileError = 0x00690009
	Errors.DataStore.OverCapacity = 0x0069000A
	Errors.DataStore.OperationNotAllowed = 0x0069000B
	Errors.DataStore.InvalidPassword = 0x0069000C
	Errors.DataStore.ValueNotEqual = 0x0069000D

	Errors.ServiceItem.Unknown = 0x006C0001
	Errors.ServiceItem.InvalidArgument = 0x006C0002
	Errors.ServiceItem.EShopUnknownHttpError = 0x006C0003
	Errors.ServiceItem.EShopResponseParseError = 0x006C0004
	Errors.ServiceItem.NotOwned = 0x006C0005
	Errors.ServiceItem.InvalidLimitationType = 0x006C0006
	Errors.ServiceItem.ConsumptionRightShortage = 0x006C0007

	Errors.MatchmakeReferee.Unknown = 0x006F0001
	Errors.MatchmakeReferee.InvalidArgument = 0x006F0002
	Errors.MatchmakeReferee.AlreadyExists = 0x006F0003
	Errors.MatchmakeReferee.NotParticipatedGathering = 0x006F0004
	Errors.MatchmakeReferee.NotParticipatedRound = 0x006F0005
	Errors.MatchmakeReferee.StatsNotFound = 0x006F0006
	Errors.MatchmakeReferee.RoundNotFound = 0x006F0007
	Errors.MatchmakeReferee.RoundArbitrated = 0x006F0008
	Errors.MatchmakeReferee.RoundNotArbitrated = 0x006F0009

	Errors.Subscriber.Unknown = 0x00700001
	Errors.Subscriber.InvalidArgument = 0x00700002
	Errors.Subscriber.OverLimit = 0x00700003
	Errors.Subscriber.PermissionDenied = 0x00700004

	Errors.Ranking2.Unknown = 0x00710001
	Errors.Ranking2.InvalidArgument = 0x00710002
	Errors.Ranking2.InvalidScore = 0x00710003

	Errors.SmartDeviceVoiceChat.Unknown = 0x00720001
	Errors.SmartDeviceVoiceChat.InvalidArgument = 0x00720002
	Errors.SmartDeviceVoiceChat.InvalidResponse = 0x00720003
	Errors.SmartDeviceVoiceChat.InvalidAccessToken = 0x00720004
	Errors.SmartDeviceVoiceChat.Unauthorized = 0x00720005
	Errors.SmartDeviceVoiceChat.AccessError = 0x00720006
	Errors.SmartDeviceVoiceChat.UserNotFound = 0x00720007
	Errors.SmartDeviceVoiceChat.RoomNotFound = 0x00720008
	Errors.SmartDeviceVoiceChat.RoomNotActivated = 0x00720009
	Errors.SmartDeviceVoiceChat.ApplicationNotSupported = 0x0072000A
	Errors.SmartDeviceVoiceChat.InternalServerError = 0x0072000B
	Errors.SmartDeviceVoiceChat.ServiceUnavailable = 0x0072000C
	Errors.SmartDeviceVoiceChat.UnexpectedError = 0x0072000D
	Errors.SmartDeviceVoiceChat.UnderMaintenance = 0x0072000E
	Errors.SmartDeviceVoiceChat.ServiceNoLongerAvailable = 0x0072000F
	Errors.SmartDeviceVoiceChat.AccountTemporarilyDisabled = 0x00720010
	Errors.SmartDeviceVoiceChat.PermissionDenied = 0x00720011
	Errors.SmartDeviceVoiceChat.NetworkServiceLicenseRequired = 0x00720012
	Errors.SmartDeviceVoiceChat.AccountLibraryError = 0x00720013
	Errors.SmartDeviceVoiceChat.GameModeNotFound = 0x00720014

	Errors.Screening.Unknown = 0x00730001
	Errors.Screening.InvalidArgument = 0x00730002
	Errors.Screening.NotFound = 0x00730003

	Errors.Custom.Unknown = 0x00740001

	Errors.Ess.Unknown = 0x00750001
	Errors.Ess.GameSessionError = 0x00750002
	Errors.Ess.GameSessionMaintenance = 0x00750003

	valueOfErrors := reflect.ValueOf(Errors)
	typeOfErrors := valueOfErrors.Type()

	for i := 0; i < valueOfErrors.NumField(); i++ {
		category := typeOfErrors.Field(i).Name

		valueOfCategory := reflect.ValueOf(valueOfErrors.Field(i).Interface())
		typeOfCategory := valueOfCategory.Type()

		for j := 0; j < valueOfCategory.NumField(); j++ {
			errorName := typeOfCategory.Field(j).Name
			errorCode := valueOfCategory.Field(j).Interface().(uint32)

			ErrorNames[errorCode] = category + "::" + errorName
		}
	}
}

func ErrorNameFromCode(errorCode uint32) string {
	name := ErrorNames[errorCode]

	if name == "" {
		return "Invalid Error Code: " + strconv.Itoa(int(errorCode))
	}

	return name
}
